---
title: stats
weight: 31
---
#### Section profile.**stats**

This section configures restic command `stats` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "stats" command walks one or multiple snapshots in a repository
and accumulates statistics about the data stored therein. It reports
on the number of unique files and their sizes, according to one of
the counting modes as given by the --mode flag.

It operates on all snapshots matching the selection criteria or all
snapshots if nothing is specified. The special snapshot ID "latest"
is also supported. Some modes make more sense over
just a single snapshot, while others are useful across all snapshots,
depending on what you are trying to calculate.

The modes are:

* restore-size: (default) Counts the size of the restored files.
* files-by-contents: Counts total size of unique files, where a file is
considered unique if it has unique contents.
* raw-data: Counts the size of blobs in the repository, regardless of
how many files reference them.
* blobs-per-file: A combination of files-by-contents and raw-data.


Refer to the online manual for more details about each mode.

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
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **mode** |`string` |`"restore-size"` |counting mode: restore-size (default), files-by-contents, blobs-per-file or raw-data |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.10.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.10.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

