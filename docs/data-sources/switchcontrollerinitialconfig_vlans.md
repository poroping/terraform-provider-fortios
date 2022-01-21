---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerinitialconfig_vlans"
description: |-
  Get information on a fortios Configure initial template for auto-generated VLAN interfaces.
---

# Data Source: fortios_switchcontrollerinitialconfig_vlans
Use this data source to get information on a fortios Configure initial template for auto-generated VLAN interfaces.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `default_vlan` - Default VLAN (native) assigned to all switch ports upon discovery.
* `nac` - VLAN for NAC onboarding devices.
* `nac_segment` - VLAN for NAC segemnt primary interface.
* `quarantine` - VLAN for quarantined traffic.
* `rspan` - VLAN for RSPAN/ERSPAN mirrored traffic.
* `video` - VLAN dedicated for video devices.
* `voice` - VLAN dedicated for voice devices.
