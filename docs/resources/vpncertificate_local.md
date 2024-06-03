---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_local"
description: |-
  Local keys and certificates.
---

## fortios_vpncertificate_local
Local keys and certificates.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `acme_ca_url` - The URL for the ACME CA server (Let's Encrypt is the default provider).
* `acme_domain` - A valid domain that resolves to this FortiGate unit.
* `acme_email` - Contact email address that is required by some CAs like LetsEncrypt.
* `acme_renew_window` - Beginning of the renewal window (in days before certificate expiration, 30 by default).
* `acme_rsa_key_size` - Length of the RSA private key of the generated cert (Minimum 2048 bits).
* `auto_regenerate_days` - Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
* `auto_regenerate_days_warning` - Number of days to wait before an expiry warning message is generated (0 = disabled).
* `ca_identifier` - CA identifier of the CA server for signing via SCEP.
* `certificate` - PEM format certificate.
* `cmp_path` - Path location inside CMP server.
* `cmp_regeneration_method` - CMP auto-regeneration method. Valid values: `keyupate` `renewal` .
* `cmp_server` - Address and port for CMP server (format = address:port).
* `cmp_server_cert` - CMP server certificate. This attribute must reference one of the following datasources: `vpn.certificate.ca.name` `vpn.certificate.remote.name` .
* `comments` - Comment.
* `csr` - Certificate Signing Request.
* `enroll_protocol` - Certificate enrollment protocol. Valid values: `none` `scep` `cmpv2` `acme2` .
* `ike_localid` - Local ID the FortiGate uses for authentication as a VPN client.
* `ike_localid_type` - IKE local ID type. Valid values: `asn1dn` `fqdn` .
* `name` - Name.
* `name_encoding` - Name encoding method for auto-regeneration. Valid values: `printable` `utf8` .
* `password` - Password as a PEM file.
* `private_key` - PEM format key encrypted with a password.
* `private_key_retain` - Enable/disable retention of private key during SCEP renewal (default = disable). Valid values: `enable` `disable` .
* `range` - Either a global or VDOM IP address range for the certificate. Valid values: `global` `vdom` .
* `scep_password` - SCEP server challenge password for auto-regeneration.
* `scep_url` - SCEP server URL.
* `source` - Certificate source type. Valid values: `factory` `user` `bundle` .
* `source_ip` - Source IP address for communications to the SCEP server.
* `state` - Certificate Signing Request State.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpncertificate_local can be imported using:
```sh
terraform import fortios_vpncertificate_local.labelname {{mkey}}
```
