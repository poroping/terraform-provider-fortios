---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

## fortios_emailfilter_bword
Configure AntiSpam banned word list.

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
* `entries` - Spam filter banned word. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Mark spam or good. Valid values: `spam` `clear` .
* `id` - Banned word entry ID.
* `language` - Language for the banned word. Valid values: `western` `simch` `trach` `japanese` `korean` `french` `thai` `spanish` .
* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard` `regexp` .
* `score` - Score value.
* `status` - Enable/disable status. Valid values: `enable` `disable` .
* `where` - Component of the email to be scanned. Valid values: `subject` `body` `all` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_emailfilter_bword can be imported using:
```sh
terraform import fortios_emailfilter_bword.labelname {{mkey}}
```
