---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_schedule"
description: |-
  Get information on a fortios Configure update schedule.
---

# Data Source: fortios_systemautoupdate_schedule
Use this data source to get information on a fortios Configure update schedule.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `day` - Update day.
* `frequency` - Update frequency.
* `status` - Enable/disable scheduled updates.
* `time` - Update time.
