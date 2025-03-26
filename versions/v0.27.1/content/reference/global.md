---
title: Section global
weight: 1
---


### Section **global**

The `global` section is at the root of the configuration file and contains the global
settings for resticprofile.

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **ca-certificates** |one or more `strings` | |Path to PEM encoded certificates to trust in addition to system certificates when resticprofile sends to a webhook - see [configuration/http_hooks/]({{% relref "/configuration/http_hooks/" %}}) |
| **command-output** |`string` |`auto` |Sets the destination for command output (stderr/stdout). "log" sends output to the log file (if specified), "console" sends it to the console instead. "auto" sends it to "both" if console is a terminal otherwise to "log" only - see [configuration/logs/]({{% relref "/configuration/logs/" %}}). Is one of `auto`, `log`, `console`, `all`  |
| **default-command** |`string` |`snapshots` |The restic or resticprofile command to use when no command was specified. **Examples**: `backup`, `cache`, `cat`, `check`, `copy`, `diff`, `dump`, `find`, `forget`, `generate`, `init`, `key`, `list`, `ls`, `migrate`, `mount`, `prune`, `rebuild-index`, `recover`, `repair`, `repair-index`, `repair-snapshots`, `restore`, `rewrite`, `self-update`, `snapshots`, `stats`, `tag`, `unlock`, `version`  |
| **group-continue-on-error** |`true` / `false` |`false` |Enable groups to continue with the next profile(s) instead of stopping at the first failure |
| **initialize** |`true` / `false` |`false` |Initialize a repository if missing |
| **ionice** |`true` / `false` |`false` |Enables setting the unix IO priority class and level for resticprofile and child processes (only on unix OS) |
| **ionice-class** |`integer` |`2` |Sets the unix "ionice-class" to apply when "ionice" is enabled. Must be >= 1 and  <= 3  |
| **ionice-level** |`integer` |`0` |Sets the unix "ionice-level" to apply when "ionice" is enabled. Must be >= 0 and  <= 7  |
| ~~legacy-arguments~~ |`true` / `false` |`false` |Legacy, broken arguments mode of resticprofile before version 0.15 |
| **log** |`string` | |Sets the default log destination to be used if not specified in "--log" or "schedule-log" - see [configuration/logs/]({{% relref "/configuration/logs/" %}}). **Examples**: `/resticprofile.log`, `syslog-tcp://syslog-server:514`, `syslog:server`, `syslog:`  |
| **min-memory** |`integer` |`100` |Minimum available memory (in MB) required to run any commands - see [usage/memory/]({{% relref "/usage/memory/" %}}) |
| **nice** |`integer` |`0` |Sets the unix "nice" value for resticprofile and child processes (on any OS). Must be >= -20 and  <= 19  |
| **prevent-auto-repository-file** |`true` / `false` |`false` |Prevents using a repository file for repository definitions containing a password |
| **prevent-sleep** |`true` / `false` |`false` |Prevent the system from sleeping while running commands - see [configuration/sleep/]({{% relref "/configuration/sleep/" %}}) |
| **priority** |`string` |`normal` |Sets process priority class for resticprofile and child processes (on any OS). Is one of `idle`, `background`, `low`, `normal`, `high`, `highest`  |
| **restic-arguments-filter** |`true` / `false` |`true` |Remove unknown flags instead of passing all configured flags to restic |
| **restic-binary** |`string` | |Full path of the restic executable (detected if not set) |
| **restic-lock-retry-after** |`integer` OR `duration` |`1m` |Time to wait before trying to get a lock on a restic repositoey - see [usage/locks/]({{% relref "/usage/locks/" %}}) |
| **restic-stale-lock-age** |`integer` OR `duration` |`1h` |The age an unused lock on a restic repository must have at least before resiticprofile attempts to unlock - see [usage/locks/]({{% relref "/usage/locks/" %}}) |
| **schedule-defaults** |nested *[ScheduleBaseConfig](../nested/schedulebaseconfig)* | |Sets defaults for all schedules |
| **scheduler** |`string` |`auto` |Selects the scheduler. Blank or "auto" uses the default scheduler of your operating system: "launchd", "systemd", "taskscheduler" or "crond" (as fallback). Alternatively you can set "crond" for cron compatible schedulers supporting the crontab executable API or "crontab:[user:]file" to write into a crontab file directly. The need for a user is detected if missing and can be set to a name, "-" (no user) or "*" (current user). **Examples**: `auto`, `launchd`, `systemd`, `taskscheduler`, `crond`, `crond:/usr/bin/crontab`, `crontab:*:/etc/cron.d/resticprofile`  |
| **send-timeout** |`integer` OR `duration` |`30s` |Timeout when sending messages to a webhook - see [configuration/http_hooks/]({{% relref "/configuration/http_hooks/" %}}). **Examples**: `15s`, `30s`, `2m30s`  |
| **shell** |one or more `strings` |`auto` |The shell that is used to run commands (default is OS specific). **Examples**: `sh`, `bash`, `pwsh`, `powershell`, `cmd`  |
| **systemd-timer-template** |`string` | |File containing the go template to generate a systemd timer - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |
| **systemd-unit-template** |`string` | |File containing the go template to generate a systemd unit - see [schedules/systemd/]({{% relref "/schedules/systemd/" %}}) |





{{< pageversions "v0.28.1" "v0.29.1" >}}
