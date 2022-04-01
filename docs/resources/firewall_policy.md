---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy"
description: |-
  Configure IPv4/IPv6 policies.
---

## fortios_firewall_policy
Configure IPv4/IPv6 policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Policy action (accept/deny/ipsec). Valid values: `accept` `deny` `ipsec` .
* `anti_replay` - Enable/disable anti-replay check. Valid values: `enable` `disable` .
* `application_list` - Name of an existing Application list. This attribute must reference one of the following datasources: `application.list.name` .
* `auth_cert` - HTTPS server certificate for policy authentication. This attribute must reference one of the following datasources: `vpn.certificate.local.name` .
* `auth_path` - Enable/disable authentication-based routing. Valid values: `enable` `disable` .
* `auth_redirect_addr` - HTTP-to-HTTPS redirect address for firewall authentication.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `enable` `disable` .
* `av_profile` - Name of an existing Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `block_notification` - Enable/disable block notification. Valid values: `enable` `disable` .
* `captive_portal_exempt` - Enable to exempt some users from the captive portal. Valid values: `enable` `disable` .
* `capture_packet` - Enable/disable capture packets. Valid values: `enable` `disable` .
* `cifs_profile` - Name of an existing CIFS profile. This attribute must reference one of the following datasources: `cifs.profile.name` .
* `comments` - Comment.
* `decrypted_traffic_mirror` - Decrypted traffic mirror. This attribute must reference one of the following datasources: `firewall.decrypted-traffic-mirror.name` .
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake. Valid values: `enable` `disable` .
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `enable` `disable` .
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `enable` `disable` .
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `disclaimer` - Enable/disable user authentication disclaimer. Valid values: `enable` `disable` .
* `dlp_profile` - Name of an existing DLP profile. This attribute must reference one of the following datasources: `dlp.profile.name` .
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dnsfilter_profile` - Name of an existing DNS filter profile. This attribute must reference one of the following datasources: `dnsfilter.profile.name` .
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `enable` `disable` .
* `dstaddr_negate` - When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be. Valid values: `enable` `disable` .
* `dynamic_shaping` - Enable/disable dynamic RADIUS defined traffic shaping. Valid values: `enable` `disable` .
* `email_collect` - Enable/disable email collection. Valid values: `enable` `disable` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `fec` - Enable/disable Forward Error Correction on traffic matching this policy on a FEC device. Valid values: `enable` `disable` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all` `check-new` .
* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `enable` `disable` .
* `fsso` - Enable/disable Fortinet Single Sign-On. Valid values: `enable` `disable` .
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication. This attribute must reference one of the following datasources: `user.fsso.name` .
* `geoip_anycast` - Enable/disable recognition of anycast IP addresses using the geography IP database. Valid values: `enable` `disable` .
* `geoip_match` - Match geography address based either on its physical location or registered location. Valid values: `physical-location` `registered-location` .
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `enable` `disable` .
* `icap_profile` - Name of an existing ICAP profile. This attribute must reference one of the following datasources: `icap.profile.name` .
* `identity_based_route` - Name of identity-based routing rule. This attribute must reference one of the following datasources: `firewall.identity-based-route.name` .
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `enable` `disable` .
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy` `flow` .
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `enable` `disable` .
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `enable` `disable` .
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `enable` `disable` .
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `enable` `disable` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `all` `utm` `disable` .
* `logtraffic_start` - Record logs when a session starts. Valid values: `enable` `disable` .
* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP. Valid values: `enable` `disable` .
* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP. Valid values: `enable` `disable` .
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `enable` `disable` .
* `nat46` - Enable/disable NAT46. Valid values: `enable` `disable` .
* `nat64` - Enable/disable NAT64. Valid values: `enable` `disable` .
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic. Valid values: `enable` `disable` .
* `natip` - Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic. Valid values: `enable` `disable` .
* `np_acceleration` - Enable/disable UTM Network Processor acceleration. Valid values: `enable` `disable` .
* `ntlm` - Enable/disable NTLM authentication. Valid values: `enable` `disable` .
* `ntlm_guest` - Enable/disable NTLM guest user access. Valid values: `enable` `disable` .
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `enable` `disable` .
* `passive_wan_health_measurement` - Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled. Valid values: `enable` `disable` .
* `per_ip_shaper` - Per-IP traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.per-ip-shaper.name` .
* `permit_any_host` - Accept UDP packets from any host. Valid values: `enable` `disable` .
* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host. Valid values: `enable` `disable` .
* `policy_expiry` - Enable/disable policy expiry. Valid values: `enable` `disable` .
* `policy_expiry_date` - Policy expiry date (YYYY-MM-DD HH:MM:SS).
* `policyid` - Policy ID (0 - 4294967294).
* `profile_group` - Name of profile group. This attribute must reference one of the following datasources: `firewall.profile-group.name` .
* `profile_protocol_options` - Name of an existing Protocol options profile. This attribute must reference one of the following datasources: `firewall.profile-protocol-options.name` .
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single` `group` .
* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server. Valid values: `enable` `disable` .
* `redirect_url` - URL users are directed to after seeing and accepting the disclaimer or authenticating.
* `replacemsg_override_group` - Override the default replacement message group for this policy. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `reputation_direction` - Direction of the initial traffic for reputation to take effect. Valid values: `source` `destination` .
* `reputation_minimum` - Minimum Reputation to take action. This attribute must reference one of the following datasources: `firewall.internet-service-reputation.id` .
* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `enable` `disable` .
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT. Valid values: `disable` `enable` .
* `schedule` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity. Valid values: `enable` `disable` .
* `sctp_filter_profile` - Name of an existing SCTP filter profile. This attribute must reference one of the following datasources: `sctp-filter.profile.name` .
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable` `enable` .
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `sgt_check` - Enable/disable security group tags (SGT) check. Valid values: `enable` `disable` .
* `srcaddr_negate` - When enabled srcaddr/srcaddr6 specifies what the source address must NOT be. Valid values: `enable` `disable` .
* `ssh_filter_profile` - Name of an existing SSH filter profile. This attribute must reference one of the following datasources: `ssh-filter.profile.name` .
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `enable` `disable` .
* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring). Valid values: `enable` `disable` .
* `ssl_ssh_profile` - Name of an existing SSL SSH profile. This attribute must reference one of the following datasources: `firewall.ssl-ssh-profile.name` .
* `status` - Enable or disable this policy. Valid values: `enable` `disable` .
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag. Valid values: `all` `data-only` `disable` .
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire. Valid values: `enable` `disable` .
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `enable` `disable` .
* `traffic_shaper` - Traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `traffic_shaper_reverse` - Reverse traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile. This attribute must reference one of the following datasources: `videofilter.profile.name` .
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1. This attribute must reference one of the following datasources: `vpn.ipsec.phase1.name` `vpn.ipsec.manualkey.name` .
* `waf_profile` - Name of an existing Web application firewall profile. This attribute must reference one of the following datasources: `waf.profile.name` .
* `wanopt` - Enable/disable WAN optimization. Valid values: `enable` `disable` .
* `wanopt_detection` - WAN optimization auto-detection mode. Valid values: `active` `passive` `off` .
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect server. Valid values: `default` `transparent` `non-transparent` .
* `wanopt_peer` - WAN optimization peer. This attribute must reference one of the following datasources: `wanopt.peer.peer-host-id` .
* `wanopt_profile` - WAN optimization profile. This attribute must reference one of the following datasources: `wanopt.profile.name` .
* `wccp` - Enable/disable forwarding traffic matching this policy to a configured WCCP server. Valid values: `enable` `disable` .
* `webcache` - Enable/disable web cache. Valid values: `enable` `disable` .
* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable` `enable` .
* `webfilter_profile` - Name of an existing Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `webproxy_forward_server` - Webproxy forward server name. This attribute must reference one of the following datasources: `web-proxy.forward-server.name` `web-proxy.forward-server-group.name` .
* `webproxy_profile` - Webproxy profile name. This attribute must reference one of the following datasources: `web-proxy.profile.name` .
* `wsso` - Enable/disable WiFi Single Sign On (WSSO). Valid values: `enable` `disable` .
* `ztna_status` - Enable/disable zero trust access. Valid values: `enable` `disable` .
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - Application group names. The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group names. This attribute must reference one of the following datasources: `application.group.name` .
* `application` - Application ID list. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `custom_log_fields` - Custom fields to append to log messages for this policy. The structure of `custom_log_fields` block is documented below.

The `custom_log_fields` block contains:

* `field_id` - Custom log field. This attribute must reference one of the following datasources: `log.custom-field.id` .
* `dstaddr` - Destination IPv4 address and address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.vip.name` `firewall.vipgrp.name` `system.external-resource.name` .
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `firewall.vipgrp6.name` `firewall.vip6.name` `system.external-resource.name` .
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below.

The `fsso_groups` block contains:

* `name` - Names of FSSO groups. This attribute must reference one of the following datasources: `user.adgrp.name` .
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below.

The `internet_service_custom` block contains:

* `name` - Custom Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-custom.name` .
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-custom-group.name` .
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below.

The `internet_service_group` block contains:

* `name` - Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-group.name` .
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.

The `internet_service_id` block contains:

* `id` - Internet Service ID. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below.

The `internet_service_name` block contains:

* `name` - Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-name.name` .
* `internet_service_src_custom` - Custom Internet Service source name. The structure of `internet_service_src_custom` block is documented below.

The `internet_service_src_custom` block contains:

* `name` - Custom Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-custom.name` .
* `internet_service_src_custom_group` - Custom Internet Service source group name. The structure of `internet_service_src_custom_group` block is documented below.

The `internet_service_src_custom_group` block contains:

* `name` - Custom Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-custom-group.name` .
* `internet_service_src_group` - Internet Service source group name. The structure of `internet_service_src_group` block is documented below.

The `internet_service_src_group` block contains:

* `name` - Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-group.name` .
* `internet_service_src_id` - Internet Service source ID. The structure of `internet_service_src_id` block is documented below.

The `internet_service_src_id` block contains:

* `id` - Internet Service ID. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `internet_service_src_name` - Internet Service source name. The structure of `internet_service_src_name` block is documented below.

The `internet_service_src_name` block contains:

* `name` - Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-name.name` .
* `ntlm_enabled_browsers` - HTTP-User-Agent value of supported browsers. The structure of `ntlm_enabled_browsers` block is documented below.

The `ntlm_enabled_browsers` block contains:

* `user_agent_string` - User agent string.
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name. This attribute must reference one of the following datasources: `firewall.ippool.name` .
* `poolname6` - IPv6 pool names. The structure of `poolname6` block is documented below.

The `poolname6` block contains:

* `name` - IPv6 pool name. This attribute must reference one of the following datasources: `firewall.ippool6.name` .
* `rtp_addr` - Address names if this is an RTP NAT policy. The structure of `rtp_addr` block is documented below.

The `rtp_addr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.internet-service-custom-group.name` `firewall.addrgrp.name` .
* `service` - Service and service group names. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service and service group names. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `sgt` - Security group tags. The structure of `sgt` block is documented below.

The `sgt` block contains:

* `id` - Security group tag (1 - 65535).
* `src_vendor_mac` - Vendor MAC source ID. The structure of `src_vendor_mac` block is documented below.

The `src_vendor_mac` block contains:

* `id` - Vendor MAC ID. This attribute must reference one of the following datasources: `firewall.vendor-mac.id` .
* `srcaddr` - Source IPv4 address and address group names. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `system.external-resource.name` .
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `ssl_mirror_intf` - SSL mirror interface name. The structure of `ssl_mirror_intf` block is documented below.

The `ssl_mirror_intf` block contains:

* `name` - Mirror Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `url_category` - URL category ID list. The structure of `url_category` block is documented below.

The `url_category` block contains:

* `id` - URL category ID.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.

The `users` block contains:

* `name` - Names of individual users that can authenticate with this policy. This attribute must reference one of the following datasources: `user.local.name` .
* `ztna_ems_tag` - Source ztna-ems-tag names. The structure of `ztna_ems_tag` block is documented below.

The `ztna_ems_tag` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `ztna_geo_tag` - Source ztna-geo-tag names. The structure of `ztna_geo_tag` block is documented below.

The `ztna_geo_tag` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_policy can be imported using:
```sh
terraform import fortios_firewall_policy.labelname {{mkey}}
```
