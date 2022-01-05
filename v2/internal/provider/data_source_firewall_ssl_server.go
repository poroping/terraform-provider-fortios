// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL servers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallSslServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL servers.",

		ReadContext: dataSourceFirewallSslServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"add_header_x_forwarded_proto": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding an X-Forwarded-Proto header to forwarded requests.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the SSL server.",
				Computed:    true,
			},
			"mapped_port": {
				Type:        schema.TypeInt,
				Description: "Mapped server service port (1 - 65535, default = 80).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Server name.",
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Server service port (1 - 65535, default = 443).",
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
			"ssl_client_renegotiation": {
				Type:        schema.TypeString,
				Description: "Allow or block client renegotiation by server.",
				Computed:    true,
			},
			"ssl_dh_bits": {
				Type:        schema.TypeString,
				Description: "Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).",
				Computed:    true,
			},
			"ssl_max_version": {
				Type:        schema.TypeString,
				Description: "Highest SSL/TLS version to negotiate.",
				Computed:    true,
			},
			"ssl_min_version": {
				Type:        schema.TypeString,
				Description: "Lowest SSL/TLS version to negotiate.",
				Computed:    true,
			},
			"ssl_mode": {
				Type:        schema.TypeString,
				Description: "SSL/TLS mode for encryption and decryption of traffic.",
				Computed:    true,
			},
			"ssl_send_empty_frags": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending empty fragments to avoid attack on CBC IV.",
				Computed:    true,
			},
			"url_rewrite": {
				Type:        schema.TypeString,
				Description: "Enable/disable rewriting the URL.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallSslServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSslServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSslServer dataSource: %v", err)
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

	diags := refreshObjectFirewallSslServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
