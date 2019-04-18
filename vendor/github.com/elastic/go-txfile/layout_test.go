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
	"unsafe"
)

func TestMetaPage(t *testing.T) {
	assert := newAssertions(t)

	assert.Run("fail to validate", func(assert *assertions) {
		// var buf metaBuf
		buf := make([]byte, unsafe.Sizeof(metaPage{}))
		hdr := castMetaPage(buf[:])
		assert.Error(hdr.Validate())
		hdr.magic.Set(magic)
		assert.Error(hdr.Validate())
		hdr.version.Set(version)
		assert.Error(hdr.Validate())
	})

	assert.Run("fail if checksum not set", func(assert *assertions) {
		// var buf metabuf
		buf := make([]byte, unsafe.Sizeof(metaPage{}))
		hdr := castMetaPage(buf[:])
		hdr.Init(0, 4096, 1<<30)
		assert.Error(hdr.Validate())
	})

	assert.Run("with checksum", func(assert *assertions) {
		// var buf metabuf
		buf := make([]byte, unsafe.Sizeof(metaPage{}))
		hdr := castMetaPage(buf[:])
		hdr.Init(0, 4096, 1<<30)
		hdr.Finalize()
		assert.NoError(hdr.Validate())
	})

	assert.Run("check if contents changed", func(assert *assertions) {
		// var buf metabuf
		buf := make([]byte, unsafe.Sizeof(metaPage{}))
		hdr := castMetaPage(buf[:])
		hdr.Init(0, 4096, 1<<30)
		hdr.Finalize()
		buf[4] = 0xff
		assert.Error(hdr.Validate())
	})
}
