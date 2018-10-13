# drone-docker-deploy

[![Build Status](https://ci.convey.cf/api/badges/schrodit/drone-docker-deploy/status.svg)](https://ci.convey.cf/api/badges/schrodit/drone-docker-deploy)

drone ci plugin for faster docker deployment and automated versioning

## Documentation

The plugin uses tags presented in the .tags file where the tags are comma seperated.
The build number is added to every build version that isn't a tag event (Format: "Version-JobNumber").

| Parameter name | Description                                                                                                                           | Optional |
| -------------- | :------------------------------------------------------------------------------------------------------------------------------------ | :------: |
| repo           | Name of the docker image                                                                                                              |          |
| registry       | Name of a private registry                                                                                                            |    x     |
| dockerfile     | Path to a different Dockerfile. Default: ./Dockerfile                                                                                 |    x     |
| directory      | Use a different work directory. Default: ./                                                                                           |    x     |
| imagetagsfile      | Use a different output file for pushed tags. Default: .image_tags                                                                                           |    x     |
| usegittag      | Uses the latest git tag or the current git tag if present.<br /> If the current build is a tag event also the "latest"-tag is pushed. |    x     |
| latest         | Do also push "latest"-tag in every build.                                                                                             |    x     |
| addjobnumber   | Add the builds job number with the format "Version-JobNumber" to pushed tags                                                          |    x     |

### Note:

* Plugin has to run in privileged/trusted mode

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
