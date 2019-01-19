package pulsar

import (
    "fmt"
    "testing"
    "time"
    "sync"
    "context"

    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/outputs"
    "github.com/elastic/beats/libbeat/outputs/outest"
    "github.com/apache/pulsar/pulsar-client-go/pulsar"
)

type eventInfo struct {
    events []beat.Event
}

func makeConfig(t *testing.T, in map[string]interface{}) *common.Config {
    cfg, err := common.NewConfigFrom(in)
    if err != nil {
        t.Fatal(err)
    }
    return cfg
}

func single(fields common.MapStr) []eventInfo {
    return []eventInfo{
        {
            events: []beat.Event{
                {Timestamp: time.Now(), Fields: fields},
            },
        },
    }
}

func flatten(infos []eventInfo) []beat.Event {
    var out []beat.Event
    for _, info := range infos {
        out = append(out, info.events...)
    }
    return out
}

func TestPulsarPublishNoTls(t *testing.T) {
    pulsarConfig := map[string]interface{}{
        "url": "pulsar://localhost:6650",
        "io_threads": 5,
        "topic": "my_topic",
        "bulk_max_size": 2048,
        "max_retries": 3,
    }
    testPulsarPublish(t, pulsarConfig)
}

func testPulsarPublish(t *testing.T, cfg map[string]interface{}){
    logp.Info("start integration test pulsar")

    tests := []struct {
        title  string
        config map[string]interface{}
        topic  string
        events []eventInfo
    }{
        {
            "test single events",
            nil,
            "my-topic1",
            single(common.MapStr{
                "url": "test-host",
                "topic": "my-topic1",
                "name": "test",
            }),
        },
    }

    for i, test := range tests {
        config := makeConfig(t, cfg)
        if test.config != nil {
            config.Merge(makeConfig(t, test.config))
        }
        name := fmt.Sprintf("run test(%v): %v", i, test.title)
        t.Run(name, func(t *testing.T) {
            grp, err := makePulsar(beat.Info{Beat: "libbeat"}, outputs.NewNilObserver(), config)
            if err != nil {
                t.Fatal(err)
            }

            output := grp.Clients[0].(*client)
            if err := output.Connect(); err != nil {
                t.Fatal(err)
            }
            defer output.Close()
            // publish test events
            var wg sync.WaitGroup
            for i := range test.events {
                batch := outest.NewBatch(test.events[i].events...)
                batch.OnSignal = func(_ outest.BatchSignal) {
                    wg.Done()
                }

                wg.Add(1)
                output.Publish(batch)
            }

            // wait for all published batches to be ACKed
            wg.Wait()
            expected := flatten(test.events)

            // stored := testReadFromPulsarTopic(t, output.clientOptions, test.topic, len(expected))
            testReadFromPulsarTopic(t, output.clientOptions, test.topic, len(expected))

        })
    }
}

func testReadFromPulsarTopic (
    t *testing.T, clientOptions pulsar.ClientOptions,
    topic string, nMessages int) []pulsar.Message {
    // Instantiate a Pulsar client
    client, err := pulsar.NewClient(clientOptions)

    if err != nil { t.Fatal(err) }

    // Use the client object to instantiate a consumer
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            topic,
        SubscriptionName: "sub-1",
        Type: pulsar.Shared,
    })

    if err != nil { t.Fatal(err) }

    defer consumer.Close()

    ctx := context.Background()
    var messages []pulsar.Message
    for i := 0; i < nMessages; i++ {
        msg, err := consumer.Receive(ctx)
        if err != nil { t.Fatal(err) }

        // Do something with the message

        consumer.Ack(msg)
        messages = append(messages, msg)
    }
    
    return messages
}