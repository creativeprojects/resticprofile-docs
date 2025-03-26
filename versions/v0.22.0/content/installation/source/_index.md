---
title: Source
weight: 16
---

## Installation from source

You can download the source code and compile it, it's actually very easy! all you need to have on your machine is:
- `git` (with `git-bash` on Windows)
- [go compiler](https://golang.org/dl/)
- `GNU Make` which is installed by default on many unix boxes. On debian based distributions (Ubuntu included) the package is called `build-essential`.

To compile from sources:
```shell
$ git clone https://github.com/creativeprojects/resticprofile.git
$ cd resticprofile
$ make build
```

Your compiled binary (`resticprofile` or `resticprofile.exe`) is available in the current folder.

To install the binary in your user path:

```shell
$ make install
```

{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.21.1" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
