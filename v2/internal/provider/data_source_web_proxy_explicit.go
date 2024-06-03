// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure explicit Web proxy settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWebProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Description: "Configure explicit Web proxy settings.",

		ReadContext: dataSourceWebProxyExplicitRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ftp_incoming_port": {
				Type:        schema.TypeString,
				Description: "Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).",
				Computed:    true,
			},
			"ftp_over_http": {
				Type:        schema.TypeString,
				Description: "Enable to proxy FTP-over-HTTP sessions sent from a web browser.",
				Computed:    true,
			},
			"http_connection_mode": {
				Type:        schema.TypeString,
				Description: "HTTP connection mode (default = static).",
				Computed:    true,
			},
			"http_incoming_port": {
				Type:        schema.TypeString,
				Description: "Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).",
				Computed:    true,
			},
			"https_incoming_port": {
				Type:        schema.TypeString,
				Description: "Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).",
				Computed:    true,
			},
			"https_replacement_message": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending the client a replacement message for HTTPS requests.",
				Computed:    true,
			},
			"incoming_ip": {
				Type:        schema.TypeString,
				Description: "Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.",
				Computed:    true,
			},
			"incoming_ip6": {
				Type:        schema.TypeString,
				Description: "Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.",
				Computed:    true,
			},
			"ipv6_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command.",
				Computed:    true,
			},
			"message_upon_server_error": {
				Type:        schema.TypeString,
				Description: "Enable/disable displaying a replacement message when a server error is detected.",
				Computed:    true,
			},
			"outgoing_ip": {
				Type:        schema.TypeString,
				Description: "Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.",
				Computed:    true,
			},
			"outgoing_ip6": {
				Type:        schema.TypeString,
				Description: "Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.",
				Computed:    true,
			},
			"pac_file_data": {
				Type:        schema.TypeString,
				Description: "PAC file contents enclosed in quotes (maximum of 256K bytes).",
				Computed:    true,
			},
			"pac_file_name": {
				Type:        schema.TypeString,
				Description: "Pac file name.",
				Computed:    true,
			},
			"pac_file_server_port": {
				Type:        schema.TypeString,
				Description: "Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).",
				Computed:    true,
			},
			"pac_file_server_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile.",
				Computed:    true,
			},
			"pac_file_through_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable to get Proxy Auto-Configuration (PAC) through HTTPS.",
				Computed:    true,
			},
			"pac_file_url": {
				Type:        schema.TypeString,
				Description: "PAC file access URL.",
				Computed:    true,
			},
			"pac_policy": {
				Type:        schema.TypeList,
				Description: "PAC policies.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": {
							Type:        schema.TypeString,
							Description: "Optional comments.",
							Computed:    true,
						},
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination address objects.",
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
						"pac_file_data": {
							Type:        schema.TypeString,
							Description: "PAC file contents enclosed in quotes (maximum of 256K bytes).",
							Computed:    true,
						},
						"pac_file_name": {
							Type:        schema.TypeString,
							Description: "Pac file name.",
							Computed:    true,
						},
						"policyid": {
							Type:        schema.TypeInt,
							Description: "Policy ID.",
							Computed:    true,
						},
						"srcaddr": {
							Type:        schema.TypeList,
							Description: "Source address objects.",
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
						"srcaddr6": {
							Type:        schema.TypeList,
							Description: "Source address6 objects.",
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
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable policy.",
							Computed:    true,
						},
					},
				},
			},
			"pref_dns_result": {
				Type:        schema.TypeString,
				Description: "Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4).",
				Computed:    true,
			},
			"realm": {
				Type:        schema.TypeString,
				Description: "Authentication realm used to identify the explicit web proxy (maximum of 63 characters).",
				Computed:    true,
			},
			"sec_default_action": {
				Type:        schema.TypeString,
				Description: "Accept or deny explicit web proxy sessions when no web proxy firewall policy exists.",
				Computed:    true,
			},
			"socks": {
				Type:        schema.TypeString,
				Description: "Enable/disable the SOCKS proxy.",
				Computed:    true,
			},
			"socks_incoming_port": {
				Type:        schema.TypeString,
				Description: "Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).",
				Computed:    true,
			},
			"ssl_algorithm": {
				Type:        schema.TypeString,
				Description: "Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the explicit Web proxy for HTTP and HTTPS session.",
				Computed:    true,
			},
			"strict_guest": {
				Type:        schema.TypeString,
				Description: "Enable/disable strict guest user checking by the explicit web proxy.",
				Computed:    true,
			},
			"trace_auth_no_rsp": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging timed-out authentication requests.",
				Computed:    true,
			},
			"unknown_http_version": {
				Type:        schema.TypeString,
				Description: "How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebProxyExplicitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WebProxyExplicit"

	o, err := c.Cmdb.ReadWebProxyExplicit(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyExplicit dataSource: %v", err)
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

	diags := refreshObjectWebProxyExplicit(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
