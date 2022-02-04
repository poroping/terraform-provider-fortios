// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Filters for null device logging.

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

func resourceLogNullDeviceFilter() *schema.Resource {
	return &schema.Resource{
		Description: "Filters for null device logging.",

		CreateContext: resourceLogNullDeviceFilterCreate,
		ReadContext:   resourceLogNullDeviceFilterRead,
		UpdateContext: resourceLogNullDeviceFilterUpdate,
		DeleteContext: resourceLogNullDeviceFilterDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"anomaly": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable anomaly logging.",
				Optional:    true,
				Computed:    true,
			},
			"filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Null-device log filter.",
				Optional:    true,
				Computed:    true,
			},
			"filter_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),

				Description: "Include/exclude logs that match the filter.",
				Optional:    true,
				Computed:    true,
			},
			"forward_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable forward traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"free_style": {
				Type:        schema.TypeList,
				Description: "Free style filters.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"traffic", "event", "virus", "webfilter", "attack", "spam", "anomaly", "voip", "dlp", "app-ctrl", "waf", "gtp", "dns", "ssh", "ssl", "file-filter", "icap"}, false),

							Description: "Log category.",
							Optional:    true,
							Computed:    true,
						},
						"filter": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),

							Description: "Free style filter string.",
							Optional:    true,
							Computed:    true,
						},
						"filter_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),

							Description: "Include/exclude logs that match the filter.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"gtp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable GTP messages logging.",
				Optional:    true,
				Computed:    true,
			},
			"local_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local in or out traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"multicast_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multicast traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"emergency", "alert", "critical", "error", "warning", "notification", "information", "debug"}, false),

				Description: "Lowest severity level to log.",
				Optional:    true,
				Computed:    true,
			},
			"sniffer_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sniffer traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"voip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VoIP logging.",
				Optional:    true,
				Computed:    true,
			},
			"ztna_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ztna traffic logging.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogNullDeviceFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogNullDeviceFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogNullDeviceFilter(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogNullDeviceFilter")
	}

	return resourceLogNullDeviceFilterRead(ctx, d, meta)
}

func resourceLogNullDeviceFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogNullDeviceFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogNullDeviceFilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogNullDeviceFilter resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogNullDeviceFilter")
	}

	return resourceLogNullDeviceFilterRead(ctx, d, meta)
}

func resourceLogNullDeviceFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectLogNullDeviceFilter(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateLogNullDeviceFilter(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogNullDeviceFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogNullDeviceFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogNullDeviceFilter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogNullDeviceFilter resource: %v", err)
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

	diags := refreshObjectLogNullDeviceFilter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenLogNullDeviceFilterFreeStyle(v *[]models.LogNullDeviceFilterFreeStyle, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Filter; tmp != nil {
				v["filter"] = *tmp
			}

			if tmp := cfg.FilterType; tmp != nil {
				v["filter_type"] = *tmp
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

func refreshObjectLogNullDeviceFilter(d *schema.ResourceData, o *models.LogNullDeviceFilter, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Anomaly != nil {
		v := *o.Anomaly

		if err = d.Set("anomaly", v); err != nil {
			return diag.Errorf("error reading anomaly: %v", err)
		}
	}

	if o.Filter != nil {
		v := *o.Filter

		if err = d.Set("filter", v); err != nil {
			return diag.Errorf("error reading filter: %v", err)
		}
	}

	if o.FilterType != nil {
		v := *o.FilterType

		if err = d.Set("filter_type", v); err != nil {
			return diag.Errorf("error reading filter_type: %v", err)
		}
	}

	if o.ForwardTraffic != nil {
		v := *o.ForwardTraffic

		if err = d.Set("forward_traffic", v); err != nil {
			return diag.Errorf("error reading forward_traffic: %v", err)
		}
	}

	if o.FreeStyle != nil {
		if err = d.Set("free_style", flattenLogNullDeviceFilterFreeStyle(o.FreeStyle, sort)); err != nil {
			return diag.Errorf("error reading free_style: %v", err)
		}
	}

	if o.Gtp != nil {
		v := *o.Gtp

		if err = d.Set("gtp", v); err != nil {
			return diag.Errorf("error reading gtp: %v", err)
		}
	}

	if o.LocalTraffic != nil {
		v := *o.LocalTraffic

		if err = d.Set("local_traffic", v); err != nil {
			return diag.Errorf("error reading local_traffic: %v", err)
		}
	}

	if o.MulticastTraffic != nil {
		v := *o.MulticastTraffic

		if err = d.Set("multicast_traffic", v); err != nil {
			return diag.Errorf("error reading multicast_traffic: %v", err)
		}
	}

	if o.Severity != nil {
		v := *o.Severity

		if err = d.Set("severity", v); err != nil {
			return diag.Errorf("error reading severity: %v", err)
		}
	}

	if o.SnifferTraffic != nil {
		v := *o.SnifferTraffic

		if err = d.Set("sniffer_traffic", v); err != nil {
			return diag.Errorf("error reading sniffer_traffic: %v", err)
		}
	}

	if o.Voip != nil {
		v := *o.Voip

		if err = d.Set("voip", v); err != nil {
			return diag.Errorf("error reading voip: %v", err)
		}
	}

	if o.ZtnaTraffic != nil {
		v := *o.ZtnaTraffic

		if err = d.Set("ztna_traffic", v); err != nil {
			return diag.Errorf("error reading ztna_traffic: %v", err)
		}
	}

	return nil
}

func expandLogNullDeviceFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogNullDeviceFilterFreeStyle, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogNullDeviceFilterFreeStyle

	for i := range l {
		tmp := models.LogNullDeviceFilterFreeStyle{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Filter = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.filter_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FilterType = &v2
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

func getObjectLogNullDeviceFilter(d *schema.ResourceData, sv string) (*models.LogNullDeviceFilter, diag.Diagnostics) {
	obj := models.LogNullDeviceFilter{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("anomaly"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("anomaly", sv)
				diags = append(diags, e)
			}
			obj.Anomaly = &v2
		}
	}
	if v1, ok := d.GetOk("filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("filter", sv)
				diags = append(diags, e)
			}
			obj.Filter = &v2
		}
	}
	if v1, ok := d.GetOk("filter_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("filter_type", sv)
				diags = append(diags, e)
			}
			obj.FilterType = &v2
		}
	}
	if v1, ok := d.GetOk("forward_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forward_traffic", sv)
				diags = append(diags, e)
			}
			obj.ForwardTraffic = &v2
		}
	}
	if v, ok := d.GetOk("free_style"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("free_style", sv)
			diags = append(diags, e)
		}
		t, err := expandLogNullDeviceFilterFreeStyle(d, v, "free_style", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FreeStyle = t
		}
	} else if d.HasChange("free_style") {
		old, new := d.GetChange("free_style")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FreeStyle = &[]models.LogNullDeviceFilterFreeStyle{}
		}
	}
	if v1, ok := d.GetOk("gtp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("gtp", sv)
				diags = append(diags, e)
			}
			obj.Gtp = &v2
		}
	}
	if v1, ok := d.GetOk("local_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_traffic", sv)
				diags = append(diags, e)
			}
			obj.LocalTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_traffic", sv)
				diags = append(diags, e)
			}
			obj.MulticastTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("severity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("severity", sv)
				diags = append(diags, e)
			}
			obj.Severity = &v2
		}
	}
	if v1, ok := d.GetOk("sniffer_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sniffer_traffic", sv)
				diags = append(diags, e)
			}
			obj.SnifferTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("voip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip", sv)
				diags = append(diags, e)
			}
			obj.Voip = &v2
		}
	}
	if v1, ok := d.GetOk("ztna_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("ztna_traffic", sv)
				diags = append(diags, e)
			}
			obj.ZtnaTraffic = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectLogNullDeviceFilter(d *schema.ResourceData, sv string) (*models.LogNullDeviceFilter, diag.Diagnostics) {
	obj := models.LogNullDeviceFilter{}
	diags := diag.Diagnostics{}

	obj.FreeStyle = &[]models.LogNullDeviceFilterFreeStyle{}

	return &obj, diags
}
