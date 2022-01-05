// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure protocol options.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallProfileProtocolOptions() *schema.Resource {
	return &schema.Resource{
		Description: "Configure protocol options.",

		ReadContext: dataSourceFirewallProfileProtocolOptionsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cifs": {
				Type:        schema.TypeList,
				Description: "Configure CIFS protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_controller": {
							Type:        schema.TypeString,
							Description: "Domain for which to decrypt CIFS traffic.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 445).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"server_credential_type": {
							Type:        schema.TypeString,
							Description: "CIFS server credential type.",
							Computed:    true,
						},
						"server_keytab": {
							Type:        schema.TypeList,
							Description: "Server keytab.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keytab": {
										Type:        schema.TypeString,
										Description: "Base64 encoded keytab file containing credential of the server.",
										Computed:    true,
									},
									"principal": {
										Type:        schema.TypeString,
										Description: "Service principal.  For example, \"host/cifsserver.example.com@example.com\".",
										Computed:    true,
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:        schema.TypeInt,
							Description: "Maximum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:        schema.TypeInt,
							Description: "Minimum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_size": {
							Type:        schema.TypeInt,
							Description: "Set TCP static window size.",
							Computed:    true,
						},
						"tcp_window_type": {
							Type:        schema.TypeString,
							Description: "TCP window type to use for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"dns": {
				Type:        schema.TypeList,
				Description: "Configure DNS protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 53).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
					},
				},
			},
			"feature_set": {
				Type:        schema.TypeString,
				Description: "Flow/proxy feature set.",
				Computed:    true,
			},
			"ftp": {
				Type:        schema.TypeList,
				Description: "Configure FTP protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": {
							Type:        schema.TypeInt,
							Description: "Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).",
							Computed:    true,
						},
						"comfort_interval": {
							Type:        schema.TypeInt,
							Description: "Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).",
							Computed:    true,
						},
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable the inspection of all ports for the protocol.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 21).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:        schema.TypeString,
							Description: "SSL decryption and encryption performed by an external device.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"stream_based_uncompressed_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).",
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:        schema.TypeInt,
							Description: "Maximum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:        schema.TypeInt,
							Description: "Minimum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_size": {
							Type:        schema.TypeInt,
							Description: "Set TCP static window size.",
							Computed:    true,
						},
						"tcp_window_type": {
							Type:        schema.TypeString,
							Description: "TCP window type to use for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"http": {
				Type:        schema.TypeList,
				Description: "Configure HTTP protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_page_status_code": {
							Type:        schema.TypeInt,
							Description: "Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).",
							Computed:    true,
						},
						"comfort_amount": {
							Type:        schema.TypeInt,
							Description: "Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).",
							Computed:    true,
						},
						"comfort_interval": {
							Type:        schema.TypeInt,
							Description: "Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).",
							Computed:    true,
						},
						"fortinet_bar": {
							Type:        schema.TypeString,
							Description: "Enable/disable Fortinet bar on HTML content.",
							Computed:    true,
						},
						"fortinet_bar_port": {
							Type:        schema.TypeInt,
							Description: "Port for use by Fortinet Bar (1 - 65535, default = 8011).",
							Computed:    true,
						},
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable the inspection of all ports for the protocol.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 80).",
							Computed:    true,
						},
						"post_lang": {
							Type:        schema.TypeString,
							Description: "ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"range_block": {
							Type:        schema.TypeString,
							Description: "Enable/disable blocking of partial downloads.",
							Computed:    true,
						},
						"retry_count": {
							Type:        schema.TypeInt,
							Description: "Number of attempts to retry HTTP connection (0 - 100, default = 0).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:        schema.TypeString,
							Description: "SSL decryption and encryption performed by an external device.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"stream_based_uncompressed_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).",
							Computed:    true,
						},
						"streaming_content_bypass": {
							Type:        schema.TypeString,
							Description: "Enable/disable bypassing of streaming content from buffering.",
							Computed:    true,
						},
						"strip_x_forwarded_for": {
							Type:        schema.TypeString,
							Description: "Enable/disable stripping of HTTP X-Forwarded-For header.",
							Computed:    true,
						},
						"switching_protocols": {
							Type:        schema.TypeString,
							Description: "Bypass from scanning, or block a connection that attempts to switch protocol.",
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:        schema.TypeInt,
							Description: "Maximum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:        schema.TypeInt,
							Description: "Minimum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_size": {
							Type:        schema.TypeInt,
							Description: "Set TCP static window size.",
							Computed:    true,
						},
						"tcp_window_type": {
							Type:        schema.TypeString,
							Description: "TCP window type to use for this protocol.",
							Computed:    true,
						},
						"tunnel_non_http": {
							Type:        schema.TypeString,
							Description: "Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"unknown_http_version": {
							Type:        schema.TypeString,
							Description: "How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1.",
							Computed:    true,
						},
					},
				},
			},
			"imap": {
				Type:        schema.TypeList,
				Description: "Configure IMAP protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable the inspection of all ports for the protocol.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 143).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:        schema.TypeString,
							Description: "SSL decryption and encryption performed by an external device.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"mail_signature": {
				Type:        schema.TypeList,
				Description: "Configure Mail signature.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature": {
							Type:        schema.TypeString,
							Description: "Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.",
							Computed:    true,
						},
					},
				},
			},
			"mapi": {
				Type:        schema.TypeList,
				Description: "Configure MAPI protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 135).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"nntp": {
				Type:        schema.TypeList,
				Description: "Configure NNTP protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable the inspection of all ports for the protocol.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 119).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"oversize_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging for antivirus oversize file blocking.",
				Computed:    true,
			},
			"pop3": {
				Type:        schema.TypeList,
				Description: "Configure POP3 protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable the inspection of all ports for the protocol.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 110).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:        schema.TypeString,
							Description: "SSL decryption and encryption performed by an external device.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Name of the replacement message group to be used",
				Computed:    true,
			},
			"rpc_over_http": {
				Type:        schema.TypeString,
				Description: "Enable/disable inspection of RPC over HTTP.",
				Computed:    true,
			},
			"smtp": {
				Type:        schema.TypeList,
				Description: "Configure SMTP protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": {
							Type:        schema.TypeString,
							Description: "Enable/disable the inspection of all ports for the protocol.",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"ports": {
							Type:        schema.TypeInt,
							Description: "Ports to scan for content (1 - 65535, default = 25).",
							Computed:    true,
						},
						"proxy_after_tcp_handshake": {
							Type:        schema.TypeString,
							Description: "Proxy traffic after the TCP 3-way handshake has been established (not before).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"server_busy": {
							Type:        schema.TypeString,
							Description: "Enable/disable SMTP server busy when server not available.",
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:        schema.TypeString,
							Description: "SSL decryption and encryption performed by an external device.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable the active status of scanning for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"ssh": {
				Type:        schema.TypeList,
				Description: "Configure SFTP and SCP protocol options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comfort_amount": {
							Type:        schema.TypeInt,
							Description: "Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).",
							Computed:    true,
						},
						"comfort_interval": {
							Type:        schema.TypeInt,
							Description: "Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeString,
							Description: "One or more options that can be applied to the session.",
							Computed:    true,
						},
						"oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
						"scan_bzip2": {
							Type:        schema.TypeString,
							Description: "Enable/disable scanning of BZip2 compressed files.",
							Computed:    true,
						},
						"ssl_offloaded": {
							Type:        schema.TypeString,
							Description: "SSL decryption and encryption performed by an external device.",
							Computed:    true,
						},
						"stream_based_uncompressed_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).",
							Computed:    true,
						},
						"tcp_window_maximum": {
							Type:        schema.TypeInt,
							Description: "Maximum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_minimum": {
							Type:        schema.TypeInt,
							Description: "Minimum dynamic TCP window size.",
							Computed:    true,
						},
						"tcp_window_size": {
							Type:        schema.TypeInt,
							Description: "Set TCP static window size.",
							Computed:    true,
						},
						"tcp_window_type": {
							Type:        schema.TypeString,
							Description: "TCP window type to use for this protocol.",
							Computed:    true,
						},
						"uncompressed_nest_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).",
							Computed:    true,
						},
						"uncompressed_oversize_limit": {
							Type:        schema.TypeInt,
							Description: "Maximum in-memory uncompressed file size that can be scanned (1 - 383 MB, default = 10).",
							Computed:    true,
						},
					},
				},
			},
			"switching_protocols_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging for HTTP/HTTPS switching protocols.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallProfileProtocolOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProfileProtocolOptions(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProfileProtocolOptions dataSource: %v", err)
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

	diags := refreshObjectFirewallProfileProtocolOptions(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
