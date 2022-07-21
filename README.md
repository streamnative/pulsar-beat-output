## Beat Output Pulsar

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fstreamnative%2Fpulsar-beat-output.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fstreamnative%2Fpulsar-beat-output?ref=badge_shield)

This is an output implementation of [elastic beats](https://github.com/elastic/beats) for support [Filebeat](https://github.com/elastic/beats/tree/master/filebeat), [Metricbeat](https://github.com/elastic/beats/tree/master/metricbeat), [Functionbeat](https://github.com/elastic/beats/tree/master/x-pack/functionbeat), [Winlogbeat](https://github.com/elastic/beats/tree/master/winlogbeat), [Journalbeat](https://github.com/elastic/beats/tree/master/journalbeat), [Auditbeat](https://github.com/elastic/beats/tree/master/auditbeat) to [Apache Pulsar](https://github.com/apache/pulsar)

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

#### Build beats

```
go build -o filebeat filebeat/filebeat.go
go build -o functionbeat functionbeat/functionbeat.go
go build -o journalbeat journalbeat/journalbeat.go
go build -o winlogbeat winlogbeat/winlogbeat.go
go build -o packetbeat packetbeat/packetbeat.go
```

### Usage

In this section, you can use the sample config file in the directory [./sample/config/], or you can create it as follow steps.

#### example
#### Add following configuration to beat.yml 
```yml
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
docker run -d -it --network pulsar-beat -p 6650:6650 -p 8080:8080 -v $PWD/data:/pulsar/data --name pulsar-beat-standalone apachepulsar/pulsar:2.7.0 bin/pulsar standalone
```

#### Add following configuration to filebeat.yml
```yml
output.pulsar:
  url: "pulsar://pulsar-beat-standalone:6650"
  topic: my_topic
  name: test123
```

#### Start Filebeat
```
docker pull golang:1.17
docker run -it --network pulsar-beat --name filebeat golang:1.17 /bin/bash
git clone https://github.com/streamnative/pulsar-beat-output
cd pulsar-beat-output
go build -o filebeat filebeat/filebeat.go
chown -R root:root filebeat.yml test_module/modules.d/system.yml test_module/module/system
cp test_module/module/system/auth/test/test.log /var/log/messages.log
cp filebeat/filebeat filebeat.yml test_module
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
|token| Access token information of cluster | "" |
|token_file_path| The file path where token is saved | "" |
|log_level| Setting the log level, available options(panic, fatal, error, warn, info, debug, trace) | info |
|oauth2.enabled| Enabled or disabled oauth2 authentication | false |
|oauth2.clientId| client ID | "" |
|oauth2.issuerUrl| URL of the authentication provider which allows the Pulsar client to obtain an access token | "" |
|oauth2.privateKey| URL of a JSON credentials file | "" |
|oauth2.audience| The audience value is either the application (`Client ID`) for an ID Token or the API that is being called (`API Identifier`) for an Access Token | "" |
|oauth2.scope| Scope is a mechanism in OAuth 2.0 to limit an application's access to a user's account | "" |



#### Producer
|Name| Description                                                                                                                                                                       | Default |
|---|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
|topic| Specify the topic this producer will be publishing on. You can set the topic dynamically by using a format string to access any event field. For example `%{[fields.log_topic]}`. | ""      |
|partition_key| Specify the message key. You can set the message key dynamically by using a format string to access any event field. For example `%{[fields.partition_key]}`                      | ""      |
|name| Specify a name for the producer                                                                                                                                                   | ""      |
|send_timeout| Set the send timeout                                                                                                                                                              | 30s     |
|block_if_queue_full| Set whether the send and sendAsync operations should block when the outgoing message queue is full.                                                                               | false   |
|batching_max_messages| maximum number of messages in a batch                                                                                                                                             | 1000    |
|batching_max_publish_delay| the batch delay                                                                                                                                                                   | 1ms     |
|message_routing_mode| the message routing mode, SinglePartition,RoundRobinPartition, CustomPartition(0,1,2)                                                                                             | 1       |
|hashing_schema| JavaStringHash,Murmur3_32Hash(0,1)                                                                                                                                                | 0       |
|compression_type| NONE,LZ4,ZLIB,ZSTD(0,1,2,3)                                                                                                                                                       | 0       |
|max_cache_producers| Specify the maximun cache(lru) producers of dynamic topic.                                                                                                                                                        | 8       |

### FAQ

#### case-insensitive import collision: "github.com/datadog/zstd" and "github.com/DataDog/zstd"

```
/root/go/pkg/mod/github.com/apache/pulsar-client-go@v0.3.0/pulsar/internal/compression/zstd_cgo.go:27:2: case-insensitive import collision: "github.com/datadog/zstd" and "github.com/DataDog/zstd"
```

Replace zstd_cgo.go file
```
cp zstd_cgo.go /root/go/pkg/mod/github.com/apache/pulsar-client-go@v0.3.0/pulsar/internal/compression/zstd_cgo.go
```

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
