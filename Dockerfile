# Beats dockerfile used for testing
FROM golang:1.11.4
MAINTAINER Nicolas Ruflin <ruflin@elastic.co>

RUN set -x && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
         netcat python-pip virtualenv libpcap-dev && \
    apt-get clean

ENV PYTHON_ENV=/tmp/python-env

RUN test -d ${PYTHON_ENV} || virtualenv ${PYTHON_ENV}
COPY ./requirements.txt /tmp/requirements.txt

# Upgrade pip to make sure to have the most recent version
RUN . ${PYTHON_ENV}/bin/activate && pip install -U pip
RUN . ${PYTHON_ENV}/bin/activate && pip install -Ur /tmp/requirements.txt

# Apache pulsar go client dependency package
RUN wget http://mirrors.shu.edu.cn/apache/pulsar/pulsar-2.2.1/DEB/apache-pulsar-client.deb
RUN wget http://mirrors.shu.edu.cn/apache/pulsar/pulsar-2.2.1/DEB/apache-pulsar-client-dev.deb
RUN dpkg -i apache-pulsar-client.deb
RUN dpkg -i apache-pulsar-client-dev.deb

WORKDIR /go

# Libbeat specific
RUN mkdir -p /etc/pki/tls/certs

RUN git clone -b branch-2.1 https://github.com/apache/pulsar $GOPATH/src/github.com/apache/pulsar
RUN go get github.com/AmateurEvents/filebeat-ouput-pulsar
RUN go build $GOPATH/src/github.com/AmateurEvents/filebeat-ouput-pulsar
COPY ./filebeat.yml /go/
RUN ls -l
RUN pwd
RUN chown root /go/filebeat.yml
#CMD /go/filebeat-ouput-pulsar -e -c /go/filebeat.yml
CMD /bin/bash
