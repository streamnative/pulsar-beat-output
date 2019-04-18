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

package sftest

import structform "github.com/elastic/go-structform"

func Arr(l int, t structform.BaseType, elems ...interface{}) []Record {
	a := []Record{ArrayStartRec{l, t}}
	for _, elem := range elems {
		switch v := elem.(type) {
		case Record:
			a = append(a, v)
		case []Record:
			a = append(a, v...)
		case Recording:
			a = append(a, v...)
		default:
			panic("invalid key type")
		}
	}

	return append(a, ArrayFinishRec{})
}

func Obj(l int, t structform.BaseType, kv ...interface{}) []Record {
	if len(kv)%2 != 0 {
		panic("invalid object")
	}

	a := []Record{ObjectStartRec{l, t}}
	for i := 0; i < len(kv); i += 2 {
		k := kv[i].(string)
		a = append(a, ObjectKeyRec{k})

		switch v := kv[i+1].(type) {
		case Record:
			a = append(a, v)
		case []Record:
			a = append(a, v...)
		case Recording:
			a = append(a, v...)
		default:
			panic("invalid key type")
		}
	}

	return append(a, ObjectFinishRec{})
}
