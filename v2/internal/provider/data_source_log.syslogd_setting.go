// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global settings for remote syslog server.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogSyslogdSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Global settings for remote syslog server.",

		ReadContext: dataSourceLogSyslogdSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "Certificate used to communicate with Syslog server.",
				Computed:    true,
			},
			"custom_field_name": {
				Type:        schema.TypeList,
				Description: "Custom field name for CEF format logging.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom": {
							Type:        schema.TypeString,
							Description: "Field custom name.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Field name.",
							Computed:    true,
						},
					},
				},
			},
			"enc_algorithm": {
				Type:        schema.TypeString,
				Description: "Enable/disable reliable syslogging with TLS encryption.",
				Computed:    true,
			},
			"facility": {
				Type:        schema.TypeString,
				Description: "Remote syslog facility.",
				Computed:    true,
			},
			"format": {
				Type:        schema.TypeString,
				Description: "Log format.",
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
				Description: "Syslog maximum log rate in MBps (0 = unlimited).",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Remote syslog logging over UDP/Reliable TCP.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Server listen port.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeString,
				Description: "Set log transmission priority.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Address of remote syslog server.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address of syslog.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable remote syslog logging.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogSyslogdSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogSyslogdSetting"

	o, err := c.Cmdb.ReadLogSyslogdSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogSyslogdSetting dataSource: %v", err)
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

	diags := refreshObjectLogSyslogdSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
