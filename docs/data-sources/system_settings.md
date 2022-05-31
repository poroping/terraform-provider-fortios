---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_settings"
description: |-
  Get information on a fortios Configure VDOM settings.
---

# Data Source: fortios_system_settings
Use this data source to get information on a fortios Configure VDOM settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allow_linkdown_path` - Enable/disable link down path.
* `allow_subnet_overlap` - Enable/disable allowing interface subnets to use overlapping IP addresses.
* `application_bandwidth_tracking` - Enable/disable application bandwidth tracking.
* `asymroute` - Enable/disable IPv4 asymmetric routing.
* `asymroute_icmp` - Enable/disable ICMP asymmetric routing.
* `asymroute6` - Enable/disable asymmetric IPv6 routing.
* `asymroute6_icmp` - Enable/disable asymmetric ICMPv6 routing.
* `auxiliary_session` - Enable/disable auxiliary session.
* `bfd` - Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (1 - 100000 ms, default = 250).
* `bfd_detect_mult` - BFD detection multiplier (1 - 50, default = 3).
* `bfd_dont_enforce_src_port` - Enable to not enforce verifying the source port of BFD Packets.
* `bfd_required_min_rx` - BFD required minimal receive interval (1 - 100000 ms, default = 250).
* `block_land_attack` - Enable/disable blocking of land attacks.
* `central_nat` - Enable/disable central NAT.
* `comments` - VDOM comments.
* `consolidated_firewall_mode` - Consolidated firewall mode.
* `default_app_port_as_service` - Enable/disable policy service enforcement based on application default ports.
* `default_policy_expiry_days` - Default policy expiry in days (0 - 365 days, default = 30).
* `default_voip_alg_mode` - Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile.
* `deny_tcp_with_icmp` - Enable/disable denying TCP by sending an ICMP communication prohibited packet.
* `device` - Interface to use for management access for NAT mode.
* `dhcp_proxy` - Enable/disable the DHCP Proxy.
* `dhcp_proxy_interface` - Specify outgoing interface to reach server.
* `dhcp_proxy_interface_select_method` - Specify how to select outgoing interface to reach server.
* `dhcp_server_ip` - DHCP Server IPv4 address.
* `dhcp6_server_ip` - DHCPv6 server IPv6 address.
* `discovered_device_timeout` - Timeout for discovered devices (1 - 365 days, default = 28).
* `ecmp_max_paths` - Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 255, default = 255).
* `email_portal_check_dns` - Enable/disable using DNS to validate email addresses collected by a captive portal.
* `firewall_session_dirty` - Select how to manage sessions affected by firewall policy configuration changes.
* `fw_session_hairpin` - Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate.
* `gateway` - Transparent mode IPv4 default gateway IP address.
* `gateway6` - Transparent mode IPv4 default gateway IP address.
* `gui_advanced_policy` - Enable/disable advanced policy configuration on the GUI.
* `gui_allow_unnamed_policy` - Enable/disable the requirement for policy naming on the GUI.
* `gui_antivirus` - Enable/disable AntiVirus on the GUI.
* `gui_ap_profile` - Enable/disable FortiAP profiles on the GUI.
* `gui_application_control` - Enable/disable application control on the GUI.
* `gui_dhcp_advanced` - Enable/disable advanced DHCP options on the GUI.
* `gui_dns_database` - Enable/disable DNS database settings on the GUI.
* `gui_dnsfilter` - Enable/disable DNS Filtering on the GUI.
* `gui_domain_ip_reputation` - Enable/disable Domain and IP Reputation on the GUI.
* `gui_dos_policy` - Enable/disable DoS policies on the GUI.
* `gui_dynamic_profile_display` - Enable/disable RADIUS Single Sign On (RSSO) on the GUI.
* `gui_dynamic_routing` - Enable/disable dynamic routing on the GUI.
* `gui_email_collection` - Enable/disable email collection on the GUI.
* `gui_endpoint_control` - Enable/disable endpoint control on the GUI.
* `gui_endpoint_control_advanced` - Enable/disable advanced endpoint control options on the GUI.
* `gui_enforce_change_summary` - Enforce change summaries for select tables in the GUI.
* `gui_explicit_proxy` - Enable/disable the explicit proxy on the GUI.
* `gui_file_filter` - Enable/disable File-filter on the GUI.
* `gui_fortiap_split_tunneling` - Enable/disable FortiAP split tunneling on the GUI.
* `gui_fortiextender_controller` - Enable/disable FortiExtender on the GUI.
* `gui_icap` - Enable/disable ICAP on the GUI.
* `gui_implicit_policy` - Enable/disable implicit firewall policies on the GUI.
* `gui_ips` - Enable/disable IPS on the GUI.
* `gui_load_balance` - Enable/disable server load balancing on the GUI.
* `gui_local_in_policy` - Enable/disable Local-In policies on the GUI.
* `gui_local_reports` - Enable/disable local reports on the GUI.
* `gui_multicast_policy` - Enable/disable multicast firewall policies on the GUI.
* `gui_multiple_interface_policy` - Enable/disable adding multiple interfaces to a policy on the GUI.
* `gui_multiple_utm_profiles` - Enable/disable multiple UTM profiles on the GUI.
* `gui_nat46_64` - Enable/disable NAT46 and NAT64 settings on the GUI.
* `gui_object_colors` - Enable/disable object colors on the GUI.
* `gui_ot` - Enable/disable Show Operational Technology Purdue Model.
* `gui_per_policy_disclaimer` - Enable/disable policy disclaimer on the GUI.
* `gui_policy_based_ipsec` - Enable/disable policy-based IPsec VPN on the GUI.
* `gui_policy_disclaimer` - Enable/disable policy disclaimer on the GUI.
* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI.
* `gui_security_profile_group` - Enable/disable Security Profile Groups on the GUI.
* `gui_spamfilter` - Enable/disable Antispam on the GUI.
* `gui_sslvpn_personal_bookmarks` - Enable/disable SSL-VPN personal bookmark management on the GUI.
* `gui_sslvpn_realms` - Enable/disable SSL-VPN realms on the GUI.
* `gui_switch_controller` - Enable/disable the switch controller on the GUI.
* `gui_threat_weight` - Enable/disable threat weight on the GUI.
* `gui_traffic_shaping` - Enable/disable traffic shaping on the GUI.
* `gui_videofilter` - Enable/disable Video filtering on the GUI.
* `gui_voip_profile` - Enable/disable VoIP profiles on the GUI.
* `gui_vpn` - Enable/disable VPN tunnels on the GUI.
* `gui_waf_profile` - Enable/disable Web Application Firewall on the GUI.
* `gui_wan_load_balancing` - Enable/disable SD-WAN on the GUI.
* `gui_wanopt_cache` - Enable/disable WAN Optimization and Web Caching on the GUI.
* `gui_webfilter` - Enable/disable Web filtering on the GUI.
* `gui_webfilter_advanced` - Enable/disable advanced web filtering on the GUI.
* `gui_wireless_controller` - Enable/disable the wireless controller on the GUI.
* `gui_ztna` - Enable/disable Zero Trust Network Access features on the GUI.
* `h323_direct_model` - Enable/disable H323 direct model.
* `http_external_dest` - Offload HTTP traffic to FortiWeb or FortiCache.
* `ike_dn_format` - Configure IKE ASN.1 Distinguished Name format conventions.
* `ike_policy_route` - Enable/disable IKE Policy Based Routing (PBR).
* `ike_port` - UDP port for IKE/IPsec traffic (default 500).
* `ike_quick_crash_detect` - Enable/disable IKE quick crash detection (RFC 6290).
* `ike_session_resume` - Enable/disable IKEv2 session resumption (RFC 5723).
* `implicit_allow_dns` - Enable/disable implicitly allowing DNS traffic.
* `ip` - IP address and netmask.
* `ip6` - IPv6 address prefix for NAT mode.
* `link_down_access` - Enable/disable link down access traffic.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM.
* `location_id` - Local location ID in the form of an IPv4 address.
* `mac_ttl` - Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).
* `manageip` - Transparent mode IPv4 management IP address and netmask.
* `manageip6` - Transparent mode IPv6 management IP address and netmask.
* `multicast_forward` - Enable/disable multicast forwarding.
* `multicast_skip_policy` - Enable/disable allowing multicast traffic through the FortiGate without a policy check.
* `multicast_ttl_notchange` - Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets.
* `ngfw_mode` - Next Generation Firewall (NGFW) mode.
* `opmode` - Firewall operation mode (NAT or Transparent).
* `prp_trailer_action` - Enable/disable action to take on PRP trailer.
* `sccp_port` - TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).
* `sctp_session_without_init` - Enable/disable SCTP session creation without SCTP INIT.
* `ses_denied_traffic` - Enable/disable including denied session in the session table.
* `sip_expectation` - Enable/disable the SIP kernel session helper to create an expectation for port 5060.
* `sip_nat_trace` - Enable/disable recording the original SIP source IP address when NAT is used.
* `sip_ssl_port` - TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).
* `sip_tcp_port` - TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_udp_port` - UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `snat_hairpin_traffic` - Enable/disable source NAT (SNAT) for hairpin traffic.
* `status` - Enable/disable this VDOM.
* `strict_src_check` - Enable/disable strict source verification.
* `tcp_session_without_syn` - Enable/disable allowing TCP session without SYN flags.
* `utf8_spam_tagging` - Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support.
* `v4_ecmp_mode` - IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode.
* `vdom_type` - VDOM type (traffic or admin).
* `vpn_stats_log` - Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space.
* `vpn_stats_period` - Period to send VPN log statistics (0 or 60 - 86400 sec).
* `wccp_cache_engine` - Enable/disable WCCP cache engine.
* `gui_default_policy_columns` - Default columns to display for policy lists on GUI.The structure of `gui_default_policy_columns` block is documented below.

The `gui_default_policy_columns` block contains:

* `name` - Select column name.
