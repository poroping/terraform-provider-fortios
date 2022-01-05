// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report layout configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceReportLayout() *schema.Resource {
	return &schema.Resource{
		Description: "Report layout configuration.",

		ReadContext: dataSourceReportLayoutRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"body_item": {
				Type:        schema.TypeList,
				Description: "Configure report body item.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chart": {
							Type:        schema.TypeString,
							Description: "Report item chart name.",
							Computed:    true,
						},
						"chart_options": {
							Type:        schema.TypeString,
							Description: "Report chart options.",
							Computed:    true,
						},
						"column": {
							Type:        schema.TypeInt,
							Description: "Report section column number.",
							Computed:    true,
						},
						"content": {
							Type:        schema.TypeString,
							Description: "Report item text content.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"drill_down_items": {
							Type:        schema.TypeString,
							Description: "Control how drill down charts are shown.",
							Computed:    true,
						},
						"drill_down_types": {
							Type:        schema.TypeString,
							Description: "Control whether keys from the parent being combined or not.",
							Computed:    true,
						},
						"hide": {
							Type:        schema.TypeString,
							Description: "Enable/disable hide item in report.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Report item ID.",
							Computed:    true,
						},
						"img_src": {
							Type:        schema.TypeString,
							Description: "Report item image file name.",
							Computed:    true,
						},
						"list": {
							Type:        schema.TypeList,
							Description: "Configure report list item.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Description: "List entry content.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "List entry ID.",
										Computed:    true,
									},
								},
							},
						},
						"list_component": {
							Type:        schema.TypeString,
							Description: "Report item list component.",
							Computed:    true,
						},
						"misc_component": {
							Type:        schema.TypeString,
							Description: "Report item miscellaneous component.",
							Computed:    true,
						},
						"parameters": {
							Type:        schema.TypeList,
							Description: "Parameters.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "ID.",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Field name that match field of parameters defined in dataset.",
										Computed:    true,
									},
									"value": {
										Type:        schema.TypeString,
										Description: "Value to replace corresponding field of parameters defined in dataset.",
										Computed:    true,
									},
								},
							},
						},
						"style": {
							Type:        schema.TypeString,
							Description: "Report item style.",
							Computed:    true,
						},
						"table_caption_style": {
							Type:        schema.TypeString,
							Description: "Table chart caption style.",
							Computed:    true,
						},
						"table_column_widths": {
							Type:        schema.TypeString,
							Description: "Report item table column widths.",
							Computed:    true,
						},
						"table_even_row_style": {
							Type:        schema.TypeString,
							Description: "Table chart even row style.",
							Computed:    true,
						},
						"table_head_style": {
							Type:        schema.TypeString,
							Description: "Table chart head style.",
							Computed:    true,
						},
						"table_odd_row_style": {
							Type:        schema.TypeString,
							Description: "Table chart odd row style.",
							Computed:    true,
						},
						"text_component": {
							Type:        schema.TypeString,
							Description: "Report item text component.",
							Computed:    true,
						},
						"title": {
							Type:        schema.TypeString,
							Description: "Report section title.",
							Computed:    true,
						},
						"top_n": {
							Type:        schema.TypeInt,
							Description: "Value of top.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Report item type.",
							Computed:    true,
						},
					},
				},
			},
			"cutoff_option": {
				Type:        schema.TypeString,
				Description: "Cutoff-option is either run-time or custom.",
				Computed:    true,
			},
			"cutoff_time": {
				Type:        schema.TypeString,
				Description: "Custom cutoff time to generate report [hh:mm].",
				Computed:    true,
			},
			"day": {
				Type:        schema.TypeString,
				Description: "Schedule days of week to generate report.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"email_recipients": {
				Type:        schema.TypeString,
				Description: "Email recipients for generated reports.",
				Computed:    true,
			},
			"email_send": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending emails after reports are generated.",
				Computed:    true,
			},
			"format": {
				Type:        schema.TypeString,
				Description: "Report format.",
				Computed:    true,
			},
			"max_pdf_report": {
				Type:        schema.TypeInt,
				Description: "Maximum number of PDF reports to keep at one time (oldest report is overwritten).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Report layout name.",
				Required:    true,
			},
			"options": {
				Type:        schema.TypeString,
				Description: "Report layout options.",
				Computed:    true,
			},
			"page": {
				Type:        schema.TypeList,
				Description: "Configure report page.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"column_break_before": {
							Type:        schema.TypeString,
							Description: "Report page auto column break before heading.",
							Computed:    true,
						},
						"footer": {
							Type:        schema.TypeList,
							Description: "Configure report page footer.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"footer_item": {
										Type:        schema.TypeList,
										Description: "Configure report footer item.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"content": {
													Type:        schema.TypeString,
													Description: "Report item text content.",
													Computed:    true,
												},
												"description": {
													Type:        schema.TypeString,
													Description: "Description.",
													Computed:    true,
												},
												"id": {
													Type:        schema.TypeInt,
													Description: "Report item ID.",
													Computed:    true,
												},
												"img_src": {
													Type:        schema.TypeString,
													Description: "Report item image file name.",
													Computed:    true,
												},
												"style": {
													Type:        schema.TypeString,
													Description: "Report item style.",
													Computed:    true,
												},
												"type": {
													Type:        schema.TypeString,
													Description: "Report item type.",
													Computed:    true,
												},
											},
										},
									},
									"style": {
										Type:        schema.TypeString,
										Description: "Report footer style.",
										Computed:    true,
									},
								},
							},
						},
						"header": {
							Type:        schema.TypeList,
							Description: "Configure report page header.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"header_item": {
										Type:        schema.TypeList,
										Description: "Configure report header item.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"content": {
													Type:        schema.TypeString,
													Description: "Report item text content.",
													Computed:    true,
												},
												"description": {
													Type:        schema.TypeString,
													Description: "Description.",
													Computed:    true,
												},
												"id": {
													Type:        schema.TypeInt,
													Description: "Report item ID.",
													Computed:    true,
												},
												"img_src": {
													Type:        schema.TypeString,
													Description: "Report item image file name.",
													Computed:    true,
												},
												"style": {
													Type:        schema.TypeString,
													Description: "Report item style.",
													Computed:    true,
												},
												"type": {
													Type:        schema.TypeString,
													Description: "Report item type.",
													Computed:    true,
												},
											},
										},
									},
									"style": {
										Type:        schema.TypeString,
										Description: "Report header style.",
										Computed:    true,
									},
								},
							},
						},
						"options": {
							Type:        schema.TypeString,
							Description: "Report page options.",
							Computed:    true,
						},
						"page_break_before": {
							Type:        schema.TypeString,
							Description: "Report page auto page break before heading.",
							Computed:    true,
						},
						"paper": {
							Type:        schema.TypeString,
							Description: "Report page paper.",
							Computed:    true,
						},
					},
				},
			},
			"schedule_type": {
				Type:        schema.TypeString,
				Description: "Report schedule type.",
				Computed:    true,
			},
			"style_theme": {
				Type:        schema.TypeString,
				Description: "Report style theme.",
				Computed:    true,
			},
			"subtitle": {
				Type:        schema.TypeString,
				Description: "Report subtitle.",
				Computed:    true,
			},
			"time": {
				Type:        schema.TypeString,
				Description: "Schedule time to generate report [hh:mm].",
				Computed:    true,
			},
			"title": {
				Type:        schema.TypeString,
				Description: "Report title.",
				Computed:    true,
			},
		},
	}
}

func dataSourceReportLayoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportLayout(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportLayout dataSource: %v", err)
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

	diags := refreshObjectReportLayout(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
