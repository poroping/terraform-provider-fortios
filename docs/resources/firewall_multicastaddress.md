---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress"
description: |-
  Configure multicast addresses.
---

## fortios_firewall_multicastaddress
Configure multicast addresses.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available. This attribute must reference one of the following datasources: `system.interface.name` .
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Comment.
* `end_ip` - Final IPv4 address (inclusive) in the range for the address.
* `name` - Multicast address name.
* `start_ip` - First IPv4 address (inclusive) in the range for the address.
* `subnet` - Broadcast address and subnet.
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange` `broadcastmask` .
* `visibility` - Enable/disable visibility of the multicast address on the GUI. Valid values: `enable` `disable` .
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category. This attribute must reference one of the following datasources: `system.object-tagging.category` .
* `name` - Tagging entry name.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name. This attribute must reference one of the following datasources: `system.object-tagging.tags.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_multicastaddress can be imported using:
```sh
terraform import fortios_firewall_multicastaddress.labelname {{mkey}}
```
