---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd_neighbor"
description: |-
  Get information on a fortios Neighbor.
---

# Data Source: fortios_router_bfd_neighbor
Use this data source to get information on a fortios Neighbor.


## Example Usage

```hcl

```

## Argument Reference

* `ip` - (Required) IPv4 address of the BFD neighbor.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `interface` - Interface name.
* `ip` - IPv4 address of the BFD neighbor.
