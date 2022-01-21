---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollersecuritypolicy_8021x"
description: |-
  Get information on a fortios Configure 802.1x MAC Authentication Bypass (MAB) policies.
---

# Data Source: fortios_switchcontrollersecuritypolicy_8021x
Use this data source to get information on a fortios Configure 802.1x MAC Authentication Bypass (MAB) policies.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Policy name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_fail_vlan` - Enable to allow limited access to clients that cannot authenticate.
* `auth_fail_vlan_id` - VLAN ID on which authentication failed.
* `authserver_timeout_period` - Authentication server timeout period (3 - 15 sec, default = 3).
* `authserver_timeout_vlan` - Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable. 
* `authserver_timeout_vlanid` - Authentication server timeout VLAN name.
* `eap_auto_untagged_vlans` - Enable/disable automatic inclusion of untagged VLANs.
* `eap_passthru` - Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication.
* `framevid_apply` - Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN.
* `guest_auth_delay` - Guest authentication delay (1 - 900  sec, default = 30).
* `guest_vlan` - Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients.
* `guest_vlan_id` - Guest VLAN name.
* `mac_auth_bypass` - Enable/disable MAB for this policy.
* `name` - Policy name.
* `open_auth` - Enable/disable open authentication for this policy.
* `policy_type` - Policy type.
* `radius_timeout_overwrite` - Enable to override the global RADIUS session timeout.
* `security_mode` - Port or MAC based 802.1X security mode.
* `user_group` - Name of user-group to assign to this MAC Authentication Bypass (MAB) policy.The structure of `user_group` block is documented below.

The `user_group` block contains:

* `name` - Group name.
