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
	"math"
	"math/rand"
	"testing"

	"github.com/urso/qcgen"
)

type allocatorTest func(*allocator) bool

type allocatorTestRunner func(assert *assertions, cfg allocatorConfig, tst allocatorTest)

type allocatorConfig struct {
	// Minimum continuous pages that must be available in new allocation area. At
	// least one region will be >= MinContPages.
	MinContPages uint

	// Minium number of free pages that must be available.
	MinFreePages uint

	// propability of a freespace region to be assigned to the meta area.
	PropMetaPage float64

	// EndMarker and MaxPages. One must be set.
	EndMarker uint
	MaxPages  uint
}

func TestDataAllocator(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("fresh allocator", func(assert *assertions) {
		prepareAlloc := func(pages uint) (*allocator, *dataAllocator, txAllocState) {
			state := makeFreshAllocator(allocatorConfig{MinFreePages: pages})
			allocator := state.DataAllocator()
			tx := state.makeTxAllocState(false, 0)
			return state, allocator, tx
		}

		assert.Run("available space", func(assert *assertions) {
			assert.Run("bound allocator", func(assert *assertions) {
				_, allocator, tx := prepareAlloc(100)
				assert.Equal(100, int(allocator.Avail(&tx)))
			})

			assert.Run("unbound allocator", func(assert *assertions) {
				_, allocator, tx := prepareAlloc(0)
				assert.Equal(noLimit, allocator.Avail(&tx))
			})
		})

		// common data allocator tests
		testDataAllocator(assert,
			func(assert *assertions, cfg allocatorConfig, tst allocatorTest) {
				tst(makeFreshAllocator(cfg))
			})
	})

	// QuickCheck based allocator tests, with randomized initial allocation state
	assert.Run("with freelist entries", func(assert *assertions) {
		// Run randomized data allocator tests. The allocator state is randomized
		// on each run.
		testDataAllocator(assert,
			func(assert *assertions, cfg allocatorConfig, tst allocatorTest) {
				assert.QuickCheck(makeAllocatorState(cfg), tst)
			})
	})
}

func testDataAllocator(assert *assertions, runner allocatorTestRunner) {
	prepareAlloc := func(st *allocator) (*dataAllocator, txAllocState) {
		return st.DataAllocator(), st.makeTxAllocState(false, 0)
	}

	run := func(title string, cfg allocatorConfig, fn func(*assertions, *allocator) bool) {
		assert.Run(title, func(assert *assertions) {
			runner(assert, cfg, func(st *allocator) bool {
				return fn(assert, st)
			})
		})
	}

	run("alloc regions", allocatorConfig{
		MaxPages:     500,
		MinFreePages: 20,
	}, func(assert *assertions, state *allocator) bool {
		var regions regionList
		allocator, tx := prepareAlloc(state)
		n := allocator.AllocRegionsWith(&tx, 20, regions.Add)
		return assert.Equal(20, int(n)) &&
			assert.Equal(20, int(regions.CountPages()))
	})

	run("alloc continuous", allocatorConfig{
		MinContPages: 100,
		MaxPages:     1 << 20,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		avail := allocator.Avail(&tx)
		r := allocator.AllocContinuousRegion(&tx, 100)
		return assert.Equal(100, int(r.count)) &&
			assert.Equal(avail-100, allocator.Avail(&tx))
	})

	run("impossible alloc regions", allocatorConfig{
		MaxPages: 100,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		var regions regionList
		n := allocator.AllocRegionsWith(&tx, 200, regions.Add)
		return assert.Equal(0, int(n)) &&
			assert.Nil(regions)
	})

	run("impossible alloc continuous", allocatorConfig{
		MaxPages: 100,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		r := allocator.AllocContinuousRegion(&tx, 200)
		return assert.Equal(PageID(0), r.id)
	})

	run("allocate and free regions", allocatorConfig{
		MaxPages:     10000,
		MinFreePages: 100,
		MinContPages: 100,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		r := allocator.AllocContinuousRegion(&tx, 100)
		if !assert.Equal(100, int(r.count), "alloc mismatch") {
			return false
		}

		// free region
		r.EachPage(func(id PageID) { allocator.Free(&tx, id) })

		// try to allocate same freed region again
		expected := r
		r = allocator.AllocContinuousRegion(&tx, 100)
		return assert.Equal(expected, r)
	})

	run("rollback restores original data region after allocations", allocatorConfig{
		MinFreePages: 50,
		MaxPages:     500,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		orig := copyAllocArea(state.data)

		var regions regionList
		allocator.AllocRegionsWith(&tx, 50, regions.Add)
		if !assert.Equal(50, int(regions.CountPages())) {
			return false
		}

		// check state has been changed
		assert.NotEqual(orig, state.data)

		// undo allocations
		state.Rollback(&tx)
		return assert.Equal(orig, state.data)
	})

	run("free pages after commit", allocatorConfig{
		MinFreePages: 50,
		MaxPages:     500,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		origAvail := allocator.Avail(&tx)

		// allocate regions in first transaction
		var regions regionList
		allocator.AllocRegionsWith(&tx, 50, regions.Add)
		if !assert.Equal(origAvail-50, allocator.Avail(&tx), "invalid allocator state") {
			return false
		}

		// simulate commit by creating a new tx allocation state:
		tx = state.makeTxAllocState(false, 0)

		// return all pages in new transaction
		regions.EachPage(func(id PageID) { allocator.Free(&tx, id) })

		// commit transaction state:
		state.data.commit(
			state.data.endMarker,
			mergeRegionLists(state.data.freelist.regions, tx.data.freed.Regions()))

		// validate
		tx = state.makeTxAllocState(false, 0)
		return assert.Equal(int(origAvail), int(allocator.Avail(&tx)))
	})
}

func TestMetaAllocator(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("fresh allocator", func(assert *assertions) {
		prepareAlloc := func(pages uint) (*allocator, *metaAllocator, txAllocState) {
			state := makeFreshAllocator(allocatorConfig{MinFreePages: pages})
			allocator := state.MetaAllocator()
			tx := state.makeTxAllocState(false, 0)
			return state, allocator, tx
		}

		assert.Run("available space", func(assert *assertions) {
			assert.Run("bound allocator", func(assert *assertions) {
				_, allocator, tx := prepareAlloc(100)
				assert.Equal(100, int(allocator.Avail(&tx)))
			})

			assert.Run("unbound allocator", func(assert *assertions) {
				_, allocator, tx := prepareAlloc(0)
				assert.Equal(noLimit, allocator.Avail(&tx))
			})
		})

		// common data allocator tests
		testMetaAllocator(assert,
			func(assert *assertions, cfg allocatorConfig, tst allocatorTest) {
				tst(makeFreshAllocator(cfg))
			})
	})

	// QuickCheck based allocator tests, with randomized initial allocation state
	assert.Run("with freelist entries", func(assert *assertions) {
		// Run randomized data allocator tests. The allocator state is randomized
		// on each run.
		testMetaAllocator(assert,
			func(assert *assertions, cfg allocatorConfig, tst allocatorTest) {
				assert.QuickCheck(makeAllocatorState(cfg), tst)
			})
	})

	assert.Run("overflow release", func(assert *assertions) {
		type input struct {
			regions             regionList
			maxPages, endMarker uint64
		}

		type expected struct {
			regions regionList
			freed   uint64
		}

		testcases := []struct {
			title    string
			input    input
			expected expected
		}{
			{
				"do not free anything if no file limit",
				input{
					regions:   regionList{{1, 1000}, {2000, 5000}},
					maxPages:  0,
					endMarker: 7000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 5000}},
				},
			},
			{
				"do not free if end marker < max pages, freelist ends at end marker",
				input{
					regions:   regionList{{1, 1000}, {2000, 5000}},
					maxPages:  10000,
					endMarker: 7000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 5000}},
				},
			},
			{
				"do not free if end marker < max pages, freelist ends before end marker",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}},
					maxPages:  10000,
					endMarker: 7000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}},
				},
			},
			{
				"do not free if end marker > max pages, but last region not adjacent to end marker",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}, {4200, 500}},
					maxPages:  4000,
					endMarker: 5000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}, {4200, 500}},
				},
			},
			{
				"free, remove complete region",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}, {4500, 500}},
					maxPages:  4000,
					endMarker: 5000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}},
					freed:   500,
				},
			},
			{
				"free, remove complete region, region adjacent to max pages",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}, {4000, 1000}},
					maxPages:  4000,
					endMarker: 5000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}},
					freed:   1000,
				},
			},
			{
				"free, remove adjacent complete regions",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}, {4000, 500}, {4500, 500}},
					maxPages:  4000,
					endMarker: 5000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}},
					freed:   1000,
				},
			},
			{
				"free, split region",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}, {3000, 2000}},
					maxPages:  4000,
					endMarker: 5000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}, {3000, 1000}},
					freed:   1000,
				},
			},
			{
				"free, split region with adjacent/unmerged regions",
				input{
					regions:   regionList{{1, 1000}, {2000, 1000}, {3000, 1500}, {4500, 500}},
					maxPages:  4000,
					endMarker: 5000,
				},
				expected{
					regions: regionList{{1, 1000}, {2000, 1000}, {3000, 1000}},
					freed:   1000,
				},
			},
		}

		for _, test := range testcases {
			input, expected := test.input, test.expected
			assert.Run(test.title, func(assert *assertions) {
				list := make(regionList, len(input.regions))
				copy(list, test.input.regions)

				actual, freed := releaseOverflowPages(list, uint(input.maxPages), PageID(input.endMarker))
				assert.Equal(expected.freed, uint64(freed))
				assert.Equal(expected.regions, actual)
			})
		}
	})
}

func testMetaAllocator(assert *assertions, runner allocatorTestRunner) {
	prepareAlloc := func(st *allocator) (*metaAllocator, txAllocState) {
		return st.MetaAllocator(), st.makeTxAllocState(false, 0)
	}

	run := func(title string, cfg allocatorConfig, fn func(*assertions, *allocator) bool) {
		assert.Run(title, func(assert *assertions) {
			runner(assert, cfg, func(st *allocator) bool {
				return fn(assert, st)
			})
		})
	}

	run("alloc regions", allocatorConfig{
		MaxPages:     100,
		MinFreePages: 20,
		PropMetaPage: 0.3,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		avail := allocator.Avail(&tx)
		regions := allocator.AllocRegions(&tx, 20)
		ok := assert.Equal(20, int(regions.CountPages()))
		ok = ok && assert.Equal(avail-20, allocator.Avail(&tx))
		return ok
	})

	run("impossible alloc regions", allocatorConfig{
		MaxPages:     100,
		PropMetaPage: 0.3,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		regions := allocator.AllocRegions(&tx, 200)
		return assert.Nil(regions)
	})

	run("alloc regions with overflow area", allocatorConfig{
		MaxPages:     100,
		PropMetaPage: 0.3,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		tx.options.overflowAreaEnabled = true
		regions := allocator.AllocRegions(&tx, 150)
		return assert.Equal(150, int(regions.CountPages()))
	})

	run("rollback returns pages into meta area", allocatorConfig{
		MinFreePages: 50,
		MaxPages:     500,
		PropMetaPage: 0.3,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		origData := copyAllocArea(state.data)
		origMeta := copyAllocArea(state.meta)

		regions := allocator.AllocRegions(&tx, 50)
		if !assert.Equal(50, int(regions.CountPages())) {
			return false
		}

		state.Rollback(&tx)

		ok := true
		ok = assert.Equal(origData, state.data, "data area mismatch") && ok
		if len(origMeta.freelist.regions) > 0 {
			ok = assert.Equal(origMeta, state.meta, "meta area mismatch") && ok
		} else {
			ok = assert.Equal(0, len(state.meta.freelist.regions), "meta area mismatch") && ok
		}
		return ok
	})

	run("free pages after commit", allocatorConfig{
		MinFreePages: 50,
		MaxPages:     500,
	}, func(assert *assertions, state *allocator) bool {
		allocator, tx := prepareAlloc(state)
		origAvail := allocator.Avail(&tx)

		// allocate regions in first transaction
		regions := allocator.AllocRegions(&tx, 50)
		if !assert.Equal(origAvail-50, allocator.Avail(&tx), "invalid allocator state") {
			return false
		}

		// simulate commit by creating a new tx allocation state:
		tx = state.makeTxAllocState(false, 0)

		// return all pages in new transaction
		allocator.FreeRegions(&tx, regions)

		// commit transaction state:
		state.data.commit(
			state.data.endMarker,
			mergeRegionLists(state.data.freelist.regions, tx.data.freed.Regions()))
		state.meta.commit(
			state.meta.endMarker,
			mergeRegionLists(state.meta.freelist.regions, tx.meta.freed.Regions()))

		// validate
		tx = state.makeTxAllocState(false, 0)
		return assert.Equal(int(origAvail), int(allocator.Avail(&tx)))
	})
}

func makeFreshAllocator(cfg allocatorConfig) *allocator {
	pages := cfg.MaxPages
	if cfg.MaxPages == 0 {
		pages = cfg.EndMarker + cfg.MinContPages + cfg.MinFreePages
	}

	pageSize := uint(64)
	if pages > 0 {
		pages += 2
	}
	a := &allocator{
		maxSize:  pages * pageSize,
		maxPages: pages,
		pageSize: pageSize,
		data: allocArea{
			endMarker: 2,
		},
		meta: allocArea{
			endMarker: 2,
		},
	}

	// ensure we never return nil lists, such that asserts on list contents will
	// not complain about nil vs list of length 0
	a.data.freelist.regions = regionList{}
	a.meta.freelist.regions = regionList{}

	return a
}

func makeAllocatorState(config allocatorConfig) func(*rand.Rand) *allocator {
	if config.EndMarker == 0 && config.MaxPages == 0 {
		panic("either end marker or max size required")
	}
	if config.EndMarker > config.MaxPages {
		panic("end marker must be less then max size")
	}

	const pageSize = 64

	return func(rng *rand.Rand) *allocator {
		maxPages := config.MaxPages
		endMarker := config.EndMarker
		minContPages := config.MinContPages

		// if max pages is set, compute randomized end marker.
		var freePages uint // number of pages available after end of freelist
		if endMarker == 0 {
			maxEndMarker := config.MaxPages
			if qcgen.GenBool(rng) { // continuous space from end of file?
				maxEndMarker -= config.MinContPages
				minContPages = 0
				endMarker = qcgen.GenUintRange(rng, 2, maxEndMarker)
			} else {
				endMarker = maxEndMarker
			}

			freePages = maxPages - endMarker

			traceln("set endmarker to: ", endMarker)
			traceln("free pages at end of file: ", freePages)
		}

		// determine minimum number of pages required to be available in the freelist
		minFree := config.MinFreePages
		if minFree > 0 {
			if freePages >= minFree {
				minFree = 0
			} else {
				minFree -= freePages
			}
		}

		// Split file into allocated and free regions. Free regions are inserted
		// into the free list.
		metaList, dataList := splitRegions(rng, 2, endMarker, minFree, minContPages, config.PropMetaPage)

		// create allocator state
		a := &allocator{
			maxPages: maxPages,
			maxSize:  maxPages * pageSize,
			data: allocArea{
				endMarker: PageID(endMarker),
			},
			meta: allocArea{
				endMarker: PageID(endMarker),
			},
		}
		a.data.freelist.AddRegions(dataList)
		a.meta.freelist.AddRegions(metaList)
		a.metaTotal = metaList.CountPages()
		return a
	}
}

func splitRegions(
	rng *rand.Rand,
	start, end uint,
	minFree, minContPages uint,
	propMeta float64,
) (meta regionList, data regionList) {
	traceln("split regions with propability", propMeta)

	// ensure we never return nil, such that asserts on list contents will
	// not complain about nil vs list of length 0
	meta = regionList{}
	data = regionList{}

	addRegion := func(r region) {
		list := &data
		if propMeta > 0 && rng.Uint64() < uint64(math.MaxUint64*propMeta) {
			list = &meta
		}
		*list = append(*list, r)
	}

	count := end - start
	if count == minFree {
		addRegion(region{PageID(start), uint32(count)})
		return
	}

	if minFree < minContPages {
		minFree = minContPages
	}

	allocMode := qcgen.GenBool(rng)
	for start <= end {
		avail := end - start
		if avail == 0 {
			break
		}

		if avail == minFree || avail == minContPages {
			addRegion(region{PageID(start), uint32(avail)})
			break
		}

		if allocMode {
			// allocMode = true => select are being actively allocated/used

			if avail > 1 {
				// advance start pointer, such that at least minFree/minContPages is
				// still available after advancing the pointer
				required := minFree
				if required < minContPages {
					required = minContPages
				}

				// allocate from allocStart to allocEnd, guaranteeing enough we will
				// have enough space to fullfill MinFreePages and MinContPages
				// conditions.
				allocEnd := end - required
				allocStart := start + 1
				if allocStart < allocEnd {
					start = qcgen.GenUintRange(rng, allocStart, allocEnd)
				}
			}
		} else {
			// allocMode = false => select area of available space to be put into the
			// freelist.

			// add freelist entry
			split := end
			if avail > 1 {
				split = qcgen.GenUintRange(rng, start+1, end)
			}
			count := split - start
			if count >= minContPages {
				minContPages = 0
			}
			addRegion(region{PageID(start), uint32(count)})
			start = split

			if minFree < count {
				minFree = 0
			} else {
				minFree -= count
			}
		}

		// switch allocation mode
		allocMode = !allocMode
	}

	optimizeRegionList(&meta)
	optimizeRegionList(&data)
	traceln("new testing meta freelist: ", meta)
	traceln("new testing data freelist: ", data)
	return
}

func copyAllocArea(st allocArea) allocArea {
	new := st
	L := len(st.freelist.regions)
	if L > 0 {
		new.freelist.regions = make(regionList, L)
		copy(new.freelist.regions, st.freelist.regions)
	}

	return new
}
