---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_contentheader"
description: |-
  Get information on a fortios Configure content types used by Web filter.
---

# Data Source: fortios_webfilter_contentheader
Use this data source to get information on a fortios Configure content types used by Web filter.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - Configure content types used by web filter.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Action to take for this content type.
* `category` - Categories that this content type applies to.
* `pattern` - Content type (regular expression).
