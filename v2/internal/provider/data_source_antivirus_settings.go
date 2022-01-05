// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiVirus settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceAntivirusSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiVirus settings.",

		ReadContext: dataSourceAntivirusSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cache_clean_result": {
				Type:        schema.TypeString,
				Description: "Enable/disable cache of clean scan results (default = enable).",
				Computed:    true,
			},
			"cache_infected_result": {
				Type:        schema.TypeString,
				Description: "Enable/disable cache of infected scan results (default = enable).",
				Computed:    true,
			},
			"default_db": {
				Type:        schema.TypeString,
				Description: "Select the AV database to be used for AV scanning.",
				Computed:    true,
			},
			"grayware": {
				Type:        schema.TypeString,
				Description: "Enable/disable grayware detection when an AntiVirus profile is applied to traffic.",
				Computed:    true,
			},
			"machine_learning_detection": {
				Type:        schema.TypeString,
				Description: "Use machine learning based malware detection.",
				Computed:    true,
			},
			"override_timeout": {
				Type:        schema.TypeInt,
				Description: "Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.",
				Computed:    true,
			},
			"use_extreme_db": {
				Type:        schema.TypeString,
				Description: "Enable/disable the use of Extreme AVDB.",
				Computed:    true,
			},
		},
	}
}

func dataSourceAntivirusSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAntivirusSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AntivirusSettings dataSource: %v", err)
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

	diags := refreshObjectAntivirusSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
