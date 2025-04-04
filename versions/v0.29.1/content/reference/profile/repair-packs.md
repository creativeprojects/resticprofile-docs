---
title: repair-packs
weight: 25
---
#### Section profile.**repair-packs**

This section configures restic command `repair-packs`  available since `0.16.1` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

WARNING: The CLI for this command is experimental and will likely change in the future!

The "repair packs" command extracts intact blobs from the specified pack files, rebuilds
the index to remove the damaged pack files and removes the pack files from the repository.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |





{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" >}}
