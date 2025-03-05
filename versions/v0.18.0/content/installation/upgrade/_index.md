---
title: Upgrade
weight: 20
---

Once installed, you can easily upgrade resticprofile to the latest release using this command:

```shell
$ resticprofile self-update
```

_Please note on versions before 0.10.0, there was an issue with self-updating from linux with ARM processors (like a raspberry pi). This was fixed in version 0.10.0_

resticprofile will check for a new version from GitHub releases and asks you if you want to update to the new version. If you add the flag `-q` or `--quiet` to the command line, it will update automatically without asking.

```shell
$ resticprofile --quiet self-update
```

and since version 0.11.0:

```shell
$ resticprofile self-update --quiet
```

{{< pageversions "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" >}}
