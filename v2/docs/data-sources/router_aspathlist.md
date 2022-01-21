---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_aspathlist"
description: |-
  Get information on a fortios Configure Autonomous System (AS) path lists.
---

# Data Source: fortios_router_aspathlist
Use this data source to get information on a fortios Configure Autonomous System (AS) path lists.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) AS path list name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - AS path list name.
* `rule` - AS path list rule.The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Permit or deny route-based operations, based on the route's AS_PATH attribute.
* `id` - ID.
* `regexp` - Regular-expression to match the Border Gateway Protocol (BGP) AS paths.
