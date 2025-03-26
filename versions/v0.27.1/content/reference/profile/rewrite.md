---
title: rewrite
weight: 24
---
#### Section profile.**rewrite**

This section configures restic command `rewrite`  available since `0.15.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "rewrite" command excludes files from existing snapshots. It creates new
snapshots containing the same data as the original ones, but without the files
you specify to exclude. All metadata (time, host, tags) will be preserved.

The snapshots to rewrite are specified using the --host, --tag and --path options,
or by providing a list of snapshot IDs. Please note that specifying neither any of
these options nor a snapshot ID will cause the command to rewrite all snapshots.

The special tag 'rewrite' will be added to the new snapshots to distinguish
them from the original ones, unless --forget is used. If the --forget option is
used, the original snapshots will instead be directly removed from the repository.

Please note that the --forget option only removes the snapshots and not the actual
data stored in the repository. In order to delete the no longer referenced data,
use the "prune" command.

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
| **dry-run** |`true` / `false` |`false` |do not do anything, just print what would be done. `restic >= 0.15.0`  |
| **exclude** |one or more `strings` | |exclude a pattern. `restic >= 0.15.0`  |
| **exclude-file** |one or more `strings` | |read exclude patterns from a file. `restic >= 0.15.0`  |
| **forget** |`true` / `false` |`false` |remove original snapshots after creating new ones. `restic >= 0.15.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.15.0`  |
| **iexclude** |one or more `strings` | |same as --exclude pattern but ignores the casing of filenames. `restic >= 0.15.0`  |
| **iexclude-file** |one or more `strings` | |same as --exclude-file but ignores casing of filenames in patterns. `restic >= 0.15.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.15.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.15.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" "v0.29.1" >}}
