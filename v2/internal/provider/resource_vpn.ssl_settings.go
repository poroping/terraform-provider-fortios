// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL-VPN.

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

func resourceVpnSslSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL-VPN.",

		CreateContext: resourceVpnSslSettingsCreate,
		ReadContext:   resourceVpnSslSettingsRead,
		UpdateContext: resourceVpnSslSettingsUpdate,
		DeleteContext: resourceVpnSslSettingsDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high", "medium", "default", "low"}, false),

				Description: "Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any.",
				Optional:    true,
				Computed:    true,
			},
			"auth_session_check_source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable checking of source IP for authentication session.",
				Optional:    true,
				Computed:    true,
			},
			"auth_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 259200),

				Description: "SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).",
				Optional:    true,
				Computed:    true,
			},
			"authentication_rule": {
				Type:        schema.TypeList,
				Description: "Authentication rule for SSL-VPN.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"any", "local", "radius", "tacacs+", "ldap", "peer"}, false),

							Description: "SSL-VPN authentication method restriction.",
							Optional:    true,
							Computed:    true,
						},
						"cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"any", "high", "medium"}, false),

							Description: "SSL-VPN cipher strength.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SSL-VPN client certificate restrictive.",
							Optional:    true,
							Computed:    true,
						},
						"groups": {
							Type:        schema.TypeList,
							Description: "User groups.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Group name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID (0 - 4294967295).",
							Optional:    true,
							Computed:    true,
						},
						"portal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SSL-VPN portal.",
							Optional:    true,
							Computed:    true,
						},
						"realm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "SSL-VPN realm.",
							Optional:    true,
							Computed:    true,
						},
						"source_address": {
							Type:        schema.TypeList,
							Description: "Source address of incoming traffic.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Address name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"source_address_negate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable negated source address match.",
							Optional:    true,
							Computed:    true,
						},
						"source_address6": {
							Type:        schema.TypeList,
							Description: "IPv6 source address of incoming traffic.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "IPv6 address name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"source_address6_negate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable negated source IPv6 address match.",
							Optional:    true,
							Computed:    true,
						},
						"source_interface": {
							Type:        schema.TypeList,
							Description: "SSL-VPN source interface of incoming traffic.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Interface name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"user_peer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name of user peer.",
							Optional:    true,
							Computed:    true,
						},
						"users": {
							Type:        schema.TypeList,
							Description: "User name.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "User name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"auto_tunnel_static_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable to auto-create static routes for the SSL-VPN tunnel IP addresses.",
				Optional:    true,
				Computed:    true,
			},
			"banned_cipher": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more cipher technologies that cannot be used in SSL-VPN negotiations. Only applies to TLS 1.2 and below.",
				Optional:         true,
				Computed:         true,
			},
			"check_referer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable verification of referer field in HTTP request header.",
				Optional:    true,
				Computed:    true,
			},
			"ciphersuite": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below.",
				Optional:         true,
				Computed:         true,
			},
			"client_sigalgs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"no-rsa-pss", "all"}, false),

				Description: "Set signature algorithms related to client authentication. Affects TLS version <= 1.2 only.",
				Optional:    true,
				Computed:    true,
			},
			"default_portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default SSL-VPN portal.",
				Optional:    true,
				Computed:    true,
			},
			"deflate_compression_level": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9),

				Description: "Compression level (0~9).",
				Optional:    true,
				Computed:    true,
			},
			"deflate_min_data_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(200, 65535),

				Description: "Minimum amount of data that triggers compression (200 - 65535 bytes).",
				Optional:    true,
				Computed:    true,
			},
			"dns_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DNS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "DNS server 2.",
				Optional:    true,
				Computed:    true,
			},
			"dns_suffix": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 253),

				Description: "DNS suffix used for SSL-VPN clients.",
				Optional:    true,
				Computed:    true,
			},
			"dtls_hello_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 60),

				Description: "SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"dtls_max_proto_ver": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dtls1-0", "dtls1-2"}, false),

				Description: "DTLS maximum protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"dtls_min_proto_ver": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dtls1-0", "dtls1-2"}, false),

				Description: "DTLS minimum protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"dtls_tunnel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DTLS to prevent eavesdropping, tampering, or message forgery.",
				Optional:    true,
				Computed:    true,
			},
			"dual_stack_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal.",
				Optional:    true,
				Computed:    true,
			},
			"encode_2f_sequence": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Encode \\2F sequence to forward slash in URLs.",
				Optional:    true,
				Computed:    true,
			},
			"encrypt_and_store_password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Encrypt and store user passwords for SSL-VPN web sessions.",
				Optional:    true,
				Computed:    true,
			},
			"force_two_factor_auth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable only PKI users with two-factor authentication for SSL-VPNs.",
				Optional:    true,
				Computed:    true,
			},
			"header_x_forwarded_for": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"pass", "add", "remove"}, false),

				Description: "Forward the same, add, or remove HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"hsts_include_subdomains": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Add HSTS includeSubDomains response header.",
				Optional:    true,
				Computed:    true,
			},
			"http_compression": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable to allow HTTP compression over SSL-VPN tunnels.",
				Optional:    true,
				Computed:    true,
			},
			"http_only_cookie": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN support for HttpOnly cookies.",
				Optional:    true,
				Computed:    true,
			},
			"http_request_body_timeout": {
				Type: schema.TypeInt,

				Description: "SSL-VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"http_request_header_timeout": {
				Type: schema.TypeInt,

				Description: "SSL-VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"https_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable redirect of port 80 to SSL-VPN port.",
				Optional:    true,
				Computed:    true,
			},
			"idle_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 259200),

				Description: "SSL-VPN disconnects if idle for specified time in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"ipv6_dns_server1": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 1.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_dns_server2": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 DNS server 2.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_wins_server1": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 WINS server 1.",
				Optional:         true,
				Computed:         true,
			},
			"ipv6_wins_server2": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 WINS server 2.",
				Optional:         true,
				Computed:         true,
			},
			"login_attempt_limit": {
				Type: schema.TypeInt,

				Description: "SSL-VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).",
				Optional:    true,
				Computed:    true,
			},
			"login_block_time": {
				Type: schema.TypeInt,

				Description: "Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"login_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 180),

				Description: "SSLVPN maximum login timeout (10 - 180 sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "SSL-VPN access port (1 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"port_precedence": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable, Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface.",
				Optional:    true,
				Computed:    true,
			},
			"reqclientcert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable to require client certificates for all SSL-VPN users.",
				Optional:    true,
				Computed:    true,
			},
			"route_source_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to allow SSL-VPN sessions to bypass routing and bind to the incoming interface.",
				Optional:    true,
				Computed:    true,
			},
			"saml_redirect_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "SAML local redirect port in the machine running FortiClient (0 - 65535). 0 is to disable redirection on FGT side.",
				Optional:    true,
				Computed:    true,
			},
			"servercert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the server certificate to be used for SSL-VPNs.",
				Optional:    true,
				Computed:    true,
			},
			"source_address": {
				Type:        schema.TypeList,
				Description: "Source address of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"source_address_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negated source address match.",
				Optional:    true,
				Computed:    true,
			},
			"source_address6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPv6 address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"source_address6_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negated source IPv6 address match.",
				Optional:    true,
				Computed:    true,
			},
			"source_interface": {
				Type:        schema.TypeList,
				Description: "SSL-VPN source interface of incoming traffic.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl_client_renegotiation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable to allow client renegotiation by the server if the tunnel goes down.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_insert_empty_fragment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable insertion of empty fragment.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_max_proto_ver": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tls1-0", "tls1-1", "tls1-2", "tls1-3"}, false),

				Description: "SSL maximum protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_ver": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tls1-0", "tls1-1", "tls1-2", "tls1-3"}, false),

				Description: "SSL minimum protocol version.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN.",
				Optional:    true,
				Computed:    true,
			},
			"tlsv1_0": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "tlsv1-0",
				Optional:    true,
				Computed:    true,
			},
			"tlsv1_1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "tlsv1-1",
				Optional:    true,
				Computed:    true,
			},
			"tlsv1_2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "tlsv1-2",
				Optional:    true,
				Computed:    true,
			},
			"tlsv1_3": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "tlsv1-3",
				Optional:    true,
				Computed:    true,
			},
			"transform_backward_slashes": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Transform backward slashes to forward slashes in URLs.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_addr_assigned_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"first-available", "round-robin"}, false),

				Description: "Method used for assigning address for tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_connect_without_reauth": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable tunnel connection without re-authorization if previous connection dropped.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_ip_pools": {
				Type:        schema.TypeList,
				Description: "Names of the IPv4 IP Pool firewall objects that define the IP addresses reserved for remote clients.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"tunnel_ipv6_pools": {
				Type:        schema.TypeList,
				Description: "Names of the IPv6 IP Pool firewall objects that define the IP addresses reserved for remote clients.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"tunnel_user_session_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Time out value to clean up user session after tunnel connection is dropped (1 - 255 sec, default=30).",
				Optional:    true,
				Computed:    true,
			},
			"unsafe_legacy_renegotiation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable unsafe legacy re-negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"url_obscuration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable to obscure the host name of the URL of the web browser display.",
				Optional:    true,
				Computed:    true,
			},
			"user_peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of user peer.",
				Optional:    true,
				Computed:    true,
			},
			"wins_server1": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server 1.",
				Optional:    true,
				Computed:    true,
			},
			"wins_server2": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "WINS server 2.",
				Optional:    true,
				Computed:    true,
			},
			"x_content_type_options": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Add HTTP X-Content-Type-Options header.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnSslSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnSslSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnSslSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslSettings")
	}

	return resourceVpnSslSettingsRead(ctx, d, meta)
}

func resourceVpnSslSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnSslSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnSslSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnSslSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnSslSettings")
	}

	return resourceVpnSslSettingsRead(ctx, d, meta)
}

func resourceVpnSslSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectVpnSslSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateVpnSslSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnSslSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnSslSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslSettings resource: %v", err)
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

	diags := refreshObjectVpnSslSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnSslSettingsAuthenticationRule(v *[]models.VpnSslSettingsAuthenticationRule, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Auth; tmp != nil {
				v["auth"] = *tmp
			}

			if tmp := cfg.Cipher; tmp != nil {
				v["cipher"] = *tmp
			}

			if tmp := cfg.ClientCert; tmp != nil {
				v["client_cert"] = *tmp
			}

			if tmp := cfg.Groups; tmp != nil {
				v["groups"] = flattenVpnSslSettingsAuthenticationRuleGroups(tmp, sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Portal; tmp != nil {
				v["portal"] = *tmp
			}

			if tmp := cfg.Realm; tmp != nil {
				v["realm"] = *tmp
			}

			if tmp := cfg.SourceAddress; tmp != nil {
				v["source_address"] = flattenVpnSslSettingsAuthenticationRuleSourceAddress(tmp, sort)
			}

			if tmp := cfg.SourceAddressNegate; tmp != nil {
				v["source_address_negate"] = *tmp
			}

			if tmp := cfg.SourceAddress6; tmp != nil {
				v["source_address6"] = flattenVpnSslSettingsAuthenticationRuleSourceAddress6(tmp, sort)
			}

			if tmp := cfg.SourceAddress6Negate; tmp != nil {
				v["source_address6_negate"] = *tmp
			}

			if tmp := cfg.SourceInterface; tmp != nil {
				v["source_interface"] = flattenVpnSslSettingsAuthenticationRuleSourceInterface(tmp, sort)
			}

			if tmp := cfg.UserPeer; tmp != nil {
				v["user_peer"] = *tmp
			}

			if tmp := cfg.Users; tmp != nil {
				v["users"] = flattenVpnSslSettingsAuthenticationRuleUsers(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenVpnSslSettingsAuthenticationRuleGroups(v *[]models.VpnSslSettingsAuthenticationRuleGroups, sort bool) interface{} {
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

func flattenVpnSslSettingsAuthenticationRuleSourceAddress(v *[]models.VpnSslSettingsAuthenticationRuleSourceAddress, sort bool) interface{} {
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

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6(v *[]models.VpnSslSettingsAuthenticationRuleSourceAddress6, sort bool) interface{} {
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

func flattenVpnSslSettingsAuthenticationRuleSourceInterface(v *[]models.VpnSslSettingsAuthenticationRuleSourceInterface, sort bool) interface{} {
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

func flattenVpnSslSettingsAuthenticationRuleUsers(v *[]models.VpnSslSettingsAuthenticationRuleUsers, sort bool) interface{} {
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

func flattenVpnSslSettingsSourceAddress(v *[]models.VpnSslSettingsSourceAddress, sort bool) interface{} {
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

func flattenVpnSslSettingsSourceAddress6(v *[]models.VpnSslSettingsSourceAddress6, sort bool) interface{} {
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

func flattenVpnSslSettingsSourceInterface(v *[]models.VpnSslSettingsSourceInterface, sort bool) interface{} {
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

func flattenVpnSslSettingsTunnelIpPools(v *[]models.VpnSslSettingsTunnelIpPools, sort bool) interface{} {
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

func flattenVpnSslSettingsTunnelIpv6Pools(v *[]models.VpnSslSettingsTunnelIpv6Pools, sort bool) interface{} {
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

func refreshObjectVpnSslSettings(d *schema.ResourceData, o *models.VpnSslSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Algorithm != nil {
		v := *o.Algorithm

		if err = d.Set("algorithm", v); err != nil {
			return diag.Errorf("error reading algorithm: %v", err)
		}
	}

	if o.AuthSessionCheckSourceIp != nil {
		v := *o.AuthSessionCheckSourceIp

		if err = d.Set("auth_session_check_source_ip", v); err != nil {
			return diag.Errorf("error reading auth_session_check_source_ip: %v", err)
		}
	}

	if o.AuthTimeout != nil {
		v := *o.AuthTimeout

		if err = d.Set("auth_timeout", v); err != nil {
			return diag.Errorf("error reading auth_timeout: %v", err)
		}
	}

	if o.AuthenticationRule != nil {
		if err = d.Set("authentication_rule", flattenVpnSslSettingsAuthenticationRule(o.AuthenticationRule, sort)); err != nil {
			return diag.Errorf("error reading authentication_rule: %v", err)
		}
	}

	if o.AutoTunnelStaticRoute != nil {
		v := *o.AutoTunnelStaticRoute

		if err = d.Set("auto_tunnel_static_route", v); err != nil {
			return diag.Errorf("error reading auto_tunnel_static_route: %v", err)
		}
	}

	if o.BannedCipher != nil {
		v := *o.BannedCipher

		if err = d.Set("banned_cipher", v); err != nil {
			return diag.Errorf("error reading banned_cipher: %v", err)
		}
	}

	if o.CheckReferer != nil {
		v := *o.CheckReferer

		if err = d.Set("check_referer", v); err != nil {
			return diag.Errorf("error reading check_referer: %v", err)
		}
	}

	if o.Ciphersuite != nil {
		v := *o.Ciphersuite

		if err = d.Set("ciphersuite", v); err != nil {
			return diag.Errorf("error reading ciphersuite: %v", err)
		}
	}

	if o.ClientSigalgs != nil {
		v := *o.ClientSigalgs

		if err = d.Set("client_sigalgs", v); err != nil {
			return diag.Errorf("error reading client_sigalgs: %v", err)
		}
	}

	if o.DefaultPortal != nil {
		v := *o.DefaultPortal

		if err = d.Set("default_portal", v); err != nil {
			return diag.Errorf("error reading default_portal: %v", err)
		}
	}

	if o.DeflateCompressionLevel != nil {
		v := *o.DeflateCompressionLevel

		if err = d.Set("deflate_compression_level", v); err != nil {
			return diag.Errorf("error reading deflate_compression_level: %v", err)
		}
	}

	if o.DeflateMinDataSize != nil {
		v := *o.DeflateMinDataSize

		if err = d.Set("deflate_min_data_size", v); err != nil {
			return diag.Errorf("error reading deflate_min_data_size: %v", err)
		}
	}

	if o.DnsServer1 != nil {
		v := *o.DnsServer1

		if err = d.Set("dns_server1", v); err != nil {
			return diag.Errorf("error reading dns_server1: %v", err)
		}
	}

	if o.DnsServer2 != nil {
		v := *o.DnsServer2

		if err = d.Set("dns_server2", v); err != nil {
			return diag.Errorf("error reading dns_server2: %v", err)
		}
	}

	if o.DnsSuffix != nil {
		v := *o.DnsSuffix

		if err = d.Set("dns_suffix", v); err != nil {
			return diag.Errorf("error reading dns_suffix: %v", err)
		}
	}

	if o.DtlsHelloTimeout != nil {
		v := *o.DtlsHelloTimeout

		if err = d.Set("dtls_hello_timeout", v); err != nil {
			return diag.Errorf("error reading dtls_hello_timeout: %v", err)
		}
	}

	if o.DtlsMaxProtoVer != nil {
		v := *o.DtlsMaxProtoVer

		if err = d.Set("dtls_max_proto_ver", v); err != nil {
			return diag.Errorf("error reading dtls_max_proto_ver: %v", err)
		}
	}

	if o.DtlsMinProtoVer != nil {
		v := *o.DtlsMinProtoVer

		if err = d.Set("dtls_min_proto_ver", v); err != nil {
			return diag.Errorf("error reading dtls_min_proto_ver: %v", err)
		}
	}

	if o.DtlsTunnel != nil {
		v := *o.DtlsTunnel

		if err = d.Set("dtls_tunnel", v); err != nil {
			return diag.Errorf("error reading dtls_tunnel: %v", err)
		}
	}

	if o.DualStackMode != nil {
		v := *o.DualStackMode

		if err = d.Set("dual_stack_mode", v); err != nil {
			return diag.Errorf("error reading dual_stack_mode: %v", err)
		}
	}

	if o.Encode2fSequence != nil {
		v := *o.Encode2fSequence

		if err = d.Set("encode_2f_sequence", v); err != nil {
			return diag.Errorf("error reading encode_2f_sequence: %v", err)
		}
	}

	if o.EncryptAndStorePassword != nil {
		v := *o.EncryptAndStorePassword

		if err = d.Set("encrypt_and_store_password", v); err != nil {
			return diag.Errorf("error reading encrypt_and_store_password: %v", err)
		}
	}

	if o.ForceTwoFactorAuth != nil {
		v := *o.ForceTwoFactorAuth

		if err = d.Set("force_two_factor_auth", v); err != nil {
			return diag.Errorf("error reading force_two_factor_auth: %v", err)
		}
	}

	if o.HeaderXForwardedFor != nil {
		v := *o.HeaderXForwardedFor

		if err = d.Set("header_x_forwarded_for", v); err != nil {
			return diag.Errorf("error reading header_x_forwarded_for: %v", err)
		}
	}

	if o.HstsIncludeSubdomains != nil {
		v := *o.HstsIncludeSubdomains

		if err = d.Set("hsts_include_subdomains", v); err != nil {
			return diag.Errorf("error reading hsts_include_subdomains: %v", err)
		}
	}

	if o.HttpCompression != nil {
		v := *o.HttpCompression

		if err = d.Set("http_compression", v); err != nil {
			return diag.Errorf("error reading http_compression: %v", err)
		}
	}

	if o.HttpOnlyCookie != nil {
		v := *o.HttpOnlyCookie

		if err = d.Set("http_only_cookie", v); err != nil {
			return diag.Errorf("error reading http_only_cookie: %v", err)
		}
	}

	if o.HttpRequestBodyTimeout != nil {
		v := *o.HttpRequestBodyTimeout

		if err = d.Set("http_request_body_timeout", v); err != nil {
			return diag.Errorf("error reading http_request_body_timeout: %v", err)
		}
	}

	if o.HttpRequestHeaderTimeout != nil {
		v := *o.HttpRequestHeaderTimeout

		if err = d.Set("http_request_header_timeout", v); err != nil {
			return diag.Errorf("error reading http_request_header_timeout: %v", err)
		}
	}

	if o.HttpsRedirect != nil {
		v := *o.HttpsRedirect

		if err = d.Set("https_redirect", v); err != nil {
			return diag.Errorf("error reading https_redirect: %v", err)
		}
	}

	if o.IdleTimeout != nil {
		v := *o.IdleTimeout

		if err = d.Set("idle_timeout", v); err != nil {
			return diag.Errorf("error reading idle_timeout: %v", err)
		}
	}

	if o.Ipv6DnsServer1 != nil {
		v := *o.Ipv6DnsServer1

		if err = d.Set("ipv6_dns_server1", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server1: %v", err)
		}
	}

	if o.Ipv6DnsServer2 != nil {
		v := *o.Ipv6DnsServer2

		if err = d.Set("ipv6_dns_server2", v); err != nil {
			return diag.Errorf("error reading ipv6_dns_server2: %v", err)
		}
	}

	if o.Ipv6WinsServer1 != nil {
		v := *o.Ipv6WinsServer1

		if err = d.Set("ipv6_wins_server1", v); err != nil {
			return diag.Errorf("error reading ipv6_wins_server1: %v", err)
		}
	}

	if o.Ipv6WinsServer2 != nil {
		v := *o.Ipv6WinsServer2

		if err = d.Set("ipv6_wins_server2", v); err != nil {
			return diag.Errorf("error reading ipv6_wins_server2: %v", err)
		}
	}

	if o.LoginAttemptLimit != nil {
		v := *o.LoginAttemptLimit

		if err = d.Set("login_attempt_limit", v); err != nil {
			return diag.Errorf("error reading login_attempt_limit: %v", err)
		}
	}

	if o.LoginBlockTime != nil {
		v := *o.LoginBlockTime

		if err = d.Set("login_block_time", v); err != nil {
			return diag.Errorf("error reading login_block_time: %v", err)
		}
	}

	if o.LoginTimeout != nil {
		v := *o.LoginTimeout

		if err = d.Set("login_timeout", v); err != nil {
			return diag.Errorf("error reading login_timeout: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.PortPrecedence != nil {
		v := *o.PortPrecedence

		if err = d.Set("port_precedence", v); err != nil {
			return diag.Errorf("error reading port_precedence: %v", err)
		}
	}

	if o.Reqclientcert != nil {
		v := *o.Reqclientcert

		if err = d.Set("reqclientcert", v); err != nil {
			return diag.Errorf("error reading reqclientcert: %v", err)
		}
	}

	if o.RouteSourceInterface != nil {
		v := *o.RouteSourceInterface

		if err = d.Set("route_source_interface", v); err != nil {
			return diag.Errorf("error reading route_source_interface: %v", err)
		}
	}

	if o.SamlRedirectPort != nil {
		v := *o.SamlRedirectPort

		if err = d.Set("saml_redirect_port", v); err != nil {
			return diag.Errorf("error reading saml_redirect_port: %v", err)
		}
	}

	if o.Servercert != nil {
		v := *o.Servercert

		if err = d.Set("servercert", v); err != nil {
			return diag.Errorf("error reading servercert: %v", err)
		}
	}

	if o.SourceAddress != nil {
		if err = d.Set("source_address", flattenVpnSslSettingsSourceAddress(o.SourceAddress, sort)); err != nil {
			return diag.Errorf("error reading source_address: %v", err)
		}
	}

	if o.SourceAddressNegate != nil {
		v := *o.SourceAddressNegate

		if err = d.Set("source_address_negate", v); err != nil {
			return diag.Errorf("error reading source_address_negate: %v", err)
		}
	}

	if o.SourceAddress6 != nil {
		if err = d.Set("source_address6", flattenVpnSslSettingsSourceAddress6(o.SourceAddress6, sort)); err != nil {
			return diag.Errorf("error reading source_address6: %v", err)
		}
	}

	if o.SourceAddress6Negate != nil {
		v := *o.SourceAddress6Negate

		if err = d.Set("source_address6_negate", v); err != nil {
			return diag.Errorf("error reading source_address6_negate: %v", err)
		}
	}

	if o.SourceInterface != nil {
		if err = d.Set("source_interface", flattenVpnSslSettingsSourceInterface(o.SourceInterface, sort)); err != nil {
			return diag.Errorf("error reading source_interface: %v", err)
		}
	}

	if o.SslClientRenegotiation != nil {
		v := *o.SslClientRenegotiation

		if err = d.Set("ssl_client_renegotiation", v); err != nil {
			return diag.Errorf("error reading ssl_client_renegotiation: %v", err)
		}
	}

	if o.SslInsertEmptyFragment != nil {
		v := *o.SslInsertEmptyFragment

		if err = d.Set("ssl_insert_empty_fragment", v); err != nil {
			return diag.Errorf("error reading ssl_insert_empty_fragment: %v", err)
		}
	}

	if o.SslMaxProtoVer != nil {
		v := *o.SslMaxProtoVer

		if err = d.Set("ssl_max_proto_ver", v); err != nil {
			return diag.Errorf("error reading ssl_max_proto_ver: %v", err)
		}
	}

	if o.SslMinProtoVer != nil {
		v := *o.SslMinProtoVer

		if err = d.Set("ssl_min_proto_ver", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_ver: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tlsv10 != nil {
		v := *o.Tlsv10

		if err = d.Set("tlsv1_0", v); err != nil {
			return diag.Errorf("error reading tlsv1_0: %v", err)
		}
	}

	if o.Tlsv11 != nil {
		v := *o.Tlsv11

		if err = d.Set("tlsv1_1", v); err != nil {
			return diag.Errorf("error reading tlsv1_1: %v", err)
		}
	}

	if o.Tlsv12 != nil {
		v := *o.Tlsv12

		if err = d.Set("tlsv1_2", v); err != nil {
			return diag.Errorf("error reading tlsv1_2: %v", err)
		}
	}

	if o.Tlsv13 != nil {
		v := *o.Tlsv13

		if err = d.Set("tlsv1_3", v); err != nil {
			return diag.Errorf("error reading tlsv1_3: %v", err)
		}
	}

	if o.TransformBackwardSlashes != nil {
		v := *o.TransformBackwardSlashes

		if err = d.Set("transform_backward_slashes", v); err != nil {
			return diag.Errorf("error reading transform_backward_slashes: %v", err)
		}
	}

	if o.TunnelAddrAssignedMethod != nil {
		v := *o.TunnelAddrAssignedMethod

		if err = d.Set("tunnel_addr_assigned_method", v); err != nil {
			return diag.Errorf("error reading tunnel_addr_assigned_method: %v", err)
		}
	}

	if o.TunnelConnectWithoutReauth != nil {
		v := *o.TunnelConnectWithoutReauth

		if err = d.Set("tunnel_connect_without_reauth", v); err != nil {
			return diag.Errorf("error reading tunnel_connect_without_reauth: %v", err)
		}
	}

	if o.TunnelIpPools != nil {
		if err = d.Set("tunnel_ip_pools", flattenVpnSslSettingsTunnelIpPools(o.TunnelIpPools, sort)); err != nil {
			return diag.Errorf("error reading tunnel_ip_pools: %v", err)
		}
	}

	if o.TunnelIpv6Pools != nil {
		if err = d.Set("tunnel_ipv6_pools", flattenVpnSslSettingsTunnelIpv6Pools(o.TunnelIpv6Pools, sort)); err != nil {
			return diag.Errorf("error reading tunnel_ipv6_pools: %v", err)
		}
	}

	if o.TunnelUserSessionTimeout != nil {
		v := *o.TunnelUserSessionTimeout

		if err = d.Set("tunnel_user_session_timeout", v); err != nil {
			return diag.Errorf("error reading tunnel_user_session_timeout: %v", err)
		}
	}

	if o.UnsafeLegacyRenegotiation != nil {
		v := *o.UnsafeLegacyRenegotiation

		if err = d.Set("unsafe_legacy_renegotiation", v); err != nil {
			return diag.Errorf("error reading unsafe_legacy_renegotiation: %v", err)
		}
	}

	if o.UrlObscuration != nil {
		v := *o.UrlObscuration

		if err = d.Set("url_obscuration", v); err != nil {
			return diag.Errorf("error reading url_obscuration: %v", err)
		}
	}

	if o.UserPeer != nil {
		v := *o.UserPeer

		if err = d.Set("user_peer", v); err != nil {
			return diag.Errorf("error reading user_peer: %v", err)
		}
	}

	if o.WinsServer1 != nil {
		v := *o.WinsServer1

		if err = d.Set("wins_server1", v); err != nil {
			return diag.Errorf("error reading wins_server1: %v", err)
		}
	}

	if o.WinsServer2 != nil {
		v := *o.WinsServer2

		if err = d.Set("wins_server2", v); err != nil {
			return diag.Errorf("error reading wins_server2: %v", err)
		}
	}

	if o.XContentTypeOptions != nil {
		v := *o.XContentTypeOptions

		if err = d.Set("x_content_type_options", v); err != nil {
			return diag.Errorf("error reading x_content_type_options: %v", err)
		}
	}

	return nil
}

func expandVpnSslSettingsAuthenticationRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsAuthenticationRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsAuthenticationRule

	for i := range l {
		tmp := models.VpnSslSettingsAuthenticationRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Auth = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Cipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.groups", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslSettingsAuthenticationRuleGroups(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslSettingsAuthenticationRuleGroups
			// 	}
			tmp.Groups = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.portal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Portal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.realm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Realm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslSettingsAuthenticationRuleSourceAddress(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslSettingsAuthenticationRuleSourceAddress
			// 	}
			tmp.SourceAddress = v2

		}

		pre_append = fmt.Sprintf("%s.%d.source_address_negate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceAddressNegate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_address6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslSettingsAuthenticationRuleSourceAddress6(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslSettingsAuthenticationRuleSourceAddress6
			// 	}
			tmp.SourceAddress6 = v2

		}

		pre_append = fmt.Sprintf("%s.%d.source_address6_negate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceAddress6Negate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslSettingsAuthenticationRuleSourceInterface(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslSettingsAuthenticationRuleSourceInterface
			// 	}
			tmp.SourceInterface = v2

		}

		pre_append = fmt.Sprintf("%s.%d.user_peer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UserPeer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.users", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVpnSslSettingsAuthenticationRuleUsers(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VpnSslSettingsAuthenticationRuleUsers
			// 	}
			tmp.Users = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVpnSslSettingsAuthenticationRuleGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsAuthenticationRuleGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsAuthenticationRuleGroups

	for i := range l {
		tmp := models.VpnSslSettingsAuthenticationRuleGroups{}
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

func expandVpnSslSettingsAuthenticationRuleSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsAuthenticationRuleSourceAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsAuthenticationRuleSourceAddress

	for i := range l {
		tmp := models.VpnSslSettingsAuthenticationRuleSourceAddress{}
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

func expandVpnSslSettingsAuthenticationRuleSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsAuthenticationRuleSourceAddress6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsAuthenticationRuleSourceAddress6

	for i := range l {
		tmp := models.VpnSslSettingsAuthenticationRuleSourceAddress6{}
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

func expandVpnSslSettingsAuthenticationRuleSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsAuthenticationRuleSourceInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsAuthenticationRuleSourceInterface

	for i := range l {
		tmp := models.VpnSslSettingsAuthenticationRuleSourceInterface{}
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

func expandVpnSslSettingsAuthenticationRuleUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsAuthenticationRuleUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsAuthenticationRuleUsers

	for i := range l {
		tmp := models.VpnSslSettingsAuthenticationRuleUsers{}
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

func expandVpnSslSettingsSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsSourceAddress, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsSourceAddress

	for i := range l {
		tmp := models.VpnSslSettingsSourceAddress{}
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

func expandVpnSslSettingsSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsSourceAddress6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsSourceAddress6

	for i := range l {
		tmp := models.VpnSslSettingsSourceAddress6{}
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

func expandVpnSslSettingsSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsSourceInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsSourceInterface

	for i := range l {
		tmp := models.VpnSslSettingsSourceInterface{}
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

func expandVpnSslSettingsTunnelIpPools(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsTunnelIpPools, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsTunnelIpPools

	for i := range l {
		tmp := models.VpnSslSettingsTunnelIpPools{}
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

func expandVpnSslSettingsTunnelIpv6Pools(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnSslSettingsTunnelIpv6Pools, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnSslSettingsTunnelIpv6Pools

	for i := range l {
		tmp := models.VpnSslSettingsTunnelIpv6Pools{}
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

func getObjectVpnSslSettings(d *schema.ResourceData, sv string) (*models.VpnSslSettings, diag.Diagnostics) {
	obj := models.VpnSslSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("algorithm", sv)
				diags = append(diags, e)
			}
			obj.Algorithm = &v2
		}
	}
	if v1, ok := d.GetOk("auth_session_check_source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_session_check_source_ip", sv)
				diags = append(diags, e)
			}
			obj.AuthSessionCheckSourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("auth_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AuthTimeout = &tmp
		}
	}
	if v, ok := d.GetOk("authentication_rule"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("authentication_rule", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslSettingsAuthenticationRule(d, v, "authentication_rule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AuthenticationRule = t
		}
	} else if d.HasChange("authentication_rule") {
		old, new := d.GetChange("authentication_rule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AuthenticationRule = &[]models.VpnSslSettingsAuthenticationRule{}
		}
	}
	if v1, ok := d.GetOk("auto_tunnel_static_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_tunnel_static_route", sv)
				diags = append(diags, e)
			}
			obj.AutoTunnelStaticRoute = &v2
		}
	}
	if v1, ok := d.GetOk("banned_cipher"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("banned_cipher", sv)
				diags = append(diags, e)
			}
			obj.BannedCipher = &v2
		}
	}
	if v1, ok := d.GetOk("check_referer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("check_referer", sv)
				diags = append(diags, e)
			}
			obj.CheckReferer = &v2
		}
	}
	if v1, ok := d.GetOk("ciphersuite"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ciphersuite", sv)
				diags = append(diags, e)
			}
			obj.Ciphersuite = &v2
		}
	}
	if v1, ok := d.GetOk("client_sigalgs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("client_sigalgs", sv)
				diags = append(diags, e)
			}
			obj.ClientSigalgs = &v2
		}
	}
	if v1, ok := d.GetOk("default_portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_portal", sv)
				diags = append(diags, e)
			}
			obj.DefaultPortal = &v2
		}
	}
	if v1, ok := d.GetOk("deflate_compression_level"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deflate_compression_level", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeflateCompressionLevel = &tmp
		}
	}
	if v1, ok := d.GetOk("deflate_min_data_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deflate_min_data_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeflateMinDataSize = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server1", sv)
				diags = append(diags, e)
			}
			obj.DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server2", sv)
				diags = append(diags, e)
			}
			obj.DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_suffix"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_suffix", sv)
				diags = append(diags, e)
			}
			obj.DnsSuffix = &v2
		}
	}
	if v1, ok := d.GetOk("dtls_hello_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dtls_hello_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DtlsHelloTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("dtls_max_proto_ver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dtls_max_proto_ver", sv)
				diags = append(diags, e)
			}
			obj.DtlsMaxProtoVer = &v2
		}
	}
	if v1, ok := d.GetOk("dtls_min_proto_ver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dtls_min_proto_ver", sv)
				diags = append(diags, e)
			}
			obj.DtlsMinProtoVer = &v2
		}
	}
	if v1, ok := d.GetOk("dtls_tunnel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dtls_tunnel", sv)
				diags = append(diags, e)
			}
			obj.DtlsTunnel = &v2
		}
	}
	if v1, ok := d.GetOk("dual_stack_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("dual_stack_mode", sv)
				diags = append(diags, e)
			}
			obj.DualStackMode = &v2
		}
	}
	if v1, ok := d.GetOk("encode_2f_sequence"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("encode_2f_sequence", sv)
				diags = append(diags, e)
			}
			obj.Encode2fSequence = &v2
		}
	}
	if v1, ok := d.GetOk("encrypt_and_store_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("encrypt_and_store_password", sv)
				diags = append(diags, e)
			}
			obj.EncryptAndStorePassword = &v2
		}
	}
	if v1, ok := d.GetOk("force_two_factor_auth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("force_two_factor_auth", sv)
				diags = append(diags, e)
			}
			obj.ForceTwoFactorAuth = &v2
		}
	}
	if v1, ok := d.GetOk("header_x_forwarded_for"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_x_forwarded_for", sv)
				diags = append(diags, e)
			}
			obj.HeaderXForwardedFor = &v2
		}
	}
	if v1, ok := d.GetOk("hsts_include_subdomains"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hsts_include_subdomains", sv)
				diags = append(diags, e)
			}
			obj.HstsIncludeSubdomains = &v2
		}
	}
	if v1, ok := d.GetOk("http_compression"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_compression", sv)
				diags = append(diags, e)
			}
			obj.HttpCompression = &v2
		}
	}
	if v1, ok := d.GetOk("http_only_cookie"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_only_cookie", sv)
				diags = append(diags, e)
			}
			obj.HttpOnlyCookie = &v2
		}
	}
	if v1, ok := d.GetOk("http_request_body_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_request_body_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HttpRequestBodyTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("http_request_header_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_request_header_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HttpRequestHeaderTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("https_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("https_redirect", sv)
				diags = append(diags, e)
			}
			obj.HttpsRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("idle_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("idle_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IdleTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_dns_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv6DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_wins_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_wins_server1", sv)
				diags = append(diags, e)
			}
			obj.Ipv6WinsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("ipv6_wins_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv6_wins_server2", sv)
				diags = append(diags, e)
			}
			obj.Ipv6WinsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("login_attempt_limit"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_attempt_limit", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LoginAttemptLimit = &tmp
		}
	}
	if v1, ok := d.GetOk("login_block_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_block_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LoginBlockTime = &tmp
		}
	}
	if v1, ok := d.GetOk("login_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("login_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LoginTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("port_precedence"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port_precedence", sv)
				diags = append(diags, e)
			}
			obj.PortPrecedence = &v2
		}
	}
	if v1, ok := d.GetOk("reqclientcert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reqclientcert", sv)
				diags = append(diags, e)
			}
			obj.Reqclientcert = &v2
		}
	}
	if v1, ok := d.GetOk("route_source_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("route_source_interface", sv)
				diags = append(diags, e)
			}
			obj.RouteSourceInterface = &v2
		}
	}
	if v1, ok := d.GetOk("saml_redirect_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("saml_redirect_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SamlRedirectPort = &tmp
		}
	}
	if v1, ok := d.GetOk("servercert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("servercert", sv)
				diags = append(diags, e)
			}
			obj.Servercert = &v2
		}
	}
	if v, ok := d.GetOk("source_address"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("source_address", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslSettingsSourceAddress(d, v, "source_address", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SourceAddress = t
		}
	} else if d.HasChange("source_address") {
		old, new := d.GetChange("source_address")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SourceAddress = &[]models.VpnSslSettingsSourceAddress{}
		}
	}
	if v1, ok := d.GetOk("source_address_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_address_negate", sv)
				diags = append(diags, e)
			}
			obj.SourceAddressNegate = &v2
		}
	}
	if v, ok := d.GetOk("source_address6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("source_address6", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslSettingsSourceAddress6(d, v, "source_address6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SourceAddress6 = t
		}
	} else if d.HasChange("source_address6") {
		old, new := d.GetChange("source_address6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SourceAddress6 = &[]models.VpnSslSettingsSourceAddress6{}
		}
	}
	if v1, ok := d.GetOk("source_address6_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_address6_negate", sv)
				diags = append(diags, e)
			}
			obj.SourceAddress6Negate = &v2
		}
	}
	if v, ok := d.GetOk("source_interface"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("source_interface", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslSettingsSourceInterface(d, v, "source_interface", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SourceInterface = t
		}
	} else if d.HasChange("source_interface") {
		old, new := d.GetChange("source_interface")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SourceInterface = &[]models.VpnSslSettingsSourceInterface{}
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
	if v1, ok := d.GetOk("ssl_insert_empty_fragment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_insert_empty_fragment", sv)
				diags = append(diags, e)
			}
			obj.SslInsertEmptyFragment = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_max_proto_ver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_max_proto_ver", sv)
				diags = append(diags, e)
			}
			obj.SslMaxProtoVer = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_proto_ver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_proto_ver", sv)
				diags = append(diags, e)
			}
			obj.SslMinProtoVer = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("tlsv1_0"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("tlsv1_0", sv)
				diags = append(diags, e)
			}
			obj.Tlsv10 = &v2
		}
	}
	if v1, ok := d.GetOk("tlsv1_1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("tlsv1_1", sv)
				diags = append(diags, e)
			}
			obj.Tlsv11 = &v2
		}
	}
	if v1, ok := d.GetOk("tlsv1_2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("tlsv1_2", sv)
				diags = append(diags, e)
			}
			obj.Tlsv12 = &v2
		}
	}
	if v1, ok := d.GetOk("tlsv1_3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("tlsv1_3", sv)
				diags = append(diags, e)
			}
			obj.Tlsv13 = &v2
		}
	}
	if v1, ok := d.GetOk("transform_backward_slashes"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("transform_backward_slashes", sv)
				diags = append(diags, e)
			}
			obj.TransformBackwardSlashes = &v2
		}
	}
	if v1, ok := d.GetOk("tunnel_addr_assigned_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("tunnel_addr_assigned_method", sv)
				diags = append(diags, e)
			}
			obj.TunnelAddrAssignedMethod = &v2
		}
	}
	if v1, ok := d.GetOk("tunnel_connect_without_reauth"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_connect_without_reauth", sv)
				diags = append(diags, e)
			}
			obj.TunnelConnectWithoutReauth = &v2
		}
	}
	if v, ok := d.GetOk("tunnel_ip_pools"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tunnel_ip_pools", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslSettingsTunnelIpPools(d, v, "tunnel_ip_pools", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TunnelIpPools = t
		}
	} else if d.HasChange("tunnel_ip_pools") {
		old, new := d.GetChange("tunnel_ip_pools")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TunnelIpPools = &[]models.VpnSslSettingsTunnelIpPools{}
		}
	}
	if v, ok := d.GetOk("tunnel_ipv6_pools"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tunnel_ipv6_pools", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnSslSettingsTunnelIpv6Pools(d, v, "tunnel_ipv6_pools", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.TunnelIpv6Pools = t
		}
	} else if d.HasChange("tunnel_ipv6_pools") {
		old, new := d.GetChange("tunnel_ipv6_pools")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.TunnelIpv6Pools = &[]models.VpnSslSettingsTunnelIpv6Pools{}
		}
	}
	if v1, ok := d.GetOk("tunnel_user_session_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_user_session_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TunnelUserSessionTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("unsafe_legacy_renegotiation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unsafe_legacy_renegotiation", sv)
				diags = append(diags, e)
			}
			obj.UnsafeLegacyRenegotiation = &v2
		}
	}
	if v1, ok := d.GetOk("url_obscuration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url_obscuration", sv)
				diags = append(diags, e)
			}
			obj.UrlObscuration = &v2
		}
	}
	if v1, ok := d.GetOk("user_peer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_peer", sv)
				diags = append(diags, e)
			}
			obj.UserPeer = &v2
		}
	}
	if v1, ok := d.GetOk("wins_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_server1", sv)
				diags = append(diags, e)
			}
			obj.WinsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("wins_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wins_server2", sv)
				diags = append(diags, e)
			}
			obj.WinsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("x_content_type_options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("x_content_type_options", sv)
				diags = append(diags, e)
			}
			obj.XContentTypeOptions = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectVpnSslSettings(d *schema.ResourceData, sv string) (*models.VpnSslSettings, diag.Diagnostics) {
	obj := models.VpnSslSettings{}
	diags := diag.Diagnostics{}

	obj.AuthenticationRule = &[]models.VpnSslSettingsAuthenticationRule{}
	obj.SourceAddress = &[]models.VpnSslSettingsSourceAddress{}
	obj.SourceAddress6 = &[]models.VpnSslSettingsSourceAddress6{}
	obj.SourceInterface = &[]models.VpnSslSettingsSourceInterface{}
	obj.TunnelIpPools = &[]models.VpnSslSettingsTunnelIpPools{}
	obj.TunnelIpv6Pools = &[]models.VpnSslSettingsTunnelIpv6Pools{}

	return &obj, diags
}
