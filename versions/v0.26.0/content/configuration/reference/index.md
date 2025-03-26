---
title: Reference
weight: 50
---













{{% notice note %}}
The configuration file reference is generated from resticprofile's data model and restic's manual pages.
{{% /notice %}}


* [Sections](#sections)
  * [Section <strong>global</strong>](#section-global)
  * [Profile sections](#profile-sections)
    * [Section <strong>profile</strong>](#section-profile)
    * [Section profile\.<strong>backup</strong>](#section-profilebackup)
    * [Section profile\.<strong>cache</strong>](#section-profilecache)
    * [Section profile\.<strong>cat</strong>](#section-profilecat)
    * [Section profile\.<strong>check</strong>](#section-profilecheck)
    * [Section profile\.<strong>copy</strong>](#section-profilecopy)
    * [Section profile\.<strong>diff</strong>](#section-profilediff)
    * [Section profile\.<strong>dump</strong>](#section-profiledump)
    * [Section profile\.<strong>find</strong>](#section-profilefind)
    * [Section profile\.<strong>forget</strong>](#section-profileforget)
    * [Section profile\.<strong>init</strong>](#section-profileinit)
    * [Section profile\.<strong>key</strong>](#section-profilekey)
    * [Section profile\.<strong>list</strong>](#section-profilelist)
    * [Section profile\.<strong>ls</strong>](#section-profilels)
    * [Section profile\.<strong>migrate</strong>](#section-profilemigrate)
    * [Section profile\.<strong>mount</strong>](#section-profilemount)
    * [Section profile\.<strong>prune</strong>](#section-profileprune)
    * [Section profile\.<strong>rebuild-index</strong>](#section-profilerebuild-index)
    * [Section profile\.<strong>recover</strong>](#section-profilerecover)
    * [Section profile\.<strong>repair</strong>](#section-profilerepair)
    * [Section profile\.<strong>repair-index</strong>](#section-profilerepair-index)
    * [Section profile\.<strong>repair-snapshots</strong>](#section-profilerepair-snapshots)
    * [Section profile\.<strong>restore</strong>](#section-profilerestore)
    * [Section profile\.<strong>retention</strong>](#section-profileretention)
    * [Section profile\.<strong>rewrite</strong>](#section-profilerewrite)
    * [Section profile\.<strong>snapshots</strong>](#section-profilesnapshots)
    * [Section profile\.<strong>stats</strong>](#section-profilestats)
    * [Section profile\.<strong>tag</strong>](#section-profiletag)
    * [Section profile\.<strong>unlock</strong>](#section-profileunlock)
  * [Nested profile sections](#nested-profile-sections)
    * [Nested <em>SendMonitoringHeader</em>](#nested-sendmonitoringheader)
    * [Nested <em>SendMonitoringSection</em>](#nested-sendmonitoringsection)
    * [Nested <em>StreamErrorSection</em>](#nested-streamerrorsection)
  * [Section <strong>groups</strong>](#section-groups)
* [Value types](#value-types)
* [JSON schema](#json-schema)

## Sections

### Section **global**

The `global` section is at the root of the configuration file and contains the global
settings for resticprofile.

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **ca-certificates** |one or more `strings` | |Path to PEM encoded certificates to trust in addition to system certificates when resticprofile sends to a webhook - see [configuration/http_hooks/]({{% relref "/configuration/http_hooks/" %}}) |
| **default-command** |`string` |`snapshots` |The restic or resticprofile command to use when no command was specified. **Examples**: `backup`, `cache`, `cat`, `check`, `copy`, `diff`, `dump`, `find`, `forget`, `generate`, `init`, `key`, `list`, `ls`, `migrate`, `mount`, `prune`, `rebuild-index`, `recover`, `repair`, `repair-index`, `repair-snapshots`, `restore`, `rewrite`, `self-update`, `snapshots`, `stats`, `tag`, `unlock`, `version`  |
| **group-continue-on-error** |`true` / `false` |`false` |Enable groups to continue with the next profile(s) instead of stopping at the first failure |
| **initialize** |`true` / `false` |`false` |Initialize a repository if missing |
| **ionice** |`true` / `false` |`false` |Enables setting the unix IO priority class and level for resticprofile and child processes (only on unix OS) |
| **ionice-class** |`integer` |`2` |Sets the unix "ionice-class" to apply when "ionice" is enabled. Must be >= 1 and  <= 3  |
| **ionice-level** |`integer` |`0` |Sets the unix "ionice-level" to apply when "ionice" is enabled. Must be >= 0 and  <= 7  |
| ~~legacy-arguments~~ |`true` / `false` |`false` |Legacy, broken arguments mode of resticprofile before version 0.15 |
| **log** |`string` | |Sets the default log destination to be used if not specified in '--log' or 'schedule-log' - see [configuration/logs/]({{% relref "/configuration/logs/" %}}) |
| **min-memory** |`integer` |`100` |Minimum available memory (in MB) required to run any commands - see [usage/memory/]({{% relref "/usage/memory/" %}}) |
| **nice** |`integer` |`0` |Sets the unix "nice" value for resticprofile and child processes (on any OS). Must be >= -20 and  <= 19  |
| **prevent-sleep** |`true` / `false` |`false` |Prevent the system from sleeping while running commands - see [configuration/sleep/]({{% relref "/configuration/sleep/" %}}) |
| **priority** |`string` |`normal` |Sets process priority class for resticprofile and child processes (on any OS). Is one of `idle`, `background`, `low`, `normal`, `high`, `highest`  |
| **restic-arguments-filter** |`true` / `false` |`true` |Remove unknown flags instead of passing all configured flags to restic |
| **restic-binary** |`string` | |Full path of the restic executable (detected if not set) |
| **restic-lock-retry-after** |`integer` OR `duration` |`1m` |Time to wait before trying to get a lock on a restic repositoey - see [usage/locks/]({{% relref "/usage/locks/" %}}) |
| **restic-stale-lock-age** |`integer` OR `duration` |`1h` |The age an unused lock on a restic repository must have at least before resiticprofile attempts to unlock - see [usage/locks/]({{% relref "/usage/locks/" %}}) |
| **scheduler** |`string` | |Leave blank for the default scheduler or use "crond" to select cron on supported operating systems |
| **send-timeout** |`integer` OR `duration` |`30s` |Timeout when sending messages to a webhook - see [configuration/http_hooks/]({{% relref "/configuration/http_hooks/" %}}). **Examples**: `15s`, `30s`, `2m30s`  |
| **shell** |one or more `strings` |`auto` |The shell that is used to run commands (default is OS specific). **Examples**: `sh`, `bash`, `pwsh`, `powershell`, `cmd`  |
| **systemd-timer-template** |`string` | |File containing the go template to generate a systemd timer - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |
| **systemd-unit-template** |`string` | |File containing the go template to generate a systemd unit - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |




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
| **env** |`key` => `value` | |Additional environment variables to set in any child process |
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
| **stream-error** | one or more nested *[StreamErrorSection](#nested-streamerrorsection)* | |Run shell command(s) when a pattern matches the stderr of restic |
| **systemd-drop-in-files** |one or more `strings` | |Files containing systemd drop-in (override) files - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **cacert** |`string` | |file to load root certificates from (default: use system certificates or $RESTIC_CACERT) |
| **cache-dir** |`string` |`""` |set the cache directory&. (default: use system default cache directory) |
| **cleanup-cache** |`true` / `false` |`false` |auto remove old cache directories |
| **compression** |`string` |`auto` |compression mode (only available for repository format version 2), one of (auto/off/max) (default: $RESTIC_COMPRESSION). `restic >= 0.14.0`  |
| **extended-status** |`true` / `false` |`false` |set output mode to JSON for commands that support it |
| **insecure-tls** |`true` / `false` |`false` |skip TLS certificate verification when connecting to the repository (insecure). `restic >= 0.13.0`  |
| **key-hint** |`string` |`""` |key ID of key to try decrypting first (default: $RESTIC_KEY_HINT) |
| **limit-download** |`integer` |`0` |limits downloads to a maximum rate in KiB/s. (default: unlimited) |
| **limit-upload** |`integer` |`0` |limits uploads to a maximum rate in KiB/s. (default: unlimited) |
| **no-cache** |`true` / `false` |`false` |do not use a local cache |
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




{{% notice note %}}
Most **restic** command flags defined in profile sections below can also be set at profile level.
They will be inherited in all sections that define these flags and ignored in all others.
{{% /notice %}}

[go to top](#reference)

#### Section profile.**backup**

This section configures restic command `backup` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "backup" command creates a new snapshot and saves the files and directories
given as the arguments.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **check-after** |`true` / `false` |`false` |Check the repository after the backup command succeeded |
| **check-before** |`true` / `false` |`false` |Check the repository before starting the backup command |
| **no-error-on-warning** |`true` / `false` |`false` |Do not fail the backup when some files could not be read |
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |
| **schedule** |one or more `strings` | |Set the times at which the scheduled command is run (times are specified in systemd timer format). **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` |`false` |Don't schedule the start of this profile when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't schedule the start of this profile when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't schedule the start of this profile when running on battery, and the battery charge left is less than the value |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `tcp://localhost:514`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |
| **source** |one or more `strings` | |The paths to backup. **Examples**: `/opt/`, `/home/user/`, `C:\Users\User\Documents`  |
| **source-base** |`string` | |The base path to resolve relative backup paths against. Defaults to current directory if unset or empty (see also "base-dir" in profile). **Examples**: `/`, `$PWD`, `C:\`, `%cd%`  |
| **stdin-command** |one or more `strings` | |Shell command(s) that generate content to redirect into the stdin of restic. When set, the flag "stdin" is always set to "true" |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **dry-run** |`true` / `false` |`false` |do not upload or write any data, just show what would be done. `restic >= 0.13.0`  |
| **exclude** |one or more `strings` | |exclude a pattern |
| **exclude-caches** |`true` / `false` |`false` |excludes cache directories that are marked with a CACHEDIR.TAG file. See https://bford.info/cachedir/ for the Cache Directory Tagging Standard |
| **exclude-file** |one or more `strings` | |read exclude patterns from a file |
| **exclude-if-present** |one or more `strings` | |takes filename[:header], exclude contents of directories containing filename (except filename itself) if header of that file is as provided |
| **exclude-larger-than** |`string` |`""` |max size of the files to be backed up (allowed suffixes: k/K, m/M, g/G, t/T). `restic >= 0.10.0`  |
| **extended-status** |`true` / `false` |`false` |set output mode to JSON for commands that support it |
| **files-from** |one or more `strings` | |read the files to backup from file (can be combined with file args) |
| **files-from-raw** |one or more `strings` | |read the files to backup from file (can be combined with file args). `restic >= 0.12.0`  |
| **files-from-verbatim** |one or more `strings` | |read the files to backup from file (can be combined with file args). `restic >= 0.12.0`  |
| **force** |`true` / `false` |`false` |force re-reading the target files/directories (overrides the "parent" flag) |
| **group-by** |`string` |`host,paths` |group snapshots by host, paths and/or tags, separated by comma (disable grouping with ''). `restic >= 0.16.0`  |
| **host** |`true` / `false` OR `hostname` |`""` |set the hostname for the snapshot manually. To prevent an expensive rescan use the "parent" flag. Boolean true is replaced with the hostname of the system. Defaults to true for config version 2 in "backup". **Examples**: `true`, `false`, `"host"`  |
| **iexclude** |one or more `strings` | |same as --exclude pattern but ignores the casing of filenames. `restic >= 0.10.0`  |
| **iexclude-file** |one or more `strings` | |same as --exclude-file but ignores casing of filenames in patterns. `restic >= 0.10.0`  |
| **ignore-ctime** |`true` / `false` |`false` |ignore ctime changes when checking for modified files. `restic >= 0.12.0`  |
| **ignore-inode** |`true` / `false` |`false` |ignore inode number changes when checking for modified files. `restic >= 0.10.0`  |
| **no-scan** |`true` / `false` |`false` |do not run scanner to estimate size of backup. `restic >= 0.15.0`  |
| **one-file-system** |`true` / `false` |`false` |exclude other file systems, don't cross filesystem boundaries and subvolumes |
| **parent** |`string` |`""` |use this parent snapshot (default: latest snapshot in the group determined by --group-by and not newer than the timestamp determined by --time) |
| **read-concurrency** |`integer` |`0` |read n files concurrently (default: $RESTIC_READ_CONCURRENCY or 2). `restic >= 0.15.0`  |
| **stdin** |`true` / `false` |`false` |read backup from stdin |
| **stdin-filename** |`string` |`"stdin"` |filename to use when reading from stdin |
| **tag** |`true` / `false` OR one or more `strings` | |add tags for the new snapshot in the format tag[,tag,...]. Boolean true is unsupported in section "backup". **Examples**: `false`, `"tag"`  |
| **time** |`string` |`""` |time of the backup (ex. '2012-11-01 22:08:41') (default: now) |
| **use-fs-snapshot** |`true` / `false` |`false` |use filesystem snapshot where possible (currently only Windows VSS). `restic >= 0.12.0` . Available only for `windows`  |
| **with-atime** |`true` / `false` |`false` |store the atime for all files and directories |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**cache**

This section configures restic command `cache` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "cache" command allows listing and cleaning local cache directories.

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
| **cleanup** |`true` / `false` |`false` |remove old cache directories |
| **max-age** |`integer` |`30` |max age in days for cache directories to be considered old |
| **no-size** |`true` / `false` |`false` |do not output the size of the cache directories |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**cat**

This section configures restic command `cat` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "cat" command is used to print internal objects to stdout.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**check**

This section configures restic command `check` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "check" command tests the repository for errors and reports any errors it
finds. It can also be used to read all data and therefore simulate a restore.

By default, the "check" command will always load all data directly from the
repository and not use a local cache.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **schedule** |one or more `strings` | |Set the times at which the scheduled command is run (times are specified in systemd timer format). **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` |`false` |Don't schedule the start of this profile when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't schedule the start of this profile when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't schedule the start of this profile when running on battery, and the battery charge left is less than the value |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `tcp://localhost:514`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| ~~check-unused~~ |`true` / `false` |`false` |find unused blobs. `restic < 0.14.0`  |
| **read-data** |`true` / `false` |`false` |read all data blobs |
| **read-data-subset** |`string` |`""` |read a subset of data packs, specified as 'n/t' for specific part, or either 'x%' or 'x.y%' or a size in bytes with suffixes k/K, m/M, g/G, t/T for a random subset |
| **with-cache** |`true` / `false` |`false` |use existing cache, only read uncached data from repository |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**copy**

This section configures restic command `copy`  available since `0.10.0` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "copy" command copies one or more snapshots from one repository to another.

NOTE: This process will have to both download (read) and upload (write) the
entire snapshot(s) due to the different encryption keys used in the source and
destination repositories. This /may incur higher bandwidth usage and costs/ than
expected during normal backup runs.

NOTE: The copying process does not re-chunk files, which may break deduplication
between the files copied and files already stored in the destination repository.
This means that copied files, which existed in both the source and destination
repository, /may occupy up to twice their space/ in the destination repository.
This can be mitigated by the "--copy-chunker-params" option when initializing a
new destination repository using the "init" command.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **initialize** |`true` / `false` |`false` |Initialize the secondary repository if missing |
| **initialize-copy-chunker-params** |`true` / `false` |`true` |Copy chunker parameters when initializing the secondary repository |
| **key-hint** |`string` | |Key ID of key to try decrypting the destination repository first |
| **password-command** |`string` | |Shell command to obtain the destination repository password from |
| **password-file** |`string` | |File to read the destination repository password from |
| **repository** |`string` | |Destination repository to copy snapshots to |
| **repository-file** |`string` | |File from which to read the destination repository location to copy snapshots to |
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |
| **schedule** |one or more `strings` | |Set the times at which the scheduled command is run (times are specified in systemd timer format). **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` |`false` |Don't schedule the start of this profile when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't schedule the start of this profile when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't schedule the start of this profile when running on battery, and the battery charge left is less than the value |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `tcp://localhost:514`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |
| **snapshot** |one or more `strings` | |Snapshot IDs to copy (if empty, all snapshots are copied) |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **from-key-hint** |`string` |`""` |key ID of key to try decrypting the source repository first (default: $RESTIC_FROM_KEY_HINT). `restic >= 0.14.0`  |
| **from-password-command** |`string` |`""` |shell command to obtain the source repository password from (default: $RESTIC_FROM_PASSWORD_COMMAND). `restic >= 0.14.0`  |
| **from-password-file** |`string` |`""` |file to read the source repository password from (default: $RESTIC_FROM_PASSWORD_FILE). `restic >= 0.14.0`  |
| **from-repository** |`string` |`""` |source repository to copy snapshots from (default: $RESTIC_FROM_REPOSITORY). `restic >= 0.14.0`  |
| **from-repository-file** |`string` |`""` |file from which to read the source repository location to copy snapshots from (default: $RESTIC_FROM_REPOSITORY_FILE). `restic >= 0.14.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.10.0`  |
| ~~key-hint2~~ |`string` |`""` |key ID of key to try decrypting the destination repository first (default: $RESTIC_KEY_HINT2). `restic >= 0.10.0 < 0.14.0`  |
| ~~password-command2~~ |`string` |`""` |shell command to obtain the destination repository password from (default: $RESTIC_PASSWORD_COMMAND2). `restic >= 0.10.0 < 0.14.0`  |
| ~~password-file2~~ |`string` |`""` |file to read the destination repository password from (default: $RESTIC_PASSWORD_FILE2). `restic >= 0.10.0 < 0.14.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.10.0`  |
| ~~repo2~~ |`string` |`""` |destination repository to copy snapshots to (default: $RESTIC_REPOSITORY2). `restic >= 0.10.0 < 0.14.0`  |
| ~~repository-file2~~ |`string` |`""` |file from which to read the destination repository location to copy snapshots to (default: $RESTIC_REPOSITORY_FILE2). `restic >= 0.13.0 < 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.10.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
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
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**dump**

This section configures restic command `dump` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "dump" command extracts files from a snapshot from the repository. If a
single file is selected, it prints its contents to stdout. Folders are output
as a tar (default) or zip file containing the contents of the specified folder.
Pass "/" as file name to dump the whole snapshot as an archive file.

The special snapshot "latest" can be used to use the latest snapshot in the
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
| **archive** |`string` |`"tar"` |set archive format as "tar" or "zip". `restic >= 0.12.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**find**

This section configures restic command `find` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "find" command searches for files or directories in snapshots stored in the
repo.
It can also be used to search for restic blobs or trees for troubleshooting.

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
| **blob** |`true` / `false` |`false` |pattern is a blob-ID |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **human-readable** |`true` / `false` |`false` |print sizes in human readable format. `restic >= 0.16.0`  |
| **ignore-case** |`true` / `false` |`false` |ignore case for pattern |
| **long** |`true` / `false` |`false` |use a long listing format showing size and mode |
| **newest** |`string` |`""` |newest modification date/time |
| **oldest** |`string` |`""` |oldest modification date/time |
| **pack** |`true` / `false` |`false` |pattern is a pack-ID |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **show-pack-id** |`true` / `false` |`false` |display the pack-ID the blobs belong to (with --blob or --tree) |
| **snapshot** |one or more `strings` | |snapshot id to search in |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **tree** |`true` / `false` |`false` |pattern is a tree-ID |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**forget**

This section configures restic command `forget` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "forget" command removes snapshots according to a policy. All snapshots are
first divided into groups according to "--group-by", and after that the policy
specified by the "--keep-*" options is applied to each group individually.

Please note that this command really only deletes the snapshot object in the
repository, which is a reference to data stored there. In order to remove the
unreferenced data after "forget" was run successfully, see the "prune" command.

Please also read the documentation for "forget" to learn about some important
security considerations.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **schedule** |one or more `strings` | |Set the times at which the scheduled command is run (times are specified in systemd timer format). **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` |`false` |Don't schedule the start of this profile when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't schedule the start of this profile when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't schedule the start of this profile when running on battery, and the battery charge left is less than the value |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `tcp://localhost:514`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **compact** |`true` / `false` |`false` |use compact output format |
| **dry-run** |`true` / `false` |`false` |do not delete anything, just print what would be done |
| **group-by** |`string` |`host,paths` |group snapshots by host, paths and/or tags, separated by comma (disable grouping with '') |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **keep-daily** |`integer` |`0` |keep the last n daily snapshots (use 'unlimited' to keep all daily snapshots) |
| **keep-hourly** |`integer` |`0` |keep the last n hourly snapshots (use 'unlimited' to keep all hourly snapshots) |
| **keep-last** |`integer` |`0` |keep the last n snapshots (use 'unlimited' to keep all snapshots) |
| **keep-monthly** |`integer` |`0` |keep the last n monthly snapshots (use 'unlimited' to keep all monthly snapshots) |
| **keep-tag** |one or more `strings` | |keep snapshots with this taglist |
| **keep-weekly** |`integer` |`0` |keep the last n weekly snapshots (use 'unlimited' to keep all weekly snapshots) |
| **keep-within** |`string` | |keep snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot |
| **keep-within-daily** |`string` | |keep daily snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-hourly** |`string` | |keep hourly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-monthly** |`string` | |keep monthly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-weekly** |`string` | |keep weekly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-yearly** |`string` | |keep yearly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-yearly** |`integer` |`0` |keep the last n yearly snapshots (use 'unlimited' to keep all yearly snapshots) |
| **max-repack-size** |`string` |`""` |maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T). `restic >= 0.12.0`  |
| **max-unused** |`string` |`"5%"` |tolerate given limit of unused data (absolute value in bytes with suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited'). `restic >= 0.12.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **prune** |`true` / `false` |`false` |automatically run the 'prune' command if snapshots have been removed |
| **repack-cacheable-only** |`true` / `false` |`false` |only repack packs which are cacheable. `restic >= 0.12.0`  |
| **repack-small** |`true` / `false` |`false` |repack pack files below 80% of target pack size. `restic >= 0.14.0`  |
| **repack-uncompressed** |`true` / `false` |`false` |repack all uncompressed data. `restic >= 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
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
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
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
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**list**

This section configures restic command `list` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "list" command allows listing objects in the repository based on type.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**ls**

This section configures restic command `ls` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "ls" command lists files and directories in a snapshot.

The special snapshot ID "latest" can be used to list files and
directories of the latest snapshot in the repository. The
--host flag can be used in conjunction to select the latest
snapshot originating from a certain host only.

File listings can optionally be filtered by directories. Any
positional arguments after the snapshot ID are interpreted as
absolute directory paths, and only files inside those directories
will be listed. If the --recursive flag is used, then the filter
will allow traversing into matching directories' subfolders.
Any directory paths specified must be absolute (starting with
a path separator); paths use the forward slash '/' as separator.

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
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **human-readable** |`true` / `false` |`false` |print sizes in human readable format. `restic >= 0.16.0`  |
| **long** |`true` / `false` |`false` |use a long listing format showing size and mode |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **recursive** |`true` / `false` |`false` |include files in subfolders of the listed directories |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
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
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**mount**

This section configures restic command `mount` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "mount" command mounts the repository via fuse to a directory. This is a
read-only mount.


If you need a different template for directories that contain snapshots,
you can pass a time template via --time-template and path templates via
--path-template.

Example time template without colons:


--time-template "2006-01-02_15-04-05"


You need to specify a sample format for exactly the following timestamp:


Mon Jan 2 15:04:05 -0700 MST 2006


For details please see the documentation for time.Format() at:
  https://godoc.org/time#Time.Format

For path templates, you can use the following patterns which will be replaced:
    %i by short snapshot ID
    %I by long snapshot ID
    %u by username
    %h by hostname
    %t by tags
    %T by timestamp as specified by --time-template

The default path templates are:
    "ids/%i"
    "snapshots/%T"
    "hosts/%h/%T"
    "tags/%t/%T"

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
| **allow-other** |`true` / `false` |`false` |allow other users to access the data in the mounted directory |
| ~~allow-root~~ |`true` / `false` |`false` |allow root user to access the data in the mounted directory. `restic < 0.10.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **no-default-permissions** |`true` / `false` |`false` |for 'allow-other', ignore Unix permissions and allow users to read all snapshot files |
| **owner-root** |`true` / `false` |`false` |use 'root' as the owner of files and dirs |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **path-template** |one or more `strings` | |set template for path names. `restic >= 0.14.0`  |
| ~~snapshot-template~~ |`string` |`"2006-01-02T15:04:05Z07:00"` |set template to use for snapshot dirs. `restic < 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **time-template** |`string` |`"2006-01-02T15:04:05Z07:00"` |set template to use for times. `restic >= 0.14.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**prune**

This section configures restic command `prune` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "prune" command checks the repository and removes data that is not
referenced and therefore not needed any more.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **schedule** |one or more `strings` | |Set the times at which the scheduled command is run (times are specified in systemd timer format). **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` |`false` |Don't schedule the start of this profile when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't schedule the start of this profile when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't schedule the start of this profile when running on battery, and the battery charge left is less than the value |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `tcp://localhost:514`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](#nested-sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **dry-run** |`true` / `false` |`false` |do not modify the repository, just print what would be done. `restic >= 0.12.0`  |
| **max-repack-size** |`string` |`""` |maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T). `restic >= 0.12.0`  |
| **max-unused** |`string` |`"5%"` |tolerate given limit of unused data (absolute value in bytes with suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited'). `restic >= 0.12.0`  |
| **repack-cacheable-only** |`true` / `false` |`false` |only repack packs which are cacheable. `restic >= 0.12.0`  |
| **repack-small** |`true` / `false` |`false` |repack pack files below 80% of target pack size. `restic >= 0.14.0`  |
| **repack-uncompressed** |`true` / `false` |`false` |repack all uncompressed data. `restic >= 0.14.0`  |
| **unsafe-recover-no-free-space** |`string` |`""` |UNSAFE, READ THE DOCUMENTATION BEFORE USING! Try to recover a repository stuck with no free space. Do not use without trying out 'prune --max-repack-size 0' first. `restic >= 0.14.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**rebuild-index**

This section configures restic command `rebuild-index`  removed in `0.16.0` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "rebuild-index" command creates a new index based on the pack files in the
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
| **read-all-packs** |`true` / `false` |`false` |read all pack files to generate new index from scratch. `restic >= 0.12.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**recover**

This section configures restic command `recover` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "recover" command builds a new snapshot from all directories it can find in
the raw data of the repository which are not referenced in an existing snapshot.
It can be used if, for example, a snapshot has been removed by accident with "forget".

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**repair**

This section configures restic command `repair`  available since `0.16.0` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

Repair the repository

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
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
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**repair-snapshots**

This section configures restic command `repair-snapshots`  available since `0.16.0` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "repair snapshots" command repairs broken snapshots. It scans the given
snapshots and generates new ones with damaged directories and file contents
removed. If the broken snapshots are deleted, a prune run will be able to
clean up the repository.

The command depends on a correct index, thus make sure to run "repair index"
first!

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
| **dry-run** |`true` / `false` |`false` |do not do anything, just print what would be done. `restic >= 0.16.0`  |
| **forget** |`true` / `false` |`false` |remove original snapshots after creating new ones. `restic >= 0.16.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.16.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.16.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.16.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**restore**

This section configures restic command `restore` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "restore" command extracts the data from a snapshot from the repository to
a directory.

The special snapshot "latest" can be used to restore the latest snapshot in the
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
| **exclude** |one or more `strings` | |exclude a pattern |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host, when snapshot ID "latest" is given. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **iexclude** |one or more `strings` | |same as --exclude but ignores the casing of filenames. `restic >= 0.10.0`  |
| **iinclude** |one or more `strings` | |same as --include but ignores the casing of filenames. `restic >= 0.10.0`  |
| **include** |one or more `strings` | |include a pattern, exclude everything else |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path, when snapshot ID "latest" is given. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **sparse** |`true` / `false` |`false` |restore files as sparse. `restic >= 0.15.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...], when snapshot ID "latest" is given. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |
| **target** |`string` |`""` |directory to extract data to |
| **verify** |`true` / `false` |`false` |verify restored files content |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**retention**

This section configures restic command `forget` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "forget" command removes snapshots according to a policy. All snapshots are
first divided into groups according to "--group-by", and after that the policy
specified by the "--keep-*" options is applied to each group individually.

Please note that this command really only deletes the snapshot object in the
repository, which is a reference to data stored there. In order to remove the
unreferenced data after "forget" was run successfully, see the "prune" command.

Please also read the documentation for "forget" to learn about some important
security considerations.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **after-backup** |`true` / `false` | |Apply retention after the backup command succeeded. Defaults to true in configuration format v2 if any "keep-*" flag is set and "before-backup" is unset |
| **before-backup** |`true` / `false` | |Apply retention before starting the backup command |
| ~~schedule~~ |one or more `strings` | |Set the times at which the scheduled command is run (times are specified in systemd timer format). **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| ~~schedule-after-network-online~~ |`true` / `false` |`false` |Don't schedule the start of this profile when the network is offline (supported in "systemd") |
| ~~schedule-capture-environment~~ |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| ~~schedule-ignore-on-battery~~ |`true` / `false` |`false` |Don't schedule the start of this profile when running on battery |
| ~~schedule-ignore-on-battery-less-than~~ |`integer` | |Don't schedule the start of this profile when running on battery, and the battery charge left is less than the value |
| ~~schedule-lock-mode~~ |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| ~~schedule-lock-wait~~ |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| ~~schedule-log~~ |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `tcp://localhost:514`  |
| ~~schedule-permission~~ |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| ~~schedule-priority~~ |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **compact** |`true` / `false` |`false` |use compact output format |
| **dry-run** |`true` / `false` |`false` |do not delete anything, just print what would be done |
| **group-by** |`string` |`host,paths` |group snapshots by host, paths and/or tags, separated by comma (disable grouping with '') |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname that applies in section "backup". Defaults to true for config version 2 in "retention". **Examples**: `true`, `false`, `"host"`  |
| **keep-daily** |`integer` |`0` |keep the last n daily snapshots (use 'unlimited' to keep all daily snapshots) |
| **keep-hourly** |`integer` |`0` |keep the last n hourly snapshots (use 'unlimited' to keep all hourly snapshots) |
| **keep-last** |`integer` |`0` |keep the last n snapshots (use 'unlimited' to keep all snapshots) |
| **keep-monthly** |`integer` |`0` |keep the last n monthly snapshots (use 'unlimited' to keep all monthly snapshots) |
| **keep-tag** |one or more `strings` | |keep snapshots with this taglist |
| **keep-weekly** |`integer` |`0` |keep the last n weekly snapshots (use 'unlimited' to keep all weekly snapshots) |
| **keep-within** |`string` | |keep snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot |
| **keep-within-daily** |`string` | |keep daily snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-hourly** |`string` | |keep hourly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-monthly** |`string` | |keep monthly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-weekly** |`string` | |keep weekly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-within-yearly** |`string` | |keep yearly snapshots that are newer than duration (eg. 1y5m7d2h) relative to the latest snapshot. `restic >= 0.13.0`  |
| **keep-yearly** |`integer` |`0` |keep the last n yearly snapshots (use 'unlimited' to keep all yearly snapshots) |
| **max-repack-size** |`string` |`""` |maximum size to repack (allowed suffixes: k/K, m/M, g/G, t/T). `restic >= 0.12.0`  |
| **max-unused** |`string` |`"5%"` |tolerate given limit of unused data (absolute value in bytes with suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited'). `restic >= 0.12.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". Defaults to true in "retention". **Examples**: `true`, `false`, `"path"`  |
| **prune** |`true` / `false` |`false` |automatically run the 'prune' command if snapshots have been removed |
| **repack-cacheable-only** |`true` / `false` |`false` |only repack packs which are cacheable. `restic >= 0.12.0`  |
| **repack-small** |`true` / `false` |`false` |repack pack files below 80% of target pack size. `restic >= 0.14.0`  |
| **repack-uncompressed** |`true` / `false` |`false` |repack all uncompressed data. `restic >= 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". Defaults to true for config version 2 in "retention". **Examples**: `true`, `false`, `"tag"`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**rewrite**

This section configures restic command `rewrite`  available since `0.15.0` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "rewrite" command excludes files from existing snapshots. It creates new
snapshots containing the same data as the original ones, but without the files
you specify to exclude. All metadata (time, host, tags) will be preserved.

The snapshots to rewrite are specified using the --host, --tag and --path options,
or by providing a list of snapshot IDs. Please note that specifying neither any of
these options nor a snapshot ID will cause the command to rewrite all snapshots.

The special tag 'rewrite' will be added to the new snapshots to distinguish
them from the original ones, unless --forget is used. If the --forget option is
used, the original snapshots will instead be directly removed from the repository.

Please note that the --forget option only removes the snapshots and not the actual
data stored in the repository. In order to delete the no longer referenced data,
use the "prune" command.

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
| **dry-run** |`true` / `false` |`false` |do not do anything, just print what would be done. `restic >= 0.15.0`  |
| **exclude** |one or more `strings` | |exclude a pattern. `restic >= 0.15.0`  |
| **exclude-file** |one or more `strings` | |read exclude patterns from a file. `restic >= 0.15.0`  |
| **forget** |`true` / `false` |`false` |remove original snapshots after creating new ones. `restic >= 0.15.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.15.0`  |
| **iexclude** |one or more `strings` | |same as --exclude pattern but ignores the casing of filenames. `restic >= 0.15.0`  |
| **iexclude-file** |one or more `strings` | |same as --exclude-file but ignores casing of filenames in patterns. `restic >= 0.15.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.15.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.15.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**snapshots**

This section configures restic command `snapshots` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "snapshots" command lists all snapshots stored in the repository.

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
| **compact** |`true` / `false` |`false` |use compact output format |
| **group-by** |`string` | |group snapshots by host, paths and/or tags, separated by comma. `restic >= 0.10.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| ~~last~~ |`true` / `false` |`false` |only show the last snapshot for each host and path. `restic < 0.13.0`  |
| **latest** |`integer` |`0` |only show the last n snapshots for each host and path. `restic >= 0.13.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**stats**

This section configures restic command `stats` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "stats" command walks one or multiple snapshots in a repository
and accumulates statistics about the data stored therein. It reports
on the number of unique files and their sizes, according to one of
the counting modes as given by the --mode flag.

It operates on all snapshots matching the selection criteria or all
snapshots if nothing is specified. The special snapshot ID "latest"
is also supported. Some modes make more sense over
just a single snapshot, while others are useful across all snapshots,
depending on what you are trying to calculate.

The modes are:

restore-size: (default) Counts the size of the restored files.
files-by-contents: Counts total size of files, where a file is
considered unique if it has unique contents.
raw-data: Counts the size of blobs in the repository, regardless of
how many files reference them.
blobs-per-file: A combination of files-by-contents and raw-data.


Refer to the online manual for more details about each mode.

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
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **mode** |`string` |`"restore-size"` |counting mode: restore-size (default), files-by-contents, blobs-per-file or raw-data |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.10.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.10.0`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**tag**

This section configures restic command `tag` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "tag" command allows you to modify tags on exiting snapshots.

You can either set/replace the entire set of tags on a snapshot, or
add tags to/remove tags from the existing set.

When no snapshot-ID is given, all snapshots matching the host, tag and path filter criteria are modified.

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
| **add** |one or more `strings` | |tags which will be added to the existing tags in the format tag[,tag,...] |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host. Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path. Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"`  |
| **remove** |one or more `strings` | |tags which will be removed from the existing tags in the format tag[,tag,...] |
| **set** |one or more `strings` | |tags which will replace the existing tags in the format tag[,tag,...] |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"`  |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)
#### Section profile.**unlock**

This section configures restic command `unlock` .
  Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "unlock" command removes stale locks that have been created by other restic processes.

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
| **remove-all** |`true` / `false` |`false` |remove all locks, even non-stale ones |



{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](#section-profile)*
can be overridden in this section.
{{% /notice %}}

[go to top](#reference)


### Nested profile sections

Nested sections describe configuration structure that is assigned to flags within the
configuration, see [HTTP Hooks]({{% relref "/configuration/http_hooks/" %}}) as an example.

#### Nested *SendMonitoringHeader*



| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **name** |`string` | |Name of the HTTP header. **Examples**: `"Authorization"`, `"Cache-Control"`, `"Content-Disposition"`, `"Content-Type"`  |
| **value** |`string` | |Value of the header. **Examples**: `"Bearer ..."`, `"Basic ..."`, `"no-cache"`, `"attachment; filename=stats.txt"`, `"application/json"`, `"text/plain"`, `"text/xml"`  |


#### Nested *SendMonitoringSection*



| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **body** |`string` | |Request body, overrides "body-template" |
| **body-template** |`string` | |Path to a file containing the request body (go template). See [configuration/http_hooks/#body-template]({{% relref "/configuration/http_hooks/#body-template" %}}) |
| **headers** | one or more nested *[SendMonitoringHeader](#nested-sendmonitoringheader)* | |Additional HTTP headers to send with the request |
| **method** |`string` |`GET` |HTTP method of the request. Is one of `GET`, `DELETE`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT`, `TRACE`  |
| **skip-tls-verification** |`true` / `false` |`false` |Enables insecure TLS (without verification), see also "global.ca-certificates" |
| **url** |`uri` | |URL of the target to send to |


#### Nested *StreamErrorSection*



| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **max-runs** |`integer` |`0` |Maximum amount of times that "run" is started ; 0 for no limit. Must be >= 0  |
| **min-matches** |`integer` |`0` |Minimum amount of times the "pattern" must match before "run" is started ; 0 for no limit. Must be >= 0  |
| **pattern** |`regex` | |A regular expression pattern that is tested against stderr of a running restic command |
| **run** |`string` | |The shell command to run when the pattern matches |




[go to top](#reference)

### Section **groups**

Config file format v1 uses a simplified groups section. Every named item below `groups`
maps to one or more `profile` names (list of strings).

From file format v2, every named item in the groups section is configuration structure
following the format described below (see [Configuration v2]({{% relref "/configuration/v2/" %}})
for details):

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **continue-on-error** |nested *[Bool](#nested-bool)* |`auto` |Continue with the next profile on a failure, overrides "global.group-continue-on-error" |
| **description** |`string` | |Describe the group |
| **profiles** |one or more `strings` | |Names of the profiles belonging to this group |




[go to top](#reference)

## Value types

**Type**: `true` / `false`
: Is a boolean value to toggle a flag or specific behaviour that has
to match the syntax for booleans in the used file format.

**Type**: `numeric` & `integer`
: Is a numeric value (whole-number for `integer`) that has to match the
syntax for numbers in the used file format.

**Type**: `string`
: Is a sequence of UTF-8 characters that usually have to be placed in
quotes and must match the syntax for strings in the used file format.

**Type**: `duration`
: A duration `string` is a sequence of decimal numbers, each with optional
fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid
time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

**Type**: `uri`
: An uri `string` must be a valid URI or URL.

**Type**: `hostname`
: A hostname `string` must be a valid hostname or FQDN.

**Type**: one or more `[type]s`
: Indicates that multiple values of `[type]` can be specified as a list.
The list syntax depends on the used file format.

**Type**: one or more nested `[type]`
: Is a value or a list of values that follow the configuration structure declared
in `[type]`.

**Type**: `key` => `[type]`
: Is a value that is configuration structure of `string` keys and values of `[type]`.

**Type**: `key` => `value`
: Is a value that is configuration structure of `string` keys and values of any type.

[go to top](#reference)

## JSON schema

resticprofile provides a JSON schema for v1 & v2 configuration files. The schema may be
used to validate configuration files in JSON format (and possibly others), see
[JSON schema]({{% relref "/configuration/jsonschema/" %}}) for details.

JSON schema URLs for **any** *restic* version:

* Config V1: https://creativeprojects.github.io/resticprofile/jsonschema/config-1.json
* Config V2: https://creativeprojects.github.io/resticprofile/jsonschema/config-2.json

JSON schema URLs for a specific *restic* version:

* `.../config-1-restic-{MAJOR}-{MINOR}.json`
* `.../config-2-restic-{MAJOR}-{MINOR}.json`

Available URLs:

 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-16.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-16.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-15.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-15.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-14.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-14.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-13.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-13.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-12.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-12.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-11.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-11.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-10.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-10.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-2-restic-0-9.json
 * https://creativeprojects.github.io/resticprofile/jsonschema/config-1-restic-0-9.json

[go to top](#reference)
