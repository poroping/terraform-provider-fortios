---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceaddition"
description: |-
  Get information on a fortios Configure Internet Services Addition.
---

# Data Source: fortios_firewall_internetserviceaddition
Use this data source to get information on a fortios Configure Internet Services Addition.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Internet Service ID in the Internet Service database.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `id` - Internet Service ID in the Internet Service database.
* `entry` - Entries added to the Internet Service addition database.The structure of `entry` block is documented below.

The `entry` block contains:

* `addr_mode` - Address mode (IPv4 or IPv6).
* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry.The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).
