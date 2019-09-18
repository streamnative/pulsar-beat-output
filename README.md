## Beat Output Pulsar

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fstreamnative%2Fpulsar-beat-output.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fstreamnative%2Fpulsar-beat-output?ref=badge_shield)

This is a output implementation of [elastic beats](https://github.com/elastic/beats) for support [Filebeat](https://github.com/elastic/beats/tree/master/filebeat), [Metricbeat](https://github.com/elastic/beats/tree/master/metricbeat), [Functionbeat](https://github.com/elastic/beats/tree/master/x-pack/functionbeat), [Winlogbeat](https://github.com/elastic/beats/tree/master/winlogbeat), [Journalbeat](https://github.com/elastic/beats/tree/master/journalbeat), [Auditbeat](https://github.com/elastic/beats/tree/master/auditbeat) to [Apache Pulsar](https://github.com/apache/pulsar)

### Compatibility
This output is developed and tested using Apache Pulsar Client 2.4.0 and Beats 7.3.1

### Download pulsar-beat-output

```
mkdir -p $GOPATH/src/github.com/streamnative/
cd $GOPATH/src/github.com/streamnative/
git clone https://github.com/streamnative/pulsar-beat-output
cd pulsar-beat-output
```

### Build

#### Build Filebeat

Edit main.go file

```
package main

import (
    "os"
    _ "github.com/streamnative/pulsar-beat-output/pulsar"
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
  url: "pulsar://localhost:6650"
  topic: my_topic
  name: test123
```
#### Start filebeat
```
./filebeat modules enable system
./filebeat modules list
./filebeat -c filebeat.yml -e
```

#### Build other beat

```
go build -o metricbeat metricbeat.go
go build -o filebeat filebeat.go
go build -o functionbeat functionbeat.go
go build -o journalbeat journalbeat.go
go build -o auditbeat auditbeat.go
go build -o winlogbeat winlogbeat.go
go build -o packetbeat packetbeat.go
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
docker run -d -it --network pulsar-beat -p 6650:6650 -p 8080:8080 -v $PWD/data:/pulsar/data --name pulsar-beat-standalone apachepulsar/pulsar:2.4.0 bin/pulsar standalone
```

#### Add following configuration to filebeat.yml
```
output.pulsar:
  url: "pulsar://pulsar-beat-standalone:6650"
  topic: my_topic
  name: test123
```

#### Start Filebeat
```
docker pull golang:1.12.4
docker run -it --network pulsar-beat --name filebeat golang:1.12.4 /bin/bash
mkdir -p $GOPATH/src/github.com/streamnative/
cd $GOPATH/src/github.com/streamnative/
git clone https://github.com/streamnative/pulsar-beat-output
cd pulsar-beat-output
go build -o filebeat main.go
chown -R root:root filebeat.yml test_module/modules.d/system.yml test_module/module/system
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

#### Producer
|Name|Description|Default|
|---|---|---|
|topic| Specify the topic this producer will be publishing on. |""|
|name| Specify a name for the producer |""|
|send_timeout| Set the send timeout |30s|
|block_if_queue_full| Set whether the send and sendAsync operations should block when the outgoing message queue is full. |false|
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

#### Build journalbeat.go

```
systemd/sd-journal.h: No such file or directory
```

```
apt-get update
apt-get install libsystemd-dev
```

#### Build auditbeat.go
```
vendor/github.com/elastic/beats/x-pack/auditbeat/module/system/package/rpm_linux.go:23:24: fatal error: rpm/rpmlib.h: No such file or directory
```

```
aapt-get install librpm-dev
```

#### Start beat
```
Exiting: error loading config file: config file ("filebeat.yml") must be owned by the user identifier (uid=0) or root
```
```
chown -R root:root filebeat.yml
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fstreamnative%2Fpulsar-beat-output.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fstreamnative%2Fpulsar-beat-output?ref=badge_large)