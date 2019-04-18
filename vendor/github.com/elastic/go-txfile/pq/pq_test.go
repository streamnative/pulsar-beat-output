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
	"github.com/elastic/go-txfile/internal/mint"
)

func TestQueueOperations(testing *testing.T) {
	t := mint.NewWith(testing, func(sub *mint.T) func() {
		pushTracer(mint.NewTestLogTracer(sub, logTracer))
		return popTracer
	})

	defaultConfig := config{
		File: txfile.Options{
			MaxSize:  128 * 1024, // default file size of 128 pages
			PageSize: 1024,
		},
		Queue: Settings{
			WriteBuffer: 4096, // buffer up to 4 pages
		},
	}

	t.Run("no events in new file", withQueue(defaultConfig, func(t *mint.T, qu *testQueue) {
		t.Equal(0, qu.len())
	}))

	t.Run("no events without flush", withQueue(defaultConfig, func(t *mint.T, qu *testQueue) {
		qu.append("event1", "event2")
		t.Equal(0, qu.len())
	}))

	t.Run("events available with flush", withQueue(defaultConfig, func(t *mint.T, qu *testQueue) {
		qu.append("event1", "event2")
		qu.flush()
		t.Equal(2, qu.len())
	}))

	t.Run("no more events after ack", withQueue(defaultConfig, func(t *mint.T, qu *testQueue) {
		qu.append("event1", "event2")
		qu.flush()
		qu.ack(2)
		t.Equal(0, qu.len())
	}))

	t.Run("read written events", withQueue(defaultConfig, func(t *mint.T, qu *testQueue) {
		events := []string{"event1", "event2"}
		qu.append(events...)
		qu.flush()
		actual := qu.read(-1)
		qu.ack(uint(len(actual)))

		t.Equal(events, actual)
		t.Equal(0, qu.len())
	}))

	t.Run("flush callback is called", func(t *mint.T) {
		var count int

		cfg := defaultConfig
		cfg.Queue.Flushed = func(n uint) {
			count += int(n)
		}

		qu, teardown := setupQueue(t, cfg)
		defer teardown()

		qu.append("a", "b", "c")
		qu.flush()
		t.Equal(3, count)
	})

	t.Run("scenarios", func(t *mint.T) {
		sz := int(defaultConfig.File.PageSize)

		qcCconfig := defaultConfig
		qcCconfig.File.MaxSize = 5000 * 500

		testcases := []struct {
			ops    scenario
			config config
		}{
			// randomized event lengths
			{
				ops: scenario{
					opRandEvents(1, sz, sz/2+sz, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(1, sz*2, sz*2+sz/2, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(1, sz*2, sz*3+sz/2, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(2, sz, sz/2+sz, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(2, sz*2, sz*2+sz/2, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(2, sz*2, sz*3+sz/2, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(10, sz/10, sz/2, 0),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(10, sz/10, sz*3, 0),
				},
				config: defaultConfig,
			},

			// randomized events with known seed (former failing tests)
			{
				ops: scenario{
					opRandEvents(2, sz*2, sz*2+sz/2, 1519332567522137000),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(2, sz*2, sz*3+sz/2, 1519334194758363000),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(2, 2048, 2560, 1519336773321643000),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(10, sz/10, sz*3, 1519336716061915000),
				},
				config: defaultConfig,
			}, {
				ops: scenario{
					opRandEvents(10, 102, 2073, 1519398160271388000),
				},
				config: defaultConfig,
			},

			// failed scenarios from quick check
			{
				ops: scenario{
					opWriteEvents(4344, 1293, 1164, 2154, 3104, 4282, 2241, 3944),
				},
				config: qcCconfig,
			}, {
				ops: scenario{
					opWriteEvents(2982, 2877, 317, 2319, 2098, 2752, 1445, 3653),
				},
				config: qcCconfig,
			}, {
				ops: scenario{
					opWriteEvents(983, 242, 166, 789, 2780, 2806, 3616, 4777, 3569),
				},
				config: qcCconfig,
			}, {
				ops: scenario{
					opWriteEvents(4770, 4139),
					opWriteEvents(3966, 231),
				},
				config: qcCconfig,
			}, {
				ops: scenario{
					opWriteEvents(4770, 4139), opReopen,
					opWriteEvents(3966, 231), opReopen,
				},
				config: qcCconfig,
			}, {
				ops: scenario{
					opWriteEvents(400, 400, 400, 400, 400),
					opReadAll,
					opACKEvents(1),
					opACKAll,
				},
				config: qcCconfig,
			},
			{
				ops: scenario{
					opWriteEvents(400, 400, 400, 400, 400),
					opReadAll,
					opACKEvents(1),
					opACKEvents(2),
					opACKEvents(1),
					opACKEvents(1),
				},
				config: qcCconfig,
			}, {
				ops: scenario{
					opWriteEvents(1500, 1500, 1500, 1500, 1500),
					opReadAll,
					opACKEvents(1),
					opACKAll,
				},
				config: qcCconfig,
			},
			{
				ops: scenario{
					opWriteEvents(1500, 1500, 1500, 1500, 1500),
					opReadAll,
					opACKEvents(1),
					opACKEvents(2),
					opACKEvents(1),
					opACKEvents(1),
				},
				config: qcCconfig,
			},

			// special failing scenarios:
			{
				ops: scenario{
					// initial writes so to reconstruct queue state
					opWriteEvents(concatSizes(
						nOfSize(24, 49),
						nOfSize(256-24, 50),
						nOfSize(374-256, 51))...),
					opReadAll,
					opACKEvents(374),
					opWriteEvents(nOfSize(296, 51)...),
					opReadAll,
					opACKEvents(296),
					opWriteEvents(nOfSize(296, 51)...),
					opReadAll,
					opACKEvents(296),
					opWriteEvents(nOfSize(296, 51)...),
					opReadAll,
					opACKEvents(296),
					opWriteEvents(nOfSize(296, 51)...),
					opReadAll,
					opACKEvents(296),
					opWriteEvents(nOfSize(296, 51)...),
					opReadAll,
					opACKEvents(296),

					// write that requires additinal allocations
					opWriteEvents(nOfSize(296, 51)...),
					opReadAll, // <- missing ACK and next push

					// write with delayed ACK
					opWriteEvents(nOfSize(295, 51)...),
					opACKEvents(296), // <- delayed ACK
					opReadAll,        // <- fails
					opACKEvents(295),
				},
				config: config{
					File: txfile.Options{
						MaxSize:  128 * 1024,
						PageSize: 4 * 1024,
						Prealloc: true,
					},
					Queue: Settings{
						WriteBuffer: 16 * 1024,
					},
				},
			},
		}

		for _, test := range testcases {
			test := test
			t.Run(test.ops.String(), func(t *mint.T) {
				runScenario(t, test.config, append(test.ops, opFinalize))
			})
		}
	})

	t.Run("quick check queue operations", func(t *mint.T) {
		config := defaultConfig
		config.File.MaxSize = 5000 * 500

		// rounds := between(1, 3)
		rounds := between(1, 4)
		// rounds := exactly(1)
		events := between(5, 50)
		// events := between(1, 10)
		sizes := between(10, 5000)

		t.Run("randomized writes", func(t *mint.T) {
			t.QuickCheck(qcGenRandWrites(rounds, events, sizes), func(ops scenario) bool {
				return runScenario(t, config, append(ops, opReadAll))
			})
		})

		t.Run("randomized acks", func(t *mint.T) {
			t.QuickCheck(qcGenRandACKs(events, sizes), func(ops scenario) bool {
				return runScenario(t, config, append(ops, opReadAll))
			})
		})

	})
}

func withQueue(cfg config, fn func(*mint.T, *testQueue)) func(*mint.T) {
	return func(t *mint.T) {
		qu, teardown := setupQueue(t, cfg)
		defer teardown()
		fn(t, qu)
	}
}

func qcGenRandWrites(rounds, N, Len testRange) func(*rand.Rand) scenario {
	return func(rng *rand.Rand) scenario {
		var ops scenario

		W := rounds.rand(rng)
		reopen := W > 1

		for i := 0; i < W; i++ {
			ops = append(ops, qcGenWriteOp(N.rand(rng), Len)(rng))
			if reopen {
				ops = append(ops, opReopen)
			}
		}
		return ops
	}
}

func qcGenWriteOp(N int, Len testRange) func(*rand.Rand) testOp {
	return func(rng *rand.Rand) testOp {
		events := make([]int, N)
		for i := range events {
			events[i] = Len.rand(rng)
		}
		return opWriteEvents(events...)
	}
}

func qcGenRandACKs(events, Len testRange) func(*rand.Rand) scenario {
	return func(rng *rand.Rand) scenario {
		var ops scenario

		N := events.rand(rng)
		ops = append(ops, qcGenWriteOp(N, Len)(rng), opReadAll)

		for N > 0 {
			acks := testRange{min: 1, max: N}.rand(rng)
			ops = append(ops, opACKEvents(acks))
			N -= acks
		}

		return ops
	}
}

func genQcRandString(L testRange) func(*rand.Rand) string {
	return func(rng *rand.Rand) string {
		codePoints := make([]byte, L.rand(rng))
		for i := range codePoints {
			codePoints[i] = byte(rng.Intn(int('Z'-'0'))) + '0'
		}
		return string(codePoints)
	}
}
