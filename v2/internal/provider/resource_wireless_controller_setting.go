// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: VDOM wireless controller configuration.

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

func resourceWirelessControllerSetting() *schema.Resource {
	return &schema.Resource{
		Description: "VDOM wireless controller configuration.",

		CreateContext: resourceWirelessControllerSettingCreate,
		ReadContext:   resourceWirelessControllerSettingRead,
		UpdateContext: resourceWirelessControllerSettingUpdate,
		DeleteContext: resourceWirelessControllerSettingDelete,

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
			"account_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "FortiCloud customer account ID.",
				Optional:    true,
				Computed:    true,
			},
			"country": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"--", "AF", "AL", "DZ", "AS", "AO", "AR", "AM", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BA", "BW", "BR", "BN", "BG", "BF", "KH", "CM", "KY", "CF", "TD", "CL", "CN", "CX", "CO", "CG", "CD", "CR", "HR", "CY", "CZ", "DK", "DM", "DO", "EC", "EG", "SV", "ET", "EE", "GF", "PF", "FO", "FJ", "FI", "FR", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GY", "HT", "HN", "HK", "HU", "IS", "IN", "ID", "IQ", "IE", "IM", "IL", "IT", "CI", "JM", "JO", "KZ", "KE", "KR", "KW", "LA", "LV", "LB", "LS", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MA", "MZ", "MM", "NA", "NP", "NL", "AN", "AW", "NZ", "NI", "NE", "NO", "MP", "OM", "PK", "PW", "PA", "PG", "PY", "PE", "PH", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "KN", "LC", "MF", "PM", "VC", "SA", "SN", "RS", "ME", "SL", "SG", "SK", "SI", "ZA", "ES", "LK", "SE", "SR", "CH", "TW", "TZ", "TH", "TG", "TT", "TN", "TR", "TM", "AE", "TC", "UG", "UA", "GB", "US", "PS", "UY", "UZ", "VU", "VE", "VN", "VI", "WF", "YE", "ZM", "ZW", "JP", "CA"}, false),

				Description: "Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.",
				Optional:    true,
				Computed:    true,
			},
			"darrp_optimize": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"darrp_optimize_schedules": {
				Type:        schema.TypeList,
				Description: "Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Schedule name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"device_holdoff": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),

				Description: "Lower limit of creation time of device for identification in minutes (0 - 60, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"device_idle": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 14400),

				Description: "Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440).",
				Optional:    true,
				Computed:    true,
			},
			"device_weight": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"duplicate_ssid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM.",
				Optional:    true,
				Computed:    true,
			},
			"fake_ssid_action": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Actions taken for detected fake SSID.",
				Optional:         true,
				Computed:         true,
			},
			"fapc_compatibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FAP-C series compatibility.",
				Optional:    true,
				Computed:    true,
			},
			"firmware_provision_on_authorization": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable automatic provisioning of latest firmware on authorization.",
				Optional:    true,
				Computed:    true,
			},
			"offending_ssid": {
				Type:        schema.TypeList,
				Description: "Configure offending SSID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "Actions taken for detected offending SSID.",
							Optional:         true,
							Computed:         true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"ssid_pattern": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 33),

							Description: "Define offending SSID pattern (case insensitive), eg: word, word*, *word, wo*rd.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"phishing_ssid_detect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable phishing SSID detection.",
				Optional:    true,
				Computed:    true,
			},
			"wfa_compatibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WFA compatibility.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerSetting")
	}

	return resourceWirelessControllerSettingRead(ctx, d, meta)
}

func resourceWirelessControllerSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerSetting")
	}

	return resourceWirelessControllerSettingRead(ctx, d, meta)
}

func resourceWirelessControllerSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWirelessControllerSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWirelessControllerSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerSetting resource: %v", err)
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

	diags := refreshObjectWirelessControllerSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerSettingDarrpOptimizeSchedules(v *[]models.WirelessControllerSettingDarrpOptimizeSchedules, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerSettingOffendingSsid(v *[]models.WirelessControllerSettingOffendingSsid, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SsidPattern; tmp != nil {
				v["ssid_pattern"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerSetting(d *schema.ResourceData, o *models.WirelessControllerSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AccountId != nil {
		v := *o.AccountId

		if err = d.Set("account_id", v); err != nil {
			return diag.Errorf("error reading account_id: %v", err)
		}
	}

	if o.Country != nil {
		v := *o.Country

		if err = d.Set("country", v); err != nil {
			return diag.Errorf("error reading country: %v", err)
		}
	}

	if o.DarrpOptimize != nil {
		v := *o.DarrpOptimize

		if err = d.Set("darrp_optimize", v); err != nil {
			return diag.Errorf("error reading darrp_optimize: %v", err)
		}
	}

	if o.DarrpOptimizeSchedules != nil {
		if err = d.Set("darrp_optimize_schedules", flattenWirelessControllerSettingDarrpOptimizeSchedules(o.DarrpOptimizeSchedules, sort)); err != nil {
			return diag.Errorf("error reading darrp_optimize_schedules: %v", err)
		}
	}

	if o.DeviceHoldoff != nil {
		v := *o.DeviceHoldoff

		if err = d.Set("device_holdoff", v); err != nil {
			return diag.Errorf("error reading device_holdoff: %v", err)
		}
	}

	if o.DeviceIdle != nil {
		v := *o.DeviceIdle

		if err = d.Set("device_idle", v); err != nil {
			return diag.Errorf("error reading device_idle: %v", err)
		}
	}

	if o.DeviceWeight != nil {
		v := *o.DeviceWeight

		if err = d.Set("device_weight", v); err != nil {
			return diag.Errorf("error reading device_weight: %v", err)
		}
	}

	if o.DuplicateSsid != nil {
		v := *o.DuplicateSsid

		if err = d.Set("duplicate_ssid", v); err != nil {
			return diag.Errorf("error reading duplicate_ssid: %v", err)
		}
	}

	if o.FakeSsidAction != nil {
		v := *o.FakeSsidAction

		if err = d.Set("fake_ssid_action", v); err != nil {
			return diag.Errorf("error reading fake_ssid_action: %v", err)
		}
	}

	if o.FapcCompatibility != nil {
		v := *o.FapcCompatibility

		if err = d.Set("fapc_compatibility", v); err != nil {
			return diag.Errorf("error reading fapc_compatibility: %v", err)
		}
	}

	if o.FirmwareProvisionOnAuthorization != nil {
		v := *o.FirmwareProvisionOnAuthorization

		if err = d.Set("firmware_provision_on_authorization", v); err != nil {
			return diag.Errorf("error reading firmware_provision_on_authorization: %v", err)
		}
	}

	if o.OffendingSsid != nil {
		if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o.OffendingSsid, sort)); err != nil {
			return diag.Errorf("error reading offending_ssid: %v", err)
		}
	}

	if o.PhishingSsidDetect != nil {
		v := *o.PhishingSsidDetect

		if err = d.Set("phishing_ssid_detect", v); err != nil {
			return diag.Errorf("error reading phishing_ssid_detect: %v", err)
		}
	}

	if o.WfaCompatibility != nil {
		v := *o.WfaCompatibility

		if err = d.Set("wfa_compatibility", v); err != nil {
			return diag.Errorf("error reading wfa_compatibility: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerSettingDarrpOptimizeSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerSettingDarrpOptimizeSchedules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerSettingDarrpOptimizeSchedules

	for i := range l {
		tmp := models.WirelessControllerSettingDarrpOptimizeSchedules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerSettingOffendingSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerSettingOffendingSsid, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerSettingOffendingSsid

	for i := range l {
		tmp := models.WirelessControllerSettingOffendingSsid{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssid_pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SsidPattern = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerSetting(d *schema.ResourceData, sv string) (*models.WirelessControllerSetting, diag.Diagnostics) {
	obj := models.WirelessControllerSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("account_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("account_id", sv)
				diags = append(diags, e)
			}
			obj.AccountId = &v2
		}
	}
	if v1, ok := d.GetOk("country"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("country", sv)
				diags = append(diags, e)
			}
			obj.Country = &v2
		}
	}
	if v1, ok := d.GetOk("darrp_optimize"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("darrp_optimize", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DarrpOptimize = &tmp
		}
	}
	if v, ok := d.GetOk("darrp_optimize_schedules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("darrp_optimize_schedules", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerSettingDarrpOptimizeSchedules(d, v, "darrp_optimize_schedules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DarrpOptimizeSchedules = t
		}
	} else if d.HasChange("darrp_optimize_schedules") {
		old, new := d.GetChange("darrp_optimize_schedules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DarrpOptimizeSchedules = &[]models.WirelessControllerSettingDarrpOptimizeSchedules{}
		}
	}
	if v1, ok := d.GetOk("device_holdoff"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("device_holdoff", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeviceHoldoff = &tmp
		}
	}
	if v1, ok := d.GetOk("device_idle"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("device_idle", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeviceIdle = &tmp
		}
	}
	if v1, ok := d.GetOk("device_weight"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("device_weight", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeviceWeight = &tmp
		}
	}
	if v1, ok := d.GetOk("duplicate_ssid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("duplicate_ssid", sv)
				diags = append(diags, e)
			}
			obj.DuplicateSsid = &v2
		}
	}
	if v1, ok := d.GetOk("fake_ssid_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fake_ssid_action", sv)
				diags = append(diags, e)
			}
			obj.FakeSsidAction = &v2
		}
	}
	if v1, ok := d.GetOk("fapc_compatibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fapc_compatibility", sv)
				diags = append(diags, e)
			}
			obj.FapcCompatibility = &v2
		}
	}
	if v1, ok := d.GetOk("firmware_provision_on_authorization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("firmware_provision_on_authorization", sv)
				diags = append(diags, e)
			}
			obj.FirmwareProvisionOnAuthorization = &v2
		}
	}
	if v, ok := d.GetOk("offending_ssid"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("offending_ssid", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerSettingOffendingSsid(d, v, "offending_ssid", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OffendingSsid = t
		}
	} else if d.HasChange("offending_ssid") {
		old, new := d.GetChange("offending_ssid")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OffendingSsid = &[]models.WirelessControllerSettingOffendingSsid{}
		}
	}
	if v1, ok := d.GetOk("phishing_ssid_detect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("phishing_ssid_detect", sv)
				diags = append(diags, e)
			}
			obj.PhishingSsidDetect = &v2
		}
	}
	if v1, ok := d.GetOk("wfa_compatibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wfa_compatibility", sv)
				diags = append(diags, e)
			}
			obj.WfaCompatibility = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWirelessControllerSetting(d *schema.ResourceData, sv string) (*models.WirelessControllerSetting, diag.Diagnostics) {
	obj := models.WirelessControllerSetting{}
	diags := diag.Diagnostics{}

	obj.DarrpOptimizeSchedules = &[]models.WirelessControllerSettingDarrpOptimizeSchedules{}
	obj.OffendingSsid = &[]models.WirelessControllerSettingOffendingSsid{}

	return &obj, diags
}
