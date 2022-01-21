---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_switchgroup"
description: |-
  Get information on a fortios Configure FortiSwitch switch groups.
---

# Data Source: fortios_switchcontroller_switchgroup
Use this data source to get information on a fortios Configure FortiSwitch switch groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Switch group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Optional switch group description.
* `fortilink` - FortiLink interface to which switch group members belong.
* `name` - Switch group name.
* `members` - FortiSwitch members belonging to this switch group.The structure of `members` block is documented below.

The `members` block contains:

* `name` - Managed device ID.
* `switch_id` - Managed device ID.
