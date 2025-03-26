---
title: Include
weight: 15
---

The configuration may be split into multiple files by adding `includes = "glob-pattern"` to the main configuration file. 
E.g. the following `profiles.conf` loads configurations from `conf.d` and `profiles.d`:

{{< tabs groupid="config-with-json" >}}
{{% tab title="toml" %}}

```toml
# Includes
includes = ["conf.d/*.conf", "profiles.d/*.yaml", "profiles.d/*.toml"]

# Defaults
[global]
  initialize = true
```


{{% /tab %}}
{{% tab title="yaml" %}}

```yaml
        
includes:
  - "conf.d/*.conf"
  - "profiles.d/*.yaml"
  - "profiles.d/*.toml"

global:
  initialize: true

```

{{% /tab %}}
{{% tab title="hcl" %}}

```hcl

includes = ["conf.d/*.conf", "profiles.d/*.yaml", "profiles.d/*.toml"]

global {
  initialize = true
}
```

{{% /tab %}}
{{% tab title="json" %}}

```json
{
  "includes": [
    "conf.d/*.conf",
    "profiles.d/*.yaml",
    "profiles.d/*.toml"
  ],
  "global": {
    "initialize": true
  }
}
```

{{% /tab %}}
{{% /tabs %}}


Included configuration files may use any supported format and settings are merged so that multiple files can extend the same profiles.
The HCL format is special in that it cannot be mixed with other formats.

Included files cannot include nested files. Specifying `includes` inside an included file has no effect.

Within included files, the current [configuration path]({{% relref "/configuration/#path-resolution-in-configuration" %}}) is not changed. Path resolution remains relative to the path of the main configuration file.

## Configuration Merging

Loading a configuration file involves loading the physical file from disk and applying all [variables]({{% relref "/configuration/variables" %}}) and [templates]({{% relref "/configuration/templates" %}}) prior to parsing the file in a supported format `hcl`, `json`, `toml` and `yaml`. This means [variables]({{% relref "/configuration/variables" %}}) and [templates]({{% relref "/configuration/templates" %}}) must create valid configuration markup that can be parsed or loading will fail.

Configuration files are loaded and applied in a fixed order:

1. The main configuration file is loaded first
2. `includes` are iterated in declaration order:
   * Every item may be a single file path or glob expression
   * Glob expressions are resolved and iterated in alphabetical order
   * All paths are resolved relative to [configuration path]({{% relref "/configuration/#path-resolution-in-configuration" %}})

Configuration files are loaded in the following order when assuming `/etc/resticprofile/profiles.conf` with `includes = ["first.conf", "conf.d/*.conf", "last.conf"]`:
```
/etc/resticprofile/profiles.conf
/etc/resticprofile/first.conf
/etc/resticprofile/conf.d/00_a.conf
/etc/resticprofile/conf.d/01_a.conf
/etc/resticprofile/conf.d/01_b.conf
/etc/resticprofile/last.conf
```

Configuration **merging** follows the logic:

* Configuration properties are replaced
* Configuration structure (tree) is merged
* What includes later overrides what defines earlier
* Lists of values or lists of objects are considered properties not config structure and will be replaced


{{< tabs groupid="include-merging-example" >}}
{{% tab title="Final configuration" %}}

```yaml
includes:
  - first.yaml
  - second.yaml

default:
  initialize: true
  backup:
     exclude:
        - .*
     source:
        - /etc
        - /opt
```

{{% /tab %}}
{{% tab title="profiles.yaml" %}}

```yaml
includes:
  - first.yaml
  - second.yaml

default:
   
   backup:
      source:
         - /usr


        
```

{{% /tab %}}
{{% tab title="first.yaml" %}}

```yaml
        



default:
   initialize: false
   backup:
      source:
         - /etc
         - /opt

        
```

{{% /tab %}}
{{% tab title="second.yaml" %}}

```yaml
        



default:
   initialize: true
   backup:
      exclude:
         - .*


        
```

{{% /tab %}}
{{% /tabs %}}


{{% notice style="note" %}}

`resticprofile` prior to v0.18.0 had a slightly different behavior when merging configuration properties of a different type (e.g. number <-> text or list <-> single value). In such cases the existing value was not overridden by an included file, breaking the rule "what includes later overrides what defines earlier".

{{% /notice %}}


{{< pageversions "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" >}}
