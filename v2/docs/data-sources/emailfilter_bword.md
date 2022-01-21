---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_bword"
description: |-
  Get information on a fortios Configure AntiSpam banned word list.
---

# Data Source: fortios_emailfilter_bword
Use this data source to get information on a fortios Configure AntiSpam banned word list.


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
* `entries` - Spam filter banned word.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Mark spam or good.
* `id` - Banned word entry ID.
* `language` - Language for the banned word.
* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression.
* `score` - Score value.
* `status` - Enable/disable status.
* `where` - Component of the email to be scanned.
