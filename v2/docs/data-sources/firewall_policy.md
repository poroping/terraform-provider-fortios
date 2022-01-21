---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy"
description: |-
  Get information on a fortios Configure IPv4/IPv6 policies.
---

# Data Source: fortios_firewall_policy
Use this data source to get information on a fortios Configure IPv4/IPv6 policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID (0 - 4294967294).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Policy action (accept/deny/ipsec).
* `anti_replay` - Enable/disable anti-replay check.
* `application_list` - Name of an existing Application list.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_path` - Enable/disable authentication-based routing.
* `auth_redirect_addr` - HTTP-to-HTTPS redirect address for firewall authentication.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading.
* `av_profile` - Name of an existing Antivirus profile.
* `block_notification` - Enable/disable block notification.
* `captive_portal_exempt` - Enable to exempt some users from the captive portal.
* `capture_packet` - Enable/disable capture packets.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `disclaimer` - Enable/disable user authentication disclaimer.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dsri` - Enable DSRI to ignore HTTP server responses.
* `dstaddr_negate` - When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be.
* `dynamic_shaping` - Enable/disable dynamic RADIUS defined traffic shaping.
* `email_collect` - Enable/disable email collection.
* `emailfilter_profile` - Name of an existing email filter profile.
* `fec` - Enable/disable Forward Error Correction on traffic matching this policy on a FEC device.
* `file_filter_profile` - Name of an existing file-filter profile.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `fsso` - Enable/disable Fortinet Single Sign-On.
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `geoip_anycast` - Enable/disable recognition of anycast IP addresses using the geography IP database.
* `geoip_match` - Match geography address based either on its physical location or registered location.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy.
* `icap_profile` - Name of an existing ICAP profile.
* `identity_based_route` - Name of identity-based routing rule.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. 
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. 
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be.
* `ippool` - Enable to use IP Pools for source NAT.
* `ips_sensor` - Name of an existing IPS sensor.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP.
* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP.
* `name` - Policy name.
* `nat` - Enable/disable source NAT.
* `nat46` - Enable/disable NAT46.
* `nat64` - Enable/disable NAT64.
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic.
* `natip` - Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic.
* `np_acceleration` - Enable/disable UTM Network Processor acceleration.
* `ntlm` - Enable/disable NTLM authentication.
* `ntlm_guest` - Enable/disable NTLM guest user access.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `passive_wan_health_measurement` - Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled.
* `per_ip_shaper` - Per-IP traffic shaper.
* `permit_any_host` - Accept UDP packets from any host.
* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host.
* `policyid` - Policy ID (0 - 4294967294).
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server.
* `redirect_url` - URL users are directed to after seeing and accepting the disclaimer or authenticating.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `reputation_direction` - Direction of the initial traffic for reputation to take effect.
* `reputation_minimum` - Minimum Reputation to take action.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO).
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT.
* `schedule` - Schedule name.
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `sgt_check` - Enable/disable security group tags (SGT) check.
* `srcaddr_negate` - When enabled srcaddr/srcaddr6 specifies what the source address must NOT be.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy.
* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy.
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag.
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `waf_profile` - Name of an existing Web application firewall profile.
* `wanopt` - Enable/disable WAN optimization.
* `wanopt_detection` - WAN optimization auto-detection mode.
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect server.
* `wanopt_peer` - WAN optimization peer.
* `wanopt_profile` - WAN optimization profile.
* `wccp` - Enable/disable forwarding traffic matching this policy to a configured WCCP server.
* `webcache` - Enable/disable web cache.
* `webcache_https` - Enable/disable web cache for HTTPS.
* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Webproxy forward server name.
* `webproxy_profile` - Webproxy profile name.
* `wsso` - Enable/disable WiFi Single Sign On (WSSO).
* `ztna_status` - Enable/disable zero trust access.
* `app_category` - Application category ID list.The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - Application group names.The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group names.
* `application` - Application ID list.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `custom_log_fields` - Custom fields to append to log messages for this policy.The structure of `custom_log_fields` block is documented below.

The `custom_log_fields` block contains:

* `field_id` - Custom log field.
* `dstaddr` - Destination IPv4 address and address group names.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `dstaddr6` - Destination IPv6 address name and address group names.The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name.
* `dstintf` - Outgoing (egress) interface.The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name.
* `fsso_groups` - Names of FSSO groups.The structure of `fsso_groups` block is documented below.

The `fsso_groups` block contains:

* `name` - Names of FSSO groups.
* `groups` - Names of user groups that can authenticate with this policy.The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name.
* `internet_service_custom` - Custom Internet Service name.The structure of `internet_service_custom` block is documented below.

The `internet_service_custom` block contains:

* `name` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.The structure of `internet_service_custom_group` block is documented below.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group name.The structure of `internet_service_group` block is documented below.

The `internet_service_group` block contains:

* `name` - Internet Service group name.
* `internet_service_id` - Internet Service ID.The structure of `internet_service_id` block is documented below.

The `internet_service_id` block contains:

* `id` - Internet Service ID.
* `internet_service_name` - Internet Service name.The structure of `internet_service_name` block is documented below.

The `internet_service_name` block contains:

* `name` - Internet Service name.
* `internet_service_src_custom` - Custom Internet Service source name.The structure of `internet_service_src_custom` block is documented below.

The `internet_service_src_custom` block contains:

* `name` - Custom Internet Service name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.The structure of `internet_service_src_custom_group` block is documented below.

The `internet_service_src_custom_group` block contains:

* `name` - Custom Internet Service group name.
* `internet_service_src_group` - Internet Service source group name.The structure of `internet_service_src_group` block is documented below.

The `internet_service_src_group` block contains:

* `name` - Internet Service group name.
* `internet_service_src_id` - Internet Service source ID.The structure of `internet_service_src_id` block is documented below.

The `internet_service_src_id` block contains:

* `id` - Internet Service ID.
* `internet_service_src_name` - Internet Service source name.The structure of `internet_service_src_name` block is documented below.

The `internet_service_src_name` block contains:

* `name` - Internet Service name.
* `ntlm_enabled_browsers` - HTTP-User-Agent value of supported browsers.The structure of `ntlm_enabled_browsers` block is documented below.

The `ntlm_enabled_browsers` block contains:

* `user_agent_string` - User agent string.
* `poolname` - IP Pool names.The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name.
* `poolname6` - IPv6 pool names.The structure of `poolname6` block is documented below.

The `poolname6` block contains:

* `name` - IPv6 pool name.
* `rtp_addr` - Address names if this is an RTP NAT policy.The structure of `rtp_addr` block is documented below.

The `rtp_addr` block contains:

* `name` - Address name.
* `service` - Service and service group names.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service and service group names.
* `sgt` - Security group tags.The structure of `sgt` block is documented below.

The `sgt` block contains:

* `id` - Security group tag (1 - 65535).
* `src_vendor_mac` - Vendor MAC source ID.The structure of `src_vendor_mac` block is documented below.

The `src_vendor_mac` block contains:

* `id` - Vendor MAC ID.
* `srcaddr` - Source IPv4 address and address group names.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
* `srcaddr6` - Source IPv6 address name and address group names.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name.
* `srcintf` - Incoming (ingress) interface.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
* `ssl_mirror_intf` - SSL mirror interface name.The structure of `ssl_mirror_intf` block is documented below.

The `ssl_mirror_intf` block contains:

* `name` - Mirror Interface name.
* `url_category` - URL category ID list.The structure of `url_category` block is documented below.

The `url_category` block contains:

* `id` - URL category ID.
* `users` - Names of individual users that can authenticate with this policy.The structure of `users` block is documented below.

The `users` block contains:

* `name` - Names of individual users that can authenticate with this policy.
* `ztna_ems_tag` - Source ztna-ems-tag names.The structure of `ztna_ems_tag` block is documented below.

The `ztna_ems_tag` block contains:

* `name` - Address name.
* `ztna_geo_tag` - Source ztna-geo-tag names.The structure of `ztna_geo_tag` block is documented below.

The `ztna_geo_tag` block contains:

* `name` - Address name.
