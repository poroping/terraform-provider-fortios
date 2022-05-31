// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure route maps.

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
)

func resourceRouterRouteMap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure route maps.",

		CreateContext: resourceRouterRouteMapCreate,
		ReadContext:   resourceRouterRouteMapRead,
		UpdateContext: resourceRouterRouteMapUpdate,
		DeleteContext: resourceRouterRouteMapDelete,

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
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "Rule.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"permit", "deny"}, false),

							Description: "Action.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Rule ID.",
							Optional:    true,
							Computed:    true,
						},
						"match_as_path": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Match BGP AS path list.",
							Optional:    true,
							Computed:    true,
						},
						"match_community": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Match BGP community list.",
							Optional:    true,
							Computed:    true,
						},
						"match_community_exact": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable exact matching of communities.",
							Optional:    true,
							Computed:    true,
						},
						"match_flags": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "BGP flag value to match (0 - 65535)",
							Optional:    true,
							Computed:    true,
						},
						"match_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Match interface configuration.",
							Optional:    true,
							Computed:    true,
						},
						"match_ip_address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Match IP address permitted by access-list or prefix-list.",
							Optional:    true,
							Computed:    true,
						},
						"match_ip_nexthop": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Match next hop IP address passed by access-list or prefix-list.",
							Optional:    true,
							Computed:    true,
						},
						"match_ip6_address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Match IPv6 address permitted by access-list6 or prefix-list6.",
							Optional:    true,
							Computed:    true,
						},
						"match_ip6_nexthop": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Match next hop IPv6 address passed by access-list6 or prefix-list6.",
							Optional:    true,
							Computed:    true,
						},
						"match_metric": {
							Type: schema.TypeInt,

							Description: "Match metric for redistribute routes.",
							Optional:    true,
							Computed:    true,
						},
						"match_origin": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "egp", "igp", "incomplete"}, false),

							Description: "Match BGP origin code.",
							Optional:    true,
							Computed:    true,
						},
						"match_route_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"external-type1", "external-type2", "none"}, false),

							Description: "Match route type.",
							Optional:    true,
							Computed:    true,
						},
						"match_tag": {
							Type: schema.TypeInt,

							Description: "Match tag.",
							Optional:    true,
							Computed:    true,
						},
						"match_vrf": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "Match VRF ID.",
							Optional:    true,
							Computed:    true,
						},
						"set_aggregator_as": {
							Type: schema.TypeInt,

							Description: "BGP aggregator AS.",
							Optional:    true,
							Computed:    true,
						},
						"set_aggregator_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "BGP aggregator IP.",
							Optional:    true,
							Computed:    true,
						},
						"set_aspath": {
							Type:        schema.TypeList,
							Description: "Prepend BGP AS path attribute.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"as": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "AS number (0 - 4294967295). Use quotes for repeating numbers, For example, \"1 1 2\".",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"set_aspath_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"prepend", "replace"}, false),

							Description: "Specify preferred action of set-aspath.",
							Optional:    true,
							Computed:    true,
						},
						"set_atomic_aggregate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable BGP atomic aggregate attribute.",
							Optional:    true,
							Computed:    true,
						},
						"set_community": {
							Type:        schema.TypeList,
							Description: "BGP community attribute.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Attribute: AA|AA:NN|internet|local-AS|no-advertise|no-export.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"set_community_additive": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable adding set-community to existing community.",
							Optional:    true,
							Computed:    true,
						},
						"set_community_delete": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Delete communities matching community list.",
							Optional:    true,
							Computed:    true,
						},
						"set_dampening_max_suppress": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Maximum duration to suppress a route (1 - 255 min, 0 = unset).",
							Optional:    true,
							Computed:    true,
						},
						"set_dampening_reachability_half_life": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 45),

							Description: "Reachability half-life time for the penalty (1 - 45 min, 0 = unset).",
							Optional:    true,
							Computed:    true,
						},
						"set_dampening_reuse": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 20000),

							Description: "Value to start reusing a route (1 - 20000, 0 = unset).",
							Optional:    true,
							Computed:    true,
						},
						"set_dampening_suppress": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 20000),

							Description: "Value to start suppressing a route (1 - 20000, 0 = unset).",
							Optional:    true,
							Computed:    true,
						},
						"set_dampening_unreachability_half_life": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 45),

							Description: "Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset).",
							Optional:    true,
							Computed:    true,
						},
						"set_extcommunity_rt": {
							Type:        schema.TypeList,
							Description: "Route Target extended community.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "AA:NN.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"set_extcommunity_soo": {
							Type:        schema.TypeList,
							Description: "Site-of-Origin extended community.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Community (format = AA:NN).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"set_flags": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "BGP flags value (0 - 65535)",
							Optional:    true,
							Computed:    true,
						},
						"set_ip_nexthop": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IP address of next hop.",
							Optional:    true,
							Computed:    true,
						},
						"set_ip6_nexthop": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 global address of next hop.",
							Optional:         true,
							Computed:         true,
						},
						"set_ip6_nexthop_local": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 local address of next hop.",
							Optional:         true,
							Computed:         true,
						},
						"set_local_preference": {
							Type: schema.TypeInt,

							Description: "BGP local preference path attribute.",
							Optional:    true,
							Computed:    true,
						},
						"set_metric": {
							Type: schema.TypeInt,

							Description: "Metric value.",
							Optional:    true,
							Computed:    true,
						},
						"set_metric_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"external-type1", "external-type2", "none"}, false),

							Description: "Metric type.",
							Optional:    true,
							Computed:    true,
						},
						"set_origin": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "egp", "igp", "incomplete"}, false),

							Description: "BGP origin code.",
							Optional:    true,
							Computed:    true,
						},
						"set_originator_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "BGP originator ID attribute.",
							Optional:    true,
							Computed:    true,
						},
						"set_priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Priority for routing table.",
							Optional:    true,
							Computed:    true,
						},
						"set_route_tag": {
							Type: schema.TypeInt,

							Description: "Route tag for routing table.",
							Optional:    true,
							Computed:    true,
						},
						"set_tag": {
							Type: schema.TypeInt,

							Description: "Tag value.",
							Optional:    true,
							Computed:    true,
						},
						"set_weight": {
							Type: schema.TypeInt,

							Description: "BGP weight for routing table.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterRouteMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterRouteMap resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterRouteMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterRouteMap(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRouteMap")
	}

	return resourceRouterRouteMapRead(ctx, d, meta)
}

func resourceRouterRouteMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterRouteMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterRouteMap(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterRouteMap resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRouteMap")
	}

	return resourceRouterRouteMapRead(ctx, d, meta)
}

func resourceRouterRouteMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterRouteMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterRouteMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRouteMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterRouteMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRouteMap resource: %v", err)
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

	diags := refreshObjectRouterRouteMap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterRouteMapRule(d *schema.ResourceData, v *[]models.RouterRouteMapRule, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.MatchAsPath; tmp != nil {
				v["match_as_path"] = *tmp
			}

			if tmp := cfg.MatchCommunity; tmp != nil {
				v["match_community"] = *tmp
			}

			if tmp := cfg.MatchCommunityExact; tmp != nil {
				v["match_community_exact"] = *tmp
			}

			if tmp := cfg.MatchFlags; tmp != nil {
				v["match_flags"] = *tmp
			}

			if tmp := cfg.MatchInterface; tmp != nil {
				v["match_interface"] = *tmp
			}

			if tmp := cfg.MatchIpAddress; tmp != nil {
				v["match_ip_address"] = *tmp
			}

			if tmp := cfg.MatchIpNexthop; tmp != nil {
				v["match_ip_nexthop"] = *tmp
			}

			if tmp := cfg.MatchIp6Address; tmp != nil {
				v["match_ip6_address"] = *tmp
			}

			if tmp := cfg.MatchIp6Nexthop; tmp != nil {
				v["match_ip6_nexthop"] = *tmp
			}

			if tmp := cfg.MatchMetric; tmp != nil {
				v["match_metric"] = *tmp
			}

			if tmp := cfg.MatchOrigin; tmp != nil {
				v["match_origin"] = *tmp
			}

			if tmp := cfg.MatchRouteType; tmp != nil {
				v["match_route_type"] = *tmp
			}

			if tmp := cfg.MatchTag; tmp != nil {
				v["match_tag"] = *tmp
			}

			if tmp := cfg.MatchVrf; tmp != nil {
				v["match_vrf"] = *tmp
			}

			if tmp := cfg.SetAggregatorAs; tmp != nil {
				v["set_aggregator_as"] = *tmp
			}

			if tmp := cfg.SetAggregatorIp; tmp != nil {
				v["set_aggregator_ip"] = *tmp
			}

			if tmp := cfg.SetAspath; tmp != nil {
				v["set_aspath"] = flattenRouterRouteMapRuleSetAspath(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "set_aspath"), sort)
			}

			if tmp := cfg.SetAspathAction; tmp != nil {
				v["set_aspath_action"] = *tmp
			}

			if tmp := cfg.SetAtomicAggregate; tmp != nil {
				v["set_atomic_aggregate"] = *tmp
			}

			if tmp := cfg.SetCommunity; tmp != nil {
				v["set_community"] = flattenRouterRouteMapRuleSetCommunity(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "set_community"), sort)
			}

			if tmp := cfg.SetCommunityAdditive; tmp != nil {
				v["set_community_additive"] = *tmp
			}

			if tmp := cfg.SetCommunityDelete; tmp != nil {
				v["set_community_delete"] = *tmp
			}

			if tmp := cfg.SetDampeningMaxSuppress; tmp != nil {
				v["set_dampening_max_suppress"] = *tmp
			}

			if tmp := cfg.SetDampeningReachabilityHalfLife; tmp != nil {
				v["set_dampening_reachability_half_life"] = *tmp
			}

			if tmp := cfg.SetDampeningReuse; tmp != nil {
				v["set_dampening_reuse"] = *tmp
			}

			if tmp := cfg.SetDampeningSuppress; tmp != nil {
				v["set_dampening_suppress"] = *tmp
			}

			if tmp := cfg.SetDampeningUnreachabilityHalfLife; tmp != nil {
				v["set_dampening_unreachability_half_life"] = *tmp
			}

			if tmp := cfg.SetExtcommunityRt; tmp != nil {
				v["set_extcommunity_rt"] = flattenRouterRouteMapRuleSetExtcommunityRt(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "set_extcommunity_rt"), sort)
			}

			if tmp := cfg.SetExtcommunitySoo; tmp != nil {
				v["set_extcommunity_soo"] = flattenRouterRouteMapRuleSetExtcommunitySoo(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "set_extcommunity_soo"), sort)
			}

			if tmp := cfg.SetFlags; tmp != nil {
				v["set_flags"] = *tmp
			}

			if tmp := cfg.SetIpNexthop; tmp != nil {
				v["set_ip_nexthop"] = *tmp
			}

			if tmp := cfg.SetIp6Nexthop; tmp != nil {
				v["set_ip6_nexthop"] = *tmp
			}

			if tmp := cfg.SetIp6NexthopLocal; tmp != nil {
				v["set_ip6_nexthop_local"] = *tmp
			}

			if tmp := cfg.SetLocalPreference; tmp != nil {
				v["set_local_preference"] = *tmp
			}

			if tmp := cfg.SetMetric; tmp != nil {
				v["set_metric"] = *tmp
			}

			if tmp := cfg.SetMetricType; tmp != nil {
				v["set_metric_type"] = *tmp
			}

			if tmp := cfg.SetOrigin; tmp != nil {
				v["set_origin"] = *tmp
			}

			if tmp := cfg.SetOriginatorId; tmp != nil {
				v["set_originator_id"] = *tmp
			}

			if tmp := cfg.SetPriority; tmp != nil {
				v["set_priority"] = *tmp
			}

			if tmp := cfg.SetRouteTag; tmp != nil {
				v["set_route_tag"] = *tmp
			}

			if tmp := cfg.SetTag; tmp != nil {
				v["set_tag"] = *tmp
			}

			if tmp := cfg.SetWeight; tmp != nil {
				v["set_weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterRouteMapRuleSetAspath(d *schema.ResourceData, v *[]models.RouterRouteMapRuleSetAspath, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.As; tmp != nil {
				v["as"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "as")
	}

	return flat
}

func flattenRouterRouteMapRuleSetCommunity(d *schema.ResourceData, v *[]models.RouterRouteMapRuleSetCommunity, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Community; tmp != nil {
				v["community"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "community")
	}

	return flat
}

func flattenRouterRouteMapRuleSetExtcommunityRt(d *schema.ResourceData, v *[]models.RouterRouteMapRuleSetExtcommunityRt, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Community; tmp != nil {
				v["community"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "community")
	}

	return flat
}

func flattenRouterRouteMapRuleSetExtcommunitySoo(d *schema.ResourceData, v *[]models.RouterRouteMapRuleSetExtcommunitySoo, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Community; tmp != nil {
				v["community"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "community")
	}

	return flat
}

func refreshObjectRouterRouteMap(d *schema.ResourceData, o *models.RouterRouteMap, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Rule != nil {
		if err = d.Set("rule", flattenRouterRouteMapRule(d, o.Rule, "rule", sort)); err != nil {
			return diag.Errorf("error reading rule: %v", err)
		}
	}

	return nil
}

func expandRouterRouteMapRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRouteMapRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRouteMapRule

	for i := range l {
		tmp := models.RouterRouteMapRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_as_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchAsPath = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchCommunity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_community_exact", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchCommunityExact = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_flags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MatchFlags = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_ip_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchIpAddress = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_ip_nexthop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchIpNexthop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_ip6_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchIp6Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_ip6_nexthop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchIp6Nexthop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MatchMetric = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_origin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchOrigin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_route_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MatchRouteType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MatchTag = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.match_vrf", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MatchVrf = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_aggregator_as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetAggregatorAs = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_aggregator_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetAggregatorIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_aspath", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterRouteMapRuleSetAspath(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterRouteMapRuleSetAspath
			// 	}
			tmp.SetAspath = v2

		}

		pre_append = fmt.Sprintf("%s.%d.set_aspath_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetAspathAction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_atomic_aggregate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetAtomicAggregate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterRouteMapRuleSetCommunity(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterRouteMapRuleSetCommunity
			// 	}
			tmp.SetCommunity = v2

		}

		pre_append = fmt.Sprintf("%s.%d.set_community_additive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetCommunityAdditive = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_community_delete", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetCommunityDelete = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_dampening_max_suppress", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetDampeningMaxSuppress = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_dampening_reachability_half_life", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetDampeningReachabilityHalfLife = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_dampening_reuse", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetDampeningReuse = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_dampening_suppress", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetDampeningSuppress = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_dampening_unreachability_half_life", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetDampeningUnreachabilityHalfLife = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_extcommunity_rt", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterRouteMapRuleSetExtcommunityRt(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterRouteMapRuleSetExtcommunityRt
			// 	}
			tmp.SetExtcommunityRt = v2

		}

		pre_append = fmt.Sprintf("%s.%d.set_extcommunity_soo", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandRouterRouteMapRuleSetExtcommunitySoo(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.RouterRouteMapRuleSetExtcommunitySoo
			// 	}
			tmp.SetExtcommunitySoo = v2

		}

		pre_append = fmt.Sprintf("%s.%d.set_flags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetFlags = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_ip_nexthop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetIpNexthop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_ip6_nexthop", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetIp6Nexthop = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_ip6_nexthop_local", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetIp6NexthopLocal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_local_preference", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetLocalPreference = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_metric", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetMetric = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_metric_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetMetricType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_origin", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetOrigin = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_originator_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SetOriginatorId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetPriority = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_route_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetRouteTag = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_tag", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetTag = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.set_weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SetWeight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRouteMapRuleSetAspath(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRouteMapRuleSetAspath, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRouteMapRuleSetAspath

	for i := range l {
		tmp := models.RouterRouteMapRuleSetAspath{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.as", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.As = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRouteMapRuleSetCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRouteMapRuleSetCommunity, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRouteMapRuleSetCommunity

	for i := range l {
		tmp := models.RouterRouteMapRuleSetCommunity{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Community = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRouteMapRuleSetExtcommunityRt(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRouteMapRuleSetExtcommunityRt, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRouteMapRuleSetExtcommunityRt

	for i := range l {
		tmp := models.RouterRouteMapRuleSetExtcommunityRt{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Community = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterRouteMapRuleSetExtcommunitySoo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterRouteMapRuleSetExtcommunitySoo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterRouteMapRuleSetExtcommunitySoo

	for i := range l {
		tmp := models.RouterRouteMapRuleSetExtcommunitySoo{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Community = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterRouteMap(d *schema.ResourceData, sv string) (*models.RouterRouteMap, diag.Diagnostics) {
	obj := models.RouterRouteMap{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v, ok := d.GetOk("rule"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("rule", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterRouteMapRule(d, v, "rule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Rule = t
		}
	} else if d.HasChange("rule") {
		old, new := d.GetChange("rule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Rule = &[]models.RouterRouteMapRule{}
		}
	}
	return &obj, diags
}
