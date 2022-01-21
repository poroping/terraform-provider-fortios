---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf_neighbor"
description: |-
  Get information on a fortios OSPF neighbor configuration are used when OSPF runs on non-broadcast media
---

# Data Source: fortios_router_ospf_neighbor
Use this data source to get information on a fortios OSPF neighbor configuration are used when OSPF runs on non-broadcast media


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Neighbor entry ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.
