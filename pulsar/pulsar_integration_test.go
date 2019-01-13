

func makeConfig(t *testing.T, in map[string]interface{}) *common.Config {
    cfg, err := common.NewConfigFrom(in)
    if err != nil {
        t.Fatal(err)
    }
    return cfg
}

func TestPulsarPublish(t *testing.T, cfg interface{}) (outputs.Client, *Client) {
    logp.Info("start integration test pulsar")

    tests := []struct {
        title  string
        config map[string]interface{}
        topic  string
        events []eventInfo
    }{
        {

        }
    }

    config := defaultConfig()
    for i, test := range tests {
        if test.config != nil {
            cfg.Merge(makeConfig(t, test.config))
        }
        clientOptions, producerOptions, err := initOptions(&config)
        name := fmt.Sprintf("run test(%v): %v", i, test.title)
        t.Run(name, func(t *testing.T) {
            grp, err := makePulsar(beat.Info{Beat: "libbeat"}, outputs.NewNilObserver(), cfg)
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

            stored := testReadFromPulsarTopic(t, test.topic, len(expected), timeout)

        })
    }
}

func testReadFromPulsarTopic (
    t *testing.T, clientOptions pulsar.ClientOptions,
    topic string, nMessages int,
    timeout time.Duration) []*pulsar.Message {
    // Instantiate a Pulsar client
    client, err := pulsar.NewClient(clientOptions)

    if err != nil { log.Fatal(err) }

    // Use the client object to instantiate a consumer
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            topic,
        SubscriptionName: "sub-1",
        SubscriptionType: pulsar.Exclusive,
    })

    if err != nil { log.Fatal(err) }

    defer consumer.Close()

    ctx := context.Background()
    var messages []*pulsar.Message
    for i := 0; i < nMessages; i++ {
        msg, err := consumer.Receive(ctx)
        if err != nil { log.Fatal(err) }

        // Do something with the message

        consumer.Ack(msg)
        messages = append(messages, msg)
    }
    
    return messages
}