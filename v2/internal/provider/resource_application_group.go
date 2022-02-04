// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure firewall application groups.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceApplicationGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Configure firewall application groups.",

		CreateContext: resourceApplicationGroupCreate,
		ReadContext:   resourceApplicationGroupRead,
		UpdateContext: resourceApplicationGroupUpdate,
		DeleteContext: resourceApplicationGroupDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Application IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"behavior": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Application behavior filter.",
				Optional:         true,
				Computed:         true,
			},
			"category": {
				Type:        schema.TypeList,
				Description: "Application category ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Category IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comments.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Application group name.",
				ForceNew:    true,
				Required:    true,
			},
			"popularity": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Application popularity filter (1 - 5, from least to most popular).",
				Optional:         true,
				Computed:         true,
			},
			"protocols": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Application protocol filter.",
				Optional:         true,
				Computed:         true,
			},
			"risk": {
				Type:        schema.TypeList,
				Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type: schema.TypeInt,

							Description: "Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"technology": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Application technology filter.",
				Optional:         true,
				Computed:         true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"application", "filter"}, false),

				Description: "Application group type.",
				Optional:    true,
				Computed:    true,
			},
			"vendor": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Application vendor filter.",
				Optional:         true,
				Computed:         true,
			},
		},
	}
}

func resourceApplicationGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ApplicationGroup resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectApplicationGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateApplicationGroup(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationGroup")
	}

	return resourceApplicationGroupRead(ctx, d, meta)
}

func resourceApplicationGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectApplicationGroup(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateApplicationGroup(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ApplicationGroup resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ApplicationGroup")
	}

	return resourceApplicationGroupRead(ctx, d, meta)
}

func resourceApplicationGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteApplicationGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ApplicationGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadApplicationGroup(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ApplicationGroup resource: %v", err)
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

	diags := refreshObjectApplicationGroup(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenApplicationGroupApplication(v *[]models.ApplicationGroupApplication, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationGroupCategory(v *[]models.ApplicationGroupCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenApplicationGroupRisk(v *[]models.ApplicationGroupRisk, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Level; tmp != nil {
				v["level"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "level")
	}

	return flat
}

func refreshObjectApplicationGroup(d *schema.ResourceData, o *models.ApplicationGroup, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Application != nil {
		if err = d.Set("application", flattenApplicationGroupApplication(o.Application, sort)); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.Behavior != nil {
		v := *o.Behavior

		if err = d.Set("behavior", v); err != nil {
			return diag.Errorf("error reading behavior: %v", err)
		}
	}

	if o.Category != nil {
		if err = d.Set("category", flattenApplicationGroupCategory(o.Category, sort)); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Popularity != nil {
		v := *o.Popularity

		if err = d.Set("popularity", v); err != nil {
			return diag.Errorf("error reading popularity: %v", err)
		}
	}

	if o.Protocols != nil {
		v := *o.Protocols

		if err = d.Set("protocols", v); err != nil {
			return diag.Errorf("error reading protocols: %v", err)
		}
	}

	if o.Risk != nil {
		if err = d.Set("risk", flattenApplicationGroupRisk(o.Risk, sort)); err != nil {
			return diag.Errorf("error reading risk: %v", err)
		}
	}

	if o.Technology != nil {
		v := *o.Technology

		if err = d.Set("technology", v); err != nil {
			return diag.Errorf("error reading technology: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
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

func expandApplicationGroupApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationGroupApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationGroupApplication

	for i := range l {
		tmp := models.ApplicationGroupApplication{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationGroupCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationGroupCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationGroupCategory

	for i := range l {
		tmp := models.ApplicationGroupCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandApplicationGroupRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ApplicationGroupRisk, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ApplicationGroupRisk

	for i := range l {
		tmp := models.ApplicationGroupRisk{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Level = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectApplicationGroup(d *schema.ResourceData, sv string) (*models.ApplicationGroup, diag.Diagnostics) {
	obj := models.ApplicationGroup{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("application"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("application", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationGroupApplication(d, v, "application", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Application = t
		}
	} else if d.HasChange("application") {
		old, new := d.GetChange("application")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Application = &[]models.ApplicationGroupApplication{}
		}
	}
	if v1, ok := d.GetOk("behavior"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("behavior", sv)
				diags = append(diags, e)
			}
			obj.Behavior = &v2
		}
	}
	if v, ok := d.GetOk("category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("category", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationGroupCategory(d, v, "category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Category = t
		}
	} else if d.HasChange("category") {
		old, new := d.GetChange("category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Category = &[]models.ApplicationGroupCategory{}
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
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("popularity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("popularity", sv)
				diags = append(diags, e)
			}
			obj.Popularity = &v2
		}
	}
	if v1, ok := d.GetOk("protocols"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("protocols", sv)
				diags = append(diags, e)
			}
			obj.Protocols = &v2
		}
	}
	if v, ok := d.GetOk("risk"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("risk", sv)
			diags = append(diags, e)
		}
		t, err := expandApplicationGroupRisk(d, v, "risk", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Risk = t
		}
	} else if d.HasChange("risk") {
		old, new := d.GetChange("risk")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Risk = &[]models.ApplicationGroupRisk{}
		}
	}
	if v1, ok := d.GetOk("technology"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("technology", sv)
				diags = append(diags, e)
			}
			obj.Technology = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("vendor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("vendor", sv)
				diags = append(diags, e)
			}
			obj.Vendor = &v2
		}
	}
	return &obj, diags
}
