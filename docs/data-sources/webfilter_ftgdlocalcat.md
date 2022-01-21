---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdlocalcat"
description: |-
  Get information on a fortios Configure FortiGuard Web Filter local categories.
---

# Data Source: fortios_webfilter_ftgdlocalcat
Use this data source to get information on a fortios Configure FortiGuard Web Filter local categories.


## Example Usage

```hcl

```

## Argument Reference

* `desc` - (Required) Local category description.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `desc` - Local category description.
* `id` - Local category ID.
* `status` - Enable/disable the local category.
