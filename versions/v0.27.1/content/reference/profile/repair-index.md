---
title: repair-index
weight: 20
---
#### Section profile.**repair-index**

This section configures restic command `repair-index`  available since `0.16.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "repair index" command creates a new index based on the pack files in the
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
| **read-all-packs** |`true` / `false` |`false` |read all pack files to generate new index from scratch. `restic >= 0.16.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

