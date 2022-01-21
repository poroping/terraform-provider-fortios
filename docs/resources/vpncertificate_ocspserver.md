---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_ocspserver"
description: |-
  OCSP server configuration.
---

## fortios_vpncertificate_ocspserver
OCSP server configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `cert` - OCSP server certificate. This attribute must reference one of the following datasources: `vpn.certificate.remote.name` `vpn.certificate.ca.name` .
* `name` - OCSP server entry name.
* `secondary_cert` - Secondary OCSP server certificate. This attribute must reference one of the following datasources: `vpn.certificate.remote.name` `vpn.certificate.ca.name` .
* `secondary_url` - Secondary OCSP server URL.
* `source_ip` - Source IP address for communications to the OCSP server.
* `unavail_action` - Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke` `ignore` .
* `url` - OCSP server URL.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_vpncertificate_ocspserver can be imported using:
```sh
terraform import fortios_vpncertificate_ocspserver.labelname {{mkey}}
```
