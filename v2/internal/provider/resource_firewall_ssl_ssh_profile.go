// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL/SSH protocol options.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallSslSshProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL/SSH protocol options.",

		CreateContext: resourceFirewallSslSshProfileCreate,
		ReadContext:   resourceFirewallSslSshProfileRead,
		UpdateContext: resourceFirewallSslSshProfileUpdate,
		DeleteContext: resourceFirewallSslSshProfileDelete,

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
			"allowlist": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable exempting servers by FortiGuard allowlist.",
				Optional:    true,
				Computed:    true,
			},
			"block_blacklisted_certificates": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist.",
				Optional:    true,
				Computed:    true,
			},
			"block_blocklisted_certificates": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist.",
				Optional:    true,
				Computed:    true,
			},
			"caname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "CA certificate used by SSL Inspection.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"dot": {
				Type:        schema.TypeList,
				Description: "Configure DNS over TLS options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ftps": {
				Type:        schema.TypeList,
				Description: "Configure FTPS options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Allow or block the invalid SSL session server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"min_allowed_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Minimum SSL version to be allowed.",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on the SSL encryption used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"https": {
				Type:        schema.TypeList,
				Description: "Configure HTTPS options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on certificate probe failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Allow or block the invalid SSL session server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"min_allowed_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Minimum SSL version to be allowed.",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Optional:    true,
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "certificate-inspection", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on the SSL encryption used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"imaps": {
				Type:        schema.TypeList,
				Description: "Configure IMAPS options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Allow or block the invalid SSL session server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Optional:    true,
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on the SSL encryption used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mapi_over_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable inspection of MAPI over HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"pop3s": {
				Type:        schema.TypeList,
				Description: "Configure POP3S options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Allow or block the invalid SSL session server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Optional:    true,
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on the SSL encryption used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"rpc_over_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable inspection of RPC over HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"server_cert": {
				Type:        schema.TypeList,
				Description: "Certificate used by SSL Inspection to replace server certificate.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Certificate list.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_cert_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"re-sign", "replace"}, false),

				Description: "Re-sign or replace the server's certificate.",
				Optional:    true,
				Computed:    true,
			},
			"smtps": {
				Type:        schema.TypeList,
				Description: "Configure SMTPS options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Allow or block the invalid SSL session server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Optional:    true,
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on the SSL encryption used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Description: "Configure SSH options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Level of SSL inspection.",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to use for scanning (1 - 65535, default = 443).",
							Optional:    true,
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"ssh_algorithm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"compatible", "high-encryption"}, false),

							Description: "Relative strength of encryption algorithms accepted during negotiation.",
							Optional:    true,
							Computed:    true,
						},
						"ssh_tun_policy_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SSH tunnel policy check.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "deep-inspection"}, false),

							Description: "Configure protocol inspection status.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "block"}, false),

							Description: "Action based on SSH version being unsupported.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl": {
				Type:        schema.TypeList,
				Description: "Configure SSL options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_probe_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on certificate probe failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_failure": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation failure.",
							Optional:    true,
							Computed:    true,
						},
						"cert_validation_timeout": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on certificate validation timeout.",
							Optional:    true,
							Computed:    true,
						},
						"client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request.",
							Optional:    true,
							Computed:    true,
						},
						"client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate.",
							Optional:    true,
							Computed:    true,
						},
						"expired_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is expired.",
							Optional:    true,
							Computed:    true,
						},
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "certificate-inspection", "deep-inspection"}, false),

							Description: "Level of SSL inspection.",
							Optional:    true,
							Computed:    true,
						},
						"invalid_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Allow or block the invalid SSL session server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"min_allowed_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Minimum SSL version to be allowed.",
							Optional:    true,
							Computed:    true,
						},
						"revoked_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is revoked.",
							Optional:    true,
							Computed:    true,
						},
						"sni_server_cert_check": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "strict", "disable"}, false),

							Description: "Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on the SSL encryption used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_cipher": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL cipher used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_negotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL negotiation used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"unsupported_ssl_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block"}, false),

							Description: "Action based on the SSL version used being unsupported.",
							Optional:    true,
							Computed:    true,
						},
						"untrusted_server_cert": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "block", "ignore"}, false),

							Description: "Action based on server certificate is not issued by a trusted CA.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl_anomalies_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging SSL anomalies.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_anomaly_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging SSL anomalies.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_exempt": {
				Type:        schema.TypeList,
				Description: "Servers to exempt from SSL inspection.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPv4 address object.",
							Optional:    true,
							Computed:    true,
						},
						"address6": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPv6 address object.",
							Optional:    true,
							Computed:    true,
						},
						"fortiguard_category": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "FortiGuard category ID.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 512),

							Description: "ID number.",
							Optional:    true,
							Computed:    true,
						},
						"regex": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Exempt servers by regular expression.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"fortiguard-category", "address", "address6", "wildcard-fqdn", "regex"}, false),

							Description: "Type of address object (IPv4 or IPv6) or FortiGuard category.",
							Optional:    true,
							Computed:    true,
						},
						"wildcard_fqdn": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Exempt servers by wildcard FQDN.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl_exemption_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging SSL exemptions.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_exemptions_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging SSL exemptions.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_handshake_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging of TLS handshakes.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_negotiation_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging SSL negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_server": {
				Type:        schema.TypeList,
				Description: "SSL server settings used for client certificate request.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ftps_client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request during the FTPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"ftps_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate during the FTPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"https_client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request during the HTTPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"https_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate during the HTTPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "SSL server ID.",
							Optional:    true,
							Computed:    true,
						},
						"imaps_client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request during the IMAPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"imaps_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate during the IMAPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IPv4 address of the SSL server.",
							Optional:    true,
							Computed:    true,
						},
						"pop3s_client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request during the POP3S handshake.",
							Optional:    true,
							Computed:    true,
						},
						"pop3s_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate during the POP3S handshake.",
							Optional:    true,
							Computed:    true,
						},
						"smtps_client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request during the SMTPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"smtps_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate during the SMTPS handshake.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_other_client_cert_request": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on client certificate request during an SSL protocol handshake.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_other_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "inspect", "block"}, false),

							Description: "Action based on received client certificate during an SSL protocol handshake.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl_server_cert_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging of server certificate information.",
				Optional:    true,
				Computed:    true,
			},
			"supported_alpn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"http1-1", "http2", "all", "none"}, false),

				Description: "Configure ALPN option.",
				Optional:    true,
				Computed:    true,
			},
			"untrusted_caname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Untrusted CA certificate used by SSL Inspection.",
				Optional:    true,
				Computed:    true,
			},
			"use_ssl_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable the use of SSL server table for SSL offloading.",
				Optional:    true,
				Computed:    true,
			},
			"whitelist": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable exempting servers by FortiGuard whitelist.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSslSshProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallSslSshProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallSslSshProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSslSshProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSslSshProfile")
	}

	return resourceFirewallSslSshProfileRead(ctx, d, meta)
}

func resourceFirewallSslSshProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSslSshProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSslSshProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSslSshProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSslSshProfile")
	}

	return resourceFirewallSslSshProfileRead(ctx, d, meta)
}

func resourceFirewallSslSshProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallSslSshProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSslSshProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSslSshProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSslSshProfile resource: %v", err)
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

	diags := refreshObjectFirewallSslSshProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallSslSshProfileDot(v *[]models.FirewallSslSshProfileDot, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileFtps(v *[]models.FirewallSslSshProfileFtps, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertRequest; tmp != nil {
				v["client_cert_request"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.InvalidServerCert; tmp != nil {
				v["invalid_server_cert"] = *tmp
			}

			if tmp := cfg.MinAllowedSslVersion; tmp != nil {
				v["min_allowed_ssl_version"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedSsl; tmp != nil {
				v["unsupported_ssl"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileHttps(v *[]models.FirewallSslSshProfileHttps, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertProbeFailure; tmp != nil {
				v["cert_probe_failure"] = *tmp
			}

			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertRequest; tmp != nil {
				v["client_cert_request"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.InvalidServerCert; tmp != nil {
				v["invalid_server_cert"] = *tmp
			}

			if tmp := cfg.MinAllowedSslVersion; tmp != nil {
				v["min_allowed_ssl_version"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedSsl; tmp != nil {
				v["unsupported_ssl"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileImaps(v *[]models.FirewallSslSshProfileImaps, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertRequest; tmp != nil {
				v["client_cert_request"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.InvalidServerCert; tmp != nil {
				v["invalid_server_cert"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedSsl; tmp != nil {
				v["unsupported_ssl"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfilePop3s(v *[]models.FirewallSslSshProfilePop3s, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertRequest; tmp != nil {
				v["client_cert_request"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.InvalidServerCert; tmp != nil {
				v["invalid_server_cert"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedSsl; tmp != nil {
				v["unsupported_ssl"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileServerCert(v *[]models.FirewallSslSshProfileServerCert, sort bool) interface{} {
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

func flattenFirewallSslSshProfileSmtps(v *[]models.FirewallSslSshProfileSmtps, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertRequest; tmp != nil {
				v["client_cert_request"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.InvalidServerCert; tmp != nil {
				v["invalid_server_cert"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedSsl; tmp != nil {
				v["unsupported_ssl"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileSsh(v *[]models.FirewallSslSshProfileSsh, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.SshAlgorithm; tmp != nil {
				v["ssh_algorithm"] = *tmp
			}

			if tmp := cfg.SshTunPolicyCheck; tmp != nil {
				v["ssh_tun_policy_check"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UnsupportedVersion; tmp != nil {
				v["unsupported_version"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileSsl(v *[]models.FirewallSslSshProfileSsl, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CertProbeFailure; tmp != nil {
				v["cert_probe_failure"] = *tmp
			}

			if tmp := cfg.CertValidationFailure; tmp != nil {
				v["cert_validation_failure"] = *tmp
			}

			if tmp := cfg.CertValidationTimeout; tmp != nil {
				v["cert_validation_timeout"] = *tmp
			}

			if tmp := cfg.ClientCertRequest; tmp != nil {
				v["client_cert_request"] = *tmp
			}

			if tmp := cfg.ClientCertificate; tmp != nil {
				v["client_certificate"] = *tmp
			}

			if tmp := cfg.ExpiredServerCert; tmp != nil {
				v["expired_server_cert"] = *tmp
			}

			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.InvalidServerCert; tmp != nil {
				v["invalid_server_cert"] = *tmp
			}

			if tmp := cfg.MinAllowedSslVersion; tmp != nil {
				v["min_allowed_ssl_version"] = *tmp
			}

			if tmp := cfg.RevokedServerCert; tmp != nil {
				v["revoked_server_cert"] = *tmp
			}

			if tmp := cfg.SniServerCertCheck; tmp != nil {
				v["sni_server_cert_check"] = *tmp
			}

			if tmp := cfg.UnsupportedSsl; tmp != nil {
				v["unsupported_ssl"] = *tmp
			}

			if tmp := cfg.UnsupportedSslCipher; tmp != nil {
				v["unsupported_ssl_cipher"] = *tmp
			}

			if tmp := cfg.UnsupportedSslNegotiation; tmp != nil {
				v["unsupported_ssl_negotiation"] = *tmp
			}

			if tmp := cfg.UnsupportedSslVersion; tmp != nil {
				v["unsupported_ssl_version"] = *tmp
			}

			if tmp := cfg.UntrustedServerCert; tmp != nil {
				v["untrusted_server_cert"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallSslSshProfileSslExempt(v *[]models.FirewallSslSshProfileSslExempt, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Address; tmp != nil {
				v["address"] = *tmp
			}

			if tmp := cfg.Address6; tmp != nil {
				v["address6"] = *tmp
			}

			if tmp := cfg.FortiguardCategory; tmp != nil {
				v["fortiguard_category"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Regex; tmp != nil {
				v["regex"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.WildcardFqdn; tmp != nil {
				v["wildcard_fqdn"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallSslSshProfileSslServer(v *[]models.FirewallSslSshProfileSslServer, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.FtpsClientCertRequest; tmp != nil {
				v["ftps_client_cert_request"] = *tmp
			}

			if tmp := cfg.FtpsClientCertificate; tmp != nil {
				v["ftps_client_certificate"] = *tmp
			}

			if tmp := cfg.HttpsClientCertRequest; tmp != nil {
				v["https_client_cert_request"] = *tmp
			}

			if tmp := cfg.HttpsClientCertificate; tmp != nil {
				v["https_client_certificate"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.ImapsClientCertRequest; tmp != nil {
				v["imaps_client_cert_request"] = *tmp
			}

			if tmp := cfg.ImapsClientCertificate; tmp != nil {
				v["imaps_client_certificate"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Pop3sClientCertRequest; tmp != nil {
				v["pop3s_client_cert_request"] = *tmp
			}

			if tmp := cfg.Pop3sClientCertificate; tmp != nil {
				v["pop3s_client_certificate"] = *tmp
			}

			if tmp := cfg.SmtpsClientCertRequest; tmp != nil {
				v["smtps_client_cert_request"] = *tmp
			}

			if tmp := cfg.SmtpsClientCertificate; tmp != nil {
				v["smtps_client_certificate"] = *tmp
			}

			if tmp := cfg.SslOtherClientCertRequest; tmp != nil {
				v["ssl_other_client_cert_request"] = *tmp
			}

			if tmp := cfg.SslOtherClientCertificate; tmp != nil {
				v["ssl_other_client_certificate"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectFirewallSslSshProfile(d *schema.ResourceData, o *models.FirewallSslSshProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Allowlist != nil {
		v := *o.Allowlist

		if err = d.Set("allowlist", v); err != nil {
			return diag.Errorf("error reading allowlist: %v", err)
		}
	}

	if o.BlockBlacklistedCertificates != nil {
		v := *o.BlockBlacklistedCertificates

		if err = d.Set("block_blacklisted_certificates", v); err != nil {
			return diag.Errorf("error reading block_blacklisted_certificates: %v", err)
		}
	}

	if o.BlockBlocklistedCertificates != nil {
		v := *o.BlockBlocklistedCertificates

		if err = d.Set("block_blocklisted_certificates", v); err != nil {
			return diag.Errorf("error reading block_blocklisted_certificates: %v", err)
		}
	}

	if o.Caname != nil {
		v := *o.Caname

		if err = d.Set("caname", v); err != nil {
			return diag.Errorf("error reading caname: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Dot != nil {
		if err = d.Set("dot", flattenFirewallSslSshProfileDot(o.Dot, sort)); err != nil {
			return diag.Errorf("error reading dot: %v", err)
		}
	}

	if o.Ftps != nil {
		if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o.Ftps, sort)); err != nil {
			return diag.Errorf("error reading ftps: %v", err)
		}
	}

	if o.Https != nil {
		if err = d.Set("https", flattenFirewallSslSshProfileHttps(o.Https, sort)); err != nil {
			return diag.Errorf("error reading https: %v", err)
		}
	}

	if o.Imaps != nil {
		if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o.Imaps, sort)); err != nil {
			return diag.Errorf("error reading imaps: %v", err)
		}
	}

	if o.MapiOverHttps != nil {
		v := *o.MapiOverHttps

		if err = d.Set("mapi_over_https", v); err != nil {
			return diag.Errorf("error reading mapi_over_https: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Pop3s != nil {
		if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3s(o.Pop3s, sort)); err != nil {
			return diag.Errorf("error reading pop3s: %v", err)
		}
	}

	if o.RpcOverHttps != nil {
		v := *o.RpcOverHttps

		if err = d.Set("rpc_over_https", v); err != nil {
			return diag.Errorf("error reading rpc_over_https: %v", err)
		}
	}

	if o.ServerCert != nil {
		if err = d.Set("server_cert", flattenFirewallSslSshProfileServerCert(o.ServerCert, sort)); err != nil {
			return diag.Errorf("error reading server_cert: %v", err)
		}
	}

	if o.ServerCertMode != nil {
		v := *o.ServerCertMode

		if err = d.Set("server_cert_mode", v); err != nil {
			return diag.Errorf("error reading server_cert_mode: %v", err)
		}
	}

	if o.Smtps != nil {
		if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o.Smtps, sort)); err != nil {
			return diag.Errorf("error reading smtps: %v", err)
		}
	}

	if o.Ssh != nil {
		if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o.Ssh, sort)); err != nil {
			return diag.Errorf("error reading ssh: %v", err)
		}
	}

	if o.Ssl != nil {
		if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o.Ssl, sort)); err != nil {
			return diag.Errorf("error reading ssl: %v", err)
		}
	}

	if o.SslAnomaliesLog != nil {
		v := *o.SslAnomaliesLog

		if err = d.Set("ssl_anomalies_log", v); err != nil {
			return diag.Errorf("error reading ssl_anomalies_log: %v", err)
		}
	}

	if o.SslAnomalyLog != nil {
		v := *o.SslAnomalyLog

		if err = d.Set("ssl_anomaly_log", v); err != nil {
			return diag.Errorf("error reading ssl_anomaly_log: %v", err)
		}
	}

	if o.SslExempt != nil {
		if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o.SslExempt, sort)); err != nil {
			return diag.Errorf("error reading ssl_exempt: %v", err)
		}
	}

	if o.SslExemptionLog != nil {
		v := *o.SslExemptionLog

		if err = d.Set("ssl_exemption_log", v); err != nil {
			return diag.Errorf("error reading ssl_exemption_log: %v", err)
		}
	}

	if o.SslExemptionsLog != nil {
		v := *o.SslExemptionsLog

		if err = d.Set("ssl_exemptions_log", v); err != nil {
			return diag.Errorf("error reading ssl_exemptions_log: %v", err)
		}
	}

	if o.SslHandshakeLog != nil {
		v := *o.SslHandshakeLog

		if err = d.Set("ssl_handshake_log", v); err != nil {
			return diag.Errorf("error reading ssl_handshake_log: %v", err)
		}
	}

	if o.SslNegotiationLog != nil {
		v := *o.SslNegotiationLog

		if err = d.Set("ssl_negotiation_log", v); err != nil {
			return diag.Errorf("error reading ssl_negotiation_log: %v", err)
		}
	}

	if o.SslServer != nil {
		if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o.SslServer, sort)); err != nil {
			return diag.Errorf("error reading ssl_server: %v", err)
		}
	}

	if o.SslServerCertLog != nil {
		v := *o.SslServerCertLog

		if err = d.Set("ssl_server_cert_log", v); err != nil {
			return diag.Errorf("error reading ssl_server_cert_log: %v", err)
		}
	}

	if o.SupportedAlpn != nil {
		v := *o.SupportedAlpn

		if err = d.Set("supported_alpn", v); err != nil {
			return diag.Errorf("error reading supported_alpn: %v", err)
		}
	}

	if o.UntrustedCaname != nil {
		v := *o.UntrustedCaname

		if err = d.Set("untrusted_caname", v); err != nil {
			return diag.Errorf("error reading untrusted_caname: %v", err)
		}
	}

	if o.UseSslServer != nil {
		v := *o.UseSslServer

		if err = d.Set("use_ssl_server", v); err != nil {
			return diag.Errorf("error reading use_ssl_server: %v", err)
		}
	}

	if o.Whitelist != nil {
		v := *o.Whitelist

		if err = d.Set("whitelist", v); err != nil {
			return diag.Errorf("error reading whitelist: %v", err)
		}
	}

	return nil
}

func expandFirewallSslSshProfileDot(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileDot, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileDot

	for i := range l {
		tmp := models.FirewallSslSshProfileDot{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileFtps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileFtps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileFtps

	for i := range l {
		tmp := models.FirewallSslSshProfileFtps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invalid_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InvalidServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.min_allowed_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MinAllowedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSsl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileHttps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileHttps

	for i := range l {
		tmp := models.FirewallSslSshProfileHttps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_probe_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertProbeFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invalid_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InvalidServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.min_allowed_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MinAllowedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSsl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileImaps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileImaps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileImaps

	for i := range l {
		tmp := models.FirewallSslSshProfileImaps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invalid_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InvalidServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSsl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfilePop3s(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfilePop3s, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfilePop3s

	for i := range l {
		tmp := models.FirewallSslSshProfilePop3s{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invalid_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InvalidServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSsl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileServerCert, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileServerCert

	for i := range l {
		tmp := models.FirewallSslSshProfileServerCert{}
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

func expandFirewallSslSshProfileSmtps(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileSmtps, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileSmtps

	for i := range l {
		tmp := models.FirewallSslSshProfileSmtps{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invalid_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InvalidServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSsl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileSsh, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileSsh

	for i := range l {
		tmp := models.FirewallSslSshProfileSsh{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssh_algorithm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SshAlgorithm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssh_tun_policy_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SshTunPolicyCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedVersion = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileSsl, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileSsl

	for i := range l {
		tmp := models.FirewallSslSshProfileSsl{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.cert_probe_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertProbeFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_failure", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationFailure = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cert_validation_timeout", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CertValidationTimeout = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expired_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ExpiredServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invalid_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InvalidServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.min_allowed_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MinAllowedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.revoked_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RevokedServerCert = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sni_server_cert_check", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SniServerCertCheck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSsl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_cipher", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslCipher = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_negotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslNegotiation = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unsupported_ssl_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnsupportedSslVersion = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.untrusted_server_cert", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UntrustedServerCert = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileSslExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileSslExempt, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileSslExempt

	for i := range l {
		tmp := models.FirewallSslSshProfileSslExempt{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.address", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.address6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Address6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortiguard_category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FortiguardCategory = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.regex", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Regex = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.wildcard_fqdn", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.WildcardFqdn = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallSslSshProfileSslServer(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallSslSshProfileSslServer, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallSslSshProfileSslServer

	for i := range l {
		tmp := models.FirewallSslSshProfileSslServer{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ftps_client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FtpsClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ftps_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FtpsClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.https_client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpsClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.https_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HttpsClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.imaps_client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImapsClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.imaps_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ImapsClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pop3s_client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pop3sClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pop3s_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pop3sClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.smtps_client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SmtpsClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.smtps_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SmtpsClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_other_client_cert_request", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOtherClientCertRequest = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_other_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOtherClientCertificate = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallSslSshProfile(d *schema.ResourceData, sv string) (*models.FirewallSslSshProfile, diag.Diagnostics) {
	obj := models.FirewallSslSshProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allowlist"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("allowlist", sv)
				diags = append(diags, e)
			}
			obj.Allowlist = &v2
		}
	}
	if v1, ok := d.GetOk("block_blacklisted_certificates"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("block_blacklisted_certificates", sv)
				diags = append(diags, e)
			}
			obj.BlockBlacklistedCertificates = &v2
		}
	}
	if v1, ok := d.GetOk("block_blocklisted_certificates"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("block_blocklisted_certificates", sv)
				diags = append(diags, e)
			}
			obj.BlockBlocklistedCertificates = &v2
		}
	}
	if v1, ok := d.GetOk("caname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("caname", sv)
				diags = append(diags, e)
			}
			obj.Caname = &v2
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
	if v, ok := d.GetOk("dot"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("dot", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileDot(d, v, "dot", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dot = t
		}
	} else if d.HasChange("dot") {
		old, new := d.GetChange("dot")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dot = &[]models.FirewallSslSshProfileDot{}
		}
	}
	if v, ok := d.GetOk("ftps"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftps", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileFtps(d, v, "ftps", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ftps = t
		}
	} else if d.HasChange("ftps") {
		old, new := d.GetChange("ftps")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ftps = &[]models.FirewallSslSshProfileFtps{}
		}
	}
	if v, ok := d.GetOk("https"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("https", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileHttps(d, v, "https", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Https = t
		}
	} else if d.HasChange("https") {
		old, new := d.GetChange("https")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Https = &[]models.FirewallSslSshProfileHttps{}
		}
	}
	if v, ok := d.GetOk("imaps"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("imaps", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileImaps(d, v, "imaps", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Imaps = t
		}
	} else if d.HasChange("imaps") {
		old, new := d.GetChange("imaps")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Imaps = &[]models.FirewallSslSshProfileImaps{}
		}
	}
	if v1, ok := d.GetOk("mapi_over_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mapi_over_https", sv)
				diags = append(diags, e)
			}
			obj.MapiOverHttps = &v2
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
	if v, ok := d.GetOk("pop3s"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pop3s", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfilePop3s(d, v, "pop3s", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Pop3s = t
		}
	} else if d.HasChange("pop3s") {
		old, new := d.GetChange("pop3s")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Pop3s = &[]models.FirewallSslSshProfilePop3s{}
		}
	}
	if v1, ok := d.GetOk("rpc_over_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rpc_over_https", sv)
				diags = append(diags, e)
			}
			obj.RpcOverHttps = &v2
		}
	}
	if v, ok := d.GetOk("server_cert"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("server_cert", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileServerCert(d, v, "server_cert", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ServerCert = t
		}
	} else if d.HasChange("server_cert") {
		old, new := d.GetChange("server_cert")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ServerCert = &[]models.FirewallSslSshProfileServerCert{}
		}
	}
	if v1, ok := d.GetOk("server_cert_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_cert_mode", sv)
				diags = append(diags, e)
			}
			obj.ServerCertMode = &v2
		}
	}
	if v, ok := d.GetOk("smtps"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("smtps", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileSmtps(d, v, "smtps", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Smtps = t
		}
	} else if d.HasChange("smtps") {
		old, new := d.GetChange("smtps")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Smtps = &[]models.FirewallSslSshProfileSmtps{}
		}
	}
	if v, ok := d.GetOk("ssh"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssh", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileSsh(d, v, "ssh", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ssh = t
		}
	} else if d.HasChange("ssh") {
		old, new := d.GetChange("ssh")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ssh = &[]models.FirewallSslSshProfileSsh{}
		}
	}
	if v, ok := d.GetOk("ssl"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssl", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ssl = t
		}
	} else if d.HasChange("ssl") {
		old, new := d.GetChange("ssl")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ssl = &[]models.FirewallSslSshProfileSsl{}
		}
	}
	if v1, ok := d.GetOk("ssl_anomalies_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("ssl_anomalies_log", sv)
				diags = append(diags, e)
			}
			obj.SslAnomaliesLog = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_anomaly_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ssl_anomaly_log", sv)
				diags = append(diags, e)
			}
			obj.SslAnomalyLog = &v2
		}
	}
	if v, ok := d.GetOk("ssl_exempt"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssl_exempt", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileSslExempt(d, v, "ssl_exempt", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SslExempt = t
		}
	} else if d.HasChange("ssl_exempt") {
		old, new := d.GetChange("ssl_exempt")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SslExempt = &[]models.FirewallSslSshProfileSslExempt{}
		}
	}
	if v1, ok := d.GetOk("ssl_exemption_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ssl_exemption_log", sv)
				diags = append(diags, e)
			}
			obj.SslExemptionLog = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_exemptions_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.2") {
				e := utils.AttributeVersionWarning("ssl_exemptions_log", sv)
				diags = append(diags, e)
			}
			obj.SslExemptionsLog = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_handshake_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ssl_handshake_log", sv)
				diags = append(diags, e)
			}
			obj.SslHandshakeLog = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_negotiation_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("ssl_negotiation_log", sv)
				diags = append(diags, e)
			}
			obj.SslNegotiationLog = &v2
		}
	}
	if v, ok := d.GetOk("ssl_server"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssl_server", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallSslSshProfileSslServer(d, v, "ssl_server", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SslServer = t
		}
	} else if d.HasChange("ssl_server") {
		old, new := d.GetChange("ssl_server")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SslServer = &[]models.FirewallSslSshProfileSslServer{}
		}
	}
	if v1, ok := d.GetOk("ssl_server_cert_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("ssl_server_cert_log", sv)
				diags = append(diags, e)
			}
			obj.SslServerCertLog = &v2
		}
	}
	if v1, ok := d.GetOk("supported_alpn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("supported_alpn", sv)
				diags = append(diags, e)
			}
			obj.SupportedAlpn = &v2
		}
	}
	if v1, ok := d.GetOk("untrusted_caname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("untrusted_caname", sv)
				diags = append(diags, e)
			}
			obj.UntrustedCaname = &v2
		}
	}
	if v1, ok := d.GetOk("use_ssl_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("use_ssl_server", sv)
				diags = append(diags, e)
			}
			obj.UseSslServer = &v2
		}
	}
	if v1, ok := d.GetOk("whitelist"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("whitelist", sv)
				diags = append(diags, e)
			}
			obj.Whitelist = &v2
		}
	}
	return &obj, diags
}
