// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.

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

func resourceSwitchControllerStpInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.",

		CreateContext: resourceSwitchControllerStpInstanceCreate,
		ReadContext:   resourceSwitchControllerStpInstanceRead,
		UpdateContext: resourceSwitchControllerStpInstanceUpdate,
		DeleteContext: resourceSwitchControllerStpInstanceDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"fosid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),

				Description: "Instance ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"vlan_range": {
				Type:        schema.TypeList,
				Description: "Configure VLAN range for STP instance.",
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
		},
	}
}

func resourceSwitchControllerStpInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SwitchControllerStpInstance resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerStpInstance(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerStpInstance(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerStpInstance")
	}

	return resourceSwitchControllerStpInstanceRead(ctx, d, meta)
}

func resourceSwitchControllerStpInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerStpInstance(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerStpInstance(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerStpInstance resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerStpInstance")
	}

	return resourceSwitchControllerStpInstanceRead(ctx, d, meta)
}

func resourceSwitchControllerStpInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerStpInstance(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerStpInstance resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStpInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerStpInstance(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerStpInstance resource: %v", err)
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

	diags := refreshObjectSwitchControllerStpInstance(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerStpInstanceVlanRange(d *schema.ResourceData, v *[]models.SwitchControllerStpInstanceVlanRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func refreshObjectSwitchControllerStpInstance(d *schema.ResourceData, o *models.SwitchControllerStpInstance, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.VlanRange != nil {
		if err = d.Set("vlan_range", flattenSwitchControllerStpInstanceVlanRange(d, o.VlanRange, "vlan_range", sort)); err != nil {
			return diag.Errorf("error reading vlan_range: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerStpInstanceVlanRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerStpInstanceVlanRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerStpInstanceVlanRange

	for i := range l {
		tmp := models.SwitchControllerStpInstanceVlanRange{}
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

func getObjectSwitchControllerStpInstance(d *schema.ResourceData, sv string) (*models.SwitchControllerStpInstance, diag.Diagnostics) {
	obj := models.SwitchControllerStpInstance{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			obj.Id = &v2
		}
	}
	if v, ok := d.GetOk("vlan_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("vlan_range", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerStpInstanceVlanRange(d, v, "vlan_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.VlanRange = t
		}
	} else if d.HasChange("vlan_range") {
		old, new := d.GetChange("vlan_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.VlanRange = &[]models.SwitchControllerStpInstanceVlanRange{}
		}
	}
	return &obj, diags
}
