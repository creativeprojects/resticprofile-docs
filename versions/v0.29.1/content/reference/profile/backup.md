---
title: backup
weight: 1
---
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
| **schedule** |one or more `strings` OR nested *[ScheduleConfig](../nested/scheduleconfig)* | |Configures the scheduled execution of this profile section. Can be times in systemd timer format or a config structure. **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` | |Don't start this schedule when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't start this schedule when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't start this schedule when running on battery and the state of charge is less than this percentage. **Examples**: `20`, `33`, `50`, `75`  |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `syslog-tcp://syslog-server:514`, `syslog:server`, `syslog:`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`standard` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |
| **source** |one or more `strings` | |The paths to backup. **Examples**: `/opt/`, `/home/user/`, `C:\Users\User\Documents`  |
| **source-base** |`string` | |The base path to resolve relative backup paths against. Defaults to current directory if unset or empty (see also "base-dir" in profile). **Examples**: `/`, `$PWD`, `C:\`, `%cd%`  |
| **source-relative** |`true` / `false` |`false` |Enable backup with relative source paths. This will change the working directory of the "restic backup" command to "source-base", and will not expand "source" to an absolute path |
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
| **force** |`true` / `false` |`false` |force re-reading the source files/directories (overrides the "parent" flag) |
| **group-by** |`string` |`host,paths` |group snapshots by host, paths and/or tags, separated by comma (disable grouping with ''). `restic >= 0.16.0`  |
| **host** |`true` / `false` OR `hostname` |`""` |set the hostname for the snapshot manually (default: $RESTIC_HOST). To prevent an expensive rescan use the "parent" flag. Boolean true is replaced with the hostname of the system. Defaults to true for config version 2 in "backup". **Examples**: `true`, `false`, `"host"`  |
| **iexclude** |one or more `strings` | |same as --exclude pattern but ignores the casing of filenames. `restic >= 0.10.0`  |
| **iexclude-file** |one or more `strings` | |same as --exclude-file but ignores casing of filenames in patterns. `restic >= 0.10.0`  |
| **ignore-ctime** |`true` / `false` |`false` |ignore ctime changes when checking for modified files. `restic >= 0.12.0`  |
| **ignore-inode** |`true` / `false` |`false` |ignore inode number and ctime changes when checking for modified files. `restic >= 0.10.0`  |
| **no-scan** |`true` / `false` |`false` |do not run scanner to estimate size of backup. `restic >= 0.15.0`  |
| **one-file-system** |`true` / `false` |`false` |exclude other file systems, don't cross filesystem boundaries and subvolumes |
| **parent** |`string` |`""` |use this parent snapshot (default: latest snapshot in the group determined by --group-by and not newer than the timestamp determined by --time) |
| **read-concurrency** |`integer` |`0` |read n files concurrently (default: $RESTIC_READ_CONCURRENCY or 2). `restic >= 0.15.0`  |
| **skip-if-unchanged** |`true` / `false` |`false` |skip snapshot creation if identical to parent snapshot. `restic >= 0.17.0`  |
| **stdin** |`true` / `false` |`false` |read backup from stdin |
| **stdin-filename** |`string` |`"stdin"` |filename to use when reading from stdin |
| **stdin-from-command** |`true` / `false` |`false` |interpret arguments as command to execute and store its stdout. `restic >= 0.17.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |add tags for the new snapshot in the format tag[,tag,...]. Boolean true is unsupported in section "backup". **Examples**: `false`, `"tag"`  |
| **time** |`string` |`""` |time of the backup (ex. '2012-11-01 22:08:41') (default: now) |
| **use-fs-snapshot** |`true` / `false` |`false` |use filesystem snapshot where possible (currently only Windows VSS). `restic >= 0.12.0` . Available only for `windows`  |
| **with-atime** |`true` / `false` |`false` |store the atime for all files and directories |




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.28.1" >}}
