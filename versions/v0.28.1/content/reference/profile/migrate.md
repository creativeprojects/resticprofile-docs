---
title: migrate
weight: 18
---
#### Section profile.**migrate**

This section configures restic command `migrate` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "migrate" command checks which migrations can be applied for a repository
and prints a list with available migration names. If one or more migration
names are specified, these migrations are applied.

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
| **force** |`true` / `false` |`false` |apply a migration a second time |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

