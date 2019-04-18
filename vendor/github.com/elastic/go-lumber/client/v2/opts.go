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

package v2

import (
	"encoding/json"
	"errors"
	"time"
)

// Option type to be passed to New/Dial functions.
type Option func(*options) error

type options struct {
	timeout     time.Duration
	encoder     jsonEncoder
	compressLvl int
}

type jsonEncoder func(interface{}) ([]byte, error)

// JSONEncoder client option configuring the encoder used to convert events
// to json. The default is `json.Marshal`.
func JSONEncoder(encoder func(interface{}) ([]byte, error)) Option {
	return func(opt *options) error {
		opt.encoder = encoder
		return nil
	}
}

// Timeout client option configuring read/write timeout.
func Timeout(to time.Duration) Option {
	return func(opt *options) error {
		if to < 0 {
			return errors.New("timeouts must not be negative")
		}
		opt.timeout = to
		return nil
	}
}

// CompressionLevel client option setting the gzip compression level (0 to 9).
func CompressionLevel(l int) Option {
	return func(opt *options) error {
		if !(0 <= l && l <= 9) {
			return errors.New("compression level must be within 0 and 9")
		}
		opt.compressLvl = l
		return nil
	}
}

func applyOptions(opts []Option) (options, error) {
	o := options{
		encoder: json.Marshal,
		timeout: 30 * time.Second,
	}

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return o, err
		}
	}
	return o, nil
}
