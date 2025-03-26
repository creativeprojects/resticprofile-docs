---
title: Docker
weight: 30
---



## Using resticprofile from a docker image ##

You can run resticprofile inside a docker container. It is probably the easiest way to install resticprofile (and restic at the same time) and keep it updated.

**But** be aware that you will need to mount your backup source (and destination if it's local) as a docker volume.
Depending on your operating system, the backup might be **slower**. Volumes mounted on a mac OS host are well known for being quite slow.

By default, the resticprofile container starts at `/resticprofile`. So you can feed a configuration this way:

```shell
$ docker run -it --rm -v $PWD/examples:/resticprofile creativeprojects/resticprofile
```

You can list your profiles:
```shell
$ docker run -it --rm -v $PWD/examples:/resticprofile creativeprojects/resticprofile profiles
```

### Container host name

Each time a container is started, it gets assigned a new random name.

You might want to force a hostname when starting your container via docker run (flags `-h` or `--hostname`):

```shell
$ docker run -it --rm -v $PWD:/resticprofile -h my-hostname creativeprojects/resticprofile -n profile backup
```

### Platforms

Starting from version `0.18.0`, the resticprofile docker image is available in these 2 platforms:
- linux/amd64
- linux/arm64/v8 (compatible with raspberry pi 64bits)

### rclone

Starting from version `0.18.0`, the resticprofile docker image also includes [rclone][1].

## Scheduling with docker compose

There's an example in the contribution section how to schedule backups in a long running container.
The configuration needs to specify `crond` as a scheduler.

See [contrib][2]

[1]: https://rclone.org/
[2]: https://github.com/creativeprojects/resticprofile/tree/master/contrib/schedule-in-docker

{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.21.1" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
