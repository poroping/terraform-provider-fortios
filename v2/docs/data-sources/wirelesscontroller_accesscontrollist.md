---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_accesscontrollist"
description: |-
  Get information on a fortios Configure WiFi bridge access control list.
---

# Data Source: fortios_wirelesscontroller_accesscontrollist
Use this data source to get information on a fortios Configure WiFi bridge access control list.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) AP access control list name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Description.
* `name` - AP access control list name.
* `layer3_ipv4_rules` - AP ACL layer3 ipv4 rule list.The structure of `layer3_ipv4_rules` block is documented below.

The `layer3_ipv4_rules` block contains:

* `action` - Policy action (allow | deny).
* `comment` - Description.
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).
* `layer3_ipv6_rules` - AP ACL layer3 ipv6 rule list.The structure of `layer3_ipv6_rules` block is documented below.

The `layer3_ipv6_rules` block contains:

* `action` - Policy action (allow | deny).
* `comment` - Description.
* `dstaddr` - Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `rule_id` - Rule ID (1 - 65535).
* `srcaddr` - Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `srcport` - Source port (0 - 65535, default = 0, meaning any).
