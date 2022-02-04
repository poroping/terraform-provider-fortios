// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Web proxy global settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWebProxyGlobal() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Web proxy global settings.",

		ReadContext: dataSourceWebProxyGlobalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fast_policy_match": {
				Type:        schema.TypeString,
				Description: "Enable/disable fast matching algorithm for explicit and transparent proxy policy.",
				Computed:    true,
			},
			"forward_proxy_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable forwarding proxy authentication headers.",
				Computed:    true,
			},
			"forward_server_affinity_timeout": {
				Type:        schema.TypeInt,
				Description: "Period of time before the source IP's traffic is no longer assigned to the forwarding server (6 - 60 min, default = 30).",
				Computed:    true,
			},
			"ldap_user_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable LDAP user cache for explicit and transparent proxy user.",
				Computed:    true,
			},
			"learn_client_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable learning the client's IP address from headers.",
				Computed:    true,
			},
			"learn_client_ip_from_header": {
				Type:        schema.TypeString,
				Description: "Learn client IP address from the specified headers.",
				Computed:    true,
			},
			"learn_client_ip_srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name (srcaddr or srcaddr6 must be set).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"learn_client_ip_srcaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 Source address name (srcaddr or srcaddr6 must be set).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"max_message_length": {
				Type:        schema.TypeInt,
				Description: "Maximum length of HTTP message, not including body (16 - 256 Kbytes, default = 32).",
				Computed:    true,
			},
			"max_request_length": {
				Type:        schema.TypeInt,
				Description: "Maximum length of HTTP request line (2 - 64 Kbytes, default = 8).",
				Computed:    true,
			},
			"max_waf_body_cache_length": {
				Type:        schema.TypeInt,
				Description: "Maximum length of HTTP messages processed by Web Application Firewall (WAF) (10 - 1024 Kbytes, default = 32).",
				Computed:    true,
			},
			"proxy_fqdn": {
				Type:        schema.TypeString,
				Description: "Fully Qualified Domain Name (FQDN) that clients connect to (default = default.fqdn) to connect to the explicit web proxy.",
				Computed:    true,
			},
			"src_affinity_exempt_addr": {
				Type:        schema.TypeString,
				Description: "IPv4 source addresses to exempt proxy affinity.",
				Computed:    true,
			},
			"src_affinity_exempt_addr6": {
				Type:        schema.TypeString,
				Description: "IPv6 source addresses to exempt proxy affinity.",
				Computed:    true,
			},
			"ssl_ca_cert": {
				Type:        schema.TypeString,
				Description: "SSL CA certificate for SSL interception.",
				Computed:    true,
			},
			"ssl_cert": {
				Type:        schema.TypeString,
				Description: "SSL certificate for SSL interception.",
				Computed:    true,
			},
			"strict_web_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable strict web checking to block web sites that send incorrect headers that don't conform to HTTP 1.1.",
				Computed:    true,
			},
			"tunnel_non_http": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing non-HTTP traffic. Allowed non-HTTP traffic is tunneled.",
				Computed:    true,
			},
			"unknown_http_version": {
				Type:        schema.TypeString,
				Description: "Action to take when an unknown version of HTTP is encountered: reject, allow (tunnel), or proceed with best-effort.",
				Computed:    true,
			},
			"webproxy_profile": {
				Type:        schema.TypeString,
				Description: "Name of the web proxy profile to apply when explicit proxy traffic is allowed by default and traffic is accepted that does not match an explicit proxy policy.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebProxyGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WebProxyGlobal"

	o, err := c.Cmdb.ReadWebProxyGlobal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyGlobal dataSource: %v", err)
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

	diags := refreshObjectWebProxyGlobal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
