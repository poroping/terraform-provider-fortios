---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpvenueurl"
description: |-
  Configure venue URL.
---

## fortios_wirelesscontrollerhotspot20_anqpvenueurl
Configure venue URL.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Name of venue url.
* `value_list` - URL list. The structure of `value_list` block is documented below.

The `value_list` block contains:

* `index` - URL index.
* `number` - Venue number.
* `value` - Venue URL value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_anqpvenueurl can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_anqpvenueurl.labelname {{mkey}}
```
