---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_zone"
description: |-
  Get information on a fortios Configure SD-WAN zones.
---

# Data Source: fortios_system_sdwan_zone
Use this data source to get information on a fortios Configure SD-WAN zones.

## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Zone name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA.
