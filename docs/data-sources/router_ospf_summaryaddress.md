---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_summary_address"
description: |-
  Get information on a fortios IP address summary configuration.
---

# Data Source: fortios_router_ospf_summaryaddress
Use this data source to get information on a fortios IP address summary configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Summary address entry ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `advertise` - Enable/disable advertise status.
* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.
