FROM golang:1.12.4

RUN \
    apt-get update \
      && apt-get install -y --no-install-recommends \
         netcat \
         python-pip \
         rsync \
         virtualenv \
         libpcap-dev \
      && rm -rf /var/lib/apt/lists/*

# Apache pulsar go client dependency package
RUN wget https://archive.apache.org/dist/pulsar/pulsar-2.3.0/DEB/apache-pulsar-client.deb
RUN wget https://archive.apache.org/dist/pulsar/pulsar-2.3.0/DEB/apache-pulsar-client-dev.deb
RUN dpkg -i apache-pulsar-client.deb
RUN dpkg -i apache-pulsar-client-dev.deb

RUN git clone https://github.com/streamnative/beat-ouput-pulsar
RUN go build $GOPATH/src/github.com/streamnative/beat-ouput-pulsar
COPY ./filebeat.yml /go/
RUN ls -l
RUN pwd
RUN chown root /go/filebeat.yml
CMD /bin/bash