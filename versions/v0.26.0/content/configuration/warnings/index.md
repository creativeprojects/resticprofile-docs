---
title: Warnings
weight: 30
---

## Warnings from restic

Until version 0.13.0, resticprofile was always considering a restic warning as an error. This will remain the **default**.
But the version 0.13.0 introduced a parameter to avoid this behaviour and consider that the backup was successful instead.

A restic warning occurs when it cannot read some files, but a snapshot was successfully created.

### no-error-on-warning

{{< tabs groupid="config-with-json" >}}
{{% tab title="toml" %}}

```toml
version = "1"

[profile]

  [profile.backup]
    no-error-on-warning = true

```

{{% /tab %}}
{{% tab title="yaml" %}}


```yaml
version: "1"

profile:
    backup:
        no-error-on-warning: true
```

{{% /tab %}}
{{% tab title="hcl" %}}

```hcl
"profile" = {

  "backup" = {
    "no-error-on-warning" = true
  }
}
```

{{% /tab %}}
{{% tab title="json" %}}

```json
{
  "version": "1",
  "profile": {
    "backup": {
      "no-error-on-warning": true
    }
  }
}
```

{{% /tab %}}
{{< /tabs >}}

{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
