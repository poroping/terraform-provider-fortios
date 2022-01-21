---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallschedule_group"
description: |-
  Get information on a fortios Schedule group configuration.
---

# Data Source: fortios_firewallschedule_group
Use this data source to get information on a fortios Schedule group configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Schedule group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Color of icon on the GUI.
* `fabric_object` - Security Fabric global object setting.
* `name` - Schedule group name.
* `member` - Schedules added to the schedule group.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Schedule name.
