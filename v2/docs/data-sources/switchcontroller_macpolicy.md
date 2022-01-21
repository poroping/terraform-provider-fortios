---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_macpolicy"
description: |-
  Get information on a fortios Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.
---

# Data Source: fortios_switchcontroller_macpolicy
Use this data source to get information on a fortios Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) MAC policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this mac-policy is applied.
* `count` - Enable/disable packet count on the NAC device.
* `description` - Description for the MAC policy.
* `drop` - Enable/disable dropping of NAC device traffic.
* `fortilink` - FortiLink interface for which this MAC policy belongs to.
* `name` - MAC policy name.
* `traffic_policy` - Traffic policy to be applied when using this MAC policy.
* `vlan` - Ingress traffic VLAN assignment for the MAC address matching this MAC policy.
