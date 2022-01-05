// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless intrusion detection system (WIDS) profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerWidsProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless intrusion detection system (WIDS) profiles.",

		ReadContext: dataSourceWirelessControllerWidsProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ap_auto_suppress": {
				Type:        schema.TypeString,
				Description: "Enable/disable on-wire rogue AP auto-suppression (default = disable).",
				Computed:    true,
			},
			"ap_bgscan_disable_schedules": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space.",
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
			"ap_bgscan_duration": {
				Type:        schema.TypeInt,
				Description: "Listen time on scanning a channel (10 - 1000 msec, default = 30).",
				Computed:    true,
			},
			"ap_bgscan_idle": {
				Type:        schema.TypeInt,
				Description: "Wait time for channel inactivity before scanning this channel (0 - 1000 msec, default = 20).",
				Computed:    true,
			},
			"ap_bgscan_intv": {
				Type:        schema.TypeInt,
				Description: "Period between successive channel scans (1 - 600 sec, default = 3).",
				Computed:    true,
			},
			"ap_bgscan_period": {
				Type:        schema.TypeInt,
				Description: "Period between background scans (10 - 3600 sec, default = 600).",
				Computed:    true,
			},
			"ap_bgscan_report_intv": {
				Type:        schema.TypeInt,
				Description: "Period between background scan reports (15 - 600 sec, default = 30).",
				Computed:    true,
			},
			"ap_fgscan_report_intv": {
				Type:        schema.TypeInt,
				Description: "Period between foreground scan reports (15 - 600 sec, default = 15).",
				Computed:    true,
			},
			"ap_scan": {
				Type:        schema.TypeString,
				Description: "Enable/disable rogue AP detection.",
				Computed:    true,
			},
			"ap_scan_passive": {
				Type:        schema.TypeString,
				Description: "Enable/disable passive scanning. Enable means do not send probe request on any channels (default = disable).",
				Computed:    true,
			},
			"ap_scan_threshold": {
				Type:        schema.TypeString,
				Description: "Minimum signal level/threshold in dBm required for the AP to report detected rogue AP (-95 to -20, default = -90).",
				Computed:    true,
			},
			"asleap_attack": {
				Type:        schema.TypeString,
				Description: "Enable/disable asleap attack detection (default = disable).",
				Computed:    true,
			},
			"assoc_flood_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for association frame flooding.",
				Computed:    true,
			},
			"assoc_flood_time": {
				Type:        schema.TypeInt,
				Description: "Number of seconds after which a station is considered not connected.",
				Computed:    true,
			},
			"assoc_frame_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable association frame flooding detection (default = disable).",
				Computed:    true,
			},
			"auth_flood_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for authentication frame flooding.",
				Computed:    true,
			},
			"auth_flood_time": {
				Type:        schema.TypeInt,
				Description: "Number of seconds after which a station is considered not connected.",
				Computed:    true,
			},
			"auth_frame_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication frame flooding detection (default = disable).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"deauth_broadcast": {
				Type:        schema.TypeString,
				Description: "Enable/disable broadcasting de-authentication detection (default = disable).",
				Computed:    true,
			},
			"deauth_unknown_src_thresh": {
				Type:        schema.TypeInt,
				Description: "Threshold value per second to deauth unknown src for DoS attack (0: no limit).",
				Computed:    true,
			},
			"eapol_fail_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable EAPOL-Failure flooding (to AP) detection (default = disable).",
				Computed:    true,
			},
			"eapol_fail_intv": {
				Type:        schema.TypeInt,
				Description: "The detection interval for EAPOL-Failure flooding (1 - 3600 sec).",
				Computed:    true,
			},
			"eapol_fail_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for EAPOL-Failure flooding in specified interval.",
				Computed:    true,
			},
			"eapol_logoff_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable EAPOL-Logoff flooding (to AP) detection (default = disable).",
				Computed:    true,
			},
			"eapol_logoff_intv": {
				Type:        schema.TypeInt,
				Description: "The detection interval for EAPOL-Logoff flooding (1 - 3600 sec).",
				Computed:    true,
			},
			"eapol_logoff_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for EAPOL-Logoff flooding in specified interval.",
				Computed:    true,
			},
			"eapol_pre_fail_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable premature EAPOL-Failure flooding (to STA) detection (default = disable).",
				Computed:    true,
			},
			"eapol_pre_fail_intv": {
				Type:        schema.TypeInt,
				Description: "The detection interval for premature EAPOL-Failure flooding (1 - 3600 sec).",
				Computed:    true,
			},
			"eapol_pre_fail_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for premature EAPOL-Failure flooding in specified interval.",
				Computed:    true,
			},
			"eapol_pre_succ_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable premature EAPOL-Success flooding (to STA) detection (default = disable).",
				Computed:    true,
			},
			"eapol_pre_succ_intv": {
				Type:        schema.TypeInt,
				Description: "The detection interval for premature EAPOL-Success flooding (1 - 3600 sec).",
				Computed:    true,
			},
			"eapol_pre_succ_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for premature EAPOL-Success flooding in specified interval.",
				Computed:    true,
			},
			"eapol_start_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable EAPOL-Start flooding (to AP) detection (default = disable).",
				Computed:    true,
			},
			"eapol_start_intv": {
				Type:        schema.TypeInt,
				Description: "The detection interval for EAPOL-Start flooding (1 - 3600 sec).",
				Computed:    true,
			},
			"eapol_start_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for EAPOL-Start flooding in specified interval.",
				Computed:    true,
			},
			"eapol_succ_flood": {
				Type:        schema.TypeString,
				Description: "Enable/disable EAPOL-Success flooding (to AP) detection (default = disable).",
				Computed:    true,
			},
			"eapol_succ_intv": {
				Type:        schema.TypeInt,
				Description: "The detection interval for EAPOL-Success flooding (1 - 3600 sec).",
				Computed:    true,
			},
			"eapol_succ_thresh": {
				Type:        schema.TypeInt,
				Description: "The threshold value for EAPOL-Success flooding in specified interval.",
				Computed:    true,
			},
			"invalid_mac_oui": {
				Type:        schema.TypeString,
				Description: "Enable/disable invalid MAC OUI detection.",
				Computed:    true,
			},
			"long_duration_attack": {
				Type:        schema.TypeString,
				Description: "Enable/disable long duration attack detection based on user configured threshold (default = disable).",
				Computed:    true,
			},
			"long_duration_thresh": {
				Type:        schema.TypeInt,
				Description: "Threshold value for long duration attack detection (1000 - 32767 usec, default = 8200).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WIDS profile name.",
				Required:    true,
			},
			"null_ssid_probe_resp": {
				Type:        schema.TypeString,
				Description: "Enable/disable null SSID probe response detection (default = disable).",
				Computed:    true,
			},
			"sensor_mode": {
				Type:        schema.TypeString,
				Description: "Scan nearby WiFi stations (default = disable).",
				Computed:    true,
			},
			"spoofed_deauth": {
				Type:        schema.TypeString,
				Description: "Enable/disable spoofed de-authentication attack detection (default = disable).",
				Computed:    true,
			},
			"weak_wep_iv": {
				Type:        schema.TypeString,
				Description: "Enable/disable weak WEP IV (Initialization Vector) detection (default = disable).",
				Computed:    true,
			},
			"wireless_bridge": {
				Type:        schema.TypeString,
				Description: "Enable/disable wireless bridge detection (default = disable).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerWidsProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerWidsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWidsProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerWidsProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
