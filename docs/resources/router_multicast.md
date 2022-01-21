---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicast"
description: |-
  Configure router multicast.
---

## fortios_router_multicast
Configure router multicast.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `multicast_routing` - Enable/disable IP multicast routing. Valid values: `enable` `disable` .
* `route_limit` - Maximum number of multicast routes.
* `route_threshold` - Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.
* `interface` - PIM interfaces. The structure of `interface` block is documented below.

The `interface` block contains:

* `bfd` - Enable/disable Protocol Independent Multicast (PIM) Bidirectional Forwarding Detection (BFD). Valid values: `enable` `disable` .
* `cisco_exclude_genid` - Exclude GenID from hello packets (compatibility with old Cisco IOS). Valid values: `enable` `disable` .
* `dr_priority` - DR election priority.
* `hello_holdtime` - Time before old neighbor information expires (0 - 65535 sec, default = 105).
* `hello_interval` - Interval between sending PIM hello messages (0 - 65535 sec, default = 30).
* `multicast_flow` - Acceptable source for multicast group. This attribute must reference one of the following datasources: `router.multicast-flow.name` .
* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `neighbour_filter` - Routers acknowledged as neighbor routers. This attribute must reference one of the following datasources: `router.access-list.name` .
* `passive` - Enable/disable listening to IGMP but not participating in PIM. Valid values: `enable` `disable` .
* `pim_mode` - PIM operation mode. Valid values: `sparse-mode` `dense-mode` .
* `propagation_delay` - Delay flooding packets on this interface (100 - 5000 msec, default = 500).
* `rp_candidate` - Enable/disable compete to become RP in elections. Valid values: `enable` `disable` .
* `rp_candidate_group` - Multicast groups managed by this RP. This attribute must reference one of the following datasources: `router.access-list.name` .
* `rp_candidate_interval` - RP candidate advertisement interval (1 - 16383 sec, default = 60).
* `rp_candidate_priority` - Router's priority as RP.
* `rpf_nbr_fail_back` - Enable/disable fail back for RPF neighbor query. Valid values: `enable` `disable` .
* `rpf_nbr_fail_back_filter` - Filter for fail back RPF neighbors. This attribute must reference one of the following datasources: `router.access-list.name` .
* `state_refresh_interval` - Interval between sending state-refresh packets (1 - 100 sec, default = 60).
* `static_group` - Statically set multicast groups to forward out. This attribute must reference one of the following datasources: `router.multicast-flow.name` .
* `ttl_threshold` - Minimum TTL of multicast packets that will be forwarded (applied only to new multicast routes) (1 - 255, default = 1).
* `igmp` - IGMP configuration options. The structure of `igmp` block is documented below.

The `igmp` block contains:

* `access_group` - Groups IGMP hosts are allowed to join. This attribute must reference one of the following datasources: `router.access-list.name` .
* `immediate_leave_group` - Groups to drop membership for immediately after receiving IGMPv2 leave. This attribute must reference one of the following datasources: `router.access-list.name` .
* `last_member_query_count` - Number of group specific queries before removing group (2 - 7, default = 2).
* `last_member_query_interval` - Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).
* `query_interval` - Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).
* `query_max_response_time` - Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).
* `query_timeout` - Timeout between queries before becoming querier for network (60 - 900, default = 255).
* `router_alert_check` - Enable/disable require IGMP packets contain router alert option. Valid values: `enable` `disable` .
* `version` - Maximum version of IGMP to support. Valid values: `3` `2` `1` .
* `join_group` - Join multicast groups. The structure of `join_group` block is documented below.

The `join_group` block contains:

* `address` - Multicast group IP address.
* `pim_sm_global` - PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.

The `pim_sm_global` block contains:

* `accept_register_list` - Sources allowed to register packets with this Rendezvous Point (RP). This attribute must reference one of the following datasources: `router.access-list.name` .
* `accept_source_list` - Sources allowed to send multicast traffic. This attribute must reference one of the following datasources: `router.access-list.name` .
* `bsr_allow_quick_refresh` - Enable/disable accept BSR quick refresh packets from neighbors. Valid values: `enable` `disable` .
* `bsr_candidate` - Enable/disable allowing this router to become a bootstrap router (BSR). Valid values: `enable` `disable` .
* `bsr_hash` - BSR hash length (0 - 32, default = 10).
* `bsr_interface` - Interface to advertise as candidate BSR. This attribute must reference one of the following datasources: `system.interface.name` .
* `bsr_priority` - BSR priority (0 - 255, default = 0).
* `cisco_crp_prefix` - Enable/disable making candidate RP compatible with old Cisco IOS. Valid values: `enable` `disable` .
* `cisco_ignore_rp_set_priority` - Use only hash for RP selection (compatibility with old Cisco IOS). Valid values: `enable` `disable` .
* `cisco_register_checksum` - Checksum entire register packet(for old Cisco IOS compatibility). Valid values: `enable` `disable` .
* `cisco_register_checksum_group` - Cisco register checksum only these groups. This attribute must reference one of the following datasources: `router.access-list.name` .
* `join_prune_holdtime` - Join/prune holdtime (1 - 65535, default = 210).
* `message_interval` - Period of time between sending periodic PIM join/prune messages in seconds (1 - 65535, default = 60).
* `null_register_retries` - Maximum retries of null register (1 - 20, default = 1).
* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 - 65535, default = 0 which means unlimited).
* `register_rp_reachability` - Enable/disable check RP is reachable before registering packets. Valid values: `enable` `disable` .
* `register_source` - Override source address in register packets. Valid values: `disable` `interface` `ip-address` .
* `register_source_interface` - Override with primary interface address. This attribute must reference one of the following datasources: `system.interface.name` .
* `register_source_ip` - Override with local IP address.
* `register_supression` - Period of time to honor register-stop message (1 - 65535 sec, default = 60).
* `rp_register_keepalive` - Timeout for RP receiving data on (S,G) tree (1 - 65535 sec, default = 185).
* `spt_threshold` - Enable/disable switching to source specific trees. Valid values: `enable` `disable` .
* `spt_threshold_group` - Groups allowed to switch to source tree. This attribute must reference one of the following datasources: `router.access-list.name` .
* `ssm` - Enable/disable source specific multicast. Valid values: `enable` `disable` .
* `ssm_range` - Groups allowed to source specific multicast. This attribute must reference one of the following datasources: `router.access-list.name` .
* `rp_address` - Statically configure RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block contains:

* `group` - Groups to use this RP. This attribute must reference one of the following datasources: `router.access-list.name` .
* `id` - ID.
* `ip_address` - RP router address.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_multicast can be imported using:
```sh
terraform import fortios_router_multicast.labelname {{mkey}}
```
