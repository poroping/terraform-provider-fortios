---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_physicalswitch"
description: |-
  Get information on a fortios Configure physical switches.
---

# Data Source: fortios_system_physicalswitch
Use this data source to get information on a fortios Configure physical switches.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `age_enable` - Enable/disable layer 2 age timer.
* `age_val` - Layer 2 table age timer value.
* `name` - Name.
