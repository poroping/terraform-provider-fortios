---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peer"
description: |-
  Get information on a fortios Configure peer users.
---

# Data Source: fortios_user_peer
Use this data source to get information on a fortios Configure peer users.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Peer name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ca` - Name of the CA certificate.
* `cn` - Peer certificate common name.
* `cn_type` - Peer certificate common name type.
* `ldap_mode` - Mode for LDAP peer authentication.
* `ldap_password` - Password for LDAP server bind.
* `ldap_server` - Name of an LDAP server defined under the user ldap command. Performs client access rights check.
* `ldap_username` - Username for LDAP server bind.
* `mandatory_ca_verify` - Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid.
* `name` - Peer name.
* `ocsp_override_server` - Online Certificate Status Protocol (OCSP) server for certificate retrieval.
* `passwd` - Peer's password used for two-factor authentication.
* `subject` - Peer certificate name constraints.
* `two_factor` - Enable/disable two-factor authentication, applying certificate and password-based authentication.
