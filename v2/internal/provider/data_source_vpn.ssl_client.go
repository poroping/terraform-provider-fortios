// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: client

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnSslClient() *schema.Resource {
	return &schema.Resource{
		Description: "client",

		ReadContext: dataSourceVpnSslClientRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "Certificate to offer to SSL-VPN server if it requests one.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Distance for routes added by SSL-VPN (1 - 255).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "SSL interface to send/receive traffic over.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SSL-VPN tunnel name.",
				Required:    true,
			},
			"peer": {
				Type:        schema.TypeString,
				Description: "Authenticate peer's certificate with the peer/peergrp.",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN server port.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority for routes added by SSL-VPN (0 - 4294967295).",
				Computed:    true,
			},
			"psk": {
				Type:        schema.TypeString,
				Description: "Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).",
				Computed:    true,
				Sensitive:   true,
			},
			"realm": {
				Type:        schema.TypeString,
				Description: "Realm name configured on SSL-VPN server.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "IPv4, IPv6 or DNS address of the SSL-VPN server.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this SSL-VPN client configuration.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "Username to offer to the peer to authenticate the client.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnSslClientRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslClient(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslClient dataSource: %v", err)
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

	diags := refreshObjectVpnSslClient(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
