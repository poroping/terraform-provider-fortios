---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicecustomgroup"
description: |-
  Get information on a fortios Configure custom Internet Service group.
---

# Data Source: fortios_firewall_internetservicecustomgroup
Use this data source to get information on a fortios Configure custom Internet Service group.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Custom Internet Service group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - Custom Internet Service group name.
* `member` - Custom Internet Service group members.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Group member name.
