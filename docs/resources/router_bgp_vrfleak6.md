---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_vrfleak6"
description: |-
  BGP IPv6 VRF leaking table.
---

## fortios_router_bgp_vrfleak6
BGP IPv6 VRF leaking table.

~> This resource is configuring a child table of the parent resource: `fortios_router_bgp`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `vrf` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `vrf` - Origin VRF ID (0 - 31).
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block contains:

* `interface` - Interface which is used to leak routes to target VRF. This attribute must reference one of the following datasources: `system.interface.name` .
* `route_map` - Route map of VRF leaking. This attribute must reference one of the following datasources: `router.route-map.name` .
* `vrf` - Target VRF ID (0 - 31).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_vrfleak6 can be imported using:
```sh
terraform import fortios_router_vrfleak6.labelname {{mkey}}
```
