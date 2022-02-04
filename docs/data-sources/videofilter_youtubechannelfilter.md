---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_youtubechannelfilter"
description: |-
  Get information on a fortios Configure YouTube channel filter.
---

# Data Source: fortios_videofilter_youtubechannelfilter
Use this data source to get information on a fortios Configure YouTube channel filter.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `default_action` - YouTube channel filter default action.
* `id` - ID.
* `log` - Enable/disable logging.
* `name` - Name.
* `entries` - YouTube filter entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - YouTube channel filter action.
* `channel_id` - Channel ID.
* `comment` - Comment.
* `id` - ID.
