---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stpinstance"
description: |-
  Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.
---

## fortios_switchcontroller_stpinstance
Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `id` - Instance ID.
* `vlan_range` - Configure VLAN range for STP instance. The structure of `vlan_range` block is documented below.

The `vlan_range` block contains:

* `vlan_name` - VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_stpinstance can be imported using:
```sh
terraform import fortios_switchcontroller_stpinstance.labelname {{mkey}}
```
