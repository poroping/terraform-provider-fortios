// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSL proxy settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallSslSetting() *schema.Resource {
	return &schema.Resource{
		Description: "SSL proxy settings.",

		ReadContext: dataSourceFirewallSslSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"abbreviate_handshake": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of SSL abbreviated handshake.",
				Computed:    true,
			},
			"cert_cache_capacity": {
				Type:        schema.TypeInt,
				Description: "Maximum capacity of the host certificate cache (0 - 500, default = 200).",
				Computed:    true,
			},
			"cert_cache_timeout": {
				Type:        schema.TypeInt,
				Description: "Time limit to keep certificate cache (1 - 120 min, default = 10).",
				Computed:    true,
			},
			"kxp_queue_threshold": {
				Type:        schema.TypeInt,
				Description: "Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).",
				Computed:    true,
			},
			"no_matching_cipher_action": {
				Type:        schema.TypeString,
				Description: "Bypass or drop the connection when no matching cipher is found.",
				Computed:    true,
			},
			"proxy_connect_timeout": {
				Type:        schema.TypeInt,
				Description: "Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).",
				Computed:    true,
			},
			"session_cache_capacity": {
				Type:        schema.TypeInt,
				Description: "Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).",
				Computed:    true,
			},
			"session_cache_timeout": {
				Type:        schema.TypeInt,
				Description: "Time limit to keep SSL session state (1 - 60 min, default = 20).",
				Computed:    true,
			},
			"ssl_dh_bits": {
				Type:        schema.TypeString,
				Description: "Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).",
				Computed:    true,
			},
			"ssl_queue_threshold": {
				Type:        schema.TypeInt,
				Description: "Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).",
				Computed:    true,
			},
			"ssl_send_empty_frags": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallSslSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "FirewallSslSetting"

	o, err := c.Cmdb.ReadFirewallSslSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSslSetting dataSource: %v", err)
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

	diags := refreshObjectFirewallSslSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
