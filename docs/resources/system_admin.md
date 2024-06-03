---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_admin"
description: |-
  Configure admin users.
---

## fortios_system_admin
Configure admin users.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features. This attribute must reference one of the following datasources: `system.accprofile.name` .
* `accprofile_override` - Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access. Valid values: `enable` `disable` .
* `allow_remove_admin_session` - Enable/disable allow admin session to be removed by privileged admin users. Valid values: `enable` `disable` .
* `comments` - Comment.
* `email_to` - This administrator's email address.
* `force_password_change` - Enable/disable force password change on next login. Valid values: `enable` `disable` .
* `fortitoken` - This administrator's FortiToken serial number.
* `guest_auth` - Enable/disable guest authentication. Valid values: `disable` `enable` .
* `guest_lang` - Guest management portal language. This attribute must reference one of the following datasources: `system.custom-language.name` .
* `ip6_trusthost1` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost10` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost2` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost3` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost4` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost5` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost6` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost7` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost8` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `ip6_trusthost9` - Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
* `name` - User name.
* `password` - Admin user password.
* `password_expire` - Password expire time.
* `peer_auth` - Set to enable peer certificate authentication (for HTTPS admin access). Valid values: `enable` `disable` .
* `peer_group` - Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).
* `radius_vdom_override` - Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access. Valid values: `enable` `disable` .
* `remote_auth` - Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server. Valid values: `enable` `disable` .
* `remote_group` - User group name used for remote auth.
* `schedule` - Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.
* `sms_custom_server` - Custom SMS server to send SMS messages to. This attribute must reference one of the following datasources: `system.sms-server.name` .
* `sms_phone` - Phone number on which the administrator receives SMS messages.
* `sms_server` - Send SMS messages using the FortiGuard SMS server or a custom server. Valid values: `fortiguard` `custom` .
* `ssh_certificate` - Select the certificate to be used by the FortiGate for authentication with an SSH client. This attribute must reference one of the following datasources: `certificate.remote.name` .
* `ssh_public_key1` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_public_key2` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `ssh_public_key3` - Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
* `trusthost1` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost10` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost2` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost3` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost4` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost5` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost6` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost7` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost8` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `trusthost9` - Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
* `two_factor` - Enable/disable two-factor authentication. Valid values: `disable` `fortitoken` `fortitoken-cloud` `email` `sms` .
* `two_factor_authentication` - Authentication method by FortiToken Cloud. Valid values: `fortitoken` `email` `sms` .
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud. Valid values: `email` `sms` .
* `vdom_override` - Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access. Valid values: `enable` `disable` .
* `wildcard` - Enable/disable wildcard RADIUS authentication. Valid values: `enable` `disable` .
* `guest_usergroups` - Select guest user groups. The structure of `guest_usergroups` block is documented below.

The `guest_usergroups` block contains:

* `name` - Select guest user groups.
* `vdom` - Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_admin can be imported using:
```sh
terraform import fortios_system_admin.labelname {{mkey}}
```
