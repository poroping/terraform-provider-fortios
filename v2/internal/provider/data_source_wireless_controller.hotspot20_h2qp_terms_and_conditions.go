// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure terms and conditions.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerHotspot20H2qpTermsAndConditions() *schema.Resource {
	return &schema.Resource{
		Description: "Configure terms and conditions.",

		ReadContext: dataSourceWirelessControllerHotspot20H2qpTermsAndConditionsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"filename": {
				Type:        schema.TypeString,
				Description: "Filename.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Terms and Conditions ID.",
				Required:    true,
			},
			"timestamp": {
				Type:        schema.TypeInt,
				Description: "Timestamp.",
				Computed:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "URL.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerHotspot20H2qpTermsAndConditionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpTermsAndConditions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpTermsAndConditions dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpTermsAndConditions(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
