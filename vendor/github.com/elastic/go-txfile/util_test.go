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
)

// TestHelpers tests some critical internal utilities for being somewhat correct
func TestInternalHelpers(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("is power of 2", func(assert *assertions) {
		for bit := uint64(0); bit < 64; bit++ {
			value := uint64(1) << bit
			assert.True(isPowerOf2(value))

			// Check masking with another arbitrary bit pattern, the result does not
			// satisfy the isPowerOf2 predicate.
			assert.QuickCheck(func(mask uint64) bool {
				// mask must modify the value, otherwise continue testing
				return (mask == 0 || mask == value) || !isPowerOf2(mask|value)
			})
			bit++
		}
	})

	assert.Run("check nextPowerOf2 consistent", makeQuickCheck(func(u uint64) bool {
		if u >= (1 << 63) {
			//  if highest bit is already set, we can not compute next power of 2
			//  (which would be 0)
			return true
		}

		v := nextPowerOf2(u)
		if !isPowerOf2(v) {
			return false // v must be power of 2
		} else if !(v/2 <= u && u < v) {
			return false // u must be in range [v/2, v[
		} else if isPowerOf2(u) && v != 2*u {
			return false // if u is power of 2, v = 2*u
		}

		return true
	}))
}

func TestPagingWriter(t *testing.T) {
	assert := newAssertions(t)

	recordPages := func(ids *idList, pages *[][]byte) func(PageID, []byte) reason {
		return func(id PageID, buf []byte) reason {
			*ids, *pages = append(*ids, id), append(*pages, buf)
			return nil
		}
	}

	assert.Run("requires pages", func(assert *assertions) {
		w := newPagingWriter(make(idList, 0), 4096, 0, nil)
		assert.Nil(w)
	})

	assert.Run("errors on write, if not enough pages", func(assert *assertions) {
		w := newPagingWriter(idList{2}, 64, 0, nil)
		var err error
		for i := 0; i < 60; i++ {
			var tmp [10]byte
			if err = w.Write(tmp[:]); err != nil {
				break
			}
		}
		assert.Error(err)
	})

	assert.Run("write 1 page", func(assert *assertions) {
		var ids idList
		var pages [][]byte
		w := newPagingWriter(idList{2}, 64, 0, recordPages(&ids, &pages))

		tmp := "abcdefghijklmn"
		assert.NoError(w.Write([]byte(tmp)))
		assert.NoError(w.Flush())

		assert.Equal(idList{2}, ids)
		if assert.Len(pages, 1) {
			assert.Len(pages[0], 64)

			hdr, payload := castListPage(pages[0])
			assert.Equal(PageID(0), hdr.next.Get())
			assert.Equal(1, int(hdr.count.Get()))
			assert.Equal(tmp, string(payload[:len(tmp)]))
		}
	})

	assert.Run("pages with multiple ids will be linked", func(assert *assertions) {
		var ids idList
		var pages [][]byte
		w := newPagingWriter(idList{0, 1, 2}, 64, 0, recordPages(&ids, &pages))

		assert.NoError(w.Flush()) // flush all pages
		assert.Equal(idList{0, 1, 2}, ids)
		if assert.Len(pages, 3) {
			expectedNext := idList{1, 2, 0}
			for i, page := range pages {
				assert.Len(page, 64)
				hdr, _ := castListPage(page)
				assert.Equal(expectedNext[i], hdr.next.Get())
				assert.Equal(0, int(hdr.count.Get()))
			}
		}
	})

	assert.Run("write 3 pages", func(assert *assertions) {
		var ids idList
		var pages [][]byte
		w := newPagingWriter(idList{0, 1, 2}, 64, 0, recordPages(&ids, &pages))

		for i := 0; i < 11; i++ {
			var buf [10]byte
			if !assert.NoError(w.Write(buf[:])) {
				assert.FailNow("all writes must be ok")
			}
		}
		assert.NoError(w.Flush())

		if assert.Len(pages, 3) {
			expectedCounts := []uint32{5, 5, 1}
			for i, page := range pages {
				hdr, _ := castListPage(page)
				assert.Equal(expectedCounts[i], hdr.count.Get())
			}
		}
	})
}
