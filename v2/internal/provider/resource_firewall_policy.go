// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4/IPv6 policies.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4/IPv6 policies.",

		CreateContext: resourceFirewallPolicyCreate,
		ReadContext:   resourceFirewallPolicyRead,
		UpdateContext: resourceFirewallPolicyUpdate,
		DeleteContext: resourceFirewallPolicyDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "deny", "ipsec"}, false),

				Description: "Policy action (accept/deny/ipsec).",
				Optional:    true,
				Computed:    true,
			},
			"anti_replay": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable anti-replay check.",
				Optional:    true,
				Computed:    true,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "Application category ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Category IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"app_group": {
				Type:        schema.TypeList,
				Description: "Application group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Application group names.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Application IDs.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"application_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Application list.",
				Optional:    true,
				Computed:    true,
			},
			"auth_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "HTTPS server certificate for policy authentication.",
				Optional:    true,
				Computed:    true,
			},
			"auth_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authentication-based routing.",
				Optional:    true,
				Computed:    true,
			},
			"auth_redirect_addr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "HTTP-to-HTTPS redirect address for firewall authentication.",
				Optional:    true,
				Computed:    true,
			},
			"auto_asic_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy traffic ASIC offloading.",
				Optional:    true,
				Computed:    true,
			},
			"av_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Antivirus profile.",
				Optional:    true,
				Computed:    true,
			},
			"block_notification": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable block notification.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal_exempt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to exempt some users from the captive portal.",
				Optional:    true,
				Computed:    true,
			},
			"capture_packet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable capture packets.",
				Optional:    true,
				Computed:    true,
			},
			"cifs_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing CIFS profile.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"custom_log_fields": {
				Type:        schema.TypeList,
				Description: "Custom fields to append to log messages for this policy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Custom log field.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"decrypted_traffic_mirror": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Decrypted traffic mirror.",
				Optional:    true,
				Computed:    true,
			},
			"delay_tcp_npu_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable TCP NPU session delay to guarantee packet order of 3-way handshake.",
				Optional:    true,
				Computed:    true,
			},
			"diffserv_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to change packet's DiffServ values to the specified diffservcode-forward value.",
				Optional:    true,
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode_forward": {
				Type: schema.TypeString,

				Description: "Change packet's DiffServ to this value.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode_rev": {
				Type: schema.TypeString,

				Description: "Change packet's reverse (reply) DiffServ to this value.",
				Optional:    true,
				Computed:    true,
			},
			"disclaimer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable user authentication disclaimer.",
				Optional:    true,
				Computed:    true,
			},
			"dlp_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DLP sensor.",
				Optional:    true,
				Computed:    true,
			},
			"dnsfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing DNS filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"dsri": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable DSRI to ignore HTTP server responses.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address and address group names.",
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
			"dstaddr_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Destination IPv6 address name and address group names.",
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
			"dstintf": {
				Type:        schema.TypeList,
				Description: "Outgoing (egress) interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dynamic_shaping": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable dynamic RADIUS defined traffic shaping.",
				Optional:    true,
				Computed:    true,
			},
			"email_collect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable email collection.",
				Optional:    true,
				Computed:    true,
			},
			"emailfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing email filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"fec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Forward Error Correction on traffic matching this policy on a FEC device.",
				Optional:    true,
				Computed:    true,
			},
			"file_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing file-filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"firewall_session_dirty": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"check-all", "check-new"}, false),

				Description: "How to handle sessions if the configuration of this firewall policy changes.",
				Optional:    true,
				Computed:    true,
			},
			"fixedport": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to prevent source NAT from changing a session's source port.",
				Optional:    true,
				Computed:    true,
			},
			"fsso": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Fortinet Single Sign-On.",
				Optional:    true,
				Computed:    true,
			},
			"fsso_agent_for_ntlm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "FSSO agent to use for NTLM authentication.",
				Optional:    true,
				Computed:    true,
			},
			"fsso_groups": {
				Type:        schema.TypeList,
				Description: "Names of FSSO groups.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "Names of FSSO groups.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"geoip_anycast": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable recognition of anycast IP addresses using the geography IP database.",
				Optional:    true,
				Computed:    true,
			},
			"geoip_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"physical-location", "registered-location"}, false),

				Description: "Match geography address based either on its physical location or registered location.",
				Optional:    true,
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Names of user groups that can authenticate with this policy.",
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
			"http_policy_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Redirect HTTP(S) traffic to matching transparent web proxy policy.",
				Optional:    true,
				Computed:    true,
			},
			"icap_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing ICAP profile.",
				Optional:    true,
				Computed:    true,
			},
			"identity_based_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of identity-based routing rule.",
				Optional:    true,
				Computed:    true,
			},
			"inbound": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.",
				Optional:    true,
				Computed:    true,
			},
			"inspection_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"proxy", "flow"}, false),

				Description: "Policy inspection mode (Flow/proxy). Default is Flow mode.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. ",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Internet Service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled internet-service specifies what the service must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_src": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. ",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_src_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_group": {
				Type:        schema.TypeList,
				Description: "Internet Service source group name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_id": {
				Type:        schema.TypeList,
				Description: "Internet Service source ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Internet Service ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_name": {
				Type:        schema.TypeList,
				Description: "Internet Service source name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled internet-service-src specifies what the service must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"ippool": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to use IP Pools for source NAT.",
				Optional:    true,
				Computed:    true,
			},
			"ips_sensor": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing IPS sensor.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "utm", "disable"}, false),

				Description: "Enable or disable logging. Log all sessions or security profile sessions.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic_start": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Record logs when a session starts.",
				Optional:    true,
				Computed:    true,
			},
			"match_vip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to match packets that have had their destination addresses changed by a VIP.",
				Optional:    true,
				Computed:    true,
			},
			"match_vip_only": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable matching of only those packets that have had their destination addresses changed by a VIP.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Policy name.",
				Optional:    true,
				Computed:    true,
			},
			"nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable source NAT.",
				Optional:    true,
				Computed:    true,
			},
			"nat46": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT46.",
				Optional:    true,
				Computed:    true,
			},
			"nat64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT64.",
				Optional:    true,
				Computed:    true,
			},
			"natinbound": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Policy-based IPsec VPN: apply destination NAT to inbound traffic.",
				Optional:    true,
				Computed:    true,
			},
			"natip": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Policy-based IPsec VPN: source NAT IP address for outgoing traffic.",
				Optional:    true,
				Computed:    true,
			},
			"natoutbound": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Policy-based IPsec VPN: apply source NAT to outbound traffic.",
				Optional:    true,
				Computed:    true,
			},
			"np_acceleration": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable UTM Network Processor acceleration.",
				Optional:    true,
				Computed:    true,
			},
			"ntlm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NTLM authentication.",
				Optional:    true,
				Computed:    true,
			},
			"ntlm_enabled_browsers": {
				Type:        schema.TypeList,
				Description: "HTTP-User-Agent value of supported browsers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_agent_string": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "User agent string.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ntlm_guest": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NTLM guest user access.",
				Optional:    true,
				Computed:    true,
			},
			"outbound": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.",
				Optional:    true,
				Computed:    true,
			},
			"passive_wan_health_measurement": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled.",
				Optional:    true,
				Computed:    true,
			},
			"per_ip_shaper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Per-IP traffic shaper.",
				Optional:    true,
				Computed:    true,
			},
			"permit_any_host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Accept UDP packets from any host.",
				Optional:    true,
				Computed:    true,
			},
			"permit_stun_host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host.",
				Optional:    true,
				Computed:    true,
			},
			"policyid": {
				Type: schema.TypeInt,

				Description: "Policy ID (0 - 4294967294).",
				ForceNew:    true,
				Required:    true,
			},
			"poolname": {
				Type:        schema.TypeList,
				Description: "IP Pool names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP pool name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"poolname6": {
				Type:        schema.TypeList,
				Description: "IPv6 pool names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPv6 pool name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"profile_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of profile group.",
				Optional:    true,
				Computed:    true,
			},
			"profile_protocol_options": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Protocol options profile.",
				Optional:    true,
				Computed:    true,
			},
			"profile_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"single", "group"}, false),

				Description: "Determine whether the firewall policy allows security profile groups or single profiles only.",
				Optional:    true,
				Computed:    true,
			},
			"radius_mac_auth_bypass": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server.",
				Optional:    true,
				Computed:    true,
			},
			"redirect_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "URL users are directed to after seeing and accepting the disclaimer or authenticating.",
				Optional:    true,
				Computed:    true,
			},
			"replacemsg_override_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Override the default replacement message group for this policy.",
				Optional:    true,
				Computed:    true,
			},
			"reputation_direction": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"source", "destination"}, false),

				Description: "Direction of the initial traffic for reputation to take effect.",
				Optional:    true,
				Computed:    true,
			},
			"reputation_minimum": {
				Type: schema.TypeInt,

				Description: "Minimum Reputation to take action.",
				Optional:    true,
				Computed:    true,
			},
			"rsso": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RADIUS single sign-on (RSSO).",
				Optional:    true,
				Computed:    true,
			},
			"rtp_addr": {
				Type:        schema.TypeList,
				Description: "Address names if this is an RTP NAT policy.",
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
			"rtp_nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable Real Time Protocol (RTP) NAT.",
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Schedule name.",
				Optional:    true,
				Computed:    true,
			},
			"schedule_timeout": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity.",
				Optional:    true,
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SCTP filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"send_deny_packet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to send a reply when a session is denied or blocked by a firewall policy.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Service and service group names.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"service_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled service specifies what the service must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"session_ttl": {
				Type: schema.TypeString,

				Description: "TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).",
				Optional:    true,
				Computed:    true,
			},
			"sgt": {
				Type:        schema.TypeList,
				Description: "Security group tags.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Security group tag (1 - 65535).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"sgt_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable security group tags (SGT) check.",
				Optional:    true,
				Computed:    true,
			},
			"src_vendor_mac": {
				Type:        schema.TypeList,
				Description: "Vendor MAC source ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Vendor MAC ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address and address group names.",
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
			"srcaddr_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled srcaddr/srcaddr6 specifies what the source address must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Source IPv6 address name and address group names.",
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
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Incoming (ingress) interface.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssh_filter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SSH filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_policy_redirect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Redirect SSH traffic to matching transparent proxy policy.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_mirror": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_mirror_intf": {
				Type:        schema.TypeList,
				Description: "SSL mirror interface name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Mirror Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ssl_ssh_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing SSL SSH profile.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable or disable this policy.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_mss_receiver": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Receiver TCP maximum segment size (MSS).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_mss_sender": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Sender TCP maximum segment size (MSS).",
				Optional:    true,
				Computed:    true,
			},
			"tcp_session_without_syn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"all", "data-only", "disable"}, false),

				Description: "Enable/disable creation of TCP session without SYN flag.",
				Optional:    true,
				Computed:    true,
			},
			"timeout_send_rst": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending RST packets when TCP sessions expire.",
				Optional:    true,
				Computed:    true,
			},
			"tos": {
				Type: schema.TypeString,

				Description: "ToS (Type of Service) value used for comparison.",
				Optional:    true,
				Computed:    true,
			},
			"tos_mask": {
				Type: schema.TypeString,

				Description: "Non-zero bit positions are used for comparison while zero bit positions are ignored.",
				Optional:    true,
				Computed:    true,
			},
			"tos_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable negated TOS match.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_shaper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Traffic shaper.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_shaper_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Reverse traffic shaper.",
				Optional:    true,
				Computed:    true,
			},
			"url_category": {
				Type:        schema.TypeList,
				Description: "URL category ID list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "URL category ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Names of individual users that can authenticate with this policy.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Names of individual users that can authenticate with this policy.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"utm_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.",
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"videofilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VideoFilter profile.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_cos_fwd": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),

				Description: "VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_cos_rev": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),

				Description: "VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.",
				Optional:    true,
				Computed:    true,
			},
			"vlan_filter": {
				Type: schema.TypeString,

				Description: "Set VLAN filters.",
				Optional:    true,
				Computed:    true,
			},
			"voip_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing VoIP profile.",
				Optional:    true,
				Computed:    true,
			},
			"vpntunnel": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Policy-based IPsec VPN: name of the IPsec VPN Phase 1.",
				Optional:    true,
				Computed:    true,
			},
			"waf_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web application firewall profile.",
				Optional:    true,
				Computed:    true,
			},
			"wanopt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WAN optimization.",
				Optional:    true,
				Computed:    true,
			},
			"wanopt_detection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"active", "passive", "off"}, false),

				Description: "WAN optimization auto-detection mode.",
				Optional:    true,
				Computed:    true,
			},
			"wanopt_passive_opt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "transparent", "non-transparent"}, false),

				Description: "WAN optimization passive mode options. This option decides what IP address will be used to connect server.",
				Optional:    true,
				Computed:    true,
			},
			"wanopt_peer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WAN optimization peer.",
				Optional:    true,
				Computed:    true,
			},
			"wanopt_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "WAN optimization profile.",
				Optional:    true,
				Computed:    true,
			},
			"wccp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable forwarding traffic matching this policy to a configured WCCP server.",
				Optional:    true,
				Computed:    true,
			},
			"webcache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable web cache.",
				Optional:    true,
				Computed:    true,
			},
			"webcache_https": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable web cache for HTTPS.",
				Optional:    true,
				Computed:    true,
			},
			"webfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Web filter profile.",
				Optional:    true,
				Computed:    true,
			},
			"webproxy_forward_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Webproxy forward server name.",
				Optional:    true,
				Computed:    true,
			},
			"webproxy_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Webproxy profile name.",
				Optional:    true,
				Computed:    true,
			},
			"wsso": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WiFi Single Sign On (WSSO).",
				Optional:    true,
				Computed:    true,
			},
			"ztna_ems_tag": {
				Type:        schema.TypeList,
				Description: "Source ztna-ems-tag names.",
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
			"ztna_geo_tag": {
				Type:        schema.TypeList,
				Description: "Source ztna-geo-tag names.",
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
			"ztna_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable zero trust access.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "policyid"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallPolicy")
	}

	return resourceFirewallPolicyRead(ctx, d, meta)
}

func resourceFirewallPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallPolicy")
	}

	return resourceFirewallPolicyRead(ctx, d, meta)
}

func resourceFirewallPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallPolicy resource: %v", err)
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

	diags := refreshObjectFirewallPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallPolicyAppCategory(v *[]models.FirewallPolicyAppCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicyAppGroup(v *[]models.FirewallPolicyAppGroup, sort bool) interface{} {
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

func flattenFirewallPolicyApplication(v *[]models.FirewallPolicyApplication, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicyCustomLogFields(v *[]models.FirewallPolicyCustomLogFields, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.FieldId; tmp != nil {
				v["field_id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "field_id")
	}

	return flat
}

func flattenFirewallPolicyDstaddr(v *[]models.FirewallPolicyDstaddr, sort bool) interface{} {
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

func flattenFirewallPolicyDstaddr6(v *[]models.FirewallPolicyDstaddr6, sort bool) interface{} {
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

func flattenFirewallPolicyDstintf(v *[]models.FirewallPolicyDstintf, sort bool) interface{} {
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

func flattenFirewallPolicyFssoGroups(v *[]models.FirewallPolicyFssoGroups, sort bool) interface{} {
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

func flattenFirewallPolicyGroups(v *[]models.FirewallPolicyGroups, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceCustom(v *[]models.FirewallPolicyInternetServiceCustom, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceCustomGroup(v *[]models.FirewallPolicyInternetServiceCustomGroup, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceGroup(v *[]models.FirewallPolicyInternetServiceGroup, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceId(v *[]models.FirewallPolicyInternetServiceId, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicyInternetServiceName(v *[]models.FirewallPolicyInternetServiceName, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceSrcCustom(v *[]models.FirewallPolicyInternetServiceSrcCustom, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceSrcCustomGroup(v *[]models.FirewallPolicyInternetServiceSrcCustomGroup, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceSrcGroup(v *[]models.FirewallPolicyInternetServiceSrcGroup, sort bool) interface{} {
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

func flattenFirewallPolicyInternetServiceSrcId(v *[]models.FirewallPolicyInternetServiceSrcId, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicyInternetServiceSrcName(v *[]models.FirewallPolicyInternetServiceSrcName, sort bool) interface{} {
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

func flattenFirewallPolicyNtlmEnabledBrowsers(v *[]models.FirewallPolicyNtlmEnabledBrowsers, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.UserAgentString; tmp != nil {
				v["user_agent_string"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "user_agent_string")
	}

	return flat
}

func flattenFirewallPolicyPoolname(v *[]models.FirewallPolicyPoolname, sort bool) interface{} {
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

func flattenFirewallPolicyPoolname6(v *[]models.FirewallPolicyPoolname6, sort bool) interface{} {
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

func flattenFirewallPolicyRtpAddr(v *[]models.FirewallPolicyRtpAddr, sort bool) interface{} {
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

func flattenFirewallPolicyService(v *[]models.FirewallPolicyService, sort bool) interface{} {
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

func flattenFirewallPolicySgt(v *[]models.FirewallPolicySgt, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicySrcVendorMac(v *[]models.FirewallPolicySrcVendorMac, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicySrcaddr(v *[]models.FirewallPolicySrcaddr, sort bool) interface{} {
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

func flattenFirewallPolicySrcaddr6(v *[]models.FirewallPolicySrcaddr6, sort bool) interface{} {
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

func flattenFirewallPolicySrcintf(v *[]models.FirewallPolicySrcintf, sort bool) interface{} {
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

func flattenFirewallPolicySslMirrorIntf(v *[]models.FirewallPolicySslMirrorIntf, sort bool) interface{} {
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

func flattenFirewallPolicyUrlCategory(v *[]models.FirewallPolicyUrlCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallPolicyUsers(v *[]models.FirewallPolicyUsers, sort bool) interface{} {
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

func flattenFirewallPolicyZtnaEmsTag(v *[]models.FirewallPolicyZtnaEmsTag, sort bool) interface{} {
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

func flattenFirewallPolicyZtnaGeoTag(v *[]models.FirewallPolicyZtnaGeoTag, sort bool) interface{} {
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

func refreshObjectFirewallPolicy(d *schema.ResourceData, o *models.FirewallPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.AntiReplay != nil {
		v := *o.AntiReplay

		if err = d.Set("anti_replay", v); err != nil {
			return diag.Errorf("error reading anti_replay: %v", err)
		}
	}

	if o.AppCategory != nil {
		if err = d.Set("app_category", flattenFirewallPolicyAppCategory(o.AppCategory, sort)); err != nil {
			return diag.Errorf("error reading app_category: %v", err)
		}
	}

	if o.AppGroup != nil {
		if err = d.Set("app_group", flattenFirewallPolicyAppGroup(o.AppGroup, sort)); err != nil {
			return diag.Errorf("error reading app_group: %v", err)
		}
	}

	if o.Application != nil {
		if err = d.Set("application", flattenFirewallPolicyApplication(o.Application, sort)); err != nil {
			return diag.Errorf("error reading application: %v", err)
		}
	}

	if o.ApplicationList != nil {
		v := *o.ApplicationList

		if err = d.Set("application_list", v); err != nil {
			return diag.Errorf("error reading application_list: %v", err)
		}
	}

	if o.AuthCert != nil {
		v := *o.AuthCert

		if err = d.Set("auth_cert", v); err != nil {
			return diag.Errorf("error reading auth_cert: %v", err)
		}
	}

	if o.AuthPath != nil {
		v := *o.AuthPath

		if err = d.Set("auth_path", v); err != nil {
			return diag.Errorf("error reading auth_path: %v", err)
		}
	}

	if o.AuthRedirectAddr != nil {
		v := *o.AuthRedirectAddr

		if err = d.Set("auth_redirect_addr", v); err != nil {
			return diag.Errorf("error reading auth_redirect_addr: %v", err)
		}
	}

	if o.AutoAsicOffload != nil {
		v := *o.AutoAsicOffload

		if err = d.Set("auto_asic_offload", v); err != nil {
			return diag.Errorf("error reading auto_asic_offload: %v", err)
		}
	}

	if o.AvProfile != nil {
		v := *o.AvProfile

		if err = d.Set("av_profile", v); err != nil {
			return diag.Errorf("error reading av_profile: %v", err)
		}
	}

	if o.BlockNotification != nil {
		v := *o.BlockNotification

		if err = d.Set("block_notification", v); err != nil {
			return diag.Errorf("error reading block_notification: %v", err)
		}
	}

	if o.CaptivePortalExempt != nil {
		v := *o.CaptivePortalExempt

		if err = d.Set("captive_portal_exempt", v); err != nil {
			return diag.Errorf("error reading captive_portal_exempt: %v", err)
		}
	}

	if o.CapturePacket != nil {
		v := *o.CapturePacket

		if err = d.Set("capture_packet", v); err != nil {
			return diag.Errorf("error reading capture_packet: %v", err)
		}
	}

	if o.CifsProfile != nil {
		v := *o.CifsProfile

		if err = d.Set("cifs_profile", v); err != nil {
			return diag.Errorf("error reading cifs_profile: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.CustomLogFields != nil {
		if err = d.Set("custom_log_fields", flattenFirewallPolicyCustomLogFields(o.CustomLogFields, sort)); err != nil {
			return diag.Errorf("error reading custom_log_fields: %v", err)
		}
	}

	if o.DecryptedTrafficMirror != nil {
		v := *o.DecryptedTrafficMirror

		if err = d.Set("decrypted_traffic_mirror", v); err != nil {
			return diag.Errorf("error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if o.DelayTcpNpuSession != nil {
		v := *o.DelayTcpNpuSession

		if err = d.Set("delay_tcp_npu_session", v); err != nil {
			return diag.Errorf("error reading delay_tcp_npu_session: %v", err)
		}
	}

	if o.DiffservForward != nil {
		v := *o.DiffservForward

		if err = d.Set("diffserv_forward", v); err != nil {
			return diag.Errorf("error reading diffserv_forward: %v", err)
		}
	}

	if o.DiffservReverse != nil {
		v := *o.DiffservReverse

		if err = d.Set("diffserv_reverse", v); err != nil {
			return diag.Errorf("error reading diffserv_reverse: %v", err)
		}
	}

	if o.DiffservcodeForward != nil {
		v := *o.DiffservcodeForward

		if err = d.Set("diffservcode_forward", v); err != nil {
			return diag.Errorf("error reading diffservcode_forward: %v", err)
		}
	}

	if o.DiffservcodeRev != nil {
		v := *o.DiffservcodeRev

		if err = d.Set("diffservcode_rev", v); err != nil {
			return diag.Errorf("error reading diffservcode_rev: %v", err)
		}
	}

	if o.Disclaimer != nil {
		v := *o.Disclaimer

		if err = d.Set("disclaimer", v); err != nil {
			return diag.Errorf("error reading disclaimer: %v", err)
		}
	}

	if o.DlpSensor != nil {
		v := *o.DlpSensor

		if err = d.Set("dlp_sensor", v); err != nil {
			return diag.Errorf("error reading dlp_sensor: %v", err)
		}
	}

	if o.DnsfilterProfile != nil {
		v := *o.DnsfilterProfile

		if err = d.Set("dnsfilter_profile", v); err != nil {
			return diag.Errorf("error reading dnsfilter_profile: %v", err)
		}
	}

	if o.Dsri != nil {
		v := *o.Dsri

		if err = d.Set("dsri", v); err != nil {
			return diag.Errorf("error reading dsri: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenFirewallPolicyDstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.DstaddrNegate != nil {
		v := *o.DstaddrNegate

		if err = d.Set("dstaddr_negate", v); err != nil {
			return diag.Errorf("error reading dstaddr_negate: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenFirewallPolicyDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenFirewallPolicyDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.DynamicShaping != nil {
		v := *o.DynamicShaping

		if err = d.Set("dynamic_shaping", v); err != nil {
			return diag.Errorf("error reading dynamic_shaping: %v", err)
		}
	}

	if o.EmailCollect != nil {
		v := *o.EmailCollect

		if err = d.Set("email_collect", v); err != nil {
			return diag.Errorf("error reading email_collect: %v", err)
		}
	}

	if o.EmailfilterProfile != nil {
		v := *o.EmailfilterProfile

		if err = d.Set("emailfilter_profile", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile: %v", err)
		}
	}

	if o.Fec != nil {
		v := *o.Fec

		if err = d.Set("fec", v); err != nil {
			return diag.Errorf("error reading fec: %v", err)
		}
	}

	if o.FileFilterProfile != nil {
		v := *o.FileFilterProfile

		if err = d.Set("file_filter_profile", v); err != nil {
			return diag.Errorf("error reading file_filter_profile: %v", err)
		}
	}

	if o.FirewallSessionDirty != nil {
		v := *o.FirewallSessionDirty

		if err = d.Set("firewall_session_dirty", v); err != nil {
			return diag.Errorf("error reading firewall_session_dirty: %v", err)
		}
	}

	if o.Fixedport != nil {
		v := *o.Fixedport

		if err = d.Set("fixedport", v); err != nil {
			return diag.Errorf("error reading fixedport: %v", err)
		}
	}

	if o.Fsso != nil {
		v := *o.Fsso

		if err = d.Set("fsso", v); err != nil {
			return diag.Errorf("error reading fsso: %v", err)
		}
	}

	if o.FssoAgentForNtlm != nil {
		v := *o.FssoAgentForNtlm

		if err = d.Set("fsso_agent_for_ntlm", v); err != nil {
			return diag.Errorf("error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if o.FssoGroups != nil {
		if err = d.Set("fsso_groups", flattenFirewallPolicyFssoGroups(o.FssoGroups, sort)); err != nil {
			return diag.Errorf("error reading fsso_groups: %v", err)
		}
	}

	if o.GeoipAnycast != nil {
		v := *o.GeoipAnycast

		if err = d.Set("geoip_anycast", v); err != nil {
			return diag.Errorf("error reading geoip_anycast: %v", err)
		}
	}

	if o.GeoipMatch != nil {
		v := *o.GeoipMatch

		if err = d.Set("geoip_match", v); err != nil {
			return diag.Errorf("error reading geoip_match: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenFirewallPolicyGroups(o.Groups, sort)); err != nil {
			return diag.Errorf("error reading groups: %v", err)
		}
	}

	if o.HttpPolicyRedirect != nil {
		v := *o.HttpPolicyRedirect

		if err = d.Set("http_policy_redirect", v); err != nil {
			return diag.Errorf("error reading http_policy_redirect: %v", err)
		}
	}

	if o.IcapProfile != nil {
		v := *o.IcapProfile

		if err = d.Set("icap_profile", v); err != nil {
			return diag.Errorf("error reading icap_profile: %v", err)
		}
	}

	if o.IdentityBasedRoute != nil {
		v := *o.IdentityBasedRoute

		if err = d.Set("identity_based_route", v); err != nil {
			return diag.Errorf("error reading identity_based_route: %v", err)
		}
	}

	if o.Inbound != nil {
		v := *o.Inbound

		if err = d.Set("inbound", v); err != nil {
			return diag.Errorf("error reading inbound: %v", err)
		}
	}

	if o.InspectionMode != nil {
		v := *o.InspectionMode

		if err = d.Set("inspection_mode", v); err != nil {
			return diag.Errorf("error reading inspection_mode: %v", err)
		}
	}

	if o.InternetService != nil {
		v := *o.InternetService

		if err = d.Set("internet_service", v); err != nil {
			return diag.Errorf("error reading internet_service: %v", err)
		}
	}

	if o.InternetServiceCustom != nil {
		if err = d.Set("internet_service_custom", flattenFirewallPolicyInternetServiceCustom(o.InternetServiceCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceCustomGroup != nil {
		if err = d.Set("internet_service_custom_group", flattenFirewallPolicyInternetServiceCustomGroup(o.InternetServiceCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if o.InternetServiceGroup != nil {
		if err = d.Set("internet_service_group", flattenFirewallPolicyInternetServiceGroup(o.InternetServiceGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		if err = d.Set("internet_service_id", flattenFirewallPolicyInternetServiceId(o.InternetServiceId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
		}
	}

	if o.InternetServiceName != nil {
		if err = d.Set("internet_service_name", flattenFirewallPolicyInternetServiceName(o.InternetServiceName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_name: %v", err)
		}
	}

	if o.InternetServiceNegate != nil {
		v := *o.InternetServiceNegate

		if err = d.Set("internet_service_negate", v); err != nil {
			return diag.Errorf("error reading internet_service_negate: %v", err)
		}
	}

	if o.InternetServiceSrc != nil {
		v := *o.InternetServiceSrc

		if err = d.Set("internet_service_src", v); err != nil {
			return diag.Errorf("error reading internet_service_src: %v", err)
		}
	}

	if o.InternetServiceSrcCustom != nil {
		if err = d.Set("internet_service_src_custom", flattenFirewallPolicyInternetServiceSrcCustom(o.InternetServiceSrcCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom: %v", err)
		}
	}

	if o.InternetServiceSrcCustomGroup != nil {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallPolicyInternetServiceSrcCustomGroup(o.InternetServiceSrcCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom_group: %v", err)
		}
	}

	if o.InternetServiceSrcGroup != nil {
		if err = d.Set("internet_service_src_group", flattenFirewallPolicyInternetServiceSrcGroup(o.InternetServiceSrcGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_group: %v", err)
		}
	}

	if o.InternetServiceSrcId != nil {
		if err = d.Set("internet_service_src_id", flattenFirewallPolicyInternetServiceSrcId(o.InternetServiceSrcId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_id: %v", err)
		}
	}

	if o.InternetServiceSrcName != nil {
		if err = d.Set("internet_service_src_name", flattenFirewallPolicyInternetServiceSrcName(o.InternetServiceSrcName, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_name: %v", err)
		}
	}

	if o.InternetServiceSrcNegate != nil {
		v := *o.InternetServiceSrcNegate

		if err = d.Set("internet_service_src_negate", v); err != nil {
			return diag.Errorf("error reading internet_service_src_negate: %v", err)
		}
	}

	if o.Ippool != nil {
		v := *o.Ippool

		if err = d.Set("ippool", v); err != nil {
			return diag.Errorf("error reading ippool: %v", err)
		}
	}

	if o.IpsSensor != nil {
		v := *o.IpsSensor

		if err = d.Set("ips_sensor", v); err != nil {
			return diag.Errorf("error reading ips_sensor: %v", err)
		}
	}

	if o.Logtraffic != nil {
		v := *o.Logtraffic

		if err = d.Set("logtraffic", v); err != nil {
			return diag.Errorf("error reading logtraffic: %v", err)
		}
	}

	if o.LogtrafficStart != nil {
		v := *o.LogtrafficStart

		if err = d.Set("logtraffic_start", v); err != nil {
			return diag.Errorf("error reading logtraffic_start: %v", err)
		}
	}

	if o.MatchVip != nil {
		v := *o.MatchVip

		if err = d.Set("match_vip", v); err != nil {
			return diag.Errorf("error reading match_vip: %v", err)
		}
	}

	if o.MatchVipOnly != nil {
		v := *o.MatchVipOnly

		if err = d.Set("match_vip_only", v); err != nil {
			return diag.Errorf("error reading match_vip_only: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Nat != nil {
		v := *o.Nat

		if err = d.Set("nat", v); err != nil {
			return diag.Errorf("error reading nat: %v", err)
		}
	}

	if o.Nat46 != nil {
		v := *o.Nat46

		if err = d.Set("nat46", v); err != nil {
			return diag.Errorf("error reading nat46: %v", err)
		}
	}

	if o.Nat64 != nil {
		v := *o.Nat64

		if err = d.Set("nat64", v); err != nil {
			return diag.Errorf("error reading nat64: %v", err)
		}
	}

	if o.Natinbound != nil {
		v := *o.Natinbound

		if err = d.Set("natinbound", v); err != nil {
			return diag.Errorf("error reading natinbound: %v", err)
		}
	}

	if o.Natip != nil {
		v := *o.Natip
		if current, ok := d.GetOk("natip"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("natip", v); err != nil {
			return diag.Errorf("error reading natip: %v", err)
		}
	}

	if o.Natoutbound != nil {
		v := *o.Natoutbound

		if err = d.Set("natoutbound", v); err != nil {
			return diag.Errorf("error reading natoutbound: %v", err)
		}
	}

	if o.NpAcceleration != nil {
		v := *o.NpAcceleration

		if err = d.Set("np_acceleration", v); err != nil {
			return diag.Errorf("error reading np_acceleration: %v", err)
		}
	}

	if o.Ntlm != nil {
		v := *o.Ntlm

		if err = d.Set("ntlm", v); err != nil {
			return diag.Errorf("error reading ntlm: %v", err)
		}
	}

	if o.NtlmEnabledBrowsers != nil {
		if err = d.Set("ntlm_enabled_browsers", flattenFirewallPolicyNtlmEnabledBrowsers(o.NtlmEnabledBrowsers, sort)); err != nil {
			return diag.Errorf("error reading ntlm_enabled_browsers: %v", err)
		}
	}

	if o.NtlmGuest != nil {
		v := *o.NtlmGuest

		if err = d.Set("ntlm_guest", v); err != nil {
			return diag.Errorf("error reading ntlm_guest: %v", err)
		}
	}

	if o.Outbound != nil {
		v := *o.Outbound

		if err = d.Set("outbound", v); err != nil {
			return diag.Errorf("error reading outbound: %v", err)
		}
	}

	if o.PassiveWanHealthMeasurement != nil {
		v := *o.PassiveWanHealthMeasurement

		if err = d.Set("passive_wan_health_measurement", v); err != nil {
			return diag.Errorf("error reading passive_wan_health_measurement: %v", err)
		}
	}

	if o.PerIpShaper != nil {
		v := *o.PerIpShaper

		if err = d.Set("per_ip_shaper", v); err != nil {
			return diag.Errorf("error reading per_ip_shaper: %v", err)
		}
	}

	if o.PermitAnyHost != nil {
		v := *o.PermitAnyHost

		if err = d.Set("permit_any_host", v); err != nil {
			return diag.Errorf("error reading permit_any_host: %v", err)
		}
	}

	if o.PermitStunHost != nil {
		v := *o.PermitStunHost

		if err = d.Set("permit_stun_host", v); err != nil {
			return diag.Errorf("error reading permit_stun_host: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
		}
	}

	if o.Poolname != nil {
		if err = d.Set("poolname", flattenFirewallPolicyPoolname(o.Poolname, sort)); err != nil {
			return diag.Errorf("error reading poolname: %v", err)
		}
	}

	if o.Poolname6 != nil {
		if err = d.Set("poolname6", flattenFirewallPolicyPoolname6(o.Poolname6, sort)); err != nil {
			return diag.Errorf("error reading poolname6: %v", err)
		}
	}

	if o.ProfileGroup != nil {
		v := *o.ProfileGroup

		if err = d.Set("profile_group", v); err != nil {
			return diag.Errorf("error reading profile_group: %v", err)
		}
	}

	if o.ProfileProtocolOptions != nil {
		v := *o.ProfileProtocolOptions

		if err = d.Set("profile_protocol_options", v); err != nil {
			return diag.Errorf("error reading profile_protocol_options: %v", err)
		}
	}

	if o.ProfileType != nil {
		v := *o.ProfileType

		if err = d.Set("profile_type", v); err != nil {
			return diag.Errorf("error reading profile_type: %v", err)
		}
	}

	if o.RadiusMacAuthBypass != nil {
		v := *o.RadiusMacAuthBypass

		if err = d.Set("radius_mac_auth_bypass", v); err != nil {
			return diag.Errorf("error reading radius_mac_auth_bypass: %v", err)
		}
	}

	if o.RedirectUrl != nil {
		v := *o.RedirectUrl

		if err = d.Set("redirect_url", v); err != nil {
			return diag.Errorf("error reading redirect_url: %v", err)
		}
	}

	if o.ReplacemsgOverrideGroup != nil {
		v := *o.ReplacemsgOverrideGroup

		if err = d.Set("replacemsg_override_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_override_group: %v", err)
		}
	}

	if o.ReputationDirection != nil {
		v := *o.ReputationDirection

		if err = d.Set("reputation_direction", v); err != nil {
			return diag.Errorf("error reading reputation_direction: %v", err)
		}
	}

	if o.ReputationMinimum != nil {
		v := *o.ReputationMinimum

		if err = d.Set("reputation_minimum", v); err != nil {
			return diag.Errorf("error reading reputation_minimum: %v", err)
		}
	}

	if o.Rsso != nil {
		v := *o.Rsso

		if err = d.Set("rsso", v); err != nil {
			return diag.Errorf("error reading rsso: %v", err)
		}
	}

	if o.RtpAddr != nil {
		if err = d.Set("rtp_addr", flattenFirewallPolicyRtpAddr(o.RtpAddr, sort)); err != nil {
			return diag.Errorf("error reading rtp_addr: %v", err)
		}
	}

	if o.RtpNat != nil {
		v := *o.RtpNat

		if err = d.Set("rtp_nat", v); err != nil {
			return diag.Errorf("error reading rtp_nat: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.ScheduleTimeout != nil {
		v := *o.ScheduleTimeout

		if err = d.Set("schedule_timeout", v); err != nil {
			return diag.Errorf("error reading schedule_timeout: %v", err)
		}
	}

	if o.SctpFilterProfile != nil {
		v := *o.SctpFilterProfile

		if err = d.Set("sctp_filter_profile", v); err != nil {
			return diag.Errorf("error reading sctp_filter_profile: %v", err)
		}
	}

	if o.SendDenyPacket != nil {
		v := *o.SendDenyPacket

		if err = d.Set("send_deny_packet", v); err != nil {
			return diag.Errorf("error reading send_deny_packet: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallPolicyService(o.Service, sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.ServiceNegate != nil {
		v := *o.ServiceNegate

		if err = d.Set("service_negate", v); err != nil {
			return diag.Errorf("error reading service_negate: %v", err)
		}
	}

	if o.SessionTtl != nil {
		v := *o.SessionTtl

		if err = d.Set("session_ttl", v); err != nil {
			return diag.Errorf("error reading session_ttl: %v", err)
		}
	}

	if o.Sgt != nil {
		if err = d.Set("sgt", flattenFirewallPolicySgt(o.Sgt, sort)); err != nil {
			return diag.Errorf("error reading sgt: %v", err)
		}
	}

	if o.SgtCheck != nil {
		v := *o.SgtCheck

		if err = d.Set("sgt_check", v); err != nil {
			return diag.Errorf("error reading sgt_check: %v", err)
		}
	}

	if o.SrcVendorMac != nil {
		if err = d.Set("src_vendor_mac", flattenFirewallPolicySrcVendorMac(o.SrcVendorMac, sort)); err != nil {
			return diag.Errorf("error reading src_vendor_mac: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenFirewallPolicySrcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.SrcaddrNegate != nil {
		v := *o.SrcaddrNegate

		if err = d.Set("srcaddr_negate", v); err != nil {
			return diag.Errorf("error reading srcaddr_negate: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenFirewallPolicySrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenFirewallPolicySrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.SshFilterProfile != nil {
		v := *o.SshFilterProfile

		if err = d.Set("ssh_filter_profile", v); err != nil {
			return diag.Errorf("error reading ssh_filter_profile: %v", err)
		}
	}

	if o.SshPolicyRedirect != nil {
		v := *o.SshPolicyRedirect

		if err = d.Set("ssh_policy_redirect", v); err != nil {
			return diag.Errorf("error reading ssh_policy_redirect: %v", err)
		}
	}

	if o.SslMirror != nil {
		v := *o.SslMirror

		if err = d.Set("ssl_mirror", v); err != nil {
			return diag.Errorf("error reading ssl_mirror: %v", err)
		}
	}

	if o.SslMirrorIntf != nil {
		if err = d.Set("ssl_mirror_intf", flattenFirewallPolicySslMirrorIntf(o.SslMirrorIntf, sort)); err != nil {
			return diag.Errorf("error reading ssl_mirror_intf: %v", err)
		}
	}

	if o.SslSshProfile != nil {
		v := *o.SslSshProfile

		if err = d.Set("ssl_ssh_profile", v); err != nil {
			return diag.Errorf("error reading ssl_ssh_profile: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TcpMssReceiver != nil {
		v := *o.TcpMssReceiver

		if err = d.Set("tcp_mss_receiver", v); err != nil {
			return diag.Errorf("error reading tcp_mss_receiver: %v", err)
		}
	}

	if o.TcpMssSender != nil {
		v := *o.TcpMssSender

		if err = d.Set("tcp_mss_sender", v); err != nil {
			return diag.Errorf("error reading tcp_mss_sender: %v", err)
		}
	}

	if o.TcpSessionWithoutSyn != nil {
		v := *o.TcpSessionWithoutSyn

		if err = d.Set("tcp_session_without_syn", v); err != nil {
			return diag.Errorf("error reading tcp_session_without_syn: %v", err)
		}
	}

	if o.TimeoutSendRst != nil {
		v := *o.TimeoutSendRst

		if err = d.Set("timeout_send_rst", v); err != nil {
			return diag.Errorf("error reading timeout_send_rst: %v", err)
		}
	}

	if o.Tos != nil {
		v := *o.Tos

		if err = d.Set("tos", v); err != nil {
			return diag.Errorf("error reading tos: %v", err)
		}
	}

	if o.TosMask != nil {
		v := *o.TosMask

		if err = d.Set("tos_mask", v); err != nil {
			return diag.Errorf("error reading tos_mask: %v", err)
		}
	}

	if o.TosNegate != nil {
		v := *o.TosNegate

		if err = d.Set("tos_negate", v); err != nil {
			return diag.Errorf("error reading tos_negate: %v", err)
		}
	}

	if o.TrafficShaper != nil {
		v := *o.TrafficShaper

		if err = d.Set("traffic_shaper", v); err != nil {
			return diag.Errorf("error reading traffic_shaper: %v", err)
		}
	}

	if o.TrafficShaperReverse != nil {
		v := *o.TrafficShaperReverse

		if err = d.Set("traffic_shaper_reverse", v); err != nil {
			return diag.Errorf("error reading traffic_shaper_reverse: %v", err)
		}
	}

	if o.UrlCategory != nil {
		if err = d.Set("url_category", flattenFirewallPolicyUrlCategory(o.UrlCategory, sort)); err != nil {
			return diag.Errorf("error reading url_category: %v", err)
		}
	}

	if o.Users != nil {
		if err = d.Set("users", flattenFirewallPolicyUsers(o.Users, sort)); err != nil {
			return diag.Errorf("error reading users: %v", err)
		}
	}

	if o.UtmStatus != nil {
		v := *o.UtmStatus

		if err = d.Set("utm_status", v); err != nil {
			return diag.Errorf("error reading utm_status: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	if o.VideofilterProfile != nil {
		v := *o.VideofilterProfile

		if err = d.Set("videofilter_profile", v); err != nil {
			return diag.Errorf("error reading videofilter_profile: %v", err)
		}
	}

	if o.VlanCosFwd != nil {
		v := *o.VlanCosFwd

		if err = d.Set("vlan_cos_fwd", v); err != nil {
			return diag.Errorf("error reading vlan_cos_fwd: %v", err)
		}
	}

	if o.VlanCosRev != nil {
		v := *o.VlanCosRev

		if err = d.Set("vlan_cos_rev", v); err != nil {
			return diag.Errorf("error reading vlan_cos_rev: %v", err)
		}
	}

	if o.VlanFilter != nil {
		v := *o.VlanFilter

		if err = d.Set("vlan_filter", v); err != nil {
			return diag.Errorf("error reading vlan_filter: %v", err)
		}
	}

	if o.VoipProfile != nil {
		v := *o.VoipProfile

		if err = d.Set("voip_profile", v); err != nil {
			return diag.Errorf("error reading voip_profile: %v", err)
		}
	}

	if o.Vpntunnel != nil {
		v := *o.Vpntunnel

		if err = d.Set("vpntunnel", v); err != nil {
			return diag.Errorf("error reading vpntunnel: %v", err)
		}
	}

	if o.WafProfile != nil {
		v := *o.WafProfile

		if err = d.Set("waf_profile", v); err != nil {
			return diag.Errorf("error reading waf_profile: %v", err)
		}
	}

	if o.Wanopt != nil {
		v := *o.Wanopt

		if err = d.Set("wanopt", v); err != nil {
			return diag.Errorf("error reading wanopt: %v", err)
		}
	}

	if o.WanoptDetection != nil {
		v := *o.WanoptDetection

		if err = d.Set("wanopt_detection", v); err != nil {
			return diag.Errorf("error reading wanopt_detection: %v", err)
		}
	}

	if o.WanoptPassiveOpt != nil {
		v := *o.WanoptPassiveOpt

		if err = d.Set("wanopt_passive_opt", v); err != nil {
			return diag.Errorf("error reading wanopt_passive_opt: %v", err)
		}
	}

	if o.WanoptPeer != nil {
		v := *o.WanoptPeer

		if err = d.Set("wanopt_peer", v); err != nil {
			return diag.Errorf("error reading wanopt_peer: %v", err)
		}
	}

	if o.WanoptProfile != nil {
		v := *o.WanoptProfile

		if err = d.Set("wanopt_profile", v); err != nil {
			return diag.Errorf("error reading wanopt_profile: %v", err)
		}
	}

	if o.Wccp != nil {
		v := *o.Wccp

		if err = d.Set("wccp", v); err != nil {
			return diag.Errorf("error reading wccp: %v", err)
		}
	}

	if o.Webcache != nil {
		v := *o.Webcache

		if err = d.Set("webcache", v); err != nil {
			return diag.Errorf("error reading webcache: %v", err)
		}
	}

	if o.WebcacheHttps != nil {
		v := *o.WebcacheHttps

		if err = d.Set("webcache_https", v); err != nil {
			return diag.Errorf("error reading webcache_https: %v", err)
		}
	}

	if o.WebfilterProfile != nil {
		v := *o.WebfilterProfile

		if err = d.Set("webfilter_profile", v); err != nil {
			return diag.Errorf("error reading webfilter_profile: %v", err)
		}
	}

	if o.WebproxyForwardServer != nil {
		v := *o.WebproxyForwardServer

		if err = d.Set("webproxy_forward_server", v); err != nil {
			return diag.Errorf("error reading webproxy_forward_server: %v", err)
		}
	}

	if o.WebproxyProfile != nil {
		v := *o.WebproxyProfile

		if err = d.Set("webproxy_profile", v); err != nil {
			return diag.Errorf("error reading webproxy_profile: %v", err)
		}
	}

	if o.Wsso != nil {
		v := *o.Wsso

		if err = d.Set("wsso", v); err != nil {
			return diag.Errorf("error reading wsso: %v", err)
		}
	}

	if o.ZtnaEmsTag != nil {
		if err = d.Set("ztna_ems_tag", flattenFirewallPolicyZtnaEmsTag(o.ZtnaEmsTag, sort)); err != nil {
			return diag.Errorf("error reading ztna_ems_tag: %v", err)
		}
	}

	if o.ZtnaGeoTag != nil {
		if err = d.Set("ztna_geo_tag", flattenFirewallPolicyZtnaGeoTag(o.ZtnaGeoTag, sort)); err != nil {
			return diag.Errorf("error reading ztna_geo_tag: %v", err)
		}
	}

	if o.ZtnaStatus != nil {
		v := *o.ZtnaStatus

		if err = d.Set("ztna_status", v); err != nil {
			return diag.Errorf("error reading ztna_status: %v", err)
		}
	}

	return nil
}

func expandFirewallPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyAppCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyAppCategory

	for i := range l {
		tmp := models.FirewallPolicyAppCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyAppGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyAppGroup

	for i := range l {
		tmp := models.FirewallPolicyAppGroup{}
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

func expandFirewallPolicyApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyApplication, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyApplication

	for i := range l {
		tmp := models.FirewallPolicyApplication{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyCustomLogFields(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyCustomLogFields, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyCustomLogFields

	for i := range l {
		tmp := models.FirewallPolicyCustomLogFields{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.field_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.FieldId = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyDstaddr

	for i := range l {
		tmp := models.FirewallPolicyDstaddr{}
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

func expandFirewallPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyDstaddr6

	for i := range l {
		tmp := models.FirewallPolicyDstaddr6{}
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

func expandFirewallPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyDstintf

	for i := range l {
		tmp := models.FirewallPolicyDstintf{}
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

func expandFirewallPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyFssoGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyFssoGroups

	for i := range l {
		tmp := models.FirewallPolicyFssoGroups{}
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

func expandFirewallPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyGroups

	for i := range l {
		tmp := models.FirewallPolicyGroups{}
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

func expandFirewallPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceCustom

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceCustom{}
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

func expandFirewallPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceCustomGroup

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceCustomGroup{}
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

func expandFirewallPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceGroup

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceGroup{}
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

func expandFirewallPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceId

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceName

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceName{}
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

func expandFirewallPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceSrcCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceSrcCustom

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceSrcCustom{}
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

func expandFirewallPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceSrcCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceSrcCustomGroup

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceSrcCustomGroup{}
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

func expandFirewallPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceSrcGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceSrcGroup

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceSrcGroup{}
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

func expandFirewallPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceSrcId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceSrcId

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceSrcId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyInternetServiceSrcName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyInternetServiceSrcName

	for i := range l {
		tmp := models.FirewallPolicyInternetServiceSrcName{}
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

func expandFirewallPolicyNtlmEnabledBrowsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyNtlmEnabledBrowsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyNtlmEnabledBrowsers

	for i := range l {
		tmp := models.FirewallPolicyNtlmEnabledBrowsers{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.user_agent_string", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.UserAgentString = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyPoolname(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyPoolname, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyPoolname

	for i := range l {
		tmp := models.FirewallPolicyPoolname{}
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

func expandFirewallPolicyPoolname6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyPoolname6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyPoolname6

	for i := range l {
		tmp := models.FirewallPolicyPoolname6{}
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

func expandFirewallPolicyRtpAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyRtpAddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyRtpAddr

	for i := range l {
		tmp := models.FirewallPolicyRtpAddr{}
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

func expandFirewallPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyService

	for i := range l {
		tmp := models.FirewallPolicyService{}
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

func expandFirewallPolicySgt(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicySgt, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicySgt

	for i := range l {
		tmp := models.FirewallPolicySgt{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicySrcVendorMac(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicySrcVendorMac, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicySrcVendorMac

	for i := range l {
		tmp := models.FirewallPolicySrcVendorMac{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicySrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicySrcaddr

	for i := range l {
		tmp := models.FirewallPolicySrcaddr{}
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

func expandFirewallPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicySrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicySrcaddr6

	for i := range l {
		tmp := models.FirewallPolicySrcaddr6{}
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

func expandFirewallPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicySrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicySrcintf

	for i := range l {
		tmp := models.FirewallPolicySrcintf{}
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

func expandFirewallPolicySslMirrorIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicySslMirrorIntf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicySslMirrorIntf

	for i := range l {
		tmp := models.FirewallPolicySslMirrorIntf{}
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

func expandFirewallPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyUrlCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyUrlCategory

	for i := range l {
		tmp := models.FirewallPolicyUrlCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyUsers

	for i := range l {
		tmp := models.FirewallPolicyUsers{}
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

func expandFirewallPolicyZtnaEmsTag(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyZtnaEmsTag, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyZtnaEmsTag

	for i := range l {
		tmp := models.FirewallPolicyZtnaEmsTag{}
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

func expandFirewallPolicyZtnaGeoTag(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicyZtnaGeoTag, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicyZtnaGeoTag

	for i := range l {
		tmp := models.FirewallPolicyZtnaGeoTag{}
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

func getObjectFirewallPolicy(d *schema.ResourceData, sv string) (*models.FirewallPolicy, diag.Diagnostics) {
	obj := models.FirewallPolicy{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("anti_replay"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("anti_replay", sv)
				diags = append(diags, e)
			}
			obj.AntiReplay = &v2
		}
	}
	if v, ok := d.GetOk("app_category"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("app_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppCategory = t
		}
	} else if d.HasChange("app_category") {
		old, new := d.GetChange("app_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppCategory = &[]models.FirewallPolicyAppCategory{}
		}
	}
	if v, ok := d.GetOk("app_group"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("app_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyAppGroup(d, v, "app_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AppGroup = t
		}
	} else if d.HasChange("app_group") {
		old, new := d.GetChange("app_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AppGroup = &[]models.FirewallPolicyAppGroup{}
		}
	}
	if v, ok := d.GetOk("application"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("application", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyApplication(d, v, "application", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Application = t
		}
	} else if d.HasChange("application") {
		old, new := d.GetChange("application")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Application = &[]models.FirewallPolicyApplication{}
		}
	}
	if v1, ok := d.GetOk("application_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application_list", sv)
				diags = append(diags, e)
			}
			obj.ApplicationList = &v2
		}
	}
	if v1, ok := d.GetOk("auth_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_cert", sv)
				diags = append(diags, e)
			}
			obj.AuthCert = &v2
		}
	}
	if v1, ok := d.GetOk("auth_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_path", sv)
				diags = append(diags, e)
			}
			obj.AuthPath = &v2
		}
	}
	if v1, ok := d.GetOk("auth_redirect_addr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_redirect_addr", sv)
				diags = append(diags, e)
			}
			obj.AuthRedirectAddr = &v2
		}
	}
	if v1, ok := d.GetOk("auto_asic_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_asic_offload", sv)
				diags = append(diags, e)
			}
			obj.AutoAsicOffload = &v2
		}
	}
	if v1, ok := d.GetOk("av_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("av_profile", sv)
				diags = append(diags, e)
			}
			obj.AvProfile = &v2
		}
	}
	if v1, ok := d.GetOk("block_notification"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_notification", sv)
				diags = append(diags, e)
			}
			obj.BlockNotification = &v2
		}
	}
	if v1, ok := d.GetOk("captive_portal_exempt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_exempt", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalExempt = &v2
		}
	}
	if v1, ok := d.GetOk("capture_packet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("capture_packet", sv)
				diags = append(diags, e)
			}
			obj.CapturePacket = &v2
		}
	}
	if v1, ok := d.GetOk("cifs_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cifs_profile", sv)
				diags = append(diags, e)
			}
			obj.CifsProfile = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v, ok := d.GetOk("custom_log_fields"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("custom_log_fields", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyCustomLogFields(d, v, "custom_log_fields", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CustomLogFields = t
		}
	} else if d.HasChange("custom_log_fields") {
		old, new := d.GetChange("custom_log_fields")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CustomLogFields = &[]models.FirewallPolicyCustomLogFields{}
		}
	}
	if v1, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("decrypted_traffic_mirror", sv)
				diags = append(diags, e)
			}
			obj.DecryptedTrafficMirror = &v2
		}
	}
	if v1, ok := d.GetOk("delay_tcp_npu_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("delay_tcp_npu_session", sv)
				diags = append(diags, e)
			}
			obj.DelayTcpNpuSession = &v2
		}
	}
	if v1, ok := d.GetOk("diffserv_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv_forward", sv)
				diags = append(diags, e)
			}
			obj.DiffservForward = &v2
		}
	}
	if v1, ok := d.GetOk("diffserv_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffserv_reverse", sv)
				diags = append(diags, e)
			}
			obj.DiffservReverse = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode_forward", sv)
				diags = append(diags, e)
			}
			obj.DiffservcodeForward = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode_rev"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode_rev", sv)
				diags = append(diags, e)
			}
			obj.DiffservcodeRev = &v2
		}
	}
	if v1, ok := d.GetOk("disclaimer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("disclaimer", sv)
				diags = append(diags, e)
			}
			obj.Disclaimer = &v2
		}
	}
	if v1, ok := d.GetOk("dlp_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dlp_sensor", sv)
				diags = append(diags, e)
			}
			obj.DlpSensor = &v2
		}
	}
	if v1, ok := d.GetOk("dnsfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dnsfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.DnsfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("dsri"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dsri", sv)
				diags = append(diags, e)
			}
			obj.Dsri = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.FirewallPolicyDstaddr{}
		}
	}
	if v1, ok := d.GetOk("dstaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.DstaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.FirewallPolicyDstaddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.FirewallPolicyDstintf{}
		}
	}
	if v1, ok := d.GetOk("dynamic_shaping"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.6", "") {
				e := utils.AttributeVersionWarning("dynamic_shaping", sv)
				diags = append(diags, e)
			}
			obj.DynamicShaping = &v2
		}
	}
	if v1, ok := d.GetOk("email_collect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_collect", sv)
				diags = append(diags, e)
			}
			obj.EmailCollect = &v2
		}
	}
	if v1, ok := d.GetOk("emailfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("emailfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.EmailfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("fec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("fec", sv)
				diags = append(diags, e)
			}
			obj.Fec = &v2
		}
	}
	if v1, ok := d.GetOk("file_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("file_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.FileFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("firewall_session_dirty"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("firewall_session_dirty", sv)
				diags = append(diags, e)
			}
			obj.FirewallSessionDirty = &v2
		}
	}
	if v1, ok := d.GetOk("fixedport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fixedport", sv)
				diags = append(diags, e)
			}
			obj.Fixedport = &v2
		}
	}
	if v1, ok := d.GetOk("fsso"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("fsso", sv)
				diags = append(diags, e)
			}
			obj.Fsso = &v2
		}
	}
	if v1, ok := d.GetOk("fsso_agent_for_ntlm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fsso_agent_for_ntlm", sv)
				diags = append(diags, e)
			}
			obj.FssoAgentForNtlm = &v2
		}
	}
	if v, ok := d.GetOk("fsso_groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fsso_groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyFssoGroups(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FssoGroups = t
		}
	} else if d.HasChange("fsso_groups") {
		old, new := d.GetChange("fsso_groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FssoGroups = &[]models.FirewallPolicyFssoGroups{}
		}
	}
	if v1, ok := d.GetOk("geoip_anycast"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("geoip_anycast", sv)
				diags = append(diags, e)
			}
			obj.GeoipAnycast = &v2
		}
	}
	if v1, ok := d.GetOk("geoip_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("geoip_match", sv)
				diags = append(diags, e)
			}
			obj.GeoipMatch = &v2
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.FirewallPolicyGroups{}
		}
	}
	if v1, ok := d.GetOk("http_policy_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_policy_redirect", sv)
				diags = append(diags, e)
			}
			obj.HttpPolicyRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("icap_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icap_profile", sv)
				diags = append(diags, e)
			}
			obj.IcapProfile = &v2
		}
	}
	if v1, ok := d.GetOk("identity_based_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("identity_based_route", sv)
				diags = append(diags, e)
			}
			obj.IdentityBasedRoute = &v2
		}
	}
	if v1, ok := d.GetOk("inbound"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inbound", sv)
				diags = append(diags, e)
			}
			obj.Inbound = &v2
		}
	}
	if v1, ok := d.GetOk("inspection_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("inspection_mode", sv)
				diags = append(diags, e)
			}
			obj.InspectionMode = &v2
		}
	}
	if v1, ok := d.GetOk("internet_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service", sv)
				diags = append(diags, e)
			}
			obj.InternetService = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.FirewallPolicyInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustomGroup = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustomGroup = &[]models.FirewallPolicyInternetServiceCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceGroup = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceGroup = &[]models.FirewallPolicyInternetServiceGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("internet_service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceId = t
		}
	} else if d.HasChange("internet_service_id") {
		old, new := d.GetChange("internet_service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceId = &[]models.FirewallPolicyInternetServiceId{}
		}
	}
	if v, ok := d.GetOk("internet_service_name"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("internet_service_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceName = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceName = &[]models.FirewallPolicyInternetServiceName{}
		}
	}
	if v1, ok := d.GetOk("internet_service_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_negate", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceNegate = &v2
		}
	}
	if v1, ok := d.GetOk("internet_service_src"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_src", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceSrc = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustom = t
		}
	} else if d.HasChange("internet_service_src_custom") {
		old, new := d.GetChange("internet_service_src_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustom = &[]models.FirewallPolicyInternetServiceSrcCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustomGroup = t
		}
	} else if d.HasChange("internet_service_src_custom_group") {
		old, new := d.GetChange("internet_service_src_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustomGroup = &[]models.FirewallPolicyInternetServiceSrcCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcGroup = t
		}
	} else if d.HasChange("internet_service_src_group") {
		old, new := d.GetChange("internet_service_src_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcGroup = &[]models.FirewallPolicyInternetServiceSrcGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("internet_service_src_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceSrcId(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcId = t
		}
	} else if d.HasChange("internet_service_src_id") {
		old, new := d.GetChange("internet_service_src_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcId = &[]models.FirewallPolicyInternetServiceSrcId{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_name"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("internet_service_src_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyInternetServiceSrcName(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcName = t
		}
	} else if d.HasChange("internet_service_src_name") {
		old, new := d.GetChange("internet_service_src_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcName = &[]models.FirewallPolicyInternetServiceSrcName{}
		}
	}
	if v1, ok := d.GetOk("internet_service_src_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("internet_service_src_negate", sv)
				diags = append(diags, e)
			}
			obj.InternetServiceSrcNegate = &v2
		}
	}
	if v1, ok := d.GetOk("ippool"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ippool", sv)
				diags = append(diags, e)
			}
			obj.Ippool = &v2
		}
	}
	if v1, ok := d.GetOk("ips_sensor"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ips_sensor", sv)
				diags = append(diags, e)
			}
			obj.IpsSensor = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic", sv)
				diags = append(diags, e)
			}
			obj.Logtraffic = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic_start"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic_start", sv)
				diags = append(diags, e)
			}
			obj.LogtrafficStart = &v2
		}
	}
	if v1, ok := d.GetOk("match_vip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_vip", sv)
				diags = append(diags, e)
			}
			obj.MatchVip = &v2
		}
	}
	if v1, ok := d.GetOk("match_vip_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("match_vip_only", sv)
				diags = append(diags, e)
			}
			obj.MatchVipOnly = &v2
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
	if v1, ok := d.GetOk("nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nat", sv)
				diags = append(diags, e)
			}
			obj.Nat = &v2
		}
	}
	if v1, ok := d.GetOk("nat46"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nat46", sv)
				diags = append(diags, e)
			}
			obj.Nat46 = &v2
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
	if v1, ok := d.GetOk("natinbound"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("natinbound", sv)
				diags = append(diags, e)
			}
			obj.Natinbound = &v2
		}
	}
	if v1, ok := d.GetOk("natip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("natip", sv)
				diags = append(diags, e)
			}
			obj.Natip = &v2
		}
	}
	if v1, ok := d.GetOk("natoutbound"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("natoutbound", sv)
				diags = append(diags, e)
			}
			obj.Natoutbound = &v2
		}
	}
	if v1, ok := d.GetOk("np_acceleration"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("np_acceleration", sv)
				diags = append(diags, e)
			}
			obj.NpAcceleration = &v2
		}
	}
	if v1, ok := d.GetOk("ntlm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntlm", sv)
				diags = append(diags, e)
			}
			obj.Ntlm = &v2
		}
	}
	if v, ok := d.GetOk("ntlm_enabled_browsers"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ntlm_enabled_browsers", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyNtlmEnabledBrowsers(d, v, "ntlm_enabled_browsers", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NtlmEnabledBrowsers = t
		}
	} else if d.HasChange("ntlm_enabled_browsers") {
		old, new := d.GetChange("ntlm_enabled_browsers")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NtlmEnabledBrowsers = &[]models.FirewallPolicyNtlmEnabledBrowsers{}
		}
	}
	if v1, ok := d.GetOk("ntlm_guest"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ntlm_guest", sv)
				diags = append(diags, e)
			}
			obj.NtlmGuest = &v2
		}
	}
	if v1, ok := d.GetOk("outbound"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbound", sv)
				diags = append(diags, e)
			}
			obj.Outbound = &v2
		}
	}
	if v1, ok := d.GetOk("passive_wan_health_measurement"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("passive_wan_health_measurement", sv)
				diags = append(diags, e)
			}
			obj.PassiveWanHealthMeasurement = &v2
		}
	}
	if v1, ok := d.GetOk("per_ip_shaper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("per_ip_shaper", sv)
				diags = append(diags, e)
			}
			obj.PerIpShaper = &v2
		}
	}
	if v1, ok := d.GetOk("permit_any_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_any_host", sv)
				diags = append(diags, e)
			}
			obj.PermitAnyHost = &v2
		}
	}
	if v1, ok := d.GetOk("permit_stun_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_stun_host", sv)
				diags = append(diags, e)
			}
			obj.PermitStunHost = &v2
		}
	}
	if v1, ok := d.GetOk("policyid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policyid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Policyid = &tmp
		}
	}
	if v, ok := d.GetOk("poolname"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("poolname", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyPoolname(d, v, "poolname", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Poolname = t
		}
	} else if d.HasChange("poolname") {
		old, new := d.GetChange("poolname")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Poolname = &[]models.FirewallPolicyPoolname{}
		}
	}
	if v, ok := d.GetOk("poolname6"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("poolname6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyPoolname6(d, v, "poolname6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Poolname6 = t
		}
	} else if d.HasChange("poolname6") {
		old, new := d.GetChange("poolname6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Poolname6 = &[]models.FirewallPolicyPoolname6{}
		}
	}
	if v1, ok := d.GetOk("profile_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_group", sv)
				diags = append(diags, e)
			}
			obj.ProfileGroup = &v2
		}
	}
	if v1, ok := d.GetOk("profile_protocol_options"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_protocol_options", sv)
				diags = append(diags, e)
			}
			obj.ProfileProtocolOptions = &v2
		}
	}
	if v1, ok := d.GetOk("profile_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("profile_type", sv)
				diags = append(diags, e)
			}
			obj.ProfileType = &v2
		}
	}
	if v1, ok := d.GetOk("radius_mac_auth_bypass"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("radius_mac_auth_bypass", sv)
				diags = append(diags, e)
			}
			obj.RadiusMacAuthBypass = &v2
		}
	}
	if v1, ok := d.GetOk("redirect_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redirect_url", sv)
				diags = append(diags, e)
			}
			obj.RedirectUrl = &v2
		}
	}
	if v1, ok := d.GetOk("replacemsg_override_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("replacemsg_override_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgOverrideGroup = &v2
		}
	}
	if v1, ok := d.GetOk("reputation_direction"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reputation_direction", sv)
				diags = append(diags, e)
			}
			obj.ReputationDirection = &v2
		}
	}
	if v1, ok := d.GetOk("reputation_minimum"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reputation_minimum", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ReputationMinimum = &tmp
		}
	}
	if v1, ok := d.GetOk("rsso"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("rsso", sv)
				diags = append(diags, e)
			}
			obj.Rsso = &v2
		}
	}
	if v, ok := d.GetOk("rtp_addr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("rtp_addr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyRtpAddr(d, v, "rtp_addr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.RtpAddr = t
		}
	} else if d.HasChange("rtp_addr") {
		old, new := d.GetChange("rtp_addr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.RtpAddr = &[]models.FirewallPolicyRtpAddr{}
		}
	}
	if v1, ok := d.GetOk("rtp_nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rtp_nat", sv)
				diags = append(diags, e)
			}
			obj.RtpNat = &v2
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v1, ok := d.GetOk("schedule_timeout"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule_timeout", sv)
				diags = append(diags, e)
			}
			obj.ScheduleTimeout = &v2
		}
	}
	if v1, ok := d.GetOk("sctp_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("sctp_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.SctpFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("send_deny_packet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_deny_packet", sv)
				diags = append(diags, e)
			}
			obj.SendDenyPacket = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallPolicyService{}
		}
	}
	if v1, ok := d.GetOk("service_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("service_negate", sv)
				diags = append(diags, e)
			}
			obj.ServiceNegate = &v2
		}
	}
	if v1, ok := d.GetOk("session_ttl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_ttl", sv)
				diags = append(diags, e)
			}
			obj.SessionTtl = &v2
		}
	}
	if v, ok := d.GetOk("sgt"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("sgt", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicySgt(d, v, "sgt", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Sgt = t
		}
	} else if d.HasChange("sgt") {
		old, new := d.GetChange("sgt")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Sgt = &[]models.FirewallPolicySgt{}
		}
	}
	if v1, ok := d.GetOk("sgt_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("sgt_check", sv)
				diags = append(diags, e)
			}
			obj.SgtCheck = &v2
		}
	}
	if v, ok := d.GetOk("src_vendor_mac"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("src_vendor_mac", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicySrcVendorMac(d, v, "src_vendor_mac", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SrcVendorMac = t
		}
	} else if d.HasChange("src_vendor_mac") {
		old, new := d.GetChange("src_vendor_mac")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SrcVendorMac = &[]models.FirewallPolicySrcVendorMac{}
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.FirewallPolicySrcaddr{}
		}
	}
	if v1, ok := d.GetOk("srcaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("srcaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.SrcaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.FirewallPolicySrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.FirewallPolicySrcintf{}
		}
	}
	if v1, ok := d.GetOk("ssh_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.SshFilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_policy_redirect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_policy_redirect", sv)
				diags = append(diags, e)
			}
			obj.SshPolicyRedirect = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_mirror"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("ssl_mirror", sv)
				diags = append(diags, e)
			}
			obj.SslMirror = &v2
		}
	}
	if v, ok := d.GetOk("ssl_mirror_intf"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("ssl_mirror_intf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicySslMirrorIntf(d, v, "ssl_mirror_intf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SslMirrorIntf = t
		}
	} else if d.HasChange("ssl_mirror_intf") {
		old, new := d.GetChange("ssl_mirror_intf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SslMirrorIntf = &[]models.FirewallPolicySslMirrorIntf{}
		}
	}
	if v1, ok := d.GetOk("ssl_ssh_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_ssh_profile", sv)
				diags = append(diags, e)
			}
			obj.SslSshProfile = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("tcp_mss_receiver"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_mss_receiver", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpMssReceiver = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_mss_sender"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_mss_sender", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpMssSender = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_session_without_syn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_session_without_syn", sv)
				diags = append(diags, e)
			}
			obj.TcpSessionWithoutSyn = &v2
		}
	}
	if v1, ok := d.GetOk("timeout_send_rst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("timeout_send_rst", sv)
				diags = append(diags, e)
			}
			obj.TimeoutSendRst = &v2
		}
	}
	if v1, ok := d.GetOk("tos"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos", sv)
				diags = append(diags, e)
			}
			obj.Tos = &v2
		}
	}
	if v1, ok := d.GetOk("tos_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos_mask", sv)
				diags = append(diags, e)
			}
			obj.TosMask = &v2
		}
	}
	if v1, ok := d.GetOk("tos_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos_negate", sv)
				diags = append(diags, e)
			}
			obj.TosNegate = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_shaper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_shaper", sv)
				diags = append(diags, e)
			}
			obj.TrafficShaper = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_shaper_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_shaper_reverse", sv)
				diags = append(diags, e)
			}
			obj.TrafficShaperReverse = &v2
		}
	}
	if v, ok := d.GetOk("url_category"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("url_category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyUrlCategory(d, v, "url_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.UrlCategory = t
		}
	} else if d.HasChange("url_category") {
		old, new := d.GetChange("url_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.UrlCategory = &[]models.FirewallPolicyUrlCategory{}
		}
	}
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.FirewallPolicyUsers{}
		}
	}
	if v1, ok := d.GetOk("utm_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("utm_status", sv)
				diags = append(diags, e)
			}
			obj.UtmStatus = &v2
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
	if v1, ok := d.GetOk("videofilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("videofilter_profile", sv)
				diags = append(diags, e)
			}
			obj.VideofilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("vlan_cos_fwd"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_cos_fwd", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VlanCosFwd = &tmp
		}
	}
	if v1, ok := d.GetOk("vlan_cos_rev"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_cos_rev", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VlanCosRev = &tmp
		}
	}
	if v1, ok := d.GetOk("vlan_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vlan_filter", sv)
				diags = append(diags, e)
			}
			obj.VlanFilter = &v2
		}
	}
	if v1, ok := d.GetOk("voip_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip_profile", sv)
				diags = append(diags, e)
			}
			obj.VoipProfile = &v2
		}
	}
	if v1, ok := d.GetOk("vpntunnel"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vpntunnel", sv)
				diags = append(diags, e)
			}
			obj.Vpntunnel = &v2
		}
	}
	if v1, ok := d.GetOk("waf_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("waf_profile", sv)
				diags = append(diags, e)
			}
			obj.WafProfile = &v2
		}
	}
	if v1, ok := d.GetOk("wanopt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanopt", sv)
				diags = append(diags, e)
			}
			obj.Wanopt = &v2
		}
	}
	if v1, ok := d.GetOk("wanopt_detection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanopt_detection", sv)
				diags = append(diags, e)
			}
			obj.WanoptDetection = &v2
		}
	}
	if v1, ok := d.GetOk("wanopt_passive_opt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanopt_passive_opt", sv)
				diags = append(diags, e)
			}
			obj.WanoptPassiveOpt = &v2
		}
	}
	if v1, ok := d.GetOk("wanopt_peer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanopt_peer", sv)
				diags = append(diags, e)
			}
			obj.WanoptPeer = &v2
		}
	}
	if v1, ok := d.GetOk("wanopt_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wanopt_profile", sv)
				diags = append(diags, e)
			}
			obj.WanoptProfile = &v2
		}
	}
	if v1, ok := d.GetOk("wccp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wccp", sv)
				diags = append(diags, e)
			}
			obj.Wccp = &v2
		}
	}
	if v1, ok := d.GetOk("webcache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webcache", sv)
				diags = append(diags, e)
			}
			obj.Webcache = &v2
		}
	}
	if v1, ok := d.GetOk("webcache_https"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webcache_https", sv)
				diags = append(diags, e)
			}
			obj.WebcacheHttps = &v2
		}
	}
	if v1, ok := d.GetOk("webfilter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webfilter_profile", sv)
				diags = append(diags, e)
			}
			obj.WebfilterProfile = &v2
		}
	}
	if v1, ok := d.GetOk("webproxy_forward_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webproxy_forward_server", sv)
				diags = append(diags, e)
			}
			obj.WebproxyForwardServer = &v2
		}
	}
	if v1, ok := d.GetOk("webproxy_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("webproxy_profile", sv)
				diags = append(diags, e)
			}
			obj.WebproxyProfile = &v2
		}
	}
	if v1, ok := d.GetOk("wsso"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("wsso", sv)
				diags = append(diags, e)
			}
			obj.Wsso = &v2
		}
	}
	if v, ok := d.GetOk("ztna_ems_tag"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("ztna_ems_tag", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyZtnaEmsTag(d, v, "ztna_ems_tag", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ZtnaEmsTag = t
		}
	} else if d.HasChange("ztna_ems_tag") {
		old, new := d.GetChange("ztna_ems_tag")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ZtnaEmsTag = &[]models.FirewallPolicyZtnaEmsTag{}
		}
	}
	if v, ok := d.GetOk("ztna_geo_tag"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("ztna_geo_tag", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicyZtnaGeoTag(d, v, "ztna_geo_tag", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ZtnaGeoTag = t
		}
	} else if d.HasChange("ztna_geo_tag") {
		old, new := d.GetChange("ztna_geo_tag")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ZtnaGeoTag = &[]models.FirewallPolicyZtnaGeoTag{}
		}
	}
	if v1, ok := d.GetOk("ztna_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ztna_status", sv)
				diags = append(diags, e)
			}
			obj.ZtnaStatus = &v2
		}
	}
	return &obj, diags
}
