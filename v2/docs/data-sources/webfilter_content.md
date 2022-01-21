---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_content"
description: |-
  Get information on a fortios Configure Web filter banned word table.
---

# Data Source: fortios_webfilter_content
Use this data source to get information on a fortios Configure Web filter banned word table.


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
* `entries` - Configure banned word entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Block or exempt word when a match is found.
* `lang` - Language of banned word.
* `name` - Banned word.
* `pattern_type` - Banned word pattern type: wildcard pattern or Perl regular expression.
* `score` - Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
* `status` - Enable/disable banned word.
