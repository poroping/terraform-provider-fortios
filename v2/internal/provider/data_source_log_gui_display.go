// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure how log messages are displayed on the GUI.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogGuiDisplay() *schema.Resource {
	return &schema.Resource{
		Description: "Configure how log messages are displayed on the GUI.",

		ReadContext: dataSourceLogGuiDisplayRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fortiview_unscanned_apps": {
				Type:        schema.TypeString,
				Description: "Enable/disable showing unscanned traffic in FortiView application charts.",
				Computed:    true,
			},
			"resolve_apps": {
				Type:        schema.TypeString,
				Description: "Resolve unknown applications on the GUI using Fortinet's remote application database.",
				Computed:    true,
			},
			"resolve_hosts": {
				Type:        schema.TypeString,
				Description: "Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogGuiDisplayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogGuiDisplay"

	o, err := c.Cmdb.ReadLogGuiDisplay(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogGuiDisplay dataSource: %v", err)
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

	diags := refreshObjectLogGuiDisplay(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
