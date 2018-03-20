FROM alpine
RUN apk add docker && \
    apk -Uuv add ca-certificates
VOLUME [ "/var/run/docker.sock:/var/run/docker.sock:ro" ]
ADD src/dev-deploy /bin/
ENTRYPOINT /bin/docker-deploy