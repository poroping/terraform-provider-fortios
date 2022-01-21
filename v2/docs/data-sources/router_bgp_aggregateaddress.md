---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_aggregateaddress"
description: |-
  Get information on a fortios BGP aggregate address table.
---

# Data Source: fortios_router_bgp_aggregateaddress
Use this data source to get information on a fortios BGP aggregate address table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `as_set` - Enable/disable generate AS set path information.
* `id` - ID.
* `prefix` - Aggregate prefix.
* `summary_only` - Enable/disable filter more specific routes from updates.
