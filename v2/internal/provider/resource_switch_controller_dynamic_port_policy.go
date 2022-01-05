// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

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

func resourceSwitchControllerDynamicPortPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.",

		CreateContext: resourceSwitchControllerDynamicPortPolicyCreate,
		ReadContext:   resourceSwitchControllerDynamicPortPolicyRead,
		UpdateContext: resourceSwitchControllerDynamicPortPolicyUpdate,
		DeleteContext: resourceSwitchControllerDynamicPortPolicyDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Description for the Dynamic port policy.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "FortiLink interface for which this Dynamic port policy belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Dynamic port policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"policy": {
				Type:        schema.TypeList,
				Description: "Port policies with matching criteria and actions.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"802_1x": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "802.1x security policy to be applied when using this policy.",
							Optional:    true,
							Computed:    true,
						},
						"bounce_port_link": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile.",
							Optional:    true,
							Computed:    true,
						},
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"device", "interface-tag"}, false),

							Description: "Category of Dynamic port policy.",
							Optional:    true,
							Computed:    true,
						},
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description for the policy.",
							Optional:    true,
							Computed:    true,
						},
						"family": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Policy matching family.",
							Optional:    true,
							Computed:    true,
						},
						"host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),

							Description: "Policy matching host.",
							Optional:    true,
							Computed:    true,
						},
						"interface_tags": {
							Type:        schema.TypeList,
							Description: "Policy matching the FortiSwitch interface object tags.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "FortiSwitch port tag name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"lldp_profile": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "LLDP profile to be applied when using this policy.",
							Optional:    true,
							Computed:    true,
						},
						"mac": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 17),

							Description: "Policy matching MAC address.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Policy name.",
							Optional:    true,
							Computed:    true,
						},
						"qos_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "QoS policy to be applied when using this policy.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable policy.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Policy matching type.",
							Optional:    true,
							Computed:    true,
						},
						"vlan_policy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "VLAN policy to be applied when using this policy.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerDynamicPortPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerDynamicPortPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerDynamicPortPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerDynamicPortPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerDynamicPortPolicy")
	}

	return resourceSwitchControllerDynamicPortPolicyRead(ctx, d, meta)
}

func resourceSwitchControllerDynamicPortPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerDynamicPortPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerDynamicPortPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerDynamicPortPolicy")
	}

	return resourceSwitchControllerDynamicPortPolicyRead(ctx, d, meta)
}

func resourceSwitchControllerDynamicPortPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerDynamicPortPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerDynamicPortPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerDynamicPortPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerDynamicPortPolicy resource: %v", err)
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

	diags := refreshObjectSwitchControllerDynamicPortPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerDynamicPortPolicyPolicy(v *[]models.SwitchControllerDynamicPortPolicyPolicy, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.The8021x; tmp != nil {
				v["802_1x"] = *tmp
			}

			if tmp := cfg.BouncePortLink; tmp != nil {
				v["bounce_port_link"] = *tmp
			}

			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Family; tmp != nil {
				v["family"] = *tmp
			}

			if tmp := cfg.Host; tmp != nil {
				v["host"] = *tmp
			}

			if tmp := cfg.InterfaceTags; tmp != nil {
				v["interface_tags"] = flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags(tmp, sort)
			}

			if tmp := cfg.LldpProfile; tmp != nil {
				v["lldp_profile"] = *tmp
			}

			if tmp := cfg.Mac; tmp != nil {
				v["mac"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.QosPolicy; tmp != nil {
				v["qos_policy"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.VlanPolicy; tmp != nil {
				v["vlan_policy"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags(v *[]models.SwitchControllerDynamicPortPolicyPolicyInterfaceTags, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.TagName; tmp != nil {
				v["tag_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "tag_name")
	}

	return flat
}

func refreshObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData, o *models.SwitchControllerDynamicPortPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
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

	if o.Policy != nil {
		if err = d.Set("policy", flattenSwitchControllerDynamicPortPolicyPolicy(o.Policy, sort)); err != nil {
			return diag.Errorf("error reading policy: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerDynamicPortPolicyPolicy, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerDynamicPortPolicyPolicy

	for i := range l {
		tmp := models.SwitchControllerDynamicPortPolicyPolicy{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.802_1x", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.The8021x = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bounce_port_link", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BouncePortLink = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.family", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Family = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Host = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface_tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerDynamicPortPolicyPolicyInterfaceTags
			// 	}
			tmp.InterfaceTags = v2

		}

		pre_append = fmt.Sprintf("%s.%d.lldp_profile", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LldpProfile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mac", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mac = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.qos_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QosPolicy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.vlan_policy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VlanPolicy = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerDynamicPortPolicyPolicyInterfaceTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerDynamicPortPolicyPolicyInterfaceTags

	for i := range l {
		tmp := models.SwitchControllerDynamicPortPolicyPolicyInterfaceTags{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.tag_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TagName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData, sv string) (*models.SwitchControllerDynamicPortPolicy, diag.Diagnostics) {
	obj := models.SwitchControllerDynamicPortPolicy{}
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
	if v, ok := d.GetOk("policy"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("policy", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerDynamicPortPolicyPolicy(d, v, "policy", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Policy = t
		}
	} else if d.HasChange("policy") {
		old, new := d.GetChange("policy")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Policy = &[]models.SwitchControllerDynamicPortPolicyPolicy{}
		}
	}
	return &obj, diags
}
