// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure quarantine options.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceAntivirusQuarantine() *schema.Resource {
	return &schema.Resource{
		Description: "Configure quarantine options.",

		CreateContext: resourceAntivirusQuarantineCreate,
		ReadContext:   resourceAntivirusQuarantineRead,
		UpdateContext: resourceAntivirusQuarantineUpdate,
		DeleteContext: resourceAntivirusQuarantineDelete,

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
			"agelimit": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 479),

				Description: "Age limit for quarantined files (0 - 479 hours, 0 means forever).",
				Optional:    true,
				Computed:    true,
			},
			"destination": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"NULL", "disk", "FortiAnalyzer"}, false),

				Description: "Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them.",
				Optional:    true,
				Computed:    true,
			},
			"drop_blocked": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Optional:         true,
				Computed:         true,
			},
			"drop_heuristic": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Optional:         true,
				Computed:         true,
			},
			"drop_infected": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Optional:         true,
				Computed:         true,
			},
			"drop_machine_learning": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Do not quarantine files detected by machine learning found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.",
				Optional:         true,
				Computed:         true,
			},
			"lowspace": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"drop-new", "ovrw-old"}, false),

				Description: "Select the method for handling additional files when running low on disk space.",
				Optional:    true,
				Computed:    true,
			},
			"maxfilesize": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 500),

				Description: "Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).",
				Optional:    true,
				Computed:    true,
			},
			"quarantine_quota": {
				Type: schema.TypeInt,

				Description: "The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).",
				Optional:    true,
				Computed:    true,
			},
			"store_blocked": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Quarantine blocked files found in sessions using the selected protocols.",
				Optional:         true,
				Computed:         true,
			},
			"store_heuristic": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Quarantine files detected by heuristics found in sessions using the selected protocols.",
				Optional:         true,
				Computed:         true,
			},
			"store_infected": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Quarantine infected files found in sessions using the selected protocols.",
				Optional:         true,
				Computed:         true,
			},
			"store_machine_learning": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Quarantine files detected by machine learning found in sessions using the selected protocols.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceAntivirusQuarantineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAntivirusQuarantine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateAntivirusQuarantine(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AntivirusQuarantine")
	}

	return resourceAntivirusQuarantineRead(ctx, d, meta)
}

func resourceAntivirusQuarantineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectAntivirusQuarantine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateAntivirusQuarantine(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating AntivirusQuarantine resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("AntivirusQuarantine")
	}

	return resourceAntivirusQuarantineRead(ctx, d, meta)
}

func resourceAntivirusQuarantineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectAntivirusQuarantine(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateAntivirusQuarantine(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting AntivirusQuarantine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusQuarantineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadAntivirusQuarantine(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading AntivirusQuarantine resource: %v", err)
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

	diags := refreshObjectAntivirusQuarantine(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectAntivirusQuarantine(d *schema.ResourceData, o *models.AntivirusQuarantine, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Agelimit != nil {
		v := *o.Agelimit

		if err = d.Set("agelimit", v); err != nil {
			return diag.Errorf("error reading agelimit: %v", err)
		}
	}

	if o.Destination != nil {
		v := *o.Destination

		if err = d.Set("destination", v); err != nil {
			return diag.Errorf("error reading destination: %v", err)
		}
	}

	if o.DropBlocked != nil {
		v := *o.DropBlocked

		if err = d.Set("drop_blocked", v); err != nil {
			return diag.Errorf("error reading drop_blocked: %v", err)
		}
	}

	if o.DropHeuristic != nil {
		v := *o.DropHeuristic

		if err = d.Set("drop_heuristic", v); err != nil {
			return diag.Errorf("error reading drop_heuristic: %v", err)
		}
	}

	if o.DropInfected != nil {
		v := *o.DropInfected

		if err = d.Set("drop_infected", v); err != nil {
			return diag.Errorf("error reading drop_infected: %v", err)
		}
	}

	if o.DropMachineLearning != nil {
		v := *o.DropMachineLearning

		if err = d.Set("drop_machine_learning", v); err != nil {
			return diag.Errorf("error reading drop_machine_learning: %v", err)
		}
	}

	if o.Lowspace != nil {
		v := *o.Lowspace

		if err = d.Set("lowspace", v); err != nil {
			return diag.Errorf("error reading lowspace: %v", err)
		}
	}

	if o.Maxfilesize != nil {
		v := *o.Maxfilesize

		if err = d.Set("maxfilesize", v); err != nil {
			return diag.Errorf("error reading maxfilesize: %v", err)
		}
	}

	if o.QuarantineQuota != nil {
		v := *o.QuarantineQuota

		if err = d.Set("quarantine_quota", v); err != nil {
			return diag.Errorf("error reading quarantine_quota: %v", err)
		}
	}

	if o.StoreBlocked != nil {
		v := *o.StoreBlocked

		if err = d.Set("store_blocked", v); err != nil {
			return diag.Errorf("error reading store_blocked: %v", err)
		}
	}

	if o.StoreHeuristic != nil {
		v := *o.StoreHeuristic

		if err = d.Set("store_heuristic", v); err != nil {
			return diag.Errorf("error reading store_heuristic: %v", err)
		}
	}

	if o.StoreInfected != nil {
		v := *o.StoreInfected

		if err = d.Set("store_infected", v); err != nil {
			return diag.Errorf("error reading store_infected: %v", err)
		}
	}

	if o.StoreMachineLearning != nil {
		v := *o.StoreMachineLearning

		if err = d.Set("store_machine_learning", v); err != nil {
			return diag.Errorf("error reading store_machine_learning: %v", err)
		}
	}

	return nil
}

func getObjectAntivirusQuarantine(d *schema.ResourceData, sv string) (*models.AntivirusQuarantine, diag.Diagnostics) {
	obj := models.AntivirusQuarantine{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("agelimit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("agelimit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Agelimit = &tmp
		}
	}
	if v1, ok := d.GetOk("destination"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("destination", sv)
				diags = append(diags, e)
			}
			obj.Destination = &v2
		}
	}
	if v1, ok := d.GetOk("drop_blocked"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("drop_blocked", sv)
				diags = append(diags, e)
			}
			obj.DropBlocked = &v2
		}
	}
	if v1, ok := d.GetOk("drop_heuristic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("drop_heuristic", sv)
				diags = append(diags, e)
			}
			obj.DropHeuristic = &v2
		}
	}
	if v1, ok := d.GetOk("drop_infected"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("drop_infected", sv)
				diags = append(diags, e)
			}
			obj.DropInfected = &v2
		}
	}
	if v1, ok := d.GetOk("drop_machine_learning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("drop_machine_learning", sv)
				diags = append(diags, e)
			}
			obj.DropMachineLearning = &v2
		}
	}
	if v1, ok := d.GetOk("lowspace"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lowspace", sv)
				diags = append(diags, e)
			}
			obj.Lowspace = &v2
		}
	}
	if v1, ok := d.GetOk("maxfilesize"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("maxfilesize", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Maxfilesize = &tmp
		}
	}
	if v1, ok := d.GetOk("quarantine_quota"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("quarantine_quota", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QuarantineQuota = &tmp
		}
	}
	if v1, ok := d.GetOk("store_blocked"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("store_blocked", sv)
				diags = append(diags, e)
			}
			obj.StoreBlocked = &v2
		}
	}
	if v1, ok := d.GetOk("store_heuristic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("store_heuristic", sv)
				diags = append(diags, e)
			}
			obj.StoreHeuristic = &v2
		}
	}
	if v1, ok := d.GetOk("store_infected"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("store_infected", sv)
				diags = append(diags, e)
			}
			obj.StoreInfected = &v2
		}
	}
	if v1, ok := d.GetOk("store_machine_learning"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("store_machine_learning", sv)
				diags = append(diags, e)
			}
			obj.StoreMachineLearning = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectAntivirusQuarantine(d *schema.ResourceData, sv string) (*models.AntivirusQuarantine, diag.Diagnostics) {
	obj := models.AntivirusQuarantine{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
