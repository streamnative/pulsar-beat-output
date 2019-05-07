FROM golang:1.12.4

# Apache pulsar go client dependency package
RUN wget https://archive.apache.org/dist/pulsar/pulsar-2.3.0/DEB/apache-pulsar-client.deb
RUN wget https://archive.apache.org/dist/pulsar/pulsar-2.3.0/DEB/apache-pulsar-client-dev.deb
RUN dpkg -i apache-pulsar-client.deb
RUN dpkg -i apache-pulsar-client-dev.deb

RUN go get github.com/streamnative/pulsar-beat-output
CMD /bin/bash
