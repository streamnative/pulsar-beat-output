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

import "testing"

func TestPageSet(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("query nil pageSet", func(assert *assertions) {
		var s pageSet
		assert.False(s.Has(1))
		assert.True(s.Empty())
		assert.Equal(0, s.Count())
		assert.Nil(s.IDs())
		assert.Nil(s.Regions())
	})

	assert.Run("query empty pageSet", func(assert *assertions) {
		s := pageSet{}
		assert.False(s.Has(1))
		assert.True(s.Empty())
		assert.Equal(0, s.Count())
		assert.Nil(s.IDs())
		assert.Nil(s.Regions())
	})

	assert.Run("pageSet modifications", func(assert *assertions) {
		var s pageSet

		s.Add(1)
		s.Add(2)
		s.Add(10)
		assert.False(s.Empty())
		assert.Equal(3, s.Count())

		assert.True(s.Has(1))
		assert.True(s.Has(2))
		assert.False(s.Has(3))
		assert.False(s.Has(4))
		assert.True(s.Has(10))

		ids := s.IDs()
		ids.Sort() // ids might be unsorted
		assert.Equal(idList{1, 2, 10}, ids)

		regions := s.Regions()
		assert.Equal(regionList{{1, 2}, {10, 1}}, regions)
	})
}
