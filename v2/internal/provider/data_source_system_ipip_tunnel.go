// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP in IP Tunneling.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemIpipTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP in IP Tunneling.",

		ReadContext: dataSourceSystemIpipTunnelRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auto_asic_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable tunnel ASIC offloading.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name that is associated with the incoming traffic from available options.",
				Computed:    true,
			},
			"local_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address for the local gateway.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "IPIP Tunnel name.",
				Required:    true,
			},
			"remote_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address for the remote gateway.",
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

func dataSourceSystemIpipTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpipTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpipTunnel dataSource: %v", err)
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

	diags := refreshObjectSystemIpipTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
