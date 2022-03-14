---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_offset_list"
description: |-
  Get information on a fortios Offset list.
---

# Data Source: fortios_router_rip_offsetlist
Use this data source to get information on a fortios Offset list.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Offset-list ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `access_list` - Access list name.
* `direction` - Offset list direction.
* `id` - Offset-list ID.
* `interface` - Interface name.
* `offset` - Offset.
* `status` - Status.
