---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_local"
description: |-
  Configure local users.
---

## fortios_user_local
Configure local users.

## Example Usage

```hcl
resource "fortios_user_local" "example" {
  name   = "foobar"
  passwd = "changeme"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auth_concurrent_override` - Enable/disable overriding the policy-auth-concurrent under config system global. Valid values: `enable` `disable` .
* `auth_concurrent_value` - Maximum number of concurrent logins permitted from the same user.
* `authtimeout` - Time in minutes before the authentication timeout for a user is reached.
* `email_to` - Two-factor recipient's email address.
* `fortitoken` - Two-factor recipient's FortiToken serial number. This attribute must reference one of the following datasources: `user.fortitoken.serial-number` .
* `id` - User ID.
* `ldap_server` - Name of LDAP server with which the user must authenticate. This attribute must reference one of the following datasources: `user.ldap.name` .
* `name` - Local user name.
* `passwd` - User's password.
* `passwd_policy` - Password policy to apply to this user, as defined in config user password-policy. This attribute must reference one of the following datasources: `user.password-policy.name` .
* `passwd_time` - Time of the last password update.
* `ppk_identity` - IKEv2 Postquantum Preshared Key Identity.
* `ppk_secret` - IKEv2 Postquantum Preshared Key (ASCII string or hexadecimal encoded with a leading 0x).
* `radius_server` - Name of RADIUS server with which the user must authenticate. This attribute must reference one of the following datasources: `user.radius.name` .
* `sms_custom_server` - Two-factor recipient's SMS server. This attribute must reference one of the following datasources: `system.sms-server.name` .
* `sms_phone` - Two-factor recipient's mobile phone number.
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard` `custom` .
* `status` - Enable/disable allowing the local user to authenticate with the FortiGate unit. Valid values: `enable` `disable` .
* `tacacs_server` - Name of TACACS+ server with which the user must authenticate. This attribute must reference one of the following datasources: `user.tacacs+.name` .
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable` `fortitoken` `fortitoken-cloud` `email` `sms` .
* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken` `email` `sms` .
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email` `sms` .
* `type` - Authentication method. Valid values: `password` `radius` `tacacs+` `ldap` .
* `username_case_insensitivity` - Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `enable` `disable` .
* `username_case_sensitivity` - Enable/disable case sensitivity when performing username matching (uppercase and lowercase letters are treated either as distinct or equivalent). Valid values: `disable` `enable` .
* `username_sensitivity` - Enable/disable case and accent sensitivity when performing username matching (accents are stripped and case is ignored when disabled). Valid values: `disable` `enable` .
* `workstation` - Name of the remote user workstation, if you want to limit the user to authenticate only from a particular workstation.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_local can be imported using:
```sh
terraform import fortios_user_local.labelname {{mkey}}
```
