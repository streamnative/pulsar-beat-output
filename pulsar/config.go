package pulsar

import (
    "time"
    "errors"

    "github.com/apache/pulsar/pulsar-client-go/pulsar"
    "github.com/elastic/beats/libbeat/outputs/codec"
)


type pulsarConfig struct {
    URL string `config:"url"`
    IOThreads int `config:"io_threads"`
    OperationTimeoutSeconds time.Duration `config:"operation_timeout_seconds"`
    MessageListenerThreads int `config:"message_listener_threads"`
    ConcurrentLookupRequests int `config:"concurrent_lookup_requests"`
    // Logger func(level LoggerLevel, file string, line int, message string)
    TLSTrustCertsFilePath string `config:"tls_trust_certs_file_path"`
    TLSAllowInsecureConnection bool `config:"tls_allow_insecure_connection"`
    StatsIntervalInSeconds int `config:"stats_interval_in_seconds"`

    Codec codec.Config `config:"codec"`
    BulkMaxSize int `config:"bulk_max_size"`
    MaxRetries int `config:"max_retries" validate:"min=-1,nonzero"`

    Topic string `config:"topic"`
    Name string `config:"name"`
    Properties map[string]string `config:"properties"`
    SendTimeout time.Duration `config:"send_timeout"`
    MaxPendingMessages int `config:"max_pending_messages"`
    MaxPendingMessagesAcrossPartitions int `config:"max_pending_messages_accross_partitions"`
    BlockIfQueueFull bool `config:"block_if_queue_full"`
    MessageRoutingMode int `config:"message_routing_mode"`
    HashingScheme int `config:"hashing_schema"`
    CompressionType int `config:"compression_type"`
    // MessageRouter func(Message, TopicMetadata) int `config:"message_router"`
    Batching bool `config:"batching"`
    BatchingMaxPublishDelay time.Duration `config:"batching_max_publish_delay"`
    BatchingMaxMessages uint `config:"batching_max_messages"`

}

func defaultConfig() pulsarConfig {
    return pulsarConfig{
        URL: "pulsar://localhost:6650",
        IOThreads: 5,
        Topic: "my_topic",
        BulkMaxSize: 2048,
        MaxRetries: 3,
    }
}

func (c *pulsarConfig) Validate() error {
    if len(c.URL) == 0 {
        return errors.New("no URL configured")
    }
    if len(c.Topic) == 0 {
        return errors.New("no topic configured")
    }
    return nil
}

func initOptions(
    config *pulsarConfig,
) (pulsar.ClientOptions, pulsar.ProducerOptions, error) {
    config.Validate()
    clientOptions := pulsar.ClientOptions{
        URL: config.URL,
        IOThreads: config.IOThreads,
    }
    producerOptions := pulsar.ProducerOptions{
        Topic: config.Topic,
    }
    return clientOptions, producerOptions, nil
}
