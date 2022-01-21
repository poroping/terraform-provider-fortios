---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_switchgroup"
description: |-
  Configure FortiSwitch switch groups.
---

## fortios_switchcontroller_switchgroup
Configure FortiSwitch switch groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `description` - Optional switch group description.
* `fortilink` - FortiLink interface to which switch group members belong. This attribute must reference one of the following datasources: `system.interface.name` .
* `name` - Switch group name.
* `members` - FortiSwitch members belonging to this switch group. The structure of `members` block is documented below.

The `members` block contains:

* `name` - Managed device ID. This attribute must reference one of the following datasources: `switch-controller.managed-switch.switch-id` .
* `switch_id` - Managed device ID. This attribute must reference one of the following datasources: `switch-controller.managed-switch.switch-id` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_switchgroup can be imported using:
```sh
terraform import fortios_switchcontroller_switchgroup.labelname {{mkey}}
```
