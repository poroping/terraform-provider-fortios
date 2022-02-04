---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_nacpolicy"
description: |-
  Configure NAC policy matching pattern to identify matching NAC devices.
---

## fortios_user_nacpolicy
Configure NAC policy matching pattern to identify matching NAC devices.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `category` - Category of NAC policy. Valid values: `device` `firewall-user` `ems-tag` .
* `description` - Description for the NAC policy matching pattern.
* `ems_tag` - NAC policy matching EMS tag. This attribute must reference one of the following datasources: `firewall.address.name` .
* `family` - NAC policy matching family.
* `firewall_address` - Dynamic firewall address to associate MAC which match this policy. This attribute must reference one of the following datasources: `firewall.address.name` .
* `host` - NAC policy matching host.
* `hw_vendor` - NAC policy matching hardware vendor.
* `hw_version` - NAC policy matching hardware version.
* `mac` - NAC policy matching MAC address.
* `name` - NAC policy name.
* `os` - NAC policy matching operating system.
* `src` - NAC policy matching source.
* `ssid_policy` - SSID policy to be applied on the matched NAC policy. This attribute must reference one of the following datasources: `wireless-controller.ssid-policy.name` .
* `status` - Enable/disable NAC policy. Valid values: `enable` `disable` .
* `sw_version` - NAC policy matching software version.
* `switch_auto_auth` - NAC device auto authorization when discovered and nac-policy matched. Valid values: `global` `disable` `enable` .
* `switch_fortilink` - FortiLink interface for which this NAC policy belongs to. This attribute must reference one of the following datasources: `system.interface.name` .
* `switch_mac_policy` - Switch MAC policy action to be applied on the matched NAC policy. This attribute must reference one of the following datasources: `switch-controller.mac-policy.name` .
* `switch_port_policy` - switch-port-policy to be applied on the matched NAC policy. This attribute must reference one of the following datasources: `switch-controller.port-policy.name` .
* `type` - NAC policy matching type.
* `user` - NAC policy matching user.
* `user_group` - NAC policy matching user group. This attribute must reference one of the following datasources: `user.group.name` .
* `switch_group` - List of managed FortiSwitch groups on which NAC policy can be applied. The structure of `switch_group` block is documented below.

The `switch_group` block contains:

* `name` - Managed FortiSwitch group name from available options. This attribute must reference one of the following datasources: `switch-controller.switch-group.name` .
* `switch_scope` - List of managed FortiSwitches on which NAC policy can be applied. The structure of `switch_scope` block is documented below.

The `switch_scope` block contains:

* `switch_id` - Managed FortiSwitch name from available options. This attribute must reference one of the following datasources: `switch-controller.managed-switch.switch-id` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_nacpolicy can be imported using:
```sh
terraform import fortios_user_nacpolicy.labelname {{mkey}}
```
