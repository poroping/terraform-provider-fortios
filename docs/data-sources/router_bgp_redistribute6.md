---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_redistribute6"
description: |-
  Get information on a fortios BGP IPv6 redistribute table.
---

# Data Source: fortios_router_bgp_redistribute6
Use this data source to get information on a fortios BGP IPv6 redistribute table.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Distribute list entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Distribute list entry name.
* `route_map` - Route map name.
* `status` - Status
