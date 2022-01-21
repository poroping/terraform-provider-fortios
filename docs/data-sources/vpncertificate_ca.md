---
subcategory: "FortiGate Vpn"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_ca"
description: |-
  Get information on a fortios CA certificate.
---

# Data Source: fortios_vpncertificate_ca
Use this data source to get information on a fortios CA certificate.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `ca` - CA certificate as a PEM file.
* `ca_identifier` - CA identifier of the SCEP server.
* `name` - Name.
* `range` - Either global or VDOM IP address range for the CA certificate.
* `scep_url` - URL of the SCEP server.
* `source` - CA certificate source type.
* `source_ip` - Source IP address for communications to the SCEP server.
* `ssl_inspection_trusted` - Enable/disable this CA as a trusted CA for SSL inspection.
