// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Global FortiAnalyzer settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogFortianalyzerSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Global FortiAnalyzer settings.",

		ReadContext: dataSourceLogFortianalyzerSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"access_config": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiAnalyzer access to configuration and data.",
				Computed:    true,
			},
			"alt_server": {
				Type:        schema.TypeString,
				Description: "Alternate FortiAnalyzer.",
				Computed:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "Certificate used to communicate with FortiAnalyzer.",
				Computed:    true,
			},
			"certificate_verification": {
				Type:        schema.TypeString,
				Description: "Enable/disable identity verification of FortiAnalyzer by use of certificate.",
				Computed:    true,
			},
			"conn_timeout": {
				Type:        schema.TypeInt,
				Description: "FortiAnalyzer connection time-out in seconds (for status and log buffer).",
				Computed:    true,
			},
			"enc_algorithm": {
				Type:        schema.TypeString,
				Description: "Configure the level of SSL protection for secure communication with FortiAnalyzer.",
				Computed:    true,
			},
			"fallback_to_primary": {
				Type:        schema.TypeString,
				Description: "Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available.",
				Computed:    true,
			},
			"hmac_algorithm": {
				Type:        schema.TypeString,
				Description: "OFTP login hash algorithm.",
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
			"ips_archive": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS packet archive logging.",
				Computed:    true,
			},
			"max_log_rate": {
				Type:        schema.TypeInt,
				Description: "FortiAnalyzer maximum log rate in MBps (0 = unlimited).",
				Computed:    true,
			},
			"monitor_failure_retry_period": {
				Type:        schema.TypeInt,
				Description: "Time between FortiAnalyzer connection retries in seconds (for status and log buffer).",
				Computed:    true,
			},
			"monitor_keepalive_period": {
				Type:        schema.TypeInt,
				Description: "Time between OFTP keepalives in seconds (for status and log buffer).",
				Computed:    true,
			},
			"preshared_key": {
				Type:        schema.TypeString,
				Description: "Preshared-key used for auto-authorization on FortiAnalyzer.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeString,
				Description: "Set log transmission priority.",
				Computed:    true,
			},
			"reliable": {
				Type:        schema.TypeString,
				Description: "Enable/disable reliable logging to FortiAnalyzer.",
				Computed:    true,
			},
			"serial": {
				Type:        schema.TypeList,
				Description: "Serial numbers of the FortiAnalyzer.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Serial Number.",
							Computed:    true,
						},
					},
				},
			},
			"server": {
				Type:        schema.TypeString,
				Description: "The remote FortiAnalyzer.",
				Computed:    true,
			},
			"server_cert_ca": {
				Type:        schema.TypeString,
				Description: "Mandatory CA on FortiGate in certificate chain of server.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging to FortiAnalyzer.",
				Computed:    true,
			},
			"upload_day": {
				Type:        schema.TypeString,
				Description: "Day of week (month) to upload logs.",
				Computed:    true,
			},
			"upload_interval": {
				Type:        schema.TypeString,
				Description: "Frequency to upload log files to FortiAnalyzer.",
				Computed:    true,
			},
			"upload_option": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging to hard disk and then uploading to FortiAnalyzer.",
				Computed:    true,
			},
			"upload_time": {
				Type:        schema.TypeString,
				Description: "Time to upload logs (hh:mm).",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogFortianalyzerSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogFortianalyzerSetting"

	o, err := c.Cmdb.ReadLogFortianalyzerSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogFortianalyzerSetting dataSource: %v", err)
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

	diags := refreshObjectLogFortianalyzerSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
