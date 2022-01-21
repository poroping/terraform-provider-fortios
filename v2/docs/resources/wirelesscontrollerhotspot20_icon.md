---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_icon"
description: |-
  Configure OSU provider icon.
---

## fortios_wirelesscontrollerhotspot20_icon
Configure OSU provider icon.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Icon list ID.
* `icon_list` - Icon list. The structure of `icon_list` block is documented below.

The `icon_list` block contains:

* `file` - Icon file.
* `height` - Icon height.
* `lang` - Language code.
* `name` - Icon name.
* `type` - Icon type. Valid values: `bmp` `gif` `jpeg` `png` `tiff` .
* `width` - Icon width.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_icon can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_icon.labelname {{mkey}}
```
