// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure BGP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure BGP.",

		ReadContext: dataSourceRouterBgpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"additional_path": {
				Type:        schema.TypeString,
				Description: "Enable/disable selection of BGP IPv4 additional paths.",
				Computed:    true,
			},
			"additional_path_select": {
				Type:        schema.TypeInt,
				Description: "Number of additional paths to be selected for each IPv4 NLRI.",
				Computed:    true,
			},
			"additional_path_select_vpnv4": {
				Type:        schema.TypeInt,
				Description: "Number of additional paths to be selected for each VPNv4 NLRI.",
				Computed:    true,
			},
			"additional_path_select6": {
				Type:        schema.TypeInt,
				Description: "Number of additional paths to be selected for each IPv6 NLRI.",
				Computed:    true,
			},
			"additional_path_vpnv4": {
				Type:        schema.TypeString,
				Description: "Enable/disable selection of BGP VPNv4 additional paths.",
				Computed:    true,
			},
			"additional_path6": {
				Type:        schema.TypeString,
				Description: "Enable/disable selection of BGP IPv6 additional paths.",
				Computed:    true,
			},
			"admin_distance": {
				Type:        schema.TypeList,
				Description: "Administrative distance modifications.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": {
							Type:        schema.TypeInt,
							Description: "Administrative distance to apply (1 - 255).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"neighbour_prefix": {
							Type:        schema.TypeString,
							Description: "Neighbor address prefix.",
							Computed:    true,
						},
						"route_list": {
							Type:        schema.TypeString,
							Description: "Access list of routes to apply new distance to.",
							Computed:    true,
						},
					},
				},
			},
			"aggregate_address": {
				Type:        schema.TypeList,
				Description: "BGP aggregate address table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": {
							Type:        schema.TypeString,
							Description: "Enable/disable generate AS set path information.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Aggregate prefix.",
							Computed:    true,
						},
						"summary_only": {
							Type:        schema.TypeString,
							Description: "Enable/disable filter more specific routes from updates.",
							Computed:    true,
						},
					},
				},
			},
			"aggregate_address6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 aggregate address table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": {
							Type:        schema.TypeString,
							Description: "Enable/disable generate AS set path information.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "Aggregate IPv6 prefix.",
							Computed:    true,
						},
						"summary_only": {
							Type:        schema.TypeString,
							Description: "Enable/disable filter more specific routes from updates.",
							Computed:    true,
						},
					},
				},
			},
			"always_compare_med": {
				Type:        schema.TypeString,
				Description: "Enable/disable always compare MED.",
				Computed:    true,
			},
			"as": {
				Type:        schema.TypeString,
				Description: "Router AS number, asplain/asdot/asdot+ format, 0 to disable BGP.",
				Computed:    true,
			},
			"bestpath_as_path_ignore": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignore AS path.",
				Computed:    true,
			},
			"bestpath_cmp_confed_aspath": {
				Type:        schema.TypeString,
				Description: "Enable/disable compare federation AS path length.",
				Computed:    true,
			},
			"bestpath_cmp_routerid": {
				Type:        schema.TypeString,
				Description: "Enable/disable compare router ID for identical EBGP paths.",
				Computed:    true,
			},
			"bestpath_med_confed": {
				Type:        schema.TypeString,
				Description: "Enable/disable compare MED among confederation paths.",
				Computed:    true,
			},
			"bestpath_med_missing_as_worst": {
				Type:        schema.TypeString,
				Description: "Enable/disable treat missing MED as least preferred.",
				Computed:    true,
			},
			"client_to_client_reflection": {
				Type:        schema.TypeString,
				Description: "Enable/disable client-to-client route reflection.",
				Computed:    true,
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Description: "Route reflector cluster ID.",
				Computed:    true,
			},
			"confederation_identifier": {
				Type:        schema.TypeInt,
				Description: "Confederation identifier.",
				Computed:    true,
			},
			"confederation_peers": {
				Type:        schema.TypeList,
				Description: "Confederation peers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer": {
							Type:        schema.TypeString,
							Description: "Peer ID.",
							Computed:    true,
						},
					},
				},
			},
			"dampening": {
				Type:        schema.TypeString,
				Description: "Enable/disable route-flap dampening.",
				Computed:    true,
			},
			"dampening_max_suppress_time": {
				Type:        schema.TypeInt,
				Description: "Maximum minutes a route can be suppressed.",
				Computed:    true,
			},
			"dampening_reachability_half_life": {
				Type:        schema.TypeInt,
				Description: "Reachability half-life time for penalty (min).",
				Computed:    true,
			},
			"dampening_reuse": {
				Type:        schema.TypeInt,
				Description: "Threshold to reuse routes.",
				Computed:    true,
			},
			"dampening_route_map": {
				Type:        schema.TypeString,
				Description: "Criteria for dampening.",
				Computed:    true,
			},
			"dampening_suppress": {
				Type:        schema.TypeInt,
				Description: "Threshold to suppress routes.",
				Computed:    true,
			},
			"dampening_unreachability_half_life": {
				Type:        schema.TypeInt,
				Description: "Unreachability half-life time for penalty (min).",
				Computed:    true,
			},
			"default_local_preference": {
				Type:        schema.TypeInt,
				Description: "Default local preference.",
				Computed:    true,
			},
			"deterministic_med": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforce deterministic comparison of MED.",
				Computed:    true,
			},
			"distance_external": {
				Type:        schema.TypeInt,
				Description: "Distance for routes external to the AS.",
				Computed:    true,
			},
			"distance_internal": {
				Type:        schema.TypeInt,
				Description: "Distance for routes internal to the AS.",
				Computed:    true,
			},
			"distance_local": {
				Type:        schema.TypeInt,
				Description: "Distance for routes local to the AS.",
				Computed:    true,
			},
			"ebgp_multipath": {
				Type:        schema.TypeString,
				Description: "Enable/disable EBGP multi-path.",
				Computed:    true,
			},
			"enforce_first_as": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforce first AS for EBGP routes.",
				Computed:    true,
			},
			"fast_external_failover": {
				Type:        schema.TypeString,
				Description: "Enable/disable reset peer BGP session if link goes down.",
				Computed:    true,
			},
			"graceful_end_on_timer": {
				Type:        schema.TypeString,
				Description: "Enable/disable to exit graceful restart on timer only.",
				Computed:    true,
			},
			"graceful_restart": {
				Type:        schema.TypeString,
				Description: "Enable/disable BGP graceful restart capabilities.",
				Computed:    true,
			},
			"graceful_restart_time": {
				Type:        schema.TypeInt,
				Description: "Time needed for neighbors to restart (sec).",
				Computed:    true,
			},
			"graceful_stalepath_time": {
				Type:        schema.TypeInt,
				Description: "Time to hold stale paths of restarting neighbor (sec).",
				Computed:    true,
			},
			"graceful_update_delay": {
				Type:        schema.TypeInt,
				Description: "Route advertisement/selection delay after restart (sec).",
				Computed:    true,
			},
			"holdtime_timer": {
				Type:        schema.TypeInt,
				Description: "Number of seconds to mark peer as dead.",
				Computed:    true,
			},
			"ibgp_multipath": {
				Type:        schema.TypeString,
				Description: "Enable/disable IBGP multi-path.",
				Computed:    true,
			},
			"ignore_optional_capability": {
				Type:        schema.TypeString,
				Description: "Do not send unknown optional capability notification message.",
				Computed:    true,
			},
			"keepalive_timer": {
				Type:        schema.TypeInt,
				Description: "Frequency to send keep alive requests.",
				Computed:    true,
			},
			"log_neighbour_changes": {
				Type:        schema.TypeString,
				Description: "Log BGP neighbor changes.",
				Computed:    true,
			},
			"multipath_recursive_distance": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of recursive distance to select multipath.",
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "BGP neighbor table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type:        schema.TypeString,
							Description: "Enable/disable address family IPv4 for this neighbor.",
							Computed:    true,
						},
						"activate_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable address family VPNv4 for this neighbor.",
							Computed:    true,
						},
						"activate6": {
							Type:        schema.TypeString,
							Description: "Enable/disable address family IPv6 for this neighbor.",
							Computed:    true,
						},
						"additional_path": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 additional-path capability.",
							Computed:    true,
						},
						"additional_path_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable VPNv4 additional-path capability.",
							Computed:    true,
						},
						"additional_path6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 additional-path capability.",
							Computed:    true,
						},
						"adv_additional_path": {
							Type:        schema.TypeInt,
							Description: "Number of IPv4 additional paths that can be advertised to this neighbor.",
							Computed:    true,
						},
						"adv_additional_path_vpnv4": {
							Type:        schema.TypeInt,
							Description: "Number of VPNv4 additional paths that can be advertised to this neighbor.",
							Computed:    true,
						},
						"adv_additional_path6": {
							Type:        schema.TypeInt,
							Description: "Number of IPv6 additional paths that can be advertised to this neighbor.",
							Computed:    true,
						},
						"advertisement_interval": {
							Type:        schema.TypeInt,
							Description: "Minimum interval (sec) between sending updates.",
							Computed:    true,
						},
						"allowas_in": {
							Type:        schema.TypeInt,
							Description: "IPv4 The maximum number of occurrence of my AS number allowed.",
							Computed:    true,
						},
						"allowas_in_enable": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
							Computed:    true,
						},
						"allowas_in_enable6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
							Computed:    true,
						},
						"allowas_in_vpnv4": {
							Type:        schema.TypeInt,
							Description: "The maximum number of occurrence of my AS number allowed for VPNv4 route.",
							Computed:    true,
						},
						"allowas_in6": {
							Type:        schema.TypeInt,
							Description: "IPv6 The maximum number of occurrence of my AS number allowed.",
							Computed:    true,
						},
						"as_override": {
							Type:        schema.TypeString,
							Description: "Enable/disable replace peer AS with own AS for IPv4.",
							Computed:    true,
						},
						"as_override6": {
							Type:        schema.TypeString,
							Description: "Enable/disable replace peer AS with own AS for IPv6.",
							Computed:    true,
						},
						"attribute_unchanged": {
							Type:        schema.TypeString,
							Description: "IPv4 List of attributes that should be unchanged.",
							Computed:    true,
						},
						"attribute_unchanged_vpnv4": {
							Type:        schema.TypeString,
							Description: "List of attributes that should be unchanged for VPNv4 route.",
							Computed:    true,
						},
						"attribute_unchanged6": {
							Type:        schema.TypeString,
							Description: "IPv6 List of attributes that should be unchanged.",
							Computed:    true,
						},
						"bfd": {
							Type:        schema.TypeString,
							Description: "Enable/disable BFD for this neighbor.",
							Computed:    true,
						},
						"capability_default_originate": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise default IPv4 route to this neighbor.",
							Computed:    true,
						},
						"capability_default_originate6": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise default IPv6 route to this neighbor.",
							Computed:    true,
						},
						"capability_dynamic": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise dynamic capability to this neighbor.",
							Computed:    true,
						},
						"capability_graceful_restart": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
							Computed:    true,
						},
						"capability_graceful_restart_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise VPNv4 graceful restart capability to this neighbor.",
							Computed:    true,
						},
						"capability_graceful_restart6": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
							Computed:    true,
						},
						"capability_orf": {
							Type:        schema.TypeString,
							Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
							Computed:    true,
						},
						"capability_orf6": {
							Type:        schema.TypeString,
							Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
							Computed:    true,
						},
						"capability_route_refresh": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise route refresh capability to this neighbor.",
							Computed:    true,
						},
						"conditional_advertise": {
							Type:        schema.TypeList,
							Description: "Conditional advertisement.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": {
										Type:        schema.TypeString,
										Description: "Name of advertising route map.",
										Computed:    true,
									},
									"condition_routemap": {
										Type:        schema.TypeList,
										Description: "List of conditional route maps.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Route map.",
													Computed:    true,
												},
											},
										},
									},
									"condition_type": {
										Type:        schema.TypeString,
										Description: "Type of condition.",
										Computed:    true,
									},
								},
							},
						},
						"conditional_advertise6": {
							Type:        schema.TypeList,
							Description: "IPv6 conditional advertisement.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": {
										Type:        schema.TypeString,
										Description: "Name of advertising route map.",
										Computed:    true,
									},
									"condition_routemap": {
										Type:        schema.TypeList,
										Description: "List of conditional route maps.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Route map.",
													Computed:    true,
												},
											},
										},
									},
									"condition_type": {
										Type:        schema.TypeString,
										Description: "Type of condition.",
										Computed:    true,
									},
								},
							},
						},
						"connect_timer": {
							Type:        schema.TypeInt,
							Description: "Interval (sec) for connect timer.",
							Computed:    true,
						},
						"default_originate_routemap": {
							Type:        schema.TypeString,
							Description: "Route map to specify criteria to originate IPv4 default.",
							Computed:    true,
						},
						"default_originate_routemap6": {
							Type:        schema.TypeString,
							Description: "Route map to specify criteria to originate IPv6 default.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"distribute_list_in": {
							Type:        schema.TypeString,
							Description: "Filter for IPv4 updates from this neighbor.",
							Computed:    true,
						},
						"distribute_list_in_vpnv4": {
							Type:        schema.TypeString,
							Description: "Filter for VPNv4 updates from this neighbor.",
							Computed:    true,
						},
						"distribute_list_in6": {
							Type:        schema.TypeString,
							Description: "Filter for IPv6 updates from this neighbor.",
							Computed:    true,
						},
						"distribute_list_out": {
							Type:        schema.TypeString,
							Description: "Filter for IPv4 updates to this neighbor.",
							Computed:    true,
						},
						"distribute_list_out_vpnv4": {
							Type:        schema.TypeString,
							Description: "Filter for VPNv4 updates to this neighbor.",
							Computed:    true,
						},
						"distribute_list_out6": {
							Type:        schema.TypeString,
							Description: "Filter for IPv6 updates to this neighbor.",
							Computed:    true,
						},
						"dont_capability_negotiate": {
							Type:        schema.TypeString,
							Description: "Do not negotiate capabilities with this neighbor.",
							Computed:    true,
						},
						"ebgp_enforce_multihop": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow multi-hop EBGP neighbors.",
							Computed:    true,
						},
						"ebgp_multihop_ttl": {
							Type:        schema.TypeInt,
							Description: "EBGP multihop TTL for this peer.",
							Computed:    true,
						},
						"filter_list_in": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv4 inbound routes.",
							Computed:    true,
						},
						"filter_list_in6": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv6 inbound routes.",
							Computed:    true,
						},
						"filter_list_out": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv4 outbound routes.",
							Computed:    true,
						},
						"filter_list_out6": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv6 outbound routes.",
							Computed:    true,
						},
						"holdtime_timer": {
							Type:        schema.TypeInt,
							Description: "Interval (sec) before peer considered dead.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IP/IPv6 address of neighbor.",
							Computed:    true,
						},
						"keep_alive_timer": {
							Type:        schema.TypeInt,
							Description: "Keep alive timer interval (sec).",
							Computed:    true,
						},
						"link_down_failover": {
							Type:        schema.TypeString,
							Description: "Enable/disable failover upon link down.",
							Computed:    true,
						},
						"local_as": {
							Type:        schema.TypeString,
							Description: "Local AS number of neighbor.",
							Computed:    true,
						},
						"local_as_no_prepend": {
							Type:        schema.TypeString,
							Description: "Do not prepend local-as to incoming updates.",
							Computed:    true,
						},
						"local_as_replace_as": {
							Type:        schema.TypeString,
							Description: "Replace real AS with local-as in outgoing updates.",
							Computed:    true,
						},
						"maximum_prefix": {
							Type:        schema.TypeInt,
							Description: "Maximum number of IPv4 prefixes to accept from this peer.",
							Computed:    true,
						},
						"maximum_prefix_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum IPv4 prefix threshold value (1 - 100 percent).",
							Computed:    true,
						},
						"maximum_prefix_threshold_vpnv4": {
							Type:        schema.TypeInt,
							Description: "Maximum VPNv4 prefix threshold value (1 - 100 percent).",
							Computed:    true,
						},
						"maximum_prefix_threshold6": {
							Type:        schema.TypeInt,
							Description: "Maximum IPv6 prefix threshold value (1 - 100 percent).",
							Computed:    true,
						},
						"maximum_prefix_vpnv4": {
							Type:        schema.TypeInt,
							Description: "Maximum number of VPNv4 prefixes to accept from this peer.",
							Computed:    true,
						},
						"maximum_prefix_warning_only": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
							Computed:    true,
						},
						"maximum_prefix_warning_only_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.",
							Computed:    true,
						},
						"maximum_prefix_warning_only6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 Only give warning message when limit is exceeded.",
							Computed:    true,
						},
						"maximum_prefix6": {
							Type:        schema.TypeInt,
							Description: "Maximum number of IPv6 prefixes to accept from this peer.",
							Computed:    true,
						},
						"next_hop_self": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
							Computed:    true,
						},
						"next_hop_self_rr": {
							Type:        schema.TypeString,
							Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
							Computed:    true,
						},
						"next_hop_self_rr6": {
							Type:        schema.TypeString,
							Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
							Computed:    true,
						},
						"next_hop_self_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.",
							Computed:    true,
						},
						"next_hop_self6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
							Computed:    true,
						},
						"override_capability": {
							Type:        schema.TypeString,
							Description: "Enable/disable override result of capability negotiation.",
							Computed:    true,
						},
						"passive": {
							Type:        schema.TypeString,
							Description: "Enable/disable sending of open messages to this neighbor.",
							Computed:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Password used in MD5 authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"prefix_list_in": {
							Type:        schema.TypeString,
							Description: "IPv4 Inbound filter for updates from this neighbor.",
							Computed:    true,
						},
						"prefix_list_in_vpnv4": {
							Type:        schema.TypeString,
							Description: "Inbound filter for VPNv4 updates from this neighbor.",
							Computed:    true,
						},
						"prefix_list_in6": {
							Type:        schema.TypeString,
							Description: "IPv6 Inbound filter for updates from this neighbor.",
							Computed:    true,
						},
						"prefix_list_out": {
							Type:        schema.TypeString,
							Description: "IPv4 Outbound filter for updates to this neighbor.",
							Computed:    true,
						},
						"prefix_list_out_vpnv4": {
							Type:        schema.TypeString,
							Description: "Outbound filter for VPNv4 updates to this neighbor.",
							Computed:    true,
						},
						"prefix_list_out6": {
							Type:        schema.TypeString,
							Description: "IPv6 Outbound filter for updates to this neighbor.",
							Computed:    true,
						},
						"remote_as": {
							Type:        schema.TypeString,
							Description: "AS number of neighbor.",
							Computed:    true,
						},
						"remove_private_as": {
							Type:        schema.TypeString,
							Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
							Computed:    true,
						},
						"remove_private_as_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable remove private AS number from VPNv4 outbound updates.",
							Computed:    true,
						},
						"remove_private_as6": {
							Type:        schema.TypeString,
							Description: "Enable/disable remove private AS number from IPv6 outbound updates.",
							Computed:    true,
						},
						"restart_time": {
							Type:        schema.TypeInt,
							Description: "Graceful restart delay time (sec, 0 = global default).",
							Computed:    true,
						},
						"retain_stale_time": {
							Type:        schema.TypeInt,
							Description: "Time to retain stale routes.",
							Computed:    true,
						},
						"route_map_in": {
							Type:        schema.TypeString,
							Description: "IPv4 Inbound route map filter.",
							Computed:    true,
						},
						"route_map_in_vpnv4": {
							Type:        schema.TypeString,
							Description: "VPNv4 inbound route map filter.",
							Computed:    true,
						},
						"route_map_in6": {
							Type:        schema.TypeString,
							Description: "IPv6 Inbound route map filter.",
							Computed:    true,
						},
						"route_map_out": {
							Type:        schema.TypeString,
							Description: "IPv4 outbound route map filter.",
							Computed:    true,
						},
						"route_map_out_preferable": {
							Type:        schema.TypeString,
							Description: "IPv4 outbound route map filter if the peer is preferred.",
							Computed:    true,
						},
						"route_map_out_vpnv4": {
							Type:        schema.TypeString,
							Description: "VPNv4 outbound route map filter.",
							Computed:    true,
						},
						"route_map_out_vpnv4_preferable": {
							Type:        schema.TypeString,
							Description: "VPNv4 outbound route map filter if the peer is preferred.",
							Computed:    true,
						},
						"route_map_out6": {
							Type:        schema.TypeString,
							Description: "IPv6 Outbound route map filter.",
							Computed:    true,
						},
						"route_map_out6_preferable": {
							Type:        schema.TypeString,
							Description: "IPv6 outbound route map filter if the peer is preferred.",
							Computed:    true,
						},
						"route_reflector_client": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 AS route reflector client.",
							Computed:    true,
						},
						"route_reflector_client_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable VPNv4 AS route reflector client for this neighbor.",
							Computed:    true,
						},
						"route_reflector_client6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 AS route reflector client.",
							Computed:    true,
						},
						"route_server_client": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 AS route server client.",
							Computed:    true,
						},
						"route_server_client_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable VPNv4 AS route server client for this neighbor.",
							Computed:    true,
						},
						"route_server_client6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 AS route server client.",
							Computed:    true,
						},
						"send_community": {
							Type:        schema.TypeString,
							Description: "IPv4 Send community attribute to neighbor.",
							Computed:    true,
						},
						"send_community_vpnv4": {
							Type:        schema.TypeString,
							Description: "Send community attribute to neighbor for VPNv4 address family.",
							Computed:    true,
						},
						"send_community6": {
							Type:        schema.TypeString,
							Description: "IPv6 Send community attribute to neighbor.",
							Computed:    true,
						},
						"shutdown": {
							Type:        schema.TypeString,
							Description: "Enable/disable shutdown this neighbor.",
							Computed:    true,
						},
						"soft_reconfiguration": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
							Computed:    true,
						},
						"soft_reconfiguration_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow VPNv4 inbound soft reconfiguration.",
							Computed:    true,
						},
						"soft_reconfiguration6": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
							Computed:    true,
						},
						"stale_route": {
							Type:        schema.TypeString,
							Description: "Enable/disable stale route after neighbor down.",
							Computed:    true,
						},
						"strict_capability_match": {
							Type:        schema.TypeString,
							Description: "Enable/disable strict capability matching.",
							Computed:    true,
						},
						"unsuppress_map": {
							Type:        schema.TypeString,
							Description: "IPv4 Route map to selectively unsuppress suppressed routes.",
							Computed:    true,
						},
						"unsuppress_map6": {
							Type:        schema.TypeString,
							Description: "IPv6 Route map to selectively unsuppress suppressed routes.",
							Computed:    true,
						},
						"update_source": {
							Type:        schema.TypeString,
							Description: "Interface to use as source IP/IPv6 address of TCP connections.",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "Neighbor weight.",
							Computed:    true,
						},
					},
				},
			},
			"neighbor_group": {
				Type:        schema.TypeList,
				Description: "BGP neighbor group table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type:        schema.TypeString,
							Description: "Enable/disable address family IPv4 for this neighbor.",
							Computed:    true,
						},
						"activate_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable address family VPNv4 for this neighbor.",
							Computed:    true,
						},
						"activate6": {
							Type:        schema.TypeString,
							Description: "Enable/disable address family IPv6 for this neighbor.",
							Computed:    true,
						},
						"additional_path": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 additional-path capability.",
							Computed:    true,
						},
						"additional_path_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable VPNv4 additional-path capability.",
							Computed:    true,
						},
						"additional_path6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 additional-path capability.",
							Computed:    true,
						},
						"adv_additional_path": {
							Type:        schema.TypeInt,
							Description: "Number of IPv4 additional paths that can be advertised to this neighbor.",
							Computed:    true,
						},
						"adv_additional_path_vpnv4": {
							Type:        schema.TypeInt,
							Description: "Number of VPNv4 additional paths that can be advertised to this neighbor.",
							Computed:    true,
						},
						"adv_additional_path6": {
							Type:        schema.TypeInt,
							Description: "Number of IPv6 additional paths that can be advertised to this neighbor.",
							Computed:    true,
						},
						"advertisement_interval": {
							Type:        schema.TypeInt,
							Description: "Minimum interval (sec) between sending updates.",
							Computed:    true,
						},
						"allowas_in": {
							Type:        schema.TypeInt,
							Description: "IPv4 The maximum number of occurrence of my AS number allowed.",
							Computed:    true,
						},
						"allowas_in_enable": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
							Computed:    true,
						},
						"allowas_in_enable6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
							Computed:    true,
						},
						"allowas_in_vpnv4": {
							Type:        schema.TypeInt,
							Description: "The maximum number of occurrence of my AS number allowed for VPNv4 route.",
							Computed:    true,
						},
						"allowas_in6": {
							Type:        schema.TypeInt,
							Description: "IPv6 The maximum number of occurrence of my AS number allowed.",
							Computed:    true,
						},
						"as_override": {
							Type:        schema.TypeString,
							Description: "Enable/disable replace peer AS with own AS for IPv4.",
							Computed:    true,
						},
						"as_override6": {
							Type:        schema.TypeString,
							Description: "Enable/disable replace peer AS with own AS for IPv6.",
							Computed:    true,
						},
						"attribute_unchanged": {
							Type:        schema.TypeString,
							Description: "IPv4 List of attributes that should be unchanged.",
							Computed:    true,
						},
						"attribute_unchanged_vpnv4": {
							Type:        schema.TypeString,
							Description: "List of attributes that should be unchanged for VPNv4 route.",
							Computed:    true,
						},
						"attribute_unchanged6": {
							Type:        schema.TypeString,
							Description: "IPv6 List of attributes that should be unchanged.",
							Computed:    true,
						},
						"bfd": {
							Type:        schema.TypeString,
							Description: "Enable/disable BFD for this neighbor.",
							Computed:    true,
						},
						"capability_default_originate": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise default IPv4 route to this neighbor.",
							Computed:    true,
						},
						"capability_default_originate6": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise default IPv6 route to this neighbor.",
							Computed:    true,
						},
						"capability_dynamic": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise dynamic capability to this neighbor.",
							Computed:    true,
						},
						"capability_graceful_restart": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
							Computed:    true,
						},
						"capability_graceful_restart_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise VPNv4 graceful restart capability to this neighbor.",
							Computed:    true,
						},
						"capability_graceful_restart6": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
							Computed:    true,
						},
						"capability_orf": {
							Type:        schema.TypeString,
							Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
							Computed:    true,
						},
						"capability_orf6": {
							Type:        schema.TypeString,
							Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
							Computed:    true,
						},
						"capability_route_refresh": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise route refresh capability to this neighbor.",
							Computed:    true,
						},
						"connect_timer": {
							Type:        schema.TypeInt,
							Description: "Interval (sec) for connect timer.",
							Computed:    true,
						},
						"default_originate_routemap": {
							Type:        schema.TypeString,
							Description: "Route map to specify criteria to originate IPv4 default.",
							Computed:    true,
						},
						"default_originate_routemap6": {
							Type:        schema.TypeString,
							Description: "Route map to specify criteria to originate IPv6 default.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"distribute_list_in": {
							Type:        schema.TypeString,
							Description: "Filter for IPv4 updates from this neighbor.",
							Computed:    true,
						},
						"distribute_list_in_vpnv4": {
							Type:        schema.TypeString,
							Description: "Filter for VPNv4 updates from this neighbor.",
							Computed:    true,
						},
						"distribute_list_in6": {
							Type:        schema.TypeString,
							Description: "Filter for IPv6 updates from this neighbor.",
							Computed:    true,
						},
						"distribute_list_out": {
							Type:        schema.TypeString,
							Description: "Filter for IPv4 updates to this neighbor.",
							Computed:    true,
						},
						"distribute_list_out_vpnv4": {
							Type:        schema.TypeString,
							Description: "Filter for VPNv4 updates to this neighbor.",
							Computed:    true,
						},
						"distribute_list_out6": {
							Type:        schema.TypeString,
							Description: "Filter for IPv6 updates to this neighbor.",
							Computed:    true,
						},
						"dont_capability_negotiate": {
							Type:        schema.TypeString,
							Description: "Do not negotiate capabilities with this neighbor.",
							Computed:    true,
						},
						"ebgp_enforce_multihop": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow multi-hop EBGP neighbors.",
							Computed:    true,
						},
						"ebgp_multihop_ttl": {
							Type:        schema.TypeInt,
							Description: "EBGP multihop TTL for this peer.",
							Computed:    true,
						},
						"filter_list_in": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv4 inbound routes.",
							Computed:    true,
						},
						"filter_list_in6": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv6 inbound routes.",
							Computed:    true,
						},
						"filter_list_out": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv4 outbound routes.",
							Computed:    true,
						},
						"filter_list_out6": {
							Type:        schema.TypeString,
							Description: "BGP filter for IPv6 outbound routes.",
							Computed:    true,
						},
						"holdtime_timer": {
							Type:        schema.TypeInt,
							Description: "Interval (sec) before peer considered dead.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.",
							Computed:    true,
						},
						"keep_alive_timer": {
							Type:        schema.TypeInt,
							Description: "Keep alive timer interval (sec).",
							Computed:    true,
						},
						"link_down_failover": {
							Type:        schema.TypeString,
							Description: "Enable/disable failover upon link down.",
							Computed:    true,
						},
						"local_as": {
							Type:        schema.TypeString,
							Description: "Local AS number of neighbor.",
							Computed:    true,
						},
						"local_as_no_prepend": {
							Type:        schema.TypeString,
							Description: "Do not prepend local-as to incoming updates.",
							Computed:    true,
						},
						"local_as_replace_as": {
							Type:        schema.TypeString,
							Description: "Replace real AS with local-as in outgoing updates.",
							Computed:    true,
						},
						"maximum_prefix": {
							Type:        schema.TypeInt,
							Description: "Maximum number of IPv4 prefixes to accept from this peer.",
							Computed:    true,
						},
						"maximum_prefix_threshold": {
							Type:        schema.TypeInt,
							Description: "Maximum IPv4 prefix threshold value (1 - 100 percent).",
							Computed:    true,
						},
						"maximum_prefix_threshold_vpnv4": {
							Type:        schema.TypeInt,
							Description: "Maximum VPNv4 prefix threshold value (1 - 100 percent).",
							Computed:    true,
						},
						"maximum_prefix_threshold6": {
							Type:        schema.TypeInt,
							Description: "Maximum IPv6 prefix threshold value (1 - 100 percent).",
							Computed:    true,
						},
						"maximum_prefix_vpnv4": {
							Type:        schema.TypeInt,
							Description: "Maximum number of VPNv4 prefixes to accept from this peer.",
							Computed:    true,
						},
						"maximum_prefix_warning_only": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
							Computed:    true,
						},
						"maximum_prefix_warning_only_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.",
							Computed:    true,
						},
						"maximum_prefix_warning_only6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 Only give warning message when limit is exceeded.",
							Computed:    true,
						},
						"maximum_prefix6": {
							Type:        schema.TypeInt,
							Description: "Maximum number of IPv6 prefixes to accept from this peer.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Neighbor group name.",
							Computed:    true,
						},
						"next_hop_self": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
							Computed:    true,
						},
						"next_hop_self_rr": {
							Type:        schema.TypeString,
							Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
							Computed:    true,
						},
						"next_hop_self_rr6": {
							Type:        schema.TypeString,
							Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
							Computed:    true,
						},
						"next_hop_self_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.",
							Computed:    true,
						},
						"next_hop_self6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
							Computed:    true,
						},
						"override_capability": {
							Type:        schema.TypeString,
							Description: "Enable/disable override result of capability negotiation.",
							Computed:    true,
						},
						"passive": {
							Type:        schema.TypeString,
							Description: "Enable/disable sending of open messages to this neighbor.",
							Computed:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Password used in MD5 authentication.",
							Computed:    true,
							Sensitive:   true,
						},
						"prefix_list_in": {
							Type:        schema.TypeString,
							Description: "IPv4 Inbound filter for updates from this neighbor.",
							Computed:    true,
						},
						"prefix_list_in_vpnv4": {
							Type:        schema.TypeString,
							Description: "Inbound filter for VPNv4 updates from this neighbor.",
							Computed:    true,
						},
						"prefix_list_in6": {
							Type:        schema.TypeString,
							Description: "IPv6 Inbound filter for updates from this neighbor.",
							Computed:    true,
						},
						"prefix_list_out": {
							Type:        schema.TypeString,
							Description: "IPv4 Outbound filter for updates to this neighbor.",
							Computed:    true,
						},
						"prefix_list_out_vpnv4": {
							Type:        schema.TypeString,
							Description: "Outbound filter for VPNv4 updates to this neighbor.",
							Computed:    true,
						},
						"prefix_list_out6": {
							Type:        schema.TypeString,
							Description: "IPv6 Outbound filter for updates to this neighbor.",
							Computed:    true,
						},
						"remote_as": {
							Type:        schema.TypeString,
							Description: "AS number of neighbor.",
							Computed:    true,
						},
						"remove_private_as": {
							Type:        schema.TypeString,
							Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
							Computed:    true,
						},
						"remove_private_as_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable remove private AS number from VPNv4 outbound updates.",
							Computed:    true,
						},
						"remove_private_as6": {
							Type:        schema.TypeString,
							Description: "Enable/disable remove private AS number from IPv6 outbound updates.",
							Computed:    true,
						},
						"restart_time": {
							Type:        schema.TypeInt,
							Description: "Graceful restart delay time (sec, 0 = global default).",
							Computed:    true,
						},
						"retain_stale_time": {
							Type:        schema.TypeInt,
							Description: "Time to retain stale routes.",
							Computed:    true,
						},
						"route_map_in": {
							Type:        schema.TypeString,
							Description: "IPv4 Inbound route map filter.",
							Computed:    true,
						},
						"route_map_in_vpnv4": {
							Type:        schema.TypeString,
							Description: "VPNv4 inbound route map filter.",
							Computed:    true,
						},
						"route_map_in6": {
							Type:        schema.TypeString,
							Description: "IPv6 Inbound route map filter.",
							Computed:    true,
						},
						"route_map_out": {
							Type:        schema.TypeString,
							Description: "IPv4 outbound route map filter.",
							Computed:    true,
						},
						"route_map_out_preferable": {
							Type:        schema.TypeString,
							Description: "IPv4 outbound route map filter if the peer is preferred.",
							Computed:    true,
						},
						"route_map_out_vpnv4": {
							Type:        schema.TypeString,
							Description: "VPNv4 outbound route map filter.",
							Computed:    true,
						},
						"route_map_out_vpnv4_preferable": {
							Type:        schema.TypeString,
							Description: "VPNv4 outbound route map filter if the peer is preferred.",
							Computed:    true,
						},
						"route_map_out6": {
							Type:        schema.TypeString,
							Description: "IPv6 Outbound route map filter.",
							Computed:    true,
						},
						"route_map_out6_preferable": {
							Type:        schema.TypeString,
							Description: "IPv6 outbound route map filter if the peer is preferred.",
							Computed:    true,
						},
						"route_reflector_client": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 AS route reflector client.",
							Computed:    true,
						},
						"route_reflector_client_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable VPNv4 AS route reflector client for this neighbor.",
							Computed:    true,
						},
						"route_reflector_client6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 AS route reflector client.",
							Computed:    true,
						},
						"route_server_client": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv4 AS route server client.",
							Computed:    true,
						},
						"route_server_client_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable VPNv4 AS route server client for this neighbor.",
							Computed:    true,
						},
						"route_server_client6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 AS route server client.",
							Computed:    true,
						},
						"send_community": {
							Type:        schema.TypeString,
							Description: "IPv4 Send community attribute to neighbor.",
							Computed:    true,
						},
						"send_community_vpnv4": {
							Type:        schema.TypeString,
							Description: "Send community attribute to neighbor for VPNv4 address family.",
							Computed:    true,
						},
						"send_community6": {
							Type:        schema.TypeString,
							Description: "IPv6 Send community attribute to neighbor.",
							Computed:    true,
						},
						"shutdown": {
							Type:        schema.TypeString,
							Description: "Enable/disable shutdown this neighbor.",
							Computed:    true,
						},
						"soft_reconfiguration": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
							Computed:    true,
						},
						"soft_reconfiguration_vpnv4": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow VPNv4 inbound soft reconfiguration.",
							Computed:    true,
						},
						"soft_reconfiguration6": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
							Computed:    true,
						},
						"stale_route": {
							Type:        schema.TypeString,
							Description: "Enable/disable stale route after neighbor down.",
							Computed:    true,
						},
						"strict_capability_match": {
							Type:        schema.TypeString,
							Description: "Enable/disable strict capability matching.",
							Computed:    true,
						},
						"unsuppress_map": {
							Type:        schema.TypeString,
							Description: "IPv4 Route map to selectively unsuppress suppressed routes.",
							Computed:    true,
						},
						"unsuppress_map6": {
							Type:        schema.TypeString,
							Description: "IPv6 Route map to selectively unsuppress suppressed routes.",
							Computed:    true,
						},
						"update_source": {
							Type:        schema.TypeString,
							Description: "Interface to use as source IP/IPv6 address of TCP connections.",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "Neighbor weight.",
							Computed:    true,
						},
					},
				},
			},
			"neighbor_range": {
				Type:        schema.TypeList,
				Description: "BGP neighbor range table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Neighbor range ID.",
							Computed:    true,
						},
						"max_neighbor_num": {
							Type:        schema.TypeInt,
							Description: "Maximum number of neighbors.",
							Computed:    true,
						},
						"neighbor_group": {
							Type:        schema.TypeString,
							Description: "Neighbor group name.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Neighbor range prefix.",
							Computed:    true,
						},
					},
				},
			},
			"neighbor_range6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 neighbor range table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "IPv6 neighbor range ID.",
							Computed:    true,
						},
						"max_neighbor_num": {
							Type:        schema.TypeInt,
							Description: "Maximum number of neighbors.",
							Computed:    true,
						},
						"neighbor_group": {
							Type:        schema.TypeString,
							Description: "Neighbor group name.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "IPv6 prefix.",
							Computed:    true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "BGP network table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": {
							Type:        schema.TypeString,
							Description: "Enable/disable route as backdoor.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"network_import_check": {
							Type:        schema.TypeString,
							Description: "Configure insurance of BGP network route existence in IGP.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Network prefix.",
							Computed:    true,
						},
						"route_map": {
							Type:        schema.TypeString,
							Description: "Route map to modify generated route.",
							Computed:    true,
						},
					},
				},
			},
			"network_import_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable ensure BGP network route exists in IGP.",
				Computed:    true,
			},
			"network6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 network table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": {
							Type:        schema.TypeString,
							Description: "Enable/disable route as backdoor.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"network_import_check": {
							Type:        schema.TypeString,
							Description: "Configure insurance of BGP network route existence in IGP.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "Network IPv6 prefix.",
							Computed:    true,
						},
						"route_map": {
							Type:        schema.TypeString,
							Description: "Route map to modify generated route.",
							Computed:    true,
						},
					},
				},
			},
			"recursive_inherit_priority": {
				Type:        schema.TypeString,
				Description: "Enable/disable priority inheritance for recursive resolution.",
				Computed:    true,
			},
			"recursive_next_hop": {
				Type:        schema.TypeString,
				Description: "Enable/disable recursive resolution of next-hop using BGP route.",
				Computed:    true,
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "BGP IPv4 redistribute table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Distribute list entry name.",
							Computed:    true,
						},
						"route_map": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
					},
				},
			},
			"redistribute6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 redistribute table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Distribute list entry name.",
							Computed:    true,
						},
						"route_map": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
					},
				},
			},
			"router_id": {
				Type:        schema.TypeString,
				Description: "Router ID.",
				Computed:    true,
			},
			"scan_time": {
				Type:        schema.TypeInt,
				Description: "Background scanner interval (sec), 0 to disable it.",
				Computed:    true,
			},
			"synchronization": {
				Type:        schema.TypeString,
				Description: "Enable/disable only advertise routes from iBGP if routes present in an IGP.",
				Computed:    true,
			},
			"tag_resolve_mode": {
				Type:        schema.TypeString,
				Description: "Configure tag-match mode. Resolves BGP routes with other routes containing the same tag.",
				Computed:    true,
			},
			"vrf": {
				Type:        schema.TypeList,
				Description: "BGP VRF leaking table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_rt": {
							Type:        schema.TypeList,
							Description: "List of export route target.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": {
										Type:        schema.TypeString,
										Description: "Attribute: AA|AA:NN.",
										Computed:    true,
									},
								},
							},
						},
						"import_route_map": {
							Type:        schema.TypeString,
							Description: "Import route map.",
							Computed:    true,
						},
						"import_rt": {
							Type:        schema.TypeList,
							Description: "List of import route target.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": {
										Type:        schema.TypeString,
										Description: "Attribute: AA|AA:NN.",
										Computed:    true,
									},
								},
							},
						},
						"leak_target": {
							Type:        schema.TypeList,
							Description: "Target VRF table.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:        schema.TypeString,
										Description: "Interface which is used to leak routes to target VRF.",
										Computed:    true,
									},
									"route_map": {
										Type:        schema.TypeString,
										Description: "Route map of VRF leaking.",
										Computed:    true,
									},
									"vrf": {
										Type:        schema.TypeString,
										Description: "Target VRF ID (0 - 251).",
										Computed:    true,
									},
								},
							},
						},
						"rd": {
							Type:        schema.TypeString,
							Description: "Route Distinguisher: AA|AA:NN.",
							Computed:    true,
						},
						"role": {
							Type:        schema.TypeString,
							Description: "VRF role.",
							Computed:    true,
						},
						"vrf": {
							Type:        schema.TypeString,
							Description: "Origin VRF ID (0 - 251).",
							Computed:    true,
						},
					},
				},
			},
			"vrf_leak": {
				Type:        schema.TypeList,
				Description: "BGP VRF leaking table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:        schema.TypeList,
							Description: "Target VRF table.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:        schema.TypeString,
										Description: "Interface which is used to leak routes to target VRF.",
										Computed:    true,
									},
									"route_map": {
										Type:        schema.TypeString,
										Description: "Route map of VRF leaking.",
										Computed:    true,
									},
									"vrf": {
										Type:        schema.TypeString,
										Description: "Target VRF ID (0 - 31).",
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:        schema.TypeString,
							Description: "Origin VRF ID (0 - 31).",
							Computed:    true,
						},
					},
				},
			},
			"vrf_leak6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 VRF leaking table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:        schema.TypeList,
							Description: "Target VRF table.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:        schema.TypeString,
										Description: "Interface which is used to leak routes to target VRF.",
										Computed:    true,
									},
									"route_map": {
										Type:        schema.TypeString,
										Description: "Route map of VRF leaking.",
										Computed:    true,
									},
									"vrf": {
										Type:        schema.TypeString,
										Description: "Target VRF ID (0 - 31).",
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:        schema.TypeString,
							Description: "Origin VRF ID (0 - 31).",
							Computed:    true,
						},
					},
				},
			},
			"vrf6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 VRF leaking table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leak_target": {
							Type:        schema.TypeList,
							Description: "Target VRF table.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:        schema.TypeString,
										Description: "Interface which is used to leak routes to target VRF.",
										Computed:    true,
									},
									"route_map": {
										Type:        schema.TypeString,
										Description: "Route map of VRF leaking.",
										Computed:    true,
									},
									"vrf": {
										Type:        schema.TypeString,
										Description: "Target VRF ID (0 - 251).",
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:        schema.TypeString,
							Description: "Origin VRF ID (0 - 251).",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	mkey := "RouterBgp"

	o, err := c.Cmdb.ReadRouterBgp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgp dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectRouterBgp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
