// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VPN autokey tunnel.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnIpsecPhase2Interface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VPN autokey tunnel.",

		ReadContext: dataSourceVpnIpsecPhase2InterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"add_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic route addition.",
				Computed:    true,
			},
			"auto_discovery_forwarder": {
				Type:        schema.TypeString,
				Description: "Enable/disable forwarding short-cut messages.",
				Computed:    true,
			},
			"auto_discovery_sender": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending short-cut messages.",
				Computed:    true,
			},
			"auto_negotiate": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec SA auto-negotiation.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dhcp_ipsec": {
				Type:        schema.TypeString,
				Description: "Enable/disable DHCP-IPsec.",
				Computed:    true,
			},
			"dhgrp": {
				Type:        schema.TypeString,
				Description: "Phase2 DH group.",
				Computed:    true,
			},
			"diffserv": {
				Type:        schema.TypeString,
				Description: "Enable/disable applying DSCP value to the IPsec tunnel outer IP header.",
				Computed:    true,
			},
			"diffservcode": {
				Type:        schema.TypeString,
				Description: "DSCP value to be applied to the IPsec tunnel outer IP header.",
				Computed:    true,
			},
			"dst_addr_type": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID type.",
				Computed:    true,
			},
			"dst_end_ip": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID IPv4 end.",
				Computed:    true,
			},
			"dst_end_ip6": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID IPv6 end.",
				Computed:    true,
			},
			"dst_name": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID name.",
				Computed:    true,
			},
			"dst_name6": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID name.",
				Computed:    true,
			},
			"dst_port": {
				Type:        schema.TypeInt,
				Description: "Quick mode destination port (1 - 65535 or 0 for all).",
				Computed:    true,
			},
			"dst_start_ip": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID IPv4 start.",
				Computed:    true,
			},
			"dst_start_ip6": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID IPv6 start.",
				Computed:    true,
			},
			"dst_subnet": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID IPv4 subnet.",
				Computed:    true,
			},
			"dst_subnet6": {
				Type:        schema.TypeString,
				Description: "Remote proxy ID IPv6 subnet.",
				Computed:    true,
			},
			"encapsulation": {
				Type:        schema.TypeString,
				Description: "ESP encapsulation mode.",
				Computed:    true,
			},
			"initiator_ts_narrow": {
				Type:        schema.TypeString,
				Description: "Enable/disable traffic selector narrowing for IKEv2 initiator.",
				Computed:    true,
			},
			"ipv4_df": {
				Type:        schema.TypeString,
				Description: "Enable/disable setting and resetting of IPv4 'Don't Fragment' bit.",
				Computed:    true,
			},
			"keepalive": {
				Type:        schema.TypeString,
				Description: "Enable/disable keep alive.",
				Computed:    true,
			},
			"keylife_type": {
				Type:        schema.TypeString,
				Description: "Keylife type.",
				Computed:    true,
			},
			"keylifekbs": {
				Type:        schema.TypeInt,
				Description: "Phase2 key life in number of kilobytes of traffic (5120 - 4294967295).",
				Computed:    true,
			},
			"keylifeseconds": {
				Type:        schema.TypeInt,
				Description: "Phase2 key life in time in seconds (120 - 172800).",
				Computed:    true,
			},
			"l2tp": {
				Type:        schema.TypeString,
				Description: "Enable/disable L2TP over IPsec.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "IPsec tunnel name.",
				Required:    true,
			},
			"pfs": {
				Type:        schema.TypeString,
				Description: "Enable/disable PFS feature.",
				Computed:    true,
			},
			"phase1name": {
				Type:        schema.TypeString,
				Description: "Phase 1 determines the options required for phase 2.",
				Computed:    true,
			},
			"proposal": {
				Type:        schema.TypeString,
				Description: "Phase2 proposal.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Quick mode protocol selector (1 - 255 or 0 for all).",
				Computed:    true,
			},
			"replay": {
				Type:        schema.TypeString,
				Description: "Enable/disable replay detection.",
				Computed:    true,
			},
			"route_overlap": {
				Type:        schema.TypeString,
				Description: "Action for overlapping routes.",
				Computed:    true,
			},
			"single_source": {
				Type:        schema.TypeString,
				Description: "Enable/disable single source IP restriction.",
				Computed:    true,
			},
			"src_addr_type": {
				Type:        schema.TypeString,
				Description: "Local proxy ID type.",
				Computed:    true,
			},
			"src_end_ip": {
				Type:        schema.TypeString,
				Description: "Local proxy ID end.",
				Computed:    true,
			},
			"src_end_ip6": {
				Type:        schema.TypeString,
				Description: "Local proxy ID IPv6 end.",
				Computed:    true,
			},
			"src_name": {
				Type:        schema.TypeString,
				Description: "Local proxy ID name.",
				Computed:    true,
			},
			"src_name6": {
				Type:        schema.TypeString,
				Description: "Local proxy ID name.",
				Computed:    true,
			},
			"src_port": {
				Type:        schema.TypeInt,
				Description: "Quick mode source port (1 - 65535 or 0 for all).",
				Computed:    true,
			},
			"src_start_ip": {
				Type:        schema.TypeString,
				Description: "Local proxy ID start.",
				Computed:    true,
			},
			"src_start_ip6": {
				Type:        schema.TypeString,
				Description: "Local proxy ID IPv6 start.",
				Computed:    true,
			},
			"src_subnet": {
				Type:        schema.TypeString,
				Description: "Local proxy ID subnet.",
				Computed:    true,
			},
			"src_subnet6": {
				Type:        schema.TypeString,
				Description: "Local proxy ID IPv6 subnet.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnIpsecPhase2InterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecPhase2Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecPhase2Interface dataSource: %v", err)
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

	diags := refreshObjectVpnIpsecPhase2Interface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
