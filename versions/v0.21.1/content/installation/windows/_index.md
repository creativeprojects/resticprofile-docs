---
title: Windows
weight: 12
---



## Installation using bash

You can use a script if you're using bash in Windows (via WSL, git bash, etc.)

```shell
$ curl -LO https://raw.githubusercontent.com/creativeprojects/resticprofile/master/install.sh
$ ./install.sh
```
It will create a `bin` directory under your current directory and place `resticprofile.exe` in it.

## Manual installation (Windows)

- Download the package corresponding to your system and CPU from the [release page](https://github.com/creativeprojects/resticprofile/releases)
- Once downloaded you need to open the archive and copy the binary file `resticprofile` (or `resticprofile.exe`) in your PATH.

{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" >}}
