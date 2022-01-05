// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.

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

func resourceSwitchControllerNacDevice() *schema.Resource {
	return &schema.Resource{
		Description: "Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy.",

		CreateContext: resourceSwitchControllerNacDeviceCreate,
		ReadContext:   resourceSwitchControllerNacDeviceRead,
		UpdateContext: resourceSwitchControllerNacDeviceUpdate,
		DeleteContext: resourceSwitchControllerNacDeviceDelete,

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
			"description": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Description for the learned NAC device.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Device ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"last_known_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Managed FortiSwitch port where NAC device is last learned.",
				Optional:    true,
				Computed:    true,
			},
			"last_known_switch": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Managed FortiSwitch where NAC device is last learned.",
				Optional:    true,
				Computed:    true,
			},
			"last_seen": {
				Type: schema.TypeInt,

				Description: "Device last seen.",
				Optional:    true,
				Computed:    true,
			},
			"mac": {
				Type: schema.TypeString,

				Description: "MAC address of the learned NAC device.",
				Optional:    true,
				Computed:    true,
			},
			"mac_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "MAC policy to be applied on this learned NAC device.",
				Optional:    true,
				Computed:    true,
			},
			"matched_nac_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Matched NAC policy for the learned NAC device.",
				Optional:    true,
				Computed:    true,
			},
			"port_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Port policy to be applied on this learned NAC device.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Status of the learned NAC device. Set enable to authorize the NAC device.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerNacDeviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerNacDevice resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerNacDevice(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerNacDevice(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerNacDevice")
	}

	return resourceSwitchControllerNacDeviceRead(ctx, d, meta)
}

func resourceSwitchControllerNacDeviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerNacDevice(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerNacDevice(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerNacDevice resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerNacDevice")
	}

	return resourceSwitchControllerNacDeviceRead(ctx, d, meta)
}

func resourceSwitchControllerNacDeviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerNacDevice(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerNacDevice resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNacDeviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerNacDevice(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerNacDevice resource: %v", err)
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

	diags := refreshObjectSwitchControllerNacDevice(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerNacDevice(d *schema.ResourceData, o *models.SwitchControllerNacDevice, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Description != nil {
		v := *o.Description

		if err = d.Set("description", v); err != nil {
			return diag.Errorf("error reading description: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.LastKnownPort != nil {
		v := *o.LastKnownPort

		if err = d.Set("last_known_port", v); err != nil {
			return diag.Errorf("error reading last_known_port: %v", err)
		}
	}

	if o.LastKnownSwitch != nil {
		v := *o.LastKnownSwitch

		if err = d.Set("last_known_switch", v); err != nil {
			return diag.Errorf("error reading last_known_switch: %v", err)
		}
	}

	if o.LastSeen != nil {
		v := *o.LastSeen

		if err = d.Set("last_seen", v); err != nil {
			return diag.Errorf("error reading last_seen: %v", err)
		}
	}

	if o.Mac != nil {
		v := *o.Mac

		if err = d.Set("mac", v); err != nil {
			return diag.Errorf("error reading mac: %v", err)
		}
	}

	if o.MacPolicy != nil {
		v := *o.MacPolicy

		if err = d.Set("mac_policy", v); err != nil {
			return diag.Errorf("error reading mac_policy: %v", err)
		}
	}

	if o.MatchedNacPolicy != nil {
		v := *o.MatchedNacPolicy

		if err = d.Set("matched_nac_policy", v); err != nil {
			return diag.Errorf("error reading matched_nac_policy: %v", err)
		}
	}

	if o.PortPolicy != nil {
		v := *o.PortPolicy

		if err = d.Set("port_policy", v); err != nil {
			return diag.Errorf("error reading port_policy: %v", err)
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

func getObjectSwitchControllerNacDevice(d *schema.ResourceData, sv string) (*models.SwitchControllerNacDevice, diag.Diagnostics) {
	obj := models.SwitchControllerNacDevice{}
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
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("last_known_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("last_known_port", sv)
				diags = append(diags, e)
			}
			obj.LastKnownPort = &v2
		}
	}
	if v1, ok := d.GetOk("last_known_switch"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("last_known_switch", sv)
				diags = append(diags, e)
			}
			obj.LastKnownSwitch = &v2
		}
	}
	if v1, ok := d.GetOk("last_seen"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("last_seen", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LastSeen = &tmp
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
	if v1, ok := d.GetOk("mac_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_policy", sv)
				diags = append(diags, e)
			}
			obj.MacPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("matched_nac_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("matched_nac_policy", sv)
				diags = append(diags, e)
			}
			obj.MatchedNacPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("port_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port_policy", sv)
				diags = append(diags, e)
			}
			obj.PortPolicy = &v2
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
