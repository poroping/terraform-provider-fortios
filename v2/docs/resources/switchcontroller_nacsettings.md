---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacsettings"
description: |-
  Configure integrated NAC settings for FortiSwitch.
---

## fortios_switchcontroller_nacsettings
Configure integrated NAC settings for FortiSwitch.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auto_auth` - Enable/disable NAC device auto authorization when discovered and nac-policy matched. Valid values: `disable` `enable` .
* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable` `enable` .
* `inactive_timer` - Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC devices on switch ports on link down event. Valid values: `disable` `enable` .
* `mode` - Set NAC mode to be used on the FortiSwitch ports. Valid values: `local` `global` .
* `name` - NAC settings name.
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontroller_nacsettings can be imported using:
```sh
terraform import fortios_switchcontroller_nacsettings.labelname {{mkey}}
```
