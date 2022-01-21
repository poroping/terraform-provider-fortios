---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpoperatorname"
description: |-
  Get information on a fortios Configure operator friendly name.
---

# Data Source: fortios_wirelesscontrollerhotspot20_h2qpoperatorname
Use this data source to get information on a fortios Configure operator friendly name.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Friendly name ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Friendly name ID.
* `value_list` - Name list.The structure of `value_list` block is documented below.

The `value_list` block contains:

* `index` - Value index.
* `lang` - Language code.
* `value` - Friendly name value.
