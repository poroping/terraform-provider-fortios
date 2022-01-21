---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_portpolicy"
description: |-
  Configure port policy to be applied on the managed FortiSwitch ports through NAC device.
---

## fortios_switchcontroller_portpolicy
Configure port policy to be applied on the managed FortiSwitch ports through NAC device.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `802_1x` - 802.1x security policy to be applied when using this port-policy. This attribute must reference one of the following datasources: `switch-controller.security-policy.802-1X.name` `switch-controller.security-policy.captive-portal.name` .
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this port policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable` `enable` .
* `description` - Description for the port policy.
* `fortilink` - FortiLink interface for which this port policy belongs to. This attribute must reference one of the following datasources: `system.interface.name` .
* `lldp_profile` - LLDP profile to be applied when using this port-policy. This attribute must reference one of the following datasources: `switch-controller.lldp-profile.name` .
* `name` - Port policy name.
* `qos_policy` - QoS policy to be applied when using this port-policy. This attribute must reference one of the following datasources: `switch-controller.qos.qos-policy.name` .
* `vlan_policy` - VLAN policy to be applied when using this port-policy. This attribute must reference one of the following datasources: `switch-controller.vlan-policy.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_portpolicy can be imported using:
```sh
terraform import fortios_switchcontroller_portpolicy.labelname {{mkey}}
```
