---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_macpolicy"
description: |-
  Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.
---

## fortios_switchcontroller_macpolicy
Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this mac-policy is applied. Valid values: `disable` `enable` .
* `count` - Enable/disable packet count on the NAC device. Valid values: `disable` `enable` .
* `description` - Description for the MAC policy.
* `drop` - Enable/disable dropping of NAC device traffic. Valid values: `disable` `enable` .
* `fortilink` - FortiLink interface for which this MAC policy belongs to. This attribute must reference one of the following datasources: `system.interface.name` .
* `name` - MAC policy name.
* `traffic_policy` - Traffic policy to be applied when using this MAC policy. This attribute must reference one of the following datasources: `switch-controller.traffic-policy.name` .
* `vlan` - Ingress traffic VLAN assignment for the MAC address matching this MAC policy. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_macpolicy can be imported using:
```sh
terraform import fortios_switchcontroller_macpolicy.labelname {{mkey}}
```
