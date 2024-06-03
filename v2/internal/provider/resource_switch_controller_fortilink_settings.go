// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure integrated FortiLink settings for FortiSwitch.

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

func resourceSwitchControllerFortilinkSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure integrated FortiLink settings for FortiSwitch.",

		CreateContext: resourceSwitchControllerFortilinkSettingsCreate,
		ReadContext:   resourceSwitchControllerFortilinkSettingsRead,
		UpdateContext: resourceSwitchControllerFortilinkSettingsUpdate,
		DeleteContext: resourceSwitchControllerFortilinkSettingsDelete,

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
			"fortilink": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "FortiLink interface to which this fortilink-setting belongs.",
				Optional:    true,
				Computed:    true,
			},
			"inactive_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),

				Description: "Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).",
				Optional:    true,
				Computed:    true,
			},
			"link_down_flush": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Clear NAC and dynamic devices on switch ports on link down event.",
				Optional:    true,
				Computed:    true,
			},
			"nac_ports": {
				Type:        schema.TypeList,
				Description: "NAC specific configuration.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bounce_nac_port": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.",
							Optional:    true,
							Computed:    true,
						},
						"lan_segment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),

							Description: "Enable/disable LAN segment feature on the FortiLink interface.",
							Optional:    true,
							Computed:    true,
						},
						"member_change": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Member change flag.",
							Optional:    true,
							Computed:    true,
						},
						"nac_lan_interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Configure NAC LAN interface.",
							Optional:    true,
							Computed:    true,
						},
						"nac_segment_vlans": {
							Type:        schema.TypeList,
							Description: "Configure NAC segment VLANs.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "VLAN interface name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"onboarding_vlan": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Default NAC Onboarding VLAN when NAC devices are discovered.",
							Optional:    true,
							Computed:    true,
						},
						"parent_key": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Parent key name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FortiLink settings name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSwitchControllerFortilinkSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SwitchControllerFortilinkSettings resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerFortilinkSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerFortilinkSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerFortilinkSettings")
	}

	return resourceSwitchControllerFortilinkSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerFortilinkSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerFortilinkSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerFortilinkSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerFortilinkSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerFortilinkSettings")
	}

	return resourceSwitchControllerFortilinkSettingsRead(ctx, d, meta)
}

func resourceSwitchControllerFortilinkSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSwitchControllerFortilinkSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerFortilinkSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerFortilinkSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerFortilinkSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerFortilinkSettings resource: %v", err)
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

	diags := refreshObjectSwitchControllerFortilinkSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, v *models.SwitchControllerFortilinkSettingsNacPorts, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.SwitchControllerFortilinkSettingsNacPorts{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.BounceNacPort; tmp != nil {
				v["bounce_nac_port"] = *tmp
			}

			if tmp := cfg.LanSegment; tmp != nil {
				v["lan_segment"] = *tmp
			}

			if tmp := cfg.MemberChange; tmp != nil {
				v["member_change"] = *tmp
			}

			if tmp := cfg.NacLanInterface; tmp != nil {
				v["nac_lan_interface"] = *tmp
			}

			if tmp := cfg.NacSegmentVlans; tmp != nil {
				v["nac_segment_vlans"] = flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "nac_segment_vlans"), sort)
			}

			if tmp := cfg.OnboardingVlan; tmp != nil {
				v["onboarding_vlan"] = *tmp
			}

			if tmp := cfg.ParentKey; tmp != nil {
				v["parent_key"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d *schema.ResourceData, v *[]models.SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans, prefix string, sort bool) interface{} {
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

func refreshObjectSwitchControllerFortilinkSettings(d *schema.ResourceData, o *models.SwitchControllerFortilinkSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Fortilink != nil {
		v := *o.Fortilink

		if err = d.Set("fortilink", v); err != nil {
			return diag.Errorf("error reading fortilink: %v", err)
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

	if _, ok := d.GetOk("nac_ports"); ok {
		if o.NacPorts != nil {
			if err = d.Set("nac_ports", flattenSwitchControllerFortilinkSettingsNacPorts(d, o.NacPorts, "nac_ports", sort)); err != nil {
				return diag.Errorf("error reading nac_ports: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.SwitchControllerFortilinkSettingsNacPorts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerFortilinkSettingsNacPorts

	for i := range l {
		tmp := models.SwitchControllerFortilinkSettingsNacPorts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.bounce_nac_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BounceNacPort = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lan_segment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LanSegment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.member_change", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MemberChange = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nac_lan_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NacLanInterface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nac_segment_vlans", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans
			// 	}
			tmp.NacSegmentVlans = v2

		}

		pre_append = fmt.Sprintf("%s.%d.onboarding_vlan", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OnboardingVlan = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parent_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ParentKey = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans

	for i := range l {
		tmp := models.SwitchControllerFortilinkSettingsNacPortsNacSegmentVlans{}
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

func getObjectSwitchControllerFortilinkSettings(d *schema.ResourceData, sv string) (*models.SwitchControllerFortilinkSettings, diag.Diagnostics) {
	obj := models.SwitchControllerFortilinkSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fortilink"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("fortilink", sv)
				diags = append(diags, e)
			}
			obj.Fortilink = &v2
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
	if v, ok := d.GetOk("nac_ports"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nac_ports", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerFortilinkSettingsNacPorts(d, v, "nac_ports", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NacPorts = t
		}
	} else if d.HasChange("nac_ports") {
		old, new := d.GetChange("nac_ports")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NacPorts = &models.SwitchControllerFortilinkSettingsNacPorts{}
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
	return &obj, diags
}
