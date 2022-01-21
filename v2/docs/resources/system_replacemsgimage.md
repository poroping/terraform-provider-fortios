---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_replacemsgimage"
description: |-
  Configure replacement message images.
---

## fortios_system_replacemsgimage
Configure replacement message images.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `image_base64` - Image data.
* `image_type` - Image type. Valid values: `gif` `jpg` `tiff` `png` .
* `name` - Image name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_replacemsgimage can be imported using:
```sh
terraform import fortios_system_replacemsgimage.labelname {{mkey}}
```
