---
title: Configuration v2 proposal
weight: 60
---

{{% notice style="note" %}}

The configuration file format `v2` is in preview right now. You can try to use it since `v0.17.0` but it's for testing and feedback only.

{{% /notice %}}


## Introduction

The current file format was decided at the time resticprofile was only using the `toml` format. Nesting pieces of configuration in blocks is not the easiest as you have to specify the whole path in the block:

```toml
[profile]

[profile.backup]
source = "some path"

```

Since then, I believe the `yaml` format is preferred over `toml`.

My proposal is to make a "version 2" of the configuration file, the current file format is "version 1".

**Both formats will continue to be valid**:

- if no version is specified, the "version 1" is used. This is the current format
- if a version is specified (`2`) the new format will be expected

## New format availability

The new format "version 2" will be available for:
- TOML
- YAML
- JSON

**It won't be available for HCL**. This may not be definitive, but it's not widely used and it's becoming more and more difficult to support HCL.

**HCL can still be used as is, "version = 1"**

## New format specifications

I will show the specification using the `yaml` as examples, because it's probably the most readable format.

### version

```yaml
version: 2
```

### global

The global section does not change. We'll keep all the global configuration in there.

```yaml
global:
    default-command: snapshots
    initialize: false
    priority: low

```

### profiles

All your profiles will be nested under a `profiles` section. Please note the schedules are no longer described inside the profile, but in a separate section `schedules` (see the following example).

```yaml
version: "2"

profiles:
    default:
        env:
            tmp: /tmp
        password-file: key
        repository: /backup

    documents:
        inherit: default
        backup:
            source: ~/Documents
        snapshots:
            tag:
                - documents
```

### groups

The list of profiles will be nested under a `profiles` section, so we can add more configuration to groups later.

```yaml
version: "2"

groups:
    full: # name of your group
        profiles:
            - root
            - documents
            - mysql
```

### schedules

A new schedule section could schedule either a group or a list of profiles.

```yaml
version: "2"

schedules:
    full-backup: # give a name to your schedule
        group: full
        schedule:
            - "Mon..Fri *:00,15,30,45" # every 15 minutes on weekdays
        permission: user
        run: backup # backup is the default if not specified

    other:
        profiles:
            - root
            - mysql
        schedule:
            - "Sat,Sun 0,12:00" # twice a day on week-ends
        permission: user
        run: prune
```

This format leaves more space for improvements later (like a `repos` section maybe?)

{{% notice style="tip" %}}
You can participate in designing the "version 2" [here](https://github.com/creativeprojects/resticprofile/issues/80)
{{% /notice %}}

{{< pageversions "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
