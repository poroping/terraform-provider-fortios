// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch global settings.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSwitchControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch global settings.",

		CreateContext: resourceSwitchControllerGlobalCreate,
		ReadContext:   resourceSwitchControllerGlobalRead,
		UpdateContext: resourceSwitchControllerGlobalUpdate,
		DeleteContext: resourceSwitchControllerGlobalDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"allow_multiple_interfaces": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"bounce_quarantined_link": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device.",
				Optional:    true,
				Computed:    true,
			},
			"custom_command": {
				Type:        schema.TypeList,
				Description: "List of custom commands to be pushed to all FortiSwitches in the VDOM.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "List of FortiSwitch commands.",
							Optional:    true,
							Computed:    true,
						},
						"command_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name of custom command to push to all FortiSwitches in VDOM.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"default_virtual_switch_vlan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Default VLAN for ports when added to the virtual-switch.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_server_access_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DHCP snooping server access list.",
				Optional:    true,
				Computed:    true,
			},
			"disable_discovery": {
				Type:        schema.TypeList,
				Description: "Prevent this FortiSwitch from discovering.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Managed device ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fips_enforce": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable enforcement of FIPS on managed FortiSwitch devices.",
				Optional:    true,
				Computed:    true,
			},
			"https_image_push": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable image push to FortiSwitch using HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"log_mac_limit_violations": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logs for Learning Limit Violations.",
				Optional:    true,
				Computed:    true,
			},
			"mac_aging_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000000),

				Description: "Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"mac_event_logging": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable MAC address event logging.",
				Optional:    true,
				Computed:    true,
			},
			"mac_retention_period": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 168),

				Description: "Time in hours after which an inactive MAC is removed from client DB (0 = aged out based on mac-aging-interval).",
				Optional:    true,
				Computed:    true,
			},
			"mac_violation_timer": {
				Type: schema.TypeInt,

				Description: "Set timeout for Learning Limit Violations (0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"quarantine_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"by-vlan", "by-redirect"}, false),

				Description: "Quarantine mode.",
				Optional:    true,
				Computed:    true,
			},
			"sn_dns_resolution": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number.",
				Optional:    true,
				Computed:    true,
			},
			"update_user_device": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Control which sources update the device user list.",
				Optional:         true,
				Computed:         true,
			},
			"vlan_all_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "defined"}, false),

				Description: "VLAN configuration mode, user-defined-vlans or all-possible-vlans.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_optimization": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "FortiLink VLAN optimization.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSwitchControllerGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerGlobal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerGlobal")
	}

	return resourceSwitchControllerGlobalRead(ctx, d, meta)
}

func resourceSwitchControllerGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSwitchControllerGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerGlobal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerGlobal")
	}

	return resourceSwitchControllerGlobalRead(ctx, d, meta)
}

func resourceSwitchControllerGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSwitchControllerGlobal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSwitchControllerGlobal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerGlobal resource: %v", err)
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

	diags := refreshObjectSwitchControllerGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSwitchControllerGlobalCustomCommand(v *[]models.SwitchControllerGlobalCustomCommand, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CommandEntry; tmp != nil {
				v["command_entry"] = *tmp
			}

			if tmp := cfg.CommandName; tmp != nil {
				v["command_name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "command_entry")
	}

	return flat
}

func flattenSwitchControllerGlobalDisableDiscovery(v *[]models.SwitchControllerGlobalDisableDiscovery, sort bool) interface{} {
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

func refreshObjectSwitchControllerGlobal(d *schema.ResourceData, o *models.SwitchControllerGlobal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowMultipleInterfaces != nil {
		v := *o.AllowMultipleInterfaces

		if err = d.Set("allow_multiple_interfaces", v); err != nil {
			return diag.Errorf("error reading allow_multiple_interfaces: %v", err)
		}
	}

	if o.BounceQuarantinedLink != nil {
		v := *o.BounceQuarantinedLink

		if err = d.Set("bounce_quarantined_link", v); err != nil {
			return diag.Errorf("error reading bounce_quarantined_link: %v", err)
		}
	}

	if o.CustomCommand != nil {
		if err = d.Set("custom_command", flattenSwitchControllerGlobalCustomCommand(o.CustomCommand, sort)); err != nil {
			return diag.Errorf("error reading custom_command: %v", err)
		}
	}

	if o.DefaultVirtualSwitchVlan != nil {
		v := *o.DefaultVirtualSwitchVlan

		if err = d.Set("default_virtual_switch_vlan", v); err != nil {
			return diag.Errorf("error reading default_virtual_switch_vlan: %v", err)
		}
	}

	if o.DhcpServerAccessList != nil {
		v := *o.DhcpServerAccessList

		if err = d.Set("dhcp_server_access_list", v); err != nil {
			return diag.Errorf("error reading dhcp_server_access_list: %v", err)
		}
	}

	if o.DisableDiscovery != nil {
		if err = d.Set("disable_discovery", flattenSwitchControllerGlobalDisableDiscovery(o.DisableDiscovery, sort)); err != nil {
			return diag.Errorf("error reading disable_discovery: %v", err)
		}
	}

	if o.FipsEnforce != nil {
		v := *o.FipsEnforce

		if err = d.Set("fips_enforce", v); err != nil {
			return diag.Errorf("error reading fips_enforce: %v", err)
		}
	}

	if o.HttpsImagePush != nil {
		v := *o.HttpsImagePush

		if err = d.Set("https_image_push", v); err != nil {
			return diag.Errorf("error reading https_image_push: %v", err)
		}
	}

	if o.LogMacLimitViolations != nil {
		v := *o.LogMacLimitViolations

		if err = d.Set("log_mac_limit_violations", v); err != nil {
			return diag.Errorf("error reading log_mac_limit_violations: %v", err)
		}
	}

	if o.MacAgingInterval != nil {
		v := *o.MacAgingInterval

		if err = d.Set("mac_aging_interval", v); err != nil {
			return diag.Errorf("error reading mac_aging_interval: %v", err)
		}
	}

	if o.MacEventLogging != nil {
		v := *o.MacEventLogging

		if err = d.Set("mac_event_logging", v); err != nil {
			return diag.Errorf("error reading mac_event_logging: %v", err)
		}
	}

	if o.MacRetentionPeriod != nil {
		v := *o.MacRetentionPeriod

		if err = d.Set("mac_retention_period", v); err != nil {
			return diag.Errorf("error reading mac_retention_period: %v", err)
		}
	}

	if o.MacViolationTimer != nil {
		v := *o.MacViolationTimer

		if err = d.Set("mac_violation_timer", v); err != nil {
			return diag.Errorf("error reading mac_violation_timer: %v", err)
		}
	}

	if o.QuarantineMode != nil {
		v := *o.QuarantineMode

		if err = d.Set("quarantine_mode", v); err != nil {
			return diag.Errorf("error reading quarantine_mode: %v", err)
		}
	}

	if o.SnDnsResolution != nil {
		v := *o.SnDnsResolution

		if err = d.Set("sn_dns_resolution", v); err != nil {
			return diag.Errorf("error reading sn_dns_resolution: %v", err)
		}
	}

	if o.UpdateUserDevice != nil {
		v := *o.UpdateUserDevice

		if err = d.Set("update_user_device", v); err != nil {
			return diag.Errorf("error reading update_user_device: %v", err)
		}
	}

	if o.VlanAllMode != nil {
		v := *o.VlanAllMode

		if err = d.Set("vlan_all_mode", v); err != nil {
			return diag.Errorf("error reading vlan_all_mode: %v", err)
		}
	}

	if o.VlanOptimization != nil {
		v := *o.VlanOptimization

		if err = d.Set("vlan_optimization", v); err != nil {
			return diag.Errorf("error reading vlan_optimization: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerGlobalCustomCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerGlobalCustomCommand, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerGlobalCustomCommand

	for i := range l {
		tmp := models.SwitchControllerGlobalCustomCommand{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.command_entry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CommandEntry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.command_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CommandName = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerGlobalDisableDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerGlobalDisableDiscovery, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerGlobalDisableDiscovery

	for i := range l {
		tmp := models.SwitchControllerGlobalDisableDiscovery{}
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

func getObjectSwitchControllerGlobal(d *schema.ResourceData, sv string) (*models.SwitchControllerGlobal, diag.Diagnostics) {
	obj := models.SwitchControllerGlobal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allow_multiple_interfaces"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("allow_multiple_interfaces", sv)
				diags = append(diags, e)
			}
			obj.AllowMultipleInterfaces = &v2
		}
	}
	if v1, ok := d.GetOk("bounce_quarantined_link"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bounce_quarantined_link", sv)
				diags = append(diags, e)
			}
			obj.BounceQuarantinedLink = &v2
		}
	}
	if v, ok := d.GetOk("custom_command"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("custom_command", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerGlobalCustomCommand(d, v, "custom_command", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CustomCommand = t
		}
	} else if d.HasChange("custom_command") {
		old, new := d.GetChange("custom_command")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CustomCommand = &[]models.SwitchControllerGlobalCustomCommand{}
		}
	}
	if v1, ok := d.GetOk("default_virtual_switch_vlan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_virtual_switch_vlan", sv)
				diags = append(diags, e)
			}
			obj.DefaultVirtualSwitchVlan = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_server_access_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("dhcp_server_access_list", sv)
				diags = append(diags, e)
			}
			obj.DhcpServerAccessList = &v2
		}
	}
	if v, ok := d.GetOk("disable_discovery"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("disable_discovery", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerGlobalDisableDiscovery(d, v, "disable_discovery", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DisableDiscovery = t
		}
	} else if d.HasChange("disable_discovery") {
		old, new := d.GetChange("disable_discovery")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DisableDiscovery = &[]models.SwitchControllerGlobalDisableDiscovery{}
		}
	}
	if v1, ok := d.GetOk("fips_enforce"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("fips_enforce", sv)
				diags = append(diags, e)
			}
			obj.FipsEnforce = &v2
		}
	}
	if v1, ok := d.GetOk("https_image_push"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_image_push", sv)
				diags = append(diags, e)
			}
			obj.HttpsImagePush = &v2
		}
	}
	if v1, ok := d.GetOk("log_mac_limit_violations"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_mac_limit_violations", sv)
				diags = append(diags, e)
			}
			obj.LogMacLimitViolations = &v2
		}
	}
	if v1, ok := d.GetOk("mac_aging_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_aging_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MacAgingInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("mac_event_logging"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_event_logging", sv)
				diags = append(diags, e)
			}
			obj.MacEventLogging = &v2
		}
	}
	if v1, ok := d.GetOk("mac_retention_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_retention_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MacRetentionPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("mac_violation_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_violation_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MacViolationTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("quarantine_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("quarantine_mode", sv)
				diags = append(diags, e)
			}
			obj.QuarantineMode = &v2
		}
	}
	if v1, ok := d.GetOk("sn_dns_resolution"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sn_dns_resolution", sv)
				diags = append(diags, e)
			}
			obj.SnDnsResolution = &v2
		}
	}
	if v1, ok := d.GetOk("update_user_device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("update_user_device", sv)
				diags = append(diags, e)
			}
			obj.UpdateUserDevice = &v2
		}
	}
	if v1, ok := d.GetOk("vlan_all_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_all_mode", sv)
				diags = append(diags, e)
			}
			obj.VlanAllMode = &v2
		}
	}
	if v1, ok := d.GetOk("vlan_optimization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_optimization", sv)
				diags = append(diags, e)
			}
			obj.VlanOptimization = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSwitchControllerGlobal(d *schema.ResourceData, sv string) (*models.SwitchControllerGlobal, diag.Diagnostics) {
	obj := models.SwitchControllerGlobal{}
	diags := diag.Diagnostics{}

	obj.CustomCommand = &[]models.SwitchControllerGlobalCustomCommand{}
	obj.DisableDiscovery = &[]models.SwitchControllerGlobalDisableDiscovery{}

	return &obj, diags
}
