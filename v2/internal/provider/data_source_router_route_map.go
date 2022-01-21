// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure route maps.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterRouteMap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure route maps.",

		ReadContext: dataSourceRouterRouteMapRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "Rule.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Rule ID.",
							Computed:    true,
						},
						"match_as_path": {
							Type:        schema.TypeString,
							Description: "Match BGP AS path list.",
							Computed:    true,
						},
						"match_community": {
							Type:        schema.TypeString,
							Description: "Match BGP community list.",
							Computed:    true,
						},
						"match_community_exact": {
							Type:        schema.TypeString,
							Description: "Enable/disable exact matching of communities.",
							Computed:    true,
						},
						"match_flags": {
							Type:        schema.TypeInt,
							Description: "BGP flag value to match (0 - 65535)",
							Computed:    true,
						},
						"match_interface": {
							Type:        schema.TypeString,
							Description: "Match interface configuration.",
							Computed:    true,
						},
						"match_ip_address": {
							Type:        schema.TypeString,
							Description: "Match IP address permitted by access-list or prefix-list.",
							Computed:    true,
						},
						"match_ip_nexthop": {
							Type:        schema.TypeString,
							Description: "Match next hop IP address passed by access-list or prefix-list.",
							Computed:    true,
						},
						"match_ip6_address": {
							Type:        schema.TypeString,
							Description: "Match IPv6 address permitted by access-list6 or prefix-list6.",
							Computed:    true,
						},
						"match_ip6_nexthop": {
							Type:        schema.TypeString,
							Description: "Match next hop IPv6 address passed by access-list6 or prefix-list6.",
							Computed:    true,
						},
						"match_metric": {
							Type:        schema.TypeInt,
							Description: "Match metric for redistribute routes.",
							Computed:    true,
						},
						"match_origin": {
							Type:        schema.TypeString,
							Description: "Match BGP origin code.",
							Computed:    true,
						},
						"match_route_type": {
							Type:        schema.TypeString,
							Description: "Match route type.",
							Computed:    true,
						},
						"match_tag": {
							Type:        schema.TypeInt,
							Description: "Match tag.",
							Computed:    true,
						},
						"match_vrf": {
							Type:        schema.TypeInt,
							Description: "Match VRF ID.",
							Computed:    true,
						},
						"set_aggregator_as": {
							Type:        schema.TypeInt,
							Description: "BGP aggregator AS.",
							Computed:    true,
						},
						"set_aggregator_ip": {
							Type:        schema.TypeString,
							Description: "BGP aggregator IP.",
							Computed:    true,
						},
						"set_aspath": {
							Type:        schema.TypeList,
							Description: "Prepend BGP AS path attribute.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"as": {
										Type:        schema.TypeString,
										Description: "AS number (0 - 4294967295). NOTE: Use quotes for repeating numbers, e.g.: \"1 1 2\"\n",
										Computed:    true,
									},
								},
							},
						},
						"set_aspath_action": {
							Type:        schema.TypeString,
							Description: "Specify preferred action of set-aspath.",
							Computed:    true,
						},
						"set_atomic_aggregate": {
							Type:        schema.TypeString,
							Description: "Enable/disable BGP atomic aggregate attribute.",
							Computed:    true,
						},
						"set_community": {
							Type:        schema.TypeList,
							Description: "BGP community attribute.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": {
										Type:        schema.TypeString,
										Description: "Attribute: AA|AA:NN|internet|local-AS|no-advertise|no-export.",
										Computed:    true,
									},
								},
							},
						},
						"set_community_additive": {
							Type:        schema.TypeString,
							Description: "Enable/disable adding set-community to existing community.",
							Computed:    true,
						},
						"set_community_delete": {
							Type:        schema.TypeString,
							Description: "Delete communities matching community list.",
							Computed:    true,
						},
						"set_dampening_max_suppress": {
							Type:        schema.TypeInt,
							Description: "Maximum duration to suppress a route (1 - 255 min, 0 = unset).",
							Computed:    true,
						},
						"set_dampening_reachability_half_life": {
							Type:        schema.TypeInt,
							Description: "Reachability half-life time for the penalty (1 - 45 min, 0 = unset).",
							Computed:    true,
						},
						"set_dampening_reuse": {
							Type:        schema.TypeInt,
							Description: "Value to start reusing a route (1 - 20000, 0 = unset).",
							Computed:    true,
						},
						"set_dampening_suppress": {
							Type:        schema.TypeInt,
							Description: "Value to start suppressing a route (1 - 20000, 0 = unset).",
							Computed:    true,
						},
						"set_dampening_unreachability_half_life": {
							Type:        schema.TypeInt,
							Description: "Unreachability Half-life time for the penalty (1 - 45 min, 0 = unset)",
							Computed:    true,
						},
						"set_extcommunity_rt": {
							Type:        schema.TypeList,
							Description: "Route Target extended community.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": {
										Type:        schema.TypeString,
										Description: "AA:NN.",
										Computed:    true,
									},
								},
							},
						},
						"set_extcommunity_soo": {
							Type:        schema.TypeList,
							Description: "Site-of-Origin extended community.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"community": {
										Type:        schema.TypeString,
										Description: "AA:NN",
										Computed:    true,
									},
								},
							},
						},
						"set_flags": {
							Type:        schema.TypeInt,
							Description: "BGP flags value (0 - 65535)",
							Computed:    true,
						},
						"set_ip_nexthop": {
							Type:        schema.TypeString,
							Description: "IP address of next hop.",
							Computed:    true,
						},
						"set_ip6_nexthop": {
							Type:        schema.TypeString,
							Description: "IPv6 global address of next hop.",
							Computed:    true,
						},
						"set_ip6_nexthop_local": {
							Type:        schema.TypeString,
							Description: "IPv6 local address of next hop.",
							Computed:    true,
						},
						"set_local_preference": {
							Type:        schema.TypeInt,
							Description: "BGP local preference path attribute.",
							Computed:    true,
						},
						"set_metric": {
							Type:        schema.TypeInt,
							Description: "Metric value.",
							Computed:    true,
						},
						"set_metric_type": {
							Type:        schema.TypeString,
							Description: "Metric type.",
							Computed:    true,
						},
						"set_origin": {
							Type:        schema.TypeString,
							Description: "BGP origin code.",
							Computed:    true,
						},
						"set_originator_id": {
							Type:        schema.TypeString,
							Description: "BGP originator ID attribute.",
							Computed:    true,
						},
						"set_route_tag": {
							Type:        schema.TypeInt,
							Description: "Route tag for routing table.",
							Computed:    true,
						},
						"set_tag": {
							Type:        schema.TypeInt,
							Description: "Tag value.",
							Computed:    true,
						},
						"set_weight": {
							Type:        schema.TypeInt,
							Description: "BGP weight for routing table.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterRouteMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadRouterRouteMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRouteMap dataSource: %v", err)
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

	diags := refreshObjectRouterRouteMap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
