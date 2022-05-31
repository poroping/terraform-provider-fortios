// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WiFi quality of service (QoS) profiles.

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

func resourceWirelessControllerQosProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WiFi quality of service (QoS) profiles.",

		CreateContext: resourceWirelessControllerQosProfileCreate,
		ReadContext:   resourceWirelessControllerQosProfileRead,
		UpdateContext: resourceWirelessControllerQosProfileUpdate,
		DeleteContext: resourceWirelessControllerQosProfileDelete,

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
			"bandwidth_admission_control": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WMM bandwidth admission control.",
				Optional:    true,
				Computed:    true,
			},
			"bandwidth_capacity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 600000),

				Description: "Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).",
				Optional:    true,
				Computed:    true,
			},
			"burst": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable client rate burst.",
				Optional:    true,
				Computed:    true,
			},
			"call_admission_control": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WMM call admission control.",
				Optional:    true,
				Computed:    true,
			},
			"call_capacity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),

				Description: "Maximum number of Voice over WLAN (VoWLAN) phones allowed (0 - 60, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"downlink": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),

				Description: "Maximum downlink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Optional:    true,
				Computed:    true,
			},
			"downlink_sta": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),

				Description: "Maximum downlink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Optional:    true,
				Computed:    true,
			},
			"dscp_wmm_be": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for best effort access (default = 0 24).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP WMM mapping numbers (0 - 63).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dscp_wmm_bk": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for background access (default = 8 16).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP WMM mapping numbers (0 - 63).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dscp_wmm_mapping": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Differentiated Services Code Point (DSCP) mapping.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_wmm_vi": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for video access (default = 32 40).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP WMM mapping numbers (0 - 63).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dscp_wmm_vo": {
				Type:        schema.TypeList,
				Description: "DSCP mapping for voice access (default = 48 56).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),

							Description: "DSCP WMM mapping numbers (0 - 63).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WiFi QoS profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"uplink": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),

				Description: "Maximum uplink bandwidth for Virtual Access Points (VAPs) (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Optional:    true,
				Computed:    true,
			},
			"uplink_sta": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2097152),

				Description: "Maximum uplink bandwidth for clients (0 - 2097152 Kbps, default = 0, 0 means no limit).",
				Optional:    true,
				Computed:    true,
			},
			"wmm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WiFi multi-media (WMM) control.",
				Optional:    true,
				Computed:    true,
			},
			"wmm_be_dscp": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),

				Description: "DSCP marking for best effort access (default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"wmm_bk_dscp": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),

				Description: "DSCP marking for background access (default = 8).",
				Optional:    true,
				Computed:    true,
			},
			"wmm_dscp_marking": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WMM Differentiated Services Code Point (DSCP) marking.",
				Optional:    true,
				Computed:    true,
			},
			"wmm_uapsd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WMM Unscheduled Automatic Power Save Delivery (U-APSD) power save mode.",
				Optional:    true,
				Computed:    true,
			},
			"wmm_vi_dscp": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),

				Description: "DSCP marking for video access (default = 32).",
				Optional:    true,
				Computed:    true,
			},
			"wmm_vo_dscp": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 63),

				Description: "DSCP marking for voice access (default = 48).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerQosProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerQosProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerQosProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerQosProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerQosProfile")
	}

	return resourceWirelessControllerQosProfileRead(ctx, d, meta)
}

func resourceWirelessControllerQosProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerQosProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerQosProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerQosProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerQosProfile")
	}

	return resourceWirelessControllerQosProfileRead(ctx, d, meta)
}

func resourceWirelessControllerQosProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerQosProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerQosProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerQosProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerQosProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerQosProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerQosProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerQosProfileDscpWmmBe(d *schema.ResourceData, v *[]models.WirelessControllerQosProfileDscpWmmBe, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
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

func flattenWirelessControllerQosProfileDscpWmmBk(d *schema.ResourceData, v *[]models.WirelessControllerQosProfileDscpWmmBk, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
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

func flattenWirelessControllerQosProfileDscpWmmVi(d *schema.ResourceData, v *[]models.WirelessControllerQosProfileDscpWmmVi, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
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

func flattenWirelessControllerQosProfileDscpWmmVo(d *schema.ResourceData, v *[]models.WirelessControllerQosProfileDscpWmmVo, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
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

func refreshObjectWirelessControllerQosProfile(d *schema.ResourceData, o *models.WirelessControllerQosProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BandwidthAdmissionControl != nil {
		v := *o.BandwidthAdmissionControl

		if err = d.Set("bandwidth_admission_control", v); err != nil {
			return diag.Errorf("error reading bandwidth_admission_control: %v", err)
		}
	}

	if o.BandwidthCapacity != nil {
		v := *o.BandwidthCapacity

		if err = d.Set("bandwidth_capacity", v); err != nil {
			return diag.Errorf("error reading bandwidth_capacity: %v", err)
		}
	}

	if o.Burst != nil {
		v := *o.Burst

		if err = d.Set("burst", v); err != nil {
			return diag.Errorf("error reading burst: %v", err)
		}
	}

	if o.CallAdmissionControl != nil {
		v := *o.CallAdmissionControl

		if err = d.Set("call_admission_control", v); err != nil {
			return diag.Errorf("error reading call_admission_control: %v", err)
		}
	}

	if o.CallCapacity != nil {
		v := *o.CallCapacity

		if err = d.Set("call_capacity", v); err != nil {
			return diag.Errorf("error reading call_capacity: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Downlink != nil {
		v := *o.Downlink

		if err = d.Set("downlink", v); err != nil {
			return diag.Errorf("error reading downlink: %v", err)
		}
	}

	if o.DownlinkSta != nil {
		v := *o.DownlinkSta

		if err = d.Set("downlink_sta", v); err != nil {
			return diag.Errorf("error reading downlink_sta: %v", err)
		}
	}

	if o.DscpWmmBe != nil {
		if err = d.Set("dscp_wmm_be", flattenWirelessControllerQosProfileDscpWmmBe(d, o.DscpWmmBe, "dscp_wmm_be", sort)); err != nil {
			return diag.Errorf("error reading dscp_wmm_be: %v", err)
		}
	}

	if o.DscpWmmBk != nil {
		if err = d.Set("dscp_wmm_bk", flattenWirelessControllerQosProfileDscpWmmBk(d, o.DscpWmmBk, "dscp_wmm_bk", sort)); err != nil {
			return diag.Errorf("error reading dscp_wmm_bk: %v", err)
		}
	}

	if o.DscpWmmMapping != nil {
		v := *o.DscpWmmMapping

		if err = d.Set("dscp_wmm_mapping", v); err != nil {
			return diag.Errorf("error reading dscp_wmm_mapping: %v", err)
		}
	}

	if o.DscpWmmVi != nil {
		if err = d.Set("dscp_wmm_vi", flattenWirelessControllerQosProfileDscpWmmVi(d, o.DscpWmmVi, "dscp_wmm_vi", sort)); err != nil {
			return diag.Errorf("error reading dscp_wmm_vi: %v", err)
		}
	}

	if o.DscpWmmVo != nil {
		if err = d.Set("dscp_wmm_vo", flattenWirelessControllerQosProfileDscpWmmVo(d, o.DscpWmmVo, "dscp_wmm_vo", sort)); err != nil {
			return diag.Errorf("error reading dscp_wmm_vo: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Uplink != nil {
		v := *o.Uplink

		if err = d.Set("uplink", v); err != nil {
			return diag.Errorf("error reading uplink: %v", err)
		}
	}

	if o.UplinkSta != nil {
		v := *o.UplinkSta

		if err = d.Set("uplink_sta", v); err != nil {
			return diag.Errorf("error reading uplink_sta: %v", err)
		}
	}

	if o.Wmm != nil {
		v := *o.Wmm

		if err = d.Set("wmm", v); err != nil {
			return diag.Errorf("error reading wmm: %v", err)
		}
	}

	if o.WmmBeDscp != nil {
		v := *o.WmmBeDscp

		if err = d.Set("wmm_be_dscp", v); err != nil {
			return diag.Errorf("error reading wmm_be_dscp: %v", err)
		}
	}

	if o.WmmBkDscp != nil {
		v := *o.WmmBkDscp

		if err = d.Set("wmm_bk_dscp", v); err != nil {
			return diag.Errorf("error reading wmm_bk_dscp: %v", err)
		}
	}

	if o.WmmDscpMarking != nil {
		v := *o.WmmDscpMarking

		if err = d.Set("wmm_dscp_marking", v); err != nil {
			return diag.Errorf("error reading wmm_dscp_marking: %v", err)
		}
	}

	if o.WmmUapsd != nil {
		v := *o.WmmUapsd

		if err = d.Set("wmm_uapsd", v); err != nil {
			return diag.Errorf("error reading wmm_uapsd: %v", err)
		}
	}

	if o.WmmViDscp != nil {
		v := *o.WmmViDscp

		if err = d.Set("wmm_vi_dscp", v); err != nil {
			return diag.Errorf("error reading wmm_vi_dscp: %v", err)
		}
	}

	if o.WmmVoDscp != nil {
		v := *o.WmmVoDscp

		if err = d.Set("wmm_vo_dscp", v); err != nil {
			return diag.Errorf("error reading wmm_vo_dscp: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerQosProfileDscpWmmBe(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerQosProfileDscpWmmBe, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerQosProfileDscpWmmBe

	for i := range l {
		tmp := models.WirelessControllerQosProfileDscpWmmBe{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerQosProfileDscpWmmBk(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerQosProfileDscpWmmBk, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerQosProfileDscpWmmBk

	for i := range l {
		tmp := models.WirelessControllerQosProfileDscpWmmBk{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerQosProfileDscpWmmVi(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerQosProfileDscpWmmVi, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerQosProfileDscpWmmVi

	for i := range l {
		tmp := models.WirelessControllerQosProfileDscpWmmVi{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerQosProfileDscpWmmVo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerQosProfileDscpWmmVo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerQosProfileDscpWmmVo

	for i := range l {
		tmp := models.WirelessControllerQosProfileDscpWmmVo{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerQosProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerQosProfile, diag.Diagnostics) {
	obj := models.WirelessControllerQosProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("bandwidth_admission_control"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bandwidth_admission_control", sv)
				diags = append(diags, e)
			}
			obj.BandwidthAdmissionControl = &v2
		}
	}
	if v1, ok := d.GetOk("bandwidth_capacity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bandwidth_capacity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BandwidthCapacity = &tmp
		}
	}
	if v1, ok := d.GetOk("burst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("burst", sv)
				diags = append(diags, e)
			}
			obj.Burst = &v2
		}
	}
	if v1, ok := d.GetOk("call_admission_control"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("call_admission_control", sv)
				diags = append(diags, e)
			}
			obj.CallAdmissionControl = &v2
		}
	}
	if v1, ok := d.GetOk("call_capacity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("call_capacity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CallCapacity = &tmp
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("downlink"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("downlink", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Downlink = &tmp
		}
	}
	if v1, ok := d.GetOk("downlink_sta"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("downlink_sta", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DownlinkSta = &tmp
		}
	}
	if v, ok := d.GetOk("dscp_wmm_be"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dscp_wmm_be", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerQosProfileDscpWmmBe(d, v, "dscp_wmm_be", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DscpWmmBe = t
		}
	} else if d.HasChange("dscp_wmm_be") {
		old, new := d.GetChange("dscp_wmm_be")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DscpWmmBe = &[]models.WirelessControllerQosProfileDscpWmmBe{}
		}
	}
	if v, ok := d.GetOk("dscp_wmm_bk"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dscp_wmm_bk", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerQosProfileDscpWmmBk(d, v, "dscp_wmm_bk", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DscpWmmBk = t
		}
	} else if d.HasChange("dscp_wmm_bk") {
		old, new := d.GetChange("dscp_wmm_bk")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DscpWmmBk = &[]models.WirelessControllerQosProfileDscpWmmBk{}
		}
	}
	if v1, ok := d.GetOk("dscp_wmm_mapping"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_wmm_mapping", sv)
				diags = append(diags, e)
			}
			obj.DscpWmmMapping = &v2
		}
	}
	if v, ok := d.GetOk("dscp_wmm_vi"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dscp_wmm_vi", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerQosProfileDscpWmmVi(d, v, "dscp_wmm_vi", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DscpWmmVi = t
		}
	} else if d.HasChange("dscp_wmm_vi") {
		old, new := d.GetChange("dscp_wmm_vi")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DscpWmmVi = &[]models.WirelessControllerQosProfileDscpWmmVi{}
		}
	}
	if v, ok := d.GetOk("dscp_wmm_vo"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dscp_wmm_vo", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerQosProfileDscpWmmVo(d, v, "dscp_wmm_vo", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DscpWmmVo = t
		}
	} else if d.HasChange("dscp_wmm_vo") {
		old, new := d.GetChange("dscp_wmm_vo")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DscpWmmVo = &[]models.WirelessControllerQosProfileDscpWmmVo{}
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
	if v1, ok := d.GetOk("uplink"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uplink", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Uplink = &tmp
		}
	}
	if v1, ok := d.GetOk("uplink_sta"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uplink_sta", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UplinkSta = &tmp
		}
	}
	if v1, ok := d.GetOk("wmm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm", sv)
				diags = append(diags, e)
			}
			obj.Wmm = &v2
		}
	}
	if v1, ok := d.GetOk("wmm_be_dscp"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm_be_dscp", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WmmBeDscp = &tmp
		}
	}
	if v1, ok := d.GetOk("wmm_bk_dscp"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm_bk_dscp", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WmmBkDscp = &tmp
		}
	}
	if v1, ok := d.GetOk("wmm_dscp_marking"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm_dscp_marking", sv)
				diags = append(diags, e)
			}
			obj.WmmDscpMarking = &v2
		}
	}
	if v1, ok := d.GetOk("wmm_uapsd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm_uapsd", sv)
				diags = append(diags, e)
			}
			obj.WmmUapsd = &v2
		}
	}
	if v1, ok := d.GetOk("wmm_vi_dscp"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm_vi_dscp", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WmmViDscp = &tmp
		}
	}
	if v1, ok := d.GetOk("wmm_vo_dscp"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wmm_vo_dscp", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WmmVoDscp = &tmp
		}
	}
	return &obj, diags
}
