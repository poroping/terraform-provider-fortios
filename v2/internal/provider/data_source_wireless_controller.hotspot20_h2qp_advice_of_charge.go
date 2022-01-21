// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure advice of charge.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerHotspot20H2qpAdviceOfCharge() *schema.Resource {
	return &schema.Resource{
		Description: "Configure advice of charge.",

		ReadContext: dataSourceWirelessControllerHotspot20H2qpAdviceOfChargeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"aoc_list": {
				Type:        schema.TypeList,
				Description: "AOC list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nai_realm": {
							Type:        schema.TypeString,
							Description: "NAI realm list name.",
							Computed:    true,
						},
						"nai_realm_encoding": {
							Type:        schema.TypeString,
							Description: "NAI realm encoding.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Advice of charge ID.",
							Computed:    true,
						},
						"plan_info": {
							Type:        schema.TypeList,
							Description: "Plan info.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currency": {
										Type:        schema.TypeString,
										Description: "Currency code.",
										Computed:    true,
									},
									"info_file": {
										Type:        schema.TypeString,
										Description: "Info file.",
										Computed:    true,
									},
									"lang": {
										Type:        schema.TypeString,
										Description: "Languague code.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Plan name.",
										Computed:    true,
									},
								},
							},
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Usage charge type.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Plan name.",
				Required:    true,
			},
		},
	}
}

func dataSourceWirelessControllerHotspot20H2qpAdviceOfChargeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpAdviceOfCharge(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpAdviceOfCharge dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpAdviceOfCharge(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
