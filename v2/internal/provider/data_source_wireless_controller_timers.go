// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure CAPWAP timers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerTimers() *schema.Resource {
	return &schema.Resource{
		Description: "Configure CAPWAP timers.",

		ReadContext: dataSourceWirelessControllerTimersRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ble_scan_report_intv": {
				Type:        schema.TypeInt,
				Description: "Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).",
				Computed:    true,
			},
			"client_idle_timeout": {
				Type:        schema.TypeInt,
				Description: "Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).",
				Computed:    true,
			},
			"discovery_interval": {
				Type:        schema.TypeInt,
				Description: "Time between discovery requests (2 - 180 sec, default = 5).",
				Computed:    true,
			},
			"drma_interval": {
				Type:        schema.TypeInt,
				Description: "Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).",
				Computed:    true,
			},
			"echo_interval": {
				Type:        schema.TypeInt,
				Description: "Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).",
				Computed:    true,
			},
			"fake_ap_log": {
				Type:        schema.TypeInt,
				Description: "Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).",
				Computed:    true,
			},
			"ipsec_intf_cleanup": {
				Type:        schema.TypeInt,
				Description: "Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).",
				Computed:    true,
			},
			"radio_stats_interval": {
				Type:        schema.TypeInt,
				Description: "Time between running radio reports (1 - 255 sec, default = 15).",
				Computed:    true,
			},
			"rogue_ap_log": {
				Type:        schema.TypeInt,
				Description: "Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).",
				Computed:    true,
			},
			"sta_capability_interval": {
				Type:        schema.TypeInt,
				Description: "Time between running station capability reports (1 - 255 sec, default = 30).",
				Computed:    true,
			},
			"sta_locate_timer": {
				Type:        schema.TypeInt,
				Description: "Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).",
				Computed:    true,
			},
			"sta_stats_interval": {
				Type:        schema.TypeInt,
				Description: "Time between running client (station) reports (1 - 255 sec, default = 1).",
				Computed:    true,
			},
			"vap_stats_interval": {
				Type:        schema.TypeInt,
				Description: "Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerTimersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WirelessControllerTimers"

	o, err := c.Cmdb.ReadWirelessControllerTimers(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerTimers dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerTimers(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
