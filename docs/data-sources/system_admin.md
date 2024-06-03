---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_admin"
description: |-
  Get information on a fortios Configure admin users.
---

# Data Source: fortios_system_admin
Use this data source to get information on a fortios Configure admin users.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) User name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `accprofile` - Access profile for this administrator. Access profiles control administrator access to FortiGate features.
* `accprofile_override` - Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access.
* `allow_remove_admin_session` - Enable/disable allow admin session to be removed by privileged admin users.
* `comments` - Comment.
* `email_to` - This administrator's email address.
* `force_password_change` - Enable/disable force password change on next login.
* `fortitoken` - This administrator's FortiToken serial number.
* `guest_auth` - Enable/disable guest authentication.
* `guest_lang` - Guest management portal language.
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
* `peer_auth` - Set to enable peer certificate authentication (for HTTPS admin access).
* `peer_group` - Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).
* `radius_vdom_override` - Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access.
* `remote_auth` - Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server.
* `remote_group` - User group name used for remote auth.
* `schedule` - Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.
* `sms_custom_server` - Custom SMS server to send SMS messages to.
* `sms_phone` - Phone number on which the administrator receives SMS messages.
* `sms_server` - Send SMS messages using the FortiGuard SMS server or a custom server.
* `ssh_certificate` - Select the certificate to be used by the FortiGate for authentication with an SSH client.
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
* `two_factor` - Enable/disable two-factor authentication.
* `two_factor_authentication` - Authentication method by FortiToken Cloud.
* `two_factor_notification` - Notification method for user activation by FortiToken Cloud.
* `vdom_override` - Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access.
* `wildcard` - Enable/disable wildcard RADIUS authentication.
* `guest_usergroups` - Select guest user groups.The structure of `guest_usergroups` block is documented below.

The `guest_usergroups` block contains:

* `name` - Select guest user groups.
* `vdom` - Virtual domain(s) that the administrator can access.The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name.
