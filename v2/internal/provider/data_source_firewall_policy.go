// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4/IPv6 policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4/IPv6 policies.",

		ReadContext: dataSourceFirewallPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Policy action (accept/deny/ipsec).",
				Computed:    true,
			},
			"anti_replay": {
				Type:        schema.TypeString,
				Description: "Enable/disable anti-replay check.",
				Computed:    true,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "Application category ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Category IDs.",
							Computed:    true,
						},
					},
				},
			},
			"app_group": {
				Type:        schema.TypeList,
				Description: "Application group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Application group names.",
							Computed:    true,
						},
					},
				},
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application IDs.",
							Computed:    true,
						},
					},
				},
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Name of an existing Application list.",
				Computed:    true,
			},
			"auth_cert": {
				Type:        schema.TypeString,
				Description: "HTTPS server certificate for policy authentication.",
				Computed:    true,
			},
			"auth_path": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication-based routing.",
				Computed:    true,
			},
			"auth_redirect_addr": {
				Type:        schema.TypeString,
				Description: "HTTP-to-HTTPS redirect address for firewall authentication.",
				Computed:    true,
			},
			"auto_asic_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy traffic ASIC offloading.",
				Computed:    true,
			},
			"av_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Antivirus profile.",
				Computed:    true,
			},
			"block_notification": {
				Type:        schema.TypeString,
				Description: "Enable/disable block notification.",
				Computed:    true,
			},
			"captive_portal_exempt": {
				Type:        schema.TypeString,
				Description: "Enable to exempt some users from the captive portal.",
				Computed:    true,
			},
			"capture_packet": {
				Type:        schema.TypeString,
				Description: "Enable/disable capture packets.",
				Computed:    true,
			},
			"cifs_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing CIFS profile.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"custom_log_fields": {
				Type:        schema.TypeList,
				Description: "Custom fields to append to log messages for this policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": {
							Type:        schema.TypeString,
							Description: "Custom log field.",
							Computed:    true,
						},
					},
				},
			},
			"decrypted_traffic_mirror": {
				Type:        schema.TypeString,
				Description: "Decrypted traffic mirror.",
				Computed:    true,
			},
			"delay_tcp_npu_session": {
				Type:        schema.TypeString,
				Description: "Enable TCP NPU session delay to guarantee packet order of 3-way handshake.",
				Computed:    true,
			},
			"diffserv_forward": {
				Type:        schema.TypeString,
				Description: "Enable to change packet's DiffServ values to the specified diffservcode-forward value.",
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:        schema.TypeString,
				Description: "Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.",
				Computed:    true,
			},
			"diffservcode_forward": {
				Type:        schema.TypeString,
				Description: "Change packet's DiffServ to this value.",
				Computed:    true,
			},
			"diffservcode_rev": {
				Type:        schema.TypeString,
				Description: "Change packet's reverse (reply) DiffServ to this value.",
				Computed:    true,
			},
			"disclaimer": {
				Type:        schema.TypeString,
				Description: "Enable/disable user authentication disclaimer.",
				Computed:    true,
			},
			"dlp_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing DLP sensor.",
				Computed:    true,
			},
			"dnsfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing DNS filter profile.",
				Computed:    true,
			},
			"dsri": {
				Type:        schema.TypeString,
				Description: "Enable DSRI to ignore HTTP server responses.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address and address group names.",
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
			"dstaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be.",
				Computed:    true,
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "Destination IPv6 address name and address group names.",
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
			"dstintf": {
				Type:        schema.TypeList,
				Description: "Outgoing (egress) interface.",
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
			"dynamic_shaping": {
				Type:        schema.TypeString,
				Description: "Enable/disable dynamic RADIUS defined traffic shaping.",
				Computed:    true,
			},
			"email_collect": {
				Type:        schema.TypeString,
				Description: "Enable/disable email collection.",
				Computed:    true,
			},
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing email filter profile.",
				Computed:    true,
			},
			"fec": {
				Type:        schema.TypeString,
				Description: "Enable/disable Forward Error Correction on traffic matching this policy on a FEC device.",
				Computed:    true,
			},
			"file_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing file-filter profile.",
				Computed:    true,
			},
			"firewall_session_dirty": {
				Type:        schema.TypeString,
				Description: "How to handle sessions if the configuration of this firewall policy changes.",
				Computed:    true,
			},
			"fixedport": {
				Type:        schema.TypeString,
				Description: "Enable to prevent source NAT from changing a session's source port.",
				Computed:    true,
			},
			"fsso": {
				Type:        schema.TypeString,
				Description: "Enable/disable Fortinet Single Sign-On.",
				Computed:    true,
			},
			"fsso_agent_for_ntlm": {
				Type:        schema.TypeString,
				Description: "FSSO agent to use for NTLM authentication.",
				Computed:    true,
			},
			"fsso_groups": {
				Type:        schema.TypeList,
				Description: "Names of FSSO groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Names of FSSO groups.",
							Computed:    true,
						},
					},
				},
			},
			"geoip_anycast": {
				Type:        schema.TypeString,
				Description: "Enable/disable recognition of anycast IP addresses using the geography IP database.",
				Computed:    true,
			},
			"geoip_match": {
				Type:        schema.TypeString,
				Description: "Match geography address based either on its physical location or registered location.",
				Computed:    true,
			},
			"groups": {
				Type:        schema.TypeList,
				Description: "Names of user groups that can authenticate with this policy.",
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
			"http_policy_redirect": {
				Type:        schema.TypeString,
				Description: "Redirect HTTP(S) traffic to matching transparent web proxy policy.",
				Computed:    true,
			},
			"icap_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing ICAP profile.",
				Computed:    true,
			},
			"identity_based_route": {
				Type:        schema.TypeString,
				Description: "Name of identity-based routing rule.",
				Computed:    true,
			},
			"inbound": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.",
				Computed:    true,
			},
			"inspection_mode": {
				Type:        schema.TypeString,
				Description: "Policy inspection mode (Flow/proxy). Default is Flow mode.",
				Computed:    true,
			},
			"internet_service": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. ",
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_group": {
				Type:        schema.TypeList,
				Description: "Internet Service group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Internet Service ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_name": {
				Type:        schema.TypeList,
				Description: "Internet Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_negate": {
				Type:        schema.TypeString,
				Description: "When enabled internet-service specifies what the service must NOT be.",
				Computed:    true,
			},
			"internet_service_src": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. ",
				Computed:    true,
			},
			"internet_service_src_custom": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_custom_group": {
				Type:        schema.TypeList,
				Description: "Custom Internet Service source group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Custom Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_group": {
				Type:        schema.TypeList,
				Description: "Internet Service source group name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service group name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_id": {
				Type:        schema.TypeList,
				Description: "Internet Service source ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Internet Service ID.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_name": {
				Type:        schema.TypeList,
				Description: "Internet Service source name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Internet Service name.",
							Computed:    true,
						},
					},
				},
			},
			"internet_service_src_negate": {
				Type:        schema.TypeString,
				Description: "When enabled internet-service-src specifies what the service must NOT be.",
				Computed:    true,
			},
			"ippool": {
				Type:        schema.TypeString,
				Description: "Enable to use IP Pools for source NAT.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "Name of an existing IPS sensor.",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Enable or disable logging. Log all sessions or security profile sessions.",
				Computed:    true,
			},
			"logtraffic_start": {
				Type:        schema.TypeString,
				Description: "Record logs when a session starts.",
				Computed:    true,
			},
			"match_vip": {
				Type:        schema.TypeString,
				Description: "Enable to match packets that have had their destination addresses changed by a VIP.",
				Computed:    true,
			},
			"match_vip_only": {
				Type:        schema.TypeString,
				Description: "Enable/disable matching of only those packets that have had their destination addresses changed by a VIP.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Policy name.",
				Computed:    true,
			},
			"nat": {
				Type:        schema.TypeString,
				Description: "Enable/disable source NAT.",
				Computed:    true,
			},
			"nat46": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT46.",
				Computed:    true,
			},
			"nat64": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT64.",
				Computed:    true,
			},
			"natinbound": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: apply destination NAT to inbound traffic.",
				Computed:    true,
			},
			"natip": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: source NAT IP address for outgoing traffic.",
				Computed:    true,
			},
			"natoutbound": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: apply source NAT to outbound traffic.",
				Computed:    true,
			},
			"np_acceleration": {
				Type:        schema.TypeString,
				Description: "Enable/disable UTM Network Processor acceleration.",
				Computed:    true,
			},
			"ntlm": {
				Type:        schema.TypeString,
				Description: "Enable/disable NTLM authentication.",
				Computed:    true,
			},
			"ntlm_enabled_browsers": {
				Type:        schema.TypeList,
				Description: "HTTP-User-Agent value of supported browsers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_agent_string": {
							Type:        schema.TypeString,
							Description: "User agent string.",
							Computed:    true,
						},
					},
				},
			},
			"ntlm_guest": {
				Type:        schema.TypeString,
				Description: "Enable/disable NTLM guest user access.",
				Computed:    true,
			},
			"outbound": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.",
				Computed:    true,
			},
			"passive_wan_health_measurement": {
				Type:        schema.TypeString,
				Description: "Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled.",
				Computed:    true,
			},
			"per_ip_shaper": {
				Type:        schema.TypeString,
				Description: "Per-IP traffic shaper.",
				Computed:    true,
			},
			"permit_any_host": {
				Type:        schema.TypeString,
				Description: "Accept UDP packets from any host.",
				Computed:    true,
			},
			"permit_stun_host": {
				Type:        schema.TypeString,
				Description: "Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host.",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID (0 - 4294967294).",
				Required:    true,
			},
			"poolname": {
				Type:        schema.TypeList,
				Description: "IP Pool names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IP pool name.",
							Computed:    true,
						},
					},
				},
			},
			"poolname6": {
				Type:        schema.TypeList,
				Description: "IPv6 pool names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IPv6 pool name.",
							Computed:    true,
						},
					},
				},
			},
			"profile_group": {
				Type:        schema.TypeString,
				Description: "Name of profile group.",
				Computed:    true,
			},
			"profile_protocol_options": {
				Type:        schema.TypeString,
				Description: "Name of an existing Protocol options profile.",
				Computed:    true,
			},
			"profile_type": {
				Type:        schema.TypeString,
				Description: "Determine whether the firewall policy allows security profile groups or single profiles only.",
				Computed:    true,
			},
			"radius_mac_auth_bypass": {
				Type:        schema.TypeString,
				Description: "Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server.",
				Computed:    true,
			},
			"redirect_url": {
				Type:        schema.TypeString,
				Description: "URL users are directed to after seeing and accepting the disclaimer or authenticating.",
				Computed:    true,
			},
			"replacemsg_override_group": {
				Type:        schema.TypeString,
				Description: "Override the default replacement message group for this policy.",
				Computed:    true,
			},
			"reputation_direction": {
				Type:        schema.TypeString,
				Description: "Direction of the initial traffic for reputation to take effect.",
				Computed:    true,
			},
			"reputation_minimum": {
				Type:        schema.TypeInt,
				Description: "Minimum Reputation to take action.",
				Computed:    true,
			},
			"rsso": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS single sign-on (RSSO).",
				Computed:    true,
			},
			"rtp_addr": {
				Type:        schema.TypeList,
				Description: "Address names if this is an RTP NAT policy.",
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
			"rtp_nat": {
				Type:        schema.TypeString,
				Description: "Enable Real Time Protocol (RTP) NAT.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
				Computed:    true,
			},
			"schedule_timeout": {
				Type:        schema.TypeString,
				Description: "Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity.",
				Computed:    true,
			},
			"sctp_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SCTP filter profile.",
				Computed:    true,
			},
			"send_deny_packet": {
				Type:        schema.TypeString,
				Description: "Enable to send a reply when a session is denied or blocked by a firewall policy.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service and service group names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Service and service group names.",
							Computed:    true,
						},
					},
				},
			},
			"service_negate": {
				Type:        schema.TypeString,
				Description: "When enabled service specifies what the service must NOT be.",
				Computed:    true,
			},
			"session_ttl": {
				Type:        schema.TypeString,
				Description: "TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).",
				Computed:    true,
			},
			"sgt": {
				Type:        schema.TypeList,
				Description: "Security group tags.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Security group tag (1 - 65535).",
							Computed:    true,
						},
					},
				},
			},
			"sgt_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable security group tags (SGT) check.",
				Computed:    true,
			},
			"src_vendor_mac": {
				Type:        schema.TypeList,
				Description: "Vendor MAC source ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Vendor MAC ID.",
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address and address group names.",
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
			"srcaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled srcaddr/srcaddr6 specifies what the source address must NOT be.",
				Computed:    true,
			},
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "Source IPv6 address name and address group names.",
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
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Incoming (ingress) interface.",
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
			"ssh_filter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSH filter profile.",
				Computed:    true,
			},
			"ssh_policy_redirect": {
				Type:        schema.TypeString,
				Description: "Redirect SSH traffic to matching transparent proxy policy.",
				Computed:    true,
			},
			"ssl_mirror": {
				Type:        schema.TypeString,
				Description: "Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).",
				Computed:    true,
			},
			"ssl_mirror_intf": {
				Type:        schema.TypeList,
				Description: "SSL mirror interface name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Mirror Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"ssl_ssh_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing SSL SSH profile.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable or disable this policy.",
				Computed:    true,
			},
			"tcp_mss_receiver": {
				Type:        schema.TypeInt,
				Description: "Receiver TCP maximum segment size (MSS).",
				Computed:    true,
			},
			"tcp_mss_sender": {
				Type:        schema.TypeInt,
				Description: "Sender TCP maximum segment size (MSS).",
				Computed:    true,
			},
			"tcp_session_without_syn": {
				Type:        schema.TypeString,
				Description: "Enable/disable creation of TCP session without SYN flag.",
				Computed:    true,
			},
			"timeout_send_rst": {
				Type:        schema.TypeString,
				Description: "Enable/disable sending RST packets when TCP sessions expire.",
				Computed:    true,
			},
			"tos": {
				Type:        schema.TypeString,
				Description: "ToS (Type of Service) value used for comparison.",
				Computed:    true,
			},
			"tos_mask": {
				Type:        schema.TypeString,
				Description: "Non-zero bit positions are used for comparison while zero bit positions are ignored.",
				Computed:    true,
			},
			"tos_negate": {
				Type:        schema.TypeString,
				Description: "Enable negated TOS match.",
				Computed:    true,
			},
			"traffic_shaper": {
				Type:        schema.TypeString,
				Description: "Traffic shaper.",
				Computed:    true,
			},
			"traffic_shaper_reverse": {
				Type:        schema.TypeString,
				Description: "Reverse traffic shaper.",
				Computed:    true,
			},
			"url_category": {
				Type:        schema.TypeList,
				Description: "URL category ID list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "URL category ID.",
							Computed:    true,
						},
					},
				},
			},
			"users": {
				Type:        schema.TypeList,
				Description: "Names of individual users that can authenticate with this policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Names of individual users that can authenticate with this policy.",
							Computed:    true,
						},
					},
				},
			},
			"utm_status": {
				Type:        schema.TypeString,
				Description: "Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"videofilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VideoFilter profile.",
				Computed:    true,
			},
			"vlan_cos_fwd": {
				Type:        schema.TypeInt,
				Description: "VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.",
				Computed:    true,
			},
			"vlan_cos_rev": {
				Type:        schema.TypeInt,
				Description: "VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.",
				Computed:    true,
			},
			"vlan_filter": {
				Type:        schema.TypeString,
				Description: "Set VLAN filters.",
				Computed:    true,
			},
			"voip_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing VoIP profile.",
				Computed:    true,
			},
			"vpntunnel": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: name of the IPsec VPN Phase 1.",
				Computed:    true,
			},
			"waf_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web application firewall profile.",
				Computed:    true,
			},
			"wanopt": {
				Type:        schema.TypeString,
				Description: "Enable/disable WAN optimization.",
				Computed:    true,
			},
			"wanopt_detection": {
				Type:        schema.TypeString,
				Description: "WAN optimization auto-detection mode.",
				Computed:    true,
			},
			"wanopt_passive_opt": {
				Type:        schema.TypeString,
				Description: "WAN optimization passive mode options. This option decides what IP address will be used to connect server.",
				Computed:    true,
			},
			"wanopt_peer": {
				Type:        schema.TypeString,
				Description: "WAN optimization peer.",
				Computed:    true,
			},
			"wanopt_profile": {
				Type:        schema.TypeString,
				Description: "WAN optimization profile.",
				Computed:    true,
			},
			"wccp": {
				Type:        schema.TypeString,
				Description: "Enable/disable forwarding traffic matching this policy to a configured WCCP server.",
				Computed:    true,
			},
			"webcache": {
				Type:        schema.TypeString,
				Description: "Enable/disable web cache.",
				Computed:    true,
			},
			"webcache_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable web cache for HTTPS.",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing Web filter profile.",
				Computed:    true,
			},
			"webproxy_forward_server": {
				Type:        schema.TypeString,
				Description: "Webproxy forward server name.",
				Computed:    true,
			},
			"webproxy_profile": {
				Type:        schema.TypeString,
				Description: "Webproxy profile name.",
				Computed:    true,
			},
			"wsso": {
				Type:        schema.TypeString,
				Description: "Enable/disable WiFi Single Sign On (WSSO).",
				Computed:    true,
			},
			"ztna_ems_tag": {
				Type:        schema.TypeList,
				Description: "Source ztna-ems-tag names.",
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
			"ztna_geo_tag": {
				Type:        schema.TypeList,
				Description: "Source ztna-geo-tag names.",
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
			"ztna_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable zero trust access.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("policyid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadFirewallPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
