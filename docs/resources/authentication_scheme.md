---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_scheme"
description: |-
  Configure Authentication Schemes.
---

## fortios_authentication_scheme
Configure Authentication Schemes.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `domain_controller` - Domain controller setting. This attribute must reference one of the following datasources: `user.domain-controller.name` .
* `ems_device_owner` - Enable/disable SSH public-key authentication with device owner (default = disable). Valid values: `enable` `disable` .
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication. This attribute must reference one of the following datasources: `user.fsso.name` .
* `fsso_guest` - Enable/disable user fsso-guest authentication (default = disable). Valid values: `enable` `disable` .
* `kerberos_keytab` - Kerberos keytab setting. This attribute must reference one of the following datasources: `user.krb-keytab.name` .
* `method` - Authentication methods (default = basic). Valid values: `ntlm` `basic` `digest` `form` `negotiate` `fsso` `rsso` `ssh-publickey` `cert` `saml` .
* `name` - Authentication scheme name.
* `negotiate_ntlm` - Enable/disable negotiate authentication for NTLM (default = disable). Valid values: `enable` `disable` .
* `require_tfa` - Enable/disable two-factor authentication (default = disable). Valid values: `enable` `disable` .
* `saml_server` - SAML configuration. This attribute must reference one of the following datasources: `user.saml.name` .
* `saml_timeout` - SAML authentication timeout in seconds.
* `ssh_ca` - SSH CA name. This attribute must reference one of the following datasources: `firewall.ssh.local-ca.name` .
* `user_cert` - Enable/disable authentication with user certificate (default = disable). Valid values: `enable` `disable` .
* `user_database` - Authentication server to contain user information; "local" (default) or "123" (for LDAP). The structure of `user_database` block is documented below.

The `user_database` block contains:

* `name` - Authentication server name. This attribute must reference one of the following datasources: `system.datasource.name` `user.radius.name` `user.tacacs+.name` `user.ldap.name` `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_authentication_scheme can be imported using:
```sh
terraform import fortios_authentication_scheme.labelname {{mkey}}
```
