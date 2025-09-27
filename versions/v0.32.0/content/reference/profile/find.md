---
title: find
weight: 9
---
#### Section profile.**find**

This section configures restic command `find` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "find" command searches for files or directories in snapshots stored in the
repo.
It can also be used to search for restic blobs or trees for troubleshooting.
The default sort option for the snapshots is youngest to oldest. To sort the
output from oldest to youngest specify --reverse.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |
| **send-after** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **blob** |`true` / `false` |`false` |pattern is a blob-ID |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **human-readable** |`true` / `false` |`false` |print sizes in human readable format. `restic >= 0.16.0`  |
| **ignore-case** |`true` / `false` |`false` |ignore case for pattern |
| **long** |`true` / `false` |`false` |use a long listing format showing size and mode |
| **newest** |`string` |`""` |newest modification date/time |
| **oldest** |`string` |`""` |oldest modification date/time |
| **pack** |`true` / `false` |`false` |pattern is a pack-ID |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **reverse** |`true` / `false` |`false` |reverse sort order oldest to newest. `restic >= 0.18.0`  |
| **show-pack-id** |`true` / `false` |`false` |display the pack-ID the blobs belong to (with --blob or --tree) |
| **snapshot** |one or more `strings` | |snapshot id to search in |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **tree** |`true` / `false` |`false` |pattern is a tree-ID |




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

