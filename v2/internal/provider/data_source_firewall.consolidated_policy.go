// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure consolidated IPv4/IPv6 policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallConsolidatedPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure consolidated IPv4/IPv6 policies.",

		ReadContext: dataSourceFirewallConsolidatedPolicyRead,

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
			"captive_portal_exempt": {
				Type:        schema.TypeString,
				Description: "Enable exemption of some users from the captive portal.",
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
			"diffserv_forward": {
				Type:        schema.TypeString,
				Description: "Enable to change packet's DiffServ values to the specified diffservcode-forward value.",
				Computed:    true,
			},
			"diffserv_reverse": {
				Type:        schema.TypeString,
				Description: "Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. ",
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
			"dstaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled dstaddr specifies what the destination address must NOT be.",
				Computed:    true,
			},
			"dstaddr4": {
				Type:        schema.TypeList,
				Description: "Destination IPv4 address name and address group names.",
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
				Computed:    true,
			},
			"poolname4": {
				Type:        schema.TypeList,
				Description: "IPv4 pool names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IPv4 pool name.",
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
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
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
							Description: "Service name.",
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
				Type:        schema.TypeInt,
				Description: "TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).",
				Computed:    true,
			},
			"srcaddr_negate": {
				Type:        schema.TypeString,
				Description: "When enabled srcaddr specifies what the source address must NOT be.",
				Computed:    true,
			},
			"srcaddr4": {
				Type:        schema.TypeList,
				Description: "Source IPv4 address name and address group names.",
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
			"users": {
				Type:        schema.TypeList,
				Description: "Names of individual users that can authenticate with this policy.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "User name.",
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
				Description: "WAN optimization passive mode options. This option decides what IP address will be used to connect to server.",
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
		},
	}
}

func dataSourceFirewallConsolidatedPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallConsolidatedPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallConsolidatedPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallConsolidatedPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
