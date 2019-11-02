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
	"errors"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/elastic/beats/libbeat/outputs/codec"
)

type pulsarConfig struct {
	URL                        string        `config:"url"`
	IOThreads                  int           `config:"io_threads"`
	OperationTimeoutSeconds    time.Duration `config:"operation_timeout_seconds"`
	MessageListenerThreads     int           `config:"message_listener_threads"`
	ConcurrentLookupRequests   int           `config:"concurrent_lookup_requests"`
	UseTLS                    bool          `config:"use_tls"`
	TLSTrustCertsFilePath      string        `config:"tls_trust_certs_file_path"`
	TLSAllowInsecureConnection bool          `config:"tls_allow_insecure_connection"`
	CertificatePath            string        `config:"certificate_path"`
	PrivateKeyPath             string        `config:"private_key_path"`
	StatsIntervalInSeconds     int           `config:"stats_interval_in_seconds"`

	Codec       codec.Config `config:"codec"`
	BulkMaxSize int          `config:"bulk_max_size"`
	MaxRetries  int          `config:"max_retries" validate:"min=-1,nonzero"`

	Topic                              string                    `config:"topic"`
	Name                               string                    `config:"name"`
	Properties                         map[string]string         `config:"properties"`
	SendTimeout                        time.Duration             `config:"send_timeout"`
	MaxPendingMessages                 int                       `config:"max_pending_messages"`
	MaxPendingMessagesAcrossPartitions int                       `config:"max_pending_messages_accross_partitions"`
	BlockIfQueueFull                   bool                      `config:"block_if_queue_full"`
	HashingScheme                      pulsar.HashingScheme      `config:"hashing_schema"`
	CompressionType                    pulsar.CompressionType    `config:"compression_type"`
	DisableBatching                           bool                      `config:"disable_batching"`
	BatchingMaxPublishDelay            int             `config:"batching_max_publish_delay"`
	BatchingMaxMessages                uint                      `config:"batching_max_messages"`
}

func defaultConfig() pulsarConfig {
	return pulsarConfig{
		URL:         "pulsar://localhost:6650",
		IOThreads:   5,
		Topic:       "my_topic",
		BulkMaxSize: 2048,
		MaxRetries:  3,
	}
}

func (c *pulsarConfig) Validate() error {
	if len(c.URL) == 0 {
		return errors.New("no URL configured")
	}
	if len(c.Topic) == 0 {
		return errors.New("no topic configured")
	}
	if c.UseTLS {
		if len(c.TLSTrustCertsFilePath) == 0 {
			return errors.New("no tls_trust_certs_file_path configured")
		}
		if len(c.CertificatePath) > 0 {
			if len(c.PrivateKeyPath) == 0 {
				return errors.New("no private_key_path configured")
			}
		}
	}
	if c.BulkMaxSize < 0 {
		return errors.New("bulk max size is incorrect")
	}
	if c.CompressionType < 0 {
		return errors.New("compression_type is incorrect")
	}
	return nil
}

func initOptions(
	config *pulsarConfig,
) (pulsar.ClientOptions, pulsar.ProducerOptions, error) {
	config.Validate()
	clientOptions := pulsar.ClientOptions{
		URL:       config.URL,
	}
	if config.UseTLS {
		clientOptions.TLSTrustCertsFilePath = config.TLSTrustCertsFilePath
		if len(config.CertificatePath) > 0 {
			clientOptions.Authentication = pulsar.NewAuthenticationTLS(config.CertificatePath, config.PrivateKeyPath)
		}
	}
	// The client has not added these options yet.
	// if config.IOThreads > 0 {
	// 	clientOptions.IOThreads = config.IOThreads
	// }
	// if config.OperationTimeoutSeconds > 0 {
	// 	clientOptions.OperationTimeoutSeconds = config.OperationTimeoutSeconds * time.Second
	// }
	// if config.MessageListenerThreads > 0 {
	// 	clientOptions.MessageListenerThreads = config.MessageListenerThreads
	// }
	// if config.ConcurrentLookupRequests > 0 {
	// 	clientOptions.ConcurrentLookupRequests = config.ConcurrentLookupRequests
	// }
	// if config.TLSAllowInsecureConnection {
	// 	clientOptions.TLSAllowInsecureConnection = config.TLSAllowInsecureConnection
	// }
	// if config.StatsIntervalInSeconds > 0 {
	// 	clientOptions.StatsIntervalInSeconds = config.StatsIntervalInSeconds
	// }
	producerOptions := pulsar.ProducerOptions{
		Topic: config.Topic,
	}
	if len(config.Name) > 0 {
		producerOptions.Name = config.Name
	}
	if config.SendTimeout > 0 {
		producerOptions.SendTimeout = config.SendTimeout * time.Second
	}
	if len(config.Properties) > 0 {
		producerOptions.Properties = config.Properties
	}
	if config.MaxPendingMessages > 0 {
		producerOptions.MaxPendingMessages = config.MaxPendingMessages
	}
	if config.MaxPendingMessagesAcrossPartitions > 0 {
		producerOptions.MaxPendingMessagesAcrossPartitions = config.MaxPendingMessagesAcrossPartitions
	}
	if config.BlockIfQueueFull {
		producerOptions.BlockIfQueueFull = config.BlockIfQueueFull
	}
	if config.HashingScheme > 0 {
		producerOptions.HashingScheme = config.HashingScheme
	}
	if config.CompressionType > 0 {
		producerOptions.CompressionType = config.CompressionType
	}
	if config.DisableBatching {
		producerOptions.DisableBatching = config.DisableBatching
	}
	if config.BatchingMaxPublishDelay > 0 {
		producerOptions.BatchingMaxPublishDelay = time.Duration(config.BatchingMaxPublishDelay) * time.Second
	} else {
		producerOptions.BatchingMaxPublishDelay = time.Duration(1) * time.Second
	}
	if config.BatchingMaxMessages > 0 {
		producerOptions.BatchingMaxMessages = config.BatchingMaxMessages
	}
	return clientOptions, producerOptions, nil
}
