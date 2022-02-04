---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_nacpolicy"
description: |-
  Get information on a fortios Configure NAC policy matching pattern to identify matching NAC devices.
---

# Data Source: fortios_user_nacpolicy
Use this data source to get information on a fortios Configure NAC policy matching pattern to identify matching NAC devices.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) NAC policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `category` - Category of NAC policy.
* `description` - Description for the NAC policy matching pattern.
* `ems_tag` - NAC policy matching EMS tag.
* `family` - NAC policy matching family.
* `firewall_address` - Dynamic firewall address to associate MAC which match this policy.
* `host` - NAC policy matching host.
* `hw_vendor` - NAC policy matching hardware vendor.
* `hw_version` - NAC policy matching hardware version.
* `mac` - NAC policy matching MAC address.
* `name` - NAC policy name.
* `os` - NAC policy matching operating system.
* `src` - NAC policy matching source.
* `ssid_policy` - SSID policy to be applied on the matched NAC policy.
* `status` - Enable/disable NAC policy.
* `sw_version` - NAC policy matching software version.
* `switch_auto_auth` - NAC device auto authorization when discovered and nac-policy matched.
* `switch_fortilink` - FortiLink interface for which this NAC policy belongs to.
* `switch_mac_policy` - Switch MAC policy action to be applied on the matched NAC policy.
* `switch_port_policy` - switch-port-policy to be applied on the matched NAC policy.
* `type` - NAC policy matching type.
* `user` - NAC policy matching user.
* `user_group` - NAC policy matching user group.
* `switch_group` - List of managed FortiSwitch groups on which NAC policy can be applied.The structure of `switch_group` block is documented below.

The `switch_group` block contains:

* `name` - Managed FortiSwitch group name from available options.
* `switch_scope` - List of managed FortiSwitches on which NAC policy can be applied.The structure of `switch_scope` block is documented below.

The `switch_scope` block contains:

* `switch_id` - Managed FortiSwitch name from available options.
