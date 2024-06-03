// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch LLDP settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerLldpSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch LLDP settings.",

		ReadContext: dataSourceSwitchControllerLldpSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"device_detection": {
				Type:        schema.TypeString,
				Description: "Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment.",
				Computed:    true,
			},
			"fast_start_interval": {
				Type:        schema.TypeInt,
				Description: "Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).",
				Computed:    true,
			},
			"management_interface": {
				Type:        schema.TypeString,
				Description: "Primary management interface to be advertised in LLDP and CDP PDUs.",
				Computed:    true,
			},
			"tx_hold": {
				Type:        schema.TypeInt,
				Description: "Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.",
				Computed:    true,
			},
			"tx_interval": {
				Type:        schema.TypeInt,
				Description: "Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerLldpSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SwitchControllerLldpSettings"

	o, err := c.Cmdb.ReadSwitchControllerLldpSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerLldpSettings dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerLldpSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
