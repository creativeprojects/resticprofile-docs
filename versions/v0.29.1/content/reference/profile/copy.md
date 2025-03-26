---
title: copy
weight: 5
---
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
| **snapshot** |one or more `strings` | |Snapshot IDs to copy (if empty, all snapshots are copied) |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **from-insecure-no-password** |`true` / `false` |`false` |use an empty password for the source repository, must be passed to every restic command (insecure). `restic >= 0.17.0`  |
| **from-key-hint** |`string` |`""` |key ID of key to try decrypting the source repository first (default: $RESTIC_FROM_KEY_HINT). `restic >= 0.14.0`  |
| **from-password-command** |`string` |`""` |shell command to obtain the source repository password from (default: $RESTIC_FROM_PASSWORD_COMMAND). `restic >= 0.14.0`  |
| **from-password-file** |`string` |`""` |file to read the source repository password from (default: $RESTIC_FROM_PASSWORD_FILE). `restic >= 0.14.0`  |
| **from-repository** |`string` |`""` |source repository to copy snapshots from (default: $RESTIC_FROM_REPOSITORY). `restic >= 0.14.0`  |
| **from-repository-file** |`string` |`""` |file from which to read the source repository location to copy snapshots from (default: $RESTIC_FROM_REPOSITORY_FILE). `restic >= 0.14.0`  |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname of the system. **Examples**: `true`, `false`, `"host"` . `restic >= 0.10.0`  |
| ~~key-hint2~~ |`string` |`""` |key ID of key to try decrypting the destination repository first (default: $RESTIC_KEY_HINT2). `restic >= 0.10.0 < 0.14.0`  |
| ~~password-command2~~ |`string` |`""` |shell command to obtain the destination repository password from (default: $RESTIC_PASSWORD_COMMAND2). `restic >= 0.10.0 < 0.14.0`  |
| ~~password-file2~~ |`string` |`""` |file to read the destination repository password from (default: $RESTIC_PASSWORD_FILE2). `restic >= 0.10.0 < 0.14.0`  |
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". **Examples**: `true`, `false`, `"path"` . `restic >= 0.10.0`  |
| ~~repo2~~ |`string` |`""` |destination repository to copy snapshots to (default: $RESTIC_REPOSITORY2). `restic >= 0.10.0 < 0.14.0`  |
| ~~repository-file2~~ |`string` |`""` |file from which to read the destination repository location to copy snapshots to (default: $RESTIC_REPOSITORY_FILE2). `restic >= 0.13.0 < 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". **Examples**: `true`, `false`, `"tag"` . `restic >= 0.10.0`  |




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.28.1" >}}
