// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report dataset configuration.

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

func resourceReportDataset() *schema.Resource {
	return &schema.Resource{
		Description: "Report dataset configuration.",

		CreateContext: resourceReportDatasetCreate,
		ReadContext:   resourceReportDatasetRead,
		UpdateContext: resourceReportDatasetUpdate,
		DeleteContext: resourceReportDatasetDelete,

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
			"field": {
				Type:        schema.TypeList,
				Description: "Fields.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"displayname": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Display name.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Field ID (1 to number of columns in SQL result).",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"text", "integer", "double"}, false),

							Description: "Field type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"parameters": {
				Type:        schema.TypeList,
				Description: "Parameters.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"text", "integer", "double", "long-integer", "date-time"}, false),

							Description: "Data type.",
							Optional:    true,
							Computed:    true,
						},
						"display_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Display name.",
							Optional:    true,
							Computed:    true,
						},
						"field": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "SQL field name.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Parameter ID (1 to number of columns in SQL result).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"policy": {
				Type: schema.TypeInt,

				Description: "Used by monitor policy.",
				Optional:    true,
				Computed:    true,
			},
			"query": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2303),

				Description: "SQL query statement.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceReportDatasetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ReportDataset resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectReportDataset(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateReportDataset(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportDataset")
	}

	return resourceReportDatasetRead(ctx, d, meta)
}

func resourceReportDatasetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportDataset(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateReportDataset(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ReportDataset resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportDataset")
	}

	return resourceReportDatasetRead(ctx, d, meta)
}

func resourceReportDatasetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteReportDataset(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ReportDataset resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportDatasetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportDataset(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportDataset resource: %v", err)
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

	diags := refreshObjectReportDataset(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenReportDatasetField(v *[]models.ReportDatasetField, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Displayname; tmp != nil {
				v["displayname"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenReportDatasetParameters(v *[]models.ReportDatasetParameters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DataType; tmp != nil {
				v["data_type"] = *tmp
			}

			if tmp := cfg.DisplayName; tmp != nil {
				v["display_name"] = *tmp
			}

			if tmp := cfg.Field; tmp != nil {
				v["field"] = *tmp
			}

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

func refreshObjectReportDataset(d *schema.ResourceData, o *models.ReportDataset, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Field != nil {
		if err = d.Set("field", flattenReportDatasetField(o.Field, sort)); err != nil {
			return diag.Errorf("error reading field: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Parameters != nil {
		if err = d.Set("parameters", flattenReportDatasetParameters(o.Parameters, sort)); err != nil {
			return diag.Errorf("error reading parameters: %v", err)
		}
	}

	if o.Policy != nil {
		v := *o.Policy

		if err = d.Set("policy", v); err != nil {
			return diag.Errorf("error reading policy: %v", err)
		}
	}

	if o.Query != nil {
		v := *o.Query

		if err = d.Set("query", v); err != nil {
			return diag.Errorf("error reading query: %v", err)
		}
	}

	return nil
}

func expandReportDatasetField(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportDatasetField, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportDatasetField

	for i := range l {
		tmp := models.ReportDatasetField{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.displayname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Displayname = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportDatasetParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportDatasetParameters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportDatasetParameters

	for i := range l {
		tmp := models.ReportDatasetParameters{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.data_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DataType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.display_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DisplayName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.field", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Field = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectReportDataset(d *schema.ResourceData, sv string) (*models.ReportDataset, diag.Diagnostics) {
	obj := models.ReportDataset{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("field"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("field", sv)
			diags = append(diags, e)
		}
		t, err := expandReportDatasetField(d, v, "field", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Field = t
		}
	} else if d.HasChange("field") {
		old, new := d.GetChange("field")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Field = &[]models.ReportDatasetField{}
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
	if v, ok := d.GetOk("parameters"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("parameters", sv)
			diags = append(diags, e)
		}
		t, err := expandReportDatasetParameters(d, v, "parameters", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Parameters = t
		}
	} else if d.HasChange("parameters") {
		old, new := d.GetChange("parameters")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Parameters = &[]models.ReportDatasetParameters{}
		}
	}
	if v1, ok := d.GetOk("policy"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policy", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Policy = &tmp
		}
	}
	if v1, ok := d.GetOk("query"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query", sv)
				diags = append(diags, e)
			}
			obj.Query = &v2
		}
	}
	return &obj, diags
}
