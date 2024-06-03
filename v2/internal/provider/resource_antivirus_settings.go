// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiVirus settings.

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

func resourceAntivirusSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiVirus settings.",

		CreateContext: resourceAntivirusSettingsCreate,
		ReadContext:   resourceAntivirusSettingsRead,
		UpdateContext: resourceAntivirusSettingsUpdate,
		DeleteContext: resourceAntivirusSettingsDelete,

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
			"cache_clean_result": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable cache of clean scan results (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"cache_infected_result": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable cache of infected scan results (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"default_db": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "extended", "extreme"}, false),

				Description: "Select the AV database to be used for AV scanning.",
				Optional:    true,
				Computed:    true,
			},
			"grayware": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable grayware detection when an AntiVirus profile is applied to traffic.",
				Optional:    true,
				Computed:    true,
			},
			"machine_learning_detection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "monitor", "disable"}, false),

				Description: "Use machine learning based malware detection.",
				Optional:    true,
				Computed:    true,
			},
			"override_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 3600),

				Description: "Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.",
				Optional:    true,
				Computed:    true,
			},
			"use_extreme_db": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the use of Extreme AVDB.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceAntivirusSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAntivirusSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAntivirusSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AntivirusSettings")
	}

	return resourceAntivirusSettingsRead(ctx, d, meta)
}

func resourceAntivirusSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAntivirusSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAntivirusSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AntivirusSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AntivirusSettings")
	}

	return resourceAntivirusSettingsRead(ctx, d, meta)
}

func resourceAntivirusSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectAntivirusSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateAntivirusSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AntivirusSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAntivirusSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AntivirusSettings resource: %v", err)
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

	diags := refreshObjectAntivirusSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectAntivirusSettings(d *schema.ResourceData, o *models.AntivirusSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CacheCleanResult != nil {
		v := *o.CacheCleanResult

		if err = d.Set("cache_clean_result", v); err != nil {
			return diag.Errorf("error reading cache_clean_result: %v", err)
		}
	}

	if o.CacheInfectedResult != nil {
		v := *o.CacheInfectedResult

		if err = d.Set("cache_infected_result", v); err != nil {
			return diag.Errorf("error reading cache_infected_result: %v", err)
		}
	}

	if o.DefaultDb != nil {
		v := *o.DefaultDb

		if err = d.Set("default_db", v); err != nil {
			return diag.Errorf("error reading default_db: %v", err)
		}
	}

	if o.Grayware != nil {
		v := *o.Grayware

		if err = d.Set("grayware", v); err != nil {
			return diag.Errorf("error reading grayware: %v", err)
		}
	}

	if o.MachineLearningDetection != nil {
		v := *o.MachineLearningDetection

		if err = d.Set("machine_learning_detection", v); err != nil {
			return diag.Errorf("error reading machine_learning_detection: %v", err)
		}
	}

	if o.OverrideTimeout != nil {
		v := *o.OverrideTimeout

		if err = d.Set("override_timeout", v); err != nil {
			return diag.Errorf("error reading override_timeout: %v", err)
		}
	}

	if o.UseExtremeDb != nil {
		v := *o.UseExtremeDb

		if err = d.Set("use_extreme_db", v); err != nil {
			return diag.Errorf("error reading use_extreme_db: %v", err)
		}
	}

	return nil
}

func getObjectAntivirusSettings(d *schema.ResourceData, sv string) (*models.AntivirusSettings, diag.Diagnostics) {
	obj := models.AntivirusSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cache_clean_result"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "v7.2.1") {
				e := utils.AttributeVersionWarning("cache_clean_result", sv)
				diags = append(diags, e)
			}
			obj.CacheCleanResult = &v2
		}
	}
	if v1, ok := d.GetOk("cache_infected_result"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("cache_infected_result", sv)
				diags = append(diags, e)
			}
			obj.CacheInfectedResult = &v2
		}
	}
	if v1, ok := d.GetOk("default_db"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("default_db", sv)
				diags = append(diags, e)
			}
			obj.DefaultDb = &v2
		}
	}
	if v1, ok := d.GetOk("grayware"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("grayware", sv)
				diags = append(diags, e)
			}
			obj.Grayware = &v2
		}
	}
	if v1, ok := d.GetOk("machine_learning_detection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("machine_learning_detection", sv)
				diags = append(diags, e)
			}
			obj.MachineLearningDetection = &v2
		}
	}
	if v1, ok := d.GetOk("override_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("override_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OverrideTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("use_extreme_db"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("use_extreme_db", sv)
				diags = append(diags, e)
			}
			obj.UseExtremeDb = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectAntivirusSettings(d *schema.ResourceData, sv string) (*models.AntivirusSettings, diag.Diagnostics) {
	obj := models.AntivirusSettings{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
