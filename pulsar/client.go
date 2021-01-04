/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package pulsar

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/outputs/codec"
	"github.com/elastic/beats/libbeat/publisher"
)

type client struct {
	clientOptions   pulsar.ClientOptions
	producerOptions pulsar.ProducerOptions
	pulsarClient    pulsar.Client
	producer        pulsar.Producer
	observer        outputs.Observer
	beat            beat.Info
	config          *pulsarConfig
	codec           codec.Codec
}

func newPulsarClient(
	beat beat.Info,
	observer outputs.Observer,
	clientOptions pulsar.ClientOptions,
	producerOptions pulsar.ProducerOptions,
	config *pulsarConfig,
) (*client, error) {
	c := &client{
		clientOptions:   clientOptions,
		producerOptions: producerOptions,
		observer:        observer,
		beat:            beat,
		config:          config,
	}
	return c, nil
}

func (c *client) Connect() error {
	var err error
	c.pulsarClient, err = pulsar.NewClient(c.clientOptions)
	logp.Info("start create pulsar client")
	if err != nil {
		logp.Debug("pulsar", "Create pulsar client failed: %v", err)
		return err
	}
	logp.Info("start create pulsar producer")
	c.producer, err = c.pulsarClient.CreateProducer(c.producerOptions)
	if err != nil {
		logp.Debug("pulsar", "Create pulsar producer failed: %v", err)
		return err
	}
	logp.Info("start create encoder")
	c.codec, err = codec.CreateEncoder(c.beat, c.config.Codec)
	if err != nil {
		logp.Debug("pulsar", "Create encoder failed: %v", err)
		return err
	}

	return nil
}

func (c *client) Close() error {
	c.pulsarClient.Close()
	return nil
}

func (c *client) Publish(batch publisher.Batch) error {
	defer batch.ACK()
	events := batch.Events()
	c.observer.NewBatch(len(events))
	logp.Debug("Pulsar received events: %d", len(events))
	for i := range events {
		event := &events[i]
		serializedEvent, err := c.codec.Encode(c.beat.Beat, &event.Content)
		if err != nil {
			c.observer.Dropped(1)
			logp.Error("Failed event: %v, error: %v", event, err)
			continue
		}
		buf := make([]byte, len(serializedEvent))
		copy(buf, serializedEvent)
		logp.Debug("Pulsar success encode events: %d", i)
		pTime := time.Now()
		c.producer.SendAsync(ctx, &pulsar.ProducerMessage{
			EventTime: pTime,
			Key:       fmt.Sprintf("%d", pTime.Nanosecond()),
			Payload:   buf,
		}, func(msgId pulsar.MessageID, prodMsg *pulsar.ProducerMessage, err error) {
			if err != nil {
				c.observer.Dropped(1)
				logp.Error("produce send failed: %v", err)
			} else {
				logp.Debug("Pulsar success send events: messageID: %s ", msgId)
				c.observer.Acked(1)
			}
		})
	}
	logp.Debug("Pulsar success send events: %d", len(events))
	return nil
}

func (c *client) String() string {
	return "file(" + c.clientOptions.URL + ")"
}
