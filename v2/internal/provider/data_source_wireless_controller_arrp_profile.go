// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerArrpProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.",

		ReadContext: dataSourceWirelessControllerArrpProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"darrp_optimize": {
				Type:        schema.TypeInt,
				Description: "Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).",
				Computed:    true,
			},
			"darrp_optimize_schedules": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Schedule name.",
							Computed:    true,
						},
					},
				},
			},
			"include_dfs_channel": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).",
				Computed:    true,
			},
			"include_weather_channel": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).",
				Computed:    true,
			},
			"monitor_period": {
				Type:        schema.TypeInt,
				Description: "Period in seconds to measure average transmit retries and receive errors (default = 300).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WiFi ARRP profile name.",
				Required:    true,
			},
			"override_darrp_optimize": {
				Type:        schema.TypeString,
				Description: "Enable to override setting darrp-optimize and darrp-optimize-schedules (default = disable).",
				Computed:    true,
			},
			"selection_period": {
				Type:        schema.TypeInt,
				Description: "Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).",
				Computed:    true,
			},
			"threshold_ap": {
				Type:        schema.TypeInt,
				Description: "Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).",
				Computed:    true,
			},
			"threshold_channel_load": {
				Type:        schema.TypeInt,
				Description: "Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).",
				Computed:    true,
			},
			"threshold_noise_floor": {
				Type:        schema.TypeString,
				Description: "Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).",
				Computed:    true,
			},
			"threshold_rx_errors": {
				Type:        schema.TypeInt,
				Description: "Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).",
				Computed:    true,
			},
			"threshold_spectral_rssi": {
				Type:        schema.TypeString,
				Description: "Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).",
				Computed:    true,
			},
			"threshold_tx_retries": {
				Type:        schema.TypeInt,
				Description: "Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).",
				Computed:    true,
			},
			"weight_channel_load": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).",
				Computed:    true,
			},
			"weight_dfs_channel": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).",
				Computed:    true,
			},
			"weight_managed_ap": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).",
				Computed:    true,
			},
			"weight_noise_floor": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).",
				Computed:    true,
			},
			"weight_rogue_ap": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).",
				Computed:    true,
			},
			"weight_spectral_rssi": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).",
				Computed:    true,
			},
			"weight_weather_channel": {
				Type:        schema.TypeInt,
				Description: "Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerArrpProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerArrpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerArrpProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerArrpProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
