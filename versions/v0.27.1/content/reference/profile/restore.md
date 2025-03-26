---
title: restore
weight: 22
---
#### Section profile.**restore**

This section configures restic command `restore` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "restore" command extracts the data from a snapshot from the repository to
a directory.

The special snapshot "latest" can be used to restore the latest snapshot in the
repository.

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
| **exclude** |one or more `strings` | |exclude a pattern |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **iexclude** |one or more `strings` | |same as --exclude but ignores the casing of filenames. `restic >= 0.10.0`  |
| **iinclude** |one or more `strings` | |same as --include but ignores the casing of filenames. `restic >= 0.10.0`  |
| **include** |one or more `strings` | |include a pattern, exclude everything else |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **sparse** |`true` / `false` |`false` |restore files as sparse. `restic >= 0.15.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **target** |`string` |`""` |directory to extract data to |
| **verify** |`true` / `false` |`false` |verify restored files content |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

