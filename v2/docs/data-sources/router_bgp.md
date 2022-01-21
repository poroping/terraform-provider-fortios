---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp"
description: |-
  Get information on a fortios Configure BGP.
---

# Data Source: fortios_router_bgp
Use this data source to get information on a fortios Configure BGP.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `additional_path` - Enable/disable selection of BGP IPv4 additional paths.
* `additional_path_select` - Number of additional paths to be selected for each IPv4 NLRI.
* `additional_path_select6` - Number of additional paths to be selected for each IPv6 NLRI.
* `additional_path6` - Enable/disable selection of BGP IPv6 additional paths.
* `always_compare_med` - Enable/disable always compare MED.
* `as` - Router AS number, valid from 1 to 4294967295, 0 to disable BGP.
* `bestpath_as_path_ignore` - Enable/disable ignore AS path.
* `bestpath_cmp_confed_aspath` - Enable/disable compare federation AS path length.
* `bestpath_cmp_routerid` - Enable/disable compare router ID for identical EBGP paths.
* `bestpath_med_confed` - Enable/disable compare MED among confederation paths.
* `bestpath_med_missing_as_worst` - Enable/disable treat missing MED as least preferred.
* `client_to_client_reflection` - Enable/disable client-to-client route reflection.
* `cluster_id` - Route reflector cluster ID.
* `confederation_identifier` - Confederation identifier.
* `dampening` - Enable/disable route-flap dampening.
* `dampening_max_suppress_time` - Maximum minutes a route can be suppressed.
* `dampening_reachability_half_life` - Reachability half-life time for penalty (min).
* `dampening_reuse` - Threshold to reuse routes.
* `dampening_route_map` - Criteria for dampening.
* `dampening_suppress` - Threshold to suppress routes.
* `dampening_unreachability_half_life` - Unreachability half-life time for penalty (min).
* `default_local_preference` - Default local preference.
* `deterministic_med` - Enable/disable enforce deterministic comparison of MED.
* `distance_external` - Distance for routes external to the AS.
* `distance_internal` - Distance for routes internal to the AS.
* `distance_local` - Distance for routes local to the AS.
* `ebgp_multipath` - Enable/disable EBGP multi-path.
* `enforce_first_as` - Enable/disable enforce first AS for EBGP routes.
* `fast_external_failover` - Enable/disable reset peer BGP session if link goes down.
* `graceful_end_on_timer` - Enable/disable to exit graceful restart on timer only.
* `graceful_restart` - Enable/disable BGP graceful restart capabilities.
* `graceful_restart_time` - Time needed for neighbors to restart (sec).
* `graceful_stalepath_time` - Time to hold stale paths of restarting neighbor (sec).
* `graceful_update_delay` - Route advertisement/selection delay after restart (sec).
* `holdtime_timer` - Number of seconds to mark peer as dead.
* `ibgp_multipath` - Enable/disable IBGP multi-path.
* `ignore_optional_capability` - Don't send unknown optional capability notification message
* `keepalive_timer` - Frequency to send keep alive requests.
* `log_neighbour_changes` - Enable logging of BGP neighbour's changes
* `multipath_recursive_distance` - Enable/disable use of recursive distance to select multipath.
* `network_import_check` - Enable/disable ensure BGP network route exists in IGP.
* `recursive_next_hop` - Enable/disable recursive resolution of next-hop using BGP route.
* `router_id` - Router ID.
* `scan_time` - Background scanner interval (sec), 0 to disable it.
* `synchronization` - Enable/disable only advertise routes from iBGP if routes present in an IGP.
* `admin_distance` - Administrative distance modifications.The structure of `admin_distance` block is documented below.

The `admin_distance` block contains:

* `distance` - Administrative distance to apply (1 - 255).
* `id` - ID.
* `neighbour_prefix` - Neighbor address prefix.
* `route_list` - Access list of routes to apply new distance to.
* `aggregate_address` - BGP aggregate address table.The structure of `aggregate_address` block is documented below.

The `aggregate_address` block contains:

* `as_set` - Enable/disable generate AS set path information.
* `id` - ID.
* `prefix` - Aggregate prefix.
* `summary_only` - Enable/disable filter more specific routes from updates.
* `aggregate_address6` - BGP IPv6 aggregate address table.The structure of `aggregate_address6` block is documented below.

The `aggregate_address6` block contains:

* `as_set` - Enable/disable generate AS set path information.
* `id` - ID.
* `prefix6` - Aggregate IPv6 prefix.
* `summary_only` - Enable/disable filter more specific routes from updates.
* `confederation_peers` - Confederation peers.The structure of `confederation_peers` block is documented below.

The `confederation_peers` block contains:

* `peer` - Peer ID.
* `neighbor` - BGP neighbor table.The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `additional_path` - Enable/disable IPv4 additional-path capability.
* `additional_path6` - Enable/disable IPv6 additional-path capability.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor.
* `connect_timer` - Interval (sec) for connect timer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `interface` - Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.
* `ip` - IP/IPv6 address of neighbor.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `link_down_failover` - Enable/disable failover upon link down.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates.
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `password` - Password used in MD5 authentication.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `stale_route` - Enable/disable stale route after neighbor down.
* `strict_capability_match` - Enable/disable strict capability matching.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.
* `conditional_advertise` - Conditional advertisement.The structure of `conditional_advertise` block is documented below.

The `conditional_advertise` block contains:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition.
* `conditional_advertise6` - IPv6 conditional advertisement.The structure of `conditional_advertise6` block is documented below.

The `conditional_advertise6` block contains:

* `advertise_routemap` - Name of advertising route map.
* `condition_routemap` - Name of condition route map.
* `condition_type` - Type of condition.
* `neighbor_group` - BGP neighbor group table.The structure of `neighbor_group` block is documented below.

The `neighbor_group` block contains:

* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `additional_path` - Enable/disable IPv4 additional-path capability.
* `additional_path6` - Enable/disable IPv6 additional-path capability.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor.
* `connect_timer` - Interval (sec) for connect timer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `dont_capability_negotiate` - Don't negotiate capabilities with this neighbor
* `ebgp_enforce_multihop` - Enable/disable allow multi-hop EBGP neighbors.
* `ebgp_multihop_ttl` - EBGP multihop TTL for this peer.
* `filter_list_in` - BGP filter for IPv4 inbound routes.
* `filter_list_in6` - BGP filter for IPv6 inbound routes.
* `filter_list_out` - BGP filter for IPv4 outbound routes.
* `filter_list_out6` - BGP filter for IPv6 outbound routes.
* `holdtime_timer` - Interval (sec) before peer considered dead.
* `interface` - Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.
* `keep_alive_timer` - Keep alive timer interval (sec).
* `link_down_failover` - Enable/disable failover upon link down.
* `local_as` - Local AS number of neighbor.
* `local_as_no_prepend` - Do not prepend local-as to incoming updates.
* `local_as_replace_as` - Replace real AS with local-as in outgoing updates.
* `maximum_prefix` - Maximum number of IPv4 prefixes to accept from this peer.
* `maximum_prefix_threshold` - Maximum IPv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `name` - Neighbor group name.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration6` - Enable/disable allow IPv6 inbound soft reconfiguration.
* `stale_route` - Enable/disable stale route after neighbor down.
* `strict_capability_match` - Enable/disable strict capability matching.
* `unsuppress_map` - IPv4 Route map to selectively unsuppress suppressed routes.
* `unsuppress_map6` - IPv6 Route map to selectively unsuppress suppressed routes.
* `update_source` - Interface to use as source IP/IPv6 address of TCP connections.
* `weight` - Neighbor weight.
* `neighbor_range` - BGP neighbor range table.The structure of `neighbor_range` block is documented below.

The `neighbor_range` block contains:

* `id` - Neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.
* `prefix` - Neighbor range prefix.
* `neighbor_range6` - BGP IPv6 neighbor range table.The structure of `neighbor_range6` block is documented below.

The `neighbor_range6` block contains:

* `id` - IPv6 neighbor range ID.
* `max_neighbor_num` - Maximum number of neighbors.
* `neighbor_group` - Neighbor group name.
* `prefix6` - IPv6 prefix.
* `network` - BGP network table.The structure of `network` block is documented below.

The `network` block contains:

* `backdoor` - Enable/disable route as backdoor.
* `id` - ID.
* `prefix` - Network prefix.
* `route_map` - Route map to modify generated route.
* `network6` - BGP IPv6 network table.The structure of `network6` block is documented below.

The `network6` block contains:

* `backdoor` - Enable/disable route as backdoor.
* `id` - ID.
* `prefix6` - Network IPv6 prefix.
* `route_map` - Route map to modify generated route.
* `redistribute` - BGP IPv4 redistribute table.The structure of `redistribute` block is documented below.

The `redistribute` block contains:

* `name` - Distribute list entry name.
* `route_map` - Route map name.
* `status` - Status
* `redistribute6` - BGP IPv6 redistribute table.The structure of `redistribute6` block is documented below.

The `redistribute6` block contains:

* `name` - Distribute list entry name.
* `route_map` - Route map name.
* `status` - Status
* `vrf_leak` - BGP VRF leaking table.The structure of `vrf_leak` block is documented below.

The `vrf_leak` block contains:

* `vrf` - Origin VRF ID <0 - 31>.
* `target` - Target VRF table.The structure of `target` block is documented below.

The `target` block contains:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID <0 - 31>.
* `vrf_leak6` - BGP IPv6 VRF leaking table.The structure of `vrf_leak6` block is documented below.

The `vrf_leak6` block contains:

* `vrf` - Origin VRF ID <0 - 31>.
* `target` - Target VRF table.The structure of `target` block is documented below.

The `target` block contains:

* `interface` - Interface which is used to leak routes to target VRF.
* `route_map` - Route map of VRF leaking.
* `vrf` - Target VRF ID <0 - 31>.
