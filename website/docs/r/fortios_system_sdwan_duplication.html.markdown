---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_duplication"
description: |-
  Create SD-WAN duplication rule.
---

## fortios_system_sdwan_duplication
Create SD-WAN duplication rule.

~> This resource is configuring a child table of the parent resource: `fortios_system_sdwan`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.

## Example Usage

WIP

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `fosid` - Duplication rule ID (1 - 255).
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated. Valid values: `enable` `disable` .
* `packet_duplication` - Configure packet duplication method. Valid values: `disable` `force` `on-demand` .
* `dstaddr` - Destination address or address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address or address group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dstaddr6` - Destination address6 or address6 group names. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address6 or address6 group name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `dstintf` - Outgoing (egress) interfaces or zones. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface, zone or SDWAN zone name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `service` - Service and service group name. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service and service group name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `service_id` - SD-WAN service rule ID list. The structure of `service_id` block is documented below.

The `service_id` block contains:

* `id` - SD-WAN service rule ID. This attribute must reference one of the following datasources: `system.sdwan.service.id` .
* `srcaddr` - Source address or address group names. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address or address group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `srcaddr6` - Source address6 or address6 group names. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address6 or address6 group name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `srcintf` - Incoming (ingress) interfaces or zones. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface, zone or SDWAN zone name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

system_duplication can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_duplication.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
