---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_group"
description: |-
  Configure firewall application groups.
---

## fortios_application_group
Configure firewall application groups.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `behavior` - Application behavior filter.
* `comment` - Comment
* `name` - Application group name.
* `popularity` - Application popularity filter (1 - 5, from least to most popular). Valid values: `1` `2` `3` `4` `5` .
* `protocols` - Application protocol filter.
* `technology` - Application technology filter.
* `type` - Application group type. Valid values: `application` `filter` .
* `vendor` - Application vendor filter.
* `application` - Application ID list. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `category` - Application category ID list. The structure of `category` block is documented below.

The `category` block contains:

* `id` - Category IDs.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.

The `risk` block contains:

* `level` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_application_group can be imported using:
```sh
terraform import fortios_application_group.labelname {{mkey}}
```
