// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure connection capability.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerHotspot20H2qpConnCapability() *schema.Resource {
	return &schema.Resource{
		Description: "Configure connection capability.",

		ReadContext: dataSourceWirelessControllerHotspot20H2qpConnCapabilityRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"esp_port": {
				Type:        schema.TypeString,
				Description: "Set ESP port service (used by IPsec VPNs) status.",
				Computed:    true,
			},
			"ftp_port": {
				Type:        schema.TypeString,
				Description: "Set FTP port service status.",
				Computed:    true,
			},
			"http_port": {
				Type:        schema.TypeString,
				Description: "Set HTTP port service status.",
				Computed:    true,
			},
			"icmp_port": {
				Type:        schema.TypeString,
				Description: "Set ICMP port service status.",
				Computed:    true,
			},
			"ikev2_port": {
				Type:        schema.TypeString,
				Description: "Set IKEv2 port service for IPsec VPN status.",
				Computed:    true,
			},
			"ikev2_xx_port": {
				Type:        schema.TypeString,
				Description: "Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Connection capability name.",
				Required:    true,
			},
			"pptp_vpn_port": {
				Type:        schema.TypeString,
				Description: "Set Point to Point Tunneling Protocol (PPTP) VPN port service status.",
				Computed:    true,
			},
			"ssh_port": {
				Type:        schema.TypeString,
				Description: "Set SSH port service status.",
				Computed:    true,
			},
			"tls_port": {
				Type:        schema.TypeString,
				Description: "Set TLS VPN (HTTPS) port service status.",
				Computed:    true,
			},
			"voip_tcp_port": {
				Type:        schema.TypeString,
				Description: "Set VoIP TCP port service status.",
				Computed:    true,
			},
			"voip_udp_port": {
				Type:        schema.TypeString,
				Description: "Set VoIP UDP port service status.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerHotspot20H2qpConnCapabilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpConnCapability(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpConnCapability dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpConnCapability(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
