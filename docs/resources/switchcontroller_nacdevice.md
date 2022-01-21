---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacdevice"
description: |-
  Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.
---

## fortios_switchcontroller_nacdevice
Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `description` - Description for the learned NAC device.
* `id` - Device ID.
* `last_known_port` - Managed FortiSwitch port where NAC device is last learned.
* `last_known_switch` - Managed FortiSwitch where NAC device is last learned. This attribute must reference one of the following datasources: `switch-controller.managed-switch.switch-id` .
* `last_seen` - Device last seen.
* `mac` - MAC address of the learned NAC device.
* `mac_policy` - MAC policy to be applied on this learned NAC device. This attribute must reference one of the following datasources: `switch-controller.mac-policy.name` .
* `matched_nac_policy` - Matched NAC policy for the learned NAC device. This attribute must reference one of the following datasources: `user.nac-policy.name` .
* `port_policy` - Port policy to be applied on this learned NAC device. This attribute must reference one of the following datasources: `switch-controller.port-policy.name` .
* `status` - Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_nacdevice can be imported using:
```sh
terraform import fortios_switchcontroller_nacdevice.labelname {{mkey}}
```
