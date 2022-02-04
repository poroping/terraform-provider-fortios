// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure OSPF.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Description: "Configure OSPF.",

		CreateContext: resourceRouterOspfCreate,
		ReadContext:   resourceRouterOspfRead,
		UpdateContext: resourceRouterOspfUpdate,
		DeleteContext: resourceRouterOspfDelete,

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
			"abr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cisco", "ibm", "shortcut", "standard"}, false),

				Description: "Area border router type.",
				Optional:    true,
				Computed:    true,
			},
			"area": {
				Type:        schema.TypeList,
				Description: "OSPF area configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "message-digest"}, false),

							Description: "Authentication type.",
							Optional:    true,
							Computed:    true,
						},
						"comments": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comment.",
							Optional:    true,
							Computed:    true,
						},
						"default_cost": {
							Type: schema.TypeInt,

							Description: "Summary default cost of stub or NSSA area.",
							Optional:    true,
							Computed:    true,
						},
						"filter_list": {
							Type:        schema.TypeList,
							Description: "OSPF area filter-list configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"in", "out"}, false),

										Description: "Direction.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Filter list entry ID.",
										Optional:    true,
										Computed:    true,
									},
									"list": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Access-list or prefix-list name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Area entry IP address.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_default_information_originate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "always", "disable"}, false),

							Description: "Redistribute, advertise, or do not originate Type-7 default route into NSSA area.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_default_information_originate_metric": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),

							Description: "OSPF default metric.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_default_information_originate_metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

							Description: "OSPF metric type for default routes.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_redistribution": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable redistribute into NSSA area.",
							Optional:    true,
							Computed:    true,
						},
						"nssa_translator_role": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"candidate", "never", "always"}, false),

							Description: "NSSA translator role type.",
							Optional:    true,
							Computed:    true,
						},
						"range": {
							Type:        schema.TypeList,
							Description: "OSPF area range configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable advertise status.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Range entry ID.",
										Optional:    true,
										Computed:    true,
									},
									"prefix": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

										Description: "Prefix.",
										Optional:    true,
										Computed:    true,
									},
									"substitute": {
										Type:         schema.TypeString,
										ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

										Description: "Substitute prefix.",
										Optional:    true,
										Computed:    true,
									},
									"substitute_status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable substitute status.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"shortcut": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "default"}, false),

							Description: "Enable/disable shortcut option.",
							Optional:    true,
							Computed:    true,
						},
						"stub_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no-summary", "summary"}, false),

							Description: "Stub summary setting.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"regular", "nssa", "stub"}, false),

							Description: "Area type setting.",
							Optional:    true,
							Computed:    true,
						},
						"virtual_link": {
							Type:        schema.TypeList,
							Description: "OSPF virtual link configuration.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"none", "text", "message-digest"}, false),

										Description: "Authentication type.",
										Optional:    true,
										Computed:    true,
									},
									"authentication_key": {
										Type: schema.TypeString,

										Description: "Authentication key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
									"dead_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Dead interval.",
										Optional:    true,
										Computed:    true,
									},
									"hello_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Hello interval.",
										Optional:    true,
										Computed:    true,
									},
									"keychain": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Message-digest key-chain name.",
										Optional:    true,
										Computed:    true,
									},
									"md5_keychain": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Authentication MD5 key-chain name.",
										Optional:    true,
										Computed:    true,
									},
									"md5_keys": {
										Type:        schema.TypeList,
										Description: "MD5 key.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 255),

													Description: "Key ID (1 - 255).",
													Optional:    true,
													Computed:    true,
												},
												"key_string": {
													Type: schema.TypeString,

													Description: "Password for the key.",
													Optional:    true,
													Computed:    true,
													Sensitive:   true,
												},
											},
										},
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Virtual link entry name.",
										Optional:    true,
										Computed:    true,
									},
									"peer": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "Peer IP.",
										Optional:    true,
										Computed:    true,
									},
									"retransmit_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Retransmit interval.",
										Optional:    true,
										Computed:    true,
									},
									"transmit_delay": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Transmit delay.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"auto_cost_ref_bandwidth": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000000),

				Description: "Reference bandwidth in terms of megabits per second.",
				Optional:    true,
				Computed:    true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Bidirectional Forwarding Detection (BFD).",
				Optional:    true,
				Computed:    true,
			},
			"database_overflow": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable database overflow.",
				Optional:    true,
				Computed:    true,
			},
			"database_overflow_max_lsas": {
				Type: schema.TypeInt,

				Description: "Database overflow maximum LSAs.",
				Optional:    true,
				Computed:    true,
			},
			"database_overflow_time_to_recover": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Database overflow time to recover (sec).",
				Optional:    true,
				Computed:    true,
			},
			"default_information_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),

				Description: "Default information metric.",
				Optional:    true,
				Computed:    true,
			},
			"default_information_metric_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

				Description: "Default information metric type.",
				Optional:    true,
				Computed:    true,
			},
			"default_information_originate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "always", "disable"}, false),

				Description: "Enable/disable generation of default route.",
				Optional:    true,
				Computed:    true,
			},
			"default_information_route_map": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default information route map.",
				Optional:    true,
				Computed:    true,
			},
			"default_metric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),

				Description: "Default metric of redistribute routes.",
				Optional:    true,
				Computed:    true,
			},
			"distance": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Distance of the route.",
				Optional:    true,
				Computed:    true,
			},
			"distance_external": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative external distance.",
				Optional:    true,
				Computed:    true,
			},
			"distance_inter_area": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative inter-area distance.",
				Optional:    true,
				Computed:    true,
			},
			"distance_intra_area": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Administrative intra-area distance.",
				Optional:    true,
				Computed:    true,
			},
			"distribute_list": {
				Type:        schema.TypeList,
				Description: "Distribute list configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Access list name.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Distribute list entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"connected", "static", "rip"}, false),

							Description: "Protocol type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"distribute_list_in": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Filter incoming routes.",
				Optional:    true,
				Computed:    true,
			},
			"distribute_route_map_in": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Filter incoming external routes by route-map.",
				Optional:    true,
				Computed:    true,
			},
			"log_neighbour_changes": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Log of OSPF neighbor changes.",
				Optional:    true,
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "OSPF neighbor configuration are used when OSPF runs on non-broadcast media.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Neighbor entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Interface IP address of the neighbor.",
							Optional:    true,
							Computed:    true,
						},
						"poll_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Poll interval time in seconds.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Priority.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "OSPF network configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Attach the network to area.",
							Optional:    true,
							Computed:    true,
						},
						"comments": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comment.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Network entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Prefix.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ospf_interface": {
				Type:        schema.TypeList,
				Description: "OSPF interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "text", "message-digest"}, false),

							Description: "Authentication type.",
							Optional:    true,
							Computed:    true,
						},
						"authentication_key": {
							Type: schema.TypeString,

							Description: "Authentication key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

							Description: "Bidirectional Forwarding Detection (BFD).",
							Optional:    true,
							Computed:    true,
						},
						"comments": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Comment.",
							Optional:    true,
							Computed:    true,
						},
						"cost": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Optional:    true,
							Computed:    true,
						},
						"database_filter_out": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable control of flooding out LSAs.",
							Optional:    true,
							Computed:    true,
						},
						"dead_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Dead interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Hello interval.",
							Optional:    true,
							Computed:    true,
						},
						"hello_multiplier": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 10),

							Description: "Number of hello packets within dead interval.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Configuration interface name.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address.",
							Optional:    true,
							Computed:    true,
						},
						"keychain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Message-digest key-chain name.",
							Optional:    true,
							Computed:    true,
						},
						"md5_keychain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Authentication MD5 key-chain name.",
							Optional:    true,
							Computed:    true,
						},
						"md5_keys": {
							Type:        schema.TypeList,
							Description: "MD5 key.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Key ID (1 - 255).",
										Optional:    true,
										Computed:    true,
									},
									"key_string": {
										Type: schema.TypeString,

										Description: "Password for the key.",
										Optional:    true,
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"mtu": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(576, 65535),

							Description: "MTU for database description packets.",
							Optional:    true,
							Computed:    true,
						},
						"mtu_ignore": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable ignore MTU.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Interface entry name.",
							Optional:    true,
							Computed:    true,
						},
						"network_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"broadcast", "non-broadcast", "point-to-point", "point-to-multipoint", "point-to-multipoint-non-broadcast"}, false),

							Description: "Network type.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_length": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),

							Description: "Prefix length.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Priority.",
							Optional:    true,
							Computed:    true,
						},
						"resync_timeout": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),

							Description: "Graceful restart neighbor resynchronization timeout.",
							Optional:    true,
							Computed:    true,
						},
						"retransmit_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Retransmit interval.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable status.",
							Optional:    true,
							Computed:    true,
						},
						"transmit_delay": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Transmit delay.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"passive_interface": {
				Type:        schema.TypeList,
				Description: "Passive interface configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Passive interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "Redistribute configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),

							Description: "Redistribute metric setting.",
							Optional:    true,
							Computed:    true,
						},
						"metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"1", "2"}, false),

							Description: "Metric type.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Redistribute name.",
							Optional:    true,
							Computed:    true,
						},
						"routemap": {
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
						"tag": {
							Type: schema.TypeInt,

							Description: "Tag value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"restart_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "lls", "graceful-restart"}, false),

				Description: "OSPF restart mode (graceful or LLS).",
				Optional:    true,
				Computed:    true,
			},
			"restart_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Graceful restart period.",
				Optional:    true,
				Computed:    true,
			},
			"rfc1583_compatible": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RFC1583 compatibility.",
				Optional:    true,
				Computed:    true,
			},
			"router_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Router ID.",
				Optional:    true,
				Computed:    true,
			},
			"spf_timers": {
				Type: schema.TypeString,

				Description: "SPF calculation frequency.",
				Optional:    true,
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IP address summary configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable advertise status.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Summary address entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Prefix.",
							Optional:    true,
							Computed:    true,
						},
						"tag": {
							Type: schema.TypeInt,

							Description: "Tag value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspf(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf")
	}

	return resourceRouterOspfRead(ctx, d, meta)
}

func resourceRouterOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspf(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspf resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf")
	}

	return resourceRouterOspfRead(ctx, d, meta)
}

func resourceRouterOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterOspf(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterOspf(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf resource: %v", err)
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

	diags := refreshObjectRouterOspf(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterOspfArea(v *[]models.RouterOspfArea, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.Comments; tmp != nil {
				v["comments"] = *tmp
			}

			if tmp := cfg.DefaultCost; tmp != nil {
				v["default_cost"] = *tmp
			}

			if tmp := cfg.FilterList; tmp != nil {
				v["filter_list"] = flattenRouterOspfAreaFilterList(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.NssaDefaultInformationOriginate; tmp != nil {
				v["nssa_default_information_originate"] = *tmp
			}

			if tmp := cfg.NssaDefaultInformationOriginateMetric; tmp != nil {
				v["nssa_default_information_originate_metric"] = *tmp
			}

			if tmp := cfg.NssaDefaultInformationOriginateMetricType; tmp != nil {
				v["nssa_default_information_originate_metric_type"] = *tmp
			}

			if tmp := cfg.NssaRedistribution; tmp != nil {
				v["nssa_redistribution"] = *tmp
			}

			if tmp := cfg.NssaTranslatorRole; tmp != nil {
				v["nssa_translator_role"] = *tmp
			}

			if tmp := cfg.Range; tmp != nil {
				v["range"] = flattenRouterOspfAreaRange(tmp, sort)
			}

			if tmp := cfg.Shortcut; tmp != nil {
				v["shortcut"] = *tmp
			}

			if tmp := cfg.StubType; tmp != nil {
				v["stub_type"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.VirtualLink; tmp != nil {
				v["virtual_link"] = flattenRouterOspfAreaVirtualLink(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfAreaFilterList(v *[]models.RouterOspfAreaFilterList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.List; tmp != nil {
				v["list"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfAreaRange(v *[]models.RouterOspfAreaRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Advertise; tmp != nil {
				v["advertise"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			if tmp := cfg.Substitute; tmp != nil {
				v["substitute"] = *tmp
			}

			if tmp := cfg.SubstituteStatus; tmp != nil {
				v["substitute_status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfAreaVirtualLink(v *[]models.RouterOspfAreaVirtualLink, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.AuthenticationKey; tmp != nil {
				v["authentication_key"] = *tmp
			}

			if tmp := cfg.DeadInterval; tmp != nil {
				v["dead_interval"] = *tmp
			}

			if tmp := cfg.HelloInterval; tmp != nil {
				v["hello_interval"] = *tmp
			}

			if tmp := cfg.Keychain; tmp != nil {
				v["keychain"] = *tmp
			}

			if tmp := cfg.Md5Keychain; tmp != nil {
				v["md5_keychain"] = *tmp
			}

			if tmp := cfg.Md5Keys; tmp != nil {
				v["md5_keys"] = flattenRouterOspfAreaVirtualLinkMd5Keys(tmp, sort)
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Peer; tmp != nil {
				v["peer"] = *tmp
			}

			if tmp := cfg.RetransmitInterval; tmp != nil {
				v["retransmit_interval"] = *tmp
			}

			if tmp := cfg.TransmitDelay; tmp != nil {
				v["transmit_delay"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterOspfAreaVirtualLinkMd5Keys(v *[]models.RouterOspfAreaVirtualLinkMd5Keys, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.KeyString; tmp != nil {
				v["key_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfDistributeList(v *[]models.RouterOspfDistributeList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AccessList; tmp != nil {
				v["access_list"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfNeighbor(v *[]models.RouterOspfNeighbor, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Cost; tmp != nil {
				v["cost"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.PollInterval; tmp != nil {
				v["poll_interval"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfNetwork(v *[]models.RouterOspfNetwork, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Area; tmp != nil {
				v["area"] = *tmp
			}

			if tmp := cfg.Comments; tmp != nil {
				v["comments"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfOspfInterface(v *[]models.RouterOspfOspfInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Authentication; tmp != nil {
				v["authentication"] = *tmp
			}

			if tmp := cfg.AuthenticationKey; tmp != nil {
				v["authentication_key"] = *tmp
			}

			if tmp := cfg.Bfd; tmp != nil {
				v["bfd"] = *tmp
			}

			if tmp := cfg.Comments; tmp != nil {
				v["comments"] = *tmp
			}

			if tmp := cfg.Cost; tmp != nil {
				v["cost"] = *tmp
			}

			if tmp := cfg.DatabaseFilterOut; tmp != nil {
				v["database_filter_out"] = *tmp
			}

			if tmp := cfg.DeadInterval; tmp != nil {
				v["dead_interval"] = *tmp
			}

			if tmp := cfg.HelloInterval; tmp != nil {
				v["hello_interval"] = *tmp
			}

			if tmp := cfg.HelloMultiplier; tmp != nil {
				v["hello_multiplier"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Keychain; tmp != nil {
				v["keychain"] = *tmp
			}

			if tmp := cfg.Md5Keychain; tmp != nil {
				v["md5_keychain"] = *tmp
			}

			if tmp := cfg.Md5Keys; tmp != nil {
				v["md5_keys"] = flattenRouterOspfOspfInterfaceMd5Keys(tmp, sort)
			}

			if tmp := cfg.Mtu; tmp != nil {
				v["mtu"] = *tmp
			}

			if tmp := cfg.MtuIgnore; tmp != nil {
				v["mtu_ignore"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.NetworkType; tmp != nil {
				v["network_type"] = *tmp
			}

			if tmp := cfg.PrefixLength; tmp != nil {
				v["prefix_length"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.ResyncTimeout; tmp != nil {
				v["resync_timeout"] = *tmp
			}

			if tmp := cfg.RetransmitInterval; tmp != nil {
				v["retransmit_interval"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TransmitDelay; tmp != nil {
				v["transmit_delay"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterOspfOspfInterfaceMd5Keys(v *[]models.RouterOspfOspfInterfaceMd5Keys, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.KeyString; tmp != nil {
				v["key_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterOspfPassiveInterface(v *[]models.RouterOspfPassiveInterface, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenRouterOspfRedistribute(v *[]models.RouterOspfRedistribute, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Metric; tmp != nil {
				v["metric"] = *tmp
			}

			if tmp := cfg.MetricType; tmp != nil {
				v["metric_type"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Routemap; tmp != nil {
				v["routemap"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterOspfSummaryAddress(v *[]models.RouterOspfSummaryAddress, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Advertise; tmp != nil {
				v["advertise"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			if tmp := cfg.Tag; tmp != nil {
				v["tag"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterOspf(d *schema.ResourceData, o *models.RouterOspf, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AbrType != nil {
		v := *o.AbrType

		if err = d.Set("abr_type", v); err != nil {
			return diag.Errorf("error reading abr_type: %v", err)
		}
	}

	if o.Area != nil {
		if err = d.Set("area", flattenRouterOspfArea(o.Area, sort)); err != nil {
			return diag.Errorf("error reading area: %v", err)
		}
	}

	if o.AutoCostRefBandwidth != nil {
		v := *o.AutoCostRefBandwidth

		if err = d.Set("auto_cost_ref_bandwidth", v); err != nil {
			return diag.Errorf("error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.DatabaseOverflow != nil {
		v := *o.DatabaseOverflow

		if err = d.Set("database_overflow", v); err != nil {
			return diag.Errorf("error reading database_overflow: %v", err)
		}
	}

	if o.DatabaseOverflowMaxLsas != nil {
		v := *o.DatabaseOverflowMaxLsas

		if err = d.Set("database_overflow_max_lsas", v); err != nil {
			return diag.Errorf("error reading database_overflow_max_lsas: %v", err)
		}
	}

	if o.DatabaseOverflowTimeToRecover != nil {
		v := *o.DatabaseOverflowTimeToRecover

		if err = d.Set("database_overflow_time_to_recover", v); err != nil {
			return diag.Errorf("error reading database_overflow_time_to_recover: %v", err)
		}
	}

	if o.DefaultInformationMetric != nil {
		v := *o.DefaultInformationMetric

		if err = d.Set("default_information_metric", v); err != nil {
			return diag.Errorf("error reading default_information_metric: %v", err)
		}
	}

	if o.DefaultInformationMetricType != nil {
		v := *o.DefaultInformationMetricType

		if err = d.Set("default_information_metric_type", v); err != nil {
			return diag.Errorf("error reading default_information_metric_type: %v", err)
		}
	}

	if o.DefaultInformationOriginate != nil {
		v := *o.DefaultInformationOriginate

		if err = d.Set("default_information_originate", v); err != nil {
			return diag.Errorf("error reading default_information_originate: %v", err)
		}
	}

	if o.DefaultInformationRouteMap != nil {
		v := *o.DefaultInformationRouteMap

		if err = d.Set("default_information_route_map", v); err != nil {
			return diag.Errorf("error reading default_information_route_map: %v", err)
		}
	}

	if o.DefaultMetric != nil {
		v := *o.DefaultMetric

		if err = d.Set("default_metric", v); err != nil {
			return diag.Errorf("error reading default_metric: %v", err)
		}
	}

	if o.Distance != nil {
		v := *o.Distance

		if err = d.Set("distance", v); err != nil {
			return diag.Errorf("error reading distance: %v", err)
		}
	}

	if o.DistanceExternal != nil {
		v := *o.DistanceExternal

		if err = d.Set("distance_external", v); err != nil {
			return diag.Errorf("error reading distance_external: %v", err)
		}
	}

	if o.DistanceInterArea != nil {
		v := *o.DistanceInterArea

		if err = d.Set("distance_inter_area", v); err != nil {
			return diag.Errorf("error reading distance_inter_area: %v", err)
		}
	}

	if o.DistanceIntraArea != nil {
		v := *o.DistanceIntraArea

		if err = d.Set("distance_intra_area", v); err != nil {
			return diag.Errorf("error reading distance_intra_area: %v", err)
		}
	}

	if o.DistributeList != nil {
		if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o.DistributeList, sort)); err != nil {
			return diag.Errorf("error reading distribute_list: %v", err)
		}
	}

	if o.DistributeListIn != nil {
		v := *o.DistributeListIn

		if err = d.Set("distribute_list_in", v); err != nil {
			return diag.Errorf("error reading distribute_list_in: %v", err)
		}
	}

	if o.DistributeRouteMapIn != nil {
		v := *o.DistributeRouteMapIn

		if err = d.Set("distribute_route_map_in", v); err != nil {
			return diag.Errorf("error reading distribute_route_map_in: %v", err)
		}
	}

	if o.LogNeighbourChanges != nil {
		v := *o.LogNeighbourChanges

		if err = d.Set("log_neighbour_changes", v); err != nil {
			return diag.Errorf("error reading log_neighbour_changes: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenRouterOspfNeighbor(o.Neighbor, sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	if o.Network != nil {
		if err = d.Set("network", flattenRouterOspfNetwork(o.Network, sort)); err != nil {
			return diag.Errorf("error reading network: %v", err)
		}
	}

	if o.OspfInterface != nil {
		if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o.OspfInterface, sort)); err != nil {
			return diag.Errorf("error reading ospf_interface: %v", err)
		}
	}

	if o.PassiveInterface != nil {
		if err = d.Set("passive_interface", flattenRouterOspfPassiveInterface(o.PassiveInterface, sort)); err != nil {
			return diag.Errorf("error reading passive_interface: %v", err)
		}
	}

	if o.Redistribute != nil {
		if err = d.Set("redistribute", flattenRouterOspfRedistribute(o.Redistribute, sort)); err != nil {
			return diag.Errorf("error reading redistribute: %v", err)
		}
	}

	if o.RestartMode != nil {
		v := *o.RestartMode

		if err = d.Set("restart_mode", v); err != nil {
			return diag.Errorf("error reading restart_mode: %v", err)
		}
	}

	if o.RestartPeriod != nil {
		v := *o.RestartPeriod

		if err = d.Set("restart_period", v); err != nil {
			return diag.Errorf("error reading restart_period: %v", err)
		}
	}

	if o.Rfc1583Compatible != nil {
		v := *o.Rfc1583Compatible

		if err = d.Set("rfc1583_compatible", v); err != nil {
			return diag.Errorf("error reading rfc1583_compatible: %v", err)
		}
	}

	if o.RouterId != nil {
		v := *o.RouterId

		if err = d.Set("router_id", v); err != nil {
			return diag.Errorf("error reading router_id: %v", err)
		}
	}

	if o.SpfTimers != nil {
		v := *o.SpfTimers

		if err = d.Set("spf_timers", v); err != nil {
			return diag.Errorf("error reading spf_timers: %v", err)
		}
	}

	if o.SummaryAddress != nil {
		if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o.SummaryAddress, sort)); err != nil {
			return diag.Errorf("error reading summary_address: %v", err)
		}
	}

	return nil
}

func expandRouterOspfArea(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfArea, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfArea

	for i := range l {
		tmp := models.RouterOspfArea{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comments", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comments = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.default_cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DefaultCost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspfAreaFilterList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspfAreaFilterList
			// 	}
			tmp.FilterList = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_default_information_originate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaDefaultInformationOriginate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_default_information_originate_metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.NssaDefaultInformationOriginateMetric = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_default_information_originate_metric_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaDefaultInformationOriginateMetricType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_redistribution", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaRedistribution = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nssa_translator_role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NssaTranslatorRole = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspfAreaRange(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspfAreaRange
			// 	}
			tmp.Range = v2

		}

		pre_append = fmt.Sprintf("%s.%d.shortcut", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Shortcut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stub_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StubType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virtual_link", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspfAreaVirtualLink(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspfAreaVirtualLink
			// 	}
			tmp.VirtualLink = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfAreaFilterList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfAreaFilterList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfAreaFilterList

	for i := range l {
		tmp := models.RouterOspfAreaFilterList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.List = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfAreaRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfAreaRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfAreaRange

	for i := range l {
		tmp := models.RouterOspfAreaRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.advertise", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Advertise = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.substitute", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Substitute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.substitute_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SubstituteStatus = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfAreaVirtualLink(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfAreaVirtualLink, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfAreaVirtualLink

	for i := range l {
		tmp := models.RouterOspfAreaVirtualLink{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.authentication_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthenticationKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dead_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DeadInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.keychain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Keychain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.md5_keychain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Md5Keychain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.md5_keys", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspfAreaVirtualLinkMd5Keys(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspfAreaVirtualLinkMd5Keys
			// 	}
			tmp.Md5Keys = v2

		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.peer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Peer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retransmit_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RetransmitInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_delay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TransmitDelay = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfAreaVirtualLinkMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfAreaVirtualLinkMd5Keys, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfAreaVirtualLinkMd5Keys

	for i := range l {
		tmp := models.RouterOspfAreaVirtualLinkMd5Keys{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.key_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeyString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfDistributeList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfDistributeList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfDistributeList

	for i := range l {
		tmp := models.RouterOspfDistributeList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessList = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Protocol = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfNeighbor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfNeighbor

	for i := range l {
		tmp := models.RouterOspfNeighbor{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Cost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.poll_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PollInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfNetwork, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfNetwork

	for i := range l {
		tmp := models.RouterOspfNetwork{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.area", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Area = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comments", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comments = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
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

func expandRouterOspfOspfInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfOspfInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfOspfInterface

	for i := range l {
		tmp := models.RouterOspfOspfInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.authentication", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Authentication = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.authentication_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthenticationKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bfd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bfd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comments", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comments = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cost", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Cost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.database_filter_out", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DatabaseFilterOut = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dead_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DeadInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_multiplier", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HelloMultiplier = &v2
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

		pre_append = fmt.Sprintf("%s.%d.keychain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Keychain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.md5_keychain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Md5Keychain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.md5_keys", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterOspfOspfInterfaceMd5Keys(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterOspfOspfInterfaceMd5Keys
			// 	}
			tmp.Md5Keys = v2

		}

		pre_append = fmt.Sprintf("%s.%d.mtu", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Mtu = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mtu_ignore", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MtuIgnore = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.network_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NetworkType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PrefixLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.resync_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.ResyncTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retransmit_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.RetransmitInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.transmit_delay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TransmitDelay = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfOspfInterfaceMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfOspfInterfaceMd5Keys, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfOspfInterfaceMd5Keys

	for i := range l {
		tmp := models.RouterOspfOspfInterfaceMd5Keys{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.key_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.KeyString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfPassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfPassiveInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfPassiveInterface

	for i := range l {
		tmp := models.RouterOspfPassiveInterface{}
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

func expandRouterOspfRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfRedistribute, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfRedistribute

	for i := range l {
		tmp := models.RouterOspfRedistribute{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Metric = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metric_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MetricType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.routemap", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Routemap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Tag = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterOspfSummaryAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterOspfSummaryAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterOspfSummaryAddress

	for i := range l {
		tmp := models.RouterOspfSummaryAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.advertise", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Advertise = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Tag = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterOspf(d *schema.ResourceData, sv string) (*models.RouterOspf, diag.Diagnostics) {
	obj := models.RouterOspf{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("abr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("abr_type", sv)
				diags = append(diags, e)
			}
			obj.AbrType = &v2
		}
	}
	if v, ok := d.GetOk("area"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("area", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfArea(d, v, "area", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Area = t
		}
	} else if d.HasChange("area") {
		old, new := d.GetChange("area")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Area = &[]models.RouterOspfArea{}
		}
	}
	if v1, ok := d.GetOk("auto_cost_ref_bandwidth"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_cost_ref_bandwidth", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoCostRefBandwidth = &tmp
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
	if v1, ok := d.GetOk("database_overflow"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("database_overflow", sv)
				diags = append(diags, e)
			}
			obj.DatabaseOverflow = &v2
		}
	}
	if v1, ok := d.GetOk("database_overflow_max_lsas"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("database_overflow_max_lsas", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DatabaseOverflowMaxLsas = &tmp
		}
	}
	if v1, ok := d.GetOk("database_overflow_time_to_recover"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("database_overflow_time_to_recover", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DatabaseOverflowTimeToRecover = &tmp
		}
	}
	if v1, ok := d.GetOk("default_information_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultInformationMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("default_information_metric_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_metric_type", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationMetricType = &v2
		}
	}
	if v1, ok := d.GetOk("default_information_originate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_originate", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationOriginate = &v2
		}
	}
	if v1, ok := d.GetOk("default_information_route_map"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_information_route_map", sv)
				diags = append(diags, e)
			}
			obj.DefaultInformationRouteMap = &v2
		}
	}
	if v1, ok := d.GetOk("default_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DefaultMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("distance"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Distance = &tmp
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
	if v1, ok := d.GetOk("distance_inter_area"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance_inter_area", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DistanceInterArea = &tmp
		}
	}
	if v1, ok := d.GetOk("distance_intra_area"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distance_intra_area", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DistanceIntraArea = &tmp
		}
	}
	if v, ok := d.GetOk("distribute_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("distribute_list", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfDistributeList(d, v, "distribute_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DistributeList = t
		}
	} else if d.HasChange("distribute_list") {
		old, new := d.GetChange("distribute_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DistributeList = &[]models.RouterOspfDistributeList{}
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
	if v1, ok := d.GetOk("distribute_route_map_in"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("distribute_route_map_in", sv)
				diags = append(diags, e)
			}
			obj.DistributeRouteMapIn = &v2
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
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.RouterOspfNeighbor{}
		}
	}
	if v, ok := d.GetOk("network"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("network", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Network = t
		}
	} else if d.HasChange("network") {
		old, new := d.GetChange("network")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Network = &[]models.RouterOspfNetwork{}
		}
	}
	if v, ok := d.GetOk("ospf_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ospf_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfOspfInterface(d, v, "ospf_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OspfInterface = t
		}
	} else if d.HasChange("ospf_interface") {
		old, new := d.GetChange("ospf_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OspfInterface = &[]models.RouterOspfOspfInterface{}
		}
	}
	if v, ok := d.GetOk("passive_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("passive_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfPassiveInterface(d, v, "passive_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PassiveInterface = t
		}
	} else if d.HasChange("passive_interface") {
		old, new := d.GetChange("passive_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PassiveInterface = &[]models.RouterOspfPassiveInterface{}
		}
	}
	if v, ok := d.GetOk("redistribute"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("redistribute", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfRedistribute(d, v, "redistribute", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Redistribute = t
		}
	} else if d.HasChange("redistribute") {
		old, new := d.GetChange("redistribute")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Redistribute = &[]models.RouterOspfRedistribute{}
		}
	}
	if v1, ok := d.GetOk("restart_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("restart_mode", sv)
				diags = append(diags, e)
			}
			obj.RestartMode = &v2
		}
	}
	if v1, ok := d.GetOk("restart_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("restart_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RestartPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("rfc1583_compatible"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rfc1583_compatible", sv)
				diags = append(diags, e)
			}
			obj.Rfc1583Compatible = &v2
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
	if v1, ok := d.GetOk("spf_timers"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spf_timers", sv)
				diags = append(diags, e)
			}
			obj.SpfTimers = &v2
		}
	}
	if v, ok := d.GetOk("summary_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("summary_address", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfSummaryAddress(d, v, "summary_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SummaryAddress = t
		}
	} else if d.HasChange("summary_address") {
		old, new := d.GetChange("summary_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SummaryAddress = &[]models.RouterOspfSummaryAddress{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterOspf(d *schema.ResourceData, sv string) (*models.RouterOspf, diag.Diagnostics) {
	obj := models.RouterOspf{}
	diags := diag.Diagnostics{}

	obj.Area = &[]models.RouterOspfArea{}
	obj.DistributeList = &[]models.RouterOspfDistributeList{}
	obj.Neighbor = &[]models.RouterOspfNeighbor{}
	obj.Network = &[]models.RouterOspfNetwork{}
	obj.OspfInterface = &[]models.RouterOspfOspfInterface{}
	obj.PassiveInterface = &[]models.RouterOspfPassiveInterface{}
	obj.Redistribute = &[]models.RouterOspfRedistribute{}
	obj.SummaryAddress = &[]models.RouterOspfSummaryAddress{}

	return &obj, diags
}
