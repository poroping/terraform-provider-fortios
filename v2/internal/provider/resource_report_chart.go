// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report chart widget configuration.

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

func resourceReportChart() *schema.Resource {
	return &schema.Resource{
		Description: "Report chart widget configuration.",

		CreateContext: resourceReportChartCreate,
		ReadContext:   resourceReportChartRead,
		UpdateContext: resourceReportChartUpdate,
		DeleteContext: resourceReportChartDelete,

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
			"background": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),

				Description: "Chart background.",
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"misc", "traffic", "event", "virus", "webfilter", "attack", "spam", "dlp", "app-ctrl", "vulnerability"}, false),

				Description: "Category.",
				Optional:    true,
				Computed:    true,
			},
			"category_series": {
				Type:        schema.TypeList,
				Description: "Category series of pie chart.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Category series value expression.",
							Optional:    true,
							Computed:    true,
						},
						"font_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),

							Description: "Font size of category-series title.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"color_palette": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),

				Description: "Color palette (system will pick color automatically by default).",
				Optional:    true,
				Computed:    true,
			},
			"column": {
				Type:        schema.TypeList,
				Description: "Table column definition.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"detail_unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Detail unit of column.",
							Optional:    true,
							Computed:    true,
						},
						"detail_value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Detail value of column.",
							Optional:    true,
							Computed:    true,
						},
						"footer_unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Footer unit of column.",
							Optional:    true,
							Computed:    true,
						},
						"footer_value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Footer value of column.",
							Optional:    true,
							Computed:    true,
						},
						"header_value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Display name of table header.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"mapping": {
							Type:        schema.TypeList,
							Description: "Show detail in certain display value for certain condition.",
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

										Description: "id",
										Optional:    true,
										Computed:    true,
									},
									"op": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"none", "greater", "greater-equal", "less", "less-equal", "equal", "between"}, false),

										Description: "Comparision operater.",
										Optional:    true,
										Computed:    true,
									},
									"value_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"integer", "string"}, false),

										Description: "Value type.",
										Optional:    true,
										Computed:    true,
									},
									"value1": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),

										Description: "Value 1.",
										Optional:    true,
										Computed:    true,
									},
									"value2": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),

										Description: "Value 2.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dataset": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Bind dataset to chart.",
				Optional:    true,
				Computed:    true,
			},
			"dimension": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"2D", "3D"}, false),

				Description: "Dimension.",
				Optional:    true,
				Computed:    true,
			},
			"drill_down_charts": {
				Type:        schema.TypeList,
				Description: "Drill down charts.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chart_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Drill down chart name.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Drill down chart ID.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable this drill down chart.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"favorite": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

				Description: "Favorite.",
				Optional:    true,
				Computed:    true,
			},
			"graph_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "bar", "pie", "line", "flow"}, false),

				Description: "Graph type.",
				Optional:    true,
				Computed:    true,
			},
			"legend": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/Disable Legend area.",
				Optional:    true,
				Computed:    true,
			},
			"legend_font_size": {
				Type: schema.TypeInt,

				Description: "Font size of legend area.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Chart Widget Name",
				ForceNew:    true,
				Required:    true,
			},
			"period": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"last24h", "last7d"}, false),

				Description: "Time period.",
				Optional:    true,
				Computed:    true,
			},
			"policy": {
				Type: schema.TypeInt,

				Description: "Used by monitor policy.",
				Optional:    true,
				Computed:    true,
			},
			"style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "manual"}, false),

				Description: "Style.",
				Optional:    true,
				Computed:    true,
			},
			"title": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Chart title.",
				Optional:    true,
				Computed:    true,
			},
			"title_font_size": {
				Type: schema.TypeInt,

				Description: "Font size of chart title.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"graph", "table"}, false),

				Description: "Chart type.",
				Optional:    true,
				Computed:    true,
			},
			"value_series": {
				Type:        schema.TypeList,
				Description: "Value series of pie chart.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Value series value expression.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"x_series": {
				Type:        schema.TypeList,
				Description: "X-series of chart.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caption": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "X-series caption.",
							Optional:    true,
							Computed:    true,
						},
						"caption_font_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),

							Description: "X-series caption font size.",
							Optional:    true,
							Computed:    true,
						},
						"databind": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "X-series value expression.",
							Optional:    true,
							Computed:    true,
						},
						"font_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),

							Description: "X-series label font size.",
							Optional:    true,
							Computed:    true,
						},
						"is_category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),

							Description: "X-series represent category or not.",
							Optional:    true,
							Computed:    true,
						},
						"label_angle": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"45-degree", "vertical", "horizontal"}, false),

							Description: "X-series label angle.",
							Optional:    true,
							Computed:    true,
						},
						"scale_direction": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"decrease", "increase"}, false),

							Description: "Scale increase or decrease.",
							Optional:    true,
							Computed:    true,
						},
						"scale_format": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"YYYY-MM-DD-HH-MM", "YYYY-MM-DD HH", "YYYY-MM-DD", "YYYY-MM", "YYYY", "HH-MM", "MM-DD"}, false),

							Description: "Date/time format.",
							Optional:    true,
							Computed:    true,
						},
						"scale_step": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Scale step.",
							Optional:    true,
							Computed:    true,
						},
						"scale_unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"minute", "hour", "day", "month", "year"}, false),

							Description: "Scale unit.",
							Optional:    true,
							Computed:    true,
						},
						"unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "X-series unit.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"y_series": {
				Type:        schema.TypeList,
				Description: "Y-series of chart.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caption": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Y-series caption.",
							Optional:    true,
							Computed:    true,
						},
						"caption_font_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),

							Description: "Y-series caption font size.",
							Optional:    true,
							Computed:    true,
						},
						"databind": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Y-series value expression.",
							Optional:    true,
							Computed:    true,
						},
						"extra_databind": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Extra Y-series value.",
							Optional:    true,
							Computed:    true,
						},
						"extra_y": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Allow another Y-series value",
							Optional:    true,
							Computed:    true,
						},
						"extra_y_legend": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Extra Y-series legend type/name.",
							Optional:    true,
							Computed:    true,
						},
						"font_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),

							Description: "Y-series label font size.",
							Optional:    true,
							Computed:    true,
						},
						"group": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Y-series group option.",
							Optional:    true,
							Computed:    true,
						},
						"label_angle": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"45-degree", "vertical", "horizontal"}, false),

							Description: "Y-series label angle.",
							Optional:    true,
							Computed:    true,
						},
						"unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Y-series unit.",
							Optional:    true,
							Computed:    true,
						},
						"y_legend": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "First Y-series legend type/name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceReportChartCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ReportChart resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectReportChart(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateReportChart(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportChart")
	}

	return resourceReportChartRead(ctx, d, meta)
}

func resourceReportChartUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportChart(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateReportChart(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ReportChart resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportChart")
	}

	return resourceReportChartRead(ctx, d, meta)
}

func resourceReportChartDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteReportChart(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ReportChart resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportChartRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportChart(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportChart resource: %v", err)
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

	diags := refreshObjectReportChart(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenReportChartCategorySeries(d *schema.ResourceData, v *[]models.ReportChartCategorySeries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Databind; tmp != nil {
				v["databind"] = *tmp
			}

			if tmp := cfg.FontSize; tmp != nil {
				v["font_size"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenReportChartColumn(d *schema.ResourceData, v *[]models.ReportChartColumn, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DetailUnit; tmp != nil {
				v["detail_unit"] = *tmp
			}

			if tmp := cfg.DetailValue; tmp != nil {
				v["detail_value"] = *tmp
			}

			if tmp := cfg.FooterUnit; tmp != nil {
				v["footer_unit"] = *tmp
			}

			if tmp := cfg.FooterValue; tmp != nil {
				v["footer_value"] = *tmp
			}

			if tmp := cfg.HeaderValue; tmp != nil {
				v["header_value"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Mapping; tmp != nil {
				v["mapping"] = flattenReportChartColumnMapping(d, tmp, prefix+"mapping", sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenReportChartColumnMapping(d *schema.ResourceData, v *[]models.ReportChartColumnMapping, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Displayname; tmp != nil {
				v["displayname"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Op; tmp != nil {
				v["op"] = *tmp
			}

			if tmp := cfg.ValueType; tmp != nil {
				v["value_type"] = *tmp
			}

			if tmp := cfg.Value1; tmp != nil {
				v["value1"] = *tmp
			}

			if tmp := cfg.Value2; tmp != nil {
				v["value2"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenReportChartDrillDownCharts(d *schema.ResourceData, v *[]models.ReportChartDrillDownCharts, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ChartName; tmp != nil {
				v["chart_name"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenReportChartValueSeries(d *schema.ResourceData, v *[]models.ReportChartValueSeries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Databind; tmp != nil {
				v["databind"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenReportChartXSeries(d *schema.ResourceData, v *[]models.ReportChartXSeries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Caption; tmp != nil {
				v["caption"] = *tmp
			}

			if tmp := cfg.CaptionFontSize; tmp != nil {
				v["caption_font_size"] = *tmp
			}

			if tmp := cfg.Databind; tmp != nil {
				v["databind"] = *tmp
			}

			if tmp := cfg.FontSize; tmp != nil {
				v["font_size"] = *tmp
			}

			if tmp := cfg.IsCategory; tmp != nil {
				v["is_category"] = *tmp
			}

			if tmp := cfg.LabelAngle; tmp != nil {
				v["label_angle"] = *tmp
			}

			if tmp := cfg.ScaleDirection; tmp != nil {
				v["scale_direction"] = *tmp
			}

			if tmp := cfg.ScaleFormat; tmp != nil {
				v["scale_format"] = *tmp
			}

			if tmp := cfg.ScaleStep; tmp != nil {
				v["scale_step"] = *tmp
			}

			if tmp := cfg.ScaleUnit; tmp != nil {
				v["scale_unit"] = *tmp
			}

			if tmp := cfg.Unit; tmp != nil {
				v["unit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenReportChartYSeries(d *schema.ResourceData, v *[]models.ReportChartYSeries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Caption; tmp != nil {
				v["caption"] = *tmp
			}

			if tmp := cfg.CaptionFontSize; tmp != nil {
				v["caption_font_size"] = *tmp
			}

			if tmp := cfg.Databind; tmp != nil {
				v["databind"] = *tmp
			}

			if tmp := cfg.ExtraDatabind; tmp != nil {
				v["extra_databind"] = *tmp
			}

			if tmp := cfg.ExtraY; tmp != nil {
				v["extra_y"] = *tmp
			}

			if tmp := cfg.ExtraYLegend; tmp != nil {
				v["extra_y_legend"] = *tmp
			}

			if tmp := cfg.FontSize; tmp != nil {
				v["font_size"] = *tmp
			}

			if tmp := cfg.Group; tmp != nil {
				v["group"] = *tmp
			}

			if tmp := cfg.LabelAngle; tmp != nil {
				v["label_angle"] = *tmp
			}

			if tmp := cfg.Unit; tmp != nil {
				v["unit"] = *tmp
			}

			if tmp := cfg.YLegend; tmp != nil {
				v["y_legend"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectReportChart(d *schema.ResourceData, o *models.ReportChart, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Background != nil {
		v := *o.Background

		if err = d.Set("background", v); err != nil {
			return diag.Errorf("error reading background: %v", err)
		}
	}

	if o.Category != nil {
		v := *o.Category

		if err = d.Set("category", v); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.CategorySeries != nil {
		if err = d.Set("category_series", flattenReportChartCategorySeries(d, o.CategorySeries, "category_series", sort)); err != nil {
			return diag.Errorf("error reading category_series: %v", err)
		}
	}

	if o.ColorPalette != nil {
		v := *o.ColorPalette

		if err = d.Set("color_palette", v); err != nil {
			return diag.Errorf("error reading color_palette: %v", err)
		}
	}

	if o.Column != nil {
		if err = d.Set("column", flattenReportChartColumn(d, o.Column, "column", sort)); err != nil {
			return diag.Errorf("error reading column: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Dataset != nil {
		v := *o.Dataset

		if err = d.Set("dataset", v); err != nil {
			return diag.Errorf("error reading dataset: %v", err)
		}
	}

	if o.Dimension != nil {
		v := *o.Dimension

		if err = d.Set("dimension", v); err != nil {
			return diag.Errorf("error reading dimension: %v", err)
		}
	}

	if o.DrillDownCharts != nil {
		if err = d.Set("drill_down_charts", flattenReportChartDrillDownCharts(d, o.DrillDownCharts, "drill_down_charts", sort)); err != nil {
			return diag.Errorf("error reading drill_down_charts: %v", err)
		}
	}

	if o.Favorite != nil {
		v := *o.Favorite

		if err = d.Set("favorite", v); err != nil {
			return diag.Errorf("error reading favorite: %v", err)
		}
	}

	if o.GraphType != nil {
		v := *o.GraphType

		if err = d.Set("graph_type", v); err != nil {
			return diag.Errorf("error reading graph_type: %v", err)
		}
	}

	if o.Legend != nil {
		v := *o.Legend

		if err = d.Set("legend", v); err != nil {
			return diag.Errorf("error reading legend: %v", err)
		}
	}

	if o.LegendFontSize != nil {
		v := *o.LegendFontSize

		if err = d.Set("legend_font_size", v); err != nil {
			return diag.Errorf("error reading legend_font_size: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Period != nil {
		v := *o.Period

		if err = d.Set("period", v); err != nil {
			return diag.Errorf("error reading period: %v", err)
		}
	}

	if o.Policy != nil {
		v := *o.Policy

		if err = d.Set("policy", v); err != nil {
			return diag.Errorf("error reading policy: %v", err)
		}
	}

	if o.Style != nil {
		v := *o.Style

		if err = d.Set("style", v); err != nil {
			return diag.Errorf("error reading style: %v", err)
		}
	}

	if o.Title != nil {
		v := *o.Title

		if err = d.Set("title", v); err != nil {
			return diag.Errorf("error reading title: %v", err)
		}
	}

	if o.TitleFontSize != nil {
		v := *o.TitleFontSize

		if err = d.Set("title_font_size", v); err != nil {
			return diag.Errorf("error reading title_font_size: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.ValueSeries != nil {
		if err = d.Set("value_series", flattenReportChartValueSeries(d, o.ValueSeries, "value_series", sort)); err != nil {
			return diag.Errorf("error reading value_series: %v", err)
		}
	}

	if o.XSeries != nil {
		if err = d.Set("x_series", flattenReportChartXSeries(d, o.XSeries, "x_series", sort)); err != nil {
			return diag.Errorf("error reading x_series: %v", err)
		}
	}

	if o.YSeries != nil {
		if err = d.Set("y_series", flattenReportChartYSeries(d, o.YSeries, "y_series", sort)); err != nil {
			return diag.Errorf("error reading y_series: %v", err)
		}
	}

	return nil
}

func expandReportChartCategorySeries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartCategorySeries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartCategorySeries

	for i := range l {
		tmp := models.ReportChartCategorySeries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.databind", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Databind = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.font_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FontSize = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportChartColumn(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartColumn, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartColumn

	for i := range l {
		tmp := models.ReportChartColumn{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.detail_unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DetailUnit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.detail_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DetailValue = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.footer_unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FooterUnit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.footer_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FooterValue = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HeaderValue = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mapping", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportChartColumnMapping(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportChartColumnMapping
			// 	}
			tmp.Mapping = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportChartColumnMapping(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartColumnMapping, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartColumnMapping

	for i := range l {
		tmp := models.ReportChartColumnMapping{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.displayname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Displayname = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.op", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Op = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ValueType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value1", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value1 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value2 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportChartDrillDownCharts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartDrillDownCharts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartDrillDownCharts

	for i := range l {
		tmp := models.ReportChartDrillDownCharts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chart_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChartName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportChartValueSeries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartValueSeries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartValueSeries

	for i := range l {
		tmp := models.ReportChartValueSeries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.databind", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Databind = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportChartXSeries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartXSeries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartXSeries

	for i := range l {
		tmp := models.ReportChartXSeries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.caption", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Caption = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.caption_font_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CaptionFontSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.databind", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Databind = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.font_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FontSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.is_category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IsCategory = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.label_angle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LabelAngle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scale_direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScaleDirection = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scale_format", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScaleFormat = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scale_step", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ScaleStep = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scale_unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScaleUnit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Unit = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportChartYSeries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportChartYSeries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportChartYSeries

	for i := range l {
		tmp := models.ReportChartYSeries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.caption", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Caption = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.caption_font_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CaptionFontSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.databind", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Databind = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.extra_databind", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExtraDatabind = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.extra_y", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExtraY = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.extra_y_legend", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExtraYLegend = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.font_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FontSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.group", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Group = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.label_angle", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LabelAngle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Unit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.y_legend", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.YLegend = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectReportChart(d *schema.ResourceData, sv string) (*models.ReportChart, diag.Diagnostics) {
	obj := models.ReportChart{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("background"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("background", sv)
				diags = append(diags, e)
			}
			obj.Background = &v2
		}
	}
	if v1, ok := d.GetOk("category"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("category", sv)
				diags = append(diags, e)
			}
			obj.Category = &v2
		}
	}
	if v, ok := d.GetOk("category_series"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("category_series", sv)
			diags = append(diags, e)
		}
		t, err := expandReportChartCategorySeries(d, v, "category_series", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CategorySeries = t
		}
	} else if d.HasChange("category_series") {
		old, new := d.GetChange("category_series")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CategorySeries = &[]models.ReportChartCategorySeries{}
		}
	}
	if v1, ok := d.GetOk("color_palette"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color_palette", sv)
				diags = append(diags, e)
			}
			obj.ColorPalette = &v2
		}
	}
	if v, ok := d.GetOk("column"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("column", sv)
			diags = append(diags, e)
		}
		t, err := expandReportChartColumn(d, v, "column", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Column = t
		}
	} else if d.HasChange("column") {
		old, new := d.GetChange("column")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Column = &[]models.ReportChartColumn{}
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("dataset"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dataset", sv)
				diags = append(diags, e)
			}
			obj.Dataset = &v2
		}
	}
	if v1, ok := d.GetOk("dimension"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dimension", sv)
				diags = append(diags, e)
			}
			obj.Dimension = &v2
		}
	}
	if v, ok := d.GetOk("drill_down_charts"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("drill_down_charts", sv)
			diags = append(diags, e)
		}
		t, err := expandReportChartDrillDownCharts(d, v, "drill_down_charts", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DrillDownCharts = t
		}
	} else if d.HasChange("drill_down_charts") {
		old, new := d.GetChange("drill_down_charts")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DrillDownCharts = &[]models.ReportChartDrillDownCharts{}
		}
	}
	if v1, ok := d.GetOk("favorite"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("favorite", sv)
				diags = append(diags, e)
			}
			obj.Favorite = &v2
		}
	}
	if v1, ok := d.GetOk("graph_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graph_type", sv)
				diags = append(diags, e)
			}
			obj.GraphType = &v2
		}
	}
	if v1, ok := d.GetOk("legend"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("legend", sv)
				diags = append(diags, e)
			}
			obj.Legend = &v2
		}
	}
	if v1, ok := d.GetOk("legend_font_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("legend_font_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LegendFontSize = &tmp
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
	if v1, ok := d.GetOk("period"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("period", sv)
				diags = append(diags, e)
			}
			obj.Period = &v2
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
	if v1, ok := d.GetOk("style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("style", sv)
				diags = append(diags, e)
			}
			obj.Style = &v2
		}
	}
	if v1, ok := d.GetOk("title"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("title", sv)
				diags = append(diags, e)
			}
			obj.Title = &v2
		}
	}
	if v1, ok := d.GetOk("title_font_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("title_font_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TitleFontSize = &tmp
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
	if v, ok := d.GetOk("value_series"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("value_series", sv)
			diags = append(diags, e)
		}
		t, err := expandReportChartValueSeries(d, v, "value_series", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ValueSeries = t
		}
	} else if d.HasChange("value_series") {
		old, new := d.GetChange("value_series")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ValueSeries = &[]models.ReportChartValueSeries{}
		}
	}
	if v, ok := d.GetOk("x_series"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("x_series", sv)
			diags = append(diags, e)
		}
		t, err := expandReportChartXSeries(d, v, "x_series", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.XSeries = t
		}
	} else if d.HasChange("x_series") {
		old, new := d.GetChange("x_series")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.XSeries = &[]models.ReportChartXSeries{}
		}
	}
	if v, ok := d.GetOk("y_series"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("y_series", sv)
			diags = append(diags, e)
		}
		t, err := expandReportChartYSeries(d, v, "y_series", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.YSeries = t
		}
	} else if d.HasChange("y_series") {
		old, new := d.GetChange("y_series")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.YSeries = &[]models.ReportChartYSeries{}
		}
	}
	return &obj, diags
}
