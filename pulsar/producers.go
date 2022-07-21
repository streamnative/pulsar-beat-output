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
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/elastic/beats/v7/libbeat/logp"
)

type Producers struct {
	cache map[string]pulsar.Producer
	lock  sync.RWMutex
}

func NewProducers() *Producers {
	return &Producers{
		cache: make(map[string]pulsar.Producer),
	}
}

func (p *Producers) LoadProducer(topic string, client *client) (pulsar.Producer, error) {
	cacheProducer := p.getProducer(topic)
	if cacheProducer != nil {
		return cacheProducer, nil
	}

	return p.createProducer(topic, client)
}

func (p *Producers) getProducer(topic string) pulsar.Producer {
	p.lock.RLock()
	producer := p.cache[topic]
	p.lock.RUnlock()

	return producer
}

func (p *Producers) createProducer(topic string, client *client) (pulsar.Producer, error) {
	p.lock.Lock()

	logp.Info("creating pulsar producer for topic: %s", topic)
	newProducerOptions := pulsar.ProducerOptions{
		Topic:                           topic,
		Name:                            client.producerOptions.Name,
		Properties:                      client.producerOptions.Properties,
		SendTimeout:                     client.producerOptions.SendTimeout,
		DisableBlockIfQueueFull:         client.producerOptions.DisableBlockIfQueueFull,
		MaxPendingMessages:              client.producerOptions.MaxPendingMessages,
		HashingScheme:                   client.producerOptions.HashingScheme,
		CompressionType:                 client.producerOptions.CompressionType,
		CompressionLevel:                client.producerOptions.CompressionLevel,
		MessageRouter:                   client.producerOptions.MessageRouter,
		DisableBatching:                 client.producerOptions.DisableBatching,
		BatchingMaxPublishDelay:         client.producerOptions.BatchingMaxPublishDelay,
		BatchingMaxMessages:             client.producerOptions.BatchingMaxMessages,
		Interceptors:                    client.producerOptions.Interceptors,
		Schema:                          client.producerOptions.Schema,
		MaxReconnectToBroker:            client.producerOptions.MaxReconnectToBroker,
		BatcherBuilderType:              client.producerOptions.BatcherBuilderType,
		PartitionsAutoDiscoveryInterval: client.producerOptions.PartitionsAutoDiscoveryInterval,
		Encryption:                      client.producerOptions.Encryption,
	}
	producer, err := client.pulsarClient.CreateProducer(newProducerOptions)
	if err == nil {
		p.cache[topic] = producer
		logp.Info("created pulsar producer for topic: %s", topic)
	} else {
		logp.Err("creating pulsar producer{topic=%s} failed: %v", topic, err)
	}

	p.lock.Unlock()

	return producer, err
}

func (p *Producers) Close() {
	p.lock.Lock()
	for topic := range p.cache {
		producer := p.cache[topic]

		logp.Info("closing pulsar producer{topic=%s}", topic)
		err := producer.Flush()
		if err != nil {
			logp.Err("flush pulsar producer{topic=%s} failed: %v", topic, err)
		}

		producer.Close()
		logp.Info("closed pulsar producer{topic=%s}", topic)
	}
	p.lock.Unlock()
}
