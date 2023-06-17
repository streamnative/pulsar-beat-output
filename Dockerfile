FROM golang:1.18 as build-env

WORKDIR /go/src/github.com/streamnative/pulsar-beat-output
ADD . /go/src/github.com/streamnative/pulsar-beat-output

RUN CGO_ENABLED=0 go build -o /go/bin/filebeat github.com/streamnative/pulsar-beat-output/filebeat
RUN CGO_ENABLED=0 go build -o /go/bin/heartbeat github.com/streamnative/pulsar-beat-output/heartbeat
RUN CGO_ENABLED=0 go build -o /go/bin/metricbeat github.com/streamnative/pulsar-beat-output/metricbeat
RUN CGO_ENABLED=0 go build -o /go/bin/packetbeat github.com/streamnative/pulsar-beat-output/packetbeat
RUN CGO_ENABLED=0 go build -o /go/bin/auditbeat github.com/streamnative/pulsar-beat-output/auditbeat
RUN CGO_ENABLED=0 go build -o /go/bin/winlogbeat github.com/streamnative/pulsar-beat-output/winlogbeat

