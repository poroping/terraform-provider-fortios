---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis_summaryaddress"
description: |-
  Get information on a fortios IS-IS summary addresses.
---

# Data Source: fortios_router_isis_summaryaddress
Use this data source to get information on a fortios IS-IS summary addresses.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Summary address entry ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Summary address entry ID.
* `level` - Level.
* `prefix` - Prefix.
