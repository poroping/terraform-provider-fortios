---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype"
description: |-
  Get information on a fortios Configure network authentication type.
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype
Use this data source to get information on a fortios Configure network authentication type.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Authentication type name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_type` - Network authentication type.
* `name` - Authentication type name.
* `url` - Redirect URL.
