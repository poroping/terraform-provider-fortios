---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_vapgroup"
description: |-
  Get information on a fortios Configure virtual Access Point (VAP) groups.
---

# Data Source: fortios_wirelesscontroller_vapgroup
Use this data source to get information on a fortios Configure virtual Access Point (VAP) groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Group Name
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `name` - Group Name
* `vaps` - List of SSIDs to be included in the VAP group.The structure of `vaps` block is documented below.

The `vaps` block contains:

* `name` - vap name
