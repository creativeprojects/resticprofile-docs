---
title: Preventing system sleep
weight: 40
---

This feature is available for:
- macOS
- Windows
- Linux with systemd ([logind](https://www.freedesktop.org/software/systemd/man/systemd-logind.service.html))

There's a `global` parameter called `prevent-sleep` that you can set to `true`, and resticprofile will prevent your system from idle sleeping.

Please note:
- it will not prevent a sleep if the system is running on batteries
- it will not prevent a sleep triggered by a user action: using the sleep button, closing the laptop lid, etc.

{{< pageversions "v0.19.0" "v0.20.0" "v0.22.0" "v0.23.0" "v0.24.0" "v0.25.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
