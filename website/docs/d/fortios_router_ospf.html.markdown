---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf"
description: |-
  Get information on a fortios Configure OSPF.
---

# Data Source: fortios_router_ospf
Use this data source to get information on a fortios Configure OSPF.

## Example Usage

```hcl
data "fortios_router_ospf" sample1 {
}

output output1 {
  value = data.fortios_router_ospf.sample1
}
```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `abr_type` - Area border router type.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `database_overflow` - Enable/disable database overflow.
* `database_overflow_max_lsas` - Database overflow maximum LSAs.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type.
* `default_information_originate` - Enable/disable generation of default route.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `distance` - Distance of the route.
* `distance_external` - Administrative external distance.
* `distance_inter_area` - Administrative inter-area distance.
* `distance_intra_area` - Administrative intra-area distance.
* `distribute_list_in` - Filter incoming routes.
* `distribute_route_map_in` - Filter incoming external routes by route-map.
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes
* `restart_mode` - OSPF restart mode (graceful or LLS).
* `restart_period` - Graceful restart period.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility.
* `router_id` - Router ID.
* `spf_timers` - SPF calculation frequency.
* `area` - OSPF area configuration.The structure of `area` block is documented below.

The `area` block contains:

* `authentication` - Authentication type.
* `comments` - Comment.
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area.
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `shortcut` - Enable/disable shortcut option.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `filter_list` - OSPF area filter-list configuration.The structure of `filter_list` block is documented below.

The `filter_list` block contains:

* `direction` - Direction.
* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name.
* `range` - OSPF area range configuration.The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Range entry ID.
* `prefix` - Prefix.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status.
* `virtual_link` - OSPF virtual link configuration.The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `keychain` - Message-digest key-chain name.
* `name` - Virtual link entry name.
* `peer` - Peer IP.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `md5_keys` - MD5 key.The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.
* `distribute_list` - Distribute list configuration.The structure of `distribute_list` block is documented below.

The `distribute_list` block contains:

* `access_list` - Access list name.
* `id` - Distribute list entry ID.
* `protocol` - Protocol type.
* `neighbor` - OSPF neighbor configuration are used when OSPF runs on non-broadcast mediaThe structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.
* `network` - OSPF network configuration.The structure of `network` block is documented below.

The `network` block contains:

* `area` - Attach the network to area.
* `comments` - Comment.
* `id` - Network entry ID.
* `prefix` - Prefix.
* `ospf_interface` - OSPF interface configuration.The structure of `ospf_interface` block is documented below.

The `ospf_interface` block contains:

* `authentication` - Authentication type.
* `authentication_key` - Authentication key.
* `bfd` - Bidirectional Forwarding Detection (BFD).
* `comments` - Comment.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `database_filter_out` - Enable/disable control of flooding out LSAs.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `hello_multiplier` - Number of hello packets within dead interval.
* `interface` - Configuration interface name.
* `ip` - IP address.
* `keychain` - Message-digest key-chain name.
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU.
* `name` - Interface entry name.
* `network_type` - Network type.
* `prefix_length` - Prefix length.
* `priority` - Priority.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable status.
* `transmit_delay` - Transmit delay.
* `md5_keys` - MD5 key.The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.
* `passive_interface` - Passive interface configuration.The structure of `passive_interface` block is documented below.

The `passive_interface` block contains:

* `name` - Passive interface name.
* `redistribute` - Redistribute configuration.The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `metric` - Redistribute metric setting.
* `metric_type` - Metric type.
* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - status
* `tag` - Tag value.
* `summary_address` - IP address summary configuration.The structure of `summary_address` block is documented below.

The `summary_address` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.
