---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bgp_neighbor"
description: |-
  Get information on a fortios BGP neighbor table.
---

# Data Source: fortios_router_bgp_neighbor
Use this data source to get information on a fortios BGP neighbor table.


## Example Usage

```hcl

```

## Argument Reference

* `ip` - (Required) IP/IPv6 address of neighbor.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `activate` - Enable/disable address family IPv4 for this neighbor.
* `activate_vpnv4` - Enable/disable address family VPNv4 for this neighbor.
* `activate6` - Enable/disable address family IPv6 for this neighbor.
* `additional_path` - Enable/disable IPv4 additional-path capability.
* `additional_path_vpnv4` - Enable/disable VPNv4 additional-path capability.
* `additional_path6` - Enable/disable IPv6 additional-path capability.
* `adv_additional_path` - Number of IPv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path_vpnv4` - Number of VPNv4 additional paths that can be advertised to this neighbor.
* `adv_additional_path6` - Number of IPv6 additional paths that can be advertised to this neighbor.
* `advertisement_interval` - Minimum interval (sec) between sending updates.
* `allowas_in` - IPv4 The maximum number of occurrence of my AS number allowed.
* `allowas_in_enable` - Enable/disable IPv4 Enable to allow my AS in AS path.
* `allowas_in_enable6` - Enable/disable IPv6 Enable to allow my AS in AS path.
* `allowas_in_vpnv4` - The maximum number of occurrence of my AS number allowed for VPNv4 route.
* `allowas_in6` - IPv6 The maximum number of occurrence of my AS number allowed.
* `as_override` - Enable/disable replace peer AS with own AS for IPv4.
* `as_override6` - Enable/disable replace peer AS with own AS for IPv6.
* `attribute_unchanged` - IPv4 List of attributes that should be unchanged.
* `attribute_unchanged_vpnv4` - List of attributes that should be unchanged for VPNv4 route.
* `attribute_unchanged6` - IPv6 List of attributes that should be unchanged.
* `bfd` - Enable/disable BFD for this neighbor.
* `capability_default_originate` - Enable/disable advertise default IPv4 route to this neighbor.
* `capability_default_originate6` - Enable/disable advertise default IPv6 route to this neighbor.
* `capability_dynamic` - Enable/disable advertise dynamic capability to this neighbor.
* `capability_graceful_restart` - Enable/disable advertise IPv4 graceful restart capability to this neighbor.
* `capability_graceful_restart_vpnv4` - Enable/disable advertise VPNv4 graceful restart capability to this neighbor.
* `capability_graceful_restart6` - Enable/disable advertise IPv6 graceful restart capability to this neighbor.
* `capability_orf` - Accept/Send IPv4 ORF lists to/from this neighbor.
* `capability_orf6` - Accept/Send IPv6 ORF lists to/from this neighbor.
* `capability_route_refresh` - Enable/disable advertise route refresh capability to this neighbor.
* `connect_timer` - Interval (sec) for connect timer.
* `default_originate_routemap` - Route map to specify criteria to originate IPv4 default.
* `default_originate_routemap6` - Route map to specify criteria to originate IPv6 default.
* `description` - Description.
* `distribute_list_in` - Filter for IPv4 updates from this neighbor.
* `distribute_list_in_vpnv4` - Filter for VPNv4 updates from this neighbor.
* `distribute_list_in6` - Filter for IPv6 updates from this neighbor.
* `distribute_list_out` - Filter for IPv4 updates to this neighbor.
* `distribute_list_out_vpnv4` - Filter for VPNv4 updates to this neighbor.
* `distribute_list_out6` - Filter for IPv6 updates to this neighbor.
* `dont_capability_negotiate` - Do not negotiate capabilities with this neighbor.
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
* `maximum_prefix_threshold_vpnv4` - Maximum VPNv4 prefix threshold value (1 - 100 percent).
* `maximum_prefix_threshold6` - Maximum IPv6 prefix threshold value (1 - 100 percent).
* `maximum_prefix_vpnv4` - Maximum number of VPNv4 prefixes to accept from this peer.
* `maximum_prefix_warning_only` - Enable/disable IPv4 Only give warning message when limit is exceeded.
* `maximum_prefix_warning_only_vpnv4` - Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.
* `maximum_prefix_warning_only6` - Enable/disable IPv6 Only give warning message when limit is exceeded.
* `maximum_prefix6` - Maximum number of IPv6 prefixes to accept from this peer.
* `next_hop_self` - Enable/disable IPv4 next-hop calculation for this neighbor.
* `next_hop_self_rr` - Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.
* `next_hop_self_rr6` - Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.
* `next_hop_self_vpnv4` - Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.
* `next_hop_self6` - Enable/disable IPv6 next-hop calculation for this neighbor.
* `override_capability` - Enable/disable override result of capability negotiation.
* `passive` - Enable/disable sending of open messages to this neighbor.
* `password` - Password used in MD5 authentication.
* `prefix_list_in` - IPv4 Inbound filter for updates from this neighbor.
* `prefix_list_in_vpnv4` - Inbound filter for VPNv4 updates from this neighbor.
* `prefix_list_in6` - IPv6 Inbound filter for updates from this neighbor.
* `prefix_list_out` - IPv4 Outbound filter for updates to this neighbor.
* `prefix_list_out_vpnv4` - Outbound filter for VPNv4 updates to this neighbor.
* `prefix_list_out6` - IPv6 Outbound filter for updates to this neighbor.
* `remote_as` - AS number of neighbor.
* `remove_private_as` - Enable/disable remove private AS number from IPv4 outbound updates.
* `remove_private_as_vpnv4` - Enable/disable remove private AS number from VPNv4 outbound updates.
* `remove_private_as6` - Enable/disable remove private AS number from IPv6 outbound updates.
* `restart_time` - Graceful restart delay time (sec, 0 = global default).
* `retain_stale_time` - Time to retain stale routes.
* `route_map_in` - IPv4 Inbound route map filter.
* `route_map_in_vpnv4` - VPNv4 inbound route map filter.
* `route_map_in6` - IPv6 Inbound route map filter.
* `route_map_out` - IPv4 outbound route map filter.
* `route_map_out_preferable` - IPv4 outbound route map filter if the peer is preferred.
* `route_map_out_vpnv4` - VPNv4 outbound route map filter.
* `route_map_out_vpnv4_preferable` - VPNv4 outbound route map filter if the peer is preferred.
* `route_map_out6` - IPv6 Outbound route map filter.
* `route_map_out6_preferable` - IPv6 outbound route map filter if the peer is preferred.
* `route_reflector_client` - Enable/disable IPv4 AS route reflector client.
* `route_reflector_client_vpnv4` - Enable/disable VPNv4 AS route reflector client for this neighbor.
* `route_reflector_client6` - Enable/disable IPv6 AS route reflector client.
* `route_server_client` - Enable/disable IPv4 AS route server client.
* `route_server_client_vpnv4` - Enable/disable VPNv4 AS route server client for this neighbor.
* `route_server_client6` - Enable/disable IPv6 AS route server client.
* `send_community` - IPv4 Send community attribute to neighbor.
* `send_community_vpnv4` - Send community attribute to neighbor for VPNv4 address family.
* `send_community6` - IPv6 Send community attribute to neighbor.
* `shutdown` - Enable/disable shutdown this neighbor.
* `soft_reconfiguration` - Enable/disable allow IPv4 inbound soft reconfiguration.
* `soft_reconfiguration_vpnv4` - Enable/disable allow VPNv4 inbound soft reconfiguration.
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
* `condition_type` - Type of condition.
* `condition_routemap` - List of conditional route maps.The structure of `condition_routemap` block is documented below.

The `condition_routemap` block contains:

* `name` - route map
* `conditional_advertise6` - IPv6 conditional advertisement.The structure of `conditional_advertise6` block is documented below.

The `conditional_advertise6` block contains:

* `advertise_routemap` - Name of advertising route map.
* `condition_type` - Type of condition.
* `condition_routemap` - List of conditional route maps.The structure of `condition_routemap` block is documented below.

The `condition_routemap` block contains:

* `name` - route map
