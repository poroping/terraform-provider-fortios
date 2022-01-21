---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_region"
description: |-
  Get information on a fortios Configure FortiAP regions (for floor plans and maps).
---

# Data Source: fortios_wirelesscontroller_region
Use this data source to get information on a fortios Configure FortiAP regions (for floor plans and maps).


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) FortiAP region name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Comments.
* `grayscale` - Region image grayscale.
* `name` - FortiAP region name.
* `opacity` - Region image opacity (0 - 100).
