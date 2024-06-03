// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.

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

func resourceSwitchControllerAutoConfigDefault() *schema.Resource {
	return &schema.Resource{
		Description: "Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.",

		CreateContext: resourceSwitchControllerAutoConfigDefaultCreate,
		ReadContext:   resourceSwitchControllerAutoConfigDefaultRead,
		UpdateContext: resourceSwitchControllerAutoConfigDefaultUpdate,
		DeleteContext: resourceSwitchControllerAutoConfigDefaultDelete,

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
			"fgt_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Default FortiLink auto-config policy.",
				Optional:    true,
				Computed:    true,
			},
			"icl_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Default ICL auto-config policy.",
				Optional:    true,
				Computed:    true,
			},
			"isl_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Default ISL auto-config policy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigDefaultCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerAutoConfigDefault(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerAutoConfigDefault(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerAutoConfigDefault")
	}

	return resourceSwitchControllerAutoConfigDefaultRead(ctx, d, meta)
}

func resourceSwitchControllerAutoConfigDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerAutoConfigDefault(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerAutoConfigDefault(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerAutoConfigDefault resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerAutoConfigDefault")
	}

	return resourceSwitchControllerAutoConfigDefaultRead(ctx, d, meta)
}

func resourceSwitchControllerAutoConfigDefaultDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSwitchControllerAutoConfigDefault(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSwitchControllerAutoConfigDefault(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerAutoConfigDefault resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigDefaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerAutoConfigDefault(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerAutoConfigDefault resource: %v", err)
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

	diags := refreshObjectSwitchControllerAutoConfigDefault(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, o *models.SwitchControllerAutoConfigDefault, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FgtPolicy != nil {
		v := *o.FgtPolicy

		if err = d.Set("fgt_policy", v); err != nil {
			return diag.Errorf("error reading fgt_policy: %v", err)
		}
	}

	if o.IclPolicy != nil {
		v := *o.IclPolicy

		if err = d.Set("icl_policy", v); err != nil {
			return diag.Errorf("error reading icl_policy: %v", err)
		}
	}

	if o.IslPolicy != nil {
		v := *o.IslPolicy

		if err = d.Set("isl_policy", v); err != nil {
			return diag.Errorf("error reading isl_policy: %v", err)
		}
	}

	return nil
}

func getObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, sv string) (*models.SwitchControllerAutoConfigDefault, diag.Diagnostics) {
	obj := models.SwitchControllerAutoConfigDefault{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fgt_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fgt_policy", sv)
				diags = append(diags, e)
			}
			obj.FgtPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("icl_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icl_policy", sv)
				diags = append(diags, e)
			}
			obj.IclPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("isl_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("isl_policy", sv)
				diags = append(diags, e)
			}
			obj.IslPolicy = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, sv string) (*models.SwitchControllerAutoConfigDefault, diag.Diagnostics) {
	obj := models.SwitchControllerAutoConfigDefault{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
