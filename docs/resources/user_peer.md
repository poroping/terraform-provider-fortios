---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peer"
description: |-
  Configure peer users.
---

## fortios_user_peer
Configure peer users.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `ca` - Name of the CA certificate. This attribute must reference one of the following datasources: `vpn.certificate.ca.name` .
* `cn` - Peer certificate common name.
* `cn_type` - Peer certificate common name type. Valid values: `string` `email` `FQDN` `ipv4` `ipv6` .
* `ldap_mode` - Mode for LDAP peer authentication. Valid values: `password` `principal-name` .
* `ldap_password` - Password for LDAP server bind.
* `ldap_server` - Name of an LDAP server defined under the user ldap command. Performs client access rights check. This attribute must reference one of the following datasources: `user.ldap.name` .
* `ldap_username` - Username for LDAP server bind.
* `mandatory_ca_verify` - Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable` `disable` .
* `name` - Peer name.
* `ocsp_override_server` - Online Certificate Status Protocol (OCSP) server for certificate retrieval. This attribute must reference one of the following datasources: `vpn.certificate.ocsp-server.name` .
* `passwd` - Peer's password used for two-factor authentication.
* `subject` - Peer certificate name constraints.
* `two_factor` - Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_user_peer can be imported using:
```sh
terraform import fortios_user_peer.labelname {{mkey}}
```
