---
title: repair-snapshots
weight: 26
---
#### Section profile.**repair-snapshots**

This section configures restic command `repair-snapshots`  available since `0.16.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "repair snapshots" command repairs broken snapshots. It scans the given
snapshots and generates new ones with damaged directories and file contents
removed. If the broken snapshots are deleted, a prune run will be able to
clean up the repository.

The command depends on a correct index, thus make sure to run "repair index"
first!

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
| **dry-run** |`true` / `false` |`false` |do not do anything, just print what would be done. `restic >= 0.16.0`  |
| **forget** |`true` / `false` |`false` |remove original snapshots after creating new ones. `restic >= 0.16.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.16.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.16.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.16.0`  |




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.28.1" >}}
