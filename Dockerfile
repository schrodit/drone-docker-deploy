FROM alpine
RUN apk --update add docker && \
    service docker start
VOLUME [ "/var/run/docker.sock:/var/run/docker.sock:ro" ]
ADD src/docker-deploy /bin/
ENTRYPOINT /bin/docker-deploy