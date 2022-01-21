---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_aggregateaddress6"
description: |-
  BGP IPv6 aggregate address table.
---

## fortios_router_bgp_aggregateaddress6
BGP IPv6 aggregate address table.

~> This resource is configuring a child table of the parent resource: `fortios_router_bgp`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `as_set` - Enable/disable generate AS set path information. Valid values: `enable` `disable` .
* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_aggregateaddress6 can be imported using:
```sh
terraform import fortios_router_aggregateaddress6.labelname {{mkey}}
```
