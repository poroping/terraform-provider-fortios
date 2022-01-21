---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpvenuename"
description: |-
  Get information on a fortios Configure venue name duple.
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqpvenuename
Use this data source to get information on a fortios Configure venue name duple.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of venue name duple.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Name of venue name duple.
* `value_list` - Name list.The structure of `value_list` block is documented below.

The `value_list` block contains:

* `index` - Value index.
* `lang` - Language code.
* `value` - Venue name value.
