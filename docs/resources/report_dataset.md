---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_dataset"
description: |-
  Report dataset configuration.
---

## fortios_report_dataset
Report dataset configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Name.
* `policy` - Used by monitor policy.
* `query` - SQL query statement.
* `field` - Fields. The structure of `field` block is documented below.

The `field` block contains:

* `displayname` - Display name.
* `id` - Field ID (1 to number of columns in SQL result).
* `name` - Name.
* `type` - Field type. Valid values: `text` `integer` `double` .
* `parameters` - Parameters. The structure of `parameters` block is documented below.

The `parameters` block contains:

* `data_type` - Data type. Valid values: `text` `integer` `double` `long-integer` `date-time` .
* `display_name` - Display name.
* `field` - SQL field name.
* `id` - Parameter ID (1 to number of columns in SQL result).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_report_dataset can be imported using:
```sh
terraform import fortios_report_dataset.labelname {{mkey}}
```
