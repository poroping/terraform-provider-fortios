// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report layout configuration.

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

func resourceReportLayout() *schema.Resource {
	return &schema.Resource{
		Description: "Report layout configuration.",

		CreateContext: resourceReportLayoutCreate,
		ReadContext:   resourceReportLayoutRead,
		UpdateContext: resourceReportLayoutUpdate,
		DeleteContext: resourceReportLayoutDelete,

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
			"body_item": {
				Type:        schema.TypeList,
				Description: "Configure report body item.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chart": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Report item chart name.",
							Optional:    true,
							Computed:    true,
						},
						"chart_options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Report chart options.",
							Optional:         true,
							Computed:         true,
						},
						"column": {
							Type: schema.TypeInt,

							Description: "Report section column number.",
							Optional:    true,
							Computed:    true,
						},
						"content": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Report item text content.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"drill_down_items": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 11),

							Description: "Control how drill down charts are shown.",
							Optional:    true,
							Computed:    true,
						},
						"drill_down_types": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),

							Description: "Control whether keys from the parent being combined or not.",
							Optional:    true,
							Computed:    true,
						},
						"hide": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable hide item in report.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Report item ID.",
							Optional:    true,
							Computed:    true,
						},
						"img_src": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Report item image file name.",
							Optional:    true,
							Computed:    true,
						},
						"list": {
							Type:        schema.TypeList,
							Description: "Configure report list item.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),

										Description: "List entry content.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "List entry ID.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"list_component": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bullet", "numbered"}, false),

							Description: "Report item list component.",
							Optional:    true,
							Computed:    true,
						},
						"misc_component": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"hline", "page-break", "column-break", "section-start"}, false),

							Description: "Report item miscellaneous component.",
							Optional:    true,
							Computed:    true,
						},
						"parameters": {
							Type:        schema.TypeList,
							Description: "Parameters.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "ID.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),

										Description: "Field name that match field of parameters defined in dataset.",
										Optional:    true,
										Computed:    true,
									},
									"value": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 1023),

										Description: "Value to replace corresponding field of parameters defined in dataset.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"style": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Report item style.",
							Optional:    true,
							Computed:    true,
						},
						"table_caption_style": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Table chart caption style.",
							Optional:    true,
							Computed:    true,
						},
						"table_column_widths": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 179),

							Description: "Report item table column widths.",
							Optional:    true,
							Computed:    true,
						},
						"table_even_row_style": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Table chart even row style.",
							Optional:    true,
							Computed:    true,
						},
						"table_head_style": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Table chart head style.",
							Optional:    true,
							Computed:    true,
						},
						"table_odd_row_style": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),

							Description: "Table chart odd row style.",
							Optional:    true,
							Computed:    true,
						},
						"text_component": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"text", "heading1", "heading2", "heading3"}, false),

							Description: "Report item text component.",
							Optional:    true,
							Computed:    true,
						},
						"title": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Report section title.",
							Optional:    true,
							Computed:    true,
						},
						"top_n": {
							Type: schema.TypeInt,

							Description: "Value of top.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"text", "image", "chart", "misc"}, false),

							Description: "Report item type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"cutoff_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"run-time", "custom"}, false),

				Description: "Cutoff-option is either run-time or custom.",
				Optional:    true,
				Computed:    true,
			},
			"cutoff_time": {
				Type: schema.TypeString,

				Description: "Custom cutoff time to generate report (format = hh:mm).",
				Optional:    true,
				Computed:    true,
			},
			"day": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}, false),

				Description: "Schedule days of week to generate report.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"email_recipients": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Email recipients for generated reports.",
				Optional:    true,
				Computed:    true,
			},
			"email_send": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending emails after reports are generated.",
				Optional:    true,
				Computed:    true,
			},
			"format": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Report format.",
				Optional:         true,
				Computed:         true,
			},
			"max_pdf_report": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 365),

				Description: "Maximum number of PDF reports to keep at one time (oldest report is overwritten).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Report layout name.",
				ForceNew:    true,
				Required:    true,
			},
			"options": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Report layout options.",
				Optional:         true,
				Computed:         true,
			},
			"page": {
				Type:        schema.TypeList,
				Description: "Configure report page.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"column_break_before": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Report page auto column break before heading.",
							Optional:         true,
							Computed:         true,
						},
						"footer": {
							Type:        schema.TypeList,
							Description: "Configure report page footer.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"footer_item": {
										Type:        schema.TypeList,
										Description: "Configure report footer item.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"content": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 511),

													Description: "Report item text content.",
													Optional:    true,
													Computed:    true,
												},
												"description": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Description.",
													Optional:    true,
													Computed:    true,
												},
												"id": {
													Type: schema.TypeInt,

													Description: "Report item ID.",
													Optional:    true,
													Computed:    true,
												},
												"img_src": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 127),

													Description: "Report item image file name.",
													Optional:    true,
													Computed:    true,
												},
												"style": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 71),

													Description: "Report item style.",
													Optional:    true,
													Computed:    true,
												},
												"type": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"text", "image"}, false),

													Description: "Report item type.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"style": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 71),

										Description: "Report footer style.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"header": {
							Type:        schema.TypeList,
							Description: "Configure report page header.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"header_item": {
										Type:        schema.TypeList,
										Description: "Configure report header item.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"content": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 511),

													Description: "Report item text content.",
													Optional:    true,
													Computed:    true,
												},
												"description": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),

													Description: "Description.",
													Optional:    true,
													Computed:    true,
												},
												"id": {
													Type: schema.TypeInt,

													Description: "Report item ID.",
													Optional:    true,
													Computed:    true,
												},
												"img_src": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 127),

													Description: "Report item image file name.",
													Optional:    true,
													Computed:    true,
												},
												"style": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 71),

													Description: "Report item style.",
													Optional:    true,
													Computed:    true,
												},
												"type": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{"text", "image"}, false),

													Description: "Report item type.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"style": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 71),

										Description: "Report header style.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Report page options.",
							Optional:         true,
							Computed:         true,
						},
						"page_break_before": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Report page auto page break before heading.",
							Optional:         true,
							Computed:         true,
						},
						"paper": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"a4", "letter"}, false),

							Description: "Report page paper.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"schedule_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"demand", "daily", "weekly"}, false),

				Description: "Report schedule type.",
				Optional:    true,
				Computed:    true,
			},
			"style_theme": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Report style theme.",
				Optional:    true,
				Computed:    true,
			},
			"subtitle": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Report subtitle.",
				Optional:    true,
				Computed:    true,
			},
			"time": {
				Type: schema.TypeString,

				Description: "Schedule time to generate report (format = hh:mm).",
				Optional:    true,
				Computed:    true,
			},
			"title": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Report title.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceReportLayoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ReportLayout resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectReportLayout(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateReportLayout(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportLayout")
	}

	return resourceReportLayoutRead(ctx, d, meta)
}

func resourceReportLayoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportLayout(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateReportLayout(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ReportLayout resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportLayout")
	}

	return resourceReportLayoutRead(ctx, d, meta)
}

func resourceReportLayoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteReportLayout(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ReportLayout resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportLayout(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportLayout resource: %v", err)
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

	diags := refreshObjectReportLayout(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenReportLayoutBodyItem(v *[]models.ReportLayoutBodyItem, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Chart; tmp != nil {
				v["chart"] = *tmp
			}

			if tmp := cfg.ChartOptions; tmp != nil {
				v["chart_options"] = *tmp
			}

			if tmp := cfg.Column; tmp != nil {
				v["column"] = *tmp
			}

			if tmp := cfg.Content; tmp != nil {
				v["content"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.DrillDownItems; tmp != nil {
				v["drill_down_items"] = *tmp
			}

			if tmp := cfg.DrillDownTypes; tmp != nil {
				v["drill_down_types"] = *tmp
			}

			if tmp := cfg.Hide; tmp != nil {
				v["hide"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ImgSrc; tmp != nil {
				v["img_src"] = *tmp
			}

			if tmp := cfg.List; tmp != nil {
				v["list"] = flattenReportLayoutBodyItemList(tmp, sort)
			}

			if tmp := cfg.ListComponent; tmp != nil {
				v["list_component"] = *tmp
			}

			if tmp := cfg.MiscComponent; tmp != nil {
				v["misc_component"] = *tmp
			}

			if tmp := cfg.Parameters; tmp != nil {
				v["parameters"] = flattenReportLayoutBodyItemParameters(tmp, sort)
			}

			if tmp := cfg.Style; tmp != nil {
				v["style"] = *tmp
			}

			if tmp := cfg.TableCaptionStyle; tmp != nil {
				v["table_caption_style"] = *tmp
			}

			if tmp := cfg.TableColumnWidths; tmp != nil {
				v["table_column_widths"] = *tmp
			}

			if tmp := cfg.TableEvenRowStyle; tmp != nil {
				v["table_even_row_style"] = *tmp
			}

			if tmp := cfg.TableHeadStyle; tmp != nil {
				v["table_head_style"] = *tmp
			}

			if tmp := cfg.TableOddRowStyle; tmp != nil {
				v["table_odd_row_style"] = *tmp
			}

			if tmp := cfg.TextComponent; tmp != nil {
				v["text_component"] = *tmp
			}

			if tmp := cfg.Title; tmp != nil {
				v["title"] = *tmp
			}

			if tmp := cfg.TopN; tmp != nil {
				v["top_n"] = *tmp
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

func flattenReportLayoutBodyItemList(v *[]models.ReportLayoutBodyItemList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Content; tmp != nil {
				v["content"] = *tmp
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

func flattenReportLayoutBodyItemParameters(v *[]models.ReportLayoutBodyItemParameters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenReportLayoutPage(v *[]models.ReportLayoutPage, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ColumnBreakBefore; tmp != nil {
				v["column_break_before"] = *tmp
			}

			if tmp := cfg.Footer; tmp != nil {
				v["footer"] = flattenReportLayoutPageFooter(tmp, sort)
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = flattenReportLayoutPageHeader(tmp, sort)
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.PageBreakBefore; tmp != nil {
				v["page_break_before"] = *tmp
			}

			if tmp := cfg.Paper; tmp != nil {
				v["paper"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenReportLayoutPageFooter(v *[]models.ReportLayoutPageFooter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.FooterItem; tmp != nil {
				v["footer_item"] = flattenReportLayoutPageFooterFooterItem(tmp, sort)
			}

			if tmp := cfg.Style; tmp != nil {
				v["style"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenReportLayoutPageFooterFooterItem(v *[]models.ReportLayoutPageFooterFooterItem, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Content; tmp != nil {
				v["content"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ImgSrc; tmp != nil {
				v["img_src"] = *tmp
			}

			if tmp := cfg.Style; tmp != nil {
				v["style"] = *tmp
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

func flattenReportLayoutPageHeader(v *[]models.ReportLayoutPageHeader, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.HeaderItem; tmp != nil {
				v["header_item"] = flattenReportLayoutPageHeaderHeaderItem(tmp, sort)
			}

			if tmp := cfg.Style; tmp != nil {
				v["style"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenReportLayoutPageHeaderHeaderItem(v *[]models.ReportLayoutPageHeaderHeaderItem, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Content; tmp != nil {
				v["content"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ImgSrc; tmp != nil {
				v["img_src"] = *tmp
			}

			if tmp := cfg.Style; tmp != nil {
				v["style"] = *tmp
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

func refreshObjectReportLayout(d *schema.ResourceData, o *models.ReportLayout, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BodyItem != nil {
		if err = d.Set("body_item", flattenReportLayoutBodyItem(o.BodyItem, sort)); err != nil {
			return diag.Errorf("error reading body_item: %v", err)
		}
	}

	if o.CutoffOption != nil {
		v := *o.CutoffOption

		if err = d.Set("cutoff_option", v); err != nil {
			return diag.Errorf("error reading cutoff_option: %v", err)
		}
	}

	if o.CutoffTime != nil {
		v := *o.CutoffTime

		if err = d.Set("cutoff_time", v); err != nil {
			return diag.Errorf("error reading cutoff_time: %v", err)
		}
	}

	if o.Day != nil {
		v := *o.Day

		if err = d.Set("day", v); err != nil {
			return diag.Errorf("error reading day: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.EmailRecipients != nil {
		v := *o.EmailRecipients

		if err = d.Set("email_recipients", v); err != nil {
			return diag.Errorf("error reading email_recipients: %v", err)
		}
	}

	if o.EmailSend != nil {
		v := *o.EmailSend

		if err = d.Set("email_send", v); err != nil {
			return diag.Errorf("error reading email_send: %v", err)
		}
	}

	if o.Format != nil {
		v := *o.Format

		if err = d.Set("format", v); err != nil {
			return diag.Errorf("error reading format: %v", err)
		}
	}

	if o.MaxPdfReport != nil {
		v := *o.MaxPdfReport

		if err = d.Set("max_pdf_report", v); err != nil {
			return diag.Errorf("error reading max_pdf_report: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Options != nil {
		v := *o.Options

		if err = d.Set("options", v); err != nil {
			return diag.Errorf("error reading options: %v", err)
		}
	}

	if o.Page != nil {
		if err = d.Set("page", flattenReportLayoutPage(o.Page, sort)); err != nil {
			return diag.Errorf("error reading page: %v", err)
		}
	}

	if o.ScheduleType != nil {
		v := *o.ScheduleType

		if err = d.Set("schedule_type", v); err != nil {
			return diag.Errorf("error reading schedule_type: %v", err)
		}
	}

	if o.StyleTheme != nil {
		v := *o.StyleTheme

		if err = d.Set("style_theme", v); err != nil {
			return diag.Errorf("error reading style_theme: %v", err)
		}
	}

	if o.Subtitle != nil {
		v := *o.Subtitle

		if err = d.Set("subtitle", v); err != nil {
			return diag.Errorf("error reading subtitle: %v", err)
		}
	}

	if o.Time != nil {
		v := *o.Time

		if err = d.Set("time", v); err != nil {
			return diag.Errorf("error reading time: %v", err)
		}
	}

	if o.Title != nil {
		v := *o.Title

		if err = d.Set("title", v); err != nil {
			return diag.Errorf("error reading title: %v", err)
		}
	}

	return nil
}

func expandReportLayoutBodyItem(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutBodyItem, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutBodyItem

	for i := range l {
		tmp := models.ReportLayoutBodyItem{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chart", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Chart = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.chart_options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChartOptions = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.column", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Column = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.content", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Content = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drill_down_items", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrillDownItems = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.drill_down_types", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DrillDownTypes = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hide", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Hide = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.img_src", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImgSrc = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.list", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportLayoutBodyItemList(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportLayoutBodyItemList
			// 	}
			tmp.List = v2

		}

		pre_append = fmt.Sprintf("%s.%d.list_component", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ListComponent = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.misc_component", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MiscComponent = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parameters", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportLayoutBodyItemParameters(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportLayoutBodyItemParameters
			// 	}
			tmp.Parameters = v2

		}

		pre_append = fmt.Sprintf("%s.%d.style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Style = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.table_caption_style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TableCaptionStyle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.table_column_widths", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TableColumnWidths = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.table_even_row_style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TableEvenRowStyle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.table_head_style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TableHeadStyle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.table_odd_row_style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TableOddRowStyle = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.text_component", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TextComponent = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.title", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Title = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.top_n", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.TopN = &v2
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

func expandReportLayoutBodyItemList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutBodyItemList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutBodyItemList

	for i := range l {
		tmp := models.ReportLayoutBodyItemList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.content", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Content = &v2
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

func expandReportLayoutBodyItemParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutBodyItemParameters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutBodyItemParameters

	for i := range l {
		tmp := models.ReportLayoutBodyItemParameters{}
		var pre_append string

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

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportLayoutPage(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutPage, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutPage

	for i := range l {
		tmp := models.ReportLayoutPage{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.column_break_before", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ColumnBreakBefore = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.footer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportLayoutPageFooter(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportLayoutPageFooter
			// 	}
			tmp.Footer = v2

		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportLayoutPageHeader(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportLayoutPageHeader
			// 	}
			tmp.Header = v2

		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.page_break_before", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PageBreakBefore = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.paper", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Paper = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportLayoutPageFooter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutPageFooter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutPageFooter

	for i := range l {
		tmp := models.ReportLayoutPageFooter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.footer_item", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportLayoutPageFooterFooterItem(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportLayoutPageFooterFooterItem
			// 	}
			tmp.FooterItem = v2

		}

		pre_append = fmt.Sprintf("%s.%d.style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Style = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportLayoutPageFooterFooterItem(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutPageFooterFooterItem, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutPageFooterFooterItem

	for i := range l {
		tmp := models.ReportLayoutPageFooterFooterItem{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.content", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Content = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.img_src", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImgSrc = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Style = &v2
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

func expandReportLayoutPageHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutPageHeader, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutPageHeader

	for i := range l {
		tmp := models.ReportLayoutPageHeader{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.header_item", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandReportLayoutPageHeaderHeaderItem(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.ReportLayoutPageHeaderHeaderItem
			// 	}
			tmp.HeaderItem = v2

		}

		pre_append = fmt.Sprintf("%s.%d.style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Style = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandReportLayoutPageHeaderHeaderItem(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.ReportLayoutPageHeaderHeaderItem, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.ReportLayoutPageHeaderHeaderItem

	for i := range l {
		tmp := models.ReportLayoutPageHeaderHeaderItem{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.content", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Content = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.img_src", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImgSrc = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.style", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Style = &v2
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

func getObjectReportLayout(d *schema.ResourceData, sv string) (*models.ReportLayout, diag.Diagnostics) {
	obj := models.ReportLayout{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("body_item"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("body_item", sv)
			diags = append(diags, e)
		}
		t, err := expandReportLayoutBodyItem(d, v, "body_item", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.BodyItem = t
		}
	} else if d.HasChange("body_item") {
		old, new := d.GetChange("body_item")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.BodyItem = &[]models.ReportLayoutBodyItem{}
		}
	}
	if v1, ok := d.GetOk("cutoff_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cutoff_option", sv)
				diags = append(diags, e)
			}
			obj.CutoffOption = &v2
		}
	}
	if v1, ok := d.GetOk("cutoff_time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cutoff_time", sv)
				diags = append(diags, e)
			}
			obj.CutoffTime = &v2
		}
	}
	if v1, ok := d.GetOk("day"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("day", sv)
				diags = append(diags, e)
			}
			obj.Day = &v2
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("email_recipients"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_recipients", sv)
				diags = append(diags, e)
			}
			obj.EmailRecipients = &v2
		}
	}
	if v1, ok := d.GetOk("email_send"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_send", sv)
				diags = append(diags, e)
			}
			obj.EmailSend = &v2
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
	if v1, ok := d.GetOk("max_pdf_report"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_pdf_report", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxPdfReport = &tmp
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
	if v1, ok := d.GetOk("options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("options", sv)
				diags = append(diags, e)
			}
			obj.Options = &v2
		}
	}
	if v, ok := d.GetOk("page"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("page", sv)
			diags = append(diags, e)
		}
		t, err := expandReportLayoutPage(d, v, "page", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Page = t
		}
	} else if d.HasChange("page") {
		old, new := d.GetChange("page")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Page = &[]models.ReportLayoutPage{}
		}
	}
	if v1, ok := d.GetOk("schedule_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule_type", sv)
				diags = append(diags, e)
			}
			obj.ScheduleType = &v2
		}
	}
	if v1, ok := d.GetOk("style_theme"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("style_theme", sv)
				diags = append(diags, e)
			}
			obj.StyleTheme = &v2
		}
	}
	if v1, ok := d.GetOk("subtitle"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subtitle", sv)
				diags = append(diags, e)
			}
			obj.Subtitle = &v2
		}
	}
	if v1, ok := d.GetOk("time"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("time", sv)
				diags = append(diags, e)
			}
			obj.Time = &v2
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
	return &obj, diags
}
