// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report style configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceReportStyle() *schema.Resource {
	return &schema.Resource{
		Description: "Report style configuration.",

		CreateContext: resourceReportStyleCreate,
		ReadContext:   resourceReportStyleRead,
		UpdateContext: resourceReportStyleUpdate,
		DeleteContext: resourceReportStyleDelete,

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
			"align": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"left", "center", "right", "justify"}, false),

				Description: "Alignment.",
				Optional:    true,
				Computed:    true,
			},
			"bg_color": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Background color.",
				Optional:    true,
				Computed:    true,
			},
			"border_bottom": {
				Type: schema.TypeString,

				Description: "Border bottom.",
				Optional:    true,
				Computed:    true,
			},
			"border_left": {
				Type: schema.TypeString,

				Description: "Border left.",
				Optional:    true,
				Computed:    true,
			},
			"border_right": {
				Type: schema.TypeString,

				Description: "Border right.",
				Optional:    true,
				Computed:    true,
			},
			"border_top": {
				Type: schema.TypeString,

				Description: "Border top.",
				Optional:    true,
				Computed:    true,
			},
			"column_gap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Column gap.",
				Optional:    true,
				Computed:    true,
			},
			"column_span": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "all"}, false),

				Description: "Column span.",
				Optional:    true,
				Computed:    true,
			},
			"fg_color": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Foreground color.",
				Optional:    true,
				Computed:    true,
			},
			"font_family": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Verdana", "Arial", "Helvetica", "Courier", "Times"}, false),

				Description: "Font family.",
				Optional:    true,
				Computed:    true,
			},
			"font_size": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Font size.",
				Optional:    true,
				Computed:    true,
			},
			"font_style": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "italic"}, false),

				Description: "Font style.",
				Optional:    true,
				Computed:    true,
			},
			"font_weight": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"normal", "bold"}, false),

				Description: "Font weight.",
				Optional:    true,
				Computed:    true,
			},
			"height": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Height.",
				Optional:    true,
				Computed:    true,
			},
			"line_height": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Text line height.",
				Optional:    true,
				Computed:    true,
			},
			"margin_bottom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Margin bottom.",
				Optional:    true,
				Computed:    true,
			},
			"margin_left": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Margin left.",
				Optional:    true,
				Computed:    true,
			},
			"margin_right": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Margin right.",
				Optional:    true,
				Computed:    true,
			},
			"margin_top": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Margin top.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),

				Description: "Report style name.",
				ForceNew:    true,
				Required:    true,
			},
			"options": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Report style options.",
				Optional:         true,
				Computed:         true,
			},
			"padding_bottom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Padding bottom.",
				Optional:    true,
				Computed:    true,
			},
			"padding_left": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Padding left.",
				Optional:    true,
				Computed:    true,
			},
			"padding_right": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Padding right.",
				Optional:    true,
				Computed:    true,
			},
			"padding_top": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Padding top.",
				Optional:    true,
				Computed:    true,
			},
			"width": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Width.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceReportStyleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating ReportStyle resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectReportStyle(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateReportStyle(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportStyle")
	}

	return resourceReportStyleRead(ctx, d, meta)
}

func resourceReportStyleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportStyle(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateReportStyle(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ReportStyle resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportStyle")
	}

	return resourceReportStyleRead(ctx, d, meta)
}

func resourceReportStyleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteReportStyle(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ReportStyle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportStyleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportStyle(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportStyle resource: %v", err)
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

	diags := refreshObjectReportStyle(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectReportStyle(d *schema.ResourceData, o *models.ReportStyle, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Align != nil {
		v := *o.Align

		if err = d.Set("align", v); err != nil {
			return diag.Errorf("error reading align: %v", err)
		}
	}

	if o.BgColor != nil {
		v := *o.BgColor

		if err = d.Set("bg_color", v); err != nil {
			return diag.Errorf("error reading bg_color: %v", err)
		}
	}

	if o.BorderBottom != nil {
		v := *o.BorderBottom

		if err = d.Set("border_bottom", v); err != nil {
			return diag.Errorf("error reading border_bottom: %v", err)
		}
	}

	if o.BorderLeft != nil {
		v := *o.BorderLeft

		if err = d.Set("border_left", v); err != nil {
			return diag.Errorf("error reading border_left: %v", err)
		}
	}

	if o.BorderRight != nil {
		v := *o.BorderRight

		if err = d.Set("border_right", v); err != nil {
			return diag.Errorf("error reading border_right: %v", err)
		}
	}

	if o.BorderTop != nil {
		v := *o.BorderTop

		if err = d.Set("border_top", v); err != nil {
			return diag.Errorf("error reading border_top: %v", err)
		}
	}

	if o.ColumnGap != nil {
		v := *o.ColumnGap

		if err = d.Set("column_gap", v); err != nil {
			return diag.Errorf("error reading column_gap: %v", err)
		}
	}

	if o.ColumnSpan != nil {
		v := *o.ColumnSpan

		if err = d.Set("column_span", v); err != nil {
			return diag.Errorf("error reading column_span: %v", err)
		}
	}

	if o.FgColor != nil {
		v := *o.FgColor

		if err = d.Set("fg_color", v); err != nil {
			return diag.Errorf("error reading fg_color: %v", err)
		}
	}

	if o.FontFamily != nil {
		v := *o.FontFamily

		if err = d.Set("font_family", v); err != nil {
			return diag.Errorf("error reading font_family: %v", err)
		}
	}

	if o.FontSize != nil {
		v := *o.FontSize

		if err = d.Set("font_size", v); err != nil {
			return diag.Errorf("error reading font_size: %v", err)
		}
	}

	if o.FontStyle != nil {
		v := *o.FontStyle

		if err = d.Set("font_style", v); err != nil {
			return diag.Errorf("error reading font_style: %v", err)
		}
	}

	if o.FontWeight != nil {
		v := *o.FontWeight

		if err = d.Set("font_weight", v); err != nil {
			return diag.Errorf("error reading font_weight: %v", err)
		}
	}

	if o.Height != nil {
		v := *o.Height

		if err = d.Set("height", v); err != nil {
			return diag.Errorf("error reading height: %v", err)
		}
	}

	if o.LineHeight != nil {
		v := *o.LineHeight

		if err = d.Set("line_height", v); err != nil {
			return diag.Errorf("error reading line_height: %v", err)
		}
	}

	if o.MarginBottom != nil {
		v := *o.MarginBottom

		if err = d.Set("margin_bottom", v); err != nil {
			return diag.Errorf("error reading margin_bottom: %v", err)
		}
	}

	if o.MarginLeft != nil {
		v := *o.MarginLeft

		if err = d.Set("margin_left", v); err != nil {
			return diag.Errorf("error reading margin_left: %v", err)
		}
	}

	if o.MarginRight != nil {
		v := *o.MarginRight

		if err = d.Set("margin_right", v); err != nil {
			return diag.Errorf("error reading margin_right: %v", err)
		}
	}

	if o.MarginTop != nil {
		v := *o.MarginTop

		if err = d.Set("margin_top", v); err != nil {
			return diag.Errorf("error reading margin_top: %v", err)
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

	if o.PaddingBottom != nil {
		v := *o.PaddingBottom

		if err = d.Set("padding_bottom", v); err != nil {
			return diag.Errorf("error reading padding_bottom: %v", err)
		}
	}

	if o.PaddingLeft != nil {
		v := *o.PaddingLeft

		if err = d.Set("padding_left", v); err != nil {
			return diag.Errorf("error reading padding_left: %v", err)
		}
	}

	if o.PaddingRight != nil {
		v := *o.PaddingRight

		if err = d.Set("padding_right", v); err != nil {
			return diag.Errorf("error reading padding_right: %v", err)
		}
	}

	if o.PaddingTop != nil {
		v := *o.PaddingTop

		if err = d.Set("padding_top", v); err != nil {
			return diag.Errorf("error reading padding_top: %v", err)
		}
	}

	if o.Width != nil {
		v := *o.Width

		if err = d.Set("width", v); err != nil {
			return diag.Errorf("error reading width: %v", err)
		}
	}

	return nil
}

func getObjectReportStyle(d *schema.ResourceData, sv string) (*models.ReportStyle, diag.Diagnostics) {
	obj := models.ReportStyle{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("align"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("align", sv)
				diags = append(diags, e)
			}
			obj.Align = &v2
		}
	}
	if v1, ok := d.GetOk("bg_color"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bg_color", sv)
				diags = append(diags, e)
			}
			obj.BgColor = &v2
		}
	}
	if v1, ok := d.GetOk("border_bottom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("border_bottom", sv)
				diags = append(diags, e)
			}
			obj.BorderBottom = &v2
		}
	}
	if v1, ok := d.GetOk("border_left"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("border_left", sv)
				diags = append(diags, e)
			}
			obj.BorderLeft = &v2
		}
	}
	if v1, ok := d.GetOk("border_right"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("border_right", sv)
				diags = append(diags, e)
			}
			obj.BorderRight = &v2
		}
	}
	if v1, ok := d.GetOk("border_top"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("border_top", sv)
				diags = append(diags, e)
			}
			obj.BorderTop = &v2
		}
	}
	if v1, ok := d.GetOk("column_gap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("column_gap", sv)
				diags = append(diags, e)
			}
			obj.ColumnGap = &v2
		}
	}
	if v1, ok := d.GetOk("column_span"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("column_span", sv)
				diags = append(diags, e)
			}
			obj.ColumnSpan = &v2
		}
	}
	if v1, ok := d.GetOk("fg_color"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fg_color", sv)
				diags = append(diags, e)
			}
			obj.FgColor = &v2
		}
	}
	if v1, ok := d.GetOk("font_family"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("font_family", sv)
				diags = append(diags, e)
			}
			obj.FontFamily = &v2
		}
	}
	if v1, ok := d.GetOk("font_size"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("font_size", sv)
				diags = append(diags, e)
			}
			obj.FontSize = &v2
		}
	}
	if v1, ok := d.GetOk("font_style"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("font_style", sv)
				diags = append(diags, e)
			}
			obj.FontStyle = &v2
		}
	}
	if v1, ok := d.GetOk("font_weight"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("font_weight", sv)
				diags = append(diags, e)
			}
			obj.FontWeight = &v2
		}
	}
	if v1, ok := d.GetOk("height"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("height", sv)
				diags = append(diags, e)
			}
			obj.Height = &v2
		}
	}
	if v1, ok := d.GetOk("line_height"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("line_height", sv)
				diags = append(diags, e)
			}
			obj.LineHeight = &v2
		}
	}
	if v1, ok := d.GetOk("margin_bottom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("margin_bottom", sv)
				diags = append(diags, e)
			}
			obj.MarginBottom = &v2
		}
	}
	if v1, ok := d.GetOk("margin_left"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("margin_left", sv)
				diags = append(diags, e)
			}
			obj.MarginLeft = &v2
		}
	}
	if v1, ok := d.GetOk("margin_right"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("margin_right", sv)
				diags = append(diags, e)
			}
			obj.MarginRight = &v2
		}
	}
	if v1, ok := d.GetOk("margin_top"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("margin_top", sv)
				diags = append(diags, e)
			}
			obj.MarginTop = &v2
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
	if v1, ok := d.GetOk("padding_bottom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("padding_bottom", sv)
				diags = append(diags, e)
			}
			obj.PaddingBottom = &v2
		}
	}
	if v1, ok := d.GetOk("padding_left"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("padding_left", sv)
				diags = append(diags, e)
			}
			obj.PaddingLeft = &v2
		}
	}
	if v1, ok := d.GetOk("padding_right"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("padding_right", sv)
				diags = append(diags, e)
			}
			obj.PaddingRight = &v2
		}
	}
	if v1, ok := d.GetOk("padding_top"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("padding_top", sv)
				diags = append(diags, e)
			}
			obj.PaddingTop = &v2
		}
	}
	if v1, ok := d.GetOk("width"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("width", sv)
				diags = append(diags, e)
			}
			obj.Width = &v2
		}
	}
	return &obj, diags
}
