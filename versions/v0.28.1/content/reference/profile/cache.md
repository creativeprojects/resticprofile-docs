---
title: cache
weight: 2
---
#### Section profile.**cache**

This section configures restic command `cache` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "cache" command allows listing and cleaning local cache directories.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **cleanup** |`true` / `false` |`false` |remove old cache directories |
| **max-age** |`integer` |`30` |max age in days for cache directories to be considered old |
| **no-size** |`true` / `false` |`false` |do not output the size of the cache directories |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.29.1" >}}
