---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vipgrp"
description: |-
  Get information on a fortios Configure IPv4 virtual IP groups.
---

# Data Source: fortios_firewall_vipgrp
Use this data source to get information on a fortios Configure IPv4 virtual IP groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) VIP group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `interface` - Interface.
* `name` - VIP group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - Member VIP objects of the group (Separate multiple objects with a space).The structure of `member` block is documented below.

The `member` block contains:

* `name` - VIP name.
