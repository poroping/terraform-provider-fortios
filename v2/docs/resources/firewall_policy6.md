---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy6"
description: |-
  Configure IPv6 policies.
---

## fortios_firewall_policy6
Configure IPv6 policies.

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
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `enable` `disable` .
* `av_profile` - Name of an existing Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `cifs_profile` - Name of an existing CIFS profile. This attribute must reference one of the following datasources: `cifs.profile.name` .
* `comments` - Comment.
* `decrypted_traffic_mirror` - Decrypted traffic mirror. This attribute must reference one of the following datasources: `firewall.decrypted-traffic-mirror.name` .
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `enable` `disable` .
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `enable` `disable` .
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dnsfilter_profile` - Name of an existing DNS filter profile. This attribute must reference one of the following datasources: `dnsfilter.profile.name` .
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `enable` `disable` .
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable` `disable` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all` `check-new` .
* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `enable` `disable` .
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `enable` `disable` .
* `icap_profile` - Name of an existing ICAP profile. This attribute must reference one of the following datasources: `icap.profile.name` .
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `enable` `disable` .
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy` `flow` .
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `enable` `disable` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `all` `utm` `disable` .
* `logtraffic_start` - Record logs when a session starts. Valid values: `enable` `disable` .
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `enable` `disable` .
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic. Valid values: `enable` `disable` .
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic. Valid values: `enable` `disable` .
* `np_acceleration` - Enable/disable UTM Network Processor acceleration. Valid values: `enable` `disable` .
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `enable` `disable` .
* `per_ip_shaper` - Per-IP traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.per-ip-shaper.name` .
* `policyid` - Policy ID (0 - 4294967294).
* `profile_group` - Name of profile group. This attribute must reference one of the following datasources: `firewall.profile-group.name` .
* `profile_protocol_options` - Name of an existing Protocol options profile. This attribute must reference one of the following datasources: `firewall.profile-protocol-options.name` .
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single` `group` .
* `replacemsg_override_group` - Override the default replacement message group for this policy. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `enable` `disable` .
* `schedule` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `send_deny_packet` - Enable/disable return of deny-packet. Valid values: `enable` `disable` .
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `session_ttl` - Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable` `disable` .
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
* `traffic_shaper` - Reverse traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `traffic_shaper_reverse` - Reverse traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `utm_status` - Enable AV/web/ips protection profile. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_filter` - Set VLAN filters.
* `voip_profile` - Name of an existing VoIP profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1. This attribute must reference one of the following datasources: `vpn.ipsec.phase1.name` `vpn.ipsec.manualkey.name` .
* `waf_profile` - Name of an existing Web application firewall profile. This attribute must reference one of the following datasources: `waf.profile.name` .
* `webcache` - Enable/disable web cache. Valid values: `enable` `disable` .
* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable` `enable` .
* `webfilter_profile` - Name of an existing Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `webproxy_forward_server` - Web proxy forward server name. This attribute must reference one of the following datasources: `web-proxy.forward-server.name` `web-proxy.forward-server-group.name` .
* `webproxy_profile` - Webproxy profile name. This attribute must reference one of the following datasources: `web-proxy.profile.name` .
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - Application group names. The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group names. This attribute must reference one of the following datasources: `application.group.name` .
* `application` - Application ID list. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `custom_log_fields` - Log field index numbers to append custom log fields to log messages for this policy. The structure of `custom_log_fields` block is documented below.

The `custom_log_fields` block contains:

* `field_id` - Custom log field. This attribute must reference one of the following datasources: `log.custom-field.id` .
* `dstaddr` - Destination address and address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `firewall.vip6.name` `firewall.vipgrp6.name` `system.external-resource.name` .
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below.

The `fsso_groups` block contains:

* `name` - Names of FSSO groups. This attribute must reference one of the following datasources: `user.adgrp.name` .
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name. This attribute must reference one of the following datasources: `firewall.ippool6.name` .
* `service` - Service and service group names. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source address and address group names. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` `system.sdwan.zone.name` .
* `ssl_mirror_intf` - SSL mirror interface name. The structure of `ssl_mirror_intf` block is documented below.

The `ssl_mirror_intf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` .
* `url_category` - URL category ID list. The structure of `url_category` block is documented below.

The `url_category` block contains:

* `id` - URL category ID.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.

The `users` block contains:

* `name` - Names of individual users that can authenticate with this policy. This attribute must reference one of the following datasources: `user.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_policy6 can be imported using:
```sh
terraform import fortios_firewall_policy6.labelname {{mkey}}
```
