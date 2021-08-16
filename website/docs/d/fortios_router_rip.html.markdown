---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip"
description: |-
  Get information on a fortios Configure RIP.
---

# Data Source: fortios_router_rip
Use this data source to get information on a fortios Configure RIP.

## Example Usage

```hcl
data "fortios_router_rip" sample1 {
}

output output1 {
  value = data.fortios_router_rip.sample1
}
```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `default_information_originate` - Enable/disable generation of default route.
* `default_metric` - Default metric.
* `garbage_timer` - Garbage timer in seconds.
* `max_out_metric` - Maximum metric allowed to output(0 means 'not set').
* `recv_buffer_size` - Receiving buffer size.
* `timeout_timer` - Timeout timer in seconds.
* `update_timer` - Update timer in seconds.
* `version` - RIP version.
* `distance` - distanceThe structure of `distance` block is documented below.

The `distance` block contains:

* `access_list` - Access list for route destination.
* `distance` - Distance (1 - 255).
* `id` - Distance ID.
* `prefix` - Distance prefix.
* `distribute_list` - Distribute list.The structure of `distribute_list` block is documented below.

The `distribute_list` block contains:

* `direction` - Distribute list direction.
* `id` - Distribute list ID.
* `interface` - Distribute list interface name.
* `listname` - Distribute access/prefix list name.
* `status` - status
* `interface` - RIP interface configuration.The structure of `interface` block is documented below.

The `interface` block contains:

* `auth_keychain` - Authentication key-chain name.
* `auth_mode` - Authentication mode.
* `auth_string` - Authentication string/password.
* `flags` - flags
* `name` - Interface name.
* `receive_version` - Receive version.
* `send_version` - Send version.
* `send_version2_broadcast` - Enable/disable broadcast version 1 compatible packets.
* `split_horizon` - Enable/disable split horizon.
* `split_horizon_status` - Enable/disable split horizon.
* `neighbor` - neighborThe structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `id` - Neighbor entry ID.
* `ip` - IP address.
* `network` - networkThe structure of `network` block is documented below.

The `network` block contains:

* `id` - Network entry ID.
* `prefix` - Network prefix.
* `offset_list` - Offset list.The structure of `offset_list` block is documented below.

The `offset_list` block contains:

* `access_list` - Access list name.
* `direction` - Offset list direction.
* `id` - Offset-list ID.
* `interface` - Interface name.
* `offset` - offset
* `status` - status
* `passive_interface` - Passive interface configuration.The structure of `passive_interface` block is documented below.

The `passive_interface` block contains:

* `name` - Passive interface name.
* `redistribute` - Redistribute configuration.The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `metric` - Redistribute metric setting.
* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - status
