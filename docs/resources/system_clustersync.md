---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_clustersync"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.
---

## fortios_system_clustersync
Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `sync_id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `hb_interval` - Heartbeat interval (1 - 10 sec).
* `hb_lost_threshold` - Lost heartbeat threshold (1 - 10).
* `ike_heartbeat_interval` - IKE heartbeat interval (1 - 60 secs).
* `ike_monitor` - Enable/disable IKE HA monitor. Valid values: `enable` `disable` .
* `ike_monitor_interval` - IKE HA monitor interval (10 - 300 secs).
* `ike_seqjump_speed` - ESP jump ahead factor (1G - 10G pps equivalent).
* `ipsec_tunnel_sync` - Enable/disable IPsec tunnel synchronization. Valid values: `enable` `disable` .
* `peerip` - IP address of the interface on the peer unit that is used for the session synchronization link.
* `peervd` - VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd. This attribute must reference one of the following datasources: `system.vdom.name` .
* `secondary_add_ipsec_routes` - Enable/disable IKE route announcement on the backup unit. Valid values: `enable` `disable` .
* `slave_add_ike_routes` - Enable/disable IKE route announcement on the backup unit. Valid values: `enable` `disable` .
* `sync_id` - Sync ID.
* `down_intfs_before_sess_sync` - List of interfaces to be turned down before session synchronization is complete. The structure of `down_intfs_before_sess_sync` block is documented below.

The `down_intfs_before_sess_sync` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `session_sync_filter` - Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize. The structure of `session_sync_filter` block is documented below.

The `session_sync_filter` block contains:

* `dstaddr` - Only sessions to this IPv4 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized. You can only enter one address. To synchronize sessions for multiple destination addresses, add multiple filters.
* `dstintf` - Only sessions to this interface are synchronized. You can only enter one interface name. To synchronize sessions to multiple destination interfaces, add multiple filters. This attribute must reference one of the following datasources: `system.interface.name` .
* `srcaddr` - Only sessions from this IPv4 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized. You can only enter one address. To synchronize sessions from multiple source addresses, add multiple filters.
* `srcintf` - Only sessions from this interface are synchronized. You can only enter one interface name. To synchronize sessions for multiple source interfaces, add multiple filters. This attribute must reference one of the following datasources: `system.interface.name` .
* `custom_service` - Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services. The structure of `custom_service` block is documented below.

The `custom_service` block contains:

* `dst_port_range` - Custom service destination port range.
* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.
* `syncvd` - Sessions from these VDOMs are synchronized using this session synchronization configuration. The structure of `syncvd` block is documented below.

The `syncvd` block contains:

* `name` - VDOM name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_clustersync can be imported using:
```sh
terraform import fortios_system_clustersync.labelname {{mkey}}
```
