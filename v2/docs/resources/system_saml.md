---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_saml"
description: |-
  Global settings for SAML authentication.
---

## fortios_system_saml
Global settings for SAML authentication.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `binding_protocol` - IdP Binding protocol. Valid values: `post` `redirect` .
* `cert` - Certificate to sign SAML messages. This attribute must reference one of the following datasources: `certificate.local.name` .
* `default_login_page` - Choose default login page. Valid values: `normal` `sso` .
* `default_profile` - Default profile for new SSO admin. This attribute must reference one of the following datasources: `system.accprofile.name` .
* `entity_id` - SP entity ID.
* `idp_cert` - IDP certificate name. This attribute must reference one of the following datasources: `certificate.remote.name` .
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `life` - Length of the range of time when the assertion is valid (in minutes).
* `portal_url` - SP portal URL.
* `role` - SAML role. Valid values: `identity-provider` `service-provider` .
* `server_address` - Server address.
* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `status` - Enable/disable SAML authentication (default = disable). Valid values: `enable` `disable` .
* `tolerance` - Tolerance to the range of time when the assertion is valid (in minutes).
* `service_providers` - Authorized service providers. The structure of `service_providers` block is documented below.

The `service_providers` block contains:

* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `name` - Name.
* `prefix` - Prefix.
* `sp_binding_protocol` - SP binding protocol. Valid values: `post` `redirect` .
* `sp_cert` - SP certificate name. This attribute must reference one of the following datasources: `certificate.remote.name` .
* `sp_entity_id` - SP entity ID.
* `sp_portal_url` - SP portal URL.
* `sp_single_logout_url` - SP single logout URL.
* `sp_single_sign_on_url` - SP single sign-on URL.
* `assertion_attributes` - Customized SAML attributes to send along with assertion. The structure of `assertion_attributes` block is documented below.

The `assertion_attributes` block contains:

* `name` - Name.
* `type` - Type. Valid values: `username` `email` `profile-name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_saml can be imported using:
```sh
terraform import fortios_system_saml.labelname {{mkey}}
```
