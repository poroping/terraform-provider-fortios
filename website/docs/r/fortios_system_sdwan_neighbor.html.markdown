---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_neighbor"
description: |-
  Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.
---

## fortios_system_sdwan_neighbor
Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

~> This resource is configuring a child table of the parent resource: `fortios_system_sdwan`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.

## Example Usage

WIP

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `health_check` - SD-WAN health-check name. This attribute must reference one of the following datasources: `system.sdwan.health-check.name` .
* `ip` - IP/IPv6 address of neighbor. This attribute must reference one of the following datasources: `router.bgp.neighbor.ip` .
* `member` - Member sequence number. This attribute must reference one of the following datasources: `system.sdwan.members.seq-num` .
* `mode` - What metric to select the neighbor. Valid values: `sla` `speedtest` .
* `role` - Role of neighbor. Valid values: `standalone` `primary` `secondary` .
* `sla_id` - SLA ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

system_neighbor can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_neighbor.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
