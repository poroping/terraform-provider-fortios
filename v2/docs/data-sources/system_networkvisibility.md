---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_networkvisibility"
description: |-
  Get information on a fortios Configure network visibility settings.
---

# Data Source: fortios_system_networkvisibility
Use this data source to get information on a fortios Configure network visibility settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `destination_hostname_visibility` - Enable/disable logging of destination hostname visibility.
* `destination_location` - Enable/disable logging of destination geographical location visibility.
* `destination_visibility` - Enable/disable logging of destination visibility.
* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `source_location` - Enable/disable logging of source geographical location visibility.
