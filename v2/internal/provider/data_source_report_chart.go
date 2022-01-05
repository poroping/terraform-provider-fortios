// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report chart widget configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceReportChart() *schema.Resource {
	return &schema.Resource{
		Description: "Report chart widget configuration.",

		ReadContext: dataSourceReportChartRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"background": {
				Type:        schema.TypeString,
				Description: "Chart background.",
				Computed:    true,
			},
			"category": {
				Type:        schema.TypeString,
				Description: "Category.",
				Computed:    true,
			},
			"category_series": {
				Type:        schema.TypeList,
				Description: "Category series of pie chart.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": {
							Type:        schema.TypeString,
							Description: "Category series value expression.",
							Computed:    true,
						},
						"font_size": {
							Type:        schema.TypeInt,
							Description: "Font size of category-series title.",
							Computed:    true,
						},
					},
				},
			},
			"color_palette": {
				Type:        schema.TypeString,
				Description: "Color palette (system will pick color automatically by default).",
				Computed:    true,
			},
			"column": {
				Type:        schema.TypeList,
				Description: "Table column definition.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"detail_unit": {
							Type:        schema.TypeString,
							Description: "Detail unit of column.",
							Computed:    true,
						},
						"detail_value": {
							Type:        schema.TypeString,
							Description: "Detail value of column.",
							Computed:    true,
						},
						"footer_unit": {
							Type:        schema.TypeString,
							Description: "Footer unit of column.",
							Computed:    true,
						},
						"footer_value": {
							Type:        schema.TypeString,
							Description: "Footer value of column.",
							Computed:    true,
						},
						"header_value": {
							Type:        schema.TypeString,
							Description: "Display name of table header.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"mapping": {
							Type:        schema.TypeList,
							Description: "Show detail in certain display value for certain condition.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"displayname": {
										Type:        schema.TypeString,
										Description: "Display name.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "id",
										Computed:    true,
									},
									"op": {
										Type:        schema.TypeString,
										Description: "Comparision operater.",
										Computed:    true,
									},
									"value_type": {
										Type:        schema.TypeString,
										Description: "Value type.",
										Computed:    true,
									},
									"value1": {
										Type:        schema.TypeString,
										Description: "Value 1.",
										Computed:    true,
									},
									"value2": {
										Type:        schema.TypeString,
										Description: "Value 2.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dataset": {
				Type:        schema.TypeString,
				Description: "Bind dataset to chart.",
				Computed:    true,
			},
			"dimension": {
				Type:        schema.TypeString,
				Description: "Dimension.",
				Computed:    true,
			},
			"drill_down_charts": {
				Type:        schema.TypeList,
				Description: "Drill down charts.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chart_name": {
							Type:        schema.TypeString,
							Description: "Drill down chart name.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Drill down chart ID.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this drill down chart.",
							Computed:    true,
						},
					},
				},
			},
			"favorite": {
				Type:        schema.TypeString,
				Description: "Favorite.",
				Computed:    true,
			},
			"graph_type": {
				Type:        schema.TypeString,
				Description: "Graph type.",
				Computed:    true,
			},
			"legend": {
				Type:        schema.TypeString,
				Description: "Enable/Disable Legend area.",
				Computed:    true,
			},
			"legend_font_size": {
				Type:        schema.TypeInt,
				Description: "Font size of legend area.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Chart Widget Name",
				Required:    true,
			},
			"period": {
				Type:        schema.TypeString,
				Description: "Time period.",
				Computed:    true,
			},
			"policy": {
				Type:        schema.TypeInt,
				Description: "Used by monitor policy.",
				Computed:    true,
			},
			"style": {
				Type:        schema.TypeString,
				Description: "Style.",
				Computed:    true,
			},
			"title": {
				Type:        schema.TypeString,
				Description: "Chart title.",
				Computed:    true,
			},
			"title_font_size": {
				Type:        schema.TypeInt,
				Description: "Font size of chart title.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Chart type.",
				Computed:    true,
			},
			"value_series": {
				Type:        schema.TypeList,
				Description: "Value series of pie chart.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": {
							Type:        schema.TypeString,
							Description: "Value series value expression.",
							Computed:    true,
						},
					},
				},
			},
			"x_series": {
				Type:        schema.TypeList,
				Description: "X-series of chart.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caption": {
							Type:        schema.TypeString,
							Description: "X-series caption.",
							Computed:    true,
						},
						"caption_font_size": {
							Type:        schema.TypeInt,
							Description: "X-series caption font size.",
							Computed:    true,
						},
						"databind": {
							Type:        schema.TypeString,
							Description: "X-series value expression.",
							Computed:    true,
						},
						"font_size": {
							Type:        schema.TypeInt,
							Description: "X-series label font size.",
							Computed:    true,
						},
						"is_category": {
							Type:        schema.TypeString,
							Description: "X-series represent category or not.",
							Computed:    true,
						},
						"label_angle": {
							Type:        schema.TypeString,
							Description: "X-series label angle.",
							Computed:    true,
						},
						"scale_direction": {
							Type:        schema.TypeString,
							Description: "Scale increase or decrease.",
							Computed:    true,
						},
						"scale_format": {
							Type:        schema.TypeString,
							Description: "Date/time format.",
							Computed:    true,
						},
						"scale_step": {
							Type:        schema.TypeInt,
							Description: "Scale step.",
							Computed:    true,
						},
						"scale_unit": {
							Type:        schema.TypeString,
							Description: "Scale unit.",
							Computed:    true,
						},
						"unit": {
							Type:        schema.TypeString,
							Description: "X-series unit.",
							Computed:    true,
						},
					},
				},
			},
			"y_series": {
				Type:        schema.TypeList,
				Description: "Y-series of chart.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caption": {
							Type:        schema.TypeString,
							Description: "Y-series caption.",
							Computed:    true,
						},
						"caption_font_size": {
							Type:        schema.TypeInt,
							Description: "Y-series caption font size.",
							Computed:    true,
						},
						"databind": {
							Type:        schema.TypeString,
							Description: "Y-series value expression.",
							Computed:    true,
						},
						"extra_databind": {
							Type:        schema.TypeString,
							Description: "Extra Y-series value.",
							Computed:    true,
						},
						"extra_y": {
							Type:        schema.TypeString,
							Description: "Allow another Y-series value",
							Computed:    true,
						},
						"extra_y_legend": {
							Type:        schema.TypeString,
							Description: "Extra Y-series legend type/name.",
							Computed:    true,
						},
						"font_size": {
							Type:        schema.TypeInt,
							Description: "Y-series label font size.",
							Computed:    true,
						},
						"group": {
							Type:        schema.TypeString,
							Description: "Y-series group option.",
							Computed:    true,
						},
						"label_angle": {
							Type:        schema.TypeString,
							Description: "Y-series label angle.",
							Computed:    true,
						},
						"unit": {
							Type:        schema.TypeString,
							Description: "Y-series unit.",
							Computed:    true,
						},
						"y_legend": {
							Type:        schema.TypeString,
							Description: "First Y-series legend type/name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceReportChartRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportChart(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportChart dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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
