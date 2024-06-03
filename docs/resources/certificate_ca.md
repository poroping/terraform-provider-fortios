---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_ca"
description: |-
  CA certificate.
---

## fortios_certificate_ca
CA certificate.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auto_update_days` - Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
* `auto_update_days_warning` - Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
* `ca` - CA certificate as a PEM file.
* `ca_identifier` - CA identifier of the SCEP server.
* `name` - Name.
* `obsolete` - Enable/disable this CA as obsoleted. Valid values: `disable` `enable` .
* `range` - Either global or VDOM IP address range for the CA certificate. Valid values: `global` `vdom` .
* `scep_url` - URL of the SCEP server.
* `source` - CA certificate source type. Valid values: `factory` `user` `bundle` .
* `source_ip` - Source IP address for communications to the SCEP server.
* `ssl_inspection_trusted` - Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_certificate_ca can be imported using:
```sh
terraform import fortios_certificate_ca.labelname {{mkey}}
```
