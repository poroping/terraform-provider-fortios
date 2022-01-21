---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_region"
description: |-
  Configure FortiAP regions (for floor plans and maps).
---

## fortios_wirelesscontroller_region
Configure FortiAP regions (for floor plans and maps).

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comments` - Comments.
* `grayscale` - Region image grayscale. Valid values: `enable` `disable` .
* `name` - FortiAP region name.
* `opacity` - Region image opacity (0 - 100).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_region can be imported using:
```sh
terraform import fortios_wirelesscontroller_region.labelname {{mkey}}
```
