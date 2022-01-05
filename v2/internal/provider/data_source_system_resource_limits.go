// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure resource limits.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemResourceLimits() *schema.Resource {
	return &schema.Resource{
		Description: "Configure resource limits.",

		ReadContext: dataSourceSystemResourceLimitsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"custom_service": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall custom services.",
				Computed:    true,
			},
			"dialup_tunnel": {
				Type:        schema.TypeInt,
				Description: "Maximum number of dial-up tunnels.",
				Computed:    true,
			},
			"firewall_address": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall addresses (IPv4, IPv6, multicast).",
				Computed:    true,
			},
			"firewall_addrgrp": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall address groups (IPv4, IPv6).",
				Computed:    true,
			},
			"firewall_policy": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall policies (policy, DoS-policy4, DoS-policy6, multicast).",
				Computed:    true,
			},
			"ipsec_phase1": {
				Type:        schema.TypeInt,
				Description: "Maximum number of VPN IPsec phase1 tunnels.",
				Computed:    true,
			},
			"ipsec_phase1_interface": {
				Type:        schema.TypeInt,
				Description: "Maximum number of VPN IPsec phase1 interface tunnels.",
				Computed:    true,
			},
			"ipsec_phase2": {
				Type:        schema.TypeInt,
				Description: "Maximum number of VPN IPsec phase2 tunnels.",
				Computed:    true,
			},
			"ipsec_phase2_interface": {
				Type:        schema.TypeInt,
				Description: "Maximum number of VPN IPsec phase2 interface tunnels.",
				Computed:    true,
			},
			"log_disk_quota": {
				Type:        schema.TypeInt,
				Description: "Log disk quota in MiB.",
				Computed:    true,
			},
			"onetime_schedule": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall one-time schedules.",
				Computed:    true,
			},
			"proxy": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent proxy users.",
				Computed:    true,
			},
			"recurring_schedule": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall recurring schedules.",
				Computed:    true,
			},
			"service_group": {
				Type:        schema.TypeInt,
				Description: "Maximum number of firewall service groups.",
				Computed:    true,
			},
			"session": {
				Type:        schema.TypeInt,
				Description: "Maximum number of sessions.",
				Computed:    true,
			},
			"sslvpn": {
				Type:        schema.TypeInt,
				Description: "Maximum number of SSL-VPN.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeInt,
				Description: "Maximum number of local users.",
				Computed:    true,
			},
			"user_group": {
				Type:        schema.TypeInt,
				Description: "Maximum number of user groups.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemResourceLimitsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemResourceLimits"

	o, err := c.Cmdb.ReadSystemResourceLimits(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemResourceLimits dataSource: %v", err)
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

	diags := refreshObjectSystemResourceLimits(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
