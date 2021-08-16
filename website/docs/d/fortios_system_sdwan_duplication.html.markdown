---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_duplication"
description: |-
  Get information on a fortios Create SD-WAN duplication rule.
---

# Data Source: fortios_system_sdwan_duplication
Use this data source to get information on a fortios Create SD-WAN duplication rule.

## Example Usage

```hcl

```

## Argument Reference

* `fosid` - (Required) Duplication rule ID (1 - 255).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fosid` - Duplication rule ID (1 - 255).
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated.
* `packet_duplication` - Configure packet duplication method.
* `dstaddr` - Destination address or address group names.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address or address group name.
* `dstaddr6` - Destination address6 or address6 group names.The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address6 or address6 group name.
* `dstintf` - Outgoing (egress) interfaces or zones.The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface, zone or SDWAN zone name.
* `service` - Service and service group name.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service and service group name.
* `service_id` - SD-WAN service rule ID list.The structure of `service_id` block is documented below.

The `service_id` block contains:

* `id` - SD-WAN service rule ID.
* `srcaddr` - Source address or address group names.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address or address group name.
* `srcaddr6` - Source address6 or address6 group names.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address6 or address6 group name.
* `srcintf` - Incoming (ingress) interfaces or zones.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface, zone or SDWAN zone name.
