// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure system-wide switch controller settings.

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

func resourceSwitchControllerSystem() *schema.Resource {
	return &schema.Resource{
		Description: "Configure system-wide switch controller settings.",

		CreateContext: resourceSwitchControllerSystemCreate,
		ReadContext:   resourceSwitchControllerSystemRead,
		UpdateContext: resourceSwitchControllerSystemUpdate,
		DeleteContext: resourceSwitchControllerSystemDelete,

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
			"data_sync_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1800),

				Description: "Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"dynamic_periodic_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),

				Description: "Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
			"iot_holdoff": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10080),

				Description: "MAC entry's creation time. Time must be greater than this value for an entry to be created (0 - 10080 mins, default = 5 mins).",
				Optional:    true,
				Computed:    true,
			},
			"iot_mac_idle": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10080),

				Description: "MAC entry's idle time. MAC entry is removed after this value (0 - 10080 mins, default = 1440 mins).",
				Optional:    true,
				Computed:    true,
			},
			"iot_scan_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 10080),

				Description: "IoT scan interval (2 - 10080 mins, default = 60 mins, 0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"iot_weight_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"nac_periodic_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),

				Description: "Periodic time interval to run NAC engine (5 - 60 sec, default = 15).",
				Optional:    true,
				Computed:    true,
			},
			"parallel_process": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),

				Description: "Maximum number of parallel processes.",
				Optional:    true,
				Computed:    true,
			},
			"parallel_process_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable parallel process override.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"compatible", "strict"}, false),

				Description: "Compatible/strict tunnel mode.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerSystem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerSystem(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerSystem")
	}

	return resourceSwitchControllerSystemRead(ctx, d, meta)
}

func resourceSwitchControllerSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerSystem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerSystem(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerSystem resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerSystem")
	}

	return resourceSwitchControllerSystemRead(ctx, d, meta)
}

func resourceSwitchControllerSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSwitchControllerSystem(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSwitchControllerSystem(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerSystem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerSystem(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerSystem resource: %v", err)
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

	diags := refreshObjectSwitchControllerSystem(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerSystem(d *schema.ResourceData, o *models.SwitchControllerSystem, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DataSyncInterval != nil {
		v := *o.DataSyncInterval

		if err = d.Set("data_sync_interval", v); err != nil {
			return diag.Errorf("error reading data_sync_interval: %v", err)
		}
	}

	if o.DynamicPeriodicInterval != nil {
		v := *o.DynamicPeriodicInterval

		if err = d.Set("dynamic_periodic_interval", v); err != nil {
			return diag.Errorf("error reading dynamic_periodic_interval: %v", err)
		}
	}

	if o.IotHoldoff != nil {
		v := *o.IotHoldoff

		if err = d.Set("iot_holdoff", v); err != nil {
			return diag.Errorf("error reading iot_holdoff: %v", err)
		}
	}

	if o.IotMacIdle != nil {
		v := *o.IotMacIdle

		if err = d.Set("iot_mac_idle", v); err != nil {
			return diag.Errorf("error reading iot_mac_idle: %v", err)
		}
	}

	if o.IotScanInterval != nil {
		v := *o.IotScanInterval

		if err = d.Set("iot_scan_interval", v); err != nil {
			return diag.Errorf("error reading iot_scan_interval: %v", err)
		}
	}

	if o.IotWeightThreshold != nil {
		v := *o.IotWeightThreshold

		if err = d.Set("iot_weight_threshold", v); err != nil {
			return diag.Errorf("error reading iot_weight_threshold: %v", err)
		}
	}

	if o.NacPeriodicInterval != nil {
		v := *o.NacPeriodicInterval

		if err = d.Set("nac_periodic_interval", v); err != nil {
			return diag.Errorf("error reading nac_periodic_interval: %v", err)
		}
	}

	if o.ParallelProcess != nil {
		v := *o.ParallelProcess

		if err = d.Set("parallel_process", v); err != nil {
			return diag.Errorf("error reading parallel_process: %v", err)
		}
	}

	if o.ParallelProcessOverride != nil {
		v := *o.ParallelProcessOverride

		if err = d.Set("parallel_process_override", v); err != nil {
			return diag.Errorf("error reading parallel_process_override: %v", err)
		}
	}

	if o.TunnelMode != nil {
		v := *o.TunnelMode

		if err = d.Set("tunnel_mode", v); err != nil {
			return diag.Errorf("error reading tunnel_mode: %v", err)
		}
	}

	return nil
}

func getObjectSwitchControllerSystem(d *schema.ResourceData, sv string) (*models.SwitchControllerSystem, diag.Diagnostics) {
	obj := models.SwitchControllerSystem{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("data_sync_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("data_sync_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DataSyncInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("dynamic_periodic_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dynamic_periodic_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DynamicPeriodicInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("iot_holdoff"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("iot_holdoff", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IotHoldoff = &tmp
		}
	}
	if v1, ok := d.GetOk("iot_mac_idle"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("iot_mac_idle", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IotMacIdle = &tmp
		}
	}
	if v1, ok := d.GetOk("iot_scan_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("iot_scan_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IotScanInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("iot_weight_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("iot_weight_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IotWeightThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("nac_periodic_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("nac_periodic_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NacPeriodicInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("parallel_process"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("parallel_process", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ParallelProcess = &tmp
		}
	}
	if v1, ok := d.GetOk("parallel_process_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("parallel_process_override", sv)
				diags = append(diags, e)
			}
			obj.ParallelProcessOverride = &v2
		}
	}
	if v1, ok := d.GetOk("tunnel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("tunnel_mode", sv)
				diags = append(diags, e)
			}
			obj.TunnelMode = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSwitchControllerSystem(d *schema.ResourceData, sv string) (*models.SwitchControllerSystem, diag.Diagnostics) {
	obj := models.SwitchControllerSystem{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
