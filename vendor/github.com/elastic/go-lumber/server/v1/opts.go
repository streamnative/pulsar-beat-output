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

package v1

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/elastic/go-lumber/lj"
)

// Option type for configuring server run options.
type Option func(*options) error

type options struct {
	timeout time.Duration
	tls     *tls.Config
	ch      chan *lj.Batch
}

// Timeout configures server network timeouts.
func Timeout(to time.Duration) Option {
	return func(opt *options) error {
		if to < 0 {
			return errors.New("timeouts must not be negative")
		}
		opt.timeout = to
		return nil
	}
}

// TLS enables and configures TLS support in lumberjack server.
// Protocol version 1 mandates TLS being enabled.
func TLS(tls *tls.Config) Option {
	return func(opt *options) error {
		opt.tls = tls
		return nil
	}
}

// Channel option is used to register custom channel received batches will be
// forwarded to.
func Channel(c chan *lj.Batch) Option {
	return func(opt *options) error {
		opt.ch = c
		return nil
	}
}

func applyOptions(opts []Option) (options, error) {
	o := options{
		timeout: 30 * time.Second,
		tls:     nil,
	}

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return o, err
		}
	}
	return o, nil
}
