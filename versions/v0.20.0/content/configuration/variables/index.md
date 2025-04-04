---
title: Variables
weight: 25
---


## Variable expansion in configuration file

You might want to reuse the same configuration (or bits of it) on different environments. One way of doing it is to create a generic configuration where specific bits will be replaced by a variable.

## Pre-defined variables

The syntax for using a pre-defined variable is:
```
{{ .VariableName }}
```


The list of pre-defined variables is:

| Variable | Type | Description |
|----------|------|-------------|
| **.Profile.Name** | string | Profile name |
| **.Now** | [time.Time](https://golang.org/pkg/time/) object | Now object: see explanation bellow |
| **.CurrentDir** | string | Current directory at the time resticprofile was started |
| **.ConfigDir** | string | Directory where the configuration was loaded from |
| **.TempDir** | string | OS temporary directory (might not exist) |
| **.BinaryDir** | string | Directory where resticprofile was started from (added in `v0.18.0`) |
| **.Hostname** | string | Host name |
| **.Env.{NAME}** | string | Environment variable `${NAME}` |

Environment variables are accessible using `.Env.` followed by the name of the environment variable.

Example: `{{ .Env.HOME }}` will be replaced by your home directory (on unixes). The equivalent on Windows would be `{{ .Env.USERPROFILE }}`.

For variables that are objects, you can call all public field or method on it.
For example, for the variable `.Now` you can use:
- `.Now.Day`
- `.Now.Format layout`
- `.Now.Hour`
- `.Now.Minute`
- `.Now.Month`
- `.Now.Second`
- `.Now.UTC`
- `.Now.Unix`
- `.Now.Weekday`
- `.Now.Year`
- `.Now.YearDay`


### Example

You can use a combination of inheritance and variables in the resticprofile configuration file like so:

{{< tabs groupid="config-with-json" >}}
{{% tab title="toml" %}}

```toml
[generic]
  password-file = "{{ .ConfigDir }}/{{ .Profile.Name }}-key"
  repository = "/backup/{{ .Now.Weekday }}"
  lock = "$HOME/resticprofile-profile-{{ .Profile.Name }}.lock"
  initialize = true

  [generic.backup]
    check-before = true
    exclude = [ "/**/.git" ]
    exclude-caches = true
    one-file-system = false
    run-after = "echo All Done!"
    run-before = [
        "echo Hello {{ .Env.LOGNAME }}",
        "echo current dir: {{ .CurrentDir }}",
        "echo config dir: {{ .ConfigDir }}",
        "echo profile started at {{ .Now.Format "02 Jan 06 15:04 MST" }}"
    ]
    tag = [ "{{ .Profile.Name }}", "dev" ]

  [generic.retention]
    after-backup = true
    before-backup = false
    compact = false
    keep-within = "30d"
    prune = true
    tag = [ "{{ .Profile.Name }}", "dev" ]

  [generic.snapshots]
    tag = [ "{{ .Profile.Name }}", "dev" ]

[src]
  inherit = "generic"

  [src.backup]
    source = [ "{{ .Env.HOME }}/go/src" ]
  
  [src.check]
    # Weekday is an integer from 0 to 6
    # Nice trick to add 1 to an integer: https://stackoverflow.com/a/72465098
    read-data-subset = "{{ len (printf "a%*s" .Now.Weekday "") }}/7"

```

{{% /tab %}}
{{% tab title="yaml" %}}

```yaml
---
generic:
    password-file: "{{ .ConfigDir }}/{{ .Profile.Name }}-key"
    repository: "/backup/{{ .Now.Weekday }}"
    lock: "$HOME/resticprofile-profile-{{ .Profile.Name }}.lock"
    initialize: true

    backup:
        check-before: true
        exclude:
        - /**/.git
        exclude-caches: true
        one-file-system: false
        run-after: echo All Done!
        run-before:
          - "echo Hello {{ .Env.LOGNAME }}"
          - "echo current dir: {{ .CurrentDir }}"
          - "echo config dir: {{ .ConfigDir }}"
          - "echo profile started at {{ .Now.Format "02 Jan 06 15:04 MST" }}"
        tag:
          - "{{ .Profile.Name }}"
          - dev

    retention:
        after-backup: true
        before-backup: false
        compact: false
        keep-within: 30d
        prune: true
        tag:
          - "{{ .Profile.Name }}"
          - dev

    snapshots:
        tag:
          - "{{ .Profile.Name }}"
          - dev

src:
    inherit: generic

    backup:
        source:
          - "{{ .Env.HOME }}/go/src"

    check:
        # Weekday is an integer from 0 to 6
        # Nice trick to add 1 to an integer: https://stackoverflow.com/a/72465098
        read-data-subset: "{{ len (printf "a%*s" .Now.Weekday "") }}/7"

```

{{% /tab %}}
{{% tab title="hcl" %}}

```hcl
"generic" = {
  "password-file" = "{{ .ConfigDir }}/{{ .Profile.Name }}-key"
  "repository" = "/backup/{{ .Now.Weekday }}"
  "lock" = "$HOME/resticprofile-profile-{{ .Profile.Name }}.lock"
  "initialize" = true

  "backup" = {
    "check-before" = true
    "exclude" = ["/**/.git"]
    "exclude-caches" = true
    "one-file-system" = false
    "run-after" = "echo All Done!"
    "run-before" = ["echo Hello {{ .Env.LOGNAME }}", "echo current dir: {{ .CurrentDir }}", "echo config dir: {{ .ConfigDir }}", "echo profile started at {{ .Now.Format "02 Jan 06 15:04 MST" }}"]
    "tag" = ["{{ .Profile.Name }}", "dev"]
  }

  "retention" = {
    "after-backup" = true
    "before-backup" = false
    "compact" = false
    "keep-within" = "30d"
    "prune" = true
    "tag" = ["{{ .Profile.Name }}", "dev"]
  }

  "snapshots" = {
    "tag" = ["{{ .Profile.Name }}", "dev"]
  }
}

"src" = {
  "inherit" = "generic"

  "backup" = {
    "source" = ["{{ .Env.HOME }}/go/src"]
  }

  "check" = {
    # Weekday is an integer from 0 to 6
    # Nice trick to add 1 to an integer: https://stackoverflow.com/a/72465098
    "read-data-subset" = "{{ len (printf "a%*s" .Now.Weekday "") }}/7"
  }
}
```

{{% /tab %}}
{{% tab title="json" %}}

```json
{
  "generic": {
    "password-file": "{{ .ConfigDir }}/{{ .Profile.Name }}-key",
    "repository": "/backup/{{ .Now.Weekday }}",
    "lock": "$HOME/resticprofile-profile-{{ .Profile.Name }}.lock",
    "initialize": true,
    "backup": {
      "check-before": true,
      "exclude": [
        "/**/.git"
      ],
      "exclude-caches": true,
      "one-file-system": false,
      "run-after": "echo All Done!",
      "run-before": [
        "echo Hello {{ .Env.LOGNAME }}",
        "echo current dir: {{ .CurrentDir }}",
        "echo config dir: {{ .ConfigDir }}",
        "echo profile started at {{ .Now.Format "02 Jan 06 15:04 MST" }}"
      ],
      "tag": [
        "{{ .Profile.Name }}",
        "dev"
      ]
    },
    "retention": {
      "after-backup": true,
      "before-backup": false,
      "compact": false,
      "keep-within": "30d",
      "prune": true,
      "tag": [
        "{{ .Profile.Name }}",
        "dev"
      ]
    },
    "snapshots": {
      "tag": [
        "{{ .Profile.Name }}",
        "dev"
      ]
    }
  },
  "src": {
    "inherit": "generic",
    "backup": {
      "source": [
        "{{ .Env.HOME }}/go/src"
      ]
    },
    "check": {
      "read-data-subset": "{{ len (printf "a%*s" .Now.Weekday "") }}/7"
    }
  }
}
```

{{% /tab %}}
{{% /tabs %}}

This is obviously not a real world example, but it shows many of the possibilities you can do with variable expansion.

To check the generated configuration, you can use the resticprofile `show` command:

```shell
% resticprofile -c examples/template.yaml -n src show

global:
    default-command:          snapshots
    restic-lock-retry-after:  1m0s
    restic-stale-lock-age:    2h0m0s
    min-memory:               100
    send-timeout:             30s

profile src:
    repository:     /backup/Monday
    password-file:  /Users/CP/go/src/resticprofile/examples/src-key
    initialize:     true
    lock:           /Users/CP/resticprofile-profile-src.lock

    backup:
        check-before:    true
        run-before:      echo Hello CP
                         echo current dir: /Users/CP/go/src/resticprofile
                         echo config dir: /Users/CP/go/src/resticprofile/examples
                         echo profile started at 05 Sep 22 17:39 BST
        run-after:       echo All Done!
        source:          /Users/CP/go/src
        exclude:         /**/.git
        exclude-caches:  true
        tag:             src
                         dev

    retention:
        after-backup:  true
        keep-within:   30d
        path:          /Users/CP/go/src
        prune:         true
        tag:           src
                       dev

    check:
        read-data-subset:  2/7

    snapshots:
        tag:  src
              dev
```

As you can see, the `src` profile inherited from the `generic` profile. The tags `{{ .Profile.Name }}` got replaced by the name of the current profile `src`. Now you can reuse the same generic configuration in another profile.

You might have noticed the `read-data-subset` in the `check` section which will read a seventh of the data every day, meaning the whole repository data will be checked over a week. You can find [more information about this trick](https://stackoverflow.com/a/72465098).

## Hand-made variables

But you can also define variables yourself. Hand-made variables starts with a `$` ([PHP](https://en.wikipedia.org/wiki/PHP) anyone?) and get declared and assigned with the `:=` operator ([Pascal](https://en.wikipedia.org/wiki/Pascal_(programming_language)) anyone?). Here's an example:

```yaml
# declare and assign a value to the variable
{{ $name := "something" }}

# put the content of the variable here
tag: "{{ $name }}"
```


### Example


Here's an example of a configuration on Linux where I use a variable `$mountpoint` set to a USB drive mount point:

{{< tabs groupid="config-with-json" >}}
{{% tab title="toml" %}}

```toml
[global]
  priority = "low"

{{ $mountpoint := "/mnt/external" }}

[default]
  repository = "local:{{ $mountpoint }}/backup"
  password-file = "key"
  run-before = "mount {{ $mountpoint }}"
  run-after = "umount {{ $mountpoint }}"
  run-after-fail = "umount {{ $mountpoint }}"

  [default.backup]
    exclude-caches = true
    source = [ "/etc", "/var/lib/libvirt" ]
    check-after = true
```

{{% /tab %}}
{{% tab title="yaml" %}}

```yaml
global:
  priority: low

{{ $mountpoint := "/mnt/external" }}

default:
  repository: 'local:{{ $mountpoint }}/backup'
  password-file: key
  run-before: 'mount {{ $mountpoint }}'
  run-after: 'umount {{ $mountpoint }}'
  run-after-fail: 'umount {{ $mountpoint }}'

  backup:
    exclude-caches: true
    source:
      - /etc
      - /var/lib/libvirt
    check-after: true
```

{{% /tab %}}
{{% tab title="hcl" %}}


```hcl
global {
    priority = "low"
}

{{ $mountpoint := "/mnt/external" }}

default {
    repository = "local:{{ $mountpoint }}/backup"
    password-file = "key"
    run-before = "mount {{ $mountpoint }}"
    run-after = "umount {{ $mountpoint }}"
    run-after-fail = "umount {{ $mountpoint }}"

    backup {
        exclude-caches = true
        source = [ "/etc", "/var/lib/libvirt" ]
        check-after = true
    }
}

```

{{% /tab %}}
{{% tab title="json" %}}

```json
{{ $mountpoint := "/mnt/external" }}
{
  "global": {
    "priority": "low"
  },
  "default": {
    "repository": "local:{{ $mountpoint }}/backup",
    "password-file": "key",
    "run-before": "mount {{ $mountpoint }}",
    "run-after": "umount {{ $mountpoint }}",
    "run-after-fail": "umount {{ $mountpoint }}",
    "backup": {
      "exclude-caches": true,
      "source": [
        "/etc",
        "/var/lib/libvirt"
      ],
      "check-after": true
    }
  }
}
```

{{% /tab %}}
{{% /tabs %}}

{{< pageversions "v0.18.0" "v0.19.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
