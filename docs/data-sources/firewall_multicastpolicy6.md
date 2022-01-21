---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastpolicy6"
description: |-
  Get information on a fortios Configure IPv6 multicast NAT policies.
---

# Data Source: fortios_firewall_multicastpolicy6
Use this data source to get information on a fortios Configure IPv6 multicast NAT policies.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Policy ID (0 - 4294967294).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Accept or deny traffic matching the policy.
* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration.
* `comments` - Comment.
* `dstintf` - IPv6 destination interface name.
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
* `id` - Policy ID (0 - 4294967294).
* `logtraffic` - Enable/disable logging traffic accepted by this policy.
* `name` - Policy name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `srcintf` - IPv6 source interface name.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `status` - Enable/disable this policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dstaddr` - IPv6 destination address name.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `srcaddr` - IPv6 source address name.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
