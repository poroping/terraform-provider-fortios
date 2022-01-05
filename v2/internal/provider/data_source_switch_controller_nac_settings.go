// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure integrated NAC settings for FortiSwitch.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerNacSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure integrated NAC settings for FortiSwitch.",

		ReadContext: dataSourceSwitchControllerNacSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auto_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAC device auto authorization when discovered and nac-policy matched.",
				Computed:    true,
			},
			"bounce_nac_port": {
				Type:        schema.TypeString,
				Description: "Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.",
				Computed:    true,
			},
			"inactive_timer": {
				Type:        schema.TypeInt,
				Description: "Time interval(minutes, 0 = no expiry) to be included in the inactive NAC devices expiry calculation (mac age-out + inactive-time + periodic scan interval).",
				Computed:    true,
			},
			"link_down_flush": {
				Type:        schema.TypeString,
				Description: "Clear NAC devices on switch ports on link down event.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "Set NAC mode to be used on the FortiSwitch ports.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "NAC settings name.",
				Required:    true,
			},
			"onboarding_vlan": {
				Type:        schema.TypeString,
				Description: "Default NAC Onboarding VLAN when NAC devices are discovered.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerNacSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.Errorf("error reading SwitchControllerNacSettings dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerNacSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
