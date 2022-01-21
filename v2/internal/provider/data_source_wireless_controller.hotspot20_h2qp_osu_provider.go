// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure online sign up (OSU) provider list.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerHotspot20H2qpOsuProvider() *schema.Resource {
	return &schema.Resource{
		Description: "Configure online sign up (OSU) provider list.",

		ReadContext: dataSourceWirelessControllerHotspot20H2qpOsuProviderRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"friendly_name": {
				Type:        schema.TypeList,
				Description: "OSU provider friendly name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"friendly_name": {
							Type:        schema.TypeString,
							Description: "OSU provider friendly name.",
							Computed:    true,
						},
						"index": {
							Type:        schema.TypeInt,
							Description: "OSU provider friendly name index.",
							Computed:    true,
						},
						"lang": {
							Type:        schema.TypeString,
							Description: "Language code.",
							Computed:    true,
						},
					},
				},
			},
			"icon": {
				Type:        schema.TypeString,
				Description: "OSU provider icon.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "OSU provider ID.",
				Required:    true,
			},
			"osu_method": {
				Type:        schema.TypeString,
				Description: "OSU method list.",
				Computed:    true,
			},
			"osu_nai": {
				Type:        schema.TypeString,
				Description: "OSU NAI.",
				Computed:    true,
			},
			"server_uri": {
				Type:        schema.TypeString,
				Description: "Server URI.",
				Computed:    true,
			},
			"service_description": {
				Type:        schema.TypeList,
				Description: "OSU service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lang": {
							Type:        schema.TypeString,
							Description: "Language code.",
							Computed:    true,
						},
						"service_description": {
							Type:        schema.TypeString,
							Description: "Service description.",
							Computed:    true,
						},
						"service_id": {
							Type:        schema.TypeInt,
							Description: "OSU service ID.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessControllerHotspot20H2qpOsuProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpOsuProvider(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpOsuProvider dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpOsuProvider(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
