---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_remote"
description: |-
  Get information on a fortios Remote certificate as a PEM file.
---

# Data Source: fortios_certificate_remote
Use this data source to get information on a fortios Remote certificate as a PEM file.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `range` - Either the global or VDOM IP address range for the remote certificate.
* `remote` - Remote certificate.
* `source` - Remote certificate source type.
