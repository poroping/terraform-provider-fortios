// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure integrated NAC settings for FortiSwitch.

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

func resourceSwitchControllerNacSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure integrated NAC settings for FortiSwitch.",

		CreateContext: resourceSwitchControllerNacSettingsCreate,
		ReadContext:   resourceSwitchControllerNacSettingsRead,
		UpdateContext: resourceSwitchControllerNacSettingsUpdate,
		DeleteContext: resourceSwitchControllerNacSettingsDelete,

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
			"auto_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable NAC device auto authorization when discovered and nac-policy matched.",
				Optional:    true,
				Computed:    true,
			},
			"bounce_nac_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.",
				Optional:    true,
				Computed:    true,
			},
			"inactive_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1440),

				Description: "Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).",
				Optional:    true,
				Computed:    true,
			},
			"link_down_flush": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Clear NAC devices on switch ports on link down event.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"local", "global"}, false),

				Description: "Set NAC mode to be used on the FortiSwitch ports.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "NAC settings name.",
				ForceNew:    true,
				Required:    true,
			},
			"onboarding_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Default NAC Onboarding VLAN when NAC devices are discovered.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerNacSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerNacSettings resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerNacSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerNacSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerNacSettings")
	}

	return resourceSwitchControllerNacSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerNacSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerNacSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerNacSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerNacSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerNacSettings")
	}

	return resourceSwitchControllerNacSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerNacSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerNacSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerNacSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerNacSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerNacSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerNacSettings resource: %v", err)
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

	diags := refreshObjectSwitchControllerNacSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSwitchControllerNacSettings(d *schema.ResourceData, o *models.SwitchControllerNacSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoAuth != nil {
		v := *o.AutoAuth

		if err = d.Set("auto_auth", v); err != nil {
			return diag.Errorf("error reading auto_auth: %v", err)
		}
	}

	if o.BounceNacPort != nil {
		v := *o.BounceNacPort

		if err = d.Set("bounce_nac_port", v); err != nil {
			return diag.Errorf("error reading bounce_nac_port: %v", err)
		}
	}

	if o.InactiveTimer != nil {
		v := *o.InactiveTimer

		if err = d.Set("inactive_timer", v); err != nil {
			return diag.Errorf("error reading inactive_timer: %v", err)
		}
	}

	if o.LinkDownFlush != nil {
		v := *o.LinkDownFlush

		if err = d.Set("link_down_flush", v); err != nil {
			return diag.Errorf("error reading link_down_flush: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.OnboardingVlan != nil {
		v := *o.OnboardingVlan

		if err = d.Set("onboarding_vlan", v); err != nil {
			return diag.Errorf("error reading onboarding_vlan: %v", err)
		}
	}

	return nil
}

func getObjectSwitchControllerNacSettings(d *schema.ResourceData, sv string) (*models.SwitchControllerNacSettings, diag.Diagnostics) {
	obj := models.SwitchControllerNacSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_auth", sv)
				diags = append(diags, e)
			}
			obj.AutoAuth = &v2
		}
	}
	if v1, ok := d.GetOk("bounce_nac_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bounce_nac_port", sv)
				diags = append(diags, e)
			}
			obj.BounceNacPort = &v2
		}
	}
	if v1, ok := d.GetOk("inactive_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inactive_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.InactiveTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("link_down_flush"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_down_flush", sv)
				diags = append(diags, e)
			}
			obj.LinkDownFlush = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
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
	if v1, ok := d.GetOk("onboarding_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("onboarding_vlan", sv)
				diags = append(diags, e)
			}
			obj.OnboardingVlan = &v2
		}
	}
	return &obj, diags
}
