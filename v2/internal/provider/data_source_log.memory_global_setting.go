// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global settings for memory logging.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogMemoryGlobalSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Global settings for memory logging.",

		ReadContext: dataSourceLogMemoryGlobalSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"full_final_warning_threshold": {
				Type:        schema.TypeInt,
				Description: "Log full final warning threshold as a percent (3 - 100, default = 95).",
				Computed:    true,
			},
			"full_first_warning_threshold": {
				Type:        schema.TypeInt,
				Description: "Log full first warning threshold as a percent (1 - 98, default = 75).",
				Computed:    true,
			},
			"full_second_warning_threshold": {
				Type:        schema.TypeInt,
				Description: "Log full second warning threshold as a percent (2 - 99, default = 90).",
				Computed:    true,
			},
			"max_size": {
				Type:        schema.TypeInt,
				Description: "Maximum amount of memory that can be used for memory logging in bytes.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogMemoryGlobalSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogMemoryGlobalSetting"

	o, err := c.Cmdb.ReadLogMemoryGlobalSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogMemoryGlobalSetting dataSource: %v", err)
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

	diags := refreshObjectLogMemoryGlobalSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
