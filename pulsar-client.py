import pulsar

client = pulsar.Client('pulsar://localhost:6650')
consumer = client.subscribe('my_topic', subscription_name='test123')

while True:
    msg = consumer.receive()
    print msg
    print dir(msg)
    print("Received message: '%s'" % msg.data())
    consumer.acknowledge(msg)

client.close()