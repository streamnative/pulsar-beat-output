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

package server

import (
	"errors"
	"net"
)

type muxListener struct {
	net.Listener
	ch chan net.Conn
}

type muxConn struct {
	net.Conn
}

type versionConn struct {
	net.Conn
	parent *muxConn
	v      byte
}

var (
	// ErrListenerClosed indicates the multiplexing network listener being closed.
	ErrListenerClosed = errors.New("listener closed")
)

func newMuxListener(l net.Listener) *muxListener {
	return &muxListener{l, make(chan net.Conn, 1)}
}

// Accept waits for and returns the next connection to the listener.
func (l *muxListener) Accept() (net.Conn, error) {
	conn, ok := <-l.ch
	if !ok {
		return nil, ErrListenerClosed
	}
	return conn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (l *muxListener) Close() error {
	close(l.ch)
	return nil
}

func newMuxConn(v byte, c net.Conn) *muxConn {
	mc := &muxConn{}
	vc := &versionConn{c, mc, v}
	mc.Conn = vc
	return mc
}

func (vc *versionConn) Read(buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}

	buf[0] = vc.v
	vc.parent.Conn = vc.Conn
	n, err := vc.Conn.Read(buf[1:])
	return n + 1, err
}
