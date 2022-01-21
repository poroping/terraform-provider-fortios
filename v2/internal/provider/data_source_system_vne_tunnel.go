// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual network enabler tunnel.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemVneTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual network enabler tunnel.",

		ReadContext: dataSourceSystemVneTunnelRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bmr_hostname": {
				Type:        schema.TypeString,
				Description: "BMR hostname.",
				Computed:    true,
				Sensitive:   true,
			},
			"br": {
				Type:        schema.TypeString,
				Description: "Border relay IPv6 address.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Computed:    true,
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Description: "Tunnel IPv4 address and netmask.",
				Computed:    true,
			},
			"mode": {
				Type:        schema.TypeString,
				Description: "VNE tunnel mode.",
				Computed:    true,
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Description: "Name of local certificate for SSL connections.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable VNE tunnel.",
				Computed:    true,
			},
			"update_url": {
				Type:        schema.TypeString,
				Description: "URL of provisioning server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemVneTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemVneTunnel"

	o, err := c.Cmdb.ReadSystemVneTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVneTunnel dataSource: %v", err)
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

	diags := refreshObjectSystemVneTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
