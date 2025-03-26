---
title: diff
weight: 6
---
#### Section profile.**diff**

This section configures restic command `diff` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "diff" command shows differences from the first to the second snapshot. The
first characters in each line display what has happened to a particular file or
directory:

+  The item was added
-  The item was removed
U  The metadata (access mode, timestamps, ...) for the item was updated
M  The file's content was modified
T  The type was changed, e.g. a file was made a symlink

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
| **metadata** |`true` / `false` |`false` |print changes in metadata |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

