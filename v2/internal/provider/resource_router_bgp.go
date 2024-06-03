// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure BGP.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure BGP.",

		CreateContext: resourceRouterBgpCreate,
		ReadContext:   resourceRouterBgpRead,
		UpdateContext: resourceRouterBgpUpdate,
		DeleteContext: resourceRouterBgpDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"additional_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
			"additional_path_select_vpnv4": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),

				Description: "Number of additional paths to be selected for each VPNv4 NLRI.",
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
			"additional_path_vpnv4": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable selection of BGP VPNv4 additional paths.",
				Optional:    true,
				Computed:    true,
			},
			"additional_path6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validators.FortiValidateIPv4Classnet,

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

							Description: "Aggregate prefix.",
							Optional:    true,
							Computed:    true,
						},
						"summary_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "Aggregate IPv6 prefix.",
							Optional:         true,
							Computed:         true,
						},
						"summary_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable filter more specific routes from updates.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"always_compare_med": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable always compare MED.",
				Optional:    true,
				Computed:    true,
			},
			"as": {
				Type: schema.TypeString,

				Description: "Router AS number, asplain/asdot/asdot+ format, 0 to disable BGP.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_as_path_ignore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignore AS path.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_cmp_confed_aspath": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable compare federation AS path length.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_cmp_routerid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable compare router ID for identical EBGP paths.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_med_confed": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable compare MED among confederation paths.",
				Optional:    true,
				Computed:    true,
			},
			"bestpath_med_missing_as_worst": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable treat missing MED as least preferred.",
				Optional:    true,
				Computed:    true,
			},
			"client_to_client_reflection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EBGP multi-path.",
				Optional:    true,
				Computed:    true,
			},
			"enforce_first_as": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable enforce first AS for EBGP routes.",
				Optional:    true,
				Computed:    true,
			},
			"fast_external_failover": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable reset peer BGP session if link goes down.",
				Optional:    true,
				Computed:    true,
			},
			"graceful_end_on_timer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable to exit graceful restart on timer only.",
				Optional:    true,
				Computed:    true,
			},
			"graceful_restart": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IBGP multi-path.",
				Optional:    true,
				Computed:    true,
			},
			"ignore_optional_capability": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Do not send unknown optional capability notification message.",
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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Log BGP neighbor changes.",
				Optional:    true,
				Computed:    true,
			},
			"multipath_recursive_distance": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address family IPv4 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"activate_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address family VPNv4 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"activate6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address family IPv6 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"send", "receive", "both", "disable"}, false),

							Description: "Enable/disable IPv4 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"send", "receive", "both", "disable"}, false),

							Description: "Enable/disable VPNv4 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"send", "receive", "both", "disable"}, false),

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
						"adv_additional_path_vpnv4": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 255),

							Description: "Number of VPNv4 additional paths that can be advertised to this neighbor.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_enable6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_vpnv4": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "The maximum number of occurrence of my AS number allowed for VPNv4 route.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable replace peer AS with own AS for IPv4.",
							Optional:    true,
							Computed:    true,
						},
						"as_override6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable replace peer AS with own AS for IPv6.",
							Optional:    true,
							Computed:    true,
						},
						"attribute_unchanged": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "IPv4 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"attribute_unchanged_vpnv4": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "List of attributes that should be unchanged for VPNv4 route.",
							Optional:         true,
							Computed:         true,
						},
						"attribute_unchanged6": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "IPv6 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable BFD for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise default IPv4 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise default IPv6 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_dynamic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise dynamic capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise VPNv4 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "receive", "send", "both"}, false),

							Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "receive", "send", "both"}, false),

							Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_route_refresh": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
										Type:        schema.TypeList,
										Description: "List of conditional route maps.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Route map.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"condition_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"exist", "non-exist"}, false),

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
										Type:        schema.TypeList,
										Description: "List of conditional route maps.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Route map.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"condition_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"exist", "non-exist"}, false),

										Description: "Type of condition.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"connect_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

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
						"distribute_list_in_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for VPNv4 updates from this neighbor.",
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
						"distribute_list_out_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for VPNv4 updates to this neighbor.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Do not negotiate capabilities with this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"ebgp_enforce_multihop": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable failover upon link down.",
							Optional:    true,
							Computed:    true,
						},
						"local_as": {
							Type: schema.TypeString,

							Description: "Local AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_no_prepend": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Do not prepend local-as to incoming updates.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_replace_as": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"maximum_prefix_threshold_vpnv4": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum VPNv4 prefix threshold value (1 - 100 percent).",
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
						"maximum_prefix_vpnv4": {
							Type: schema.TypeInt,

							Description: "Maximum number of VPNv4 prefixes to accept from this peer.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"override_capability": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable override result of capability negotiation.",
							Optional:    true,
							Computed:    true,
						},
						"passive": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"prefix_list_in_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Inbound filter for VPNv4 updates from this neighbor.",
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
						"prefix_list_out_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Outbound filter for VPNv4 updates to this neighbor.",
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
							Type: schema.TypeString,

							Description: "AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable remove private AS number from VPNv4 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"route_map_in_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VPNv4 inbound route map filter.",
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
						"route_map_out_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VPNv4 outbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out_vpnv4_preferable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VPNv4 outbound route map filter if the peer is preferred.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable VPNv4 AS route reflector client for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable VPNv4 AS route server client for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"send_community": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standard", "extended", "both", "disable"}, false),

							Description: "IPv4 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"send_community_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standard", "extended", "both", "disable"}, false),

							Description: "Send community attribute to neighbor for VPNv4 address family.",
							Optional:    true,
							Computed:    true,
						},
						"send_community6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standard", "extended", "both", "disable"}, false),

							Description: "IPv6 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"shutdown": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable shutdown this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allow VPNv4 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"stale_route": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable stale route after neighbor down.",
							Optional:    true,
							Computed:    true,
						},
						"strict_capability_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address family IPv4 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"activate_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address family VPNv4 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"activate6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable address family IPv6 for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"send", "receive", "both", "disable"}, false),

							Description: "Enable/disable IPv4 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"send", "receive", "both", "disable"}, false),

							Description: "Enable/disable VPNv4 additional-path capability.",
							Optional:    true,
							Computed:    true,
						},
						"additional_path6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"send", "receive", "both", "disable"}, false),

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
						"adv_additional_path_vpnv4": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 255),

							Description: "Number of VPNv4 additional paths that can be advertised to this neighbor.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_enable6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 Enable to allow my AS in AS path.",
							Optional:    true,
							Computed:    true,
						},
						"allowas_in_vpnv4": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),

							Description: "The maximum number of occurrence of my AS number allowed for VPNv4 route.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable replace peer AS with own AS for IPv4.",
							Optional:    true,
							Computed:    true,
						},
						"as_override6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable replace peer AS with own AS for IPv6.",
							Optional:    true,
							Computed:    true,
						},
						"attribute_unchanged": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "IPv4 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"attribute_unchanged_vpnv4": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "List of attributes that should be unchanged for VPNv4 route.",
							Optional:         true,
							Computed:         true,
						},
						"attribute_unchanged6": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "IPv6 List of attributes that should be unchanged.",
							Optional:         true,
							Computed:         true,
						},
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable BFD for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise default IPv4 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_default_originate6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise default IPv6 route to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_dynamic": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise dynamic capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise IPv4 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise VPNv4 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_graceful_restart6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise IPv6 graceful restart capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "receive", "send", "both"}, false),

							Description: "Accept/Send IPv4 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_orf6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "receive", "send", "both"}, false),

							Description: "Accept/Send IPv6 ORF lists to/from this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"capability_route_refresh": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable advertise route refresh capability to this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"connect_timer": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

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
						"distribute_list_in_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for VPNv4 updates from this neighbor.",
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
						"distribute_list_out_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for VPNv4 updates to this neighbor.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Do not negotiate capabilities with this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"ebgp_enforce_multihop": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable failover upon link down.",
							Optional:    true,
							Computed:    true,
						},
						"local_as": {
							Type: schema.TypeString,

							Description: "Local AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_no_prepend": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Do not prepend local-as to incoming updates.",
							Optional:    true,
							Computed:    true,
						},
						"local_as_replace_as": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"maximum_prefix_threshold_vpnv4": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Maximum VPNv4 prefix threshold value (1 - 100 percent).",
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
						"maximum_prefix_vpnv4": {
							Type: schema.TypeInt,

							Description: "Maximum number of VPNv4 prefixes to accept from this peer.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 Only give warning message when limit is exceeded.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable only giving warning message when limit is exceeded for VPNv4 routes.",
							Optional:    true,
							Computed:    true,
						},
						"maximum_prefix_warning_only6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_rr6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"next_hop_self6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 next-hop calculation for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"override_capability": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable override result of capability negotiation.",
							Optional:    true,
							Computed:    true,
						},
						"passive": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"prefix_list_in_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Inbound filter for VPNv4 updates from this neighbor.",
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
						"prefix_list_out_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Outbound filter for VPNv4 updates to this neighbor.",
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
							Type: schema.TypeString,

							Description: "AS number of neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable remove private AS number from VPNv4 outbound updates.",
							Optional:    true,
							Computed:    true,
						},
						"remove_private_as6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"route_map_in_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VPNv4 inbound route map filter.",
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
						"route_map_out_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VPNv4 outbound route map filter.",
							Optional:    true,
							Computed:    true,
						},
						"route_map_out_vpnv4_preferable": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "VPNv4 outbound route map filter if the peer is preferred.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable VPNv4 AS route reflector client for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"route_reflector_client6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 AS route reflector client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv4 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable VPNv4 AS route server client for this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"route_server_client6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable IPv6 AS route server client.",
							Optional:    true,
							Computed:    true,
						},
						"send_community": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standard", "extended", "both", "disable"}, false),

							Description: "IPv4 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"send_community_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standard", "extended", "both", "disable"}, false),

							Description: "Send community attribute to neighbor for VPNv4 address family.",
							Optional:    true,
							Computed:    true,
						},
						"send_community6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standard", "extended", "both", "disable"}, false),

							Description: "IPv6 Send community attribute to neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"shutdown": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable shutdown this neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allow IPv4 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration_vpnv4": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allow VPNv4 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"soft_reconfiguration6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allow IPv6 inbound soft reconfiguration.",
							Optional:    true,
							Computed:    true,
						},
						"stale_route": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable stale route after neighbor down.",
							Optional:    true,
							Computed:    true,
						},
						"strict_capability_match": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validators.FortiValidateIPv4Classnet,

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
							ValidateFunc:     validators.FortiValidateIPv6Network,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"network_import_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

							Description: "Configure insurance of BGP network route existence in IGP.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
						"network_import_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

							Description: "Configure insurance of BGP network route existence in IGP.",
							Optional:    true,
							Computed:    true,
						},
						"prefix6": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Network,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
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
			"recursive_inherit_priority": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable priority inheritance for recursive resolution.",
				Optional:    true,
				Computed:    true,
			},
			"recursive_next_hop": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
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
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Status.",
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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable only advertise routes from iBGP if routes present in an IGP.",
				Optional:    true,
				Computed:    true,
			},
			"tag_resolve_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "preferred", "merge"}, false),

				Description: "Configure tag-match mode. Resolves BGP routes with other routes containing the same tag.",
				Optional:    true,
				Computed:    true,
			},
			"vrf": {
				Type:        schema.TypeList,
				Description: "BGP VRF leaking table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"export_rt": {
							Type:        schema.TypeList,
							Description: "List of export route target.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Attribute: AA|AA:NN.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"import_route_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Import route map.",
							Optional:    true,
							Computed:    true,
						},
						"import_rt": {
							Type:        schema.TypeList,
							Description: "List of import route target.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Attribute: AA|AA:NN.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"leak_target": {
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

										Description: "Target VRF ID (0 - 251).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"rd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Route Distinguisher: AA|AA:NN.",
							Optional:    true,
							Computed:    true,
						},
						"role": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"standalone", "ce", "pe"}, false),

							Description: "VRF role.",
							Optional:    true,
							Computed:    true,
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Origin VRF ID (0 - 251).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
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

										Description: "Target VRF ID (0 - 31).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Origin VRF ID (0 - 31).",
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

										Description: "Target VRF ID (0 - 31).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Origin VRF ID (0 - 31).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vrf6": {
				Type:        schema.TypeList,
				Description: "BGP IPv6 VRF leaking table.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leak_target": {
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

										Description: "Target VRF ID (0 - 251).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"vrf": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Origin VRF ID (0 - 251).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	obj, diags := getObjectRouterBgp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgp")
	}

	return resourceRouterBgpRead(ctx, d, meta)
}

func resourceRouterBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectRouterBgp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgp")
	}

	return resourceRouterBgpRead(ctx, d, meta)
}

func resourceRouterBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	obj, diags := getEmptyObjectRouterBgp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterBgp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadRouterBgp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenRouterBgpAdminDistance(d *schema.ResourceData, v *[]models.RouterBgpAdminDistance, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Distance; tmp != nil {
				v["distance"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.NeighbourPrefix; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.neighbour_prefix", prefix, i), *tmp); tmp != nil {
					v["neighbour_prefix"] = *tmp
				}
			}

			if tmp := cfg.RouteList; tmp != nil {
				v["route_list"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpAggregateAddress(d *schema.ResourceData, v *[]models.RouterBgpAggregateAddress, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AsSet; tmp != nil {
				v["as_set"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.prefix", prefix, i), *tmp); tmp != nil {
					v["prefix"] = *tmp
				}
			}

			if tmp := cfg.SummaryOnly; tmp != nil {
				v["summary_only"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpAggregateAddress6(d *schema.ResourceData, v *[]models.RouterBgpAggregateAddress6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AsSet; tmp != nil {
				v["as_set"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			if tmp := cfg.SummaryOnly; tmp != nil {
				v["summary_only"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpConfederationPeers(d *schema.ResourceData, v *[]models.RouterBgpConfederationPeers, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Peer; tmp != nil {
				v["peer"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "peer")
	}

	return flat
}

func flattenRouterBgpNeighbor(d *schema.ResourceData, v *[]models.RouterBgpNeighbor, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Activate; tmp != nil {
				v["activate"] = *tmp
			}

			if tmp := cfg.ActivateVpnv4; tmp != nil {
				v["activate_vpnv4"] = *tmp
			}

			if tmp := cfg.Activate6; tmp != nil {
				v["activate6"] = *tmp
			}

			if tmp := cfg.AdditionalPath; tmp != nil {
				v["additional_path"] = *tmp
			}

			if tmp := cfg.AdditionalPathVpnv4; tmp != nil {
				v["additional_path_vpnv4"] = *tmp
			}

			if tmp := cfg.AdditionalPath6; tmp != nil {
				v["additional_path6"] = *tmp
			}

			if tmp := cfg.AdvAdditionalPath; tmp != nil {
				v["adv_additional_path"] = *tmp
			}

			if tmp := cfg.AdvAdditionalPathVpnv4; tmp != nil {
				v["adv_additional_path_vpnv4"] = *tmp
			}

			if tmp := cfg.AdvAdditionalPath6; tmp != nil {
				v["adv_additional_path6"] = *tmp
			}

			if tmp := cfg.AdvertisementInterval; tmp != nil {
				v["advertisement_interval"] = *tmp
			}

			if tmp := cfg.AllowasIn; tmp != nil {
				v["allowas_in"] = *tmp
			}

			if tmp := cfg.AllowasInEnable; tmp != nil {
				v["allowas_in_enable"] = *tmp
			}

			if tmp := cfg.AllowasInEnable6; tmp != nil {
				v["allowas_in_enable6"] = *tmp
			}

			if tmp := cfg.AllowasInVpnv4; tmp != nil {
				v["allowas_in_vpnv4"] = *tmp
			}

			if tmp := cfg.AllowasIn6; tmp != nil {
				v["allowas_in6"] = *tmp
			}

			if tmp := cfg.AsOverride; tmp != nil {
				v["as_override"] = *tmp
			}

			if tmp := cfg.AsOverride6; tmp != nil {
				v["as_override6"] = *tmp
			}

			if tmp := cfg.AttributeUnchanged; tmp != nil {
				v["attribute_unchanged"] = *tmp
			}

			if tmp := cfg.AttributeUnchangedVpnv4; tmp != nil {
				v["attribute_unchanged_vpnv4"] = *tmp
			}

			if tmp := cfg.AttributeUnchanged6; tmp != nil {
				v["attribute_unchanged6"] = *tmp
			}

			if tmp := cfg.Bfd; tmp != nil {
				v["bfd"] = *tmp
			}

			if tmp := cfg.CapabilityDefaultOriginate; tmp != nil {
				v["capability_default_originate"] = *tmp
			}

			if tmp := cfg.CapabilityDefaultOriginate6; tmp != nil {
				v["capability_default_originate6"] = *tmp
			}

			if tmp := cfg.CapabilityDynamic; tmp != nil {
				v["capability_dynamic"] = *tmp
			}

			if tmp := cfg.CapabilityGracefulRestart; tmp != nil {
				v["capability_graceful_restart"] = *tmp
			}

			if tmp := cfg.CapabilityGracefulRestartVpnv4; tmp != nil {
				v["capability_graceful_restart_vpnv4"] = *tmp
			}

			if tmp := cfg.CapabilityGracefulRestart6; tmp != nil {
				v["capability_graceful_restart6"] = *tmp
			}

			if tmp := cfg.CapabilityOrf; tmp != nil {
				v["capability_orf"] = *tmp
			}

			if tmp := cfg.CapabilityOrf6; tmp != nil {
				v["capability_orf6"] = *tmp
			}

			if tmp := cfg.CapabilityRouteRefresh; tmp != nil {
				v["capability_route_refresh"] = *tmp
			}

			if tmp := cfg.ConditionalAdvertise; tmp != nil {
				v["conditional_advertise"] = flattenRouterBgpNeighborConditionalAdvertise(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "conditional_advertise"), sort)
			}

			if tmp := cfg.ConditionalAdvertise6; tmp != nil {
				v["conditional_advertise6"] = flattenRouterBgpNeighborConditionalAdvertise6(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "conditional_advertise6"), sort)
			}

			if tmp := cfg.ConnectTimer; tmp != nil {
				v["connect_timer"] = *tmp
			}

			if tmp := cfg.DefaultOriginateRoutemap; tmp != nil {
				v["default_originate_routemap"] = *tmp
			}

			if tmp := cfg.DefaultOriginateRoutemap6; tmp != nil {
				v["default_originate_routemap6"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.DistributeListIn; tmp != nil {
				v["distribute_list_in"] = *tmp
			}

			if tmp := cfg.DistributeListInVpnv4; tmp != nil {
				v["distribute_list_in_vpnv4"] = *tmp
			}

			if tmp := cfg.DistributeListIn6; tmp != nil {
				v["distribute_list_in6"] = *tmp
			}

			if tmp := cfg.DistributeListOut; tmp != nil {
				v["distribute_list_out"] = *tmp
			}

			if tmp := cfg.DistributeListOutVpnv4; tmp != nil {
				v["distribute_list_out_vpnv4"] = *tmp
			}

			if tmp := cfg.DistributeListOut6; tmp != nil {
				v["distribute_list_out6"] = *tmp
			}

			if tmp := cfg.DontCapabilityNegotiate; tmp != nil {
				v["dont_capability_negotiate"] = *tmp
			}

			if tmp := cfg.EbgpEnforceMultihop; tmp != nil {
				v["ebgp_enforce_multihop"] = *tmp
			}

			if tmp := cfg.EbgpMultihopTtl; tmp != nil {
				v["ebgp_multihop_ttl"] = *tmp
			}

			if tmp := cfg.FilterListIn; tmp != nil {
				v["filter_list_in"] = *tmp
			}

			if tmp := cfg.FilterListIn6; tmp != nil {
				v["filter_list_in6"] = *tmp
			}

			if tmp := cfg.FilterListOut; tmp != nil {
				v["filter_list_out"] = *tmp
			}

			if tmp := cfg.FilterListOut6; tmp != nil {
				v["filter_list_out6"] = *tmp
			}

			if tmp := cfg.HoldtimeTimer; tmp != nil {
				v["holdtime_timer"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.KeepAliveTimer; tmp != nil {
				v["keep_alive_timer"] = *tmp
			}

			if tmp := cfg.LinkDownFailover; tmp != nil {
				v["link_down_failover"] = *tmp
			}

			if tmp := cfg.LocalAs; tmp != nil {
				v["local_as"] = *tmp
			}

			if tmp := cfg.LocalAsNoPrepend; tmp != nil {
				v["local_as_no_prepend"] = *tmp
			}

			if tmp := cfg.LocalAsReplaceAs; tmp != nil {
				v["local_as_replace_as"] = *tmp
			}

			if tmp := cfg.MaximumPrefix; tmp != nil {
				v["maximum_prefix"] = *tmp
			}

			if tmp := cfg.MaximumPrefixThreshold; tmp != nil {
				v["maximum_prefix_threshold"] = *tmp
			}

			if tmp := cfg.MaximumPrefixThresholdVpnv4; tmp != nil {
				v["maximum_prefix_threshold_vpnv4"] = *tmp
			}

			if tmp := cfg.MaximumPrefixThreshold6; tmp != nil {
				v["maximum_prefix_threshold6"] = *tmp
			}

			if tmp := cfg.MaximumPrefixVpnv4; tmp != nil {
				v["maximum_prefix_vpnv4"] = *tmp
			}

			if tmp := cfg.MaximumPrefixWarningOnly; tmp != nil {
				v["maximum_prefix_warning_only"] = *tmp
			}

			if tmp := cfg.MaximumPrefixWarningOnlyVpnv4; tmp != nil {
				v["maximum_prefix_warning_only_vpnv4"] = *tmp
			}

			if tmp := cfg.MaximumPrefixWarningOnly6; tmp != nil {
				v["maximum_prefix_warning_only6"] = *tmp
			}

			if tmp := cfg.MaximumPrefix6; tmp != nil {
				v["maximum_prefix6"] = *tmp
			}

			if tmp := cfg.NextHopSelf; tmp != nil {
				v["next_hop_self"] = *tmp
			}

			if tmp := cfg.NextHopSelfRr; tmp != nil {
				v["next_hop_self_rr"] = *tmp
			}

			if tmp := cfg.NextHopSelfRr6; tmp != nil {
				v["next_hop_self_rr6"] = *tmp
			}

			if tmp := cfg.NextHopSelfVpnv4; tmp != nil {
				v["next_hop_self_vpnv4"] = *tmp
			}

			if tmp := cfg.NextHopSelf6; tmp != nil {
				v["next_hop_self6"] = *tmp
			}

			if tmp := cfg.OverrideCapability; tmp != nil {
				v["override_capability"] = *tmp
			}

			if tmp := cfg.Passive; tmp != nil {
				v["passive"] = *tmp
			}

			if tmp := cfg.Password; tmp != nil {
				v["password"] = *tmp
			}

			if tmp := cfg.PrefixListIn; tmp != nil {
				v["prefix_list_in"] = *tmp
			}

			if tmp := cfg.PrefixListInVpnv4; tmp != nil {
				v["prefix_list_in_vpnv4"] = *tmp
			}

			if tmp := cfg.PrefixListIn6; tmp != nil {
				v["prefix_list_in6"] = *tmp
			}

			if tmp := cfg.PrefixListOut; tmp != nil {
				v["prefix_list_out"] = *tmp
			}

			if tmp := cfg.PrefixListOutVpnv4; tmp != nil {
				v["prefix_list_out_vpnv4"] = *tmp
			}

			if tmp := cfg.PrefixListOut6; tmp != nil {
				v["prefix_list_out6"] = *tmp
			}

			if tmp := cfg.RemoteAs; tmp != nil {
				v["remote_as"] = *tmp
			}

			if tmp := cfg.RemovePrivateAs; tmp != nil {
				v["remove_private_as"] = *tmp
			}

			if tmp := cfg.RemovePrivateAsVpnv4; tmp != nil {
				v["remove_private_as_vpnv4"] = *tmp
			}

			if tmp := cfg.RemovePrivateAs6; tmp != nil {
				v["remove_private_as6"] = *tmp
			}

			if tmp := cfg.RestartTime; tmp != nil {
				v["restart_time"] = *tmp
			}

			if tmp := cfg.RetainStaleTime; tmp != nil {
				v["retain_stale_time"] = *tmp
			}

			if tmp := cfg.RouteMapIn; tmp != nil {
				v["route_map_in"] = *tmp
			}

			if tmp := cfg.RouteMapInVpnv4; tmp != nil {
				v["route_map_in_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteMapIn6; tmp != nil {
				v["route_map_in6"] = *tmp
			}

			if tmp := cfg.RouteMapOut; tmp != nil {
				v["route_map_out"] = *tmp
			}

			if tmp := cfg.RouteMapOutPreferable; tmp != nil {
				v["route_map_out_preferable"] = *tmp
			}

			if tmp := cfg.RouteMapOutVpnv4; tmp != nil {
				v["route_map_out_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteMapOutVpnv4Preferable; tmp != nil {
				v["route_map_out_vpnv4_preferable"] = *tmp
			}

			if tmp := cfg.RouteMapOut6; tmp != nil {
				v["route_map_out6"] = *tmp
			}

			if tmp := cfg.RouteMapOut6Preferable; tmp != nil {
				v["route_map_out6_preferable"] = *tmp
			}

			if tmp := cfg.RouteReflectorClient; tmp != nil {
				v["route_reflector_client"] = *tmp
			}

			if tmp := cfg.RouteReflectorClientVpnv4; tmp != nil {
				v["route_reflector_client_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteReflectorClient6; tmp != nil {
				v["route_reflector_client6"] = *tmp
			}

			if tmp := cfg.RouteServerClient; tmp != nil {
				v["route_server_client"] = *tmp
			}

			if tmp := cfg.RouteServerClientVpnv4; tmp != nil {
				v["route_server_client_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteServerClient6; tmp != nil {
				v["route_server_client6"] = *tmp
			}

			if tmp := cfg.SendCommunity; tmp != nil {
				v["send_community"] = *tmp
			}

			if tmp := cfg.SendCommunityVpnv4; tmp != nil {
				v["send_community_vpnv4"] = *tmp
			}

			if tmp := cfg.SendCommunity6; tmp != nil {
				v["send_community6"] = *tmp
			}

			if tmp := cfg.Shutdown; tmp != nil {
				v["shutdown"] = *tmp
			}

			if tmp := cfg.SoftReconfiguration; tmp != nil {
				v["soft_reconfiguration"] = *tmp
			}

			if tmp := cfg.SoftReconfigurationVpnv4; tmp != nil {
				v["soft_reconfiguration_vpnv4"] = *tmp
			}

			if tmp := cfg.SoftReconfiguration6; tmp != nil {
				v["soft_reconfiguration6"] = *tmp
			}

			if tmp := cfg.StaleRoute; tmp != nil {
				v["stale_route"] = *tmp
			}

			if tmp := cfg.StrictCapabilityMatch; tmp != nil {
				v["strict_capability_match"] = *tmp
			}

			if tmp := cfg.UnsuppressMap; tmp != nil {
				v["unsuppress_map"] = *tmp
			}

			if tmp := cfg.UnsuppressMap6; tmp != nil {
				v["unsuppress_map6"] = *tmp
			}

			if tmp := cfg.UpdateSource; tmp != nil {
				v["update_source"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func flattenRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData, v *[]models.RouterBgpNeighborConditionalAdvertise, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AdvertiseRoutemap; tmp != nil {
				v["advertise_routemap"] = *tmp
			}

			if tmp := cfg.ConditionRoutemap; tmp != nil {
				v["condition_routemap"] = flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "condition_routemap"), sort)
			}

			if tmp := cfg.ConditionType; tmp != nil {
				v["condition_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "advertise_routemap")
	}

	return flat
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d *schema.ResourceData, v *[]models.RouterBgpNeighborConditionalAdvertiseConditionRoutemap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterBgpNeighborConditionalAdvertise6(d *schema.ResourceData, v *[]models.RouterBgpNeighborConditionalAdvertise6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AdvertiseRoutemap; tmp != nil {
				v["advertise_routemap"] = *tmp
			}

			if tmp := cfg.ConditionRoutemap; tmp != nil {
				v["condition_routemap"] = flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "condition_routemap"), sort)
			}

			if tmp := cfg.ConditionType; tmp != nil {
				v["condition_type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "advertise_routemap")
	}

	return flat
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d *schema.ResourceData, v *[]models.RouterBgpNeighborConditionalAdvertise6ConditionRoutemap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterBgpNeighborGroup(d *schema.ResourceData, v *[]models.RouterBgpNeighborGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Activate; tmp != nil {
				v["activate"] = *tmp
			}

			if tmp := cfg.ActivateVpnv4; tmp != nil {
				v["activate_vpnv4"] = *tmp
			}

			if tmp := cfg.Activate6; tmp != nil {
				v["activate6"] = *tmp
			}

			if tmp := cfg.AdditionalPath; tmp != nil {
				v["additional_path"] = *tmp
			}

			if tmp := cfg.AdditionalPathVpnv4; tmp != nil {
				v["additional_path_vpnv4"] = *tmp
			}

			if tmp := cfg.AdditionalPath6; tmp != nil {
				v["additional_path6"] = *tmp
			}

			if tmp := cfg.AdvAdditionalPath; tmp != nil {
				v["adv_additional_path"] = *tmp
			}

			if tmp := cfg.AdvAdditionalPathVpnv4; tmp != nil {
				v["adv_additional_path_vpnv4"] = *tmp
			}

			if tmp := cfg.AdvAdditionalPath6; tmp != nil {
				v["adv_additional_path6"] = *tmp
			}

			if tmp := cfg.AdvertisementInterval; tmp != nil {
				v["advertisement_interval"] = *tmp
			}

			if tmp := cfg.AllowasIn; tmp != nil {
				v["allowas_in"] = *tmp
			}

			if tmp := cfg.AllowasInEnable; tmp != nil {
				v["allowas_in_enable"] = *tmp
			}

			if tmp := cfg.AllowasInEnable6; tmp != nil {
				v["allowas_in_enable6"] = *tmp
			}

			if tmp := cfg.AllowasInVpnv4; tmp != nil {
				v["allowas_in_vpnv4"] = *tmp
			}

			if tmp := cfg.AllowasIn6; tmp != nil {
				v["allowas_in6"] = *tmp
			}

			if tmp := cfg.AsOverride; tmp != nil {
				v["as_override"] = *tmp
			}

			if tmp := cfg.AsOverride6; tmp != nil {
				v["as_override6"] = *tmp
			}

			if tmp := cfg.AttributeUnchanged; tmp != nil {
				v["attribute_unchanged"] = *tmp
			}

			if tmp := cfg.AttributeUnchangedVpnv4; tmp != nil {
				v["attribute_unchanged_vpnv4"] = *tmp
			}

			if tmp := cfg.AttributeUnchanged6; tmp != nil {
				v["attribute_unchanged6"] = *tmp
			}

			if tmp := cfg.Bfd; tmp != nil {
				v["bfd"] = *tmp
			}

			if tmp := cfg.CapabilityDefaultOriginate; tmp != nil {
				v["capability_default_originate"] = *tmp
			}

			if tmp := cfg.CapabilityDefaultOriginate6; tmp != nil {
				v["capability_default_originate6"] = *tmp
			}

			if tmp := cfg.CapabilityDynamic; tmp != nil {
				v["capability_dynamic"] = *tmp
			}

			if tmp := cfg.CapabilityGracefulRestart; tmp != nil {
				v["capability_graceful_restart"] = *tmp
			}

			if tmp := cfg.CapabilityGracefulRestartVpnv4; tmp != nil {
				v["capability_graceful_restart_vpnv4"] = *tmp
			}

			if tmp := cfg.CapabilityGracefulRestart6; tmp != nil {
				v["capability_graceful_restart6"] = *tmp
			}

			if tmp := cfg.CapabilityOrf; tmp != nil {
				v["capability_orf"] = *tmp
			}

			if tmp := cfg.CapabilityOrf6; tmp != nil {
				v["capability_orf6"] = *tmp
			}

			if tmp := cfg.CapabilityRouteRefresh; tmp != nil {
				v["capability_route_refresh"] = *tmp
			}

			if tmp := cfg.ConnectTimer; tmp != nil {
				v["connect_timer"] = *tmp
			}

			if tmp := cfg.DefaultOriginateRoutemap; tmp != nil {
				v["default_originate_routemap"] = *tmp
			}

			if tmp := cfg.DefaultOriginateRoutemap6; tmp != nil {
				v["default_originate_routemap6"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.DistributeListIn; tmp != nil {
				v["distribute_list_in"] = *tmp
			}

			if tmp := cfg.DistributeListInVpnv4; tmp != nil {
				v["distribute_list_in_vpnv4"] = *tmp
			}

			if tmp := cfg.DistributeListIn6; tmp != nil {
				v["distribute_list_in6"] = *tmp
			}

			if tmp := cfg.DistributeListOut; tmp != nil {
				v["distribute_list_out"] = *tmp
			}

			if tmp := cfg.DistributeListOutVpnv4; tmp != nil {
				v["distribute_list_out_vpnv4"] = *tmp
			}

			if tmp := cfg.DistributeListOut6; tmp != nil {
				v["distribute_list_out6"] = *tmp
			}

			if tmp := cfg.DontCapabilityNegotiate; tmp != nil {
				v["dont_capability_negotiate"] = *tmp
			}

			if tmp := cfg.EbgpEnforceMultihop; tmp != nil {
				v["ebgp_enforce_multihop"] = *tmp
			}

			if tmp := cfg.EbgpMultihopTtl; tmp != nil {
				v["ebgp_multihop_ttl"] = *tmp
			}

			if tmp := cfg.FilterListIn; tmp != nil {
				v["filter_list_in"] = *tmp
			}

			if tmp := cfg.FilterListIn6; tmp != nil {
				v["filter_list_in6"] = *tmp
			}

			if tmp := cfg.FilterListOut; tmp != nil {
				v["filter_list_out"] = *tmp
			}

			if tmp := cfg.FilterListOut6; tmp != nil {
				v["filter_list_out6"] = *tmp
			}

			if tmp := cfg.HoldtimeTimer; tmp != nil {
				v["holdtime_timer"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.KeepAliveTimer; tmp != nil {
				v["keep_alive_timer"] = *tmp
			}

			if tmp := cfg.LinkDownFailover; tmp != nil {
				v["link_down_failover"] = *tmp
			}

			if tmp := cfg.LocalAs; tmp != nil {
				v["local_as"] = *tmp
			}

			if tmp := cfg.LocalAsNoPrepend; tmp != nil {
				v["local_as_no_prepend"] = *tmp
			}

			if tmp := cfg.LocalAsReplaceAs; tmp != nil {
				v["local_as_replace_as"] = *tmp
			}

			if tmp := cfg.MaximumPrefix; tmp != nil {
				v["maximum_prefix"] = *tmp
			}

			if tmp := cfg.MaximumPrefixThreshold; tmp != nil {
				v["maximum_prefix_threshold"] = *tmp
			}

			if tmp := cfg.MaximumPrefixThresholdVpnv4; tmp != nil {
				v["maximum_prefix_threshold_vpnv4"] = *tmp
			}

			if tmp := cfg.MaximumPrefixThreshold6; tmp != nil {
				v["maximum_prefix_threshold6"] = *tmp
			}

			if tmp := cfg.MaximumPrefixVpnv4; tmp != nil {
				v["maximum_prefix_vpnv4"] = *tmp
			}

			if tmp := cfg.MaximumPrefixWarningOnly; tmp != nil {
				v["maximum_prefix_warning_only"] = *tmp
			}

			if tmp := cfg.MaximumPrefixWarningOnlyVpnv4; tmp != nil {
				v["maximum_prefix_warning_only_vpnv4"] = *tmp
			}

			if tmp := cfg.MaximumPrefixWarningOnly6; tmp != nil {
				v["maximum_prefix_warning_only6"] = *tmp
			}

			if tmp := cfg.MaximumPrefix6; tmp != nil {
				v["maximum_prefix6"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.NextHopSelf; tmp != nil {
				v["next_hop_self"] = *tmp
			}

			if tmp := cfg.NextHopSelfRr; tmp != nil {
				v["next_hop_self_rr"] = *tmp
			}

			if tmp := cfg.NextHopSelfRr6; tmp != nil {
				v["next_hop_self_rr6"] = *tmp
			}

			if tmp := cfg.NextHopSelfVpnv4; tmp != nil {
				v["next_hop_self_vpnv4"] = *tmp
			}

			if tmp := cfg.NextHopSelf6; tmp != nil {
				v["next_hop_self6"] = *tmp
			}

			if tmp := cfg.OverrideCapability; tmp != nil {
				v["override_capability"] = *tmp
			}

			if tmp := cfg.Passive; tmp != nil {
				v["passive"] = *tmp
			}

			if tmp := cfg.Password; tmp != nil {
				v["password"] = *tmp
			}

			if tmp := cfg.PrefixListIn; tmp != nil {
				v["prefix_list_in"] = *tmp
			}

			if tmp := cfg.PrefixListInVpnv4; tmp != nil {
				v["prefix_list_in_vpnv4"] = *tmp
			}

			if tmp := cfg.PrefixListIn6; tmp != nil {
				v["prefix_list_in6"] = *tmp
			}

			if tmp := cfg.PrefixListOut; tmp != nil {
				v["prefix_list_out"] = *tmp
			}

			if tmp := cfg.PrefixListOutVpnv4; tmp != nil {
				v["prefix_list_out_vpnv4"] = *tmp
			}

			if tmp := cfg.PrefixListOut6; tmp != nil {
				v["prefix_list_out6"] = *tmp
			}

			if tmp := cfg.RemoteAs; tmp != nil {
				v["remote_as"] = *tmp
			}

			if tmp := cfg.RemovePrivateAs; tmp != nil {
				v["remove_private_as"] = *tmp
			}

			if tmp := cfg.RemovePrivateAsVpnv4; tmp != nil {
				v["remove_private_as_vpnv4"] = *tmp
			}

			if tmp := cfg.RemovePrivateAs6; tmp != nil {
				v["remove_private_as6"] = *tmp
			}

			if tmp := cfg.RestartTime; tmp != nil {
				v["restart_time"] = *tmp
			}

			if tmp := cfg.RetainStaleTime; tmp != nil {
				v["retain_stale_time"] = *tmp
			}

			if tmp := cfg.RouteMapIn; tmp != nil {
				v["route_map_in"] = *tmp
			}

			if tmp := cfg.RouteMapInVpnv4; tmp != nil {
				v["route_map_in_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteMapIn6; tmp != nil {
				v["route_map_in6"] = *tmp
			}

			if tmp := cfg.RouteMapOut; tmp != nil {
				v["route_map_out"] = *tmp
			}

			if tmp := cfg.RouteMapOutPreferable; tmp != nil {
				v["route_map_out_preferable"] = *tmp
			}

			if tmp := cfg.RouteMapOutVpnv4; tmp != nil {
				v["route_map_out_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteMapOutVpnv4Preferable; tmp != nil {
				v["route_map_out_vpnv4_preferable"] = *tmp
			}

			if tmp := cfg.RouteMapOut6; tmp != nil {
				v["route_map_out6"] = *tmp
			}

			if tmp := cfg.RouteMapOut6Preferable; tmp != nil {
				v["route_map_out6_preferable"] = *tmp
			}

			if tmp := cfg.RouteReflectorClient; tmp != nil {
				v["route_reflector_client"] = *tmp
			}

			if tmp := cfg.RouteReflectorClientVpnv4; tmp != nil {
				v["route_reflector_client_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteReflectorClient6; tmp != nil {
				v["route_reflector_client6"] = *tmp
			}

			if tmp := cfg.RouteServerClient; tmp != nil {
				v["route_server_client"] = *tmp
			}

			if tmp := cfg.RouteServerClientVpnv4; tmp != nil {
				v["route_server_client_vpnv4"] = *tmp
			}

			if tmp := cfg.RouteServerClient6; tmp != nil {
				v["route_server_client6"] = *tmp
			}

			if tmp := cfg.SendCommunity; tmp != nil {
				v["send_community"] = *tmp
			}

			if tmp := cfg.SendCommunityVpnv4; tmp != nil {
				v["send_community_vpnv4"] = *tmp
			}

			if tmp := cfg.SendCommunity6; tmp != nil {
				v["send_community6"] = *tmp
			}

			if tmp := cfg.Shutdown; tmp != nil {
				v["shutdown"] = *tmp
			}

			if tmp := cfg.SoftReconfiguration; tmp != nil {
				v["soft_reconfiguration"] = *tmp
			}

			if tmp := cfg.SoftReconfigurationVpnv4; tmp != nil {
				v["soft_reconfiguration_vpnv4"] = *tmp
			}

			if tmp := cfg.SoftReconfiguration6; tmp != nil {
				v["soft_reconfiguration6"] = *tmp
			}

			if tmp := cfg.StaleRoute; tmp != nil {
				v["stale_route"] = *tmp
			}

			if tmp := cfg.StrictCapabilityMatch; tmp != nil {
				v["strict_capability_match"] = *tmp
			}

			if tmp := cfg.UnsuppressMap; tmp != nil {
				v["unsuppress_map"] = *tmp
			}

			if tmp := cfg.UnsuppressMap6; tmp != nil {
				v["unsuppress_map6"] = *tmp
			}

			if tmp := cfg.UpdateSource; tmp != nil {
				v["update_source"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterBgpNeighborRange(d *schema.ResourceData, v *[]models.RouterBgpNeighborRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.MaxNeighborNum; tmp != nil {
				v["max_neighbor_num"] = *tmp
			}

			if tmp := cfg.NeighborGroup; tmp != nil {
				v["neighbor_group"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.prefix", prefix, i), *tmp); tmp != nil {
					v["prefix"] = *tmp
				}
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpNeighborRange6(d *schema.ResourceData, v *[]models.RouterBgpNeighborRange6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.MaxNeighborNum; tmp != nil {
				v["max_neighbor_num"] = *tmp
			}

			if tmp := cfg.NeighborGroup; tmp != nil {
				v["neighbor_group"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpNetwork(d *schema.ResourceData, v *[]models.RouterBgpNetwork, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Backdoor; tmp != nil {
				v["backdoor"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.NetworkImportCheck; tmp != nil {
				v["network_import_check"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.prefix", prefix, i), *tmp); tmp != nil {
					v["prefix"] = *tmp
				}
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpNetwork6(d *schema.ResourceData, v *[]models.RouterBgpNetwork6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Backdoor; tmp != nil {
				v["backdoor"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.NetworkImportCheck; tmp != nil {
				v["network_import_check"] = *tmp
			}

			if tmp := cfg.Prefix6; tmp != nil {
				v["prefix6"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterBgpRedistribute(d *schema.ResourceData, v *[]models.RouterBgpRedistribute, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterBgpRedistribute6(d *schema.ResourceData, v *[]models.RouterBgpRedistribute6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterBgpVrf(d *schema.ResourceData, v *[]models.RouterBgpVrf, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ExportRt; tmp != nil {
				v["export_rt"] = flattenRouterBgpVrfExportRt(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "export_rt"), sort)
			}

			if tmp := cfg.ImportRouteMap; tmp != nil {
				v["import_route_map"] = *tmp
			}

			if tmp := cfg.ImportRt; tmp != nil {
				v["import_rt"] = flattenRouterBgpVrfImportRt(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "import_rt"), sort)
			}

			if tmp := cfg.LeakTarget; tmp != nil {
				v["leak_target"] = flattenRouterBgpVrfLeakTarget(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "leak_target"), sort)
			}

			if tmp := cfg.Rd; tmp != nil {
				v["rd"] = *tmp
			}

			if tmp := cfg.Role; tmp != nil {
				v["role"] = *tmp
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrfExportRt(d *schema.ResourceData, v *[]models.RouterBgpVrfExportRt, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.RouteTarget; tmp != nil {
				v["route_target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "route_target")
	}

	return flat
}

func flattenRouterBgpVrfImportRt(d *schema.ResourceData, v *[]models.RouterBgpVrfImportRt, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.RouteTarget; tmp != nil {
				v["route_target"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "route_target")
	}

	return flat
}

func flattenRouterBgpVrfLeakTarget(d *schema.ResourceData, v *[]models.RouterBgpVrfLeakTarget, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrfLeak(d *schema.ResourceData, v *[]models.RouterBgpVrfLeak, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Target; tmp != nil {
				v["target"] = flattenRouterBgpVrfLeakTarget(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "target"), sort)
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrfLeakTarget(d *schema.ResourceData, v *[]models.RouterBgpVrfLeakTarget, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrfLeak6(d *schema.ResourceData, v *[]models.RouterBgpVrfLeak6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Target; tmp != nil {
				v["target"] = flattenRouterBgpVrfLeak6Target(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "target"), sort)
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrfLeak6Target(d *schema.ResourceData, v *[]models.RouterBgpVrfLeak6Target, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrf6(d *schema.ResourceData, v *[]models.RouterBgpVrf6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.LeakTarget; tmp != nil {
				v["leak_target"] = flattenRouterBgpVrf6LeakTarget(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "leak_target"), sort)
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func flattenRouterBgpVrf6LeakTarget(d *schema.ResourceData, v *[]models.RouterBgpVrf6LeakTarget, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.RouteMap; tmp != nil {
				v["route_map"] = *tmp
			}

			if tmp := cfg.Vrf; tmp != nil {
				v["vrf"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vrf")
	}

	return flat
}

func refreshObjectRouterBgp(d *schema.ResourceData, o *models.RouterBgp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AdditionalPath != nil {
		v := *o.AdditionalPath

		if err = d.Set("additional_path", v); err != nil {
			return diag.Errorf("error reading additional_path: %v", err)
		}
	}

	if o.AdditionalPathSelect != nil {
		v := *o.AdditionalPathSelect

		if err = d.Set("additional_path_select", v); err != nil {
			return diag.Errorf("error reading additional_path_select: %v", err)
		}
	}

	if o.AdditionalPathSelectVpnv4 != nil {
		v := *o.AdditionalPathSelectVpnv4

		if err = d.Set("additional_path_select_vpnv4", v); err != nil {
			return diag.Errorf("error reading additional_path_select_vpnv4: %v", err)
		}
	}

	if o.AdditionalPathSelect6 != nil {
		v := *o.AdditionalPathSelect6

		if err = d.Set("additional_path_select6", v); err != nil {
			return diag.Errorf("error reading additional_path_select6: %v", err)
		}
	}

	if o.AdditionalPathVpnv4 != nil {
		v := *o.AdditionalPathVpnv4

		if err = d.Set("additional_path_vpnv4", v); err != nil {
			return diag.Errorf("error reading additional_path_vpnv4: %v", err)
		}
	}

	if o.AdditionalPath6 != nil {
		v := *o.AdditionalPath6

		if err = d.Set("additional_path6", v); err != nil {
			return diag.Errorf("error reading additional_path6: %v", err)
		}
	}

	if o.AdminDistance != nil {
		if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(d, o.AdminDistance, "admin_distance", sort)); err != nil {
			return diag.Errorf("error reading admin_distance: %v", err)
		}
	}

	if o.AggregateAddress != nil {
		if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(d, o.AggregateAddress, "aggregate_address", sort)); err != nil {
			return diag.Errorf("error reading aggregate_address: %v", err)
		}
	}

	if o.AggregateAddress6 != nil {
		if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(d, o.AggregateAddress6, "aggregate_address6", sort)); err != nil {
			return diag.Errorf("error reading aggregate_address6: %v", err)
		}
	}

	if o.AlwaysCompareMed != nil {
		v := *o.AlwaysCompareMed

		if err = d.Set("always_compare_med", v); err != nil {
			return diag.Errorf("error reading always_compare_med: %v", err)
		}
	}

	if o.As != nil {
		v := *o.As

		if err = d.Set("as", v); err != nil {
			return diag.Errorf("error reading as: %v", err)
		}
	}

	if o.BestpathAsPathIgnore != nil {
		v := *o.BestpathAsPathIgnore

		if err = d.Set("bestpath_as_path_ignore", v); err != nil {
			return diag.Errorf("error reading bestpath_as_path_ignore: %v", err)
		}
	}

	if o.BestpathCmpConfedAspath != nil {
		v := *o.BestpathCmpConfedAspath

		if err = d.Set("bestpath_cmp_confed_aspath", v); err != nil {
			return diag.Errorf("error reading bestpath_cmp_confed_aspath: %v", err)
		}
	}

	if o.BestpathCmpRouterid != nil {
		v := *o.BestpathCmpRouterid

		if err = d.Set("bestpath_cmp_routerid", v); err != nil {
			return diag.Errorf("error reading bestpath_cmp_routerid: %v", err)
		}
	}

	if o.BestpathMedConfed != nil {
		v := *o.BestpathMedConfed

		if err = d.Set("bestpath_med_confed", v); err != nil {
			return diag.Errorf("error reading bestpath_med_confed: %v", err)
		}
	}

	if o.BestpathMedMissingAsWorst != nil {
		v := *o.BestpathMedMissingAsWorst

		if err = d.Set("bestpath_med_missing_as_worst", v); err != nil {
			return diag.Errorf("error reading bestpath_med_missing_as_worst: %v", err)
		}
	}

	if o.ClientToClientReflection != nil {
		v := *o.ClientToClientReflection

		if err = d.Set("client_to_client_reflection", v); err != nil {
			return diag.Errorf("error reading client_to_client_reflection: %v", err)
		}
	}

	if o.ClusterId != nil {
		v := *o.ClusterId

		if err = d.Set("cluster_id", v); err != nil {
			return diag.Errorf("error reading cluster_id: %v", err)
		}
	}

	if o.ConfederationIdentifier != nil {
		v := *o.ConfederationIdentifier

		if err = d.Set("confederation_identifier", v); err != nil {
			return diag.Errorf("error reading confederation_identifier: %v", err)
		}
	}

	if o.ConfederationPeers != nil {
		if err = d.Set("confederation_peers", flattenRouterBgpConfederationPeers(d, o.ConfederationPeers, "confederation_peers", sort)); err != nil {
			return diag.Errorf("error reading confederation_peers: %v", err)
		}
	}

	if o.Dampening != nil {
		v := *o.Dampening

		if err = d.Set("dampening", v); err != nil {
			return diag.Errorf("error reading dampening: %v", err)
		}
	}

	if o.DampeningMaxSuppressTime != nil {
		v := *o.DampeningMaxSuppressTime

		if err = d.Set("dampening_max_suppress_time", v); err != nil {
			return diag.Errorf("error reading dampening_max_suppress_time: %v", err)
		}
	}

	if o.DampeningReachabilityHalfLife != nil {
		v := *o.DampeningReachabilityHalfLife

		if err = d.Set("dampening_reachability_half_life", v); err != nil {
			return diag.Errorf("error reading dampening_reachability_half_life: %v", err)
		}
	}

	if o.DampeningReuse != nil {
		v := *o.DampeningReuse

		if err = d.Set("dampening_reuse", v); err != nil {
			return diag.Errorf("error reading dampening_reuse: %v", err)
		}
	}

	if o.DampeningRouteMap != nil {
		v := *o.DampeningRouteMap

		if err = d.Set("dampening_route_map", v); err != nil {
			return diag.Errorf("error reading dampening_route_map: %v", err)
		}
	}

	if o.DampeningSuppress != nil {
		v := *o.DampeningSuppress

		if err = d.Set("dampening_suppress", v); err != nil {
			return diag.Errorf("error reading dampening_suppress: %v", err)
		}
	}

	if o.DampeningUnreachabilityHalfLife != nil {
		v := *o.DampeningUnreachabilityHalfLife

		if err = d.Set("dampening_unreachability_half_life", v); err != nil {
			return diag.Errorf("error reading dampening_unreachability_half_life: %v", err)
		}
	}

	if o.DefaultLocalPreference != nil {
		v := *o.DefaultLocalPreference

		if err = d.Set("default_local_preference", v); err != nil {
			return diag.Errorf("error reading default_local_preference: %v", err)
		}
	}

	if o.DeterministicMed != nil {
		v := *o.DeterministicMed

		if err = d.Set("deterministic_med", v); err != nil {
			return diag.Errorf("error reading deterministic_med: %v", err)
		}
	}

	if o.DistanceExternal != nil {
		v := *o.DistanceExternal

		if err = d.Set("distance_external", v); err != nil {
			return diag.Errorf("error reading distance_external: %v", err)
		}
	}

	if o.DistanceInternal != nil {
		v := *o.DistanceInternal

		if err = d.Set("distance_internal", v); err != nil {
			return diag.Errorf("error reading distance_internal: %v", err)
		}
	}

	if o.DistanceLocal != nil {
		v := *o.DistanceLocal

		if err = d.Set("distance_local", v); err != nil {
			return diag.Errorf("error reading distance_local: %v", err)
		}
	}

	if o.EbgpMultipath != nil {
		v := *o.EbgpMultipath

		if err = d.Set("ebgp_multipath", v); err != nil {
			return diag.Errorf("error reading ebgp_multipath: %v", err)
		}
	}

	if o.EnforceFirstAs != nil {
		v := *o.EnforceFirstAs

		if err = d.Set("enforce_first_as", v); err != nil {
			return diag.Errorf("error reading enforce_first_as: %v", err)
		}
	}

	if o.FastExternalFailover != nil {
		v := *o.FastExternalFailover

		if err = d.Set("fast_external_failover", v); err != nil {
			return diag.Errorf("error reading fast_external_failover: %v", err)
		}
	}

	if o.GracefulEndOnTimer != nil {
		v := *o.GracefulEndOnTimer

		if err = d.Set("graceful_end_on_timer", v); err != nil {
			return diag.Errorf("error reading graceful_end_on_timer: %v", err)
		}
	}

	if o.GracefulRestart != nil {
		v := *o.GracefulRestart

		if err = d.Set("graceful_restart", v); err != nil {
			return diag.Errorf("error reading graceful_restart: %v", err)
		}
	}

	if o.GracefulRestartTime != nil {
		v := *o.GracefulRestartTime

		if err = d.Set("graceful_restart_time", v); err != nil {
			return diag.Errorf("error reading graceful_restart_time: %v", err)
		}
	}

	if o.GracefulStalepathTime != nil {
		v := *o.GracefulStalepathTime

		if err = d.Set("graceful_stalepath_time", v); err != nil {
			return diag.Errorf("error reading graceful_stalepath_time: %v", err)
		}
	}

	if o.GracefulUpdateDelay != nil {
		v := *o.GracefulUpdateDelay

		if err = d.Set("graceful_update_delay", v); err != nil {
			return diag.Errorf("error reading graceful_update_delay: %v", err)
		}
	}

	if o.HoldtimeTimer != nil {
		v := *o.HoldtimeTimer

		if err = d.Set("holdtime_timer", v); err != nil {
			return diag.Errorf("error reading holdtime_timer: %v", err)
		}
	}

	if o.IbgpMultipath != nil {
		v := *o.IbgpMultipath

		if err = d.Set("ibgp_multipath", v); err != nil {
			return diag.Errorf("error reading ibgp_multipath: %v", err)
		}
	}

	if o.IgnoreOptionalCapability != nil {
		v := *o.IgnoreOptionalCapability

		if err = d.Set("ignore_optional_capability", v); err != nil {
			return diag.Errorf("error reading ignore_optional_capability: %v", err)
		}
	}

	if o.KeepaliveTimer != nil {
		v := *o.KeepaliveTimer

		if err = d.Set("keepalive_timer", v); err != nil {
			return diag.Errorf("error reading keepalive_timer: %v", err)
		}
	}

	if o.LogNeighbourChanges != nil {
		v := *o.LogNeighbourChanges

		if err = d.Set("log_neighbour_changes", v); err != nil {
			return diag.Errorf("error reading log_neighbour_changes: %v", err)
		}
	}

	if o.MultipathRecursiveDistance != nil {
		v := *o.MultipathRecursiveDistance

		if err = d.Set("multipath_recursive_distance", v); err != nil {
			return diag.Errorf("error reading multipath_recursive_distance: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenRouterBgpNeighbor(d, o.Neighbor, "neighbor", sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	if o.NeighborGroup != nil {
		if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(d, o.NeighborGroup, "neighbor_group", sort)); err != nil {
			return diag.Errorf("error reading neighbor_group: %v", err)
		}
	}

	if o.NeighborRange != nil {
		if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(d, o.NeighborRange, "neighbor_range", sort)); err != nil {
			return diag.Errorf("error reading neighbor_range: %v", err)
		}
	}

	if o.NeighborRange6 != nil {
		if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(d, o.NeighborRange6, "neighbor_range6", sort)); err != nil {
			return diag.Errorf("error reading neighbor_range6: %v", err)
		}
	}

	if o.Network != nil {
		if err = d.Set("network", flattenRouterBgpNetwork(d, o.Network, "network", sort)); err != nil {
			return diag.Errorf("error reading network: %v", err)
		}
	}

	if o.NetworkImportCheck != nil {
		v := *o.NetworkImportCheck

		if err = d.Set("network_import_check", v); err != nil {
			return diag.Errorf("error reading network_import_check: %v", err)
		}
	}

	if o.Network6 != nil {
		if err = d.Set("network6", flattenRouterBgpNetwork6(d, o.Network6, "network6", sort)); err != nil {
			return diag.Errorf("error reading network6: %v", err)
		}
	}

	if o.RecursiveInheritPriority != nil {
		v := *o.RecursiveInheritPriority

		if err = d.Set("recursive_inherit_priority", v); err != nil {
			return diag.Errorf("error reading recursive_inherit_priority: %v", err)
		}
	}

	if o.RecursiveNextHop != nil {
		v := *o.RecursiveNextHop

		if err = d.Set("recursive_next_hop", v); err != nil {
			return diag.Errorf("error reading recursive_next_hop: %v", err)
		}
	}

	if o.Redistribute != nil {
		if err = d.Set("redistribute", flattenRouterBgpRedistribute(d, o.Redistribute, "redistribute", sort)); err != nil {
			return diag.Errorf("error reading redistribute: %v", err)
		}
	}

	if o.Redistribute6 != nil {
		if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(d, o.Redistribute6, "redistribute6", sort)); err != nil {
			return diag.Errorf("error reading redistribute6: %v", err)
		}
	}

	if o.RouterId != nil {
		v := *o.RouterId

		if err = d.Set("router_id", v); err != nil {
			return diag.Errorf("error reading router_id: %v", err)
		}
	}

	if o.ScanTime != nil {
		v := *o.ScanTime

		if err = d.Set("scan_time", v); err != nil {
			return diag.Errorf("error reading scan_time: %v", err)
		}
	}

	if o.Synchronization != nil {
		v := *o.Synchronization

		if err = d.Set("synchronization", v); err != nil {
			return diag.Errorf("error reading synchronization: %v", err)
		}
	}

	if o.TagResolveMode != nil {
		v := *o.TagResolveMode

		if err = d.Set("tag_resolve_mode", v); err != nil {
			return diag.Errorf("error reading tag_resolve_mode: %v", err)
		}
	}

	if o.Vrf != nil {
		if err = d.Set("vrf", flattenRouterBgpVrf(d, o.Vrf, "vrf", sort)); err != nil {
			return diag.Errorf("error reading vrf: %v", err)
		}
	}

	if o.VrfLeak != nil {
		if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(d, o.VrfLeak, "vrf_leak", sort)); err != nil {
			return diag.Errorf("error reading vrf_leak: %v", err)
		}
	}

	if o.VrfLeak6 != nil {
		if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(d, o.VrfLeak6, "vrf_leak6", sort)); err != nil {
			return diag.Errorf("error reading vrf_leak6: %v", err)
		}
	}

	if o.Vrf6 != nil {
		if err = d.Set("vrf6", flattenRouterBgpVrf6(d, o.Vrf6, "vrf6", sort)); err != nil {
			return diag.Errorf("error reading vrf6: %v", err)
		}
	}

	return nil
}

func expandRouterBgpAdminDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpAdminDistance, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpAdminDistance

	for i := range l {
		tmp := models.RouterBgpAdminDistance{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.distance", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Distance = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.neighbour_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NeighbourPrefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteList = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpAggregateAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpAggregateAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpAggregateAddress

	for i := range l {
		tmp := models.RouterBgpAggregateAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.as_set", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AsSet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.summary_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SummaryOnly = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpAggregateAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpAggregateAddress6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpAggregateAddress6

	for i := range l {
		tmp := models.RouterBgpAggregateAddress6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.as_set", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AsSet = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.summary_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SummaryOnly = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpConfederationPeers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpConfederationPeers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpConfederationPeers

	for i := range l {
		tmp := models.RouterBgpConfederationPeers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.peer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Peer = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighbor

	for i := range l {
		tmp := models.RouterBgpNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.activate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Activate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.activate_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ActivateVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.activate6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Activate6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalPath = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_path_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalPathVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_path6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalPath6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_additional_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvAdditionalPath = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_additional_path_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvAdditionalPathVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_additional_path6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvAdditionalPath6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.advertisement_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvertisementInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllowasIn = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in_enable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AllowasInEnable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in_enable6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AllowasInEnable6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllowasInVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllowasIn6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.as_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AsOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.as_override6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AsOverride6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.attribute_unchanged", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AttributeUnchanged = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.attribute_unchanged_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AttributeUnchangedVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.attribute_unchanged6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AttributeUnchanged6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bfd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_default_originate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityDefaultOriginate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_default_originate6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityDefaultOriginate6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_dynamic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityDynamic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_graceful_restart", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityGracefulRestart = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_graceful_restart_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityGracefulRestartVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_graceful_restart6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityGracefulRestart6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_orf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityOrf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_orf6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityOrf6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_route_refresh", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityRouteRefresh = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.conditional_advertise", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpNeighborConditionalAdvertise(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpNeighborConditionalAdvertise
			// 	}
			tmp.ConditionalAdvertise = v2

		}

		pre_append = fmt.Sprintf("%s.%d.conditional_advertise6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpNeighborConditionalAdvertise6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpNeighborConditionalAdvertise6
			// 	}
			tmp.ConditionalAdvertise6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.connect_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ConnectTimer = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_originate_routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultOriginateRoutemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_originate_routemap6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultOriginateRoutemap6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListInVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_out_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListOutVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dont_capability_negotiate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DontCapabilityNegotiate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ebgp_enforce_multihop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EbgpEnforceMultihop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ebgp_multihop_ttl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EbgpMultihopTtl = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.holdtime_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HoldtimeTimer = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keep_alive_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeepAliveTimer = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.link_down_failover", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LinkDownFailover = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_as_no_prepend", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalAsNoPrepend = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_as_replace_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalAsReplaceAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefix = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_threshold_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixThresholdVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_threshold6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixThreshold6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_warning_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaximumPrefixWarningOnly = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_warning_only_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaximumPrefixWarningOnlyVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_warning_only6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaximumPrefixWarningOnly6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefix6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self_rr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelfRr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self_rr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelfRr6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelfVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelf6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_capability", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideCapability = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.passive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Passive = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Password = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListInVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_out_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListOutVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remote_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemoteAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remove_private_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemovePrivateAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remove_private_as_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemovePrivateAsVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remove_private_as6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemovePrivateAs6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.restart_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RestartTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retain_stale_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RetainStaleTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapInVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out_preferable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOutPreferable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOutVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out_vpnv4_preferable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOutVpnv4Preferable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out6_preferable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOut6Preferable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_reflector_client", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteReflectorClient = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_reflector_client_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteReflectorClientVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_reflector_client6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteReflectorClient6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_server_client", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteServerClient = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_server_client_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteServerClientVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_server_client6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteServerClient6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendCommunity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_community_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendCommunityVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_community6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendCommunity6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.shutdown", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Shutdown = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.soft_reconfiguration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SoftReconfiguration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.soft_reconfiguration_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SoftReconfigurationVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.soft_reconfiguration6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SoftReconfiguration6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stale_route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StaleRoute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.strict_capability_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StrictCapabilityMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsuppress_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsuppressMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsuppress_map6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsuppressMap6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.update_source", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpdateSource = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Weight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborConditionalAdvertise, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborConditionalAdvertise

	for i := range l {
		tmp := models.RouterBgpNeighborConditionalAdvertise{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.advertise_routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdvertiseRoutemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.condition_routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpNeighborConditionalAdvertiseConditionRoutemap
			// 	}
			tmp.ConditionRoutemap = v2

		}

		pre_append = fmt.Sprintf("%s.%d.condition_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ConditionType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborConditionalAdvertiseConditionRoutemap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborConditionalAdvertiseConditionRoutemap

	for i := range l {
		tmp := models.RouterBgpNeighborConditionalAdvertiseConditionRoutemap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborConditionalAdvertise6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborConditionalAdvertise6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborConditionalAdvertise6

	for i := range l {
		tmp := models.RouterBgpNeighborConditionalAdvertise6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.advertise_routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdvertiseRoutemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.condition_routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpNeighborConditionalAdvertise6ConditionRoutemap
			// 	}
			tmp.ConditionRoutemap = v2

		}

		pre_append = fmt.Sprintf("%s.%d.condition_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ConditionType = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborConditionalAdvertise6ConditionRoutemap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborConditionalAdvertise6ConditionRoutemap

	for i := range l {
		tmp := models.RouterBgpNeighborConditionalAdvertise6ConditionRoutemap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborGroup

	for i := range l {
		tmp := models.RouterBgpNeighborGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.activate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Activate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.activate_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ActivateVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.activate6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Activate6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalPath = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_path_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalPathVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_path6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalPath6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_additional_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvAdditionalPath = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_additional_path_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvAdditionalPathVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.adv_additional_path6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvAdditionalPath6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.advertisement_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdvertisementInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllowasIn = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in_enable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AllowasInEnable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in_enable6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AllowasInEnable6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllowasInVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.allowas_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AllowasIn6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.as_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AsOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.as_override6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AsOverride6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.attribute_unchanged", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AttributeUnchanged = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.attribute_unchanged_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AttributeUnchangedVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.attribute_unchanged6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AttributeUnchanged6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bfd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_default_originate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityDefaultOriginate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_default_originate6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityDefaultOriginate6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_dynamic", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityDynamic = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_graceful_restart", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityGracefulRestart = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_graceful_restart_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityGracefulRestartVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_graceful_restart6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityGracefulRestart6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_orf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityOrf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_orf6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityOrf6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.capability_route_refresh", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CapabilityRouteRefresh = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.connect_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ConnectTimer = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_originate_routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultOriginateRoutemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_originate_routemap6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultOriginateRoutemap6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListInVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_out_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListOutVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.distribute_list_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DistributeListOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dont_capability_negotiate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DontCapabilityNegotiate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ebgp_enforce_multihop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EbgpEnforceMultihop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ebgp_multihop_ttl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EbgpMultihopTtl = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterListOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.holdtime_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HoldtimeTimer = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keep_alive_timer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.KeepAliveTimer = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.link_down_failover", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LinkDownFailover = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_as_no_prepend", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalAsNoPrepend = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.local_as_replace_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LocalAsReplaceAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefix = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_threshold_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixThresholdVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_threshold6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixThreshold6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefixVpnv4 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_warning_only", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaximumPrefixWarningOnly = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_warning_only_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaximumPrefixWarningOnlyVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix_warning_only6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaximumPrefixWarningOnly6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.maximum_prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaximumPrefix6 = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelf = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self_rr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelfRr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self_rr6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelfRr6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelfVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.next_hop_self6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NextHopSelf6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.override_capability", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OverrideCapability = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.passive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Passive = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.password", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Password = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListInVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_out_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListOutVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_list_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrefixListOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remote_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemoteAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remove_private_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemovePrivateAs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remove_private_as_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemovePrivateAsVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.remove_private_as6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RemovePrivateAs6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.restart_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RestartTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retain_stale_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RetainStaleTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_in", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapIn = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_in_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapInVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_in6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapIn6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out_preferable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOutPreferable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOutVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out_vpnv4_preferable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOutVpnv4Preferable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOut6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map_out6_preferable", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMapOut6Preferable = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_reflector_client", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteReflectorClient = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_reflector_client_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteReflectorClientVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_reflector_client6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteReflectorClient6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_server_client", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteServerClient = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_server_client_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteServerClientVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_server_client6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteServerClient6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendCommunity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_community_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendCommunityVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.send_community6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SendCommunity6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.shutdown", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Shutdown = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.soft_reconfiguration", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SoftReconfiguration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.soft_reconfiguration_vpnv4", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SoftReconfigurationVpnv4 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.soft_reconfiguration6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SoftReconfiguration6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stale_route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StaleRoute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.strict_capability_match", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StrictCapabilityMatch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsuppress_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsuppressMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsuppress_map6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsuppressMap6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.update_source", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpdateSource = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Weight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborRange

	for i := range l {
		tmp := models.RouterBgpNeighborRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_neighbor_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxNeighborNum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.neighbor_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NeighborGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNeighborRange6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNeighborRange6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNeighborRange6

	for i := range l {
		tmp := models.RouterBgpNeighborRange6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_neighbor_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxNeighborNum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.neighbor_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NeighborGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNetwork, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNetwork

	for i := range l {
		tmp := models.RouterBgpNetwork{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.backdoor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Backdoor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.network_import_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NetworkImportCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpNetwork6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpNetwork6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpNetwork6

	for i := range l {
		tmp := models.RouterBgpNetwork6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.backdoor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Backdoor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.network_import_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NetworkImportCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpRedistribute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpRedistribute

	for i := range l {
		tmp := models.RouterBgpRedistribute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpRedistribute6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpRedistribute6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpRedistribute6

	for i := range l {
		tmp := models.RouterBgpRedistribute6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrf

	for i := range l {
		tmp := models.RouterBgpVrf{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.export_rt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpVrfExportRt(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpVrfExportRt
			// 	}
			tmp.ExportRt = v2

		}

		pre_append = fmt.Sprintf("%s.%d.import_route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImportRouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.import_rt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpVrfImportRt(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpVrfImportRt
			// 	}
			tmp.ImportRt = v2

		}

		pre_append = fmt.Sprintf("%s.%d.leak_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpVrfLeakTarget(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpVrfLeakTarget
			// 	}
			tmp.LeakTarget = v2

		}

		pre_append = fmt.Sprintf("%s.%d.rd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Rd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Role = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfExportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfExportRt, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfExportRt

	for i := range l {
		tmp := models.RouterBgpVrfExportRt{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.route_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteTarget = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfImportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfImportRt, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfImportRt

	for i := range l {
		tmp := models.RouterBgpVrfImportRt{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.route_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteTarget = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfLeakTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfLeakTarget, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfLeakTarget

	for i := range l {
		tmp := models.RouterBgpVrfLeakTarget{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfLeak(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfLeak, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfLeak

	for i := range l {
		tmp := models.RouterBgpVrfLeak{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpVrfLeakTarget(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpVrfLeakTarget
			// 	}
			tmp.Target = v2

		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfLeakTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfLeakTarget, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfLeakTarget

	for i := range l {
		tmp := models.RouterBgpVrfLeakTarget{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfLeak6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfLeak6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfLeak6

	for i := range l {
		tmp := models.RouterBgpVrfLeak6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpVrfLeak6Target(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpVrfLeak6Target
			// 	}
			tmp.Target = v2

		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrfLeak6Target(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrfLeak6Target, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrfLeak6Target

	for i := range l {
		tmp := models.RouterBgpVrfLeak6Target{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrf6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrf6

	for i := range l {
		tmp := models.RouterBgpVrf6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.leak_target", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterBgpVrf6LeakTarget(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterBgpVrf6LeakTarget
			// 	}
			tmp.LeakTarget = v2

		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterBgpVrf6LeakTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterBgpVrf6LeakTarget, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterBgpVrf6LeakTarget

	for i := range l {
		tmp := models.RouterBgpVrf6LeakTarget{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.route_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouteMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Vrf = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterBgp(d *schema.ResourceData, sv string) (*models.RouterBgp, diag.Diagnostics) {
	obj := models.RouterBgp{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("additional_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("additional_path", sv)
				diags = append(diags, e)
			}
			obj.AdditionalPath = &v2
		}
	}
	if v1, ok := d.GetOk("additional_path_select"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("additional_path_select", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdditionalPathSelect = &tmp
		}
	}
	if v1, ok := d.GetOk("additional_path_select_vpnv4"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("additional_path_select_vpnv4", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdditionalPathSelectVpnv4 = &tmp
		}
	}
	if v1, ok := d.GetOk("additional_path_select6"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("additional_path_select6", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdditionalPathSelect6 = &tmp
		}
	}
	if v1, ok := d.GetOk("additional_path_vpnv4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("additional_path_vpnv4", sv)
				diags = append(diags, e)
			}
			obj.AdditionalPathVpnv4 = &v2
		}
	}
	if v1, ok := d.GetOk("additional_path6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("additional_path6", sv)
				diags = append(diags, e)
			}
			obj.AdditionalPath6 = &v2
		}
	}
	if v, ok := d.GetOk("admin_distance"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("admin_distance", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpAdminDistance(d, v, "admin_distance", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AdminDistance = t
		}
	} else if d.HasChange("admin_distance") {
		old, new := d.GetChange("admin_distance")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AdminDistance = &[]models.RouterBgpAdminDistance{}
		}
	}
	if v, ok := d.GetOk("aggregate_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("aggregate_address", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpAggregateAddress(d, v, "aggregate_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AggregateAddress = t
		}
	} else if d.HasChange("aggregate_address") {
		old, new := d.GetChange("aggregate_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AggregateAddress = &[]models.RouterBgpAggregateAddress{}
		}
	}
	if v, ok := d.GetOk("aggregate_address6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("aggregate_address6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpAggregateAddress6(d, v, "aggregate_address6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AggregateAddress6 = t
		}
	} else if d.HasChange("aggregate_address6") {
		old, new := d.GetChange("aggregate_address6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AggregateAddress6 = &[]models.RouterBgpAggregateAddress6{}
		}
	}
	if v1, ok := d.GetOk("always_compare_med"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("always_compare_med", sv)
				diags = append(diags, e)
			}
			obj.AlwaysCompareMed = &v2
		}
	}
	if v1, ok := d.GetOk("as"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("as", sv)
				diags = append(diags, e)
			}
			obj.As = &v2
		}
	}
	if v1, ok := d.GetOk("bestpath_as_path_ignore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bestpath_as_path_ignore", sv)
				diags = append(diags, e)
			}
			obj.BestpathAsPathIgnore = &v2
		}
	}
	if v1, ok := d.GetOk("bestpath_cmp_confed_aspath"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bestpath_cmp_confed_aspath", sv)
				diags = append(diags, e)
			}
			obj.BestpathCmpConfedAspath = &v2
		}
	}
	if v1, ok := d.GetOk("bestpath_cmp_routerid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bestpath_cmp_routerid", sv)
				diags = append(diags, e)
			}
			obj.BestpathCmpRouterid = &v2
		}
	}
	if v1, ok := d.GetOk("bestpath_med_confed"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bestpath_med_confed", sv)
				diags = append(diags, e)
			}
			obj.BestpathMedConfed = &v2
		}
	}
	if v1, ok := d.GetOk("bestpath_med_missing_as_worst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bestpath_med_missing_as_worst", sv)
				diags = append(diags, e)
			}
			obj.BestpathMedMissingAsWorst = &v2
		}
	}
	if v1, ok := d.GetOk("client_to_client_reflection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_to_client_reflection", sv)
				diags = append(diags, e)
			}
			obj.ClientToClientReflection = &v2
		}
	}
	if v1, ok := d.GetOk("cluster_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cluster_id", sv)
				diags = append(diags, e)
			}
			obj.ClusterId = &v2
		}
	}
	if v1, ok := d.GetOk("confederation_identifier"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("confederation_identifier", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ConfederationIdentifier = &tmp
		}
	}
	if v, ok := d.GetOk("confederation_peers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("confederation_peers", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpConfederationPeers(d, v, "confederation_peers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ConfederationPeers = t
		}
	} else if d.HasChange("confederation_peers") {
		old, new := d.GetChange("confederation_peers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ConfederationPeers = &[]models.RouterBgpConfederationPeers{}
		}
	}
	if v1, ok := d.GetOk("dampening"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening", sv)
				diags = append(diags, e)
			}
			obj.Dampening = &v2
		}
	}
	if v1, ok := d.GetOk("dampening_max_suppress_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening_max_suppress_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DampeningMaxSuppressTime = &tmp
		}
	}
	if v1, ok := d.GetOk("dampening_reachability_half_life"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening_reachability_half_life", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DampeningReachabilityHalfLife = &tmp
		}
	}
	if v1, ok := d.GetOk("dampening_reuse"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening_reuse", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DampeningReuse = &tmp
		}
	}
	if v1, ok := d.GetOk("dampening_route_map"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening_route_map", sv)
				diags = append(diags, e)
			}
			obj.DampeningRouteMap = &v2
		}
	}
	if v1, ok := d.GetOk("dampening_suppress"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening_suppress", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DampeningSuppress = &tmp
		}
	}
	if v1, ok := d.GetOk("dampening_unreachability_half_life"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dampening_unreachability_half_life", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DampeningUnreachabilityHalfLife = &tmp
		}
	}
	if v1, ok := d.GetOk("default_local_preference"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_local_preference", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultLocalPreference = &tmp
		}
	}
	if v1, ok := d.GetOk("deterministic_med"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deterministic_med", sv)
				diags = append(diags, e)
			}
			obj.DeterministicMed = &v2
		}
	}
	if v1, ok := d.GetOk("distance_external"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance_external", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DistanceExternal = &tmp
		}
	}
	if v1, ok := d.GetOk("distance_internal"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance_internal", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DistanceInternal = &tmp
		}
	}
	if v1, ok := d.GetOk("distance_local"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance_local", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DistanceLocal = &tmp
		}
	}
	if v1, ok := d.GetOk("ebgp_multipath"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ebgp_multipath", sv)
				diags = append(diags, e)
			}
			obj.EbgpMultipath = &v2
		}
	}
	if v1, ok := d.GetOk("enforce_first_as"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enforce_first_as", sv)
				diags = append(diags, e)
			}
			obj.EnforceFirstAs = &v2
		}
	}
	if v1, ok := d.GetOk("fast_external_failover"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fast_external_failover", sv)
				diags = append(diags, e)
			}
			obj.FastExternalFailover = &v2
		}
	}
	if v1, ok := d.GetOk("graceful_end_on_timer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graceful_end_on_timer", sv)
				diags = append(diags, e)
			}
			obj.GracefulEndOnTimer = &v2
		}
	}
	if v1, ok := d.GetOk("graceful_restart"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graceful_restart", sv)
				diags = append(diags, e)
			}
			obj.GracefulRestart = &v2
		}
	}
	if v1, ok := d.GetOk("graceful_restart_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graceful_restart_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GracefulRestartTime = &tmp
		}
	}
	if v1, ok := d.GetOk("graceful_stalepath_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graceful_stalepath_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GracefulStalepathTime = &tmp
		}
	}
	if v1, ok := d.GetOk("graceful_update_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graceful_update_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GracefulUpdateDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("holdtime_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("holdtime_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HoldtimeTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("ibgp_multipath"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ibgp_multipath", sv)
				diags = append(diags, e)
			}
			obj.IbgpMultipath = &v2
		}
	}
	if v1, ok := d.GetOk("ignore_optional_capability"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ignore_optional_capability", sv)
				diags = append(diags, e)
			}
			obj.IgnoreOptionalCapability = &v2
		}
	}
	if v1, ok := d.GetOk("keepalive_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keepalive_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeepaliveTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("log_neighbour_changes"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_neighbour_changes", sv)
				diags = append(diags, e)
			}
			obj.LogNeighbourChanges = &v2
		}
	}
	if v1, ok := d.GetOk("multipath_recursive_distance"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("multipath_recursive_distance", sv)
				diags = append(diags, e)
			}
			obj.MultipathRecursiveDistance = &v2
		}
	}
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.RouterBgpNeighbor{}
		}
	}
	if v, ok := d.GetOk("neighbor_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor_group", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNeighborGroup(d, v, "neighbor_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NeighborGroup = t
		}
	} else if d.HasChange("neighbor_group") {
		old, new := d.GetChange("neighbor_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NeighborGroup = &[]models.RouterBgpNeighborGroup{}
		}
	}
	if v, ok := d.GetOk("neighbor_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor_range", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNeighborRange(d, v, "neighbor_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NeighborRange = t
		}
	} else if d.HasChange("neighbor_range") {
		old, new := d.GetChange("neighbor_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NeighborRange = &[]models.RouterBgpNeighborRange{}
		}
	}
	if v, ok := d.GetOk("neighbor_range6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor_range6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNeighborRange6(d, v, "neighbor_range6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NeighborRange6 = t
		}
	} else if d.HasChange("neighbor_range6") {
		old, new := d.GetChange("neighbor_range6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NeighborRange6 = &[]models.RouterBgpNeighborRange6{}
		}
	}
	if v, ok := d.GetOk("network"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("network", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Network = t
		}
	} else if d.HasChange("network") {
		old, new := d.GetChange("network")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Network = &[]models.RouterBgpNetwork{}
		}
	}
	if v1, ok := d.GetOk("network_import_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("network_import_check", sv)
				diags = append(diags, e)
			}
			obj.NetworkImportCheck = &v2
		}
	}
	if v, ok := d.GetOk("network6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("network6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNetwork6(d, v, "network6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Network6 = t
		}
	} else if d.HasChange("network6") {
		old, new := d.GetChange("network6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Network6 = &[]models.RouterBgpNetwork6{}
		}
	}
	if v1, ok := d.GetOk("recursive_inherit_priority"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.1", "") {
				e := utils.AttributeVersionWarning("recursive_inherit_priority", sv)
				diags = append(diags, e)
			}
			obj.RecursiveInheritPriority = &v2
		}
	}
	if v1, ok := d.GetOk("recursive_next_hop"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("recursive_next_hop", sv)
				diags = append(diags, e)
			}
			obj.RecursiveNextHop = &v2
		}
	}
	if v, ok := d.GetOk("redistribute"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpRedistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute = &[]models.RouterBgpRedistribute{}
		}
	}
	if v, ok := d.GetOk("redistribute6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpRedistribute6(d, v, "redistribute6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute6 = t
		}
	} else if d.HasChange("redistribute6") {
		old, new := d.GetChange("redistribute6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute6 = &[]models.RouterBgpRedistribute6{}
		}
	}
	if v1, ok := d.GetOk("router_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("router_id", sv)
				diags = append(diags, e)
			}
			obj.RouterId = &v2
		}
	}
	if v1, ok := d.GetOk("scan_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scan_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ScanTime = &tmp
		}
	}
	if v1, ok := d.GetOk("synchronization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("synchronization", sv)
				diags = append(diags, e)
			}
			obj.Synchronization = &v2
		}
	}
	if v1, ok := d.GetOk("tag_resolve_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("tag_resolve_mode", sv)
				diags = append(diags, e)
			}
			obj.TagResolveMode = &v2
		}
	}
	if v, ok := d.GetOk("vrf"); ok {
		if !utils.CheckVer(sv, "v7.2.0", "") {
			e := utils.AttributeVersionWarning("vrf", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrf(d, v, "vrf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vrf = t
		}
	} else if d.HasChange("vrf") {
		old, new := d.GetChange("vrf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vrf = &[]models.RouterBgpVrf{}
		}
	}
	if v, ok := d.GetOk("vrf_leak"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "v7.2.0") {
			e := utils.AttributeVersionWarning("vrf_leak", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrfLeak(d, v, "vrf_leak", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VrfLeak = t
		}
	} else if d.HasChange("vrf_leak") {
		old, new := d.GetChange("vrf_leak")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VrfLeak = &[]models.RouterBgpVrfLeak{}
		}
	}
	if v, ok := d.GetOk("vrf_leak6"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "v7.2.0") {
			e := utils.AttributeVersionWarning("vrf_leak6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrfLeak6(d, v, "vrf_leak6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VrfLeak6 = t
		}
	} else if d.HasChange("vrf_leak6") {
		old, new := d.GetChange("vrf_leak6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VrfLeak6 = &[]models.RouterBgpVrfLeak6{}
		}
	}
	if v, ok := d.GetOk("vrf6"); ok {
		if !utils.CheckVer(sv, "v7.2.0", "") {
			e := utils.AttributeVersionWarning("vrf6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpVrf6(d, v, "vrf6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vrf6 = t
		}
	} else if d.HasChange("vrf6") {
		old, new := d.GetChange("vrf6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vrf6 = &[]models.RouterBgpVrf6{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterBgp(d *schema.ResourceData, sv string) (*models.RouterBgp, diag.Diagnostics) {
	obj := models.RouterBgp{}
	diags := diag.Diagnostics{}

	obj.AdminDistance = &[]models.RouterBgpAdminDistance{}
	obj.AggregateAddress = &[]models.RouterBgpAggregateAddress{}
	obj.AggregateAddress6 = &[]models.RouterBgpAggregateAddress6{}
	obj.ConfederationPeers = &[]models.RouterBgpConfederationPeers{}
	obj.Neighbor = &[]models.RouterBgpNeighbor{}
	obj.NeighborGroup = &[]models.RouterBgpNeighborGroup{}
	obj.NeighborRange = &[]models.RouterBgpNeighborRange{}
	obj.NeighborRange6 = &[]models.RouterBgpNeighborRange6{}
	obj.Network = &[]models.RouterBgpNetwork{}
	obj.Network6 = &[]models.RouterBgpNetwork6{}
	obj.Redistribute = &[]models.RouterBgpRedistribute{}
	obj.Redistribute6 = &[]models.RouterBgpRedistribute6{}
	obj.Vrf = &[]models.RouterBgpVrf{}
	obj.VrfLeak = &[]models.RouterBgpVrfLeak{}
	obj.VrfLeak6 = &[]models.RouterBgpVrfLeak6{}
	obj.Vrf6 = &[]models.RouterBgpVrf6{}

	return &obj, diags
}
