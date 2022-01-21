---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system3gmodem_custom"
description: |-
  3G MODEM custom.
---

## fortios_system3gmodem_custom
3G MODEM custom.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `class_id` - USB interface class in hexadecimal format (00-ff).
* `id` - ID.
* `init_string` - Init string in hexadecimal format (even length).
* `model` - MODEM model name.
* `modeswitch_string` - USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'
* `product_id` - USB product ID in hexadecimal format (0000-ffff).
* `vendor` - MODEM vendor name.
* `vendor_id` - USB vendor ID in hexadecimal format (0000-ffff).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system3gmodem_custom can be imported using:
```sh
terraform import fortios_system3gmodem_custom.labelname {{mkey}}
```
