---
title: key-remove
weight: 15
---
#### Section profile.**key-remove**

This section configures restic command `key-remove`  available since `0.17.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "remove" sub-command removes the selected key ID. The "remove" command does not allow
removing the current key being used to access the repository.

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
