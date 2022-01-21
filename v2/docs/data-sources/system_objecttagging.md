---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_objecttagging"
description: |-
  Get information on a fortios Configure object tagging.
---

# Data Source: fortios_system_objecttagging
Use this data source to get information on a fortios Configure object tagging.


## Example Usage

```hcl

```

## Argument Reference

* `category` - (Required) Tag Category.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `address` - Address.
* `category` - Tag Category.
* `color` - Color of icon on the GUI.
* `device` - Device.
* `interface` - Interface.
* `multiple` - Allow multiple tag selection.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
