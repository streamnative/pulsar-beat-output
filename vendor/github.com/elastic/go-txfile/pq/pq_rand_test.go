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
	"fmt"
	"math/rand"
	"strings"

	"github.com/elastic/go-txfile/internal/mint"
)

type testCtx struct {
	t   *mint.T
	rng *rand.Rand

	written []string
	read    []string
	acked   int
}

type (
	scenario []testOp

	testOp struct {
		describe string
		validate validator
		exec     execer
	}

	execer    func(*testCtx, *testQueue)
	validator func(*testCtx, *testQueue)
)

func (s scenario) String() string {
	ops := make([]string, len(s))
	for i, op := range s {
		ops[i] = op.describe
	}
	return strings.Join(ops, "\n")
}

var opReopen = testOp{
	describe: "opReopen",
	validate: checkAvail,
	exec:     func(_ *testCtx, qu *testQueue) { qu.Reopen() },
}

var opReadAll = testOp{
	describe: "opReadAll",
	validate: checkAll(checkNoAvail, checkAllEvents),
	exec: func(ctx *testCtx, qu *testQueue) {
		ctx.read = append(ctx.read, qu.read(-1)...)
	},
}

var opFinalize = testOp{
	describe: "opFinalize",
	validate: checkAll(checkNoAvail, checkAllEvents),
	exec: func(ctx *testCtx, qu *testQueue) {
		ctx.read = append(ctx.read, qu.read(-1)...)
	},
}

func opWriteEvents(lens ...int) testOp {
	return testOp{
		describe: fmt.Sprintf("opWriteEvents %v: %v", len(lens), lens),
		validate: checkAvail,
		exec: func(ctx *testCtx, qu *testQueue) {
			for _, l := range lens {
				event := genQcRandString(exactly(l))(ctx.rng)
				qu.append(event)
				ctx.written = append(ctx.written, event)
			}
			qu.flush()
		},
	}
}

func opRandEvents(N, minSz, maxSz int, seed int64) testOp {
	name := fmt.Sprintf("opRandEvents %v: min size=%v, max size=%v, seed=%v",
		N, minSz, maxSz, seed)
	if seed == 0 {
		seed = mint.RngSeed()
		name = fmt.Sprintf("%v, run seed=%v", name, seed)
	}

	return testOp{
		describe: name,
		validate: checkAvail,
		exec: func(ctx *testCtx, qu *testQueue) {
			rng := mint.NewRng(seed)
			for i := 0; i < N; i++ {
				event := genQcRandString(testRange{minSz, maxSz})(rng)
				qu.append(event)
				ctx.written = append(ctx.written, event)
			}
			qu.flush()
		},
	}
}

var opACKAll = testOp{
	describe: "opACKAll",
	validate: checkNoPending,
	exec: func(ctx *testCtx, qu *testQueue) {
		delta := len(ctx.written) - ctx.acked
		qu.ack(uint(delta))
	},
}

func opACKEvents(n int) testOp {
	return testOp{
		describe: fmt.Sprintf("opACKEvents %v", n),
		validate: checkPending,
		exec: func(ctx *testCtx, qu *testQueue) {
			ctx.acked += n
			qu.ack(uint(n))
		},
	}
}

func checkAll(checks ...validator) validator {
	return func(ctx *testCtx, qu *testQueue) {
		for _, check := range checks {
			check(ctx, qu)
		}
	}
}

func checkAllEvents(ctx *testCtx, _ *testQueue) {
	ctx.t.Equal(ctx.written, ctx.read, "comparing written with read events")
}

func checkNoAvail(ctx *testCtx, qu *testQueue) {
	ctx.t.Equal(0, qu.len(), "expected no events to be available")
}

func checkAvail(ctx *testCtx, qu *testQueue) {
	ctx.t.Equal(len(ctx.written), len(ctx.read)+qu.len(), "available events mismatch")
}

func checkNoPending(ctx *testCtx, qu *testQueue) {
	cnt, err := qu.Pending()
	ctx.t.NoError(err)
	ctx.t.Equal(0, cnt)
}

func checkPending(ctx *testCtx, qu *testQueue) {
	cnt, err := qu.Pending()
	ctx.t.NoError(err)
	ctx.t.Equal(len(ctx.written)-ctx.acked, cnt)
}

func runScenario(t *mint.T, cfg config, ops scenario) bool {
	qu, teardown := setupQueue(t, cfg)
	defer teardown()

	rng := mint.NewRng(-1)
	ctx := &testCtx{t: t, rng: rng}

	for _, op := range ops {
		traceln("run op:", op.describe)

		if op.exec != nil {
			op.exec(ctx, qu)
			if t.Failed() {
				return false
			}
		}

		if op.validate != nil {
			op.validate(ctx, qu)
			if t.Failed() {
				return false
			}
		}
	}

	return !t.Failed()
}

func nOfSize(n int, sz int) []int {
	ret := make([]int, n)
	for i := range ret {
		ret[i] = sz
	}
	return ret
}

func concatSizes(all ...[]int) []int {
	var ret []int
	for _, other := range all {
		ret = append(ret, other...)
	}
	return ret
}
