---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6"
description: |-
  Get information on a fortios Configure IPv6 OSPF.
---

# Data Source: fortios_router_ospf6
Use this data source to get information on a fortios Configure IPv6 OSPF.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `abr_type` - Area border router type.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type.
* `default_information_originate` - Enable/disable generation of default route.
* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `log_neighbour_changes` - Log OSPFv3 neighbor changes.
* `restart_mode` - OSPFv3 restart mode (graceful or none).
* `restart_on_topology_change` - Enable/disable continuing graceful restart upon topology change.
* `restart_period` - Graceful restart period in seconds.
* `router_id` - A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `area` - OSPF6 area configuration.The structure of `area` block is documented below.

The `area` block contains:

* `authentication` - Authentication mode.
* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `key_rollover_interval` - Key roll-over interval.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area.
* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes.
* `nssa_redistribution` - Enable/disable redistribute into NSSA area.
* `nssa_translator_role` - NSSA translator role type.
* `stub_type` - Stub summary setting.
* `type` - Area type setting.
* `ipsec_keys` - IPsec authentication and encryption keys.The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `range` - OSPF6 area range configuration.The structure of `range` block is documented below.

The `range` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.
* `virtual_link` - OSPF6 virtual link configuration.The structure of `virtual_link` block is documented below.

The `virtual_link` block contains:

* `authentication` - Authentication mode.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `key_rollover_interval` - Key roll-over interval.
* `name` - Virtual link entry name.
* `peer` - A.B.C.D, peer router ID.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `ipsec_keys` - IPsec authentication and encryption keys.The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `ospf6_interface` - OSPF6 interface configuration.The structure of `ospf6_interface` block is documented below.

The `ospf6_interface` block contains:

* `area_id` - A.B.C.D, in IPv4 address format.
* `authentication` - Authentication mode.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `interface` - Configuration interface name.
* `ipsec_auth_alg` - Authentication algorithm.
* `ipsec_enc_alg` - Encryption algorithm.
* `key_rollover_interval` - Key roll-over interval.
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets.
* `name` - Interface entry name.
* `network_type` - Network type.
* `priority` - Priority.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable OSPF6 routing on this interface.
* `transmit_delay` - Transmit delay.
* `ipsec_keys` - IPsec authentication and encryption keys.The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media.The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.
* `passive_interface` - Passive interface configuration.The structure of `passive_interface` block is documented below.

The `passive_interface` block contains:

* `name` - Passive interface name.
* `redistribute` - Redistribute configuration.The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `metric` - Redistribute metric setting.
* `metric_type` - Metric type.
* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status.
* `summary_address` - IPv6 address summary configuration.The structure of `summary_address` block is documented below.

The `summary_address` block contains:

* `advertise` - Enable/disable advertise status.
* `id` - Summary address entry ID.
* `prefix6` - IPv6 prefix.
* `tag` - Tag value.
