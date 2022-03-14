---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_neighbor_range6"
description: |-
  Get information on a fortios BGP IPv6 neighbor range table.
---

# Data Source: fortios_router_bgp_neighborrange6
Use this data source to get information on a fortios BGP IPv6 neighbor range table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) IPv6 neighbor range ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - IPv6 neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.
* `prefix6` - IPv6 prefix.
