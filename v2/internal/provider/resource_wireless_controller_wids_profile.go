// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless intrusion detection system (WIDS) profiles.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWirelessControllerWidsProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless intrusion detection system (WIDS) profiles.",

		CreateContext: resourceWirelessControllerWidsProfileCreate,
		ReadContext:   resourceWirelessControllerWidsProfileRead,
		UpdateContext: resourceWirelessControllerWidsProfileUpdate,
		DeleteContext: resourceWirelessControllerWidsProfileDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"ap_auto_suppress": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable on-wire rogue AP auto-suppression (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ap_bgscan_disable_schedules": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Schedule name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ap_bgscan_duration": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000),

				Description: "Listen time on scanning a channel (10 - 1000 msec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"ap_bgscan_idle": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000),

				Description: "Wait time for channel inactivity before scanning this channel (0 - 1000 msec, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"ap_bgscan_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 600),

				Description: "Period between successive channel scans (1 - 600 sec, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"ap_bgscan_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),

				Description: "Period between background scans (10 - 3600 sec, default = 600).",
				Optional:    true,
				Computed:    true,
			},
			"ap_bgscan_report_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 600),

				Description: "Period between background scan reports (15 - 600 sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"ap_fgscan_report_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 600),

				Description: "Period between foreground scan reports (15 - 600 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
			"ap_scan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable rogue AP detection.",
				Optional:    true,
				Computed:    true,
			},
			"ap_scan_passive": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable passive scanning. Enable means do not send probe request on any channels (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ap_scan_threshold": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Minimum signal level/threshold in dBm required for the AP to report detected rogue AP (-95 to -20, default = -90).",
				Optional:    true,
				Computed:    true,
			},
			"asleap_attack": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable asleap attack detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"assoc_flood_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "The threshold value for association frame flooding.",
				Optional:    true,
				Computed:    true,
			},
			"assoc_flood_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 120),

				Description: "Number of seconds after which a station is considered not connected.",
				Optional:    true,
				Computed:    true,
			},
			"assoc_frame_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable association frame flooding detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"auth_flood_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),

				Description: "The threshold value for authentication frame flooding.",
				Optional:    true,
				Computed:    true,
			},
			"auth_flood_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 120),

				Description: "Number of seconds after which a station is considered not connected.",
				Optional:    true,
				Computed:    true,
			},
			"auth_frame_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication frame flooding detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"deauth_broadcast": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable broadcasting de-authentication detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"deauth_unknown_src_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Threshold value per second to deauth unknown src for DoS attack (0: no limit).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_fail_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EAPOL-Failure flooding (to AP) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_fail_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "The detection interval for EAPOL-Failure flooding (1 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_fail_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "The threshold value for EAPOL-Failure flooding in specified interval.",
				Optional:    true,
				Computed:    true,
			},
			"eapol_logoff_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EAPOL-Logoff flooding (to AP) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_logoff_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "The detection interval for EAPOL-Logoff flooding (1 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_logoff_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "The threshold value for EAPOL-Logoff flooding in specified interval.",
				Optional:    true,
				Computed:    true,
			},
			"eapol_pre_fail_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable premature EAPOL-Failure flooding (to STA) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_pre_fail_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "The detection interval for premature EAPOL-Failure flooding (1 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_pre_fail_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "The threshold value for premature EAPOL-Failure flooding in specified interval.",
				Optional:    true,
				Computed:    true,
			},
			"eapol_pre_succ_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable premature EAPOL-Success flooding (to STA) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_pre_succ_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "The detection interval for premature EAPOL-Success flooding (1 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_pre_succ_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "The threshold value for premature EAPOL-Success flooding in specified interval.",
				Optional:    true,
				Computed:    true,
			},
			"eapol_start_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EAPOL-Start flooding (to AP) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_start_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "The detection interval for EAPOL-Start flooding (1 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_start_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "The threshold value for EAPOL-Start flooding in specified interval.",
				Optional:    true,
				Computed:    true,
			},
			"eapol_succ_flood": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable EAPOL-Success flooding (to AP) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_succ_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "The detection interval for EAPOL-Success flooding (1 - 3600 sec).",
				Optional:    true,
				Computed:    true,
			},
			"eapol_succ_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),

				Description: "The threshold value for EAPOL-Success flooding in specified interval.",
				Optional:    true,
				Computed:    true,
			},
			"invalid_mac_oui": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable invalid MAC OUI detection.",
				Optional:    true,
				Computed:    true,
			},
			"long_duration_attack": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable long duration attack detection based on user configured threshold (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"long_duration_thresh": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1000, 32767),

				Description: "Threshold value for long duration attack detection (1000 - 32767 usec, default = 8200).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WIDS profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"null_ssid_probe_resp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable null SSID probe response detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"sensor_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "foreign", "both"}, false),

				Description: "Scan nearby WiFi stations (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"spoofed_deauth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable spoofed de-authentication attack detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"weak_wep_iv": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable weak WEP IV (Initialization Vector) detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"wireless_bridge": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wireless bridge detection (default = disable).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerWidsProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WirelessControllerWidsProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerWidsProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerWidsProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWidsProfile")
	}

	return resourceWirelessControllerWidsProfileRead(ctx, d, meta)
}

func resourceWirelessControllerWidsProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerWidsProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerWidsProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerWidsProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWidsProfile")
	}

	return resourceWirelessControllerWidsProfileRead(ctx, d, meta)
}

func resourceWirelessControllerWidsProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerWidsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerWidsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWidsProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadWirelessControllerWidsProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWidsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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

func flattenWirelessControllerWidsProfileApBgscanDisableSchedules(d *schema.ResourceData, v *[]models.WirelessControllerWidsProfileApBgscanDisableSchedules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectWirelessControllerWidsProfile(d *schema.ResourceData, o *models.WirelessControllerWidsProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ApAutoSuppress != nil {
		v := *o.ApAutoSuppress

		if err = d.Set("ap_auto_suppress", v); err != nil {
			return diag.Errorf("error reading ap_auto_suppress: %v", err)
		}
	}

	if o.ApBgscanDisableSchedules != nil {
		if err = d.Set("ap_bgscan_disable_schedules", flattenWirelessControllerWidsProfileApBgscanDisableSchedules(d, o.ApBgscanDisableSchedules, "ap_bgscan_disable_schedules", sort)); err != nil {
			return diag.Errorf("error reading ap_bgscan_disable_schedules: %v", err)
		}
	}

	if o.ApBgscanDuration != nil {
		v := *o.ApBgscanDuration

		if err = d.Set("ap_bgscan_duration", v); err != nil {
			return diag.Errorf("error reading ap_bgscan_duration: %v", err)
		}
	}

	if o.ApBgscanIdle != nil {
		v := *o.ApBgscanIdle

		if err = d.Set("ap_bgscan_idle", v); err != nil {
			return diag.Errorf("error reading ap_bgscan_idle: %v", err)
		}
	}

	if o.ApBgscanIntv != nil {
		v := *o.ApBgscanIntv

		if err = d.Set("ap_bgscan_intv", v); err != nil {
			return diag.Errorf("error reading ap_bgscan_intv: %v", err)
		}
	}

	if o.ApBgscanPeriod != nil {
		v := *o.ApBgscanPeriod

		if err = d.Set("ap_bgscan_period", v); err != nil {
			return diag.Errorf("error reading ap_bgscan_period: %v", err)
		}
	}

	if o.ApBgscanReportIntv != nil {
		v := *o.ApBgscanReportIntv

		if err = d.Set("ap_bgscan_report_intv", v); err != nil {
			return diag.Errorf("error reading ap_bgscan_report_intv: %v", err)
		}
	}

	if o.ApFgscanReportIntv != nil {
		v := *o.ApFgscanReportIntv

		if err = d.Set("ap_fgscan_report_intv", v); err != nil {
			return diag.Errorf("error reading ap_fgscan_report_intv: %v", err)
		}
	}

	if o.ApScan != nil {
		v := *o.ApScan

		if err = d.Set("ap_scan", v); err != nil {
			return diag.Errorf("error reading ap_scan: %v", err)
		}
	}

	if o.ApScanPassive != nil {
		v := *o.ApScanPassive

		if err = d.Set("ap_scan_passive", v); err != nil {
			return diag.Errorf("error reading ap_scan_passive: %v", err)
		}
	}

	if o.ApScanThreshold != nil {
		v := *o.ApScanThreshold

		if err = d.Set("ap_scan_threshold", v); err != nil {
			return diag.Errorf("error reading ap_scan_threshold: %v", err)
		}
	}

	if o.AsleapAttack != nil {
		v := *o.AsleapAttack

		if err = d.Set("asleap_attack", v); err != nil {
			return diag.Errorf("error reading asleap_attack: %v", err)
		}
	}

	if o.AssocFloodThresh != nil {
		v := *o.AssocFloodThresh

		if err = d.Set("assoc_flood_thresh", v); err != nil {
			return diag.Errorf("error reading assoc_flood_thresh: %v", err)
		}
	}

	if o.AssocFloodTime != nil {
		v := *o.AssocFloodTime

		if err = d.Set("assoc_flood_time", v); err != nil {
			return diag.Errorf("error reading assoc_flood_time: %v", err)
		}
	}

	if o.AssocFrameFlood != nil {
		v := *o.AssocFrameFlood

		if err = d.Set("assoc_frame_flood", v); err != nil {
			return diag.Errorf("error reading assoc_frame_flood: %v", err)
		}
	}

	if o.AuthFloodThresh != nil {
		v := *o.AuthFloodThresh

		if err = d.Set("auth_flood_thresh", v); err != nil {
			return diag.Errorf("error reading auth_flood_thresh: %v", err)
		}
	}

	if o.AuthFloodTime != nil {
		v := *o.AuthFloodTime

		if err = d.Set("auth_flood_time", v); err != nil {
			return diag.Errorf("error reading auth_flood_time: %v", err)
		}
	}

	if o.AuthFrameFlood != nil {
		v := *o.AuthFrameFlood

		if err = d.Set("auth_frame_flood", v); err != nil {
			return diag.Errorf("error reading auth_frame_flood: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DeauthBroadcast != nil {
		v := *o.DeauthBroadcast

		if err = d.Set("deauth_broadcast", v); err != nil {
			return diag.Errorf("error reading deauth_broadcast: %v", err)
		}
	}

	if o.DeauthUnknownSrcThresh != nil {
		v := *o.DeauthUnknownSrcThresh

		if err = d.Set("deauth_unknown_src_thresh", v); err != nil {
			return diag.Errorf("error reading deauth_unknown_src_thresh: %v", err)
		}
	}

	if o.EapolFailFlood != nil {
		v := *o.EapolFailFlood

		if err = d.Set("eapol_fail_flood", v); err != nil {
			return diag.Errorf("error reading eapol_fail_flood: %v", err)
		}
	}

	if o.EapolFailIntv != nil {
		v := *o.EapolFailIntv

		if err = d.Set("eapol_fail_intv", v); err != nil {
			return diag.Errorf("error reading eapol_fail_intv: %v", err)
		}
	}

	if o.EapolFailThresh != nil {
		v := *o.EapolFailThresh

		if err = d.Set("eapol_fail_thresh", v); err != nil {
			return diag.Errorf("error reading eapol_fail_thresh: %v", err)
		}
	}

	if o.EapolLogoffFlood != nil {
		v := *o.EapolLogoffFlood

		if err = d.Set("eapol_logoff_flood", v); err != nil {
			return diag.Errorf("error reading eapol_logoff_flood: %v", err)
		}
	}

	if o.EapolLogoffIntv != nil {
		v := *o.EapolLogoffIntv

		if err = d.Set("eapol_logoff_intv", v); err != nil {
			return diag.Errorf("error reading eapol_logoff_intv: %v", err)
		}
	}

	if o.EapolLogoffThresh != nil {
		v := *o.EapolLogoffThresh

		if err = d.Set("eapol_logoff_thresh", v); err != nil {
			return diag.Errorf("error reading eapol_logoff_thresh: %v", err)
		}
	}

	if o.EapolPreFailFlood != nil {
		v := *o.EapolPreFailFlood

		if err = d.Set("eapol_pre_fail_flood", v); err != nil {
			return diag.Errorf("error reading eapol_pre_fail_flood: %v", err)
		}
	}

	if o.EapolPreFailIntv != nil {
		v := *o.EapolPreFailIntv

		if err = d.Set("eapol_pre_fail_intv", v); err != nil {
			return diag.Errorf("error reading eapol_pre_fail_intv: %v", err)
		}
	}

	if o.EapolPreFailThresh != nil {
		v := *o.EapolPreFailThresh

		if err = d.Set("eapol_pre_fail_thresh", v); err != nil {
			return diag.Errorf("error reading eapol_pre_fail_thresh: %v", err)
		}
	}

	if o.EapolPreSuccFlood != nil {
		v := *o.EapolPreSuccFlood

		if err = d.Set("eapol_pre_succ_flood", v); err != nil {
			return diag.Errorf("error reading eapol_pre_succ_flood: %v", err)
		}
	}

	if o.EapolPreSuccIntv != nil {
		v := *o.EapolPreSuccIntv

		if err = d.Set("eapol_pre_succ_intv", v); err != nil {
			return diag.Errorf("error reading eapol_pre_succ_intv: %v", err)
		}
	}

	if o.EapolPreSuccThresh != nil {
		v := *o.EapolPreSuccThresh

		if err = d.Set("eapol_pre_succ_thresh", v); err != nil {
			return diag.Errorf("error reading eapol_pre_succ_thresh: %v", err)
		}
	}

	if o.EapolStartFlood != nil {
		v := *o.EapolStartFlood

		if err = d.Set("eapol_start_flood", v); err != nil {
			return diag.Errorf("error reading eapol_start_flood: %v", err)
		}
	}

	if o.EapolStartIntv != nil {
		v := *o.EapolStartIntv

		if err = d.Set("eapol_start_intv", v); err != nil {
			return diag.Errorf("error reading eapol_start_intv: %v", err)
		}
	}

	if o.EapolStartThresh != nil {
		v := *o.EapolStartThresh

		if err = d.Set("eapol_start_thresh", v); err != nil {
			return diag.Errorf("error reading eapol_start_thresh: %v", err)
		}
	}

	if o.EapolSuccFlood != nil {
		v := *o.EapolSuccFlood

		if err = d.Set("eapol_succ_flood", v); err != nil {
			return diag.Errorf("error reading eapol_succ_flood: %v", err)
		}
	}

	if o.EapolSuccIntv != nil {
		v := *o.EapolSuccIntv

		if err = d.Set("eapol_succ_intv", v); err != nil {
			return diag.Errorf("error reading eapol_succ_intv: %v", err)
		}
	}

	if o.EapolSuccThresh != nil {
		v := *o.EapolSuccThresh

		if err = d.Set("eapol_succ_thresh", v); err != nil {
			return diag.Errorf("error reading eapol_succ_thresh: %v", err)
		}
	}

	if o.InvalidMacOui != nil {
		v := *o.InvalidMacOui

		if err = d.Set("invalid_mac_oui", v); err != nil {
			return diag.Errorf("error reading invalid_mac_oui: %v", err)
		}
	}

	if o.LongDurationAttack != nil {
		v := *o.LongDurationAttack

		if err = d.Set("long_duration_attack", v); err != nil {
			return diag.Errorf("error reading long_duration_attack: %v", err)
		}
	}

	if o.LongDurationThresh != nil {
		v := *o.LongDurationThresh

		if err = d.Set("long_duration_thresh", v); err != nil {
			return diag.Errorf("error reading long_duration_thresh: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NullSsidProbeResp != nil {
		v := *o.NullSsidProbeResp

		if err = d.Set("null_ssid_probe_resp", v); err != nil {
			return diag.Errorf("error reading null_ssid_probe_resp: %v", err)
		}
	}

	if o.SensorMode != nil {
		v := *o.SensorMode

		if err = d.Set("sensor_mode", v); err != nil {
			return diag.Errorf("error reading sensor_mode: %v", err)
		}
	}

	if o.SpoofedDeauth != nil {
		v := *o.SpoofedDeauth

		if err = d.Set("spoofed_deauth", v); err != nil {
			return diag.Errorf("error reading spoofed_deauth: %v", err)
		}
	}

	if o.WeakWepIv != nil {
		v := *o.WeakWepIv

		if err = d.Set("weak_wep_iv", v); err != nil {
			return diag.Errorf("error reading weak_wep_iv: %v", err)
		}
	}

	if o.WirelessBridge != nil {
		v := *o.WirelessBridge

		if err = d.Set("wireless_bridge", v); err != nil {
			return diag.Errorf("error reading wireless_bridge: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerWidsProfileApBgscanDisableSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerWidsProfileApBgscanDisableSchedules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerWidsProfileApBgscanDisableSchedules

	for i := range l {
		tmp := models.WirelessControllerWidsProfileApBgscanDisableSchedules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerWidsProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerWidsProfile, diag.Diagnostics) {
	obj := models.WirelessControllerWidsProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ap_auto_suppress"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_auto_suppress", sv)
				diags = append(diags, e)
			}
			obj.ApAutoSuppress = &v2
		}
	}
	if v, ok := d.GetOk("ap_bgscan_disable_schedules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ap_bgscan_disable_schedules", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerWidsProfileApBgscanDisableSchedules(d, v, "ap_bgscan_disable_schedules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ApBgscanDisableSchedules = t
		}
	} else if d.HasChange("ap_bgscan_disable_schedules") {
		old, new := d.GetChange("ap_bgscan_disable_schedules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ApBgscanDisableSchedules = &[]models.WirelessControllerWidsProfileApBgscanDisableSchedules{}
		}
	}
	if v1, ok := d.GetOk("ap_bgscan_duration"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_bgscan_duration", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApBgscanDuration = &tmp
		}
	}
	if v1, ok := d.GetOk("ap_bgscan_idle"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_bgscan_idle", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApBgscanIdle = &tmp
		}
	}
	if v1, ok := d.GetOk("ap_bgscan_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_bgscan_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApBgscanIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("ap_bgscan_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_bgscan_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApBgscanPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("ap_bgscan_report_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_bgscan_report_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApBgscanReportIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("ap_fgscan_report_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_fgscan_report_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ApFgscanReportIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("ap_scan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_scan", sv)
				diags = append(diags, e)
			}
			obj.ApScan = &v2
		}
	}
	if v1, ok := d.GetOk("ap_scan_passive"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_scan_passive", sv)
				diags = append(diags, e)
			}
			obj.ApScanPassive = &v2
		}
	}
	if v1, ok := d.GetOk("ap_scan_threshold"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ap_scan_threshold", sv)
				diags = append(diags, e)
			}
			obj.ApScanThreshold = &v2
		}
	}
	if v1, ok := d.GetOk("asleap_attack"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("asleap_attack", sv)
				diags = append(diags, e)
			}
			obj.AsleapAttack = &v2
		}
	}
	if v1, ok := d.GetOk("assoc_flood_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assoc_flood_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AssocFloodThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("assoc_flood_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assoc_flood_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AssocFloodTime = &tmp
		}
	}
	if v1, ok := d.GetOk("assoc_frame_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("assoc_frame_flood", sv)
				diags = append(diags, e)
			}
			obj.AssocFrameFlood = &v2
		}
	}
	if v1, ok := d.GetOk("auth_flood_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_flood_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthFloodThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_flood_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_flood_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthFloodTime = &tmp
		}
	}
	if v1, ok := d.GetOk("auth_frame_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_frame_flood", sv)
				diags = append(diags, e)
			}
			obj.AuthFrameFlood = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("deauth_broadcast"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deauth_broadcast", sv)
				diags = append(diags, e)
			}
			obj.DeauthBroadcast = &v2
		}
	}
	if v1, ok := d.GetOk("deauth_unknown_src_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deauth_unknown_src_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeauthUnknownSrcThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_fail_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_fail_flood", sv)
				diags = append(diags, e)
			}
			obj.EapolFailFlood = &v2
		}
	}
	if v1, ok := d.GetOk("eapol_fail_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_fail_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolFailIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_fail_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_fail_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolFailThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_logoff_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_logoff_flood", sv)
				diags = append(diags, e)
			}
			obj.EapolLogoffFlood = &v2
		}
	}
	if v1, ok := d.GetOk("eapol_logoff_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_logoff_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolLogoffIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_logoff_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_logoff_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolLogoffThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_pre_fail_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_pre_fail_flood", sv)
				diags = append(diags, e)
			}
			obj.EapolPreFailFlood = &v2
		}
	}
	if v1, ok := d.GetOk("eapol_pre_fail_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_pre_fail_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolPreFailIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_pre_fail_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_pre_fail_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolPreFailThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_pre_succ_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_pre_succ_flood", sv)
				diags = append(diags, e)
			}
			obj.EapolPreSuccFlood = &v2
		}
	}
	if v1, ok := d.GetOk("eapol_pre_succ_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_pre_succ_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolPreSuccIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_pre_succ_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_pre_succ_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolPreSuccThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_start_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_start_flood", sv)
				diags = append(diags, e)
			}
			obj.EapolStartFlood = &v2
		}
	}
	if v1, ok := d.GetOk("eapol_start_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_start_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolStartIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_start_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_start_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolStartThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_succ_flood"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_succ_flood", sv)
				diags = append(diags, e)
			}
			obj.EapolSuccFlood = &v2
		}
	}
	if v1, ok := d.GetOk("eapol_succ_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_succ_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolSuccIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("eapol_succ_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eapol_succ_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EapolSuccThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("invalid_mac_oui"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("invalid_mac_oui", sv)
				diags = append(diags, e)
			}
			obj.InvalidMacOui = &v2
		}
	}
	if v1, ok := d.GetOk("long_duration_attack"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("long_duration_attack", sv)
				diags = append(diags, e)
			}
			obj.LongDurationAttack = &v2
		}
	}
	if v1, ok := d.GetOk("long_duration_thresh"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("long_duration_thresh", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LongDurationThresh = &tmp
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("null_ssid_probe_resp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("null_ssid_probe_resp", sv)
				diags = append(diags, e)
			}
			obj.NullSsidProbeResp = &v2
		}
	}
	if v1, ok := d.GetOk("sensor_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sensor_mode", sv)
				diags = append(diags, e)
			}
			obj.SensorMode = &v2
		}
	}
	if v1, ok := d.GetOk("spoofed_deauth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("spoofed_deauth", sv)
				diags = append(diags, e)
			}
			obj.SpoofedDeauth = &v2
		}
	}
	if v1, ok := d.GetOk("weak_wep_iv"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weak_wep_iv", sv)
				diags = append(diags, e)
			}
			obj.WeakWepIv = &v2
		}
	}
	if v1, ok := d.GetOk("wireless_bridge"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wireless_bridge", sv)
				diags = append(diags, e)
			}
			obj.WirelessBridge = &v2
		}
	}
	return &obj, diags
}
