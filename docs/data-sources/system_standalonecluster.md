---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_standalonecluster"
description: |-
  Get information on a fortios Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.
---

# Data Source: fortios_system_standalonecluster
Use this data source to get information on a fortios Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `encryption` - Enable/disable encryption when synchronizing sessions.
* `group_member_id` - Cluster member ID (0 - 15).
* `layer2_connection` - Indicate whether layer 2 connections are present among FGSP members.
* `psksecret` - Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `standalone_group_id` - Cluster group ID (0 - 255). Must be the same for all members.
* `cluster_peer` - Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.The structure of `cluster_peer` block is documented below.

The `cluster_peer` block contains:

* `hb_interval` - Heartbeat interval (1 - 20 (100*ms). Increase to reduce false positives.
* `hb_lost_threshold` - Lost heartbeat threshold (1 - 60). Increase to reduce false positives.
* `ipsec_tunnel_sync` - Enable/disable IPsec tunnel synchronization.
* `peerip` - IP address of the interface on the peer unit that is used for the session synchronization link.
* `peervd` - VDOM that contains the session synchronization link interface on the peer unit. Usually both peers would have the same peervd.
* `secondary_add_ipsec_routes` - Enable/disable IKE route announcement on the backup unit.
* `sync_id` - Sync ID.
* `down_intfs_before_sess_sync` - List of interfaces to be turned down before session synchronization is complete.The structure of `down_intfs_before_sess_sync` block is documented below.

The `down_intfs_before_sess_sync` block contains:

* `name` - Interface name.
* `session_sync_filter` - Add one or more filters if you only want to synchronize some sessions. Use the filter to configure the types of sessions to synchronize.The structure of `session_sync_filter` block is documented below.

The `session_sync_filter` block contains:

* `dstaddr` - Only sessions to this IPv4 address are synchronized.
* `dstaddr6` - Only sessions to this IPv6 address are synchronized.
* `dstintf` - Only sessions to this interface are synchronized.
* `srcaddr` - Only sessions from this IPv4 address are synchronized.
* `srcaddr6` - Only sessions from this IPv6 address are synchronized.
* `srcintf` - Only sessions from this interface are synchronized.
* `custom_service` - Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services.The structure of `custom_service` block is documented below.

The `custom_service` block contains:

* `dst_port_range` - Custom service destination port range.
* `id` - Custom service ID.
* `src_port_range` - Custom service source port range.
* `syncvd` - Sessions from these VDOMs are synchronized using this session synchronization configuration.The structure of `syncvd` block is documented below.

The `syncvd` block contains:

* `name` - VDOM name.
