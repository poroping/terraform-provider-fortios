---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stpinstance"
description: |-
  Get information on a fortios Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.
---

# Data Source: fortios_switchcontroller_stpinstance
Use this data source to get information on a fortios Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Instance ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Instance ID.
* `vlan_range` - Configure VLAN range for STP instance.The structure of `vlan_range` block is documented below.

The `vlan_range` block contains:

* `vlan_name` - VLAN name.
