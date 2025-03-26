---
title: check
weight: 4
---
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
| **schedule** |one or more `strings` OR nested *[ScheduleConfig](../nested/scheduleconfig)* | |Configures the scheduled execution of this profile section. Can be times in systemd timer format or a config structure. **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| **schedule-after-network-online** |`true` / `false` | |Don't start this schedule when the network is offline (supported in "systemd") |
| **schedule-capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **schedule-ignore-on-battery** |`true` / `false` |`false` |Don't start this schedule when running on battery |
| **schedule-ignore-on-battery-less-than** |`integer` | |Don't start this schedule when running on battery and the state of charge is less than this percentage. **Examples**: `20`, `33`, `50`, `75`  |
| **schedule-lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **schedule-lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **schedule-log** |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `syslog-tcp://syslog-server:514`, `syslog:server`, `syslog:`  |
| **schedule-permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **schedule-priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| ~~check-unused~~ |`true` / `false` |`false` |find unused blobs. `restic < 0.14.0`  |
| **read-data** |`true` / `false` |`false` |read all data blobs |
| **read-data-subset** |`string` |`""` |read a subset of data packs, specified as 'n/t' for specific part, or either 'x%' or 'x.y%' or a size in bytes with suffixes k/K, m/M, g/G, t/T for a random subset |
| **with-cache** |`true` / `false` |`false` |use existing cache, only read uncached data from repository |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.29.1" >}}
