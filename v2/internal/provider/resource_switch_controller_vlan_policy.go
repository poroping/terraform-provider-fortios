// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

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

func resourceSwitchControllerVlanPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.",

		CreateContext: resourceSwitchControllerVlanPolicyCreate,
		ReadContext:   resourceSwitchControllerVlanPolicyRead,
		UpdateContext: resourceSwitchControllerVlanPolicyUpdate,
		DeleteContext: resourceSwitchControllerVlanPolicyDelete,

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
			"allowed_vlans": {
				Type:        schema.TypeList,
				Description: "Allowed VLANs to be applied when using this VLAN policy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "VLAN name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"allowed_vlans_all": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable all defined VLANs when using this VLAN policy.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Description for the VLAN policy.",
				Optional:    true,
				Computed:    true,
			},
			"discard_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "all-untagged", "all-tagged"}, false),

				Description: "Discard mode to be applied when using this VLAN policy.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "FortiLink interface for which this VLAN policy belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "VLAN policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"untagged_vlans": {
				Type:        schema.TypeList,
				Description: "Untagged VLANs to be applied when using this VLAN policy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "VLAN name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Native VLAN to be applied when using this VLAN policy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerVlanPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerVlanPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerVlanPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerVlanPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerVlanPolicy")
	}

	return resourceSwitchControllerVlanPolicyRead(ctx, d, meta)
}

func resourceSwitchControllerVlanPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerVlanPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerVlanPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerVlanPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerVlanPolicy")
	}

	return resourceSwitchControllerVlanPolicyRead(ctx, d, meta)
}

func resourceSwitchControllerVlanPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerVlanPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerVlanPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerVlanPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerVlanPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerVlanPolicy resource: %v", err)
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

	diags := refreshObjectSwitchControllerVlanPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerVlanPolicyAllowedVlans(v *[]models.SwitchControllerVlanPolicyAllowedVlans, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.VlanName; tmp != nil {
				v["vlan_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vlan_name")
	}

	return flat
}

func flattenSwitchControllerVlanPolicyUntaggedVlans(v *[]models.SwitchControllerVlanPolicyUntaggedVlans, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.VlanName; tmp != nil {
				v["vlan_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "vlan_name")
	}

	return flat
}

func refreshObjectSwitchControllerVlanPolicy(d *schema.ResourceData, o *models.SwitchControllerVlanPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowedVlans != nil {
		if err = d.Set("allowed_vlans", flattenSwitchControllerVlanPolicyAllowedVlans(o.AllowedVlans, sort)); err != nil {
			return diag.Errorf("error reading allowed_vlans: %v", err)
		}
	}

	if o.AllowedVlansAll != nil {
		v := *o.AllowedVlansAll

		if err = d.Set("allowed_vlans_all", v); err != nil {
			return diag.Errorf("error reading allowed_vlans_all: %v", err)
		}
	}

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.DiscardMode != nil {
		v := *o.DiscardMode

		if err = d.Set("discard_mode", v); err != nil {
			return diag.Errorf("error reading discard_mode: %v", err)
		}
	}

	if o.Fortilink != nil {
		v := *o.Fortilink

		if err = d.Set("fortilink", v); err != nil {
			return diag.Errorf("error reading fortilink: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.UntaggedVlans != nil {
		if err = d.Set("untagged_vlans", flattenSwitchControllerVlanPolicyUntaggedVlans(o.UntaggedVlans, sort)); err != nil {
			return diag.Errorf("error reading untagged_vlans: %v", err)
		}
	}

	if o.Vlan != nil {
		v := *o.Vlan

		if err = d.Set("vlan", v); err != nil {
			return diag.Errorf("error reading vlan: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerVlanPolicyAllowedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerVlanPolicyAllowedVlans, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerVlanPolicyAllowedVlans

	for i := range l {
		tmp := models.SwitchControllerVlanPolicyAllowedVlans{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vlan_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerVlanPolicyUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerVlanPolicyUntaggedVlans, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerVlanPolicyUntaggedVlans

	for i := range l {
		tmp := models.SwitchControllerVlanPolicyUntaggedVlans{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.vlan_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSwitchControllerVlanPolicy(d *schema.ResourceData, sv string) (*models.SwitchControllerVlanPolicy, diag.Diagnostics) {
	obj := models.SwitchControllerVlanPolicy{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("allowed_vlans"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("allowed_vlans", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerVlanPolicyAllowedVlans(d, v, "allowed_vlans", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AllowedVlans = t
		}
	} else if d.HasChange("allowed_vlans") {
		old, new := d.GetChange("allowed_vlans")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AllowedVlans = &[]models.SwitchControllerVlanPolicyAllowedVlans{}
		}
	}
	if v1, ok := d.GetOk("allowed_vlans_all"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allowed_vlans_all", sv)
				diags = append(diags, e)
			}
			obj.AllowedVlansAll = &v2
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
	if v1, ok := d.GetOk("discard_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("discard_mode", sv)
				diags = append(diags, e)
			}
			obj.DiscardMode = &v2
		}
	}
	if v1, ok := d.GetOk("fortilink"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fortilink", sv)
				diags = append(diags, e)
			}
			obj.Fortilink = &v2
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
	if v, ok := d.GetOk("untagged_vlans"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("untagged_vlans", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerVlanPolicyUntaggedVlans(d, v, "untagged_vlans", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UntaggedVlans = t
		}
	} else if d.HasChange("untagged_vlans") {
		old, new := d.GetChange("untagged_vlans")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UntaggedVlans = &[]models.SwitchControllerVlanPolicyUntaggedVlans{}
		}
	}
	if v1, ok := d.GetOk("vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan", sv)
				diags = append(diags, e)
			}
			obj.Vlan = &v2
		}
	}
	return &obj, diags
}
