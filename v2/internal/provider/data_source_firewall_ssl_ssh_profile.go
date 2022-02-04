// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL/SSH protocol options.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallSslSshProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL/SSH protocol options.",

		ReadContext: dataSourceFirewallSslSshProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allowlist": {
				Type:        schema.TypeString,
				Description: "Enable/disable exempting servers by FortiGuard allowlist.",
				Computed:    true,
			},
			"block_blacklisted_certificates": {
				Type:        schema.TypeString,
				Description: "Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist.",
				Computed:    true,
			},
			"block_blocklisted_certificates": {
				Type:        schema.TypeString,
				Description: "Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist.",
				Computed:    true,
			},
			"caname": {
				Type:        schema.TypeString,
				Description: "CA certificate used by SSL Inspection.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"dot": {
				Type:        schema.TypeList,
				Description: "Configure DNS over TLS options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"ftps": {
				Type:        schema.TypeList,
				Description: "Configure FTPS options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:        schema.TypeString,
							Description: "Allow or block the invalid SSL session server certificate.",
							Computed:    true,
						},
						"min_allowed_ssl_version": {
							Type:        schema.TypeString,
							Description: "Minimum SSL version to be allowed.",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL encryption used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"https": {
				Type:        schema.TypeList,
				Description: "Configure HTTPS options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate probe failure.",
							Computed:    true,
						},
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:        schema.TypeString,
							Description: "Allow or block the invalid SSL session server certificate.",
							Computed:    true,
						},
						"min_allowed_ssl_version": {
							Type:        schema.TypeString,
							Description: "Minimum SSL version to be allowed.",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL encryption used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"imaps": {
				Type:        schema.TypeList,
				Description: "Configure IMAPS options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:        schema.TypeString,
							Description: "Allow or block the invalid SSL session server certificate.",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL encryption used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"mapi_over_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable inspection of MAPI over HTTPS.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"pop3s": {
				Type:        schema.TypeList,
				Description: "Configure POP3S options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:        schema.TypeString,
							Description: "Allow or block the invalid SSL session server certificate.",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL encryption used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"rpc_over_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable inspection of RPC over HTTPS.",
				Computed:    true,
			},
			"server_cert": {
				Type:        schema.TypeList,
				Description: "Certificate used by SSL Inspection to replace server certificate.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Certificate list.",
							Computed:    true,
						},
					},
				},
			},
			"server_cert_mode": {
				Type:        schema.TypeString,
				Description: "Re-sign or replace the server's certificate.",
				Computed:    true,
			},
			"smtps": {
				Type:        schema.TypeList,
				Description: "Configure SMTPS options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:        schema.TypeString,
							Description: "Allow or block the invalid SSL session server certificate.",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL encryption used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Description: "Configure SSH options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Level of SSL inspection.",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"ssh_algorithm": {
							Type:        schema.TypeString,
							Description: "Relative strength of encryption algorithms accepted during negotiation.",
							Computed:    true,
						},
						"ssh_tun_policy_check": {
							Type:        schema.TypeString,
							Description: "Enable/disable SSH tunnel policy check.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Configure protocol inspection status.",
							Computed:    true,
						},
						"unsupported_version": {
							Type:        schema.TypeString,
							Description: "Action based on SSH version being unsupported.",
							Computed:    true,
						},
					},
				},
			},
			"ssl": {
				Type:        schema.TypeList,
				Description: "Configure SSL options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate probe failure.",
							Computed:    true,
						},
						"cert_validation_failure": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation failure.",
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:        schema.TypeString,
							Description: "Action based on certificate validation timeout.",
							Computed:    true,
						},
						"client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request.",
							Computed:    true,
						},
						"client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate.",
							Computed:    true,
						},
						"expired_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is expired.",
							Computed:    true,
						},
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Level of SSL inspection.",
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:        schema.TypeString,
							Description: "Allow or block the invalid SSL session server certificate.",
							Computed:    true,
						},
						"min_allowed_ssl_version": {
							Type:        schema.TypeString,
							Description: "Minimum SSL version to be allowed.",
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is revoked.",
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:        schema.TypeString,
							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL encryption used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL cipher used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL negotiation used being unsupported.",
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:        schema.TypeString,
							Description: "Action based on the SSL version used being unsupported.",
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:        schema.TypeString,
							Description: "Action based on server certificate is not issued by a trusted CA.",
							Computed:    true,
						},
					},
				},
			},
			"ssl_anomalies_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging SSL anomalies.",
				Computed:    true,
			},
			"ssl_anomaly_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging SSL anomalies.",
				Computed:    true,
			},
			"ssl_exempt": {
				Type:        schema.TypeList,
				Description: "Servers to exempt from SSL inspection.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Description: "IPv4 address object.",
							Computed:    true,
						},
						"address6": {
							Type:        schema.TypeString,
							Description: "IPv6 address object.",
							Computed:    true,
						},
						"fortiguard_category": {
							Type:        schema.TypeInt,
							Description: "FortiGuard category ID.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID number.",
							Computed:    true,
						},
						"regex": {
							Type:        schema.TypeString,
							Description: "Exempt servers by regular expression.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Type of address object (IPv4 or IPv6) or FortiGuard category.",
							Computed:    true,
						},
						"wildcard_fqdn": {
							Type:        schema.TypeString,
							Description: "Exempt servers by wildcard FQDN.",
							Computed:    true,
						},
					},
				},
			},
			"ssl_exemption_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging SSL exemptions.",
				Computed:    true,
			},
			"ssl_exemptions_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging SSL exemptions.",
				Computed:    true,
			},
			"ssl_handshake_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of TLS handshakes.",
				Computed:    true,
			},
			"ssl_negotiation_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging SSL negotiation.",
				Computed:    true,
			},
			"ssl_server": {
				Type:        schema.TypeList,
				Description: "SSL server settings used for client certificate request.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ftps_client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request during the FTPS handshake.",
							Computed:    true,
						},
						"ftps_client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate during the FTPS handshake.",
							Computed:    true,
						},
						"https_client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request during the HTTPS handshake.",
							Computed:    true,
						},
						"https_client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate during the HTTPS handshake.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "SSL server ID.",
							Computed:    true,
						},
						"imaps_client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request during the IMAPS handshake.",
							Computed:    true,
						},
						"imaps_client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate during the IMAPS handshake.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "IPv4 address of the SSL server.",
							Computed:    true,
						},
						"pop3s_client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request during the POP3S handshake.",
							Computed:    true,
						},
						"pop3s_client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate during the POP3S handshake.",
							Computed:    true,
						},
						"smtps_client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request during the SMTPS handshake.",
							Computed:    true,
						},
						"smtps_client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate during the SMTPS handshake.",
							Computed:    true,
						},
						"ssl_other_client_cert_request": {
							Type:        schema.TypeString,
							Description: "Action based on client certificate request during an SSL protocol handshake.",
							Computed:    true,
						},
						"ssl_other_client_certificate": {
							Type:        schema.TypeString,
							Description: "Action based on received client certificate during an SSL protocol handshake.",
							Computed:    true,
						},
					},
				},
			},
			"ssl_server_cert_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging of server certificate information.",
				Computed:    true,
			},
			"supported_alpn": {
				Type:        schema.TypeString,
				Description: "Configure ALPN option.",
				Computed:    true,
			},
			"untrusted_caname": {
				Type:        schema.TypeString,
				Description: "Untrusted CA certificate used by SSL Inspection.",
				Computed:    true,
			},
			"use_ssl_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable the use of SSL server table for SSL offloading.",
				Computed:    true,
			},
			"whitelist": {
				Type:        schema.TypeString,
				Description: "Enable/disable exempting servers by FortiGuard whitelist.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallSslSshProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSslSshProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSslSshProfile dataSource: %v", err)
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

	diags := refreshObjectFirewallSslSshProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
