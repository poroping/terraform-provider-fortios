---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_fortilinksettings"
description: |-
  Configure integrated FortiLink settings for FortiSwitch.
---

## fortios_switchcontroller_fortilinksettings
Configure integrated FortiLink settings for FortiSwitch.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `fortilink` - FortiLink interface to which this fortilink-setting belongs. This attribute must reference one of the following datasources: `system.interface.name` .
* `inactive_timer` - Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable` `enable` .
* `name` - FortiLink settings name.
* `nac_ports` - NAC specific configuration. The structure of `nac_ports` block is documented below.

The `nac_ports` block contains:

* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable` `enable` .
* `lan_segment` - Enable/disable LAN segment feature on the FortiLink interface. Valid values: `enabled` `disabled` .
* `member_change` - Member change flag.
* `nac_lan_interface` - Configure NAC LAN interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered. This attribute must reference one of the following datasources: `system.interface.name` .
* `parent_key` - Parent key name.
* `nac_segment_vlans` - Configure NAC segment VLANs. The structure of `nac_segment_vlans` block is documented below.

The `nac_segment_vlans` block contains:

* `vlan_name` - VLAN interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_fortilinksettings can be imported using:
```sh
terraform import fortios_switchcontroller_fortilinksettings.labelname {{mkey}}
```
