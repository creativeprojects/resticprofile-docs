---
title: dump
weight: 7
---
#### Section profile.**dump**

This section configures restic command `dump` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "dump" command extracts files from a snapshot from the repository. If a
single file is selected, it prints its contents to stdout. Folders are output
as a tar (default) or zip file containing the contents of the specified folder.
Pass "/" as file name to dump the whole snapshot as an archive file.

The special snapshot "latest" can be used to use the latest snapshot in the
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
| **archive** |`string` |`"tar"` |set archive format as "tar" or "zip". `restic >= 0.12.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" "v0.29.1" >}}
