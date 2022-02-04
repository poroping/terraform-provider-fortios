---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6_ospf6interface"
description: |-
  OSPF6 interface configuration.
---

## fortios_router_ospf6_ospf6interface
OSPF6 interface configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf6`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `area_id` - A.B.C.D, in IPv4 address format.
* `authentication` - Authentication mode. Valid values: `none` `ah` `esp` `area` .
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global` `enable` `disable` .
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `interface` - Configuration interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5` `sha1` `sha256` `sha384` `sha512` .
* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null` `des` `3des` `aes128` `aes192` `aes256` .
* `key_rollover_interval` - Key roll-over interval.
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets. Valid values: `enable` `disable` .
* `name` - Interface entry name.
* `network_type` - Network type. Valid values: `broadcast` `point-to-point` `non-broadcast` `point-to-multipoint` `point-to-multipoint-non-broadcast` .
* `priority` - Priority.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable OSPF6 routing on this interface. Valid values: `disable` `enable` .
* `transmit_delay` - Transmit delay.
* `ipsec_keys` - IPsec authentication and encryption keys. The structure of `ipsec_keys` block is documented below.

The `ipsec_keys` block contains:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.
* `neighbor` - OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media. The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_ospf6interface can be imported using:
```sh
terraform import fortios_router_ospf6interface.labelname {{mkey}}
```
