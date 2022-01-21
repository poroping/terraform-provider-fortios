---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_youtubekey"
description: |-
  Configure YouTube API keys.
---

## fortios_videofilter_youtubekey
Configure YouTube API keys.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - ID.
* `key` - Key.
* `status` - Enable/disable YouTube API service. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_videofilter_youtubekey can be imported using:
```sh
terraform import fortios_videofilter_youtubekey.labelname {{mkey}}
```
