// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report themes configuration

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

func resourceReportTheme() *schema.Resource {
	return &schema.Resource{
		Description: "Report themes configuration",

		CreateContext: resourceReportThemeCreate,
		ReadContext:   resourceReportThemeRead,
		UpdateContext: resourceReportThemeUpdate,
		DeleteContext: resourceReportThemeDelete,

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
			"bullet_list_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Bullet list style.",
				Optional:    true,
				Computed:    true,
			},
			"column_count": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1", "2", "3"}, false),

				Description: "Report page column count.",
				Optional:    true,
				Computed:    true,
			},
			"default_html_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Default HTML report style.",
				Optional:    true,
				Computed:    true,
			},
			"default_pdf_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Default PDF report style.",
				Optional:    true,
				Computed:    true,
			},
			"graph_chart_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Graph chart style.",
				Optional:    true,
				Computed:    true,
			},
			"heading1_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report heading style.",
				Optional:    true,
				Computed:    true,
			},
			"heading2_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report heading style.",
				Optional:    true,
				Computed:    true,
			},
			"heading3_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report heading style.",
				Optional:    true,
				Computed:    true,
			},
			"heading4_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report heading style.",
				Optional:    true,
				Computed:    true,
			},
			"hline_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Horizontal line style.",
				Optional:    true,
				Computed:    true,
			},
			"image_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Image style.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Report theme name.",
				ForceNew:    true,
				Required:    true,
			},
			"normal_text_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Normal text style.",
				Optional:    true,
				Computed:    true,
			},
			"numbered_list_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Numbered list style.",
				Optional:    true,
				Computed:    true,
			},
			"page_footer_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report page footer style.",
				Optional:    true,
				Computed:    true,
			},
			"page_header_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report page header style.",
				Optional:    true,
				Computed:    true,
			},
			"page_orient": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"portrait", "landscape"}, false),

				Description: "Report page orientation.",
				Optional:    true,
				Computed:    true,
			},
			"page_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report page style.",
				Optional:    true,
				Computed:    true,
			},
			"report_subtitle_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report subtitle style.",
				Optional:    true,
				Computed:    true,
			},
			"report_title_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report title style.",
				Optional:    true,
				Computed:    true,
			},
			"table_chart_caption_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table chart caption style.",
				Optional:    true,
				Computed:    true,
			},
			"table_chart_even_row_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table chart even row style.",
				Optional:    true,
				Computed:    true,
			},
			"table_chart_head_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table chart head row style.",
				Optional:    true,
				Computed:    true,
			},
			"table_chart_odd_row_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table chart odd row style.",
				Optional:    true,
				Computed:    true,
			},
			"table_chart_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table chart style.",
				Optional:    true,
				Computed:    true,
			},
			"toc_heading1_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table of contents heading style.",
				Optional:    true,
				Computed:    true,
			},
			"toc_heading2_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table of contents heading style.",
				Optional:    true,
				Computed:    true,
			},
			"toc_heading3_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table of contents heading style.",
				Optional:    true,
				Computed:    true,
			},
			"toc_heading4_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table of contents heading style.",
				Optional:    true,
				Computed:    true,
			},
			"toc_title_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Table of contents title style.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceReportThemeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ReportTheme resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectReportTheme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateReportTheme(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportTheme")
	}

	return resourceReportThemeRead(ctx, d, meta)
}

func resourceReportThemeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportTheme(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateReportTheme(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ReportTheme resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportTheme")
	}

	return resourceReportThemeRead(ctx, d, meta)
}

func resourceReportThemeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteReportTheme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ReportTheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportThemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportTheme(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportTheme resource: %v", err)
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

	diags := refreshObjectReportTheme(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectReportTheme(d *schema.ResourceData, o *models.ReportTheme, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BulletListStyle != nil {
		v := *o.BulletListStyle

		if err = d.Set("bullet_list_style", v); err != nil {
			return diag.Errorf("error reading bullet_list_style: %v", err)
		}
	}

	if o.ColumnCount != nil {
		v := *o.ColumnCount

		if err = d.Set("column_count", v); err != nil {
			return diag.Errorf("error reading column_count: %v", err)
		}
	}

	if o.DefaultHtmlStyle != nil {
		v := *o.DefaultHtmlStyle

		if err = d.Set("default_html_style", v); err != nil {
			return diag.Errorf("error reading default_html_style: %v", err)
		}
	}

	if o.DefaultPdfStyle != nil {
		v := *o.DefaultPdfStyle

		if err = d.Set("default_pdf_style", v); err != nil {
			return diag.Errorf("error reading default_pdf_style: %v", err)
		}
	}

	if o.GraphChartStyle != nil {
		v := *o.GraphChartStyle

		if err = d.Set("graph_chart_style", v); err != nil {
			return diag.Errorf("error reading graph_chart_style: %v", err)
		}
	}

	if o.Heading1Style != nil {
		v := *o.Heading1Style

		if err = d.Set("heading1_style", v); err != nil {
			return diag.Errorf("error reading heading1_style: %v", err)
		}
	}

	if o.Heading2Style != nil {
		v := *o.Heading2Style

		if err = d.Set("heading2_style", v); err != nil {
			return diag.Errorf("error reading heading2_style: %v", err)
		}
	}

	if o.Heading3Style != nil {
		v := *o.Heading3Style

		if err = d.Set("heading3_style", v); err != nil {
			return diag.Errorf("error reading heading3_style: %v", err)
		}
	}

	if o.Heading4Style != nil {
		v := *o.Heading4Style

		if err = d.Set("heading4_style", v); err != nil {
			return diag.Errorf("error reading heading4_style: %v", err)
		}
	}

	if o.HlineStyle != nil {
		v := *o.HlineStyle

		if err = d.Set("hline_style", v); err != nil {
			return diag.Errorf("error reading hline_style: %v", err)
		}
	}

	if o.ImageStyle != nil {
		v := *o.ImageStyle

		if err = d.Set("image_style", v); err != nil {
			return diag.Errorf("error reading image_style: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NormalTextStyle != nil {
		v := *o.NormalTextStyle

		if err = d.Set("normal_text_style", v); err != nil {
			return diag.Errorf("error reading normal_text_style: %v", err)
		}
	}

	if o.NumberedListStyle != nil {
		v := *o.NumberedListStyle

		if err = d.Set("numbered_list_style", v); err != nil {
			return diag.Errorf("error reading numbered_list_style: %v", err)
		}
	}

	if o.PageFooterStyle != nil {
		v := *o.PageFooterStyle

		if err = d.Set("page_footer_style", v); err != nil {
			return diag.Errorf("error reading page_footer_style: %v", err)
		}
	}

	if o.PageHeaderStyle != nil {
		v := *o.PageHeaderStyle

		if err = d.Set("page_header_style", v); err != nil {
			return diag.Errorf("error reading page_header_style: %v", err)
		}
	}

	if o.PageOrient != nil {
		v := *o.PageOrient

		if err = d.Set("page_orient", v); err != nil {
			return diag.Errorf("error reading page_orient: %v", err)
		}
	}

	if o.PageStyle != nil {
		v := *o.PageStyle

		if err = d.Set("page_style", v); err != nil {
			return diag.Errorf("error reading page_style: %v", err)
		}
	}

	if o.ReportSubtitleStyle != nil {
		v := *o.ReportSubtitleStyle

		if err = d.Set("report_subtitle_style", v); err != nil {
			return diag.Errorf("error reading report_subtitle_style: %v", err)
		}
	}

	if o.ReportTitleStyle != nil {
		v := *o.ReportTitleStyle

		if err = d.Set("report_title_style", v); err != nil {
			return diag.Errorf("error reading report_title_style: %v", err)
		}
	}

	if o.TableChartCaptionStyle != nil {
		v := *o.TableChartCaptionStyle

		if err = d.Set("table_chart_caption_style", v); err != nil {
			return diag.Errorf("error reading table_chart_caption_style: %v", err)
		}
	}

	if o.TableChartEvenRowStyle != nil {
		v := *o.TableChartEvenRowStyle

		if err = d.Set("table_chart_even_row_style", v); err != nil {
			return diag.Errorf("error reading table_chart_even_row_style: %v", err)
		}
	}

	if o.TableChartHeadStyle != nil {
		v := *o.TableChartHeadStyle

		if err = d.Set("table_chart_head_style", v); err != nil {
			return diag.Errorf("error reading table_chart_head_style: %v", err)
		}
	}

	if o.TableChartOddRowStyle != nil {
		v := *o.TableChartOddRowStyle

		if err = d.Set("table_chart_odd_row_style", v); err != nil {
			return diag.Errorf("error reading table_chart_odd_row_style: %v", err)
		}
	}

	if o.TableChartStyle != nil {
		v := *o.TableChartStyle

		if err = d.Set("table_chart_style", v); err != nil {
			return diag.Errorf("error reading table_chart_style: %v", err)
		}
	}

	if o.TocHeading1Style != nil {
		v := *o.TocHeading1Style

		if err = d.Set("toc_heading1_style", v); err != nil {
			return diag.Errorf("error reading toc_heading1_style: %v", err)
		}
	}

	if o.TocHeading2Style != nil {
		v := *o.TocHeading2Style

		if err = d.Set("toc_heading2_style", v); err != nil {
			return diag.Errorf("error reading toc_heading2_style: %v", err)
		}
	}

	if o.TocHeading3Style != nil {
		v := *o.TocHeading3Style

		if err = d.Set("toc_heading3_style", v); err != nil {
			return diag.Errorf("error reading toc_heading3_style: %v", err)
		}
	}

	if o.TocHeading4Style != nil {
		v := *o.TocHeading4Style

		if err = d.Set("toc_heading4_style", v); err != nil {
			return diag.Errorf("error reading toc_heading4_style: %v", err)
		}
	}

	if o.TocTitleStyle != nil {
		v := *o.TocTitleStyle

		if err = d.Set("toc_title_style", v); err != nil {
			return diag.Errorf("error reading toc_title_style: %v", err)
		}
	}

	return nil
}

func getObjectReportTheme(d *schema.ResourceData, sv string) (*models.ReportTheme, diag.Diagnostics) {
	obj := models.ReportTheme{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("bullet_list_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bullet_list_style", sv)
				diags = append(diags, e)
			}
			obj.BulletListStyle = &v2
		}
	}
	if v1, ok := d.GetOk("column_count"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("column_count", sv)
				diags = append(diags, e)
			}
			obj.ColumnCount = &v2
		}
	}
	if v1, ok := d.GetOk("default_html_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_html_style", sv)
				diags = append(diags, e)
			}
			obj.DefaultHtmlStyle = &v2
		}
	}
	if v1, ok := d.GetOk("default_pdf_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_pdf_style", sv)
				diags = append(diags, e)
			}
			obj.DefaultPdfStyle = &v2
		}
	}
	if v1, ok := d.GetOk("graph_chart_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("graph_chart_style", sv)
				diags = append(diags, e)
			}
			obj.GraphChartStyle = &v2
		}
	}
	if v1, ok := d.GetOk("heading1_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("heading1_style", sv)
				diags = append(diags, e)
			}
			obj.Heading1Style = &v2
		}
	}
	if v1, ok := d.GetOk("heading2_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("heading2_style", sv)
				diags = append(diags, e)
			}
			obj.Heading2Style = &v2
		}
	}
	if v1, ok := d.GetOk("heading3_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("heading3_style", sv)
				diags = append(diags, e)
			}
			obj.Heading3Style = &v2
		}
	}
	if v1, ok := d.GetOk("heading4_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("heading4_style", sv)
				diags = append(diags, e)
			}
			obj.Heading4Style = &v2
		}
	}
	if v1, ok := d.GetOk("hline_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hline_style", sv)
				diags = append(diags, e)
			}
			obj.HlineStyle = &v2
		}
	}
	if v1, ok := d.GetOk("image_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("image_style", sv)
				diags = append(diags, e)
			}
			obj.ImageStyle = &v2
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
	if v1, ok := d.GetOk("normal_text_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("normal_text_style", sv)
				diags = append(diags, e)
			}
			obj.NormalTextStyle = &v2
		}
	}
	if v1, ok := d.GetOk("numbered_list_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("numbered_list_style", sv)
				diags = append(diags, e)
			}
			obj.NumberedListStyle = &v2
		}
	}
	if v1, ok := d.GetOk("page_footer_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("page_footer_style", sv)
				diags = append(diags, e)
			}
			obj.PageFooterStyle = &v2
		}
	}
	if v1, ok := d.GetOk("page_header_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("page_header_style", sv)
				diags = append(diags, e)
			}
			obj.PageHeaderStyle = &v2
		}
	}
	if v1, ok := d.GetOk("page_orient"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("page_orient", sv)
				diags = append(diags, e)
			}
			obj.PageOrient = &v2
		}
	}
	if v1, ok := d.GetOk("page_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("page_style", sv)
				diags = append(diags, e)
			}
			obj.PageStyle = &v2
		}
	}
	if v1, ok := d.GetOk("report_subtitle_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("report_subtitle_style", sv)
				diags = append(diags, e)
			}
			obj.ReportSubtitleStyle = &v2
		}
	}
	if v1, ok := d.GetOk("report_title_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("report_title_style", sv)
				diags = append(diags, e)
			}
			obj.ReportTitleStyle = &v2
		}
	}
	if v1, ok := d.GetOk("table_chart_caption_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("table_chart_caption_style", sv)
				diags = append(diags, e)
			}
			obj.TableChartCaptionStyle = &v2
		}
	}
	if v1, ok := d.GetOk("table_chart_even_row_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("table_chart_even_row_style", sv)
				diags = append(diags, e)
			}
			obj.TableChartEvenRowStyle = &v2
		}
	}
	if v1, ok := d.GetOk("table_chart_head_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("table_chart_head_style", sv)
				diags = append(diags, e)
			}
			obj.TableChartHeadStyle = &v2
		}
	}
	if v1, ok := d.GetOk("table_chart_odd_row_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("table_chart_odd_row_style", sv)
				diags = append(diags, e)
			}
			obj.TableChartOddRowStyle = &v2
		}
	}
	if v1, ok := d.GetOk("table_chart_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("table_chart_style", sv)
				diags = append(diags, e)
			}
			obj.TableChartStyle = &v2
		}
	}
	if v1, ok := d.GetOk("toc_heading1_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("toc_heading1_style", sv)
				diags = append(diags, e)
			}
			obj.TocHeading1Style = &v2
		}
	}
	if v1, ok := d.GetOk("toc_heading2_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("toc_heading2_style", sv)
				diags = append(diags, e)
			}
			obj.TocHeading2Style = &v2
		}
	}
	if v1, ok := d.GetOk("toc_heading3_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("toc_heading3_style", sv)
				diags = append(diags, e)
			}
			obj.TocHeading3Style = &v2
		}
	}
	if v1, ok := d.GetOk("toc_heading4_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("toc_heading4_style", sv)
				diags = append(diags, e)
			}
			obj.TocHeading4Style = &v2
		}
	}
	if v1, ok := d.GetOk("toc_title_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("toc_title_style", sv)
				diags = append(diags, e)
			}
			obj.TocTitleStyle = &v2
		}
	}
	return &obj, diags
}
