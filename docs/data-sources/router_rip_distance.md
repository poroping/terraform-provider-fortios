---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_distance"
description: |-
  Get information on a fortios distance
---

# Data Source: fortios_router_rip_distance
Use this data source to get information on a fortios distance


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Distance ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_list` - Access list for route destination.
* `distance` - Distance (1 - 255).
* `id` - Distance ID.
* `prefix` - Distance prefix.
