---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_accesscontrollist"
description: |-
  Configure WiFi bridge access control list.
---

## fortios_wirelesscontroller_accesscontrollist
Configure WiFi bridge access control list.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Description.
* `name` - AP access control list name.
* `layer3_ipv4_rules` - AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.

The `layer3_ipv4_rules` block contains:

* `action` - Policy action (allow | deny). Valid values: `allow` `deny` .
* `comment` - Description.
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).
* `layer3_ipv6_rules` - AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.

The `layer3_ipv6_rules` block contains:

* `action` - Policy action (allow | deny). Valid values: `allow` `deny` .
* `comment` - Description.
* `dstaddr` - Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `srcport` - Source port (0 - 65535, default = 0, meaning any).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_accesscontrollist can be imported using:
```sh
terraform import fortios_wirelesscontroller_accesscontrollist.labelname {{mkey}}
```
