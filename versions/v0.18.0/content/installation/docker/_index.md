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

[1]: https://rclone.org/

{{< pageversions "v0.19.0" "v0.20.0" "v0.21.0" >}}
