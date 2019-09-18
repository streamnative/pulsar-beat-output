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
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/go-structform/json"
)

type unfoldCase struct {
	json  string
	input interface{}
	value interface{}
}

type stateStringUnfolder struct {
	BaseUnfoldState
	fn func(string) error
}

type intFromString int

func (i *intFromString) Expand() UnfoldState {
	return &stateStringUnfolder{
		fn: func(str string) error {
			tmp, err := strconv.Atoi(str)
			*i = intFromString(tmp)
			return err
		},
	}
}

func (u *stateStringUnfolder) OnString(ctx UnfoldCtx, in string) error {
	defer ctx.Done()
	return u.fn(in)
}

func TestFoldUnfoldConsistent(t *testing.T) {
	for name, test := range unfoldSamples() {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			u, err := NewUnfolder(test.value)
			if err != nil {
				t.Fatalf("NewUnfolder failed with: %v", err)
			}

			gcCheck := newGCCheckVisitor(u)
			if err := Fold(test.input, gcCheck); err != nil {
				t.Fatalf("Fold-Unfold failed with: %v", err)
			}

			if st := &u.unfolder; len(st.stack) > 0 {
				t.Fatalf("Unfolder state stack not empty: %v, %v", st.stack, st.current)
			}

			// serialize to json
			var buf bytes.Buffer
			if err := Fold(test.value, json.NewVisitor(&buf)); err != nil {
				t.Fatalf("serialize to json failed with: %v", err)
			}

			// compare conversions did preserve type
			assertJSON(t, test.json, buf.String())
		})
	}
}

func TestUnfoldJsonInto(t *testing.T) {
	for name, test := range unfoldSamples() {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			un, err := NewUnfolder(test.value)
			if err != nil {
				t.Fatal(err)
			}

			dec := json.NewParser(un)
			input := test.json

			err = dec.ParseString(input)
			if err != nil {
				t.Fatal(err)
			}

			// check state valid by processing a second time
			if err = un.SetTarget(test.value); err != nil {
				t.Fatal(err)
			}

			err = dec.ParseString(input)
			if err != nil {
				t.Fatal(err)
			}

			// clear unfolder state.
			un.Reset()
		})
	}
}

func BenchmarkUnfoldJsonInto(b *testing.B) {
	for name, test := range unfoldSamples() {
		un, err := NewUnfolder(test.value)
		if err != nil {
			b.Fatal(err)
		}

		dec := json.NewParser(un)
		input := test.json
		// parse once to reset state for use of 'setTarget' in benchmark
		if err := dec.ParseString(input); err != nil {
			b.Fatal(err)
		}

		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				un.SetTarget(test.value)
				err := dec.ParseString(input)
				if err != nil {
					b.Error(err)
				}
			}
		})
	}
}

func unfoldSamples() map[string]unfoldCase {
	samples := []unfoldCase{
		// primitives
		{`null`, nil, new(interface{})},
		{`""`, nil, new(string)},
		{`true`, true, new(bool)},
		{`true`, true, new(interface{})},
		{`false`, false, new(bool)},
		{`false`, false, new(interface{})},
		{`null`, nil, new(*int)},
		{`10`, int8(10), new(int8)},
		{`10`, int8(10), new(int)},
		{`10`, int8(10), new(*int8)},
		{`10`, int8(10), new(*int)},
		{`10`, int8(10), new(interface{})},
		{`10`, int32(10), new(int64)},
		{`10`, int32(10), new(int16)},
		{`10`, int32(10), new(interface{})},
		{`10`, int(10), new(int)},
		{`10`, int(10), new(uint)},
		{`10`, uint(10), new(uint)},
		{`10`, uint(10), new(uint16)},
		{`10`, uint(10), new(interface{})},
		{`10`, uint8(10), new(int)},
		{`10`, uint16(10), new(uint64)},
		{`10`, uint32(10), new(uint8)},
		{`12340`, uint16(12340), new(uint16)},
		{`12340`, uint16(12340), new(interface{})},
		{`1234567`, uint32(1234567), new(uint32)},
		{`1234567`, uint32(1234567), new(int64)},
		{`12345678190`, uint64(12345678190), new(uint64)},
		{`12345678190`, uint64(12345678190), new(int64)},
		{`-10`, int8(-10), new(int)},
		{`-10`, int8(-10), new(int8)},
		{`-10`, int8(-10), new(int32)},
		{`-10`, int8(-10), new(int64)},
		{`-10`, int8(-10), new(interface{})},
		{`3.14`, float32(3.14), new(float32)},
		{`3.14`, float32(3.14), new(interface{})},
		{`3.14`, float64(3.14), new(float32)},
		{`3.14`, float64(3.14), new(float64)},
		{`3.14`, float64(3.14), new(interface{})},
		{`"test"`, "test", new(string)},
		{`"test"`, "test", new(interface{})},
		{`"test with \" being escaped"`, "test with \" being escaped", new(string)},
		{`"test with \" being escaped"`, "test with \" being escaped", new(interface{})},

		// arrays
		{`[]`, []uint8{}, new(interface{})},
		{`[]`, []uint8{}, new([]uint8)},
		{`[]`, []interface{}{}, new([]interface{})},
		{`[]`, []interface{}{}, &[]struct{ A string }{}},
		{`[1,2,3]`, []uint8{1, 2, 3}, new(interface{})},
		{`[1,2,3]`, []uint8{1, 2, 3}, &[]uint8{0, 0, 0}},
		{`[1,2,3]`, []uint8{1, 2, 3}, &[]uint8{0, 0, 0, 4}},
		{`[1,2,3]`, []uint8{1, 2, 3}, &[]uint8{}},
		{`[1,2,3]`, []interface{}{1, 2, 3}, &[]uint8{}},
		{`[1,2,3]`, []interface{}{1, 2, 3}, new([]uint8)},
		{`[1,2,3]`, []int{1, 2, 3}, &[]interface{}{}},
		{`[1,2,3]`, []int{1, 2, 3}, &[]interface{}{nil, nil, nil, 4}},

		{`[1]`, []uint8{1}, &[]uint8{0, 2, 3}},
		{`[1,2]`, []uint8{1, 2}, &[]uint8{0, 0, 3}},
		{`[1,2]`, []interface{}{1, 2}, &[]uint{0, 0, 3}},
		{`[1,2]`, []interface{}{1, 2}, &[]interface{}{0, 0, 3}},
		{`["a","b"]`, []string{"a", "b"}, &[]interface{}{}},
		{`["a","b"]`, []string{"a", "b"}, &[]interface{}{nil, nil}},
		{`["a","b"]`, []string{"a", "b"}, &[]string{}},
		{`["a","b"]`, []string{"a", "b"}, new([]string)},
		{`[null,true,false]`,
			[]interface{}{nil, true, false},
			new(interface{})},
		{`[null,true]`,
			[]interface{}{nil, true},
			new(interface{})},
		{`[null,true,false,123,3.14,"test"]`,
			[]interface{}{nil, true, false, 123, 3.14, "test"},
			new(interface{})},
		{`[null,true,false,123,3.14,"test"]`,
			[]interface{}{nil, true, false, 123, 3.14, "test"},
			&[]interface{}{}},

		{`[1,2,3]`, []int{1, 2, 3}, &[]*int{}},
		{`[1,null,3]`, []interface{}{1, nil, 3}, &[]*int{}},

		// nested arrays
		{`[[]]`, []interface{}{[]uint{}}, new(interface{})},
		{`[]`, []interface{}{}, &[]interface{}{[]interface{}{1}}},
		{`[]`, []interface{}{}, &[][]int{}},
		{`[]`, []interface{}{}, &[][]int{{1}}},
		{`[[1]]`, []interface{}{[]interface{}{1}}, &[][]int{}},
		{`[[1,2,3],[4,5,6]]`,
			[]interface{}{[]interface{}{1, 2, 3}, []interface{}{4, 5, 6}},
			new(interface{})},
		{`[[1],[4]]`,
			[]interface{}{[]interface{}{1}, []interface{}{4}},
			&[]interface{}{[]interface{}{0, 2, 3}}},
		{`[[1],[4],[6]]`,
			[]interface{}{[]interface{}{1}, []interface{}{4}, []interface{}{6}},
			&[]interface{}{
				[]interface{}{0, 2, 3},
				[]interface{}{0, 5},
			}},
		{`[[1],[4],[6]]`,
			[]interface{}{[]interface{}{1}, []interface{}{4}, []interface{}{6}},
			&[][]int{
				{0, 2, 3},
				{0, 5},
			},
		},

		// maps
		{`{}`, map[string]interface{}{}, new(interface{})},
		{`{}`, map[string]interface{}{}, &map[string]interface{}{}},
		{`{"a":1}`, map[string]int{"a": 1}, new(interface{})},
		{`{"a":1}`, map[string]int{"a": 1}, &map[string]interface{}{}},
		{`{"a":1}`, map[string]int{"a": 1}, &map[string]interface{}{"a": 2}},
		{`{"a":1}`, map[string]int{"a": 1}, &map[string]int{}},
		{`{"a":1}`, map[string]int{"a": 1}, &map[string]int{"a": 2}},
		{`{"a":1}`, struct{ A int }{1}, &map[string]interface{}{}},
		{`{"a":1}`, struct{ A int }{1}, &map[string]int{}},
		{`{"a":1,"b":"b","c":true}`,
			map[string]interface{}{"a": 1, "b": "b", "c": true},
			new(interface{}),
		},

		// nested maps
		{`{"a":{}}`, map[string]interface{}{"a": map[string]interface{}{}}, new(interface{})},
		{`{"a":{}}`,
			map[string]interface{}{"a": map[string]interface{}{}},
			&map[string]interface{}{}},
		{`{"a":{}}`,
			map[string]interface{}{"a": map[string]interface{}{}},
			&map[string]interface{}{"a": nil}},
		{`{"a":{}}`,
			map[string]interface{}{"a": map[string]interface{}{}},
			&map[string]map[string]string{"a": nil}},
		{`{"0":{"a":1,"b":2,"c":3},"1":{"e":5,"f":6}}`,
			map[string]map[string]int{
				"0": {"a": 1, "b": 2, "c": 3},
				"1": {"e": 5, "f": 6},
			},
			new(interface{})},
		{`{"0":{"a":1,"b":2,"c":3},"1":{"e":5,"f":6}}`,
			map[string]map[string]int{
				"0": {"a": 1, "b": 2, "c": 3},
				"1": {"e": 5, "f": 6},
			},
			&map[string]interface{}{}},
		{`{"0":{"a":1,"b":2,"c":3},"1":{"e":5,"f":6}}`,
			map[string]map[string]int{
				"0": {"a": 1, "b": 2, "c": 3},
				"1": {"e": 5, "f": 6},
			},
			&map[string]map[string]int64{}},
		{`{"0":{"a":1}}`,
			map[string]map[string]int{
				"0": {"a": 1},
			},
			&map[string]interface{}{
				"0": map[string]int{"b": 2, "c": 3},
			}},
		{`{"0":{"a":1},"1":{"e":5}}`,
			map[string]map[string]int{
				"0": {"a": 1},
				"1": {"e": 5},
			},
			&map[string]interface{}{
				"0": map[string]int{"b": 2, "c": 3},
				"1": map[string]interface{}{"f": 6, "e": 0},
			}},
		{`{"0":{"a":1},"1":{"e":5}}`,
			map[string]map[string]int{
				"0": {"a": 1},
				"1": {"e": 5},
			},
			&map[string]map[string]uint{
				"0": {"b": 2, "c": 3},
				"1": {"f": 6, "e": 0},
			}},

		// map in array
		{`[{}]`, []interface{}{map[string]interface{}{}}, new(interface{})},
		{`[{}]`, []interface{}{map[string]interface{}{}}, &[]map[string]int{}},
		{`[{"a":1}]`, []map[string]int{{"a": 1}}, new(interface{})},
		{`[{"a":1}]`, []map[string]int{{"a": 1}}, &[]interface{}{}},
		{`[{"a":1}]`, []map[string]int{{"a": 1}}, &[]map[string]interface{}{}},
		{`[{"a":1}]`, []map[string]int{{"a": 1}}, &[]map[string]interface{}{{"a": 2}}},
		{`[{"a":1,"b":2}]`, []map[string]int{{"a": 1}}, &[]map[string]int{{"b": 2}}},
		{`[{"a":1},{"b":2}]`, []map[string]int{{"a": 1}, {"b": 2}}, &[]map[string]int{}},
		{`[{"a":1},{"b":"b"},{"c":true}]`,
			[]map[string]interface{}{{"a": 1}, {"b": "b"}, {"c": true}},
			new(interface{})},

		// array in map
		{`{"a": []}`, map[string]interface{}{"a": []int{}}, new(interface{})},
		{`{"a":[1,2],"b":[3]}`, map[string][]int{"a": {1, 2}, "b": {3}}, new(interface{})},
		{`{"a":[1,2],"b":[3]}`, map[string][]int{"a": {1, 2}, "b": {3}},
			&map[string][]int{}},
		{`{"a":[1,2],"b":[3,4,5]}`, map[string][]int{"a": {1, 2}, "b": {3, 4, 5}},
			&map[string][]int{"a": {0, 2}, "b": {0}}},

		// struct
		{`{"a":1}`, map[string]int{"a": 1}, &struct{ A int }{}},
		{`{"a":1}`, map[string]int{"a": 1}, &struct{ A *int }{}},
		{`{"a": 1, "c": 2}`, map[string]int{"a": 1}, &struct{ A, b, C int }{b: 1, C: 2}},
		{`{"a": {"c": 2}, "b": 1}`,
			map[string]interface{}{"a": map[string]int{"c": 2}},
			&struct {
				A struct{ C int }
				B int
			}{B: 1},
		},
		{`{"a": 1}`,
			map[string]interface{}{"a": 1},
			&struct {
				S struct {
					A int
				} `struct:",inline"`
			}{},
		},
		{`{"a":{"b":{"c":1}}}`,
			map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]int{
						"c": 1,
					},
				},
			},
			&struct{ A struct{ B struct{ C int } } }{},
		},
		{
			`{"a": 1}`,
			map[string]interface{}{"a": 1, "b": 2},
			&struct {
				A int
				B int `struct:"-"`
			}{},
		},
		{
			`{"a": 1}`,
			map[string]interface{}{"a": 1, "b": 2},
			&struct {
				A int
				B int `struct:",omit"`
			}{},
		},
	}

	m := map[string]unfoldCase{}
	for i, test := range samples {
		title := fmt.Sprintf("%v: %v (%T -> %T)", i, test.json, test.input, test.value)
		m[title] = test
	}
	return m
}

func TestUserUnfold(t *testing.T) {
	type myint int
	type addInt struct{ A, B int }

	type complexUnfolder func(*myint, interface{}) error
	type initUnfolder func(*myint) (interface{}, complexUnfolder)

	unfoldFromString := func(to *myint, in string) error {
		i, err := strconv.Atoi(in)
		*to = myint(i)
		return err
	}

	unfoldFromBool := func(to *myint, in bool) error {
		*to = 0
		if in {
			*to = 1
		}
		return nil
	}

	makeCellUnfolder := func(cell interface{}, proc func(*myint)) initUnfolder {
		return func(_ *myint) (interface{}, complexUnfolder) {
			return cell, func(to *myint, _ interface{}) error {
				proc(to)
				return nil
			}
		}
	}

	makeUnfoldWithIntCell := func(extra int) initUnfolder {
		cell := new(int)
		return makeCellUnfolder(cell, func(to *myint) {
			*to = myint(*cell + extra)
		})
	}

	makeUnfoldReuseVar := func(extra int) initUnfolder {
		return func(to *myint) (interface{}, complexUnfolder) {
			return to, func(to *myint, _ interface{}) error {
				*to += myint(extra)
				return nil
			}
		}
	}

	makeUnfoldStructAdd := func() initUnfolder {
		cell := &addInt{} // prealloc cell for reuse
		return makeCellUnfolder(cell, func(to *myint) {
			*to = myint(cell.A + cell.B)
		})
	}

	unfoldWithState := func(to *myint) UnfoldState {
		return &stateStringUnfolder{
			fn: func(in string) error {
				i, err := strconv.Atoi(in)
				*to = myint(i)
				return err
			},
		}
	}

	tests := map[string]struct {
		input    interface{}
		want     interface{}
		unfolder interface{}
	}{
		"parse from string": {
			input:    "12345",
			want:     myint(12345),
			unfolder: unfoldFromString,
		},
		"parse from bool": {
			input:    true,
			want:     myint(1),
			unfolder: unfoldFromBool,
		},
		"custom post processing with temporary cell": {
			input:    3,
			want:     myint(23),
			unfolder: makeUnfoldWithIntCell(20),
		},
		"reuse cell and post process": {
			input:    13,
			want:     myint(23),
			unfolder: makeUnfoldReuseVar(10),
		},
		"value from custom structure": {
			input:    addInt{1, 2},
			want:     myint(3),
			unfolder: makeUnfoldStructAdd(),
		},
		"expander value": {
			input: "42",
			want:  intFromString(42),
		},
		"parse with UnfoldState": {
			input:    "1234",
			want:     myint(1234),
			unfolder: unfoldWithState,
		},
		"parse array values from strings": {
			input:    []string{"1", "2", "3"},
			want:     []myint{1, 2, 3},
			unfolder: unfoldFromString,
		},
		"array with temporary cell": {
			input:    []int{1, 2, 3},
			want:     []myint{21, 22, 23},
			unfolder: makeUnfoldWithIntCell(20),
		},
		"array post processing with cell reuse": {
			input:    []int{1, 2, 3},
			want:     []myint{11, 12, 13},
			unfolder: makeUnfoldReuseVar(10),
		},
		"array with values from custom structure": {
			input:    []addInt{{1, 2}, {20, 3}},
			want:     []myint{3, 23},
			unfolder: makeUnfoldStructAdd(),
		},
		"array with UnfoldState": {
			input:    []string{"1", "2", "3"},
			want:     []myint{1, 2, 3},
			unfolder: unfoldWithState,
		},
		"array of expanders": {
			input: []string{"1", "2", "3"},
			want:  []intFromString{1, 2, 3},
		},
		"parse map values from strings": {
			input:    map[string]string{"a": "1", "b": "2", "c": "3"},
			want:     map[string]myint{"a": 1, "b": 2, "c": 3},
			unfolder: unfoldFromString,
		},
		"map with temporary cell": {
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			want:     map[string]myint{"a": 21, "b": 22, "c": 23},
			unfolder: makeUnfoldWithIntCell(20),
		},
		"map post processing with cell reuse": {
			input:    map[string]int{"a": 1, "b": 2, "c": 3},
			want:     map[string]myint{"a": 21, "b": 22, "c": 23},
			unfolder: makeUnfoldReuseVar(20),
		},
		"map with values from custom structure": {
			input:    map[string]addInt{"a": {1, 2}, "b": {20, 3}},
			want:     map[string]myint{"a": 3, "b": 23},
			unfolder: makeUnfoldStructAdd(),
		},
		"map with UnfoldState": {
			input:    map[string]string{"a": "1", "b": "2", "c": "3"},
			want:     map[string]myint{"a": 1, "b": 2, "c": 3},
			unfolder: unfoldWithState,
		},
		"map of expanders": {
			input: map[string]string{"a": "1", "b": "2", "c": "3"},
			want:  map[string]intFromString{"a": 1, "b": 2, "c": 3},
		},
		"struct field from string": {
			input:    map[string]string{"a": "1"},
			want:     struct{ A myint }{1},
			unfolder: unfoldFromString,
		},
		"struct field with temporary cell": {
			input:    map[string]int{"a": 1},
			want:     struct{ A myint }{11},
			unfolder: makeUnfoldWithIntCell(10),
		},
		"struct field reuse and post process": {
			input:    map[string]int{"a": 1},
			want:     struct{ A myint }{11},
			unfolder: makeUnfoldReuseVar(10),
		},
		"struct field with custom structure": {
			input:    map[string]addInt{"a": {1, 2}},
			want:     struct{ A myint }{3},
			unfolder: makeUnfoldStructAdd(),
		},
		"struct with UnfoldState": {
			input:    map[string]string{"a": "1"},
			want:     struct{ A myint }{1},
			unfolder: unfoldWithState,
		},
		"struct with expander": {
			input: map[string]string{"a": "1"},
			want:  struct{ A intFromString }{1},
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {
			var cell reflect.Value
			wantType := reflect.TypeOf(test.want)

			if wantType.Kind() == reflect.Map {
				tmp := reflect.MakeMap(wantType)
				cell = reflect.New(wantType)
				cell.Elem().Set(tmp)
			} else {
				cell = reflect.New(wantType)
			}

			u, err := NewUnfolder(cell.Interface(), Unfolders(test.unfolder))
			if err != nil {
				t.Fatalf("NewUnfolder failed with: %v", err)
			}

			gcCheck := newGCCheckVisitor(u)
			if err := Fold(test.input, gcCheck); err != nil {
				t.Fatalf("Fold-Unfold failed with: %v", err)
			}

			if st := &u.unfolder; len(st.stack) > 0 {
				t.Fatalf("Unfolder state stack not empty: %v, %v", st.stack, st.current)
			}

			assert.Equal(t, test.want, cell.Elem().Interface())
		})
	}
}
