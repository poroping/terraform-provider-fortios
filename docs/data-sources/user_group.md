---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_group"
description: |-
  Get information on a fortios Configure user groups.
---

# Data Source: fortios_user_group
Use this data source to get information on a fortios Configure user groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auth_concurrent_override` - Enable/disable overriding the global number of concurrent authentication sessions for this user group.
* `auth_concurrent_value` - Maximum number of concurrent authenticated connections per user (0 - 100).
* `authtimeout` - Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
* `company` - Set the action for the company guest user field.
* `email` - Enable/disable the guest user email address field.
* `expire` - Time in seconds before guest user accounts expire (1 - 31536000).
* `expire_type` - Determine when the expiration countdown begins.
* `group_type` - Set the group to be for firewall authentication, FSSO, RSSO, or guest users.
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `id` - Group ID.
* `max_accounts` - Maximum number of guest accounts that can be created for this group (0 means unlimited).
* `mobile_phone` - Enable/disable the guest user mobile phone number field.
* `multiple_guest_add` - Enable/disable addition of multiple guests.
* `name` - Group name.
* `password` - Guest user password type.
* `sms_custom_server` - SMS server.
* `sms_server` - Send SMS through FortiGuard or other external server.
* `sponsor` - Set the action for the sponsor guest user field.
* `sso_attribute_value` - Name of the RADIUS user group that this local user group represents.
* `user_id` - Guest user ID type.
* `user_name` - Enable/disable the guest user name entry.
* `guest` - Guest User.The structure of `guest` block is documented below.

The `guest` block contains:

* `comment` - Comment.
* `company` - Set the action for the company guest user field.
* `email` - Email.
* `expiration` - Expire time.
* `id` - Guest ID.
* `mobile_phone` - Mobile phone.
* `name` - Guest name.
* `password` - Guest password.
* `sponsor` - Set the action for the sponsor guest user field.
* `user_id` - Guest ID.
* `match` - Group matches.The structure of `match` block is documented below.

The `match` block contains:

* `group_name` - Name of matching user or group on remote authentication server.
* `id` - ID.
* `server_name` - Name of remote auth server.
* `member` - Names of users, peers, LDAP severs, RADIUS servers or external idp servers to add to the user group.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Group member name.
