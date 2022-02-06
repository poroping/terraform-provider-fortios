// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure protocol options.

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

func resourceFirewallProfileProtocolOptions() *schema.Resource {
	return &schema.Resource{
		Description: "Configure protocol options.",

		CreateContext: resourceFirewallProfileProtocolOptionsCreate,
		ReadContext:   resourceFirewallProfileProtocolOptionsRead,
		UpdateContext: resourceFirewallProfileProtocolOptionsUpdate,
		DeleteContext: resourceFirewallProfileProtocolOptionsDelete,

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
			"cifs": {
				Type:        schema.TypeList,
				Description: "Configure CIFS protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_controller": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Domain for which to decrypt CIFS traffic.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 445).",
							Optional:    true,
							Computed:    true,
						},
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"server_credential_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "credential-replication", "credential-keytab"}, false),

							Description: "CIFS server credential type.",
							Optional:    true,
							Computed:    true,
						},
						"server_keytab": {
							Type:        schema.TypeList,
							Description: "Server keytab.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keytab": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 8191),

										Description: "Base64 encoded keytab file containing credential of the server.",
										Optional:    true,
										Computed:    true,
									},
									"principal": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),

										Description: "Service principal. For example, host/cifsserver.example.com@example.com.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1048576, 33554432),

							Description: "Maximum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 1048576),

							Description: "Minimum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 33554432),

							Description: "Set TCP static window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"system", "static", "dynamic", "auto-tuning"}, false),

							Description: "TCP window type to use for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"dns": {
				Type:        schema.TypeList,
				Description: "Configure DNS protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 53).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"feature_set": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"flow", "proxy"}, false),

				Description: "Flow/proxy feature set.",
				Optional:    true,
				Computed:    true,
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Configure FTP protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).",
							Optional:    true,
							Computed:    true,
						},
						"comfort_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 900),

							Description: "Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the inspection of all ports for the protocol.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 21).",
							Optional:    true,
							Computed:    true,
						},
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "SSL decryption and encryption performed by an external device.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"stream_based_uncompressed_limit": {
							Type: schema.TypeInt,

							Description: "Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1048576, 33554432),

							Description: "Maximum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 1048576),

							Description: "Minimum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 33554432),

							Description: "Set TCP static window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"system", "static", "dynamic", "auto-tuning"}, false),

							Description: "TCP window type to use for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Configure HTTP protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_page_status_code": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 599),

							Description: "Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).",
							Optional:    true,
							Computed:    true,
						},
						"comfort_amount": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).",
							Optional:    true,
							Computed:    true,
						},
						"comfort_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 900),

							Description: "Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"fortinet_bar": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable Fortinet bar on HTML content.",
							Optional:    true,
							Computed:    true,
						},
						"fortinet_bar_port": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Port for use by Fortinet Bar (1 - 65535, default = 8011).",
							Optional:    true,
							Computed:    true,
						},
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the inspection of all ports for the protocol.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 80).",
							Optional:    true,
							Computed:    true,
						},
						"post_lang": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets).",
							Optional:         true,
							Computed:         true,
						},
						"proxy_after_tcp_handshake": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Optional:    true,
							Computed:    true,
						},
						"range_block": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable blocking of partial downloads.",
							Optional:    true,
							Computed:    true,
						},
						"retry_count": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),

							Description: "Number of attempts to retry HTTP connection (0 - 100, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "SSL decryption and encryption performed by an external device.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"stream_based_uncompressed_limit": {
							Type: schema.TypeInt,

							Description: "Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"streaming_content_bypass": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable bypassing of streaming content from buffering.",
							Optional:    true,
							Computed:    true,
						},
						"strip_x_forwarded_for": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable stripping of HTTP X-Forwarded-For header.",
							Optional:    true,
							Computed:    true,
						},
						"switching_protocols": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"bypass", "block"}, false),

							Description: "Bypass from scanning, or block a connection that attempts to switch protocol.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1048576, 33554432),

							Description: "Maximum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 1048576),

							Description: "Minimum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 33554432),

							Description: "Set TCP static window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"system", "static", "dynamic", "auto-tuning"}, false),

							Description: "TCP window type to use for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"tunnel_non_http": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"unknown_http_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"reject", "tunnel", "best-effort"}, false),

							Description: "How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"imap": {
				Type:        schema.TypeList,
				Description: "Configure IMAP protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the inspection of all ports for the protocol.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 143).",
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
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "SSL decryption and encryption performed by an external device.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mail_signature": {
				Type:        schema.TypeList,
				Description: "Configure Mail signature.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),

							Description: "Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "Configure MAPI protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 135).",
							Optional:    true,
							Computed:    true,
						},
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"nntp": {
				Type:        schema.TypeList,
				Description: "Configure NNTP protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the inspection of all ports for the protocol.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 119).",
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
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"oversize_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging for antivirus oversize file blocking.",
				Optional:    true,
				Computed:    true,
			},
			"pop3": {
				Type:        schema.TypeList,
				Description: "Configure POP3 protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the inspection of all ports for the protocol.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 110).",
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
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "SSL decryption and encryption performed by an external device.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the replacement message group to be used.",
				Optional:    true,
				Computed:    true,
			},
			"rpc_over_http": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable inspection of RPC over HTTP.",
				Optional:    true,
				Computed:    true,
			},
			"smtp": {
				Type:        schema.TypeList,
				Description: "Configure SMTP protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the inspection of all ports for the protocol.",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"ports": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Ports to scan for content (1 - 65535, default = 25).",
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
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"server_busy": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SMTP server busy when server not available.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "SSL decryption and encryption performed by an external device.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable the active status of scanning for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Description: "Configure SFTP and SCP protocol options.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).",
							Optional:    true,
							Computed:    true,
						},
						"comfort_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 900),

							Description: "Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "One or more options that can be applied to the session.",
							Optional:         true,
							Computed:         true,
						},
						"oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
						"scan_bzip2": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable scanning of BZip2 compressed files.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),

							Description: "SSL decryption and encryption performed by an external device.",
							Optional:    true,
							Computed:    true,
						},
						"stream_based_uncompressed_limit": {
							Type: schema.TypeInt,

							Description: "Maximum stream-based uncompressed data size that will be scanned in megabytes. Stream-based uncompression used only under certain conditions (unlimited = 0, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1048576, 33554432),

							Description: "Maximum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 1048576),

							Description: "Minimum dynamic TCP window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(65536, 33554432),

							Description: "Set TCP static window size.",
							Optional:    true,
							Computed:    true,
						},
						"tcp_window_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"system", "static", "dynamic", "auto-tuning"}, false),

							Description: "TCP window type to use for this protocol.",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(2, 100),

							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Optional:    true,
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 26214),

							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"switching_protocols_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable logging for HTTP/HTTPS switching protocols.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallProfileProtocolOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallProfileProtocolOptions resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallProfileProtocolOptions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallProfileProtocolOptions(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProfileProtocolOptions")
	}

	return resourceFirewallProfileProtocolOptionsRead(ctx, d, meta)
}

func resourceFirewallProfileProtocolOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallProfileProtocolOptions(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallProfileProtocolOptions(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallProfileProtocolOptions resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProfileProtocolOptions")
	}

	return resourceFirewallProfileProtocolOptionsRead(ctx, d, meta)
}

func resourceFirewallProfileProtocolOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallProfileProtocolOptions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallProfileProtocolOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProfileProtocolOptions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProfileProtocolOptions resource: %v", err)
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

	diags := refreshObjectFirewallProfileProtocolOptions(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsCifs(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsCifs, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsCifs{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.DomainController; tmp != nil {
				v["domain_controller"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.ServerCredentialType; tmp != nil {
				v["server_credential_type"] = *tmp
			}

			if tmp := cfg.ServerKeytab; tmp != nil {
				v["server_keytab"] = flattenFirewallProfileProtocolOptionsCifsServerKeytab(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "server_keytab"), sort)
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TcpWindowMaximum; tmp != nil {
				v["tcp_window_maximum"] = *tmp
			}

			if tmp := cfg.TcpWindowMinimum; tmp != nil {
				v["tcp_window_minimum"] = *tmp
			}

			if tmp := cfg.TcpWindowSize; tmp != nil {
				v["tcp_window_size"] = *tmp
			}

			if tmp := cfg.TcpWindowType; tmp != nil {
				v["tcp_window_type"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData, v *[]models.FirewallProfileProtocolOptionsCifsServerKeytab, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Keytab; tmp != nil {
				v["keytab"] = *tmp
			}

			if tmp := cfg.Principal; tmp != nil {
				v["principal"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "principal")
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsDns(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsDns, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsDns{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsFtp(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsFtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsFtp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ComfortAmount; tmp != nil {
				v["comfort_amount"] = *tmp
			}

			if tmp := cfg.ComfortInterval; tmp != nil {
				v["comfort_interval"] = *tmp
			}

			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.SslOffloaded; tmp != nil {
				v["ssl_offloaded"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.StreamBasedUncompressedLimit; tmp != nil {
				v["stream_based_uncompressed_limit"] = *tmp
			}

			if tmp := cfg.TcpWindowMaximum; tmp != nil {
				v["tcp_window_maximum"] = *tmp
			}

			if tmp := cfg.TcpWindowMinimum; tmp != nil {
				v["tcp_window_minimum"] = *tmp
			}

			if tmp := cfg.TcpWindowSize; tmp != nil {
				v["tcp_window_size"] = *tmp
			}

			if tmp := cfg.TcpWindowType; tmp != nil {
				v["tcp_window_type"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsHttp(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsHttp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsHttp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.BlockPageStatusCode; tmp != nil {
				v["block_page_status_code"] = *tmp
			}

			if tmp := cfg.ComfortAmount; tmp != nil {
				v["comfort_amount"] = *tmp
			}

			if tmp := cfg.ComfortInterval; tmp != nil {
				v["comfort_interval"] = *tmp
			}

			if tmp := cfg.FortinetBar; tmp != nil {
				v["fortinet_bar"] = *tmp
			}

			if tmp := cfg.FortinetBarPort; tmp != nil {
				v["fortinet_bar_port"] = *tmp
			}

			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.PostLang; tmp != nil {
				v["post_lang"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.RangeBlock; tmp != nil {
				v["range_block"] = *tmp
			}

			if tmp := cfg.RetryCount; tmp != nil {
				v["retry_count"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.SslOffloaded; tmp != nil {
				v["ssl_offloaded"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.StreamBasedUncompressedLimit; tmp != nil {
				v["stream_based_uncompressed_limit"] = *tmp
			}

			if tmp := cfg.StreamingContentBypass; tmp != nil {
				v["streaming_content_bypass"] = *tmp
			}

			if tmp := cfg.StripXForwardedFor; tmp != nil {
				v["strip_x_forwarded_for"] = *tmp
			}

			if tmp := cfg.SwitchingProtocols; tmp != nil {
				v["switching_protocols"] = *tmp
			}

			if tmp := cfg.TcpWindowMaximum; tmp != nil {
				v["tcp_window_maximum"] = *tmp
			}

			if tmp := cfg.TcpWindowMinimum; tmp != nil {
				v["tcp_window_minimum"] = *tmp
			}

			if tmp := cfg.TcpWindowSize; tmp != nil {
				v["tcp_window_size"] = *tmp
			}

			if tmp := cfg.TcpWindowType; tmp != nil {
				v["tcp_window_type"] = *tmp
			}

			if tmp := cfg.TunnelNonHttp; tmp != nil {
				v["tunnel_non_http"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			if tmp := cfg.UnknownHttpVersion; tmp != nil {
				v["unknown_http_version"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsImap(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsImap, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsImap{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.SslOffloaded; tmp != nil {
				v["ssl_offloaded"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsMailSignature, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsMailSignature{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Signature; tmp != nil {
				v["signature"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsMapi(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsMapi, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsMapi{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsNntp(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsNntp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsNntp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsPop3(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsPop3, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsPop3{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.SslOffloaded; tmp != nil {
				v["ssl_offloaded"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsSmtp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsSmtp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.InspectAll; tmp != nil {
				v["inspect_all"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.Ports; tmp != nil {
				v["ports"] = *tmp
			}

			if tmp := cfg.ProxyAfterTcpHandshake; tmp != nil {
				v["proxy_after_tcp_handshake"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.ServerBusy; tmp != nil {
				v["server_busy"] = *tmp
			}

			if tmp := cfg.SslOffloaded; tmp != nil {
				v["ssl_offloaded"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenFirewallProfileProtocolOptionsSsh(d *schema.ResourceData, v *models.FirewallProfileProtocolOptionsSsh, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.FirewallProfileProtocolOptionsSsh{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.ComfortAmount; tmp != nil {
				v["comfort_amount"] = *tmp
			}

			if tmp := cfg.ComfortInterval; tmp != nil {
				v["comfort_interval"] = *tmp
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			if tmp := cfg.OversizeLimit; tmp != nil {
				v["oversize_limit"] = *tmp
			}

			if tmp := cfg.ScanBzip2; tmp != nil {
				v["scan_bzip2"] = *tmp
			}

			if tmp := cfg.SslOffloaded; tmp != nil {
				v["ssl_offloaded"] = *tmp
			}

			if tmp := cfg.StreamBasedUncompressedLimit; tmp != nil {
				v["stream_based_uncompressed_limit"] = *tmp
			}

			if tmp := cfg.TcpWindowMaximum; tmp != nil {
				v["tcp_window_maximum"] = *tmp
			}

			if tmp := cfg.TcpWindowMinimum; tmp != nil {
				v["tcp_window_minimum"] = *tmp
			}

			if tmp := cfg.TcpWindowSize; tmp != nil {
				v["tcp_window_size"] = *tmp
			}

			if tmp := cfg.TcpWindowType; tmp != nil {
				v["tcp_window_type"] = *tmp
			}

			if tmp := cfg.UncompressedNestLimit; tmp != nil {
				v["uncompressed_nest_limit"] = *tmp
			}

			if tmp := cfg.UncompressedOversizeLimit; tmp != nil {
				v["uncompressed_oversize_limit"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectFirewallProfileProtocolOptions(d *schema.ResourceData, o *models.FirewallProfileProtocolOptions, sv string, sort bool) diag.Diagnostics {
	var err error

	if _, ok := d.GetOk("cifs"); ok {
		if o.Cifs != nil {
			if err = d.Set("cifs", flattenFirewallProfileProtocolOptionsCifs(d, o.Cifs, "cifs", sort)); err != nil {
				return diag.Errorf("error reading cifs: %v", err)
			}
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if _, ok := d.GetOk("dns"); ok {
		if o.Dns != nil {
			if err = d.Set("dns", flattenFirewallProfileProtocolOptionsDns(d, o.Dns, "dns", sort)); err != nil {
				return diag.Errorf("error reading dns: %v", err)
			}
		}
	}

	if o.FeatureSet != nil {
		v := *o.FeatureSet

		if err = d.Set("feature_set", v); err != nil {
			return diag.Errorf("error reading feature_set: %v", err)
		}
	}

	if _, ok := d.GetOk("ftp"); ok {
		if o.Ftp != nil {
			if err = d.Set("ftp", flattenFirewallProfileProtocolOptionsFtp(d, o.Ftp, "ftp", sort)); err != nil {
				return diag.Errorf("error reading ftp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("http"); ok {
		if o.Http != nil {
			if err = d.Set("http", flattenFirewallProfileProtocolOptionsHttp(d, o.Http, "http", sort)); err != nil {
				return diag.Errorf("error reading http: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("imap"); ok {
		if o.Imap != nil {
			if err = d.Set("imap", flattenFirewallProfileProtocolOptionsImap(d, o.Imap, "imap", sort)); err != nil {
				return diag.Errorf("error reading imap: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("mail_signature"); ok {
		if o.MailSignature != nil {
			if err = d.Set("mail_signature", flattenFirewallProfileProtocolOptionsMailSignature(d, o.MailSignature, "mail_signature", sort)); err != nil {
				return diag.Errorf("error reading mail_signature: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("mapi"); ok {
		if o.Mapi != nil {
			if err = d.Set("mapi", flattenFirewallProfileProtocolOptionsMapi(d, o.Mapi, "mapi", sort)); err != nil {
				return diag.Errorf("error reading mapi: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if _, ok := d.GetOk("nntp"); ok {
		if o.Nntp != nil {
			if err = d.Set("nntp", flattenFirewallProfileProtocolOptionsNntp(d, o.Nntp, "nntp", sort)); err != nil {
				return diag.Errorf("error reading nntp: %v", err)
			}
		}
	}

	if o.OversizeLog != nil {
		v := *o.OversizeLog

		if err = d.Set("oversize_log", v); err != nil {
			return diag.Errorf("error reading oversize_log: %v", err)
		}
	}

	if _, ok := d.GetOk("pop3"); ok {
		if o.Pop3 != nil {
			if err = d.Set("pop3", flattenFirewallProfileProtocolOptionsPop3(d, o.Pop3, "pop3", sort)); err != nil {
				return diag.Errorf("error reading pop3: %v", err)
			}
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.RpcOverHttp != nil {
		v := *o.RpcOverHttp

		if err = d.Set("rpc_over_http", v); err != nil {
			return diag.Errorf("error reading rpc_over_http: %v", err)
		}
	}

	if _, ok := d.GetOk("smtp"); ok {
		if o.Smtp != nil {
			if err = d.Set("smtp", flattenFirewallProfileProtocolOptionsSmtp(d, o.Smtp, "smtp", sort)); err != nil {
				return diag.Errorf("error reading smtp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("ssh"); ok {
		if o.Ssh != nil {
			if err = d.Set("ssh", flattenFirewallProfileProtocolOptionsSsh(d, o.Ssh, "ssh", sort)); err != nil {
				return diag.Errorf("error reading ssh: %v", err)
			}
		}
	}

	if o.SwitchingProtocolsLog != nil {
		v := *o.SwitchingProtocolsLog

		if err = d.Set("switching_protocols_log", v); err != nil {
			return diag.Errorf("error reading switching_protocols_log: %v", err)
		}
	}

	return nil
}

func expandFirewallProfileProtocolOptionsCifs(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsCifs, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsCifs

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsCifs{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.domain_controller", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.DomainController = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_credential_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerCredentialType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_keytab", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallProfileProtocolOptionsCifsServerKeytab(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallProfileProtocolOptionsCifsServerKeytab
			// 	}
			tmp.ServerKeytab = v2

		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_maximum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMaximum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_minimum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMinimum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpWindowType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytab(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProfileProtocolOptionsCifsServerKeytab, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsCifsServerKeytab

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsCifsServerKeytab{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.keytab", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Keytab = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.principal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Principal = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallProfileProtocolOptionsDns(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsDns, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsDns

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsDns{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsFtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsFtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsFtp

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsFtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comfort_amount", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ComfortAmount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comfort_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ComfortInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_offloaded", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOffloaded = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stream_based_uncompressed_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StreamBasedUncompressedLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_maximum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMaximum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_minimum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMinimum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpWindowType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsHttp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsHttp

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsHttp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.block_page_status_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.BlockPageStatusCode = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comfort_amount", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ComfortAmount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comfort_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ComfortInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortinet_bar", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FortinetBar = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.fortinet_bar_port", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.FortinetBarPort = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.post_lang", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PostLang = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.proxy_after_tcp_handshake", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ProxyAfterTcpHandshake = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.range_block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RangeBlock = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.retry_count", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RetryCount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_offloaded", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOffloaded = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stream_based_uncompressed_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StreamBasedUncompressedLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.streaming_content_bypass", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StreamingContentBypass = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.strip_x_forwarded_for", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StripXForwardedFor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.switching_protocols", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SwitchingProtocols = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_maximum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMaximum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_minimum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMinimum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpWindowType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tunnel_non_http", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TunnelNonHttp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unknown_http_version", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnknownHttpVersion = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsImap(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsImap, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsImap

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsImap{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
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

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_offloaded", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOffloaded = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsMailSignature, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsMailSignature

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsMailSignature{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.signature", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Signature = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsMapi(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsMapi, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsMapi

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsMapi{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ports", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ports = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsNntp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsNntp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsNntp

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsNntp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
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

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsPop3(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsPop3, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsPop3

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsPop3{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
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

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_offloaded", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOffloaded = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsSmtp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsSmtp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsSmtp

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsSmtp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.inspect_all", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InspectAll = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
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

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.server_busy", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ServerBusy = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_offloaded", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOffloaded = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandFirewallProfileProtocolOptionsSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.FirewallProfileProtocolOptionsSsh, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProfileProtocolOptionsSsh

	for i := range l {
		tmp := models.FirewallProfileProtocolOptionsSsh{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.comfort_amount", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ComfortAmount = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comfort_interval", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ComfortInterval = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OversizeLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.scan_bzip2", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ScanBzip2 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_offloaded", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslOffloaded = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.stream_based_uncompressed_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.StreamBasedUncompressedLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_maximum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMaximum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_minimum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowMinimum = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.TcpWindowSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tcp_window_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TcpWindowType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_nest_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedNestLimit = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.uncompressed_oversize_limit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UncompressedOversizeLimit = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectFirewallProfileProtocolOptions(d *schema.ResourceData, sv string) (*models.FirewallProfileProtocolOptions, diag.Diagnostics) {
	obj := models.FirewallProfileProtocolOptions{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("cifs"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("cifs", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsCifs(d, v, "cifs", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Cifs = t
		}
	} else if d.HasChange("cifs") {
		old, new := d.GetChange("cifs")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Cifs = &models.FirewallProfileProtocolOptionsCifs{}
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
	if v, ok := d.GetOk("dns"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dns", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsDns(d, v, "dns", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dns = t
		}
	} else if d.HasChange("dns") {
		old, new := d.GetChange("dns")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dns = &models.FirewallProfileProtocolOptionsDns{}
		}
	}
	if v1, ok := d.GetOk("feature_set"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "v6.4.2") {
				e := utils.AttributeVersionWarning("feature_set", sv)
				diags = append(diags, e)
			}
			obj.FeatureSet = &v2
		}
	}
	if v, ok := d.GetOk("ftp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftp", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsFtp(d, v, "ftp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ftp = t
		}
	} else if d.HasChange("ftp") {
		old, new := d.GetChange("ftp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ftp = &models.FirewallProfileProtocolOptionsFtp{}
		}
	}
	if v, ok := d.GetOk("http"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("http", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsHttp(d, v, "http", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Http = t
		}
	} else if d.HasChange("http") {
		old, new := d.GetChange("http")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Http = &models.FirewallProfileProtocolOptionsHttp{}
		}
	}
	if v, ok := d.GetOk("imap"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("imap", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsImap(d, v, "imap", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Imap = t
		}
	} else if d.HasChange("imap") {
		old, new := d.GetChange("imap")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Imap = &models.FirewallProfileProtocolOptionsImap{}
		}
	}
	if v, ok := d.GetOk("mail_signature"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mail_signature", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsMailSignature(d, v, "mail_signature", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.MailSignature = t
		}
	} else if d.HasChange("mail_signature") {
		old, new := d.GetChange("mail_signature")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.MailSignature = &models.FirewallProfileProtocolOptionsMailSignature{}
		}
	}
	if v, ok := d.GetOk("mapi"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("mapi", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsMapi(d, v, "mapi", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Mapi = t
		}
	} else if d.HasChange("mapi") {
		old, new := d.GetChange("mapi")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Mapi = &models.FirewallProfileProtocolOptionsMapi{}
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
	if v, ok := d.GetOk("nntp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nntp", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsNntp(d, v, "nntp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Nntp = t
		}
	} else if d.HasChange("nntp") {
		old, new := d.GetChange("nntp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Nntp = &models.FirewallProfileProtocolOptionsNntp{}
		}
	}
	if v1, ok := d.GetOk("oversize_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("oversize_log", sv)
				diags = append(diags, e)
			}
			obj.OversizeLog = &v2
		}
	}
	if v, ok := d.GetOk("pop3"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("pop3", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsPop3(d, v, "pop3", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Pop3 = t
		}
	} else if d.HasChange("pop3") {
		old, new := d.GetChange("pop3")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Pop3 = &models.FirewallProfileProtocolOptionsPop3{}
		}
	}
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("rpc_over_http"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rpc_over_http", sv)
				diags = append(diags, e)
			}
			obj.RpcOverHttp = &v2
		}
	}
	if v, ok := d.GetOk("smtp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("smtp", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsSmtp(d, v, "smtp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Smtp = t
		}
	} else if d.HasChange("smtp") {
		old, new := d.GetChange("smtp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Smtp = &models.FirewallProfileProtocolOptionsSmtp{}
		}
	}
	if v, ok := d.GetOk("ssh"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ssh", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProfileProtocolOptionsSsh(d, v, "ssh", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Ssh = t
		}
	} else if d.HasChange("ssh") {
		old, new := d.GetChange("ssh")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Ssh = &models.FirewallProfileProtocolOptionsSsh{}
		}
	}
	if v1, ok := d.GetOk("switching_protocols_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("switching_protocols_log", sv)
				diags = append(diags, e)
			}
			obj.SwitchingProtocolsLog = &v2
		}
	}
	return &obj, diags
}
