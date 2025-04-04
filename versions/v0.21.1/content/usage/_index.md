---
chapter: true
pre: <b>3. </b>
title: Usage
weight: 15
---




Here are a few examples how to run resticprofile (using the main example configuration file)

See all snapshots of your `[default]` profile:

```shell
$ resticprofile
```

See all available profiles in your configuration file (and the restic commands where some flags are defined):

```shell
$ resticprofile profiles

Profiles available (name, sections, description):
  root:           (backup, copy, forget, retention)
  self:           (backup, check, copy, forget, retention)
  src:            (backup, copy, retention, snapshots)

Groups available (name, profiles, description):
  full-backup:  [root, src]

```

Backup root & src profiles (using _full-backup_ group shown earlier)

```shell
$ resticprofile --name "full-backup" backup
```
or using the syntax introduced in v0.17.0:

```shell
$ resticprofile full-backup.backup
```

Assuming the _stdin_ profile from the configuration file shown before, the command to send a mysqldump to the backup is as simple as:

```shell
$ mysqldump --all-databases --order-by-primary | resticprofile --name stdin backup
```
or using the syntax introduced in v0.17.0:

```shell
$ mysqldump --all-databases --order-by-primary | resticprofile stdin.backup
```

Mount the default profile (_default_) in /mnt/restic:

```shell
$ resticprofile mount /mnt/restic
```

Display quick help

```shell
$ resticprofile --help

Usage of resticprofile:
	resticprofile [resticprofile flags] [profile name.][restic command] [restic flags]
	resticprofile [resticprofile flags] [profile name.][resticprofile command] [command specific flags]

resticprofile flags:
  -c, --config string        configuration file (default "profiles")
      --dry-run              display the restic commands instead of running them
  -f, --format string        file format of the configuration (default is to use the file extension)
  -h, --help                 display this help
      --lock-wait duration   wait up to duration to acquire a lock (syntax "1h5m30s")
  -l, --log string           logs to a target instead of the console
  -n, --name string          profile name (default "default")
      --no-ansi              disable ansi control characters (disable console colouring)
      --no-lock              skip profile lock file
      --no-prio              don't set any priority on load: used when started from a service that has already set the priority
  -q, --quiet                display only warnings and errors
      --theme string         console colouring theme (dark, light, none) (default "light")
      --trace                display even more debugging information
  -v, --verbose              display some debugging information
  -w, --wait                 wait at the end until the user presses the enter key

resticprofile own commands:
   version       display version (run in verbose mode for detailed information)
   self-update   update to latest resticprofile (use -q/--quiet flag to update without confirmation)
   profiles      display profile names from the configuration file
   show          show all the details of the current profile
   schedule      schedule jobs from a profile (use --all flag to schedule all jobs of all profiles)
   unschedule    remove scheduled jobs of a profile (use --all flag to unschedule all profiles)
   status        display the status of scheduled jobs (use --all flag for all profiles)
   generate      generate resources (--random-key [size], --bash-completion & --zsh-completion)


```

A command is either a restic command or a resticprofile own command.


## Command line reference

There are not many options on the command line, most of the options are in the configuration file.

* **[-h]**: Display quick help
* **[-c | --config] configuration_file**: Specify a configuration file other than the default
* **[-f | --format] configuration_format**: Specify the configuration file format: `toml`, `yaml`, `json` or `hcl`
* **[-n | --name] profile_name**: Profile section to use from the configuration file.
  You can also use `[profile_name].[command]` syntax instead, this will only work if `-n` is not set.
  Using `-n [profile_name] [command]` or `[profile_name].[command]` both select profile and command and are technically equivalent.
* **[--dry-run]**: Doesn't run the restic command but display the command line instead
* **[-q | --quiet]**: Force resticprofile and restic to be quiet (override any configuration from the profile)
* **[-v | --verbose]**: Force resticprofile and restic to be verbose (override any configuration from the profile)
* **[--no-ansi]**: Disable console colouring (to save output into a log file)
* **[--no-lock]**: Disable resticprofile locks, neither create nor fail on a lock. restic locks are unaffected by this option.
* **[--theme]**: Can be `light`, `dark` or `none`. The colours will adjust to a 
light or dark terminal (none to disable colouring)
* **[--lock-wait] duration**: Retry to acquire resticprofile and restic locks for up to the specified amount of time before failing on a lock failure. 
* **[-l | --log] file path or url**: To write the logs to a file or a syslog server instead of displaying on the console. 
The format of the server url is `tcp://192.168.0.1:514` or `udp://localhost:514`.
For custom log forwarding, the prefix `temp:` can be used (e.g. `temp:/t/msg.log`) to create unique log output that can be fed 
into a command or http hook by referencing it with `{{ tempDir }}/...` or `{{ tempFile "msg.log" }}` in the configuration file.
* **[-w | --wait]**: Wait at the very end of the execution for the user to press enter. 
This is only useful in Windows when resticprofile is started from explorer and the console window closes automatically at the end.
* **[resticprofile OR restic command]**: Like snapshots, backup, check, prune, forget, mount, etc.
* **[additional flags]**: Any additional flags to pass to the restic command line

{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
