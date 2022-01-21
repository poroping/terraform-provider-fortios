// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure router multicast.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Description: "Configure router multicast.",

		ReadContext: dataSourceRouterMulticastRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "PIM interfaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type:        schema.TypeString,
							Description: "Enable/disable Protocol Independent Multicast (PIM) Bidirectional Forwarding Detection (BFD).",
							Computed:    true,
						},
						"cisco_exclude_genid": {
							Type:        schema.TypeString,
							Description: "Exclude GenID from hello packets (compatibility with old Cisco IOS).",
							Computed:    true,
						},
						"dr_priority": {
							Type:        schema.TypeInt,
							Description: "DR election priority.",
							Computed:    true,
						},
						"hello_holdtime": {
							Type:        schema.TypeInt,
							Description: "Time before old neighbor information expires (0 - 65535 sec, default = 105).",
							Computed:    true,
						},
						"hello_interval": {
							Type:        schema.TypeInt,
							Description: "Interval between sending PIM hello messages (0 - 65535 sec, default = 30).",
							Computed:    true,
						},
						"igmp": {
							Type:        schema.TypeList,
							Description: "IGMP configuration options.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_group": {
										Type:        schema.TypeString,
										Description: "Groups IGMP hosts are allowed to join.",
										Computed:    true,
									},
									"immediate_leave_group": {
										Type:        schema.TypeString,
										Description: "Groups to drop membership for immediately after receiving IGMPv2 leave.",
										Computed:    true,
									},
									"last_member_query_count": {
										Type:        schema.TypeInt,
										Description: "Number of group specific queries before removing group (2 - 7, default = 2).",
										Computed:    true,
									},
									"last_member_query_interval": {
										Type:        schema.TypeInt,
										Description: "Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).",
										Computed:    true,
									},
									"query_interval": {
										Type:        schema.TypeInt,
										Description: "Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).",
										Computed:    true,
									},
									"query_max_response_time": {
										Type:        schema.TypeInt,
										Description: "Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).",
										Computed:    true,
									},
									"query_timeout": {
										Type:        schema.TypeInt,
										Description: "Timeout between queries before becoming querier for network (60 - 900, default = 255).",
										Computed:    true,
									},
									"router_alert_check": {
										Type:        schema.TypeString,
										Description: "Enable/disable require IGMP packets contain router alert option.",
										Computed:    true,
									},
									"version": {
										Type:        schema.TypeString,
										Description: "Maximum version of IGMP to support.",
										Computed:    true,
									},
								},
							},
						},
						"join_group": {
							Type:        schema.TypeList,
							Description: "Join multicast groups.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Description: "Multicast group IP address.",
										Computed:    true,
									},
								},
							},
						},
						"multicast_flow": {
							Type:        schema.TypeString,
							Description: "Acceptable source for multicast group.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"neighbour_filter": {
							Type:        schema.TypeString,
							Description: "Routers acknowledged as neighbor routers.",
							Computed:    true,
						},
						"passive": {
							Type:        schema.TypeString,
							Description: "Enable/disable listening to IGMP but not participating in PIM.",
							Computed:    true,
						},
						"pim_mode": {
							Type:        schema.TypeString,
							Description: "PIM operation mode.",
							Computed:    true,
						},
						"propagation_delay": {
							Type:        schema.TypeInt,
							Description: "Delay flooding packets on this interface (100 - 5000 msec, default = 500).",
							Computed:    true,
						},
						"rp_candidate": {
							Type:        schema.TypeString,
							Description: "Enable/disable compete to become RP in elections.",
							Computed:    true,
						},
						"rp_candidate_group": {
							Type:        schema.TypeString,
							Description: "Multicast groups managed by this RP.",
							Computed:    true,
						},
						"rp_candidate_interval": {
							Type:        schema.TypeInt,
							Description: "RP candidate advertisement interval (1 - 16383 sec, default = 60).",
							Computed:    true,
						},
						"rp_candidate_priority": {
							Type:        schema.TypeInt,
							Description: "Router's priority as RP.",
							Computed:    true,
						},
						"rpf_nbr_fail_back": {
							Type:        schema.TypeString,
							Description: "Enable/disable fail back for RPF neighbor query.",
							Computed:    true,
						},
						"rpf_nbr_fail_back_filter": {
							Type:        schema.TypeString,
							Description: "Filter for fail back RPF neighbors.",
							Computed:    true,
						},
						"state_refresh_interval": {
							Type:        schema.TypeInt,
							Description: "Interval between sending state-refresh packets (1 - 100 sec, default = 60).",
							Computed:    true,
						},
						"static_group": {
							Type:        schema.TypeString,
							Description: "Statically set multicast groups to forward out.",
							Computed:    true,
						},
						"ttl_threshold": {
							Type:        schema.TypeInt,
							Description: "Minimum TTL of multicast packets that will be forwarded (applied only to new multicast routes) (1 - 255, default = 1).",
							Computed:    true,
						},
					},
				},
			},
			"multicast_routing": {
				Type:        schema.TypeString,
				Description: "Enable/disable IP multicast routing.",
				Computed:    true,
			},
			"pim_sm_global": {
				Type:        schema.TypeList,
				Description: "PIM sparse-mode global settings.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_register_list": {
							Type:        schema.TypeString,
							Description: "Sources allowed to register packets with this Rendezvous Point (RP).",
							Computed:    true,
						},
						"accept_source_list": {
							Type:        schema.TypeString,
							Description: "Sources allowed to send multicast traffic.",
							Computed:    true,
						},
						"bsr_allow_quick_refresh": {
							Type:        schema.TypeString,
							Description: "Enable/disable accept BSR quick refresh packets from neighbors.",
							Computed:    true,
						},
						"bsr_candidate": {
							Type:        schema.TypeString,
							Description: "Enable/disable allowing this router to become a bootstrap router (BSR).",
							Computed:    true,
						},
						"bsr_hash": {
							Type:        schema.TypeInt,
							Description: "BSR hash length (0 - 32, default = 10).",
							Computed:    true,
						},
						"bsr_interface": {
							Type:        schema.TypeString,
							Description: "Interface to advertise as candidate BSR.",
							Computed:    true,
						},
						"bsr_priority": {
							Type:        schema.TypeInt,
							Description: "BSR priority (0 - 255, default = 0).",
							Computed:    true,
						},
						"cisco_crp_prefix": {
							Type:        schema.TypeString,
							Description: "Enable/disable making candidate RP compatible with old Cisco IOS.",
							Computed:    true,
						},
						"cisco_ignore_rp_set_priority": {
							Type:        schema.TypeString,
							Description: "Use only hash for RP selection (compatibility with old Cisco IOS).",
							Computed:    true,
						},
						"cisco_register_checksum": {
							Type:        schema.TypeString,
							Description: "Checksum entire register packet(for old Cisco IOS compatibility).",
							Computed:    true,
						},
						"cisco_register_checksum_group": {
							Type:        schema.TypeString,
							Description: "Cisco register checksum only these groups.",
							Computed:    true,
						},
						"join_prune_holdtime": {
							Type:        schema.TypeInt,
							Description: "Join/prune holdtime (1 - 65535, default = 210).",
							Computed:    true,
						},
						"message_interval": {
							Type:        schema.TypeInt,
							Description: "Period of time between sending periodic PIM join/prune messages in seconds (1 - 65535, default = 60).",
							Computed:    true,
						},
						"null_register_retries": {
							Type:        schema.TypeInt,
							Description: "Maximum retries of null register (1 - 20, default = 1).",
							Computed:    true,
						},
						"register_rate_limit": {
							Type:        schema.TypeInt,
							Description: "Limit of packets/sec per source registered through this RP (0 - 65535, default = 0 which means unlimited).",
							Computed:    true,
						},
						"register_rp_reachability": {
							Type:        schema.TypeString,
							Description: "Enable/disable check RP is reachable before registering packets.",
							Computed:    true,
						},
						"register_source": {
							Type:        schema.TypeString,
							Description: "Override source address in register packets.",
							Computed:    true,
						},
						"register_source_interface": {
							Type:        schema.TypeString,
							Description: "Override with primary interface address.",
							Computed:    true,
						},
						"register_source_ip": {
							Type:        schema.TypeString,
							Description: "Override with local IP address.",
							Computed:    true,
						},
						"register_supression": {
							Type:        schema.TypeInt,
							Description: "Period of time to honor register-stop message (1 - 65535 sec, default = 60).",
							Computed:    true,
						},
						"rp_address": {
							Type:        schema.TypeList,
							Description: "Statically configure RP addresses.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group": {
										Type:        schema.TypeString,
										Description: "Groups to use this RP.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "ID.",
										Computed:    true,
									},
									"ip_address": {
										Type:        schema.TypeString,
										Description: "RP router address.",
										Computed:    true,
									},
								},
							},
						},
						"rp_register_keepalive": {
							Type:        schema.TypeInt,
							Description: "Timeout for RP receiving data on (S,G) tree (1 - 65535 sec, default = 185).",
							Computed:    true,
						},
						"spt_threshold": {
							Type:        schema.TypeString,
							Description: "Enable/disable switching to source specific trees.",
							Computed:    true,
						},
						"spt_threshold_group": {
							Type:        schema.TypeString,
							Description: "Groups allowed to switch to source tree.",
							Computed:    true,
						},
						"ssm": {
							Type:        schema.TypeString,
							Description: "Enable/disable source specific multicast.",
							Computed:    true,
						},
						"ssm_range": {
							Type:        schema.TypeString,
							Description: "Groups allowed to source specific multicast.",
							Computed:    true,
						},
					},
				},
			},
			"route_limit": {
				Type:        schema.TypeInt,
				Description: "Maximum number of multicast routes.",
				Computed:    true,
			},
			"route_threshold": {
				Type:        schema.TypeInt,
				Description: "Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterMulticastRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterMulticast"

	o, err := c.Cmdb.ReadRouterMulticast(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterMulticast dataSource: %v", err)
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

	diags := refreshObjectRouterMulticast(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
