---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_aspathlist"
description: |-
  Configure Autonomous System (AS) path lists.
---

## fortios_router_aspathlist
Configure Autonomous System (AS) path lists.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - AS path list name.
* `rule` - AS path list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute. Valid values: `deny` `permit` .
* `id` - ID.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_aspathlist can be imported using:
```sh
terraform import fortios_router_aspathlist.labelname {{mkey}}
```
