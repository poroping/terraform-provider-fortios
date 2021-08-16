// Inspired by Official Fortinet Provider https://github.com/fortinetdev/terraform-provider-fortios
// Shout out to: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Generated from template using FortiOS v7.0.1 schema
// Template Authors:
// Justin Roberts (@poroping)

// Description: Configure BGP.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpCreate,
		Read:   resourceRouterBgpRead,
		Update: resourceRouterBgpUpdate,
		Delete: resourceRouterBgpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"batchid": {
				Type:        schema.TypeInt,
				Description: "Associate with batch. From 6.4.x+. Currently a WIP and broken.",
				Optional:    true,
				Default:     0,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"additional_path": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable selection of BGP IPv4 additional paths.",
				Optional:    true,
				Computed:    true,
			},
			"additional_path_select": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),

				Description: "Number of additional paths to be selected for each IPv4 NLRI.",
				Optional:    true,
				Computed:    true,
			},
			"additional_path_select6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),

				Description: "Number of additional paths to be selected for each IPv6 NLRI.",
				Optional:    true,
				Computed:    true,
			},
			"additional_path6": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable selection of BGP IPv6 additional paths.",
				Optional:    true,
				Computed:    true,
			},
			"admin_distance": {
				Type:        schema.TypeList,
				Description: "Administrative distance modifications.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Administrative distance to apply (1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"neighbour_prefix": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateIPv4Classnet,

							Description: "Neighbor address prefix.",
							Optional:    true,
							Computed:    true,
						},
						"route_list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Access list of routes to apply new distance to.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"aggregate_address": {
				Type:        schema.TypeList,
				Description: "BGP aggregate address table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable generate AS set path information.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateIPv4ClassnetAny,

							Description: "Aggregate prefix.",
							Optional:    true,
							Computed:    true,
						},
						"summary_only": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable filter more specific routes from updates.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"aggregate_address6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 aggregate address table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_set": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable generate AS set path information.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     fortiValidateIPv6Prefix,
							DiffSuppressFunc: diffCidrEqual,
							Description:      "Aggregate IPv6 prefix.",
							Optional:         true,
							Computed:         true,
						},
						"summary_only": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable filter more specific routes from updates.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"always_compare_med": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable always compare MED.",
				Optional:    true,
				Computed:    true,
			},
			"as": {
				Type: schema.TypeInt,

				Description: "Router AS number, valid from 1 to 4294967295, 0 to disable BGP.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_as_path_ignore": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable ignore AS path.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_cmp_confed_aspath": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable compare federation AS path length.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_cmp_routerid": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable compare router ID for identical EBGP paths.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_med_confed": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable compare MED among confederation paths.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_med_missing_as_worst": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable treat missing MED as least preferred.",
				Optional:    true,
				Computed:    true,
			},
			"client_to_client_reflection": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable client-to-client route reflection.",
				Optional:    true,
				Computed:    true,
			},
			"cluster_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Route reflector cluster ID.",
				Optional:    true,
				Computed:    true,
			},
			"confederation_identifier": {
				Type: schema.TypeInt,

				Description: "Confederation identifier.",
				Optional:    true,
				Computed:    true,
			},
			"confederation_peers": {
				Type:        schema.TypeList,
				Description: "Confederation peers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Peer ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dampening": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable route-flap dampening.",
				Optional:    true,
				Computed:    true,
			},
			"dampening_max_suppress_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Maximum minutes a route can be suppressed.",
				Optional:    true,
				Computed:    true,
			},
			"dampening_reachability_half_life": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 45),

				Description: "Reachability half-life time for penalty (min).",
				Optional:    true,
				Computed:    true,
			},
			"dampening_reuse": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20000),

				Description: "Threshold to reuse routes.",
				Optional:    true,
				Computed:    true,
			},
			"dampening_route_map": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Criteria for dampening.",
				Optional:    true,
				Computed:    true,
			},
			"dampening_suppress": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20000),

				Description: "Threshold to suppress routes.",
				Optional:    true,
				Computed:    true,
			},
			"dampening_unreachability_half_life": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 45),

				Description: "Unreachability half-life time for penalty (min).",
				Optional:    true,
				Computed:    true,
			},
			"default_local_preference": {
				Type: schema.TypeInt,

				Description: "Default local preference.",
				Optional:    true,
				Computed:    true,
			},
			"deterministic_med": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable enforce deterministic comparison of MED.",
				Optional:    true,
				Computed:    true,
			},
			"distance_external": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance for routes external to the AS.",
				Optional:    true,
				Computed:    true,
			},
			"distance_internal": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance for routes internal to the AS.",
				Optional:    true,
				Computed:    true,
			},
			"distance_local": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance for routes local to the AS.",
				Optional:    true,
				Computed:    true,
			},
			"ebgp_multipath": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable EBGP multi-path.",
				Optional:    true,
				Computed:    true,
			},
			"enforce_first_as": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable enforce first AS for EBGP routes.",
				Optional:    true,
				Computed:    true,
			},
			"fast_external_failover": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable reset peer BGP session if link goes down.",
				Optional:    true,
				Computed:    true,
			},
			"graceful_end_on_timer": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable to exit graceful restart on timer only.",
				Optional:    true,
				Computed:    true,
			},
			"graceful_restart": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable BGP graceful restart capabilities.",
				Optional:    true,
				Computed:    true,
			},
			"graceful_restart_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Time needed for neighbors to restart (sec).",
				Optional:    true,
				Computed:    true,
			},
			"graceful_stalepath_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Time to hold stale paths of restarting neighbor (sec).",
				Optional:    true,
				Computed:    true,
			},
			"graceful_update_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Route advertisement/selection delay after restart (sec).",
				Optional:    true,
				Computed:    true,
			},
			"holdtime_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 65535),

				Description: "Number of seconds to mark peer as dead.",
				Optional:    true,
				Computed:    true,
			},
			"ibgp_multipath": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable IBGP multi-path.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_optional_capability": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Don't send unknown optional capability notification message",
				Optional:    true,
				Computed:    true,
			},
			"keepalive_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Frequency to send keep alive requests.",
				Optional:    true,
				Computed:    true,
			},
			"log_neighbour_changes": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable logging of BGP neighbour's changes",
				Optional:    true,
				Computed:    true,
			},
			"multipath_recursive_distance": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable use of recursive distance to select multipath.",
				Optional:    true,
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "BGP neighbor table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable address family IPv4 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"activate6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable address family IPv6 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"send", "receive", "both", "disable"}),

							Description: "Enable/disable IPv4 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"send", "receive", "both", "disable"}),

							Description: "Enable/disable IPv6 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"adv_additional_path": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 255),

							Description: "Number of IPv4 additional paths that can be advertised to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"adv_additional_path6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 255),

							Description: "Number of IPv6 additional paths that can be advertised to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"advertisement_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 600),

							Description: "Minimum interval (sec) between sending updates.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "IPv4 The maximum number of occurrence of my AS number allowed.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_enable": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_enable6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "IPv6 The maximum number of occurrence of my AS number allowed.",
							Optional:    true,
							Computed:    true,
						},
						"as_override": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable replace peer AS with own AS for IPv4.",
							Optional:    true,
							Computed:    true,
						},
						"as_override6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable replace peer AS with own AS for IPv6.",
							Optional:    true,
							Computed:    true,
						},
						"attribute_unchanged": {
							Type: schema.TypeString,

							DiffSuppressFunc: diffFakeListEqual,
							Description:      "IPv4 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"attribute_unchanged6": {
							Type: schema.TypeString,

							DiffSuppressFunc: diffFakeListEqual,
							Description:      "IPv6 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable BFD for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise default IPv4 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise default IPv6 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_dynamic": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise dynamic capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"none", "receive", "send", "both"}),

							Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"none", "receive", "send", "both"}),

							Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_route_refresh": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise route refresh capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"conditional_advertise": {
							Type:        schema.TypeList,
							Description: "Conditional advertisement.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name of advertising route map.",
										Optional:    true,
										Computed:    true,
									},
									"condition_routemap": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name of condition route map.",
										Optional:    true,
										Computed:    true,
									},
									"condition_type": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"exist", "non-exist"}),

										Description: "Type of condition.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"conditional_advertise6": {
							Type:        schema.TypeList,
							Description: "IPv6 conditional advertisement.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name of advertising route map.",
										Optional:    true,
										Computed:    true,
									},
									"condition_routemap": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Name of condition route map.",
										Optional:    true,
										Computed:    true,
									},
									"condition_type": {
										Type:         schema.TypeString,
										ValidateFunc: fortiValidateEnum([]string{"exist", "non-exist"}),

										Description: "Type of condition.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"connect_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Interval (sec) for connect timer.",
							Optional:    true,
							Computed:    true,
						},
						"default_originate_routemap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map to specify criteria to originate IPv4 default.",
							Optional:    true,
							Computed:    true,
						},
						"default_originate_routemap6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map to specify criteria to originate IPv6 default.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv4 updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv6 updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv4 updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv6 updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"dont_capability_negotiate": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Don't negotiate capabilities with this neighbor",
							Optional:    true,
							Computed:    true,
						},
						"ebgp_enforce_multihop": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable allow multi-hop EBGP neighbors.",
							Optional:    true,
							Computed:    true,
						},
						"ebgp_multihop_ttl": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "EBGP multihop TTL for this peer.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv4 inbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv6 inbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv4 outbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv6 outbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"holdtime_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 65535),

							Description: "Interval (sec) before peer considered dead.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),

							Description: "IP/IPv6 address of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"keep_alive_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Keep alive timer interval (sec).",
							Optional:    true,
							Computed:    true,
						},
						"link_down_failover": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable failover upon link down.",
							Optional:    true,
							Computed:    true,
						},
						"local_as": {
							Type: schema.TypeInt,

							Description: "Local AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_no_prepend": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Do not prepend local-as to incoming updates.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_replace_as": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Replace real AS with local-as in outgoing updates.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix": {
							Type: schema.TypeInt,

							Description: "Maximum number of IPv4 prefixes to accept from this peer.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum IPv4 prefix threshold value (1 - 100 percent).",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_threshold6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum IPv6 prefix threshold value (1 - 100 percent).",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 Only give warning message when limit is exceeded.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix6": {
							Type: schema.TypeInt,

							Description: "Maximum number of IPv6 prefixes to accept from this peer.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"override_capability": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable override result of capability negotiation.",
							Optional:    true,
							Computed:    true,
						},
						"passive": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable sending of open messages to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"password": {
							Type: schema.TypeString,

							Description: "Password used in MD5 authentication.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"prefix_list_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Inbound filter for updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Inbound filter for updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Outbound filter for updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Outbound filter for updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"remote_as": {
							Type: schema.TypeInt,

							Description: "AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable remove private AS number from IPv6 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"restart_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),

							Description: "Graceful restart delay time (sec, 0 = global default).",
							Optional:    true,
							Computed:    true,
						},
						"retain_stale_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Time to retain stale routes.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Inbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Inbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 outbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out_preferable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 outbound route map filter if the peer is preferred.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Outbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out6_preferable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 outbound route map filter if the peer is preferred.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"send_community": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"standard", "extended", "both", "disable"}),

							Description: "IPv4 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"send_community6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"standard", "extended", "both", "disable"}),

							Description: "IPv6 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"shutdown": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable shutdown this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"stale_route": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable stale route after neighbor down.",
							Optional:    true,
							Computed:    true,
						},
						"strict_capability_match": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable strict capability matching.",
							Optional:    true,
							Computed:    true,
						},
						"unsuppress_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Route map to selectively unsuppress suppressed routes.",
							Optional:    true,
							Computed:    true,
						},
						"unsuppress_map6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Route map to selectively unsuppress suppressed routes.",
							Optional:    true,
							Computed:    true,
						},
						"update_source": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface to use as source IP/IPv6 address of TCP connections.",
							Optional:    true,
							Computed:    true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Neighbor weight.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"neighbor_group": {
				Type:        schema.TypeList,
				Description: "BGP neighbor group table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activate": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable address family IPv4 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"activate6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable address family IPv6 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"send", "receive", "both", "disable"}),

							Description: "Enable/disable IPv4 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"send", "receive", "both", "disable"}),

							Description: "Enable/disable IPv6 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"adv_additional_path": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 255),

							Description: "Number of IPv4 additional paths that can be advertised to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"adv_additional_path6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 255),

							Description: "Number of IPv6 additional paths that can be advertised to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"advertisement_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 600),

							Description: "Minimum interval (sec) between sending updates.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "IPv4 The maximum number of occurrence of my AS number allowed.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_enable": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_enable6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "IPv6 The maximum number of occurrence of my AS number allowed.",
							Optional:    true,
							Computed:    true,
						},
						"as_override": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable replace peer AS with own AS for IPv4.",
							Optional:    true,
							Computed:    true,
						},
						"as_override6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable replace peer AS with own AS for IPv6.",
							Optional:    true,
							Computed:    true,
						},
						"attribute_unchanged": {
							Type: schema.TypeString,

							DiffSuppressFunc: diffFakeListEqual,
							Description:      "IPv4 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"attribute_unchanged6": {
							Type: schema.TypeString,

							DiffSuppressFunc: diffFakeListEqual,
							Description:      "IPv6 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable BFD for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise default IPv4 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise default IPv6 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_dynamic": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise dynamic capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"none", "receive", "send", "both"}),

							Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"none", "receive", "send", "both"}),

							Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_route_refresh": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable advertise route refresh capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"connect_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Interval (sec) for connect timer.",
							Optional:    true,
							Computed:    true,
						},
						"default_originate_routemap": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map to specify criteria to originate IPv4 default.",
							Optional:    true,
							Computed:    true,
						},
						"default_originate_routemap6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map to specify criteria to originate IPv6 default.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv4 updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv6 updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv4 updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"distribute_list_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for IPv6 updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"dont_capability_negotiate": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Don't negotiate capabilities with this neighbor",
							Optional:    true,
							Computed:    true,
						},
						"ebgp_enforce_multihop": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable allow multi-hop EBGP neighbors.",
							Optional:    true,
							Computed:    true,
						},
						"ebgp_multihop_ttl": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "EBGP multihop TTL for this peer.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv4 inbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv6 inbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv4 outbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "BGP filter for IPv6 outbound routes.",
							Optional:    true,
							Computed:    true,
						},
						"holdtime_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 65535),

							Description: "Interval (sec) before peer considered dead.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Specify outgoing interface for peer connection. For IPv6 peer, the interface should have link-local address.",
							Optional:    true,
							Computed:    true,
						},
						"keep_alive_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Keep alive timer interval (sec).",
							Optional:    true,
							Computed:    true,
						},
						"link_down_failover": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable failover upon link down.",
							Optional:    true,
							Computed:    true,
						},
						"local_as": {
							Type: schema.TypeInt,

							Description: "Local AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_no_prepend": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Do not prepend local-as to incoming updates.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_replace_as": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Replace real AS with local-as in outgoing updates.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix": {
							Type: schema.TypeInt,

							Description: "Maximum number of IPv4 prefixes to accept from this peer.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum IPv4 prefix threshold value (1 - 100 percent).",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_threshold6": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum IPv6 prefix threshold value (1 - 100 percent).",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 Only give warning message when limit is exceeded.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix6": {
							Type: schema.TypeInt,

							Description: "Maximum number of IPv6 prefixes to accept from this peer.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),

							Description: "Neighbor group name.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"override_capability": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable override result of capability negotiation.",
							Optional:    true,
							Computed:    true,
						},
						"passive": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable sending of open messages to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Inbound filter for updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Inbound filter for updates from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Outbound filter for updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_list_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Outbound filter for updates to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"remote_as": {
							Type: schema.TypeInt,

							Description: "AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable remove private AS number from IPv6 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"restart_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),

							Description: "Graceful restart delay time (sec, 0 = global default).",
							Optional:    true,
							Computed:    true,
						},
						"retain_stale_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Time to retain stale routes.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_in": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Inbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_in6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Inbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 outbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out_preferable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 outbound route map filter if the peer is preferred.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Outbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out6_preferable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 outbound route map filter if the peer is preferred.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv4 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable IPv6 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"send_community": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"standard", "extended", "both", "disable"}),

							Description: "IPv4 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"send_community6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnum([]string{"standard", "extended", "both", "disable"}),

							Description: "IPv6 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"shutdown": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable shutdown this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration6": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"stale_route": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable stale route after neighbor down.",
							Optional:    true,
							Computed:    true,
						},
						"strict_capability_match": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable strict capability matching.",
							Optional:    true,
							Computed:    true,
						},
						"unsuppress_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv4 Route map to selectively unsuppress suppressed routes.",
							Optional:    true,
							Computed:    true,
						},
						"unsuppress_map6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "IPv6 Route map to selectively unsuppress suppressed routes.",
							Optional:    true,
							Computed:    true,
						},
						"update_source": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface to use as source IP/IPv6 address of TCP connections.",
							Optional:    true,
							Computed:    true,
						},
						"weight": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Neighbor weight.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"neighbor_range": {
				Type:        schema.TypeList,
				Description: "BGP neighbor range table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Neighbor range ID.",
							Optional:    true,
							Computed:    true,
						},
						"max_neighbor_num": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 1000),

							Description: "Maximum number of neighbors.",
							Optional:    true,
							Computed:    true,
						},
						"neighbor_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Neighbor group name.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateIPv4Classnet,

							Description: "Neighbor range prefix.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"neighbor_range6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 neighbor range table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "IPv6 neighbor range ID.",
							Optional:    true,
							Computed:    true,
						},
						"max_neighbor_num": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 1000),

							Description: "Maximum number of neighbors.",
							Optional:    true,
							Computed:    true,
						},
						"neighbor_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Neighbor group name.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     fortiValidateIPv6Network,
							DiffSuppressFunc: diffCidrEqual,
							Description:      "IPv6 prefix.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "BGP network table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable route as backdoor.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateIPv4Classnet,

							Description: "Network prefix.",
							Optional:    true,
							Computed:    true,
						},
						"route_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map to modify generated route.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"network_import_check": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable ensure BGP network route exists in IGP.",
				Optional:    true,
				Computed:    true,
			},
			"network6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 network table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backdoor": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Enable/disable route as backdoor.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     fortiValidateIPv6Network,
							DiffSuppressFunc: diffCidrEqual,
							Description:      "Network IPv6 prefix.",
							Optional:         true,
							Computed:         true,
						},
						"route_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map to modify generated route.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"recursive_next_hop": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable recursive resolution of next-hop using BGP route.",
				Optional:    true,
				Computed:    true,
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "BGP IPv4 redistribute table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distribute list entry name.",
							Optional:    true,
							Computed:    true,
						},
						"route_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Status",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"redistribute6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 redistribute table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Distribute list entry name.",
							Optional:    true,
							Computed:    true,
						},
						"route_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Route map name.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: fortiValidateEnableDisable(),

							Description: "Status",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"router_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Router ID.",
				Optional:    true,
				Computed:    true,
			},
			"scan_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),

				Description: "Background scanner interval (sec), 0 to disable it.",
				Optional:    true,
				Computed:    true,
			},
			"synchronization": {
				Type:         schema.TypeString,
				ValidateFunc: fortiValidateEnableDisable(),

				Description: "Enable/disable only advertise routes from iBGP if routes present in an IGP.",
				Optional:    true,
				Computed:    true,
			},
			"vrf_leak": {
				Type:        schema.TypeList,
				Description: "BGP VRF leaking table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:        schema.TypeList,
							Description: "Target VRF table.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Interface which is used to leak routes to target VRF.",
										Optional:    true,
										Computed:    true,
									},
									"route_map": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Route map of VRF leaking.",
										Optional:    true,
										Computed:    true,
									},
									"vrf": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 7),

										Description: "Target VRF ID <0 - 31>.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Origin VRF ID <0 - 31>.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vrf_leak6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 VRF leaking table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:        schema.TypeList,
							Description: "Target VRF table.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),

										Description: "Interface which is used to leak routes to target VRF.",
										Optional:    true,
										Computed:    true,
									},
									"route_map": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Route map of VRF leaking.",
										Optional:    true,
										Computed:    true,
									},
									"vrf": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 7),

										Description: "Target VRF ID <0 - 31>.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Origin VRF ID <0 - 31>.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterBgpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectRouterBgp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating RouterBgp resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterBgp(obj, "RouterBgp", vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating RouterBgp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {

	} else {
		d.SetId("RouterBgp")
	}

	return resourceRouterBgpRead(d, m)
}

func resourceRouterBgpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	obj, err := getObjectRouterBgp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating RouterBgp resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterBgp(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating RouterBgp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {

	} else {
		d.SetId("RouterBgp")
	}

	return resourceRouterBgpRead(d, m)
}

func resourceRouterBgpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	err := c.DeleteRouterBgp(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting RouterBgp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1
	urlparams := make(map[string][]string)

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	o, err := c.ReadRouterBgp(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading RouterBgp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading RouterBgp resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelect6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdminDistance(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {

			tmp["distance"] = flattenRouterBgpAdminDistanceDistance(i["distance"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpAdminDistanceId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := i["neighbour-prefix"]; ok {

			tmp["neighbour_prefix"] = flattenRouterBgpAdminDistanceNeighbourPrefix(i["neighbour-prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := i["route-list"]; ok {

			tmp["route_list"] = flattenRouterBgpAdminDistanceRouteList(i["route-list"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpAdminDistanceDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceNeighbourPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func flattenRouterBgpAdminDistanceRouteList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {

			tmp["as_set"] = flattenRouterBgpAggregateAddressAsSet(i["as-set"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpAggregateAddressId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterBgpAggregateAddressPrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {

			tmp["summary_only"] = flattenRouterBgpAggregateAddressSummaryOnly(i["summary-only"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpAggregateAddressAsSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func flattenRouterBgpAggregateAddressSummaryOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {

			tmp["as_set"] = flattenRouterBgpAggregateAddress6AsSet(i["as-set"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpAggregateAddress6Id(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = flattenRouterBgpAggregateAddress6Prefix6(i["prefix6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {

			tmp["summary_only"] = flattenRouterBgpAggregateAddress6SummaryOnly(i["summary-only"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpAggregateAddress6AsSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6SummaryOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathAsPathIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpConfedAspath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathMedConfed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathMedMissingAsWorst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpClientToClientReflection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpClusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpConfederationIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpConfederationPeers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {

			tmp["peer"] = flattenRouterBgpConfederationPeersPeer(i["peer"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "peer", d)
	return result
}

func flattenRouterBgpConfederationPeersPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampening(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningMaxSuppressTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningReuse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningSuppress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDefaultLocalPreference(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDistanceExternal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDistanceInternal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDistanceLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpEbgpMultipath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpFastExternalFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulEndOnTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulStalepathTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulUpdateDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpIbgpMultipath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpIgnoreOptionalCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpKeepaliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpMultipathRecursiveDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {

			tmp["activate"] = flattenRouterBgpNeighborActivate(i["activate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {

			tmp["activate6"] = flattenRouterBgpNeighborActivate6(i["activate6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := i["additional-path"]; ok {

			tmp["additional_path"] = flattenRouterBgpNeighborAdditionalPath(i["additional-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := i["additional-path6"]; ok {

			tmp["additional_path6"] = flattenRouterBgpNeighborAdditionalPath6(i["additional-path6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := i["adv-additional-path"]; ok {

			tmp["adv_additional_path"] = flattenRouterBgpNeighborAdvAdditionalPath(i["adv-additional-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := i["adv-additional-path6"]; ok {

			tmp["adv_additional_path6"] = flattenRouterBgpNeighborAdvAdditionalPath6(i["adv-additional-path6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {

			tmp["advertisement_interval"] = flattenRouterBgpNeighborAdvertisementInterval(i["advertisement-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {

			tmp["allowas_in"] = flattenRouterBgpNeighborAllowasIn(i["allowas-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {

			tmp["allowas_in_enable"] = flattenRouterBgpNeighborAllowasInEnable(i["allowas-in-enable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {

			tmp["allowas_in_enable6"] = flattenRouterBgpNeighborAllowasInEnable6(i["allowas-in-enable6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {

			tmp["allowas_in6"] = flattenRouterBgpNeighborAllowasIn6(i["allowas-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {

			tmp["as_override"] = flattenRouterBgpNeighborAsOverride(i["as-override"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {

			tmp["as_override6"] = flattenRouterBgpNeighborAsOverride6(i["as-override6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {

			tmp["attribute_unchanged"] = flattenRouterBgpNeighborAttributeUnchanged(i["attribute-unchanged"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {

			tmp["attribute_unchanged6"] = flattenRouterBgpNeighborAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {

			tmp["bfd"] = flattenRouterBgpNeighborBfd(i["bfd"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {

			tmp["capability_default_originate"] = flattenRouterBgpNeighborCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {

			tmp["capability_default_originate6"] = flattenRouterBgpNeighborCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {

			tmp["capability_dynamic"] = flattenRouterBgpNeighborCapabilityDynamic(i["capability-dynamic"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := i["capability-graceful-restart"]; ok {

			tmp["capability_graceful_restart"] = flattenRouterBgpNeighborCapabilityGracefulRestart(i["capability-graceful-restart"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := i["capability-graceful-restart6"]; ok {

			tmp["capability_graceful_restart6"] = flattenRouterBgpNeighborCapabilityGracefulRestart6(i["capability-graceful-restart6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {

			tmp["capability_orf"] = flattenRouterBgpNeighborCapabilityOrf(i["capability-orf"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {

			tmp["capability_orf6"] = flattenRouterBgpNeighborCapabilityOrf6(i["capability-orf6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := i["capability-route-refresh"]; ok {

			tmp["capability_route_refresh"] = flattenRouterBgpNeighborCapabilityRouteRefresh(i["capability-route-refresh"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if _, ok := i["conditional-advertise"]; ok {

			tmp["conditional_advertise"] = flattenRouterBgpNeighborConditionalAdvertise(i["conditional-advertise"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise6"
		if _, ok := i["conditional-advertise6"]; ok {

			tmp["conditional_advertise6"] = flattenRouterBgpNeighborConditionalAdvertise6(i["conditional-advertise6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {

			tmp["connect_timer"] = flattenRouterBgpNeighborConnectTimer(i["connect-timer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {

			tmp["default_originate_routemap"] = flattenRouterBgpNeighborDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {

			tmp["default_originate_routemap6"] = flattenRouterBgpNeighborDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenRouterBgpNeighborDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {

			tmp["distribute_list_in"] = flattenRouterBgpNeighborDistributeListIn(i["distribute-list-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {

			tmp["distribute_list_in6"] = flattenRouterBgpNeighborDistributeListIn6(i["distribute-list-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {

			tmp["distribute_list_out"] = flattenRouterBgpNeighborDistributeListOut(i["distribute-list-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {

			tmp["distribute_list_out6"] = flattenRouterBgpNeighborDistributeListOut6(i["distribute-list-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {

			tmp["dont_capability_negotiate"] = flattenRouterBgpNeighborDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {

			tmp["ebgp_enforce_multihop"] = flattenRouterBgpNeighborEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {

			tmp["ebgp_multihop_ttl"] = flattenRouterBgpNeighborEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {

			tmp["filter_list_in"] = flattenRouterBgpNeighborFilterListIn(i["filter-list-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {

			tmp["filter_list_in6"] = flattenRouterBgpNeighborFilterListIn6(i["filter-list-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {

			tmp["filter_list_out"] = flattenRouterBgpNeighborFilterListOut(i["filter-list-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {

			tmp["filter_list_out6"] = flattenRouterBgpNeighborFilterListOut6(i["filter-list-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {

			tmp["holdtime_timer"] = flattenRouterBgpNeighborHoldtimeTimer(i["holdtime-timer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterBgpNeighborInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenRouterBgpNeighborIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {

			tmp["keep_alive_timer"] = flattenRouterBgpNeighborKeepAliveTimer(i["keep-alive-timer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := i["link-down-failover"]; ok {

			tmp["link_down_failover"] = flattenRouterBgpNeighborLinkDownFailover(i["link-down-failover"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := i["local-as"]; ok {

			tmp["local_as"] = flattenRouterBgpNeighborLocalAs(i["local-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := i["local-as-no-prepend"]; ok {

			tmp["local_as_no_prepend"] = flattenRouterBgpNeighborLocalAsNoPrepend(i["local-as-no-prepend"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := i["local-as-replace-as"]; ok {

			tmp["local_as_replace_as"] = flattenRouterBgpNeighborLocalAsReplaceAs(i["local-as-replace-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {

			tmp["maximum_prefix"] = flattenRouterBgpNeighborMaximumPrefix(i["maximum-prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {

			tmp["maximum_prefix_threshold"] = flattenRouterBgpNeighborMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {

			tmp["maximum_prefix_threshold6"] = flattenRouterBgpNeighborMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {

			tmp["maximum_prefix_warning_only"] = flattenRouterBgpNeighborMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {

			tmp["maximum_prefix_warning_only6"] = flattenRouterBgpNeighborMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {

			tmp["maximum_prefix6"] = flattenRouterBgpNeighborMaximumPrefix6(i["maximum-prefix6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {

			tmp["next_hop_self"] = flattenRouterBgpNeighborNextHopSelf(i["next-hop-self"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := i["next-hop-self-rr"]; ok {

			tmp["next_hop_self_rr"] = flattenRouterBgpNeighborNextHopSelfRr(i["next-hop-self-rr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := i["next-hop-self-rr6"]; ok {

			tmp["next_hop_self_rr6"] = flattenRouterBgpNeighborNextHopSelfRr6(i["next-hop-self-rr6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {

			tmp["next_hop_self6"] = flattenRouterBgpNeighborNextHopSelf6(i["next-hop-self6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {

			tmp["override_capability"] = flattenRouterBgpNeighborOverrideCapability(i["override-capability"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {

			tmp["passive"] = flattenRouterBgpNeighborPassive(i["passive"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {

			tmp["prefix_list_in"] = flattenRouterBgpNeighborPrefixListIn(i["prefix-list-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {

			tmp["prefix_list_in6"] = flattenRouterBgpNeighborPrefixListIn6(i["prefix-list-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {

			tmp["prefix_list_out"] = flattenRouterBgpNeighborPrefixListOut(i["prefix-list-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {

			tmp["prefix_list_out6"] = flattenRouterBgpNeighborPrefixListOut6(i["prefix-list-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {

			tmp["remote_as"] = flattenRouterBgpNeighborRemoteAs(i["remote-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {

			tmp["remove_private_as"] = flattenRouterBgpNeighborRemovePrivateAs(i["remove-private-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {

			tmp["remove_private_as6"] = flattenRouterBgpNeighborRemovePrivateAs6(i["remove-private-as6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := i["restart-time"]; ok {

			tmp["restart_time"] = flattenRouterBgpNeighborRestartTime(i["restart-time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := i["retain-stale-time"]; ok {

			tmp["retain_stale_time"] = flattenRouterBgpNeighborRetainStaleTime(i["retain-stale-time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {

			tmp["route_map_in"] = flattenRouterBgpNeighborRouteMapIn(i["route-map-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {

			tmp["route_map_in6"] = flattenRouterBgpNeighborRouteMapIn6(i["route-map-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {

			tmp["route_map_out"] = flattenRouterBgpNeighborRouteMapOut(i["route-map-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := i["route-map-out-preferable"]; ok {

			tmp["route_map_out_preferable"] = flattenRouterBgpNeighborRouteMapOutPreferable(i["route-map-out-preferable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {

			tmp["route_map_out6"] = flattenRouterBgpNeighborRouteMapOut6(i["route-map-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := i["route-map-out6-preferable"]; ok {

			tmp["route_map_out6_preferable"] = flattenRouterBgpNeighborRouteMapOut6Preferable(i["route-map-out6-preferable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {

			tmp["route_reflector_client"] = flattenRouterBgpNeighborRouteReflectorClient(i["route-reflector-client"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {

			tmp["route_reflector_client6"] = flattenRouterBgpNeighborRouteReflectorClient6(i["route-reflector-client6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {

			tmp["route_server_client"] = flattenRouterBgpNeighborRouteServerClient(i["route-server-client"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {

			tmp["route_server_client6"] = flattenRouterBgpNeighborRouteServerClient6(i["route-server-client6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {

			tmp["send_community"] = flattenRouterBgpNeighborSendCommunity(i["send-community"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {

			tmp["send_community6"] = flattenRouterBgpNeighborSendCommunity6(i["send-community6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {

			tmp["shutdown"] = flattenRouterBgpNeighborShutdown(i["shutdown"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {

			tmp["soft_reconfiguration"] = flattenRouterBgpNeighborSoftReconfiguration(i["soft-reconfiguration"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {

			tmp["soft_reconfiguration6"] = flattenRouterBgpNeighborSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := i["stale-route"]; ok {

			tmp["stale_route"] = flattenRouterBgpNeighborStaleRoute(i["stale-route"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {

			tmp["strict_capability_match"] = flattenRouterBgpNeighborStrictCapabilityMatch(i["strict-capability-match"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {

			tmp["unsuppress_map"] = flattenRouterBgpNeighborUnsuppressMap(i["unsuppress-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {

			tmp["unsuppress_map6"] = flattenRouterBgpNeighborUnsuppressMap6(i["unsuppress-map6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {

			tmp["update_source"] = flattenRouterBgpNeighborUpdateSource(i["update-source"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenRouterBgpNeighborWeight(i["weight"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenRouterBgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {

			tmp["advertise_routemap"] = flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(i["advertise-routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {

			tmp["condition_routemap"] = flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(i["condition-routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {

			tmp["condition_type"] = flattenRouterBgpNeighborConditionalAdvertiseConditionType(i["condition-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "advertise-routemap", d)
	return result
}

func flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {

			tmp["advertise_routemap"] = flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(i["advertise-routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {

			tmp["condition_routemap"] = flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(i["condition-routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {

			tmp["condition_type"] = flattenRouterBgpNeighborConditionalAdvertise6ConditionType(i["condition-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "advertise-routemap", d)
	return result
}

func flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborLinkDownFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRetainStaleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborStaleRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {

			tmp["activate"] = flattenRouterBgpNeighborGroupActivate(i["activate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {

			tmp["activate6"] = flattenRouterBgpNeighborGroupActivate6(i["activate6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := i["additional-path"]; ok {

			tmp["additional_path"] = flattenRouterBgpNeighborGroupAdditionalPath(i["additional-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := i["additional-path6"]; ok {

			tmp["additional_path6"] = flattenRouterBgpNeighborGroupAdditionalPath6(i["additional-path6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := i["adv-additional-path"]; ok {

			tmp["adv_additional_path"] = flattenRouterBgpNeighborGroupAdvAdditionalPath(i["adv-additional-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := i["adv-additional-path6"]; ok {

			tmp["adv_additional_path6"] = flattenRouterBgpNeighborGroupAdvAdditionalPath6(i["adv-additional-path6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {

			tmp["advertisement_interval"] = flattenRouterBgpNeighborGroupAdvertisementInterval(i["advertisement-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {

			tmp["allowas_in"] = flattenRouterBgpNeighborGroupAllowasIn(i["allowas-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {

			tmp["allowas_in_enable"] = flattenRouterBgpNeighborGroupAllowasInEnable(i["allowas-in-enable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {

			tmp["allowas_in_enable6"] = flattenRouterBgpNeighborGroupAllowasInEnable6(i["allowas-in-enable6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {

			tmp["allowas_in6"] = flattenRouterBgpNeighborGroupAllowasIn6(i["allowas-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {

			tmp["as_override"] = flattenRouterBgpNeighborGroupAsOverride(i["as-override"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {

			tmp["as_override6"] = flattenRouterBgpNeighborGroupAsOverride6(i["as-override6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {

			tmp["attribute_unchanged"] = flattenRouterBgpNeighborGroupAttributeUnchanged(i["attribute-unchanged"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {

			tmp["attribute_unchanged6"] = flattenRouterBgpNeighborGroupAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {

			tmp["bfd"] = flattenRouterBgpNeighborGroupBfd(i["bfd"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {

			tmp["capability_default_originate"] = flattenRouterBgpNeighborGroupCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {

			tmp["capability_default_originate6"] = flattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {

			tmp["capability_dynamic"] = flattenRouterBgpNeighborGroupCapabilityDynamic(i["capability-dynamic"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := i["capability-graceful-restart"]; ok {

			tmp["capability_graceful_restart"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestart(i["capability-graceful-restart"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := i["capability-graceful-restart6"]; ok {

			tmp["capability_graceful_restart6"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestart6(i["capability-graceful-restart6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {

			tmp["capability_orf"] = flattenRouterBgpNeighborGroupCapabilityOrf(i["capability-orf"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {

			tmp["capability_orf6"] = flattenRouterBgpNeighborGroupCapabilityOrf6(i["capability-orf6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := i["capability-route-refresh"]; ok {

			tmp["capability_route_refresh"] = flattenRouterBgpNeighborGroupCapabilityRouteRefresh(i["capability-route-refresh"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {

			tmp["connect_timer"] = flattenRouterBgpNeighborGroupConnectTimer(i["connect-timer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {

			tmp["default_originate_routemap"] = flattenRouterBgpNeighborGroupDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {

			tmp["default_originate_routemap6"] = flattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenRouterBgpNeighborGroupDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {

			tmp["distribute_list_in"] = flattenRouterBgpNeighborGroupDistributeListIn(i["distribute-list-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {

			tmp["distribute_list_in6"] = flattenRouterBgpNeighborGroupDistributeListIn6(i["distribute-list-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {

			tmp["distribute_list_out"] = flattenRouterBgpNeighborGroupDistributeListOut(i["distribute-list-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {

			tmp["distribute_list_out6"] = flattenRouterBgpNeighborGroupDistributeListOut6(i["distribute-list-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {

			tmp["dont_capability_negotiate"] = flattenRouterBgpNeighborGroupDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {

			tmp["ebgp_enforce_multihop"] = flattenRouterBgpNeighborGroupEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {

			tmp["ebgp_multihop_ttl"] = flattenRouterBgpNeighborGroupEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {

			tmp["filter_list_in"] = flattenRouterBgpNeighborGroupFilterListIn(i["filter-list-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {

			tmp["filter_list_in6"] = flattenRouterBgpNeighborGroupFilterListIn6(i["filter-list-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {

			tmp["filter_list_out"] = flattenRouterBgpNeighborGroupFilterListOut(i["filter-list-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {

			tmp["filter_list_out6"] = flattenRouterBgpNeighborGroupFilterListOut6(i["filter-list-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {

			tmp["holdtime_timer"] = flattenRouterBgpNeighborGroupHoldtimeTimer(i["holdtime-timer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterBgpNeighborGroupInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {

			tmp["keep_alive_timer"] = flattenRouterBgpNeighborGroupKeepAliveTimer(i["keep-alive-timer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := i["link-down-failover"]; ok {

			tmp["link_down_failover"] = flattenRouterBgpNeighborGroupLinkDownFailover(i["link-down-failover"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := i["local-as"]; ok {

			tmp["local_as"] = flattenRouterBgpNeighborGroupLocalAs(i["local-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := i["local-as-no-prepend"]; ok {

			tmp["local_as_no_prepend"] = flattenRouterBgpNeighborGroupLocalAsNoPrepend(i["local-as-no-prepend"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := i["local-as-replace-as"]; ok {

			tmp["local_as_replace_as"] = flattenRouterBgpNeighborGroupLocalAsReplaceAs(i["local-as-replace-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {

			tmp["maximum_prefix"] = flattenRouterBgpNeighborGroupMaximumPrefix(i["maximum-prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {

			tmp["maximum_prefix_threshold"] = flattenRouterBgpNeighborGroupMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {

			tmp["maximum_prefix_threshold6"] = flattenRouterBgpNeighborGroupMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {

			tmp["maximum_prefix_warning_only"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {

			tmp["maximum_prefix_warning_only6"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {

			tmp["maximum_prefix6"] = flattenRouterBgpNeighborGroupMaximumPrefix6(i["maximum-prefix6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenRouterBgpNeighborGroupName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {

			tmp["next_hop_self"] = flattenRouterBgpNeighborGroupNextHopSelf(i["next-hop-self"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := i["next-hop-self-rr"]; ok {

			tmp["next_hop_self_rr"] = flattenRouterBgpNeighborGroupNextHopSelfRr(i["next-hop-self-rr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := i["next-hop-self-rr6"]; ok {

			tmp["next_hop_self_rr6"] = flattenRouterBgpNeighborGroupNextHopSelfRr6(i["next-hop-self-rr6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {

			tmp["next_hop_self6"] = flattenRouterBgpNeighborGroupNextHopSelf6(i["next-hop-self6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {

			tmp["override_capability"] = flattenRouterBgpNeighborGroupOverrideCapability(i["override-capability"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {

			tmp["passive"] = flattenRouterBgpNeighborGroupPassive(i["passive"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {

			tmp["prefix_list_in"] = flattenRouterBgpNeighborGroupPrefixListIn(i["prefix-list-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {

			tmp["prefix_list_in6"] = flattenRouterBgpNeighborGroupPrefixListIn6(i["prefix-list-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {

			tmp["prefix_list_out"] = flattenRouterBgpNeighborGroupPrefixListOut(i["prefix-list-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {

			tmp["prefix_list_out6"] = flattenRouterBgpNeighborGroupPrefixListOut6(i["prefix-list-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {

			tmp["remote_as"] = flattenRouterBgpNeighborGroupRemoteAs(i["remote-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {

			tmp["remove_private_as"] = flattenRouterBgpNeighborGroupRemovePrivateAs(i["remove-private-as"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {

			tmp["remove_private_as6"] = flattenRouterBgpNeighborGroupRemovePrivateAs6(i["remove-private-as6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := i["restart-time"]; ok {

			tmp["restart_time"] = flattenRouterBgpNeighborGroupRestartTime(i["restart-time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := i["retain-stale-time"]; ok {

			tmp["retain_stale_time"] = flattenRouterBgpNeighborGroupRetainStaleTime(i["retain-stale-time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {

			tmp["route_map_in"] = flattenRouterBgpNeighborGroupRouteMapIn(i["route-map-in"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {

			tmp["route_map_in6"] = flattenRouterBgpNeighborGroupRouteMapIn6(i["route-map-in6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {

			tmp["route_map_out"] = flattenRouterBgpNeighborGroupRouteMapOut(i["route-map-out"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := i["route-map-out-preferable"]; ok {

			tmp["route_map_out_preferable"] = flattenRouterBgpNeighborGroupRouteMapOutPreferable(i["route-map-out-preferable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {

			tmp["route_map_out6"] = flattenRouterBgpNeighborGroupRouteMapOut6(i["route-map-out6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := i["route-map-out6-preferable"]; ok {

			tmp["route_map_out6_preferable"] = flattenRouterBgpNeighborGroupRouteMapOut6Preferable(i["route-map-out6-preferable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {

			tmp["route_reflector_client"] = flattenRouterBgpNeighborGroupRouteReflectorClient(i["route-reflector-client"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {

			tmp["route_reflector_client6"] = flattenRouterBgpNeighborGroupRouteReflectorClient6(i["route-reflector-client6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {

			tmp["route_server_client"] = flattenRouterBgpNeighborGroupRouteServerClient(i["route-server-client"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {

			tmp["route_server_client6"] = flattenRouterBgpNeighborGroupRouteServerClient6(i["route-server-client6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {

			tmp["send_community"] = flattenRouterBgpNeighborGroupSendCommunity(i["send-community"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {

			tmp["send_community6"] = flattenRouterBgpNeighborGroupSendCommunity6(i["send-community6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {

			tmp["shutdown"] = flattenRouterBgpNeighborGroupShutdown(i["shutdown"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {

			tmp["soft_reconfiguration"] = flattenRouterBgpNeighborGroupSoftReconfiguration(i["soft-reconfiguration"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {

			tmp["soft_reconfiguration6"] = flattenRouterBgpNeighborGroupSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := i["stale-route"]; ok {

			tmp["stale_route"] = flattenRouterBgpNeighborGroupStaleRoute(i["stale-route"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {

			tmp["strict_capability_match"] = flattenRouterBgpNeighborGroupStrictCapabilityMatch(i["strict-capability-match"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {

			tmp["unsuppress_map"] = flattenRouterBgpNeighborGroupUnsuppressMap(i["unsuppress-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {

			tmp["unsuppress_map6"] = flattenRouterBgpNeighborGroupUnsuppressMap6(i["unsuppress-map6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {

			tmp["update_source"] = flattenRouterBgpNeighborGroupUpdateSource(i["update-source"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenRouterBgpNeighborGroupWeight(i["weight"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterBgpNeighborGroupActivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupConnectTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLinkDownFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupOverrideCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRetainStaleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupShutdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStaleRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUnsuppressMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUpdateSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpNeighborRangeId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := i["max-neighbor-num"]; ok {

			tmp["max_neighbor_num"] = flattenRouterBgpNeighborRangeMaxNeighborNum(i["max-neighbor-num"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := i["neighbor-group"]; ok {

			tmp["neighbor_group"] = flattenRouterBgpNeighborRangeNeighborGroup(i["neighbor-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterBgpNeighborRangePrefix(i["prefix"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNeighborRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangeMaxNeighborNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangeNeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRangePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func flattenRouterBgpNeighborRange6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpNeighborRange6Id(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := i["max-neighbor-num"]; ok {

			tmp["max_neighbor_num"] = flattenRouterBgpNeighborRange6MaxNeighborNum(i["max-neighbor-num"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := i["neighbor-group"]; ok {

			tmp["neighbor_group"] = flattenRouterBgpNeighborRange6NeighborGroup(i["neighbor-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = flattenRouterBgpNeighborRange6Prefix6(i["prefix6"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNeighborRange6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6MaxNeighborNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6NeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {

			tmp["backdoor"] = flattenRouterBgpNetworkBackdoor(i["backdoor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpNetworkId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterBgpNetworkPrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterBgpNetworkRouteMap(i["route-map"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNetworkBackdoor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}
	return v
}

func flattenRouterBgpNetworkRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkImportCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {

			tmp["backdoor"] = flattenRouterBgpNetwork6Backdoor(i["backdoor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterBgpNetwork6Id(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {

			tmp["prefix6"] = flattenRouterBgpNetwork6Prefix6(i["prefix6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterBgpNetwork6RouteMap(i["route-map"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNetwork6Backdoor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6RouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRecursiveNextHop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenRouterBgpRedistributeName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterBgpRedistributeRouteMap(i["route-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenRouterBgpRedistributeStatus(i["status"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterBgpRedistributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenRouterBgpRedistribute6Name(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterBgpRedistribute6RouteMap(i["route-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenRouterBgpRedistribute6Status(i["status"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterBgpRedistribute6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6RouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpScanTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpSynchronization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {

			tmp["target"] = flattenRouterBgpVrfLeakTarget(i["target"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {

			tmp["vrf"] = flattenRouterBgpVrfLeakVrf(i["vrf"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeakTarget(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterBgpVrfLeakTargetInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterBgpVrfLeakTargetRouteMap(i["route-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {

			tmp["vrf"] = flattenRouterBgpVrfLeakTargetVrf(i["vrf"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeakTargetInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {

			tmp["target"] = flattenRouterBgpVrfLeak6Target(i["target"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {

			tmp["vrf"] = flattenRouterBgpVrfLeak6Vrf(i["vrf"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeak6Target(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table
		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterBgpVrfLeak6TargetInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {

			tmp["route_map"] = flattenRouterBgpVrfLeak6TargetRouteMap(i["route-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {

			tmp["vrf"] = flattenRouterBgpVrfLeak6TargetVrf(i["vrf"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeak6TargetInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6TargetRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6TargetVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6Vrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("additional_path", flattenRouterBgpAdditionalPath(o["additional-path"], d, "additional_path", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path_select", flattenRouterBgpAdditionalPathSelect(o["additional-path-select"], d, "additional_path_select", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-select"]) {
			return fmt.Errorf("error reading additional_path_select: %v", err)
		}
	}

	if err = d.Set("additional_path_select6", flattenRouterBgpAdditionalPathSelect6(o["additional-path-select6"], d, "additional_path_select6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-select6"]) {
			return fmt.Errorf("error reading additional_path_select6: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterBgpAdditionalPath6(o["additional-path6"], d, "additional_path6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("error reading additional_path6: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance", sv)); err != nil {
			if !fortiAPIPatch(o["admin-distance"]) {
				return fmt.Errorf("error reading admin_distance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("admin_distance"); ok {
			if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance", sv)); err != nil {
				if !fortiAPIPatch(o["admin-distance"]) {
					return fmt.Errorf("error reading admin_distance: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address", sv)); err != nil {
			if !fortiAPIPatch(o["aggregate-address"]) {
				return fmt.Errorf("error reading aggregate_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address"); ok {
			if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address", sv)); err != nil {
				if !fortiAPIPatch(o["aggregate-address"]) {
					return fmt.Errorf("error reading aggregate_address: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6", sv)); err != nil {
			if !fortiAPIPatch(o["aggregate-address6"]) {
				return fmt.Errorf("error reading aggregate_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address6"); ok {
			if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6", sv)); err != nil {
				if !fortiAPIPatch(o["aggregate-address6"]) {
					return fmt.Errorf("error reading aggregate_address6: %v", err)
				}
			}
		}
	}

	if err = d.Set("always_compare_med", flattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med", sv)); err != nil {
		if !fortiAPIPatch(o["always-compare-med"]) {
			return fmt.Errorf("error reading always_compare_med: %v", err)
		}
	}

	if err = d.Set("as", flattenRouterBgpAs(o["as"], d, "as", sv)); err != nil {
		if !fortiAPIPatch(o["as"]) {
			return fmt.Errorf("error reading as: %v", err)
		}
	}

	if err = d.Set("bestpath_as_path_ignore", flattenRouterBgpBestpathAsPathIgnore(o["bestpath-as-path-ignore"], d, "bestpath_as_path_ignore", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-as-path-ignore"]) {
			return fmt.Errorf("error reading bestpath_as_path_ignore: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_confed_aspath", flattenRouterBgpBestpathCmpConfedAspath(o["bestpath-cmp-confed-aspath"], d, "bestpath_cmp_confed_aspath", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-confed-aspath"]) {
			return fmt.Errorf("error reading bestpath_cmp_confed_aspath: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", flattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-routerid"]) {
			return fmt.Errorf("error reading bestpath_cmp_routerid: %v", err)
		}
	}

	if err = d.Set("bestpath_med_confed", flattenRouterBgpBestpathMedConfed(o["bestpath-med-confed"], d, "bestpath_med_confed", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-med-confed"]) {
			return fmt.Errorf("error reading bestpath_med_confed: %v", err)
		}
	}

	if err = d.Set("bestpath_med_missing_as_worst", flattenRouterBgpBestpathMedMissingAsWorst(o["bestpath-med-missing-as-worst"], d, "bestpath_med_missing_as_worst", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-med-missing-as-worst"]) {
			return fmt.Errorf("error reading bestpath_med_missing_as_worst: %v", err)
		}
	}

	if err = d.Set("client_to_client_reflection", flattenRouterBgpClientToClientReflection(o["client-to-client-reflection"], d, "client_to_client_reflection", sv)); err != nil {
		if !fortiAPIPatch(o["client-to-client-reflection"]) {
			return fmt.Errorf("error reading client_to_client_reflection: %v", err)
		}
	}

	if err = d.Set("cluster_id", flattenRouterBgpClusterId(o["cluster-id"], d, "cluster_id", sv)); err != nil {
		if !fortiAPIPatch(o["cluster-id"]) {
			return fmt.Errorf("error reading cluster_id: %v", err)
		}
	}

	if err = d.Set("confederation_identifier", flattenRouterBgpConfederationIdentifier(o["confederation-identifier"], d, "confederation_identifier", sv)); err != nil {
		if !fortiAPIPatch(o["confederation-identifier"]) {
			return fmt.Errorf("error reading confederation_identifier: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("confederation_peers", flattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers", sv)); err != nil {
			if !fortiAPIPatch(o["confederation-peers"]) {
				return fmt.Errorf("error reading confederation_peers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("confederation_peers"); ok {
			if err = d.Set("confederation_peers", flattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers", sv)); err != nil {
				if !fortiAPIPatch(o["confederation-peers"]) {
					return fmt.Errorf("error reading confederation_peers: %v", err)
				}
			}
		}
	}

	if err = d.Set("dampening", flattenRouterBgpDampening(o["dampening"], d, "dampening", sv)); err != nil {
		if !fortiAPIPatch(o["dampening"]) {
			return fmt.Errorf("error reading dampening: %v", err)
		}
	}

	if err = d.Set("dampening_max_suppress_time", flattenRouterBgpDampeningMaxSuppressTime(o["dampening-max-suppress-time"], d, "dampening_max_suppress_time", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-max-suppress-time"]) {
			return fmt.Errorf("error reading dampening_max_suppress_time: %v", err)
		}
	}

	if err = d.Set("dampening_reachability_half_life", flattenRouterBgpDampeningReachabilityHalfLife(o["dampening-reachability-half-life"], d, "dampening_reachability_half_life", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-reachability-half-life"]) {
			return fmt.Errorf("error reading dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("dampening_reuse", flattenRouterBgpDampeningReuse(o["dampening-reuse"], d, "dampening_reuse", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-reuse"]) {
			return fmt.Errorf("error reading dampening_reuse: %v", err)
		}
	}

	if err = d.Set("dampening_route_map", flattenRouterBgpDampeningRouteMap(o["dampening-route-map"], d, "dampening_route_map", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-route-map"]) {
			return fmt.Errorf("error reading dampening_route_map: %v", err)
		}
	}

	if err = d.Set("dampening_suppress", flattenRouterBgpDampeningSuppress(o["dampening-suppress"], d, "dampening_suppress", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-suppress"]) {
			return fmt.Errorf("error reading dampening_suppress: %v", err)
		}
	}

	if err = d.Set("dampening_unreachability_half_life", flattenRouterBgpDampeningUnreachabilityHalfLife(o["dampening-unreachability-half-life"], d, "dampening_unreachability_half_life", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-unreachability-half-life"]) {
			return fmt.Errorf("error reading dampening_unreachability_half_life: %v", err)
		}
	}

	if err = d.Set("default_local_preference", flattenRouterBgpDefaultLocalPreference(o["default-local-preference"], d, "default_local_preference", sv)); err != nil {
		if !fortiAPIPatch(o["default-local-preference"]) {
			return fmt.Errorf("error reading default_local_preference: %v", err)
		}
	}

	if err = d.Set("deterministic_med", flattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med", sv)); err != nil {
		if !fortiAPIPatch(o["deterministic-med"]) {
			return fmt.Errorf("error reading deterministic_med: %v", err)
		}
	}

	if err = d.Set("distance_external", flattenRouterBgpDistanceExternal(o["distance-external"], d, "distance_external", sv)); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_internal", flattenRouterBgpDistanceInternal(o["distance-internal"], d, "distance_internal", sv)); err != nil {
		if !fortiAPIPatch(o["distance-internal"]) {
			return fmt.Errorf("error reading distance_internal: %v", err)
		}
	}

	if err = d.Set("distance_local", flattenRouterBgpDistanceLocal(o["distance-local"], d, "distance_local", sv)); err != nil {
		if !fortiAPIPatch(o["distance-local"]) {
			return fmt.Errorf("error reading distance_local: %v", err)
		}
	}

	if err = d.Set("ebgp_multipath", flattenRouterBgpEbgpMultipath(o["ebgp-multipath"], d, "ebgp_multipath", sv)); err != nil {
		if !fortiAPIPatch(o["ebgp-multipath"]) {
			return fmt.Errorf("error reading ebgp_multipath: %v", err)
		}
	}

	if err = d.Set("enforce_first_as", flattenRouterBgpEnforceFirstAs(o["enforce-first-as"], d, "enforce_first_as", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-first-as"]) {
			return fmt.Errorf("error reading enforce_first_as: %v", err)
		}
	}

	if err = d.Set("fast_external_failover", flattenRouterBgpFastExternalFailover(o["fast-external-failover"], d, "fast_external_failover", sv)); err != nil {
		if !fortiAPIPatch(o["fast-external-failover"]) {
			return fmt.Errorf("error reading fast_external_failover: %v", err)
		}
	}

	if err = d.Set("graceful_end_on_timer", flattenRouterBgpGracefulEndOnTimer(o["graceful-end-on-timer"], d, "graceful_end_on_timer", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-end-on-timer"]) {
			return fmt.Errorf("error reading graceful_end_on_timer: %v", err)
		}
	}

	if err = d.Set("graceful_restart", flattenRouterBgpGracefulRestart(o["graceful-restart"], d, "graceful_restart", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-restart"]) {
			return fmt.Errorf("error reading graceful_restart: %v", err)
		}
	}

	if err = d.Set("graceful_restart_time", flattenRouterBgpGracefulRestartTime(o["graceful-restart-time"], d, "graceful_restart_time", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-restart-time"]) {
			return fmt.Errorf("error reading graceful_restart_time: %v", err)
		}
	}

	if err = d.Set("graceful_stalepath_time", flattenRouterBgpGracefulStalepathTime(o["graceful-stalepath-time"], d, "graceful_stalepath_time", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-stalepath-time"]) {
			return fmt.Errorf("error reading graceful_stalepath_time: %v", err)
		}
	}

	if err = d.Set("graceful_update_delay", flattenRouterBgpGracefulUpdateDelay(o["graceful-update-delay"], d, "graceful_update_delay", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-update-delay"]) {
			return fmt.Errorf("error reading graceful_update_delay: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterBgpHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer", sv)); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("ibgp_multipath", flattenRouterBgpIbgpMultipath(o["ibgp-multipath"], d, "ibgp_multipath", sv)); err != nil {
		if !fortiAPIPatch(o["ibgp-multipath"]) {
			return fmt.Errorf("error reading ibgp_multipath: %v", err)
		}
	}

	if err = d.Set("ignore_optional_capability", flattenRouterBgpIgnoreOptionalCapability(o["ignore-optional-capability"], d, "ignore_optional_capability", sv)); err != nil {
		if !fortiAPIPatch(o["ignore-optional-capability"]) {
			return fmt.Errorf("error reading ignore_optional_capability: %v", err)
		}
	}

	if err = d.Set("keepalive_timer", flattenRouterBgpKeepaliveTimer(o["keepalive-timer"], d, "keepalive_timer", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive-timer"]) {
			return fmt.Errorf("error reading keepalive_timer: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterBgpLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes", sv)); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("multipath_recursive_distance", flattenRouterBgpMultipathRecursiveDistance(o["multipath-recursive-distance"], d, "multipath_recursive_distance", sv)); err != nil {
		if !fortiAPIPatch(o["multipath-recursive-distance"]) {
			return fmt.Errorf("error reading multipath_recursive_distance: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterBgpNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBgpNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor-group"]) {
				return fmt.Errorf("error reading neighbor_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_group"); ok {
			if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor-group"]) {
					return fmt.Errorf("error reading neighbor_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor-range"]) {
				return fmt.Errorf("error reading neighbor_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_range"); ok {
			if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor-range"]) {
					return fmt.Errorf("error reading neighbor_range: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor-range6"]) {
				return fmt.Errorf("error reading neighbor_range6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_range6"); ok {
			if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor-range6"]) {
					return fmt.Errorf("error reading neighbor_range6: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenRouterBgpNetwork(o["network"], d, "network", sv)); err != nil {
			if !fortiAPIPatch(o["network"]) {
				return fmt.Errorf("error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterBgpNetwork(o["network"], d, "network", sv)); err != nil {
				if !fortiAPIPatch(o["network"]) {
					return fmt.Errorf("error reading network: %v", err)
				}
			}
		}
	}

	if err = d.Set("network_import_check", flattenRouterBgpNetworkImportCheck(o["network-import-check"], d, "network_import_check", sv)); err != nil {
		if !fortiAPIPatch(o["network-import-check"]) {
			return fmt.Errorf("error reading network_import_check: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("network6", flattenRouterBgpNetwork6(o["network6"], d, "network6", sv)); err != nil {
			if !fortiAPIPatch(o["network6"]) {
				return fmt.Errorf("error reading network6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network6"); ok {
			if err = d.Set("network6", flattenRouterBgpNetwork6(o["network6"], d, "network6", sv)); err != nil {
				if !fortiAPIPatch(o["network6"]) {
					return fmt.Errorf("error reading network6: %v", err)
				}
			}
		}
	}

	if err = d.Set("recursive_next_hop", flattenRouterBgpRecursiveNextHop(o["recursive-next-hop"], d, "recursive_next_hop", sv)); err != nil {
		if !fortiAPIPatch(o["recursive-next-hop"]) {
			return fmt.Errorf("error reading recursive_next_hop: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterBgpRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterBgpRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute"]) {
					return fmt.Errorf("error reading redistribute: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute6"]) {
				return fmt.Errorf("error reading redistribute6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute6"); ok {
			if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute6"]) {
					return fmt.Errorf("error reading redistribute6: %v", err)
				}
			}
		}
	}

	if err = d.Set("router_id", flattenRouterBgpRouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("error reading router_id: %v", err)
		}
	}

	if err = d.Set("scan_time", flattenRouterBgpScanTime(o["scan-time"], d, "scan_time", sv)); err != nil {
		if !fortiAPIPatch(o["scan-time"]) {
			return fmt.Errorf("error reading scan_time: %v", err)
		}
	}

	if err = d.Set("synchronization", flattenRouterBgpSynchronization(o["synchronization"], d, "synchronization", sv)); err != nil {
		if !fortiAPIPatch(o["synchronization"]) {
			return fmt.Errorf("error reading synchronization: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak", sv)); err != nil {
			if !fortiAPIPatch(o["vrf-leak"]) {
				return fmt.Errorf("error reading vrf_leak: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf_leak"); ok {
			if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak", sv)); err != nil {
				if !fortiAPIPatch(o["vrf-leak"]) {
					return fmt.Errorf("error reading vrf_leak: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(o["vrf-leak6"], d, "vrf_leak6", sv)); err != nil {
			if !fortiAPIPatch(o["vrf-leak6"]) {
				return fmt.Errorf("error reading vrf_leak6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf_leak6"); ok {
			if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(o["vrf-leak6"], d, "vrf_leak6", sv)); err != nil {
				if !fortiAPIPatch(o["vrf-leak6"]) {
					return fmt.Errorf("error reading vrf_leak6: %v", err)
				}
			}
		}
	}

	return nil
}

func expandRouterBgpAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelect6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distance"], _ = expandRouterBgpAdminDistanceDistance(d, i["distance"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpAdminDistanceId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["neighbour-prefix"], _ = expandRouterBgpAdminDistanceNeighbourPrefix(d, i["neighbour_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-list"], _ = expandRouterBgpAdminDistanceRouteList(d, i["route_list"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpAdminDistanceDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceNeighbourPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceRouteList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["as-set"], _ = expandRouterBgpAggregateAddressAsSet(d, i["as_set"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpAggregateAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterBgpAggregateAddressPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["summary-only"], _ = expandRouterBgpAggregateAddressSummaryOnly(d, i["summary_only"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpAggregateAddressAsSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressSummaryOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["as-set"], _ = expandRouterBgpAggregateAddress6AsSet(d, i["as_set"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpAggregateAddress6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix6"], _ = expandRouterBgpAggregateAddress6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["summary-only"], _ = expandRouterBgpAggregateAddress6SummaryOnly(d, i["summary_only"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpAggregateAddress6AsSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6SummaryOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAlwaysCompareMed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathAsPathIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpConfedAspath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpRouterid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathMedConfed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathMedMissingAsWorst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpClientToClientReflection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpClusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpConfederationIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpConfederationPeers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["peer"], _ = expandRouterBgpConfederationPeersPeer(d, i["peer"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpConfederationPeersPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampening(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningMaxSuppressTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningReachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningReuse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningSuppress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningUnreachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDefaultLocalPreference(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDeterministicMed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceExternal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceInternal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpEbgpMultipath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpEnforceFirstAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpFastExternalFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulEndOnTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulStalepathTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulUpdateDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpIbgpMultipath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpIgnoreOptionalCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpKeepaliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpMultipathRecursiveDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["activate"], _ = expandRouterBgpNeighborActivate(d, i["activate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["activate6"], _ = expandRouterBgpNeighborActivate6(d, i["activate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["additional-path"], _ = expandRouterBgpNeighborAdditionalPath(d, i["additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["additional-path6"], _ = expandRouterBgpNeighborAdditionalPath6(d, i["additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["adv-additional-path"], _ = expandRouterBgpNeighborAdvAdditionalPath(d, i["adv_additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["adv-additional-path6"], _ = expandRouterBgpNeighborAdvAdditionalPath6(d, i["adv_additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["advertisement-interval"], _ = expandRouterBgpNeighborAdvertisementInterval(d, i["advertisement_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in"], _ = expandRouterBgpNeighborAllowasIn(d, i["allowas_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in-enable"], _ = expandRouterBgpNeighborAllowasInEnable(d, i["allowas_in_enable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in-enable6"], _ = expandRouterBgpNeighborAllowasInEnable6(d, i["allowas_in_enable6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in6"], _ = expandRouterBgpNeighborAllowasIn6(d, i["allowas_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["as-override"], _ = expandRouterBgpNeighborAsOverride(d, i["as_override"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["as-override6"], _ = expandRouterBgpNeighborAsOverride6(d, i["as_override6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["attribute-unchanged"], _ = expandRouterBgpNeighborAttributeUnchanged(d, i["attribute_unchanged"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["attribute-unchanged6"], _ = expandRouterBgpNeighborAttributeUnchanged6(d, i["attribute_unchanged6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd"], _ = expandRouterBgpNeighborBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-default-originate"], _ = expandRouterBgpNeighborCapabilityDefaultOriginate(d, i["capability_default_originate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-default-originate6"], _ = expandRouterBgpNeighborCapabilityDefaultOriginate6(d, i["capability_default_originate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-dynamic"], _ = expandRouterBgpNeighborCapabilityDynamic(d, i["capability_dynamic"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-graceful-restart"], _ = expandRouterBgpNeighborCapabilityGracefulRestart(d, i["capability_graceful_restart"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-graceful-restart6"], _ = expandRouterBgpNeighborCapabilityGracefulRestart6(d, i["capability_graceful_restart6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-orf"], _ = expandRouterBgpNeighborCapabilityOrf(d, i["capability_orf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-orf6"], _ = expandRouterBgpNeighborCapabilityOrf6(d, i["capability_orf6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-route-refresh"], _ = expandRouterBgpNeighborCapabilityRouteRefresh(d, i["capability_route_refresh"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["conditional-advertise"], _ = expandRouterBgpNeighborConditionalAdvertise(d, i["conditional_advertise"], pre_append, sv)
		} else {
			tmp["conditional-advertise"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["conditional-advertise6"], _ = expandRouterBgpNeighborConditionalAdvertise6(d, i["conditional_advertise6"], pre_append, sv)
		} else {
			tmp["conditional-advertise6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["connect-timer"], _ = expandRouterBgpNeighborConnectTimer(d, i["connect_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default-originate-routemap"], _ = expandRouterBgpNeighborDefaultOriginateRoutemap(d, i["default_originate_routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default-originate-routemap6"], _ = expandRouterBgpNeighborDefaultOriginateRoutemap6(d, i["default_originate_routemap6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandRouterBgpNeighborDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-in"], _ = expandRouterBgpNeighborDistributeListIn(d, i["distribute_list_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-in6"], _ = expandRouterBgpNeighborDistributeListIn6(d, i["distribute_list_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-out"], _ = expandRouterBgpNeighborDistributeListOut(d, i["distribute_list_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-out6"], _ = expandRouterBgpNeighborDistributeListOut6(d, i["distribute_list_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dont-capability-negotiate"], _ = expandRouterBgpNeighborDontCapabilityNegotiate(d, i["dont_capability_negotiate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ebgp-enforce-multihop"], _ = expandRouterBgpNeighborEbgpEnforceMultihop(d, i["ebgp_enforce_multihop"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ebgp-multihop-ttl"], _ = expandRouterBgpNeighborEbgpMultihopTtl(d, i["ebgp_multihop_ttl"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-in"], _ = expandRouterBgpNeighborFilterListIn(d, i["filter_list_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-in6"], _ = expandRouterBgpNeighborFilterListIn6(d, i["filter_list_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-out"], _ = expandRouterBgpNeighborFilterListOut(d, i["filter_list_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-out6"], _ = expandRouterBgpNeighborFilterListOut6(d, i["filter_list_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["holdtime-timer"], _ = expandRouterBgpNeighborHoldtimeTimer(d, i["holdtime_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterBgpNeighborInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandRouterBgpNeighborIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["keep-alive-timer"], _ = expandRouterBgpNeighborKeepAliveTimer(d, i["keep_alive_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["link-down-failover"], _ = expandRouterBgpNeighborLinkDownFailover(d, i["link_down_failover"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["local-as"], _ = expandRouterBgpNeighborLocalAs(d, i["local_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["local-as-no-prepend"], _ = expandRouterBgpNeighborLocalAsNoPrepend(d, i["local_as_no_prepend"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["local-as-replace-as"], _ = expandRouterBgpNeighborLocalAsReplaceAs(d, i["local_as_replace_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix"], _ = expandRouterBgpNeighborMaximumPrefix(d, i["maximum_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-threshold"], _ = expandRouterBgpNeighborMaximumPrefixThreshold(d, i["maximum_prefix_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-threshold6"], _ = expandRouterBgpNeighborMaximumPrefixThreshold6(d, i["maximum_prefix_threshold6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-warning-only"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnly(d, i["maximum_prefix_warning_only"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-warning-only6"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnly6(d, i["maximum_prefix_warning_only6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix6"], _ = expandRouterBgpNeighborMaximumPrefix6(d, i["maximum_prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self"], _ = expandRouterBgpNeighborNextHopSelf(d, i["next_hop_self"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self-rr"], _ = expandRouterBgpNeighborNextHopSelfRr(d, i["next_hop_self_rr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self-rr6"], _ = expandRouterBgpNeighborNextHopSelfRr6(d, i["next_hop_self_rr6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self6"], _ = expandRouterBgpNeighborNextHopSelf6(d, i["next_hop_self6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["override-capability"], _ = expandRouterBgpNeighborOverrideCapability(d, i["override_capability"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["passive"], _ = expandRouterBgpNeighborPassive(d, i["passive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["password"], _ = expandRouterBgpNeighborPassword(d, i["password"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-in"], _ = expandRouterBgpNeighborPrefixListIn(d, i["prefix_list_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-in6"], _ = expandRouterBgpNeighborPrefixListIn6(d, i["prefix_list_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-out"], _ = expandRouterBgpNeighborPrefixListOut(d, i["prefix_list_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-out6"], _ = expandRouterBgpNeighborPrefixListOut6(d, i["prefix_list_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remote-as"], _ = expandRouterBgpNeighborRemoteAs(d, i["remote_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remove-private-as"], _ = expandRouterBgpNeighborRemovePrivateAs(d, i["remove_private_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remove-private-as6"], _ = expandRouterBgpNeighborRemovePrivateAs6(d, i["remove_private_as6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["restart-time"], _ = expandRouterBgpNeighborRestartTime(d, i["restart_time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["retain-stale-time"], _ = expandRouterBgpNeighborRetainStaleTime(d, i["retain_stale_time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-in"], _ = expandRouterBgpNeighborRouteMapIn(d, i["route_map_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-in6"], _ = expandRouterBgpNeighborRouteMapIn6(d, i["route_map_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out"], _ = expandRouterBgpNeighborRouteMapOut(d, i["route_map_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out-preferable"], _ = expandRouterBgpNeighborRouteMapOutPreferable(d, i["route_map_out_preferable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out6"], _ = expandRouterBgpNeighborRouteMapOut6(d, i["route_map_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out6-preferable"], _ = expandRouterBgpNeighborRouteMapOut6Preferable(d, i["route_map_out6_preferable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-reflector-client"], _ = expandRouterBgpNeighborRouteReflectorClient(d, i["route_reflector_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-reflector-client6"], _ = expandRouterBgpNeighborRouteReflectorClient6(d, i["route_reflector_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-server-client"], _ = expandRouterBgpNeighborRouteServerClient(d, i["route_server_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-server-client6"], _ = expandRouterBgpNeighborRouteServerClient6(d, i["route_server_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-community"], _ = expandRouterBgpNeighborSendCommunity(d, i["send_community"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-community6"], _ = expandRouterBgpNeighborSendCommunity6(d, i["send_community6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["shutdown"], _ = expandRouterBgpNeighborShutdown(d, i["shutdown"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["soft-reconfiguration"], _ = expandRouterBgpNeighborSoftReconfiguration(d, i["soft_reconfiguration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["soft-reconfiguration6"], _ = expandRouterBgpNeighborSoftReconfiguration6(d, i["soft_reconfiguration6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["stale-route"], _ = expandRouterBgpNeighborStaleRoute(d, i["stale_route"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["strict-capability-match"], _ = expandRouterBgpNeighborStrictCapabilityMatch(d, i["strict_capability_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["unsuppress-map"], _ = expandRouterBgpNeighborUnsuppressMap(d, i["unsuppress_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["unsuppress-map6"], _ = expandRouterBgpNeighborUnsuppressMap6(d, i["unsuppress_map6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["update-source"], _ = expandRouterBgpNeighborUpdateSource(d, i["update_source"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandRouterBgpNeighborWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborActivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(d, i["advertise_routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d, i["condition_routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionType(d, i["condition_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(d, i["advertise_routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d, i["condition_routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionType(d, i["condition_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConnectTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLinkDownFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborOverrideCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRetainStaleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborShutdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStaleRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUnsuppressMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUpdateSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["activate"], _ = expandRouterBgpNeighborGroupActivate(d, i["activate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["activate6"], _ = expandRouterBgpNeighborGroupActivate6(d, i["activate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["additional-path"], _ = expandRouterBgpNeighborGroupAdditionalPath(d, i["additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["additional-path6"], _ = expandRouterBgpNeighborGroupAdditionalPath6(d, i["additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["adv-additional-path"], _ = expandRouterBgpNeighborGroupAdvAdditionalPath(d, i["adv_additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["adv-additional-path6"], _ = expandRouterBgpNeighborGroupAdvAdditionalPath6(d, i["adv_additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["advertisement-interval"], _ = expandRouterBgpNeighborGroupAdvertisementInterval(d, i["advertisement_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in"], _ = expandRouterBgpNeighborGroupAllowasIn(d, i["allowas_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in-enable"], _ = expandRouterBgpNeighborGroupAllowasInEnable(d, i["allowas_in_enable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in-enable6"], _ = expandRouterBgpNeighborGroupAllowasInEnable6(d, i["allowas_in_enable6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowas-in6"], _ = expandRouterBgpNeighborGroupAllowasIn6(d, i["allowas_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["as-override"], _ = expandRouterBgpNeighborGroupAsOverride(d, i["as_override"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["as-override6"], _ = expandRouterBgpNeighborGroupAsOverride6(d, i["as_override6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["attribute-unchanged"], _ = expandRouterBgpNeighborGroupAttributeUnchanged(d, i["attribute_unchanged"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["attribute-unchanged6"], _ = expandRouterBgpNeighborGroupAttributeUnchanged6(d, i["attribute_unchanged6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd"], _ = expandRouterBgpNeighborGroupBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-default-originate"], _ = expandRouterBgpNeighborGroupCapabilityDefaultOriginate(d, i["capability_default_originate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-default-originate6"], _ = expandRouterBgpNeighborGroupCapabilityDefaultOriginate6(d, i["capability_default_originate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-dynamic"], _ = expandRouterBgpNeighborGroupCapabilityDynamic(d, i["capability_dynamic"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-graceful-restart"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestart(d, i["capability_graceful_restart"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-graceful-restart6"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestart6(d, i["capability_graceful_restart6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-orf"], _ = expandRouterBgpNeighborGroupCapabilityOrf(d, i["capability_orf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-orf6"], _ = expandRouterBgpNeighborGroupCapabilityOrf6(d, i["capability_orf6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["capability-route-refresh"], _ = expandRouterBgpNeighborGroupCapabilityRouteRefresh(d, i["capability_route_refresh"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["connect-timer"], _ = expandRouterBgpNeighborGroupConnectTimer(d, i["connect_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default-originate-routemap"], _ = expandRouterBgpNeighborGroupDefaultOriginateRoutemap(d, i["default_originate_routemap"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default-originate-routemap6"], _ = expandRouterBgpNeighborGroupDefaultOriginateRoutemap6(d, i["default_originate_routemap6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandRouterBgpNeighborGroupDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-in"], _ = expandRouterBgpNeighborGroupDistributeListIn(d, i["distribute_list_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-in6"], _ = expandRouterBgpNeighborGroupDistributeListIn6(d, i["distribute_list_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-out"], _ = expandRouterBgpNeighborGroupDistributeListOut(d, i["distribute_list_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distribute-list-out6"], _ = expandRouterBgpNeighborGroupDistributeListOut6(d, i["distribute_list_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dont-capability-negotiate"], _ = expandRouterBgpNeighborGroupDontCapabilityNegotiate(d, i["dont_capability_negotiate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ebgp-enforce-multihop"], _ = expandRouterBgpNeighborGroupEbgpEnforceMultihop(d, i["ebgp_enforce_multihop"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ebgp-multihop-ttl"], _ = expandRouterBgpNeighborGroupEbgpMultihopTtl(d, i["ebgp_multihop_ttl"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-in"], _ = expandRouterBgpNeighborGroupFilterListIn(d, i["filter_list_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-in6"], _ = expandRouterBgpNeighborGroupFilterListIn6(d, i["filter_list_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-out"], _ = expandRouterBgpNeighborGroupFilterListOut(d, i["filter_list_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-list-out6"], _ = expandRouterBgpNeighborGroupFilterListOut6(d, i["filter_list_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["holdtime-timer"], _ = expandRouterBgpNeighborGroupHoldtimeTimer(d, i["holdtime_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterBgpNeighborGroupInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["keep-alive-timer"], _ = expandRouterBgpNeighborGroupKeepAliveTimer(d, i["keep_alive_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["link-down-failover"], _ = expandRouterBgpNeighborGroupLinkDownFailover(d, i["link_down_failover"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["local-as"], _ = expandRouterBgpNeighborGroupLocalAs(d, i["local_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["local-as-no-prepend"], _ = expandRouterBgpNeighborGroupLocalAsNoPrepend(d, i["local_as_no_prepend"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["local-as-replace-as"], _ = expandRouterBgpNeighborGroupLocalAsReplaceAs(d, i["local_as_replace_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix"], _ = expandRouterBgpNeighborGroupMaximumPrefix(d, i["maximum_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-threshold"], _ = expandRouterBgpNeighborGroupMaximumPrefixThreshold(d, i["maximum_prefix_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-threshold6"], _ = expandRouterBgpNeighborGroupMaximumPrefixThreshold6(d, i["maximum_prefix_threshold6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-warning-only"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnly(d, i["maximum_prefix_warning_only"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix-warning-only6"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnly6(d, i["maximum_prefix_warning_only6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["maximum-prefix6"], _ = expandRouterBgpNeighborGroupMaximumPrefix6(d, i["maximum_prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterBgpNeighborGroupName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self"], _ = expandRouterBgpNeighborGroupNextHopSelf(d, i["next_hop_self"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self-rr"], _ = expandRouterBgpNeighborGroupNextHopSelfRr(d, i["next_hop_self_rr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self-rr6"], _ = expandRouterBgpNeighborGroupNextHopSelfRr6(d, i["next_hop_self_rr6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["next-hop-self6"], _ = expandRouterBgpNeighborGroupNextHopSelf6(d, i["next_hop_self6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["override-capability"], _ = expandRouterBgpNeighborGroupOverrideCapability(d, i["override_capability"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["passive"], _ = expandRouterBgpNeighborGroupPassive(d, i["passive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-in"], _ = expandRouterBgpNeighborGroupPrefixListIn(d, i["prefix_list_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-in6"], _ = expandRouterBgpNeighborGroupPrefixListIn6(d, i["prefix_list_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-out"], _ = expandRouterBgpNeighborGroupPrefixListOut(d, i["prefix_list_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix-list-out6"], _ = expandRouterBgpNeighborGroupPrefixListOut6(d, i["prefix_list_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remote-as"], _ = expandRouterBgpNeighborGroupRemoteAs(d, i["remote_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remove-private-as"], _ = expandRouterBgpNeighborGroupRemovePrivateAs(d, i["remove_private_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["remove-private-as6"], _ = expandRouterBgpNeighborGroupRemovePrivateAs6(d, i["remove_private_as6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["restart-time"], _ = expandRouterBgpNeighborGroupRestartTime(d, i["restart_time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["retain-stale-time"], _ = expandRouterBgpNeighborGroupRetainStaleTime(d, i["retain_stale_time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-in"], _ = expandRouterBgpNeighborGroupRouteMapIn(d, i["route_map_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-in6"], _ = expandRouterBgpNeighborGroupRouteMapIn6(d, i["route_map_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out"], _ = expandRouterBgpNeighborGroupRouteMapOut(d, i["route_map_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutPreferable(d, i["route_map_out_preferable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out6"], _ = expandRouterBgpNeighborGroupRouteMapOut6(d, i["route_map_out6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map-out6-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOut6Preferable(d, i["route_map_out6_preferable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-reflector-client"], _ = expandRouterBgpNeighborGroupRouteReflectorClient(d, i["route_reflector_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-reflector-client6"], _ = expandRouterBgpNeighborGroupRouteReflectorClient6(d, i["route_reflector_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-server-client"], _ = expandRouterBgpNeighborGroupRouteServerClient(d, i["route_server_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-server-client6"], _ = expandRouterBgpNeighborGroupRouteServerClient6(d, i["route_server_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-community"], _ = expandRouterBgpNeighborGroupSendCommunity(d, i["send_community"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-community6"], _ = expandRouterBgpNeighborGroupSendCommunity6(d, i["send_community6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["shutdown"], _ = expandRouterBgpNeighborGroupShutdown(d, i["shutdown"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["soft-reconfiguration"], _ = expandRouterBgpNeighborGroupSoftReconfiguration(d, i["soft_reconfiguration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["soft-reconfiguration6"], _ = expandRouterBgpNeighborGroupSoftReconfiguration6(d, i["soft_reconfiguration6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["stale-route"], _ = expandRouterBgpNeighborGroupStaleRoute(d, i["stale_route"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["strict-capability-match"], _ = expandRouterBgpNeighborGroupStrictCapabilityMatch(d, i["strict_capability_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["unsuppress-map"], _ = expandRouterBgpNeighborGroupUnsuppressMap(d, i["unsuppress_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["unsuppress-map6"], _ = expandRouterBgpNeighborGroupUnsuppressMap6(d, i["unsuppress_map6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["update-source"], _ = expandRouterBgpNeighborGroupUpdateSource(d, i["update_source"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandRouterBgpNeighborGroupWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborGroupActivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupConnectTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLinkDownFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupOverrideCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRetainStaleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupShutdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStaleRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUnsuppressMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUpdateSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpNeighborRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max-neighbor-num"], _ = expandRouterBgpNeighborRangeMaxNeighborNum(d, i["max_neighbor_num"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["neighbor-group"], _ = expandRouterBgpNeighborRangeNeighborGroup(d, i["neighbor_group"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterBgpNeighborRangePrefix(d, i["prefix"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeMaxNeighborNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpNeighborRange6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max-neighbor-num"], _ = expandRouterBgpNeighborRange6MaxNeighborNum(d, i["max_neighbor_num"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["neighbor-group"], _ = expandRouterBgpNeighborRange6NeighborGroup(d, i["neighbor_group"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix6"], _ = expandRouterBgpNeighborRange6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborRange6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6MaxNeighborNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6NeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["backdoor"], _ = expandRouterBgpNetworkBackdoor(d, i["backdoor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpNetworkId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterBgpNetworkPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterBgpNetworkRouteMap(d, i["route_map"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNetworkBackdoor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkImportCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["backdoor"], _ = expandRouterBgpNetwork6Backdoor(d, i["backdoor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterBgpNetwork6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix6"], _ = expandRouterBgpNetwork6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterBgpNetwork6RouteMap(d, i["route_map"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNetwork6Backdoor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6RouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRecursiveNextHop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterBgpRedistributeName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterBgpRedistributeRouteMap(d, i["route_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandRouterBgpRedistributeStatus(d, i["status"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpRedistributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterBgpRedistribute6Name(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterBgpRedistribute6RouteMap(d, i["route_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandRouterBgpRedistribute6Status(d, i["status"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpRedistribute6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6RouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpScanTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpSynchronization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["target"], _ = expandRouterBgpVrfLeakTarget(d, i["target"], pre_append, sv)
		} else {
			tmp["target"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vrf"], _ = expandRouterBgpVrfLeakVrf(d, i["vrf"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterBgpVrfLeakTargetInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterBgpVrfLeakTargetRouteMap(d, i["route_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vrf"], _ = expandRouterBgpVrfLeakTargetVrf(d, i["vrf"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTargetInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["target"], _ = expandRouterBgpVrfLeak6Target(d, i["target"], pre_append, sv)
		} else {
			tmp["target"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vrf"], _ = expandRouterBgpVrfLeak6Vrf(d, i["vrf"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeak6Target(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterBgpVrfLeak6TargetInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-map"], _ = expandRouterBgpVrfLeak6TargetRouteMap(d, i["route_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vrf"], _ = expandRouterBgpVrfLeak6TargetVrf(d, i["vrf"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeak6TargetInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6TargetRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6TargetVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6Vrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("additional_path"); ok {
		t, err := expandRouterBgpAdditionalPath(d, v, "additional_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_select"); ok {
		t, err := expandRouterBgpAdditionalPathSelect(d, v, "additional_path_select", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-select"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_select6"); ok {
		t, err := expandRouterBgpAdditionalPathSelect6(d, v, "additional_path_select6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-select6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok {
		t, err := expandRouterBgpAdditionalPath6(d, v, "additional_path6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("admin_distance"); ok {
		t, err := expandRouterBgpAdminDistance(d, v, "admin_distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-distance"] = t
		}
	} else if d.HasChange("admin_distance") {
		old, new := d.GetChange("admin_distance")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["admin-distance"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("aggregate_address"); ok {
		t, err := expandRouterBgpAggregateAddress(d, v, "aggregate_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-address"] = t
		}
	} else if d.HasChange("aggregate_address") {
		old, new := d.GetChange("aggregate_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["aggregate-address"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("aggregate_address6"); ok {
		t, err := expandRouterBgpAggregateAddress6(d, v, "aggregate_address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-address6"] = t
		}
	} else if d.HasChange("aggregate_address6") {
		old, new := d.GetChange("aggregate_address6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["aggregate-address6"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("always_compare_med"); ok {
		t, err := expandRouterBgpAlwaysCompareMed(d, v, "always_compare_med", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["always-compare-med"] = t
		}
	}

	if v, ok := d.GetOk("as"); ok {
		t, err := expandRouterBgpAs(d, v, "as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_as_path_ignore"); ok {
		t, err := expandRouterBgpBestpathAsPathIgnore(d, v, "bestpath_as_path_ignore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-as-path-ignore"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_confed_aspath"); ok {
		t, err := expandRouterBgpBestpathCmpConfedAspath(d, v, "bestpath_cmp_confed_aspath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-cmp-confed-aspath"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_routerid"); ok {
		t, err := expandRouterBgpBestpathCmpRouterid(d, v, "bestpath_cmp_routerid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-cmp-routerid"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_med_confed"); ok {
		t, err := expandRouterBgpBestpathMedConfed(d, v, "bestpath_med_confed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-med-confed"] = t
		}
	}

	if v, ok := d.GetOk("bestpath_med_missing_as_worst"); ok {
		t, err := expandRouterBgpBestpathMedMissingAsWorst(d, v, "bestpath_med_missing_as_worst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bestpath-med-missing-as-worst"] = t
		}
	}

	if v, ok := d.GetOk("client_to_client_reflection"); ok {
		t, err := expandRouterBgpClientToClientReflection(d, v, "client_to_client_reflection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-to-client-reflection"] = t
		}
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		t, err := expandRouterBgpClusterId(d, v, "cluster_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cluster-id"] = t
		}
	}

	if v, ok := d.GetOk("confederation_identifier"); ok {
		t, err := expandRouterBgpConfederationIdentifier(d, v, "confederation_identifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confederation-identifier"] = t
		}
	}

	if v, ok := d.GetOk("confederation_peers"); ok {
		t, err := expandRouterBgpConfederationPeers(d, v, "confederation_peers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confederation-peers"] = t
		}
	} else if d.HasChange("confederation_peers") {
		old, new := d.GetChange("confederation_peers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["confederation-peers"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("dampening"); ok {
		t, err := expandRouterBgpDampening(d, v, "dampening", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening"] = t
		}
	}

	if v, ok := d.GetOk("dampening_max_suppress_time"); ok {
		t, err := expandRouterBgpDampeningMaxSuppressTime(d, v, "dampening_max_suppress_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-max-suppress-time"] = t
		}
	}

	if v, ok := d.GetOk("dampening_reachability_half_life"); ok {
		t, err := expandRouterBgpDampeningReachabilityHalfLife(d, v, "dampening_reachability_half_life", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-reachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("dampening_reuse"); ok {
		t, err := expandRouterBgpDampeningReuse(d, v, "dampening_reuse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-reuse"] = t
		}
	}

	if v, ok := d.GetOk("dampening_route_map"); ok {
		t, err := expandRouterBgpDampeningRouteMap(d, v, "dampening_route_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-route-map"] = t
		}
	}

	if v, ok := d.GetOk("dampening_suppress"); ok {
		t, err := expandRouterBgpDampeningSuppress(d, v, "dampening_suppress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-suppress"] = t
		}
	}

	if v, ok := d.GetOk("dampening_unreachability_half_life"); ok {
		t, err := expandRouterBgpDampeningUnreachabilityHalfLife(d, v, "dampening_unreachability_half_life", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dampening-unreachability-half-life"] = t
		}
	}

	if v, ok := d.GetOk("default_local_preference"); ok {
		t, err := expandRouterBgpDefaultLocalPreference(d, v, "default_local_preference", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-local-preference"] = t
		}
	}

	if v, ok := d.GetOk("deterministic_med"); ok {
		t, err := expandRouterBgpDeterministicMed(d, v, "deterministic_med", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deterministic-med"] = t
		}
	}

	if v, ok := d.GetOk("distance_external"); ok {
		t, err := expandRouterBgpDistanceExternal(d, v, "distance_external", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-external"] = t
		}
	}

	if v, ok := d.GetOk("distance_internal"); ok {
		t, err := expandRouterBgpDistanceInternal(d, v, "distance_internal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-internal"] = t
		}
	}

	if v, ok := d.GetOk("distance_local"); ok {
		t, err := expandRouterBgpDistanceLocal(d, v, "distance_local", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-local"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_multipath"); ok {
		t, err := expandRouterBgpEbgpMultipath(d, v, "ebgp_multipath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-multipath"] = t
		}
	}

	if v, ok := d.GetOk("enforce_first_as"); ok {
		t, err := expandRouterBgpEnforceFirstAs(d, v, "enforce_first_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-first-as"] = t
		}
	}

	if v, ok := d.GetOk("fast_external_failover"); ok {
		t, err := expandRouterBgpFastExternalFailover(d, v, "fast_external_failover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-external-failover"] = t
		}
	}

	if v, ok := d.GetOk("graceful_end_on_timer"); ok {
		t, err := expandRouterBgpGracefulEndOnTimer(d, v, "graceful_end_on_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-end-on-timer"] = t
		}
	}

	if v, ok := d.GetOk("graceful_restart"); ok {
		t, err := expandRouterBgpGracefulRestart(d, v, "graceful_restart", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-restart"] = t
		}
	}

	if v, ok := d.GetOk("graceful_restart_time"); ok {
		t, err := expandRouterBgpGracefulRestartTime(d, v, "graceful_restart_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-restart-time"] = t
		}
	}

	if v, ok := d.GetOk("graceful_stalepath_time"); ok {
		t, err := expandRouterBgpGracefulStalepathTime(d, v, "graceful_stalepath_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-stalepath-time"] = t
		}
	}

	if v, ok := d.GetOk("graceful_update_delay"); ok {
		t, err := expandRouterBgpGracefulUpdateDelay(d, v, "graceful_update_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graceful-update-delay"] = t
		}
	}

	if v, ok := d.GetOk("holdtime_timer"); ok {
		t, err := expandRouterBgpHoldtimeTimer(d, v, "holdtime_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holdtime-timer"] = t
		}
	}

	if v, ok := d.GetOk("ibgp_multipath"); ok {
		t, err := expandRouterBgpIbgpMultipath(d, v, "ibgp_multipath", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibgp-multipath"] = t
		}
	}

	if v, ok := d.GetOk("ignore_optional_capability"); ok {
		t, err := expandRouterBgpIgnoreOptionalCapability(d, v, "ignore_optional_capability", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-optional-capability"] = t
		}
	}

	if v, ok := d.GetOk("keepalive_timer"); ok {
		t, err := expandRouterBgpKeepaliveTimer(d, v, "keepalive_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive-timer"] = t
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		t, err := expandRouterBgpLogNeighbourChanges(d, v, "log_neighbour_changes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-neighbour-changes"] = t
		}
	}

	if v, ok := d.GetOk("multipath_recursive_distance"); ok {
		t, err := expandRouterBgpMultipathRecursiveDistance(d, v, "multipath_recursive_distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipath-recursive-distance"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok {
		t, err := expandRouterBgpNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["neighbor"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("neighbor_group"); ok {
		t, err := expandRouterBgpNeighborGroup(d, v, "neighbor_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-group"] = t
		}
	} else if d.HasChange("neighbor_group") {
		old, new := d.GetChange("neighbor_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["neighbor-group"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("neighbor_range"); ok {
		t, err := expandRouterBgpNeighborRange(d, v, "neighbor_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-range"] = t
		}
	} else if d.HasChange("neighbor_range") {
		old, new := d.GetChange("neighbor_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["neighbor-range"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("neighbor_range6"); ok {
		t, err := expandRouterBgpNeighborRange6(d, v, "neighbor_range6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-range6"] = t
		}
	} else if d.HasChange("neighbor_range6") {
		old, new := d.GetChange("neighbor_range6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["neighbor-range6"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("network"); ok {
		t, err := expandRouterBgpNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	} else if d.HasChange("network") {
		old, new := d.GetChange("network")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["network"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("network_import_check"); ok {
		t, err := expandRouterBgpNetworkImportCheck(d, v, "network_import_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-import-check"] = t
		}
	}

	if v, ok := d.GetOk("network6"); ok {
		t, err := expandRouterBgpNetwork6(d, v, "network6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network6"] = t
		}
	} else if d.HasChange("network6") {
		old, new := d.GetChange("network6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["network6"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("recursive_next_hop"); ok {
		t, err := expandRouterBgpRecursiveNextHop(d, v, "recursive_next_hop", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recursive-next-hop"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok {
		t, err := expandRouterBgpRedistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["redistribute"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("redistribute6"); ok {
		t, err := expandRouterBgpRedistribute6(d, v, "redistribute6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6"] = t
		}
	} else if d.HasChange("redistribute6") {
		old, new := d.GetChange("redistribute6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["redistribute6"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		t, err := expandRouterBgpRouterId(d, v, "router_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("scan_time"); ok {
		t, err := expandRouterBgpScanTime(d, v, "scan_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-time"] = t
		}
	}

	if v, ok := d.GetOk("synchronization"); ok {
		t, err := expandRouterBgpSynchronization(d, v, "synchronization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synchronization"] = t
		}
	}

	if v, ok := d.GetOk("vrf_leak"); ok {
		t, err := expandRouterBgpVrfLeak(d, v, "vrf_leak", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-leak"] = t
		}
	} else if d.HasChange("vrf_leak") {
		old, new := d.GetChange("vrf_leak")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["vrf-leak"] = make([]struct{}, 0)
		}
	}

	if v, ok := d.GetOk("vrf_leak6"); ok {
		t, err := expandRouterBgpVrfLeak6(d, v, "vrf_leak6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-leak6"] = t
		}
	} else if d.HasChange("vrf_leak6") {
		old, new := d.GetChange("vrf_leak6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj["vrf-leak6"] = make([]struct{}, 0)
		}
	}

	return &obj, nil
}
