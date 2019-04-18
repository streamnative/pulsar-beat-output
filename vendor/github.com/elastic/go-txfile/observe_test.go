// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package txfile

import (
	"testing"
	"time"
)

type statEntry struct {
	kind     statKind
	readonly bool
	file     FileStats
	tx       TxStats
}

type statKind uint8

const (
	statOnOpen statKind = iota
	statOnTxBegin
	statOnTxClose
)

type testObserveLast statEntry

func TestObserveStats(t *testing.T) {
	assert := newAssertions(t)

	setupFileWith := func(assert *assertions, stat *statEntry) (*testFile, func()) {
		return setupTestFile(assert, Options{
			Observer: newTestObserver(stat),
			MaxSize:  10 * 1 << 20, // 10 MB
		})
	}

	initStat := FileStats{
		Version:  1,
		Size:     4096 * 2,
		MaxSize:  10 * 1 << 20,
		PageSize: 4096,
	}

	assert.Run("new file", func(assert *assertions) {
		var stat statEntry
		_, teardown := setupFileWith(assert, &stat)
		defer teardown()

		checkOpenStat(assert, initStat, stat)
	})

	assert.Run("allocate and free pages", func(assert *assertions) {
		var stat statEntry
		f, teardown := setupFileWith(assert, &stat)
		defer teardown()

		activeFStat := initStat

		// 1. transaction allocating 4 pages
		var inuse []PageID
		{
			tx := f.Begin()
			assert.Equal(statOnTxBegin, stat.kind)
			defer tx.Close()

			// allocate + commit
			pages, err := tx.AllocN(4)
			if !assert.NoError(err) {
				return
			}
			for _, page := range pages {
				inuse = append(inuse, page.ID())
				page.SetBytes([]byte{1, 2, 3, 4})
			}

			time.Sleep(500 * time.Millisecond)
			assert.NoError(tx.Commit())

			// validate last tx stats
			activeFStat = updFileStats(activeFStat, FileStats{
				Size:          6 * 4096,
				DataAllocated: 4,
			})
			checkFileStat(assert, activeFStat, stat, "invalid file stats after alloc only tx")
			checkTxStat(assert, TxStats{
				Readonly:  false,
				Commit:    true,
				Total:     4,
				Accessed:  0,
				Allocated: 4,
				Written:   4,
			}, stat, "invalid tx stats after alloc only tx")

			f.Reopen()
			checkOpenStat(assert, activeFStat, stat, "invalid file stats after reopen")
			if assert.Failed() {
				return
			}
		}

		// 2. transaction releasing the first 2 allocated pages -> force meta data
		{
			tx := f.Begin()
			assert.Equal(statOnTxBegin, stat.kind)
			defer tx.Close()

			// free first 2 allocated pages + commit
			for _, id := range inuse[:2] {
				p, err := tx.Page(id)
				if !assert.NoError(err) {
					return
				}

				assert.NoError(p.Free())
			}
			inuse = inuse[2:]

			time.Sleep(500 * time.Millisecond)
			assert.NoError(tx.Commit())

			// validate
			activeFStat = updFileStats(activeFStat, FileStats{
				DataAllocated: 2,
				MetaAllocated: 1,
				MetaArea:      2,
				Size:          activeFStat.Size + 2*4096,
			})
			checkFileStat(assert, activeFStat, stat, "invalid file stat after free tx")
			checkTxStat(assert, TxStats{
				Readonly:  false,
				Commit:    true,
				Total:     1,
				Accessed:  0,
				Allocated: 0,
				Freed:     2,
				Written:   1, // note: write of meta data page
			}, stat, "invalid tx stats after free only tx")

			f.Reopen()
			checkOpenStat(assert, activeFStat, stat, "invalid file stats after reopen with free list in middle of data area")
			if assert.Failed() {
				return
			}
		}
	})

	assert.Run("readonly transaction", func(assert *assertions) {
		var stat statEntry
		f, teardown := setupFileWith(assert, &stat)
		defer teardown()

		// prepare
		var inuse []PageID
		f.withTx(true, func(tx *Tx) {
			pages, err := tx.AllocN(4)
			if !assert.NoError(err) {
				return
			}

			for _, page := range pages {
				inuse = append(inuse, page.ID())
				page.SetBytes([]byte{1, 2, 3, 4})
			}

			assert.NoError(tx.Commit())
		})
		if assert.Failed() {
			return
		}

		// read a few pages
		f.withTx(false, func(tx *Tx) {
			for _, id := range inuse {
				p, err := tx.Page(id)
				if assert.NoError(err) {
					_, err := p.Bytes()
					assert.NoError(err)
				}
			}

			time.Sleep(500 * time.Millisecond)
		})
		if assert.Failed() {
			return
		}

		// validate
		checkTxStat(assert, TxStats{
			Readonly: true,
			Commit:   false,
			Total:    4,
			Accessed: 4,
		}, stat, "invalid tx stats after alloc only tx")

	})

	assert.Run("rollback does not change file stats", func(assert *assertions) {
		var stat statEntry
		f, teardown := setupFileWith(assert, &stat)
		defer teardown()

		tx := f.Begin()
		assert.Equal(statOnTxBegin, stat.kind)
		defer tx.Close()

		// allocate + write
		pages, err := tx.AllocN(4)
		if !assert.NoError(err) {
			return
		}
		for _, page := range pages {
			page.SetBytes([]byte{1, 2, 3, 4})
		}
		assert.NoError(tx.Flush())

		time.Sleep(500 * time.Millisecond)
		assert.NoError(tx.Rollback()) // rollback after write

		// validate
		checkFileStat(assert, initStat, stat, "invalid file stats after rollback tx")
		checkTxStat(assert, TxStats{
			Readonly:  false,
			Commit:    false,
			Total:     4,
			Accessed:  0,
			Allocated: 4,
			Written:   4,
		}, stat, "invalid tx stats after rollback with writes")
	})
}

func newTestObserver(stat *statEntry) *testObserveLast {
	return (*testObserveLast)(stat)
}

func (t *testObserveLast) OnOpen(stats FileStats) {
	*t = testObserveLast(statEntry{
		kind: statOnOpen,
		file: stats,
	})
}

func (t *testObserveLast) OnTxBegin(readonly bool) {
	*t = testObserveLast(statEntry{
		kind:     statOnTxBegin,
		readonly: readonly,
	})
}

func (t *testObserveLast) OnTxClose(file FileStats, tx TxStats) {
	*t = testObserveLast(statEntry{
		kind: statOnTxClose,
		file: file,
		tx:   tx,
	})
}

func checkOpenStat(assert *assertions, expected FileStats, actual statEntry, msg ...interface{}) {
	assert.Equal(statOnOpen, actual.kind, msg...)
	checkFileStat(assert, expected, actual, msg...)
}

func checkFileStat(assert *assertions, expected FileStats, actual statEntry, msg ...interface{}) {
	assert.Equal(expected, actual.file, msg...)
}

func checkTxStat(assert *assertions, expected TxStats, actual statEntry, msg ...interface{}) {
	txActual := actual.tx

	assert.Equal(statOnTxClose, actual.kind, "invalid stat type")

	assert.True(txActual.Duration > 0, "duration should not be 0")
	txActual.Duration = 0

	assert.Equal(expected, txActual, msg...)
}

func updFileStats(orig, upd FileStats) FileStats {
	sel64 := func(orig, upd uint64) uint64 {
		if upd != 0 {
			return upd
		}
		return orig
	}

	sel := func(orig, upd uint) uint { return uint(sel64(uint64(orig), uint64(upd))) }
	sel32 := func(orig, upd uint32) uint32 { return uint32(sel64(uint64(orig), uint64(upd))) }

	return FileStats{
		Version:       sel32(orig.Version, upd.Version),
		Size:          sel64(orig.Size, upd.Size),
		MaxSize:       sel64(orig.MaxSize, upd.MaxSize),
		PageSize:      sel32(orig.PageSize, upd.PageSize),
		MetaArea:      sel(orig.MetaArea, upd.MetaArea),
		DataAllocated: sel(orig.DataAllocated, upd.DataAllocated),
		MetaAllocated: sel(orig.MetaAllocated, upd.MetaAllocated),
	}
}

func updTxStats(orig, upd TxStats) TxStats {
	sel := func(orig, upd uint) uint {
		if upd != 0 {
			return upd
		}
		return orig
	}

	return TxStats{
		Readonly:  upd.Readonly,
		Commit:    upd.Commit,
		Duration:  time.Duration(sel(uint(orig.Duration), uint(upd.Duration))),
		Total:     sel(orig.Total, upd.Total),
		Accessed:  sel(orig.Accessed, upd.Accessed),
		Allocated: sel(orig.Allocated, upd.Allocated),
		Freed:     sel(orig.Freed, upd.Freed),
		Written:   sel(orig.Written, upd.Written),
		Updated:   sel(orig.Updated, upd.Updated),
	}
}
