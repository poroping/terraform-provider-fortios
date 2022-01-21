---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_vlanpolicy"
description: |-
  Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.
---

## fortios_switchcontroller_vlanpolicy
Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allowed_vlans_all` - Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable` `disable` .
* `description` - Description for the VLAN policy.
* `discard_mode` - Discard mode to be applied when using this VLAN policy. Valid values: `none` `all-untagged` `all-tagged` .
* `fortilink` - FortiLink interface for which this VLAN policy belongs to. This attribute must reference one of the following datasources: `system.interface.name` .
* `name` - VLAN policy name.
* `vlan` - Native VLAN to be applied when using this VLAN policy. This attribute must reference one of the following datasources: `system.interface.name` .
* `allowed_vlans` - Allowed VLANs to be applied when using this VLAN policy. The structure of `allowed_vlans` block is documented below.

The `allowed_vlans` block contains:

* `vlan_name` - VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .
* `untagged_vlans` - Untagged VLANs to be applied when using this VLAN policy. The structure of `untagged_vlans` block is documented below.

The `untagged_vlans` block contains:

* `vlan_name` - VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_vlanpolicy can be imported using:
```sh
terraform import fortios_switchcontroller_vlanpolicy.labelname {{mkey}}
```
