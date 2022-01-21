---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge"
description: |-
  Configure advice of charge.
---

## fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge
Configure advice of charge.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Plan name.
* `aoc_list` - AOC list. The structure of `aoc_list` block is documented below.

The `aoc_list` block contains:

* `nai_realm` - NAI realm list name.
* `nai_realm_encoding` - NAI realm encoding.
* `name` - Advice of charge ID.
* `type` - Usage charge type. Valid values: `time-based` `volume-based` `time-and-volume-based` `unlimited` .
* `plan_info` - Plan info. The structure of `plan_info` block is documented below.

The `plan_info` block contains:

* `currency` - Currency code.
* `info_file` - Info file.
* `lang` - Languague code.
* `name` - Plan name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge.labelname {{mkey}}
```
