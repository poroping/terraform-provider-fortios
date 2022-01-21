---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_distributelist"
description: |-
  Get information on a fortios Distribute list.
---

# Data Source: fortios_router_rip_distributelist
Use this data source to get information on a fortios Distribute list.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Distribute list ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `direction` - Distribute list direction.
* `id` - Distribute list ID.
* `interface` - Distribute list interface name.
* `listname` - Distribute access/prefix list name.
* `status` - status
