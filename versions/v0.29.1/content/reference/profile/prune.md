---
title: prune
weight: 20
---
#### Section profile.**prune**

This section configures restic command `prune` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "prune" command checks the repository and removes data that is not
referenced and therefore not needed any more.

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
| **schedule-priority** |`string` |`standard` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **send-after** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |



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




{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}


{{< pageversions "v0.27.1" "v0.28.1" >}}
