// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure L2TP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnL2tp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure L2TP.",

		ReadContext: dataSourceVpnL2tpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"compress": {
				Type:        schema.TypeString,
				Description: "Enable/disable data compression.",
				Computed:    true,
			},
			"eip": {
				Type:        schema.TypeString,
				Description: "End IP.",
				Computed:    true,
			},
			"enforce_ipsec": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPsec enforcement.",
				Computed:    true,
			},
			"hello_interval": {
				Type:        schema.TypeInt,
				Description: "L2TP hello message interval in seconds (0 - 3600 sec, default = 60).",
				Computed:    true,
			},
			"lcp_echo_interval": {
				Type:        schema.TypeInt,
				Description: "Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.",
				Computed:    true,
			},
			"lcp_max_echo_fails": {
				Type:        schema.TypeInt,
				Description: "Maximum number of missed LCP echo messages before disconnect.",
				Computed:    true,
			},
			"sip": {
				Type:        schema.TypeString,
				Description: "Start IP.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGate as a L2TP gateway.",
				Computed:    true,
			},
			"usrgrp": {
				Type:        schema.TypeString,
				Description: "User group.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnL2tpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "VpnL2tp"

	o, err := c.Cmdb.ReadVpnL2tp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnL2tp dataSource: %v", err)
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

	diags := refreshObjectVpnL2tp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
