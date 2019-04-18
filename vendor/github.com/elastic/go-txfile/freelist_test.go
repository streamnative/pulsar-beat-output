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
	"testing"
)

func TestFreelist(t *testing.T) {
	assert := newAssertions(t)

	mkList := func(r regionList) freelist {
		var l freelist
		l.AddRegions(r)
		return l
	}

	assert.Run("empty list", func(assert *assertions) {
		var l freelist

		assert.Nil(l.AllocAllRegions())
		assert.Equal(region{}, l.AllocContinuousRegion(allocFromBeginning, 1))

		var regions regionList
		l.AllocAllRegionsWith(regions.Add)
		assert.Nil(regions)

		regions = nil
		l.AllocRegionsWith(allocFromBeginning, 1, regions.Add)
		assert.Nil(regions)
	})

	assert.Run("adding regions consistent", func(assert *assertions) {
		var l freelist
		l.AddRegions(regionList{{1, 5}, {10, 4}})
		assert.Equal(9, int(l.Avail()))
		assert.Equal(regionList{{1, 5}, {10, 4}}, l.regions)
	})

	assert.Run("alloc all regions", func(assert *assertions) {
		regions := regionList{{1, 5}, {10, 4}}
		l := mkList(regions)
		assert.Equal(regions, l.AllocAllRegions())
		assert.Equal(0, int(l.Avail()))
	})

	assert.Run("alloc continuous", func(assert *assertions) {
		testcases := []struct {
			title     string
			init      regionList
			n         uint
			order     *allocOrder
			allocates region
			after     regionList
		}{
			{
				title: "no continuous region available",
				init:  regionList{{1, 3}, {8, 6}},
				n:     8,
				order: allocFromBeginning,
				after: regionList{{1, 3}, {8, 6}},
			},
			{
				title:     "allocate matching region from beginning",
				init:      regionList{{1, 3}, {100, 10}, {200, 10}, {300, 3}},
				n:         10,
				order:     allocFromBeginning,
				allocates: region{100, 10},
				after:     regionList{{1, 3}, {200, 10}, {300, 3}},
			},
			{
				title:     "allocate matching region from end",
				init:      regionList{{1, 3}, {100, 10}, {200, 10}, {300, 3}},
				n:         10,
				order:     allocFromEnd,
				allocates: region{200, 10},
				after:     regionList{{1, 3}, {100, 10}, {300, 3}},
			},
			{
				title:     "allocate best fit",
				init:      regionList{{1, 3}, {100, 10}, {150, 5}, {200, 10}, {300, 3}},
				n:         5,
				order:     allocFromBeginning,
				allocates: region{150, 5},
				after:     regionList{{1, 3}, {100, 10}, {200, 10}, {300, 3}},
			},
			{
				title:     "split first smallest fitting region",
				init:      regionList{{1, 3}, {100, 10}, {150, 15}, {200, 10}, {300, 3}},
				n:         8,
				order:     allocFromBeginning,
				allocates: region{100, 8},
				after:     regionList{{1, 3}, {108, 2}, {150, 15}, {200, 10}, {300, 3}},
			},
			{
				title:     "split last smallest fitting region",
				init:      regionList{{1, 3}, {100, 10}, {150, 15}, {200, 10}, {300, 3}},
				n:         8,
				order:     allocFromEnd,
				allocates: region{202, 8},
				after:     regionList{{1, 3}, {100, 10}, {150, 15}, {200, 2}, {300, 3}},
			},
		}

		for _, test := range testcases {
			test := test
			assert.Run(test.title, func(assert *assertions) {
				l := mkList(test.init)
				assert.Equal(test.allocates, l.AllocContinuousRegion(test.order, test.n), "unexpected allocation result")
				assert.Equal(int(test.after.CountPages()), int(l.Avail()), "invalid number of available pages")
				assert.Equal(test.after, l.regions, "freelist regions not matching expected freelist")
			})
		}
	})

	assert.Run("alloc non-continuous", func(assert *assertions) {
		testcases := []struct {
			title     string
			init      regionList
			n         uint
			order     *allocOrder
			allocates regionList
			after     regionList
		}{
			{
				title: "not enough space available",
				init:  regionList{{1, 3}, {8, 6}},
				n:     20,
				order: allocFromBeginning,
				after: regionList{{1, 3}, {8, 6}},
			},
			{
				title:     "allocate sub-region from first region only",
				init:      regionList{{1, 10}, {20, 5}},
				n:         5,
				order:     allocFromBeginning,
				allocates: regionList{{1, 5}},
				after:     regionList{{6, 5}, {20, 5}},
			},
			{
				title:     "allocate complete first region found",
				init:      regionList{{1, 10}, {20, 5}},
				n:         10,
				order:     allocFromBeginning,
				allocates: regionList{{1, 10}},
				after:     regionList{{20, 5}},
			},
			{
				title:     "allocate sub-region from last region only",
				init:      regionList{{1, 10}, {20, 10}},
				n:         5,
				order:     allocFromEnd,
				allocates: regionList{{25, 5}},
				after:     regionList{{1, 10}, {20, 5}},
			},
			{
				title:     "allocate complete last region found",
				init:      regionList{{1, 10}, {20, 10}},
				n:         10,
				order:     allocFromEnd,
				allocates: regionList{{20, 10}},
				after:     regionList{{1, 10}},
			},
			{
				title:     "allocate multiple regions from the beginning",
				init:      regionList{{1, 3}, {5, 3}, {10, 5}, {20, 10}},
				n:         9,
				order:     allocFromBeginning,
				allocates: regionList{{1, 3}, {5, 3}, {10, 3}},
				after:     regionList{{13, 2}, {20, 10}},
			},
			{
				title:     "allocate multiple regions from the end",
				init:      regionList{{1, 3}, {5, 4}, {10, 3}, {20, 3}},
				n:         9,
				order:     allocFromEnd,
				allocates: regionList{{6, 3}, {10, 3}, {20, 3}},
				after:     regionList{{1, 3}, {5, 1}},
			},
		}

		for _, test := range testcases {
			test := test
			assert.Run(test.title, func(assert *assertions) {
				var allocated regionList
				l := mkList(test.init)
				l.AllocRegionsWith(test.order, test.n, allocated.Add)
				assert.Equal(test.allocates, allocated, "unexpected allocation result")
				assert.Equal(int(test.after.CountPages()), int(l.Avail()), "invalid number of available pages")
				assert.Equal(test.after, l.regions, "freelist regions not matching expected freelist")
			})
		}
	})

	assert.Run("remove region", func(assert *assertions) {
		testcases := []struct {
			title  string
			init   regionList
			remove region
			after  regionList
		}{
			{
				title:  "remove region from end, not in freelist",
				init:   regionList{{1, 10}, {30, 10}, {50, 10}},
				remove: region{60, 10},
				after:  regionList{{1, 10}, {30, 10}, {50, 10}},
			},
			{
				title:  "remove region from beginning, not in freelist",
				init:   regionList{{10, 10}, {30, 10}, {50, 10}},
				remove: region{1, 5},
				after:  regionList{{10, 10}, {30, 10}, {50, 10}},
			},
			{
				title:  "remove region from middle, not in freelist",
				init:   regionList{{10, 10}, {30, 10}, {50, 10}},
				remove: region{20, 5},
				after:  regionList{{10, 10}, {30, 10}, {50, 10}},
			},
			{
				title:  "remove first region from list",
				init:   regionList{{1, 10}, {20, 10}, {50, 10}},
				remove: region{1, 10},
				after:  regionList{{20, 10}, {50, 10}},
			},
			{
				title:  "remove from beginning of first region",
				init:   regionList{{1, 10}, {20, 10}, {50, 10}},
				remove: region{1, 5},
				after:  regionList{{6, 5}, {20, 10}, {50, 10}},
			},
			{
				title:  "remove from end of first region",
				init:   regionList{{1, 10}, {20, 10}, {50, 10}},
				remove: region{6, 5},
				after:  regionList{{1, 5}, {20, 10}, {50, 10}},
			},
			{
				title:  "remove from middle of first region",
				init:   regionList{{1, 10}, {20, 10}, {50, 10}},
				remove: region{6, 2},
				after:  regionList{{1, 5}, {8, 3}, {20, 10}, {50, 10}},
			},
			{
				title:  "remove multiple regions",
				init:   regionList{{1, 10}, {20, 10}, {50, 10}},
				remove: region{1, 30},
				after:  regionList{{50, 10}},
			},
			{
				title:  "remove sub-regions from multiple regions",
				init:   regionList{{1, 10}, {12, 5}, {20, 10}, {50, 10}},
				remove: region{6, 20},
				after:  regionList{{1, 5}, {26, 4}, {50, 10}},
			},
		}

		for _, test := range testcases {
			test := test
			assert.Run(test.title, func(assert *assertions) {
				l := mkList(test.init)
				l.RemoveRegion(test.remove)
				assert.Equal(int(test.after.CountPages()), int(l.Avail()), "invalid number of available pages")
				assert.Equal(test.after, l.regions, "freelist regions not matching expected freelist")
			})
		}
	})

	assert.Run("add-merge new region", func(assert *assertions) {
		testcases := []struct {
			title            string
			init, add, after regionList
			requires64       bool // mark a test requires 64bit arch to run
		}{
			{
				title: "into empty list",
				add:   regionList{{10, 2}},
				after: regionList{{10, 2}},
			},
			{
				title: "add to end of list",
				init:  regionList{{1, 10}},
				add:   regionList{{100, 10}},
				after: regionList{{1, 10}, {100, 10}},
			},
			{
				title: "merge to end of list",
				init:  regionList{{1, 10}},
				add:   regionList{{11, 10}},
				after: regionList{{1, 20}},
			},
			{
				title: "add to beginning of list",
				init:  regionList{{100, 10}},
				add:   regionList{{10, 10}},
				after: regionList{{10, 10}, {100, 10}},
			},
			{
				title: "merge with beginning of list",
				init:  regionList{{11, 10}},
				add:   regionList{{1, 10}},
				after: regionList{{1, 20}},
			},
			{
				title: "add in middle of list",
				init:  regionList{{1, 10}, {1000, 10}},
				add:   regionList{{100, 10}},
				after: regionList{{1, 10}, {100, 10}, {1000, 10}},
			},
			{
				title: "merge end of entry in middle of list",
				init:  regionList{{1, 10}, {100, 10}, {1001, 10}},
				add:   regionList{{110, 10}},
				after: regionList{{1, 10}, {100, 20}, {1001, 10}},
			},
			{
				title: "merge start of entry in middle of list",
				init:  regionList{{1, 10}, {110, 10}, {1001, 10}},
				add:   regionList{{100, 10}},
				after: regionList{{1, 10}, {100, 20}, {1001, 10}},
			},
			{
				title: "combine regions on merge",
				init:  regionList{{1, 10}, {100, 10}, {120, 10}, {1001, 10}},
				add:   regionList{{110, 10}},
				after: regionList{{1, 10}, {100, 30}, {1001, 10}},
			},
			{
				title:      "do not combine unmergable regions into full region",
				requires64: true,
				init:       regionList{{0, 1<<32 - 10}, {1<<32 - 5, 100}},
				add:        regionList{{1<<32 - 10, 5}},
				after:      regionList{{0, 1<<32 - 5}, {1<<32 - 5, 100}},
			},
		}
		for _, test := range testcases {
			test := test
			assert.Run(test.title, func(assert *assertions) {
				if test.requires64 {
					if uint64(maxUint) != math.MaxUint64 {
						assert.Skip("test requires 64bit system to run")
					}
				}

				var l freelist
				l.AddRegions(test.init)
				for _, reg := range test.add {
					l.AddRegion(reg)
				}
				assert.Equal(int(test.after.CountPages()), int(l.Avail()), "invalid number of available pages")
				assert.Equal(test.after, l.regions, "freelist regions not matching expected freelist")
			})
		}
	})

	assert.Run("serialization", func(assert *assertions) {
		testcases := []struct {
			title string
			pages uint
			meta  regionList
			data  regionList
		}{
			{
				title: "empty lists",
				pages: 0,
			},
			{
				title: "empty lists with allocated write pages",
				pages: 3,
			},
			{
				title: "small data list only, with estimate",
				data:  regionList{{100, 10}, {200, 20}, {1000, 500}},
			},
			{
				title: "small meta list only, with estimate",
				meta:  regionList{{100, 10}, {200, 20}, {1000, 500}},
			},
			{
				title: "small data and meta list, with estimate",
				meta:  regionList{{100, 10}, {200, 20}, {1000, 500}, {1500, 300}},
				data:  regionList{{500, 10}, {2000, 20}, {10000, 500}, {40000, 200}},
			},
			{
				title: "data and meta list, with extra pages",
				pages: 10,
				meta:  regionList{{100, 10}, {200, 20}, {1000, 500}, {1500, 300}},
				data:  regionList{{500, 10}, {2000, 20}, {10000, 500}, {40000, 200}},
			},
		}

		pageSize := uint(64)
		for _, test := range testcases {
			N := test.pages
			meta, data := test.meta, test.data

			assert.Run(test.title, func(assert *assertions) {
				store := newTestPageStore(pageSize)

				if N == 0 {
					predictor := prepareFreelistEncPagePrediction(listPageHeaderSize, pageSize)
					predictor.AddRegions(meta)
					predictor.AddRegions(data)
					N = predictor.Estimate()
				}

				root := PageID(2)
				writePages := regionList{{root, uint32(N)}}
				if N == 0 { // no pages -> set root to nil
					root = 0
					writePages = nil
				}

				tracef("freelist test write regions: %v\n", writePages)

				err := writeFreeLists(writePages, pageSize, meta, data, store.Set)
				assert.FatalOnError(err, "write freelist failed")

				var actualMeta, actualData regionList
				ids, err := readFreeList(store.Get, root, func(isMeta bool, reg region) {
					if isMeta {
						actualMeta.Add(reg)
					} else {
						actualData.Add(reg)
					}
				})
				assert.FatalOnError(err, "reading the list failed")

				assert.Equal(writePages, ids.Regions())
				assert.Equal(meta, actualMeta, "meta regions mismatch")
				assert.Equal(data, actualData, "data regions mismatch")
			})
		}
	})
}
