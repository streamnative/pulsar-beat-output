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
	"math/rand"
	"testing"

	"github.com/elastic/go-txfile"
	"github.com/elastic/go-txfile/internal/cleanup"
	"github.com/elastic/go-txfile/internal/mint"
	"github.com/elastic/go-txfile/txfiletest"
)

type testQueue struct {
	t      *mint.T
	config config
	*Queue
	*txfiletest.TestFile
}

type config struct {
	File  txfile.Options
	Queue Settings
}

type testRange struct {
	min, max int
}

func exactly(n int) testRange        { return testRange{n, n} }
func between(min, max int) testRange { return testRange{min, max} }

func setupQueue(t *mint.T, cfg config) (*testQueue, func()) {
	if testing.Short() {
		cfg.File.Sync = txfile.SyncNone
	}

	tf, teardown := txfiletest.SetupTestFile(t.T, cfg.File)

	ok := false
	defer cleanup.IfNot(&ok, teardown)

	tq := &testQueue{
		t:        t,
		config:   cfg,
		TestFile: tf,
		Queue:    nil,
	}

	tq.Open()
	ok = true
	return tq, func() {
		tq.Close()
		teardown()
	}
}

func (q *testQueue) Reopen() {
	q.Close()
	q.Open()
}

func (q *testQueue) Open() {
	if q.Queue != nil {
		return
	}

	q.TestFile.Open()

	d, err := NewStandaloneDelegate(q.TestFile.File)
	if err != nil {
		q.t.Fatal(err)
	}

	tmp, err := New(d, q.config.Queue)
	if err != nil {
		q.t.Fatal(err)
	}

	q.Queue = tmp
}

func (q *testQueue) Close() {
	if q.Queue == nil {
		return
	}

	q.t.FatalOnError(q.Queue.Close())
	q.TestFile.Close()
	q.Queue = nil
}

func (q *testQueue) len() int {
	reader := q.Reader()
	q.t.FatalOnError(reader.Begin())
	defer reader.Done()

	i, err := q.Reader().Available()
	q.t.NoError(err)
	return int(i)
}

func (q *testQueue) append(events ...string) {
	w, err := q.Queue.Writer()
	q.t.FatalOnError(err)

	for _, event := range events {
		_, err := w.Write([]byte(event))
		q.t.FatalOnError(err)
		q.t.FatalOnError(w.Next())
	}
}

func (q *testQueue) readWith(fn func(*Reader)) {
	r := q.Reader()
	q.t.FatalOnError(r.Begin())
	defer r.Done()

	fn(r)
}

// read reads up to n events from the queue.
func (q *testQueue) read(n int) []string {
	var out []string
	if n > 0 {
		out = make([]string, 0, n)
	}

	reader := q.Reader()
	q.t.FatalOnError(reader.Begin())
	defer reader.Done()

	for n < 0 || len(out) < n {
		sz, err := reader.Next()
		q.t.FatalOnError(err)
		if sz <= 0 {
			break
		}

		buf := make([]byte, sz)
		_, err = reader.Read(buf)
		q.t.FatalOnError(err)

		out = append(out, string(buf))
	}
	return out
}

func (q *testQueue) flush() {
	w, err := q.Queue.Writer()
	q.t.FatalOnError(err)

	err = w.Flush()
	q.t.NoError(err)
}

func (q *testQueue) ack(n uint) {
	err := q.Queue.ACK(n)
	q.t.NoError(err)
}

func (r testRange) rand(rng *rand.Rand) int {
	if r.min == r.max {
		return r.min
	}
	return rng.Intn(r.max-r.min) + r.min
}
