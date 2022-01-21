---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdlocalrating"
description: |-
  Configure local FortiGuard Web Filter local ratings.
---

## fortios_webfilter_ftgdlocalrating
Configure local FortiGuard Web Filter local ratings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `url` to be defined.

* `comment` - Comment.
* `rating` - Local rating.
* `status` - Enable/disable local rating. Valid values: `enable` `disable` .
* `url` - URL to rate locally.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_ftgdlocalrating can be imported using:
```sh
terraform import fortios_webfilter_ftgdlocalrating.labelname {{mkey}}
```
