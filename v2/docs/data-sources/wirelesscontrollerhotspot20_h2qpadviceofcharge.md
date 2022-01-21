---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge"
description: |-
  Get information on a fortios Configure advice of charge.
---

# Data Source: fortios_wirelesscontrollerhotspot20_h2qpadviceofcharge
Use this data source to get information on a fortios Configure advice of charge.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Plan name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Plan name.
* `aoc_list` - AOC list.The structure of `aoc_list` block is documented below.

The `aoc_list` block contains:

* `nai_realm` - NAI realm list name.
* `nai_realm_encoding` - NAI realm encoding.
* `name` - Advice of charge ID.
* `type` - Usage charge type.
* `plan_info` - Plan info.The structure of `plan_info` block is documented below.

The `plan_info` block contains:

* `currency` - Currency code.
* `info_file` - Info file.
* `lang` - Languague code.
* `name` - Plan name.
