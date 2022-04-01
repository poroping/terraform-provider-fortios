// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch global settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch global settings.",

		ReadContext: dataSourceSwitchControllerGlobalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_multiple_interfaces": {
				Type:        schema.TypeString,
				Description: "Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate.",
				Computed:    true,
			},
			"bounce_quarantined_link": {
				Type:        schema.TypeString,
				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device.",
				Computed:    true,
			},
			"custom_command": {
				Type:        schema.TypeList,
				Description: "List of custom commands to be pushed to all FortiSwitches in the VDOM.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": {
							Type:        schema.TypeString,
							Description: "List of FortiSwitch commands.",
							Computed:    true,
						},
						"command_name": {
							Type:        schema.TypeString,
							Description: "Name of custom command to push to all FortiSwitches in VDOM.",
							Computed:    true,
						},
					},
				},
			},
			"default_virtual_switch_vlan": {
				Type:        schema.TypeString,
				Description: "Default VLAN for ports when added to the virtual-switch.",
				Computed:    true,
			},
			"dhcp_server_access_list": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP snooping server access list.",
				Computed:    true,
			},
			"disable_discovery": {
				Type:        schema.TypeList,
				Description: "Prevent this FortiSwitch from discovering.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Managed device ID.",
							Computed:    true,
						},
					},
				},
			},
			"fips_enforce": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforcement of FIPS on managed FortiSwitch devices.",
				Computed:    true,
			},
			"firmware_provision_on_authorization": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic provisioning of latest firmware on authorization.",
				Computed:    true,
			},
			"https_image_push": {
				Type:        schema.TypeString,
				Description: "Enable/disable image push to FortiSwitch using HTTPS.",
				Computed:    true,
			},
			"log_mac_limit_violations": {
				Type:        schema.TypeString,
				Description: "Enable/disable logs for Learning Limit Violations.",
				Computed:    true,
			},
			"mac_aging_interval": {
				Type:        schema.TypeInt,
				Description: "Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).",
				Computed:    true,
			},
			"mac_event_logging": {
				Type:        schema.TypeString,
				Description: "Enable/disable MAC address event logging.",
				Computed:    true,
			},
			"mac_retention_period": {
				Type:        schema.TypeInt,
				Description: "Time in hours after which an inactive MAC is removed from client DB (0 = aged out based on mac-aging-interval).",
				Computed:    true,
			},
			"mac_violation_timer": {
				Type:        schema.TypeInt,
				Description: "Set timeout for Learning Limit Violations (0 = disabled).",
				Computed:    true,
			},
			"quarantine_mode": {
				Type:        schema.TypeString,
				Description: "Quarantine mode.",
				Computed:    true,
			},
			"sn_dns_resolution": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number.",
				Computed:    true,
			},
			"update_user_device": {
				Type:        schema.TypeString,
				Description: "Control which sources update the device user list.",
				Computed:    true,
			},
			"vlan_all_mode": {
				Type:        schema.TypeString,
				Description: "VLAN configuration mode, user-defined-vlans or all-possible-vlans.",
				Computed:    true,
			},
			"vlan_optimization": {
				Type:        schema.TypeString,
				Description: "FortiLink VLAN optimization.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SwitchControllerGlobal"

	o, err := c.Cmdb.ReadSwitchControllerGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerGlobal dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
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

	d.SetId(mkey)

	return nil
}
