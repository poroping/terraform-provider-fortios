// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VDOM property.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemVdomProperty() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VDOM property.",

		ReadContext: dataSourceSystemVdomPropertyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"custom_service": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall custom services.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"dialup_tunnel": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of dial-up tunnels.",
				Computed:    true,
			},
			"firewall_address": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).",
				Computed:    true,
			},
			"firewall_addrgrp": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall address groups (IPv4, IPv6).",
				Computed:    true,
			},
			"firewall_policy": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).",
				Computed:    true,
			},
			"ipsec_phase1": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of VPN IPsec phase 1 tunnels.",
				Computed:    true,
			},
			"ipsec_phase1_interface": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of VPN IPsec phase1 interface tunnels.",
				Computed:    true,
			},
			"ipsec_phase2": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of VPN IPsec phase 2 tunnels.",
				Computed:    true,
			},
			"ipsec_phase2_interface": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of VPN IPsec phase2 interface tunnels.",
				Computed:    true,
			},
			"log_disk_quota": {
				Type:        schema.TypeString,
				Description: "Log disk quota in MiB (range depends on how much disk space is available).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "VDOM name.",
				Required:    true,
			},
			"onetime_schedule": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall one-time schedules.",
				Computed:    true,
			},
			"proxy": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of concurrent proxy users.",
				Computed:    true,
			},
			"recurring_schedule": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall recurring schedules.",
				Computed:    true,
			},
			"service_group": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of firewall service groups.",
				Computed:    true,
			},
			"session": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of sessions.",
				Computed:    true,
			},
			"snmp_index": {
				Type:        schema.TypeInt,
				Description: "Permanent SNMP Index of the virtual domain (1 - 2147483647).",
				Computed:    true,
			},
			"sslvpn": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of SSL-VPNs.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of local users.",
				Computed:    true,
			},
			"user_group": {
				Type:        schema.TypeString,
				Description: "Maximum guaranteed number of user groups.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemVdomPropertyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVdomProperty(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVdomProperty dataSource: %v", err)
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

	diags := refreshObjectSystemVdomProperty(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
