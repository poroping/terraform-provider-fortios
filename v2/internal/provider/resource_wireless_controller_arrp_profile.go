// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.

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

func resourceWirelessControllerArrpProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.",

		CreateContext: resourceWirelessControllerArrpProfileCreate,
		ReadContext:   resourceWirelessControllerArrpProfileRead,
		UpdateContext: resourceWirelessControllerArrpProfileUpdate,
		DeleteContext: resourceWirelessControllerArrpProfileDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"include_dfs_channel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of DFS channel in DARRP channel selection phase 1 (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"include_weather_channel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of weather channel in DARRP channel selection phase 1 (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"monitor_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Period in seconds to measure average transmit retries and receive errors (default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WiFi ARRP profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"selection_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Period in seconds to measure average channel load, noise floor, spectral RSSI (default = 3600).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_ap": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 500),

				Description: "Threshold to reject channel in DARRP channel selection phase 1 due to surrounding APs (0 - 500, default = 250).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_channel_load": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Threshold in percentage to reject channel in DARRP channel selection phase 1 due to channel load (0 - 100, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_noise_floor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Threshold in dBm to reject channel in DARRP channel selection phase 1 due to noise floor (-95 to -20, default = -85).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_rx_errors": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),

				Description: "Threshold in percentage for receive errors to trigger channel reselection in DARRP monitor stage (0 - 100, default = 50).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_spectral_rssi": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),

				Description: "Threshold in dBm to reject channel in DARRP channel selection phase 1 due to spectral RSSI (-95 to -20, default = -65).",
				Optional:    true,
				Computed:    true,
			},
			"threshold_tx_retries": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000),

				Description: "Threshold in percentage for transmit retries to trigger channel reselection in DARRP monitor stage (0 - 1000, default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"weight_channel_load": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for channel load (0 - 2000, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"weight_dfs_channel": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for DFS channel (0 - 2000, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"weight_managed_ap": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for managed APs (0 - 2000, default = 50).",
				Optional:    true,
				Computed:    true,
			},
			"weight_noise_floor": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for noise floor (0 - 2000, default = 40).",
				Optional:    true,
				Computed:    true,
			},
			"weight_rogue_ap": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for rogue APs (0 - 2000, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"weight_spectral_rssi": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for spectral RSSI (0 - 2000, default = 40).",
				Optional:    true,
				Computed:    true,
			},
			"weight_weather_channel": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),

				Description: "Weight in DARRP channel score calculation for weather channel (0 - 2000, default = 1000).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerArrpProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerArrpProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerArrpProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerArrpProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerArrpProfile")
	}

	return resourceWirelessControllerArrpProfileRead(ctx, d, meta)
}

func resourceWirelessControllerArrpProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerArrpProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerArrpProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerArrpProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerArrpProfile")
	}

	return resourceWirelessControllerArrpProfileRead(ctx, d, meta)
}

func resourceWirelessControllerArrpProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerArrpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerArrpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerArrpProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerArrpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerArrpProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerArrpProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerArrpProfile(d *schema.ResourceData, o *models.WirelessControllerArrpProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.IncludeDfsChannel != nil {
		v := *o.IncludeDfsChannel

		if err = d.Set("include_dfs_channel", v); err != nil {
			return diag.Errorf("error reading include_dfs_channel: %v", err)
		}
	}

	if o.IncludeWeatherChannel != nil {
		v := *o.IncludeWeatherChannel

		if err = d.Set("include_weather_channel", v); err != nil {
			return diag.Errorf("error reading include_weather_channel: %v", err)
		}
	}

	if o.MonitorPeriod != nil {
		v := *o.MonitorPeriod

		if err = d.Set("monitor_period", v); err != nil {
			return diag.Errorf("error reading monitor_period: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.SelectionPeriod != nil {
		v := *o.SelectionPeriod

		if err = d.Set("selection_period", v); err != nil {
			return diag.Errorf("error reading selection_period: %v", err)
		}
	}

	if o.ThresholdAp != nil {
		v := *o.ThresholdAp

		if err = d.Set("threshold_ap", v); err != nil {
			return diag.Errorf("error reading threshold_ap: %v", err)
		}
	}

	if o.ThresholdChannelLoad != nil {
		v := *o.ThresholdChannelLoad

		if err = d.Set("threshold_channel_load", v); err != nil {
			return diag.Errorf("error reading threshold_channel_load: %v", err)
		}
	}

	if o.ThresholdNoiseFloor != nil {
		v := *o.ThresholdNoiseFloor

		if err = d.Set("threshold_noise_floor", v); err != nil {
			return diag.Errorf("error reading threshold_noise_floor: %v", err)
		}
	}

	if o.ThresholdRxErrors != nil {
		v := *o.ThresholdRxErrors

		if err = d.Set("threshold_rx_errors", v); err != nil {
			return diag.Errorf("error reading threshold_rx_errors: %v", err)
		}
	}

	if o.ThresholdSpectralRssi != nil {
		v := *o.ThresholdSpectralRssi

		if err = d.Set("threshold_spectral_rssi", v); err != nil {
			return diag.Errorf("error reading threshold_spectral_rssi: %v", err)
		}
	}

	if o.ThresholdTxRetries != nil {
		v := *o.ThresholdTxRetries

		if err = d.Set("threshold_tx_retries", v); err != nil {
			return diag.Errorf("error reading threshold_tx_retries: %v", err)
		}
	}

	if o.WeightChannelLoad != nil {
		v := *o.WeightChannelLoad

		if err = d.Set("weight_channel_load", v); err != nil {
			return diag.Errorf("error reading weight_channel_load: %v", err)
		}
	}

	if o.WeightDfsChannel != nil {
		v := *o.WeightDfsChannel

		if err = d.Set("weight_dfs_channel", v); err != nil {
			return diag.Errorf("error reading weight_dfs_channel: %v", err)
		}
	}

	if o.WeightManagedAp != nil {
		v := *o.WeightManagedAp

		if err = d.Set("weight_managed_ap", v); err != nil {
			return diag.Errorf("error reading weight_managed_ap: %v", err)
		}
	}

	if o.WeightNoiseFloor != nil {
		v := *o.WeightNoiseFloor

		if err = d.Set("weight_noise_floor", v); err != nil {
			return diag.Errorf("error reading weight_noise_floor: %v", err)
		}
	}

	if o.WeightRogueAp != nil {
		v := *o.WeightRogueAp

		if err = d.Set("weight_rogue_ap", v); err != nil {
			return diag.Errorf("error reading weight_rogue_ap: %v", err)
		}
	}

	if o.WeightSpectralRssi != nil {
		v := *o.WeightSpectralRssi

		if err = d.Set("weight_spectral_rssi", v); err != nil {
			return diag.Errorf("error reading weight_spectral_rssi: %v", err)
		}
	}

	if o.WeightWeatherChannel != nil {
		v := *o.WeightWeatherChannel

		if err = d.Set("weight_weather_channel", v); err != nil {
			return diag.Errorf("error reading weight_weather_channel: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerArrpProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerArrpProfile, diag.Diagnostics) {
	obj := models.WirelessControllerArrpProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("include_dfs_channel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("include_dfs_channel", sv)
				diags = append(diags, e)
			}
			obj.IncludeDfsChannel = &v2
		}
	}
	if v1, ok := d.GetOk("include_weather_channel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("include_weather_channel", sv)
				diags = append(diags, e)
			}
			obj.IncludeWeatherChannel = &v2
		}
	}
	if v1, ok := d.GetOk("monitor_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("monitor_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MonitorPeriod = &tmp
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
	if v1, ok := d.GetOk("selection_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("selection_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SelectionPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_ap"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_ap", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdAp = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_channel_load"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_channel_load", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdChannelLoad = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_noise_floor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_noise_floor", sv)
				diags = append(diags, e)
			}
			obj.ThresholdNoiseFloor = &v2
		}
	}
	if v1, ok := d.GetOk("threshold_rx_errors"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_rx_errors", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdRxErrors = &tmp
		}
	}
	if v1, ok := d.GetOk("threshold_spectral_rssi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_spectral_rssi", sv)
				diags = append(diags, e)
			}
			obj.ThresholdSpectralRssi = &v2
		}
	}
	if v1, ok := d.GetOk("threshold_tx_retries"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("threshold_tx_retries", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ThresholdTxRetries = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_channel_load"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_channel_load", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightChannelLoad = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_dfs_channel"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_dfs_channel", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightDfsChannel = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_managed_ap"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_managed_ap", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightManagedAp = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_noise_floor"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_noise_floor", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightNoiseFloor = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_rogue_ap"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_rogue_ap", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightRogueAp = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_spectral_rssi"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_spectral_rssi", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightSpectralRssi = &tmp
		}
	}
	if v1, ok := d.GetOk("weight_weather_channel"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight_weather_channel", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WeightWeatherChannel = &tmp
		}
	}
	return &obj, diags
}
