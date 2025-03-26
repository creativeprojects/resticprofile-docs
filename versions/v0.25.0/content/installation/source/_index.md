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
git clone https://github.com/creativeprojects/resticprofile.git
cd resticprofile
make build
```

Your compiled binary (`resticprofile` or `resticprofile.exe`) is available in the current folder.

To install the binary in your user path:

```shell
make install
```

To build all common platforms (`build-mac`, `build-linux`, `build-pi` & `build-windows`):

```shell
make build-all
```

Alternatively, a **go-only** build (without `GNU Make`) is accomplished with:

```shell
git clone https://github.com/creativeprojects/resticprofile.git
cd resticprofile
go mod download
go generate ./... 
go build -o resticprofile .
```


{{% notice style="note" %}}

The build step `go generate ./...` installs additional binaries `github.com/zyedidia/eget` and `github.com/vektra/mockery` into `$GOPATH/bin`.
See `generate.go`, `Makefile` and `go.mod` for details on additional dependencies.

{{% /notice %}}


{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
