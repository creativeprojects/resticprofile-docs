---
title: StreamErrorSection
weight: 5
---
#### Nested *StreamErrorSection*



| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **max-runs** |`integer` |`0` |Maximum amount of times that "run" is started ; 0 for no limit. Must be >= 0  |
| **min-matches** |`integer` |`0` |Minimum amount of times the "pattern" must match before "run" is started ; 0 for no limit. Must be >= 0  |
| **pattern** |`regex` | |A regular expression pattern that is tested against stderr of a running restic command |
| **run** |`string` | |The shell command to run when the pattern matches |



{{< pageversions "v0.27.1" "v0.28.1" >}}
