---
title: ls
weight: 13
---
#### Section profile.**ls**

This section configures restic command `ls` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "ls" command lists files and directories in a snapshot.

The special snapshot ID "latest" can be used to list files and
directories of the latest snapshot in the repository. The
--host flag can be used in conjunction to select the latest
snapshot originating from a certain host only.

File listings can optionally be filtered by directories. Any
positional arguments after the snapshot ID are interpreted as
absolute directory paths, and only files inside those directories
will be listed. If the --recursive flag is used, then the filter
will allow traversing into matching directories' subfolders.
Any directory paths specified must be absolute (starting with
a path separator); paths use the forward slash '/' as separator.

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
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **human-readable** |`true` / `false` |`false` |print sizes in human readable format. `restic >= 0.16.0`  |
| **long** |`true` / `false` |`false` |use a long listing format showing size and mode |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **recursive** |`true` / `false` |`false` |include files in subfolders of the listed directories |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" "v0.29.1" >}}
