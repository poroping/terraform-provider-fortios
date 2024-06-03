---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_profile"
description: |-
  Configure VideoFilter profile.
---

## fortios_videofilter_profile
Configure VideoFilter profile.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comment` - Comment.
* `dailymotion` - Enable/disable Dailymotion video source. Valid values: `enable` `disable` .
* `default_action` - Video filter default action. Valid values: `allow` `monitor` `block` .
* `log` - Enable/disable logging. Valid values: `enable` `disable` .
* `name` - Name.
* `replacemsg_group` - Replacement message group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `vimeo` - Enable/disable Vimeo video source. Valid values: `enable` `disable` .
* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `youtube` - Enable/disable YouTube video source. Valid values: `enable` `disable` .
* `youtube_channel_filter` - Set YouTube channel filter. This attribute must reference one of the following datasources: `videofilter.youtube-channel-filter.id` .
* `youtube_restrict` - Set YouTube-restrict mode. Valid values: `none` `strict` `moderate` .
* `fortiguard_category` - Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.

The `fortiguard_category` block contains:

* `filters` - Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.

The `filters` block contains:

* `action` - VideoFilter action. Valid values: `allow` `monitor` `block` .
* `category_id` - Category ID.
* `id` - ID.
* `log` - Enable/disable logging. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_videofilter_profile can be imported using:
```sh
terraform import fortios_videofilter_profile.labelname {{mkey}}
```
