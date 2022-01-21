---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionhelper"
description: |-
  Get information on a fortios Configure session helper.
---

# Data Source: fortios_system_sessionhelper
Use this data source to get information on a fortios Configure session helper.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Session helper ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Session helper ID.
* `name` - Helper name.
* `port` - Protocol port.
* `protocol` - Protocol number.
