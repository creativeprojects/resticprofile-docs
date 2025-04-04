---
title: Ansible
weight: 15
---

## Installation using Ansible

Installation using Ansible is not streamlined, but here's the playbook I'm using on my servers:

<!-- {{%resources title="Playbooks" color="green" pattern=".*"/%}} -->

{{% notice style="secondary" title="Playbooks" icon="paperclip" %}}
* [resticprofile.yml](files/resticprofile.yml)
{{% /notice %}}


This is very much work in progress. Once I get a stable ansible script I should publish it to Ansible Galaxy.

The playbook is installing (or upgrading):

* latest restic binary to `/usr/local/bin`
* latest resticprofile binary to `/usr/local/bin`
* the resticprofile configuration file from a template file found in `./resticprofile/{{ inventory_hostname }}/profiles.*` to `/root/resticprofile/profiles.*`
* password files that can be encrypted using ansible vault. These files are located in `./resticprofile/{{ inventory_hostname }}/keys/*`: they will be decrypted and saved to `/root/resticprofile/`.
* other files (like files needed for `--exclude-file`, `--files-from` or anything else you need) from `./resticprofile/{{ inventory_hostname }}/copy/*` to `/root/resticprofile/`

{{% notice style="warning" %}}
The playbook needs the python package `github3`, which in turn needs a whole bunch of package including `cryptography`, which in turn needs the `libssl-dev` package (on debian) and a `rust` compiler.

I can tell you it takes forever to compile and install all this on a Raspberry Pi 3.

Please note the installation of a rust compiler is not part of this playbook.
{{% /notice %}}

### Requirements

Each target machine must have one variable named `arch` containing the resticprofile OS & Arch. You can see a list of all the available OS & Arch couples on the [releases page](https://github.com/creativeprojects/resticprofile/releases).

Typically, a binary will be distributed using this convention:

`resticprofile-[VERSION]_[OS]_[ARCH].tar.gz`

Your host variables file should declare a `arch` variable containing the `[OS]_[ARCH]` part of the file name.

#### Examples:

<!-- checkdoc-ignore -->
```yaml
arch: linux_amd64
```

or for a Raspberry pi 3+:

<!-- checkdoc-ignore -->
```yaml
arch: linux_armv7
```

Note: _I might find a way to detect this automatically at some point_

{{< pageversions "v0.18.0" "v0.19.0" "v0.20.0" "v0.21.1" "v0.22.0" "v0.23.0" "v0.24.0" "v0.26.0" "v0.27.1" "v0.28.1" "v0.29.1" >}}
