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
	"math/rand"
	"testing"

	"github.com/urso/qcgen"
)

func TestWAL(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("tx updates", func(assert *assertions) {
		assert.Run("add entries to empty wal", func(assert *assertions) {
			wal := makeWALog()
			tx := wal.makeTxWALState(0)
			tx.Set(1, 10)
			tx.Set(2, 100)

			upd := createMappingUpdate(wal.mapping, &tx)
			assert.Equal(walMapping{1: 10, 2: 100}, upd)
		})

		assert.Run("remove entries from empty wal", func(assert *assertions) {
			wal := makeWALog()
			tx := wal.makeTxWALState(0)
			tx.Release(10)
			assert.Equal(walMapping{}, createMappingUpdate(wal.mapping, &tx))
		})

		assert.Run("add to non empty wal", func(assert *assertions) {
			wal := makeWALog()
			wal.mapping = walMapping{10: 100}
			tx := wal.makeTxWALState(0)
			tx.Set(10, 1000)
			tx.Set(11, 101)
			upd := createMappingUpdate(wal.mapping, &tx)
			assert.Equal(walMapping{10: 1000, 11: 101}, upd, "new mapping mismatch")
			assert.Equal(walMapping{10: 100}, wal.mapping, "original mapping has been updated")
		})

		assert.Run("remove from non empty wal", func(assert *assertions) {
			wal := makeWALog()
			wal.mapping = walMapping{10: 100, 20: 2000}
			tx := wal.makeTxWALState(0)
			tx.Release(10)
			upd := createMappingUpdate(wal.mapping, &tx)
			assert.Equal(walMapping{20: 2000}, upd, "new mapping mismatch")
			assert.Equal(walMapping{10: 100, 20: 2000}, wal.mapping, "original mapping has been updated")
		})

		assert.Run("add and remove same mapping from non empty wal", func(assert *assertions) {
			wal := makeWALog()
			wal.mapping = walMapping{10: 100}
			tx := wal.makeTxWALState(0)
			tx.Set(10, 1000)
			tx.Set(11, 101)
			tx.Release(10)
			upd := createMappingUpdate(wal.mapping, &tx)
			assert.Equal(walMapping{11: 101}, upd, "new mapping mismatch")
			assert.Equal(walMapping{10: 100}, wal.mapping, "original mapping has been updated")
		})
	})

	assert.Run("serialization", func(assert *assertions) {
		assert.Run("empty mapping", func(assert *assertions) {
			store := newTestPageStore(64)
			mapping, ids, err := readWAL(store.Get, PageID(0))
			assert.NoError(err)
			assert.Nil(ids)
			assert.Equal(walMapping{}, mapping)
		})

		// uses fatal on error, so we don't have to see the complete input on failure
		assert.Run("randomized", makeQuickCheck(makeGenMapping(0, 1500), func(mapping walMapping) bool {
			pageSize := uint(64)
			store := newTestPageStore(pageSize)

			updateWAL := len(mapping) != 0

			root := PageID(0)
			pageCount := uint32(predictWALMappingPages(mapping, pageSize))
			var regions regionList

			tracef("mapping size=%v, page size=%v, page count=%v", len(mapping), pageSize, pageCount)

			if updateWAL {
				root = 3
				regions = regionList{{root, pageCount}}
				err := writeWAL(regions, pageSize, mapping, store.Set)
				assert.FatalOnError(err, "writing wal mapping failed")
			}

			newMapping, ids, err := readWAL(store.Get, root)
			assert.FatalOnError(err, "reading wal mapping failed")

			ok := assert.Equal(regions, ids.Regions(), "meta pages use doesn't match up")
			ok = ok && assert.Equal(mapping, newMapping, "mappings not equal after read")
			return ok
		}))
	})
}

func makeGenMapping(min, max uint64) func(*rand.Rand) walMapping {
	return func(rng *rand.Rand) walMapping {
		mapping := walMapping{}
		count := qcgen.GenUint64Range(rng, min, max)
		for i := count; i > 0; i-- {
			mapping[qcMakePageID(rng)] = qcMakePageID(rng)
		}
		return mapping
	}
}
