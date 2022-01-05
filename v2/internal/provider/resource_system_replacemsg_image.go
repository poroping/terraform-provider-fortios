// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure replacement message images.

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

func resourceSystemReplacemsgImage() *schema.Resource {
	return &schema.Resource{
		Description: "Configure replacement message images.",

		CreateContext: resourceSystemReplacemsgImageCreate,
		ReadContext:   resourceSystemReplacemsgImageRead,
		UpdateContext: resourceSystemReplacemsgImageUpdate,
		DeleteContext: resourceSystemReplacemsgImageDelete,

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
			"image_base64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),

				Description: "Image data.",
				Optional:    true,
				Computed:    true,
			},
			"image_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"gif", "jpg", "tiff", "png"}, false),

				Description: "Image type.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 23),

				Description: "Image name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSystemReplacemsgImageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemReplacemsgImage resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemReplacemsgImage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemReplacemsgImage(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemReplacemsgImage")
	}

	return resourceSystemReplacemsgImageRead(ctx, d, meta)
}

func resourceSystemReplacemsgImageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemReplacemsgImage(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemReplacemsgImage(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemReplacemsgImage resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemReplacemsgImage")
	}

	return resourceSystemReplacemsgImageRead(ctx, d, meta)
}

func resourceSystemReplacemsgImageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemReplacemsgImage(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemReplacemsgImage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgImageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemReplacemsgImage(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemReplacemsgImage resource: %v", err)
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

	diags := refreshObjectSystemReplacemsgImage(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemReplacemsgImage(d *schema.ResourceData, o *models.SystemReplacemsgImage, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ImageBase64 != nil {
		v := *o.ImageBase64

		if err = d.Set("image_base64", v); err != nil {
			return diag.Errorf("error reading image_base64: %v", err)
		}
	}

	if o.ImageType != nil {
		v := *o.ImageType

		if err = d.Set("image_type", v); err != nil {
			return diag.Errorf("error reading image_type: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func getObjectSystemReplacemsgImage(d *schema.ResourceData, sv string) (*models.SystemReplacemsgImage, diag.Diagnostics) {
	obj := models.SystemReplacemsgImage{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("image_base64"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("image_base64", sv)
				diags = append(diags, e)
			}
			obj.ImageBase64 = &v2
		}
	}
	if v1, ok := d.GetOk("image_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("image_type", sv)
				diags = append(diags, e)
			}
			obj.ImageType = &v2
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
	return &obj, diags
}
