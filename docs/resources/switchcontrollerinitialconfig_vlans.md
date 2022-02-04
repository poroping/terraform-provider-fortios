---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerinitialconfig_vlans"
description: |-
  Configure initial template for auto-generated VLAN interfaces.
---

## fortios_switchcontrollerinitialconfig_vlans
Configure initial template for auto-generated VLAN interfaces.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `default_vlan` - Default VLAN (native) assigned to all switch ports upon discovery. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .
* `nac` - VLAN for NAC onboarding devices. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .
* `nac_segment` - VLAN for NAC segment primary interface. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .
* `quarantine` - VLAN for quarantined traffic. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .
* `rspan` - VLAN for RSPAN/ERSPAN mirrored traffic. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .
* `video` - VLAN dedicated for video devices. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .
* `voice` - VLAN dedicated for voice devices. This attribute must reference one of the following datasources: `switch-controller.initial-config.template.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontrollerinitialconfig_vlans can be imported using:
```sh
terraform import fortios_switchcontrollerinitialconfig_vlans.labelname {{mkey}}
```
