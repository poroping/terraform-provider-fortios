---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_icon"
description: |-
  Get information on a fortios Configure OSU provider icon.
---

# Data Source: fortios_wirelesscontrollerhotspot20_icon
Use this data source to get information on a fortios Configure OSU provider icon.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Icon list ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Icon list ID.
* `icon_list` - Icon list.The structure of `icon_list` block is documented below.

The `icon_list` block contains:

* `file` - Icon file.
* `height` - Icon height.
* `lang` - Language code.
* `name` - Icon name.
* `type` - Icon type.
* `width` - Icon width.
