---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallconsolidated_policy"
description: |-
  Configure consolidated IPv4/IPv6 policies.
---

## fortios_firewallconsolidated_policy
Configure consolidated IPv4/IPv6 policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Policy action (accept/deny/ipsec). Valid values: `accept` `deny` `ipsec` .
* `application_list` - Name of an existing Application list. This attribute must reference one of the following datasources: `application.list.name` .
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `enable` `disable` .
* `av_profile` - Name of an existing Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `captive_portal_exempt` - Enable exemption of some users from the captive portal. Valid values: `enable` `disable` .
* `cifs_profile` - Name of an existing CIFS profile. This attribute must reference one of the following datasources: `cifs.profile.name` .
* `comments` - Comment.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `enable` `disable` .
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.  Valid values: `enable` `disable` .
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dnsfilter_profile` - Name of an existing DNS filter profile. This attribute must reference one of the following datasources: `dnsfilter.profile.name` .
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable` `disable` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `enable` `disable` .
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `enable` `disable` .
* `icap_profile` - Name of an existing ICAP profile. This attribute must reference one of the following datasources: `icap.profile.name` .
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `enable` `disable` .
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy` `flow` .
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.  Valid values: `enable` `disable` .
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.  Valid values: `enable` `disable` .
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `enable` `disable` .
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `enable` `disable` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `all` `utm` `disable` .
* `logtraffic_start` - Record logs when a session starts. Valid values: `enable` `disable` .
* `name` - Policy name.
* `nat` - Enable/disable source NAT. Valid values: `enable` `disable` .
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `enable` `disable` .
* `per_ip_shaper` - Per-IP traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.per-ip-shaper.name` .
* `policyid` - Policy ID (0 - 4294967294).
* `profile_group` - Name of profile group. This attribute must reference one of the following datasources: `firewall.profile-group.name` .
* `profile_protocol_options` - Name of an existing Protocol options profile. This attribute must reference one of the following datasources: `firewall.profile-protocol-options.name` .
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single` `group` .
* `schedule` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable` `disable` .
* `ssh_filter_profile` - Name of an existing SSH filter profile. This attribute must reference one of the following datasources: `ssh-filter.profile.name` .
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `enable` `disable` .
* `ssl_ssh_profile` - Name of an existing SSL SSH profile. This attribute must reference one of the following datasources: `firewall.ssl-ssh-profile.name` .
* `status` - Enable or disable this policy. Valid values: `enable` `disable` .
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `traffic_shaper` - Traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `traffic_shaper_reverse` - Reverse traffic shaper. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `voip_profile` - Name of an existing VoIP profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1. This attribute must reference one of the following datasources: `vpn.ipsec.phase1.name` `vpn.ipsec.manualkey.name` .
* `waf_profile` - Name of an existing Web application firewall profile. This attribute must reference one of the following datasources: `waf.profile.name` .
* `wanopt` - Enable/disable WAN optimization. Valid values: `enable` `disable` .
* `wanopt_detection` - WAN optimization auto-detection mode. Valid values: `active` `passive` `off` .
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect to server. Valid values: `default` `transparent` `non-transparent` .
* `wanopt_peer` - WAN optimization peer. This attribute must reference one of the following datasources: `wanopt.peer.peer-host-id` .
* `wanopt_profile` - WAN optimization profile. This attribute must reference one of the following datasources: `wanopt.profile.name` .
* `webcache` - Enable/disable web cache. Valid values: `enable` `disable` .
* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable` `enable` .
* `webfilter_profile` - Name of an existing Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `webproxy_forward_server` - Webproxy forward server name. This attribute must reference one of the following datasources: `web-proxy.forward-server.name` `web-proxy.forward-server-group.name` .
* `webproxy_profile` - Webproxy profile name. This attribute must reference one of the following datasources: `web-proxy.profile.name` .
* `dstaddr4` - Destination IPv4 address name and address group names. The structure of `dstaddr4` block is documented below.

The `dstaddr4` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.vip.name` `firewall.vipgrp.name` `system.external-resource.name` .
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `firewall.vip6.name` `firewall.vipgrp6.name` `system.external-resource.name` .
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
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
* `poolname4` - IPv4 pool names. The structure of `poolname4` block is documented below.

The `poolname4` block contains:

* `name` - IPv4 pool name. This attribute must reference one of the following datasources: `firewall.ippool.name` .
* `poolname6` - IPv6 pool names. The structure of `poolname6` block is documented below.

The `poolname6` block contains:

* `name` - IPv6 pool name. This attribute must reference one of the following datasources: `firewall.ippool6.name` .
* `service` - Service and service group names. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr4` - Source IPv4 address name and address group names. The structure of `srcaddr4` block is documented below.

The `srcaddr4` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `system.external-resource.name` .
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name. This attribute must reference one of the following datasources: `user.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewallconsolidated_policy can be imported using:
```sh
terraform import fortios_firewallconsolidated_policy.labelname {{mkey}}
```
