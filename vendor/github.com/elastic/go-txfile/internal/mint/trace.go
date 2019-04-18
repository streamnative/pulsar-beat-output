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

package mint

type TestLogTracer struct {
	backends []backend
}

type backend struct {
	print  func(...interface{})
	printf func(string, ...interface{})
}

func NewTestLogTracer(loggers ...interface{}) *TestLogTracer {
	type testLogger interface {
		Log(...interface{})
		Logf(string, ...interface{})
	}

	type tracer interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}

	bs := make([]backend, 0, len(loggers))
	for _, logger := range loggers {
		var to backend

		switch v := logger.(type) {
		case testLogger:
			to = backend{print: v.Log, printf: v.Logf}
		case tracer:
			to = backend{print: v.Println, printf: v.Printf}
		}

		if to.print != nil {
			bs = append(bs, to)
		}
	}

	return &TestLogTracer{bs}
}

func (t *TestLogTracer) Println(vs ...interface{}) {
	for _, b := range t.backends {
		b.print(vs...)
	}
}

func (t *TestLogTracer) Printf(fmt string, vs ...interface{}) {
	for _, b := range t.backends {
		b.printf(fmt, vs...)
	}
}
