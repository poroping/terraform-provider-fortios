---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_admindistance"
description: |-
  Get information on a fortios Administrative distance modifications.
---

# Data Source: fortios_router_bgp_admindistance
Use this data source to get information on a fortios Administrative distance modifications.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `distance` - Administrative distance to apply (1 - 255).
* `id` - ID.
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.
