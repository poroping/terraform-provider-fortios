---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_saml"
description: |-
  Get information on a fortios Global settings for SAML authentication.
---

# Data Source: fortios_system_saml
Use this data source to get information on a fortios Global settings for SAML authentication.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `binding_protocol` - IdP Binding protocol.
* `cert` - Certificate to sign SAML messages.
* `default_login_page` - Choose default login page.
* `default_profile` - Default profile for new SSO admin.
* `entity_id` - SP entity ID.
* `idp_cert` - IDP certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `life` - Length of the range of time when the assertion is valid (in minutes).
* `portal_url` - SP portal URL.
* `role` - SAML role.
* `server_address` - Server address.
* `single_logout_url` - SP single logout URL.
* `single_sign_on_url` - SP single sign-on URL.
* `status` - Enable/disable SAML authentication (default = disable).
* `tolerance` - Tolerance to the range of time when the assertion is valid (in minutes).
* `service_providers` - Authorized service providers.The structure of `service_providers` block is documented below.

The `service_providers` block contains:

* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout URL.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `name` - Name.
* `prefix` - Prefix.
* `sp_binding_protocol` - SP binding protocol.
* `sp_cert` - SP certificate name.
* `sp_entity_id` - SP entity ID.
* `sp_portal_url` - SP portal URL.
* `sp_single_logout_url` - SP single logout URL.
* `sp_single_sign_on_url` - SP single sign-on URL.
* `assertion_attributes` - Customized SAML attributes to send along with assertion.The structure of `assertion_attributes` block is documented below.

The `assertion_attributes` block contains:

* `name` - Name.
* `type` - Type.
