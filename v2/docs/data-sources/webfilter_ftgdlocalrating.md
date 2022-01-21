---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdlocalrating"
description: |-
  Get information on a fortios Configure local FortiGuard Web Filter local ratings.
---

# Data Source: fortios_webfilter_ftgdlocalrating
Use this data source to get information on a fortios Configure local FortiGuard Web Filter local ratings.


## Example Usage

```hcl

```

## Argument Reference

* `url` - (Required) URL to rate locally.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `rating` - Local rating.
* `status` - Enable/disable local rating.
* `url` - URL to rate locally.
