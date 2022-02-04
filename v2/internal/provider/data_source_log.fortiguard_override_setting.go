// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Override global FortiCloud logging settings for this VDOM.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogFortiguardOverrideSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Override global FortiCloud logging settings for this VDOM.",

		ReadContext: dataSourceLogFortiguardOverrideSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"access_config": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiCloud access to configuration and data.",
				Computed:    true,
			},
			"max_log_rate": {
				Type:        schema.TypeInt,
				Description: "FortiCloud maximum log rate in MBps (0 = unlimited).",
				Computed:    true,
			},
			"override": {
				Type:        schema.TypeString,
				Description: "Overriding FortiCloud settings for this VDOM or use global settings.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeString,
				Description: "Set log transmission priority.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging to FortiCloud.",
				Computed:    true,
			},
			"upload_day": {
				Type:        schema.TypeString,
				Description: "Day of week to roll logs.",
				Computed:    true,
			},
			"upload_interval": {
				Type:        schema.TypeString,
				Description: "Frequency of uploading log files to FortiCloud.",
				Computed:    true,
			},
			"upload_option": {
				Type:        schema.TypeString,
				Description: "Configure how log messages are sent to FortiCloud.",
				Computed:    true,
			},
			"upload_time": {
				Type:        schema.TypeString,
				Description: "Time of day to roll logs (hh:mm).",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogFortiguardOverrideSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogFortiguardOverrideSetting"

	o, err := c.Cmdb.ReadLogFortiguardOverrideSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogFortiguardOverrideSetting dataSource: %v", err)
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

	diags := refreshObjectLogFortiguardOverrideSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
