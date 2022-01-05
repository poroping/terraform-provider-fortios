// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure integrated FortiLink settings for FortiSwitch.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerFortilinkSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure integrated FortiLink settings for FortiSwitch.",

		ReadContext: dataSourceSwitchControllerFortilinkSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fortilink": {
				Type:        schema.TypeString,
				Description: "FortiLink interface to which this fortilink-setting belongs.",
				Computed:    true,
			},
			"inactive_timer": {
				Type:        schema.TypeInt,
				Description: "Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).",
				Computed:    true,
			},
			"link_down_flush": {
				Type:        schema.TypeString,
				Description: "Clear NAC and dynamic devices on switch ports on link down event.",
				Computed:    true,
			},
			"nac_ports": {
				Type:        schema.TypeList,
				Description: "NAC specific configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bounce_nac_port": {
							Type:        schema.TypeString,
							Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.",
							Computed:    true,
						},
						"lan_segment": {
							Type:        schema.TypeString,
							Description: "Enable/disable LAN segment feature on the FortiLink interface.",
							Computed:    true,
						},
						"member_change": {
							Type:        schema.TypeInt,
							Description: "Member change flag.",
							Computed:    true,
						},
						"nac_lan_interface": {
							Type:        schema.TypeString,
							Description: "Configure NAC LAN interface.",
							Computed:    true,
						},
						"nac_segment_vlans": {
							Type:        schema.TypeList,
							Description: "Configure NAC segment VLANs.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": {
										Type:        schema.TypeString,
										Description: "VLAN interface name.",
										Computed:    true,
									},
								},
							},
						},
						"onboarding_vlan": {
							Type:        schema.TypeString,
							Description: "Default NAC Onboarding VLAN when NAC devices are discovered.",
							Computed:    true,
						},
						"parent_key": {
							Type:        schema.TypeString,
							Description: "Parent key name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "FortiLink settings name.",
				Required:    true,
			},
		},
	}
}

func dataSourceSwitchControllerFortilinkSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSwitchControllerFortilinkSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerFortilinkSettings dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerFortilinkSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
