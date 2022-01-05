// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: 3G MODEM custom.

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

func resourceSystem3gModemCustom() *schema.Resource {
	return &schema.Resource{
		Description: "3G MODEM custom.",

		CreateContext: resourceSystem3gModemCustomCreate,
		ReadContext:   resourceSystem3gModemCustomRead,
		UpdateContext: resourceSystem3gModemCustomUpdate,
		DeleteContext: resourceSystem3gModemCustomDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"class_id": {
				Type: schema.TypeString,

				Description: "USB interface class in hexadecimal format (00-ff).",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"init_string": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Init string in hexadecimal format (even length).",
				Optional:    true,
				Computed:    true,
			},
			"model": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "MODEM model name.",
				Optional:    true,
				Computed:    true,
			},
			"modeswitch_string": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "USB modeswitch arguments. e.g: '-v 1410 -p 9030 -V 1410 -P 9032 -u 3'",
				Optional:    true,
				Computed:    true,
			},
			"product_id": {
				Type: schema.TypeString,

				Description: "USB product ID in hexadecimal format (0000-ffff).",
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "MODEM vendor name.",
				Optional:    true,
				Computed:    true,
			},
			"vendor_id": {
				Type: schema.TypeString,

				Description: "USB vendor ID in hexadecimal format (0000-ffff).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystem3gModemCustomCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating System3gModemCustom resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystem3gModemCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystem3gModemCustom(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("System3gModemCustom")
	}

	return resourceSystem3gModemCustomRead(ctx, d, meta)
}

func resourceSystem3gModemCustomUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystem3gModemCustom(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystem3gModemCustom(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating System3gModemCustom resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("System3gModemCustom")
	}

	return resourceSystem3gModemCustomRead(ctx, d, meta)
}

func resourceSystem3gModemCustomDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystem3gModemCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting System3gModemCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystem3gModemCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystem3gModemCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading System3gModemCustom resource: %v", err)
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

	diags := refreshObjectSystem3gModemCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystem3gModemCustom(d *schema.ResourceData, o *models.System3gModemCustom, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ClassId != nil {
		v := *o.ClassId

		if err = d.Set("class_id", v); err != nil {
			return diag.Errorf("error reading class_id: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.InitString != nil {
		v := *o.InitString

		if err = d.Set("init_string", v); err != nil {
			return diag.Errorf("error reading init_string: %v", err)
		}
	}

	if o.Model != nil {
		v := *o.Model

		if err = d.Set("model", v); err != nil {
			return diag.Errorf("error reading model: %v", err)
		}
	}

	if o.ModeswitchString != nil {
		v := *o.ModeswitchString

		if err = d.Set("modeswitch_string", v); err != nil {
			return diag.Errorf("error reading modeswitch_string: %v", err)
		}
	}

	if o.ProductId != nil {
		v := *o.ProductId

		if err = d.Set("product_id", v); err != nil {
			return diag.Errorf("error reading product_id: %v", err)
		}
	}

	if o.Vendor != nil {
		v := *o.Vendor

		if err = d.Set("vendor", v); err != nil {
			return diag.Errorf("error reading vendor: %v", err)
		}
	}

	if o.VendorId != nil {
		v := *o.VendorId

		if err = d.Set("vendor_id", v); err != nil {
			return diag.Errorf("error reading vendor_id: %v", err)
		}
	}

	return nil
}

func getObjectSystem3gModemCustom(d *schema.ResourceData, sv string) (*models.System3gModemCustom, diag.Diagnostics) {
	obj := models.System3gModemCustom{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("class_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("class_id", sv)
				diags = append(diags, e)
			}
			obj.ClassId = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("init_string"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("init_string", sv)
				diags = append(diags, e)
			}
			obj.InitString = &v2
		}
	}
	if v1, ok := d.GetOk("model"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("model", sv)
				diags = append(diags, e)
			}
			obj.Model = &v2
		}
	}
	if v1, ok := d.GetOk("modeswitch_string"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("modeswitch_string", sv)
				diags = append(diags, e)
			}
			obj.ModeswitchString = &v2
		}
	}
	if v1, ok := d.GetOk("product_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("product_id", sv)
				diags = append(diags, e)
			}
			obj.ProductId = &v2
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
	if v1, ok := d.GetOk("vendor_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vendor_id", sv)
				diags = append(diags, e)
			}
			obj.VendorId = &v2
		}
	}
	return &obj, diags
}
