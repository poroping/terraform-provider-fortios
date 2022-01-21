---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_onetime"
description: |-
  Onetime schedule configuration.
---

## fortios_firewallschedule_onetime
Onetime schedule configuration.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `color` - Color of icon on the GUI.
* `end` - Schedule end date and time, format hh:mm yyyy/mm/dd.
* `expiration_days` - Write an event log message this many days before the schedule expires.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `name` - Onetime schedule name.
* `start` - Schedule start date and time, format hh:mm yyyy/mm/dd.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallschedule_onetime can be imported using:
```sh
terraform import fortios_firewallschedule_onetime.labelname {{mkey}}
```
