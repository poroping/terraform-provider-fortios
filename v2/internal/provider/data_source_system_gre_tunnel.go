// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure GRE tunnel.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemGreTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure GRE tunnel.",

		ReadContext: dataSourceSystemGreTunnelRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"checksum_reception": {
				Type:        schema.TypeString,
				Description: "Enable/disable validating checksums in received GRE packets.",
				Computed:    true,
			},
			"checksum_transmission": {
				Type:        schema.TypeString,
				Description: "Enable/disable including checksums in transmitted GRE packets.",
				Computed:    true,
			},
			"diffservcode": {
				Type:        schema.TypeString,
				Description: "DiffServ setting to be applied to GRE tunnel outer IP header.",
				Computed:    true,
			},
			"dscp_copying": {
				Type:        schema.TypeString,
				Description: "Enable/disable DSCP copying.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Computed:    true,
			},
			"ip_version": {
				Type:        schema.TypeString,
				Description: "IP version to use for VPN interface.",
				Computed:    true,
			},
			"keepalive_failtimes": {
				Type:        schema.TypeInt,
				Description: "Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).",
				Computed:    true,
			},
			"keepalive_interval": {
				Type:        schema.TypeInt,
				Description: "Keepalive message interval (0 - 32767, 0 = disabled).",
				Computed:    true,
			},
			"key_inbound": {
				Type:        schema.TypeInt,
				Description: "Require received GRE packets contain this key (0 - 4294967295).",
				Computed:    true,
			},
			"key_outbound": {
				Type:        schema.TypeInt,
				Description: "Include this key in transmitted GRE packets (0 - 4294967295).",
				Computed:    true,
			},
			"local_gw": {
				Type:        schema.TypeString,
				Description: "IP address of the local gateway.",
				Computed:    true,
			},
			"local_gw6": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the local gateway.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Tunnel name.",
				Required:    true,
			},
			"remote_gw": {
				Type:        schema.TypeString,
				Description: "IP address of the remote gateway.",
				Computed:    true,
			},
			"remote_gw6": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the remote gateway.",
				Computed:    true,
			},
			"sequence_number_reception": {
				Type:        schema.TypeString,
				Description: "Enable/disable validating sequence numbers in received GRE packets.",
				Computed:    true,
			},
			"sequence_number_transmission": {
				Type:        schema.TypeString,
				Description: "Enable/disable including of sequence numbers in transmitted GRE packets.",
				Computed:    true,
			},
			"use_sdwan": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of SD-WAN to reach remote gateway.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemGreTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemGreTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemGreTunnel dataSource: %v", err)
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

	diags := refreshObjectSystemGreTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
