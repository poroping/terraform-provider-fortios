// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure custom application signatures.

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

func resourceApplicationCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure custom application signatures.",

		CreateContext: resourceApplicationCustomCreate,
		ReadContext:   resourceApplicationCustomRead,
		UpdateContext: resourceApplicationCustomUpdate,
		DeleteContext: resourceApplicationCustomDelete,

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
			"behavior": {
				Type: schema.TypeString,

				Description: "Custom application signature behavior.",
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Type: schema.TypeInt,

				Description: "Custom application category ID (use ? to view available options).",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Custom application category ID (use ? to view available options).",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type: schema.TypeString,

				Description: "Custom application signature protocol.",
				Optional:    true,
				Computed:    true,
			},
			"signature": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),

				Description: "The text that makes up the actual custom application signature.",
				Optional:    true,
				Computed:    true,
			},
			"tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Signature tag.",
				ForceNew:    true,
				Required:    true,
			},
			"technology": {
				Type: schema.TypeString,

				Description: "Custom application signature technology.",
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Type: schema.TypeString,

				Description: "Custom application signature vendor.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceApplicationCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "tag"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating ApplicationCustom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectApplicationCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateApplicationCustom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationCustom")
	}

	return resourceApplicationCustomRead(ctx, d, meta)
}

func resourceApplicationCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectApplicationCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateApplicationCustom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ApplicationCustom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationCustom")
	}

	return resourceApplicationCustomRead(ctx, d, meta)
}

func resourceApplicationCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteApplicationCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ApplicationCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationCustom resource: %v", err)
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

	diags := refreshObjectApplicationCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectApplicationCustom(d *schema.ResourceData, o *models.ApplicationCustom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Behavior != nil {
		v := *o.Behavior

		if err = d.Set("behavior", v); err != nil {
			return diag.Errorf("error reading behavior: %v", err)
		}
	}

	if o.Category != nil {
		v := *o.Category

		if err = d.Set("category", v); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Signature != nil {
		v := *o.Signature

		if err = d.Set("signature", v); err != nil {
			return diag.Errorf("error reading signature: %v", err)
		}
	}

	if o.Tag != nil {
		v := *o.Tag

		if err = d.Set("tag", v); err != nil {
			return diag.Errorf("error reading tag: %v", err)
		}
	}

	if o.Technology != nil {
		v := *o.Technology

		if err = d.Set("technology", v); err != nil {
			return diag.Errorf("error reading technology: %v", err)
		}
	}

	if o.Vendor != nil {
		v := *o.Vendor

		if err = d.Set("vendor", v); err != nil {
			return diag.Errorf("error reading vendor: %v", err)
		}
	}

	return nil
}

func getObjectApplicationCustom(d *schema.ResourceData, sv string) (*models.ApplicationCustom, diag.Diagnostics) {
	obj := models.ApplicationCustom{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("behavior"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("behavior", sv)
				diags = append(diags, e)
			}
			obj.Behavior = &v2
		}
	}
	if v1, ok := d.GetOk("category"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("category", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Category = &tmp
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Fosid = &tmp
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v1, ok := d.GetOk("signature"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("signature", sv)
				diags = append(diags, e)
			}
			obj.Signature = &v2
		}
	}
	if v1, ok := d.GetOk("tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tag", sv)
				diags = append(diags, e)
			}
			obj.Tag = &v2
		}
	}
	if v1, ok := d.GetOk("technology"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("technology", sv)
				diags = append(diags, e)
			}
			obj.Technology = &v2
		}
	}
	if v1, ok := d.GetOk("vendor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vendor", sv)
				diags = append(diags, e)
			}
			obj.Vendor = &v2
		}
	}
	return &obj, diags
}
