---
title: "Schedule Commands"
weight: 20
---


resticprofile accepts these internal commands:
- `schedule`
- `unschedule`
- `status`

All internal commands either operate on the profile selected by `--name`, on the profiles selected by a group, or on all profiles when the flag `--all` is passed.

Examples:
```shell
$ resticprofile --name profile schedule 
$ resticprofile --name group schedule 
$ resticprofile schedule --all 
```

Please note, schedules are always independent of each other no matter whether they have been created with `--all`, by group or from a single profile.

### schedule command

Install all the schedules defined on the selected profile or profiles.

Please note on systemd, we need to `start` the timer once to enable it. Otherwise it will only be enabled on the next reboot. If you **dont' want** to start (and enable) it now, pass the `--no-start` flag to the command line.

Also if you use the `--all` flag to schedule all your profiles at once, make sure you use only the `user` mode or `system` mode. A combination of both would not schedule the tasks properly:
- if the user is not privileged, only the `user` tasks will be scheduled
- if the user **is** privileged, **all schedule will end-up as a `system` schedule**

### unschedule command

Remove all the schedules defined on the selected profile or profiles.

### status command

Print the status on all the installed schedules of the selected profile or profiles. 

The display of the `status` command will be OS dependant. Please see the examples below on which output you can expect from it.

{{< pageversions "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
