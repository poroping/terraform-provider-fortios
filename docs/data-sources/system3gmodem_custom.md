---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system3gmodem_custom"
description: |-
  Get information on a fortios 3G MODEM custom.
---

# Data Source: fortios_system3gmodem_custom
Use this data source to get information on a fortios 3G MODEM custom.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `class_id` - USB interface class in hexadecimal format (00-ff).
* `id` - ID.
* `init_string` - Init string in hexadecimal format (even length).
* `model` - MODEM model name.
* `modeswitch_string` - USB modeswitch arguments. For example: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'.
* `product_id` - USB product ID in hexadecimal format (0000-ffff).
* `vendor` - MODEM vendor name.
* `vendor_id` - USB vendor ID in hexadecimal format (0000-ffff).
