// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure RIPng.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterRipng() *schema.Resource {
	return &schema.Resource{
		Description: "Configure RIPng.",

		ReadContext: dataSourceRouterRipngRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"aggregate_address": {
				Type:        schema.TypeList,
				Description: "Aggregate address.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Aggregate address entry ID.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "Aggregate address prefix.",
							Computed:    true,
						},
					},
				},
			},
			"default_information_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable generation of default route.",
				Computed:    true,
			},
			"default_metric": {
				Type:        schema.TypeInt,
				Description: "Default metric.",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeList,
				Description: "distance",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list6": {
							Type:        schema.TypeString,
							Description: "Access list for route destination.",
							Computed:    true,
						},
						"distance": {
							Type:        schema.TypeInt,
							Description: "Distance (1 - 255).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Distance ID.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "Distance prefix6.",
							Computed:    true,
						},
					},
				},
			},
			"distribute_list": {
				Type:        schema.TypeList,
				Description: "Distribute list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:        schema.TypeString,
							Description: "Distribute list direction.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Distribute list ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Distribute list interface name.",
							Computed:    true,
						},
						"listname": {
							Type:        schema.TypeString,
							Description: "Distribute access/prefix list name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "status",
							Computed:    true,
						},
					},
				},
			},
			"garbage_timer": {
				Type:        schema.TypeInt,
				Description: "Garbage timer.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeList,
				Description: "RIPng interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flags": {
							Type:        schema.TypeInt,
							Description: "Flags.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"split_horizon": {
							Type:        schema.TypeString,
							Description: "Enable/disable split horizon.",
							Computed:    true,
						},
						"split_horizon_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable split horizon.",
							Computed:    true,
						},
					},
				},
			},
			"max_out_metric": {
				Type:        schema.TypeInt,
				Description: "Maximum metric allowed to output(0 means 'not set').",
				Computed:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "neighbor",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Neighbor entry ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"ip6": {
							Type:        schema.TypeString,
							Description: "IPv6 link-local address.",
							Computed:    true,
						},
					},
				},
			},
			"network": {
				Type:        schema.TypeList,
				Description: "Network.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Network entry ID.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Network IPv6 link-local prefix.",
							Computed:    true,
						},
					},
				},
			},
			"offset_list": {
				Type:        schema.TypeList,
				Description: "Offset list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_list6": {
							Type:        schema.TypeString,
							Description: "IPv6 access list name.",
							Computed:    true,
						},
						"direction": {
							Type:        schema.TypeString,
							Description: "Offset list direction.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Offset-list ID.",
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
						"offset": {
							Type:        schema.TypeInt,
							Description: "offset",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "status",
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
							Description: "status",
							Computed:    true,
						},
					},
				},
			},
			"timeout_timer": {
				Type:        schema.TypeInt,
				Description: "Timeout timer.",
				Computed:    true,
			},
			"update_timer": {
				Type:        schema.TypeInt,
				Description: "Update timer.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterRipngRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterRipng"

	o, err := c.Cmdb.ReadRouterRipng(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRipng dataSource: %v", err)
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

	diags := refreshObjectRouterRipng(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
