---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_neighborgroup"
description: |-
  BGP neighbor group table.
---

## fortios_router_bgp_neighborgroup
BGP neighbor group table.

~> This resource is configuring a child table of the parent resource: `fortios_router_bgp`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `activate` - Enable/disable address family IPv4 for this neighbor. Valid values: `enable` `disable` .
* `activate6` - Enable/disable address family IPv6 for this neighbor. Valid values: `enable` `disable` .
* `additional_path` - Enable/disable IPv4 additional-path capability. Valid values: `send` `receive` `both` `disable` .
* `additional_path6` - Enable/disable IPv6 additional-path capability. Valid values: `send` `receive` `both` `disable` .
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable` `disable` .
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable` `disable` .
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable` `disable` .
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable` `disable` .
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged. Valid values: `as-path` `med` `next-hop` .
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged. Valid values: `as-path` `med` `next-hop` .
* `bfd` - Enable/disable BFD for this neighbor. Valid values: `enable` `disable` .
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor. Valid values: `enable` `disable` .
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable` `disable` .
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable` `disable` .
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `enable` `disable` .
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `enable` `disable` .
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none` `receive` `send` `both` .
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none` `receive` `send` `both` .
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor. Valid values: `enable` `disable` .
* `connect_timer` - Interval (sec) for connect timer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default. This attribute must reference one of the following datasources: `router.route-map.name` .
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default. This attribute must reference one of the following datasources: `router.route-map.name` .
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor. This attribute must reference one of the following datasources: `router.access-list.name` .
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor. This attribute must reference one of the following datasources: `router.access-list6.name` .
* `distribute_list_out` - Filter for IPv4 updates to this neighbor. This attribute must reference one of the following datasources: `router.access-list.name` .
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor. This attribute must reference one of the following datasources: `router.access-list6.name` .
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor Valid values: `enable` `disable` .
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors. Valid values: `enable` `disable` .
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes. This attribute must reference one of the following datasources: `router.aspath-list.name` .
* `filter_list_in6` - BGP filter for IPv6 inbound routes. This attribute must reference one of the following datasources: `router.aspath-list.name` .
* `filter_list_out` - BGP filter for IPv4 outbound routes. This attribute must reference one of the following datasources: `router.aspath-list.name` .
* `filter_list_out6` - BGP filter for IPv6 outbound routes. This attribute must reference one of the following datasources: `router.aspath-list.name` .
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `interface` - Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address. This attribute must reference one of the following datasources: `system.interface.name` .
* `keep_alive_timer` - Keep alive timer interval (sec).
* `link_down_failover` - Enable/disable failover upon link down. Valid values: `enable` `disable` .
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates. Valid values: `enable` `disable` .
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates. Valid values: `enable` `disable` .
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `enable` `disable` .
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `enable` `disable` .
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `name` - Neighbor group name.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable` `disable` .
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `enable` `disable` .
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `enable` `disable` .
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable` `disable` .
* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `enable` `disable` .
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable` `disable` .
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor. This attribute must reference one of the following datasources: `router.prefix-list.name` .
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor. This attribute must reference one of the following datasources: `router.prefix-list6.name` .
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor. This attribute must reference one of the following datasources: `router.prefix-list.name` .
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor. This attribute must reference one of the following datasources: `router.prefix-list6.name` .
* `remote_as` - AS number of neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable` `disable` .
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable` `disable` .
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter. This attribute must reference one of the following datasources: `router.route-map.name` .
* `route_map_in6` - IPv6 Inbound route map filter. This attribute must reference one of the following datasources: `router.route-map.name` .
* `route_map_out` - IPv4 outbound route map filter. This attribute must reference one of the following datasources: `router.route-map.name` .
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred. This attribute must reference one of the following datasources: `router.route-map.name` .
* `route_map_out6` - IPv6 Outbound route map filter. This attribute must reference one of the following datasources: `router.route-map.name` .
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred. This attribute must reference one of the following datasources: `router.route-map.name` .
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client. Valid values: `enable` `disable` .
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client. Valid values: `enable` `disable` .
* `route_server_client` - Enable/disable IPv4 AS route server client. Valid values: `enable` `disable` .
* `route_server_client6` - Enable/disable IPv6 AS route server client. Valid values: `enable` `disable` .
* `send_community` - IPv4 Send community attribute to neighbor. Valid values: `standard` `extended` `both` `disable` .
* `send_community6` - IPv6 Send community attribute to neighbor. Valid values: `standard` `extended` `both` `disable` .
* `shutdown` - Enable/disable shutdown this neighbor. Valid values: `enable` `disable` .
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable` `disable` .
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable` `disable` .
* `stale_route` - Enable/disable stale route after neighbor down. Valid values: `enable` `disable` .
* `strict_capability_match` - Enable/disable strict capability matching. Valid values: `enable` `disable` .
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes. This attribute must reference one of the following datasources: `router.route-map.name` .
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes. This attribute must reference one of the following datasources: `router.route-map.name` .
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections. This attribute must reference one of the following datasources: `system.interface.name` .
* `weight` - Neighbor weight.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_neighborgroup can be imported using:
```sh
terraform import fortios_router_neighborgroup.labelname {{mkey}}
```
