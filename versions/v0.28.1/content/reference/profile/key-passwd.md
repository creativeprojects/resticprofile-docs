---
title: key-passwd
weight: 14
---
#### Section profile.**key-passwd**

This section configures restic command `key-passwd`  available since `0.17.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "passwd" sub-command creates a new key, validates the key and remove the old key ID.
Returns the new key ID.

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
| **host** |`true` / `false` OR `hostname` |`""` |the hostname for new key. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.17.0`  |
| **new-insecure-no-password** |`true` / `false` |`false` |add an empty password for the repository (insecure). `restic >= 0.17.0`  |
| **new-password-file** |`string` |`""` |file from which to read the new password. `restic >= 0.17.0`  |
| **user** |`string` |`""` |the username for new key. `restic >= 0.17.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

