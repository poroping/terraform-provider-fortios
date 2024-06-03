---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustom"
description: |-
  Configure custom Internet Services.
---

## fortios_firewall_internetservicecustom
Configure custom Internet Services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `name` - Internet Service name.
* `reputation` - Reputation level of the custom Internet Service. This attribute must reference one of the following datasources: `firewall.internet-service-reputation.id` .
* `entry` - Entries added to the Internet Service database and custom database. The structure of `entry` block is documented below.

The `entry` block contains:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4` `ipv6` .
* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `dst` - Destination address or address group name. The structure of `dst` block is documented below.

The `dst` block contains:

* `name` - Select the destination address or address group object from available options. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dst6` - Destination address6 or address6 group name. The structure of `dst6` block is documented below.

The `dst6` block contains:

* `name` - Select the destination address6 or address group object from available options. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_internetservicecustom can be imported using:
```sh
terraform import fortios_firewall_internetservicecustom.labelname {{mkey}}
```
