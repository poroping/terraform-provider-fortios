---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacsettings"
description: |-
  Get information on a fortios Configure integrated NAC settings for FortiSwitch.
---

# Data Source: fortios_switchcontroller_nacsettings
Use this data source to get information on a fortios Configure integrated NAC settings for FortiSwitch.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) NAC settings name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auto_auth` - Enable/disable NAC device auto authorization when discovered and nac-policy matched.
* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.
* `inactive_timer` - Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC devices on switch ports on link down event.
* `mode` - Set NAC mode to be used on the FortiSwitch ports.
* `name` - NAC settings name.
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
