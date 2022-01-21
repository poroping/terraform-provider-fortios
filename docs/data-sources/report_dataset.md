---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_dataset"
description: |-
  Get information on a fortios Report dataset configuration.
---

# Data Source: fortios_report_dataset
Use this data source to get information on a fortios Report dataset configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `policy` - Used by monitor policy.
* `query` - SQL query statement.
* `field` - Fields.The structure of `field` block is documented below.

The `field` block contains:

* `displayname` - Display name.
* `id` - Field ID (1 to number of columns in SQL result).
* `name` - Name.
* `type` - Field type.
* `parameters` - Parameters.The structure of `parameters` block is documented below.

The `parameters` block contains:

* `data_type` - Data type.
* `display_name` - Display name.
* `field` - SQL field name.
* `id` - Parameter ID (1 to number of columns in SQL result).
