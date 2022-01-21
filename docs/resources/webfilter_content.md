---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_content"
description: |-
  Configure Web filter banned word table.
---

## fortios_webfilter_content
Configure Web filter banned word table.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - Configure banned word entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Block or exempt word when a match is found. Valid values: `block` `exempt` .
* `lang` - Language of banned word. Valid values: `western` `simch` `trach` `japanese` `korean` `french` `thai` `spanish` `cyrillic` .
* `name` - Banned word.
* `pattern_type` - Banned word pattern type: wildcard pattern or Perl regular expression. Valid values: `wildcard` `regexp` .
* `score` - Score, to be applied every time the word appears on a web page (0 - 4294967295, default = 10).
* `status` - Enable/disable banned word. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_content can be imported using:
```sh
terraform import fortios_webfilter_content.labelname {{mkey}}
```
