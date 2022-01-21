---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_neighbor"
description: |-
  Get information on a fortios neighbor
---

# Data Source: fortios_router_rip_neighbor
Use this data source to get information on a fortios neighbor


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Neighbor entry ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Neighbor entry ID.
* `ip` - IP address.
