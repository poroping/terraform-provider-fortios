---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_centralsnatmap"
description: |-
  Configure IPv4 and IPv6 central SNAT policies.
---

## fortios_firewall_centralsnatmap
Configure IPv4 and IPv6 central SNAT policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comments` - Comment.
* `nat` - Enable/disable source NAT. Valid values: `disable` `enable` .
* `nat_port` - Translated port or port range (1 to 65535, 0 means any port).
* `nat46` - Enable/disable NAT46. Valid values: `enable` `disable` .
* `nat64` - Enable/disable NAT64. Valid values: `enable` `disable` .
* `orig_port` - Original TCP port (1 to 65535, 0 means any port).
* `policyid` - Policy ID.
* `protocol` - Integer value for the protocol type (0 - 255).
* `status` - Enable/disable the active status of this policy. Valid values: `enable` `disable` .
* `type` - IPv4/IPv6 source NAT. Valid values: `ipv4` `ipv6` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dst_addr` - IPv4 Destination address. The structure of `dst_addr` block is documented below.

The `dst_addr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dst_addr6` - IPv6 Destination address. The structure of `dst_addr6` block is documented below.

The `dst_addr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `dstintf` - Destination interface name from available interfaces. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `nat_ippool` block is documented below.

The `nat_ippool` block contains:

* `name` - IP pool name. This attribute must reference one of the following datasources: `firewall.ippool.name` .
* `nat_ippool6` - IPv6 pools to be used for source NAT. The structure of `nat_ippool6` block is documented below.

The `nat_ippool6` block contains:

* `name` - IPv6 pool name. This attribute must reference one of the following datasources: `firewall.ippool6.name` .
* `orig_addr` - IPv4 Original address. The structure of `orig_addr` block is documented below.

The `orig_addr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `orig_addr6` - IPv6 Original address. The structure of `orig_addr6` block is documented below.

The `orig_addr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `srcintf` - Source interface name from available interfaces. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_centralsnatmap can be imported using:
```sh
terraform import fortios_firewall_centralsnatmap.labelname {{mkey}}
```
