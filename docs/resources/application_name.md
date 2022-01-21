---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_name"
description: |-
  Configure application signatures.
---

## fortios_application_name
Configure application signatures.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `behavior` - Application behavior.
* `category` - Application category ID.
* `id` - Application ID.
* `name` - Application name.
* `parameter` - Application parameter name.
* `popularity` - Application popularity.
* `protocol` - Application protocol.
* `risk` - Application risk.
* `sub_category` - Application sub-category ID.
* `technology` - Application technology.
* `vendor` - Application vendor.
* `weight` - Application weight.
* `metadata` - Meta data. The structure of `metadata` block is documented below.

The `metadata` block contains:

* `id` - ID.
* `metaid` - Meta ID.
* `valueid` - Value ID.
* `parameters` - Application parameters. The structure of `parameters` block is documented below.

The `parameters` block contains:

* `defaultvalue` - Parameter default value.
* `name` - Parameter name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_application_name can be imported using:
```sh
terraform import fortios_application_name.labelname {{mkey}}
```
