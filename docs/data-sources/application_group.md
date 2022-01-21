---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_group"
description: |-
  Get information on a fortios Configure firewall application groups.
---

# Data Source: fortios_application_group
Use this data source to get information on a fortios Configure firewall application groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Application group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `behavior` - Application behavior filter.
* `comment` - Comment
* `name` - Application group name.
* `popularity` - Application popularity filter (1 - 5, from least to most popular).
* `protocols` - Application protocol filter.
* `technology` - Application technology filter.
* `type` - Application group type.
* `vendor` - Application vendor filter.
* `application` - Application ID list.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `category` - Application category ID list.The structure of `category` block is documented below.

The `category` block contains:

* `id` - Category IDs.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).The structure of `risk` block is documented below.

The `risk` block contains:

* `level` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).
