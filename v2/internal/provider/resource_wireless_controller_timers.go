// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure CAPWAP timers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWirelessControllerTimers() *schema.Resource {
	return &schema.Resource{
		Description: "Configure CAPWAP timers.",

		CreateContext: resourceWirelessControllerTimersCreate,
		ReadContext:   resourceWirelessControllerTimersRead,
		UpdateContext: resourceWirelessControllerTimersUpdate,
		DeleteContext: resourceWirelessControllerTimersDelete,

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
			"ble_scan_report_intv": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),

				Description: "Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"client_idle_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 3600),

				Description: "Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).",
				Optional:    true,
				Computed:    true,
			},
			"discovery_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 180),

				Description: "Time between discovery requests (2 - 180 sec, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"drma_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"echo_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"fake_ap_log": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_intf_cleanup": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 3600),

				Description: "Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).",
				Optional:    true,
				Computed:    true,
			},
			"radio_stats_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time between running radio reports (1 - 255 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
			"rogue_ap_log": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1440),

				Description: "Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"sta_capability_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time between running station capability reports (1 - 255 sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"sta_locate_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).",
				Optional:    true,
				Computed:    true,
			},
			"sta_stats_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time between running client (station) reports (1 - 255 sec, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"vap_stats_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerTimersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerTimers(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerTimers(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerTimers")
	}

	return resourceWirelessControllerTimersRead(ctx, d, meta)
}

func resourceWirelessControllerTimersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerTimers(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerTimers(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerTimers resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerTimers")
	}

	return resourceWirelessControllerTimersRead(ctx, d, meta)
}

func resourceWirelessControllerTimersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerTimers(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerTimers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerTimersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerTimers(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerTimers resource: %v", err)
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

	diags := refreshObjectWirelessControllerTimers(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerTimers(d *schema.ResourceData, o *models.WirelessControllerTimers, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BleScanReportIntv != nil {
		v := *o.BleScanReportIntv

		if err = d.Set("ble_scan_report_intv", v); err != nil {
			return diag.Errorf("error reading ble_scan_report_intv: %v", err)
		}
	}

	if o.ClientIdleTimeout != nil {
		v := *o.ClientIdleTimeout

		if err = d.Set("client_idle_timeout", v); err != nil {
			return diag.Errorf("error reading client_idle_timeout: %v", err)
		}
	}

	if o.DiscoveryInterval != nil {
		v := *o.DiscoveryInterval

		if err = d.Set("discovery_interval", v); err != nil {
			return diag.Errorf("error reading discovery_interval: %v", err)
		}
	}

	if o.DrmaInterval != nil {
		v := *o.DrmaInterval

		if err = d.Set("drma_interval", v); err != nil {
			return diag.Errorf("error reading drma_interval: %v", err)
		}
	}

	if o.EchoInterval != nil {
		v := *o.EchoInterval

		if err = d.Set("echo_interval", v); err != nil {
			return diag.Errorf("error reading echo_interval: %v", err)
		}
	}

	if o.FakeApLog != nil {
		v := *o.FakeApLog

		if err = d.Set("fake_ap_log", v); err != nil {
			return diag.Errorf("error reading fake_ap_log: %v", err)
		}
	}

	if o.IpsecIntfCleanup != nil {
		v := *o.IpsecIntfCleanup

		if err = d.Set("ipsec_intf_cleanup", v); err != nil {
			return diag.Errorf("error reading ipsec_intf_cleanup: %v", err)
		}
	}

	if o.RadioStatsInterval != nil {
		v := *o.RadioStatsInterval

		if err = d.Set("radio_stats_interval", v); err != nil {
			return diag.Errorf("error reading radio_stats_interval: %v", err)
		}
	}

	if o.RogueApLog != nil {
		v := *o.RogueApLog

		if err = d.Set("rogue_ap_log", v); err != nil {
			return diag.Errorf("error reading rogue_ap_log: %v", err)
		}
	}

	if o.StaCapabilityInterval != nil {
		v := *o.StaCapabilityInterval

		if err = d.Set("sta_capability_interval", v); err != nil {
			return diag.Errorf("error reading sta_capability_interval: %v", err)
		}
	}

	if o.StaLocateTimer != nil {
		v := *o.StaLocateTimer

		if err = d.Set("sta_locate_timer", v); err != nil {
			return diag.Errorf("error reading sta_locate_timer: %v", err)
		}
	}

	if o.StaStatsInterval != nil {
		v := *o.StaStatsInterval

		if err = d.Set("sta_stats_interval", v); err != nil {
			return diag.Errorf("error reading sta_stats_interval: %v", err)
		}
	}

	if o.VapStatsInterval != nil {
		v := *o.VapStatsInterval

		if err = d.Set("vap_stats_interval", v); err != nil {
			return diag.Errorf("error reading vap_stats_interval: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerTimers(d *schema.ResourceData, sv string) (*models.WirelessControllerTimers, diag.Diagnostics) {
	obj := models.WirelessControllerTimers{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ble_scan_report_intv"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ble_scan_report_intv", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BleScanReportIntv = &tmp
		}
	}
	if v1, ok := d.GetOk("client_idle_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_idle_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ClientIdleTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("discovery_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("discovery_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DiscoveryInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("drma_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("drma_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DrmaInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("echo_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("echo_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EchoInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("fake_ap_log"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fake_ap_log", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.FakeApLog = &tmp
		}
	}
	if v1, ok := d.GetOk("ipsec_intf_cleanup"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_intf_cleanup", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IpsecIntfCleanup = &tmp
		}
	}
	if v1, ok := d.GetOk("radio_stats_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radio_stats_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RadioStatsInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("rogue_ap_log"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rogue_ap_log", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RogueApLog = &tmp
		}
	}
	if v1, ok := d.GetOk("sta_capability_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sta_capability_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StaCapabilityInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("sta_locate_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sta_locate_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StaLocateTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("sta_stats_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sta_stats_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StaStatsInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("vap_stats_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vap_stats_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VapStatsInterval = &tmp
		}
	}
	return &obj, diags
}
