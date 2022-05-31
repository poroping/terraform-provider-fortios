// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 access proxy.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallAccessProxy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 access proxy.",

		ReadContext: dataSourceFirewallAccessProxy6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"api_gateway": {
				Type:        schema.TypeList,
				Description: "Set IPv4 API Gateway.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_cookie_age": {
							Type:        schema.TypeInt,
							Description: "Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.",
							Computed:    true,
						},
						"http_cookie_domain": {
							Type:        schema.TypeString,
							Description: "Domain that HTTP cookie persistence should apply to.",
							Computed:    true,
						},
						"http_cookie_domain_from_host": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of HTTP cookie domain from host field in HTTP.",
							Computed:    true,
						},
						"http_cookie_generation": {
							Type:        schema.TypeInt,
							Description: "Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.",
							Computed:    true,
						},
						"http_cookie_path": {
							Type:        schema.TypeString,
							Description: "Limit HTTP cookie persistence to the specified path.",
							Computed:    true,
						},
						"http_cookie_share": {
							Type:        schema.TypeString,
							Description: "Control sharing of cookies across API Gateway. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
							Computed:    true,
						},
						"https_cookie_secure": {
							Type:        schema.TypeString,
							Description: "Enable/disable verification that inserted HTTPS cookies are secure.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "API Gateway ID.",
							Computed:    true,
						},
						"ldb_method": {
							Type:        schema.TypeString,
							Description: "Method used to distribute sessions to real servers.",
							Computed:    true,
						},
						"persistence": {
							Type:        schema.TypeString,
							Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
							Computed:    true,
						},
						"realservers": {
							Type:        schema.TypeList,
							Description: "Select the real servers that this Access Proxy will distribute traffic to.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr_type": {
										Type:        schema.TypeString,
										Description: "Type of address.",
										Computed:    true,
									},
									"address": {
										Type:        schema.TypeString,
										Description: "Address or address group of the real server.",
										Computed:    true,
									},
									"domain": {
										Type:        schema.TypeString,
										Description: "Wildcard domain name of the real server.",
										Computed:    true,
									},
									"health_check": {
										Type:        schema.TypeString,
										Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
										Computed:    true,
									},
									"health_check_proto": {
										Type:        schema.TypeString,
										Description: "Protocol of the health check monitor to use when polling to determine server's connectivity status.",
										Computed:    true,
									},
									"holddown_interval": {
										Type:        schema.TypeString,
										Description: "Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds).",
										Computed:    true,
									},
									"http_host": {
										Type:        schema.TypeString,
										Description: "HTTP server domain name in HTTP header.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Real server ID.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "IP address of the real server.",
										Computed:    true,
									},
									"mappedport": {
										Type:        schema.TypeString,
										Description: "Port for communicating with the real server.",
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeInt,
										Description: "Port for communicating with the real server.",
										Computed:    true,
									},
									"ssh_client_cert": {
										Type:        schema.TypeString,
										Description: "Set access-proxy SSH client certificate profile.",
										Computed:    true,
									},
									"ssh_host_key": {
										Type:        schema.TypeList,
										Description: "One or more server host key.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Server host key name.",
													Computed:    true,
												},
											},
										},
									},
									"ssh_host_key_validation": {
										Type:        schema.TypeString,
										Description: "Enable/disable SSH real server host key validation.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "TCP forwarding server type.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
										Computed:    true,
									},
								},
							},
						},
						"saml_redirect": {
							Type:        schema.TypeString,
							Description: "Enable/disable SAML redirection after successful authentication.",
							Computed:    true,
						},
						"saml_server": {
							Type:        schema.TypeString,
							Description: "SAML service provider configuration for VIP authentication.",
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeString,
							Description: "Service.",
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:        schema.TypeString,
							Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
							Computed:    true,
						},
						"ssl_cipher_suites": {
							Type:        schema.TypeList,
							Description: "SSL/TLS cipher suites to offer to a server, ordered by priority.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": {
										Type:        schema.TypeString,
										Description: "Cipher suite name.",
										Computed:    true,
									},
									"priority": {
										Type:        schema.TypeInt,
										Description: "SSL/TLS cipher suites priority.",
										Computed:    true,
									},
									"versions": {
										Type:        schema.TypeString,
										Description: "SSL/TLS versions that the cipher suite can be used with.",
										Computed:    true,
									},
								},
							},
						},
						"ssl_dh_bits": {
							Type:        schema.TypeString,
							Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
							Computed:    true,
						},
						"ssl_max_version": {
							Type:        schema.TypeString,
							Description: "Highest SSL/TLS version acceptable from a server.",
							Computed:    true,
						},
						"ssl_min_version": {
							Type:        schema.TypeString,
							Description: "Lowest SSL/TLS version acceptable from a server.",
							Computed:    true,
						},
						"ssl_vpn_web_portal": {
							Type:        schema.TypeString,
							Description: "SSL-VPN web portal.",
							Computed:    true,
						},
						"url_map": {
							Type:        schema.TypeString,
							Description: "URL pattern to match.",
							Computed:    true,
						},
						"url_map_type": {
							Type:        schema.TypeString,
							Description: "Type of url-map.",
							Computed:    true,
						},
						"virtual_host": {
							Type:        schema.TypeString,
							Description: "Virtual host.",
							Computed:    true,
						},
					},
				},
			},
			"api_gateway6": {
				Type:        schema.TypeList,
				Description: "Set IPv6 API Gateway.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_cookie_age": {
							Type:        schema.TypeInt,
							Description: "Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.",
							Computed:    true,
						},
						"http_cookie_domain": {
							Type:        schema.TypeString,
							Description: "Domain that HTTP cookie persistence should apply to.",
							Computed:    true,
						},
						"http_cookie_domain_from_host": {
							Type:        schema.TypeString,
							Description: "Enable/disable use of HTTP cookie domain from host field in HTTP.",
							Computed:    true,
						},
						"http_cookie_generation": {
							Type:        schema.TypeInt,
							Description: "Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.",
							Computed:    true,
						},
						"http_cookie_path": {
							Type:        schema.TypeString,
							Description: "Limit HTTP cookie persistence to the specified path.",
							Computed:    true,
						},
						"http_cookie_share": {
							Type:        schema.TypeString,
							Description: "Control sharing of cookies across API Gateway. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
							Computed:    true,
						},
						"https_cookie_secure": {
							Type:        schema.TypeString,
							Description: "Enable/disable verification that inserted HTTPS cookies are secure.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "API Gateway ID.",
							Computed:    true,
						},
						"ldb_method": {
							Type:        schema.TypeString,
							Description: "Method used to distribute sessions to real servers.",
							Computed:    true,
						},
						"persistence": {
							Type:        schema.TypeString,
							Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
							Computed:    true,
						},
						"realservers": {
							Type:        schema.TypeList,
							Description: "Select the real servers that this Access Proxy will distribute traffic to.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr_type": {
										Type:        schema.TypeString,
										Description: "Type of address.",
										Computed:    true,
									},
									"address": {
										Type:        schema.TypeString,
										Description: "Address or address group of the real server.",
										Computed:    true,
									},
									"domain": {
										Type:        schema.TypeString,
										Description: "Wildcard domain name of the real server.",
										Computed:    true,
									},
									"health_check": {
										Type:        schema.TypeString,
										Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
										Computed:    true,
									},
									"health_check_proto": {
										Type:        schema.TypeString,
										Description: "Protocol of the health check monitor to use when polling to determine server's connectivity status.",
										Computed:    true,
									},
									"holddown_interval": {
										Type:        schema.TypeString,
										Description: "Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds).",
										Computed:    true,
									},
									"http_host": {
										Type:        schema.TypeString,
										Description: "HTTP server domain name in HTTP header.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Real server ID.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "IPv6 address of the real server.",
										Computed:    true,
									},
									"mappedport": {
										Type:        schema.TypeString,
										Description: "Port for communicating with the real server.",
										Computed:    true,
									},
									"port": {
										Type:        schema.TypeInt,
										Description: "Port for communicating with the real server.",
										Computed:    true,
									},
									"ssh_client_cert": {
										Type:        schema.TypeString,
										Description: "Set access-proxy SSH client certificate profile.",
										Computed:    true,
									},
									"ssh_host_key": {
										Type:        schema.TypeList,
										Description: "One or more server host key.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Server host key name.",
													Computed:    true,
												},
											},
										},
									},
									"ssh_host_key_validation": {
										Type:        schema.TypeString,
										Description: "Enable/disable SSH real server host key validation.",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
										Computed:    true,
									},
									"type": {
										Type:        schema.TypeString,
										Description: "TCP forwarding server type.",
										Computed:    true,
									},
									"weight": {
										Type:        schema.TypeInt,
										Description: "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
										Computed:    true,
									},
								},
							},
						},
						"saml_redirect": {
							Type:        schema.TypeString,
							Description: "Enable/disable SAML redirection after successful authentication.",
							Computed:    true,
						},
						"saml_server": {
							Type:        schema.TypeString,
							Description: "SAML service provider configuration for VIP authentication.",
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeString,
							Description: "Service.",
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:        schema.TypeString,
							Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
							Computed:    true,
						},
						"ssl_cipher_suites": {
							Type:        schema.TypeList,
							Description: "SSL/TLS cipher suites to offer to a server, ordered by priority.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": {
										Type:        schema.TypeString,
										Description: "Cipher suite name.",
										Computed:    true,
									},
									"priority": {
										Type:        schema.TypeInt,
										Description: "SSL/TLS cipher suites priority.",
										Computed:    true,
									},
									"versions": {
										Type:        schema.TypeString,
										Description: "SSL/TLS versions that the cipher suite can be used with.",
										Computed:    true,
									},
								},
							},
						},
						"ssl_dh_bits": {
							Type:        schema.TypeString,
							Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
							Computed:    true,
						},
						"ssl_max_version": {
							Type:        schema.TypeString,
							Description: "Highest SSL/TLS version acceptable from a server.",
							Computed:    true,
						},
						"ssl_min_version": {
							Type:        schema.TypeString,
							Description: "Lowest SSL/TLS version acceptable from a server.",
							Computed:    true,
						},
						"ssl_vpn_web_portal": {
							Type:        schema.TypeString,
							Description: "SSL-VPN web portal.",
							Computed:    true,
						},
						"url_map": {
							Type:        schema.TypeString,
							Description: "URL pattern to match.",
							Computed:    true,
						},
						"url_map_type": {
							Type:        schema.TypeString,
							Description: "Type of url-map.",
							Computed:    true,
						},
						"virtual_host": {
							Type:        schema.TypeString,
							Description: "Virtual host.",
							Computed:    true,
						},
					},
				},
			},
			"auth_portal": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication portal.",
				Computed:    true,
			},
			"auth_virtual_host": {
				Type:        schema.TypeString,
				Description: "Virtual host for authentication portal.",
				Computed:    true,
			},
			"client_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable to request client certificate.",
				Computed:    true,
			},
			"decrypted_traffic_mirror": {
				Type:        schema.TypeString,
				Description: "Decrypted traffic mirror.",
				Computed:    true,
			},
			"empty_cert_action": {
				Type:        schema.TypeString,
				Description: "Action of an empty client certificate.",
				Computed:    true,
			},
			"log_blocked_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of blocked traffic.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Access Proxy name.",
				Required:    true,
			},
			"vip": {
				Type:        schema.TypeString,
				Description: "Virtual IP name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAccessProxy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAccessProxy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAccessProxy6 dataSource: %v", err)
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

	diags := refreshObjectFirewallAccessProxy6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
