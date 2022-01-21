---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationdestination"
description: |-
  Automation destinations.
---

## fortios_system_automationdestination
Automation destinations.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `ha_group_id` - Cluster group ID set for this destination (default = 0).
* `name` - Name.
* `type` - Destination type. Valid values: `fortigate` `ha-cluster` .
* `destination` - Destinations. The structure of `destination` block is documented below.

The `destination` block contains:

* `name` - Destination.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_automationdestination can be imported using:
```sh
terraform import fortios_system_automationdestination.labelname {{mkey}}
```
