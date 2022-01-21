---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_routemap"
description: |-
  Configure route maps.
---

## fortios_router_routemap
Configure route maps.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comments` - Optional comments.
* `name` - Name.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Action. Valid values: `permit` `deny` .
* `id` - Rule ID.
* `match_as_path` - Match BGP AS path list. This attribute must reference one of the following datasources: `router.aspath-list.name` .
* `match_community` - Match BGP community list. This attribute must reference one of the following datasources: `router.community-list.name` .
* `match_community_exact` - Enable/disable exact matching of communities. Valid values: `enable` `disable` .
* `match_flags` - BGP flag value to match (0 - 65535)
* `match_interface` - Match interface configuration. This attribute must reference one of the following datasources: `system.interface.name` .
* `match_ip_address` - Match IP address permitted by access-list or prefix-list. This attribute must reference one of the following datasources: `router.access-list.name` `router.prefix-list.name` .
* `match_ip_nexthop` - Match next hop IP address passed by access-list or prefix-list. This attribute must reference one of the following datasources: `router.access-list.name` `router.prefix-list.name` .
* `match_ip6_address` - Match IPv6 address permitted by access-list6 or prefix-list6. This attribute must reference one of the following datasources: `router.access-list6.name` `router.prefix-list6.name` .
* `match_ip6_nexthop` - Match next hop IPv6 address passed by access-list6 or prefix-list6. This attribute must reference one of the following datasources: `router.access-list6.name` `router.prefix-list6.name` .
* `match_metric` - Match metric for redistribute routes.
* `match_origin` - Match BGP origin code. Valid values: `none` `egp` `igp` `incomplete` .
* `match_route_type` - Match route type. Valid values: `external-type1` `external-type2` `none` .
* `match_tag` - Match tag.
* `match_vrf` - Match VRF ID.
* `set_aggregator_as` - BGP aggregator AS.
* `set_aggregator_ip` - BGP aggregator IP.
* `set_aspath_action` - Specify preferred action of set-aspath. Valid values: `prepend` `replace` .
* `set_atomic_aggregate` - Enable/disable BGP atomic aggregate attribute. Valid values: `enable` `disable` .
* `set_community_additive` - Enable/disable adding set-community to existing community. Valid values: `enable` `disable` .
* `set_community_delete` - Delete communities matching community list. This attribute must reference one of the following datasources: `router.community-list.name` .
* `set_dampening_max_suppress` - Maximum duration to suppress a route (1 - 255 min, 0 = unset).
* `set_dampening_reachability_half_life` - Reachability half-life time for the penalty (1 - 45 min, 0 = unset).
* `set_dampening_reuse` - Value to start reusing a route (1 - 20000, 0 = unset).
* `set_dampening_suppress` - Value to start suppressing a route (1 - 20000, 0 = unset).
* `set_dampening_unreachability_half_life` - Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)
* `set_flags` - BGP flags value (0 - 65535)
* `set_ip_nexthop` - IP address of next hop.
* `set_ip6_nexthop` - IPv6 global address of next hop.
* `set_ip6_nexthop_local` - IPv6 local address of next hop.
* `set_local_preference` - BGP local preference path attribute.
* `set_metric` - Metric value.
* `set_metric_type` - Metric type. Valid values: `external-type1` `external-type2` `none` .
* `set_origin` - BGP origin code. Valid values: `none` `egp` `igp` `incomplete` .
* `set_originator_id` - BGP originator ID attribute.
* `set_route_tag` - Route tag for routing table.
* `set_tag` - Tag value.
* `set_weight` - BGP weight for routing table.
* `set_aspath` - Prepend BGP AS path attribute. The structure of `set_aspath` block is documented below.

The `set_aspath` block contains:

* `as` - AS number (0 - 4294967295). NOTE: Use quotes for repeating numbers, e.g.: "1 1 2"

* `set_community` - BGP community attribute. The structure of `set_community` block is documented below.

The `set_community` block contains:

* `community` - Attribute: AA|AA:NN|internet|local-AS|no-advertise|no-export.
* `set_extcommunity_rt` - Route Target extended community. The structure of `set_extcommunity_rt` block is documented below.

The `set_extcommunity_rt` block contains:

* `community` - AA:NN.
* `set_extcommunity_soo` - Site-of-Origin extended community. The structure of `set_extcommunity_soo` block is documented below.

The `set_extcommunity_soo` block contains:

* `community` - AA:NN

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_routemap can be imported using:
```sh
terraform import fortios_router_routemap.labelname {{mkey}}
```
