FROM tuteng/pulsar-authentication-authorization:latest

EXPOSE 6650
EXPOSE 6651
EXPOSE 8080
EXPOSE 8443

COPY ./functions_worker.yml conf/
COPY ./standalone.conf conf/

# HEALTHCHECK --interval=1s --retries=600 CMD nc -z localhost 6651

ENTRYPOINT ["supervisord", "-n"]
