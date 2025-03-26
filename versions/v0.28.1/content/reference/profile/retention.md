---
title: retention
weight: 28
---
#### Section profile.**retention**

This section configures restic command `forget` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "forget" command removes snapshots according to a policy. All snapshots are
first divided into groups according to "--group-by", and after that the policy
specified by the "--keep-" options is applied to each group individually.
If there are not enough snapshots to keep one for each duration related
"--keep-{within-,}" option, the oldest snapshot in the group is kept
additionally.

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
| ~~schedule~~ |one or more `strings` OR nested *[ScheduleConfig](../nested/scheduleconfig)* | |Configures the scheduled execution of this profile section. Can be times in systemd timer format or a config structure. **Examples**: `hourly`, `daily`, `weekly`, `monthly`, `10:00,14:00,18:00,22:00`, `Wed,Fri 17:48`, `*-*-15 02:45`, `Mon..Fri 00:30`  |
| ~~schedule-after-network-online~~ |`true` / `false` | |Don't start this schedule when the network is offline (supported in "systemd") |
| ~~schedule-capture-environment~~ |one or more `strings` |`RESTIC_*` |Set names (or glob expressions) of environment variables to capture during schedule creation. The captured environment is applied prior to "profile.env" when running the schedule. Whether capturing is supported depends on the type of scheduler being used (supported in "systemd" and "launchd") |
| ~~schedule-ignore-on-battery~~ |`true` / `false` |`false` |Don't start this schedule when running on battery |
| ~~schedule-ignore-on-battery-less-than~~ |`integer` | |Don't start this schedule when running on battery and the state of charge is less than this percentage. **Examples**: `20`, `33`, `50`, `75`  |
| ~~schedule-lock-mode~~ |`string` |`default` |Specify how locks are used when running on schedule - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `default`, `fail`, `ignore`  |
| ~~schedule-lock-wait~~ |`integer` OR `duration` | |Set the maximum time to wait for acquiring locks when running on schedule. **Examples**: `150s`, `15m`, `30m`, `45m`, `1h`, `2h30m`  |
| ~~schedule-log~~ |`string` | |Redirect the output into a log file or to syslog when running on schedule. **Examples**: `/resticprofile.log`, `syslog-tcp://syslog-server:514`, `syslog:server`, `syslog:`  |
| ~~schedule-permission~~ |`string` |`auto` |Specify whether the schedule runs with system or user privileges - see [schedules/configuration/]({{% relref "/schedules/configuration/" %}}). Is one of `auto`, `system`, `user`, `user_logged_on`  |
| ~~schedule-priority~~ |`string` |`background` |Set the priority at which the schedule is run. Is one of `background`, `standard`  |



##### Flags passed to the **restic** command line:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **compact** |`true` / `false` |`false` |use compact output format |
| **dry-run** |`true` / `false` |`false` |do not delete anything, just print what would be done |
| **group-by** |`string` |`host,paths` |group snapshots by host, paths and/or tags, separated by comma (disable grouping with '') |
| **host** |`true` / `false` OR one or more `hostnames` | |only consider snapshots for this host (can be specified multiple times) (default: $RESTIC_HOST). Boolean true is replaced with the hostname that applies in section "backup". Defaults to true for config version 2 in "retention". **Examples**: `true`, `false`, `"host"`  |
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
| **path** |`true` / `false` OR one or more `strings` | |only consider snapshots including this (absolute) path (can be specified multiple times, snapshots must include all specified paths). Boolean true is replaced with the paths from section "backup". Defaults to true in "retention". **Examples**: `true`, `false`, `"path"`  |
| **prune** |`true` / `false` |`false` |automatically run the 'prune' command if snapshots have been removed |
| **repack-cacheable-only** |`true` / `false` |`false` |only repack packs which are cacheable. `restic >= 0.12.0`  |
| **repack-small** |`true` / `false` |`false` |repack pack files below 80% of target pack size. `restic >= 0.14.0`  |
| **repack-uncompressed** |`true` / `false` |`false` |repack all uncompressed data. `restic >= 0.14.0`  |
| **tag** |`true` / `false` OR one or more `strings` | |only consider snapshots including tag[,tag,...]. Boolean true is replaced with the tags from section "backup". Defaults to true for config version 2 in "retention". **Examples**: `true`, `false`, `"tag"`  |
| **unsafe-allow-remove-all** |`true` / `false` |`false` |allow deleting all snapshots of a snapshot group. `restic >= 0.17.0`  |




{{% notice hint %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

