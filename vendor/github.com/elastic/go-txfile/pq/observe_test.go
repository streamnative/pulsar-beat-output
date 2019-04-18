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

package pq

import (
	"runtime"
	"testing"

	txfile "github.com/elastic/go-txfile"
	"github.com/elastic/go-txfile/internal/mint"
)

type testObserveLast statEntry

type statEntry struct {
	kind statKind
	off  uintptr

	// OnInit statis
	version   uint32
	available uint

	// per callback stats
	flush FlushStats
	read  ReadStats
	ack   ACKStats
}

type statKind uint8

const (
	statNone statKind = iota
	statOnOpen
	statOnFlush
	statOnRead
	statOnACK
)

var isWindows bool

func init() {
	isWindows = runtime.GOOS == "windows"
}

func TestObserveStats(testing *testing.T) {
	const testPageSize = 1024
	const testMaxPages = 128
	const testMaxSize = testMaxPages * testPageSize

	t := mint.NewWith(testing, func(sub *mint.T) func() {
		pushTracer(mint.NewTestLogTracer(sub, logTracer))
		return popTracer
	})

	withQueue := func(fn func(*mint.T, *testQueue, *statEntry)) func(*mint.T) {
		return func(t *mint.T) {
			stat := &statEntry{}
			qu, teardown := setupQueue(t, config{
				File: txfile.Options{
					MaxSize:  testMaxSize, // default file size of 128 pages
					PageSize: testPageSize,
				},
				Queue: Settings{
					WriteBuffer: defaultMinPages, // buffer up to 5 pages
					Observer:    newTestObserver(stat),
				},
			})
			defer teardown()

			stat.reset()
			fn(t, qu, stat)
		}
	}

	t.Run("open", func(t *mint.T) {
		t.Run("empty queue", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.Reopen()
			t.Equal(stat.kind, statOnOpen)
			t.Equal(int(stat.version), queueVersion) // current version
			t.Equal(int(stat.available), 0)
		}))

		t.Run("non-empty queue", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			// write 3 events
			qu.append("a", "b", "c")
			qu.flush()

			// validate
			qu.Reopen()
			t.Equal(stat.kind, statOnOpen)
			t.Equal(int(stat.version), queueVersion) // current version
			t.Equal(int(stat.available), 3)
		}))
	})

	t.Run("write", func(t *mint.T) {
		t.Run("explicit flush", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("a", "bc", "def")
			t.Equal(statNone, stat.kind, "unexpected stats update")

			qu.flush()
			t.Equal(statOnFlush, stat.kind)
			t.Equal(FlushStats{
				// do not compare duration and timestamps ;)
				Duration: stat.flush.Duration,
				Oldest:   stat.flush.Oldest,
				Newest:   stat.flush.Newest,

				// validated fields
				Failed:      false,
				OutOfMemory: false,
				Pages:       1,
				Allocate:    1,
				Events:      3,
				BytesTotal:  6,
				BytesMin:    1,
				BytesMax:    3,
			}, stat.flush)
			t.True(stat.flush.Duration > 0, "flush duration should be > 0")
			t.False(stat.flush.Oldest.IsZero(), "oldest timestamp must not be 0")
			t.False(stat.flush.Newest.IsZero(), "newest timestamp must not be 0")

			if !isWindows {
				t.True(stat.flush.Oldest != stat.flush.Newest, "timestamps do match")
			}
		}))

		t.Run("implicit flush", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			var msg [5000]byte
			qu.append(string(msg[:]))

			t.Equal(statOnFlush, stat.kind)
			t.Equal(FlushStats{
				// do not compare duration and timestamps ;)
				Duration: stat.flush.Duration,
				Oldest:   stat.flush.Oldest,
				Newest:   stat.flush.Newest,

				// validated fields
				Failed:      false,
				OutOfMemory: false,
				Pages:       6,
				Allocate:    6,
				Events:      1,
				BytesTotal:  5000,
				BytesMin:    5000,
				BytesMax:    5000,
			}, stat.flush)
			t.True(stat.flush.Duration > 0, "flush duration should be > 0")
			t.False(stat.flush.Oldest.IsZero(), "oldest timestamp must not be 0")
			t.False(stat.flush.Newest.IsZero(), "newest timestamp must not be 0")
			t.True(stat.flush.Oldest == stat.flush.Newest, "timestamps do not match")
		}))

		t.Run("flush on close", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("a", "bc", "def")
			qu.Close()

			t.Equal(statOnFlush, stat.kind)
			t.Equal(FlushStats{
				// do not compare duration and timestamps ;)
				Duration: stat.flush.Duration,
				Oldest:   stat.flush.Oldest,
				Newest:   stat.flush.Newest,

				// validated fields
				Failed:      false,
				OutOfMemory: false,
				Pages:       1,
				Allocate:    1,
				Events:      3,
				BytesTotal:  6,
				BytesMin:    1,
				BytesMax:    3,
			}, stat.flush)
			t.True(stat.flush.Duration > 0, "flush duration should be > 0")
			t.False(stat.flush.Oldest.IsZero(), "oldest timestamp must not be 0")
			t.False(stat.flush.Newest.IsZero(), "newest timestamp must not be 0")

			if !isWindows {
				t.True(stat.flush.Oldest != stat.flush.Newest, "timestamps do match")
			}
		}))
	})

	t.Run("read", func(t *mint.T) {
		t.Run("empty queue", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			data := qu.read(10)
			t.Equal(0, len(data), "did read data")

			if t.Equal(statOnRead, stat.kind, "no read stat upon empty read") {
				t.Equal(ReadStats{
					Duration: stat.read.Duration,
				}, stat.read)
			}
		}))

		t.Run("one entry", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("entry one", "entry two", "entry three")
			qu.flush()

			data := qu.read(1)
			t.Equal(1, len(data), "failed to 1 entry")

			if t.Equal(statOnRead, stat.kind, "no read state") {
				t.Equal(ReadStats{
					Duration:   stat.read.Duration,
					Read:       1,
					BytesTotal: 9,
					BytesMin:   9,
					BytesMax:   9,
				}, stat.read)
			}
		}))

		t.Run("multiple entries", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("entry one", "entry two", "entry three", "entry four")
			qu.flush()

			data := qu.read(3)
			t.Equal(3, len(data), "failed to 3 entries")

			if t.Equal(statOnRead, stat.kind, "no read state") {
				t.Equal(ReadStats{
					Duration:   stat.read.Duration,
					Read:       3,
					BytesTotal: 9 + 9 + 11,
					BytesMin:   9,
					BytesMax:   11,
				}, stat.read)
			}
		}))

		t.Run("skip", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("entry one", "entry two", "entry three", "entry four")
			qu.flush()

			qu.readWith(func(r *Reader) {
				var tmp [2]byte
				_, err := r.Next()
				t.FatalOnError(err)

				_, err = r.Read(tmp[:])
				t.FatalOnError(err)

				_, err = r.Next()
				t.FatalOnError(err)
			})

			if t.Equal(statOnRead, stat.kind, "no read state") {
				t.Equal(ReadStats{
					Duration:     stat.read.Duration,
					Skipped:      1,
					BytesTotal:   2,
					BytesSkipped: 7,
				}, stat.read)
			}
		}))

		t.Run("1 event in 2 transactions", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("entry one", "entry two", "entry three", "entry four")
			qu.flush()

			qu.readWith(func(r *Reader) {
				var tmp [2]byte
				_, err := r.Next()
				t.FatalOnError(err)

				_, err = r.Read(tmp[:])
				t.FatalOnError(err)
			})
			if t.Equal(statOnRead, stat.kind, "no read stat upon empty read") {
				t.Equal(ReadStats{
					Duration: stat.read.Duration,
				}, stat.read)
			}

			qu.readWith(func(r *Reader) {
				var tmp [7]byte
				_, err := r.Read(tmp[:])
				t.FatalOnError(err)
			})
			if t.Equal(statOnRead, stat.kind, "no read stat upon empty read") {
				t.Equal(ReadStats{
					Duration:   stat.read.Duration,
					Read:       1,
					BytesTotal: 9,
					BytesMin:   9,
					BytesMax:   9,
				}, stat.read)
			}
		}))
	})

	t.Run("ack", func(t *mint.T) {
		t.Run("fail on empty queue", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			err := qu.ACK(2)
			t.Error(err)

			t.Equal(statOnACK, stat.kind)
			t.Equal(ACKStats{
				Duration: stat.ack.Duration,
				Failed:   true,
				Events:   2,
			}, stat.ack)
		}))

		t.Run("event in active page", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("entry one", "entry two", "entry three", "entry four")
			qu.flush()

			t.FatalOnError(qu.ACK(1))
			t.Equal(ACKStats{
				Duration: stat.ack.Duration,
				Failed:   false,
				Events:   1,
			}, stat.ack)
		}))

		t.Run("all events - single page", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append("entry one", "entry two", "entry three", "entry four")
			qu.flush()

			t.FatalOnError(qu.ACK(4))
			t.Equal(ACKStats{
				Duration: stat.ack.Duration,
				Failed:   false,
				Events:   4,
				Pages:    0, // we don't free last page so to not interfere with concurrent write
			}, stat.ack)
		}))

		t.Run("large events", withQueue(func(t *mint.T, qu *testQueue, stat *statEntry) {
			qu.append(genEventLen(testPageSize*2.5), genEventLen(testPageSize*1.5))
			qu.flush()

			t.FatalOnError(qu.ACK(2))
			t.Equal(ACKStats{
				Duration: stat.ack.Duration,
				Failed:   false,
				Events:   2,
				Pages:    2, // we only free pages for the first event, so to not mess with the readers local state
			}, stat.ack)
		}))
	})
}

func newTestObserver(stat *statEntry) *testObserveLast {
	return (*testObserveLast)(stat)
}

func (t *testObserveLast) OnQueueInit(off uintptr, version uint32, available uint) {
	t.set(off, statOnOpen, statEntry{version: version, available: available})
}

func (t *testObserveLast) OnQueueFlush(off uintptr, stats FlushStats) {
	t.set(off, statOnFlush, statEntry{flush: stats})
}

func (t *testObserveLast) OnQueueRead(off uintptr, stats ReadStats) {
	t.set(off, statOnRead, statEntry{read: stats})
}

func (t *testObserveLast) OnQueueACK(off uintptr, stats ACKStats) {
	t.set(off, statOnACK, statEntry{ack: stats})
}

func (t *testObserveLast) set(off uintptr, kind statKind, e statEntry) {
	*t = testObserveLast(e)
	t.off, t.kind = off, kind
}

func (s *statEntry) reset() { *s = statEntry{} }

func genEventLen(n int) string {
	return string(make([]byte, n))
}
