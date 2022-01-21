---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_communitylist"
description: |-
  Get information on a fortios Configure community lists.
---

# Data Source: fortios_router_communitylist
Use this data source to get information on a fortios Configure community lists.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Community list name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Community list name.
* `type` - Community list type (standard or expanded).
* `rule` - Community list rule.The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Permit or deny route-based operations, based on the route's COMMUNITY attribute.
* `id` - ID.
* `match` - Community specifications for matching a reserved community.
* `regexp` - Ordered list of COMMUNITY attributes as a regular expression.
