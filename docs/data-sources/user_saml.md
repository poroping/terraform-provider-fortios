---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_saml"
description: |-
  Get information on a fortios SAML server entry configuration.
---

# Data Source: fortios_user_saml
Use this data source to get information on a fortios SAML server entry configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) SAML server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `adfs_claim` - Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable).
* `cert` - Certificate to sign SAML messages.
* `digest_method` - Digest Method Algorithm. (default = sha1).
* `entity_id` - SP entity ID.
* `group_claim_type` - Group claim in assertion statement.
* `group_name` - Group name in assertion statement.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `limit_relaystate` - Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes).
* `name` - SAML server entry name.
* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `user_claim_type` - User name claim in assertion statement.
* `user_name` - User name in assertion statement.
