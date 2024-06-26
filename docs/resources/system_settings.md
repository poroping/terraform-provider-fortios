---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_settings"
description: |-
  Configure VDOM settings.
---

## fortios_system_settings
Configure VDOM settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allow_linkdown_path` - Enable/disable link down path. Valid values: `enable` `disable` .
* `allow_subnet_overlap` - Enable/disable allowing interface subnets to use overlapping IP addresses. Valid values: `enable` `disable` .
* `application_bandwidth_tracking` - Enable/disable application bandwidth tracking. Valid values: `disable` `enable` .
* `asymroute` - Enable/disable IPv4 asymmetric routing. Valid values: `enable` `disable` .
* `asymroute_icmp` - Enable/disable ICMP asymmetric routing. Valid values: `enable` `disable` .
* `asymroute6` - Enable/disable asymmetric IPv6 routing. Valid values: `enable` `disable` .
* `asymroute6_icmp` - Enable/disable asymmetric ICMPv6 routing. Valid values: `enable` `disable` .
* `auxiliary_session` - Enable/disable auxiliary session. Valid values: `enable` `disable` .
* `bfd` - Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces. Valid values: `enable` `disable` .
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (1 - 100000 ms, default = 250).
* `bfd_detect_mult` - BFD detection multiplier (1 - 50, default = 3).
* `bfd_dont_enforce_src_port` - Enable to not enforce verifying the source port of BFD Packets. Valid values: `enable` `disable` .
* `bfd_required_min_rx` - BFD required minimal receive interval (1 - 100000 ms, default = 250).
* `block_land_attack` - Enable/disable blocking of land attacks. Valid values: `disable` `enable` .
* `central_nat` - Enable/disable central NAT. Valid values: `enable` `disable` .
* `comments` - VDOM comments.
* `consolidated_firewall_mode` - Consolidated firewall mode. Valid values: .
* `default_app_port_as_service` - Enable/disable policy service enforcement based on application default ports. Valid values: `enable` `disable` .
* `default_policy_expiry_days` - Default policy expiry in days (0 - 365 days, default = 30).
* `default_voip_alg_mode` - Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile. Valid values: `proxy-based` `kernel-helper-based` .
* `deny_tcp_with_icmp` - Enable/disable denying TCP by sending an ICMP communication prohibited packet. Valid values: `enable` `disable` .
* `detect_unknown_esp` - Enable/disable detection of unknown ESP packets (default = enable). Valid values: `enable` `disable` .
* `device` - Interface to use for management access for NAT mode. This attribute must reference one of the following datasources: `system.interface.name` .
* `dhcp_proxy` - Enable/disable the DHCP Proxy. Valid values: `enable` `disable` .
* `dhcp_proxy_interface` - Specify outgoing interface to reach server. This attribute must reference one of the following datasources: `system.interface.name` .
* `dhcp_proxy_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto` `sdwan` `specify` .
* `dhcp_server_ip` - DHCP Server IPv4 address.
* `dhcp6_server_ip` - DHCPv6 server IPv6 address.
* `discovered_device_timeout` - Timeout for discovered devices (1 - 365 days, default = 28).
* `dyn_addr_session_check` - Enable/disable dirty session check caused by dynamic address updates. Valid values: `enable` `disable` .
* `ecmp_max_paths` - Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 255, default = 255).
* `email_portal_check_dns` - Enable/disable using DNS to validate email addresses collected by a captive portal. Valid values: `disable` `enable` .
* `ext_resource_session_check` - Enable/disable dirty session check caused by external resource updates. Valid values: `enable` `disable` .
* `firewall_session_dirty` - Select how to manage sessions affected by firewall policy configuration changes. Valid values: `check-all` `check-new` `check-policy-option` .
* `fqdn_session_check` - Enable/disable dirty session check caused by FQDN updates. Valid values: `enable` `disable` .
* `fw_session_hairpin` - Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate. Valid values: `enable` `disable` .
* `gateway` - Transparent mode IPv4 default gateway IP address.
* `gateway6` - Transparent mode IPv4 default gateway IP address.
* `gui_advanced_policy` - Enable/disable advanced policy configuration on the GUI. Valid values: `enable` `disable` .
* `gui_advanced_wireless_features` - Enable/disable advanced wireless features in GUI. Valid values: `enable` `disable` .
* `gui_allow_unnamed_policy` - Enable/disable the requirement for policy naming on the GUI. Valid values: `enable` `disable` .
* `gui_antivirus` - Enable/disable AntiVirus on the GUI. Valid values: `enable` `disable` .
* `gui_ap_profile` - Enable/disable FortiAP profiles on the GUI. Valid values: `enable` `disable` .
* `gui_application_control` - Enable/disable application control on the GUI. Valid values: `enable` `disable` .
* `gui_dhcp_advanced` - Enable/disable advanced DHCP options on the GUI. Valid values: `enable` `disable` .
* `gui_dlp_profile` - Enable/disable Data Leak Prevention on the GUI. Valid values: `enable` `disable` .
* `gui_dns_database` - Enable/disable DNS database settings on the GUI. Valid values: `enable` `disable` .
* `gui_dnsfilter` - Enable/disable DNS Filtering on the GUI. Valid values: `enable` `disable` .
* `gui_domain_ip_reputation` - Enable/disable Domain and IP Reputation on the GUI. Valid values: `enable` `disable` .
* `gui_dos_policy` - Enable/disable DoS policies on the GUI. Valid values: `enable` `disable` .
* `gui_dynamic_profile_display` - Enable/disable RADIUS Single Sign On (RSSO) on the GUI. Valid values: `enable` `disable` .
* `gui_dynamic_routing` - Enable/disable dynamic routing on the GUI. Valid values: `enable` `disable` .
* `gui_email_collection` - Enable/disable email collection on the GUI. Valid values: `enable` `disable` .
* `gui_endpoint_control` - Enable/disable endpoint control on the GUI. Valid values: `enable` `disable` .
* `gui_endpoint_control_advanced` - Enable/disable advanced endpoint control options on the GUI. Valid values: `enable` `disable` .
* `gui_enforce_change_summary` - Enforce change summaries for select tables in the GUI. Valid values: `disable` `require` `optional` .
* `gui_explicit_proxy` - Enable/disable the explicit proxy on the GUI. Valid values: `enable` `disable` .
* `gui_file_filter` - Enable/disable File-filter on the GUI. Valid values: `enable` `disable` .
* `gui_fortiap_split_tunneling` - Enable/disable FortiAP split tunneling on the GUI. Valid values: `enable` `disable` .
* `gui_fortiextender_controller` - Enable/disable FortiExtender on the GUI. Valid values: `enable` `disable` .
* `gui_icap` - Enable/disable ICAP on the GUI. Valid values: `enable` `disable` .
* `gui_implicit_policy` - Enable/disable implicit firewall policies on the GUI. Valid values: `enable` `disable` .
* `gui_ips` - Enable/disable IPS on the GUI. Valid values: `enable` `disable` .
* `gui_load_balance` - Enable/disable server load balancing on the GUI. Valid values: `enable` `disable` .
* `gui_local_in_policy` - Enable/disable Local-In policies on the GUI. Valid values: `enable` `disable` .
* `gui_local_reports` - Enable/disable local reports on the GUI. Valid values: `enable` `disable` .
* `gui_multicast_policy` - Enable/disable multicast firewall policies on the GUI. Valid values: `enable` `disable` .
* `gui_multiple_interface_policy` - Enable/disable adding multiple interfaces to a policy on the GUI. Valid values: `enable` `disable` .
* `gui_multiple_utm_profiles` - Enable/disable multiple UTM profiles on the GUI. Valid values: `enable` `disable` .
* `gui_nat46_64` - Enable/disable NAT46 and NAT64 settings on the GUI. Valid values: `enable` `disable` .
* `gui_object_colors` - Enable/disable object colors on the GUI. Valid values: `enable` `disable` .
* `gui_ot` - Enable/disable Operational technology features on the GUI. Valid values: `enable` `disable` .
* `gui_per_policy_disclaimer` - Enable/disable policy disclaimer on the GUI. Valid values: `enable` `disable` .
* `gui_policy_based_ipsec` - Enable/disable policy-based IPsec VPN on the GUI. Valid values: `enable` `disable` .
* `gui_policy_disclaimer` - Enable/disable policy disclaimer on the GUI. Valid values: `enable` `disable` .
* `gui_proxy_inspection` - Enable/disable the proxy features on the GUI. Valid values: `enable` `disable` .
* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI. Valid values: `enable` `disable` .
* `gui_security_profile_group` - Enable/disable Security Profile Groups on the GUI. Valid values: `enable` `disable` .
* `gui_spamfilter` - Enable/disable Antispam on the GUI. Valid values: `enable` `disable` .
* `gui_sslvpn_personal_bookmarks` - Enable/disable SSL-VPN personal bookmark management on the GUI. Valid values: `enable` `disable` .
* `gui_sslvpn_realms` - Enable/disable SSL-VPN realms on the GUI. Valid values: `enable` `disable` .
* `gui_switch_controller` - Enable/disable the switch controller on the GUI. Valid values: `enable` `disable` .
* `gui_threat_weight` - Enable/disable threat weight on the GUI. Valid values: `enable` `disable` .
* `gui_traffic_shaping` - Enable/disable traffic shaping on the GUI. Valid values: `enable` `disable` .
* `gui_videofilter` - Enable/disable Video filtering on the GUI. Valid values: `enable` `disable` .
* `gui_voip_profile` - Enable/disable VoIP profiles on the GUI. Valid values: `enable` `disable` .
* `gui_vpn` - Enable/disable VPN tunnels on the GUI. Valid values: `enable` `disable` .
* `gui_waf_profile` - Enable/disable Web Application Firewall on the GUI. Valid values: `enable` `disable` .
* `gui_wan_load_balancing` - Enable/disable SD-WAN on the GUI. Valid values: `enable` `disable` .
* `gui_wanopt_cache` - Enable/disable WAN Optimization and Web Caching on the GUI. Valid values: `enable` `disable` .
* `gui_webfilter` - Enable/disable Web filtering on the GUI. Valid values: `enable` `disable` .
* `gui_webfilter_advanced` - Enable/disable advanced web filtering on the GUI. Valid values: `enable` `disable` .
* `gui_wireless_controller` - Enable/disable the wireless controller on the GUI. Valid values: `enable` `disable` .
* `gui_ztna` - Enable/disable Zero Trust Network Access features on the GUI. Valid values: `enable` `disable` .
* `h323_direct_model` - Enable/disable H323 direct model. Valid values: `disable` `enable` .
* `http_external_dest` - Offload HTTP traffic to FortiWeb or FortiCache. Valid values: `fortiweb` `forticache` .
* `ike_dn_format` - Configure IKE ASN.1 Distinguished Name format conventions. Valid values: `with-space` `no-space` .
* `ike_policy_route` - Enable/disable IKE Policy Based Routing (PBR). Valid values: `enable` `disable` .
* `ike_port` - UDP port for IKE/IPsec traffic (default 500).
* `ike_quick_crash_detect` - Enable/disable IKE quick crash detection (RFC 6290). Valid values: `enable` `disable` .
* `ike_session_resume` - Enable/disable IKEv2 session resumption (RFC 5723). Valid values: `enable` `disable` .
* `implicit_allow_dns` - Enable/disable implicitly allowing DNS traffic. Valid values: `enable` `disable` .
* `internet_service_database_cache` - Enable/disable Internet Service database caching. Valid values: `disable` `enable` .
* `ip` - IP address and netmask.
* `ip6` - IPv6 address prefix for NAT mode.
* `lan_extension_controller_addr` - Controller IP address or FQDN to connect.
* `link_down_access` - Enable/disable link down access traffic. Valid values: `enable` `disable` .
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM. Valid values: `enable` `disable` `global` .
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM. Valid values: `enable` `disable` `global` .
* `location_id` - Local location ID in the form of an IPv4 address.
* `mac_ttl` - Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).
* `manageip` - Transparent mode IPv4 management IP address and netmask.
* `manageip6` - Transparent mode IPv6 management IP address and netmask.
* `multicast_forward` - Enable/disable multicast forwarding. Valid values: `enable` `disable` .
* `multicast_skip_policy` - Enable/disable allowing multicast traffic through the FortiGate without a policy check. Valid values: `enable` `disable` .
* `multicast_ttl_notchange` - Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets. Valid values: `enable` `disable` .
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in NAT46. Valid values: `enable` `disable` .
* `nat46_generate_ipv6_fragment_header` - Enable/disable NAT46 IPv6 fragment header generation. Valid values: `enable` `disable` .
* `nat64_force_ipv6_packet_forwarding` - Enable/disable mandatory IPv6 packet forwarding in NAT64. Valid values: `enable` `disable` .
* `ngfw_mode` - Next Generation Firewall (NGFW) mode. Valid values: `profile-based` `policy-based` .
* `opmode` - Firewall operation mode (NAT or Transparent). Valid values: `nat` `transparent` .
* `policy_offload_level` - Configure firewall policy offload level. Valid values: `disable` `dos-offload` .
* `prp_trailer_action` - Enable/disable action to take on PRP trailer. Valid values: `enable` `disable` .
* `sccp_port` - TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).
* `sctp_session_without_init` - Enable/disable SCTP session creation without SCTP INIT. Valid values: `enable` `disable` .
* `ses_denied_traffic` - Enable/disable including denied session in the session table. Valid values: `enable` `disable` .
* `sip_expectation` - Enable/disable the SIP kernel session helper to create an expectation for port 5060. Valid values: `enable` `disable` .
* `sip_nat_trace` - Enable/disable recording the original SIP source IP address when NAT is used. Valid values: `enable` `disable` .
* `sip_ssl_port` - TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).
* `sip_tcp_port` - TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_udp_port` - UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `snat_hairpin_traffic` - Enable/disable source NAT (SNAT) for hairpin traffic. Valid values: `enable` `disable` .
* `status` - Enable/disable this VDOM. Valid values: `enable` `disable` .
* `strict_src_check` - Enable/disable strict source verification. Valid values: `enable` `disable` .
* `tcp_session_without_syn` - Enable/disable allowing TCP session without SYN flags. Valid values: `enable` `disable` .
* `utf8_spam_tagging` - Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support. Valid values: `enable` `disable` .
* `v4_ecmp_mode` - IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode. Valid values: `source-ip-based` `weight-based` `usage-based` `source-dest-ip-based` .
* `vdom_type` - Vdom type (traffic, lan-extension or admin). Valid values: `traffic` `lan-extension` `admin` .
* `vpn_stats_log` - Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space. Valid values: `ipsec` `pptp` `l2tp` `ssl` .
* `vpn_stats_period` - Period to send VPN log statistics (0 or 60 - 86400 sec).
* `wccp_cache_engine` - Enable/disable WCCP cache engine. Valid values: `enable` `disable` .
* `gui_default_policy_columns` - Default columns to display for policy lists on GUI. The structure of `gui_default_policy_columns` block is documented below.

The `gui_default_policy_columns` block contains:

* `name` - Select column name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_settings can be imported using:
```sh
terraform import fortios_system_settings.labelname {{mkey}}
```
