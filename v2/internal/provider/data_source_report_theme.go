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
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceReportTheme() *schema.Resource {
	return &schema.Resource{
		Description: "Report themes configuration",

		ReadContext: dataSourceReportThemeRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bullet_list_style": {
				Type:        schema.TypeString,
				Description: "Bullet list style.",
				Computed:    true,
			},
			"column_count": {
				Type:        schema.TypeString,
				Description: "Report page column count.",
				Computed:    true,
			},
			"default_html_style": {
				Type:        schema.TypeString,
				Description: "Default HTML report style.",
				Computed:    true,
			},
			"default_pdf_style": {
				Type:        schema.TypeString,
				Description: "Default PDF report style.",
				Computed:    true,
			},
			"graph_chart_style": {
				Type:        schema.TypeString,
				Description: "Graph chart style.",
				Computed:    true,
			},
			"heading1_style": {
				Type:        schema.TypeString,
				Description: "Report heading style.",
				Computed:    true,
			},
			"heading2_style": {
				Type:        schema.TypeString,
				Description: "Report heading style.",
				Computed:    true,
			},
			"heading3_style": {
				Type:        schema.TypeString,
				Description: "Report heading style.",
				Computed:    true,
			},
			"heading4_style": {
				Type:        schema.TypeString,
				Description: "Report heading style.",
				Computed:    true,
			},
			"hline_style": {
				Type:        schema.TypeString,
				Description: "Horizontal line style.",
				Computed:    true,
			},
			"image_style": {
				Type:        schema.TypeString,
				Description: "Image style.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Report theme name.",
				Required:    true,
			},
			"normal_text_style": {
				Type:        schema.TypeString,
				Description: "Normal text style.",
				Computed:    true,
			},
			"numbered_list_style": {
				Type:        schema.TypeString,
				Description: "Numbered list style.",
				Computed:    true,
			},
			"page_footer_style": {
				Type:        schema.TypeString,
				Description: "Report page footer style.",
				Computed:    true,
			},
			"page_header_style": {
				Type:        schema.TypeString,
				Description: "Report page header style.",
				Computed:    true,
			},
			"page_orient": {
				Type:        schema.TypeString,
				Description: "Report page orientation.",
				Computed:    true,
			},
			"page_style": {
				Type:        schema.TypeString,
				Description: "Report page style.",
				Computed:    true,
			},
			"report_subtitle_style": {
				Type:        schema.TypeString,
				Description: "Report subtitle style.",
				Computed:    true,
			},
			"report_title_style": {
				Type:        schema.TypeString,
				Description: "Report title style.",
				Computed:    true,
			},
			"table_chart_caption_style": {
				Type:        schema.TypeString,
				Description: "Table chart caption style.",
				Computed:    true,
			},
			"table_chart_even_row_style": {
				Type:        schema.TypeString,
				Description: "Table chart even row style.",
				Computed:    true,
			},
			"table_chart_head_style": {
				Type:        schema.TypeString,
				Description: "Table chart head row style.",
				Computed:    true,
			},
			"table_chart_odd_row_style": {
				Type:        schema.TypeString,
				Description: "Table chart odd row style.",
				Computed:    true,
			},
			"table_chart_style": {
				Type:        schema.TypeString,
				Description: "Table chart style.",
				Computed:    true,
			},
			"toc_heading1_style": {
				Type:        schema.TypeString,
				Description: "Table of contents heading style.",
				Computed:    true,
			},
			"toc_heading2_style": {
				Type:        schema.TypeString,
				Description: "Table of contents heading style.",
				Computed:    true,
			},
			"toc_heading3_style": {
				Type:        schema.TypeString,
				Description: "Table of contents heading style.",
				Computed:    true,
			},
			"toc_heading4_style": {
				Type:        schema.TypeString,
				Description: "Table of contents heading style.",
				Computed:    true,
			},
			"toc_title_style": {
				Type:        schema.TypeString,
				Description: "Table of contents title style.",
				Computed:    true,
			},
		},
	}
}

func dataSourceReportThemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.Errorf("error reading ReportTheme dataSource: %v", err)
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

	diags := refreshObjectReportTheme(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
