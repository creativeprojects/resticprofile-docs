---
title: Copy command
weight: 17
---



## Special case for the `copy` command section

The copy command needs two repositories (and quite likely 2 different set of keys). You can configure a `copy` section like this:


{{< tabs groupid="config-with-hcl" >}}
{{% tab title="toml" %}}

```toml
[default]
  initialize = false
  repository = "/backup/original"
  password-file = "key"

  [default.copy]
    initialize = true
    repository = "/backup/copy"
    password-file = "other_key"
```

{{% /tab %}}
{{% tab title="yaml" %}}

```yaml
default:
    initialize: false
    repository: "/backup/original"
    password-file: key
    copy:
        initialize: true
        repository: "/backup/copy"
        password-file: other_key
```

{{% /tab %}}
{{% tab title="hcl" %}}


```hcl

default {
    initialize = false
    repository = "/backup/original"
    password-file = "key"

    copy = {
        initialize = true
        repository = "/backup/copy"
        password-file = "other_key"
    }
}
```

{{% /tab %}}
{{% /tabs %}}

You will note that the secondary repository doesn't need to have a `2` behind its flags (`repository2`, `password-file2`, etc.). It's because the flags are well separated in the configuration and there's no ambiguity.

## Initialisation

If you want to initialize the *copy* repository using the `copy-chunker-params` flag, it needs to be called `initialize-copy-chunker-params` instead. As such, this flag does not exist on the `copy` target which is why we need to prefix it.


{{< tabs groupid="config-with-hcl" >}}
{{% tab title="toml" %}}

```toml
[profile]
  initialize = false
  repository = "/backup/original"
  password-file = "key"

  [profile.copy]
    initialize = true
    initialize-copy-chunker-params = true
    repository = "/backup/copy"
    password-file = "other_key"
```

{{% /tab %}}
{{% tab title="yaml" %}}

```yaml
profile:
    initialize: false
    repository: "/backup/original"
    password-file: key
    copy:
        initialize: true
        initialize-copy-chunker-params: true
        repository: "/backup/copy"
        password-file: other_key
```

{{% /tab %}}
{{% tab title="hcl" %}}


```hcl

profile {
    initialize = false
    repository = "/backup/original"
    password-file = "key"

    copy = {
        initialize = true
        initialize-copy-chunker-params = true
        repository = "/backup/copy"
        password-file = "other_key"
    }
}
```

{{% /tab %}}
{{% /tabs %}}

{{< pageversions "v0.18.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" >}}
