---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_neighbor"
description: |-
  Get information on a fortios Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.
---

# Data Source: fortios_system_sdwan_neighbor
Use this data source to get information on a fortios Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

## Example Usage

WIP

## Argument Reference

* `ip` - (Required) IP/IPv6 address of neighbor.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `health_check` - SD-WAN health-check name.
* `ip` - IP/IPv6 address of neighbor.
* `member` - Member sequence number.
* `mode` - What metric to select the neighbor.
* `role` - Role of neighbor.
* `sla_id` - SLA ID.
