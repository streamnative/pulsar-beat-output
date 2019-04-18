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
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/elastic/go-txfile/internal/cleanup"
	"github.com/elastic/go-txfile/internal/vfs/osfs/osfstest"
)

type testFile struct {
	*File
	path   string
	assert *assertions
	opts   Options

	// validation overwrites
	disableEndMarkerCheck bool // disable checking file end markers against max size
}

// global bools with randomized order
var bools [2]bool

func init() {
	now := time.Now().Unix()
	rng := rand.New(rand.NewSource(now))

	bools[rng.Intn(2)] = true
}

func TestTxFile(t *testing.T) {
	assert := newAssertions(t)

	sampleContents := []string{
		"Hello World",
		"The quick brown fox jumps over the lazy dog",
		"Lorem ipsum dolor sit amet, consectetur adipisici elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua.",
	}

	assert.Run("mmap sizing", func(assert *assertions) {
		const (
			_ = 1 << (10 * iota)
			KB
			MB
			GB
		)

		testcases := []struct {
			min, max, expected uint64
		}{
			{64, 0, 64 * KB},
			{4 * KB, 0, 64 * KB},
			{100 * KB, 0, 128 * KB},
			{5 * MB, 0, 8 * MB},
			{300 * MB, 0, 512 * MB},
			{1200 * MB, 0, 2 * GB},
			{2100 * MB, 0, 3 * GB},
		}

		for _, test := range testcases {
			min, max, expected := test.min, test.max, test.expected
			title := fmt.Sprintf("min=%v,max=%v,expected=%v", min, max, expected)

			assert.Run(title, func(assert *assertions) {
				fmt.Printf("maxUint: %v, expected: %v\n", maxUint, expected)
				if expected > uint64(maxMmapSize) {
					assert.Skip("unsatisfyable tests on 32bit system")
				}

				sz, err := computeMmapSize(uint(min), uint(max), 4096)
				assert.NoError(err)
				assert.Equal(int(expected), int(sz))
			})
		}
	})

	assert.Run("open/close", func(assert *assertions) {
		path, teardown := setupPath(assert, "")
		defer teardown()

		assert.Log("create and close file")
		f, err := Open(path, os.ModePerm, Options{
			MaxSize:  10 * 1 << 20, // 10MB
			PageSize: 4096,
		})
		assert.FatalOnError(err)
		assert.NoError(f.Close())

		// check if we can re-open the file:

		assert.Log("open and close existing file")
		f, err = Open(path, os.ModePerm, Options{})
		assert.FatalOnError(err)
		assert.NoError(f.Close())
	})
	if assert.Failed() {
		// if we find a file can not even be created correctly, no need to run more
		// tests, that rely on file creation during test setup
		return
	}

	assert.Run("start and close readonly transaction without reads", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		tx := f.BeginReadonly()
		assert.NotNil(tx)
		assert.True(tx.Readonly())
		assert.False(tx.Writable())
		assert.True(tx.Active())
		assert.NoError(tx.Close())
		assert.False(tx.Active())
	})

	assert.Run("start and close read-write transaction without reads/writes", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		tx := f.Begin()
		assert.NotNil(tx)
		assert.False(tx.Readonly())
		assert.True(tx.Writable())
		assert.True(tx.Active())
		assert.NoError(tx.Close())
		assert.False(tx.Active())
	})

	assert.Run("readonly transaction can not allocate pages", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		tx := f.BeginReadonly()
		defer tx.Close()

		page, err := tx.Alloc()
		assert.Nil(page)
		assert.Error(err)
	})

	assert.Run("pre-allocate meta area", func(assert *assertions) {
		assert.Run("available pages do not exceeds total number of pages", func(assert *assertions) {
			const (
				totalPages = 100
				metaArea   = 64
				dataArea   = totalPages - metaArea - 2
			)
			f, teardown := setupTestFile(assert, Options{
				PageSize:     4096,
				MaxSize:      totalPages * 4096,
				InitMetaArea: metaArea,
			})
			defer teardown()

			f.withTx(true, func(tx *Tx) {
				// Allocate all available data pages.
				_, err := tx.AllocN(dataArea)
				assert.FatalOnError(err)

				// Allocating more pages must fail, due to most space being used by the
				// meta area.
				_, err = tx.Alloc()
				assert.Error(err)
			})
		})

		assert.Run("check meta area is used", func(assert *assertions) {
			/* Allocate all data pages and free some, forcing the file to create a
			 * new freelist, that must be written to the existing meta area. Trying to
			 * allocate more pages from the data area into the meta area would fail
			 * with OOM.
			 */

			const (
				totalPages = 100
				metaPages  = 2
				metaArea   = 64
				dataArea   = totalPages - metaArea - metaPages
			)
			f, teardown := setupTestFile(assert, Options{
				PageSize:     4096,
				MaxSize:      totalPages * 4096,
				InitMetaArea: metaArea,
			})
			defer teardown()

			// allocate all free space
			f.withTx(true, func(tx *Tx) {
				_, err := tx.AllocN(dataArea)
				assert.FatalOnError(err)
				assert.FatalOnError(tx.Commit())
			})

			// return some pages, forcing a freelist update
			f.withTx(true, func(tx *Tx) {
				// find ID in middle of allocated data area
				page, err := tx.Page(metaArea + metaPages + 5)
				assert.FatalOnError(err)
				page.Free()

				assert.FatalOnError(tx.Commit())
			})

		})
	})

	assert.Run("write transaction with modifications on new file with rollback", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		tx := f.Begin()
		defer tx.Close()

		ids := pageSet{}

		page, err := tx.Alloc()
		assert.FatalOnError(err)
		if assert.NotNil(page) {
			ids.Add(page.ID())
			assert.NotEqual(PageID(0), page.ID())
		}

		pages, err := tx.AllocN(5)
		assert.FatalOnError(err)
		assert.Len(pages, 5)
		for _, page := range pages {
			if assert.NotNil(page) {
				assert.NotEqual(PageID(0), page.ID())
				ids.Add(page.ID())
			}
		}

		// check we didn't get the same ID twice:
		assert.Len(ids, 6)

		f.Rollback(tx)
	})

	assert.Run("comitting write transaction without modifications", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		tx := f.Begin()
		defer tx.Close()
		f.Commit(tx)
	})

	assert.Run("committing write transaction on new file with page writes", func(assert *assertions) {
		contents := sampleContents
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		var ids idList
		func() {
			tx := f.Begin()
			defer tx.Close()
			ids = f.txAppend(tx, contents)
			tx.SetRoot(ids[0])
			f.Commit(tx, "failed to commit the initial transaction")
		}()

		traceln("transaction page ids: ", ids)

		f.Reopen()
		tx := f.BeginReadonly()
		defer tx.Close()
		assert.Equal(ids[0], tx.Root())
		assert.Equal(contents, f.readIDs(ids))
	})

	assert.Run("read contents after reopen and page contents has been overwritten", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		writeSample := func(msg string, id PageID) (to PageID) {
			f.withTx(true, func(tx *Tx) {
				to = id
				if to == 0 {
					page, err := tx.Alloc()
					assert.FatalOnError(err)
					to = page.ID()
				}

				f.txWriteAt(tx, to, msg)
				f.Commit(tx)
			})
			return
		}

		msgs := sampleContents[:2]
		id := writeSample(msgs[0], 0)
		writeSample(msgs[1], id)

		// reopen and check new contents is really available
		f.Reopen()
		assert.Equal(msgs[1], f.read(id))
	})

	assert.Run("check allocates smallest page possible", func(assert *assertions) {
		bools := [2]bool{false, true}

		for _, reopen := range bools {
			for _, withFragmentation := range bools {
				reopen := reopen
				withFragmentation := withFragmentation
				title := fmt.Sprintf("reopen=%v, fragmentation=%v", reopen, withFragmentation)
				assert.Run(title, func(assert *assertions) {

					f, teardown := setupTestFile(assert, Options{})
					defer teardown()

					// allocate 2 pages
					var id PageID
					f.withTx(true, func(tx *Tx) {
						page, err := tx.Alloc()
						assert.FatalOnError(err)

						// allocate dummy page, so to ensure some fragmentation on free
						if withFragmentation {
							_, err = tx.Alloc()
							assert.FatalOnError(err)
						}

						id = page.ID()
						f.Commit(tx)
					})

					if reopen {
						f.Reopen()
					}

					// free first allocated page
					f.withTx(true, func(tx *Tx) {
						page, err := tx.Page(id)
						assert.FatalOnError(err)
						assert.FatalOnError(page.Free())
						f.Commit(tx)
					})

					if reopen {
						f.Reopen()
					}

					// expect just freed page can be allocated again
					var newID PageID
					f.withTx(true, func(tx *Tx) {
						page, err := tx.Alloc()
						assert.FatalOnError(err)

						newID = page.ID()
						f.Commit(tx)
					})

					// verify
					assert.Equal(id, newID)
				})
			}
		}

	})

	assert.Run("file open old transaction if verification fails", func(assert *assertions) {
		msgs := sampleContents[:2]
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		// start first transaction with first message

		writeSample := func(msg string, id PageID) PageID {
			tx := f.Begin()
			defer tx.Close()

			if id == 0 {
				page, err := tx.Alloc()
				assert.FatalOnError(err)
				id = page.ID()
			}

			f.txWriteAt(tx, id, msg)
			tx.SetRoot(id)
			f.Commit(tx)
			return id
		}

		id := writeSample(msgs[0], 0)
		writeSample(msgs[1], id) // overwrite contents

		// get active meta page id
		metaID := PageID(f.File.metaActive)
		metaOff := int64(metaID) * int64(f.allocator.pageSize)

		// close file and write invalid contents into last tx meta page
		f.Close()

		func() {
			tmp, err := os.OpenFile(f.path, os.O_RDWR, 0777)
			assert.FatalOnError(err)
			defer tmp.Close()

			_, err = tmp.WriteAt([]byte{1, 2, 3, 4, 5}, metaOff)
			assert.FatalOnError(err)
		}()

		// Open db file again and check recent transaction contents is still
		// available.
		f.Open()
		assert.Equal(msgs[0], f.read(id))
	})

	assert.Run("concurrent read transaction can not access not comitted contents", func(assert *assertions) {
		orig, modified := sampleContents[0], sampleContents[1:]
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		// commit original contents first
		id := f.append([]string{orig})[0]

		current := orig
		for _, newmsg := range modified {
			func() {
				// create write transaction, used to update the contents
				writer := f.Begin()
				defer writer.Close()
				f.txWriteAt(writer, id, newmsg)

				// check concurrent read will not have access to new contents
				reader := f.BeginReadonly()
				defer reader.Close()
				assert.Equal(current, f.txRead(reader, id))
				f.Commit(reader) // close reader transaction

				// check new reader will read modified contents
				reader = f.BeginReadonly()
				defer reader.Close()
				assert.Equal(current, f.txRead(reader, id), "read unexpected message")
				f.Commit(reader)

				// finish transaction
				f.Commit(writer)
				current = newmsg
			}()
		}
	})

	assert.Run("execute wal checkpoint", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		// first tx with original contents
		ids := f.append(sampleContents)

		// overwrite original contents
		overwrites := append(sampleContents[1:], sampleContents[0])
		f.writeAllAt(ids, overwrites)
		assert.Len(f.wal.mapping, 3, "expected wal mapping entries")

		// run wal checkpoint
		f.withTx(true, func(tx *Tx) {
			assert.FatalOnError(tx.CheckpointWAL())
			f.Commit(tx)
		})
		assert.Len(f.wal.mapping, 0, "mapping must be empty after checkpointing")
	})

	assert.Run("force remap", func(assert *assertions) {
		N, i := 64/3+1, 0
		contents := make([]string, 0, N)
		for len(contents) < N {
			contents = append(contents, sampleContents[i])
			if i++; i == len(sampleContents) {
				i = 0
			}
		}

		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		ids := f.append(contents) // write enough contents to enforce an munmap/mmap
		assert.Equal(len(contents), len(ids))

		// check we can still read all contents
		assert.Equal(contents, f.readIDs(ids))
	})

	assert.Run("inplace update", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		id := f.append(sampleContents[:1])[0]

		func() {
			tx := f.Begin()
			defer tx.Close()

			page, err := tx.Page(id)
			assert.FatalOnError(err)
			assert.FatalOnError(page.Load())
			buf, err := page.Bytes()
			assert.FatalOnError(err)

			// validate contents buffer
			assert.Len(buf, 4096)
			L := int(castU32(buf).Get())
			assert.Equal(len(sampleContents[0]), L)
			assert.Equal(sampleContents[0], string(buf[4:4+L]))

			// overwrite buffer
			castU32(buf).Set(uint32(len(sampleContents[1])))
			copy(buf[4:], sampleContents[1])
			assert.FatalOnError(page.MarkDirty())

			// commit
			f.Commit(tx)
		}()

		// read/validate new contents
		assert.Equal(sampleContents[1], f.read(id))
	})

	assert.Run("multiple inplace updates", func(assert *assertions) {
		f, teardown := setupTestFile(assert, Options{})
		defer teardown()

		id := f.append(sampleContents[:1])[0]
		assert.Equal(sampleContents[0], f.read(id))

		f.writeAt(id, sampleContents[1])
		assert.Equal(sampleContents[1], f.read(id))

		f.writeAt(id, sampleContents[2])
		assert.Equal(sampleContents[2], f.read(id))
	})
}

func TestResizeFile(t *testing.T) {
	assert := newAssertions(t)

	type (
		st struct {
			allocated pageSet
		}

		op func(f *testFile, st *st)

		scenario struct {
			opts Options
			ops  []op
		}
	)

	noop := func(_ *testFile, _ *st) {}

	opt := func(b bool, o op) op {
		if b {
			return o
		}
		return noop
	}

	allocateAll := func(f *testFile, st *st) {
		var ids pageSet
		f.assert.Log("Allocate all available pages")
		f.withTx(true, func(tx *Tx) {
			for {
				page, err := tx.Alloc()
				if err != nil {
					break
				}
				ids.Add(page.ID())
			}

			assert.NoError(tx.Commit())
		})

		st.allocated.AddSet(ids)
	}

	allocate := func(n int, success bool) op {
		return func(f *testFile, st *st) {
			if success {
				f.assert.Logf("Must Allocate %v pages", n)
			} else {
				f.assert.Logf("Try to allocate %v pages, expecting a failure", n)
			}

			f.withTx(true, func(tx *Tx) {
				pages, err := tx.AllocN(n)
				if success && err != nil {
					f.assert.FatalOnError(err)
				}
				if !success && err == nil {
					f.assert.Fatal("Expected allocation to fail")
				}

				var ids pageSet
				for _, page := range pages {
					ids.Add(page.ID())
				}

				f.assert.FatalOnError(tx.Commit())

				st.allocated.AddSet(ids)
			})
		}
	}

	releasePages := func(ids idList) op {
		return func(f *testFile, st *st) {
			f.assert.Logf("release %v pages", len(ids))
			for _, id := range ids {
				if !st.allocated.Has(id) {
					panic(fmt.Errorf("incorrect test, trying to free unallocated page %v", id))
				}
			}

			f.withTx(true, func(tx *Tx) {
				for _, id := range ids {
					page, err := tx.Page(id)
					f.assert.FatalOnError(err)

					page.Free()
				}

				f.assert.FatalOnError(tx.Commit())
			})

			for _, id := range ids {
				st.allocated.Remove(id)
			}
		}
	}

	/* reopenFile := func(f *testFile, _ *st) {
		  f.assert.Logf("reopen file")
		  f.Reopen()
	  } */
	closeFile := func(f *testFile, _ *st) {
		f.assert.Logf("close test file")
		f.Close()
	}
	openWith := func(opts Options) op {
		return func(f *testFile, _ *st) {
			f.assert.Logf("open test file with %#v", opts)
			f.OpenWith(opts)
		}
	}
	reopenWith := func(opts Options) op {
		open := openWith(opts)
		return func(f *testFile, st *st) {
			closeFile(f, st)
			open(f, st)
		}
	}
	optReopenWith := func(reopen bool, opts Options) op {
		return opt(reopen, reopenWith(opts))
	}

	checkSize := func(expected int64) op {
		return func(f *testFile, st *st) {
			f.assert.Equal(expected, f.size)
		}
	}

	runScenario := func(assert *assertions, scenario scenario) {
		assert.Logf("create file with: %#v", scenario.opts)

		f, teardown := setupTestFile(assert, scenario.opts)
		defer teardown()

		// disable end marker checks on validation, as resizing files can
		// temporary violate the endMarker <= maxSize invariant.
		f.disableEndMarkerCheck = true

		var st st
		for _, op := range scenario.ops {
			op(f, &st)
		}
	}

	runScenarios := func(assert *assertions, scenarios map[string]scenario) {
		for name, scenario := range scenarios {
			assert.Run(name, func(assert *assertions) {
				runScenario(assert, scenario)
			})
		}
	}

	assert.Run("grow file on re-open", func(assert *assertions) {
		pageSize := uint(4096)
		initSize := 50 * pageSize

		cases := map[string]uint{
			"new limit":     2 * initSize,
			"same max size": initSize,
			"unbound limit": 0,
		}

		for _, reopen := range bools {
			reopen := reopen
			assert.Run(fmt.Sprintf("reopen:%v", reopen), func(assert *assertions) {
				for _, prealloc := range bools {
					prealloc := prealloc
					assert.Run(fmt.Sprintf("prealloc:%v", prealloc), func(assert *assertions) {
						scenarios := map[string]scenario{}
						for name, sz := range cases {
							canAllocMore := sz == 0 || sz > initSize
							expectedSz := sz
							if expectedSz == 0 {
								expectedSz = initSize
							}

							scenarios[name] = scenario{
								opts: Options{
									MaxSize:  uint64(initSize),
									PageSize: uint32(pageSize),
									Prealloc: prealloc,
								},
								ops: []op{
									allocateAll,

									// resize
									closeFile,
									openWith(Options{
										MaxSize:  uint64(sz),
										PageSize: uint32(pageSize),
										Flags:    FlagUpdMaxSize,
										Prealloc: prealloc,
									}),
									opt(prealloc, checkSize(int64(expectedSz))),

									// optional: close and open file without extra Options
									optReopenWith(reopen, Options{}),
									opt(reopen && prealloc, checkSize(int64(expectedSz))),

									// validate by allocating one more page
									allocate(1, canAllocMore),

									opt(prealloc, checkSize(int64(expectedSz))),
								},
							}
						}
						runScenarios(assert, scenarios)

					})
				}
			})
		}
	})

	assert.Run("shrink file", func(assert *assertions) {
		const initPages = 100
		const pageSize = 4096

		initOpts := Options{
			MaxSize:  pageSize * initPages,
			PageSize: pageSize,
		}

		for _, reopen := range bools {
			const numHeaderPages = 2
			reopen := reopen

			assert.Run(fmt.Sprintf("reopen:%v", reopen), func(assert *assertions) {
				runScenarios(assert, map[string]scenario{
					"fail with too small size": scenario{
						opts: initOpts,
						ops: []op{
							closeFile,
							func(f *testFile, _ *st) {
								opts := Options{
									MaxSize: pageSize * 4,
									Flags:   FlagUpdMaxSize,
								}
								f.assert.Logf("open test file with invalid %#v", opts)

								tmp, err := Open(f.path, os.ModePerm, opts)
								if err == nil {
									tmp.Close()
									f.assert.Fail("expected resize to fail")
								}
							},
						},
					},

					"empty file": scenario{
						opts: initOpts,
						ops: []op{
							// resize file
							closeFile,
							openWith(Options{
								MaxSize:  pageSize * initPages / 2,
								Flags:    FlagUpdMaxSize,
								Prealloc: true,
							}),

							// optional: close and open file without extra Options
							optReopenWith(reopen, Options{}),

							// try to allocate all available pages
							allocate(initPages/2-numHeaderPages, true),

							// check we can not allocate any new pages
							allocate(1, false),
						},
					},

					"pages already allocated": scenario{
						opts: initOpts,
						ops: []op{
							// allocate enough pages, such that the page count exceeds the new maxSize upon resize
							allocate(initPages/2-1, true),

							// update max size
							closeFile,
							openWith(Options{
								MaxSize: pageSize * initPages / 2,
								Flags:   FlagUpdMaxSize,
							}),

							// optional: close and open file without extra Options
							optReopenWith(reopen, Options{}),

							// check allocation fails
							allocate(1, false),
						},
					},

					"bounding free data regions at end of file, including new size": scenario{
						opts: Options{
							MaxSize:      100 * pageSize,
							PageSize:     pageSize,
							InitMetaArea: 16,
						},
						ops: []op{
							allocateAll,
							releasePages(region{id: 40, count: 60}.PageIDs()),

							// update max size
							closeFile,
							openWith(Options{
								MaxSize: pageSize * 50,
								Flags:   FlagUpdMaxSize,
							}),

							// optional: close and open file without extra Options
							optReopenWith(reopen, Options{}),

							// alocate free pages
							allocate(10, true),

							// check we can not allocate any new pages
							allocate(1, false),
						},
					},

					"bounding free data region at end of file, excluding new size": scenario{
						opts: Options{
							MaxSize:      100 * pageSize,
							PageSize:     pageSize,
							InitMetaArea: 16,
						},
						ops: []op{
							allocateAll,
							releasePages(region{id: 50, count: 50}.PageIDs()),

							// update max size
							closeFile,
							openWith(Options{
								MaxSize: pageSize * 50,
								Flags:   FlagUpdMaxSize,
							}),

							// optional: close and open file without extra Options
							optReopenWith(reopen, Options{}),

							// check we can not allocate any new pages
							allocate(1, false),
						},
					},

					"bounding free data region at end of file, overflow": scenario{
						opts: Options{
							MaxSize:      100 * pageSize,
							PageSize:     pageSize,
							InitMetaArea: 16,
						},
						ops: []op{
							allocateAll,
							releasePages(region{id: 60, count: 40}.PageIDs()),

							// update max size
							closeFile,
							openWith(Options{
								MaxSize: pageSize * 50,
								Flags:   FlagUpdMaxSize,
							}),

							// optional: close and open file without extra Options
							optReopenWith(reopen, Options{}),

							// check we can not allocate any new pages
							allocate(1, false),

							// free overflow pages
							releasePages(region{id: 50, count: 10}.PageIDs()),

							// check we still can not allocate any new pages
							allocate(1, false),

							// free some more space in the reduced data area
							releasePages(region{id: 40, count: 10}.PageIDs()),

							// check we can allocate freed pages, but no more:
							allocate(10, true),
							allocate(1, false),
						},
					},
				})
			})
		}
	})
}

func setupPath(assert *assertions, file string) (string, func()) {
	return osfstest.SetupPath(assert, file)
}

func setupTestFile(assert *assertions, opts Options) (*testFile, func()) {
	// if opts.MaxSize == 0 {
	// 	opts.MaxSize = 10 * 1 << 20 // 10 MB
	// }
	if opts.PageSize == 0 {
		opts.PageSize = 4096
	}

	if testing.Short() {
		opts.Sync = SyncNone
	}

	ok := false
	path, teardown := setupPath(assert, "")
	defer cleanup.IfNot(&ok, teardown)

	tf := &testFile{path: path, assert: assert, opts: opts}
	tf.Open()

	ok = true
	return tf, func() {
		tf.Close()
		teardown()
	}
}

func (f *testFile) Reopen() {
	f.Close()
	f.Open()
}

func (f *testFile) Close() {
	if f.File != nil {
		f.assert.NoError(f.File.Close(), "close failed on reopen")
		f.File = nil
	}
}

func (f *testFile) Open() {
	f.OpenWith(f.opts)
}

func (f *testFile) OpenWith(opts Options) {
	tmp, err := Open(f.path, os.ModePerm, opts)
	f.assert.FatalOnError(err, "reopen failed")
	f.File = tmp

	f.checkConsistency()
}

func (f *testFile) Commit(tx *Tx, msgAndArgs ...interface{}) {
	needsCheck := tx.Writable()
	f.assert.FatalOnError(tx.Commit(), msgAndArgs...)
	if needsCheck {
		if !f.checkConsistency() {
			f.assert.Fatal("inconsistent file state")
		}
	}
}

func (f *testFile) Rollback(tx *Tx, msgAndArgs ...interface{}) {
	needsCheck := tx.Writable()
	f.assert.FatalOnError(tx.Rollback(), msgAndArgs...)
	if needsCheck {
		if !f.checkConsistency() {
			f.assert.Fatal("inconsistent file state")
		}
	}
}

// checkConsistency checks on disk file state is correct and consistent with in
// memory state.
func (f *testFile) checkConsistency() bool {
	if err := f.getMetaPage().Validate(); err != nil {
		f.assert.Error(err, "meta page validation")
		return false
	}

	meta := f.getMetaPage()
	ok := true

	// check wal:
	walMapping, walPages, err := readWAL(f.mmapedPage, meta.wal.Get())
	f.assert.FatalOnError(err, "reading wal state")
	ok = ok && f.assert.Equal(walMapping, f.wal.mapping, "wal mapping")
	ok = ok && f.assert.Equal(walPages.Regions(), f.wal.metaPages, "wal meta pages")

	maxSize := meta.maxSize.Get() / uint64(meta.pageSize.Get())
	dataEnd := meta.dataEndMarker.Get()
	metaEnd := meta.metaEndMarker.Get()

	// validate meta end markers state
	if !f.disableEndMarkerCheck {
		ok = ok && f.assert.True(maxSize == 0 || uint64(dataEnd) <= maxSize, "data end marker in bounds")
	}

	// compare alloc markers and counters with allocator state
	ok = ok && f.assert.Equal(dataEnd, f.allocator.data.endMarker, "data end marker mismatch")
	ok = ok && f.assert.Equal(metaEnd, f.allocator.meta.endMarker, "meta end marker mismatch")
	ok = ok && f.assert.Equal(uint(meta.metaTotal.Get()), f.allocator.metaTotal, "meta area size mismatch")

	// compare free lists
	var metaList, dataList regionList
	flPages, err := readFreeList(f.mmapedPage, meta.freelist.Get(), func(isMeta bool, r region) {
		if isMeta {
			metaList.Add(r)
		} else {
			dataList.Add(r)
		}
	})
	optimizeRegionList(&metaList)
	optimizeRegionList(&dataList)
	f.assert.FatalOnError(err, "reading freelist")
	ok = ok && f.assert.Equal(metaList, f.allocator.meta.freelist.regions, "meta area freelist mismatch")
	ok = ok && f.assert.Equal(dataList, f.allocator.data.freelist.regions, "data area freelist mismatch")
	ok = ok && f.assert.Equal(flPages.Regions(), f.allocator.freelistPages, "freelist meta pages")

	return ok
}

func (f *testFile) BeginWith(opts TxOptions) *Tx {
	tx, err := f.File.BeginWith(opts)
	f.assert.FatalOnError(err)
	return tx
}

func (f *testFile) Begin() *Tx {
	tx, err := f.File.Begin()
	f.assert.FatalOnError(err)
	return tx
}

func (f *testFile) BeginReadonly() *Tx {
	tx, err := f.File.BeginReadonly()
	f.assert.FatalOnError(err)
	return tx
}

func (f *testFile) withTx(write bool, fn func(tx *Tx)) {
	tx := f.BeginWith(TxOptions{Readonly: !write})
	defer func() {
		f.assert.FatalOnError(tx.Close())
	}()
	fn(tx)
}

func (f *testFile) freePages(ids idList) {
	f.withTx(true, func(tx *Tx) {
		f.txFreePages(tx, ids)
		f.assert.FatalOnError(tx.Commit())
	})
}

func (f *testFile) txFreePages(tx *Tx, ids idList) {
	for _, id := range ids {
		page, err := tx.Page(id)
		f.assert.FatalOnError(err)

		page.Free()
	}
}

func (f *testFile) append(contents []string) (ids idList) {
	f.withTx(true, func(tx *Tx) {
		ids = f.txAppend(tx, contents)
		f.assert.FatalOnError(tx.Commit())
	})
	return
}

func (f *testFile) txAppend(tx *Tx, contents []string) idList {
	ids := make(idList, len(contents))
	pages, err := tx.AllocN(len(contents))
	f.assert.FatalOnError(err)
	for i, page := range pages {
		buf := make([]byte, tx.PageSize())
		data := contents[i]
		castU32(buf).Set(uint32(len(data)))
		copy(buf[4:], data)
		f.assert.NoError(page.SetBytes(buf))

		ids[i] = page.ID()
	}
	return ids
}

func (f *testFile) read(id PageID) (s string) {
	f.withTx(false, func(tx *Tx) { s = f.txRead(tx, id) })
	return
}

func (f *testFile) txRead(tx *Tx, id PageID) string {
	page, err := tx.Page(id)
	f.assert.FatalOnError(err)

	buf, err := page.Bytes()
	f.assert.FatalOnError(err)

	count := castU32(buf).Get()
	return string(buf[4 : 4+count])
}

func (f *testFile) readIDs(ids idList) (contents []string) {
	f.withTx(false, func(tx *Tx) { contents = f.txReadIDs(tx, ids) })
	return
}

func (f *testFile) txReadIDs(tx *Tx, ids idList) []string {
	contents := make([]string, len(ids))
	for i, id := range ids {
		contents[i] = f.txRead(tx, id)
	}
	return contents
}

func (f *testFile) writeAt(id PageID, msg string) {
	f.withTx(true, func(tx *Tx) {
		f.txWriteAt(tx, id, msg)
		f.assert.FatalOnError(tx.Commit())
	})
}

func (f *testFile) txWriteAt(tx *Tx, id PageID, msg string) {
	page, err := tx.Page(id)
	f.assert.FatalOnError(err)
	buf := make([]byte, tx.PageSize())
	castU32(buf).Set(uint32(len(msg)))
	copy(buf[4:], msg)
	f.assert.NoError(page.SetBytes(buf))
}

func (f *testFile) writeAllAt(ids idList, msgs []string) {
	f.withTx(true, func(tx *Tx) {
		f.txWriteAllAt(tx, ids, msgs)
		f.assert.FatalOnError(tx.Commit())
	})
}

func (f *testFile) txWriteAllAt(tx *Tx, ids idList, msgs []string) {
	for i, id := range ids {
		f.txWriteAt(tx, id, msgs[i])
	}
}
