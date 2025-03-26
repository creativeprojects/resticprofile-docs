---
title: snapshots
weight: 30
---
#### Section profile.**snapshots**

This section configures restic command `snapshots` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "snapshots" command lists all snapshots stored in the repository.

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
| **compact** |`true` / `false` |`false` |use compact output format |
| **group-by** |`string` | |group snapshots by host, paths and/or tags, separated by comma. `restic >= 0.10.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| ~~last~~ |`true` / `false` |`false` |only show the last snapshot for each host and path. `restic < 0.13.0`  |
| **latest** |`integer` |`0` |only show the last n snapshots for each host and path. `restic >= 0.13.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.29.1" >}}
