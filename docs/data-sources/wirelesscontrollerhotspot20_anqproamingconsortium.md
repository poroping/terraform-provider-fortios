---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqproamingconsortium"
description: |-
  Get information on a fortios Configure roaming consortium.
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqproamingconsortium
Use this data source to get information on a fortios Configure roaming consortium.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Roaming consortium name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Roaming consortium name.
* `oi_list` - Organization identifier list.The structure of `oi_list` block is documented below.

The `oi_list` block contains:

* `comment` - Comment.
* `index` - OI index.
* `oi` - Organization identifier.
