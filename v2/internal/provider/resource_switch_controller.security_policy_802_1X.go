// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure 802.1x MAC Authentication Bypass (MAB) policies.

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

func resourceSwitchControllerSecurityPolicy8021X() *schema.Resource {
	return &schema.Resource{
		Description: "Configure 802.1x MAC Authentication Bypass (MAB) policies.",

		CreateContext: resourceSwitchControllerSecurityPolicy8021XCreate,
		ReadContext:   resourceSwitchControllerSecurityPolicy8021XRead,
		UpdateContext: resourceSwitchControllerSecurityPolicy8021XUpdate,
		DeleteContext: resourceSwitchControllerSecurityPolicy8021XDelete,

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
			"auth_fail_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to allow limited access to clients that cannot authenticate.",
				Optional:    true,
				Computed:    true,
			},
			"auth_fail_vlan_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "VLAN ID on which authentication failed.",
				Optional:    true,
				Computed:    true,
			},
			"authserver_timeout_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 15),

				Description: "Authentication server timeout period (3 - 15 sec, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"authserver_timeout_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable the authentication server timeout VLAN to allow limited access when RADIUS is unavailable. ",
				Optional:    true,
				Computed:    true,
			},
			"authserver_timeout_vlanid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Authentication server timeout VLAN name.",
				Optional:    true,
				Computed:    true,
			},
			"eap_auto_untagged_vlans": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable automatic inclusion of untagged VLANs.",
				Optional:    true,
				Computed:    true,
			},
			"eap_passthru": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable EAP pass-through mode, allowing protocols (such as LLDP) to pass through ports for more flexible authentication.",
				Optional:    true,
				Computed:    true,
			},
			"framevid_apply": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable the capability to apply the EAP/MAB frame VLAN to the port native VLAN.",
				Optional:    true,
				Computed:    true,
			},
			"guest_auth_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 900),

				Description: "Guest authentication delay (1 - 900  sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"guest_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable the guest VLAN feature to allow limited access to non-802.1X-compliant clients.",
				Optional:    true,
				Computed:    true,
			},
			"guest_vlan_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Guest VLAN name.",
				Optional:    true,
				Computed:    true,
			},
			"mac_auth_bypass": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable MAB for this policy.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"open_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable open authentication for this policy.",
				Optional:    true,
				Computed:    true,
			},
			"policy_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"802.1X"}, false),

				Description: "Policy type.",
				Optional:    true,
				Computed:    true,
			},
			"radius_timeout_overwrite": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to override the global RADIUS session timeout.",
				Optional:    true,
				Computed:    true,
			},
			"security_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"802.1X", "802.1X-mac-based"}, false),

				Description: "Port or MAC based 802.1X security mode.",
				Optional:    true,
				Computed:    true,
			},
			"user_group": {
				Type:        schema.TypeList,
				Description: "Name of user-group to assign to this MAC Authentication Bypass (MAB) policy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicy8021XCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerSecurityPolicy8021X resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerSecurityPolicy8021X(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerSecurityPolicy8021X(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerSecurityPolicy8021X")
	}

	return resourceSwitchControllerSecurityPolicy8021XRead(ctx, d, meta)
}

func resourceSwitchControllerSecurityPolicy8021XUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerSecurityPolicy8021X(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerSecurityPolicy8021X(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerSecurityPolicy8021X")
	}

	return resourceSwitchControllerSecurityPolicy8021XRead(ctx, d, meta)
}

func resourceSwitchControllerSecurityPolicy8021XDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerSecurityPolicy8021X(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicy8021XRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerSecurityPolicy8021X(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerSecurityPolicy8021X resource: %v", err)
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

	diags := refreshObjectSwitchControllerSecurityPolicy8021X(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerSecurityPolicy8021XUserGroup(v *[]models.SwitchControllerSecurityPolicy8021XUserGroup, sort bool) interface{} {
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

func refreshObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData, o *models.SwitchControllerSecurityPolicy8021X, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthFailVlan != nil {
		v := *o.AuthFailVlan

		if err = d.Set("auth_fail_vlan", v); err != nil {
			return diag.Errorf("error reading auth_fail_vlan: %v", err)
		}
	}

	if o.AuthFailVlanId != nil {
		v := *o.AuthFailVlanId

		if err = d.Set("auth_fail_vlan_id", v); err != nil {
			return diag.Errorf("error reading auth_fail_vlan_id: %v", err)
		}
	}

	if o.AuthserverTimeoutPeriod != nil {
		v := *o.AuthserverTimeoutPeriod

		if err = d.Set("authserver_timeout_period", v); err != nil {
			return diag.Errorf("error reading authserver_timeout_period: %v", err)
		}
	}

	if o.AuthserverTimeoutVlan != nil {
		v := *o.AuthserverTimeoutVlan

		if err = d.Set("authserver_timeout_vlan", v); err != nil {
			return diag.Errorf("error reading authserver_timeout_vlan: %v", err)
		}
	}

	if o.AuthserverTimeoutVlanid != nil {
		v := *o.AuthserverTimeoutVlanid

		if err = d.Set("authserver_timeout_vlanid", v); err != nil {
			return diag.Errorf("error reading authserver_timeout_vlanid: %v", err)
		}
	}

	if o.EapAutoUntaggedVlans != nil {
		v := *o.EapAutoUntaggedVlans

		if err = d.Set("eap_auto_untagged_vlans", v); err != nil {
			return diag.Errorf("error reading eap_auto_untagged_vlans: %v", err)
		}
	}

	if o.EapPassthru != nil {
		v := *o.EapPassthru

		if err = d.Set("eap_passthru", v); err != nil {
			return diag.Errorf("error reading eap_passthru: %v", err)
		}
	}

	if o.FramevidApply != nil {
		v := *o.FramevidApply

		if err = d.Set("framevid_apply", v); err != nil {
			return diag.Errorf("error reading framevid_apply: %v", err)
		}
	}

	if o.GuestAuthDelay != nil {
		v := *o.GuestAuthDelay

		if err = d.Set("guest_auth_delay", v); err != nil {
			return diag.Errorf("error reading guest_auth_delay: %v", err)
		}
	}

	if o.GuestVlan != nil {
		v := *o.GuestVlan

		if err = d.Set("guest_vlan", v); err != nil {
			return diag.Errorf("error reading guest_vlan: %v", err)
		}
	}

	if o.GuestVlanId != nil {
		v := *o.GuestVlanId

		if err = d.Set("guest_vlan_id", v); err != nil {
			return diag.Errorf("error reading guest_vlan_id: %v", err)
		}
	}

	if o.MacAuthBypass != nil {
		v := *o.MacAuthBypass

		if err = d.Set("mac_auth_bypass", v); err != nil {
			return diag.Errorf("error reading mac_auth_bypass: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OpenAuth != nil {
		v := *o.OpenAuth

		if err = d.Set("open_auth", v); err != nil {
			return diag.Errorf("error reading open_auth: %v", err)
		}
	}

	if o.PolicyType != nil {
		v := *o.PolicyType

		if err = d.Set("policy_type", v); err != nil {
			return diag.Errorf("error reading policy_type: %v", err)
		}
	}

	if o.RadiusTimeoutOverwrite != nil {
		v := *o.RadiusTimeoutOverwrite

		if err = d.Set("radius_timeout_overwrite", v); err != nil {
			return diag.Errorf("error reading radius_timeout_overwrite: %v", err)
		}
	}

	if o.SecurityMode != nil {
		v := *o.SecurityMode

		if err = d.Set("security_mode", v); err != nil {
			return diag.Errorf("error reading security_mode: %v", err)
		}
	}

	if o.UserGroup != nil {
		if err = d.Set("user_group", flattenSwitchControllerSecurityPolicy8021XUserGroup(o.UserGroup, sort)); err != nil {
			return diag.Errorf("error reading user_group: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerSecurityPolicy8021XUserGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerSecurityPolicy8021XUserGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerSecurityPolicy8021XUserGroup

	for i := range l {
		tmp := models.SwitchControllerSecurityPolicy8021XUserGroup{}
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

func getObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData, sv string) (*models.SwitchControllerSecurityPolicy8021X, diag.Diagnostics) {
	obj := models.SwitchControllerSecurityPolicy8021X{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_fail_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_fail_vlan", sv)
				diags = append(diags, e)
			}
			obj.AuthFailVlan = &v2
		}
	}
	if v1, ok := d.GetOk("auth_fail_vlan_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_fail_vlan_id", sv)
				diags = append(diags, e)
			}
			obj.AuthFailVlanId = &v2
		}
	}
	if v1, ok := d.GetOk("authserver_timeout_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("authserver_timeout_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthserverTimeoutPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("authserver_timeout_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("authserver_timeout_vlan", sv)
				diags = append(diags, e)
			}
			obj.AuthserverTimeoutVlan = &v2
		}
	}
	if v1, ok := d.GetOk("authserver_timeout_vlanid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("authserver_timeout_vlanid", sv)
				diags = append(diags, e)
			}
			obj.AuthserverTimeoutVlanid = &v2
		}
	}
	if v1, ok := d.GetOk("eap_auto_untagged_vlans"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("eap_auto_untagged_vlans", sv)
				diags = append(diags, e)
			}
			obj.EapAutoUntaggedVlans = &v2
		}
	}
	if v1, ok := d.GetOk("eap_passthru"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("eap_passthru", sv)
				diags = append(diags, e)
			}
			obj.EapPassthru = &v2
		}
	}
	if v1, ok := d.GetOk("framevid_apply"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("framevid_apply", sv)
				diags = append(diags, e)
			}
			obj.FramevidApply = &v2
		}
	}
	if v1, ok := d.GetOk("guest_auth_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guest_auth_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.GuestAuthDelay = &tmp
		}
	}
	if v1, ok := d.GetOk("guest_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guest_vlan", sv)
				diags = append(diags, e)
			}
			obj.GuestVlan = &v2
		}
	}
	if v1, ok := d.GetOk("guest_vlan_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("guest_vlan_id", sv)
				diags = append(diags, e)
			}
			obj.GuestVlanId = &v2
		}
	}
	if v1, ok := d.GetOk("mac_auth_bypass"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_auth_bypass", sv)
				diags = append(diags, e)
			}
			obj.MacAuthBypass = &v2
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
	if v1, ok := d.GetOk("open_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("open_auth", sv)
				diags = append(diags, e)
			}
			obj.OpenAuth = &v2
		}
	}
	if v1, ok := d.GetOk("policy_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policy_type", sv)
				diags = append(diags, e)
			}
			obj.PolicyType = &v2
		}
	}
	if v1, ok := d.GetOk("radius_timeout_overwrite"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_timeout_overwrite", sv)
				diags = append(diags, e)
			}
			obj.RadiusTimeoutOverwrite = &v2
		}
	}
	if v1, ok := d.GetOk("security_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("security_mode", sv)
				diags = append(diags, e)
			}
			obj.SecurityMode = &v2
		}
	}
	if v, ok := d.GetOk("user_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("user_group", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerSecurityPolicy8021XUserGroup(d, v, "user_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UserGroup = t
		}
	} else if d.HasChange("user_group") {
		old, new := d.GetChange("user_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UserGroup = &[]models.SwitchControllerSecurityPolicy8021XUserGroup{}
		}
	}
	return &obj, diags
}
