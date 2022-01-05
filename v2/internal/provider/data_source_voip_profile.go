// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VoIP profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVoipProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VoIP profiles.",

		ReadContext: dataSourceVoipProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"feature_set": {
				Type:        schema.TypeString,
				Description: "Flow or proxy inspection feature set.",
				Computed:    true,
			},
			"msrp": {
				Type:        schema.TypeList,
				Description: "MSRP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_violations": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of MSRP violations.",
							Computed:    true,
						},
						"max_msg_size": {
							Type:        schema.TypeInt,
							Description: "Maximum allowable MSRP message size (1-65535).",
							Computed:    true,
						},
						"max_msg_size_action": {
							Type:        schema.TypeString,
							Description: "Action for violation of max-msg-size.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable MSRP.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
			"sccp": {
				Type:        schema.TypeList,
				Description: "SCCP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_mcast": {
							Type:        schema.TypeString,
							Description: "Enable/disable block multicast RTP connections.",
							Computed:    true,
						},
						"log_call_summary": {
							Type:        schema.TypeString,
							Description: "Enable/disable log summary of SCCP calls.",
							Computed:    true,
						},
						"log_violations": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of SCCP violations.",
							Computed:    true,
						},
						"max_calls": {
							Type:        schema.TypeInt,
							Description: "Maximum calls per minute per SCCP client (max 65535).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SCCP.",
							Computed:    true,
						},
						"verify_header": {
							Type:        schema.TypeString,
							Description: "Enable/disable verify SCCP header content.",
							Computed:    true,
						},
					},
				},
			},
			"sip": {
				Type:        schema.TypeList,
				Description: "SIP.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ack_rate": {
							Type:        schema.TypeInt,
							Description: "ACK request rate limit (per second, per policy).",
							Computed:    true,
						},
						"ack_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"block_ack": {
							Type:        schema.TypeString,
							Description: "Enable/disable block ACK requests.",
							Computed:    true,
						},
						"block_bye": {
							Type:        schema.TypeString,
							Description: "Enable/disable block BYE requests.",
							Computed:    true,
						},
						"block_cancel": {
							Type:        schema.TypeString,
							Description: "Enable/disable block CANCEL requests.",
							Computed:    true,
						},
						"block_geo_red_options": {
							Type:        schema.TypeString,
							Description: "Enable/disable block OPTIONS requests, but OPTIONS requests still notify for redundancy.",
							Computed:    true,
						},
						"block_info": {
							Type:        schema.TypeString,
							Description: "Enable/disable block INFO requests.",
							Computed:    true,
						},
						"block_invite": {
							Type:        schema.TypeString,
							Description: "Enable/disable block INVITE requests.",
							Computed:    true,
						},
						"block_long_lines": {
							Type:        schema.TypeString,
							Description: "Enable/disable block requests with headers exceeding max-line-length.",
							Computed:    true,
						},
						"block_message": {
							Type:        schema.TypeString,
							Description: "Enable/disable block MESSAGE requests.",
							Computed:    true,
						},
						"block_notify": {
							Type:        schema.TypeString,
							Description: "Enable/disable block NOTIFY requests.",
							Computed:    true,
						},
						"block_options": {
							Type:        schema.TypeString,
							Description: "Enable/disable block OPTIONS requests and no OPTIONS as notifying message for redundancy either.",
							Computed:    true,
						},
						"block_prack": {
							Type:        schema.TypeString,
							Description: "Enable/disable block prack requests.",
							Computed:    true,
						},
						"block_publish": {
							Type:        schema.TypeString,
							Description: "Enable/disable block PUBLISH requests.",
							Computed:    true,
						},
						"block_refer": {
							Type:        schema.TypeString,
							Description: "Enable/disable block REFER requests.",
							Computed:    true,
						},
						"block_register": {
							Type:        schema.TypeString,
							Description: "Enable/disable block REGISTER requests.",
							Computed:    true,
						},
						"block_subscribe": {
							Type:        schema.TypeString,
							Description: "Enable/disable block SUBSCRIBE requests.",
							Computed:    true,
						},
						"block_unknown": {
							Type:        schema.TypeString,
							Description: "Block unrecognized SIP requests (enabled by default).",
							Computed:    true,
						},
						"block_update": {
							Type:        schema.TypeString,
							Description: "Enable/disable block UPDATE requests.",
							Computed:    true,
						},
						"bye_rate": {
							Type:        schema.TypeInt,
							Description: "BYE request rate limit (per second, per policy).",
							Computed:    true,
						},
						"bye_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"call_keepalive": {
							Type:        schema.TypeInt,
							Description: "Continue tracking calls with no RTP for this many minutes.",
							Computed:    true,
						},
						"cancel_rate": {
							Type:        schema.TypeInt,
							Description: "CANCEL request rate limit (per second, per policy).",
							Computed:    true,
						},
						"cancel_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"contact_fixup": {
							Type:        schema.TypeString,
							Description: "Fixup contact anyway even if contact's IP:port doesn't match session's IP:port.",
							Computed:    true,
						},
						"hnt_restrict_source_ip": {
							Type:        schema.TypeString,
							Description: "Enable/disable restrict RTP source IP to be the same as SIP source IP when HNT is enabled.",
							Computed:    true,
						},
						"hosted_nat_traversal": {
							Type:        schema.TypeString,
							Description: "Hosted NAT Traversal (HNT).",
							Computed:    true,
						},
						"info_rate": {
							Type:        schema.TypeInt,
							Description: "INFO request rate limit (per second, per policy).",
							Computed:    true,
						},
						"info_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"invite_rate": {
							Type:        schema.TypeInt,
							Description: "INVITE request rate limit (per second, per policy).",
							Computed:    true,
						},
						"invite_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"ips_rtp": {
							Type:        schema.TypeString,
							Description: "Enable/disable allow IPS on RTP.",
							Computed:    true,
						},
						"log_call_summary": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of SIP call summary.",
							Computed:    true,
						},
						"log_violations": {
							Type:        schema.TypeString,
							Description: "Enable/disable logging of SIP violations.",
							Computed:    true,
						},
						"malformed_header_allow": {
							Type:        schema.TypeString,
							Description: "Action for malformed Allow header.",
							Computed:    true,
						},
						"malformed_header_call_id": {
							Type:        schema.TypeString,
							Description: "Action for malformed Call-ID header.",
							Computed:    true,
						},
						"malformed_header_contact": {
							Type:        schema.TypeString,
							Description: "Action for malformed Contact header.",
							Computed:    true,
						},
						"malformed_header_content_length": {
							Type:        schema.TypeString,
							Description: "Action for malformed Content-Length header.",
							Computed:    true,
						},
						"malformed_header_content_type": {
							Type:        schema.TypeString,
							Description: "Action for malformed Content-Type header.",
							Computed:    true,
						},
						"malformed_header_cseq": {
							Type:        schema.TypeString,
							Description: "Action for malformed CSeq header.",
							Computed:    true,
						},
						"malformed_header_expires": {
							Type:        schema.TypeString,
							Description: "Action for malformed Expires header.",
							Computed:    true,
						},
						"malformed_header_from": {
							Type:        schema.TypeString,
							Description: "Action for malformed From header.",
							Computed:    true,
						},
						"malformed_header_max_forwards": {
							Type:        schema.TypeString,
							Description: "Action for malformed Max-Forwards header.",
							Computed:    true,
						},
						"malformed_header_no_proxy_require": {
							Type:        schema.TypeString,
							Description: "Action for malformed SIP messages without Proxy-Require header.",
							Computed:    true,
						},
						"malformed_header_no_require": {
							Type:        schema.TypeString,
							Description: "Action for malformed SIP messages without Require header.",
							Computed:    true,
						},
						"malformed_header_p_asserted_identity": {
							Type:        schema.TypeString,
							Description: "Action for malformed P-Asserted-Identity header.",
							Computed:    true,
						},
						"malformed_header_rack": {
							Type:        schema.TypeString,
							Description: "Action for malformed RAck header.",
							Computed:    true,
						},
						"malformed_header_record_route": {
							Type:        schema.TypeString,
							Description: "Action for malformed Record-Route header.",
							Computed:    true,
						},
						"malformed_header_route": {
							Type:        schema.TypeString,
							Description: "Action for malformed Route header.",
							Computed:    true,
						},
						"malformed_header_rseq": {
							Type:        schema.TypeString,
							Description: "Action for malformed RSeq header.",
							Computed:    true,
						},
						"malformed_header_sdp_a": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP a line.",
							Computed:    true,
						},
						"malformed_header_sdp_b": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP b line.",
							Computed:    true,
						},
						"malformed_header_sdp_c": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP c line.",
							Computed:    true,
						},
						"malformed_header_sdp_i": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP i line.",
							Computed:    true,
						},
						"malformed_header_sdp_k": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP k line.",
							Computed:    true,
						},
						"malformed_header_sdp_m": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP m line.",
							Computed:    true,
						},
						"malformed_header_sdp_o": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP o line.",
							Computed:    true,
						},
						"malformed_header_sdp_r": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP r line.",
							Computed:    true,
						},
						"malformed_header_sdp_s": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP s line.",
							Computed:    true,
						},
						"malformed_header_sdp_t": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP t line.",
							Computed:    true,
						},
						"malformed_header_sdp_v": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP v line.",
							Computed:    true,
						},
						"malformed_header_sdp_z": {
							Type:        schema.TypeString,
							Description: "Action for malformed SDP z line.",
							Computed:    true,
						},
						"malformed_header_to": {
							Type:        schema.TypeString,
							Description: "Action for malformed To header.",
							Computed:    true,
						},
						"malformed_header_via": {
							Type:        schema.TypeString,
							Description: "Action for malformed VIA header.",
							Computed:    true,
						},
						"malformed_request_line": {
							Type:        schema.TypeString,
							Description: "Action for malformed request line.",
							Computed:    true,
						},
						"max_body_length": {
							Type:        schema.TypeInt,
							Description: "Maximum SIP message body length (0 meaning no limit).",
							Computed:    true,
						},
						"max_dialogs": {
							Type:        schema.TypeInt,
							Description: "Maximum number of concurrent calls/dialogs (per policy).",
							Computed:    true,
						},
						"max_idle_dialogs": {
							Type:        schema.TypeInt,
							Description: "Maximum number established but idle dialogs to retain (per policy).",
							Computed:    true,
						},
						"max_line_length": {
							Type:        schema.TypeInt,
							Description: "Maximum SIP header line length (78-4096).",
							Computed:    true,
						},
						"message_rate": {
							Type:        schema.TypeInt,
							Description: "MESSAGE request rate limit (per second, per policy).",
							Computed:    true,
						},
						"message_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"nat_port_range": {
							Type:        schema.TypeString,
							Description: "RTP NAT port range.",
							Computed:    true,
						},
						"nat_trace": {
							Type:        schema.TypeString,
							Description: "Enable/disable preservation of original IP in SDP i line.",
							Computed:    true,
						},
						"no_sdp_fixup": {
							Type:        schema.TypeString,
							Description: "Enable/disable no SDP fix-up.",
							Computed:    true,
						},
						"notify_rate": {
							Type:        schema.TypeInt,
							Description: "NOTIFY request rate limit (per second, per policy).",
							Computed:    true,
						},
						"notify_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"open_contact_pinhole": {
							Type:        schema.TypeString,
							Description: "Enable/disable open pinhole for non-REGISTER Contact port.",
							Computed:    true,
						},
						"open_record_route_pinhole": {
							Type:        schema.TypeString,
							Description: "Enable/disable open pinhole for Record-Route port.",
							Computed:    true,
						},
						"open_register_pinhole": {
							Type:        schema.TypeString,
							Description: "Enable/disable open pinhole for REGISTER Contact port.",
							Computed:    true,
						},
						"open_via_pinhole": {
							Type:        schema.TypeString,
							Description: "Enable/disable open pinhole for Via port.",
							Computed:    true,
						},
						"options_rate": {
							Type:        schema.TypeInt,
							Description: "OPTIONS request rate limit (per second, per policy).",
							Computed:    true,
						},
						"options_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"prack_rate": {
							Type:        schema.TypeInt,
							Description: "PRACK request rate limit (per second, per policy).",
							Computed:    true,
						},
						"prack_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"preserve_override": {
							Type:        schema.TypeString,
							Description: "Override i line to preserve original IPS (default: append).",
							Computed:    true,
						},
						"provisional_invite_expiry_time": {
							Type:        schema.TypeInt,
							Description: "Expiry time (10-3600, in seconds) for provisional INVITE.",
							Computed:    true,
						},
						"publish_rate": {
							Type:        schema.TypeInt,
							Description: "PUBLISH request rate limit (per second, per policy).",
							Computed:    true,
						},
						"publish_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"refer_rate": {
							Type:        schema.TypeInt,
							Description: "REFER request rate limit (per second, per policy).",
							Computed:    true,
						},
						"refer_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"register_contact_trace": {
							Type:        schema.TypeString,
							Description: "Enable/disable trace original IP/port within the contact header of REGISTER requests.",
							Computed:    true,
						},
						"register_rate": {
							Type:        schema.TypeInt,
							Description: "REGISTER request rate limit (per second, per policy).",
							Computed:    true,
						},
						"register_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"rfc2543_branch": {
							Type:        schema.TypeString,
							Description: "Enable/disable support via branch compliant with RFC 2543.",
							Computed:    true,
						},
						"rtp": {
							Type:        schema.TypeString,
							Description: "Enable/disable create pinholes for RTP traffic to traverse firewall.",
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:        schema.TypeString,
							Description: "Relative strength of encryption algorithms accepted in negotiation.",
							Computed:    true,
						},
						"ssl_auth_client": {
							Type:        schema.TypeString,
							Description: "Require a client certificate and authenticate it with the peer/peergrp.",
							Computed:    true,
						},
						"ssl_auth_server": {
							Type:        schema.TypeString,
							Description: "Authenticate the server's certificate with the peer/peergrp.",
							Computed:    true,
						},
						"ssl_client_certificate": {
							Type:        schema.TypeString,
							Description: "Name of Certificate to offer to server if requested.",
							Computed:    true,
						},
						"ssl_client_renegotiation": {
							Type:        schema.TypeString,
							Description: "Allow/block client renegotiation by server.",
							Computed:    true,
						},
						"ssl_max_version": {
							Type:        schema.TypeString,
							Description: "Highest SSL/TLS version to negotiate.",
							Computed:    true,
						},
						"ssl_min_version": {
							Type:        schema.TypeString,
							Description: "Lowest SSL/TLS version to negotiate.",
							Computed:    true,
						},
						"ssl_mode": {
							Type:        schema.TypeString,
							Description: "SSL/TLS mode for encryption & decryption of traffic.",
							Computed:    true,
						},
						"ssl_pfs": {
							Type:        schema.TypeString,
							Description: "SSL Perfect Forward Secrecy.",
							Computed:    true,
						},
						"ssl_send_empty_frags": {
							Type:        schema.TypeString,
							Description: "Send empty fragments to avoid attack on CBC IV (SSL 3.0 & TLS 1.0 only).",
							Computed:    true,
						},
						"ssl_server_certificate": {
							Type:        schema.TypeString,
							Description: "Name of Certificate return to the client in every SSL connection.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SIP.",
							Computed:    true,
						},
						"strict_register": {
							Type:        schema.TypeString,
							Description: "Enable/disable only allow the registrar to connect.",
							Computed:    true,
						},
						"subscribe_rate": {
							Type:        schema.TypeInt,
							Description: "SUBSCRIBE request rate limit (per second, per policy).",
							Computed:    true,
						},
						"subscribe_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
						"unknown_header": {
							Type:        schema.TypeString,
							Description: "Action for unknown SIP header.",
							Computed:    true,
						},
						"update_rate": {
							Type:        schema.TypeInt,
							Description: "UPDATE request rate limit (per second, per policy).",
							Computed:    true,
						},
						"update_rate_track": {
							Type:        schema.TypeString,
							Description: "Track the packet protocol field.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVoipProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVoipProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VoipProfile dataSource: %v", err)
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

	diags := refreshObjectVoipProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
