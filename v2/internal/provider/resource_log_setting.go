// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure general log settings.

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

func resourceLogSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure general log settings.",

		CreateContext: resourceLogSettingCreate,
		ReadContext:   resourceLogSettingRead,
		UpdateContext: resourceLogSettingUpdate,
		DeleteContext: resourceLogSettingDelete,

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
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"anonymization_hash": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "User name anonymization hash salt.",
				Optional:    true,
				Computed:    true,
			},
			"brief_traffic_format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable brief format traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"custom_log_fields": {
				Type:        schema.TypeList,
				Description: "Custom fields to append to all log messages.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Custom log field.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"daemon_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable daemon logging.",
				Optional:    true,
				Computed:    true,
			},
			"expolicy_implicit_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable explicit proxy firewall implicit policy logging.",
				Optional:    true,
				Computed:    true,
			},
			"faz_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable override FortiAnalyzer settings.",
				Optional:    true,
				Computed:    true,
			},
			"fortiview_weekly_data": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiView weekly data.",
				Optional:    true,
				Computed:    true,
			},
			"fwpolicy_implicit_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable implicit firewall policy logging.",
				Optional:    true,
				Computed:    true,
			},
			"fwpolicy6_implicit_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable implicit firewall policy6 logging.",
				Optional:    true,
				Computed:    true,
			},
			"local_in_allow": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local-in-allow logging.",
				Optional:    true,
				Computed:    true,
			},
			"local_in_deny_broadcast": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local-in-deny-broadcast logging.",
				Optional:    true,
				Computed:    true,
			},
			"local_in_deny_unicast": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local-in-deny-unicast logging.",
				Optional:    true,
				Computed:    true,
			},
			"local_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local-out logging.",
				Optional:    true,
				Computed:    true,
			},
			"log_invalid_packet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable invalid packet traffic logging.",
				Optional:    true,
				Computed:    true,
			},
			"log_policy_comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable inserting policy comments into traffic logs.",
				Optional:    true,
				Computed:    true,
			},
			"log_policy_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable inserting policy name into traffic logs.",
				Optional:    true,
				Computed:    true,
			},
			"log_user_in_upper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logs with user-in-upper.",
				Optional:    true,
				Computed:    true,
			},
			"neighbor_event": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable neighbor event logging.",
				Optional:    true,
				Computed:    true,
			},
			"resolve_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adding resolved domain names to traffic logs if possible.",
				Optional:    true,
				Computed:    true,
			},
			"resolve_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adding resolved service names to traffic logs.",
				Optional:    true,
				Computed:    true,
			},
			"syslog_override": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable override Syslog settings.",
				Optional:    true,
				Computed:    true,
			},
			"user_anonymize": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable anonymizing user names in log messages.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceLogSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateLogSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogSetting")
	}

	return resourceLogSettingRead(ctx, d, meta)
}

func resourceLogSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectLogSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateLogSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating LogSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("LogSetting")
	}

	return resourceLogSettingRead(ctx, d, meta)
}

func resourceLogSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteLogSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting LogSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadLogSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogSetting resource: %v", err)
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

	diags := refreshObjectLogSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenLogSettingCustomLogFields(v *[]models.LogSettingCustomLogFields, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.FieldId; tmp != nil {
				v["field_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "field_id")
	}

	return flat
}

func refreshObjectLogSetting(d *schema.ResourceData, o *models.LogSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AnonymizationHash != nil {
		v := *o.AnonymizationHash

		if err = d.Set("anonymization_hash", v); err != nil {
			return diag.Errorf("error reading anonymization_hash: %v", err)
		}
	}

	if o.BriefTrafficFormat != nil {
		v := *o.BriefTrafficFormat

		if err = d.Set("brief_traffic_format", v); err != nil {
			return diag.Errorf("error reading brief_traffic_format: %v", err)
		}
	}

	if o.CustomLogFields != nil {
		if err = d.Set("custom_log_fields", flattenLogSettingCustomLogFields(o.CustomLogFields, sort)); err != nil {
			return diag.Errorf("error reading custom_log_fields: %v", err)
		}
	}

	if o.DaemonLog != nil {
		v := *o.DaemonLog

		if err = d.Set("daemon_log", v); err != nil {
			return diag.Errorf("error reading daemon_log: %v", err)
		}
	}

	if o.ExpolicyImplicitLog != nil {
		v := *o.ExpolicyImplicitLog

		if err = d.Set("expolicy_implicit_log", v); err != nil {
			return diag.Errorf("error reading expolicy_implicit_log: %v", err)
		}
	}

	if o.FazOverride != nil {
		v := *o.FazOverride

		if err = d.Set("faz_override", v); err != nil {
			return diag.Errorf("error reading faz_override: %v", err)
		}
	}

	if o.FortiviewWeeklyData != nil {
		v := *o.FortiviewWeeklyData

		if err = d.Set("fortiview_weekly_data", v); err != nil {
			return diag.Errorf("error reading fortiview_weekly_data: %v", err)
		}
	}

	if o.FwpolicyImplicitLog != nil {
		v := *o.FwpolicyImplicitLog

		if err = d.Set("fwpolicy_implicit_log", v); err != nil {
			return diag.Errorf("error reading fwpolicy_implicit_log: %v", err)
		}
	}

	if o.Fwpolicy6ImplicitLog != nil {
		v := *o.Fwpolicy6ImplicitLog

		if err = d.Set("fwpolicy6_implicit_log", v); err != nil {
			return diag.Errorf("error reading fwpolicy6_implicit_log: %v", err)
		}
	}

	if o.LocalInAllow != nil {
		v := *o.LocalInAllow

		if err = d.Set("local_in_allow", v); err != nil {
			return diag.Errorf("error reading local_in_allow: %v", err)
		}
	}

	if o.LocalInDenyBroadcast != nil {
		v := *o.LocalInDenyBroadcast

		if err = d.Set("local_in_deny_broadcast", v); err != nil {
			return diag.Errorf("error reading local_in_deny_broadcast: %v", err)
		}
	}

	if o.LocalInDenyUnicast != nil {
		v := *o.LocalInDenyUnicast

		if err = d.Set("local_in_deny_unicast", v); err != nil {
			return diag.Errorf("error reading local_in_deny_unicast: %v", err)
		}
	}

	if o.LocalOut != nil {
		v := *o.LocalOut

		if err = d.Set("local_out", v); err != nil {
			return diag.Errorf("error reading local_out: %v", err)
		}
	}

	if o.LogInvalidPacket != nil {
		v := *o.LogInvalidPacket

		if err = d.Set("log_invalid_packet", v); err != nil {
			return diag.Errorf("error reading log_invalid_packet: %v", err)
		}
	}

	if o.LogPolicyComment != nil {
		v := *o.LogPolicyComment

		if err = d.Set("log_policy_comment", v); err != nil {
			return diag.Errorf("error reading log_policy_comment: %v", err)
		}
	}

	if o.LogPolicyName != nil {
		v := *o.LogPolicyName

		if err = d.Set("log_policy_name", v); err != nil {
			return diag.Errorf("error reading log_policy_name: %v", err)
		}
	}

	if o.LogUserInUpper != nil {
		v := *o.LogUserInUpper

		if err = d.Set("log_user_in_upper", v); err != nil {
			return diag.Errorf("error reading log_user_in_upper: %v", err)
		}
	}

	if o.NeighborEvent != nil {
		v := *o.NeighborEvent

		if err = d.Set("neighbor_event", v); err != nil {
			return diag.Errorf("error reading neighbor_event: %v", err)
		}
	}

	if o.ResolveIp != nil {
		v := *o.ResolveIp

		if err = d.Set("resolve_ip", v); err != nil {
			return diag.Errorf("error reading resolve_ip: %v", err)
		}
	}

	if o.ResolvePort != nil {
		v := *o.ResolvePort

		if err = d.Set("resolve_port", v); err != nil {
			return diag.Errorf("error reading resolve_port: %v", err)
		}
	}

	if o.SyslogOverride != nil {
		v := *o.SyslogOverride

		if err = d.Set("syslog_override", v); err != nil {
			return diag.Errorf("error reading syslog_override: %v", err)
		}
	}

	if o.UserAnonymize != nil {
		v := *o.UserAnonymize

		if err = d.Set("user_anonymize", v); err != nil {
			return diag.Errorf("error reading user_anonymize: %v", err)
		}
	}

	return nil
}

func expandLogSettingCustomLogFields(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.LogSettingCustomLogFields, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.LogSettingCustomLogFields

	for i := range l {
		tmp := models.LogSettingCustomLogFields{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.field_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FieldId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectLogSetting(d *schema.ResourceData, sv string) (*models.LogSetting, diag.Diagnostics) {
	obj := models.LogSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("anonymization_hash"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("anonymization_hash", sv)
				diags = append(diags, e)
			}
			obj.AnonymizationHash = &v2
		}
	}
	if v1, ok := d.GetOk("brief_traffic_format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("brief_traffic_format", sv)
				diags = append(diags, e)
			}
			obj.BriefTrafficFormat = &v2
		}
	}
	if v, ok := d.GetOk("custom_log_fields"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("custom_log_fields", sv)
			diags = append(diags, e)
		}
		t, err := expandLogSettingCustomLogFields(d, v, "custom_log_fields", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CustomLogFields = t
		}
	} else if d.HasChange("custom_log_fields") {
		old, new := d.GetChange("custom_log_fields")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CustomLogFields = &[]models.LogSettingCustomLogFields{}
		}
	}
	if v1, ok := d.GetOk("daemon_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("daemon_log", sv)
				diags = append(diags, e)
			}
			obj.DaemonLog = &v2
		}
	}
	if v1, ok := d.GetOk("expolicy_implicit_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expolicy_implicit_log", sv)
				diags = append(diags, e)
			}
			obj.ExpolicyImplicitLog = &v2
		}
	}
	if v1, ok := d.GetOk("faz_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("faz_override", sv)
				diags = append(diags, e)
			}
			obj.FazOverride = &v2
		}
	}
	if v1, ok := d.GetOk("fortiview_weekly_data"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("fortiview_weekly_data", sv)
				diags = append(diags, e)
			}
			obj.FortiviewWeeklyData = &v2
		}
	}
	if v1, ok := d.GetOk("fwpolicy_implicit_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fwpolicy_implicit_log", sv)
				diags = append(diags, e)
			}
			obj.FwpolicyImplicitLog = &v2
		}
	}
	if v1, ok := d.GetOk("fwpolicy6_implicit_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fwpolicy6_implicit_log", sv)
				diags = append(diags, e)
			}
			obj.Fwpolicy6ImplicitLog = &v2
		}
	}
	if v1, ok := d.GetOk("local_in_allow"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_in_allow", sv)
				diags = append(diags, e)
			}
			obj.LocalInAllow = &v2
		}
	}
	if v1, ok := d.GetOk("local_in_deny_broadcast"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_in_deny_broadcast", sv)
				diags = append(diags, e)
			}
			obj.LocalInDenyBroadcast = &v2
		}
	}
	if v1, ok := d.GetOk("local_in_deny_unicast"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_in_deny_unicast", sv)
				diags = append(diags, e)
			}
			obj.LocalInDenyUnicast = &v2
		}
	}
	if v1, ok := d.GetOk("local_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_out", sv)
				diags = append(diags, e)
			}
			obj.LocalOut = &v2
		}
	}
	if v1, ok := d.GetOk("log_invalid_packet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_invalid_packet", sv)
				diags = append(diags, e)
			}
			obj.LogInvalidPacket = &v2
		}
	}
	if v1, ok := d.GetOk("log_policy_comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_policy_comment", sv)
				diags = append(diags, e)
			}
			obj.LogPolicyComment = &v2
		}
	}
	if v1, ok := d.GetOk("log_policy_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("log_policy_name", sv)
				diags = append(diags, e)
			}
			obj.LogPolicyName = &v2
		}
	}
	if v1, ok := d.GetOk("log_user_in_upper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_user_in_upper", sv)
				diags = append(diags, e)
			}
			obj.LogUserInUpper = &v2
		}
	}
	if v1, ok := d.GetOk("neighbor_event"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("neighbor_event", sv)
				diags = append(diags, e)
			}
			obj.NeighborEvent = &v2
		}
	}
	if v1, ok := d.GetOk("resolve_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resolve_ip", sv)
				diags = append(diags, e)
			}
			obj.ResolveIp = &v2
		}
	}
	if v1, ok := d.GetOk("resolve_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resolve_port", sv)
				diags = append(diags, e)
			}
			obj.ResolvePort = &v2
		}
	}
	if v1, ok := d.GetOk("syslog_override"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("syslog_override", sv)
				diags = append(diags, e)
			}
			obj.SyslogOverride = &v2
		}
	}
	if v1, ok := d.GetOk("user_anonymize"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_anonymize", sv)
				diags = append(diags, e)
			}
			obj.UserAnonymize = &v2
		}
	}
	return &obj, diags
}
