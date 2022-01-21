---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_centralsnatmap"
description: |-
  Get information on a fortios Configure IPv4 and IPv6 central SNAT policies.
---

# Data Source: fortios_firewall_centralsnatmap
Use this data source to get information on a fortios Configure IPv4 and IPv6 central SNAT policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Comment.
* `nat` - Enable/disable source NAT.
* `nat_port` - Translated port or port range (1 to 65535, 0 means any port).
* `nat46` - Enable/disable NAT46.
* `nat64` - Enable/disable NAT64.
* `orig_port` - Original TCP port (1 to 65535, 0 means any port).
* `policyid` - Policy ID.
* `protocol` - Integer value for the protocol type (0 - 255).
* `status` - Enable/disable the active status of this policy.
* `type` - IPv4/IPv6 source NAT.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dst_addr` - IPv4 Destination address.The structure of `dst_addr` block is documented below.

The `dst_addr` block contains:

* `name` - Address name.
* `dst_addr6` - IPv6 Destination address.The structure of `dst_addr6` block is documented below.

The `dst_addr6` block contains:

* `name` - Address name.
* `dstintf` - Destination interface name from available interfaces.The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name.
* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools.The structure of `nat_ippool` block is documented below.

The `nat_ippool` block contains:

* `name` - IP pool name.
* `nat_ippool6` - IPv6 pools to be used for source NAT.The structure of `nat_ippool6` block is documented below.

The `nat_ippool6` block contains:

* `name` - IPv6 pool name.
* `orig_addr` - IPv4 Original address.The structure of `orig_addr` block is documented below.

The `orig_addr` block contains:

* `name` - Address name.
* `orig_addr6` - IPv6 Original address.The structure of `orig_addr6` block is documented below.

The `orig_addr6` block contains:

* `name` - Address name.
* `srcintf` - Source interface name from available interfaces.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
