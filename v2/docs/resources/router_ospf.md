---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf"
description: |-
  Configure OSPF.
---

## fortios_router_ospf
Configure OSPF.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `abr_type` - Area border router type. Valid values: `cisco` `ibm` `shortcut` `standard` .
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `enable` `disable` .
* `database_overflow` - Enable/disable database overflow. Valid values: `enable` `disable` .
* `database_overflow_max_lsas` - Database overflow maximum LSAs.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type. Valid values: `1` `2` .
* `default_information_originate` - Enable/disable generation of default route. Valid values: `enable` `always` `disable` .
* `default_information_route_map` - Default information route map. This attribute must reference one of the following datasources: `router.route-map.name` .
* `default_metric` - Default metric of redistribute routes.
* `distance` - Distance of the route.
* `distance_external` - Administrative external distance.
* `distance_inter_area` - Administrative inter-area distance.
* `distance_intra_area` - Administrative intra-area distance.
* `distribute_list_in` - Filter incoming routes. This attribute must reference one of the following datasources: `router.access-list.name` `router.prefix-list.name` .
* `distribute_route_map_in` - Filter incoming external routes by route-map. This attribute must reference one of the following datasources: `router.route-map.name` .
* `log_neighbour_changes` - Enable logging of OSPF neighbour's changes Valid values: `enable` `disable` .
* `restart_mode` - OSPF restart mode (graceful or LLS). Valid values: `none` `lls` `graceful-restart` .
* `restart_period` - Graceful restart period.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility. Valid values: `enable` `disable` .
* `router_id` - Router ID.
* `spf_timers` - SPF calculation frequency.
* `area` - OSPF area configuration. The structure of `area` block is documented below.

The `area` block contains:

* `authentication` - Authentication type. Valid values: `none` `text` `message-digest` .
* `comments` - Comment.
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area. Valid values: `enable` `always` `disable` .
* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes. Valid values: `1` `2` .
* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `enable` `disable` .
* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate` `never` `always` .
* `shortcut` - Enable/disable shortcut option. Valid values: `disable` `enable` `default` .
* `stub_type` - Stub summary setting. Valid values: `no-summary` `summary` .
* `type` - Area type setting. Valid values: `regular` `nssa` `stub` .
* `filter_list` - OSPF area filter-list configuration. The structure of `filter_list` block is documented below.

The `filter_list` block contains:

* `direction` - Direction. Valid values: `in` `out` .
* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name. This attribute must reference one of the following datasources: `router.access-list.name` `router.prefix-list.name` .
* `range` - OSPF area range configuration. The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status. Valid values: `disable` `enable` .
* `id` - Range entry ID.
* `prefix` - Prefix.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status. Valid values: `enable` `disable` .
* `virtual_link` - OSPF virtual link configuration. The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication type. Valid values: `none` `text` `message-digest` .
* `authentication_key` - Authentication key.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `keychain` - Message-digest key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `md5_keychain` - Authentication MD5 key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `name` - Virtual link entry name.
* `peer` - Peer IP.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.
* `distribute_list` - Distribute list configuration. The structure of `distribute_list` block is documented below.

The `distribute_list` block contains:

* `access_list` - Access list name. This attribute must reference one of the following datasources: `router.access-list.name` .
* `id` - Distribute list entry ID.
* `protocol` - Protocol type. Valid values: `connected` `static` `rip` .
* `neighbor` - OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.
* `network` - OSPF network configuration. The structure of `network` block is documented below.

The `network` block contains:

* `area` - Attach the network to area.
* `comments` - Comment.
* `id` - Network entry ID.
* `prefix` - Prefix.
* `ospf_interface` - OSPF interface configuration. The structure of `ospf_interface` block is documented below.

The `ospf_interface` block contains:

* `authentication` - Authentication type. Valid values: `none` `text` `message-digest` .
* `authentication_key` - Authentication key.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `global` `enable` `disable` .
* `comments` - Comment.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `database_filter_out` - Enable/disable control of flooding out LSAs. Valid values: `enable` `disable` .
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `hello_multiplier` - Number of hello packets within dead interval.
* `interface` - Configuration interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip` - IP address.
* `keychain` - Message-digest key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `md5_keychain` - Authentication MD5 key-chain name. This attribute must reference one of the following datasources: `router.key-chain.name` .
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU. Valid values: `enable` `disable` .
* `name` - Interface entry name.
* `network_type` - Network type. Valid values: `broadcast` `non-broadcast` `point-to-point` `point-to-multipoint` `point-to-multipoint-non-broadcast` .
* `prefix_length` - Prefix length.
* `priority` - Priority.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable status. Valid values: `disable` `enable` .
* `transmit_delay` - Transmit delay.
* `md5_keys` - MD5 key. The structure of `md5_keys` block is documented below.

The `md5_keys` block contains:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.
* `passive_interface` - Passive interface configuration. The structure of `passive_interface` block is documented below.

The `passive_interface` block contains:

* `name` - Passive interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `redistribute` - Redistribute configuration. The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `metric` - Redistribute metric setting.
* `metric_type` - Metric type. Valid values: `1` `2` .
* `name` - Redistribute name.
* `routemap` - Route map name. This attribute must reference one of the following datasources: `router.route-map.name` .
* `status` - status Valid values: `enable` `disable` .
* `tag` - Tag value.
* `summary_address` - IP address summary configuration. The structure of `summary_address` block is documented below.

The `summary_address` block contains:

* `advertise` - Enable/disable advertise status. Valid values: `disable` `enable` .
* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_ospf can be imported using:
```sh
terraform import fortios_router_ospf.labelname {{mkey}}
```
