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
	"errors"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/urso/qcgen"

	"github.com/elastic/go-txfile/internal/mint"
)

// assertions wraps the testing, assert, and quick packages in order to provide
// some unified functionality.
// The Run method pushes a new tracer to the tracer stack, such that trace prints
// will be captured by the current testing context.
type assertions = mint.T

type blocker struct {
	blocked bool
	mux     sync.Mutex
	cond    *sync.Cond
}

type closeWaiter struct {
	wg       sync.WaitGroup
	duration time.Duration
}

type testPageStore struct {
	pageSize uint
	pages    map[PageID][]byte
}

func newAssertions(t *testing.T) *assertions {
	a := mint.NewWith(t, func(sub *mint.T) func() {
		pushTracer(mint.NewTestLogTracer(sub, logTracer))
		return popTracer
	})

	a.SetDefaultGenerators(qcMakePageID)
	return a
}

func newBlocker() *blocker {
	b := &blocker{}
	b.cond = sync.NewCond(&b.mux)
	return b
}

func (b *blocker) Block() {
	b.mux.Lock()
	b.blocked = true
	b.mux.Unlock()
}

func (b *blocker) Unblock() {
	b.mux.Lock()
	b.blocked = false
	b.mux.Unlock()
	b.cond.Signal()
}

func (b *blocker) Wait() {
	b.mux.Lock()
	defer b.mux.Unlock()
	for b.blocked {
		b.cond.Wait()
	}
}

func makeQuickCheck(fns ...interface{}) func(*assertions) {
	return func(assert *assertions) {
		assert.QuickCheck(fns...)
	}
}

func makeCloseWait(timeout time.Duration) closeWaiter {
	return closeWaiter{duration: timeout}
}

func (w *closeWaiter) Add(n int) { w.wg.Add(n) }
func (w *closeWaiter) Done()     { w.wg.Done() }
func (w *closeWaiter) Wait() bool {
	err := waitFn(w.duration, func() reason {
		w.wg.Wait()
		return nil
	})
	return err == nil
}

func newTestPageStore(pageSize uint) *testPageStore {
	return &testPageStore{
		pages:    map[PageID][]byte{},
		pageSize: pageSize,
	}
}

func (p *testPageStore) Get(id PageID) []byte {
	if id < 2 {
		panic("must not access file meta region")
	}

	b := p.pages[id]
	if b == nil {
		b := make([]byte, p.pageSize)
		p.pages[id] = b
	}
	return b
}

func (p *testPageStore) Set(id PageID, b []byte) reason {
	if id < 2 {
		panic("must not overwrite file meta region")
	}
	p.pages[id] = b
	return nil
}

func waitFn(timeout time.Duration, fn func() reason) error {
	ch := make(chan error)
	go func() {
		ch <- fn()
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(timeout):
		return errors.New("wait timed out")
	}
}

// quick check generators
func qcMakePageID(rng *rand.Rand) PageID {
	max := uint64(1 << (64 - entryBits))
	if qcgen.GenBool(rng) {
		max = entryOverflow - 1 // generate 'small' id only
	}
	return PageID(qcgen.GenUint64Range(rng, 2, max))
}
