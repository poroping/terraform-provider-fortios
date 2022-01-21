---
subcategory: "FortiGate Ips"
layout: "fortios"
page_title: "FortiOS: fortios_ips_decoder"
description: |-
  Configure IPS decoder.
---

## fortios_ips_decoder
Configure IPS decoder.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Decoder name.
* `parameter` - IPS group parameters. The structure of `parameter` block is documented below.

The `parameter` block contains:

* `name` - Parameter name.
* `value` - Parameter value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_ips_decoder can be imported using:
```sh
terraform import fortios_ips_decoder.labelname {{mkey}}
```
