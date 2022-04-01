// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS rules.

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

func resourceIpsRule() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS rules.",

		CreateContext: resourceIpsRuleCreate,
		ReadContext:   resourceIpsRuleRead,
		UpdateContext: resourceIpsRuleUpdate,
		DeleteContext: resourceIpsRuleDelete,

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
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "block"}, false),

				Description: "Action.",
				Optional:    true,
				Computed:    true,
			},
			"application": {
				Type: schema.TypeString,

				Description: "Vulnerable applications.",
				Optional:    true,
				Computed:    true,
			},
			"date": {
				Type: schema.TypeInt,

				Description: "Date.",
				Optional:    true,
				Computed:    true,
			},
			"group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Group.",
				Optional:    true,
				Computed:    true,
			},
			"location": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffMultiStringEqual,
				Description:      "Vulnerable location.",
				Optional:         true,
				Computed:         true,
			},
			"log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging.",
				Optional:    true,
				Computed:    true,
			},
			"log_packet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable packet logging.",
				Optional:    true,
				Computed:    true,
			},
			"metadata": {
				Type:        schema.TypeList,
				Description: "Meta data.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"metaid": {
							Type: schema.TypeInt,

							Description: "Meta ID.",
							Optional:    true,
							Computed:    true,
						},
						"valueid": {
							Type: schema.TypeInt,

							Description: "Value ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Rule name.",
				ForceNew:    true,
				Required:    true,
			},
			"os": {
				Type: schema.TypeString,

				Description: "Vulnerable operation systems.",
				Optional:    true,
				Computed:    true,
			},
			"rev": {
				Type: schema.TypeInt,

				Description: "Revision.",
				Optional:    true,
				Computed:    true,
			},
			"rule_id": {
				Type: schema.TypeInt,

				Description: "Rule ID.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type: schema.TypeString,

				Description: "Vulnerable service.",
				Optional:    true,
				Computed:    true,
			},
			"severity": {
				Type: schema.TypeString,

				Description: "Severity.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable status.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceIpsRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating IpsRule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectIpsRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateIpsRule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsRule")
	}

	return resourceIpsRuleRead(ctx, d, meta)
}

func resourceIpsRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectIpsRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateIpsRule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating IpsRule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("IpsRule")
	}

	return resourceIpsRuleRead(ctx, d, meta)
}

func resourceIpsRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteIpsRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting IpsRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadIpsRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsRule resource: %v", err)
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

	diags := refreshObjectIpsRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenIpsRuleMetadata(d *schema.ResourceData, v *[]models.IpsRuleMetadata, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Metaid; tmp != nil {
				v["metaid"] = *tmp
			}

			if tmp := cfg.Valueid; tmp != nil {
				v["valueid"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectIpsRule(d *schema.ResourceData, o *models.IpsRule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Application != nil {
		v := *o.Application

		if err = d.Set("application", v); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.Date != nil {
		v := *o.Date

		if err = d.Set("date", v); err != nil {
			return diag.Errorf("error reading date: %v", err)
		}
	}

	if o.Group != nil {
		v := *o.Group

		if err = d.Set("group", v); err != nil {
			return diag.Errorf("error reading group: %v", err)
		}
	}

	if o.Location != nil {
		v := *o.Location

		if err = d.Set("location", v); err != nil {
			return diag.Errorf("error reading location: %v", err)
		}
	}

	if o.Log != nil {
		v := *o.Log

		if err = d.Set("log", v); err != nil {
			return diag.Errorf("error reading log: %v", err)
		}
	}

	if o.LogPacket != nil {
		v := *o.LogPacket

		if err = d.Set("log_packet", v); err != nil {
			return diag.Errorf("error reading log_packet: %v", err)
		}
	}

	if o.Metadata != nil {
		if err = d.Set("metadata", flattenIpsRuleMetadata(d, o.Metadata, "metadata", sort)); err != nil {
			return diag.Errorf("error reading metadata: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Os != nil {
		v := *o.Os

		if err = d.Set("os", v); err != nil {
			return diag.Errorf("error reading os: %v", err)
		}
	}

	if o.Rev != nil {
		v := *o.Rev

		if err = d.Set("rev", v); err != nil {
			return diag.Errorf("error reading rev: %v", err)
		}
	}

	if o.RuleId != nil {
		v := *o.RuleId

		if err = d.Set("rule_id", v); err != nil {
			return diag.Errorf("error reading rule_id: %v", err)
		}
	}

	if o.Service != nil {
		v := *o.Service

		if err = d.Set("service", v); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.Severity != nil {
		v := *o.Severity

		if err = d.Set("severity", v); err != nil {
			return diag.Errorf("error reading severity: %v", err)
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

func expandIpsRuleMetadata(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.IpsRuleMetadata, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.IpsRuleMetadata

	for i := range l {
		tmp := models.IpsRuleMetadata{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.metaid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Metaid = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.valueid", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Valueid = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectIpsRule(d *schema.ResourceData, sv string) (*models.IpsRule, diag.Diagnostics) {
	obj := models.IpsRule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("application"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application", sv)
				diags = append(diags, e)
			}
			obj.Application = &v2
		}
	}
	if v1, ok := d.GetOk("date"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("date", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Date = &tmp
		}
	}
	if v1, ok := d.GetOk("group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("group", sv)
				diags = append(diags, e)
			}
			obj.Group = &v2
		}
	}
	if v1, ok := d.GetOk("location"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("location", sv)
				diags = append(diags, e)
			}
			obj.Location = &v2
		}
	}
	if v1, ok := d.GetOk("log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log", sv)
				diags = append(diags, e)
			}
			obj.Log = &v2
		}
	}
	if v1, ok := d.GetOk("log_packet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_packet", sv)
				diags = append(diags, e)
			}
			obj.LogPacket = &v2
		}
	}
	if v, ok := d.GetOk("metadata"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("metadata", sv)
			diags = append(diags, e)
		}
		t, err := expandIpsRuleMetadata(d, v, "metadata", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Metadata = t
		}
	} else if d.HasChange("metadata") {
		old, new := d.GetChange("metadata")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Metadata = &[]models.IpsRuleMetadata{}
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
	if v1, ok := d.GetOk("os"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("os", sv)
				diags = append(diags, e)
			}
			obj.Os = &v2
		}
	}
	if v1, ok := d.GetOk("rev"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rev", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Rev = &tmp
		}
	}
	if v1, ok := d.GetOk("rule_id"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rule_id", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RuleId = &tmp
		}
	}
	if v1, ok := d.GetOk("service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service", sv)
				diags = append(diags, e)
			}
			obj.Service = &v2
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
