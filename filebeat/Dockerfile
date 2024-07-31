FROM elastic/filebeat:7.17.5

COPY dist/filebeat_linux_amd64_v1/beat /usr/share/filebeat/filebeat

USER root

RUN mkdir -p /usr/share/filebeat/config /usr/share/filebeat/secret && \
    chown -R filebeat:filebeat /usr/share/filebeat

USER filebeat