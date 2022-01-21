---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastpolicy"
description: |-
  Configure multicast NAT policies.
---

## fortios_firewall_multicastpolicy
Configure multicast NAT policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Accept or deny traffic matching the policy. Valid values: `accept` `deny` .
* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable` `disable` .
* `comments` - Comment.
* `dnat` - IPv4 DNAT address used for multicast destination addresses.
* `dstintf` - Destination interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `end_port` -  Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `id` - Policy ID ((0 - 4294967294).
* `logtraffic` - Enable/disable logging traffic accepted by this policy. Valid values: `enable` `disable` .
* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `snat` - Enable/disable substitution of the outgoing interface IP address for the original source IP address (called source NAT or SNAT). Valid values: `enable` `disable` .
* `snat_ip` - IPv4 address to be used as the source address for NATed traffic.
* `srcintf` - Source interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Destination address objects. This attribute must reference one of the following datasources: `firewall.multicast-address.name` .
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Source address objects. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_multicastpolicy can be imported using:
```sh
terraform import fortios_firewall_multicastpolicy.labelname {{mkey}}
```
