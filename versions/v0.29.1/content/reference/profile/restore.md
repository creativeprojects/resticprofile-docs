---
title: restore
weight: 27
---
#### Section profile.**restore**

This section configures restic command `restore` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "restore" command extracts the data from a snapshot from the repository to
a directory.

The special snapshotID "latest" can be used to restore the latest snapshot in the
repository.

To only restore a specific subfolder, you can use the "snapshotID:subfolder"
syntax, where "subfolder" is a path within the snapshot.

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
| **delete** |`true` / `false` |`false` |delete files from target directory if they do not exist in snapshot. Use '--dry-run -vv' to check what would be deleted. `restic >= 0.17.0`  |
| **dry-run** |`true` / `false` |`false` |do not write any data, just show what would be done. `restic >= 0.17.0`  |
| **exclude** |one or more `strings` | |exclude a pattern |
| **exclude-file** |one or more `strings` | |read exclude patterns from a file. `restic >= 0.17.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **iexclude** |one or more `strings` | |same as --exclude pattern but ignores the casing of filenames. `restic >= 0.10.0`  |
| **iexclude-file** |one or more `strings` | |same as --exclude-file but ignores casing of filenames in patterns. `restic >= 0.17.0`  |
| **iinclude** |one or more `strings` | |same as --include pattern but ignores the casing of filenames. `restic >= 0.10.0`  |
| **iinclude-file** |one or more `strings` | |same as --include-file but ignores casing of filenames in patterns. `restic >= 0.17.0`  |
| **include** |one or more `strings` | |include a pattern |
| **include-file** |one or more `strings` | |read include patterns from a file. `restic >= 0.17.0`  |
| **overwrite** |`string` |`always` |overwrite behavior, one of (always/if-changed/if-newer/never) (default: always). `restic >= 0.17.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **sparse** |`true` / `false` |`false` |restore files as sparse. `restic >= 0.15.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **target** |`string` |`""` |directory to extract data to |
| **verify** |`true` / `false` |`false` |verify restored files content |




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.28.1" >}}
