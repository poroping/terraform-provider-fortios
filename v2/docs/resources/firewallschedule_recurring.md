---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_recurring"
description: |-
  Recurring schedule configuration.
---

## fortios_firewallschedule_recurring
Recurring schedule configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `color` - Color of icon on the GUI.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday` `monday` `tuesday` `wednesday` `thursday` `friday` `saturday` `none` .
* `end` - Time of day to end the schedule, format hh:mm.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `name` - Recurring schedule name.
* `start` - Time of day to start the schedule, format hh:mm.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallschedule_recurring can be imported using:
```sh
terraform import fortios_firewallschedule_recurring.labelname {{mkey}}
```
