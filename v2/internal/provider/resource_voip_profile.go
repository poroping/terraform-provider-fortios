// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VoIP profiles.

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

func resourceVoipProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VoIP profiles.",

		CreateContext: resourceVoipProfileCreate,
		ReadContext:   resourceVoipProfileRead,
		UpdateContext: resourceVoipProfileUpdate,
		DeleteContext: resourceVoipProfileDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"feature_set": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"flow", "proxy"}, false),

				Description: "Flow or proxy inspection feature set.",
				Optional:    true,
				Computed:    true,
			},
			"msrp": {
				Type:        schema.TypeList,
				Description: "MSRP.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_violations": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of MSRP violations.",
							Optional:    true,
							Computed:    true,
						},
						"max_msg_size": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Maximum allowable MSRP message size (1-65535).",
							Optional:    true,
							Computed:    true,
						},
						"max_msg_size_action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block", "reset", "monitor"}, false),

							Description: "Action for violation of max-msg-size.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable MSRP.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"sccp": {
				Type:        schema.TypeList,
				Description: "SCCP.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_mcast": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block multicast RTP connections.",
							Optional:    true,
							Computed:    true,
						},
						"log_call_summary": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable log summary of SCCP calls.",
							Optional:    true,
							Computed:    true,
						},
						"log_violations": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of SCCP violations.",
							Optional:    true,
							Computed:    true,
						},
						"max_calls": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Maximum calls per minute per SCCP client (max 65535).",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SCCP.",
							Optional:    true,
							Computed:    true,
						},
						"verify_header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable verify SCCP header content.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sip": {
				Type:        schema.TypeList,
				Description: "SIP.",
				Optional:    true, MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ack_rate": {
							Type: schema.TypeInt,

							Description: "ACK request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"ack_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"block_ack": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block ACK requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_bye": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block BYE requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_cancel": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block CANCEL requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_geo_red_options": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block OPTIONS requests, but OPTIONS requests still notify for redundancy.",
							Optional:    true,
							Computed:    true,
						},
						"block_info": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block INFO requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_invite": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block INVITE requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_long_lines": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block requests with headers exceeding max-line-length.",
							Optional:    true,
							Computed:    true,
						},
						"block_message": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block MESSAGE requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_notify": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block NOTIFY requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_options": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block OPTIONS requests and no OPTIONS as notifying message for redundancy either.",
							Optional:    true,
							Computed:    true,
						},
						"block_prack": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block prack requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_publish": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block PUBLISH requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_refer": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block REFER requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_register": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block REGISTER requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_subscribe": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block SUBSCRIBE requests.",
							Optional:    true,
							Computed:    true,
						},
						"block_unknown": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Block unrecognized SIP requests (enabled by default).",
							Optional:    true,
							Computed:    true,
						},
						"block_update": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable block UPDATE requests.",
							Optional:    true,
							Computed:    true,
						},
						"bye_rate": {
							Type: schema.TypeInt,

							Description: "BYE request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"bye_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"call_keepalive": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10080),

							Description: "Continue tracking calls with no RTP for this many minutes.",
							Optional:    true,
							Computed:    true,
						},
						"cancel_rate": {
							Type: schema.TypeInt,

							Description: "CANCEL request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"cancel_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"contact_fixup": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Fixup contact anyway even if contact's IP:port doesn't match session's IP:port.",
							Optional:    true,
							Computed:    true,
						},
						"hnt_restrict_source_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable restrict RTP source IP to be the same as SIP source IP when HNT is enabled.",
							Optional:    true,
							Computed:    true,
						},
						"hosted_nat_traversal": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Hosted NAT Traversal (HNT).",
							Optional:    true,
							Computed:    true,
						},
						"info_rate": {
							Type: schema.TypeInt,

							Description: "INFO request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"info_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"invite_rate": {
							Type: schema.TypeInt,

							Description: "INVITE request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"invite_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"ips_rtp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable allow IPS on RTP.",
							Optional:    true,
							Computed:    true,
						},
						"log_call_summary": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of SIP call summary.",
							Optional:    true,
							Computed:    true,
						},
						"log_violations": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable logging of SIP violations.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_allow": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Allow header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_call_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Call-ID header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_contact": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Contact header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_content_length": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Content-Length header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_content_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Content-Type header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_cseq": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed CSeq header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_expires": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Expires header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_from": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed From header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_max_forwards": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Max-Forwards header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_no_proxy_require": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SIP messages without Proxy-Require header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_no_require": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SIP messages without Require header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_p_asserted_identity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed P-Asserted-Identity header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_rack": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed RAck header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_record_route": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Record-Route header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_route": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed Route header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_rseq": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed RSeq header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_a": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP a line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_b": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP b line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_c": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP c line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_i": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP i line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_k": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP k line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_m": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP m line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_o": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP o line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_r": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP r line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_s": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP s line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_t": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP t line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_v": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP v line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_sdp_z": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed SDP z line.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_to": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed To header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_header_via": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed VIA header.",
							Optional:    true,
							Computed:    true,
						},
						"malformed_request_line": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for malformed request line.",
							Optional:    true,
							Computed:    true,
						},
						"max_body_length": {
							Type: schema.TypeInt,

							Description: "Maximum SIP message body length (0 meaning no limit).",
							Optional:    true,
							Computed:    true,
						},
						"max_dialogs": {
							Type: schema.TypeInt,

							Description: "Maximum number of concurrent calls/dialogs (per policy).",
							Optional:    true,
							Computed:    true,
						},
						"max_idle_dialogs": {
							Type: schema.TypeInt,

							Description: "Maximum number established but idle dialogs to retain (per policy).",
							Optional:    true,
							Computed:    true,
						},
						"max_line_length": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(78, 4096),

							Description: "Maximum SIP header line length (78-4096).",
							Optional:    true,
							Computed:    true,
						},
						"message_rate": {
							Type: schema.TypeInt,

							Description: "MESSAGE request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"message_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"nat_port_range": {
							Type: schema.TypeString,

							Description: "RTP NAT port range.",
							Optional:    true,
							Computed:    true,
						},
						"nat_trace": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable preservation of original IP in SDP i line.",
							Optional:    true,
							Computed:    true,
						},
						"no_sdp_fixup": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable no SDP fix-up.",
							Optional:    true,
							Computed:    true,
						},
						"notify_rate": {
							Type: schema.TypeInt,

							Description: "NOTIFY request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"notify_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"open_contact_pinhole": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable open pinhole for non-REGISTER Contact port.",
							Optional:    true,
							Computed:    true,
						},
						"open_record_route_pinhole": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable open pinhole for Record-Route port.",
							Optional:    true,
							Computed:    true,
						},
						"open_register_pinhole": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable open pinhole for REGISTER Contact port.",
							Optional:    true,
							Computed:    true,
						},
						"open_via_pinhole": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable open pinhole for Via port.",
							Optional:    true,
							Computed:    true,
						},
						"options_rate": {
							Type: schema.TypeInt,

							Description: "OPTIONS request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"options_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"prack_rate": {
							Type: schema.TypeInt,

							Description: "PRACK request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"prack_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"preserve_override": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Override i line to preserve original IPS (default: append).",
							Optional:    true,
							Computed:    true,
						},
						"provisional_invite_expiry_time": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(10, 3600),

							Description: "Expiry time (10-3600, in seconds) for provisional INVITE.",
							Optional:    true,
							Computed:    true,
						},
						"publish_rate": {
							Type: schema.TypeInt,

							Description: "PUBLISH request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"publish_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"refer_rate": {
							Type: schema.TypeInt,

							Description: "REFER request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"refer_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"register_contact_trace": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable trace original IP/port within the contact header of REGISTER requests.",
							Optional:    true,
							Computed:    true,
						},
						"register_rate": {
							Type: schema.TypeInt,

							Description: "REGISTER request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"register_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"rfc2543_branch": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable support via branch compliant with RFC 2543.",
							Optional:    true,
							Computed:    true,
						},
						"rtp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable create pinholes for RTP traffic to traverse firewall.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_algorithm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

							Description: "Relative strength of encryption algorithms accepted in negotiation.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_auth_client": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Require a client certificate and authenticate it with the peer/peergrp.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_auth_server": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Authenticate the server's certificate with the peer/peergrp.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_client_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name of Certificate to offer to server if requested.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_client_renegotiation": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "secure"}, false),

							Description: "Allow/block client renegotiation by server.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_max_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Highest SSL/TLS version to negotiate.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_min_version": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ssl-3.0", "tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

							Description: "Lowest SSL/TLS version to negotiate.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_mode": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"off", "full"}, false),

							Description: "SSL/TLS mode for encryption & decryption of traffic.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_pfs": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"require", "deny", "allow"}, false),

							Description: "SSL Perfect Forward Secrecy.",
							Optional:    true,
							Computed:    true,
						},
						"ssl_send_empty_frags": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Send empty fragments to avoid attack on CBC IV (SSL 3.0 & TLS 1.0 only).",
							Optional:    true,
							Computed:    true,
						},
						"ssl_server_certificate": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Name of Certificate return to the client in every SSL connection.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable SIP.",
							Optional:    true,
							Computed:    true,
						},
						"strict_register": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable only allow the registrar to connect.",
							Optional:    true,
							Computed:    true,
						},
						"subscribe_rate": {
							Type: schema.TypeInt,

							Description: "SUBSCRIBE request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"subscribe_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
						"unknown_header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"discard", "pass", "respond"}, false),

							Description: "Action for unknown SIP header.",
							Optional:    true,
							Computed:    true,
						},
						"update_rate": {
							Type: schema.TypeInt,

							Description: "UPDATE request rate limit (per second, per policy).",
							Optional:    true,
							Computed:    true,
						},
						"update_rate_track": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "src-ip", "dest-ip"}, false),

							Description: "Track the packet protocol field.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceVoipProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VoipProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVoipProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVoipProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VoipProfile")
	}

	return resourceVoipProfileRead(ctx, d, meta)
}

func resourceVoipProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVoipProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVoipProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VoipProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VoipProfile")
	}

	return resourceVoipProfileRead(ctx, d, meta)
}

func resourceVoipProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVoipProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VoipProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVoipProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVoipProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VoipProfile resource: %v", err)
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

	diags := refreshObjectVoipProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVoipProfileMsrp(d *schema.ResourceData, v *models.VoipProfileMsrp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.VoipProfileMsrp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.LogViolations; tmp != nil {
				v["log_violations"] = *tmp
			}

			if tmp := cfg.MaxMsgSize; tmp != nil {
				v["max_msg_size"] = *tmp
			}

			if tmp := cfg.MaxMsgSizeAction; tmp != nil {
				v["max_msg_size_action"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenVoipProfileSccp(d *schema.ResourceData, v *models.VoipProfileSccp, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.VoipProfileSccp{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.BlockMcast; tmp != nil {
				v["block_mcast"] = *tmp
			}

			if tmp := cfg.LogCallSummary; tmp != nil {
				v["log_call_summary"] = *tmp
			}

			if tmp := cfg.LogViolations; tmp != nil {
				v["log_violations"] = *tmp
			}

			if tmp := cfg.MaxCalls; tmp != nil {
				v["max_calls"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.VerifyHeader; tmp != nil {
				v["verify_header"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenVoipProfileSip(d *schema.ResourceData, v *models.VoipProfileSip, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		v2 := []models.VoipProfileSip{*v}
		for i, cfg := range v2 {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AckRate; tmp != nil {
				v["ack_rate"] = *tmp
			}

			if tmp := cfg.AckRateTrack; tmp != nil {
				v["ack_rate_track"] = *tmp
			}

			if tmp := cfg.BlockAck; tmp != nil {
				v["block_ack"] = *tmp
			}

			if tmp := cfg.BlockBye; tmp != nil {
				v["block_bye"] = *tmp
			}

			if tmp := cfg.BlockCancel; tmp != nil {
				v["block_cancel"] = *tmp
			}

			if tmp := cfg.BlockGeoRedOptions; tmp != nil {
				v["block_geo_red_options"] = *tmp
			}

			if tmp := cfg.BlockInfo; tmp != nil {
				v["block_info"] = *tmp
			}

			if tmp := cfg.BlockInvite; tmp != nil {
				v["block_invite"] = *tmp
			}

			if tmp := cfg.BlockLongLines; tmp != nil {
				v["block_long_lines"] = *tmp
			}

			if tmp := cfg.BlockMessage; tmp != nil {
				v["block_message"] = *tmp
			}

			if tmp := cfg.BlockNotify; tmp != nil {
				v["block_notify"] = *tmp
			}

			if tmp := cfg.BlockOptions; tmp != nil {
				v["block_options"] = *tmp
			}

			if tmp := cfg.BlockPrack; tmp != nil {
				v["block_prack"] = *tmp
			}

			if tmp := cfg.BlockPublish; tmp != nil {
				v["block_publish"] = *tmp
			}

			if tmp := cfg.BlockRefer; tmp != nil {
				v["block_refer"] = *tmp
			}

			if tmp := cfg.BlockRegister; tmp != nil {
				v["block_register"] = *tmp
			}

			if tmp := cfg.BlockSubscribe; tmp != nil {
				v["block_subscribe"] = *tmp
			}

			if tmp := cfg.BlockUnknown; tmp != nil {
				v["block_unknown"] = *tmp
			}

			if tmp := cfg.BlockUpdate; tmp != nil {
				v["block_update"] = *tmp
			}

			if tmp := cfg.ByeRate; tmp != nil {
				v["bye_rate"] = *tmp
			}

			if tmp := cfg.ByeRateTrack; tmp != nil {
				v["bye_rate_track"] = *tmp
			}

			if tmp := cfg.CallKeepalive; tmp != nil {
				v["call_keepalive"] = *tmp
			}

			if tmp := cfg.CancelRate; tmp != nil {
				v["cancel_rate"] = *tmp
			}

			if tmp := cfg.CancelRateTrack; tmp != nil {
				v["cancel_rate_track"] = *tmp
			}

			if tmp := cfg.ContactFixup; tmp != nil {
				v["contact_fixup"] = *tmp
			}

			if tmp := cfg.HntRestrictSourceIp; tmp != nil {
				v["hnt_restrict_source_ip"] = *tmp
			}

			if tmp := cfg.HostedNatTraversal; tmp != nil {
				v["hosted_nat_traversal"] = *tmp
			}

			if tmp := cfg.InfoRate; tmp != nil {
				v["info_rate"] = *tmp
			}

			if tmp := cfg.InfoRateTrack; tmp != nil {
				v["info_rate_track"] = *tmp
			}

			if tmp := cfg.InviteRate; tmp != nil {
				v["invite_rate"] = *tmp
			}

			if tmp := cfg.InviteRateTrack; tmp != nil {
				v["invite_rate_track"] = *tmp
			}

			if tmp := cfg.IpsRtp; tmp != nil {
				v["ips_rtp"] = *tmp
			}

			if tmp := cfg.LogCallSummary; tmp != nil {
				v["log_call_summary"] = *tmp
			}

			if tmp := cfg.LogViolations; tmp != nil {
				v["log_violations"] = *tmp
			}

			if tmp := cfg.MalformedHeaderAllow; tmp != nil {
				v["malformed_header_allow"] = *tmp
			}

			if tmp := cfg.MalformedHeaderCallId; tmp != nil {
				v["malformed_header_call_id"] = *tmp
			}

			if tmp := cfg.MalformedHeaderContact; tmp != nil {
				v["malformed_header_contact"] = *tmp
			}

			if tmp := cfg.MalformedHeaderContentLength; tmp != nil {
				v["malformed_header_content_length"] = *tmp
			}

			if tmp := cfg.MalformedHeaderContentType; tmp != nil {
				v["malformed_header_content_type"] = *tmp
			}

			if tmp := cfg.MalformedHeaderCseq; tmp != nil {
				v["malformed_header_cseq"] = *tmp
			}

			if tmp := cfg.MalformedHeaderExpires; tmp != nil {
				v["malformed_header_expires"] = *tmp
			}

			if tmp := cfg.MalformedHeaderFrom; tmp != nil {
				v["malformed_header_from"] = *tmp
			}

			if tmp := cfg.MalformedHeaderMaxForwards; tmp != nil {
				v["malformed_header_max_forwards"] = *tmp
			}

			if tmp := cfg.MalformedHeaderNoProxyRequire; tmp != nil {
				v["malformed_header_no_proxy_require"] = *tmp
			}

			if tmp := cfg.MalformedHeaderNoRequire; tmp != nil {
				v["malformed_header_no_require"] = *tmp
			}

			if tmp := cfg.MalformedHeaderPAssertedIdentity; tmp != nil {
				v["malformed_header_p_asserted_identity"] = *tmp
			}

			if tmp := cfg.MalformedHeaderRack; tmp != nil {
				v["malformed_header_rack"] = *tmp
			}

			if tmp := cfg.MalformedHeaderRecordRoute; tmp != nil {
				v["malformed_header_record_route"] = *tmp
			}

			if tmp := cfg.MalformedHeaderRoute; tmp != nil {
				v["malformed_header_route"] = *tmp
			}

			if tmp := cfg.MalformedHeaderRseq; tmp != nil {
				v["malformed_header_rseq"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpA; tmp != nil {
				v["malformed_header_sdp_a"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpB; tmp != nil {
				v["malformed_header_sdp_b"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpC; tmp != nil {
				v["malformed_header_sdp_c"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpI; tmp != nil {
				v["malformed_header_sdp_i"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpK; tmp != nil {
				v["malformed_header_sdp_k"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpM; tmp != nil {
				v["malformed_header_sdp_m"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpO; tmp != nil {
				v["malformed_header_sdp_o"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpR; tmp != nil {
				v["malformed_header_sdp_r"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpS; tmp != nil {
				v["malformed_header_sdp_s"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpT; tmp != nil {
				v["malformed_header_sdp_t"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpV; tmp != nil {
				v["malformed_header_sdp_v"] = *tmp
			}

			if tmp := cfg.MalformedHeaderSdpZ; tmp != nil {
				v["malformed_header_sdp_z"] = *tmp
			}

			if tmp := cfg.MalformedHeaderTo; tmp != nil {
				v["malformed_header_to"] = *tmp
			}

			if tmp := cfg.MalformedHeaderVia; tmp != nil {
				v["malformed_header_via"] = *tmp
			}

			if tmp := cfg.MalformedRequestLine; tmp != nil {
				v["malformed_request_line"] = *tmp
			}

			if tmp := cfg.MaxBodyLength; tmp != nil {
				v["max_body_length"] = *tmp
			}

			if tmp := cfg.MaxDialogs; tmp != nil {
				v["max_dialogs"] = *tmp
			}

			if tmp := cfg.MaxIdleDialogs; tmp != nil {
				v["max_idle_dialogs"] = *tmp
			}

			if tmp := cfg.MaxLineLength; tmp != nil {
				v["max_line_length"] = *tmp
			}

			if tmp := cfg.MessageRate; tmp != nil {
				v["message_rate"] = *tmp
			}

			if tmp := cfg.MessageRateTrack; tmp != nil {
				v["message_rate_track"] = *tmp
			}

			if tmp := cfg.NatPortRange; tmp != nil {
				v["nat_port_range"] = *tmp
			}

			if tmp := cfg.NatTrace; tmp != nil {
				v["nat_trace"] = *tmp
			}

			if tmp := cfg.NoSdpFixup; tmp != nil {
				v["no_sdp_fixup"] = *tmp
			}

			if tmp := cfg.NotifyRate; tmp != nil {
				v["notify_rate"] = *tmp
			}

			if tmp := cfg.NotifyRateTrack; tmp != nil {
				v["notify_rate_track"] = *tmp
			}

			if tmp := cfg.OpenContactPinhole; tmp != nil {
				v["open_contact_pinhole"] = *tmp
			}

			if tmp := cfg.OpenRecordRoutePinhole; tmp != nil {
				v["open_record_route_pinhole"] = *tmp
			}

			if tmp := cfg.OpenRegisterPinhole; tmp != nil {
				v["open_register_pinhole"] = *tmp
			}

			if tmp := cfg.OpenViaPinhole; tmp != nil {
				v["open_via_pinhole"] = *tmp
			}

			if tmp := cfg.OptionsRate; tmp != nil {
				v["options_rate"] = *tmp
			}

			if tmp := cfg.OptionsRateTrack; tmp != nil {
				v["options_rate_track"] = *tmp
			}

			if tmp := cfg.PrackRate; tmp != nil {
				v["prack_rate"] = *tmp
			}

			if tmp := cfg.PrackRateTrack; tmp != nil {
				v["prack_rate_track"] = *tmp
			}

			if tmp := cfg.PreserveOverride; tmp != nil {
				v["preserve_override"] = *tmp
			}

			if tmp := cfg.ProvisionalInviteExpiryTime; tmp != nil {
				v["provisional_invite_expiry_time"] = *tmp
			}

			if tmp := cfg.PublishRate; tmp != nil {
				v["publish_rate"] = *tmp
			}

			if tmp := cfg.PublishRateTrack; tmp != nil {
				v["publish_rate_track"] = *tmp
			}

			if tmp := cfg.ReferRate; tmp != nil {
				v["refer_rate"] = *tmp
			}

			if tmp := cfg.ReferRateTrack; tmp != nil {
				v["refer_rate_track"] = *tmp
			}

			if tmp := cfg.RegisterContactTrace; tmp != nil {
				v["register_contact_trace"] = *tmp
			}

			if tmp := cfg.RegisterRate; tmp != nil {
				v["register_rate"] = *tmp
			}

			if tmp := cfg.RegisterRateTrack; tmp != nil {
				v["register_rate_track"] = *tmp
			}

			if tmp := cfg.Rfc2543Branch; tmp != nil {
				v["rfc2543_branch"] = *tmp
			}

			if tmp := cfg.Rtp; tmp != nil {
				v["rtp"] = *tmp
			}

			if tmp := cfg.SslAlgorithm; tmp != nil {
				v["ssl_algorithm"] = *tmp
			}

			if tmp := cfg.SslAuthClient; tmp != nil {
				v["ssl_auth_client"] = *tmp
			}

			if tmp := cfg.SslAuthServer; tmp != nil {
				v["ssl_auth_server"] = *tmp
			}

			if tmp := cfg.SslClientCertificate; tmp != nil {
				v["ssl_client_certificate"] = *tmp
			}

			if tmp := cfg.SslClientRenegotiation; tmp != nil {
				v["ssl_client_renegotiation"] = *tmp
			}

			if tmp := cfg.SslMaxVersion; tmp != nil {
				v["ssl_max_version"] = *tmp
			}

			if tmp := cfg.SslMinVersion; tmp != nil {
				v["ssl_min_version"] = *tmp
			}

			if tmp := cfg.SslMode; tmp != nil {
				v["ssl_mode"] = *tmp
			}

			if tmp := cfg.SslPfs; tmp != nil {
				v["ssl_pfs"] = *tmp
			}

			if tmp := cfg.SslSendEmptyFrags; tmp != nil {
				v["ssl_send_empty_frags"] = *tmp
			}

			if tmp := cfg.SslServerCertificate; tmp != nil {
				v["ssl_server_certificate"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.StrictRegister; tmp != nil {
				v["strict_register"] = *tmp
			}

			if tmp := cfg.SubscribeRate; tmp != nil {
				v["subscribe_rate"] = *tmp
			}

			if tmp := cfg.SubscribeRateTrack; tmp != nil {
				v["subscribe_rate_track"] = *tmp
			}

			if tmp := cfg.UnknownHeader; tmp != nil {
				v["unknown_header"] = *tmp
			}

			if tmp := cfg.UpdateRate; tmp != nil {
				v["update_rate"] = *tmp
			}

			if tmp := cfg.UpdateRateTrack; tmp != nil {
				v["update_rate_track"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectVoipProfile(d *schema.ResourceData, o *models.VoipProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.FeatureSet != nil {
		v := *o.FeatureSet

		if err = d.Set("feature_set", v); err != nil {
			return diag.Errorf("error reading feature_set: %v", err)
		}
	}

	if _, ok := d.GetOk("msrp"); ok {
		if o.Msrp != nil {
			if err = d.Set("msrp", flattenVoipProfileMsrp(d, o.Msrp, "msrp", sort)); err != nil {
				return diag.Errorf("error reading msrp: %v", err)
			}
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if _, ok := d.GetOk("sccp"); ok {
		if o.Sccp != nil {
			if err = d.Set("sccp", flattenVoipProfileSccp(d, o.Sccp, "sccp", sort)); err != nil {
				return diag.Errorf("error reading sccp: %v", err)
			}
		}
	}

	if _, ok := d.GetOk("sip"); ok {
		if o.Sip != nil {
			if err = d.Set("sip", flattenVoipProfileSip(d, o.Sip, "sip", sort)); err != nil {
				return diag.Errorf("error reading sip: %v", err)
			}
		}
	}

	return nil
}

func expandVoipProfileMsrp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.VoipProfileMsrp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VoipProfileMsrp

	for i := range l {
		tmp := models.VoipProfileMsrp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.log_violations", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogViolations = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_msg_size", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxMsgSize = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_msg_size_action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MaxMsgSizeAction = &v2
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

func expandVoipProfileSccp(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.VoipProfileSccp, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VoipProfileSccp

	for i := range l {
		tmp := models.VoipProfileSccp{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.block_mcast", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockMcast = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_call_summary", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogCallSummary = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_violations", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogViolations = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_calls", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxCalls = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.verify_header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.VerifyHeader = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func expandVoipProfileSip(d *schema.ResourceData, v interface{}, pre string, sv string) (*models.VoipProfileSip, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VoipProfileSip

	for i := range l {
		tmp := models.VoipProfileSip{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ack_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.AckRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ack_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AckRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_ack", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockAck = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_bye", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockBye = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_cancel", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockCancel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_geo_red_options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockGeoRedOptions = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_info", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockInfo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_invite", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockInvite = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_long_lines", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockLongLines = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_message", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockMessage = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_notify", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockNotify = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockOptions = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_prack", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockPrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_publish", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockPublish = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_refer", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockRefer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_register", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockRegister = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_subscribe", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockSubscribe = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_unknown", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockUnknown = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block_update", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BlockUpdate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bye_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ByeRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.bye_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ByeRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.call_keepalive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CallKeepalive = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cancel_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.CancelRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.cancel_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CancelRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.contact_fixup", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ContactFixup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hnt_restrict_source_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HntRestrictSourceIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hosted_nat_traversal", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HostedNatTraversal = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.info_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.InfoRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.info_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InfoRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invite_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.InviteRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.invite_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InviteRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ips_rtp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.IpsRtp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_call_summary", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogCallSummary = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log_violations", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LogViolations = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_allow", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderAllow = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_call_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderCallId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_contact", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderContact = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_content_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderContentLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_content_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderContentType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_cseq", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderCseq = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_expires", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderExpires = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_from", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderFrom = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_max_forwards", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderMaxForwards = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_no_proxy_require", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderNoProxyRequire = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_no_require", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderNoRequire = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_p_asserted_identity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderPAssertedIdentity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_rack", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderRack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_record_route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderRecordRoute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_route", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderRoute = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_rseq", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderRseq = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_a", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpA = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_b", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpB = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_c", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpC = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_i", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpI = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_k", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpK = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_m", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpM = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_o", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpO = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_r", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpR = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_s", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpS = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_t", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpT = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_v", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpV = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_sdp_z", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderSdpZ = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_to", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderTo = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_header_via", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedHeaderVia = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.malformed_request_line", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MalformedRequestLine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_body_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxBodyLength = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_dialogs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxDialogs = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_idle_dialogs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxIdleDialogs = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.max_line_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MaxLineLength = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.message_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.MessageRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.message_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.MessageRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nat_port_range", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NatPortRange = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nat_trace", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NatTrace = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.no_sdp_fixup", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NoSdpFixup = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.notify_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.NotifyRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.notify_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NotifyRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.open_contact_pinhole", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OpenContactPinhole = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.open_record_route_pinhole", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OpenRecordRoutePinhole = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.open_register_pinhole", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OpenRegisterPinhole = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.open_via_pinhole", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OpenViaPinhole = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.OptionsRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.options_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.OptionsRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prack_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PrackRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prack_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrackRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preserve_override", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PreserveOverride = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.provisional_invite_expiry_time", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ProvisionalInviteExpiryTime = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.publish_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PublishRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.publish_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PublishRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.refer_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.ReferRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.refer_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ReferRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_contact_trace", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RegisterContactTrace = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RegisterRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.register_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RegisterRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rfc2543_branch", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Rfc2543Branch = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rtp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Rtp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_algorithm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslAlgorithm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_auth_client", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslAuthClient = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_auth_server", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslAuthServer = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_client_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslClientCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_client_renegotiation", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslClientRenegotiation = &v2
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

		pre_append = fmt.Sprintf("%s.%d.ssl_mode", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslMode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_pfs", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslPfs = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_send_empty_frags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslSendEmptyFrags = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ssl_server_certificate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SslServerCertificate = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.strict_register", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StrictRegister = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.subscribe_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.SubscribeRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.subscribe_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SubscribeRateTrack = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unknown_header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UnknownHeader = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.update_rate", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.UpdateRate = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.update_rate_track", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UpdateRateTrack = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result[0], nil
}

func getObjectVoipProfile(d *schema.ResourceData, sv string) (*models.VoipProfile, diag.Diagnostics) {
	obj := models.VoipProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("feature_set"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("feature_set", sv)
				diags = append(diags, e)
			}
			obj.FeatureSet = &v2
		}
	}
	if v, ok := d.GetOk("msrp"); ok {
		if !utils.CheckVer(sv, "v7.0.2", "") {
			e := utils.AttributeVersionWarning("msrp", sv)
			diags = append(diags, e)
		}
		t, err := expandVoipProfileMsrp(d, v, "msrp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Msrp = t
		}
	} else if d.HasChange("msrp") {
		old, new := d.GetChange("msrp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Msrp = &models.VoipProfileMsrp{}
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
	if v, ok := d.GetOk("sccp"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sccp", sv)
			diags = append(diags, e)
		}
		t, err := expandVoipProfileSccp(d, v, "sccp", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Sccp = t
		}
	} else if d.HasChange("sccp") {
		old, new := d.GetChange("sccp")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Sccp = &models.VoipProfileSccp{}
		}
	}
	if v, ok := d.GetOk("sip"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("sip", sv)
			diags = append(diags, e)
		}
		t, err := expandVoipProfileSip(d, v, "sip", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Sip = t
		}
	} else if d.HasChange("sip") {
		old, new := d.GetChange("sip")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Sip = &models.VoipProfileSip{}
		}
	}
	return &obj, diags
}
