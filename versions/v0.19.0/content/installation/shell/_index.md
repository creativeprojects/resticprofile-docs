---
title: Shell Completion
weight: 100
---


Shell command line completions are provided for `bash` and `zsh`. 

To load the command completions in shell, use:

```shell
# bash
eval "$(resticprofile generate --bash-completion)"

# zsh
eval "$(resticprofile generate --zsh-completion)"
```

To install them permanently:

```shell
$ resticprofile generate --bash-completion > /etc/bash_completion.d/resticprofile
$ chmod +x /etc/bash_completion.d/resticprofile
```

{{< pageversions "v0.18.0" "v0.20.0" "v0.21.1" "v0.22.0" >}}
