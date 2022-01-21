---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis_summaryaddress6"
description: |-
  Get information on a fortios IS-IS IPv6 summary address.
---

# Data Source: fortios_router_isis_summaryaddress6
Use this data source to get information on a fortios IS-IS IPv6 summary address.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Prefix entry ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Prefix entry ID.
* `level` - Level.
* `prefix6` - IPv6 prefix.
