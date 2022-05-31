// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure auto script.

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

func resourceSystemAutoScript() *schema.Resource {
	return &schema.Resource{
		Description: "Configure auto script.",

		CreateContext: resourceSystemAutoScriptCreate,
		ReadContext:   resourceSystemAutoScriptRead,
		UpdateContext: resourceSystemAutoScriptUpdate,
		DeleteContext: resourceSystemAutoScriptDelete,

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
			"interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31557600),

				Description: "Repeat interval in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Auto script name.",
				ForceNew:    true,
				Required:    true,
			},
			"output_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1024),

				Description: "Number of megabytes to limit script output to (10 - 1024, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"repeat": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Number of times to repeat this script (0 = infinite).",
				Optional:    true,
				Computed:    true,
			},
			"script": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "List of FortiOS CLI commands to repeat.",
				Optional:    true,
				Computed:    true,
			},
			"start": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"manual", "auto"}, false),

				Description: "Script starting mode.",
				Optional:    true,
				Computed:    true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),

				Description: "Maximum running time for this script in seconds (0 = no timeout).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAutoScriptCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemAutoScript resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemAutoScript(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAutoScript(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutoScript")
	}

	return resourceSystemAutoScriptRead(ctx, d, meta)
}

func resourceSystemAutoScriptUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutoScript(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAutoScript(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAutoScript resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutoScript")
	}

	return resourceSystemAutoScriptRead(ctx, d, meta)
}

func resourceSystemAutoScriptDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAutoScript(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAutoScript resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoScriptRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutoScript(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutoScript resource: %v", err)
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

	diags := refreshObjectSystemAutoScript(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemAutoScript(d *schema.ResourceData, o *models.SystemAutoScript, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Interval != nil {
		v := *o.Interval

		if err = d.Set("interval", v); err != nil {
			return diag.Errorf("error reading interval: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OutputSize != nil {
		v := *o.OutputSize

		if err = d.Set("output_size", v); err != nil {
			return diag.Errorf("error reading output_size: %v", err)
		}
	}

	if o.Repeat != nil {
		v := *o.Repeat

		if err = d.Set("repeat", v); err != nil {
			return diag.Errorf("error reading repeat: %v", err)
		}
	}

	if o.Script != nil {
		v := *o.Script

		if err = d.Set("script", v); err != nil {
			return diag.Errorf("error reading script: %v", err)
		}
	}

	if o.Start != nil {
		v := *o.Start

		if err = d.Set("start", v); err != nil {
			return diag.Errorf("error reading start: %v", err)
		}
	}

	if o.Timeout != nil {
		v := *o.Timeout

		if err = d.Set("timeout", v); err != nil {
			return diag.Errorf("error reading timeout: %v", err)
		}
	}

	return nil
}

func getObjectSystemAutoScript(d *schema.ResourceData, sv string) (*models.SystemAutoScript, diag.Diagnostics) {
	obj := models.SystemAutoScript{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Interval = &tmp
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
	if v1, ok := d.GetOk("output_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("output_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.OutputSize = &tmp
		}
	}
	if v1, ok := d.GetOk("repeat"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("repeat", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Repeat = &tmp
		}
	}
	if v1, ok := d.GetOk("script"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("script", sv)
				diags = append(diags, e)
			}
			obj.Script = &v2
		}
	}
	if v1, ok := d.GetOk("start"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start", sv)
				diags = append(diags, e)
			}
			obj.Start = &v2
		}
	}
	if v1, ok := d.GetOk("timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Timeout = &tmp
		}
	}
	return &obj, diags
}
