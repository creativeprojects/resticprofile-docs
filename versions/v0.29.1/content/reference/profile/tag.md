---
title: tag
weight: 32
---
#### Section profile.**tag**

This section configures restic command `tag` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "tag" command allows you to modify tags on exiting snapshots.

You can either set/replace the entire set of tags on a snapshot, or
add tags to/remove tags from the existing set.

When no snapshot-ID is given, all snapshots matching the host, tag and path filter criteria are modified.

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
| **add** |one or more `strings` | |tags which will be added to the existing tags in the format tag[,tag,...] |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **remove** |one or more `strings` | |tags which will be removed from the existing tags in the format tag[,tag,...] |
| **set** |one or more `strings` | |tags which will replace the existing tags in the format tag[,tag,...] |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.28.1" >}}
