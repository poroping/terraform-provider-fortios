---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_group"
description: |-
  Configure user groups.
---

## fortios_user_group
Configure user groups.

## Example Usage

```hcl
resource "fortios_user_group" "example" {
  name = "foobar"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `auth_concurrent_override` - Enable/disable overriding the global number of concurrent authentication sessions for this user group. Valid values: `enable` `disable` .
* `auth_concurrent_value` - Maximum number of concurrent authenticated connections per user (0 - 100).
* `authtimeout` - Authentication timeout in minutes for this user group. 0 to use the global user setting auth-timeout.
* `company` - Set the action for the company guest user field. Valid values: `optional` `mandatory` `disabled` .
* `email` - Enable/disable the guest user email address field. Valid values: `disable` `enable` .
* `expire` - Time in seconds before guest user accounts expire (1 - 31536000).
* `expire_type` - Determine when the expiration countdown begins. Valid values: `immediately` `first-successful-login` .
* `group_type` - Set the group to be for firewall authentication, FSSO, RSSO, or guest users. Valid values: `firewall` `fsso-service` `rsso` `guest` .
* `http_digest_realm` - Realm attribute for MD5-digest authentication.
* `id` - Group ID.
* `max_accounts` - Maximum number of guest accounts that can be created for this group (0 means unlimited).
* `mobile_phone` - Enable/disable the guest user mobile phone number field. Valid values: `disable` `enable` .
* `multiple_guest_add` - Enable/disable addition of multiple guests. Valid values: `disable` `enable` .
* `name` - Group name.
* `password` - Guest user password type. Valid values: `auto-generate` `specify` `disable` .
* `sms_custom_server` - SMS server. This attribute must reference one of the following datasources: `system.sms-server.name` .
* `sms_server` - Send SMS through FortiGuard or other external server. Valid values: `fortiguard` `custom` .
* `sponsor` - Set the action for the sponsor guest user field. Valid values: `optional` `mandatory` `disabled` .
* `sso_attribute_value` - Name of the RADIUS user group that this local user group represents.
* `user_id` - Guest user ID type. Valid values: `email` `auto-generate` `specify` .
* `user_name` - Enable/disable the guest user name entry. Valid values: `disable` `enable` .
* `guest` - Guest User. The structure of `guest` block is documented below.

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
* `match` - Group matches. The structure of `match` block is documented below.

The `match` block contains:

* `group_name` - Name of matching user or group on remote authentication server.
* `id` - ID.
* `server_name` - Name of remote auth server. This attribute must reference one of the following datasources: `user.radius.name` `user.ldap.name` `user.tacacs+.name` `user.saml.name` .
* `member` - Names of users, peers, LDAP severs, or RADIUS servers to add to the user group. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Group member name. This attribute must reference one of the following datasources: `user.peer.name` `user.local.name` `user.radius.name` `user.tacacs+.name` `user.ldap.name` `user.saml.name` `user.adgrp.name` `user.pop3.name` `user.certificate.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_group can be imported using:
```sh
terraform import fortios_user_group.labelname {{mkey}}
```
