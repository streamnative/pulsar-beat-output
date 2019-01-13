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

# Libbeat specific
RUN mkdir -p /etc/pki/tls/certs
