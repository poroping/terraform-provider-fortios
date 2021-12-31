// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Rule.

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

func resourceRouterrouteMapRule() *schema.Resource {
	return &schema.Resource{
		Description: "Rule.",

		CreateContext: resourceRouterrouteMapRuleCreate,
		ReadContext:   resourceRouterrouteMapRuleRead,
		UpdateContext: resourceRouterrouteMapRuleUpdate,
		DeleteContext: resourceRouterrouteMapRuleDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"permit", "deny"}, false),

				Description: "Action.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Rule ID.",
				ForceNew:    true,
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
				ValidateFunc: validation.IntBetween(0, 31),

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

							Description: "AS number (0 - 4294967295). NOTE: Use quotes for repeating numbers, e.g.: \"1 1 2\"\n",
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

				Description: "Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)",
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

							Description: "AA:NN",
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
	}
}

func resourceRouterrouteMapRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterrouteMapRule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterrouteMapRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterrouteMapRule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterrouteMapRule")
	}

	return resourceRouterrouteMapRuleRead(ctx, d, meta)
}

func resourceRouterrouteMapRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterrouteMapRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterrouteMapRule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterrouteMapRule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterrouteMapRule")
	}

	return resourceRouterrouteMapRuleRead(ctx, d, meta)
}

func resourceRouterrouteMapRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterrouteMapRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterrouteMapRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterrouteMapRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterrouteMapRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterrouteMapRule resource: %v", err)
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

	diags := refreshObjectRouterrouteMapRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterrouteMapRuleSetAspath(v *[]models.RouterrouteMapRuleSetAspath, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenRouterrouteMapRuleSetCommunity(v *[]models.RouterrouteMapRuleSetCommunity, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenRouterrouteMapRuleSetExtcommunityRt(v *[]models.RouterrouteMapRuleSetExtcommunityRt, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenRouterrouteMapRuleSetExtcommunitySoo(v *[]models.RouterrouteMapRuleSetExtcommunitySoo, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectRouterrouteMapRule(d *schema.ResourceData, o *models.RouterrouteMapRule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.MatchAsPath != nil {
		v := *o.MatchAsPath

		if err = d.Set("match_as_path", v); err != nil {
			return diag.Errorf("error reading match_as_path: %v", err)
		}
	}

	if o.MatchCommunity != nil {
		v := *o.MatchCommunity

		if err = d.Set("match_community", v); err != nil {
			return diag.Errorf("error reading match_community: %v", err)
		}
	}

	if o.MatchCommunityExact != nil {
		v := *o.MatchCommunityExact

		if err = d.Set("match_community_exact", v); err != nil {
			return diag.Errorf("error reading match_community_exact: %v", err)
		}
	}

	if o.MatchFlags != nil {
		v := *o.MatchFlags

		if err = d.Set("match_flags", v); err != nil {
			return diag.Errorf("error reading match_flags: %v", err)
		}
	}

	if o.MatchInterface != nil {
		v := *o.MatchInterface

		if err = d.Set("match_interface", v); err != nil {
			return diag.Errorf("error reading match_interface: %v", err)
		}
	}

	if o.MatchIpAddress != nil {
		v := *o.MatchIpAddress

		if err = d.Set("match_ip_address", v); err != nil {
			return diag.Errorf("error reading match_ip_address: %v", err)
		}
	}

	if o.MatchIpNexthop != nil {
		v := *o.MatchIpNexthop

		if err = d.Set("match_ip_nexthop", v); err != nil {
			return diag.Errorf("error reading match_ip_nexthop: %v", err)
		}
	}

	if o.MatchIp6Address != nil {
		v := *o.MatchIp6Address

		if err = d.Set("match_ip6_address", v); err != nil {
			return diag.Errorf("error reading match_ip6_address: %v", err)
		}
	}

	if o.MatchIp6Nexthop != nil {
		v := *o.MatchIp6Nexthop

		if err = d.Set("match_ip6_nexthop", v); err != nil {
			return diag.Errorf("error reading match_ip6_nexthop: %v", err)
		}
	}

	if o.MatchMetric != nil {
		v := *o.MatchMetric

		if err = d.Set("match_metric", v); err != nil {
			return diag.Errorf("error reading match_metric: %v", err)
		}
	}

	if o.MatchOrigin != nil {
		v := *o.MatchOrigin

		if err = d.Set("match_origin", v); err != nil {
			return diag.Errorf("error reading match_origin: %v", err)
		}
	}

	if o.MatchRouteType != nil {
		v := *o.MatchRouteType

		if err = d.Set("match_route_type", v); err != nil {
			return diag.Errorf("error reading match_route_type: %v", err)
		}
	}

	if o.MatchTag != nil {
		v := *o.MatchTag

		if err = d.Set("match_tag", v); err != nil {
			return diag.Errorf("error reading match_tag: %v", err)
		}
	}

	if o.MatchVrf != nil {
		v := *o.MatchVrf

		if err = d.Set("match_vrf", v); err != nil {
			return diag.Errorf("error reading match_vrf: %v", err)
		}
	}

	if o.SetAggregatorAs != nil {
		v := *o.SetAggregatorAs

		if err = d.Set("set_aggregator_as", v); err != nil {
			return diag.Errorf("error reading set_aggregator_as: %v", err)
		}
	}

	if o.SetAggregatorIp != nil {
		v := *o.SetAggregatorIp

		if err = d.Set("set_aggregator_ip", v); err != nil {
			return diag.Errorf("error reading set_aggregator_ip: %v", err)
		}
	}

	if o.SetAspath != nil {
		if err = d.Set("set_aspath", flattenRouterrouteMapRuleSetAspath(o.SetAspath, sort)); err != nil {
			return diag.Errorf("error reading set_aspath: %v", err)
		}
	}

	if o.SetAspathAction != nil {
		v := *o.SetAspathAction

		if err = d.Set("set_aspath_action", v); err != nil {
			return diag.Errorf("error reading set_aspath_action: %v", err)
		}
	}

	if o.SetAtomicAggregate != nil {
		v := *o.SetAtomicAggregate

		if err = d.Set("set_atomic_aggregate", v); err != nil {
			return diag.Errorf("error reading set_atomic_aggregate: %v", err)
		}
	}

	if o.SetCommunity != nil {
		if err = d.Set("set_community", flattenRouterrouteMapRuleSetCommunity(o.SetCommunity, sort)); err != nil {
			return diag.Errorf("error reading set_community: %v", err)
		}
	}

	if o.SetCommunityAdditive != nil {
		v := *o.SetCommunityAdditive

		if err = d.Set("set_community_additive", v); err != nil {
			return diag.Errorf("error reading set_community_additive: %v", err)
		}
	}

	if o.SetCommunityDelete != nil {
		v := *o.SetCommunityDelete

		if err = d.Set("set_community_delete", v); err != nil {
			return diag.Errorf("error reading set_community_delete: %v", err)
		}
	}

	if o.SetDampeningMaxSuppress != nil {
		v := *o.SetDampeningMaxSuppress

		if err = d.Set("set_dampening_max_suppress", v); err != nil {
			return diag.Errorf("error reading set_dampening_max_suppress: %v", err)
		}
	}

	if o.SetDampeningReachabilityHalfLife != nil {
		v := *o.SetDampeningReachabilityHalfLife

		if err = d.Set("set_dampening_reachability_half_life", v); err != nil {
			return diag.Errorf("error reading set_dampening_reachability_half_life: %v", err)
		}
	}

	if o.SetDampeningReuse != nil {
		v := *o.SetDampeningReuse

		if err = d.Set("set_dampening_reuse", v); err != nil {
			return diag.Errorf("error reading set_dampening_reuse: %v", err)
		}
	}

	if o.SetDampeningSuppress != nil {
		v := *o.SetDampeningSuppress

		if err = d.Set("set_dampening_suppress", v); err != nil {
			return diag.Errorf("error reading set_dampening_suppress: %v", err)
		}
	}

	if o.SetDampeningUnreachabilityHalfLife != nil {
		v := *o.SetDampeningUnreachabilityHalfLife

		if err = d.Set("set_dampening_unreachability_half_life", v); err != nil {
			return diag.Errorf("error reading set_dampening_unreachability_half_life: %v", err)
		}
	}

	if o.SetExtcommunityRt != nil {
		if err = d.Set("set_extcommunity_rt", flattenRouterrouteMapRuleSetExtcommunityRt(o.SetExtcommunityRt, sort)); err != nil {
			return diag.Errorf("error reading set_extcommunity_rt: %v", err)
		}
	}

	if o.SetExtcommunitySoo != nil {
		if err = d.Set("set_extcommunity_soo", flattenRouterrouteMapRuleSetExtcommunitySoo(o.SetExtcommunitySoo, sort)); err != nil {
			return diag.Errorf("error reading set_extcommunity_soo: %v", err)
		}
	}

	if o.SetFlags != nil {
		v := *o.SetFlags

		if err = d.Set("set_flags", v); err != nil {
			return diag.Errorf("error reading set_flags: %v", err)
		}
	}

	if o.SetIpNexthop != nil {
		v := *o.SetIpNexthop

		if err = d.Set("set_ip_nexthop", v); err != nil {
			return diag.Errorf("error reading set_ip_nexthop: %v", err)
		}
	}

	if o.SetIp6Nexthop != nil {
		v := *o.SetIp6Nexthop

		if err = d.Set("set_ip6_nexthop", v); err != nil {
			return diag.Errorf("error reading set_ip6_nexthop: %v", err)
		}
	}

	if o.SetIp6NexthopLocal != nil {
		v := *o.SetIp6NexthopLocal

		if err = d.Set("set_ip6_nexthop_local", v); err != nil {
			return diag.Errorf("error reading set_ip6_nexthop_local: %v", err)
		}
	}

	if o.SetLocalPreference != nil {
		v := *o.SetLocalPreference

		if err = d.Set("set_local_preference", v); err != nil {
			return diag.Errorf("error reading set_local_preference: %v", err)
		}
	}

	if o.SetMetric != nil {
		v := *o.SetMetric

		if err = d.Set("set_metric", v); err != nil {
			return diag.Errorf("error reading set_metric: %v", err)
		}
	}

	if o.SetMetricType != nil {
		v := *o.SetMetricType

		if err = d.Set("set_metric_type", v); err != nil {
			return diag.Errorf("error reading set_metric_type: %v", err)
		}
	}

	if o.SetOrigin != nil {
		v := *o.SetOrigin

		if err = d.Set("set_origin", v); err != nil {
			return diag.Errorf("error reading set_origin: %v", err)
		}
	}

	if o.SetOriginatorId != nil {
		v := *o.SetOriginatorId

		if err = d.Set("set_originator_id", v); err != nil {
			return diag.Errorf("error reading set_originator_id: %v", err)
		}
	}

	if o.SetRouteTag != nil {
		v := *o.SetRouteTag

		if err = d.Set("set_route_tag", v); err != nil {
			return diag.Errorf("error reading set_route_tag: %v", err)
		}
	}

	if o.SetTag != nil {
		v := *o.SetTag

		if err = d.Set("set_tag", v); err != nil {
			return diag.Errorf("error reading set_tag: %v", err)
		}
	}

	if o.SetWeight != nil {
		v := *o.SetWeight

		if err = d.Set("set_weight", v); err != nil {
			return diag.Errorf("error reading set_weight: %v", err)
		}
	}

	return nil
}

func expandRouterrouteMapRuleSetAspath(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterrouteMapRuleSetAspath, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterrouteMapRuleSetAspath

	for i := range l {
		tmp := models.RouterrouteMapRuleSetAspath{}
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

func expandRouterrouteMapRuleSetCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterrouteMapRuleSetCommunity, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterrouteMapRuleSetCommunity

	for i := range l {
		tmp := models.RouterrouteMapRuleSetCommunity{}
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

func expandRouterrouteMapRuleSetExtcommunityRt(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterrouteMapRuleSetExtcommunityRt, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterrouteMapRuleSetExtcommunityRt

	for i := range l {
		tmp := models.RouterrouteMapRuleSetExtcommunityRt{}
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

func expandRouterrouteMapRuleSetExtcommunitySoo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterrouteMapRuleSetExtcommunitySoo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterrouteMapRuleSetExtcommunitySoo

	for i := range l {
		tmp := models.RouterrouteMapRuleSetExtcommunitySoo{}
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

func getObjectRouterrouteMapRule(d *schema.ResourceData, sv string) (*models.RouterrouteMapRule, diag.Diagnostics) {
	obj := models.RouterrouteMapRule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("match_as_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_as_path", sv)
				diags = append(diags, e)
			}
			obj.MatchAsPath = &v2
		}
	}
	if v1, ok := d.GetOk("match_community"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_community", sv)
				diags = append(diags, e)
			}
			obj.MatchCommunity = &v2
		}
	}
	if v1, ok := d.GetOk("match_community_exact"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_community_exact", sv)
				diags = append(diags, e)
			}
			obj.MatchCommunityExact = &v2
		}
	}
	if v1, ok := d.GetOk("match_flags"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("match_flags", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MatchFlags = &tmp
		}
	}
	if v1, ok := d.GetOk("match_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_interface", sv)
				diags = append(diags, e)
			}
			obj.MatchInterface = &v2
		}
	}
	if v1, ok := d.GetOk("match_ip_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_ip_address", sv)
				diags = append(diags, e)
			}
			obj.MatchIpAddress = &v2
		}
	}
	if v1, ok := d.GetOk("match_ip_nexthop"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_ip_nexthop", sv)
				diags = append(diags, e)
			}
			obj.MatchIpNexthop = &v2
		}
	}
	if v1, ok := d.GetOk("match_ip6_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_ip6_address", sv)
				diags = append(diags, e)
			}
			obj.MatchIp6Address = &v2
		}
	}
	if v1, ok := d.GetOk("match_ip6_nexthop"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_ip6_nexthop", sv)
				diags = append(diags, e)
			}
			obj.MatchIp6Nexthop = &v2
		}
	}
	if v1, ok := d.GetOk("match_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MatchMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("match_origin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_origin", sv)
				diags = append(diags, e)
			}
			obj.MatchOrigin = &v2
		}
	}
	if v1, ok := d.GetOk("match_route_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_route_type", sv)
				diags = append(diags, e)
			}
			obj.MatchRouteType = &v2
		}
	}
	if v1, ok := d.GetOk("match_tag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_tag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MatchTag = &tmp
		}
	}
	if v1, ok := d.GetOk("match_vrf"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("match_vrf", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MatchVrf = &tmp
		}
	}
	if v1, ok := d.GetOk("set_aggregator_as"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_aggregator_as", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetAggregatorAs = &tmp
		}
	}
	if v1, ok := d.GetOk("set_aggregator_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_aggregator_ip", sv)
				diags = append(diags, e)
			}
			obj.SetAggregatorIp = &v2
		}
	}
	if v, ok := d.GetOk("set_aspath"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("set_aspath", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterrouteMapRuleSetAspath(d, v, "set_aspath", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SetAspath = t
		}
	} else if d.HasChange("set_aspath") {
		old, new := d.GetChange("set_aspath")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SetAspath = &[]models.RouterrouteMapRuleSetAspath{}
		}
	}
	if v1, ok := d.GetOk("set_aspath_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_aspath_action", sv)
				diags = append(diags, e)
			}
			obj.SetAspathAction = &v2
		}
	}
	if v1, ok := d.GetOk("set_atomic_aggregate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_atomic_aggregate", sv)
				diags = append(diags, e)
			}
			obj.SetAtomicAggregate = &v2
		}
	}
	if v, ok := d.GetOk("set_community"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("set_community", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterrouteMapRuleSetCommunity(d, v, "set_community", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SetCommunity = t
		}
	} else if d.HasChange("set_community") {
		old, new := d.GetChange("set_community")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SetCommunity = &[]models.RouterrouteMapRuleSetCommunity{}
		}
	}
	if v1, ok := d.GetOk("set_community_additive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_community_additive", sv)
				diags = append(diags, e)
			}
			obj.SetCommunityAdditive = &v2
		}
	}
	if v1, ok := d.GetOk("set_community_delete"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_community_delete", sv)
				diags = append(diags, e)
			}
			obj.SetCommunityDelete = &v2
		}
	}
	if v1, ok := d.GetOk("set_dampening_max_suppress"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_dampening_max_suppress", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetDampeningMaxSuppress = &tmp
		}
	}
	if v1, ok := d.GetOk("set_dampening_reachability_half_life"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_dampening_reachability_half_life", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetDampeningReachabilityHalfLife = &tmp
		}
	}
	if v1, ok := d.GetOk("set_dampening_reuse"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_dampening_reuse", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetDampeningReuse = &tmp
		}
	}
	if v1, ok := d.GetOk("set_dampening_suppress"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_dampening_suppress", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetDampeningSuppress = &tmp
		}
	}
	if v1, ok := d.GetOk("set_dampening_unreachability_half_life"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_dampening_unreachability_half_life", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetDampeningUnreachabilityHalfLife = &tmp
		}
	}
	if v, ok := d.GetOk("set_extcommunity_rt"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("set_extcommunity_rt", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterrouteMapRuleSetExtcommunityRt(d, v, "set_extcommunity_rt", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SetExtcommunityRt = t
		}
	} else if d.HasChange("set_extcommunity_rt") {
		old, new := d.GetChange("set_extcommunity_rt")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SetExtcommunityRt = &[]models.RouterrouteMapRuleSetExtcommunityRt{}
		}
	}
	if v, ok := d.GetOk("set_extcommunity_soo"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("set_extcommunity_soo", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterrouteMapRuleSetExtcommunitySoo(d, v, "set_extcommunity_soo", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SetExtcommunitySoo = t
		}
	} else if d.HasChange("set_extcommunity_soo") {
		old, new := d.GetChange("set_extcommunity_soo")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SetExtcommunitySoo = &[]models.RouterrouteMapRuleSetExtcommunitySoo{}
		}
	}
	if v1, ok := d.GetOk("set_flags"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("set_flags", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetFlags = &tmp
		}
	}
	if v1, ok := d.GetOk("set_ip_nexthop"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_ip_nexthop", sv)
				diags = append(diags, e)
			}
			obj.SetIpNexthop = &v2
		}
	}
	if v1, ok := d.GetOk("set_ip6_nexthop"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_ip6_nexthop", sv)
				diags = append(diags, e)
			}
			obj.SetIp6Nexthop = &v2
		}
	}
	if v1, ok := d.GetOk("set_ip6_nexthop_local"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_ip6_nexthop_local", sv)
				diags = append(diags, e)
			}
			obj.SetIp6NexthopLocal = &v2
		}
	}
	if v1, ok := d.GetOk("set_local_preference"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_local_preference", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetLocalPreference = &tmp
		}
	}
	if v1, ok := d.GetOk("set_metric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_metric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetMetric = &tmp
		}
	}
	if v1, ok := d.GetOk("set_metric_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_metric_type", sv)
				diags = append(diags, e)
			}
			obj.SetMetricType = &v2
		}
	}
	if v1, ok := d.GetOk("set_origin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_origin", sv)
				diags = append(diags, e)
			}
			obj.SetOrigin = &v2
		}
	}
	if v1, ok := d.GetOk("set_originator_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_originator_id", sv)
				diags = append(diags, e)
			}
			obj.SetOriginatorId = &v2
		}
	}
	if v1, ok := d.GetOk("set_route_tag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_route_tag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetRouteTag = &tmp
		}
	}
	if v1, ok := d.GetOk("set_tag"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_tag", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetTag = &tmp
		}
	}
	if v1, ok := d.GetOk("set_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("set_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SetWeight = &tmp
		}
	}
	return &obj, diags
}
