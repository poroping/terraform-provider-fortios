// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSwitchControllerMacPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.",

		CreateContext: resourceSwitchControllerMacPolicyCreate,
		ReadContext:   resourceSwitchControllerMacPolicyRead,
		UpdateContext: resourceSwitchControllerMacPolicyUpdate,
		DeleteContext: resourceSwitchControllerMacPolicyDelete,

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
			"bounce_port_link": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where this mac-policy is applied.",
				Optional:    true,
				Computed:    true,
			},
			"foscount": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable packet count on the NAC device.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Description for the MAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"fortilink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "FortiLink interface for which this MAC policy belongs to.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "MAC policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"traffic_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Traffic policy to be applied when using this MAC policy.",
				Optional:    true,
				Computed:    true,
			},
			"vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Ingress traffic VLAN assignment for the MAC address matching this MAC policy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerMacPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerMacPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerMacPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerMacPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerMacPolicy")
	}

	return resourceSwitchControllerMacPolicyRead(ctx, d, meta)
}

func resourceSwitchControllerMacPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerMacPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerMacPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerMacPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerMacPolicy")
	}

	return resourceSwitchControllerMacPolicyRead(ctx, d, meta)
}

func resourceSwitchControllerMacPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerMacPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerMacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerMacPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerMacPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerMacPolicy resource: %v", err)
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

	diags := refreshObjectSwitchControllerMacPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerMacPolicy(d *schema.ResourceData, o *models.SwitchControllerMacPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BouncePortLink != nil {
		v := *o.BouncePortLink

		if err = d.Set("bounce_port_link", v); err != nil {
			return diag.Errorf("error reading bounce_port_link: %v", err)
		}
	}

	if o.Count != nil {
		v := *o.Count

		if err = d.Set("foscount", v); err != nil {
			return diag.Errorf("error reading foscount: %v", err)
		}
	}

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

	if o.TrafficPolicy != nil {
		v := *o.TrafficPolicy

		if err = d.Set("traffic_policy", v); err != nil {
			return diag.Errorf("error reading traffic_policy: %v", err)
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

func getObjectSwitchControllerMacPolicy(d *schema.ResourceData, sv string) (*models.SwitchControllerMacPolicy, diag.Diagnostics) {
	obj := models.SwitchControllerMacPolicy{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("bounce_port_link"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("bounce_port_link", sv)
				diags = append(diags, e)
			}
			obj.BouncePortLink = &v2
		}
	}
	if v1, ok := d.GetOk("foscount"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("foscount", sv)
				diags = append(diags, e)
			}
			obj.Count = &v2
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
	if v1, ok := d.GetOk("traffic_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_policy", sv)
				diags = append(diags, e)
			}
			obj.TrafficPolicy = &v2
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
