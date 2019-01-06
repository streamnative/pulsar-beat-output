package pulsar

import (
    "context"

    "github.com/apache/pulsar/pulsar-client-go/pulsar"
    "github.com/elastic/beats/libbeat/outputs"
    "github.com/elastic/beats/libbeat/outputs/codec"
    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/publisher"
    "github.com/elastic/beats/libbeat/logp"
)


type client struct {
    clientOptions pulsar.ClientOptions
    producerOptions pulsar.ProducerOptions
    pulsarClient pulsar.Client
    producer pulsar.Producer
    observer outputs.Observer
    beat     beat.Info
    config *pulsarConfig
    codec codec.Codec
}

func newPulsarClient(
    beat beat.Info,
    observer outputs.Observer,
    clientOptions pulsar.ClientOptions,
    producerOptions pulsar.ProducerOptions,
    config *pulsarConfig,
) (*client, error) {
    c := &client{
        clientOptions: clientOptions,
        producerOptions: producerOptions,
        observer: observer,
        beat: beat,
        config: config,
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
    for i := range events {
        event := &events[i]
        serializedEvent, err := c.codec.Encode(c.beat.Beat, &event.Content)
        if err != nil {
            logp.Debug("pulsar", "Failed event: %v, error: %v", event, err)
        }

        err = c.producer.Send(context.Background(), pulsar.ProducerMessage{
            Payload: []byte(serializedEvent),
        })
        if err != nil {
            logp.Debug("pulsar", "produce send failed: %v", err)
        }
    }
    return nil
}

func (c *client) String() string {
    return "file(" + c.clientOptions.URL + ")"
}