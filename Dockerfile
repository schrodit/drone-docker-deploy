FROM alpine
RUN apk --update add openrc && \
    apk --update add docker && \
    rc-update add docker boot
VOLUME [ "/var/run/docker.sock:/var/run/docker.sock:ro" ]
ADD src/docker-deploy /bin/
ENTRYPOINT /bin/docker-deploy