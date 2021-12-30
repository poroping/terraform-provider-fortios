// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Report setting configuration.

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

func resourceReportSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Report setting configuration.",

		CreateContext: resourceReportSettingCreate,
		ReadContext:   resourceReportSettingRead,
		UpdateContext: resourceReportSettingUpdate,
		DeleteContext: resourceReportSettingDelete,

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
			"fortiview": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable historical FortiView.",
				Optional:    true,
				Computed:    true,
			},
			"pdf_report": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable PDF report.",
				Optional:    true,
				Computed:    true,
			},
			"report_source": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Report log source.",
				Optional:         true,
				Computed:         true,
			},
			"top_n": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1000, 20000),

				Description: "Number of items to populate (1000 - 20000).",
				Optional:    true,
				Computed:    true,
			},
			"web_browsing_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 15),

				Description: "Web browsing time calculation threshold (3 - 15 min).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceReportSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateReportSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportSetting")
	}

	return resourceReportSettingRead(ctx, d, meta)
}

func resourceReportSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectReportSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateReportSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating ReportSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("ReportSetting")
	}

	return resourceReportSettingRead(ctx, d, meta)
}

func resourceReportSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteReportSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting ReportSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadReportSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading ReportSetting resource: %v", err)
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

	diags := refreshObjectReportSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectReportSetting(d *schema.ResourceData, o *models.ReportSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Fortiview != nil {
		v := *o.Fortiview

		if err = d.Set("fortiview", v); err != nil {
			return diag.Errorf("error reading fortiview: %v", err)
		}
	}

	if o.PdfReport != nil {
		v := *o.PdfReport

		if err = d.Set("pdf_report", v); err != nil {
			return diag.Errorf("error reading pdf_report: %v", err)
		}
	}

	if o.ReportSource != nil {
		v := *o.ReportSource

		if err = d.Set("report_source", v); err != nil {
			return diag.Errorf("error reading report_source: %v", err)
		}
	}

	if o.TopN != nil {
		v := *o.TopN

		if err = d.Set("top_n", v); err != nil {
			return diag.Errorf("error reading top_n: %v", err)
		}
	}

	if o.WebBrowsingThreshold != nil {
		v := *o.WebBrowsingThreshold

		if err = d.Set("web_browsing_threshold", v); err != nil {
			return diag.Errorf("error reading web_browsing_threshold: %v", err)
		}
	}

	return nil
}

func getObjectReportSetting(d *schema.ResourceData, sv string) (*models.ReportSetting, diag.Diagnostics) {
	obj := models.ReportSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fortiview"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortiview", sv)
				diags = append(diags, e)
			}
			obj.Fortiview = &v2
		}
	}
	if v1, ok := d.GetOk("pdf_report"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pdf_report", sv)
				diags = append(diags, e)
			}
			obj.PdfReport = &v2
		}
	}
	if v1, ok := d.GetOk("report_source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("report_source", sv)
				diags = append(diags, e)
			}
			obj.ReportSource = &v2
		}
	}
	if v1, ok := d.GetOk("top_n"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("top_n", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TopN = &tmp
		}
	}
	if v1, ok := d.GetOk("web_browsing_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("web_browsing_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WebBrowsingThreshold = &tmp
		}
	}
	return &obj, diags
}
