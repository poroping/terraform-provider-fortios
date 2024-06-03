// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FIPS-CC mode.

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

func resourceSystemFipsCc() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FIPS-CC mode.",

		CreateContext: resourceSystemFipsCcCreate,
		ReadContext:   resourceSystemFipsCcRead,
		UpdateContext: resourceSystemFipsCcUpdate,
		DeleteContext: resourceSystemFipsCcDelete,

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
			"entropy_token": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable"}, false),

				Description: "Enable/disable/dynamic entropy token.",
				Optional:    true,
				Computed:    true,
			},
			"key_generation_self_test": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable self tests after key generation.",
				Optional:    true,
				Computed:    true,
			},
			"self_test_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "Self test period.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ciphers for FIPS mode of operation.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemFipsCcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFipsCc(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemFipsCc(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFipsCc")
	}

	return resourceSystemFipsCcRead(ctx, d, meta)
}

func resourceSystemFipsCcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFipsCc(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemFipsCc(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemFipsCc resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFipsCc")
	}

	return resourceSystemFipsCcRead(ctx, d, meta)
}

func resourceSystemFipsCcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemFipsCc(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemFipsCc(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemFipsCc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFipsCcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFipsCc(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFipsCc resource: %v", err)
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

	diags := refreshObjectSystemFipsCc(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemFipsCc(d *schema.ResourceData, o *models.SystemFipsCc, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.EntropyToken != nil {
		v := *o.EntropyToken

		if err = d.Set("entropy_token", v); err != nil {
			return diag.Errorf("error reading entropy_token: %v", err)
		}
	}

	if o.KeyGenerationSelfTest != nil {
		v := *o.KeyGenerationSelfTest

		if err = d.Set("key_generation_self_test", v); err != nil {
			return diag.Errorf("error reading key_generation_self_test: %v", err)
		}
	}

	if o.SelfTestPeriod != nil {
		v := *o.SelfTestPeriod

		if err = d.Set("self_test_period", v); err != nil {
			return diag.Errorf("error reading self_test_period: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectSystemFipsCc(d *schema.ResourceData, sv string) (*models.SystemFipsCc, diag.Diagnostics) {
	obj := models.SystemFipsCc{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("entropy_token"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.8") {
				e := utils.AttributeVersionWarning("entropy_token", sv)
				diags = append(diags, e)
			}
			obj.EntropyToken = &v2
		}
	}
	if v1, ok := d.GetOk("key_generation_self_test"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_generation_self_test", sv)
				diags = append(diags, e)
			}
			obj.KeyGenerationSelfTest = &v2
		}
	}
	if v1, ok := d.GetOk("self_test_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("self_test_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SelfTestPeriod = &tmp
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
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemFipsCc(d *schema.ResourceData, sv string) (*models.SystemFipsCc, diag.Diagnostics) {
	obj := models.SystemFipsCc{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
