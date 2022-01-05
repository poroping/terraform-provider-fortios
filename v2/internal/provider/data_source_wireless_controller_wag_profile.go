// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless access gateway (WAG) profiles used for tunnels on AP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerWagProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless access gateway (WAG) profiles used for tunnels on AP.",

		ReadContext: dataSourceWirelessControllerWagProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dhcp_ip_addr": {
				Type:        schema.TypeString,
				Description: "IP address of the monitoring DHCP request packet sent through the tunnel",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Tunnel profile name.",
				Required:    true,
			},
			"ping_interval": {
				Type:        schema.TypeInt,
				Description: "Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).",
				Computed:    true,
			},
			"ping_number": {
				Type:        schema.TypeInt,
				Description: "Number of the tunnel mointoring echo packets (1 - 65535, default = 5).",
				Computed:    true,
			},
			"return_packet_timeout": {
				Type:        schema.TypeInt,
				Description: "Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).",
				Computed:    true,
			},
			"tunnel_type": {
				Type:        schema.TypeString,
				Description: "Tunnel type.",
				Computed:    true,
			},
			"wag_ip": {
				Type:        schema.TypeString,
				Description: "IP Address of the wireless access gateway.",
				Computed:    true,
			},
			"wag_port": {
				Type:        schema.TypeInt,
				Description: "UDP port of the wireless access gateway.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerWagProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerWagProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWagProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerWagProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
