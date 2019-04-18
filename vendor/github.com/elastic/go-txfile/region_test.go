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
	"math"
	"testing"
)

func TestRegion(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("query", func(assert *assertions) {
		r := region{1, 5} // pages 1,2,3,4,5
		first := PageID(1)
		last := PageID(5)

		assert.Equal(first, r.Start())
		assert.Equal(last+1, r.End())

		start, end := r.Range()
		assert.Equal(first, start)
		assert.Equal(last+1, end)

		assert.True(r.InRange(3))
		assert.False(r.InRange(10))
	})

	assert.Run("comparisons", func(assert *assertions) {
		r1, r2, r3 := region{1, 2}, region{10, 1}, region{3, 4}

		assert.True(r1.Before(r2))
		assert.True(r1.Before(r3))
		assert.False(r2.Before(r1))
		assert.False(r2.Before(r3))

		assert.False(r1.Precedes(r2))
		assert.True(r1.Precedes(r3))
		assert.False(r3.Precedes(r1))
		assert.False(r3.Precedes(r2))

		assert.False(regionsMergable(r1, r2))
		assert.True(regionsMergable(r1, r3))
		assert.True(regionsMergable(r3, r1))
	})

	assert.Run("iter pages from empty region", func(assert *assertions) {
		var ids idList
		region{}.EachPage(ids.Add)
		assert.Nil(ids)
	})

	assert.Run("iter pages", func(assert *assertions) {
		var ids idList
		region{id: 1, count: 4}.EachPage(ids.Add)
		assert.Equal(idList{1, 2, 3, 4}, ids)
	})

	assert.Run("split", func(assert *assertions) {
		assert.Run("split id before region", func(assert *assertions) {
			assert.Equal(region{}, region{10, 5}.SplitAt(2))
		})

		assert.Run("split id after region", func(assert *assertions) {
			assert.Equal(region{}, region{10, 5}.SplitAt(20))
		})

		assert.Run("split id within region", func(assert *assertions) {
			assert.Equal(region{10, 2}, region{10, 5}.SplitAt(12))
		})

		assert.Run("split id start of region", func(assert *assertions) {
			assert.Equal(region{}, region{10, 5}.SplitAt(10))
		})

		assert.Run("split id at end of region", func(assert *assertions) {
			assert.Equal(region{10, 5}, region{10, 5}.SplitAt(15))
		})
	})

	assert.Run("serialization", func(assert *assertions) {
		regions := []region{
			region{10, 5},
			region{10, 10000},
		}
		titles := []string{"small", "large"}
		metaFlags := []bool{false, true}
		expectedSz := []int{8, 12}

		for i, r := range regions {
			for _, isMeta := range metaFlags {
				sz := expectedSz[i]
				isMeta := isMeta
				region := r

				title := fmt.Sprintf("%v region with meta=%v", titles[i], isMeta)
				assert.Run(title, func(assert *assertions) {
					var buf [64]byte

					assert.Equal(sz, regionEncodingSize(region))

					n := encodeRegion(buf[:], isMeta, region)
					assert.Equal(sz, n)

					metaFlag, dec, m := decodeRegion(buf[:n])
					assert.Equal(isMeta, metaFlag, "wrong meta flag")
					assert.Equal(region, dec, "region does not match")
					assert.Equal(n, m, "encoding and decoding byte sizes differ")
				})
			}
		}

		assert.Run("randomized", func(assert *assertions) {
			assert.QuickCheck(func(id PageID, count uint32, meta bool) bool {
				region := region{id: id, count: count}

				sz := regionEncodingSize(region)
				buf := make([]byte, sz)
				n := encodeRegion(buf, meta, region)
				metaFlag, dec, m := decodeRegion(buf[:n])

				ok := assert.Equal(meta, metaFlag)
				ok = ok && assert.Equal(n, m)
				ok = ok && assert.Equal(region, dec)
				return ok
			})
		})
	})
}

func TestRegionList(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("query nil list", func(assert *assertions) {
		var l regionList
		assert.Equal(0, l.Len())
		assert.NotPanics(l.Sort)
		assert.NotPanics(l.MergeAdjacent)
		assert.Equal(uint(0), l.CountPages())
		assert.Equal(uint(0), l.CountPagesUpTo(PageID(math.MaxUint64)))
		assert.Nil(l.PageIDs())
	})

	assert.Run("query empty list", func(assert *assertions) {
		l := regionList{}
		assert.Equal(0, l.Len())
		assert.NotPanics(l.Sort)
		assert.NotPanics(l.MergeAdjacent)
		assert.Equal(uint(0), l.CountPages())
		assert.Equal(uint(0), l.CountPagesUpTo(PageID(math.MaxUint64)))
		assert.Nil(l.PageIDs())
	})

	assert.Run("sort", func(assert *assertions) {
		l := regionList{{10, 1}, {2, 4}, {1, 1}}
		l.Sort()
		assert.Equal(regionList{{1, 1}, {2, 4}, {10, 1}}, l)
	})

	assert.Run("merge adjacent", func(assert *assertions) {
		l := regionList{{1, 1}, {2, 3}, {5, 5}, {10, 1}, {20, 4}, {24, 2}}
		l.MergeAdjacent()
		assert.Equal(regionList{{1, 10}, {20, 6}}, l)
	})

	assert.Run("count", func(assert *assertions) {
		l := regionList{{1, 1}, {2, 3}, {5, 5}, {10, 1}, {20, 4}, {24, 2}}
		assert.Equal(16, int(l.CountPages()))
		assert.Equal(0, int(l.CountPagesUpTo(1)))
		assert.Equal(5, int(l.CountPagesUpTo(6)))
		assert.Equal(10, int(l.CountPagesUpTo(15)))
		assert.Equal(16, int(l.CountPagesUpTo(30)))
	})

	assert.Run("iter", func(assert *assertions) {
		var tmp regionList
		l := regionList{{1, 2}, {4, 3}, {10, 2}}
		l.EachRegion(tmp.Add)
		assert.Equal(tmp, l)

		assert.Equal(idList{1, 2, 4, 5, 6, 10, 11}, l.PageIDs())
	})

	assert.Run("optimize", func(assert *assertions) {
		l := regionList{{10, 1}, {1, 1}, {20, 4}, {5, 5}, {2, 3}, {24, 2}}
		optimizeRegionList(&l)
		assert.Equal(regionList{{1, 10}, {20, 6}}, l)
	})

	assert.Run("merge", func(assert *assertions) {
		assert.Run("empty lists", func(assert *assertions) {
			assert.Nil(mergeRegionLists(nil, nil))
		})

		assert.Run("list1 empty", func(assert *assertions) {
			l := regionList{{1, 2}, {4, 3}, {10, 2}}
			assert.Equal(l, mergeRegionLists(nil, l))
		})

		assert.Run("list2 empty", func(assert *assertions) {
			l := regionList{{1, 2}, {4, 3}, {10, 2}}
			assert.Equal(l, mergeRegionLists(l, nil))
		})

		assert.Run("list1 before list2", func(assert *assertions) {
			l1 := regionList{{1, 2}, {4, 3}}
			l2 := regionList{{10, 2}, {14, 3}}
			assert.Equal(regionList{{1, 2}, {4, 3}, {10, 2}, {14, 3}}, mergeRegionLists(l1, l2))
		})

		assert.Run("list2 before list1", func(assert *assertions) {
			l1 := regionList{{10, 2}, {14, 3}}
			l2 := regionList{{1, 2}, {4, 3}}
			assert.Equal(regionList{{1, 2}, {4, 3}, {10, 2}, {14, 3}}, mergeRegionLists(l1, l2))
		})

		assert.Run("interleaved", func(assert *assertions) {
			l1 := regionList{{4, 3}, {10, 2}}
			l2 := regionList{{1, 2}, {14, 3}}
			assert.Equal(regionList{{1, 2}, {4, 3}, {10, 2}, {14, 3}}, mergeRegionLists(l1, l2))
		})
	})
}
