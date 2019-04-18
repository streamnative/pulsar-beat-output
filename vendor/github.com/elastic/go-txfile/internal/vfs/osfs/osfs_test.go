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

package osfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/elastic/go-txfile/internal/mint"
	"github.com/elastic/go-txfile/internal/vfs"
)

func TestOSFileSupport(testing *testing.T) {
	t := mint.New(testing)

	setupFile := func(t *mint.T, file string) (vfs.File, func()) {
		path, teardown := setupPath(t, file)

		f, err := Open(path, os.ModePerm)
		if err != nil {
			teardown()
			t.Fatal(err)
		}

		return f, func() {
			f.Close()
			teardown()
		}
	}

	t.Run("file size", func(t *mint.T) {
		file, teardown := setupFile(t, "")
		defer teardown()

		_, err := file.WriteAt([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0)
		t.FatalOnError(err)

		sz, err := file.Size()
		t.NoError(err)
		t.Equal(10, int(sz))
	})

	t.Run("lock and unlock succeed", func(t *mint.T) {
		for _, blocking := range [2]bool{false, true} {
			blocking := blocking
			t.Run(fmt.Sprintf("blocking: %v", blocking), func(t *mint.T) {
				file, teardown := setupFile(t, "")
				defer teardown()

				err := file.Lock(true, blocking)
				t.NoError(err)

				err = file.Unlock()
				t.NoError(err)
			})
		}
	})

	t.Run("lock suceeds twice", func(t *mint.T) {
		for _, blocking := range [2]bool{false, true} {
			blocking := blocking
			t.Run(fmt.Sprintf("blocking: %v", blocking), func(t *mint.T) {
				file, teardown := setupFile(t, "")
				defer teardown()

				err := file.Lock(true, blocking)
				t.NoError(err)

				err = file.Unlock()
				t.NoError(err)

				err = file.Lock(true, blocking)
				t.NoError(err)

				err = file.Unlock()
				t.NoError(err)
			})
		}
	})

	t.Run("locking locked file fails", func(t *mint.T) {
		for _, blocking := range [2]bool{false, true} {
			blocking := blocking
			t.Run(fmt.Sprintf("blocking: %v", blocking), func(t *mint.T) {
				f1, teardown := setupFile(t, "")
				defer teardown()

				f2, err := Open(f1.Name(), os.ModePerm)
				t.FatalOnError(err)
				defer f2.Close()

				err = f1.Lock(true, false)
				t.NoError(err)

				err = f2.Lock(true, false)
				t.Error(err)

				err = f1.Unlock()
				t.NoError(err)
			})
		}
	})

	t.Run("mmap file", func(t *mint.T) {
		f, teardown := setupFile(t, "")
		defer teardown()

		var buf = [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		n, err := f.WriteAt(buf[:], 0)
		t.Equal(len(buf), n)
		t.NoError(err)

		mem, err := f.MMap(len(buf))
		t.FatalOnError(err)
		defer func() {
			t.NoError(f.MUnmap(mem))
		}()

		t.Equal(buf[:], mem[:len(buf)])
	})
}

func setupPath(t *mint.T, file string) (string, func()) {
	dir, err := ioutil.TempDir("", "")
	t.FatalOnError(err)

	if file == "" {
		file = "test.dat"
	}
	return path.Join(dir, file), func() {
		os.RemoveAll(dir)
	}
}
