---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_crl"
description: |-
  Get information on a fortios Certificate Revocation List as a PEM file.
---

# Data Source: fortios_certificate_crl
Use this data source to get information on a fortios Certificate Revocation List as a PEM file.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `crl` - Certificate Revocation List as a PEM file.
* `http_url` - HTTP server URL for CRL auto-update.
* `ldap_password` - LDAP server user password.
* `ldap_server` - LDAP server name for CRL auto-update.
* `ldap_username` - LDAP server user name.
* `name` - Name.
* `range` - Either global or VDOM IP address range for the certificate.
* `scep_cert` - Local certificate for SCEP communication for CRL auto-update.
* `scep_url` - SCEP server URL for CRL auto-update.
* `source` - Certificate source type.
* `source_ip` - Source IP address for communications to a HTTP or SCEP CA server.
* `update_interval` - Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
* `update_vdom` - VDOM for CRL update.
