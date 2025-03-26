---
title: init
weight: 10
---
#### Section profile.**init**

This section configures restic command `init` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "init" command initializes a new repository.


##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **copy-chunker-params** |`true` / `false` |`false` |copy chunker parameters from the secondary repository (useful with the copy command). `restic >= 0.10.0`  |
| **from-key-hint** |`string` |`""` |key ID of key to try decrypting the source repository first (default: $RESTIC_FROM_KEY_HINT). `restic >= 0.14.0`  |
| **from-password-command** |`string` |`""` |shell command to obtain the source repository password from (default: $RESTIC_FROM_PASSWORD_COMMAND). `restic >= 0.14.0`  |
| **from-password-file** |`string` |`""` |file to read the source repository password from (default: $RESTIC_FROM_PASSWORD_FILE). `restic >= 0.14.0`  |
| **from-repository** |`string` |`""` |source repository to copy chunker parameters from (default: $RESTIC_FROM_REPOSITORY). `restic >= 0.14.0`  |
| **from-repository-file** |`string` |`""` |file from which to read the source repository location to copy chunker parameters from (default: $RESTIC_FROM_REPOSITORY_FILE). `restic >= 0.14.0`  |
| ~~key-hint2~~ |`string` |`""` |key ID of key to try decrypting the secondary repository first (default: $RESTIC_KEY_HINT2). `restic >= 0.10.0 < 0.14.0`  |
| ~~password-command2~~ |`string` |`""` |shell command to obtain the secondary repository password from (default: $RESTIC_PASSWORD_COMMAND2). `restic >= 0.10.0 < 0.14.0`  |
| ~~password-file2~~ |`string` |`""` |file to read the secondary repository password from (default: $RESTIC_PASSWORD_FILE2). `restic >= 0.10.0 < 0.14.0`  |
| ~~repo2~~ |`string` |`""` |secondary repository to copy chunker parameters from (default: $RESTIC_REPOSITORY2). `restic >= 0.10.0 < 0.14.0`  |
| ~~repository-file2~~ |`string` |`""` |file from which to read the secondary repository location to copy chunker parameters from (default: $RESTIC_REPOSITORY_FILE2). `restic >= 0.13.0 < 0.14.0`  |
| **repository-version** |`string` |`"stable"` |repository format version to use, allowed values are a format version, 'latest' and 'stable'. `restic >= 0.14.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.28.1" "v0.29.1" >}}
