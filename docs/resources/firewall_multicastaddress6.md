---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastaddress6"
description: |-
  Configure IPv6 multicast address.
---

## fortios_firewall_multicastaddress6
Configure IPv6 multicast address.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `name` - IPv6 multicast address name.
* `visibility` - Enable/disable visibility of the IPv6 multicast address on the GUI. Valid values: `enable` `disable` .
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

fortios_firewall_multicastaddress6 can be imported using:
```sh
terraform import fortios_firewall_multicastaddress6.labelname {{mkey}}
```
