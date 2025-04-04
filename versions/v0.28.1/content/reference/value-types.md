---
title: Value types
weight: 25
---


**Type**: `true` / `false`
: Is a boolean value to toggle a flag or specific behaviour that has
to match the syntax for booleans in the used file format.

**Type**: `numeric` & `integer`
: Is a numeric value (whole-number for `integer`) that has to match the
syntax for numbers in the used file format.

**Type**: `string`
: Is a sequence of UTF-8 characters that usually have to be placed in
quotes and must match the syntax for strings in the used file format.

**Type**: `duration`
: A duration `string` is a sequence of decimal numbers, each with optional
fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid
time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

**Type**: `uri`
: An uri `string` must be a valid URI or URL.

**Type**: `hostname`
: A hostname `string` must be a valid hostname or FQDN.

**Type**: one or more `[type]s`
: Indicates that multiple values of `[type]` can be specified as a list.
The list syntax depends on the used file format.

**Type**: one or more nested `[type]`
: Is a value or a list of values that follow the configuration structure declared
in `[type]`.

**Type**: `key` => `[type]`
: Is a value that is configuration structure of `string` keys and values of `[type]`.

**Type**: `key` => `value`
: Is a value that is configuration structure of `string` keys and values of any type.

{{< pageversions "v0.27.1" "v0.29.1" >}}
