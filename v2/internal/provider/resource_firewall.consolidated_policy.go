// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure consolidated IPv4/IPv6 policies.

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

func resourceFirewallconsolidatedPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure consolidated IPv4/IPv6 policies.",

		CreateContext: resourceFirewallconsolidatedPolicyCreate,
		ReadContext:   resourceFirewallconsolidatedPolicyRead,
		UpdateContext: resourceFirewallconsolidatedPolicyUpdate,
		DeleteContext: resourceFirewallconsolidatedPolicyDelete,

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
			"application_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing Application list.",
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
			"captive_portal_exempt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable exemption of some users from the captive portal.",
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

				Description: "Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. ",
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
			"dstaddr_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled dstaddr specifies what the destination address must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr4": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address name and address group names.",
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
			"emailfilter_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of an existing email filter profile.",
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
			"fixedport": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to prevent source NAT from changing a session's source port.",
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
			"outbound": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.",
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
			"policyid": {
				Type: schema.TypeInt,

				Description: "Policy ID (0 - 4294967294).",
				ForceNew:    true,
				Required:    true,
			},
			"poolname4": {
				Type:        schema.TypeList,
				Description: "IPv4 pool names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPv4 pool name.",
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
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Schedule name.",
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

							Description: "Service name.",
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
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 2764800),

				Description: "TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "When enabled srcaddr specifies what the source address must NOT be.",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr4": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address name and address group names.",
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
			"users": {
				Type:        schema.TypeList,
				Description: "Names of individual users that can authenticate with this policy.",
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

				Description: "WAN optimization passive mode options. This option decides what IP address will be used to connect to server.",
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
		},
	}
}

func resourceFirewallconsolidatedPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallconsolidatedPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallconsolidatedPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallconsolidatedPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallconsolidatedPolicy")
	}

	return resourceFirewallconsolidatedPolicyRead(ctx, d, meta)
}

func resourceFirewallconsolidatedPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallconsolidatedPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallconsolidatedPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallconsolidatedPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallconsolidatedPolicy")
	}

	return resourceFirewallconsolidatedPolicyRead(ctx, d, meta)
}

func resourceFirewallconsolidatedPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallconsolidatedPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallconsolidatedPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallconsolidatedPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallconsolidatedPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallconsolidatedPolicy resource: %v", err)
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

	diags := refreshObjectFirewallconsolidatedPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallconsolidatedPolicyDstaddr4(v *[]models.FirewallconsolidatedPolicyDstaddr4, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyDstaddr6(v *[]models.FirewallconsolidatedPolicyDstaddr6, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyDstintf(v *[]models.FirewallconsolidatedPolicyDstintf, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyFssoGroups(v *[]models.FirewallconsolidatedPolicyFssoGroups, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyGroups(v *[]models.FirewallconsolidatedPolicyGroups, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceCustom(v *[]models.FirewallconsolidatedPolicyInternetServiceCustom, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceCustomGroup(v *[]models.FirewallconsolidatedPolicyInternetServiceCustomGroup, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceGroup(v *[]models.FirewallconsolidatedPolicyInternetServiceGroup, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceId(v *[]models.FirewallconsolidatedPolicyInternetServiceId, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceName(v *[]models.FirewallconsolidatedPolicyInternetServiceName, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceSrcCustom(v *[]models.FirewallconsolidatedPolicyInternetServiceSrcCustom, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceSrcCustomGroup(v *[]models.FirewallconsolidatedPolicyInternetServiceSrcCustomGroup, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceSrcGroup(v *[]models.FirewallconsolidatedPolicyInternetServiceSrcGroup, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceSrcId(v *[]models.FirewallconsolidatedPolicyInternetServiceSrcId, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyInternetServiceSrcName(v *[]models.FirewallconsolidatedPolicyInternetServiceSrcName, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyPoolname4(v *[]models.FirewallconsolidatedPolicyPoolname4, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyPoolname6(v *[]models.FirewallconsolidatedPolicyPoolname6, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyService(v *[]models.FirewallconsolidatedPolicyService, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicySrcaddr4(v *[]models.FirewallconsolidatedPolicySrcaddr4, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicySrcaddr6(v *[]models.FirewallconsolidatedPolicySrcaddr6, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicySrcintf(v *[]models.FirewallconsolidatedPolicySrcintf, sort bool) interface{} {
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

func flattenFirewallconsolidatedPolicyUsers(v *[]models.FirewallconsolidatedPolicyUsers, sort bool) interface{} {
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

func refreshObjectFirewallconsolidatedPolicy(d *schema.ResourceData, o *models.FirewallconsolidatedPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.ApplicationList != nil {
		v := *o.ApplicationList

		if err = d.Set("application_list", v); err != nil {
			return diag.Errorf("error reading application_list: %v", err)
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

	if o.CaptivePortalExempt != nil {
		v := *o.CaptivePortalExempt

		if err = d.Set("captive_portal_exempt", v); err != nil {
			return diag.Errorf("error reading captive_portal_exempt: %v", err)
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

	if o.DstaddrNegate != nil {
		v := *o.DstaddrNegate

		if err = d.Set("dstaddr_negate", v); err != nil {
			return diag.Errorf("error reading dstaddr_negate: %v", err)
		}
	}

	if o.Dstaddr4 != nil {
		if err = d.Set("dstaddr4", flattenFirewallconsolidatedPolicyDstaddr4(o.Dstaddr4, sort)); err != nil {
			return diag.Errorf("error reading dstaddr4: %v", err)
		}
	}

	if o.Dstaddr6 != nil {
		if err = d.Set("dstaddr6", flattenFirewallconsolidatedPolicyDstaddr6(o.Dstaddr6, sort)); err != nil {
			return diag.Errorf("error reading dstaddr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenFirewallconsolidatedPolicyDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.EmailfilterProfile != nil {
		v := *o.EmailfilterProfile

		if err = d.Set("emailfilter_profile", v); err != nil {
			return diag.Errorf("error reading emailfilter_profile: %v", err)
		}
	}

	if o.FileFilterProfile != nil {
		v := *o.FileFilterProfile

		if err = d.Set("file_filter_profile", v); err != nil {
			return diag.Errorf("error reading file_filter_profile: %v", err)
		}
	}

	if o.Fixedport != nil {
		v := *o.Fixedport

		if err = d.Set("fixedport", v); err != nil {
			return diag.Errorf("error reading fixedport: %v", err)
		}
	}

	if o.FssoGroups != nil {
		if err = d.Set("fsso_groups", flattenFirewallconsolidatedPolicyFssoGroups(o.FssoGroups, sort)); err != nil {
			return diag.Errorf("error reading fsso_groups: %v", err)
		}
	}

	if o.Groups != nil {
		if err = d.Set("groups", flattenFirewallconsolidatedPolicyGroups(o.Groups, sort)); err != nil {
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
		if err = d.Set("internet_service_custom", flattenFirewallconsolidatedPolicyInternetServiceCustom(o.InternetServiceCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceCustomGroup != nil {
		if err = d.Set("internet_service_custom_group", flattenFirewallconsolidatedPolicyInternetServiceCustomGroup(o.InternetServiceCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom_group: %v", err)
		}
	}

	if o.InternetServiceGroup != nil {
		if err = d.Set("internet_service_group", flattenFirewallconsolidatedPolicyInternetServiceGroup(o.InternetServiceGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_group: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		if err = d.Set("internet_service_id", flattenFirewallconsolidatedPolicyInternetServiceId(o.InternetServiceId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
		}
	}

	if o.InternetServiceName != nil {
		if err = d.Set("internet_service_name", flattenFirewallconsolidatedPolicyInternetServiceName(o.InternetServiceName, sort)); err != nil {
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
		if err = d.Set("internet_service_src_custom", flattenFirewallconsolidatedPolicyInternetServiceSrcCustom(o.InternetServiceSrcCustom, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom: %v", err)
		}
	}

	if o.InternetServiceSrcCustomGroup != nil {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallconsolidatedPolicyInternetServiceSrcCustomGroup(o.InternetServiceSrcCustomGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_custom_group: %v", err)
		}
	}

	if o.InternetServiceSrcGroup != nil {
		if err = d.Set("internet_service_src_group", flattenFirewallconsolidatedPolicyInternetServiceSrcGroup(o.InternetServiceSrcGroup, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_group: %v", err)
		}
	}

	if o.InternetServiceSrcId != nil {
		if err = d.Set("internet_service_src_id", flattenFirewallconsolidatedPolicyInternetServiceSrcId(o.InternetServiceSrcId, sort)); err != nil {
			return diag.Errorf("error reading internet_service_src_id: %v", err)
		}
	}

	if o.InternetServiceSrcName != nil {
		if err = d.Set("internet_service_src_name", flattenFirewallconsolidatedPolicyInternetServiceSrcName(o.InternetServiceSrcName, sort)); err != nil {
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

	if o.Outbound != nil {
		v := *o.Outbound

		if err = d.Set("outbound", v); err != nil {
			return diag.Errorf("error reading outbound: %v", err)
		}
	}

	if o.PerIpShaper != nil {
		v := *o.PerIpShaper

		if err = d.Set("per_ip_shaper", v); err != nil {
			return diag.Errorf("error reading per_ip_shaper: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
		}
	}

	if o.Poolname4 != nil {
		if err = d.Set("poolname4", flattenFirewallconsolidatedPolicyPoolname4(o.Poolname4, sort)); err != nil {
			return diag.Errorf("error reading poolname4: %v", err)
		}
	}

	if o.Poolname6 != nil {
		if err = d.Set("poolname6", flattenFirewallconsolidatedPolicyPoolname6(o.Poolname6, sort)); err != nil {
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

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallconsolidatedPolicyService(o.Service, sort)); err != nil {
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

	if o.SrcaddrNegate != nil {
		v := *o.SrcaddrNegate

		if err = d.Set("srcaddr_negate", v); err != nil {
			return diag.Errorf("error reading srcaddr_negate: %v", err)
		}
	}

	if o.Srcaddr4 != nil {
		if err = d.Set("srcaddr4", flattenFirewallconsolidatedPolicySrcaddr4(o.Srcaddr4, sort)); err != nil {
			return diag.Errorf("error reading srcaddr4: %v", err)
		}
	}

	if o.Srcaddr6 != nil {
		if err = d.Set("srcaddr6", flattenFirewallconsolidatedPolicySrcaddr6(o.Srcaddr6, sort)); err != nil {
			return diag.Errorf("error reading srcaddr6: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenFirewallconsolidatedPolicySrcintf(o.Srcintf, sort)); err != nil {
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

	if o.Users != nil {
		if err = d.Set("users", flattenFirewallconsolidatedPolicyUsers(o.Users, sort)); err != nil {
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

	return nil
}

func expandFirewallconsolidatedPolicyDstaddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyDstaddr4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyDstaddr4

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyDstaddr4{}
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

func expandFirewallconsolidatedPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyDstaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyDstaddr6

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyDstaddr6{}
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

func expandFirewallconsolidatedPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyDstintf

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyDstintf{}
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

func expandFirewallconsolidatedPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyFssoGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyFssoGroups

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyFssoGroups{}
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

func expandFirewallconsolidatedPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyGroups, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyGroups

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyGroups{}
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

func expandFirewallconsolidatedPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceCustom

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceCustom{}
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

func expandFirewallconsolidatedPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceCustomGroup

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceCustomGroup{}
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

func expandFirewallconsolidatedPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceGroup

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceGroup{}
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

func expandFirewallconsolidatedPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceId

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceId{}
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

func expandFirewallconsolidatedPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceName

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceName{}
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

func expandFirewallconsolidatedPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceSrcCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceSrcCustom

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceSrcCustom{}
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

func expandFirewallconsolidatedPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceSrcCustomGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceSrcCustomGroup

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceSrcCustomGroup{}
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

func expandFirewallconsolidatedPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceSrcGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceSrcGroup

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceSrcGroup{}
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

func expandFirewallconsolidatedPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceSrcId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceSrcId

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceSrcId{}
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

func expandFirewallconsolidatedPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyInternetServiceSrcName, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyInternetServiceSrcName

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyInternetServiceSrcName{}
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

func expandFirewallconsolidatedPolicyPoolname4(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyPoolname4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyPoolname4

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyPoolname4{}
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

func expandFirewallconsolidatedPolicyPoolname6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyPoolname6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyPoolname6

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyPoolname6{}
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

func expandFirewallconsolidatedPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyService, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyService

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyService{}
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

func expandFirewallconsolidatedPolicySrcaddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicySrcaddr4, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicySrcaddr4

	for i := range l {
		tmp := models.FirewallconsolidatedPolicySrcaddr4{}
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

func expandFirewallconsolidatedPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicySrcaddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicySrcaddr6

	for i := range l {
		tmp := models.FirewallconsolidatedPolicySrcaddr6{}
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

func expandFirewallconsolidatedPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicySrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicySrcintf

	for i := range l {
		tmp := models.FirewallconsolidatedPolicySrcintf{}
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

func expandFirewallconsolidatedPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallconsolidatedPolicyUsers, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallconsolidatedPolicyUsers

	for i := range l {
		tmp := models.FirewallconsolidatedPolicyUsers{}
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

func getObjectFirewallconsolidatedPolicy(d *schema.ResourceData, sv string) (*models.FirewallconsolidatedPolicy, diag.Diagnostics) {
	obj := models.FirewallconsolidatedPolicy{}
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
	if v1, ok := d.GetOk("application_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("application_list", sv)
				diags = append(diags, e)
			}
			obj.ApplicationList = &v2
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
	if v1, ok := d.GetOk("captive_portal_exempt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("captive_portal_exempt", sv)
				diags = append(diags, e)
			}
			obj.CaptivePortalExempt = &v2
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
	if v1, ok := d.GetOk("dstaddr_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstaddr_negate", sv)
				diags = append(diags, e)
			}
			obj.DstaddrNegate = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr4"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr4", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyDstaddr4(d, v, "dstaddr4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr4 = t
		}
	} else if d.HasChange("dstaddr4") {
		old, new := d.GetChange("dstaddr4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr4 = &[]models.FirewallconsolidatedPolicyDstaddr4{}
		}
	}
	if v, ok := d.GetOk("dstaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr6 = t
		}
	} else if d.HasChange("dstaddr6") {
		old, new := d.GetChange("dstaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr6 = &[]models.FirewallconsolidatedPolicyDstaddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.FirewallconsolidatedPolicyDstintf{}
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
	if v1, ok := d.GetOk("file_filter_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("file_filter_profile", sv)
				diags = append(diags, e)
			}
			obj.FileFilterProfile = &v2
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
	if v, ok := d.GetOk("fsso_groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fsso_groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyFssoGroups(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FssoGroups = t
		}
	} else if d.HasChange("fsso_groups") {
		old, new := d.GetChange("fsso_groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FssoGroups = &[]models.FirewallconsolidatedPolicyFssoGroups{}
		}
	}
	if v, ok := d.GetOk("groups"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("groups", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Groups = t
		}
	} else if d.HasChange("groups") {
		old, new := d.GetChange("groups")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Groups = &[]models.FirewallconsolidatedPolicyGroups{}
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
		t, err := expandFirewallconsolidatedPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.FirewallconsolidatedPolicyInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustomGroup = t
		}
	} else if d.HasChange("internet_service_custom_group") {
		old, new := d.GetChange("internet_service_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustomGroup = &[]models.FirewallconsolidatedPolicyInternetServiceCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceGroup = t
		}
	} else if d.HasChange("internet_service_group") {
		old, new := d.GetChange("internet_service_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceGroup = &[]models.FirewallconsolidatedPolicyInternetServiceGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("internet_service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceId = t
		}
	} else if d.HasChange("internet_service_id") {
		old, new := d.GetChange("internet_service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceId = &[]models.FirewallconsolidatedPolicyInternetServiceId{}
		}
	}
	if v, ok := d.GetOk("internet_service_name"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("internet_service_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceName = t
		}
	} else if d.HasChange("internet_service_name") {
		old, new := d.GetChange("internet_service_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceName = &[]models.FirewallconsolidatedPolicyInternetServiceName{}
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
		t, err := expandFirewallconsolidatedPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustom = t
		}
	} else if d.HasChange("internet_service_src_custom") {
		old, new := d.GetChange("internet_service_src_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustom = &[]models.FirewallconsolidatedPolicyInternetServiceSrcCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_custom_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcCustomGroup = t
		}
	} else if d.HasChange("internet_service_src_custom_group") {
		old, new := d.GetChange("internet_service_src_custom_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcCustomGroup = &[]models.FirewallconsolidatedPolicyInternetServiceSrcCustomGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_src_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcGroup = t
		}
	} else if d.HasChange("internet_service_src_group") {
		old, new := d.GetChange("internet_service_src_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcGroup = &[]models.FirewallconsolidatedPolicyInternetServiceSrcGroup{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_id"); ok {
		if !utils.CheckVer(sv, "", "v6.4.2") {
			e := utils.AttributeVersionWarning("internet_service_src_id", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceSrcId(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcId = t
		}
	} else if d.HasChange("internet_service_src_id") {
		old, new := d.GetChange("internet_service_src_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcId = &[]models.FirewallconsolidatedPolicyInternetServiceSrcId{}
		}
	}
	if v, ok := d.GetOk("internet_service_src_name"); ok {
		if !utils.CheckVer(sv, "v6.4.2", "") {
			e := utils.AttributeVersionWarning("internet_service_src_name", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyInternetServiceSrcName(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceSrcName = t
		}
	} else if d.HasChange("internet_service_src_name") {
		old, new := d.GetChange("internet_service_src_name")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceSrcName = &[]models.FirewallconsolidatedPolicyInternetServiceSrcName{}
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
	if v1, ok := d.GetOk("outbound"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outbound", sv)
				diags = append(diags, e)
			}
			obj.Outbound = &v2
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
	if v, ok := d.GetOk("poolname4"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("poolname4", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyPoolname4(d, v, "poolname4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Poolname4 = t
		}
	} else if d.HasChange("poolname4") {
		old, new := d.GetChange("poolname4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Poolname4 = &[]models.FirewallconsolidatedPolicyPoolname4{}
		}
	}
	if v, ok := d.GetOk("poolname6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("poolname6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyPoolname6(d, v, "poolname6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Poolname6 = t
		}
	} else if d.HasChange("poolname6") {
		old, new := d.GetChange("poolname6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Poolname6 = &[]models.FirewallconsolidatedPolicyPoolname6{}
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
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallconsolidatedPolicyService{}
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
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SessionTtl = &tmp
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
	if v, ok := d.GetOk("srcaddr4"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr4", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicySrcaddr4(d, v, "srcaddr4", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr4 = t
		}
	} else if d.HasChange("srcaddr4") {
		old, new := d.GetChange("srcaddr4")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr4 = &[]models.FirewallconsolidatedPolicySrcaddr4{}
		}
	}
	if v, ok := d.GetOk("srcaddr6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr6 = t
		}
	} else if d.HasChange("srcaddr6") {
		old, new := d.GetChange("srcaddr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr6 = &[]models.FirewallconsolidatedPolicySrcaddr6{}
		}
	}
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.FirewallconsolidatedPolicySrcintf{}
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
	if v, ok := d.GetOk("users"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("users", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallconsolidatedPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Users = t
		}
	} else if d.HasChange("users") {
		old, new := d.GetChange("users")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Users = &[]models.FirewallconsolidatedPolicyUsers{}
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
	return &obj, diags
}
