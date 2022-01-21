---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallconsolidated_policy"
description: |-
  Get information on a fortios Configure consolidated IPv4/IPv6 policies.
---

# Data Source: fortios_firewallconsolidated_policy
Use this data source to get information on a fortios Configure consolidated IPv4/IPv6 policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID (0 - 4294967294).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Policy action (accept/deny/ipsec).
* `application_list` - Name of an existing Application list.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading.
* `av_profile` - Name of an existing Antivirus profile.
* `captive_portal_exempt` - Enable exemption of some users from the captive portal.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. 
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy.
* `icap_profile` - Name of an existing ICAP profile.
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
* `name` - Policy name.
* `nat` - Enable/disable source NAT.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `per_ip_shaper` - Per-IP traffic shaper.
* `policyid` - Policy ID (0 - 4294967294).
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `schedule` - Schedule name.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy.
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `voip_profile` - Name of an existing VoIP profile.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `waf_profile` - Name of an existing Web application firewall profile.
* `wanopt` - Enable/disable WAN optimization.
* `wanopt_detection` - WAN optimization auto-detection mode.
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect to server.
* `wanopt_peer` - WAN optimization peer.
* `wanopt_profile` - WAN optimization profile.
* `webcache` - Enable/disable web cache.
* `webcache_https` - Enable/disable web cache for HTTPS.
* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Webproxy forward server name.
* `webproxy_profile` - Webproxy profile name.
* `dstaddr4` - Destination IPv4 address name and address group names.The structure of `dstaddr4` block is documented below.

The `dstaddr4` block contains:

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
* `poolname4` - IPv4 pool names.The structure of `poolname4` block is documented below.

The `poolname4` block contains:

* `name` - IPv4 pool name.
* `poolname6` - IPv6 pool names.The structure of `poolname6` block is documented below.

The `poolname6` block contains:

* `name` - IPv6 pool name.
* `service` - Service and service group names.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr4` - Source IPv4 address name and address group names.The structure of `srcaddr4` block is documented below.

The `srcaddr4` block contains:

* `name` - Address name.
* `srcaddr6` - Source IPv6 address name and address group names.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name.
* `srcintf` - Incoming (ingress) interface.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
* `users` - Names of individual users that can authenticate with this policy.The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name.
