---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_youtubechannelfilter"
description: |-
  Configure YouTube channel filter.
---

## fortios_videofilter_youtubechannelfilter
Configure YouTube channel filter.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `default_action` - YouTube channel filter default action. Valid values: `allow` `monitor` `block` .
* `id` - ID.
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `name` - Name.
* `entries` - YouTube filter entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - YouTube channel filter action. Valid values: `allow` `monitor` `block` .
* `channel_id` - Channel ID.
* `comment` - Comment.
* `id` - ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_videofilter_youtubechannelfilter can be imported using:
```sh
terraform import fortios_videofilter_youtubechannelfilter.labelname {{mkey}}
```
