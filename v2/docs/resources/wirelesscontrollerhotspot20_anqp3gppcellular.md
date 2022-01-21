---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqp3gppcellular"
description: |-
  Configure 3GPP public land mobile network (PLMN).
---

## fortios_wirelesscontrollerhotspot20_anqp3gppcellular
Configure 3GPP public land mobile network (PLMN).

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - 3GPP PLMN name.
* `mcc_mnc_list` - Mobile Country Code and Mobile Network Code configuration. The structure of `mcc_mnc_list` block is documented below.

The `mcc_mnc_list` block contains:

* `id` - ID.
* `mcc` - Mobile country code.
* `mnc` - Mobile network code.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_anqp3gppcellular can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_anqp3gppcellular.labelname {{mkey}}
```
