---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_neighborrange6"
description: |-
  BGP IPv6 neighbor range table.
---

## fortios_router_bgp_neighborrange6
BGP IPv6 neighbor range table.

~> This resource is configuring a child table of the parent resource: `fortios_router_bgp`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - IPv6 neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name. This attribute must reference one of the following datasources: `router.bgp.neighbor-group.name` .
* `prefix6` - IPv6 prefix.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_neighborrange6 can be imported using:
```sh
terraform import fortios_router_neighborrange6.labelname {{mkey}}
```
