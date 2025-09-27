---
title: features
weight: 8
---
#### Section profile.**features**

This section configures restic command `features`  available since `0.18.0` .
Information on command and flags is copied from the [restic](https://github.com/restic/restic) manual pages.

The "features" command prints a list of supported feature flags.

To pass feature flags to restic, set the RESTIC_FEATURES environment variable
to "featureA=true,featureB=false". Specifying an unknown feature flag is an error.

A feature can either be in alpha, beta, stable or deprecated state.
An alpha feature is disabled by default and may change in arbitrary ways between restic versions or be removed.
A beta feature is enabled by default, but still can change in minor ways or be removed.
A stable feature is always enabled and cannot be disabled. The flag will be removed in a future restic version.
A deprecated feature is always disabled and cannot be enabled. The flag will be removed in a future restic version.

##### Flags used by **resticprofile** only:

| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **run-after** |one or more `strings` | |Run shell command(s) after a successful restic command |
| **run-after-fail** |one or more `strings` | |Run shell command(s) after failed restic or shell commands |
| **run-before** |one or more `strings` | |Run shell command(s) before a restic command |
| **run-finally** |one or more `strings` | |Run shell command(s) always, after all other commands |
| **send-after** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after a successful restic command |
| **send-after-fail** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) after failed restic or shell commands |
| **send-before** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) before a restic command |
| **send-finally** | one or more nested *[SendMonitoringSection](../nested/sendmonitoringsection)* | |Send HTTP request(s) always, after all other commands |





{{% notice style="tip" %}}
Flags declared for the **restic** command line in section *[profile](../profile)*
can be overridden in this section.
{{% /notice %}}

