---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_setting"
description: |-
  Get information on a fortios Settings for memory buffer.
---

# Data Source: fortios_logmemory_setting
Use this data source to get information on a fortios Settings for memory buffer.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `diskfull` - Action to take when memory is full.
* `status` - Enable/disable logging to the FortiGate's memory.
