// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Replacement messages.

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

func resourceSystemReplacemsgAutomation() *schema.Resource {
	return &schema.Resource{
		Description: "Replacement messages.",

		CreateContext: resourceSystemReplacemsgAutomationCreate,
		ReadContext:   resourceSystemReplacemsgAutomationRead,
		UpdateContext: resourceSystemReplacemsgAutomationUpdate,
		DeleteContext: resourceSystemReplacemsgAutomationDelete,

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
			"buffer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),

				Description: "Message string.",
				Optional:    true,
				Computed:    true,
			},
			"format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "text", "html"}, false),

				Description: "Format flag.",
				Optional:    true,
				Computed:    true,
			},
			"header": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "http", "8bit"}, false),

				Description: "Header flag.",
				Optional:    true,
				Computed:    true,
			},
			"msg_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 28),

				Description: "Message type.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSystemReplacemsgAutomationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "msg_type"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemReplacemsgAutomation resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemReplacemsgAutomation(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemReplacemsgAutomation(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemReplacemsgAutomation")
	}

	return resourceSystemReplacemsgAutomationRead(ctx, d, meta)
}

func resourceSystemReplacemsgAutomationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemReplacemsgAutomation(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemReplacemsgAutomation(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemReplacemsgAutomation resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemReplacemsgAutomation")
	}

	return resourceSystemReplacemsgAutomationRead(ctx, d, meta)
}

func resourceSystemReplacemsgAutomationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemReplacemsgAutomation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemReplacemsgAutomation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgAutomationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemReplacemsgAutomation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemReplacemsgAutomation resource: %v", err)
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

	diags := refreshObjectSystemReplacemsgAutomation(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemReplacemsgAutomation(d *schema.ResourceData, o *models.SystemReplacemsgAutomation, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Buffer != nil {
		v := *o.Buffer

		if err = d.Set("buffer", v); err != nil {
			return diag.Errorf("error reading buffer: %v", err)
		}
	}

	if o.Format != nil {
		v := *o.Format

		if err = d.Set("format", v); err != nil {
			return diag.Errorf("error reading format: %v", err)
		}
	}

	if o.Header != nil {
		v := *o.Header

		if err = d.Set("header", v); err != nil {
			return diag.Errorf("error reading header: %v", err)
		}
	}

	if o.MsgType != nil {
		v := *o.MsgType

		if err = d.Set("msg_type", v); err != nil {
			return diag.Errorf("error reading msg_type: %v", err)
		}
	}

	return nil
}

func getObjectSystemReplacemsgAutomation(d *schema.ResourceData, sv string) (*models.SystemReplacemsgAutomation, diag.Diagnostics) {
	obj := models.SystemReplacemsgAutomation{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("buffer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("buffer", sv)
				diags = append(diags, e)
			}
			obj.Buffer = &v2
		}
	}
	if v1, ok := d.GetOk("format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("format", sv)
				diags = append(diags, e)
			}
			obj.Format = &v2
		}
	}
	if v1, ok := d.GetOk("header"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header", sv)
				diags = append(diags, e)
			}
			obj.Header = &v2
		}
	}
	if v1, ok := d.GetOk("msg_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("msg_type", sv)
				diags = append(diags, e)
			}
			obj.MsgType = &v2
		}
	}
	return &obj, diags
}
