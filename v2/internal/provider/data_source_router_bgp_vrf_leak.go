// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: BGP VRF leaking table.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterBgpVrfLeak() *schema.Resource {
	return &schema.Resource{
		Description: "BGP VRF leaking table.",

		ReadContext: dataSourceRouterBgpVrfLeakRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"target": {
				Type:        schema.TypeList,
				Description: "Target VRF table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type:        schema.TypeString,
							Description: "Interface which is used to leak routes to target VRF.",
							Computed:    true,
						},
						"route_map": {
							Type:        schema.TypeString,
							Description: "Route map of VRF leaking.",
							Computed:    true,
						},
						"vrf": {
							Type:        schema.TypeString,
							Description: "Target VRF ID (0 - 31).",
							Computed:    true,
						},
					},
				},
			},
			"vrf": {
				Type:        schema.TypeString,
				Description: "Origin VRF ID (0 - 31).",
				Required:    true,
			},
		},
	}
}

func dataSourceRouterBgpVrfLeakRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("vrf")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadRouterBgpVrfLeak(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpVrfLeak dataSource: %v", err)
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

	diags := refreshObjectRouterBgpVrfLeak(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
