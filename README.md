## Beat Output Pulsar

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

This is a output implementation of [elastic beats](https://github.com/elastic/beats) for support [Filebeat](https://github.com/elastic/beats/tree/master/filebeat), [Metricbeat](https://github.com/elastic/beats/tree/master/metricbeat), [Functionbeat](https://github.com/elastic/beats/tree/master/x-pack/functionbeat), [Winlogbeat](https://github.com/elastic/beats/tree/master/winlogbeat), [Journalbeat](https://github.com/elastic/beats/tree/master/journalbeat), [Auditbeat](https://github.com/elastic/beats/tree/master/auditbeat) to [Apache Pulsar](https://github.com/apache/pulsar)

### Compatibility
This output is developed and tested using Apache Pulsar Client 2.3.0 and Beats 7.0.0

### Download beat-output-pulsar

```
go get github.com/streamnative/beat-ouput-pulsar
cd $GOPATH/src/github.com/streamnative/beat-ouput-pulsar
```

### Build

#### Build Filebeat

Edit main.go file

```
package main

import (
    "os"
    _ "github.com/streamnative/beat-ouput-pulsar/pulsar"
    "github.com/elastic/beats/x-pack/filebeat/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

```
go build -o filebeat main.go
```

### Usage

#### Add following configuration to filebeat.yml
```
output.pulsar:
  url: ["pulsar://localhost:6650"]
  topic: my_topic
  name: test123
```
#### Start filebeat
```
./filebeat modules enable system
./filebeat modules list
./filebeat -c filebeat.yml
```

#### Build other beat, for example metricbeat

Edit main.go file

```
package main

import (
    "os"
    _ "github.com/streamnative/beat-ouput-pulsar/pulsar"
    "github.com/elastic/beats/x-pack/metricbeat/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

```
go build -o metricbeat main.go
```
### Build and test with docker

#### Requirements

- [Docker](https://docs.docker.com/docker-for-mac/install/)

#### Build Beat images

```
docker build -t pulsar-beat .
```

#### Create network

```
docker network create pulsar-beat
```

#### Start Pulsar service
```
docker run -d -it --network pulsar-beat -p 6650:6650 -p 8080:8080 -v $PWD/data:/pulsar/data --name pulsar-beat-standalone apachepulsar/pulsar:2.3.0 bin/pulsar standalone
```

#### Add following configuration to filebeat.yml
```
output.pulsar:
  url: ["pulsar://pulsar-beat-standalone:6650"]
  topic: my_topic
  name: test123
```

#### Start Filebeat
```
docker run -it --network pulsar-beat --name filebeat pulsar-beat /bin/bash
cd $GOPATH/src/github.com/streamnative/beat-ouput-pulsar
chown -R root:root filebeat.yml
chown -R root:root test_module/modules.d/system.yml
chown -R root:root test_module/module/system
cp test_module/module/system/auth/test/test.log /var/log/messages.log
cp filebeat filebeat.yml test_module
cd test_module
./filebeat modules enable system
./filebeat -c filebeat.yml -e
```

#### New open a window for consumer message
```
docker cp pulsar-client.py pulsar-beat-standalone:/pulsar
docker exec -it pulsar-beat-standalone /bin/bash
python pulsar-client.py
```
Now you can see the information collected from filebeat.

### Configurations

#### Client
|Name|Description|Default|
|---|---|---|
|url| Configure the service URL for the Pulsar service |pulsar://localhost:6650|
|certificate_path| path of tls cert file |""|
|private_key_path| path of tls key file |""|
|use_tls| Whether to turn on TLS, if to start, use protocol pulsar+ssl |false|
|operation_timeout_seconds| Set the operation timeout (default: 30 seconds) |30s|
|io_threads| Set the number of threads to be used for handling connections to brokers |1|
|message_listener_threads| Set the number of threads to be used for message listeners |1|
|tls_trust_certs_file_path| Set the path to the trusted TLS certificate file |false|
|tls_allow_insecure_connection| Configure whether the Pulsar client accept untrusted TLS certificate from broker |false|
|stats_interval_in_seconds| the interval between each stat info |60|
|concurrent_lookup_requests| Number of concurrent lookup-requests allowed to send on each broker-connection to prevent overload on broker. |60|

#### Producer
|Name|Description|Default|
|---|---|---|
|topic| Specify the topic this producer will be publishing on. |""|
|name| Specify a name for the producer |""|
|send_timeout| Set the send timeout |30s|
|block_if_queue_full| Set whether the send and sendAsync operations should block when the outgoing message queue is full. |false|
|batching| Control whether automatic batching of messages is enabled for the producer |true|
|batching_max_messages| maximum number of messages in a batch |1000|
|batching_max_publish_delay| the batch delay |1ms|
|message_routing_mode| the message routing mode, SinglePartition,RoundRobinPartition, CustomPartition(0,1,2) |1|
|hashing_schema| JavaStringHash,Murmur3_32Hash(0,1) |0|
|compression_type| NONE,LZ4,ZLIB,ZSTD(0,1,2,3) |0|

### FAQ

#### Install Pulsar Go Client
Reference https://pulsar.apache.org/docs/en/client-libraries-go/ .

If you encounter problems with dynamic libraries，please reference:https://pulsar.apache.org/docs/en/client-libraries-cpp/.

#### Build Packetbeat
Reference https://github.com/elastic/beats/issues/11054.

#### Start beat
