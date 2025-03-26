---
alwaysopen: false
title: Section Profile
weight: 5
---
### Profile sections

In config file format v1, the name of this section is the name of your profile
(excluding reserved names like `global`, `groups`, `includes` & `version`).

From config file format v2, profile sections are declared as named configuration
structure below section `profiles` (see [Configuration v2]({{% relref "/configuration/v2/" %}})
for details).

#### Section **profile**

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **base-dir** |`string` | |Sets the working directory for this profile. The profile will fail when the working directory cannot be changed. Leave empty to use the current directory instead |
| **description** |`string` | |Describes the profile |
| **env** |`key` => `value` | |Additional environment variables to set in any child process. Inline env variables take precedence over dotenv files declared with "env-file" |
| **env-file** |one or more `strings` | |Additional dotenv files to load and set as environment in any child process |
| **force-inactive-lock** |`true` / `false` |`false` |Allows to lock when the existing lock is considered stale |
| **inherit** |`string` | |Name of the profile to inherit all of the settings from |
| **initialize** |`true` / `false` | |Initialize the restic repository if missing |
| **lock** |`string` | |Path to the lock file to use with resticprofile locks |
| **prometheus-labels** |`key` => `string` | |Additional prometheus labels to set |
| **prometheus-push** |`uri` | |URL of the prometheus push gateway to send the summary of the last restic command result to |
| **prometheus-push-format** |`string` |`text` |Prometheus push gateway request format. Is one of `text`, `protobuf`  |
| **prometheus-push-job** |`string` | |Prometheus push gateway job name. $command placeholder is replaced with restic command |
| **prometheus-save-to-file** |`string` | |Path to the prometheus metrics file to update with a summary of the last restic command result |
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |
| **status-file** |`string` | |Path to the status file to update with a summary of last restic command result |
| **stream-error** | one or more nested *[StreamErrorSection](../nested/streamerrorsection)* | |Run shell command(s) when a pattern matches the stderr of restic |
| **systemd-drop-in-files** |one or more `strings` | |Files containing systemd drop-in (override) files - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **cacert** |`string` | |file to load root certificates from (default: use system certificates or $RESTIC_CACERT) |
| **cache-dir** |`string` |`""` |set the cache directory. (default: use system default cache directory) |
| **cleanup-cache** |`true` / `false` |`false` |auto remove old cache directories |
| **compression** |`string` |`auto` |compression mode (only available for repository format version 2), one of (auto/off/max) (default: $RESTIC_COMPRESSION). `restic >= 0.14.0`  |
| **extended-status** |`true` / `false` |`false` |set output mode to JSON for commands that support it |
| **http-user-agent** |`string` |`""` |set a http user agent for outgoing http requests. `restic >= 0.17.0`  |
| **insecure-no-password** |`true` / `false` |`false` |use an empty password for the repository, must be passed to every restic command (insecure). `restic >= 0.17.0`  |
| **insecure-tls** |`true` / `false` |`false` |skip TLS certificate verification when connecting to the repository (insecure). `restic >= 0.13.0`  |
| **key-hint** |`string` |`""` |key ID of key to try decrypting first (default: $RESTIC_KEY_HINT) |
| **limit-download** |`integer` |`0` |limits downloads to a maximum rate in KiB/s. (default: unlimited) |
| **limit-upload** |`integer` |`0` |limits uploads to a maximum rate in KiB/s. (default: unlimited) |
| **no-cache** |`true` / `false` |`false` |do not use a local cache |
| **no-extra-verify** |`true` / `false` |`false` |skip additional verification of data before upload (see documentation). `restic >= 0.16.4`  |
| **no-lock** |`true` / `false` |`false` |do not lock the repository, this allows some operations on read-only repositories |
| **option** |one or more `strings` | |set extended option (key=value) |
| **pack-size** |`integer` |`0` |set target pack size in MiB, created pack files may be larger (default: $RESTIC_PACK_SIZE). `restic >= 0.14.0`  |
| **password-command** |`string` |`""` |shell command to obtain the repository password from (default: $RESTIC_PASSWORD_COMMAND) |
| **password-file** |`string` |`""` |file to read the repository password from (default: $RESTIC_PASSWORD_FILE) |
| **quiet** |`true` / `false` |`false` |do not output comprehensive progress report |
| **repository** |`string` |`""` |repository to backup to or restore from (default: $RESTIC_REPOSITORY) |
| **repository-file** |`string` |`""` |file to read the repository location from (default: $RESTIC_REPOSITORY_FILE). `restic >= 0.11.0`  |
| **retry-lock** |`string` |`0s` |retry to lock the repository if it is already locked, takes a value like 5m or 2h (default: no retries). `restic >= 0.16.0`  |
| **tls-client-cert** |`string` |`""` |path to a file containing PEM encoded TLS client certificate and private key (default: $RESTIC_TLS_CLIENT_CERT) |
| **verbose** |`true` / `false` OR `integer` |`0` |be verbose (true for level 1 or a number for increased verbosity, max level is 2) |




{{% notice style="note" %}}
Most **restic** command flags defined in profile sections below can also be set at profile level.
They will be inherited in all sections that define these flags and ignored in all others.
{{% /notice %}}

{{< pageversions "v0.27.1" "v0.28.1" >}}
