// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure alarm.

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

func resourceSystemAlarm() *schema.Resource {
	return &schema.Resource{
		Description: "Configure alarm.",

		CreateContext: resourceSystemAlarmCreate,
		ReadContext:   resourceSystemAlarmRead,
		UpdateContext: resourceSystemAlarmUpdate,
		DeleteContext: resourceSystemAlarmDelete,

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
			"audible": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable audible alarm.",
				Optional:    true,
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Alarm groups.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_auth_failure_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Admin authentication failure threshold.",
							Optional:    true,
							Computed:    true,
						},
						"admin_auth_lockout_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Admin authentication lockout threshold.",
							Optional:    true,
							Computed:    true,
						},
						"decryption_failure_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Decryption failure threshold.",
							Optional:    true,
							Computed:    true,
						},
						"encryption_failure_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Encryption failure threshold.",
							Optional:    true,
							Computed:    true,
						},
						"fw_policy_id": {
							Type: schema.TypeInt,

							Description: "Firewall policy ID.",
							Optional:    true,
							Computed:    true,
						},
						"fw_policy_id_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Firewall policy ID threshold.",
							Optional:    true,
							Computed:    true,
						},
						"fw_policy_violations": {
							Type:        schema.TypeList,
							Description: "Firewall policy violations.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "Destination IP (0=all).",
										Optional:    true,
										Computed:    true,
									},
									"dst_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Destination port (0=all).",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Firewall policy violations ID.",
										Optional:    true,
										Computed:    true,
									},
									"src_ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "Source IP (0=all).",
										Optional:    true,
										Computed:    true,
									},
									"src_port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),

										Description: "Source port (0=all).",
										Optional:    true,
										Computed:    true,
									},
									"threshold": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 1024),

										Description: "Firewall policy violation threshold.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Group ID.",
							Optional:    true,
							Computed:    true,
						},
						"log_full_warning_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Log full warning threshold.",
							Optional:    true,
							Computed:    true,
						},
						"period": {
							Type: schema.TypeInt,

							Description: "Time period in seconds (0 = from start up).",
							Optional:    true,
							Computed:    true,
						},
						"replay_attempt_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "Replay attempt threshold.",
							Optional:    true,
							Computed:    true,
						},
						"self_test_failure_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),

							Description: "Self-test failure threshold.",
							Optional:    true,
							Computed:    true,
						},
						"user_auth_failure_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "User authentication failure threshold.",
							Optional:    true,
							Computed:    true,
						},
						"user_auth_lockout_threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),

							Description: "User authentication lockout threshold.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable alarm.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAlarmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAlarm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAlarm(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAlarm")
	}

	return resourceSystemAlarmRead(ctx, d, meta)
}

func resourceSystemAlarmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAlarm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAlarm(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAlarm resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAlarm")
	}

	return resourceSystemAlarmRead(ctx, d, meta)
}

func resourceSystemAlarmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemAlarm(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemAlarm(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAlarm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAlarmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAlarm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAlarm resource: %v", err)
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

	diags := refreshObjectSystemAlarm(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAlarmGroups(d *schema.ResourceData, v *[]models.SystemAlarmGroups, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AdminAuthFailureThreshold; tmp != nil {
				v["admin_auth_failure_threshold"] = *tmp
			}

			if tmp := cfg.AdminAuthLockoutThreshold; tmp != nil {
				v["admin_auth_lockout_threshold"] = *tmp
			}

			if tmp := cfg.DecryptionFailureThreshold; tmp != nil {
				v["decryption_failure_threshold"] = *tmp
			}

			if tmp := cfg.EncryptionFailureThreshold; tmp != nil {
				v["encryption_failure_threshold"] = *tmp
			}

			if tmp := cfg.FwPolicyId; tmp != nil {
				v["fw_policy_id"] = *tmp
			}

			if tmp := cfg.FwPolicyIdThreshold; tmp != nil {
				v["fw_policy_id_threshold"] = *tmp
			}

			if tmp := cfg.FwPolicyViolations; tmp != nil {
				v["fw_policy_violations"] = flattenSystemAlarmGroupsFwPolicyViolations(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "fw_policy_violations"), sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.LogFullWarningThreshold; tmp != nil {
				v["log_full_warning_threshold"] = *tmp
			}

			if tmp := cfg.Period; tmp != nil {
				v["period"] = *tmp
			}

			if tmp := cfg.ReplayAttemptThreshold; tmp != nil {
				v["replay_attempt_threshold"] = *tmp
			}

			if tmp := cfg.SelfTestFailureThreshold; tmp != nil {
				v["self_test_failure_threshold"] = *tmp
			}

			if tmp := cfg.UserAuthFailureThreshold; tmp != nil {
				v["user_auth_failure_threshold"] = *tmp
			}

			if tmp := cfg.UserAuthLockoutThreshold; tmp != nil {
				v["user_auth_lockout_threshold"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemAlarmGroupsFwPolicyViolations(d *schema.ResourceData, v *[]models.SystemAlarmGroupsFwPolicyViolations, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DstIp; tmp != nil {
				v["dst_ip"] = *tmp
			}

			if tmp := cfg.DstPort; tmp != nil {
				v["dst_port"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.SrcIp; tmp != nil {
				v["src_ip"] = *tmp
			}

			if tmp := cfg.SrcPort; tmp != nil {
				v["src_port"] = *tmp
			}

			if tmp := cfg.Threshold; tmp != nil {
				v["threshold"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemAlarm(d *schema.ResourceData, o *models.SystemAlarm, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Audible != nil {
		v := *o.Audible

		if err = d.Set("audible", v); err != nil {
			return diag.Errorf("error reading audible: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenSystemAlarmGroups(d, o.Groups, "groups", sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func expandSystemAlarmGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAlarmGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAlarmGroups

	for i := range l {
		tmp := models.SystemAlarmGroups{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.admin_auth_failure_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdminAuthFailureThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.admin_auth_lockout_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AdminAuthLockoutThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.decryption_failure_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DecryptionFailureThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.encryption_failure_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.EncryptionFailureThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fw_policy_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FwPolicyId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fw_policy_id_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FwPolicyIdThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fw_policy_violations", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemAlarmGroupsFwPolicyViolations(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemAlarmGroupsFwPolicyViolations
			// 	}
			tmp.FwPolicyViolations = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_full_warning_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.LogFullWarningThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.period", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Period = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.replay_attempt_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ReplayAttemptThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.self_test_failure_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SelfTestFailureThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.user_auth_failure_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UserAuthFailureThreshold = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.user_auth_lockout_threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UserAuthLockoutThreshold = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAlarmGroupsFwPolicyViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAlarmGroupsFwPolicyViolations, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAlarmGroupsFwPolicyViolations

	for i := range l {
		tmp := models.SystemAlarmGroupsFwPolicyViolations{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dst_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DstIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dst_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.DstPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SrcIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SrcPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Threshold = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemAlarm(d *schema.ResourceData, sv string) (*models.SystemAlarm, diag.Diagnostics) {
	obj := models.SystemAlarm{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("audible"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("audible", sv)
				diags = append(diags, e)
			}
			obj.Audible = &v2
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAlarmGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.SystemAlarmGroups{}
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemAlarm(d *schema.ResourceData, sv string) (*models.SystemAlarm, diag.Diagnostics) {
	obj := models.SystemAlarm{}
	diags := diag.Diagnostics{}

	obj.Groups = &[]models.SystemAlarmGroups{}

	return &obj, diags
}
