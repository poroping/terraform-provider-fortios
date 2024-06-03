---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_onetime"
description: |-
  Get information on a fortios Onetime schedule configuration.
---

# Data Source: fortios_firewallschedule_onetime
Use this data source to get information on a fortios Onetime schedule configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Onetime schedule name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Color of icon on the GUI.
* `end` - Schedule end date and time, format hh:mm yyyy/mm/dd.
* `end_utc` - Schedule end date and time, in epoch format.
* `expiration_days` - Write an event log message this many days before the schedule expires.
* `fabric_object` - Security Fabric global object setting.
* `name` - Onetime schedule name.
* `start` - Schedule start date and time, format hh:mm yyyy/mm/dd.
* `start_utc` - Schedule start date and time, in epoch format.
