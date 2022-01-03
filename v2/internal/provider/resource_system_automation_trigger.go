// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Trigger for automation stitches.

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

func resourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Description: "Trigger for automation stitches.",

		CreateContext: resourceSystemAutomationTriggerCreate,
		ReadContext:   resourceSystemAutomationTriggerRead,
		UpdateContext: resourceSystemAutomationTriggerUpdate,
		DeleteContext: resourceSystemAutomationTriggerDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"event_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ioc", "event-log", "reboot", "low-memory", "high-cpu", "license-near-expiry", "ha-failover", "config-change", "security-rating-summary", "virus-ips-db-updated", "faz-event", "incoming-webhook", "fabric-event"}, false),

				Description: "Event type.",
				Optional:    true,
				Computed:    true,
			},
			"fabric_event_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fabric connector event handler name.",
				Optional:    true,
				Computed:    true,
			},
			"fabric_event_severity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fabric connector event severity.",
				Optional:    true,
				Computed:    true,
			},
			"faz_event_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "FortiAnalyzer event handler name.",
				Optional:    true,
				Computed:    true,
			},
			"faz_event_severity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "FortiAnalyzer event severity.",
				Optional:    true,
				Computed:    true,
			},
			"faz_event_tags": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "FortiAnalyzer event tags.",
				Optional:    true,
				Computed:    true,
			},
			"fields": {
				Type:        schema.TypeList,
				Description: "Customized trigger field settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ioc_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"medium", "high"}, false),

				Description: "IOC threat level.",
				Optional:    true,
				Computed:    true,
			},
			"license_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"forticare-support", "fortiguard-webfilter", "fortiguard-antispam", "fortiguard-antivirus", "fortiguard-ips", "fortiguard-management", "forticloud", "any"}, false),

				Description: "License type.",
				Optional:    true,
				Computed:    true,
			},
			"logid": {
				Type:        schema.TypeList,
				Description: "Log IDs to trigger event.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Log ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"report_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"posture", "coverage", "optimization", "any"}, false),

				Description: "Security Rating report.",
				Optional:    true,
				Computed:    true,
			},
			"serial": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fabric connector serial number.",
				Optional:    true,
				Computed:    true,
			},
			"trigger_day": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),

				Description: "Day within a month to trigger.",
				Optional:    true,
				Computed:    true,
			},
			"trigger_frequency": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hourly", "daily", "weekly", "monthly"}, false),

				Description: "Scheduled trigger frequency (default = daily).",
				Optional:    true,
				Computed:    true,
			},
			"trigger_hour": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),

				Description: "Hour of the day on which to trigger (0 - 23, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"trigger_minute": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 59),

				Description: "Minute of the hour on which to trigger (0 - 59, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"trigger_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"event-based", "scheduled"}, false),

				Description: "Trigger type.",
				Optional:    true,
				Computed:    true,
			},
			"trigger_weekday": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}, false),

				Description: "Day of week for trigger.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemAutomationTriggerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemAutomationTrigger resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemAutomationTrigger(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemAutomationTrigger(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(ctx, d, meta)
}

func resourceSystemAutomationTriggerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemAutomationTrigger(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemAutomationTrigger(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemAutomationTrigger resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(ctx, d, meta)
}

func resourceSystemAutomationTriggerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemAutomationTrigger(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemAutomationTrigger resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationTriggerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemAutomationTrigger(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutomationTrigger resource: %v", err)
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

	diags := refreshObjectSystemAutomationTrigger(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemAutomationTriggerFields(v *[]models.SystemAutomationTriggerFields, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemAutomationTriggerLogid(v *[]models.SystemAutomationTriggerLogid, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func refreshObjectSystemAutomationTrigger(d *schema.ResourceData, o *models.SystemAutomationTrigger, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.EventType != nil {
		v := *o.EventType

		if err = d.Set("event_type", v); err != nil {
			return diag.Errorf("error reading event_type: %v", err)
		}
	}

	if o.FabricEventName != nil {
		v := *o.FabricEventName

		if err = d.Set("fabric_event_name", v); err != nil {
			return diag.Errorf("error reading fabric_event_name: %v", err)
		}
	}

	if o.FabricEventSeverity != nil {
		v := *o.FabricEventSeverity

		if err = d.Set("fabric_event_severity", v); err != nil {
			return diag.Errorf("error reading fabric_event_severity: %v", err)
		}
	}

	if o.FazEventName != nil {
		v := *o.FazEventName

		if err = d.Set("faz_event_name", v); err != nil {
			return diag.Errorf("error reading faz_event_name: %v", err)
		}
	}

	if o.FazEventSeverity != nil {
		v := *o.FazEventSeverity

		if err = d.Set("faz_event_severity", v); err != nil {
			return diag.Errorf("error reading faz_event_severity: %v", err)
		}
	}

	if o.FazEventTags != nil {
		v := *o.FazEventTags

		if err = d.Set("faz_event_tags", v); err != nil {
			return diag.Errorf("error reading faz_event_tags: %v", err)
		}
	}

	if o.Fields != nil {
		if err = d.Set("fields", flattenSystemAutomationTriggerFields(o.Fields, sort)); err != nil {
			return diag.Errorf("error reading fields: %v", err)
		}
	}

	if o.IocLevel != nil {
		v := *o.IocLevel

		if err = d.Set("ioc_level", v); err != nil {
			return diag.Errorf("error reading ioc_level: %v", err)
		}
	}

	if o.LicenseType != nil {
		v := *o.LicenseType

		if err = d.Set("license_type", v); err != nil {
			return diag.Errorf("error reading license_type: %v", err)
		}
	}

	if o.Logid != nil {
		if err = d.Set("logid", flattenSystemAutomationTriggerLogid(o.Logid, sort)); err != nil {
			return diag.Errorf("error reading logid: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ReportType != nil {
		v := *o.ReportType

		if err = d.Set("report_type", v); err != nil {
			return diag.Errorf("error reading report_type: %v", err)
		}
	}

	if o.Serial != nil {
		v := *o.Serial

		if err = d.Set("serial", v); err != nil {
			return diag.Errorf("error reading serial: %v", err)
		}
	}

	if o.TriggerDay != nil {
		v := *o.TriggerDay

		if err = d.Set("trigger_day", v); err != nil {
			return diag.Errorf("error reading trigger_day: %v", err)
		}
	}

	if o.TriggerFrequency != nil {
		v := *o.TriggerFrequency

		if err = d.Set("trigger_frequency", v); err != nil {
			return diag.Errorf("error reading trigger_frequency: %v", err)
		}
	}

	if o.TriggerHour != nil {
		v := *o.TriggerHour

		if err = d.Set("trigger_hour", v); err != nil {
			return diag.Errorf("error reading trigger_hour: %v", err)
		}
	}

	if o.TriggerMinute != nil {
		v := *o.TriggerMinute

		if err = d.Set("trigger_minute", v); err != nil {
			return diag.Errorf("error reading trigger_minute: %v", err)
		}
	}

	if o.TriggerType != nil {
		v := *o.TriggerType

		if err = d.Set("trigger_type", v); err != nil {
			return diag.Errorf("error reading trigger_type: %v", err)
		}
	}

	if o.TriggerWeekday != nil {
		v := *o.TriggerWeekday

		if err = d.Set("trigger_weekday", v); err != nil {
			return diag.Errorf("error reading trigger_weekday: %v", err)
		}
	}

	return nil
}

func expandSystemAutomationTriggerFields(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationTriggerFields, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationTriggerFields

	for i := range l {
		tmp := models.SystemAutomationTriggerFields{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemAutomationTriggerLogid(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemAutomationTriggerLogid, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemAutomationTriggerLogid

	for i := range l {
		tmp := models.SystemAutomationTriggerLogid{}
		var pre_append string

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

func getObjectSystemAutomationTrigger(d *schema.ResourceData, sv string) (*models.SystemAutomationTrigger, diag.Diagnostics) {
	obj := models.SystemAutomationTrigger{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("event_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("event_type", sv)
				diags = append(diags, e)
			}
			obj.EventType = &v2
		}
	}
	if v1, ok := d.GetOk("fabric_event_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("fabric_event_name", sv)
				diags = append(diags, e)
			}
			obj.FabricEventName = &v2
		}
	}
	if v1, ok := d.GetOk("fabric_event_severity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("fabric_event_severity", sv)
				diags = append(diags, e)
			}
			obj.FabricEventSeverity = &v2
		}
	}
	if v1, ok := d.GetOk("faz_event_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("faz_event_name", sv)
				diags = append(diags, e)
			}
			obj.FazEventName = &v2
		}
	}
	if v1, ok := d.GetOk("faz_event_severity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("faz_event_severity", sv)
				diags = append(diags, e)
			}
			obj.FazEventSeverity = &v2
		}
	}
	if v1, ok := d.GetOk("faz_event_tags"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("faz_event_tags", sv)
				diags = append(diags, e)
			}
			obj.FazEventTags = &v2
		}
	}
	if v, ok := d.GetOk("fields"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fields", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationTriggerFields(d, v, "fields", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Fields = t
		}
	} else if d.HasChange("fields") {
		old, new := d.GetChange("fields")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Fields = &[]models.SystemAutomationTriggerFields{}
		}
	}
	if v1, ok := d.GetOk("ioc_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ioc_level", sv)
				diags = append(diags, e)
			}
			obj.IocLevel = &v2
		}
	}
	if v1, ok := d.GetOk("license_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("license_type", sv)
				diags = append(diags, e)
			}
			obj.LicenseType = &v2
		}
	}
	if v, ok := d.GetOk("logid"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("logid", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemAutomationTriggerLogid(d, v, "logid", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Logid = t
		}
	} else if d.HasChange("logid") {
		old, new := d.GetChange("logid")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Logid = &[]models.SystemAutomationTriggerLogid{}
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
	if v1, ok := d.GetOk("report_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("report_type", sv)
				diags = append(diags, e)
			}
			obj.ReportType = &v2
		}
	}
	if v1, ok := d.GetOk("serial"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("serial", sv)
				diags = append(diags, e)
			}
			obj.Serial = &v2
		}
	}
	if v1, ok := d.GetOk("trigger_day"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger_day", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TriggerDay = &tmp
		}
	}
	if v1, ok := d.GetOk("trigger_frequency"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger_frequency", sv)
				diags = append(diags, e)
			}
			obj.TriggerFrequency = &v2
		}
	}
	if v1, ok := d.GetOk("trigger_hour"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger_hour", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TriggerHour = &tmp
		}
	}
	if v1, ok := d.GetOk("trigger_minute"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger_minute", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TriggerMinute = &tmp
		}
	}
	if v1, ok := d.GetOk("trigger_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger_type", sv)
				diags = append(diags, e)
			}
			obj.TriggerType = &v2
		}
	}
	if v1, ok := d.GetOk("trigger_weekday"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trigger_weekday", sv)
				diags = append(diags, e)
			}
			obj.TriggerWeekday = &v2
		}
	}
	return &obj, diags
}
