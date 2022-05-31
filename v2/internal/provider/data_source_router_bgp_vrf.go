// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.2.0 schemas
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

func dataSourceRouterBgpVrf() *schema.Resource {
	return &schema.Resource{
		Description: "BGP VRF leaking table.",

		ReadContext: dataSourceRouterBgpVrfRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"export_rt": {
				Type:        schema.TypeList,
				Description: "List of export route target.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_target": {
							Type:        schema.TypeString,
							Description: "Attribute: AA|AA:NN.",
							Computed:    true,
						},
					},
				},
			},
			"import_route_map": {
				Type:        schema.TypeString,
				Description: "Import route map.",
				Computed:    true,
			},
			"import_rt": {
				Type:        schema.TypeList,
				Description: "List of import route target.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_target": {
							Type:        schema.TypeString,
							Description: "Attribute: AA|AA:NN.",
							Computed:    true,
						},
					},
				},
			},
			"leak_target": {
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
							Description: "Target VRF ID (0 - 63).",
							Computed:    true,
						},
					},
				},
			},
			"rd": {
				Type:        schema.TypeString,
				Description: "Route Distinguisher: AA|AA:NN.",
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "VRF role.",
				Computed:    true,
			},
			"vrf": {
				Type:        schema.TypeString,
				Description: "Origin VRF ID (0 - 63).",
				Required:    true,
			},
		},
	}
}

func dataSourceRouterBgpVrfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterBgpVrf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterBgpVrf dataSource: %v", err)
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

	diags := refreshObjectRouterBgpVrf(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
