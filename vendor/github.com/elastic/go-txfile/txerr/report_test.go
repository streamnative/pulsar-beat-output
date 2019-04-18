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

package txerr

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFmt(t *testing.T) {
	type testCase struct {
		err      error
		format   string
		expected string
	}

	cases := map[string]testCase{
		"single error with single line format": {
			err: &testErr{
				op: "pkg/op", kind: errors.New("error kind"),
				ctx: "var=1",
				msg: "ups",
			},
			format:   "%v",
			expected: "pkg/op: var=1: error kind: ups",
		},
		"single error with multiline format": {
			err: &testErr{
				op: "pkg/op", kind: errors.New("error kind"),
				ctx: "var=1",
				msg: "ups",
			},
			format:   "%+v",
			expected: "pkg/op: var=1: error kind: ups",
		},

		"nested error with single line format": {
			err: &testErr{
				op: "pkg/op", kind: errors.New("error kind"),
				ctx: "var=1",
				msg: "ups",
				cause: &testErr{
					op: "pkg/nested-op", msg: "root cause",
				},
			},
			format:   "%v",
			expected: "pkg/op: var=1: error kind: ups",
		},

		"nested error with multiline format": {
			err: &testErr{
				op: "pkg/op", kind: errors.New("error kind"),
				ctx: "var=1",
				msg: "ups",
				cause: &testErr{
					op: "pkg/nested-op", msg: "root cause",
				},
			},
			format:   "%+v",
			expected: "pkg/op: var=1: error kind: ups\n\tpkg/nested-op: root cause",
		},

		"nested multi-error with multiline format": {
			err: &testMultErr{
				op: "pkg/op", kind: errors.New("error kind"),
				ctx: "var=1",
				msg: "ups",
				causes: []error{
					&testErr{
						op: "pkg/nested-op1", msg: "cause1",
						cause: &testErr{op: "pkg/leaf-op", msg: "root cause1"},
					},
					&testErr{op: "pkg/nested-op2", msg: "cause2"},
				},
			},
			format: "%+v",
			expected: "pkg/op: var=1: error kind: ups\n" +
				"\tpkg/nested-op1: cause1\n" +
				"\t\tpkg/leaf-op: root cause1\n" +
				"\tpkg/nested-op2: cause2",
		},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, fmt.Sprintf(test.format, test.err))
		})
	}
}
