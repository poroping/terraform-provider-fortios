// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL-VPN.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnSslSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL-VPN.",

		ReadContext: dataSourceVpnSslSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"algorithm": {
				Type:        schema.TypeString,
				Description: "Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any.",
				Computed:    true,
			},
			"auth_session_check_source_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable checking of source IP for authentication session.",
				Computed:    true,
			},
			"auth_timeout": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).",
				Computed:    true,
			},
			"authentication_rule": {
				Type:        schema.TypeList,
				Description: "Authentication rule for SSL-VPN.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth": {
							Type:        schema.TypeString,
							Description: "SSL-VPN authentication method restriction.",
							Computed:    true,
						},
						"cipher": {
							Type:        schema.TypeString,
							Description: "SSL-VPN cipher strength.",
							Computed:    true,
						},
						"client_cert": {
							Type:        schema.TypeString,
							Description: "Enable/disable SSL-VPN client certificate restrictive.",
							Computed:    true,
						},
						"groups": {
							Type:        schema.TypeList,
							Description: "User groups.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Group name.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID (0 - 4294967295).",
							Computed:    true,
						},
						"portal": {
							Type:        schema.TypeString,
							Description: "SSL-VPN portal.",
							Computed:    true,
						},
						"realm": {
							Type:        schema.TypeString,
							Description: "SSL-VPN realm.",
							Computed:    true,
						},
						"source_address": {
							Type:        schema.TypeList,
							Description: "Source address of incoming traffic.",
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
						"source_address_negate": {
							Type:        schema.TypeString,
							Description: "Enable/disable negated source address match.",
							Computed:    true,
						},
						"source_address6": {
							Type:        schema.TypeList,
							Description: "IPv6 source address of incoming traffic.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "IPv6 address name.",
										Computed:    true,
									},
								},
							},
						},
						"source_address6_negate": {
							Type:        schema.TypeString,
							Description: "Enable/disable negated source IPv6 address match.",
							Computed:    true,
						},
						"source_interface": {
							Type:        schema.TypeList,
							Description: "SSL-VPN source interface of incoming traffic.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Interface name.",
										Computed:    true,
									},
								},
							},
						},
						"user_peer": {
							Type:        schema.TypeString,
							Description: "Name of user peer.",
							Computed:    true,
						},
						"users": {
							Type:        schema.TypeList,
							Description: "User name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "User name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"auto_tunnel_static_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable to auto-create static routes for the SSL-VPN tunnel IP addresses.",
				Computed:    true,
			},
			"banned_cipher": {
				Type:        schema.TypeString,
				Description: "Select one or more cipher technologies that cannot be used in SSL-VPN negotiations. Only applies to TLS 1.2 and below.",
				Computed:    true,
			},
			"check_referer": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification of referer field in HTTP request header.",
				Computed:    true,
			},
			"ciphersuite": {
				Type:        schema.TypeString,
				Description: "Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below.",
				Computed:    true,
			},
			"client_sigalgs": {
				Type:        schema.TypeString,
				Description: "Set signature algorithms related to client authentication. Affects TLS version <= 1.2 only.",
				Computed:    true,
			},
			"default_portal": {
				Type:        schema.TypeString,
				Description: "Default SSL-VPN portal.",
				Computed:    true,
			},
			"deflate_compression_level": {
				Type:        schema.TypeInt,
				Description: "Compression level (0~9).",
				Computed:    true,
			},
			"deflate_min_data_size": {
				Type:        schema.TypeInt,
				Description: "Minimum amount of data that triggers compression (200 - 65535 bytes).",
				Computed:    true,
			},
			"dns_server1": {
				Type:        schema.TypeString,
				Description: "DNS server 1.",
				Computed:    true,
			},
			"dns_server2": {
				Type:        schema.TypeString,
				Description: "DNS server 2.",
				Computed:    true,
			},
			"dns_suffix": {
				Type:        schema.TypeString,
				Description: "DNS suffix used for SSL-VPN clients.",
				Computed:    true,
			},
			"dtls_hello_timeout": {
				Type:        schema.TypeInt,
				Description: "SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).",
				Computed:    true,
			},
			"dtls_max_proto_ver": {
				Type:        schema.TypeString,
				Description: "DTLS maximum protocol version.",
				Computed:    true,
			},
			"dtls_min_proto_ver": {
				Type:        schema.TypeString,
				Description: "DTLS minimum protocol version.",
				Computed:    true,
			},
			"dtls_tunnel": {
				Type:        schema.TypeString,
				Description: "Enable/disable DTLS to prevent eavesdropping, tampering, or message forgery.",
				Computed:    true,
			},
			"dual_stack_mode": {
				Type:        schema.TypeString,
				Description: "Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal.",
				Computed:    true,
			},
			"encode_2f_sequence": {
				Type:        schema.TypeString,
				Description: "Encode \\2F sequence to forward slash in URLs.",
				Computed:    true,
			},
			"encrypt_and_store_password": {
				Type:        schema.TypeString,
				Description: "Encrypt and store user passwords for SSL-VPN web sessions.",
				Computed:    true,
			},
			"force_two_factor_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable only PKI users with two-factor authentication for SSL-VPNs.",
				Computed:    true,
			},
			"header_x_forwarded_for": {
				Type:        schema.TypeString,
				Description: "Forward the same, add, or remove HTTP header.",
				Computed:    true,
			},
			"hsts_include_subdomains": {
				Type:        schema.TypeString,
				Description: "Add HSTS includeSubDomains response header.",
				Computed:    true,
			},
			"http_compression": {
				Type:        schema.TypeString,
				Description: "Enable/disable to allow HTTP compression over SSL-VPN tunnels.",
				Computed:    true,
			},
			"http_only_cookie": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN support for HttpOnly cookies.",
				Computed:    true,
			},
			"http_request_body_timeout": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).",
				Computed:    true,
			},
			"http_request_header_timeout": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).",
				Computed:    true,
			},
			"https_redirect": {
				Type:        schema.TypeString,
				Description: "Enable/disable redirect of port 80 to SSL-VPN port.",
				Computed:    true,
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN disconnects if idle for specified time in seconds.",
				Computed:    true,
			},
			"ipv6_dns_server1": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 1.",
				Computed:    true,
			},
			"ipv6_dns_server2": {
				Type:        schema.TypeString,
				Description: "IPv6 DNS server 2.",
				Computed:    true,
			},
			"ipv6_wins_server1": {
				Type:        schema.TypeString,
				Description: "IPv6 WINS server 1.",
				Computed:    true,
			},
			"ipv6_wins_server2": {
				Type:        schema.TypeString,
				Description: "IPv6 WINS server 2.",
				Computed:    true,
			},
			"login_attempt_limit": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).",
				Computed:    true,
			},
			"login_block_time": {
				Type:        schema.TypeInt,
				Description: "Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).",
				Computed:    true,
			},
			"login_timeout": {
				Type:        schema.TypeInt,
				Description: "SSLVPN maximum login timeout (10 - 180 sec, default = 30).",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "SSL-VPN access port (1 - 65535).",
				Computed:    true,
			},
			"port_precedence": {
				Type:        schema.TypeString,
				Description: "Enable/disable, Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface.",
				Computed:    true,
			},
			"reqclientcert": {
				Type:        schema.TypeString,
				Description: "Enable/disable to require client certificates for all SSL-VPN users.",
				Computed:    true,
			},
			"route_source_interface": {
				Type:        schema.TypeString,
				Description: "Enable to allow SSL-VPN sessions to bypass routing and bind to the incoming interface.",
				Computed:    true,
			},
			"saml_redirect_port": {
				Type:        schema.TypeInt,
				Description: "SAML local redirect port in the machine running FCT (0 - 65535). 0 is to disable redirection on FGT side.",
				Computed:    true,
			},
			"servercert": {
				Type:        schema.TypeString,
				Description: "Name of the server certificate to be used for SSL-VPNs.",
				Computed:    true,
			},
			"source_address": {
				Type:        schema.TypeList,
				Description: "Source address of incoming traffic.",
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
			"source_address_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negated source address match.",
				Computed:    true,
			},
			"source_address6": {
				Type:        schema.TypeList,
				Description: "IPv6 source address of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IPv6 address name.",
							Computed:    true,
						},
					},
				},
			},
			"source_address6_negate": {
				Type:        schema.TypeString,
				Description: "Enable/disable negated source IPv6 address match.",
				Computed:    true,
			},
			"source_interface": {
				Type:        schema.TypeList,
				Description: "SSL-VPN source interface of incoming traffic.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"ssl_client_renegotiation": {
				Type:        schema.TypeString,
				Description: "Enable/disable to allow client renegotiation by the server if the tunnel goes down.",
				Computed:    true,
			},
			"ssl_insert_empty_fragment": {
				Type:        schema.TypeString,
				Description: "Enable/disable insertion of empty fragment.",
				Computed:    true,
			},
			"ssl_max_proto_ver": {
				Type:        schema.TypeString,
				Description: "SSL maximum protocol version.",
				Computed:    true,
			},
			"ssl_min_proto_ver": {
				Type:        schema.TypeString,
				Description: "SSL minimum protocol version.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN.",
				Computed:    true,
			},
			"tlsv1_0": {
				Type:        schema.TypeString,
				Description: "tlsv1-0",
				Computed:    true,
			},
			"tlsv1_1": {
				Type:        schema.TypeString,
				Description: "tlsv1-1",
				Computed:    true,
			},
			"tlsv1_2": {
				Type:        schema.TypeString,
				Description: "tlsv1-2",
				Computed:    true,
			},
			"tlsv1_3": {
				Type:        schema.TypeString,
				Description: "tlsv1-3",
				Computed:    true,
			},
			"transform_backward_slashes": {
				Type:        schema.TypeString,
				Description: "Transform backward slashes to forward slashes in URLs.",
				Computed:    true,
			},
			"tunnel_addr_assigned_method": {
				Type:        schema.TypeString,
				Description: "Method used for assigning address for tunnel.",
				Computed:    true,
			},
			"tunnel_connect_without_reauth": {
				Type:        schema.TypeString,
				Description: "Enable/disable tunnel connection without re-authorization if previous connection dropped.",
				Computed:    true,
			},
			"tunnel_ip_pools": {
				Type:        schema.TypeList,
				Description: "Names of the IPv4 IP Pool firewall objects that define the IP addresses reserved for remote clients.",
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
			"tunnel_ipv6_pools": {
				Type:        schema.TypeList,
				Description: "Names of the IPv6 IP Pool firewall objects that define the IP addresses reserved for remote clients.",
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
			"tunnel_user_session_timeout": {
				Type:        schema.TypeInt,
				Description: "Time out value to clean up user session after tunnel connection is dropped (1 - 255 sec, default=30).",
				Computed:    true,
			},
			"unsafe_legacy_renegotiation": {
				Type:        schema.TypeString,
				Description: "Enable/disable unsafe legacy re-negotiation.",
				Computed:    true,
			},
			"url_obscuration": {
				Type:        schema.TypeString,
				Description: "Enable/disable to obscure the host name of the URL of the web browser display.",
				Computed:    true,
			},
			"user_peer": {
				Type:        schema.TypeString,
				Description: "Name of user peer.",
				Computed:    true,
			},
			"wins_server1": {
				Type:        schema.TypeString,
				Description: "WINS server 1.",
				Computed:    true,
			},
			"wins_server2": {
				Type:        schema.TypeString,
				Description: "WINS server 2.",
				Computed:    true,
			},
			"x_content_type_options": {
				Type:        schema.TypeString,
				Description: "Add HTTP X-Content-Type-Options header.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnSslSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "VpnSslSettings"

	o, err := c.Cmdb.ReadVpnSslSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslSettings dataSource: %v", err)
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

	diags := refreshObjectVpnSslSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
