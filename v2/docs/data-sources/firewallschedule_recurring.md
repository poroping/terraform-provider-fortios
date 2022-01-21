---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_recurring"
description: |-
  Get information on a fortios Recurring schedule configuration.
---

# Data Source: fortios_firewallschedule_recurring
Use this data source to get information on a fortios Recurring schedule configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Recurring schedule name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Color of icon on the GUI.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space.
* `end` - Time of day to end the schedule, format hh:mm.
* `fabric_object` - Security Fabric global object setting.
* `name` - Recurring schedule name.
* `start` - Time of day to start the schedule, format hh:mm.
