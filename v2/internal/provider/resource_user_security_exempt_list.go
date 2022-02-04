// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure security exemption list.

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

func resourceUserSecurityExemptList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure security exemption list.",

		CreateContext: resourceUserSecurityExemptListCreate,
		ReadContext:   resourceUserSecurityExemptListRead,
		UpdateContext: resourceUserSecurityExemptListUpdate,
		DeleteContext: resourceUserSecurityExemptListDelete,

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
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the exempt list.",
				ForceNew:    true,
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "Configure rules for exempting users from captive portal authentication.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination addresses or address groups.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address or group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeList,
							Description: "Destination services.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Service name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"srcaddr": {
							Type:        schema.TypeList,
							Description: "Source addresses or address groups.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address or group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceUserSecurityExemptListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserSecurityExemptList resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserSecurityExemptList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserSecurityExemptList(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserSecurityExemptList")
	}

	return resourceUserSecurityExemptListRead(ctx, d, meta)
}

func resourceUserSecurityExemptListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserSecurityExemptList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserSecurityExemptList(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserSecurityExemptList resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserSecurityExemptList")
	}

	return resourceUserSecurityExemptListRead(ctx, d, meta)
}

func resourceUserSecurityExemptListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserSecurityExemptList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserSecurityExemptList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSecurityExemptListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserSecurityExemptList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserSecurityExemptList resource: %v", err)
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

	diags := refreshObjectUserSecurityExemptList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserSecurityExemptListRule(v *[]models.UserSecurityExemptListRule, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Dstaddr; tmp != nil {
				v["dstaddr"] = flattenUserSecurityExemptListRuleDstaddr(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Service; tmp != nil {
				v["service"] = flattenUserSecurityExemptListRuleService(tmp, sort)
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				v["srcaddr"] = flattenUserSecurityExemptListRuleSrcaddr(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenUserSecurityExemptListRuleDstaddr(v *[]models.UserSecurityExemptListRuleDstaddr, sort bool) interface{} {
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

func flattenUserSecurityExemptListRuleService(v *[]models.UserSecurityExemptListRuleService, sort bool) interface{} {
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

func flattenUserSecurityExemptListRuleSrcaddr(v *[]models.UserSecurityExemptListRuleSrcaddr, sort bool) interface{} {
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

func refreshObjectUserSecurityExemptList(d *schema.ResourceData, o *models.UserSecurityExemptList, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Rule != nil {
		if err = d.Set("rule", flattenUserSecurityExemptListRule(o.Rule, sort)); err != nil {
			return diag.Errorf("error reading rule: %v", err)
		}
	}

	return nil
}

func expandUserSecurityExemptListRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserSecurityExemptListRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserSecurityExemptListRule

	for i := range l {
		tmp := models.UserSecurityExemptListRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandUserSecurityExemptListRuleDstaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.UserSecurityExemptListRuleDstaddr
			// 	}
			tmp.Dstaddr = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandUserSecurityExemptListRuleService(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.UserSecurityExemptListRuleService
			// 	}
			tmp.Service = v2

		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandUserSecurityExemptListRuleSrcaddr(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.UserSecurityExemptListRuleSrcaddr
			// 	}
			tmp.Srcaddr = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandUserSecurityExemptListRuleDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserSecurityExemptListRuleDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserSecurityExemptListRuleDstaddr

	for i := range l {
		tmp := models.UserSecurityExemptListRuleDstaddr{}
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

func expandUserSecurityExemptListRuleService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserSecurityExemptListRuleService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserSecurityExemptListRuleService

	for i := range l {
		tmp := models.UserSecurityExemptListRuleService{}
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

func expandUserSecurityExemptListRuleSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserSecurityExemptListRuleSrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserSecurityExemptListRuleSrcaddr

	for i := range l {
		tmp := models.UserSecurityExemptListRuleSrcaddr{}
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

func getObjectUserSecurityExemptList(d *schema.ResourceData, sv string) (*models.UserSecurityExemptList, diag.Diagnostics) {
	obj := models.UserSecurityExemptList{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
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
	if v, ok := d.GetOk("rule"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("rule", sv)
			diags = append(diags, e)
		}
		t, err := expandUserSecurityExemptListRule(d, v, "rule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Rule = t
		}
	} else if d.HasChange("rule") {
		old, new := d.GetChange("rule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Rule = &[]models.UserSecurityExemptListRule{}
		}
	}
	return &obj, diags
}
