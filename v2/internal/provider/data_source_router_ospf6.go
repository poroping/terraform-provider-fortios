// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 OSPF.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterOspf6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 OSPF.",

		ReadContext: dataSourceRouterOspf6Read,

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
				Description: "OSPF6 area configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"default_cost": {
							Type:        schema.TypeInt,
							Description: "Summary default cost of stub or NSSA area.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Area entry IP address.",
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:        schema.TypeString,
							Description: "Authentication algorithm.",
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:        schema.TypeString,
							Description: "Encryption algorithm.",
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type:        schema.TypeString,
										Description: "Authentication key.",
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type:        schema.TypeString,
										Description: "Encryption key.",
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type:        schema.TypeInt,
										Description: "Security Parameters Index.",
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:        schema.TypeInt,
							Description: "Key roll-over interval.",
							Computed:    true,
						},
						"nssa_default_information_originate": {
							Type:        schema.TypeString,
							Description: "Enable/disable originate type 7 default into NSSA area.",
							Computed:    true,
						},
						"nssa_default_information_originate_metric": {
							Type:        schema.TypeInt,
							Description: "OSPFv3 default metric.",
							Computed:    true,
						},
						"nssa_default_information_originate_metric_type": {
							Type:        schema.TypeString,
							Description: "OSPFv3 metric type for default routes.",
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
							Description: "OSPF6 area range configuration.",
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
									"prefix6": {
										Type:        schema.TypeString,
										Description: "IPv6 prefix.",
										Computed:    true,
									},
								},
							},
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
							Description: "OSPF6 virtual link configuration.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:        schema.TypeString,
										Description: "Authentication mode.",
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
									"ipsec_auth_alg": {
										Type:        schema.TypeString,
										Description: "Authentication algorithm.",
										Computed:    true,
									},
									"ipsec_enc_alg": {
										Type:        schema.TypeString,
										Description: "Encryption algorithm.",
										Computed:    true,
									},
									"ipsec_keys": {
										Type:        schema.TypeList,
										Description: "IPsec authentication and encryption keys.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auth_key": {
													Type:        schema.TypeString,
													Description: "Authentication key.",
													Computed:    true,
													Sensitive:   true,
												},
												"enc_key": {
													Type:        schema.TypeString,
													Description: "Encryption key.",
													Computed:    true,
													Sensitive:   true,
												},
												"spi": {
													Type:        schema.TypeInt,
													Description: "Security Parameters Index.",
													Computed:    true,
												},
											},
										},
									},
									"key_rollover_interval": {
										Type:        schema.TypeInt,
										Description: "Key roll-over interval.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Virtual link entry name.",
										Computed:    true,
									},
									"peer": {
										Type:        schema.TypeString,
										Description: "A.B.C.D, peer router ID.",
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
				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
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
			"log_neighbour_changes": {
				Type:        schema.TypeString,
				Description: "Log OSPFv3 neighbor changes.",
				Computed:    true,
			},
			"ospf6_interface": {
				Type:        schema.TypeList,
				Description: "OSPF6 interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_id": {
							Type:        schema.TypeString,
							Description: "A.B.C.D, in IPv4 address format.",
							Computed:    true,
						},
						"authentication": {
							Type:        schema.TypeString,
							Description: "Authentication mode.",
							Computed:    true,
						},
						"bfd": {
							Type:        schema.TypeString,
							Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
							Computed:    true,
						},
						"cost": {
							Type:        schema.TypeInt,
							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
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
						"interface": {
							Type:        schema.TypeString,
							Description: "Configuration interface name.",
							Computed:    true,
						},
						"ipsec_auth_alg": {
							Type:        schema.TypeString,
							Description: "Authentication algorithm.",
							Computed:    true,
						},
						"ipsec_enc_alg": {
							Type:        schema.TypeString,
							Description: "Encryption algorithm.",
							Computed:    true,
						},
						"ipsec_keys": {
							Type:        schema.TypeList,
							Description: "IPsec authentication and encryption keys.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_key": {
										Type:        schema.TypeString,
										Description: "Authentication key.",
										Computed:    true,
										Sensitive:   true,
									},
									"enc_key": {
										Type:        schema.TypeString,
										Description: "Encryption key.",
										Computed:    true,
										Sensitive:   true,
									},
									"spi": {
										Type:        schema.TypeInt,
										Description: "Security Parameters Index.",
										Computed:    true,
									},
								},
							},
						},
						"key_rollover_interval": {
							Type:        schema.TypeInt,
							Description: "Key roll-over interval.",
							Computed:    true,
						},
						"mtu": {
							Type:        schema.TypeInt,
							Description: "MTU for OSPFv3 packets.",
							Computed:    true,
						},
						"mtu_ignore": {
							Type:        schema.TypeString,
							Description: "Enable/disable ignoring MTU field in DBD packets.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface entry name.",
							Computed:    true,
						},
						"neighbor": {
							Type:        schema.TypeList,
							Description: "OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type:        schema.TypeInt,
										Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
										Computed:    true,
									},
									"ip6": {
										Type:        schema.TypeString,
										Description: "IPv6 link local address of the neighbor.",
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
						"network_type": {
							Type:        schema.TypeString,
							Description: "Network type.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Priority.",
							Computed:    true,
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Description: "Retransmit interval.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable OSPF6 routing on this interface.",
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
					},
				},
			},
			"restart_mode": {
				Type:        schema.TypeString,
				Description: "OSPFv3 restart mode (graceful or none).",
				Computed:    true,
			},
			"restart_on_topology_change": {
				Type:        schema.TypeString,
				Description: "Enable/disable continuing graceful restart upon topology change.",
				Computed:    true,
			},
			"restart_period": {
				Type:        schema.TypeInt,
				Description: "Graceful restart period in seconds.",
				Computed:    true,
			},
			"router_id": {
				Type:        schema.TypeString,
				Description: "A.B.C.D, in IPv4 address format.",
				Computed:    true,
			},
			"spf_timers": {
				Type:        schema.TypeString,
				Description: "SPF calculation frequency.",
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IPv6 address summary configuration.",
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
						"prefix6": {
							Type:        schema.TypeString,
							Description: "IPv6 prefix.",
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

func dataSourceRouterOspf6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterOspf6"

	o, err := c.Cmdb.ReadRouterOspf6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf6 dataSource: %v", err)
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

	diags := refreshObjectRouterOspf6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
