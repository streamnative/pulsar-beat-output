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

import "net"

// SyncClient synchronously publishes events to lumberjack endpoint waiting for
// ACK before allowing another send request. The client is not thread-safe.
type SyncClient struct {
	cl *Client
}

// NewSyncClientWith creates a new SyncClient from low-level lumberjack v2 Client.
func NewSyncClientWith(c *Client) (*SyncClient, error) {
	return &SyncClient{c}, nil
}

// NewSyncClientWithConn creates a new SyncClient from an active connection.
func NewSyncClientWithConn(c net.Conn, opts ...Option) (*SyncClient, error) {
	cl, err := NewWithConn(c, opts...)
	if err != nil {
		return nil, err
	}
	return NewSyncClientWith(cl)
}

// SyncDial connects to lumberjack server and returns new SyncClient. On error
// no SyncClient is being created.
func SyncDial(address string, opts ...Option) (*SyncClient, error) {
	cl, err := Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	return NewSyncClientWith(cl)
}

// SyncDialWith uses provided dialer to connect to lumberjack server. On error
// no SyncClient is being returned.
func SyncDialWith(
	dial func(network, address string) (net.Conn, error),
	address string,
	opts ...Option,
) (*SyncClient, error) {
	cl, err := DialWith(dial, address, opts...)
	if err != nil {
		return nil, err
	}
	return NewSyncClientWith(cl)
}

// Close closes the client, so no new events can be published anymore. The
// underlying network connection will be closed too. Returns an error if
// underlying net.Conn errors on Close.
func (c *SyncClient) Close() error {
	return c.cl.Close()
}

// Send publishes a new batch of events by JSON-encoding given batch.
// Send blocks until the complete batch has been ACKed by lumberjack server or
// some error happened.
func (c *SyncClient) Send(data []interface{}) (int, error) {
	if err := c.cl.Send(data); err != nil {
		return 0, err
	}

	seq, err := c.cl.AwaitACK(uint32(len(data)))
	return int(seq), err
}
