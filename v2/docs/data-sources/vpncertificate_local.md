---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_local"
description: |-
  Get information on a fortios Local keys and certificates.
---

# Data Source: fortios_vpncertificate_local
Use this data source to get information on a fortios Local keys and certificates.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `acme_ca_url` - The URL for the ACME CA server (Let's Encrypt is the default provider).
* `acme_domain` - A valid domain that resolves to this Fortigate.
* `acme_email` - Contact email address that is required by some CAs like LetsEncrypt.
* `acme_renew_window` - Beginning of the renewal window (in days before certificate expiration, 30 by default).
* `acme_rsa_key_size` - Length of the RSA private key of the generated cert (Minimum 2048 bits).
* `auto_regenerate_days` - Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).
* `auto_regenerate_days_warning` - Number of days to wait before an expiry warning message is generated (0 = disabled).
* `ca_identifier` - CA identifier of the CA server for signing via SCEP.
* `certificate` - PEM format certificate.
* `cmp_path` - Path location inside CMP server.
* `cmp_regeneration_method` - CMP auto-regeneration method.
* `cmp_server` - 'ADDRESS:PORT' for CMP server.
* `cmp_server_cert` - CMP server certificate.
* `comments` - Comment.
* `csr` - Certificate Signing Request.
* `enroll_protocol` - Certificate enrollment protocol.
* `ike_localid` - Local ID the FortiGate uses for authentication as a VPN client.
* `ike_localid_type` - IKE local ID type.
* `name` - Name.
* `name_encoding` - Name encoding method for auto-regeneration.
* `password` - Password as a PEM file.
* `private_key` - PEM format key, encrypted with a password.
* `range` - Either a global or VDOM IP address range for the certificate.
* `scep_password` - SCEP server challenge password for auto-regeneration.
* `scep_url` - SCEP server URL.
* `source` - Certificate source type.
* `source_ip` - Source IP address for communications to the SCEP server.
* `state` - Certificate Signing Request State.
