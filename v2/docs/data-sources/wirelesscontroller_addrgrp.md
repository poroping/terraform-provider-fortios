---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_addrgrp"
description: |-
  Get information on a fortios Configure the MAC address group.
---

# Data Source: fortios_wirelesscontroller_addrgrp
Use this data source to get information on a fortios Configure the MAC address group.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `default_policy` - Allow or block the clients with MAC addresses that are not in the group.
* `id` - ID.
* `addresses` - Manually selected group of addresses.The structure of `addresses` block is documented below.

The `addresses` block contains:

* `id` - Address ID.
