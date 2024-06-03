---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_saml"
description: |-
  SAML server entry configuration.
---

## fortios_user_saml
SAML server entry configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `adfs_claim` - Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable` `disable` .
* `auth_url` - URL to verify authentication.
* `cert` - Certificate to sign SAML messages. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `clock_tolerance` - Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
* `digest_method` - Digest method algorithm (default = sha1). Valid values: `sha1` `sha256` .
* `entity_id` - SP entity ID.
* `group_claim_type` - Group claim in assertion statement. Valid values: `email` `given-name` `name` `upn` `common-name` `email-adfs-1x` `group` `upn-adfs-1x` `role` `sur-name` `ppid` `name-identifier` `authentication-method` `deny-only-group-sid` `deny-only-primary-sid` `deny-only-primary-group-sid` `group-sid` `primary-group-sid` `primary-sid` `windows-account-name` .
* `group_name` - Group name in assertion statement.
* `idp_cert` - IDP Certificate name. This attribute must reference one of the following datasources: `vpn.certificate.remote.name` .
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `limit_relaystate` - Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable` `disable` .
* `name` - SAML server entry name.
* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `user_claim_type` - User name claim in assertion statement. Valid values: `email` `given-name` `name` `upn` `common-name` `email-adfs-1x` `group` `upn-adfs-1x` `role` `sur-name` `ppid` `name-identifier` `authentication-method` `deny-only-group-sid` `deny-only-primary-sid` `deny-only-primary-group-sid` `group-sid` `primary-group-sid` `primary-sid` `windows-account-name` .
* `user_name` - User name in assertion statement.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_saml can be imported using:
```sh
terraform import fortios_user_saml.labelname {{mkey}}
```
