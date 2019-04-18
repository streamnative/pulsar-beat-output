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

// Package v1 proviades common lumberjack protocol version 1 definitions.
package v1

// Version declares the protocol revision supported by this package.
const Version = 1

// Lumberjack protocol version 1 message types.
const (
	CodeVersion byte = '1'

	CodeWindowSize byte = 'W'
	CodeDataFrame  byte = 'D'
	CodeCompressed byte = 'C'
	CodeACK        byte = 'A'
)
