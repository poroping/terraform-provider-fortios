---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtpgroup"
description: |-
  Get information on a fortios Configure WTP groups.
---

# Data Source: fortios_wirelesscontroller_wtpgroup
Use this data source to get information on a fortios Configure WTP groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WTP group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - WTP group name.
* `platform_type` - FortiAP models to define the WTP group platform type.
* `wtps` - WTP list.The structure of `wtps` block is documented below.

The `wtps` block contains:

* `wtp_id` - WTP ID.
