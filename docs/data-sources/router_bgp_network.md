---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_network"
description: |-
  Get information on a fortios BGP network table.
---

# Data Source: fortios_router_bgp_network
Use this data source to get information on a fortios BGP network table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `backdoor` - Enable/disable route as backdoor.
* `id` - ID.
* `network_import_check` - Configure insurance of BGP network route existence in IGP.
* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route.
