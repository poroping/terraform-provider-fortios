---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_name"
description: |-
  Get information on a fortios Configure application signatures.
---

# Data Source: fortios_application_name
Use this data source to get information on a fortios Configure application signatures.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Application name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

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
* `metadata` - Meta data.The structure of `metadata` block is documented below.

The `metadata` block contains:

* `id` - ID.
* `metaid` - Meta ID.
* `valueid` - Value ID.
* `parameters` - Application parameters.The structure of `parameters` block is documented below.

The `parameters` block contains:

* `defaultvalue` - Parameter default value.
* `name` - Parameter name.
