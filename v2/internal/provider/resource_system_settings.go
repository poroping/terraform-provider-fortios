// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VDOM settings.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VDOM settings.",

		CreateContext: resourceSystemSettingsCreate,
		ReadContext:   resourceSystemSettingsRead,
		UpdateContext: resourceSystemSettingsUpdate,
		DeleteContext: resourceSystemSettingsDelete,

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
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allow_linkdown_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable link down path.",
				Optional:    true,
				Computed:    true,
			},
			"allow_subnet_overlap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing interface subnets to use overlapping IP addresses.",
				Optional:    true,
				Computed:    true,
			},
			"application_bandwidth_tracking": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable application bandwidth tracking.",
				Optional:    true,
				Computed:    true,
			},
			"asymroute": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPv4 asymmetric routing.",
				Optional:    true,
				Computed:    true,
			},
			"asymroute_icmp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ICMP asymmetric routing.",
				Optional:    true,
				Computed:    true,
			},
			"asymroute6": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable asymmetric IPv6 routing.",
				Optional:    true,
				Computed:    true,
			},
			"asymroute6_icmp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable asymmetric ICMPv6 routing.",
				Optional:    true,
				Computed:    true,
			},
			"auxiliary_session": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable auxiliary session.",
				Optional:    true,
				Computed:    true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces.",
				Optional:    true,
				Computed:    true,
			},
			"bfd_desired_min_tx": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),

				Description: "BFD desired minimal transmit interval (1 - 100000 ms, default = 250).",
				Optional:    true,
				Computed:    true,
			},
			"bfd_detect_mult": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),

				Description: "BFD detection multiplier (1 - 50, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"bfd_dont_enforce_src_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable to not enforce verifying the source port of BFD Packets.",
				Optional:    true,
				Computed:    true,
			},
			"bfd_required_min_rx": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),

				Description: "BFD required minimal receive interval (1 - 100000 ms, default = 250).",
				Optional:    true,
				Computed:    true,
			},
			"block_land_attack": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable blocking of land attacks.",
				Optional:    true,
				Computed:    true,
			},
			"central_nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable central NAT.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "VDOM comments.",
				Optional:    true,
				Computed:    true,
			},
			"consolidated_firewall_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{}, false),

				Description: "Consolidated firewall mode.",
				Optional:    true,
				Computed:    true,
			},
			"default_voip_alg_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"proxy-based", "kernel-helper-based"}, false),

				Description: "Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile.",
				Optional:    true,
				Computed:    true,
			},
			"deny_tcp_with_icmp": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable denying TCP by sending an ICMP communication prohibited packet.",
				Optional:    true,
				Computed:    true,
			},
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface to use for management access for NAT mode.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the DHCP Proxy.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_proxy_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_proxy_interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_server_ip": {
				Type: schema.TypeString,

				Description: "DHCP Server IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp6_server_ip": {
				Type: schema.TypeString,

				Description: "DHCPv6 server IPv6 address.",
				Optional:    true,
				Computed:    true,
			},
			"discovered_device_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 365),

				Description: "Timeout for discovered devices (1 - 365 days, default = 28).",
				Optional:    true,
				Computed:    true,
			},
			"ecmp_max_paths": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 255, default = 255).",
				Optional:    true,
				Computed:    true,
			},
			"email_portal_check_dns": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable using DNS to validate email addresses collected by a captive portal.",
				Optional:    true,
				Computed:    true,
			},
			"firewall_session_dirty": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"check-all", "check-new", "check-policy-option"}, false),

				Description: "Select how to manage sessions affected by firewall policy configuration changes.",
				Optional:    true,
				Computed:    true,
			},
			"fw_session_hairpin": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate.",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Transparent mode IPv4 default gateway IP address.",
				Optional:    true,
				Computed:    true,
			},
			"gateway6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Transparent mode IPv4 default gateway IP address.",
				Optional:         true,
				Computed:         true,
			},
			"gui_advanced_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable advanced policy configuration on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_allow_unnamed_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the requirement for policy naming on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_antivirus": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable AntiVirus on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_ap_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiAP profiles on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_application_control": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable application control on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_default_policy_columns": {
				Type:        schema.TypeList,
				Description: "Default columns to display for policy lists on GUI.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Select column name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"gui_dhcp_advanced": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable advanced DHCP options on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_dns_database": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DNS database settings on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_dnsfilter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DNS Filtering on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_domain_ip_reputation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Domain and IP Reputation on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_dos_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable DoS policies on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_dynamic_profile_display": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable RADIUS Single Sign On (RSSO) on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_dynamic_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable dynamic routing on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_email_collection": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable email collection on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_endpoint_control": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable endpoint control on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_endpoint_control_advanced": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable advanced endpoint control options on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_explicit_proxy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the explicit proxy on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_file_filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable File-filter on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_fortiap_split_tunneling": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiAP split tunneling on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_fortiextender_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiExtender on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_icap": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ICAP on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_implicit_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable implicit firewall policies on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_ips": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IPS on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_load_balance": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable server load balancing on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_local_in_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Local-In policies on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_local_reports": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable local reports on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_multicast_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multicast firewall policies on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_multiple_interface_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adding multiple interfaces to a policy on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_multiple_utm_profiles": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multiple UTM profiles on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_nat46_64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT46 and NAT64 settings on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_object_colors": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable object colors on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_per_policy_disclaimer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy disclaimer on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_policy_based_ipsec": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy-based IPsec VPN on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_policy_disclaimer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy disclaimer on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_replacement_message_groups": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable replacement message groups on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_security_profile_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Security Profile Groups on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_spamfilter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Antispam on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_sslvpn_personal_bookmarks": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN personal bookmark management on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_sslvpn_realms": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SSL-VPN realms on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_switch_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the switch controller on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_threat_weight": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable threat weight on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_traffic_shaping": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable traffic shaping on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_videofilter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Video filtering on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_voip_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VoIP profiles on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_vpn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VPN tunnels on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_waf_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Web Application Firewall on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_wan_load_balancing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SD-WAN on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_wanopt_cache": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WAN Optimization and Web Caching on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_webfilter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Web filtering on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_webfilter_advanced": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable advanced web filtering on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_wireless_controller": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the wireless controller on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"gui_ztna": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Zero Trust Network Access features on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"http_external_dest": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fortiweb", "forticache"}, false),

				Description: "Offload HTTP traffic to FortiWeb or FortiCache.",
				Optional:    true,
				Computed:    true,
			},
			"ike_dn_format": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"with-space", "no-space"}, false),

				Description: "Configure IKE ASN.1 Distinguished Name format conventions.",
				Optional:    true,
				Computed:    true,
			},
			"ike_policy_route": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKE Policy Based Routing (PBR).",
				Optional:    true,
				Computed:    true,
			},
			"ike_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 65535),

				Description: "UDP port for IKE/IPsec traffic (default 500).",
				Optional:    true,
				Computed:    true,
			},
			"ike_quick_crash_detect": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKE quick crash detection (RFC 6290).",
				Optional:    true,
				Computed:    true,
			},
			"ike_session_resume": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IKEv2 session resumption (RFC 5723).",
				Optional:    true,
				Computed:    true,
			},
			"implicit_allow_dns": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable implicitly allowing DNS traffic.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "IP address and netmask.",
				Optional:    true,
				Computed:    true,
			},
			"ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "IPv6 address prefix for NAT mode.",
				Optional:         true,
				Computed:         true,
			},
			"link_down_access": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable link down access traffic.",
				Optional:    true,
				Computed:    true,
			},
			"lldp_reception": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "global"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM.",
				Optional:    true,
				Computed:    true,
			},
			"lldp_transmission": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "global"}, false),

				Description: "Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM.",
				Optional:    true,
				Computed:    true,
			},
			"location_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Local location ID in the form of an IPv4 address.",
				Optional:    true,
				Computed:    true,
			},
			"mac_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),

				Description: "Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).",
				Optional:    true,
				Computed:    true,
			},
			"manageip": {
				Type: schema.TypeString,

				Description: "Transparent mode IPv4 management IP address and netmask.",
				Optional:    true,
				Computed:    true,
			},
			"manageip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Transparent mode IPv6 management IP address and netmask.",
				Optional:         true,
				Computed:         true,
			},
			"multicast_forward": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable multicast forwarding.",
				Optional:    true,
				Computed:    true,
			},
			"multicast_skip_policy": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing multicast traffic through the FortiGate without a policy check.",
				Optional:    true,
				Computed:    true,
			},
			"multicast_ttl_notchange": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets.",
				Optional:    true,
				Computed:    true,
			},
			"ngfw_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"profile-based", "policy-based"}, false),

				Description: "Next Generation Firewall (NGFW) mode.",
				Optional:    true,
				Computed:    true,
			},
			"opmode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"nat", "transparent"}, false),

				Description: "Firewall operation mode (NAT or Transparent).",
				Optional:    true,
				Computed:    true,
			},
			"prp_trailer_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable action to take on PRP trailer.",
				Optional:    true,
				Computed:    true,
			},
			"sccp_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).",
				Optional:    true,
				Computed:    true,
			},
			"sctp_session_without_init": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable SCTP session creation without SCTP INIT.",
				Optional:    true,
				Computed:    true,
			},
			"ses_denied_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable including denied session in the session table.",
				Optional:    true,
				Computed:    true,
			},
			"sip_expectation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the SIP kernel session helper to create an expectation for port 5060.",
				Optional:    true,
				Computed:    true,
			},
			"sip_nat_trace": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable recording the original SIP source IP address when NAT is used.",
				Optional:    true,
				Computed:    true,
			},
			"sip_ssl_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).",
				Optional:    true,
				Computed:    true,
			},
			"sip_tcp_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).",
				Optional:    true,
				Computed:    true,
			},
			"sip_udp_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).",
				Optional:    true,
				Computed:    true,
			},
			"snat_hairpin_traffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable source NAT (SNAT) for hairpin traffic.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this VDOM.",
				Optional:    true,
				Computed:    true,
			},
			"strict_src_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable strict source verification.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_session_without_syn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable allowing TCP session without SYN flags.",
				Optional:    true,
				Computed:    true,
			},
			"utf8_spam_tagging": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support.",
				Optional:    true,
				Computed:    true,
			},
			"v4_ecmp_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"source-ip-based", "weight-based", "usage-based", "source-dest-ip-based"}, false),

				Description: "IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode.",
				Optional:    true,
				Computed:    true,
			},
			"vpn_stats_log": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space.",
				Optional:         true,
				Computed:         true,
			},
			"vpn_stats_period": {
				Type: schema.TypeInt,

				Description: "Period to send VPN log statistics (0 or 60 - 86400 sec).",
				Optional:    true,
				Computed:    true,
			},
			"wccp_cache_engine": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable WCCP cache engine.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSettings(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSettings")
	}

	return resourceSystemSettingsRead(ctx, d, meta)
}

func resourceSystemSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSettings resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSettings")
	}

	return resourceSystemSettingsRead(ctx, d, meta)
}

func resourceSystemSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemSettings(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemSettings(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSettings resource: %v", err)
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

	diags := refreshObjectSystemSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSettingsGuiDefaultPolicyColumns(v *[]models.SystemSettingsGuiDefaultPolicyColumns, sort bool) interface{} {
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

func refreshObjectSystemSettings(d *schema.ResourceData, o *models.SystemSettings, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowLinkdownPath != nil {
		v := *o.AllowLinkdownPath

		if err = d.Set("allow_linkdown_path", v); err != nil {
			return diag.Errorf("error reading allow_linkdown_path: %v", err)
		}
	}

	if o.AllowSubnetOverlap != nil {
		v := *o.AllowSubnetOverlap

		if err = d.Set("allow_subnet_overlap", v); err != nil {
			return diag.Errorf("error reading allow_subnet_overlap: %v", err)
		}
	}

	if o.ApplicationBandwidthTracking != nil {
		v := *o.ApplicationBandwidthTracking

		if err = d.Set("application_bandwidth_tracking", v); err != nil {
			return diag.Errorf("error reading application_bandwidth_tracking: %v", err)
		}
	}

	if o.Asymroute != nil {
		v := *o.Asymroute

		if err = d.Set("asymroute", v); err != nil {
			return diag.Errorf("error reading asymroute: %v", err)
		}
	}

	if o.AsymrouteIcmp != nil {
		v := *o.AsymrouteIcmp

		if err = d.Set("asymroute_icmp", v); err != nil {
			return diag.Errorf("error reading asymroute_icmp: %v", err)
		}
	}

	if o.Asymroute6 != nil {
		v := *o.Asymroute6

		if err = d.Set("asymroute6", v); err != nil {
			return diag.Errorf("error reading asymroute6: %v", err)
		}
	}

	if o.Asymroute6Icmp != nil {
		v := *o.Asymroute6Icmp

		if err = d.Set("asymroute6_icmp", v); err != nil {
			return diag.Errorf("error reading asymroute6_icmp: %v", err)
		}
	}

	if o.AuxiliarySession != nil {
		v := *o.AuxiliarySession

		if err = d.Set("auxiliary_session", v); err != nil {
			return diag.Errorf("error reading auxiliary_session: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.BfdDesiredMinTx != nil {
		v := *o.BfdDesiredMinTx

		if err = d.Set("bfd_desired_min_tx", v); err != nil {
			return diag.Errorf("error reading bfd_desired_min_tx: %v", err)
		}
	}

	if o.BfdDetectMult != nil {
		v := *o.BfdDetectMult

		if err = d.Set("bfd_detect_mult", v); err != nil {
			return diag.Errorf("error reading bfd_detect_mult: %v", err)
		}
	}

	if o.BfdDontEnforceSrcPort != nil {
		v := *o.BfdDontEnforceSrcPort

		if err = d.Set("bfd_dont_enforce_src_port", v); err != nil {
			return diag.Errorf("error reading bfd_dont_enforce_src_port: %v", err)
		}
	}

	if o.BfdRequiredMinRx != nil {
		v := *o.BfdRequiredMinRx

		if err = d.Set("bfd_required_min_rx", v); err != nil {
			return diag.Errorf("error reading bfd_required_min_rx: %v", err)
		}
	}

	if o.BlockLandAttack != nil {
		v := *o.BlockLandAttack

		if err = d.Set("block_land_attack", v); err != nil {
			return diag.Errorf("error reading block_land_attack: %v", err)
		}
	}

	if o.CentralNat != nil {
		v := *o.CentralNat

		if err = d.Set("central_nat", v); err != nil {
			return diag.Errorf("error reading central_nat: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.ConsolidatedFirewallMode != nil {
		v := *o.ConsolidatedFirewallMode

		if err = d.Set("consolidated_firewall_mode", v); err != nil {
			return diag.Errorf("error reading consolidated_firewall_mode: %v", err)
		}
	}

	if o.DefaultVoipAlgMode != nil {
		v := *o.DefaultVoipAlgMode

		if err = d.Set("default_voip_alg_mode", v); err != nil {
			return diag.Errorf("error reading default_voip_alg_mode: %v", err)
		}
	}

	if o.DenyTcpWithIcmp != nil {
		v := *o.DenyTcpWithIcmp

		if err = d.Set("deny_tcp_with_icmp", v); err != nil {
			return diag.Errorf("error reading deny_tcp_with_icmp: %v", err)
		}
	}

	if o.Device != nil {
		v := *o.Device

		if err = d.Set("device", v); err != nil {
			return diag.Errorf("error reading device: %v", err)
		}
	}

	if o.DhcpProxy != nil {
		v := *o.DhcpProxy

		if err = d.Set("dhcp_proxy", v); err != nil {
			return diag.Errorf("error reading dhcp_proxy: %v", err)
		}
	}

	if o.DhcpProxyInterface != nil {
		v := *o.DhcpProxyInterface

		if err = d.Set("dhcp_proxy_interface", v); err != nil {
			return diag.Errorf("error reading dhcp_proxy_interface: %v", err)
		}
	}

	if o.DhcpProxyInterfaceSelectMethod != nil {
		v := *o.DhcpProxyInterfaceSelectMethod

		if err = d.Set("dhcp_proxy_interface_select_method", v); err != nil {
			return diag.Errorf("error reading dhcp_proxy_interface_select_method: %v", err)
		}
	}

	if o.DhcpServerIp != nil {
		v := *o.DhcpServerIp

		if err = d.Set("dhcp_server_ip", v); err != nil {
			return diag.Errorf("error reading dhcp_server_ip: %v", err)
		}
	}

	if o.Dhcp6ServerIp != nil {
		v := *o.Dhcp6ServerIp

		if err = d.Set("dhcp6_server_ip", v); err != nil {
			return diag.Errorf("error reading dhcp6_server_ip: %v", err)
		}
	}

	if o.DiscoveredDeviceTimeout != nil {
		v := *o.DiscoveredDeviceTimeout

		if err = d.Set("discovered_device_timeout", v); err != nil {
			return diag.Errorf("error reading discovered_device_timeout: %v", err)
		}
	}

	if o.EcmpMaxPaths != nil {
		v := *o.EcmpMaxPaths

		if err = d.Set("ecmp_max_paths", v); err != nil {
			return diag.Errorf("error reading ecmp_max_paths: %v", err)
		}
	}

	if o.EmailPortalCheckDns != nil {
		v := *o.EmailPortalCheckDns

		if err = d.Set("email_portal_check_dns", v); err != nil {
			return diag.Errorf("error reading email_portal_check_dns: %v", err)
		}
	}

	if o.FirewallSessionDirty != nil {
		v := *o.FirewallSessionDirty

		if err = d.Set("firewall_session_dirty", v); err != nil {
			return diag.Errorf("error reading firewall_session_dirty: %v", err)
		}
	}

	if o.FwSessionHairpin != nil {
		v := *o.FwSessionHairpin

		if err = d.Set("fw_session_hairpin", v); err != nil {
			return diag.Errorf("error reading fw_session_hairpin: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.Gateway6 != nil {
		v := *o.Gateway6

		if err = d.Set("gateway6", v); err != nil {
			return diag.Errorf("error reading gateway6: %v", err)
		}
	}

	if o.GuiAdvancedPolicy != nil {
		v := *o.GuiAdvancedPolicy

		if err = d.Set("gui_advanced_policy", v); err != nil {
			return diag.Errorf("error reading gui_advanced_policy: %v", err)
		}
	}

	if o.GuiAllowUnnamedPolicy != nil {
		v := *o.GuiAllowUnnamedPolicy

		if err = d.Set("gui_allow_unnamed_policy", v); err != nil {
			return diag.Errorf("error reading gui_allow_unnamed_policy: %v", err)
		}
	}

	if o.GuiAntivirus != nil {
		v := *o.GuiAntivirus

		if err = d.Set("gui_antivirus", v); err != nil {
			return diag.Errorf("error reading gui_antivirus: %v", err)
		}
	}

	if o.GuiApProfile != nil {
		v := *o.GuiApProfile

		if err = d.Set("gui_ap_profile", v); err != nil {
			return diag.Errorf("error reading gui_ap_profile: %v", err)
		}
	}

	if o.GuiApplicationControl != nil {
		v := *o.GuiApplicationControl

		if err = d.Set("gui_application_control", v); err != nil {
			return diag.Errorf("error reading gui_application_control: %v", err)
		}
	}

	if o.GuiDefaultPolicyColumns != nil {
		if err = d.Set("gui_default_policy_columns", flattenSystemSettingsGuiDefaultPolicyColumns(o.GuiDefaultPolicyColumns, sort)); err != nil {
			return diag.Errorf("error reading gui_default_policy_columns: %v", err)
		}
	}

	if o.GuiDhcpAdvanced != nil {
		v := *o.GuiDhcpAdvanced

		if err = d.Set("gui_dhcp_advanced", v); err != nil {
			return diag.Errorf("error reading gui_dhcp_advanced: %v", err)
		}
	}

	if o.GuiDnsDatabase != nil {
		v := *o.GuiDnsDatabase

		if err = d.Set("gui_dns_database", v); err != nil {
			return diag.Errorf("error reading gui_dns_database: %v", err)
		}
	}

	if o.GuiDnsfilter != nil {
		v := *o.GuiDnsfilter

		if err = d.Set("gui_dnsfilter", v); err != nil {
			return diag.Errorf("error reading gui_dnsfilter: %v", err)
		}
	}

	if o.GuiDomainIpReputation != nil {
		v := *o.GuiDomainIpReputation

		if err = d.Set("gui_domain_ip_reputation", v); err != nil {
			return diag.Errorf("error reading gui_domain_ip_reputation: %v", err)
		}
	}

	if o.GuiDosPolicy != nil {
		v := *o.GuiDosPolicy

		if err = d.Set("gui_dos_policy", v); err != nil {
			return diag.Errorf("error reading gui_dos_policy: %v", err)
		}
	}

	if o.GuiDynamicProfileDisplay != nil {
		v := *o.GuiDynamicProfileDisplay

		if err = d.Set("gui_dynamic_profile_display", v); err != nil {
			return diag.Errorf("error reading gui_dynamic_profile_display: %v", err)
		}
	}

	if o.GuiDynamicRouting != nil {
		v := *o.GuiDynamicRouting

		if err = d.Set("gui_dynamic_routing", v); err != nil {
			return diag.Errorf("error reading gui_dynamic_routing: %v", err)
		}
	}

	if o.GuiEmailCollection != nil {
		v := *o.GuiEmailCollection

		if err = d.Set("gui_email_collection", v); err != nil {
			return diag.Errorf("error reading gui_email_collection: %v", err)
		}
	}

	if o.GuiEndpointControl != nil {
		v := *o.GuiEndpointControl

		if err = d.Set("gui_endpoint_control", v); err != nil {
			return diag.Errorf("error reading gui_endpoint_control: %v", err)
		}
	}

	if o.GuiEndpointControlAdvanced != nil {
		v := *o.GuiEndpointControlAdvanced

		if err = d.Set("gui_endpoint_control_advanced", v); err != nil {
			return diag.Errorf("error reading gui_endpoint_control_advanced: %v", err)
		}
	}

	if o.GuiExplicitProxy != nil {
		v := *o.GuiExplicitProxy

		if err = d.Set("gui_explicit_proxy", v); err != nil {
			return diag.Errorf("error reading gui_explicit_proxy: %v", err)
		}
	}

	if o.GuiFileFilter != nil {
		v := *o.GuiFileFilter

		if err = d.Set("gui_file_filter", v); err != nil {
			return diag.Errorf("error reading gui_file_filter: %v", err)
		}
	}

	if o.GuiFortiapSplitTunneling != nil {
		v := *o.GuiFortiapSplitTunneling

		if err = d.Set("gui_fortiap_split_tunneling", v); err != nil {
			return diag.Errorf("error reading gui_fortiap_split_tunneling: %v", err)
		}
	}

	if o.GuiFortiextenderController != nil {
		v := *o.GuiFortiextenderController

		if err = d.Set("gui_fortiextender_controller", v); err != nil {
			return diag.Errorf("error reading gui_fortiextender_controller: %v", err)
		}
	}

	if o.GuiIcap != nil {
		v := *o.GuiIcap

		if err = d.Set("gui_icap", v); err != nil {
			return diag.Errorf("error reading gui_icap: %v", err)
		}
	}

	if o.GuiImplicitPolicy != nil {
		v := *o.GuiImplicitPolicy

		if err = d.Set("gui_implicit_policy", v); err != nil {
			return diag.Errorf("error reading gui_implicit_policy: %v", err)
		}
	}

	if o.GuiIps != nil {
		v := *o.GuiIps

		if err = d.Set("gui_ips", v); err != nil {
			return diag.Errorf("error reading gui_ips: %v", err)
		}
	}

	if o.GuiLoadBalance != nil {
		v := *o.GuiLoadBalance

		if err = d.Set("gui_load_balance", v); err != nil {
			return diag.Errorf("error reading gui_load_balance: %v", err)
		}
	}

	if o.GuiLocalInPolicy != nil {
		v := *o.GuiLocalInPolicy

		if err = d.Set("gui_local_in_policy", v); err != nil {
			return diag.Errorf("error reading gui_local_in_policy: %v", err)
		}
	}

	if o.GuiLocalReports != nil {
		v := *o.GuiLocalReports

		if err = d.Set("gui_local_reports", v); err != nil {
			return diag.Errorf("error reading gui_local_reports: %v", err)
		}
	}

	if o.GuiMulticastPolicy != nil {
		v := *o.GuiMulticastPolicy

		if err = d.Set("gui_multicast_policy", v); err != nil {
			return diag.Errorf("error reading gui_multicast_policy: %v", err)
		}
	}

	if o.GuiMultipleInterfacePolicy != nil {
		v := *o.GuiMultipleInterfacePolicy

		if err = d.Set("gui_multiple_interface_policy", v); err != nil {
			return diag.Errorf("error reading gui_multiple_interface_policy: %v", err)
		}
	}

	if o.GuiMultipleUtmProfiles != nil {
		v := *o.GuiMultipleUtmProfiles

		if err = d.Set("gui_multiple_utm_profiles", v); err != nil {
			return diag.Errorf("error reading gui_multiple_utm_profiles: %v", err)
		}
	}

	if o.GuiNat4664 != nil {
		v := *o.GuiNat4664

		if err = d.Set("gui_nat46_64", v); err != nil {
			return diag.Errorf("error reading gui_nat46_64: %v", err)
		}
	}

	if o.GuiObjectColors != nil {
		v := *o.GuiObjectColors

		if err = d.Set("gui_object_colors", v); err != nil {
			return diag.Errorf("error reading gui_object_colors: %v", err)
		}
	}

	if o.GuiPerPolicyDisclaimer != nil {
		v := *o.GuiPerPolicyDisclaimer

		if err = d.Set("gui_per_policy_disclaimer", v); err != nil {
			return diag.Errorf("error reading gui_per_policy_disclaimer: %v", err)
		}
	}

	if o.GuiPolicyBasedIpsec != nil {
		v := *o.GuiPolicyBasedIpsec

		if err = d.Set("gui_policy_based_ipsec", v); err != nil {
			return diag.Errorf("error reading gui_policy_based_ipsec: %v", err)
		}
	}

	if o.GuiPolicyDisclaimer != nil {
		v := *o.GuiPolicyDisclaimer

		if err = d.Set("gui_policy_disclaimer", v); err != nil {
			return diag.Errorf("error reading gui_policy_disclaimer: %v", err)
		}
	}

	if o.GuiReplacementMessageGroups != nil {
		v := *o.GuiReplacementMessageGroups

		if err = d.Set("gui_replacement_message_groups", v); err != nil {
			return diag.Errorf("error reading gui_replacement_message_groups: %v", err)
		}
	}

	if o.GuiSecurityProfileGroup != nil {
		v := *o.GuiSecurityProfileGroup

		if err = d.Set("gui_security_profile_group", v); err != nil {
			return diag.Errorf("error reading gui_security_profile_group: %v", err)
		}
	}

	if o.GuiSpamfilter != nil {
		v := *o.GuiSpamfilter

		if err = d.Set("gui_spamfilter", v); err != nil {
			return diag.Errorf("error reading gui_spamfilter: %v", err)
		}
	}

	if o.GuiSslvpnPersonalBookmarks != nil {
		v := *o.GuiSslvpnPersonalBookmarks

		if err = d.Set("gui_sslvpn_personal_bookmarks", v); err != nil {
			return diag.Errorf("error reading gui_sslvpn_personal_bookmarks: %v", err)
		}
	}

	if o.GuiSslvpnRealms != nil {
		v := *o.GuiSslvpnRealms

		if err = d.Set("gui_sslvpn_realms", v); err != nil {
			return diag.Errorf("error reading gui_sslvpn_realms: %v", err)
		}
	}

	if o.GuiSwitchController != nil {
		v := *o.GuiSwitchController

		if err = d.Set("gui_switch_controller", v); err != nil {
			return diag.Errorf("error reading gui_switch_controller: %v", err)
		}
	}

	if o.GuiThreatWeight != nil {
		v := *o.GuiThreatWeight

		if err = d.Set("gui_threat_weight", v); err != nil {
			return diag.Errorf("error reading gui_threat_weight: %v", err)
		}
	}

	if o.GuiTrafficShaping != nil {
		v := *o.GuiTrafficShaping

		if err = d.Set("gui_traffic_shaping", v); err != nil {
			return diag.Errorf("error reading gui_traffic_shaping: %v", err)
		}
	}

	if o.GuiVideofilter != nil {
		v := *o.GuiVideofilter

		if err = d.Set("gui_videofilter", v); err != nil {
			return diag.Errorf("error reading gui_videofilter: %v", err)
		}
	}

	if o.GuiVoipProfile != nil {
		v := *o.GuiVoipProfile

		if err = d.Set("gui_voip_profile", v); err != nil {
			return diag.Errorf("error reading gui_voip_profile: %v", err)
		}
	}

	if o.GuiVpn != nil {
		v := *o.GuiVpn

		if err = d.Set("gui_vpn", v); err != nil {
			return diag.Errorf("error reading gui_vpn: %v", err)
		}
	}

	if o.GuiWafProfile != nil {
		v := *o.GuiWafProfile

		if err = d.Set("gui_waf_profile", v); err != nil {
			return diag.Errorf("error reading gui_waf_profile: %v", err)
		}
	}

	if o.GuiWanLoadBalancing != nil {
		v := *o.GuiWanLoadBalancing

		if err = d.Set("gui_wan_load_balancing", v); err != nil {
			return diag.Errorf("error reading gui_wan_load_balancing: %v", err)
		}
	}

	if o.GuiWanoptCache != nil {
		v := *o.GuiWanoptCache

		if err = d.Set("gui_wanopt_cache", v); err != nil {
			return diag.Errorf("error reading gui_wanopt_cache: %v", err)
		}
	}

	if o.GuiWebfilter != nil {
		v := *o.GuiWebfilter

		if err = d.Set("gui_webfilter", v); err != nil {
			return diag.Errorf("error reading gui_webfilter: %v", err)
		}
	}

	if o.GuiWebfilterAdvanced != nil {
		v := *o.GuiWebfilterAdvanced

		if err = d.Set("gui_webfilter_advanced", v); err != nil {
			return diag.Errorf("error reading gui_webfilter_advanced: %v", err)
		}
	}

	if o.GuiWirelessController != nil {
		v := *o.GuiWirelessController

		if err = d.Set("gui_wireless_controller", v); err != nil {
			return diag.Errorf("error reading gui_wireless_controller: %v", err)
		}
	}

	if o.GuiZtna != nil {
		v := *o.GuiZtna

		if err = d.Set("gui_ztna", v); err != nil {
			return diag.Errorf("error reading gui_ztna: %v", err)
		}
	}

	if o.HttpExternalDest != nil {
		v := *o.HttpExternalDest

		if err = d.Set("http_external_dest", v); err != nil {
			return diag.Errorf("error reading http_external_dest: %v", err)
		}
	}

	if o.IkeDnFormat != nil {
		v := *o.IkeDnFormat

		if err = d.Set("ike_dn_format", v); err != nil {
			return diag.Errorf("error reading ike_dn_format: %v", err)
		}
	}

	if o.IkePolicyRoute != nil {
		v := *o.IkePolicyRoute

		if err = d.Set("ike_policy_route", v); err != nil {
			return diag.Errorf("error reading ike_policy_route: %v", err)
		}
	}

	if o.IkePort != nil {
		v := *o.IkePort

		if err = d.Set("ike_port", v); err != nil {
			return diag.Errorf("error reading ike_port: %v", err)
		}
	}

	if o.IkeQuickCrashDetect != nil {
		v := *o.IkeQuickCrashDetect

		if err = d.Set("ike_quick_crash_detect", v); err != nil {
			return diag.Errorf("error reading ike_quick_crash_detect: %v", err)
		}
	}

	if o.IkeSessionResume != nil {
		v := *o.IkeSessionResume

		if err = d.Set("ike_session_resume", v); err != nil {
			return diag.Errorf("error reading ike_session_resume: %v", err)
		}
	}

	if o.ImplicitAllowDns != nil {
		v := *o.ImplicitAllowDns

		if err = d.Set("implicit_allow_dns", v); err != nil {
			return diag.Errorf("error reading implicit_allow_dns: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip
		if current, ok := d.GetOk("ip"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Ip6 != nil {
		v := *o.Ip6

		if err = d.Set("ip6", v); err != nil {
			return diag.Errorf("error reading ip6: %v", err)
		}
	}

	if o.LinkDownAccess != nil {
		v := *o.LinkDownAccess

		if err = d.Set("link_down_access", v); err != nil {
			return diag.Errorf("error reading link_down_access: %v", err)
		}
	}

	if o.LldpReception != nil {
		v := *o.LldpReception

		if err = d.Set("lldp_reception", v); err != nil {
			return diag.Errorf("error reading lldp_reception: %v", err)
		}
	}

	if o.LldpTransmission != nil {
		v := *o.LldpTransmission

		if err = d.Set("lldp_transmission", v); err != nil {
			return diag.Errorf("error reading lldp_transmission: %v", err)
		}
	}

	if o.LocationId != nil {
		v := *o.LocationId

		if err = d.Set("location_id", v); err != nil {
			return diag.Errorf("error reading location_id: %v", err)
		}
	}

	if o.MacTtl != nil {
		v := *o.MacTtl

		if err = d.Set("mac_ttl", v); err != nil {
			return diag.Errorf("error reading mac_ttl: %v", err)
		}
	}

	if o.Manageip != nil {
		v := *o.Manageip

		if err = d.Set("manageip", v); err != nil {
			return diag.Errorf("error reading manageip: %v", err)
		}
	}

	if o.Manageip6 != nil {
		v := *o.Manageip6

		if err = d.Set("manageip6", v); err != nil {
			return diag.Errorf("error reading manageip6: %v", err)
		}
	}

	if o.MulticastForward != nil {
		v := *o.MulticastForward

		if err = d.Set("multicast_forward", v); err != nil {
			return diag.Errorf("error reading multicast_forward: %v", err)
		}
	}

	if o.MulticastSkipPolicy != nil {
		v := *o.MulticastSkipPolicy

		if err = d.Set("multicast_skip_policy", v); err != nil {
			return diag.Errorf("error reading multicast_skip_policy: %v", err)
		}
	}

	if o.MulticastTtlNotchange != nil {
		v := *o.MulticastTtlNotchange

		if err = d.Set("multicast_ttl_notchange", v); err != nil {
			return diag.Errorf("error reading multicast_ttl_notchange: %v", err)
		}
	}

	if o.NgfwMode != nil {
		v := *o.NgfwMode

		if err = d.Set("ngfw_mode", v); err != nil {
			return diag.Errorf("error reading ngfw_mode: %v", err)
		}
	}

	if o.Opmode != nil {
		v := *o.Opmode

		if err = d.Set("opmode", v); err != nil {
			return diag.Errorf("error reading opmode: %v", err)
		}
	}

	if o.PrpTrailerAction != nil {
		v := *o.PrpTrailerAction

		if err = d.Set("prp_trailer_action", v); err != nil {
			return diag.Errorf("error reading prp_trailer_action: %v", err)
		}
	}

	if o.SccpPort != nil {
		v := *o.SccpPort

		if err = d.Set("sccp_port", v); err != nil {
			return diag.Errorf("error reading sccp_port: %v", err)
		}
	}

	if o.SctpSessionWithoutInit != nil {
		v := *o.SctpSessionWithoutInit

		if err = d.Set("sctp_session_without_init", v); err != nil {
			return diag.Errorf("error reading sctp_session_without_init: %v", err)
		}
	}

	if o.SesDeniedTraffic != nil {
		v := *o.SesDeniedTraffic

		if err = d.Set("ses_denied_traffic", v); err != nil {
			return diag.Errorf("error reading ses_denied_traffic: %v", err)
		}
	}

	if o.SipExpectation != nil {
		v := *o.SipExpectation

		if err = d.Set("sip_expectation", v); err != nil {
			return diag.Errorf("error reading sip_expectation: %v", err)
		}
	}

	if o.SipNatTrace != nil {
		v := *o.SipNatTrace

		if err = d.Set("sip_nat_trace", v); err != nil {
			return diag.Errorf("error reading sip_nat_trace: %v", err)
		}
	}

	if o.SipSslPort != nil {
		v := *o.SipSslPort

		if err = d.Set("sip_ssl_port", v); err != nil {
			return diag.Errorf("error reading sip_ssl_port: %v", err)
		}
	}

	if o.SipTcpPort != nil {
		v := *o.SipTcpPort

		if err = d.Set("sip_tcp_port", v); err != nil {
			return diag.Errorf("error reading sip_tcp_port: %v", err)
		}
	}

	if o.SipUdpPort != nil {
		v := *o.SipUdpPort

		if err = d.Set("sip_udp_port", v); err != nil {
			return diag.Errorf("error reading sip_udp_port: %v", err)
		}
	}

	if o.SnatHairpinTraffic != nil {
		v := *o.SnatHairpinTraffic

		if err = d.Set("snat_hairpin_traffic", v); err != nil {
			return diag.Errorf("error reading snat_hairpin_traffic: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.StrictSrcCheck != nil {
		v := *o.StrictSrcCheck

		if err = d.Set("strict_src_check", v); err != nil {
			return diag.Errorf("error reading strict_src_check: %v", err)
		}
	}

	if o.TcpSessionWithoutSyn != nil {
		v := *o.TcpSessionWithoutSyn

		if err = d.Set("tcp_session_without_syn", v); err != nil {
			return diag.Errorf("error reading tcp_session_without_syn: %v", err)
		}
	}

	if o.Utf8SpamTagging != nil {
		v := *o.Utf8SpamTagging

		if err = d.Set("utf8_spam_tagging", v); err != nil {
			return diag.Errorf("error reading utf8_spam_tagging: %v", err)
		}
	}

	if o.V4EcmpMode != nil {
		v := *o.V4EcmpMode

		if err = d.Set("v4_ecmp_mode", v); err != nil {
			return diag.Errorf("error reading v4_ecmp_mode: %v", err)
		}
	}

	if o.VpnStatsLog != nil {
		v := *o.VpnStatsLog

		if err = d.Set("vpn_stats_log", v); err != nil {
			return diag.Errorf("error reading vpn_stats_log: %v", err)
		}
	}

	if o.VpnStatsPeriod != nil {
		v := *o.VpnStatsPeriod

		if err = d.Set("vpn_stats_period", v); err != nil {
			return diag.Errorf("error reading vpn_stats_period: %v", err)
		}
	}

	if o.WccpCacheEngine != nil {
		v := *o.WccpCacheEngine

		if err = d.Set("wccp_cache_engine", v); err != nil {
			return diag.Errorf("error reading wccp_cache_engine: %v", err)
		}
	}

	return nil
}

func expandSystemSettingsGuiDefaultPolicyColumns(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSettingsGuiDefaultPolicyColumns, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSettingsGuiDefaultPolicyColumns

	for i := range l {
		tmp := models.SystemSettingsGuiDefaultPolicyColumns{}
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

func getObjectSystemSettings(d *schema.ResourceData, sv string) (*models.SystemSettings, diag.Diagnostics) {
	obj := models.SystemSettings{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allow_linkdown_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_linkdown_path", sv)
				diags = append(diags, e)
			}
			obj.AllowLinkdownPath = &v2
		}
	}
	if v1, ok := d.GetOk("allow_subnet_overlap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_subnet_overlap", sv)
				diags = append(diags, e)
			}
			obj.AllowSubnetOverlap = &v2
		}
	}
	if v1, ok := d.GetOk("application_bandwidth_tracking"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("application_bandwidth_tracking", sv)
				diags = append(diags, e)
			}
			obj.ApplicationBandwidthTracking = &v2
		}
	}
	if v1, ok := d.GetOk("asymroute"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("asymroute", sv)
				diags = append(diags, e)
			}
			obj.Asymroute = &v2
		}
	}
	if v1, ok := d.GetOk("asymroute_icmp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("asymroute_icmp", sv)
				diags = append(diags, e)
			}
			obj.AsymrouteIcmp = &v2
		}
	}
	if v1, ok := d.GetOk("asymroute6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("asymroute6", sv)
				diags = append(diags, e)
			}
			obj.Asymroute6 = &v2
		}
	}
	if v1, ok := d.GetOk("asymroute6_icmp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("asymroute6_icmp", sv)
				diags = append(diags, e)
			}
			obj.Asymroute6Icmp = &v2
		}
	}
	if v1, ok := d.GetOk("auxiliary_session"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auxiliary_session", sv)
				diags = append(diags, e)
			}
			obj.AuxiliarySession = &v2
		}
	}
	if v1, ok := d.GetOk("bfd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd", sv)
				diags = append(diags, e)
			}
			obj.Bfd = &v2
		}
	}
	if v1, ok := d.GetOk("bfd_desired_min_tx"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_desired_min_tx", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BfdDesiredMinTx = &tmp
		}
	}
	if v1, ok := d.GetOk("bfd_detect_mult"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_detect_mult", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BfdDetectMult = &tmp
		}
	}
	if v1, ok := d.GetOk("bfd_dont_enforce_src_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_dont_enforce_src_port", sv)
				diags = append(diags, e)
			}
			obj.BfdDontEnforceSrcPort = &v2
		}
	}
	if v1, ok := d.GetOk("bfd_required_min_rx"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd_required_min_rx", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.BfdRequiredMinRx = &tmp
		}
	}
	if v1, ok := d.GetOk("block_land_attack"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_land_attack", sv)
				diags = append(diags, e)
			}
			obj.BlockLandAttack = &v2
		}
	}
	if v1, ok := d.GetOk("central_nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("central_nat", sv)
				diags = append(diags, e)
			}
			obj.CentralNat = &v2
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
	if v1, ok := d.GetOk("consolidated_firewall_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("consolidated_firewall_mode", sv)
				diags = append(diags, e)
			}
			obj.ConsolidatedFirewallMode = &v2
		}
	}
	if v1, ok := d.GetOk("default_voip_alg_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("default_voip_alg_mode", sv)
				diags = append(diags, e)
			}
			obj.DefaultVoipAlgMode = &v2
		}
	}
	if v1, ok := d.GetOk("deny_tcp_with_icmp"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("deny_tcp_with_icmp", sv)
				diags = append(diags, e)
			}
			obj.DenyTcpWithIcmp = &v2
		}
	}
	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device", sv)
				diags = append(diags, e)
			}
			obj.Device = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_proxy", sv)
				diags = append(diags, e)
			}
			obj.DhcpProxy = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_proxy_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("dhcp_proxy_interface", sv)
				diags = append(diags, e)
			}
			obj.DhcpProxyInterface = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_proxy_interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("dhcp_proxy_interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.DhcpProxyInterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_server_ip", sv)
				diags = append(diags, e)
			}
			obj.DhcpServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp6_server_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp6_server_ip", sv)
				diags = append(diags, e)
			}
			obj.Dhcp6ServerIp = &v2
		}
	}
	if v1, ok := d.GetOk("discovered_device_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("discovered_device_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DiscoveredDeviceTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ecmp_max_paths"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ecmp_max_paths", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EcmpMaxPaths = &tmp
		}
	}
	if v1, ok := d.GetOk("email_portal_check_dns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email_portal_check_dns", sv)
				diags = append(diags, e)
			}
			obj.EmailPortalCheckDns = &v2
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
	if v1, ok := d.GetOk("fw_session_hairpin"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fw_session_hairpin", sv)
				diags = append(diags, e)
			}
			obj.FwSessionHairpin = &v2
		}
	}
	if v1, ok := d.GetOk("gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway", sv)
				diags = append(diags, e)
			}
			obj.Gateway = &v2
		}
	}
	if v1, ok := d.GetOk("gateway6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway6", sv)
				diags = append(diags, e)
			}
			obj.Gateway6 = &v2
		}
	}
	if v1, ok := d.GetOk("gui_advanced_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_advanced_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiAdvancedPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_allow_unnamed_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_allow_unnamed_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiAllowUnnamedPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_antivirus"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_antivirus", sv)
				diags = append(diags, e)
			}
			obj.GuiAntivirus = &v2
		}
	}
	if v1, ok := d.GetOk("gui_ap_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_ap_profile", sv)
				diags = append(diags, e)
			}
			obj.GuiApProfile = &v2
		}
	}
	if v1, ok := d.GetOk("gui_application_control"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_application_control", sv)
				diags = append(diags, e)
			}
			obj.GuiApplicationControl = &v2
		}
	}
	if v, ok := d.GetOk("gui_default_policy_columns"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("gui_default_policy_columns", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSettingsGuiDefaultPolicyColumns(d, v, "gui_default_policy_columns", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.GuiDefaultPolicyColumns = t
		}
	} else if d.HasChange("gui_default_policy_columns") {
		old, new := d.GetChange("gui_default_policy_columns")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.GuiDefaultPolicyColumns = &[]models.SystemSettingsGuiDefaultPolicyColumns{}
		}
	}
	if v1, ok := d.GetOk("gui_dhcp_advanced"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_dhcp_advanced", sv)
				diags = append(diags, e)
			}
			obj.GuiDhcpAdvanced = &v2
		}
	}
	if v1, ok := d.GetOk("gui_dns_database"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_dns_database", sv)
				diags = append(diags, e)
			}
			obj.GuiDnsDatabase = &v2
		}
	}
	if v1, ok := d.GetOk("gui_dnsfilter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_dnsfilter", sv)
				diags = append(diags, e)
			}
			obj.GuiDnsfilter = &v2
		}
	}
	if v1, ok := d.GetOk("gui_domain_ip_reputation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("gui_domain_ip_reputation", sv)
				diags = append(diags, e)
			}
			obj.GuiDomainIpReputation = &v2
		}
	}
	if v1, ok := d.GetOk("gui_dos_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_dos_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiDosPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_dynamic_profile_display"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.3") {
				e := utils.AttributeVersionWarning("gui_dynamic_profile_display", sv)
				diags = append(diags, e)
			}
			obj.GuiDynamicProfileDisplay = &v2
		}
	}
	if v1, ok := d.GetOk("gui_dynamic_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_dynamic_routing", sv)
				diags = append(diags, e)
			}
			obj.GuiDynamicRouting = &v2
		}
	}
	if v1, ok := d.GetOk("gui_email_collection"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_email_collection", sv)
				diags = append(diags, e)
			}
			obj.GuiEmailCollection = &v2
		}
	}
	if v1, ok := d.GetOk("gui_endpoint_control"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_endpoint_control", sv)
				diags = append(diags, e)
			}
			obj.GuiEndpointControl = &v2
		}
	}
	if v1, ok := d.GetOk("gui_endpoint_control_advanced"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_endpoint_control_advanced", sv)
				diags = append(diags, e)
			}
			obj.GuiEndpointControlAdvanced = &v2
		}
	}
	if v1, ok := d.GetOk("gui_explicit_proxy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_explicit_proxy", sv)
				diags = append(diags, e)
			}
			obj.GuiExplicitProxy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_file_filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("gui_file_filter", sv)
				diags = append(diags, e)
			}
			obj.GuiFileFilter = &v2
		}
	}
	if v1, ok := d.GetOk("gui_fortiap_split_tunneling"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_fortiap_split_tunneling", sv)
				diags = append(diags, e)
			}
			obj.GuiFortiapSplitTunneling = &v2
		}
	}
	if v1, ok := d.GetOk("gui_fortiextender_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_fortiextender_controller", sv)
				diags = append(diags, e)
			}
			obj.GuiFortiextenderController = &v2
		}
	}
	if v1, ok := d.GetOk("gui_icap"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_icap", sv)
				diags = append(diags, e)
			}
			obj.GuiIcap = &v2
		}
	}
	if v1, ok := d.GetOk("gui_implicit_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_implicit_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiImplicitPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_ips"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_ips", sv)
				diags = append(diags, e)
			}
			obj.GuiIps = &v2
		}
	}
	if v1, ok := d.GetOk("gui_load_balance"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_load_balance", sv)
				diags = append(diags, e)
			}
			obj.GuiLoadBalance = &v2
		}
	}
	if v1, ok := d.GetOk("gui_local_in_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_local_in_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiLocalInPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_local_reports"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_local_reports", sv)
				diags = append(diags, e)
			}
			obj.GuiLocalReports = &v2
		}
	}
	if v1, ok := d.GetOk("gui_multicast_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_multicast_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiMulticastPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_multiple_interface_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_multiple_interface_policy", sv)
				diags = append(diags, e)
			}
			obj.GuiMultipleInterfacePolicy = &v2
		}
	}
	if v1, ok := d.GetOk("gui_multiple_utm_profiles"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("gui_multiple_utm_profiles", sv)
				diags = append(diags, e)
			}
			obj.GuiMultipleUtmProfiles = &v2
		}
	}
	if v1, ok := d.GetOk("gui_nat46_64"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("gui_nat46_64", sv)
				diags = append(diags, e)
			}
			obj.GuiNat4664 = &v2
		}
	}
	if v1, ok := d.GetOk("gui_object_colors"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_object_colors", sv)
				diags = append(diags, e)
			}
			obj.GuiObjectColors = &v2
		}
	}
	if v1, ok := d.GetOk("gui_per_policy_disclaimer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("gui_per_policy_disclaimer", sv)
				diags = append(diags, e)
			}
			obj.GuiPerPolicyDisclaimer = &v2
		}
	}
	if v1, ok := d.GetOk("gui_policy_based_ipsec"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_policy_based_ipsec", sv)
				diags = append(diags, e)
			}
			obj.GuiPolicyBasedIpsec = &v2
		}
	}
	if v1, ok := d.GetOk("gui_policy_disclaimer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("gui_policy_disclaimer", sv)
				diags = append(diags, e)
			}
			obj.GuiPolicyDisclaimer = &v2
		}
	}
	if v1, ok := d.GetOk("gui_replacement_message_groups"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("gui_replacement_message_groups", sv)
				diags = append(diags, e)
			}
			obj.GuiReplacementMessageGroups = &v2
		}
	}
	if v1, ok := d.GetOk("gui_security_profile_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.3", "") {
				e := utils.AttributeVersionWarning("gui_security_profile_group", sv)
				diags = append(diags, e)
			}
			obj.GuiSecurityProfileGroup = &v2
		}
	}
	if v1, ok := d.GetOk("gui_spamfilter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_spamfilter", sv)
				diags = append(diags, e)
			}
			obj.GuiSpamfilter = &v2
		}
	}
	if v1, ok := d.GetOk("gui_sslvpn_personal_bookmarks"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_sslvpn_personal_bookmarks", sv)
				diags = append(diags, e)
			}
			obj.GuiSslvpnPersonalBookmarks = &v2
		}
	}
	if v1, ok := d.GetOk("gui_sslvpn_realms"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_sslvpn_realms", sv)
				diags = append(diags, e)
			}
			obj.GuiSslvpnRealms = &v2
		}
	}
	if v1, ok := d.GetOk("gui_switch_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_switch_controller", sv)
				diags = append(diags, e)
			}
			obj.GuiSwitchController = &v2
		}
	}
	if v1, ok := d.GetOk("gui_threat_weight"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_threat_weight", sv)
				diags = append(diags, e)
			}
			obj.GuiThreatWeight = &v2
		}
	}
	if v1, ok := d.GetOk("gui_traffic_shaping"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_traffic_shaping", sv)
				diags = append(diags, e)
			}
			obj.GuiTrafficShaping = &v2
		}
	}
	if v1, ok := d.GetOk("gui_videofilter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gui_videofilter", sv)
				diags = append(diags, e)
			}
			obj.GuiVideofilter = &v2
		}
	}
	if v1, ok := d.GetOk("gui_voip_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_voip_profile", sv)
				diags = append(diags, e)
			}
			obj.GuiVoipProfile = &v2
		}
	}
	if v1, ok := d.GetOk("gui_vpn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_vpn", sv)
				diags = append(diags, e)
			}
			obj.GuiVpn = &v2
		}
	}
	if v1, ok := d.GetOk("gui_waf_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_waf_profile", sv)
				diags = append(diags, e)
			}
			obj.GuiWafProfile = &v2
		}
	}
	if v1, ok := d.GetOk("gui_wan_load_balancing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_wan_load_balancing", sv)
				diags = append(diags, e)
			}
			obj.GuiWanLoadBalancing = &v2
		}
	}
	if v1, ok := d.GetOk("gui_wanopt_cache"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_wanopt_cache", sv)
				diags = append(diags, e)
			}
			obj.GuiWanoptCache = &v2
		}
	}
	if v1, ok := d.GetOk("gui_webfilter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_webfilter", sv)
				diags = append(diags, e)
			}
			obj.GuiWebfilter = &v2
		}
	}
	if v1, ok := d.GetOk("gui_webfilter_advanced"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_webfilter_advanced", sv)
				diags = append(diags, e)
			}
			obj.GuiWebfilterAdvanced = &v2
		}
	}
	if v1, ok := d.GetOk("gui_wireless_controller"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gui_wireless_controller", sv)
				diags = append(diags, e)
			}
			obj.GuiWirelessController = &v2
		}
	}
	if v1, ok := d.GetOk("gui_ztna"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("gui_ztna", sv)
				diags = append(diags, e)
			}
			obj.GuiZtna = &v2
		}
	}
	if v1, ok := d.GetOk("http_external_dest"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_external_dest", sv)
				diags = append(diags, e)
			}
			obj.HttpExternalDest = &v2
		}
	}
	if v1, ok := d.GetOk("ike_dn_format"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_dn_format", sv)
				diags = append(diags, e)
			}
			obj.IkeDnFormat = &v2
		}
	}
	if v1, ok := d.GetOk("ike_policy_route"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.7", "v7.0.0") {
				e := utils.AttributeVersionWarning("ike_policy_route", sv)
				diags = append(diags, e)
			}
			obj.IkePolicyRoute = &v2
		}
	}
	if v1, ok := d.GetOk("ike_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ike_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.IkePort = &tmp
		}
	}
	if v1, ok := d.GetOk("ike_quick_crash_detect"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_quick_crash_detect", sv)
				diags = append(diags, e)
			}
			obj.IkeQuickCrashDetect = &v2
		}
	}
	if v1, ok := d.GetOk("ike_session_resume"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_session_resume", sv)
				diags = append(diags, e)
			}
			obj.IkeSessionResume = &v2
		}
	}
	if v1, ok := d.GetOk("implicit_allow_dns"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("implicit_allow_dns", sv)
				diags = append(diags, e)
			}
			obj.ImplicitAllowDns = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6", sv)
				diags = append(diags, e)
			}
			obj.Ip6 = &v2
		}
	}
	if v1, ok := d.GetOk("link_down_access"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("link_down_access", sv)
				diags = append(diags, e)
			}
			obj.LinkDownAccess = &v2
		}
	}
	if v1, ok := d.GetOk("lldp_reception"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_reception", sv)
				diags = append(diags, e)
			}
			obj.LldpReception = &v2
		}
	}
	if v1, ok := d.GetOk("lldp_transmission"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lldp_transmission", sv)
				diags = append(diags, e)
			}
			obj.LldpTransmission = &v2
		}
	}
	if v1, ok := d.GetOk("location_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("location_id", sv)
				diags = append(diags, e)
			}
			obj.LocationId = &v2
		}
	}
	if v1, ok := d.GetOk("mac_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mac_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MacTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("manageip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("manageip", sv)
				diags = append(diags, e)
			}
			obj.Manageip = &v2
		}
	}
	if v1, ok := d.GetOk("manageip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("manageip6", sv)
				diags = append(diags, e)
			}
			obj.Manageip6 = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_forward"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_forward", sv)
				diags = append(diags, e)
			}
			obj.MulticastForward = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_skip_policy"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_skip_policy", sv)
				diags = append(diags, e)
			}
			obj.MulticastSkipPolicy = &v2
		}
	}
	if v1, ok := d.GetOk("multicast_ttl_notchange"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multicast_ttl_notchange", sv)
				diags = append(diags, e)
			}
			obj.MulticastTtlNotchange = &v2
		}
	}
	if v1, ok := d.GetOk("ngfw_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ngfw_mode", sv)
				diags = append(diags, e)
			}
			obj.NgfwMode = &v2
		}
	}
	if v1, ok := d.GetOk("opmode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("opmode", sv)
				diags = append(diags, e)
			}
			obj.Opmode = &v2
		}
	}
	if v1, ok := d.GetOk("prp_trailer_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prp_trailer_action", sv)
				diags = append(diags, e)
			}
			obj.PrpTrailerAction = &v2
		}
	}
	if v1, ok := d.GetOk("sccp_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sccp_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SccpPort = &tmp
		}
	}
	if v1, ok := d.GetOk("sctp_session_without_init"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sctp_session_without_init", sv)
				diags = append(diags, e)
			}
			obj.SctpSessionWithoutInit = &v2
		}
	}
	if v1, ok := d.GetOk("ses_denied_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ses_denied_traffic", sv)
				diags = append(diags, e)
			}
			obj.SesDeniedTraffic = &v2
		}
	}
	if v1, ok := d.GetOk("sip_expectation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sip_expectation", sv)
				diags = append(diags, e)
			}
			obj.SipExpectation = &v2
		}
	}
	if v1, ok := d.GetOk("sip_nat_trace"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sip_nat_trace", sv)
				diags = append(diags, e)
			}
			obj.SipNatTrace = &v2
		}
	}
	if v1, ok := d.GetOk("sip_ssl_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sip_ssl_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SipSslPort = &tmp
		}
	}
	if v1, ok := d.GetOk("sip_tcp_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sip_tcp_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SipTcpPort = &tmp
		}
	}
	if v1, ok := d.GetOk("sip_udp_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sip_udp_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SipUdpPort = &tmp
		}
	}
	if v1, ok := d.GetOk("snat_hairpin_traffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("snat_hairpin_traffic", sv)
				diags = append(diags, e)
			}
			obj.SnatHairpinTraffic = &v2
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
	if v1, ok := d.GetOk("strict_src_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strict_src_check", sv)
				diags = append(diags, e)
			}
			obj.StrictSrcCheck = &v2
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
	if v1, ok := d.GetOk("utf8_spam_tagging"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("utf8_spam_tagging", sv)
				diags = append(diags, e)
			}
			obj.Utf8SpamTagging = &v2
		}
	}
	if v1, ok := d.GetOk("v4_ecmp_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("v4_ecmp_mode", sv)
				diags = append(diags, e)
			}
			obj.V4EcmpMode = &v2
		}
	}
	if v1, ok := d.GetOk("vpn_stats_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vpn_stats_log", sv)
				diags = append(diags, e)
			}
			obj.VpnStatsLog = &v2
		}
	}
	if v1, ok := d.GetOk("vpn_stats_period"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vpn_stats_period", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.VpnStatsPeriod = &tmp
		}
	}
	if v1, ok := d.GetOk("wccp_cache_engine"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wccp_cache_engine", sv)
				diags = append(diags, e)
			}
			obj.WccpCacheEngine = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemSettings(d *schema.ResourceData, sv string) (*models.SystemSettings, diag.Diagnostics) {
	obj := models.SystemSettings{}
	diags := diag.Diagnostics{}

	obj.GuiDefaultPolicyColumns = &[]models.SystemSettingsGuiDefaultPolicyColumns{}

	return &obj, diags
}
