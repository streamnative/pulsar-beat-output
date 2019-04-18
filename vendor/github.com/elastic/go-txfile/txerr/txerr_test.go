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

type testErr struct {
	op    string
	ctx   string
	kind  error
	msg   string
	cause error
}

type testMultErr struct {
	op     string
	ctx    string
	kind   error
	msg    string
	causes []error
}

func (e *testErr) Op() string                 { return e.op }
func (e *testErr) Kind() error                { return e.kind }
func (e *testErr) Context() string            { return e.ctx }
func (e *testErr) Message() string            { return e.msg }
func (e *testErr) Cause() error               { return e.cause }
func (e *testErr) Error() string              { return Report(e, true) }
func (e *testErr) Format(s fmt.State, c rune) { Format(e, s, c) }

func (e *testMultErr) Op() string                 { return e.op }
func (e *testMultErr) Kind() error                { return e.kind }
func (e *testMultErr) Context() string            { return e.ctx }
func (e *testMultErr) Message() string            { return e.msg }
func (e *testMultErr) Causes() []error            { return e.causes }
func (e *testMultErr) Error() string              { return Report(e, true) }
func (e *testMultErr) Format(s fmt.State, c rune) { Format(e, s, c) }

func TestErrQueries(t *testing.T) {
	var (
		kind1 = errors.New("kind1")
		kind2 = errors.New("kind2")
		kind3 = errors.New("kind3")
		kind4 = errors.New("kind4")

		err1 = &testErr{op: "err1", kind: kind1}
		err2 = &testErr{op: "err2", kind: kind2}
		err3 = &testErr{op: "err3", cause: err1}
		err4 = &testErr{op: "err4", kind: kind3, cause: err3}
		err5 = &testErr{op: "err5"}
		err6 = &testErr{kind: kind2}
		err7 = &testErr{kind: kind2, cause: err3}

		merr = &testMultErr{kind: kind4, causes: []error{
			err4,
			err2,
		}}
	)

	t.Run("find kind", func(t *testing.T) {
		type testCase struct {
			in, kind, expected error
		}

		cases := map[string]testCase{
			"kind1 in nested": {
				in:       err4,
				kind:     kind1,
				expected: err1,
			},
			"kind1 in nested multierr": {
				in:       merr,
				kind:     kind1,
				expected: err1,
			},
			"kind2 not in nested": {
				in:       err3,
				kind:     kind2,
				expected: nil,
			},
			"kind2 in nested multerr": {
				in:       merr,
				kind:     kind2,
				expected: err2,
			},
		}

		for name, test := range cases {
			test := test
			t.Run(name, func(t *testing.T) {
				found := FindKind(test.in, test.kind)
				assert.Equal(t, test.expected, found)

				is := Is(test.kind, test.in)
				assert.Equal(t, found != nil, is)
			})
		}
	})

	t.Run("get kind", func(t *testing.T) {
		type testCase struct {
			in       error
			expected error
		}

		cases := map[string]testCase{
			"error without kind":   {in: err5, expected: nil},
			"top error with kind":  {in: err4, expected: kind3},
			"kind in nested error": {in: err3, expected: kind1},
		}

		for name, test := range cases {
			test := test
			t.Run(name, func(t *testing.T) {
				assert.Equal(t, test.expected, GetKind(test.in))
			})
		}
	})

	t.Run("find op", func(t *testing.T) {
		type testCase struct {
			in, expected error
			op           string
		}

		cases := map[string]testCase{
			"op err1 in nested": {
				in:       err4,
				op:       "err1",
				expected: err1,
			},
			"op err1 in nested multierr": {
				in:       merr,
				op:       "err1",
				expected: err1,
			},
			"op err2 not in nested": {
				in:       err3,
				op:       "err2",
				expected: nil,
			},
			"op err2 in nested multerr": {
				in:       merr,
				op:       "err2",
				expected: err2,
			},
		}

		for name, test := range cases {
			test := test
			t.Run(name, func(t *testing.T) {
				found := FindOp(test.in, test.op)
				assert.Equal(t, test.expected, found)

				is := IsOp(test.op, test.in)
				assert.Equal(t, found != nil, is)
			})
		}
	})

	t.Run("get op", func(t *testing.T) {
		type testCase struct {
			in       error
			expected string
		}

		cases := map[string]testCase{
			"error without op":   {in: err6, expected: ""},
			"top error with op":  {in: err4, expected: "err4"},
			"op in nested error": {in: err7, expected: "err3"},
		}

		for name, test := range cases {
			test := test
			t.Run(name, func(t *testing.T) {
				assert.Equal(t, test.expected, GetOp(test.in))
			})
		}
	})
}
