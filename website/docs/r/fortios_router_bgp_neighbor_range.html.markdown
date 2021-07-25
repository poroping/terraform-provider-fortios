---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_neighbor_range"
description: |-
  BGP neighbor range table.
---

## fortios_router_bgp_neighbor_range
BGP neighbor range table.

~> This resource is configuring a child table of the parent resource: `fortios_router_bgp`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.

## Example Usage

WIP

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `fosid` - Neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name. This attribute must reference one of the following datasources: `router.bgp.neighbor-group.name` .
* `prefix` - Neighbor range prefix.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

router_neighbor_range can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_neighbor_range.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
