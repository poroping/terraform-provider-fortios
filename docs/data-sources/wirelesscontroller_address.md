---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_address"
description: |-
  Get information on a fortios Configure the client with its MAC address.
---

# Data Source: fortios_wirelesscontroller_address
Use this data source to get information on a fortios Configure the client with its MAC address.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - ID.
* `mac` - MAC address.
* `policy` - Allow or block the client with this MAC address.
