---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_scheme"
description: |-
  Get information on a fortios Configure Authentication Schemes.
---

# Data Source: fortios_authentication_scheme
Use this data source to get information on a fortios Configure Authentication Schemes.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Authentication scheme name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `domain_controller` - Domain controller setting.
* `ems_device_owner` - Enable/disable SSH public-key authentication with device owner (default = disable).
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `fsso_guest` - Enable/disable user fsso-guest authentication (default = disable).
* `kerberos_keytab` - Kerberos keytab setting.
* `method` - Authentication methods (default = basic).
* `name` - Authentication scheme name.
* `negotiate_ntlm` - Enable/disable negotiate authentication for NTLM (default = disable).
* `require_tfa` - Enable/disable two-factor authentication (default = disable).
* `saml_server` - SAML configuration.
* `saml_timeout` - SAML authentication timeout in seconds.
* `ssh_ca` - SSH CA name.
* `user_cert` - Enable/disable authentication with user certificate (default = disable).
* `user_database` - Authentication server to contain user information; "local" (default) or "123" (for LDAP).The structure of `user_database` block is documented below.

The `user_database` block contains:

* `name` - Authentication server name.
