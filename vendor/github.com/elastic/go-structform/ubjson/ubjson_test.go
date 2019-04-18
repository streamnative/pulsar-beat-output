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

package ubjson

import (
	"bytes"
	"io"
	"testing"

	"github.com/elastic/go-structform"
	"github.com/elastic/go-structform/sftest"
)

func TestEncParseConsistent(t *testing.T) {
	testEncParseConsistent(t, Parse)
}

func TestEncDecoderConsistent(t *testing.T) {
	testEncParseConsistent(t, func(content []byte, to structform.Visitor) error {
		dec := NewBytesDecoder(content, to)
		err := dec.Next()
		if err == io.EOF {
			err = nil
		}
		return err
	})
}

func TestEncParseBytesConsistent(t *testing.T) {
	testEncParseConsistent(t, func(content []byte, to structform.Visitor) error {
		p := NewParser(to)
		for _, b := range content {
			err := p.feed([]byte{b})
			if err != nil {
				return err
			}
		}
		return p.finalize()
	})
}

func testEncParseConsistent(
	t *testing.T,
	parse func([]byte, structform.Visitor) error,
) {
	sftest.TestEncodeParseConsistent(t, sftest.Samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			buf := bytes.NewBuffer(nil)
			vs := NewVisitor(buf)

			return vs, func(to structform.Visitor) error {
				return parse(buf.Bytes(), to)
			}
		})
}
