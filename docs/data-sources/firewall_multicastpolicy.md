---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastpolicy"
description: |-
  Get information on a fortios Configure multicast NAT policies.
---

# Data Source: fortios_firewall_multicastpolicy
Use this data source to get information on a fortios Configure multicast NAT policies.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Policy ID ((0 - 4294967294).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Accept or deny traffic matching the policy.
* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration.
* `comments` - Comment.
* `dnat` - IPv4 DNAT address used for multicast destination addresses.
* `dstintf` - Destination interface name.
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `id` - Policy ID ((0 - 4294967294).
* `logtraffic` - Enable/disable logging traffic accepted by this policy.
* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `snat` - Enable/disable substitution of the outgoing interface IP address for the original source IP address (called source NAT or SNAT).
* `snat_ip` - IPv4 address to be used as the source address for NATed traffic.
* `srcintf` - Source interface name.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - Destination address objects.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Destination address objects.
* `srcaddr` - Source address objects.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Source address objects.
