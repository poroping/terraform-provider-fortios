---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_vrf_leak"
description: |-
  Get information on a fortios BGP VRF leaking table.
---

# Data Source: fortios_router_bgp_vrfleak
Use this data source to get information on a fortios BGP VRF leaking table.


## Example Usage

```hcl

```

## Argument Reference

* `vrf` - (Required) Origin VRF ID (0 - 31).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `vrf` - Origin VRF ID (0 - 31).
* `target` - Target VRF table.The structure of `target` block is documented below.

The `target` block contains:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 31).
