---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpvenueurl"
description: |-
  Get information on a fortios Configure venue URL.
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqpvenueurl
Use this data source to get information on a fortios Configure venue URL.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of venue url.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Name of venue url.
* `value_list` - URL list.The structure of `value_list` block is documented below.

The `value_list` block contains:

* `index` - URL index.
* `number` - Venue number.
* `value` - Venue URL value.
