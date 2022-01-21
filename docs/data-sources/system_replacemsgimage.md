---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_replacemsgimage"
description: |-
  Get information on a fortios Configure replacement message images.
---

# Data Source: fortios_system_replacemsgimage
Use this data source to get information on a fortios Configure replacement message images.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Image name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `image_base64` - Image data.
* `image_type` - Image type.
* `name` - Image name.
