# drone-docker-deploy

[![Build Status](https://ci.convey.cf/api/badges/schrodit/drone-docker-deploy/status.svg)](https://ci.convey.cf/api/badges/schrodit/drone-docker-deploy)

drone ci plugin for faster docker deployment and automated versioning

## Documentation

| Name       | Description                                                                                                                           | Optional |
| ---------- | :------------------------------------------------------------------------------------------------------------------------------------ | :------: |
| repo       | Name of the docker image                                                                                                              |          |
| registry   | Name of a private registry                                                                                                            |    x     |
| dockerfile | Path to a different Dockerfile. Default: ./Dockerfile                                                                                 |    x     |
| directory  | Use a different work directory. Default: ./                                                                                           |    x     |
| usegittag  | Uses the latest git tag or the current git tag if present.<br /> If the current build is a tag event also the "latest"-tag is pushed. |    x     |
| latest     | Do also push "latest"-tag in every build.                                                                                             |    x     |

## Example

```YAML
publish:
    image: schrodit/drone-docker-deploy
    pull: true
    repo: schrodit/drone-docker-deploy
    secrets: [ docker_username, docker_password ]
    usegittag: true
    volumes:
        - /var/run/docker.sock:/var/run/docker.sock:ro
```
