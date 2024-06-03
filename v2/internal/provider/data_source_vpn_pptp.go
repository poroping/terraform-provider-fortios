// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure PPTP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnPptp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure PPTP.",

		ReadContext: dataSourceVpnPptpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"eip": {
				Type:        schema.TypeString,
				Description: "End IP.",
				Computed:    true,
			},
			"ip_mode": {
				Type:        schema.TypeString,
				Description: "IP assignment mode for PPTP client.",
				Computed:    true,
			},
			"local_ip": {
				Type:        schema.TypeString,
				Description: "Local IP to be used for peer's remote IP.",
				Computed:    true,
			},
			"sip": {
				Type:        schema.TypeString,
				Description: "Start IP.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiGate as a PPTP gateway.",
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

func dataSourceVpnPptpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "VpnPptp"

	o, err := c.Cmdb.ReadVpnPptp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnPptp dataSource: %v", err)
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

	diags := refreshObjectVpnPptp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
