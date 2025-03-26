---
title: ScheduleBaseConfig
weight: 1
---
#### Nested *ScheduleBaseConfig*



| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **after-network-online** |`true` / `false` | |Don't start this schedule when the network is offline (supported in "systemd") |
| **capture-environment** |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| **command-output** |`string` |`auto` |Sets the destination for command output (stderr/stdout). "log" sends output to the log file (if specified), "console" sends it to the console instead. "auto" sends it to "both" if console is a terminal otherwise to "log" only - see [configuration/logs/]({{% relref "/configuration/logs/" %}}). Is one of `auto`, `log`, `console`, `all`  |
| **ignore-on-battery** |`true` / `false` |`false` |Don't start this schedule when running on battery |
| **ignore-on-battery-less-than** |`integer` | |Don't start this schedule when running on battery and the state of charge is less than this percentage. **Examples**: `20`, `33`, `50`, `75`  |
| **lock-mode** |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| **lock-wait** |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| **log** |`string` | |Redirect the output into a log file or to syslog when running on schedule - see [configuration/logs/]({{% relref "/configuration/logs/" %}}). **Examples**: `/resticprofile.log`, `syslog-tcp://syslog-server:514`, `syslog:server`, `syslog:`  |
| **permission** |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| **priority** |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |
| **systemd-drop-in-files** |one or more `strings` | |Files containing systemd drop-in (override) files - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |


