---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_portpolicy"
description: |-
  Get information on a fortios Configure port policy to be applied on the managed FortiSwitch ports through NAC device.
---

# Data Source: fortios_switchcontroller_portpolicy
Use this data source to get information on a fortios Configure port policy to be applied on the managed FortiSwitch ports through NAC device.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Port policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `802_1x` - 802.1x security policy to be applied when using this port-policy.
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile.
* `description` - Description for the port policy.
* `fortilink` - FortiLink interface for which this port policy belongs to.
* `lldp_profile` - LLDP profile to be applied when using this port-policy.
* `name` - Port policy name.
* `qos_policy` - QoS policy to be applied when using this port-policy.
* `vlan_policy` - VLAN policy to be applied when using this port-policy.
