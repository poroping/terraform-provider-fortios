---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress"
description: |-
  Get information on a fortios Configure multicast addresses.
---

# Data Source: fortios_firewall_multicastaddress
Use this data source to get information on a fortios Configure multicast addresses.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Multicast address name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Comment.
* `end_ip` - Final IPv4 address (inclusive) in the range for the address.
* `name` - Multicast address name.
* `start_ip` - First IPv4 address (inclusive) in the range for the address.
* `subnet` - Broadcast address and subnet.
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address.
* `visibility` - Enable/disable visibility of the multicast address on the GUI.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
