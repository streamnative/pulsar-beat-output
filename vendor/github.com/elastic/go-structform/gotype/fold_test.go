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

package gotype

import (
	"bytes"
	gojson "encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	structform "github.com/elastic/go-structform"
	"github.com/elastic/go-structform/json"
	"github.com/elastic/go-structform/sftest"
)

type foldCase struct {
	json  string
	value interface{}
}

func TestIter2JsonConsistent(t *testing.T) {
	for name, test := range foldSamples() {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			iter, err := NewIterator(json.NewVisitor(&buf))
			if err != nil {
				panic(err)
			}

			err = iter.Fold(test.value)
			if err != nil {
				t.Fatal(err)
			}

			// compare conversions did preserve type
			assertJSON(t, test.json, buf.String())
		})
	}
}

func TestUserFold(t *testing.T) {
	ts := time.Now()
	tsStr := ts.String()
	tsInt := ts.Unix()

	foldTsString := func(t *time.Time, vs structform.ExtVisitor) error {
		return vs.OnString(t.String())
	}

	foldTsInt := func(t *time.Time, vs structform.ExtVisitor) error {
		return vs.OnInt64(t.Unix())
	}

	tests := []struct {
		in   interface{}
		fold interface{}
		want sftest.Recording
	}{
		{
			in:   ts,
			fold: foldTsString,
			want: sftest.Recording{sftest.StringRec{tsStr}},
		},
		{
			in:   ts,
			fold: foldTsInt,
			want: sftest.Recording{sftest.Int64Rec{tsInt}},
		},
		{
			in:   &ts,
			fold: foldTsString,
			want: sftest.Recording{sftest.StringRec{tsStr}},
		},
		{
			in:   &ts,
			fold: foldTsInt,
			want: sftest.Recording{sftest.Int64Rec{tsInt}},
		},
		{
			in:   map[string]interface{}{"ts": ts},
			fold: foldTsInt,
			want: sftest.Obj(1, structform.AnyType, "ts", sftest.Int64Rec{tsInt}),
		},
		{
			in:   map[string]interface{}{"ts": &ts},
			fold: foldTsInt,
			want: sftest.Obj(1, structform.AnyType, "ts", sftest.Int64Rec{tsInt}),
		},
		{
			in:   map[string]interface{}{"ts": ts},
			fold: foldTsString,
			want: sftest.Obj(1, structform.AnyType, "ts", sftest.StringRec{tsStr}),
		},
		{
			in:   map[string]interface{}{"ts": &ts},
			fold: foldTsString,
			want: sftest.Obj(1, structform.AnyType, "ts", sftest.StringRec{tsStr}),
		},
	}

	for i, test := range tests {
		title := fmt.Sprintf("%v - %#v -> %#v", i, test.in, test.want)
		test := test
		t.Run(title, func(t *testing.T) {
			t.Parallel()

			var rec sftest.Recording
			err := Fold(test.in, &rec, Folders(test.fold))
			if err != nil {
				t.Fatal(err)
			}

			rec.Assert(t, test.want)
		})
	}
}

func assertJSON(t *testing.T, expected, actual string) (err error) {
	expected, err = normalizeJSON(expected)
	if err != nil {
		t.Error(err)
		return
	}

	actual, err = normalizeJSON(actual)
	if err != nil {
		t.Error(err)
		return
	}

	// compare conversions did preserve type
	if !assert.Equal(t, expected, actual) {
		return errors.New("match failure")
	}
	return nil
}

func normalizeJSON(in string) (string, error) {
	var tmp interface{}
	if err := gojson.Unmarshal([]byte(in), &tmp); err != nil {
		return "", err
	}

	b, err := gojson.MarshalIndent(tmp, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func foldSamples() map[string]foldCase {
	samples := []foldCase{
		// primitives
		{`null`, nil},
		{`true`, true},
		{`false`, false},
		{`10`, int8(10)},
		{`10`, int32(10)},
		{`10`, int(10)},
		{`10`, uint(10)},
		{`10`, uint8(10)},
		{`10`, uint16(10)},
		{`10`, uint32(10)},
		{`12340`, uint16(12340)},
		{`1234567`, uint32(1234567)},
		{`12345678190`, uint64(12345678190)},
		{`-10`, int8(-10)},
		{`-10`, int32(-10)},
		{`-10`, int(-10)},
		{`3.14`, float32(3.14)},
		{`3.14`, float64(3.14)},
		{`"test"`, "test"},
		{`"test with \" being escaped"`, "test with \" being escaped"},

		// arrays
		{`[]`, []uint8{}},
		{`[]`, []string{}},
		{`[]`, []interface{}{}},
		{`[]`, []struct{ A string }{}},
		{`[[]]`, [][]uint8{{}}},
		{`[[]]`, [][]string{{}}},
		{`[[]]`, [][]interface{}{{}}},
		{`[[]]`, [][]struct{ A string }{{}}},
		{
			`[null,true,false,12345678910,3.14,"test"]`,
			[]interface{}{nil, true, false, uint64(12345678910), 3.14, "test"},
		},
		{`[1,2,3,4,5,6,7,8,9,10]`, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{`[1,2,3,4,5,6,7,8,9,10]`, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{`[1,2,3,4,5,6,7,8,9,10]`, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{`[1,2,3,4,5,6,7,8,9,10]`, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{`[1,2,3,4,5,6,7,8,9,10]`, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{`["testa","testb","testc"]`, []string{"testa", "testb", "testc"}},
		{`["testa","testb","testc"]`, []interface{}{"testa", "testb", "testc"}},

		// objects
		{`{}`, map[string]interface{}{}},
		{`{}`, map[string]int8{}},
		{`{}`, map[string]int64{}},
		{`{}`, map[string]struct{ A *int }{}},
		{`{}`, mapstr{}},
		{`{"a":null}`, map[string]interface{}{"a": nil}},
		{`{"a":null}`, mapstr{"a": nil}},
		{`{"a":null}`, struct{ A *int }{}},
		{`{"a":null}`, struct{ A *struct{ B int } }{}},
		{`{"a":true,"b":1,"c":"test"}`, map[string]interface{}{"a": true, "b": 1, "c": "test"}},
		{`{"a":true,"b":1,"c":"test"}`, mapstr{"a": true, "b": 1, "c": "test"}},
		{`{"a":true,"b":1,"c":"test"}`, struct {
			A bool
			B int
			C string
		}{true, 1, "test"}},

		{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
			map[string]interface{}{
				"field": []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
		{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
			map[string]interface{}{
				"field": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
		{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
			mapstr{
				"field": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
		{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
			map[string][]int{
				"field": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
		{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
			struct {
				Field []int
			}{
				Field: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},

		// structs with inlines
		{
			`{"a": 1}`,
			struct {
				X interface{} `struct:",inline"`
			}{X: map[string]int{"a": 1}},
		},
		{
			`{"a": 1}`,
			struct {
				X interface{} `struct:",inline"`
			}{X: struct{ A int }{1}},
		},
		{
			`{"a": 1}`,
			struct {
				X interface{} `struct:",inline"`
			}{X: &struct{ A int }{1}},
		},
		{
			`{"a": 1}`,
			struct {
				X map[string]interface{} `struct:",inline"`
			}{X: map[string]interface{}{"a": 1}},
		},
		{
			`{"a": 1}`,
			struct {
				X map[string]int `struct:",inline"`
			}{X: map[string]int{"a": 1}},
		},
		{
			`{"a": 1}`,
			struct {
				X struct{ A int } `struct:",inline"`
			}{X: struct{ A int }{1}},
		},

		// omit empty without values
		{
			`{"a": 1}`,
			struct {
				A int
				B interface{} `struct:",omitempty"`
			}{A: 1},
		},
		{
			`{"a": 1}`,
			struct {
				A int
				B map[string]interface{} `struct:",omitempty"`
			}{A: 1},
		},
		{
			`{"a": 1}`,
			struct {
				A int
				B []int `struct:",omitempty"`
			}{A: 1},
		},
		{
			`{"a": 1}`,
			struct {
				A int
				B string `struct:",omitempty"`
			}{A: 1},
		},
		{
			`{"a": 1}`,
			struct {
				A int
				B *int `struct:",omitempty"`
			}{A: 1},
		},
		{
			`{"a": 1}`,
			struct {
				A int
				B *struct{ C int } `struct:",omitempty"`
			}{A: 1},
		},

		// omit empty with values
		{
			`{"a": 1, "b": 2}`,
			struct {
				A int
				B interface{} `struct:",omitempty"`
			}{A: 1, B: 2},
		},
		{
			`{"a": 1, "b": {"c": 2}}`,
			struct {
				A int
				B map[string]interface{} `struct:",omitempty"`
			}{A: 1, B: map[string]interface{}{"c": 2}},
		},
		{
			`{"a": 1, "b":[2]}`,
			struct {
				A int
				B []int `struct:",omitempty"`
			}{A: 1, B: []int{2}},
		},
		{
			`{"a": 1, "b": "test"}`,
			struct {
				A int
				B string `struct:",omitempty"`
			}{A: 1, B: "test"},
		},
		{
			`{"a": 1, "b": 0}`,
			struct {
				A int
				B *int `struct:",omitempty"`
			}{A: 1, B: new(int)},
		},
		{
			`{"a": 1, "b": {"c": 2}}`,
			struct {
				A int
				B *struct{ C int } `struct:",omitempty"`
			}{A: 1, B: &struct{ C int }{2}},
		},

		// omit
		{
			`{"a": 1}`,
			struct {
				A int
				B int `struct:"-"`
			}{A: 1, B: 2},
		},
		{
			`{"a": 1}`,
			struct {
				A int
				B int `struct:",omit"`
			}{A: 1, B: 2},
		},
	}

	m := map[string]foldCase{}
	for i, test := range samples {
		title := fmt.Sprintf("%v - %v (%T)", i, test.json, test.value)
		m[title] = test
	}
	return m
}
