---
title: Warnings
weight: 30
---

## Warnings from restic

Until version 0.13.0, resticprofile was always considering a restic warning as an error. This will remain the **default**.
But the version 0.13.0 introduced a parameter to avoid this behaviour and consider that the backup was successful instead.

A restic warning occurs when it cannot read some files, but a snapshot was successfully created.

### Let me introduce no-error-on-warning

{{< tabs groupid="config-with-json" >}}
{{% tab title="toml" %}}

```toml
[profile]
  inherit = "default"

  [profile.backup]
    no-error-on-warning = true

```

{{% /tab %}}
{{% tab title="yaml" %}}


```yaml
profile:
    inherit: default
    backup:
        no-error-on-warning: true
```

{{% /tab %}}
{{% tab title="hcl" %}}

```hcl
"profile" = {
  "inherit" = "default"

  "backup" = {
    "no-error-on-warning" = true
  }
}
```

{{% /tab %}}
{{% tab title="json" %}}

```json
{
  "profile": {
    "inherit": "default",
    "backup": {
      "no-error-on-warning": true
    }
  }
}
```

{{% /tab %}}
{{% /tabs %}}

{{< pageversions "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
