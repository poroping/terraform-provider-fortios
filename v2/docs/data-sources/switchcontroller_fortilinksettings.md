---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_fortilinksettings"
description: |-
  Get information on a fortios Configure integrated FortiLink settings for FortiSwitch.
---

# Data Source: fortios_switchcontroller_fortilinksettings
Use this data source to get information on a fortios Configure integrated FortiLink settings for FortiSwitch.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) FortiLink settings name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fortilink` - FortiLink interface to which this fortilink-setting belongs.
* `inactive_timer` - Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC and dynamic devices on switch ports on link down event.
* `name` - FortiLink settings name.
* `nac_ports` - NAC specific configuration.The structure of `nac_ports` block is documented below.

The `nac_ports` block contains:

* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.
* `lan_segment` - Enable/disable LAN segment feature on the FortiLink interface.
* `member_change` - Member change flag.
* `nac_lan_interface` - Configure NAC LAN interface.
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `parent_key` - Parent key name.
* `nac_segment_vlans` - Configure NAC segment VLANs.The structure of `nac_segment_vlans` block is documented below.

The `nac_segment_vlans` block contains:

* `vlan_name` - VLAN interface name.
