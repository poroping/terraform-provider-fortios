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

```hcl

```

## Argument Reference

* `ip` - (Required) IP/IPv6 address of neighbor.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `health_check` - SD-WAN health-check name.
* `ip` - IP/IPv6 address of neighbor.
* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `mode` - What metric to select the neighbor.
* `role` - Role of neighbor.
* `sla_id` - SLA ID.
* `member` - Member sequence number list.The structure of `member` block is documented below.

The `member` block contains:

* `seq_num` - Member sequence number.
