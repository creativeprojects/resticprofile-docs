---
title: SendMonitoringSection
weight: 4
---
#### Nested *SendMonitoringSection*



| Name              | Type                    | Default  | Notes |
|:------------------|:------------------------|:---------|:------|
| **body** |`string` | |Request body, overrides "body-template" |
| **body-template** |`string` | |Path to a file containing the request body (go template). See [configuration/http_hooks/#body-template]({{% relref "/configuration/http_hooks/#body-template" %}}) |
| **headers** | one or more nested *[SendMonitoringHeader](../nested/sendmonitoringheader)* | |Additional HTTP headers to send with the request |
| **method** |`string` |`GET` |HTTP method of the request. Is one of `GET`, `DELETE`, `HEAD`, `OPTIONS`, `PATCH`, `POST`, `PUT`, `TRACE`  |
| **skip-tls-verification** |`true` / `false` |`false` |Enables insecure TLS (without verification), see also "global.ca-certificates" |
| **url** |`uri` | |URL of the target to send to |



{{< pageversions "v0.27.1" "v0.28.1" >}}
