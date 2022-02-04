// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VDOM settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VDOM settings.",

		ReadContext: dataSourceSystemSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_linkdown_path": {
				Type:        schema.TypeString,
				Description: "Enable/disable link down path.",
				Computed:    true,
			},
			"allow_subnet_overlap": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing interface subnets to use overlapping IP addresses.",
				Computed:    true,
			},
			"application_bandwidth_tracking": {
				Type:        schema.TypeString,
				Description: "Enable/disable application bandwidth tracking.",
				Computed:    true,
			},
			"asymroute": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv4 asymmetric routing.",
				Computed:    true,
			},
			"asymroute_icmp": {
				Type:        schema.TypeString,
				Description: "Enable/disable ICMP asymmetric routing.",
				Computed:    true,
			},
			"asymroute6": {
				Type:        schema.TypeString,
				Description: "Enable/disable asymmetric IPv6 routing.",
				Computed:    true,
			},
			"asymroute6_icmp": {
				Type:        schema.TypeString,
				Description: "Enable/disable asymmetric ICMPv6 routing.",
				Computed:    true,
			},
			"auxiliary_session": {
				Type:        schema.TypeString,
				Description: "Enable/disable auxiliary session.",
				Computed:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces.",
				Computed:    true,
			},
			"bfd_desired_min_tx": {
				Type:        schema.TypeInt,
				Description: "BFD desired minimal transmit interval (1 - 100000 ms, default = 250).",
				Computed:    true,
			},
			"bfd_detect_mult": {
				Type:        schema.TypeInt,
				Description: "BFD detection multiplier (1 - 50, default = 3).",
				Computed:    true,
			},
			"bfd_dont_enforce_src_port": {
				Type:        schema.TypeString,
				Description: "Enable to not enforce verifying the source port of BFD Packets.",
				Computed:    true,
			},
			"bfd_required_min_rx": {
				Type:        schema.TypeInt,
				Description: "BFD required minimal receive interval (1 - 100000 ms, default = 250).",
				Computed:    true,
			},
			"block_land_attack": {
				Type:        schema.TypeString,
				Description: "Enable/disable blocking of land attacks.",
				Computed:    true,
			},
			"central_nat": {
				Type:        schema.TypeString,
				Description: "Enable/disable central NAT.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "VDOM comments.",
				Computed:    true,
			},
			"consolidated_firewall_mode": {
				Type:        schema.TypeString,
				Description: "Consolidated firewall mode.",
				Computed:    true,
			},
			"default_voip_alg_mode": {
				Type:        schema.TypeString,
				Description: "Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile.",
				Computed:    true,
			},
			"deny_tcp_with_icmp": {
				Type:        schema.TypeString,
				Description: "Enable/disable denying TCP by sending an ICMP communication prohibited packet.",
				Computed:    true,
			},
			"device": {
				Type:        schema.TypeString,
				Description: "Interface to use for management access for NAT mode.",
				Computed:    true,
			},
			"dhcp_proxy": {
				Type:        schema.TypeString,
				Description: "Enable/disable the DHCP Proxy.",
				Computed:    true,
			},
			"dhcp_proxy_interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"dhcp_proxy_interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"dhcp_server_ip": {
				Type:        schema.TypeString,
				Description: "DHCP Server IPv4 address.",
				Computed:    true,
			},
			"dhcp6_server_ip": {
				Type:        schema.TypeString,
				Description: "DHCPv6 server IPv6 address.",
				Computed:    true,
			},
			"discovered_device_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout for discovered devices (1 - 365 days, default = 28).",
				Computed:    true,
			},
			"ecmp_max_paths": {
				Type:        schema.TypeInt,
				Description: "Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 255, default = 255).",
				Computed:    true,
			},
			"email_portal_check_dns": {
				Type:        schema.TypeString,
				Description: "Enable/disable using DNS to validate email addresses collected by a captive portal.",
				Computed:    true,
			},
			"firewall_session_dirty": {
				Type:        schema.TypeString,
				Description: "Select how to manage sessions affected by firewall policy configuration changes.",
				Computed:    true,
			},
			"fw_session_hairpin": {
				Type:        schema.TypeString,
				Description: "Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate.",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "Transparent mode IPv4 default gateway IP address.",
				Computed:    true,
			},
			"gateway6": {
				Type:        schema.TypeString,
				Description: "Transparent mode IPv4 default gateway IP address.",
				Computed:    true,
			},
			"gui_advanced_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable advanced policy configuration on the GUI.",
				Computed:    true,
			},
			"gui_allow_unnamed_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable the requirement for policy naming on the GUI.",
				Computed:    true,
			},
			"gui_antivirus": {
				Type:        schema.TypeString,
				Description: "Enable/disable AntiVirus on the GUI.",
				Computed:    true,
			},
			"gui_ap_profile": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiAP profiles on the GUI.",
				Computed:    true,
			},
			"gui_application_control": {
				Type:        schema.TypeString,
				Description: "Enable/disable application control on the GUI.",
				Computed:    true,
			},
			"gui_default_policy_columns": {
				Type:        schema.TypeList,
				Description: "Default columns to display for policy lists on GUI.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Select column name.",
							Computed:    true,
						},
					},
				},
			},
			"gui_dhcp_advanced": {
				Type:        schema.TypeString,
				Description: "Enable/disable advanced DHCP options on the GUI.",
				Computed:    true,
			},
			"gui_dns_database": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNS database settings on the GUI.",
				Computed:    true,
			},
			"gui_dnsfilter": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNS Filtering on the GUI.",
				Computed:    true,
			},
			"gui_domain_ip_reputation": {
				Type:        schema.TypeString,
				Description: "Enable/disable Domain and IP Reputation on the GUI.",
				Computed:    true,
			},
			"gui_dos_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable DoS policies on the GUI.",
				Computed:    true,
			},
			"gui_dynamic_profile_display": {
				Type:        schema.TypeString,
				Description: "Enable/disable RADIUS Single Sign On (RSSO) on the GUI.",
				Computed:    true,
			},
			"gui_dynamic_routing": {
				Type:        schema.TypeString,
				Description: "Enable/disable dynamic routing on the GUI.",
				Computed:    true,
			},
			"gui_email_collection": {
				Type:        schema.TypeString,
				Description: "Enable/disable email collection on the GUI.",
				Computed:    true,
			},
			"gui_endpoint_control": {
				Type:        schema.TypeString,
				Description: "Enable/disable endpoint control on the GUI.",
				Computed:    true,
			},
			"gui_endpoint_control_advanced": {
				Type:        schema.TypeString,
				Description: "Enable/disable advanced endpoint control options on the GUI.",
				Computed:    true,
			},
			"gui_explicit_proxy": {
				Type:        schema.TypeString,
				Description: "Enable/disable the explicit proxy on the GUI.",
				Computed:    true,
			},
			"gui_file_filter": {
				Type:        schema.TypeString,
				Description: "Enable/disable File-filter on the GUI.",
				Computed:    true,
			},
			"gui_fortiap_split_tunneling": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiAP split tunneling on the GUI.",
				Computed:    true,
			},
			"gui_fortiextender_controller": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiExtender on the GUI.",
				Computed:    true,
			},
			"gui_icap": {
				Type:        schema.TypeString,
				Description: "Enable/disable ICAP on the GUI.",
				Computed:    true,
			},
			"gui_implicit_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable implicit firewall policies on the GUI.",
				Computed:    true,
			},
			"gui_ips": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS on the GUI.",
				Computed:    true,
			},
			"gui_load_balance": {
				Type:        schema.TypeString,
				Description: "Enable/disable server load balancing on the GUI.",
				Computed:    true,
			},
			"gui_local_in_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable Local-In policies on the GUI.",
				Computed:    true,
			},
			"gui_local_reports": {
				Type:        schema.TypeString,
				Description: "Enable/disable local reports on the GUI.",
				Computed:    true,
			},
			"gui_multicast_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable multicast firewall policies on the GUI.",
				Computed:    true,
			},
			"gui_multiple_interface_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable adding multiple interfaces to a policy on the GUI.",
				Computed:    true,
			},
			"gui_multiple_utm_profiles": {
				Type:        schema.TypeString,
				Description: "Enable/disable multiple UTM profiles on the GUI.",
				Computed:    true,
			},
			"gui_nat46_64": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT46 and NAT64 settings on the GUI.",
				Computed:    true,
			},
			"gui_object_colors": {
				Type:        schema.TypeString,
				Description: "Enable/disable object colors on the GUI.",
				Computed:    true,
			},
			"gui_per_policy_disclaimer": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy disclaimer on the GUI.",
				Computed:    true,
			},
			"gui_policy_based_ipsec": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy-based IPsec VPN on the GUI.",
				Computed:    true,
			},
			"gui_policy_disclaimer": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy disclaimer on the GUI.",
				Computed:    true,
			},
			"gui_replacement_message_groups": {
				Type:        schema.TypeString,
				Description: "Enable/disable replacement message groups on the GUI.",
				Computed:    true,
			},
			"gui_security_profile_group": {
				Type:        schema.TypeString,
				Description: "Enable/disable Security Profile Groups on the GUI.",
				Computed:    true,
			},
			"gui_spamfilter": {
				Type:        schema.TypeString,
				Description: "Enable/disable Antispam on the GUI.",
				Computed:    true,
			},
			"gui_sslvpn_personal_bookmarks": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN personal bookmark management on the GUI.",
				Computed:    true,
			},
			"gui_sslvpn_realms": {
				Type:        schema.TypeString,
				Description: "Enable/disable SSL-VPN realms on the GUI.",
				Computed:    true,
			},
			"gui_switch_controller": {
				Type:        schema.TypeString,
				Description: "Enable/disable the switch controller on the GUI.",
				Computed:    true,
			},
			"gui_threat_weight": {
				Type:        schema.TypeString,
				Description: "Enable/disable threat weight on the GUI.",
				Computed:    true,
			},
			"gui_traffic_shaping": {
				Type:        schema.TypeString,
				Description: "Enable/disable traffic shaping on the GUI.",
				Computed:    true,
			},
			"gui_videofilter": {
				Type:        schema.TypeString,
				Description: "Enable/disable Video filtering on the GUI.",
				Computed:    true,
			},
			"gui_voip_profile": {
				Type:        schema.TypeString,
				Description: "Enable/disable VoIP profiles on the GUI.",
				Computed:    true,
			},
			"gui_vpn": {
				Type:        schema.TypeString,
				Description: "Enable/disable VPN tunnels on the GUI.",
				Computed:    true,
			},
			"gui_waf_profile": {
				Type:        schema.TypeString,
				Description: "Enable/disable Web Application Firewall on the GUI.",
				Computed:    true,
			},
			"gui_wan_load_balancing": {
				Type:        schema.TypeString,
				Description: "Enable/disable SD-WAN on the GUI.",
				Computed:    true,
			},
			"gui_wanopt_cache": {
				Type:        schema.TypeString,
				Description: "Enable/disable WAN Optimization and Web Caching on the GUI.",
				Computed:    true,
			},
			"gui_webfilter": {
				Type:        schema.TypeString,
				Description: "Enable/disable Web filtering on the GUI.",
				Computed:    true,
			},
			"gui_webfilter_advanced": {
				Type:        schema.TypeString,
				Description: "Enable/disable advanced web filtering on the GUI.",
				Computed:    true,
			},
			"gui_wireless_controller": {
				Type:        schema.TypeString,
				Description: "Enable/disable the wireless controller on the GUI.",
				Computed:    true,
			},
			"gui_ztna": {
				Type:        schema.TypeString,
				Description: "Enable/disable Zero Trust Network Access features on the GUI.",
				Computed:    true,
			},
			"h323_direct_model": {
				Type:        schema.TypeString,
				Description: "Enable/disable H323 direct model.",
				Computed:    true,
			},
			"http_external_dest": {
				Type:        schema.TypeString,
				Description: "Offload HTTP traffic to FortiWeb or FortiCache.",
				Computed:    true,
			},
			"ike_dn_format": {
				Type:        schema.TypeString,
				Description: "Configure IKE ASN.1 Distinguished Name format conventions.",
				Computed:    true,
			},
			"ike_policy_route": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKE Policy Based Routing (PBR).",
				Computed:    true,
			},
			"ike_port": {
				Type:        schema.TypeInt,
				Description: "UDP port for IKE/IPsec traffic (default 500).",
				Computed:    true,
			},
			"ike_quick_crash_detect": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKE quick crash detection (RFC 6290).",
				Computed:    true,
			},
			"ike_session_resume": {
				Type:        schema.TypeString,
				Description: "Enable/disable IKEv2 session resumption (RFC 5723).",
				Computed:    true,
			},
			"implicit_allow_dns": {
				Type:        schema.TypeString,
				Description: "Enable/disable implicitly allowing DNS traffic.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "IP address and netmask.",
				Computed:    true,
			},
			"ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 address prefix for NAT mode.",
				Computed:    true,
			},
			"link_down_access": {
				Type:        schema.TypeString,
				Description: "Enable/disable link down access traffic.",
				Computed:    true,
			},
			"lldp_reception": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM.",
				Computed:    true,
			},
			"lldp_transmission": {
				Type:        schema.TypeString,
				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM.",
				Computed:    true,
			},
			"location_id": {
				Type:        schema.TypeString,
				Description: "Local location ID in the form of an IPv4 address.",
				Computed:    true,
			},
			"mac_ttl": {
				Type:        schema.TypeInt,
				Description: "Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).",
				Computed:    true,
			},
			"manageip": {
				Type:        schema.TypeString,
				Description: "Transparent mode IPv4 management IP address and netmask.",
				Computed:    true,
			},
			"manageip6": {
				Type:        schema.TypeString,
				Description: "Transparent mode IPv6 management IP address and netmask.",
				Computed:    true,
			},
			"multicast_forward": {
				Type:        schema.TypeString,
				Description: "Enable/disable multicast forwarding.",
				Computed:    true,
			},
			"multicast_skip_policy": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing multicast traffic through the FortiGate without a policy check.",
				Computed:    true,
			},
			"multicast_ttl_notchange": {
				Type:        schema.TypeString,
				Description: "Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets.",
				Computed:    true,
			},
			"ngfw_mode": {
				Type:        schema.TypeString,
				Description: "Next Generation Firewall (NGFW) mode.",
				Computed:    true,
			},
			"opmode": {
				Type:        schema.TypeString,
				Description: "Firewall operation mode (NAT or Transparent).",
				Computed:    true,
			},
			"prp_trailer_action": {
				Type:        schema.TypeString,
				Description: "Enable/disable action to take on PRP trailer.",
				Computed:    true,
			},
			"sccp_port": {
				Type:        schema.TypeInt,
				Description: "TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).",
				Computed:    true,
			},
			"sctp_session_without_init": {
				Type:        schema.TypeString,
				Description: "Enable/disable SCTP session creation without SCTP INIT.",
				Computed:    true,
			},
			"ses_denied_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable including denied session in the session table.",
				Computed:    true,
			},
			"sip_expectation": {
				Type:        schema.TypeString,
				Description: "Enable/disable the SIP kernel session helper to create an expectation for port 5060.",
				Computed:    true,
			},
			"sip_nat_trace": {
				Type:        schema.TypeString,
				Description: "Enable/disable recording the original SIP source IP address when NAT is used.",
				Computed:    true,
			},
			"sip_ssl_port": {
				Type:        schema.TypeInt,
				Description: "TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).",
				Computed:    true,
			},
			"sip_tcp_port": {
				Type:        schema.TypeInt,
				Description: "TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).",
				Computed:    true,
			},
			"sip_udp_port": {
				Type:        schema.TypeInt,
				Description: "UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).",
				Computed:    true,
			},
			"snat_hairpin_traffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable source NAT (SNAT) for hairpin traffic.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this VDOM.",
				Computed:    true,
			},
			"strict_src_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable strict source verification.",
				Computed:    true,
			},
			"tcp_session_without_syn": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing TCP session without SYN flags.",
				Computed:    true,
			},
			"utf8_spam_tagging": {
				Type:        schema.TypeString,
				Description: "Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support.",
				Computed:    true,
			},
			"v4_ecmp_mode": {
				Type:        schema.TypeString,
				Description: "IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode.",
				Computed:    true,
			},
			"vpn_stats_log": {
				Type:        schema.TypeString,
				Description: "Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space.",
				Computed:    true,
			},
			"vpn_stats_period": {
				Type:        schema.TypeInt,
				Description: "Period to send VPN log statistics (0 or 60 - 86400 sec).",
				Computed:    true,
			},
			"wccp_cache_engine": {
				Type:        schema.TypeString,
				Description: "Enable/disable WCCP cache engine.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemSettings"

	o, err := c.Cmdb.ReadSystemSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSettings dataSource: %v", err)
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

	diags := refreshObjectSystemSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
