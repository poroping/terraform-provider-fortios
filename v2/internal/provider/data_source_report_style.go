// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report style configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceReportStyle() *schema.Resource {
	return &schema.Resource{
		Description: "Report style configuration.",

		ReadContext: dataSourceReportStyleRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"align": {
				Type:        schema.TypeString,
				Description: "Alignment.",
				Computed:    true,
			},
			"bg_color": {
				Type:        schema.TypeString,
				Description: "Background color.",
				Computed:    true,
			},
			"border_bottom": {
				Type:        schema.TypeString,
				Description: "Border bottom.",
				Computed:    true,
			},
			"border_left": {
				Type:        schema.TypeString,
				Description: "Border left.",
				Computed:    true,
			},
			"border_right": {
				Type:        schema.TypeString,
				Description: "Border right.",
				Computed:    true,
			},
			"border_top": {
				Type:        schema.TypeString,
				Description: "Border top.",
				Computed:    true,
			},
			"column_gap": {
				Type:        schema.TypeString,
				Description: "Column gap.",
				Computed:    true,
			},
			"column_span": {
				Type:        schema.TypeString,
				Description: "Column span.",
				Computed:    true,
			},
			"fg_color": {
				Type:        schema.TypeString,
				Description: "Foreground color.",
				Computed:    true,
			},
			"font_family": {
				Type:        schema.TypeString,
				Description: "Font family.",
				Computed:    true,
			},
			"font_size": {
				Type:        schema.TypeString,
				Description: "Font size.",
				Computed:    true,
			},
			"font_style": {
				Type:        schema.TypeString,
				Description: "Font style.",
				Computed:    true,
			},
			"font_weight": {
				Type:        schema.TypeString,
				Description: "Font weight.",
				Computed:    true,
			},
			"height": {
				Type:        schema.TypeString,
				Description: "Height.",
				Computed:    true,
			},
			"line_height": {
				Type:        schema.TypeString,
				Description: "Text line height.",
				Computed:    true,
			},
			"margin_bottom": {
				Type:        schema.TypeString,
				Description: "Margin bottom.",
				Computed:    true,
			},
			"margin_left": {
				Type:        schema.TypeString,
				Description: "Margin left.",
				Computed:    true,
			},
			"margin_right": {
				Type:        schema.TypeString,
				Description: "Margin right.",
				Computed:    true,
			},
			"margin_top": {
				Type:        schema.TypeString,
				Description: "Margin top.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Report style name.",
				Required:    true,
			},
			"options": {
				Type:        schema.TypeString,
				Description: "Report style options.",
				Computed:    true,
			},
			"padding_bottom": {
				Type:        schema.TypeString,
				Description: "Padding bottom.",
				Computed:    true,
			},
			"padding_left": {
				Type:        schema.TypeString,
				Description: "Padding left.",
				Computed:    true,
			},
			"padding_right": {
				Type:        schema.TypeString,
				Description: "Padding right.",
				Computed:    true,
			},
			"padding_top": {
				Type:        schema.TypeString,
				Description: "Padding top.",
				Computed:    true,
			},
			"width": {
				Type:        schema.TypeString,
				Description: "Width.",
				Computed:    true,
			},
		},
	}
}

func dataSourceReportStyleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadReportStyle(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportStyle dataSource: %v", err)
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

	diags := refreshObjectReportStyle(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
