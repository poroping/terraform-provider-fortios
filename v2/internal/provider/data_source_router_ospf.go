// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure OSPF.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Description: "Configure OSPF.",

		ReadContext: dataSourceRouterOspfRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"abr_type": {
				Type:        schema.TypeString,
				Description: "Area border router type.",
				Computed:    true,
			},
			"area": {
				Type:        schema.TypeList,
				Description: "OSPF area configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication type.",
							Computed:    true,
						},
						"comments": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"default_cost": {
							Type:        schema.TypeInt,
							Description: "Summary default cost of stub or NSSA area.",
							Computed:    true,
						},
						"filter_list": {
							Type:        schema.TypeList,
							Description: "OSPF area filter-list configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": {
										Type:        schema.TypeString,
										Description: "Direction.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Filter list entry ID.",
										Computed:    true,
									},
									"list": {
										Type:        schema.TypeString,
										Description: "Access-list or prefix-list name.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Area entry IP address.",
							Computed:    true,
						},
						"nssa_default_information_originate": {
							Type:        schema.TypeString,
							Description: "Redistribute, advertise, or do not originate Type-7 default route into NSSA area.",
							Computed:    true,
						},
						"nssa_default_information_originate_metric": {
							Type:        schema.TypeInt,
							Description: "OSPF default metric.",
							Computed:    true,
						},
						"nssa_default_information_originate_metric_type": {
							Type:        schema.TypeString,
							Description: "OSPF metric type for default routes.",
							Computed:    true,
						},
						"nssa_redistribution": {
							Type:        schema.TypeString,
							Description: "Enable/disable redistribute into NSSA area.",
							Computed:    true,
						},
						"nssa_translator_role": {
							Type:        schema.TypeString,
							Description: "NSSA translator role type.",
							Computed:    true,
						},
						"range": {
							Type:        schema.TypeList,
							Description: "OSPF area range configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise": {
										Type:        schema.TypeString,
										Description: "Enable/disable advertise status.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Range entry ID.",
										Computed:    true,
									},
									"prefix": {
										Type:        schema.TypeString,
										Description: "Prefix.",
										Computed:    true,
									},
									"substitute": {
										Type:        schema.TypeString,
										Description: "Substitute prefix.",
										Computed:    true,
									},
									"substitute_status": {
										Type:        schema.TypeString,
										Description: "Enable/disable substitute status.",
										Computed:    true,
									},
								},
							},
						},
						"shortcut": {
							Type:        schema.TypeString,
							Description: "Enable/disable shortcut option.",
							Computed:    true,
						},
						"stub_type": {
							Type:        schema.TypeString,
							Description: "Stub summary setting.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Area type setting.",
							Computed:    true,
						},
						"virtual_link": {
							Type:        schema.TypeList,
							Description: "OSPF virtual link configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:        schema.TypeString,
										Description: "Authentication type.",
										Computed:    true,
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Description: "Authentication key.",
										Computed:    true,
										Sensitive:   true,
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Description: "Dead interval.",
										Computed:    true,
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Description: "Hello interval.",
										Computed:    true,
									},
									"keychain": {
										Type:        schema.TypeString,
										Description: "Message-digest key-chain name.",
										Computed:    true,
									},
									"md5_keychain": {
										Type:        schema.TypeString,
										Description: "Authentication MD5 key-chain name.",
										Computed:    true,
									},
									"md5_keys": {
										Type:        schema.TypeList,
										Description: "MD5 key.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Description: "Key ID (1 - 255).",
													Computed:    true,
												},
												"key_string": {
													Type:        schema.TypeString,
													Description: "Password for the key.",
													Computed:    true,
													Sensitive:   true,
												},
											},
										},
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual link entry name.",
										Computed:    true,
									},
									"peer": {
										Type:        schema.TypeString,
										Description: "Peer IP.",
										Computed:    true,
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Description: "Retransmit interval.",
										Computed:    true,
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Description: "Transmit delay.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"auto_cost_ref_bandwidth": {
				Type:        schema.TypeInt,
				Description: "Reference bandwidth in terms of megabits per second.",
				Computed:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Bidirectional Forwarding Detection (BFD).",
				Computed:    true,
			},
			"database_overflow": {
				Type:        schema.TypeString,
				Description: "Enable/disable database overflow.",
				Computed:    true,
			},
			"database_overflow_max_lsas": {
				Type:        schema.TypeInt,
				Description: "Database overflow maximum LSAs.",
				Computed:    true,
			},
			"database_overflow_time_to_recover": {
				Type:        schema.TypeInt,
				Description: "Database overflow time to recover (sec).",
				Computed:    true,
			},
			"default_information_metric": {
				Type:        schema.TypeInt,
				Description: "Default information metric.",
				Computed:    true,
			},
			"default_information_metric_type": {
				Type:        schema.TypeString,
				Description: "Default information metric type.",
				Computed:    true,
			},
			"default_information_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable generation of default route.",
				Computed:    true,
			},
			"default_information_route_map": {
				Type:        schema.TypeString,
				Description: "Default information route map.",
				Computed:    true,
			},
			"default_metric": {
				Type:        schema.TypeInt,
				Description: "Default metric of redistribute routes.",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Distance of the route.",
				Computed:    true,
			},
			"distance_external": {
				Type:        schema.TypeInt,
				Description: "Administrative external distance.",
				Computed:    true,
			},
			"distance_inter_area": {
				Type:        schema.TypeInt,
				Description: "Administrative inter-area distance.",
				Computed:    true,
			},
			"distance_intra_area": {
				Type:        schema.TypeInt,
				Description: "Administrative intra-area distance.",
				Computed:    true,
			},
			"distribute_list": {
				Type:        schema.TypeList,
				Description: "Distribute list configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list": {
							Type:        schema.TypeString,
							Description: "Access list name.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Distribute list entry ID.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocol type.",
							Computed:    true,
						},
					},
				},
			},
			"distribute_list_in": {
				Type:        schema.TypeString,
				Description: "Filter incoming routes.",
				Computed:    true,
			},
			"distribute_route_map_in": {
				Type:        schema.TypeString,
				Description: "Filter incoming external routes by route-map.",
				Computed:    true,
			},
			"log_neighbour_changes": {
				Type:        schema.TypeString,
				Description: "Log of OSPF neighbor changes.",
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "OSPF neighbor configuration are used when OSPF runs on non-broadcast media.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": {
							Type:        schema.TypeInt,
							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Neighbor entry ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "Interface IP address of the neighbor.",
							Computed:    true,
						},
						"poll_interval": {
							Type:        schema.TypeInt,
							Description: "Poll interval time in seconds.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Priority.",
							Computed:    true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "OSPF network configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area": {
							Type:        schema.TypeString,
							Description: "Attach the network to area.",
							Computed:    true,
						},
						"comments": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Network entry ID.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Prefix.",
							Computed:    true,
						},
					},
				},
			},
			"ospf_interface": {
				Type:        schema.TypeList,
				Description: "OSPF interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication type.",
							Computed:    true,
						},
						"authentication_key": {
							Type:        schema.TypeString,
							Description: "Authentication key.",
							Computed:    true,
							Sensitive:   true,
						},
						"bfd": {
							Type:        schema.TypeString,
							Description: "Bidirectional Forwarding Detection (BFD).",
							Computed:    true,
						},
						"comments": {
							Type:        schema.TypeString,
							Description: "Comment.",
							Computed:    true,
						},
						"cost": {
							Type:        schema.TypeInt,
							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Computed:    true,
						},
						"database_filter_out": {
							Type:        schema.TypeString,
							Description: "Enable/disable control of flooding out LSAs.",
							Computed:    true,
						},
						"dead_interval": {
							Type:        schema.TypeInt,
							Description: "Dead interval.",
							Computed:    true,
						},
						"hello_interval": {
							Type:        schema.TypeInt,
							Description: "Hello interval.",
							Computed:    true,
						},
						"hello_multiplier": {
							Type:        schema.TypeInt,
							Description: "Number of hello packets within dead interval.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Configuration interface name.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IP address.",
							Computed:    true,
						},
						"keychain": {
							Type:        schema.TypeString,
							Description: "Message-digest key-chain name.",
							Computed:    true,
						},
						"md5_keychain": {
							Type:        schema.TypeString,
							Description: "Authentication MD5 key-chain name.",
							Computed:    true,
						},
						"md5_keys": {
							Type:        schema.TypeList,
							Description: "MD5 key.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Key ID (1 - 255).",
										Computed:    true,
									},
									"key_string": {
										Type:        schema.TypeString,
										Description: "Password for the key.",
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"mtu": {
							Type:        schema.TypeInt,
							Description: "MTU for database description packets.",
							Computed:    true,
						},
						"mtu_ignore": {
							Type:        schema.TypeString,
							Description: "Enable/disable ignore MTU.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface entry name.",
							Computed:    true,
						},
						"network_type": {
							Type:        schema.TypeString,
							Description: "Network type.",
							Computed:    true,
						},
						"prefix_length": {
							Type:        schema.TypeInt,
							Description: "Prefix length.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Priority.",
							Computed:    true,
						},
						"resync_timeout": {
							Type:        schema.TypeInt,
							Description: "Graceful restart neighbor resynchronization timeout.",
							Computed:    true,
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Description: "Retransmit interval.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable status.",
							Computed:    true,
						},
						"transmit_delay": {
							Type:        schema.TypeInt,
							Description: "Transmit delay.",
							Computed:    true,
						},
					},
				},
			},
			"passive_interface": {
				Type:        schema.TypeList,
				Description: "Passive interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Passive interface name.",
							Computed:    true,
						},
					},
				},
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "Redistribute configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type:        schema.TypeInt,
							Description: "Redistribute metric setting.",
							Computed:    true,
						},
						"metric_type": {
							Type:        schema.TypeString,
							Description: "Metric type.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Redistribute name.",
							Computed:    true,
						},
						"routemap": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeInt,
							Description: "Tag value.",
							Computed:    true,
						},
					},
				},
			},
			"restart_mode": {
				Type:        schema.TypeString,
				Description: "OSPF restart mode (graceful or LLS).",
				Computed:    true,
			},
			"restart_on_topology_change": {
				Type:        schema.TypeString,
				Description: "Enable/disable continuing graceful restart upon topology change.",
				Computed:    true,
			},
			"restart_period": {
				Type:        schema.TypeInt,
				Description: "Graceful restart period.",
				Computed:    true,
			},
			"rfc1583_compatible": {
				Type:        schema.TypeString,
				Description: "Enable/disable RFC1583 compatibility.",
				Computed:    true,
			},
			"router_id": {
				Type:        schema.TypeString,
				Description: "Router ID.",
				Computed:    true,
			},
			"spf_timers": {
				Type:        schema.TypeString,
				Description: "SPF calculation frequency.",
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IP address summary configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertise status.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Summary address entry ID.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Prefix.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeInt,
							Description: "Tag value.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterOspf"

	o, err := c.Cmdb.ReadRouterOspf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf dataSource: %v", err)
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

	diags := refreshObjectRouterOspf(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
