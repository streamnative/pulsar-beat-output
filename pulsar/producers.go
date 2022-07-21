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
	lru "github.com/hashicorp/golang-lru/simplelru"
)

type Producers struct {
	cache *lru.LRU
	lock  sync.RWMutex
}

func NewProducers(maxProducers int) *Producers {
	cache, err := lru.NewLRU(maxProducers, func(key interface{}, value interface{}) {
		producer := value.(pulsar.Producer)
		close(producer)
	})
	if err != nil {
		return nil
	}
	return &Producers{
		cache: cache,
	}
}

func (p *Producers) LoadProducer(topic string, client *client) (pulsar.Producer, error) {
	producer, ok := p.getProducer(topic)
	if ok {
		return producer, nil
	}

	return p.getOrCreateProducer(topic, client)
}

func (p *Producers) Len() int {
	p.lock.RLock()
	len := p.cache.Len()
	p.lock.RUnlock()
	return len
}

func (p *Producers) getProducer(topic string) (pulsar.Producer, bool) {
	p.lock.RLock()
	value, ok := p.cache.Get(topic)
	p.lock.RUnlock()

	if ok {
		return value.(pulsar.Producer), ok
	}
	return nil, false
}

func (p *Producers) getOrCreateProducer(topic string, client *client) (pulsar.Producer, error) {
	p.lock.Lock()

	// double check
	value, ok := p.cache.Get(topic)
	if ok {
		p.lock.Unlock()
		logp.Info("pulsar producer already exists for topic: %s", topic)
		return value.(pulsar.Producer), nil
	}

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
	newProducer, err := client.pulsarClient.CreateProducer(newProducerOptions)
	if err == nil {
		p.cache.Add(topic, newProducer)
		logp.Info("created pulsar producer for topic: %s", topic)
	} else {
		logp.Err("creating pulsar producer{topic=%s} failed: %v", topic, err)
	}

	p.lock.Unlock()

	return newProducer, err
}

func (p *Producers) Close() {
	p.lock.Lock()

	for key := range p.cache.Keys() {
		if value, ok := p.cache.Peek(key); ok {
			producer := value.(pulsar.Producer)
			close(producer)
		}
	}

	p.lock.Unlock()
}

func close(producer pulsar.Producer) {
	logp.Info("closing pulsar producer{topic=%s}", producer.Topic())
	err := producer.Flush()
	if err != nil {
		logp.Err("flush pulsar producer{topic=%s} failed: %v", producer.Topic(), err)
	}

	producer.Close()
	logp.Info("closed pulsar producer{topic=%s}", producer.Topic())
}
