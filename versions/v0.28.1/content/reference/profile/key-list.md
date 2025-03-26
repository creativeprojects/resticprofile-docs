---
title: key-list
weight: 13
---
#### Section profile.**key-list**

This section configures restic command `key-list`  available since `0.17.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "list" sub-command lists all the keys (passwords) associated with the repository.
Returns the key ID, username, hostname, created time and if it's the current key being
used to access the repository.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |





{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

