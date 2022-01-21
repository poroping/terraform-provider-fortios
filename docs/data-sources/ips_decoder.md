---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_decoder"
description: |-
  Get information on a fortios Configure IPS decoder.
---

# Data Source: fortios_ips_decoder
Use this data source to get information on a fortios Configure IPS decoder.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Decoder name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Decoder name.
* `parameter` - IPS group parameters.The structure of `parameter` block is documented below.

The `parameter` block contains:

* `name` - Parameter name.
* `value` - Parameter value.
