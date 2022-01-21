---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ripng"
description: |-
  Configure RIPng.
---

## fortios_router_ripng
Configure RIPng.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable` `disable` .
* `default_metric` - Default metric.
* `garbage_timer` - Garbage timer.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `timeout_timer` - Timeout timer.
* `update_timer` - Update timer.
* `aggregate_address` - Aggregate address. The structure of `aggregate_address` block is documented below.

The `aggregate_address` block contains:

* `id` - Aggregate address entry ID.
* `prefix6` - Aggregate address prefix.
* `distance` - distance The structure of `distance` block is documented below.

The `distance` block contains:

* `access_list6` - Access list for route destination. This attribute must reference one of the following datasources: `router.access-list6.name` .
* `distance` - Distance (1 - 255).
* `id` - Distance ID.
* `prefix6` - Distance prefix6.
* `distribute_list` - Distribute list. The structure of `distribute_list` block is documented below.

The `distribute_list` block contains:

* `direction` - Distribute list direction. Valid values: `in` `out` .
* `id` - Distribute list ID.
* `interface` - Distribute list interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `listname` - Distribute access/prefix list name. This attribute must reference one of the following datasources: `router.access-list6.name` `router.prefix-list6.name` .
* `status` - status Valid values: `enable` `disable` .
* `interface` - RIPng interface configuration. The structure of `interface` block is documented below.

The `interface` block contains:

* `flags` - Flags.
* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned` `regular` .
* `split_horizon_status` - Enable/disable split horizon. Valid values: `enable` `disable` .
* `neighbor` - neighbor The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `id` - Neighbor entry ID.
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip6` - IPv6 link-local address.
* `network` - Network. The structure of `network` block is documented below.

The `network` block contains:

* `id` - Network entry ID.
* `prefix` - Network IPv6 link-local prefix.
* `offset_list` - Offset list. The structure of `offset_list` block is documented below.

The `offset_list` block contains:

* `access_list6` - IPv6 access list name. This attribute must reference one of the following datasources: `router.access-list6.name` .
* `direction` - Offset list direction. Valid values: `in` `out` .
* `id` - Offset-list ID.
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `offset` - offset
* `status` - status Valid values: `enable` `disable` .
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.

The `passive_interface` block contains:

* `name` - Passive interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `metric` - Redistribute metric setting.
* `name` - Redistribute name.
* `routemap` - Route map name. This attribute must reference one of the following datasources: `router.route-map.name` .
* `status` - status Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_ripng can be imported using:
```sh
terraform import fortios_router_ripng.labelname {{mkey}}
```
