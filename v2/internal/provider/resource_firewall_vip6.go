// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual IP for IPv6.

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

func resourceFirewallVip6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual IP for IPv6.",

		CreateContext: resourceFirewallVip6Create,
		ReadContext:   resourceFirewallVip6Read,
		UpdateContext: resourceFirewallVip6Update,
		DeleteContext: resourceFirewallVip6Delete,

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
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"add_nat64_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable adding NAT64 route.",
				Optional:    true,
				Computed:    true,
			},
			"arp_reply": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to respond to ARP requests for this virtual IP address. Enabled by default.",
				Optional:    true,
				Computed:    true,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"embedded_ipv4_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of the lower 32 bits of the external IPv6 address as mapped IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"extip": {
				Type: schema.TypeString,

				Description: "IPv6 address or address range on the external interface that you want to map to an address or address range on the destination network.",
				Optional:    true,
				Computed:    true,
			},
			"extport": {
				Type: schema.TypeString,

				Description: "Incoming port number range that you want to map to a port number range on the destination network.",
				Optional:    true,
				Computed:    true,
			},
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

				Description: "Control sharing of cookies across virtual servers. Use of same-ip means a cookie from one virtual server can be used by another. Disable stops cookie sharing.",
				Optional:    true,
				Computed:    true,
			},
			"http_ip_header": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "For HTTP multiplexing, enable to add the original client IP address in the XForwarded-For HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"http_ip_header_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "For HTTP multiplexing, enter a custom HTTPS header name. The original client IP address is added to this header. If empty, X-Forwarded-For is used.",
				Optional:    true,
				Computed:    true,
			},
			"http_multiplex": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HTTP multiplexing.",
				Optional:    true,
				Computed:    true,
			},
			"http_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redirection of HTTP to HTTPS.",
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
			"fosid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Custom defined ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"ipv4_mappedip": {
				Type: schema.TypeString,

				Description: "Range of mapped IP addresses. Specify the start IP address followed by a space and the end IP address.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_mappedport": {
				Type: schema.TypeString,

				Description: "IPv4 port number range on the destination network to which the external port number range is mapped.",
				Optional:    true,
				Computed:    true,
			},
			"ldb_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static", "round-robin", "weighted", "least-session", "least-rtt", "first-alive", "http-host"}, false),

				Description: "Method used to distribute sessions to real servers.",
				Optional:    true,
				Computed:    true,
			},
			"mappedip": {
				Type: schema.TypeString,

				Description: "Mapped IPv6 address range in the format startIP-endIP.",
				Optional:    true,
				Computed:    true,
			},
			"mappedport": {
				Type: schema.TypeString,

				Description: "Port number range on the destination network to which the external port number range is mapped.",
				Optional:    true,
				Computed:    true,
			},
			"max_embryonic_connections": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),

				Description: "Maximum number of incomplete connections.",
				Optional:    true,
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeList,
				Description: "Name of the health check monitor to use when polling to determine a virtual server's connectivity status.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Health monitor name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Virtual ip6 name.",
				ForceNew:    true,
				Required:    true,
			},
			"nat_source_vip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to perform SNAT on traffic from mappedip to the extip for all egress interfaces.",
				Optional:    true,
				Computed:    true,
			},
			"nat64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable DNAT64.",
				Optional:    true,
				Computed:    true,
			},
			"nat66": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable DNAT66.",
				Optional:    true,
				Computed:    true,
			},
			"outlook_web_access": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to add the Front-End-Https header for Microsoft Outlook Web Access.",
				Optional:    true,
				Computed:    true,
			},
			"persistence": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "http-cookie", "ssl-session-id"}, false),

				Description: "Configure how to make sure that clients connect to the same server every time they make a request that is part of the same session.",
				Optional:    true,
				Computed:    true,
			},
			"portforward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable port forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp", "sctp"}, false),

				Description: "Protocol to use when forwarding packets.",
				Optional:    true,
				Computed:    true,
			},
			"realservers": {
				Type:        schema.TypeList,
				Description: "Select the real servers that this server load balancing VIP will distribute traffic to.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ip": {
							Type: schema.TypeString,

							Description: "Only clients in this IP range can connect to this real server.",
							Optional:    true,
							Computed:    true,
						},
						"healthcheck": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "vip"}, false),

							Description: "Enable to check the responsiveness of the real server before forwarding traffic.",
							Optional:    true,
							Computed:    true,
						},
						"holddown_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(30, 65535),

							Description: "Time in seconds that the health check monitor continues to monitor an unresponsive server that should be active.",
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
							Type: schema.TypeString,

							Description: "IP address of the real server.",
							Optional:    true,
							Computed:    true,
						},
						"max_connections": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2147483647),

							Description: "Max number of active connections that can directed to the real server. When reached, sessions are sent to other real servers.",
							Optional:    true,
							Computed:    true,
						},
						"monitor": {
							Type:        schema.TypeList,
							Description: "Name of the health check monitor to use when polling to determine a virtual server's connectivity status.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Health monitor name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Port for communicating with the real server. Required if port forwarding is enabled.",
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
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"http", "https", "imaps", "pop3s", "smtps", "ssl", "tcp", "udp", "ip"}, false),

				Description: "Protocol to be load balanced by the virtual server (also called the server load balance virtual IP).",
				Optional:    true,
				Computed:    true,
			},
			"src_filter": {
				Type:        schema.TypeList,
				Description: "Source IP6 filter (x:x:x:x:x:x:x:x/x). Separate addresses with spaces.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Source-filter range.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl_accept_ffdhe_groups": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FFDHE cipher suite for SSL key exchange.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low", "custom"}, false),

				Description: "Permitted encryption algorithms for SSL sessions according to encryption strength.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "The name of the certificate to use for SSL handshake.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_cipher_suites": {
				Type:        schema.TypeList,
				Description: "SSL/TLS cipher suites acceptable from a client, ordered by priority.",
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
			"ssl_client_fallback": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable support for preventing Downgrade Attacks on client connections (RFC 7507).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_client_rekey_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(200, 1048576),

				Description: "Maximum length of data in MB before triggering a client rekey (0 = disable).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_client_renegotiation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "secure"}, false),

				Description: "Allow, deny, or require secure renegotiation of client sessions to comply with RFC 5746.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_client_session_state_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10000),

				Description: "Maximum number of client to FortiGate SSL session states to keep.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_client_session_state_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 14400),

				Description: "Number of minutes to keep client to FortiGate SSL session state.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_client_session_state_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "time", "count", "both"}, false),

				Description: "How to expire SSL sessions for the segment of the SSL connection between the client and the FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_dh_bits": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"768", "1024", "1536", "2048", "3072", "4096"}, false),

				Description: "Number of bits to use in the Diffie-Hellman exchange for RSA encryption of SSL sessions.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hpkp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable", "report-only"}, false),

				Description: "Enable/disable including HPKP header in response.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hpkp_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 157680000),

				Description: "Number of minutes the web browser should keep HPKP.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hpkp_backup": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Certificate to generate backup HPKP pin from.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hpkp_include_subdomains": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Indicate that HPKP header applies to all subdomains.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hpkp_primary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Certificate to generate primary HPKP pin from.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hpkp_report_uri": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "URL to report HPKP violations to.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hsts": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable including HSTS header in response.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hsts_age": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 157680000),

				Description: "Number of seconds the client should honor the HSTS setting.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_hsts_include_subdomains": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Indicate that HSTS header applies to all subdomains.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_http_location_conversion": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to replace HTTP with HTTPS in the reply's Location HTTP header field.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_http_match_host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable HTTP host matching for location conversion.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_max_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

				Description: "Highest SSL/TLS version acceptable from a client.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

				Description: "Lowest SSL/TLS version acceptable from a client.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"half", "full"}, false),

				Description: "Apply SSL offloading between the client and the FortiGate (half) or from the client to the FortiGate and from the FortiGate to the server (full).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_pfs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"require", "deny", "allow"}, false),

				Description: "Select the cipher suites that can be used for SSL perfect forward secrecy (PFS). Applies to both client and server sessions.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_send_empty_frags": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending empty fragments to avoid CBC IV attacks (SSL 3.0 & TLS 1.0 only). May need to be disabled for compatibility with older systems.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low", "custom", "client"}, false),

				Description: "Permitted encryption algorithms for the server side of SSL full mode sessions according to encryption strength.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_cipher_suites": {
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
			"ssl_server_max_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3", "client"}, false),

				Description: "Highest SSL/TLS version acceptable from a server. Use the client setting by default.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_min_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3", "client"}, false),

				Description: "Lowest SSL/TLS version acceptable from a server. Use the client setting by default.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_session_state_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10000),

				Description: "Maximum number of FortiGate to Server SSL session states to keep.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_session_state_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 14400),

				Description: "Number of minutes to keep FortiGate to Server SSL session state.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server_session_state_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "time", "count", "both"}, false),

				Description: "How to expire SSL sessions for the segment of the SSL connection between the server and the FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"static-nat", "server-load-balance", "access-proxy"}, false),

				Description: "Configure a static NAT server load balance VIP or access proxy.",
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"weblogic_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to add an HTTP header to indicate SSL offloading for a WebLogic server.",
				Optional:    true,
				Computed:    true,
			},
			"websphere_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to add an HTTP header to indicate SSL offloading for a WebSphere server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallVip6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallVip6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallVip6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallVip6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallVip6")
	}

	return resourceFirewallVip6Read(ctx, d, meta)
}

func resourceFirewallVip6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallVip6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallVip6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallVip6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallVip6")
	}

	return resourceFirewallVip6Read(ctx, d, meta)
}

func resourceFirewallVip6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallVip6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallVip6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadFirewallVip6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallVip6 resource: %v", err)
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

	diags := refreshObjectFirewallVip6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallVip6Monitor(v *[]models.FirewallVip6Monitor, sort bool) interface{} {
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

func flattenFirewallVip6Realservers(v *[]models.FirewallVip6Realservers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ClientIp; tmp != nil {
				v["client_ip"] = *tmp
			}

			if tmp := cfg.Healthcheck; tmp != nil {
				v["healthcheck"] = *tmp
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

			if tmp := cfg.MaxConnections; tmp != nil {
				v["max_connections"] = *tmp
			}

			if tmp := cfg.Monitor; tmp != nil {
				v["monitor"] = flattenFirewallVip6RealserversMonitor(tmp, sort)
			}

			if tmp := cfg.Port; tmp != nil {
				v["port"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
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

func flattenFirewallVip6RealserversMonitor(v *[]models.FirewallVip6RealserversMonitor, sort bool) interface{} {
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

func flattenFirewallVip6SrcFilter(v *[]models.FirewallVip6SrcFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Range; tmp != nil {
				v["range"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "range")
	}

	return flat
}

func flattenFirewallVip6SslCipherSuites(v *[]models.FirewallVip6SslCipherSuites, sort bool) interface{} {
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

func flattenFirewallVip6SslServerCipherSuites(v *[]models.FirewallVip6SslServerCipherSuites, sort bool) interface{} {
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

func refreshObjectFirewallVip6(d *schema.ResourceData, o *models.FirewallVip6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddNat64Route != nil {
		v := *o.AddNat64Route

		if err = d.Set("add_nat64_route", v); err != nil {
			return diag.Errorf("error reading add_nat64_route: %v", err)
		}
	}

	if o.ArpReply != nil {
		v := *o.ArpReply

		if err = d.Set("arp_reply", v); err != nil {
			return diag.Errorf("error reading arp_reply: %v", err)
		}
	}

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.EmbeddedIpv4Address != nil {
		v := *o.EmbeddedIpv4Address

		if err = d.Set("embedded_ipv4_address", v); err != nil {
			return diag.Errorf("error reading embedded_ipv4_address: %v", err)
		}
	}

	if o.Extip != nil {
		v := *o.Extip

		if err = d.Set("extip", v); err != nil {
			return diag.Errorf("error reading extip: %v", err)
		}
	}

	if o.Extport != nil {
		v := *o.Extport

		if err = d.Set("extport", v); err != nil {
			return diag.Errorf("error reading extport: %v", err)
		}
	}

	if o.HttpCookieAge != nil {
		v := *o.HttpCookieAge

		if err = d.Set("http_cookie_age", v); err != nil {
			return diag.Errorf("error reading http_cookie_age: %v", err)
		}
	}

	if o.HttpCookieDomain != nil {
		v := *o.HttpCookieDomain

		if err = d.Set("http_cookie_domain", v); err != nil {
			return diag.Errorf("error reading http_cookie_domain: %v", err)
		}
	}

	if o.HttpCookieDomainFromHost != nil {
		v := *o.HttpCookieDomainFromHost

		if err = d.Set("http_cookie_domain_from_host", v); err != nil {
			return diag.Errorf("error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if o.HttpCookieGeneration != nil {
		v := *o.HttpCookieGeneration

		if err = d.Set("http_cookie_generation", v); err != nil {
			return diag.Errorf("error reading http_cookie_generation: %v", err)
		}
	}

	if o.HttpCookiePath != nil {
		v := *o.HttpCookiePath

		if err = d.Set("http_cookie_path", v); err != nil {
			return diag.Errorf("error reading http_cookie_path: %v", err)
		}
	}

	if o.HttpCookieShare != nil {
		v := *o.HttpCookieShare

		if err = d.Set("http_cookie_share", v); err != nil {
			return diag.Errorf("error reading http_cookie_share: %v", err)
		}
	}

	if o.HttpIpHeader != nil {
		v := *o.HttpIpHeader

		if err = d.Set("http_ip_header", v); err != nil {
			return diag.Errorf("error reading http_ip_header: %v", err)
		}
	}

	if o.HttpIpHeaderName != nil {
		v := *o.HttpIpHeaderName

		if err = d.Set("http_ip_header_name", v); err != nil {
			return diag.Errorf("error reading http_ip_header_name: %v", err)
		}
	}

	if o.HttpMultiplex != nil {
		v := *o.HttpMultiplex

		if err = d.Set("http_multiplex", v); err != nil {
			return diag.Errorf("error reading http_multiplex: %v", err)
		}
	}

	if o.HttpRedirect != nil {
		v := *o.HttpRedirect

		if err = d.Set("http_redirect", v); err != nil {
			return diag.Errorf("error reading http_redirect: %v", err)
		}
	}

	if o.HttpsCookieSecure != nil {
		v := *o.HttpsCookieSecure

		if err = d.Set("https_cookie_secure", v); err != nil {
			return diag.Errorf("error reading https_cookie_secure: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Ipv4Mappedip != nil {
		v := *o.Ipv4Mappedip

		if err = d.Set("ipv4_mappedip", v); err != nil {
			return diag.Errorf("error reading ipv4_mappedip: %v", err)
		}
	}

	if o.Ipv4Mappedport != nil {
		v := *o.Ipv4Mappedport

		if err = d.Set("ipv4_mappedport", v); err != nil {
			return diag.Errorf("error reading ipv4_mappedport: %v", err)
		}
	}

	if o.LdbMethod != nil {
		v := *o.LdbMethod

		if err = d.Set("ldb_method", v); err != nil {
			return diag.Errorf("error reading ldb_method: %v", err)
		}
	}

	if o.Mappedip != nil {
		v := *o.Mappedip

		if err = d.Set("mappedip", v); err != nil {
			return diag.Errorf("error reading mappedip: %v", err)
		}
	}

	if o.Mappedport != nil {
		v := *o.Mappedport

		if err = d.Set("mappedport", v); err != nil {
			return diag.Errorf("error reading mappedport: %v", err)
		}
	}

	if o.MaxEmbryonicConnections != nil {
		v := *o.MaxEmbryonicConnections

		if err = d.Set("max_embryonic_connections", v); err != nil {
			return diag.Errorf("error reading max_embryonic_connections: %v", err)
		}
	}

	if o.Monitor != nil {
		if err = d.Set("monitor", flattenFirewallVip6Monitor(o.Monitor, sort)); err != nil {
			return diag.Errorf("error reading monitor: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NatSourceVip != nil {
		v := *o.NatSourceVip

		if err = d.Set("nat_source_vip", v); err != nil {
			return diag.Errorf("error reading nat_source_vip: %v", err)
		}
	}

	if o.Nat64 != nil {
		v := *o.Nat64

		if err = d.Set("nat64", v); err != nil {
			return diag.Errorf("error reading nat64: %v", err)
		}
	}

	if o.Nat66 != nil {
		v := *o.Nat66

		if err = d.Set("nat66", v); err != nil {
			return diag.Errorf("error reading nat66: %v", err)
		}
	}

	if o.OutlookWebAccess != nil {
		v := *o.OutlookWebAccess

		if err = d.Set("outlook_web_access", v); err != nil {
			return diag.Errorf("error reading outlook_web_access: %v", err)
		}
	}

	if o.Persistence != nil {
		v := *o.Persistence

		if err = d.Set("persistence", v); err != nil {
			return diag.Errorf("error reading persistence: %v", err)
		}
	}

	if o.Portforward != nil {
		v := *o.Portforward

		if err = d.Set("portforward", v); err != nil {
			return diag.Errorf("error reading portforward: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Realservers != nil {
		if err = d.Set("realservers", flattenFirewallVip6Realservers(o.Realservers, sort)); err != nil {
			return diag.Errorf("error reading realservers: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
		}
	}

	if o.SrcFilter != nil {
		if err = d.Set("src_filter", flattenFirewallVip6SrcFilter(o.SrcFilter, sort)); err != nil {
			return diag.Errorf("error reading src_filter: %v", err)
		}
	}

	if o.SslAcceptFfdheGroups != nil {
		v := *o.SslAcceptFfdheGroups

		if err = d.Set("ssl_accept_ffdhe_groups", v); err != nil {
			return diag.Errorf("error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if o.SslAlgorithm != nil {
		v := *o.SslAlgorithm

		if err = d.Set("ssl_algorithm", v); err != nil {
			return diag.Errorf("error reading ssl_algorithm: %v", err)
		}
	}

	if o.SslCertificate != nil {
		v := *o.SslCertificate

		if err = d.Set("ssl_certificate", v); err != nil {
			return diag.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	if o.SslCipherSuites != nil {
		if err = d.Set("ssl_cipher_suites", flattenFirewallVip6SslCipherSuites(o.SslCipherSuites, sort)); err != nil {
			return diag.Errorf("error reading ssl_cipher_suites: %v", err)
		}
	}

	if o.SslClientFallback != nil {
		v := *o.SslClientFallback

		if err = d.Set("ssl_client_fallback", v); err != nil {
			return diag.Errorf("error reading ssl_client_fallback: %v", err)
		}
	}

	if o.SslClientRekeyCount != nil {
		v := *o.SslClientRekeyCount

		if err = d.Set("ssl_client_rekey_count", v); err != nil {
			return diag.Errorf("error reading ssl_client_rekey_count: %v", err)
		}
	}

	if o.SslClientRenegotiation != nil {
		v := *o.SslClientRenegotiation

		if err = d.Set("ssl_client_renegotiation", v); err != nil {
			return diag.Errorf("error reading ssl_client_renegotiation: %v", err)
		}
	}

	if o.SslClientSessionStateMax != nil {
		v := *o.SslClientSessionStateMax

		if err = d.Set("ssl_client_session_state_max", v); err != nil {
			return diag.Errorf("error reading ssl_client_session_state_max: %v", err)
		}
	}

	if o.SslClientSessionStateTimeout != nil {
		v := *o.SslClientSessionStateTimeout

		if err = d.Set("ssl_client_session_state_timeout", v); err != nil {
			return diag.Errorf("error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if o.SslClientSessionStateType != nil {
		v := *o.SslClientSessionStateType

		if err = d.Set("ssl_client_session_state_type", v); err != nil {
			return diag.Errorf("error reading ssl_client_session_state_type: %v", err)
		}
	}

	if o.SslDhBits != nil {
		v := *o.SslDhBits

		if err = d.Set("ssl_dh_bits", v); err != nil {
			return diag.Errorf("error reading ssl_dh_bits: %v", err)
		}
	}

	if o.SslHpkp != nil {
		v := *o.SslHpkp

		if err = d.Set("ssl_hpkp", v); err != nil {
			return diag.Errorf("error reading ssl_hpkp: %v", err)
		}
	}

	if o.SslHpkpAge != nil {
		v := *o.SslHpkpAge

		if err = d.Set("ssl_hpkp_age", v); err != nil {
			return diag.Errorf("error reading ssl_hpkp_age: %v", err)
		}
	}

	if o.SslHpkpBackup != nil {
		v := *o.SslHpkpBackup

		if err = d.Set("ssl_hpkp_backup", v); err != nil {
			return diag.Errorf("error reading ssl_hpkp_backup: %v", err)
		}
	}

	if o.SslHpkpIncludeSubdomains != nil {
		v := *o.SslHpkpIncludeSubdomains

		if err = d.Set("ssl_hpkp_include_subdomains", v); err != nil {
			return diag.Errorf("error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if o.SslHpkpPrimary != nil {
		v := *o.SslHpkpPrimary

		if err = d.Set("ssl_hpkp_primary", v); err != nil {
			return diag.Errorf("error reading ssl_hpkp_primary: %v", err)
		}
	}

	if o.SslHpkpReportUri != nil {
		v := *o.SslHpkpReportUri

		if err = d.Set("ssl_hpkp_report_uri", v); err != nil {
			return diag.Errorf("error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if o.SslHsts != nil {
		v := *o.SslHsts

		if err = d.Set("ssl_hsts", v); err != nil {
			return diag.Errorf("error reading ssl_hsts: %v", err)
		}
	}

	if o.SslHstsAge != nil {
		v := *o.SslHstsAge

		if err = d.Set("ssl_hsts_age", v); err != nil {
			return diag.Errorf("error reading ssl_hsts_age: %v", err)
		}
	}

	if o.SslHstsIncludeSubdomains != nil {
		v := *o.SslHstsIncludeSubdomains

		if err = d.Set("ssl_hsts_include_subdomains", v); err != nil {
			return diag.Errorf("error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if o.SslHttpLocationConversion != nil {
		v := *o.SslHttpLocationConversion

		if err = d.Set("ssl_http_location_conversion", v); err != nil {
			return diag.Errorf("error reading ssl_http_location_conversion: %v", err)
		}
	}

	if o.SslHttpMatchHost != nil {
		v := *o.SslHttpMatchHost

		if err = d.Set("ssl_http_match_host", v); err != nil {
			return diag.Errorf("error reading ssl_http_match_host: %v", err)
		}
	}

	if o.SslMaxVersion != nil {
		v := *o.SslMaxVersion

		if err = d.Set("ssl_max_version", v); err != nil {
			return diag.Errorf("error reading ssl_max_version: %v", err)
		}
	}

	if o.SslMinVersion != nil {
		v := *o.SslMinVersion

		if err = d.Set("ssl_min_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_version: %v", err)
		}
	}

	if o.SslMode != nil {
		v := *o.SslMode

		if err = d.Set("ssl_mode", v); err != nil {
			return diag.Errorf("error reading ssl_mode: %v", err)
		}
	}

	if o.SslPfs != nil {
		v := *o.SslPfs

		if err = d.Set("ssl_pfs", v); err != nil {
			return diag.Errorf("error reading ssl_pfs: %v", err)
		}
	}

	if o.SslSendEmptyFrags != nil {
		v := *o.SslSendEmptyFrags

		if err = d.Set("ssl_send_empty_frags", v); err != nil {
			return diag.Errorf("error reading ssl_send_empty_frags: %v", err)
		}
	}

	if o.SslServerAlgorithm != nil {
		v := *o.SslServerAlgorithm

		if err = d.Set("ssl_server_algorithm", v); err != nil {
			return diag.Errorf("error reading ssl_server_algorithm: %v", err)
		}
	}

	if o.SslServerCipherSuites != nil {
		if err = d.Set("ssl_server_cipher_suites", flattenFirewallVip6SslServerCipherSuites(o.SslServerCipherSuites, sort)); err != nil {
			return diag.Errorf("error reading ssl_server_cipher_suites: %v", err)
		}
	}

	if o.SslServerMaxVersion != nil {
		v := *o.SslServerMaxVersion

		if err = d.Set("ssl_server_max_version", v); err != nil {
			return diag.Errorf("error reading ssl_server_max_version: %v", err)
		}
	}

	if o.SslServerMinVersion != nil {
		v := *o.SslServerMinVersion

		if err = d.Set("ssl_server_min_version", v); err != nil {
			return diag.Errorf("error reading ssl_server_min_version: %v", err)
		}
	}

	if o.SslServerSessionStateMax != nil {
		v := *o.SslServerSessionStateMax

		if err = d.Set("ssl_server_session_state_max", v); err != nil {
			return diag.Errorf("error reading ssl_server_session_state_max: %v", err)
		}
	}

	if o.SslServerSessionStateTimeout != nil {
		v := *o.SslServerSessionStateTimeout

		if err = d.Set("ssl_server_session_state_timeout", v); err != nil {
			return diag.Errorf("error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if o.SslServerSessionStateType != nil {
		v := *o.SslServerSessionStateType

		if err = d.Set("ssl_server_session_state_type", v); err != nil {
			return diag.Errorf("error reading ssl_server_session_state_type: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	if o.WeblogicServer != nil {
		v := *o.WeblogicServer

		if err = d.Set("weblogic_server", v); err != nil {
			return diag.Errorf("error reading weblogic_server: %v", err)
		}
	}

	if o.WebsphereServer != nil {
		v := *o.WebsphereServer

		if err = d.Set("websphere_server", v); err != nil {
			return diag.Errorf("error reading websphere_server: %v", err)
		}
	}

	return nil
}

func expandFirewallVip6Monitor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip6Monitor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip6Monitor

	for i := range l {
		tmp := models.FirewallVip6Monitor{}
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

func expandFirewallVip6Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip6Realservers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip6Realservers

	for i := range l {
		tmp := models.FirewallVip6Realservers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.client_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.healthcheck", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Healthcheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.holddown_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.HolddownInterval = &v3
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
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_connections", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxConnections = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.monitor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallVip6RealserversMonitor(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallVip6RealserversMonitor
			// 	}
			tmp.Monitor = v2

		}

		pre_append = fmt.Sprintf("%s.%d.port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Port = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.weight", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Weight = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallVip6RealserversMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip6RealserversMonitor, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip6RealserversMonitor

	for i := range l {
		tmp := models.FirewallVip6RealserversMonitor{}
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

func expandFirewallVip6SrcFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip6SrcFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip6SrcFilter

	for i := range l {
		tmp := models.FirewallVip6SrcFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Range = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallVip6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip6SslCipherSuites, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip6SslCipherSuites

	for i := range l {
		tmp := models.FirewallVip6SslCipherSuites{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
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

func expandFirewallVip6SslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallVip6SslServerCipherSuites, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallVip6SslServerCipherSuites

	for i := range l {
		tmp := models.FirewallVip6SslServerCipherSuites{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priority", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Priority = &v3
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

func getObjectFirewallVip6(d *schema.ResourceData, sv string) (*models.FirewallVip6, diag.Diagnostics) {
	obj := models.FirewallVip6{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("add_nat64_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("add_nat64_route", sv)
				diags = append(diags, e)
			}
			obj.AddNat64Route = &v2
		}
	}
	if v1, ok := d.GetOk("arp_reply"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("arp_reply", sv)
				diags = append(diags, e)
			}
			obj.ArpReply = &v2
		}
	}
	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("embedded_ipv4_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("embedded_ipv4_address", sv)
				diags = append(diags, e)
			}
			obj.EmbeddedIpv4Address = &v2
		}
	}
	if v1, ok := d.GetOk("extip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extip", sv)
				diags = append(diags, e)
			}
			obj.Extip = &v2
		}
	}
	if v1, ok := d.GetOk("extport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extport", sv)
				diags = append(diags, e)
			}
			obj.Extport = &v2
		}
	}
	if v1, ok := d.GetOk("http_cookie_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_cookie_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HttpCookieAge = &tmp
		}
	}
	if v1, ok := d.GetOk("http_cookie_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_cookie_domain", sv)
				diags = append(diags, e)
			}
			obj.HttpCookieDomain = &v2
		}
	}
	if v1, ok := d.GetOk("http_cookie_domain_from_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_cookie_domain_from_host", sv)
				diags = append(diags, e)
			}
			obj.HttpCookieDomainFromHost = &v2
		}
	}
	if v1, ok := d.GetOk("http_cookie_generation"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_cookie_generation", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HttpCookieGeneration = &tmp
		}
	}
	if v1, ok := d.GetOk("http_cookie_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_cookie_path", sv)
				diags = append(diags, e)
			}
			obj.HttpCookiePath = &v2
		}
	}
	if v1, ok := d.GetOk("http_cookie_share"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_cookie_share", sv)
				diags = append(diags, e)
			}
			obj.HttpCookieShare = &v2
		}
	}
	if v1, ok := d.GetOk("http_ip_header"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_ip_header", sv)
				diags = append(diags, e)
			}
			obj.HttpIpHeader = &v2
		}
	}
	if v1, ok := d.GetOk("http_ip_header_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_ip_header_name", sv)
				diags = append(diags, e)
			}
			obj.HttpIpHeaderName = &v2
		}
	}
	if v1, ok := d.GetOk("http_multiplex"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_multiplex", sv)
				diags = append(diags, e)
			}
			obj.HttpMultiplex = &v2
		}
	}
	if v1, ok := d.GetOk("http_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_redirect", sv)
				diags = append(diags, e)
			}
			obj.HttpRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("https_cookie_secure"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_cookie_secure", sv)
				diags = append(diags, e)
			}
			obj.HttpsCookieSecure = &v2
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("ipv4_mappedip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ipv4_mappedip", sv)
				diags = append(diags, e)
			}
			obj.Ipv4Mappedip = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_mappedport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ipv4_mappedport", sv)
				diags = append(diags, e)
			}
			obj.Ipv4Mappedport = &v2
		}
	}
	if v1, ok := d.GetOk("ldb_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldb_method", sv)
				diags = append(diags, e)
			}
			obj.LdbMethod = &v2
		}
	}
	if v1, ok := d.GetOk("mappedip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mappedip", sv)
				diags = append(diags, e)
			}
			obj.Mappedip = &v2
		}
	}
	if v1, ok := d.GetOk("mappedport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mappedport", sv)
				diags = append(diags, e)
			}
			obj.Mappedport = &v2
		}
	}
	if v1, ok := d.GetOk("max_embryonic_connections"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("max_embryonic_connections", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MaxEmbryonicConnections = &tmp
		}
	}
	if v, ok := d.GetOk("monitor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("monitor", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip6Monitor(d, v, "monitor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Monitor = t
		}
	} else if d.HasChange("monitor") {
		old, new := d.GetChange("monitor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Monitor = &[]models.FirewallVip6Monitor{}
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
	if v1, ok := d.GetOk("nat_source_vip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("nat_source_vip", sv)
				diags = append(diags, e)
			}
			obj.NatSourceVip = &v2
		}
	}
	if v1, ok := d.GetOk("nat64"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nat64", sv)
				diags = append(diags, e)
			}
			obj.Nat64 = &v2
		}
	}
	if v1, ok := d.GetOk("nat66"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nat66", sv)
				diags = append(diags, e)
			}
			obj.Nat66 = &v2
		}
	}
	if v1, ok := d.GetOk("outlook_web_access"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outlook_web_access", sv)
				diags = append(diags, e)
			}
			obj.OutlookWebAccess = &v2
		}
	}
	if v1, ok := d.GetOk("persistence"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("persistence", sv)
				diags = append(diags, e)
			}
			obj.Persistence = &v2
		}
	}
	if v1, ok := d.GetOk("portforward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("portforward", sv)
				diags = append(diags, e)
			}
			obj.Portforward = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			obj.Protocol = &v2
		}
	}
	if v, ok := d.GetOk("realservers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("realservers", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip6Realservers(d, v, "realservers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Realservers = t
		}
	} else if d.HasChange("realservers") {
		old, new := d.GetChange("realservers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Realservers = &[]models.FirewallVip6Realservers{}
		}
	}
	if v1, ok := d.GetOk("server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_type", sv)
				diags = append(diags, e)
			}
			obj.ServerType = &v2
		}
	}
	if v, ok := d.GetOk("src_filter"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("src_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip6SrcFilter(d, v, "src_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SrcFilter = t
		}
	} else if d.HasChange("src_filter") {
		old, new := d.GetChange("src_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SrcFilter = &[]models.FirewallVip6SrcFilter{}
		}
	}
	if v1, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("ssl_accept_ffdhe_groups", sv)
				diags = append(diags, e)
			}
			obj.SslAcceptFfdheGroups = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_algorithm", sv)
				diags = append(diags, e)
			}
			obj.SslAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_certificate", sv)
				diags = append(diags, e)
			}
			obj.SslCertificate = &v2
		}
	}
	if v, ok := d.GetOk("ssl_cipher_suites"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssl_cipher_suites", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip6SslCipherSuites(d, v, "ssl_cipher_suites", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SslCipherSuites = t
		}
	} else if d.HasChange("ssl_cipher_suites") {
		old, new := d.GetChange("ssl_cipher_suites")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SslCipherSuites = &[]models.FirewallVip6SslCipherSuites{}
		}
	}
	if v1, ok := d.GetOk("ssl_client_fallback"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_fallback", sv)
				diags = append(diags, e)
			}
			obj.SslClientFallback = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_client_rekey_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_rekey_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslClientRekeyCount = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_client_renegotiation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_renegotiation", sv)
				diags = append(diags, e)
			}
			obj.SslClientRenegotiation = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_client_session_state_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_session_state_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslClientSessionStateMax = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_client_session_state_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_session_state_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslClientSessionStateTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_client_session_state_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_session_state_type", sv)
				diags = append(diags, e)
			}
			obj.SslClientSessionStateType = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_dh_bits"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_dh_bits", sv)
				diags = append(diags, e)
			}
			obj.SslDhBits = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hpkp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hpkp", sv)
				diags = append(diags, e)
			}
			obj.SslHpkp = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hpkp_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hpkp_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslHpkpAge = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_hpkp_backup"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hpkp_backup", sv)
				diags = append(diags, e)
			}
			obj.SslHpkpBackup = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hpkp_include_subdomains", sv)
				diags = append(diags, e)
			}
			obj.SslHpkpIncludeSubdomains = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hpkp_primary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hpkp_primary", sv)
				diags = append(diags, e)
			}
			obj.SslHpkpPrimary = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hpkp_report_uri"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hpkp_report_uri", sv)
				diags = append(diags, e)
			}
			obj.SslHpkpReportUri = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hsts"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hsts", sv)
				diags = append(diags, e)
			}
			obj.SslHsts = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_hsts_age"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hsts_age", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslHstsAge = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_hsts_include_subdomains"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_hsts_include_subdomains", sv)
				diags = append(diags, e)
			}
			obj.SslHstsIncludeSubdomains = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_http_location_conversion"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_http_location_conversion", sv)
				diags = append(diags, e)
			}
			obj.SslHttpLocationConversion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_http_match_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_http_match_host", sv)
				diags = append(diags, e)
			}
			obj.SslHttpMatchHost = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_max_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_max_version", sv)
				diags = append(diags, e)
			}
			obj.SslMaxVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_mode", sv)
				diags = append(diags, e)
			}
			obj.SslMode = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_pfs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_pfs", sv)
				diags = append(diags, e)
			}
			obj.SslPfs = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_send_empty_frags"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_send_empty_frags", sv)
				diags = append(diags, e)
			}
			obj.SslSendEmptyFrags = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_server_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_server_algorithm", sv)
				diags = append(diags, e)
			}
			obj.SslServerAlgorithm = &v2
		}
	}
	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssl_server_cipher_suites", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallVip6SslServerCipherSuites(d, v, "ssl_server_cipher_suites", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SslServerCipherSuites = t
		}
	} else if d.HasChange("ssl_server_cipher_suites") {
		old, new := d.GetChange("ssl_server_cipher_suites")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SslServerCipherSuites = &[]models.FirewallVip6SslServerCipherSuites{}
		}
	}
	if v1, ok := d.GetOk("ssl_server_max_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_server_max_version", sv)
				diags = append(diags, e)
			}
			obj.SslServerMaxVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_server_min_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_server_min_version", sv)
				diags = append(diags, e)
			}
			obj.SslServerMinVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_server_session_state_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_server_session_state_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslServerSessionStateMax = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_server_session_state_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_server_session_state_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslServerSessionStateTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_server_session_state_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_server_session_state_type", sv)
				diags = append(diags, e)
			}
			obj.SslServerSessionStateType = &v2
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	if v1, ok := d.GetOk("weblogic_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("weblogic_server", sv)
				diags = append(diags, e)
			}
			obj.WeblogicServer = &v2
		}
	}
	if v1, ok := d.GetOk("websphere_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("websphere_server", sv)
				diags = append(diags, e)
			}
			obj.WebsphereServer = &v2
		}
	}
	return &obj, diags
}
