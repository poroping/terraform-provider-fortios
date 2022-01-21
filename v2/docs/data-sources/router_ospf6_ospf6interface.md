---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6_ospf6interface"
description: |-
  Get information on a fortios OSPF6 interface configuration.
---

# Data Source: fortios_router_ospf6_ospf6interface
Use this data source to get information on a fortios OSPF6 interface configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Interface entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

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
* `priority` - priority
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable OSPF6 routing on this interface.
* `transmit_delay` - Transmit delay.
* `ipsec_keys` - IPsec authentication and encryption keys.The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast mediaThe structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - priority
