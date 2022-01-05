// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallPolicy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 policies.",

		ReadContext: dataSourceFirewallPolicy6Read,

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
				Description: "Log field index numbers to append custom log fields to log messages for this policy.",
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
				Description: "Destination address and address group names.",
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
				Description: "When enabled dstaddr specifies what the destination address must NOT be.",
				Computed:    true,
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
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Name of an existing email filter profile.",
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
			"natinbound": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: apply destination NAT to inbound traffic.",
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
			"outbound": {
				Type:        schema.TypeString,
				Description: "Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.",
				Computed:    true,
			},
			"per_ip_shaper": {
				Type:        schema.TypeString,
				Description: "Per-IP traffic shaper.",
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
			"replacemsg_override_group": {
				Type:        schema.TypeString,
				Description: "Override the default replacement message group for this policy.",
				Computed:    true,
			},
			"rsso": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS single sign-on (RSSO).",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
				Computed:    true,
			},
			"send_deny_packet": {
				Type:        schema.TypeString,
				Description: "Enable/disable return of deny-packet.",
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
							Description: "Address name.",
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
				Description: "Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.",
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address and address group names.",
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
				Description: "When enabled srcaddr specifies what the source address must NOT be.",
				Computed:    true,
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
							Description: "Interface name.",
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
				Description: "Reverse traffic shaper.",
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
				Description: "Enable AV/web/ips protection profile.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"vlan_cos_fwd": {
				Type:        schema.TypeInt,
				Description: "VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest",
				Computed:    true,
			},
			"vlan_cos_rev": {
				Type:        schema.TypeInt,
				Description: "VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest",
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
				Description: "Web proxy forward server name.",
				Computed:    true,
			},
			"webproxy_profile": {
				Type:        schema.TypeString,
				Description: "Webproxy profile name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallPolicy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallPolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallPolicy6 dataSource: %v", err)
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

	diags := refreshObjectFirewallPolicy6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
