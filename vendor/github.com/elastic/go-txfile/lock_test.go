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

func TestLock(t *testing.T) {
	assert := newAssertions(t)

	makeTestLock := func() (*sharedLock, *reservedLock, *pendingLock, *exclusiveLock) {
		l := newLock()
		return l.Shared(), l.Reserved(), l.Pending(), l.Exclusive()
	}

	assert.Run("multiple shared locks", func(assert *assertions) {
		shared, _, _, _ := makeTestLock()
		shared.Lock()
		defer shared.Unlock()

		assert.True(shared.check(), "shared lock can not be acquired")
	})

	assert.Run("shared lock if reserved lock is set", func(assert *assertions) {
		shared, reserved, _, _ := makeTestLock()

		reserved.Lock()
		defer reserved.Unlock()

		assert.True(shared.check(), "shared lock can not be acquired")
	})

	assert.Run("can not acquire shared lock if pending lock is set", func(assert *assertions) {
		shared, reserved, pending, _ := makeTestLock()

		reserved.Lock()
		defer reserved.Unlock()

		pending.Lock()
		defer pending.Unlock()

		assert.False(shared.check(), "shared lock can be acquired")
	})

	assert.Run("shared lock can be acquired once pending is unlocked", func(assert *assertions) {
		shared, reserved, pending, _ := makeTestLock()

		reserved.Lock()
		defer reserved.Unlock()

		pending.Lock()
		pending.Unlock()

		assert.True(shared.check(), "shared lock can not be acquired")
	})

	assert.Run("reserved lock correctly unlocks", func(assert *assertions) {
		_, reserved, _, _ := makeTestLock()

		reserved.Lock()
		reserved.Unlock()

		// this will block/fail the tests if it blocks
		reserved.Lock()
		reserved.Unlock()
	})

	assert.Run("exclusive lock can only be acquired if no shared lock is taken", func(assert *assertions) {
		_, reserved, pending, exclusive := makeTestLock()

		reserved.Lock()
		defer reserved.Unlock()

		pending.Lock()
		defer pending.Unlock()

		assert.True(exclusive.check(), "exclusive lock can not be acquired")
	})

	assert.Run("exclusive lock can not be acquired if shared lock exists", func(assert *assertions) {
		shared, reserved, pending, exclusive := makeTestLock()

		reserved.Lock()
		defer reserved.Unlock()

		shared.Lock()
		defer shared.Unlock()

		pending.Lock()
		defer pending.Unlock()

		assert.False(exclusive.check(), "exclusive lock can be acquired")
	})

	assert.Run("exclusive lock can be acquired after shared is unlocked", func(assert *assertions) {
		shared, reserved, pending, exclusive := makeTestLock()

		reserved.Lock()
		defer reserved.Unlock()

		shared.Lock()

		pending.Lock()
		defer pending.Unlock()

		shared.Unlock()
		assert.True(exclusive.check(), "exclusive lock can not be acquired")
	})
}
