---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_vrf6"
description: |-
  Get information on a fortios BGP IPv6 VRF leaking table.
---

# Data Source: fortios_router_bgp_vrf6
Use this data source to get information on a fortios BGP IPv6 VRF leaking table.


## Example Usage

```hcl

```

## Argument Reference

* `vrf` - (Required) Origin VRF ID (0 - 251).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `vrf` - Origin VRF ID (0 - 251).
* `leak_target` - Target VRF table.The structure of `leak_target` block is documented below.

The `leak_target` block contains:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID (0 - 251).
