FROM docker:17.12.1-ce
ADD src/docker-deploy /bin/
ENTRYPOINT /bin/docker-deploy