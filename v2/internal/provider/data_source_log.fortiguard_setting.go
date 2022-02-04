// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure logging to FortiCloud.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogFortiguardSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure logging to FortiCloud.",

		ReadContext: dataSourceLogFortiguardSettingRead,

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
			"conn_timeout": {
				Type:        schema.TypeInt,
				Description: "FortiGate Cloud connection timeout in seconds.",
				Computed:    true,
			},
			"enc_algorithm": {
				Type:        schema.TypeString,
				Description: "Configure the level of SSL protection for secure communication with FortiCloud.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"max_log_rate": {
				Type:        schema.TypeInt,
				Description: "FortiCloud maximum log rate in MBps (0 = unlimited).",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeString,
				Description: "Set log transmission priority.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address used to connect FortiCloud.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
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

func dataSourceLogFortiguardSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogFortiguardSetting"

	o, err := c.Cmdb.ReadLogFortiguardSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogFortiguardSetting dataSource: %v", err)
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

	diags := refreshObjectLogFortiguardSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
