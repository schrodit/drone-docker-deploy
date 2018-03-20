FROM docker:17.12.1-ce
VOLUME /var/run/docker.sock:/var/run/docker.sock:ro
ADD src/docker-deploy /bin/
ENTRYPOINT /bin/docker-deploy