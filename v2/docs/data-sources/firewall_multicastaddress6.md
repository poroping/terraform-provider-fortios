---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress6"
description: |-
  Get information on a fortios Configure IPv6 multicast address.
---

# Data Source: fortios_firewall_multicastaddress6
Use this data source to get information on a fortios Configure IPv6 multicast address.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPv6 multicast address name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `name` - IPv6 multicast address name.
* `visibility` - Enable/disable visibility of the IPv6 multicast address on the GUI.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
