---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqp3gppcellular"
description: |-
  Get information on a fortios Configure 3GPP public land mobile network (PLMN).
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqp3gppcellular
Use this data source to get information on a fortios Configure 3GPP public land mobile network (PLMN).


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) 3GPP PLMN name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - 3GPP PLMN name.
* `mcc_mnc_list` - Mobile Country Code and Mobile Network Code configuration.The structure of `mcc_mnc_list` block is documented below.

The `mcc_mnc_list` block contains:

* `id` - ID.
* `mcc` - Mobile country code.
* `mnc` - Mobile network code.
