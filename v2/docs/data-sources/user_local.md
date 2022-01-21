---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_local"
description: |-
  Get information on a fortios Configure local users.
---

# Data Source: fortios_user_local
Use this data source to get information on a fortios Configure local users.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) User name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_concurrent_override` - Enable/disable overriding the policy-auth-concurrent under config system global.
* `auth_concurrent_value` - Maximum number of concurrent logins permitted from the same user.
* `authtimeout` - Time in minutes before the authentication timeout for a user is reached.
* `email_to` - Two-factor recipient's email address.
* `fortitoken` - Two-factor recipient's FortiToken serial number.
* `id` - User ID.
* `ldap_server` - Name of LDAP server with which the user must authenticate.
* `name` - User name.
* `passwd` - User's password.
* `passwd_policy` - Password policy to apply to this user, as defined in config user password-policy.
* `passwd_time` - Time of the last password update.
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `radius_server` - Name of RADIUS server with which the user must authenticate.
* `sms_custom_server` - Two-factor recipient's SMS server.
* `sms_phone` - Two-factor recipient's mobile phone number.
* `sms_server` - Send SMS through FortiGuard or other external server.
* `status` - Enable/disable allowing the local user to authenticate with the FortiGate unit.
* `tacacs_server` - Name of TACACS+ server with which the user must authenticate.
* `two_factor` - Enable/disable two-factor authentication.
* `two_factor_authentication` - Authentication method by FortiToken Cloud.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud.
* `type` - Authentication method.
* `username_case_insensitivity` - Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent).
* `username_case_sensitivity` - Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent).
* `username_sensitivity` - Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled).
* `workstation` - Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.
