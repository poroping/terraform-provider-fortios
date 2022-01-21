// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NAC policy matching pattern to identify matching NAC devices.

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

func resourceUserNacPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NAC policy matching pattern to identify matching NAC devices.",

		CreateContext: resourceUserNacPolicyCreate,
		ReadContext:   resourceUserNacPolicyRead,
		UpdateContext: resourceUserNacPolicyUpdate,
		DeleteContext: resourceUserNacPolicyDelete,

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
			"category": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"device", "firewall-user", "ems-tag"}, false),

				Description: "Category of NAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Description for the NAC policy matching pattern.",
				Optional:    true,
				Computed:    true,
			},
			"ems_tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "NAC policy matching EMS tag.",
				Optional:    true,
				Computed:    true,
			},
			"family": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "NAC policy matching family.",
				Optional:    true,
				Computed:    true,
			},
			"firewall_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Dynamic firewall address to associate MAC which match this policy.",
				Optional:    true,
				Computed:    true,
			},
			"host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "NAC policy matching host.",
				Optional:    true,
				Computed:    true,
			},
			"hw_vendor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "NAC policy matching hardware vendor.",
				Optional:    true,
				Computed:    true,
			},
			"hw_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "NAC policy matching hardware version.",
				Optional:    true,
				Computed:    true,
			},
			"mac": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 17),

				Description: "NAC policy matching MAC address.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "NAC policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"os": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "NAC policy matching operating system.",
				Optional:    true,
				Computed:    true,
			},
			"src": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "NAC policy matching source.",
				Optional:    true,
				Computed:    true,
			},
			"ssid_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SSID policy to be applied on the matched NAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"sw_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "NAC policy matching software version.",
				Optional:    true,
				Computed:    true,
			},
			"switch_auto_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "disable", "enable"}, false),

				Description: "NAC device auto authorization when discovered and nac-policy matched.",
				Optional:    true,
				Computed:    true,
			},
			"switch_fortilink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "FortiLink interface for which this NAC policy belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"switch_group": {
				Type:        schema.TypeList,
				Description: "List of managed FortiSwitch groups on which NAC policy can be applied.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Managed FortiSwitch group name from available options.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"switch_mac_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "switch-mac-policy action to be applied on the matched NAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"switch_port_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "switch-port-policy to be applied on the matched NAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"switch_scope": {
				Type:        schema.TypeList,
				Description: "List of managed FortiSwitches on which NAC policy can be applied.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Managed FortiSwitch name from available options.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "NAC policy matching type.",
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "NAC policy matching user.",
				Optional:    true,
				Computed:    true,
			},
			"user_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "NAC policy matching user group.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserNacPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserNacPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserNacPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserNacPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserNacPolicy")
	}

	return resourceUserNacPolicyRead(ctx, d, meta)
}

func resourceUserNacPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserNacPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserNacPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserNacPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserNacPolicy")
	}

	return resourceUserNacPolicyRead(ctx, d, meta)
}

func resourceUserNacPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserNacPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserNacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserNacPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserNacPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserNacPolicy resource: %v", err)
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

	diags := refreshObjectUserNacPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenUserNacPolicySwitchGroup(v *[]models.UserNacPolicySwitchGroup, sort bool) interface{} {
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

func flattenUserNacPolicySwitchScope(v *[]models.UserNacPolicySwitchScope, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.SwitchId; tmp != nil {
				v["switch_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "switch_id")
	}

	return flat
}

func refreshObjectUserNacPolicy(d *schema.ResourceData, o *models.UserNacPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Category != nil {
		v := *o.Category

		if err = d.Set("category", v); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.EmsTag != nil {
		v := *o.EmsTag

		if err = d.Set("ems_tag", v); err != nil {
			return diag.Errorf("error reading ems_tag: %v", err)
		}
	}

	if o.Family != nil {
		v := *o.Family

		if err = d.Set("family", v); err != nil {
			return diag.Errorf("error reading family: %v", err)
		}
	}

	if o.FirewallAddress != nil {
		v := *o.FirewallAddress

		if err = d.Set("firewall_address", v); err != nil {
			return diag.Errorf("error reading firewall_address: %v", err)
		}
	}

	if o.Host != nil {
		v := *o.Host

		if err = d.Set("host", v); err != nil {
			return diag.Errorf("error reading host: %v", err)
		}
	}

	if o.HwVendor != nil {
		v := *o.HwVendor

		if err = d.Set("hw_vendor", v); err != nil {
			return diag.Errorf("error reading hw_vendor: %v", err)
		}
	}

	if o.HwVersion != nil {
		v := *o.HwVersion

		if err = d.Set("hw_version", v); err != nil {
			return diag.Errorf("error reading hw_version: %v", err)
		}
	}

	if o.Mac != nil {
		v := *o.Mac

		if err = d.Set("mac", v); err != nil {
			return diag.Errorf("error reading mac: %v", err)
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

	if o.Src != nil {
		v := *o.Src

		if err = d.Set("src", v); err != nil {
			return diag.Errorf("error reading src: %v", err)
		}
	}

	if o.SsidPolicy != nil {
		v := *o.SsidPolicy

		if err = d.Set("ssid_policy", v); err != nil {
			return diag.Errorf("error reading ssid_policy: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.SwVersion != nil {
		v := *o.SwVersion

		if err = d.Set("sw_version", v); err != nil {
			return diag.Errorf("error reading sw_version: %v", err)
		}
	}

	if o.SwitchAutoAuth != nil {
		v := *o.SwitchAutoAuth

		if err = d.Set("switch_auto_auth", v); err != nil {
			return diag.Errorf("error reading switch_auto_auth: %v", err)
		}
	}

	if o.SwitchFortilink != nil {
		v := *o.SwitchFortilink

		if err = d.Set("switch_fortilink", v); err != nil {
			return diag.Errorf("error reading switch_fortilink: %v", err)
		}
	}

	if o.SwitchGroup != nil {
		if err = d.Set("switch_group", flattenUserNacPolicySwitchGroup(o.SwitchGroup, sort)); err != nil {
			return diag.Errorf("error reading switch_group: %v", err)
		}
	}

	if o.SwitchMacPolicy != nil {
		v := *o.SwitchMacPolicy

		if err = d.Set("switch_mac_policy", v); err != nil {
			return diag.Errorf("error reading switch_mac_policy: %v", err)
		}
	}

	if o.SwitchPortPolicy != nil {
		v := *o.SwitchPortPolicy

		if err = d.Set("switch_port_policy", v); err != nil {
			return diag.Errorf("error reading switch_port_policy: %v", err)
		}
	}

	if o.SwitchScope != nil {
		if err = d.Set("switch_scope", flattenUserNacPolicySwitchScope(o.SwitchScope, sort)); err != nil {
			return diag.Errorf("error reading switch_scope: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.User != nil {
		v := *o.User

		if err = d.Set("user", v); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	if o.UserGroup != nil {
		v := *o.UserGroup

		if err = d.Set("user_group", v); err != nil {
			return diag.Errorf("error reading user_group: %v", err)
		}
	}

	return nil
}

func expandUserNacPolicySwitchGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserNacPolicySwitchGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserNacPolicySwitchGroup

	for i := range l {
		tmp := models.UserNacPolicySwitchGroup{}
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

func expandUserNacPolicySwitchScope(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.UserNacPolicySwitchScope, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.UserNacPolicySwitchScope

	for i := range l {
		tmp := models.UserNacPolicySwitchScope{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.switch_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectUserNacPolicy(d *schema.ResourceData, sv string) (*models.UserNacPolicy, diag.Diagnostics) {
	obj := models.UserNacPolicy{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("category"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("category", sv)
				diags = append(diags, e)
			}
			obj.Category = &v2
		}
	}
	if v1, ok := d.GetOk("description"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("description", sv)
				diags = append(diags, e)
			}
			obj.Description = &v2
		}
	}
	if v1, ok := d.GetOk("ems_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("ems_tag", sv)
				diags = append(diags, e)
			}
			obj.EmsTag = &v2
		}
	}
	if v1, ok := d.GetOk("family"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("family", sv)
				diags = append(diags, e)
			}
			obj.Family = &v2
		}
	}
	if v1, ok := d.GetOk("firewall_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("firewall_address", sv)
				diags = append(diags, e)
			}
			obj.FirewallAddress = &v2
		}
	}
	if v1, ok := d.GetOk("host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host", sv)
				diags = append(diags, e)
			}
			obj.Host = &v2
		}
	}
	if v1, ok := d.GetOk("hw_vendor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hw_vendor", sv)
				diags = append(diags, e)
			}
			obj.HwVendor = &v2
		}
	}
	if v1, ok := d.GetOk("hw_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hw_version", sv)
				diags = append(diags, e)
			}
			obj.HwVersion = &v2
		}
	}
	if v1, ok := d.GetOk("mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac", sv)
				diags = append(diags, e)
			}
			obj.Mac = &v2
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
	if v1, ok := d.GetOk("src"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src", sv)
				diags = append(diags, e)
			}
			obj.Src = &v2
		}
	}
	if v1, ok := d.GetOk("ssid_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ssid_policy", sv)
				diags = append(diags, e)
			}
			obj.SsidPolicy = &v2
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
	if v1, ok := d.GetOk("sw_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sw_version", sv)
				diags = append(diags, e)
			}
			obj.SwVersion = &v2
		}
	}
	if v1, ok := d.GetOk("switch_auto_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("switch_auto_auth", sv)
				diags = append(diags, e)
			}
			obj.SwitchAutoAuth = &v2
		}
	}
	if v1, ok := d.GetOk("switch_fortilink"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_fortilink", sv)
				diags = append(diags, e)
			}
			obj.SwitchFortilink = &v2
		}
	}
	if v, ok := d.GetOk("switch_group"); ok {
		if !utils.CheckVer(sv, "v7.0.2", "") {
			e := utils.AttributeVersionWarning("switch_group", sv)
			diags = append(diags, e)
		}
		t, err := expandUserNacPolicySwitchGroup(d, v, "switch_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SwitchGroup = t
		}
	} else if d.HasChange("switch_group") {
		old, new := d.GetChange("switch_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SwitchGroup = &[]models.UserNacPolicySwitchGroup{}
		}
	}
	if v1, ok := d.GetOk("switch_mac_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switch_mac_policy", sv)
				diags = append(diags, e)
			}
			obj.SwitchMacPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("switch_port_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("switch_port_policy", sv)
				diags = append(diags, e)
			}
			obj.SwitchPortPolicy = &v2
		}
	}
	if v, ok := d.GetOk("switch_scope"); ok {
		if !utils.CheckVer(sv, "", "v7.0.2") {
			e := utils.AttributeVersionWarning("switch_scope", sv)
			diags = append(diags, e)
		}
		t, err := expandUserNacPolicySwitchScope(d, v, "switch_scope", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SwitchScope = t
		}
	} else if d.HasChange("switch_scope") {
		old, new := d.GetChange("switch_scope")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SwitchScope = &[]models.UserNacPolicySwitchScope{}
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user", sv)
				diags = append(diags, e)
			}
			obj.User = &v2
		}
	}
	if v1, ok := d.GetOk("user_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_group", sv)
				diags = append(diags, e)
			}
			obj.UserGroup = &v2
		}
	}
	return &obj, diags
}
