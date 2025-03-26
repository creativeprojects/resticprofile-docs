---
title: key
weight: 11
---
#### Section profile.**key**

This section configures restic command `key` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "key" command manages keys (passwords) for accessing the repository.

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
| **host** |`true` / `false` OR `hostname` |`""` |the hostname for new keys. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.10.0`  |
| **new-password-file** |`string` |`""` |file from which to read the new password |
| **user** |`string` |`""` |the username for new keys. `restic >= 0.10.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" "v0.29.1" >}}
