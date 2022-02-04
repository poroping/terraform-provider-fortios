---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollersecuritypolicy_8021x"
description: |-
  Configure 802.1x MAC Authentication Bypass (MAB) policies.
---

## fortios_switchcontrollersecuritypolicy_8021x
Configure 802.1x MAC Authentication Bypass (MAB) policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth_fail_vlan` - Enable to allow limited access to clients that cannot authenticate. Valid values: `disable` `enable` .
* `auth_fail_vlan_id` - VLAN ID on which authentication failed. This attribute must reference one of the following datasources: `system.interface.name` .
* `authserver_timeout_period` - Authentication server timeout period (3 - 15 sec, default = 3).
* `authserver_timeout_vlan` - Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable. Valid values: `disable` `enable` .
* `authserver_timeout_vlanid` - Authentication server timeout VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .
* `eap_auto_untagged_vlans` - Enable/disable automatic inclusion of untagged VLANs. Valid values: `disable` `enable` .
* `eap_passthru` - Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication. Valid values: `disable` `enable` .
* `framevid_apply` - Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN. Valid values: `disable` `enable` .
* `guest_auth_delay` - Guest authentication delay (1 - 900  sec, default = 30).
* `guest_vlan` - Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients. Valid values: `disable` `enable` .
* `guest_vlan_id` - Guest VLAN name. This attribute must reference one of the following datasources: `system.interface.name` .
* `mac_auth_bypass` - Enable/disable MAB for this policy. Valid values: `disable` `enable` .
* `name` - Policy name.
* `open_auth` - Enable/disable open authentication for this policy. Valid values: `disable` `enable` .
* `policy_type` - Policy type. Valid values: `802.1X` .
* `radius_timeout_overwrite` - Enable to override the global RADIUS session timeout. Valid values: `disable` `enable` .
* `security_mode` - Port or MAC based 802.1X security mode. Valid values: `802.1X` `802.1X-mac-based` .
* `user_group` - Name of user-group to assign to this MAC Authentication Bypass (MAB) policy. The structure of `user_group` block is documented below.

The `user_group` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_switchcontrollersecuritypolicy_8021x can be imported using:
```sh
terraform import fortios_switchcontrollersecuritypolicy_8021x.labelname {{mkey}}
```
