---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_ospfinterface"
description: |-
  Get information on a fortios OSPF interface configuration.
---

# Data Source: fortios_router_ospf_ospfinterface
Use this data source to get information on a fortios OSPF interface configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Interface entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

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
* `md5_keychain` - Authentication MD5 key-chain name.
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
