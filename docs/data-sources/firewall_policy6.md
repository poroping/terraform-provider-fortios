---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy6"
description: |-
  Get information on a fortios Configure IPv6 policies.
---

# Data Source: fortios_firewall_policy6
Use this data source to get information on a fortios Configure IPv6 policies.


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
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dsri` - Enable DSRI to ignore HTTP server responses.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy.
* `icap_profile` - Name of an existing ICAP profile.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode.
* `ippool` - Enable to use IP Pools for source NAT.
* `ips_sensor` - Name of an existing IPS sensor.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `name` - Policy name.
* `nat` - Enable/disable source NAT.
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic.
* `np_acceleration` - Enable/disable UTM Network Processor acceleration.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `per_ip_shaper` - Per-IP traffic shaper.
* `policyid` - Policy ID (0 - 4294967294).
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO).
* `schedule` - Schedule name.
* `send_deny_packet` - Enable/disable return of deny-packet.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `session_ttl` - Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be.
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
* `traffic_shaper` - Reverse traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `utm_status` - Enable AV/web/ips protection profile.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webcache` - Enable/disable web cache.
* `webcache_https` - Enable/disable web cache for HTTPS.
* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Webproxy profile name.
* `app_category` - Application category ID list.The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - Application group names.The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group names.
* `application` - Application ID list.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `custom_log_fields` - Log field index numbers to append custom log fields to log messages for this policy.The structure of `custom_log_fields` block is documented below.

The `custom_log_fields` block contains:

* `field_id` - Custom log field.
* `dstaddr` - Destination address and address group names.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

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
* `poolname` - IP Pool names.The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name.
* `service` - Service and service group names.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Address name.
* `srcaddr` - Source address and address group names.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
* `srcintf` - Incoming (ingress) interface.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
* `ssl_mirror_intf` - SSL mirror interface name.The structure of `ssl_mirror_intf` block is documented below.

The `ssl_mirror_intf` block contains:

* `name` - Interface name.
* `url_category` - URL category ID list.The structure of `url_category` block is documented below.

The `url_category` block contains:

* `id` - URL category ID.
* `users` - Names of individual users that can authenticate with this policy.The structure of `users` block is documented below.

The `users` block contains:

* `name` - Names of individual users that can authenticate with this policy.
