---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacdevice"
description: |-
  Get information on a fortios Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.
---

# Data Source: fortios_switchcontroller_nacdevice
Use this data source to get information on a fortios Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Device ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `description` - Description for the learned NAC device.
* `id` - Device ID.
* `last_known_port` - Managed FortiSwitch port where NAC device is last learned.
* `last_known_switch` - Managed FortiSwitch where NAC device is last learned.
* `last_seen` - Device last seen.
* `mac` - MAC address of the learned NAC device.
* `mac_policy` - MAC policy to be applied on this learned NAC device.
* `matched_nac_policy` - Matched NAC policy for the learned NAC device.
* `port_policy` - Port policy to be applied on this learned NAC device.
* `status` - Status of the learned NAC device. Set enable to authorize the NAC device.
