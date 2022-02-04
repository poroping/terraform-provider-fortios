// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP neighbor table.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceRouterBgpNeighbor() *schema.Resource {
	return &schema.Resource{
		Description: "BGP neighbor table.",

		CreateContext: resourceRouterBgpNeighborCreate,
		ReadContext:   resourceRouterBgpNeighborRead,
		UpdateContext: resourceRouterBgpNeighborUpdate,
		DeleteContext: resourceRouterBgpNeighborDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"activate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable address family IPv4 for this neighbor.",
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

										Description: "route map",
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

										Description: "route map",
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
				ForceNew:    true,
				Required:    true,
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
				Type: schema.TypeInt,

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
			"maximum_prefix_threshold6": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "Maximum IPv6 prefix threshold value (1 - 100 percent).",
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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable remove private AS number from IPv4 outbound updates.",
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
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv4 AS route reflector client.",
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
	}
}

func resourceRouterBgpNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "ip"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterBgpNeighbor resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterBgpNeighbor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterBgpNeighbor(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpNeighbor")
	}

	return resourceRouterBgpNeighborRead(ctx, d, meta)
}

func resourceRouterBgpNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterBgpNeighbor(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterBgpNeighbor(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterBgpNeighbor resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterBgpNeighbor")
	}

	return resourceRouterBgpNeighborRead(ctx, d, meta)
}

func resourceRouterBgpNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterBgpNeighbor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterBgpNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpNeighbor(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpNeighbor resource: %v", err)
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

	diags := refreshObjectRouterBgpNeighbor(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterBgpNeighbor(d *schema.ResourceData, o *models.RouterBgpNeighbor, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Activate != nil {
		v := *o.Activate

		if err = d.Set("activate", v); err != nil {
			return diag.Errorf("error reading activate: %v", err)
		}
	}

	if o.Activate6 != nil {
		v := *o.Activate6

		if err = d.Set("activate6", v); err != nil {
			return diag.Errorf("error reading activate6: %v", err)
		}
	}

	if o.AdditionalPath != nil {
		v := *o.AdditionalPath

		if err = d.Set("additional_path", v); err != nil {
			return diag.Errorf("error reading additional_path: %v", err)
		}
	}

	if o.AdditionalPath6 != nil {
		v := *o.AdditionalPath6

		if err = d.Set("additional_path6", v); err != nil {
			return diag.Errorf("error reading additional_path6: %v", err)
		}
	}

	if o.AdvAdditionalPath != nil {
		v := *o.AdvAdditionalPath

		if err = d.Set("adv_additional_path", v); err != nil {
			return diag.Errorf("error reading adv_additional_path: %v", err)
		}
	}

	if o.AdvAdditionalPath6 != nil {
		v := *o.AdvAdditionalPath6

		if err = d.Set("adv_additional_path6", v); err != nil {
			return diag.Errorf("error reading adv_additional_path6: %v", err)
		}
	}

	if o.AdvertisementInterval != nil {
		v := *o.AdvertisementInterval

		if err = d.Set("advertisement_interval", v); err != nil {
			return diag.Errorf("error reading advertisement_interval: %v", err)
		}
	}

	if o.AllowasIn != nil {
		v := *o.AllowasIn

		if err = d.Set("allowas_in", v); err != nil {
			return diag.Errorf("error reading allowas_in: %v", err)
		}
	}

	if o.AllowasInEnable != nil {
		v := *o.AllowasInEnable

		if err = d.Set("allowas_in_enable", v); err != nil {
			return diag.Errorf("error reading allowas_in_enable: %v", err)
		}
	}

	if o.AllowasInEnable6 != nil {
		v := *o.AllowasInEnable6

		if err = d.Set("allowas_in_enable6", v); err != nil {
			return diag.Errorf("error reading allowas_in_enable6: %v", err)
		}
	}

	if o.AllowasIn6 != nil {
		v := *o.AllowasIn6

		if err = d.Set("allowas_in6", v); err != nil {
			return diag.Errorf("error reading allowas_in6: %v", err)
		}
	}

	if o.AsOverride != nil {
		v := *o.AsOverride

		if err = d.Set("as_override", v); err != nil {
			return diag.Errorf("error reading as_override: %v", err)
		}
	}

	if o.AsOverride6 != nil {
		v := *o.AsOverride6

		if err = d.Set("as_override6", v); err != nil {
			return diag.Errorf("error reading as_override6: %v", err)
		}
	}

	if o.AttributeUnchanged != nil {
		v := *o.AttributeUnchanged

		if err = d.Set("attribute_unchanged", v); err != nil {
			return diag.Errorf("error reading attribute_unchanged: %v", err)
		}
	}

	if o.AttributeUnchanged6 != nil {
		v := *o.AttributeUnchanged6

		if err = d.Set("attribute_unchanged6", v); err != nil {
			return diag.Errorf("error reading attribute_unchanged6: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.CapabilityDefaultOriginate != nil {
		v := *o.CapabilityDefaultOriginate

		if err = d.Set("capability_default_originate", v); err != nil {
			return diag.Errorf("error reading capability_default_originate: %v", err)
		}
	}

	if o.CapabilityDefaultOriginate6 != nil {
		v := *o.CapabilityDefaultOriginate6

		if err = d.Set("capability_default_originate6", v); err != nil {
			return diag.Errorf("error reading capability_default_originate6: %v", err)
		}
	}

	if o.CapabilityDynamic != nil {
		v := *o.CapabilityDynamic

		if err = d.Set("capability_dynamic", v); err != nil {
			return diag.Errorf("error reading capability_dynamic: %v", err)
		}
	}

	if o.CapabilityGracefulRestart != nil {
		v := *o.CapabilityGracefulRestart

		if err = d.Set("capability_graceful_restart", v); err != nil {
			return diag.Errorf("error reading capability_graceful_restart: %v", err)
		}
	}

	if o.CapabilityGracefulRestart6 != nil {
		v := *o.CapabilityGracefulRestart6

		if err = d.Set("capability_graceful_restart6", v); err != nil {
			return diag.Errorf("error reading capability_graceful_restart6: %v", err)
		}
	}

	if o.CapabilityOrf != nil {
		v := *o.CapabilityOrf

		if err = d.Set("capability_orf", v); err != nil {
			return diag.Errorf("error reading capability_orf: %v", err)
		}
	}

	if o.CapabilityOrf6 != nil {
		v := *o.CapabilityOrf6

		if err = d.Set("capability_orf6", v); err != nil {
			return diag.Errorf("error reading capability_orf6: %v", err)
		}
	}

	if o.CapabilityRouteRefresh != nil {
		v := *o.CapabilityRouteRefresh

		if err = d.Set("capability_route_refresh", v); err != nil {
			return diag.Errorf("error reading capability_route_refresh: %v", err)
		}
	}

	if o.ConditionalAdvertise != nil {
		if err = d.Set("conditional_advertise", flattenRouterBgpNeighborConditionalAdvertise(o.ConditionalAdvertise, sort)); err != nil {
			return diag.Errorf("error reading conditional_advertise: %v", err)
		}
	}

	if o.ConditionalAdvertise6 != nil {
		if err = d.Set("conditional_advertise6", flattenRouterBgpNeighborConditionalAdvertise6(o.ConditionalAdvertise6, sort)); err != nil {
			return diag.Errorf("error reading conditional_advertise6: %v", err)
		}
	}

	if o.ConnectTimer != nil {
		v := *o.ConnectTimer

		if err = d.Set("connect_timer", v); err != nil {
			return diag.Errorf("error reading connect_timer: %v", err)
		}
	}

	if o.DefaultOriginateRoutemap != nil {
		v := *o.DefaultOriginateRoutemap

		if err = d.Set("default_originate_routemap", v); err != nil {
			return diag.Errorf("error reading default_originate_routemap: %v", err)
		}
	}

	if o.DefaultOriginateRoutemap6 != nil {
		v := *o.DefaultOriginateRoutemap6

		if err = d.Set("default_originate_routemap6", v); err != nil {
			return diag.Errorf("error reading default_originate_routemap6: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.DistributeListIn != nil {
		v := *o.DistributeListIn

		if err = d.Set("distribute_list_in", v); err != nil {
			return diag.Errorf("error reading distribute_list_in: %v", err)
		}
	}

	if o.DistributeListIn6 != nil {
		v := *o.DistributeListIn6

		if err = d.Set("distribute_list_in6", v); err != nil {
			return diag.Errorf("error reading distribute_list_in6: %v", err)
		}
	}

	if o.DistributeListOut != nil {
		v := *o.DistributeListOut

		if err = d.Set("distribute_list_out", v); err != nil {
			return diag.Errorf("error reading distribute_list_out: %v", err)
		}
	}

	if o.DistributeListOut6 != nil {
		v := *o.DistributeListOut6

		if err = d.Set("distribute_list_out6", v); err != nil {
			return diag.Errorf("error reading distribute_list_out6: %v", err)
		}
	}

	if o.DontCapabilityNegotiate != nil {
		v := *o.DontCapabilityNegotiate

		if err = d.Set("dont_capability_negotiate", v); err != nil {
			return diag.Errorf("error reading dont_capability_negotiate: %v", err)
		}
	}

	if o.EbgpEnforceMultihop != nil {
		v := *o.EbgpEnforceMultihop

		if err = d.Set("ebgp_enforce_multihop", v); err != nil {
			return diag.Errorf("error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if o.EbgpMultihopTtl != nil {
		v := *o.EbgpMultihopTtl

		if err = d.Set("ebgp_multihop_ttl", v); err != nil {
			return diag.Errorf("error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if o.FilterListIn != nil {
		v := *o.FilterListIn

		if err = d.Set("filter_list_in", v); err != nil {
			return diag.Errorf("error reading filter_list_in: %v", err)
		}
	}

	if o.FilterListIn6 != nil {
		v := *o.FilterListIn6

		if err = d.Set("filter_list_in6", v); err != nil {
			return diag.Errorf("error reading filter_list_in6: %v", err)
		}
	}

	if o.FilterListOut != nil {
		v := *o.FilterListOut

		if err = d.Set("filter_list_out", v); err != nil {
			return diag.Errorf("error reading filter_list_out: %v", err)
		}
	}

	if o.FilterListOut6 != nil {
		v := *o.FilterListOut6

		if err = d.Set("filter_list_out6", v); err != nil {
			return diag.Errorf("error reading filter_list_out6: %v", err)
		}
	}

	if o.HoldtimeTimer != nil {
		v := *o.HoldtimeTimer

		if err = d.Set("holdtime_timer", v); err != nil {
			return diag.Errorf("error reading holdtime_timer: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.KeepAliveTimer != nil {
		v := *o.KeepAliveTimer

		if err = d.Set("keep_alive_timer", v); err != nil {
			return diag.Errorf("error reading keep_alive_timer: %v", err)
		}
	}

	if o.LinkDownFailover != nil {
		v := *o.LinkDownFailover

		if err = d.Set("link_down_failover", v); err != nil {
			return diag.Errorf("error reading link_down_failover: %v", err)
		}
	}

	if o.LocalAs != nil {
		v := *o.LocalAs

		if err = d.Set("local_as", v); err != nil {
			return diag.Errorf("error reading local_as: %v", err)
		}
	}

	if o.LocalAsNoPrepend != nil {
		v := *o.LocalAsNoPrepend

		if err = d.Set("local_as_no_prepend", v); err != nil {
			return diag.Errorf("error reading local_as_no_prepend: %v", err)
		}
	}

	if o.LocalAsReplaceAs != nil {
		v := *o.LocalAsReplaceAs

		if err = d.Set("local_as_replace_as", v); err != nil {
			return diag.Errorf("error reading local_as_replace_as: %v", err)
		}
	}

	if o.MaximumPrefix != nil {
		v := *o.MaximumPrefix

		if err = d.Set("maximum_prefix", v); err != nil {
			return diag.Errorf("error reading maximum_prefix: %v", err)
		}
	}

	if o.MaximumPrefixThreshold != nil {
		v := *o.MaximumPrefixThreshold

		if err = d.Set("maximum_prefix_threshold", v); err != nil {
			return diag.Errorf("error reading maximum_prefix_threshold: %v", err)
		}
	}

	if o.MaximumPrefixThreshold6 != nil {
		v := *o.MaximumPrefixThreshold6

		if err = d.Set("maximum_prefix_threshold6", v); err != nil {
			return diag.Errorf("error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if o.MaximumPrefixWarningOnly != nil {
		v := *o.MaximumPrefixWarningOnly

		if err = d.Set("maximum_prefix_warning_only", v); err != nil {
			return diag.Errorf("error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if o.MaximumPrefixWarningOnly6 != nil {
		v := *o.MaximumPrefixWarningOnly6

		if err = d.Set("maximum_prefix_warning_only6", v); err != nil {
			return diag.Errorf("error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if o.MaximumPrefix6 != nil {
		v := *o.MaximumPrefix6

		if err = d.Set("maximum_prefix6", v); err != nil {
			return diag.Errorf("error reading maximum_prefix6: %v", err)
		}
	}

	if o.NextHopSelf != nil {
		v := *o.NextHopSelf

		if err = d.Set("next_hop_self", v); err != nil {
			return diag.Errorf("error reading next_hop_self: %v", err)
		}
	}

	if o.NextHopSelfRr != nil {
		v := *o.NextHopSelfRr

		if err = d.Set("next_hop_self_rr", v); err != nil {
			return diag.Errorf("error reading next_hop_self_rr: %v", err)
		}
	}

	if o.NextHopSelfRr6 != nil {
		v := *o.NextHopSelfRr6

		if err = d.Set("next_hop_self_rr6", v); err != nil {
			return diag.Errorf("error reading next_hop_self_rr6: %v", err)
		}
	}

	if o.NextHopSelf6 != nil {
		v := *o.NextHopSelf6

		if err = d.Set("next_hop_self6", v); err != nil {
			return diag.Errorf("error reading next_hop_self6: %v", err)
		}
	}

	if o.OverrideCapability != nil {
		v := *o.OverrideCapability

		if err = d.Set("override_capability", v); err != nil {
			return diag.Errorf("error reading override_capability: %v", err)
		}
	}

	if o.Passive != nil {
		v := *o.Passive

		if err = d.Set("passive", v); err != nil {
			return diag.Errorf("error reading passive: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PrefixListIn != nil {
		v := *o.PrefixListIn

		if err = d.Set("prefix_list_in", v); err != nil {
			return diag.Errorf("error reading prefix_list_in: %v", err)
		}
	}

	if o.PrefixListIn6 != nil {
		v := *o.PrefixListIn6

		if err = d.Set("prefix_list_in6", v); err != nil {
			return diag.Errorf("error reading prefix_list_in6: %v", err)
		}
	}

	if o.PrefixListOut != nil {
		v := *o.PrefixListOut

		if err = d.Set("prefix_list_out", v); err != nil {
			return diag.Errorf("error reading prefix_list_out: %v", err)
		}
	}

	if o.PrefixListOut6 != nil {
		v := *o.PrefixListOut6

		if err = d.Set("prefix_list_out6", v); err != nil {
			return diag.Errorf("error reading prefix_list_out6: %v", err)
		}
	}

	if o.RemoteAs != nil {
		v := *o.RemoteAs

		if err = d.Set("remote_as", v); err != nil {
			return diag.Errorf("error reading remote_as: %v", err)
		}
	}

	if o.RemovePrivateAs != nil {
		v := *o.RemovePrivateAs

		if err = d.Set("remove_private_as", v); err != nil {
			return diag.Errorf("error reading remove_private_as: %v", err)
		}
	}

	if o.RemovePrivateAs6 != nil {
		v := *o.RemovePrivateAs6

		if err = d.Set("remove_private_as6", v); err != nil {
			return diag.Errorf("error reading remove_private_as6: %v", err)
		}
	}

	if o.RestartTime != nil {
		v := *o.RestartTime

		if err = d.Set("restart_time", v); err != nil {
			return diag.Errorf("error reading restart_time: %v", err)
		}
	}

	if o.RetainStaleTime != nil {
		v := *o.RetainStaleTime

		if err = d.Set("retain_stale_time", v); err != nil {
			return diag.Errorf("error reading retain_stale_time: %v", err)
		}
	}

	if o.RouteMapIn != nil {
		v := *o.RouteMapIn

		if err = d.Set("route_map_in", v); err != nil {
			return diag.Errorf("error reading route_map_in: %v", err)
		}
	}

	if o.RouteMapIn6 != nil {
		v := *o.RouteMapIn6

		if err = d.Set("route_map_in6", v); err != nil {
			return diag.Errorf("error reading route_map_in6: %v", err)
		}
	}

	if o.RouteMapOut != nil {
		v := *o.RouteMapOut

		if err = d.Set("route_map_out", v); err != nil {
			return diag.Errorf("error reading route_map_out: %v", err)
		}
	}

	if o.RouteMapOutPreferable != nil {
		v := *o.RouteMapOutPreferable

		if err = d.Set("route_map_out_preferable", v); err != nil {
			return diag.Errorf("error reading route_map_out_preferable: %v", err)
		}
	}

	if o.RouteMapOut6 != nil {
		v := *o.RouteMapOut6

		if err = d.Set("route_map_out6", v); err != nil {
			return diag.Errorf("error reading route_map_out6: %v", err)
		}
	}

	if o.RouteMapOut6Preferable != nil {
		v := *o.RouteMapOut6Preferable

		if err = d.Set("route_map_out6_preferable", v); err != nil {
			return diag.Errorf("error reading route_map_out6_preferable: %v", err)
		}
	}

	if o.RouteReflectorClient != nil {
		v := *o.RouteReflectorClient

		if err = d.Set("route_reflector_client", v); err != nil {
			return diag.Errorf("error reading route_reflector_client: %v", err)
		}
	}

	if o.RouteReflectorClient6 != nil {
		v := *o.RouteReflectorClient6

		if err = d.Set("route_reflector_client6", v); err != nil {
			return diag.Errorf("error reading route_reflector_client6: %v", err)
		}
	}

	if o.RouteServerClient != nil {
		v := *o.RouteServerClient

		if err = d.Set("route_server_client", v); err != nil {
			return diag.Errorf("error reading route_server_client: %v", err)
		}
	}

	if o.RouteServerClient6 != nil {
		v := *o.RouteServerClient6

		if err = d.Set("route_server_client6", v); err != nil {
			return diag.Errorf("error reading route_server_client6: %v", err)
		}
	}

	if o.SendCommunity != nil {
		v := *o.SendCommunity

		if err = d.Set("send_community", v); err != nil {
			return diag.Errorf("error reading send_community: %v", err)
		}
	}

	if o.SendCommunity6 != nil {
		v := *o.SendCommunity6

		if err = d.Set("send_community6", v); err != nil {
			return diag.Errorf("error reading send_community6: %v", err)
		}
	}

	if o.Shutdown != nil {
		v := *o.Shutdown

		if err = d.Set("shutdown", v); err != nil {
			return diag.Errorf("error reading shutdown: %v", err)
		}
	}

	if o.SoftReconfiguration != nil {
		v := *o.SoftReconfiguration

		if err = d.Set("soft_reconfiguration", v); err != nil {
			return diag.Errorf("error reading soft_reconfiguration: %v", err)
		}
	}

	if o.SoftReconfiguration6 != nil {
		v := *o.SoftReconfiguration6

		if err = d.Set("soft_reconfiguration6", v); err != nil {
			return diag.Errorf("error reading soft_reconfiguration6: %v", err)
		}
	}

	if o.StaleRoute != nil {
		v := *o.StaleRoute

		if err = d.Set("stale_route", v); err != nil {
			return diag.Errorf("error reading stale_route: %v", err)
		}
	}

	if o.StrictCapabilityMatch != nil {
		v := *o.StrictCapabilityMatch

		if err = d.Set("strict_capability_match", v); err != nil {
			return diag.Errorf("error reading strict_capability_match: %v", err)
		}
	}

	if o.UnsuppressMap != nil {
		v := *o.UnsuppressMap

		if err = d.Set("unsuppress_map", v); err != nil {
			return diag.Errorf("error reading unsuppress_map: %v", err)
		}
	}

	if o.UnsuppressMap6 != nil {
		v := *o.UnsuppressMap6

		if err = d.Set("unsuppress_map6", v); err != nil {
			return diag.Errorf("error reading unsuppress_map6: %v", err)
		}
	}

	if o.UpdateSource != nil {
		v := *o.UpdateSource

		if err = d.Set("update_source", v); err != nil {
			return diag.Errorf("error reading update_source: %v", err)
		}
	}

	if o.Weight != nil {
		v := *o.Weight

		if err = d.Set("weight", v); err != nil {
			return diag.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}

func getObjectRouterBgpNeighbor(d *schema.ResourceData, sv string) (*models.RouterBgpNeighbor, diag.Diagnostics) {
	obj := models.RouterBgpNeighbor{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("activate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("activate", sv)
				diags = append(diags, e)
			}
			obj.Activate = &v2
		}
	}
	if v1, ok := d.GetOk("activate6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("activate6", sv)
				diags = append(diags, e)
			}
			obj.Activate6 = &v2
		}
	}
	if v1, ok := d.GetOk("additional_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("additional_path", sv)
				diags = append(diags, e)
			}
			obj.AdditionalPath = &v2
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
	if v1, ok := d.GetOk("adv_additional_path"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("adv_additional_path", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdvAdditionalPath = &tmp
		}
	}
	if v1, ok := d.GetOk("adv_additional_path6"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("adv_additional_path6", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdvAdditionalPath6 = &tmp
		}
	}
	if v1, ok := d.GetOk("advertisement_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("advertisement_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AdvertisementInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("allowas_in"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowas_in", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AllowasIn = &tmp
		}
	}
	if v1, ok := d.GetOk("allowas_in_enable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowas_in_enable", sv)
				diags = append(diags, e)
			}
			obj.AllowasInEnable = &v2
		}
	}
	if v1, ok := d.GetOk("allowas_in_enable6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowas_in_enable6", sv)
				diags = append(diags, e)
			}
			obj.AllowasInEnable6 = &v2
		}
	}
	if v1, ok := d.GetOk("allowas_in6"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowas_in6", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AllowasIn6 = &tmp
		}
	}
	if v1, ok := d.GetOk("as_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("as_override", sv)
				diags = append(diags, e)
			}
			obj.AsOverride = &v2
		}
	}
	if v1, ok := d.GetOk("as_override6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("as_override6", sv)
				diags = append(diags, e)
			}
			obj.AsOverride6 = &v2
		}
	}
	if v1, ok := d.GetOk("attribute_unchanged"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("attribute_unchanged", sv)
				diags = append(diags, e)
			}
			obj.AttributeUnchanged = &v2
		}
	}
	if v1, ok := d.GetOk("attribute_unchanged6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("attribute_unchanged6", sv)
				diags = append(diags, e)
			}
			obj.AttributeUnchanged6 = &v2
		}
	}
	if v1, ok := d.GetOk("bfd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd", sv)
				diags = append(diags, e)
			}
			obj.Bfd = &v2
		}
	}
	if v1, ok := d.GetOk("capability_default_originate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_default_originate", sv)
				diags = append(diags, e)
			}
			obj.CapabilityDefaultOriginate = &v2
		}
	}
	if v1, ok := d.GetOk("capability_default_originate6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_default_originate6", sv)
				diags = append(diags, e)
			}
			obj.CapabilityDefaultOriginate6 = &v2
		}
	}
	if v1, ok := d.GetOk("capability_dynamic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_dynamic", sv)
				diags = append(diags, e)
			}
			obj.CapabilityDynamic = &v2
		}
	}
	if v1, ok := d.GetOk("capability_graceful_restart"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_graceful_restart", sv)
				diags = append(diags, e)
			}
			obj.CapabilityGracefulRestart = &v2
		}
	}
	if v1, ok := d.GetOk("capability_graceful_restart6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_graceful_restart6", sv)
				diags = append(diags, e)
			}
			obj.CapabilityGracefulRestart6 = &v2
		}
	}
	if v1, ok := d.GetOk("capability_orf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_orf", sv)
				diags = append(diags, e)
			}
			obj.CapabilityOrf = &v2
		}
	}
	if v1, ok := d.GetOk("capability_orf6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_orf6", sv)
				diags = append(diags, e)
			}
			obj.CapabilityOrf6 = &v2
		}
	}
	if v1, ok := d.GetOk("capability_route_refresh"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capability_route_refresh", sv)
				diags = append(diags, e)
			}
			obj.CapabilityRouteRefresh = &v2
		}
	}
	if v, ok := d.GetOk("conditional_advertise"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("conditional_advertise", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNeighborConditionalAdvertise(d, v, "conditional_advertise", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ConditionalAdvertise = t
		}
	} else if d.HasChange("conditional_advertise") {
		old, new := d.GetChange("conditional_advertise")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ConditionalAdvertise = &[]models.RouterBgpNeighborConditionalAdvertise{}
		}
	}
	if v, ok := d.GetOk("conditional_advertise6"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("conditional_advertise6", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterBgpNeighborConditionalAdvertise6(d, v, "conditional_advertise6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ConditionalAdvertise6 = t
		}
	} else if d.HasChange("conditional_advertise6") {
		old, new := d.GetChange("conditional_advertise6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ConditionalAdvertise6 = &[]models.RouterBgpNeighborConditionalAdvertise6{}
		}
	}
	if v1, ok := d.GetOk("connect_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("connect_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ConnectTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("default_originate_routemap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_originate_routemap", sv)
				diags = append(diags, e)
			}
			obj.DefaultOriginateRoutemap = &v2
		}
	}
	if v1, ok := d.GetOk("default_originate_routemap6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_originate_routemap6", sv)
				diags = append(diags, e)
			}
			obj.DefaultOriginateRoutemap6 = &v2
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("distribute_list_in"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distribute_list_in", sv)
				diags = append(diags, e)
			}
			obj.DistributeListIn = &v2
		}
	}
	if v1, ok := d.GetOk("distribute_list_in6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distribute_list_in6", sv)
				diags = append(diags, e)
			}
			obj.DistributeListIn6 = &v2
		}
	}
	if v1, ok := d.GetOk("distribute_list_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distribute_list_out", sv)
				diags = append(diags, e)
			}
			obj.DistributeListOut = &v2
		}
	}
	if v1, ok := d.GetOk("distribute_list_out6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distribute_list_out6", sv)
				diags = append(diags, e)
			}
			obj.DistributeListOut6 = &v2
		}
	}
	if v1, ok := d.GetOk("dont_capability_negotiate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dont_capability_negotiate", sv)
				diags = append(diags, e)
			}
			obj.DontCapabilityNegotiate = &v2
		}
	}
	if v1, ok := d.GetOk("ebgp_enforce_multihop"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ebgp_enforce_multihop", sv)
				diags = append(diags, e)
			}
			obj.EbgpEnforceMultihop = &v2
		}
	}
	if v1, ok := d.GetOk("ebgp_multihop_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ebgp_multihop_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EbgpMultihopTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("filter_list_in"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filter_list_in", sv)
				diags = append(diags, e)
			}
			obj.FilterListIn = &v2
		}
	}
	if v1, ok := d.GetOk("filter_list_in6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filter_list_in6", sv)
				diags = append(diags, e)
			}
			obj.FilterListIn6 = &v2
		}
	}
	if v1, ok := d.GetOk("filter_list_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filter_list_out", sv)
				diags = append(diags, e)
			}
			obj.FilterListOut = &v2
		}
	}
	if v1, ok := d.GetOk("filter_list_out6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filter_list_out6", sv)
				diags = append(diags, e)
			}
			obj.FilterListOut6 = &v2
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
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("keep_alive_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keep_alive_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeepAliveTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("link_down_failover"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_down_failover", sv)
				diags = append(diags, e)
			}
			obj.LinkDownFailover = &v2
		}
	}
	if v1, ok := d.GetOk("local_as"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_as", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LocalAs = &tmp
		}
	}
	if v1, ok := d.GetOk("local_as_no_prepend"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_as_no_prepend", sv)
				diags = append(diags, e)
			}
			obj.LocalAsNoPrepend = &v2
		}
	}
	if v1, ok := d.GetOk("local_as_replace_as"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_as_replace_as", sv)
				diags = append(diags, e)
			}
			obj.LocalAsReplaceAs = &v2
		}
	}
	if v1, ok := d.GetOk("maximum_prefix"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_prefix", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaximumPrefix = &tmp
		}
	}
	if v1, ok := d.GetOk("maximum_prefix_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_prefix_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaximumPrefixThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("maximum_prefix_threshold6"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_prefix_threshold6", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaximumPrefixThreshold6 = &tmp
		}
	}
	if v1, ok := d.GetOk("maximum_prefix_warning_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_prefix_warning_only", sv)
				diags = append(diags, e)
			}
			obj.MaximumPrefixWarningOnly = &v2
		}
	}
	if v1, ok := d.GetOk("maximum_prefix_warning_only6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_prefix_warning_only6", sv)
				diags = append(diags, e)
			}
			obj.MaximumPrefixWarningOnly6 = &v2
		}
	}
	if v1, ok := d.GetOk("maximum_prefix6"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maximum_prefix6", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaximumPrefix6 = &tmp
		}
	}
	if v1, ok := d.GetOk("next_hop_self"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("next_hop_self", sv)
				diags = append(diags, e)
			}
			obj.NextHopSelf = &v2
		}
	}
	if v1, ok := d.GetOk("next_hop_self_rr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("next_hop_self_rr", sv)
				diags = append(diags, e)
			}
			obj.NextHopSelfRr = &v2
		}
	}
	if v1, ok := d.GetOk("next_hop_self_rr6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("next_hop_self_rr6", sv)
				diags = append(diags, e)
			}
			obj.NextHopSelfRr6 = &v2
		}
	}
	if v1, ok := d.GetOk("next_hop_self6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("next_hop_self6", sv)
				diags = append(diags, e)
			}
			obj.NextHopSelf6 = &v2
		}
	}
	if v1, ok := d.GetOk("override_capability"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_capability", sv)
				diags = append(diags, e)
			}
			obj.OverrideCapability = &v2
		}
	}
	if v1, ok := d.GetOk("passive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("passive", sv)
				diags = append(diags, e)
			}
			obj.Passive = &v2
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("prefix_list_in"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix_list_in", sv)
				diags = append(diags, e)
			}
			obj.PrefixListIn = &v2
		}
	}
	if v1, ok := d.GetOk("prefix_list_in6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix_list_in6", sv)
				diags = append(diags, e)
			}
			obj.PrefixListIn6 = &v2
		}
	}
	if v1, ok := d.GetOk("prefix_list_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix_list_out", sv)
				diags = append(diags, e)
			}
			obj.PrefixListOut = &v2
		}
	}
	if v1, ok := d.GetOk("prefix_list_out6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix_list_out6", sv)
				diags = append(diags, e)
			}
			obj.PrefixListOut6 = &v2
		}
	}
	if v1, ok := d.GetOk("remote_as"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_as", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RemoteAs = &tmp
		}
	}
	if v1, ok := d.GetOk("remove_private_as"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remove_private_as", sv)
				diags = append(diags, e)
			}
			obj.RemovePrivateAs = &v2
		}
	}
	if v1, ok := d.GetOk("remove_private_as6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remove_private_as6", sv)
				diags = append(diags, e)
			}
			obj.RemovePrivateAs6 = &v2
		}
	}
	if v1, ok := d.GetOk("restart_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("restart_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RestartTime = &tmp
		}
	}
	if v1, ok := d.GetOk("retain_stale_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("retain_stale_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RetainStaleTime = &tmp
		}
	}
	if v1, ok := d.GetOk("route_map_in"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_map_in", sv)
				diags = append(diags, e)
			}
			obj.RouteMapIn = &v2
		}
	}
	if v1, ok := d.GetOk("route_map_in6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_map_in6", sv)
				diags = append(diags, e)
			}
			obj.RouteMapIn6 = &v2
		}
	}
	if v1, ok := d.GetOk("route_map_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_map_out", sv)
				diags = append(diags, e)
			}
			obj.RouteMapOut = &v2
		}
	}
	if v1, ok := d.GetOk("route_map_out_preferable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_map_out_preferable", sv)
				diags = append(diags, e)
			}
			obj.RouteMapOutPreferable = &v2
		}
	}
	if v1, ok := d.GetOk("route_map_out6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_map_out6", sv)
				diags = append(diags, e)
			}
			obj.RouteMapOut6 = &v2
		}
	}
	if v1, ok := d.GetOk("route_map_out6_preferable"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_map_out6_preferable", sv)
				diags = append(diags, e)
			}
			obj.RouteMapOut6Preferable = &v2
		}
	}
	if v1, ok := d.GetOk("route_reflector_client"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_reflector_client", sv)
				diags = append(diags, e)
			}
			obj.RouteReflectorClient = &v2
		}
	}
	if v1, ok := d.GetOk("route_reflector_client6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_reflector_client6", sv)
				diags = append(diags, e)
			}
			obj.RouteReflectorClient6 = &v2
		}
	}
	if v1, ok := d.GetOk("route_server_client"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_server_client", sv)
				diags = append(diags, e)
			}
			obj.RouteServerClient = &v2
		}
	}
	if v1, ok := d.GetOk("route_server_client6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_server_client6", sv)
				diags = append(diags, e)
			}
			obj.RouteServerClient6 = &v2
		}
	}
	if v1, ok := d.GetOk("send_community"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_community", sv)
				diags = append(diags, e)
			}
			obj.SendCommunity = &v2
		}
	}
	if v1, ok := d.GetOk("send_community6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_community6", sv)
				diags = append(diags, e)
			}
			obj.SendCommunity6 = &v2
		}
	}
	if v1, ok := d.GetOk("shutdown"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("shutdown", sv)
				diags = append(diags, e)
			}
			obj.Shutdown = &v2
		}
	}
	if v1, ok := d.GetOk("soft_reconfiguration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("soft_reconfiguration", sv)
				diags = append(diags, e)
			}
			obj.SoftReconfiguration = &v2
		}
	}
	if v1, ok := d.GetOk("soft_reconfiguration6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("soft_reconfiguration6", sv)
				diags = append(diags, e)
			}
			obj.SoftReconfiguration6 = &v2
		}
	}
	if v1, ok := d.GetOk("stale_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("stale_route", sv)
				diags = append(diags, e)
			}
			obj.StaleRoute = &v2
		}
	}
	if v1, ok := d.GetOk("strict_capability_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strict_capability_match", sv)
				diags = append(diags, e)
			}
			obj.StrictCapabilityMatch = &v2
		}
	}
	if v1, ok := d.GetOk("unsuppress_map"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unsuppress_map", sv)
				diags = append(diags, e)
			}
			obj.UnsuppressMap = &v2
		}
	}
	if v1, ok := d.GetOk("unsuppress_map6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unsuppress_map6", sv)
				diags = append(diags, e)
			}
			obj.UnsuppressMap6 = &v2
		}
	}
	if v1, ok := d.GetOk("update_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_source", sv)
				diags = append(diags, e)
			}
			obj.UpdateSource = &v2
		}
	}
	if v1, ok := d.GetOk("weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Weight = &tmp
		}
	}
	return &obj, diags
}
