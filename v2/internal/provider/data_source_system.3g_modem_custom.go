// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: 3G MODEM custom.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystem3gModemCustom() *schema.Resource {
	return &schema.Resource{
		Description: "3G MODEM custom.",

		ReadContext: dataSourceSystem3gModemCustomRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"class_id": {
				Type:        schema.TypeString,
				Description: "USB interface class in hexadecimal format (00-ff).",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID.",
				Computed:    true,
			},
			"init_string": {
				Type:        schema.TypeString,
				Description: "Init string in hexadecimal format (even length).",
				Computed:    true,
			},
			"model": {
				Type:        schema.TypeString,
				Description: "MODEM model name.",
				Computed:    true,
			},
			"modeswitch_string": {
				Type:        schema.TypeString,
				Description: "USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'",
				Computed:    true,
			},
			"product_id": {
				Type:        schema.TypeString,
				Description: "USB product ID in hexadecimal format (0000-ffff).",
				Computed:    true,
			},
			"vendor": {
				Type:        schema.TypeString,
				Description: "MODEM vendor name.",
				Computed:    true,
			},
			"vendor_id": {
				Type:        schema.TypeString,
				Description: "USB vendor ID in hexadecimal format (0000-ffff).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystem3gModemCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystem3gModemCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading System3gModemCustom dataSource: %v", err)
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

	diags := refreshObjectSystem3gModemCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
