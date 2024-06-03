---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_profile"
description: |-
  Get information on a fortios Configure VideoFilter profile.
---

# Data Source: fortios_videofilter_profile
Use this data source to get information on a fortios Configure VideoFilter profile.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `dailymotion` - Enable/disable Dailymotion video source.
* `default_action` - Video filter default action.
* `log` - Enable/disable logging.
* `name` - Name.
* `replacemsg_group` - Replacement message group.
* `vimeo` - Enable/disable Vimeo video source.
* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `youtube` - Enable/disable YouTube video source.
* `youtube_channel_filter` - Set YouTube channel filter.
* `youtube_restrict` - Set YouTube-restrict mode.
* `fortiguard_category` - Configure FortiGuard categories.The structure of `fortiguard_category` block is documented below.

The `fortiguard_category` block contains:

* `filters` - Configure VideoFilter FortiGuard category.The structure of `filters` block is documented below.

The `filters` block contains:

* `action` - VideoFilter action.
* `category_id` - Category ID.
* `id` - ID.
* `log` - Enable/disable logging.
