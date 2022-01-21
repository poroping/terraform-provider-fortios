---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_arptable"
description: |-
  Get information on a fortios Configure ARP table.
---

# Data Source: fortios_system_arptable
Use this data source to get information on a fortios Configure ARP table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Unique integer ID of the entry.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Unique integer ID of the entry.
* `interface` - Interface name.
* `ip` - IP address.
* `mac` - MAC address.
