---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtpgroup"
description: |-
  Configure WTP groups.
---

## fortios_wirelesscontroller_wtpgroup
Configure WTP groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ble_major_id` - Override BLE Major ID.
* `name` - WTP group name.
* `platform_type` - FortiAP models to define the WTP group platform type. Valid values: `AP-11N` `220B` `210B` `222B` `112B` `320B` `11C` `14C` `223B` `28C` `320C` `221C` `25D` `222C` `224D` `214B` `21D` `24D` `112D` `223C` `321C` `C220C` `C225C` `C23JD` `C24JE` `S321C` `S322C` `S323C` `S311C` `S313C` `S321CR` `S322CR` `S323CR` `S421E` `S422E` `S423E` `421E` `423E` `221E` `222E` `223E` `224E` `231E` `S221E` `S223E` `321E` `431F` `431FL` `432F` `432FR` `433F` `433FL` `231F` `231FL` `234F` `23JF` `831F` `231G` `233G` `234G` `431G` `432G` `433G` `U421E` `U422EV` `U423E` `U221EV` `U223EV` `U24JEV` `U321EV` `U323EV` `U431F` `U433F` `U231F` `U234F` `U432F` `U231G` `U441G` .
* `wtps` - WTP list. The structure of `wtps` block is documented below.

The `wtps` block contains:

* `wtp_id` - WTP ID. This attribute must reference one of the following datasources: `wireless-controller.wtp.wtp-id` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_wtpgroup can be imported using:
```sh
terraform import fortios_wirelesscontroller_wtpgroup.labelname {{mkey}}
```
