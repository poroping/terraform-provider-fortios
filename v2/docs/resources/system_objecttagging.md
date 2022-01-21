---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_objecttagging"
description: |-
  Configure object tagging.
---

## fortios_system_objecttagging
Configure object tagging.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `category` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `address` - Address. Valid values: `disable` `mandatory` `optional` .
* `category` - Tag Category.
* `color` - Color of icon on the GUI.
* `device` - Device. Valid values: `disable` `mandatory` `optional` .
* `interface` - Interface. Valid values: `disable` `mandatory` `optional` .
* `multiple` - Allow multiple tag selection. Valid values: `enable` `disable` .
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_objecttagging can be imported using:
```sh
terraform import fortios_system_objecttagging.labelname {{mkey}}
```
