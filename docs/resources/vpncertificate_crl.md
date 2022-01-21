---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_crl"
description: |-
  Certificate Revocation List as a PEM file.
---

## fortios_vpncertificate_crl
Certificate Revocation List as a PEM file.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `crl` - Certificate Revocation List as a PEM file.
* `http_url` - HTTP server URL for CRL auto-update.
* `ldap_password` - LDAP server user password.
* `ldap_server` - LDAP server name for CRL auto-update.
* `ldap_username` - LDAP server user name.
* `name` - Name.
* `range` - Either global or VDOM IP address range for the certificate. Valid values: `global` `vdom` .
* `scep_cert` - Local certificate for SCEP communication for CRL auto-update. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `scep_url` - SCEP server URL for CRL auto-update.
* `source` - Certificate source type. Valid values: `factory` `user` `bundle` .
* `source_ip` - Source IP address for communications to a HTTP or SCEP CA server.
* `update_interval` - Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
* `update_vdom` - VDOM for CRL update. This attribute must reference one of the following datasources: `system.vdom.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpncertificate_crl can be imported using:
```sh
terraform import fortios_vpncertificate_crl.labelname {{mkey}}
```
