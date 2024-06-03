// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure router multicast.

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
)

func resourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Description: "Configure router multicast.",

		CreateContext: resourceRouterMulticastCreate,
		ReadContext:   resourceRouterMulticastRead,
		UpdateContext: resourceRouterMulticastUpdate,
		DeleteContext: resourceRouterMulticastDelete,

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
			"interface": {
				Type:        schema.TypeList,
				Description: "PIM interfaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Protocol Independent Multicast (PIM) Bidirectional Forwarding Detection (BFD).",
							Optional:    true,
							Computed:    true,
						},
						"cisco_exclude_genid": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Exclude GenID from hello packets (compatibility with old Cisco IOS).",
							Optional:    true,
							Computed:    true,
						},
						"dr_priority": {
							Type: schema.TypeInt,

							Description: "DR election priority.",
							Optional:    true,
							Computed:    true,
						},
						"hello_holdtime": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Time before old neighbor information expires (0 - 65535 sec, default = 105).",
							Optional:    true,
							Computed:    true,
						},
						"hello_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Interval between sending PIM hello messages (0 - 65535 sec, default = 30).",
							Optional:    true,
							Computed:    true,
						},
						"igmp": {
							Type:        schema.TypeList,
							Description: "IGMP configuration options.",
							Optional:    true, MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_group": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Groups IGMP hosts are allowed to join.",
										Optional:    true,
										Computed:    true,
									},
									"immediate_leave_group": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Groups to drop membership for immediately after receiving IGMPv2 leave.",
										Optional:    true,
										Computed:    true,
									},
									"last_member_query_count": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(2, 7),

										Description: "Number of group specific queries before removing group (2 - 7, default = 2).",
										Optional:    true,
										Computed:    true,
									},
									"last_member_query_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Timeout between IGMPv2 leave and removing group (1 - 65535 msec, default = 1000).",
										Optional:    true,
										Computed:    true,
									},
									"query_interval": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Interval between queries to IGMP hosts (1 - 65535 sec, default = 125).",
										Optional:    true,
										Computed:    true,
									},
									"query_max_response_time": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 25),

										Description: "Maximum time to wait for a IGMP query response (1 - 25 sec, default = 10).",
										Optional:    true,
										Computed:    true,
									},
									"query_timeout": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(60, 900),

										Description: "Timeout between queries before becoming querying unit for network (60 - 900, default = 255).",
										Optional:    true,
										Computed:    true,
									},
									"router_alert_check": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable require IGMP packets contain router alert option.",
										Optional:    true,
										Computed:    true,
									},
									"version": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"3", "2", "1"}, false),

										Description: "Maximum version of IGMP to support.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"join_group": {
							Type:        schema.TypeList,
							Description: "Join multicast groups.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "Multicast group IP address.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"multicast_flow": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Acceptable source for multicast group.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
						"neighbour_filter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Routers acknowledged as neighbor routers.",
							Optional:    true,
							Computed:    true,
						},
						"passive": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable listening to IGMP but not participating in PIM.",
							Optional:    true,
							Computed:    true,
						},
						"pim_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"sparse-mode", "dense-mode"}, false),

							Description: "PIM operation mode.",
							Optional:    true,
							Computed:    true,
						},
						"propagation_delay": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 5000),

							Description: "Delay flooding packets on this interface (100 - 5000 msec, default = 500).",
							Optional:    true,
							Computed:    true,
						},
						"rp_candidate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable compete to become RP in elections.",
							Optional:    true,
							Computed:    true,
						},
						"rp_candidate_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Multicast groups managed by this RP.",
							Optional:    true,
							Computed:    true,
						},
						"rp_candidate_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16383),

							Description: "RP candidate advertisement interval (1 - 16383 sec, default = 60).",
							Optional:    true,
							Computed:    true,
						},
						"rp_candidate_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Router's priority as RP.",
							Optional:    true,
							Computed:    true,
						},
						"rpf_nbr_fail_back": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable fail back for RPF neighbor query.",
							Optional:    true,
							Computed:    true,
						},
						"rpf_nbr_fail_back_filter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Filter for fail back RPF neighbors.",
							Optional:    true,
							Computed:    true,
						},
						"state_refresh_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),

							Description: "Interval between sending state-refresh packets (1 - 100 sec, default = 60).",
							Optional:    true,
							Computed:    true,
						},
						"static_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Statically set multicast groups to forward out.",
							Optional:    true,
							Computed:    true,
						},
						"ttl_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Minimum TTL of multicast packets that will be forwarded (applied only to new multicast routes) (1 - 255, default = 1).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"multicast_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IP multicast routing.",
				Optional:    true,
				Computed:    true,
			},
			"pim_sm_global": {
				Type:        schema.TypeList,
				Description: "PIM sparse-mode global settings.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accept_register_list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Sources allowed to register packets with this Rendezvous Point (RP).",
							Optional:    true,
							Computed:    true,
						},
						"accept_source_list": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Sources allowed to send multicast traffic.",
							Optional:    true,
							Computed:    true,
						},
						"bsr_allow_quick_refresh": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable accept BSR quick refresh packets from neighbors.",
							Optional:    true,
							Computed:    true,
						},
						"bsr_candidate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable allowing this router to become a bootstrap router (BSR).",
							Optional:    true,
							Computed:    true,
						},
						"bsr_hash": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),

							Description: "BSR hash length (0 - 32, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"bsr_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Interface to advertise as candidate BSR.",
							Optional:    true,
							Computed:    true,
						},
						"bsr_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "BSR priority (0 - 255, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"cisco_crp_prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable making candidate RP compatible with old Cisco IOS.",
							Optional:    true,
							Computed:    true,
						},
						"cisco_ignore_rp_set_priority": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Use only hash for RP selection (compatibility with old Cisco IOS).",
							Optional:    true,
							Computed:    true,
						},
						"cisco_register_checksum": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Checksum entire register packet(for old Cisco IOS compatibility).",
							Optional:    true,
							Computed:    true,
						},
						"cisco_register_checksum_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Cisco register checksum only these groups.",
							Optional:    true,
							Computed:    true,
						},
						"join_prune_holdtime": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Join/prune holdtime (1 - 65535, default = 210).",
							Optional:    true,
							Computed:    true,
						},
						"message_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Period of time between sending periodic PIM join/prune messages in seconds (1 - 65535, default = 60).",
							Optional:    true,
							Computed:    true,
						},
						"null_register_retries": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),

							Description: "Maximum retries of null register (1 - 20, default = 1).",
							Optional:    true,
							Computed:    true,
						},
						"register_rate_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Limit of packets/sec per source registered through this RP (0 - 65535, default = 0 which means unlimited).",
							Optional:    true,
							Computed:    true,
						},
						"register_rp_reachability": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable check RP is reachable before registering packets.",
							Optional:    true,
							Computed:    true,
						},
						"register_source": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "interface", "ip-address"}, false),

							Description: "Override source address in register packets.",
							Optional:    true,
							Computed:    true,
						},
						"register_source_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Override with primary interface address.",
							Optional:    true,
							Computed:    true,
						},
						"register_source_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Override with local IP address.",
							Optional:    true,
							Computed:    true,
						},
						"register_supression": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Period of time to honor register-stop message (1 - 65535 sec, default = 60).",
							Optional:    true,
							Computed:    true,
						},
						"rp_address": {
							Type:        schema.TypeList,
							Description: "Statically configure RP addresses.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Groups to use this RP.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "ID.",
										Optional:    true,
										Computed:    true,
									},
									"ip_address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "RP router address.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"rp_register_keepalive": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Timeout for RP receiving data on (S,G) tree (1 - 65535 sec, default = 185).",
							Optional:    true,
							Computed:    true,
						},
						"spt_threshold": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable switching to source specific trees.",
							Optional:    true,
							Computed:    true,
						},
						"spt_threshold_group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Groups allowed to switch to source tree.",
							Optional:    true,
							Computed:    true,
						},
						"ssm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable source specific multicast.",
							Optional:    true,
							Computed:    true,
						},
						"ssm_range": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Groups allowed to source specific multicast.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"route_limit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Maximum number of multicast routes.",
				Optional:    true,
				Computed:    true,
			},
			"route_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147483647),

				Description: "Generate warnings when the number of multicast routes exceeds this number, must not be greater than route-limit.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterMulticastCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterMulticast(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterMulticast(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterMulticast")
	}

	return resourceRouterMulticastRead(ctx, d, meta)
}

func resourceRouterMulticastUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterMulticast(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterMulticast(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterMulticast resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterMulticast")
	}

	return resourceRouterMulticastRead(ctx, d, meta)
}

func resourceRouterMulticastDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectRouterMulticast(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateRouterMulticast(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterMulticast resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterMulticast(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterMulticast resource: %v", err)
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

	diags := refreshObjectRouterMulticast(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterMulticastInterface(d *schema.ResourceData, v *[]models.RouterMulticastInterface, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Bfd; tmp != nil {
				v["bfd"] = *tmp
			}

			if tmp := cfg.CiscoExcludeGenid; tmp != nil {
				v["cisco_exclude_genid"] = *tmp
			}

			if tmp := cfg.DrPriority; tmp != nil {
				v["dr_priority"] = *tmp
			}

			if tmp := cfg.HelloHoldtime; tmp != nil {
				v["hello_holdtime"] = *tmp
			}

			if tmp := cfg.HelloInterval; tmp != nil {
				v["hello_interval"] = *tmp
			}

			if tmp := cfg.Igmp; tmp != nil {
				v["igmp"] = flattenRouterMulticastInterfaceIgmp(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "igmp"), sort)
			}

			if tmp := cfg.JoinGroup; tmp != nil {
				v["join_group"] = flattenRouterMulticastInterfaceJoinGroup(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "join_group"), sort)
			}

			if tmp := cfg.MulticastFlow; tmp != nil {
				v["multicast_flow"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.NeighbourFilter; tmp != nil {
				v["neighbour_filter"] = *tmp
			}

			if tmp := cfg.Passive; tmp != nil {
				v["passive"] = *tmp
			}

			if tmp := cfg.PimMode; tmp != nil {
				v["pim_mode"] = *tmp
			}

			if tmp := cfg.PropagationDelay; tmp != nil {
				v["propagation_delay"] = *tmp
			}

			if tmp := cfg.RpCandidate; tmp != nil {
				v["rp_candidate"] = *tmp
			}

			if tmp := cfg.RpCandidateGroup; tmp != nil {
				v["rp_candidate_group"] = *tmp
			}

			if tmp := cfg.RpCandidateInterval; tmp != nil {
				v["rp_candidate_interval"] = *tmp
			}

			if tmp := cfg.RpCandidatePriority; tmp != nil {
				v["rp_candidate_priority"] = *tmp
			}

			if tmp := cfg.RpfNbrFailBack; tmp != nil {
				v["rpf_nbr_fail_back"] = *tmp
			}

			if tmp := cfg.RpfNbrFailBackFilter; tmp != nil {
				v["rpf_nbr_fail_back_filter"] = *tmp
			}

			if tmp := cfg.StateRefreshInterval; tmp != nil {
				v["state_refresh_interval"] = *tmp
			}

			if tmp := cfg.StaticGroup; tmp != nil {
				v["static_group"] = *tmp
			}

			if tmp := cfg.TtlThreshold; tmp != nil {
				v["ttl_threshold"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenRouterMulticastInterfaceIgmp(d *schema.ResourceData, v *models.RouterMulticastInterfaceIgmp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.RouterMulticastInterfaceIgmp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AccessGroup; tmp != nil {
				v["access_group"] = *tmp
			}

			if tmp := cfg.ImmediateLeaveGroup; tmp != nil {
				v["immediate_leave_group"] = *tmp
			}

			if tmp := cfg.LastMemberQueryCount; tmp != nil {
				v["last_member_query_count"] = *tmp
			}

			if tmp := cfg.LastMemberQueryInterval; tmp != nil {
				v["last_member_query_interval"] = *tmp
			}

			if tmp := cfg.QueryInterval; tmp != nil {
				v["query_interval"] = *tmp
			}

			if tmp := cfg.QueryMaxResponseTime; tmp != nil {
				v["query_max_response_time"] = *tmp
			}

			if tmp := cfg.QueryTimeout; tmp != nil {
				v["query_timeout"] = *tmp
			}

			if tmp := cfg.RouterAlertCheck; tmp != nil {
				v["router_alert_check"] = *tmp
			}

			if tmp := cfg.Version; tmp != nil {
				v["version"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenRouterMulticastInterfaceJoinGroup(d *schema.ResourceData, v *[]models.RouterMulticastInterfaceJoinGroup, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "address")
	}

	return flat
}

func flattenRouterMulticastPimSmGlobal(d *schema.ResourceData, v *models.RouterMulticastPimSmGlobal, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.RouterMulticastPimSmGlobal{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AcceptRegisterList; tmp != nil {
				v["accept_register_list"] = *tmp
			}

			if tmp := cfg.AcceptSourceList; tmp != nil {
				v["accept_source_list"] = *tmp
			}

			if tmp := cfg.BsrAllowQuickRefresh; tmp != nil {
				v["bsr_allow_quick_refresh"] = *tmp
			}

			if tmp := cfg.BsrCandidate; tmp != nil {
				v["bsr_candidate"] = *tmp
			}

			if tmp := cfg.BsrHash; tmp != nil {
				v["bsr_hash"] = *tmp
			}

			if tmp := cfg.BsrInterface; tmp != nil {
				v["bsr_interface"] = *tmp
			}

			if tmp := cfg.BsrPriority; tmp != nil {
				v["bsr_priority"] = *tmp
			}

			if tmp := cfg.CiscoCrpPrefix; tmp != nil {
				v["cisco_crp_prefix"] = *tmp
			}

			if tmp := cfg.CiscoIgnoreRpSetPriority; tmp != nil {
				v["cisco_ignore_rp_set_priority"] = *tmp
			}

			if tmp := cfg.CiscoRegisterChecksum; tmp != nil {
				v["cisco_register_checksum"] = *tmp
			}

			if tmp := cfg.CiscoRegisterChecksumGroup; tmp != nil {
				v["cisco_register_checksum_group"] = *tmp
			}

			if tmp := cfg.JoinPruneHoldtime; tmp != nil {
				v["join_prune_holdtime"] = *tmp
			}

			if tmp := cfg.MessageInterval; tmp != nil {
				v["message_interval"] = *tmp
			}

			if tmp := cfg.NullRegisterRetries; tmp != nil {
				v["null_register_retries"] = *tmp
			}

			if tmp := cfg.RegisterRateLimit; tmp != nil {
				v["register_rate_limit"] = *tmp
			}

			if tmp := cfg.RegisterRpReachability; tmp != nil {
				v["register_rp_reachability"] = *tmp
			}

			if tmp := cfg.RegisterSource; tmp != nil {
				v["register_source"] = *tmp
			}

			if tmp := cfg.RegisterSourceInterface; tmp != nil {
				v["register_source_interface"] = *tmp
			}

			if tmp := cfg.RegisterSourceIp; tmp != nil {
				v["register_source_ip"] = *tmp
			}

			if tmp := cfg.RegisterSupression; tmp != nil {
				v["register_supression"] = *tmp
			}

			if tmp := cfg.RpAddress; tmp != nil {
				v["rp_address"] = flattenRouterMulticastPimSmGlobalRpAddress(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "rp_address"), sort)
			}

			if tmp := cfg.RpRegisterKeepalive; tmp != nil {
				v["rp_register_keepalive"] = *tmp
			}

			if tmp := cfg.SptThreshold; tmp != nil {
				v["spt_threshold"] = *tmp
			}

			if tmp := cfg.SptThresholdGroup; tmp != nil {
				v["spt_threshold_group"] = *tmp
			}

			if tmp := cfg.Ssm; tmp != nil {
				v["ssm"] = *tmp
			}

			if tmp := cfg.SsmRange; tmp != nil {
				v["ssm_range"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenRouterMulticastPimSmGlobalRpAddress(d *schema.ResourceData, v *[]models.RouterMulticastPimSmGlobalRpAddress, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Group; tmp != nil {
				v["group"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.IpAddress; tmp != nil {
				v["ip_address"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterMulticast(d *schema.ResourceData, o *models.RouterMulticast, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Interface != nil {
		if err = d.Set("interface", flattenRouterMulticastInterface(d, o.Interface, "interface", sort)); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.MulticastRouting != nil {
		v := *o.MulticastRouting

		if err = d.Set("multicast_routing", v); err != nil {
			return diag.Errorf("error reading multicast_routing: %v", err)
		}
	}

	if _, ok := d.GetOk("pim_sm_global"); ok {
		if o.PimSmGlobal != nil {
			if err = d.Set("pim_sm_global", flattenRouterMulticastPimSmGlobal(d, o.PimSmGlobal, "pim_sm_global", sort)); err != nil {
				return diag.Errorf("error reading pim_sm_global: %v", err)
			}
		}
	}

	if o.RouteLimit != nil {
		v := *o.RouteLimit

		if err = d.Set("route_limit", v); err != nil {
			return diag.Errorf("error reading route_limit: %v", err)
		}
	}

	if o.RouteThreshold != nil {
		v := *o.RouteThreshold

		if err = d.Set("route_threshold", v); err != nil {
			return diag.Errorf("error reading route_threshold: %v", err)
		}
	}

	return nil
}

func expandRouterMulticastInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterMulticastInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticastInterface

	for i := range l {
		tmp := models.RouterMulticastInterface{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.bfd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Bfd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cisco_exclude_genid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CiscoExcludeGenid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dr_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DrPriority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_holdtime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HelloHoldtime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hello_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HelloInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.igmp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterMulticastInterfaceIgmp(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterMulticastInterfaceIgmp
			// 	}
			tmp.Igmp = v2

		}

		pre_append = fmt.Sprintf("%s.%d.join_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterMulticastInterfaceJoinGroup(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterMulticastInterfaceJoinGroup
			// 	}
			tmp.JoinGroup = v2

		}

		pre_append = fmt.Sprintf("%s.%d.multicast_flow", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MulticastFlow = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.neighbour_filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NeighbourFilter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.passive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Passive = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pim_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PimMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.propagation_delay", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PropagationDelay = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rp_candidate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RpCandidate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rp_candidate_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RpCandidateGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rp_candidate_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RpCandidateInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rp_candidate_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RpCandidatePriority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rpf_nbr_fail_back", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RpfNbrFailBack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rpf_nbr_fail_back_filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RpfNbrFailBackFilter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.state_refresh_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StateRefreshInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.static_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StaticGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ttl_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TtlThreshold = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterMulticastInterfaceIgmp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.RouterMulticastInterfaceIgmp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticastInterfaceIgmp

	for i := range l {
		tmp := models.RouterMulticastInterfaceIgmp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.access_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AccessGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.immediate_leave_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImmediateLeaveGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.last_member_query_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LastMemberQueryCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.last_member_query_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LastMemberQueryInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.QueryInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_max_response_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.QueryMaxResponseTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.QueryTimeout = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.router_alert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RouterAlertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Version = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandRouterMulticastInterfaceJoinGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterMulticastInterfaceJoinGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticastInterfaceJoinGroup

	for i := range l {
		tmp := models.RouterMulticastInterfaceJoinGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterMulticastPimSmGlobal(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.RouterMulticastPimSmGlobal, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticastPimSmGlobal

	for i := range l {
		tmp := models.RouterMulticastPimSmGlobal{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.accept_register_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AcceptRegisterList = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.accept_source_list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AcceptSourceList = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bsr_allow_quick_refresh", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BsrAllowQuickRefresh = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bsr_candidate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BsrCandidate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bsr_hash", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BsrHash = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bsr_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BsrInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bsr_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BsrPriority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cisco_crp_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CiscoCrpPrefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cisco_ignore_rp_set_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CiscoIgnoreRpSetPriority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cisco_register_checksum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CiscoRegisterChecksum = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cisco_register_checksum_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CiscoRegisterChecksumGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.join_prune_holdtime", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.JoinPruneHoldtime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.message_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MessageInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.null_register_retries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.NullRegisterRetries = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_rate_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RegisterRateLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_rp_reachability", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RegisterRpReachability = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_source", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RegisterSource = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_source_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RegisterSourceInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_source_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RegisterSourceIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_supression", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RegisterSupression = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rp_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterMulticastPimSmGlobalRpAddress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterMulticastPimSmGlobalRpAddress
			// 	}
			tmp.RpAddress = v2

		}

		pre_append = fmt.Sprintf("%s.%d.rp_register_keepalive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RpRegisterKeepalive = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spt_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SptThreshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.spt_threshold_group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SptThresholdGroup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ssm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssm_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsmRange = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandRouterMulticastPimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterMulticastPimSmGlobalRpAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterMulticastPimSmGlobalRpAddress

	for i := range l {
		tmp := models.RouterMulticastPimSmGlobalRpAddress{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Group = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpAddress = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterMulticast(d *schema.ResourceData, sv string) (*models.RouterMulticast, diag.Diagnostics) {
	obj := models.RouterMulticast{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("interface", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterMulticastInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Interface = t
		}
	} else if d.HasChange("interface") {
		old, new := d.GetChange("interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Interface = &[]models.RouterMulticastInterface{}
		}
	}
	if v1, ok := d.GetOk("multicast_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_routing", sv)
				diags = append(diags, e)
			}
			obj.MulticastRouting = &v2
		}
	}
	if v, ok := d.GetOk("pim_sm_global"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pim_sm_global", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterMulticastPimSmGlobal(d, v, "pim_sm_global", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PimSmGlobal = t
		}
	} else if d.HasChange("pim_sm_global") {
		old, new := d.GetChange("pim_sm_global")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PimSmGlobal = &models.RouterMulticastPimSmGlobal{}
		}
	}
	if v1, ok := d.GetOk("route_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RouteLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("route_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("route_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RouteThreshold = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectRouterMulticast(d *schema.ResourceData, sv string) (*models.RouterMulticast, diag.Diagnostics) {
	obj := models.RouterMulticast{}
	diags := diag.Diagnostics{}

	obj.Interface = &[]models.RouterMulticastInterface{}
	obj.PimSmGlobal = &models.RouterMulticastPimSmGlobal{}

	return &obj, diags
}
