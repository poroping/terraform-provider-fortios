---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_ocspserver"
description: |-
  Get information on a fortios OCSP server configuration.
---

# Data Source: fortios_vpncertificate_ocspserver
Use this data source to get information on a fortios OCSP server configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) OCSP server entry name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cert` - OCSP server certificate.
* `name` - OCSP server entry name.
* `secondary_cert` - Secondary OCSP server certificate.
* `secondary_url` - Secondary OCSP server URL.
* `source_ip` - Source IP address for communications to the OCSP server.
* `unavail_action` - Action when server is unavailable (revoke the certificate or ignore the result of the check).
* `url` - OCSP server URL.
