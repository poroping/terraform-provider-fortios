// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure explicit FTP proxy settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFtpProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Description: "Configure explicit FTP proxy settings.",

		ReadContext: dataSourceFtpProxyExplicitRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"incoming_ip": {
				Type:        schema.TypeString,
				Description: "Accept incoming FTP requests from this IP address. An interface must have this IP address.",
				Computed:    true,
			},
			"incoming_port": {
				Type:        schema.TypeString,
				Description: "Accept incoming FTP requests on one or more ports.",
				Computed:    true,
			},
			"outgoing_ip": {
				Type:        schema.TypeString,
				Description: "Outgoing FTP requests will leave from this IP address. An interface must have this IP address.",
				Computed:    true,
			},
			"sec_default_action": {
				Type:        schema.TypeString,
				Description: "Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists.",
				Computed:    true,
			},
			"ssl": {
				Type:        schema.TypeString,
				Description: "Enable/disable the explicit FTPS proxy.",
				Computed:    true,
			},
			"ssl_algorithm": {
				Type:        schema.TypeString,
				Description: "Relative strength of encryption algorithms accepted in negotiation.",
				Computed:    true,
			},
			"ssl_cert": {
				Type:        schema.TypeString,
				Description: "Name of certificate for SSL connections to this server (default = \"Fortinet_CA_SSL\").",
				Computed:    true,
			},
			"ssl_dh_bits": {
				Type:        schema.TypeString,
				Description: "Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the explicit FTP proxy.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFtpProxyExplicitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "FtpProxyExplicit"

	o, err := c.Cmdb.ReadFtpProxyExplicit(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FtpProxyExplicit dataSource: %v", err)
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

	diags := refreshObjectFtpProxyExplicit(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
