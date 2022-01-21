---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip"
description: |-
  Configure RIP.
---

## fortios_router_rip
Configure RIP.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable` `disable` .
* `default_metric` - Default metric.
* `garbage_timer` - Garbage timer in seconds.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `recv_buffer_size` - Receiving buffer size.
* `timeout_timer` - Timeout timer in seconds.
* `update_timer` - Update timer in seconds.
* `version` - RIP version. Valid values: `1` `2` .
* `distance` - distance The structure of `distance` block is documented below.

The `distance` block contains:

* `access_list` - Access list for route destination. This attribute must reference one of the following datasources: `router.access-list.name` .
* `distance` - Distance (1 - 255).
* `id` - Distance ID.
* `prefix` - Distance prefix.
* `distribute_list` - Distribute list. The structure of `distribute_list` block is documented below.

The `distribute_list` block contains:

* `direction` - Distribute list direction. Valid values: `in` `out` .
* `id` - Distribute list ID.
* `interface` - Distribute list interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `listname` - Distribute access/prefix list name. This attribute must reference one of the following datasources: `router.access-list.name` `router.prefix-list.name` .
* `status` - status Valid values: `enable` `disable` .
* `interface` - RIP interface configuration. The structure of `interface` block is documented below.

The `interface` block contains:

* `auth_keychain` - Authentication key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `auth_mode` - Authentication mode. Valid values: `none` `text` `md5` .
* `auth_string` - Authentication string/password.
* `flags` - flags
* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `receive_version` - Receive version. Valid values: `1` `2` .
* `send_version` - Send version. Valid values: `1` `2` .
* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets. Valid values: `disable` `enable` .
* `split_horizon` - Enable/disable split horizon. Valid values: `poisoned` `regular` .
* `split_horizon_status` - Enable/disable split horizon. Valid values: `enable` `disable` .
* `neighbor` - neighbor The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `id` - Neighbor entry ID.
* `ip` - IP address.
* `network` - network The structure of `network` block is documented below.

The `network` block contains:

* `id` - Network entry ID.
* `prefix` - Network prefix.
* `offset_list` - Offset list. The structure of `offset_list` block is documented below.

The `offset_list` block contains:

* `access_list` - Access list name. This attribute must reference one of the following datasources: `router.access-list.name` .
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

fortios_router_rip can be imported using:
```sh
terraform import fortios_router_rip.labelname {{mkey}}
```
