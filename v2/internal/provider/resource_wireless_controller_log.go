// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless controller event log filters.

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

func resourceWirelessControllerLog() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless controller event log filters.",

		CreateContext: resourceWirelessControllerLogCreate,
		ReadContext:   resourceWirelessControllerLogRead,
		UpdateContext: resourceWirelessControllerLogUpdate,
		DeleteContext: resourceWirelessControllerLogDelete,

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
			"addrgrp_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log address group message.",
				Optional:    true,
				Computed:    true,
			},
			"ble_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log BLE detection message.",
				Optional:    true,
				Computed:    true,
			},
			"clb_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log client load balancing message.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_starv_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log DHCP starvation event message.",
				Optional:    true,
				Computed:    true,
			},
			"led_sched_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log LED schedule event message.",
				Optional:    true,
				Computed:    true,
			},
			"radio_event_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log radio event message.",
				Optional:    true,
				Computed:    true,
			},
			"rogue_event_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log rogue AP event message.",
				Optional:    true,
				Computed:    true,
			},
			"sta_event_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log station event message.",
				Optional:    true,
				Computed:    true,
			},
			"sta_locate_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log station locate message.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable wireless event logging.",
				Optional:    true,
				Computed:    true,
			},
			"wids_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log WIDS message.",
				Optional:    true,
				Computed:    true,
			},
			"wtp_event_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log WTP event message.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerLog(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerLog(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerLog")
	}

	return resourceWirelessControllerLogRead(ctx, d, meta)
}

func resourceWirelessControllerLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerLog(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerLog(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerLog resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerLog")
	}

	return resourceWirelessControllerLogRead(ctx, d, meta)
}

func resourceWirelessControllerLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWirelessControllerLog(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWirelessControllerLog(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerLog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerLog(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerLog resource: %v", err)
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

	diags := refreshObjectWirelessControllerLog(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerLog(d *schema.ResourceData, o *models.WirelessControllerLog, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrgrpLog != nil {
		v := *o.AddrgrpLog

		if err = d.Set("addrgrp_log", v); err != nil {
			return diag.Errorf("error reading addrgrp_log: %v", err)
		}
	}

	if o.BleLog != nil {
		v := *o.BleLog

		if err = d.Set("ble_log", v); err != nil {
			return diag.Errorf("error reading ble_log: %v", err)
		}
	}

	if o.ClbLog != nil {
		v := *o.ClbLog

		if err = d.Set("clb_log", v); err != nil {
			return diag.Errorf("error reading clb_log: %v", err)
		}
	}

	if o.DhcpStarvLog != nil {
		v := *o.DhcpStarvLog

		if err = d.Set("dhcp_starv_log", v); err != nil {
			return diag.Errorf("error reading dhcp_starv_log: %v", err)
		}
	}

	if o.LedSchedLog != nil {
		v := *o.LedSchedLog

		if err = d.Set("led_sched_log", v); err != nil {
			return diag.Errorf("error reading led_sched_log: %v", err)
		}
	}

	if o.RadioEventLog != nil {
		v := *o.RadioEventLog

		if err = d.Set("radio_event_log", v); err != nil {
			return diag.Errorf("error reading radio_event_log: %v", err)
		}
	}

	if o.RogueEventLog != nil {
		v := *o.RogueEventLog

		if err = d.Set("rogue_event_log", v); err != nil {
			return diag.Errorf("error reading rogue_event_log: %v", err)
		}
	}

	if o.StaEventLog != nil {
		v := *o.StaEventLog

		if err = d.Set("sta_event_log", v); err != nil {
			return diag.Errorf("error reading sta_event_log: %v", err)
		}
	}

	if o.StaLocateLog != nil {
		v := *o.StaLocateLog

		if err = d.Set("sta_locate_log", v); err != nil {
			return diag.Errorf("error reading sta_locate_log: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.WidsLog != nil {
		v := *o.WidsLog

		if err = d.Set("wids_log", v); err != nil {
			return diag.Errorf("error reading wids_log: %v", err)
		}
	}

	if o.WtpEventLog != nil {
		v := *o.WtpEventLog

		if err = d.Set("wtp_event_log", v); err != nil {
			return diag.Errorf("error reading wtp_event_log: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerLog(d *schema.ResourceData, sv string) (*models.WirelessControllerLog, diag.Diagnostics) {
	obj := models.WirelessControllerLog{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addrgrp_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("addrgrp_log", sv)
				diags = append(diags, e)
			}
			obj.AddrgrpLog = &v2
		}
	}
	if v1, ok := d.GetOk("ble_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ble_log", sv)
				diags = append(diags, e)
			}
			obj.BleLog = &v2
		}
	}
	if v1, ok := d.GetOk("clb_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("clb_log", sv)
				diags = append(diags, e)
			}
			obj.ClbLog = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_starv_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_starv_log", sv)
				diags = append(diags, e)
			}
			obj.DhcpStarvLog = &v2
		}
	}
	if v1, ok := d.GetOk("led_sched_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("led_sched_log", sv)
				diags = append(diags, e)
			}
			obj.LedSchedLog = &v2
		}
	}
	if v1, ok := d.GetOk("radio_event_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radio_event_log", sv)
				diags = append(diags, e)
			}
			obj.RadioEventLog = &v2
		}
	}
	if v1, ok := d.GetOk("rogue_event_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rogue_event_log", sv)
				diags = append(diags, e)
			}
			obj.RogueEventLog = &v2
		}
	}
	if v1, ok := d.GetOk("sta_event_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sta_event_log", sv)
				diags = append(diags, e)
			}
			obj.StaEventLog = &v2
		}
	}
	if v1, ok := d.GetOk("sta_locate_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sta_locate_log", sv)
				diags = append(diags, e)
			}
			obj.StaLocateLog = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("wids_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wids_log", sv)
				diags = append(diags, e)
			}
			obj.WidsLog = &v2
		}
	}
	if v1, ok := d.GetOk("wtp_event_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wtp_event_log", sv)
				diags = append(diags, e)
			}
			obj.WtpEventLog = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWirelessControllerLog(d *schema.ResourceData, sv string) (*models.WirelessControllerLog, diag.Diagnostics) {
	obj := models.WirelessControllerLog{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
