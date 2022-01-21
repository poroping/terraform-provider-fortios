---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationstitch"
description: |-
  Automation stitches.
---

## fortios_system_automationstitch
Automation stitches.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `description` - Description.
* `name` - Name.
* `status` - Enable/disable this stitch. Valid values: `enable` `disable` .
* `trigger` - Trigger name. This attribute must reference one of the following datasources: `system.automation-trigger.name` .
* `action` - Action names. The structure of `action` block is documented below.

The `action` block contains:

* `name` - Action name. This attribute must reference one of the following datasources: `system.automation-action.name` .
* `actions` - Configure stitch actions. The structure of `actions` block is documented below.

The `actions` block contains:

* `action` - Action name. This attribute must reference one of the following datasources: `system.automation-action.name` .
* `delay` - Delay before execution (in seconds).
* `id` - Entry ID.
* `required` - Required in action chain. Valid values: `enable` `disable` .
* `destination` - Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.

The `destination` block contains:

* `name` - Destination name. This attribute must reference one of the following datasources: `system.automation-destination.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_automationstitch can be imported using:
```sh
terraform import fortios_system_automationstitch.labelname {{mkey}}
```
