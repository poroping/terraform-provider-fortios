---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_vlanpolicy"
description: |-
  Get information on a fortios Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.
---

# Data Source: fortios_switchcontroller_vlanpolicy
Use this data source to get information on a fortios Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) VLAN policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allowed_vlans_all` - Enable/disable all defined VLANs when using this VLAN policy.
* `description` - Description for the VLAN policy.
* `discard_mode` - Discard mode to be applied when using this VLAN policy.
* `fortilink` - FortiLink interface for which this VLAN policy belongs to.
* `name` - VLAN policy name.
* `vlan` - Native VLAN to be applied when using this VLAN policy.
* `allowed_vlans` - Allowed VLANs to be applied when using this VLAN policy.The structure of `allowed_vlans` block is documented below.

The `allowed_vlans` block contains:

* `vlan_name` - VLAN name.
* `untagged_vlans` - Untagged VLANs to be applied when using this VLAN policy.The structure of `untagged_vlans` block is documented below.

The `untagged_vlans` block contains:

* `vlan_name` - VLAN name.
