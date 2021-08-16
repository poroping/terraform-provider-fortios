---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp"
description: |-
  Configure BGP.
---

## fortios_router_bgp
Configure BGP.
## Example Usage

```hcl
resource "fortios_router_bgp" "trname" {
  additional_path_select             = 2
  additional_path_select6            = 2
  always_compare_med                 = "disable"
  as                                 = 0
  client_to_client_reflection        = "enable"
  cluster_id                         = "0.0.0.0"
  dampening                          = "disable"
  dampening_max_suppress_time        = 60
  dampening_reachability_half_life   = 15
  dampening_reuse                    = 750
  dampening_suppress                 = 2000
  dampening_unreachability_half_life = 15
  default_local_preference           = 100
  deterministic_med                  = "disable"
  distance_external                  = 20
  distance_internal                  = 200
  distance_local                     = 200
  graceful_restart_time              = 120
  graceful_stalepath_time            = 360
  graceful_update_delay              = 120
  holdtime_timer                     = 180
  ibgp_multipath                     = "disable"
  ignore_optional_capability         = "enable"
  keepalive_timer                    = 60
  log_neighbour_changes              = "enable"
  network_import_check               = "enable"
  scan_time                          = 60
  synchronization                    = "disable"

  redistribute {
    name   = "connected"
    status = "disable"
  }
  redistribute {
    name   = "rip"
    status = "disable"
  }
  redistribute {
    name   = "ospf"
    status = "disable"
  }
  redistribute {
    name   = "static"
    status = "disable"
  }
  redistribute {
    name   = "isis"
    status = "disable"
  }

  redistribute6 {
    name   = "connected"
    status = "disable"
  }
  redistribute6 {
    name   = "rip"
    status = "disable"
  }
  redistribute6 {
    name   = "ospf"
    status = "disable"
  }
  redistribute6 {
    name   = "static"
    status = "disable"
  }
  redistribute6 {
    name   = "isis"
    status = "disable"
  }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `additional_path` - Enable/disable selection of BGP IPv4 additional paths. Valid values: `enable` `disable` .
* `additional_path_select` - Number of additional paths to be selected for each IPv4 NLRI.
* `additional_path_select6` - Number of additional paths to be selected for each IPv6 NLRI.
* `additional_path6` - Enable/disable selection of BGP IPv6 additional paths. Valid values: `enable` `disable` .
* `always_compare_med` - Enable/disable always compare MED. Valid values: `enable` `disable` .
* `as` - Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path. Valid values: `enable` `disable` .
* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length. Valid values: `enable` `disable` .
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths. Valid values: `enable` `disable` .
* `bestpath_med_confed` - Enable/disable compare MED among confederation paths. Valid values: `enable` `disable` .
* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred. Valid values: `enable` `disable` .
* `client_to_client_reflection` - Enable/disable client-to-client route reflection. Valid values: `enable` `disable` .
* `cluster_id` - Route reflector cluster ID.
* `confederation_identifier` - Confederation identifier.
* `dampening` - Enable/disable route-flap dampening. Valid values: `enable` `disable` .
* `dampening_max_suppress_time` - Maximum minutes a route can be suppressed.
* `dampening_reachability_half_life` - Reachability half-life time for penalty (min).
* `dampening_reuse` - Threshold to reuse routes.
* `dampening_route_map` - Criteria for dampening. This attribute must reference one of the following datasources: `router.route-map.name` .
* `dampening_suppress` - Threshold to suppress routes.
* `dampening_unreachability_half_life` - Unreachability half-life time for penalty (min).
* `default_local_preference` - Default local preference.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED. Valid values: `enable` `disable` .
* `distance_external` - Distance for routes external to the AS.
* `distance_internal` - Distance for routes internal to the AS.
* `distance_local` - Distance for routes local to the AS.
* `ebgp_multipath` - Enable/disable EBGP multi-path. Valid values: `enable` `disable` .
* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes. Valid values: `enable` `disable` .
* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down. Valid values: `enable` `disable` .
* `graceful_end_on_timer` - Enable/disable to exit graceful restart on timer only. Valid values: `enable` `disable` .
* `graceful_restart` - Enable/disable BGP graceful restart capabilities. Valid values: `enable` `disable` .
* `graceful_restart_time` - Time needed for neighbors to restart (sec).
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbor (sec).
* `graceful_update_delay` - Route advertisement/selection delay after restart (sec).
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `ibgp_multipath` - Enable/disable IBGP multi-path. Valid values: `enable` `disable` .
* `ignore_optional_capability` - Don't send unknown optional capability notification message Valid values: `enable` `disable` .
* `keepalive_timer` - Frequency to send keep alive requests.
* `log_neighbour_changes` - Enable logging of BGP neighbour's changes Valid values: `enable` `disable` .
* `multipath_recursive_distance` - Enable/disable use of recursive distance to select multipath. Valid values: `enable` `disable` .
* `network_import_check` - Enable/disable ensure BGP network route exists in IGP. Valid values: `enable` `disable` .
* `recursive_next_hop` - Enable/disable recursive resolution of next-hop using BGP route. Valid values: `enable` `disable` .
* `router_id` - Router ID.
* `scan_time` - Background scanner interval (sec), 0 to disable it.
* `synchronization` - Enable/disable only advertise routes from iBGP if routes present in an IGP. Valid values: `enable` `disable` .
* `admin_distance` - Administrative distance modifications. The structure of `admin_distance` block is documented below.

The `admin_distance` block contains:

* `distance` - Administrative distance to apply (1 - 255).
* `id` - ID.
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to. This attribute must reference one of the following datasources: `router.access-list.name` .
* `aggregate_address` - BGP aggregate address table. The structure of `aggregate_address` block is documented below.

The `aggregate_address` block contains:

* `as_set` - Enable/disable generate AS set path information. Valid values: `enable` `disable` .
* `id` - ID.
* `prefix` - Aggregate prefix.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable` `disable` .
* `aggregate_address6` - BGP IPv6 aggregate address table. The structure of `aggregate_address6` block is documented below.

The `aggregate_address6` block contains:

* `as_set` - Enable/disable generate AS set path information. Valid values: `enable` `disable` .
* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `summary_only` - Enable/disable filter more specific routes from updates. Valid values: `enable` `disable` .
* `confederation_peers` - Confederation peers. The structure of `confederation_peers` block is documented below.

The `confederation_peers` block contains:

* `peer` - Peer ID.
* `neighbor` - BGP neighbor table. The structure of `neighbor` block is documented below.

The `neighbor` block contains:

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
* `ip` - IP/IPv6 address of neighbor.
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
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable` `disable` .
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `enable` `disable` .
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `enable` `disable` .
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable` `disable` .
* `override_capability` - Enable/disable override result of capability negotiation. Valid values: `enable` `disable` .
* `passive` - Enable/disable sending of open messages to this neighbor. Valid values: `enable` `disable` .
* `password` - Password used in MD5 authentication.
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
* `conditional_advertise` - Conditional advertisement. The structure of `conditional_advertise` block is documented below.

The `conditional_advertise` block contains:

* `advertise_routemap` - Name of advertising route map. This attribute must reference one of the following datasources: `router.route-map.name` .
* `condition_routemap` - Name of condition route map. This attribute must reference one of the following datasources: `router.route-map.name` .
* `condition_type` - Type of condition. Valid values: `exist` `non-exist` .
* `conditional_advertise6` - IPv6 conditional advertisement. The structure of `conditional_advertise6` block is documented below.

The `conditional_advertise6` block contains:

* `advertise_routemap` - Name of advertising route map. This attribute must reference one of the following datasources: `router.route-map.name` .
* `condition_routemap` - Name of condition route map. This attribute must reference one of the following datasources: `router.route-map.name` .
* `condition_type` - Type of condition. Valid values: `exist` `non-exist` .
* `neighbor_group` - BGP neighbor group table. The structure of `neighbor_group` block is documented below.

The `neighbor_group` block contains:

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
* `neighbor_range` - BGP neighbor range table. The structure of `neighbor_range` block is documented below.

The `neighbor_range` block contains:

* `id` - Neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name. This attribute must reference one of the following datasources: `router.bgp.neighbor-group.name` .
* `prefix` - Neighbor range prefix.
* `neighbor_range6` - BGP IPv6 neighbor range table. The structure of `neighbor_range6` block is documented below.

The `neighbor_range6` block contains:

* `id` - IPv6 neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name. This attribute must reference one of the following datasources: `router.bgp.neighbor-group.name` .
* `prefix6` - IPv6 prefix.
* `network` - BGP network table. The structure of `network` block is documented below.

The `network` block contains:

* `backdoor` - Enable/disable route as backdoor. Valid values: `enable` `disable` .
* `id` - ID.
* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route. This attribute must reference one of the following datasources: `router.route-map.name` .
* `network6` - BGP IPv6 network table. The structure of `network6` block is documented below.

The `network6` block contains:

* `backdoor` - Enable/disable route as backdoor. Valid values: `enable` `disable` .
* `id` - ID.
* `prefix6` - Network IPv6 prefix.
* `route_map` - Route map to modify generated route. This attribute must reference one of the following datasources: `router.route-map.name` .
* `redistribute` - BGP IPv4 redistribute table. The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `name` - Distribute list entry name.
* `route_map` - Route map name. This attribute must reference one of the following datasources: `router.route-map.name` .
* `status` - Status Valid values: `enable` `disable` .
* `redistribute6` - BGP IPv6 redistribute table. The structure of `redistribute6` block is documented below.

The `redistribute6` block contains:

* `name` - Distribute list entry name.
* `route_map` - Route map name. This attribute must reference one of the following datasources: `router.route-map.name` .
* `status` - Status Valid values: `enable` `disable` .
* `vrf_leak` - BGP VRF leaking table. The structure of `vrf_leak` block is documented below.

The `vrf_leak` block contains:

* `vrf` - Origin VRF ID <0 - 31>.
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block contains:

* `interface` - Interface which is used to leak routes to target VRF. This attribute must reference one of the following datasources: `system.interface.name` .
* `route_map` - Route map of VRF leaking. This attribute must reference one of the following datasources: `router.route-map.name` .
* `vrf` - Target VRF ID <0 - 31>.
* `vrf_leak6` - BGP IPv6 VRF leaking table. The structure of `vrf_leak6` block is documented below.

The `vrf_leak6` block contains:

* `vrf` - Origin VRF ID <0 - 31>.
* `target` - Target VRF table. The structure of `target` block is documented below.

The `target` block contains:

* `interface` - Interface which is used to leak routes to target VRF. This attribute must reference one of the following datasources: `system.interface.name` .
* `route_map` - Route map of VRF leaking. This attribute must reference one of the following datasources: `router.route-map.name` .
* `vrf` - Target VRF ID <0 - 31>.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

router_bgp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_bgp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
