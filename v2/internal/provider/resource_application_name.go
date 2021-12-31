// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure application signatures.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceApplicationName() *schema.Resource {
	return &schema.Resource{
		Description: "Configure application signatures.",

		CreateContext: resourceApplicationNameCreate,
		ReadContext:   resourceApplicationNameRead,
		UpdateContext: resourceApplicationNameUpdate,
		DeleteContext: resourceApplicationNameDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"behavior": {
				Type: schema.TypeString,

				Description: "Application behavior.",
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Type: schema.TypeInt,

				Description: "Application category ID.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Application ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"metadata": {
				Type:        schema.TypeList,
				Description: "Meta data.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"metaid": {
							Type: schema.TypeInt,

							Description: "Meta ID.",
							Optional:    true,
							Computed:    true,
						},
						"valueid": {
							Type: schema.TypeInt,

							Description: "Value ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Application name.",
				ForceNew:    true,
				Required:    true,
			},
			"parameter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Application parameter name.",
				Optional:    true,
				Computed:    true,
			},
			"parameters": {
				Type:        schema.TypeList,
				Description: "Application parameters.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"defaultvalue": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 199),

							Description: "Parameter default value.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Parameter name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"popularity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Application popularity.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type: schema.TypeString,

				Description: "Application protocol.",
				Optional:    true,
				Computed:    true,
			},
			"risk": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Application risk.",
				Optional:    true,
				Computed:    true,
			},
			"sub_category": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Application sub-category ID.",
				Optional:    true,
				Computed:    true,
			},
			"technology": {
				Type: schema.TypeString,

				Description: "Application technology.",
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Type: schema.TypeString,

				Description: "Application vendor.",
				Optional:    true,
				Computed:    true,
			},
			"weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Application weight.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceApplicationNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ApplicationName resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectApplicationName(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateApplicationName(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationName")
	}

	return resourceApplicationNameRead(ctx, d, meta)
}

func resourceApplicationNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectApplicationName(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateApplicationName(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ApplicationName resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationName")
	}

	return resourceApplicationNameRead(ctx, d, meta)
}

func resourceApplicationNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteApplicationName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ApplicationName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationName(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationName resource: %v", err)
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

	diags := refreshObjectApplicationName(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenApplicationNameMetadata(v *[]models.ApplicationNameMetadata, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Metaid; tmp != nil {
				v["metaid"] = *tmp
			}

			if tmp := cfg.Valueid; tmp != nil {
				v["valueid"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationNameParameters(v *[]models.ApplicationNameParameters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DefaultValue; tmp != nil {
				v["defaultvalue"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectApplicationName(d *schema.ResourceData, o *models.ApplicationName, sv string, sort bool) diag.Diagnostics {
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

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Metadata != nil {
		if err = d.Set("metadata", flattenApplicationNameMetadata(o.Metadata, sort)); err != nil {
			return diag.Errorf("error reading metadata: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Parameter != nil {
		v := *o.Parameter

		if err = d.Set("parameter", v); err != nil {
			return diag.Errorf("error reading parameter: %v", err)
		}
	}

	if o.Parameters != nil {
		if err = d.Set("parameters", flattenApplicationNameParameters(o.Parameters, sort)); err != nil {
			return diag.Errorf("error reading parameters: %v", err)
		}
	}

	if o.Popularity != nil {
		v := *o.Popularity

		if err = d.Set("popularity", v); err != nil {
			return diag.Errorf("error reading popularity: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Risk != nil {
		v := *o.Risk

		if err = d.Set("risk", v); err != nil {
			return diag.Errorf("error reading risk: %v", err)
		}
	}

	if o.SubCategory != nil {
		v := *o.SubCategory

		if err = d.Set("sub_category", v); err != nil {
			return diag.Errorf("error reading sub_category: %v", err)
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

	if o.Weight != nil {
		v := *o.Weight

		if err = d.Set("weight", v); err != nil {
			return diag.Errorf("error reading weight: %v", err)
		}
	}

	return nil
}

func expandApplicationNameMetadata(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationNameMetadata, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationNameMetadata

	for i := range l {
		tmp := models.ApplicationNameMetadata{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metaid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Metaid = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.valueid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Valueid = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationNameParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationNameParameters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationNameParameters

	for i := range l {
		tmp := models.ApplicationNameParameters{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.defaultvalue", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DefaultValue = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectApplicationName(d *schema.ResourceData, sv string) (*models.ApplicationName, diag.Diagnostics) {
	obj := models.ApplicationName{}
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
	if v, ok := d.GetOk("metadata"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("metadata", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationNameMetadata(d, v, "metadata", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Metadata = t
		}
	} else if d.HasChange("metadata") {
		old, new := d.GetChange("metadata")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Metadata = &[]models.ApplicationNameMetadata{}
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
	if v1, ok := d.GetOk("parameter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("parameter", sv)
				diags = append(diags, e)
			}
			obj.Parameter = &v2
		}
	}
	if v, ok := d.GetOk("parameters"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("parameters", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationNameParameters(d, v, "parameters", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Parameters = t
		}
	} else if d.HasChange("parameters") {
		old, new := d.GetChange("parameters")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Parameters = &[]models.ApplicationNameParameters{}
		}
	}
	if v1, ok := d.GetOk("popularity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("popularity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Popularity = &tmp
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
	if v1, ok := d.GetOk("risk"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("risk", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Risk = &tmp
		}
	}
	if v1, ok := d.GetOk("sub_category"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("sub_category", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SubCategory = &tmp
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
	if v1, ok := d.GetOk("weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Weight = &tmp
		}
	}
	return &obj, diags
}
