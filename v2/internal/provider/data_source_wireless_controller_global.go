// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless controller global settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless controller global settings.",

		ReadContext: dataSourceWirelessControllerGlobalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ap_log_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable configuring FortiGate to redirect wireless event log messages or FortiAPs to send UTM log messages to a syslog server (default = disable).",
				Computed:    true,
			},
			"ap_log_server_ip": {
				Type:        schema.TypeString,
				Description: "IP address that FortiGate or FortiAPs send log messages to.",
				Computed:    true,
			},
			"ap_log_server_port": {
				Type:        schema.TypeInt,
				Description: "Port that FortiGate or FortiAPs send log messages to.",
				Computed:    true,
			},
			"control_message_offload": {
				Type:        schema.TypeString,
				Description: "Configure CAPWAP control message data channel offload.",
				Computed:    true,
			},
			"data_ethernet_ii": {
				Type:        schema.TypeString,
				Description: "Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = enable).",
				Computed:    true,
			},
			"discovery_mc_addr": {
				Type:        schema.TypeString,
				Description: "Multicast IP address for AP discovery (default = 244.0.1.140).",
				Computed:    true,
			},
			"fiapp_eth_type": {
				Type:        schema.TypeInt,
				Description: "Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).",
				Computed:    true,
			},
			"image_download": {
				Type:        schema.TypeString,
				Description: "Enable/disable WTP image download at join time.",
				Computed:    true,
			},
			"ipsec_base_ip": {
				Type:        schema.TypeString,
				Description: "Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).",
				Computed:    true,
			},
			"link_aggregation": {
				Type:        schema.TypeString,
				Description: "Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable).",
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Description of the location of the wireless controller.",
				Computed:    true,
			},
			"max_clients": {
				Type:        schema.TypeInt,
				Description: "Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).",
				Computed:    true,
			},
			"max_retransmit": {
				Type:        schema.TypeInt,
				Description: "Maximum number of tunnel packet retransmissions (0 - 64, default = 3).",
				Computed:    true,
			},
			"mesh_eth_type": {
				Type:        schema.TypeInt,
				Description: "Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).",
				Computed:    true,
			},
			"nac_interval": {
				Type:        schema.TypeInt,
				Description: "Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the wireless controller.",
				Computed:    true,
			},
			"rogue_scan_mac_adjacency": {
				Type:        schema.TypeInt,
				Description: "Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).",
				Computed:    true,
			},
			"tunnel_mode": {
				Type:        schema.TypeString,
				Description: "Compatible/strict tunnel mode.",
				Computed:    true,
			},
			"wtp_share": {
				Type:        schema.TypeString,
				Description: "Enable/disable sharing of WTPs between VDOMs.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WirelessControllerGlobal"

	o, err := c.Cmdb.ReadWirelessControllerGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerGlobal dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
