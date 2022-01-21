---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_accesslist6"
description: |-
  Configure IPv6 access lists.
---

## fortios_router_accesslist6
Configure IPv6 access lists.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comments` - Comment.
* `name` - Name.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Permit or deny this IP address and netmask prefix. Valid values: `permit` `deny` .
* `exact_match` - Enable/disable exact prefix match. Valid values: `enable` `disable` .
* `flags` - Flags.
* `id` - Rule ID.
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_accesslist6 can be imported using:
```sh
terraform import fortios_router_accesslist6.labelname {{mkey}}
```
