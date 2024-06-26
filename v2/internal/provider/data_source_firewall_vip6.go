// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual IP for IPv6.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallVip6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual IP for IPv6.",

		ReadContext: dataSourceFirewallVip6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"add_nat64_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding NAT64 route.",
				Computed:    true,
			},
			"arp_reply": {
				Type:        schema.TypeString,
				Description: "Enable to respond to ARP requests for this virtual IP address. Enabled by default.",
				Computed:    true,
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Color of icon on the GUI.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"embedded_ipv4_address": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of the lower 32 bits of the external IPv6 address as mapped IPv4 address.",
				Computed:    true,
			},
			"extip": {
				Type:        schema.TypeString,
				Description: "IPv6 address or address range on the external interface that you want to map to an address or address range on the destination network.",
				Computed:    true,
			},
			"extport": {
				Type:        schema.TypeString,
				Description: "Incoming port number range that you want to map to a port number range on the destination network.",
				Computed:    true,
			},
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
				Description: "Control sharing of cookies across virtual servers. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
				Computed:    true,
			},
			"http_ip_header": {
				Type:        schema.TypeString,
				Description: "For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header.",
				Computed:    true,
			},
			"http_ip_header_name": {
				Type:        schema.TypeString,
				Description: "For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.",
				Computed:    true,
			},
			"http_multiplex": {
				Type:        schema.TypeString,
				Description: "Enable/disable HTTP multiplexing.",
				Computed:    true,
			},
			"http_redirect": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirection of HTTP to HTTPS.",
				Computed:    true,
			},
			"https_cookie_secure": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification that inserted HTTPS cookies are secure.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Custom defined ID.",
				Computed:    true,
			},
			"ipv4_mappedip": {
				Type:        schema.TypeString,
				Description: "Range of mapped IP addresses. Specify the start IP address followed by a space and the end IP address.",
				Computed:    true,
			},
			"ipv4_mappedport": {
				Type:        schema.TypeString,
				Description: "IPv4 port number range on the destination network to which the external port number range is mapped.",
				Computed:    true,
			},
			"ldb_method": {
				Type:        schema.TypeString,
				Description: "Method used to distribute sessions to real servers.",
				Computed:    true,
			},
			"mappedip": {
				Type:        schema.TypeString,
				Description: "Mapped IPv6 address range in the format startIP-endIP.",
				Computed:    true,
			},
			"mappedport": {
				Type:        schema.TypeString,
				Description: "Port number range on the destination network to which the external port number range is mapped.",
				Computed:    true,
			},
			"max_embryonic_connections": {
				Type:        schema.TypeInt,
				Description: "Maximum number of incomplete connections.",
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeList,
				Description: "Name of the health check monitor to use when polling to determine a virtual server's connectivity status.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Health monitor name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Virtual ip6 name.",
				Required:    true,
			},
			"nat_source_vip": {
				Type:        schema.TypeString,
				Description: "Enable to perform SNAT on traffic from mappedip to the extip for all egress interfaces.",
				Computed:    true,
			},
			"nat64": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNAT64.",
				Computed:    true,
			},
			"nat66": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNAT66.",
				Computed:    true,
			},
			"ndp_reply": {
				Type:        schema.TypeString,
				Description: "Enable/disable this FortiGate unit's ability to respond to NDP requests for this virtual IP address (default = enable).",
				Computed:    true,
			},
			"outlook_web_access": {
				Type:        schema.TypeString,
				Description: "Enable to add the Front-End-Https header for Microsoft Outlook Web Access.",
				Computed:    true,
			},
			"persistence": {
				Type:        schema.TypeString,
				Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
				Computed:    true,
			},
			"portforward": {
				Type:        schema.TypeString,
				Description: "Enable port forwarding.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol to use when forwarding packets.",
				Computed:    true,
			},
			"realservers": {
				Type:        schema.TypeList,
				Description: "Select the real servers that this server load balancing VIP will distribute traffic to.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ip": {
							Type:        schema.TypeString,
							Description: "Only clients in this IP range can connect to this real server.",
							Computed:    true,
						},
						"healthcheck": {
							Type:        schema.TypeString,
							Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
							Computed:    true,
						},
						"holddown_interval": {
							Type:        schema.TypeInt,
							Description: "Time in seconds that the health check monitor continues to monitor an unresponsive server that should be active.",
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
						"max_connections": {
							Type:        schema.TypeInt,
							Description: "Max number of active connections that can directed to the real server. When reached, sessions are sent to other real servers.",
							Computed:    true,
						},
						"monitor": {
							Type:        schema.TypeList,
							Description: "Name of the health check monitor to use when polling to determine a virtual server's connectivity status.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Health monitor name.",
										Computed:    true,
									},
								},
							},
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Port for communicating with the real server. Required if port forwarding is enabled.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent.",
							Computed:    true,
						},
						"translate_host": {
							Type:        schema.TypeString,
							Description: "Enable/disable translation of hostname/IP from virtual server to real server.",
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
			"server_type": {
				Type:        schema.TypeString,
				Description: "Protocol to be load balanced by the virtual server (also called the server load balance virtual IP).",
				Computed:    true,
			},
			"src_filter": {
				Type:        schema.TypeList,
				Description: "Source IP6 filter (x:x:x:x:x:x:x:x/x). Separate addresses with spaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type:        schema.TypeString,
							Description: "Source-filter range.",
							Computed:    true,
						},
					},
				},
			},
			"ssl_accept_ffdhe_groups": {
				Type:        schema.TypeString,
				Description: "Enable/disable FFDHE cipher suite for SSL key exchange.",
				Computed:    true,
			},
			"ssl_algorithm": {
				Type:        schema.TypeString,
				Description: "Permitted encryption algorithms for SSL sessions according to encryption strength.",
				Computed:    true,
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Description: "The name of the certificate to use for SSL handshake.",
				Computed:    true,
			},
			"ssl_cipher_suites": {
				Type:        schema.TypeList,
				Description: "SSL/TLS cipher suites acceptable from a client, ordered by priority.",
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
			"ssl_client_fallback": {
				Type:        schema.TypeString,
				Description: "Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507).",
				Computed:    true,
			},
			"ssl_client_rekey_count": {
				Type:        schema.TypeInt,
				Description: "Maximum length of data in MB before triggering a client rekey (0 = disable).",
				Computed:    true,
			},
			"ssl_client_renegotiation": {
				Type:        schema.TypeString,
				Description: "Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746.",
				Computed:    true,
			},
			"ssl_client_session_state_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of client to FortiGate SSL session states to keep.",
				Computed:    true,
			},
			"ssl_client_session_state_timeout": {
				Type:        schema.TypeInt,
				Description: "Number of minutes to keep client to FortiGate SSL session state.",
				Computed:    true,
			},
			"ssl_client_session_state_type": {
				Type:        schema.TypeString,
				Description: "How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate.",
				Computed:    true,
			},
			"ssl_dh_bits": {
				Type:        schema.TypeString,
				Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
				Computed:    true,
			},
			"ssl_hpkp": {
				Type:        schema.TypeString,
				Description: "Enable/disable including HPKP header in response.",
				Computed:    true,
			},
			"ssl_hpkp_age": {
				Type:        schema.TypeInt,
				Description: "Number of minutes the web browser should keep HPKP.",
				Computed:    true,
			},
			"ssl_hpkp_backup": {
				Type:        schema.TypeString,
				Description: "Certificate to generate backup HPKP pin from.",
				Computed:    true,
			},
			"ssl_hpkp_include_subdomains": {
				Type:        schema.TypeString,
				Description: "Indicate that HPKP header applies to all subdomains.",
				Computed:    true,
			},
			"ssl_hpkp_primary": {
				Type:        schema.TypeString,
				Description: "Certificate to generate primary HPKP pin from.",
				Computed:    true,
			},
			"ssl_hpkp_report_uri": {
				Type:        schema.TypeString,
				Description: "URL to report HPKP violations to.",
				Computed:    true,
			},
			"ssl_hsts": {
				Type:        schema.TypeString,
				Description: "Enable/disable including HSTS header in response.",
				Computed:    true,
			},
			"ssl_hsts_age": {
				Type:        schema.TypeInt,
				Description: "Number of seconds the client should honor the HSTS setting.",
				Computed:    true,
			},
			"ssl_hsts_include_subdomains": {
				Type:        schema.TypeString,
				Description: "Indicate that HSTS header applies to all subdomains.",
				Computed:    true,
			},
			"ssl_http_location_conversion": {
				Type:        schema.TypeString,
				Description: "Enable to replace HTTP with HTTPS in the reply's Location HTTP header field.",
				Computed:    true,
			},
			"ssl_http_match_host": {
				Type:        schema.TypeString,
				Description: "Enable/disable HTTP host matching for location conversion.",
				Computed:    true,
			},
			"ssl_max_version": {
				Type:        schema.TypeString,
				Description: "Highest SSL/TLS version acceptable from a client.",
				Computed:    true,
			},
			"ssl_min_version": {
				Type:        schema.TypeString,
				Description: "Lowest SSL/TLS version acceptable from a client.",
				Computed:    true,
			},
			"ssl_mode": {
				Type:        schema.TypeString,
				Description: "Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full).",
				Computed:    true,
			},
			"ssl_pfs": {
				Type:        schema.TypeString,
				Description: "Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions.",
				Computed:    true,
			},
			"ssl_send_empty_frags": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems.",
				Computed:    true,
			},
			"ssl_server_algorithm": {
				Type:        schema.TypeString,
				Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
				Computed:    true,
			},
			"ssl_server_cipher_suites": {
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
			"ssl_server_max_version": {
				Type:        schema.TypeString,
				Description: "Highest SSL/TLS version acceptable from a server. Use the client setting by default.",
				Computed:    true,
			},
			"ssl_server_min_version": {
				Type:        schema.TypeString,
				Description: "Lowest SSL/TLS version acceptable from a server. Use the client setting by default.",
				Computed:    true,
			},
			"ssl_server_renegotiation": {
				Type:        schema.TypeString,
				Description: "Enable/disable secure renegotiation to comply with RFC 5746.",
				Computed:    true,
			},
			"ssl_server_session_state_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of FortiGate to Server SSL session states to keep.",
				Computed:    true,
			},
			"ssl_server_session_state_timeout": {
				Type:        schema.TypeInt,
				Description: "Number of minutes to keep FortiGate to Server SSL session state.",
				Computed:    true,
			},
			"ssl_server_session_state_type": {
				Type:        schema.TypeString,
				Description: "How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Configure a static NAT server load balance VIP or access proxy.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"weblogic_server": {
				Type:        schema.TypeString,
				Description: "Enable to add an HTTP header to indicate SSL offloading for a WebLogic server.",
				Computed:    true,
			},
			"websphere_server": {
				Type:        schema.TypeString,
				Description: "Enable to add an HTTP header to indicate SSL offloading for a WebSphere server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallVip6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallVip6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallVip6 dataSource: %v", err)
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

	diags := refreshObjectFirewallVip6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
