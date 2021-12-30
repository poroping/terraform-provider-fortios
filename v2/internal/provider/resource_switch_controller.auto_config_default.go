// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
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

func resourceSwitchControllerautoConfigDefault() *schema.Resource {
	return &schema.Resource{
		Description: "Policies which are applied automatically to all ISL/ICL/FortiLink interfaces.",

		CreateContext: resourceSwitchControllerautoConfigDefaultCreate,
		ReadContext:   resourceSwitchControllerautoConfigDefaultRead,
		UpdateContext: resourceSwitchControllerautoConfigDefaultUpdate,
		DeleteContext: resourceSwitchControllerautoConfigDefaultDelete,

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

func resourceSwitchControllerautoConfigDefaultCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerautoConfigDefault(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerautoConfigDefault(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerautoConfigDefault")
	}

	return resourceSwitchControllerautoConfigDefaultRead(ctx, d, meta)
}

func resourceSwitchControllerautoConfigDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerautoConfigDefault(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerautoConfigDefault(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerautoConfigDefault resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerautoConfigDefault")
	}

	return resourceSwitchControllerautoConfigDefaultRead(ctx, d, meta)
}

func resourceSwitchControllerautoConfigDefaultDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerautoConfigDefault(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerautoConfigDefault resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerautoConfigDefaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerautoConfigDefault(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerautoConfigDefault resource: %v", err)
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

	diags := refreshObjectSwitchControllerautoConfigDefault(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerautoConfigDefault(d *schema.ResourceData, o *models.SwitchControllerautoConfigDefault, sv string, sort bool) diag.Diagnostics {
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

func getObjectSwitchControllerautoConfigDefault(d *schema.ResourceData, sv string) (*models.SwitchControllerautoConfigDefault, diag.Diagnostics) {
	obj := models.SwitchControllerautoConfigDefault{}
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
