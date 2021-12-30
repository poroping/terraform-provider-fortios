// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 access proxy.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallAccessProxy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 access proxy.",

		CreateContext: resourceFirewallAccessProxy6Create,
		ReadContext:   resourceFirewallAccessProxy6Read,
		UpdateContext: resourceFirewallAccessProxy6Update,
		DeleteContext: resourceFirewallAccessProxy6Delete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"api_gateway": {
				Type:        schema.TypeList,
				Description: "Set IPv4 API Gateway.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_cookie_age": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 525600),

							Description: "Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_domain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Domain that HTTP cookie persistence should apply to.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_domain_from_host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable use of HTTP cookie domain from host field in HTTP.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_generation": {
							Type: schema.TypeInt,

							Description: "Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_path": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Limit HTTP cookie persistence to the specified path.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_share": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "same-ip"}, false),

							Description: "Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
							Optional:    true,
							Computed:    true,
						},
						"https_cookie_secure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable verification that inserted HTTPS cookies are secure.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "API Gateway ID.",
							Optional:    true,
							Computed:    true,
						},
						"ldb_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"static", "round-robin", "weighted", "first-alive", "http-host"}, false),

							Description: "Method used to distribute sessions to real servers.",
							Optional:    true,
							Computed:    true,
						},
						"persistence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http-cookie"}, false),

							Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
							Optional:    true,
							Computed:    true,
						},
						"realservers": {
							Type:        schema.TypeList,
							Description: "Select the real servers that this Access Proxy will distribute traffic to.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

										Description: "Type of address.",
										Optional:    true,
										Computed:    true,
									},
									"address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address or address group of the real server.",
										Optional:    true,
										Computed:    true,
									},
									"health_check": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
										Optional:    true,
										Computed:    true,
									},
									"health_check_proto": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ping", "http", "tcp-connect"}, false),

										Description: "Protocol of the health check monitor to use when polling to determine server's connectivity status.",
										Optional:    true,
										Computed:    true,
									},
									"holddown_interval": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds).",
										Optional:    true,
										Computed:    true,
									},
									"http_host": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "HTTP server domain name in HTTP header.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Real server ID.",
										Optional:    true,
										Computed:    true,
									},
									"ip": {
										Type:         schema.TypeString,
										ValidateFunc: validation.IsIPv4Address,

										Description: "IP address of the real server.",
										Optional:    true,
										Computed:    true,
									},
									"mappedport": {
										Type: schema.TypeString,

										Description: "Port for communicating with the real server.",
										Optional:    true,
										Computed:    true,
									},
									"port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Port for communicating with the real server.",
										Optional:    true,
										Computed:    true,
									},
									"ssh_client_cert": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Set access-proxy SSH client certificate profile.",
										Optional:    true,
										Computed:    true,
									},
									"ssh_host_key": {
										Type:        schema.TypeList,
										Description: "One or more server host key.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Server host key name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"ssh_host_key_validation": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable SSH real server host key validation.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"active", "standby", "disable"}, false),

										Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"tcp-forwarding", "ssh"}, false),

										Description: "TCP forwarding server type.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"saml_redirect": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SAML redirection after successful authentication.",
							Optional:    true,
							Computed:    true,
						},
						"saml_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SAML service provider configuration for VIP authentication.",
							Optional:    true,
							Computed:    true,
						},
						"service": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"http", "https", "tcp-forwarding", "samlsp"}, false),

							Description: "Service.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

							Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_cipher_suites": {
							Type:        schema.TypeList,
							Description: "SSL/TLS cipher suites to offer to a server, ordered by priority.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"TLS-AES-128-GCM-SHA256", "TLS-AES-256-GCM-SHA384", "TLS-CHACHA20-POLY1305-SHA256", "TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-DHE-RSA-WITH-AES-128-CBC-SHA", "TLS-DHE-RSA-WITH-AES-256-CBC-SHA", "TLS-DHE-RSA-WITH-AES-128-CBC-SHA256", "TLS-DHE-RSA-WITH-AES-128-GCM-SHA256", "TLS-DHE-RSA-WITH-AES-256-CBC-SHA256", "TLS-DHE-RSA-WITH-AES-256-GCM-SHA384", "TLS-DHE-DSS-WITH-AES-128-CBC-SHA", "TLS-DHE-DSS-WITH-AES-256-CBC-SHA", "TLS-DHE-DSS-WITH-AES-128-CBC-SHA256", "TLS-DHE-DSS-WITH-AES-128-GCM-SHA256", "TLS-DHE-DSS-WITH-AES-256-CBC-SHA256", "TLS-DHE-DSS-WITH-AES-256-GCM-SHA384", "TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA", "TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256", "TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256", "TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA", "TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384", "TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA", "TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256", "TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256", "TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA", "TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384", "TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384", "TLS-RSA-WITH-AES-128-CBC-SHA", "TLS-RSA-WITH-AES-256-CBC-SHA", "TLS-RSA-WITH-AES-128-CBC-SHA256", "TLS-RSA-WITH-AES-128-GCM-SHA256", "TLS-RSA-WITH-AES-256-CBC-SHA256", "TLS-RSA-WITH-AES-256-GCM-SHA384", "TLS-RSA-WITH-CAMELLIA-128-CBC-SHA", "TLS-RSA-WITH-CAMELLIA-256-CBC-SHA", "TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256", "TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA", "TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA", "TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256", "TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256", "TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-RSA-WITH-SEED-CBC-SHA", "TLS-DHE-DSS-WITH-SEED-CBC-SHA", "TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256", "TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384", "TLS-RSA-WITH-SEED-CBC-SHA", "TLS-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256", "TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-RC4-128-SHA", "TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA", "TLS-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-RSA-WITH-RC4-128-MD5", "TLS-RSA-WITH-RC4-128-SHA", "TLS-DHE-RSA-WITH-DES-CBC-SHA", "TLS-DHE-DSS-WITH-DES-CBC-SHA", "TLS-RSA-WITH-DES-CBC-SHA"}, false),

										Description: "Cipher suite name.",
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Type: schema.TypeInt,

										Description: "SSL/TLS cipher suites priority.",
										Optional:    true,
										Computed:    true,
									},
									"versions": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffFakeListEqual,
										Description:      "SSL/TLS versions that the cipher suite can be used with.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
						"ssl_dh_bits": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"768", "1024", "1536", "2048", "3072", "4096"}, false),

							Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_max_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Highest SSL/TLS version acceptable from a server.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_min_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Lowest SSL/TLS version acceptable from a server.",
							Optional:    true,
							Computed:    true,
						},
						"url_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "URL pattern to match.",
							Optional:    true,
							Computed:    true,
						},
						"url_map_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"sub-string", "wildcard", "regex"}, false),

							Description: "Type of url-map.",
							Optional:    true,
							Computed:    true,
						},
						"virtual_host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Virtual host.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"api_gateway6": {
				Type:        schema.TypeList,
				Description: "Set IPv6 API Gateway.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_cookie_age": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 525600),

							Description: "Time in minutes that client web browsers should keep a cookie. Default is 60 minutes. 0 = no time limit.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_domain": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Domain that HTTP cookie persistence should apply to.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_domain_from_host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable use of HTTP cookie domain from host field in HTTP.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_generation": {
							Type: schema.TypeInt,

							Description: "Generation of HTTP cookie to be accepted. Changing invalidates all existing cookies.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_path": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Limit HTTP cookie persistence to the specified path.",
							Optional:    true,
							Computed:    true,
						},
						"http_cookie_share": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "same-ip"}, false),

							Description: "Control sharing of cookies across API Gateway. same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
							Optional:    true,
							Computed:    true,
						},
						"https_cookie_secure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable verification that inserted HTTPS cookies are secure.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "API Gateway ID.",
							Optional:    true,
							Computed:    true,
						},
						"ldb_method": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"static", "round-robin", "weighted", "first-alive", "http-host"}, false),

							Description: "Method used to distribute sessions to real servers.",
							Optional:    true,
							Computed:    true,
						},
						"persistence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "http-cookie"}, false),

							Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
							Optional:    true,
							Computed:    true,
						},
						"realservers": {
							Type:        schema.TypeList,
							Description: "Select the real servers that this Access Proxy will distribute traffic to.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr_type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ip", "fqdn"}, false),

										Description: "Type of address.",
										Optional:    true,
										Computed:    true,
									},
									"address": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address or address group of the real server.",
										Optional:    true,
										Computed:    true,
									},
									"health_check": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
										Optional:    true,
										Computed:    true,
									},
									"health_check_proto": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ping", "http", "tcp-connect"}, false),

										Description: "Protocol of the health check monitor to use when polling to determine server's connectivity status.",
										Optional:    true,
										Computed:    true,
									},
									"holddown_interval": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds).",
										Optional:    true,
										Computed:    true,
									},
									"http_host": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "HTTP server domain name in HTTP header.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "Real server ID.",
										Optional:    true,
										Computed:    true,
									},
									"ip": {
										Type:             schema.TypeString,
										ValidateFunc:     validation.IsIPv6Address,
										DiffSuppressFunc: suppressors.DiffIPEqual,
										Description:      "IPv6 address of the real server.",
										Optional:         true,
										Computed:         true,
									},
									"mappedport": {
										Type: schema.TypeString,

										Description: "Port for communicating with the real server.",
										Optional:    true,
										Computed:    true,
									},
									"port": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),

										Description: "Port for communicating with the real server.",
										Optional:    true,
										Computed:    true,
									},
									"ssh_client_cert": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Set access-proxy SSH client certificate profile.",
										Optional:    true,
										Computed:    true,
									},
									"ssh_host_key": {
										Type:        schema.TypeList,
										Description: "One or more server host key.",
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),

													Description: "Server host key name.",
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
									"ssh_host_key_validation": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

										Description: "Enable/disable SSH real server host key validation.",
										Optional:    true,
										Computed:    true,
									},
									"status": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"active", "standby", "disable"}, false),

										Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"tcp-forwarding", "ssh"}, false),

										Description: "TCP forwarding server type.",
										Optional:    true,
										Computed:    true,
									},
									"weight": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),

										Description: "Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"saml_redirect": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SAML redirection after successful authentication.",
							Optional:    true,
							Computed:    true,
						},
						"saml_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SAML service provider configuration for VIP authentication.",
							Optional:    true,
							Computed:    true,
						},
						"service": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"http", "https", "tcp-forwarding", "samlsp"}, false),

							Description: "Service.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

							Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_cipher_suites": {
							Type:        schema.TypeList,
							Description: "SSL/TLS cipher suites to offer to a server, ordered by priority.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cipher": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"TLS-AES-128-GCM-SHA256", "TLS-AES-256-GCM-SHA384", "TLS-CHACHA20-POLY1305-SHA256", "TLS-ECDHE-RSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-ECDHE-ECDSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-DHE-RSA-WITH-CHACHA20-POLY1305-SHA256", "TLS-DHE-RSA-WITH-AES-128-CBC-SHA", "TLS-DHE-RSA-WITH-AES-256-CBC-SHA", "TLS-DHE-RSA-WITH-AES-128-CBC-SHA256", "TLS-DHE-RSA-WITH-AES-128-GCM-SHA256", "TLS-DHE-RSA-WITH-AES-256-CBC-SHA256", "TLS-DHE-RSA-WITH-AES-256-GCM-SHA384", "TLS-DHE-DSS-WITH-AES-128-CBC-SHA", "TLS-DHE-DSS-WITH-AES-256-CBC-SHA", "TLS-DHE-DSS-WITH-AES-128-CBC-SHA256", "TLS-DHE-DSS-WITH-AES-128-GCM-SHA256", "TLS-DHE-DSS-WITH-AES-256-CBC-SHA256", "TLS-DHE-DSS-WITH-AES-256-GCM-SHA384", "TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA", "TLS-ECDHE-RSA-WITH-AES-128-CBC-SHA256", "TLS-ECDHE-RSA-WITH-AES-128-GCM-SHA256", "TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA", "TLS-ECDHE-RSA-WITH-AES-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-AES-256-GCM-SHA384", "TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA", "TLS-ECDHE-ECDSA-WITH-AES-128-CBC-SHA256", "TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256", "TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA", "TLS-ECDHE-ECDSA-WITH-AES-256-CBC-SHA384", "TLS-ECDHE-ECDSA-WITH-AES-256-GCM-SHA384", "TLS-RSA-WITH-AES-128-CBC-SHA", "TLS-RSA-WITH-AES-256-CBC-SHA", "TLS-RSA-WITH-AES-128-CBC-SHA256", "TLS-RSA-WITH-AES-128-GCM-SHA256", "TLS-RSA-WITH-AES-256-CBC-SHA256", "TLS-RSA-WITH-AES-256-GCM-SHA384", "TLS-RSA-WITH-CAMELLIA-128-CBC-SHA", "TLS-RSA-WITH-CAMELLIA-256-CBC-SHA", "TLS-RSA-WITH-CAMELLIA-128-CBC-SHA256", "TLS-RSA-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA", "TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA", "TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA", "TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA256", "TLS-DHE-DSS-WITH-CAMELLIA-128-CBC-SHA256", "TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-DSS-WITH-CAMELLIA-256-CBC-SHA256", "TLS-DHE-RSA-WITH-SEED-CBC-SHA", "TLS-DHE-DSS-WITH-SEED-CBC-SHA", "TLS-DHE-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-DHE-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-DHE-DSS-WITH-ARIA-128-CBC-SHA256", "TLS-DHE-DSS-WITH-ARIA-256-CBC-SHA384", "TLS-RSA-WITH-SEED-CBC-SHA", "TLS-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-ARIA-128-CBC-SHA256", "TLS-ECDHE-RSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-ECDSA-WITH-ARIA-128-CBC-SHA256", "TLS-ECDHE-ECDSA-WITH-ARIA-256-CBC-SHA384", "TLS-ECDHE-RSA-WITH-RC4-128-SHA", "TLS-ECDHE-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-DHE-DSS-WITH-3DES-EDE-CBC-SHA", "TLS-RSA-WITH-3DES-EDE-CBC-SHA", "TLS-RSA-WITH-RC4-128-MD5", "TLS-RSA-WITH-RC4-128-SHA", "TLS-DHE-RSA-WITH-DES-CBC-SHA", "TLS-DHE-DSS-WITH-DES-CBC-SHA", "TLS-RSA-WITH-DES-CBC-SHA"}, false),

										Description: "Cipher suite name.",
										Optional:    true,
										Computed:    true,
									},
									"priority": {
										Type: schema.TypeInt,

										Description: "SSL/TLS cipher suites priority.",
										Optional:    true,
										Computed:    true,
									},
									"versions": {
										Type: schema.TypeString,

										DiffSuppressFunc: suppressors.DiffFakeListEqual,
										Description:      "SSL/TLS versions that the cipher suite can be used with.",
										Optional:         true,
										Computed:         true,
									},
								},
							},
						},
						"ssl_dh_bits": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"768", "1024", "1536", "2048", "3072", "4096"}, false),

							Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_max_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Highest SSL/TLS version acceptable from a server.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_min_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Lowest SSL/TLS version acceptable from a server.",
							Optional:    true,
							Computed:    true,
						},
						"url_map": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "URL pattern to match.",
							Optional:    true,
							Computed:    true,
						},
						"url_map_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"sub-string", "wildcard", "regex"}, false),

							Description: "Type of url-map.",
							Optional:    true,
							Computed:    true,
						},
						"virtual_host": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Virtual host.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"client_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable to request client certificate.",
				Optional:    true,
				Computed:    true,
			},
			"decrypted_traffic_mirror": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Decrypted traffic mirror.",
				Optional:    true,
				Computed:    true,
			},
			"empty_cert_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "block"}, false),

				Description: "Action of an empty client certificate.",
				Optional:    true,
				Computed:    true,
			},
			"log_blocked_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging of blocked traffic.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Access Proxy name.",
				ForceNew:    true,
				Required:    true,
			},
			"vip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Virtual IP name.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAccessProxy6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallAccessProxy6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallAccessProxy6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAccessProxy6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAccessProxy6")
	}

	return resourceFirewallAccessProxy6Read(ctx, d, meta)
}

func resourceFirewallAccessProxy6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectFirewallAccessProxy6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAccessProxy6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAccessProxy6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAccessProxy6")
	}

	return resourceFirewallAccessProxy6Read(ctx, d, meta)
}

func resourceFirewallAccessProxy6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	err := c.Cmdb.DeleteFirewallAccessProxy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAccessProxy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	o, err := c.Cmdb.ReadFirewallAccessProxy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAccessProxy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenFirewallAccessProxy6ApiGateway(v *[]models.FirewallAccessProxy6ApiGateway, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.HttpCookieAge; tmp != nil {
				v["http_cookie_age"] = *tmp
			}

			if tmp := cfg.HttpCookieDomain; tmp != nil {
				v["http_cookie_domain"] = *tmp
			}

			if tmp := cfg.HttpCookieDomainFromHost; tmp != nil {
				v["http_cookie_domain_from_host"] = *tmp
			}

			if tmp := cfg.HttpCookieGeneration; tmp != nil {
				v["http_cookie_generation"] = *tmp
			}

			if tmp := cfg.HttpCookiePath; tmp != nil {
				v["http_cookie_path"] = *tmp
			}

			if tmp := cfg.HttpCookieShare; tmp != nil {
				v["http_cookie_share"] = *tmp
			}

			if tmp := cfg.HttpsCookieSecure; tmp != nil {
				v["https_cookie_secure"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.LdbMethod; tmp != nil {
				v["ldb_method"] = *tmp
			}

			if tmp := cfg.Persistence; tmp != nil {
				v["persistence"] = *tmp
			}

			if tmp := cfg.Realservers; tmp != nil {
				v["realservers"] = flattenFirewallAccessProxy6ApiGatewayRealservers(tmp, sort)
			}

			if tmp := cfg.SamlRedirect; tmp != nil {
				v["saml_redirect"] = *tmp
			}

			if tmp := cfg.SamlServer; tmp != nil {
				v["saml_server"] = *tmp
			}

			if tmp := cfg.Service; tmp != nil {
				v["service"] = *tmp
			}

			if tmp := cfg.SslAlgorithm; tmp != nil {
				v["ssl_algorithm"] = *tmp
			}

			if tmp := cfg.SslCipherSuites; tmp != nil {
				v["ssl_cipher_suites"] = flattenFirewallAccessProxy6ApiGatewaySslCipherSuites(tmp, sort)
			}

			if tmp := cfg.SslDhBits; tmp != nil {
				v["ssl_dh_bits"] = *tmp
			}

			if tmp := cfg.SslMaxVersion; tmp != nil {
				v["ssl_max_version"] = *tmp
			}

			if tmp := cfg.SslMinVersion; tmp != nil {
				v["ssl_min_version"] = *tmp
			}

			if tmp := cfg.UrlMap; tmp != nil {
				v["url_map"] = *tmp
			}

			if tmp := cfg.UrlMapType; tmp != nil {
				v["url_map_type"] = *tmp
			}

			if tmp := cfg.VirtualHost; tmp != nil {
				v["virtual_host"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGatewayRealservers(v *[]models.FirewallAccessProxy6ApiGatewayRealservers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AddrType; tmp != nil {
				v["addr_type"] = *tmp
			}

			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = *tmp
			}

			if tmp := cfg.HealthCheckProto; tmp != nil {
				v["health_check_proto"] = *tmp
			}

			if tmp := cfg.HolddownInterval; tmp != nil {
				v["holddown_interval"] = *tmp
			}

			if tmp := cfg.HttpHost; tmp != nil {
				v["http_host"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Mappedport; tmp != nil {
				v["mappedport"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.SshClientCert; tmp != nil {
				v["ssh_client_cert"] = *tmp
			}

			if tmp := cfg.SshHostKey; tmp != nil {
				v["ssh_host_key"] = flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey(tmp, sort)
			}

			if tmp := cfg.SshHostKeyValidation; tmp != nil {
				v["ssh_host_key_validation"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey(v *[]models.FirewallAccessProxy6ApiGatewayRealserversSshHostKey, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuites(v *[]models.FirewallAccessProxy6ApiGatewaySslCipherSuites, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Cipher; tmp != nil {
				v["cipher"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Versions; tmp != nil {
				v["versions"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "priority")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGateway6(v *[]models.FirewallAccessProxy6ApiGateway6, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.HttpCookieAge; tmp != nil {
				v["http_cookie_age"] = *tmp
			}

			if tmp := cfg.HttpCookieDomain; tmp != nil {
				v["http_cookie_domain"] = *tmp
			}

			if tmp := cfg.HttpCookieDomainFromHost; tmp != nil {
				v["http_cookie_domain_from_host"] = *tmp
			}

			if tmp := cfg.HttpCookieGeneration; tmp != nil {
				v["http_cookie_generation"] = *tmp
			}

			if tmp := cfg.HttpCookiePath; tmp != nil {
				v["http_cookie_path"] = *tmp
			}

			if tmp := cfg.HttpCookieShare; tmp != nil {
				v["http_cookie_share"] = *tmp
			}

			if tmp := cfg.HttpsCookieSecure; tmp != nil {
				v["https_cookie_secure"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.LdbMethod; tmp != nil {
				v["ldb_method"] = *tmp
			}

			if tmp := cfg.Persistence; tmp != nil {
				v["persistence"] = *tmp
			}

			if tmp := cfg.Realservers; tmp != nil {
				v["realservers"] = flattenFirewallAccessProxy6ApiGateway6Realservers(tmp, sort)
			}

			if tmp := cfg.SamlRedirect; tmp != nil {
				v["saml_redirect"] = *tmp
			}

			if tmp := cfg.SamlServer; tmp != nil {
				v["saml_server"] = *tmp
			}

			if tmp := cfg.Service; tmp != nil {
				v["service"] = *tmp
			}

			if tmp := cfg.SslAlgorithm; tmp != nil {
				v["ssl_algorithm"] = *tmp
			}

			if tmp := cfg.SslCipherSuites; tmp != nil {
				v["ssl_cipher_suites"] = flattenFirewallAccessProxy6ApiGateway6SslCipherSuites(tmp, sort)
			}

			if tmp := cfg.SslDhBits; tmp != nil {
				v["ssl_dh_bits"] = *tmp
			}

			if tmp := cfg.SslMaxVersion; tmp != nil {
				v["ssl_max_version"] = *tmp
			}

			if tmp := cfg.SslMinVersion; tmp != nil {
				v["ssl_min_version"] = *tmp
			}

			if tmp := cfg.UrlMap; tmp != nil {
				v["url_map"] = *tmp
			}

			if tmp := cfg.UrlMapType; tmp != nil {
				v["url_map_type"] = *tmp
			}

			if tmp := cfg.VirtualHost; tmp != nil {
				v["virtual_host"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGateway6Realservers(v *[]models.FirewallAccessProxy6ApiGateway6Realservers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AddrType; tmp != nil {
				v["addr_type"] = *tmp
			}

			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.HealthCheck; tmp != nil {
				v["health_check"] = *tmp
			}

			if tmp := cfg.HealthCheckProto; tmp != nil {
				v["health_check_proto"] = *tmp
			}

			if tmp := cfg.HolddownInterval; tmp != nil {
				v["holddown_interval"] = *tmp
			}

			if tmp := cfg.HttpHost; tmp != nil {
				v["http_host"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Mappedport; tmp != nil {
				v["mappedport"] = *tmp
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.SshClientCert; tmp != nil {
				v["ssh_client_cert"] = *tmp
			}

			if tmp := cfg.SshHostKey; tmp != nil {
				v["ssh_host_key"] = flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey(tmp, sort)
			}

			if tmp := cfg.SshHostKeyValidation; tmp != nil {
				v["ssh_host_key_validation"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Weight; tmp != nil {
				v["weight"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey(v *[]models.FirewallAccessProxy6ApiGateway6RealserversSshHostKey, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuites(v *[]models.FirewallAccessProxy6ApiGateway6SslCipherSuites, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Cipher; tmp != nil {
				v["cipher"] = *tmp
			}

			if tmp := cfg.Priority; tmp != nil {
				v["priority"] = *tmp
			}

			if tmp := cfg.Versions; tmp != nil {
				v["versions"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "priority")
	}

	return flat
}

func refreshObjectFirewallAccessProxy6(d *schema.ResourceData, o *models.FirewallAccessProxy6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ApiGateway != nil {
		if err = d.Set("api_gateway", flattenFirewallAccessProxy6ApiGateway(o.ApiGateway, sort)); err != nil {
			return diag.Errorf("error reading api_gateway: %v", err)
		}
	}

	if o.ApiGateway6 != nil {
		if err = d.Set("api_gateway6", flattenFirewallAccessProxy6ApiGateway6(o.ApiGateway6, sort)); err != nil {
			return diag.Errorf("error reading api_gateway6: %v", err)
		}
	}

	if o.ClientCert != nil {
		v := *o.ClientCert

		if err = d.Set("client_cert", v); err != nil {
			return diag.Errorf("error reading client_cert: %v", err)
		}
	}

	if o.DecryptedTrafficMirror != nil {
		v := *o.DecryptedTrafficMirror

		if err = d.Set("decrypted_traffic_mirror", v); err != nil {
			return diag.Errorf("error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if o.EmptyCertAction != nil {
		v := *o.EmptyCertAction

		if err = d.Set("empty_cert_action", v); err != nil {
			return diag.Errorf("error reading empty_cert_action: %v", err)
		}
	}

	if o.LogBlockedTraffic != nil {
		v := *o.LogBlockedTraffic

		if err = d.Set("log_blocked_traffic", v); err != nil {
			return diag.Errorf("error reading log_blocked_traffic: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Vip != nil {
		v := *o.Vip

		if err = d.Set("vip", v); err != nil {
			return diag.Errorf("error reading vip: %v", err)
		}
	}

	return nil
}

func expandFirewallAccessProxy6ApiGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGateway, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGateway

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGateway{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.http_cookie_age", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HttpCookieAge = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookieDomain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_domain_from_host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookieDomainFromHost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_generation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HttpCookieGeneration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookiePath = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_share", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookieShare = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.https_cookie_secure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpsCookieSecure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ldb_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LdbMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.persistence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Persistence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.realservers", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAccessProxy6ApiGatewayRealservers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAccessProxy6ApiGatewayRealservers
			// 	}
			tmp.Realservers = v2

		}

		pre_append = fmt.Sprintf("%s.%d.saml_redirect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamlRedirect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.saml_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamlServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Service = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_algorithm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslAlgorithm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_cipher_suites", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAccessProxy6ApiGatewaySslCipherSuites(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAccessProxy6ApiGatewaySslCipherSuites
			// 	}
			tmp.SslCipherSuites = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ssl_dh_bits", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslDhBits = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_max_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslMaxVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_min_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslMinVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UrlMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url_map_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UrlMapType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virtual_host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VirtualHost = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGatewayRealservers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGatewayRealservers

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGatewayRealservers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.health_check_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheckProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.holddown_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HolddownInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpHost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mappedport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mappedport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssh_client_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SshClientCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssh_host_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAccessProxy6ApiGatewayRealserversSshHostKey
			// 	}
			tmp.SshHostKey = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ssh_host_key_validation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SshHostKeyValidation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Weight = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGatewayRealserversSshHostKey, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGatewayRealserversSshHostKey

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGatewayRealserversSshHostKey{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGatewaySslCipherSuites, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGatewaySslCipherSuites

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGatewaySslCipherSuites{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.versions", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Versions = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGateway6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGateway6

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGateway6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.http_cookie_age", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HttpCookieAge = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_domain", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookieDomain = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_domain_from_host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookieDomainFromHost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_generation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.HttpCookieGeneration = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_path", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookiePath = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_cookie_share", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpCookieShare = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.https_cookie_secure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpsCookieSecure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ldb_method", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LdbMethod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.persistence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Persistence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.realservers", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAccessProxy6ApiGateway6Realservers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAccessProxy6ApiGateway6Realservers
			// 	}
			tmp.Realservers = v2

		}

		pre_append = fmt.Sprintf("%s.%d.saml_redirect", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamlRedirect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.saml_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SamlServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.service", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Service = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_algorithm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslAlgorithm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_cipher_suites", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAccessProxy6ApiGateway6SslCipherSuites(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAccessProxy6ApiGateway6SslCipherSuites
			// 	}
			tmp.SslCipherSuites = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ssl_dh_bits", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslDhBits = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_max_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslMaxVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_min_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslMinVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url_map", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UrlMap = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.url_map_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UrlMapType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.virtual_host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VirtualHost = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGateway6Realservers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGateway6Realservers

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGateway6Realservers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.health_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.health_check_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HealthCheckProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.holddown_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HolddownInterval = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.http_host", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpHost = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.mappedport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Mappedport = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Port = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssh_client_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SshClientCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssh_host_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAccessProxy6ApiGateway6RealserversSshHostKey
			// 	}
			tmp.SshHostKey = v2

		}

		pre_append = fmt.Sprintf("%s.%d.ssh_host_key_validation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SshHostKeyValidation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Weight = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGateway6RealserversSshHostKey, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGateway6RealserversSshHostKey

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGateway6RealserversSshHostKey{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAccessProxy6ApiGateway6SslCipherSuites, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAccessProxy6ApiGateway6SslCipherSuites

	for i := range l {
		tmp := models.FirewallAccessProxy6ApiGateway6SslCipherSuites{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Priority = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.versions", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Versions = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallAccessProxy6(d *schema.ResourceData, sv string) (*models.FirewallAccessProxy6, diag.Diagnostics) {
	obj := models.FirewallAccessProxy6{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("api_gateway"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("api_gateway", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAccessProxy6ApiGateway(d, v, "api_gateway", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ApiGateway = t
		}
	} else if d.HasChange("api_gateway") {
		old, new := d.GetChange("api_gateway")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ApiGateway = &[]models.FirewallAccessProxy6ApiGateway{}
		}
	}
	if v, ok := d.GetOk("api_gateway6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("api_gateway6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAccessProxy6ApiGateway6(d, v, "api_gateway6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ApiGateway6 = t
		}
	} else if d.HasChange("api_gateway6") {
		old, new := d.GetChange("api_gateway6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ApiGateway6 = &[]models.FirewallAccessProxy6ApiGateway6{}
		}
	}
	if v1, ok := d.GetOk("client_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("client_cert", sv)
				diags = append(diags, e)
			}
			obj.ClientCert = &v2
		}
	}
	if v1, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("decrypted_traffic_mirror", sv)
				diags = append(diags, e)
			}
			obj.DecryptedTrafficMirror = &v2
		}
	}
	if v1, ok := d.GetOk("empty_cert_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("empty_cert_action", sv)
				diags = append(diags, e)
			}
			obj.EmptyCertAction = &v2
		}
	}
	if v1, ok := d.GetOk("log_blocked_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("log_blocked_traffic", sv)
				diags = append(diags, e)
			}
			obj.LogBlockedTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("vip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vip", sv)
				diags = append(diags, e)
			}
			obj.Vip = &v2
		}
	}
	return &obj, diags
}
