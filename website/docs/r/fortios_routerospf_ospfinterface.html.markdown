---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_ospf_interface"
description: |-
  OSPF interface configuration.
---

## fortios_router_ospf_ospf_interface
OSPF interface configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `authentication` - Authentication type. Valid values: `none` `text` `message-digest` `md5` .
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

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

router_ospf_interface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_ospf_interface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
