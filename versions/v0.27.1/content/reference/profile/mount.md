---
title: mount
weight: 15
---
#### Section profile.**mount**

This section configures restic command `mount` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "mount" command mounts the repository via fuse to a directory. This is a
read-only mount.


If you need a different template for directories that contain snapshots,
you can pass a time template via --time-template and path templates via
--path-template.

Example time template without colons:


--time-template "2006-01-02_15-04-05"


You need to specify a sample format for exactly the following timestamp:


Mon Jan 2 15:04:05 -0700 MST 2006


For details please see the documentation for time.Format() at:
  https://godoc.org/time#Time.Format

For path templates, you can use the following patterns which will be replaced:
    %i by short snapshot ID
    %I by long snapshot ID
    %u by username
    %h by hostname
    %t by tags
    %T by timestamp as specified by --time-template

The default path templates are:
    "ids/%i"
    "snapshots/%T"
    "hosts/%h/%T"
    "tags/%t/%T"

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
| **allow-other** |`true` / `false` |`false` |allow other users to access the data in the mounted directory |
| ~~allow-root~~ |`true` / `false` |`false` |allow root user to access the data in the mounted directory. `restic < 0.10.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **no-default-permissions** |`true` / `false` |`false` |for 'allow-other', ignore Unix permissions and allow users to read all snapshot files |
| **owner-root** |`true` / `false` |`false` |use 'root' as the owner of files and dirs |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **path-template** |one or more `strings` | |set template for path names. `restic >= 0.14.0`  |
| ~~snapshot-template~~ |`string` |`"2006-01-02T15:04:05Z07:00"` |set template to use for snapshot dirs. `restic < 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **time-template** |`string` |`"2006-01-02T15:04:05Z07:00"` |set template to use for times. `restic >= 0.14.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" "v0.29.1" >}}
