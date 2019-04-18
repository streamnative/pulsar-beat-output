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

// Test client pushing precomputed batch of events.
//
// For list of known command line flags run:
//
//  tst-send -h
package main

import (
	"expvar"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-lumber/client/v2"
)

func main() {
	connect := flag.String("c", "localhost:5044", "Remote address")
	compress := flag.Int("compress", 3, "Compression level (0-9)")
	timeout := flag.Duration("timeout", 30*time.Second, "Connection timeouts")
	batchSize := flag.Int("batch", 2048, "Batch size")
	pipelined := flag.Int("pipeline", 0, "enabled pipeline mode with number of batches kept in pipeline")
	httpprof := flag.String("httpprof", ":6060", "HTTP profiling server address")
	flag.Parse()

	stat := expvar.NewInt("ACKed")

	batch := make([]interface{}, *batchSize)
	for i := range batch {
		batch[i] = makeEvent()
	}
	L := int64(len(batch))

	go func() {
		log.Printf("Listening: %v\n", *httpprof)
		http.ListenAndServe(*httpprof, nil)
	}()

	log.Printf("connect to: %v", *connect)
	if *pipelined == 0 {
		cl, err := v2.SyncDial(*connect,
			v2.CompressionLevel(*compress),
			v2.Timeout(*timeout))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		for {
			_, err := cl.Send(batch)
			if err != nil {
				log.Println(err)
				return
			}

			stat.Add(L)
		}
	} else {
		cl, err := v2.AsyncDial(*connect,
			*pipelined,
			v2.CompressionLevel(*compress),
			v2.Timeout(*timeout))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		for {
			cb := func(_ uint32, err error) {
				if err != nil {
					log.Println(err)
					return
				}

				stat.Add(L)
			}

			err := cl.Send(cb, batch)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

var (
	text = strings.Split(`Lorem ipsum dolor sit amet, consetetur sadipscing elitr,
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat,
sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet
clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod
tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At
vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,
no sea takimata sanctus est Lorem ipsum dolor sit amet.`, "\n")
)

func makeEvent() interface{} {
	line := text[rand.Intn(len(text))]
	return map[string]interface{}{
		"@timestamp": time.Now(),
		"type":       "filebeat",
		"message":    line,
		"offset":     1000,
	}
}
