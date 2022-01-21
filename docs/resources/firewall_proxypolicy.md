---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxypolicy"
description: |-
  Configure proxy policies.
---

## fortios_firewall_proxypolicy
Configure proxy policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Accept or deny traffic matching the policy parameters. Valid values: `accept` `deny` `redirect` .
* `application_list` - Name of an existing Application list. This attribute must reference one of the following datasources: `application.list.name` .
* `av_profile` - Name of an existing Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `cifs_profile` - Name of an existing CIFS profile. This attribute must reference one of the following datasources: `cifs.profile.name` .
* `comments` - Optional comments.
* `decrypted_traffic_mirror` - Decrypted traffic mirror. This attribute must reference one of the following datasources: `firewall.decrypted-traffic-mirror.name` .
* `device_ownership` - When enabled, the ownership enforcement will be done at policy level. Valid values: `enable` `disable` .
* `disclaimer` - Web proxy disclaimer setting: by domain, policy, or user. Valid values: `disable` `domain` `policy` `user` .
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dstaddr_negate` - When enabled, destination addresses match against any address EXCEPT the specified destination addresses. Valid values: `enable` `disable` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `http_tunnel_auth` - Enable/disable HTTP tunnel authentication. Valid values: `enable` `disable` .
* `icap_profile` - Name of an existing ICAP profile. This attribute must reference one of the following datasources: `icap.profile.name` .
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `enable` `disable` .
* `internet_service_negate` - When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service. Valid values: `enable` `disable` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `logtraffic` - Enable/disable logging traffic through the policy. Valid values: `all` `utm` `disable` .
* `logtraffic_start` - Enable/disable policy log traffic start. Valid values: `enable` `disable` .
* `name` - Policy name.
* `policyid` - Policy ID.
* `profile_group` - Name of profile group. This attribute must reference one of the following datasources: `firewall.profile-group.name` .
* `profile_protocol_options` - Name of an existing Protocol options profile. This attribute must reference one of the following datasources: `firewall.profile-protocol-options.name` .
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single` `group` .
* `proxy` - Type of explicit proxy. Valid values: `explicit-web` `transparent-web` `ftp` `ssh` `ssh-tunnel` `access-proxy` `wanopt` .
* `redirect_url` - Redirect URL for further explicit web proxy processing.
* `replacemsg_override_group` - Authentication replacement message override group. This attribute must reference one of the following datasources: `system.replacemsg-group.name` .
* `schedule` - Name of schedule object. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `sctp_filter_profile` - Name of an existing SCTP filter profile. This attribute must reference one of the following datasources: `sctp-filter.profile.name` .
* `service_negate` - When enabled, services match against any service EXCEPT the specified destination services. Valid values: `enable` `disable` .
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr_negate` - When enabled, source addresses match against any address EXCEPT the specified source addresses. Valid values: `enable` `disable` .
* `ssh_filter_profile` - Name of an existing SSH filter profile. This attribute must reference one of the following datasources: `ssh-filter.profile.name` .
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `enable` `disable` .
* `ssl_ssh_profile` - Name of an existing SSL SSH profile. This attribute must reference one of the following datasources: `firewall.ssl-ssh-profile.name` .
* `status` - Enable/disable the active status of the policy. Valid values: `enable` `disable` .
* `transparent` - Enable to use the IP address of the client to connect to the server. Valid values: `enable` `disable` .
* `utm_status` - Enable the use of UTM profiles/sensors/lists. Valid values: `enable` `disable` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile. This attribute must reference one of the following datasources: `videofilter.profile.name` .
* `voip_profile` - Name of an existing VoIP profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `waf_profile` - Name of an existing Web application firewall profile. This attribute must reference one of the following datasources: `waf.profile.name` .
* `webcache` - Enable/disable web caching. Valid values: `enable` `disable` .
* `webcache_https` - Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile). Valid values: `disable` `enable` .
* `webfilter_profile` - Name of an existing Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `webproxy_forward_server` - Web proxy forward server name. This attribute must reference one of the following datasources: `web-proxy.forward-server.name` `web-proxy.forward-server-group.name` .
* `webproxy_profile` - Name of web proxy profile. This attribute must reference one of the following datasources: `web-proxy.profile.name` .
* `ztna_tags_match_logic` - ZTNA tag matching logic. Valid values: `or` `and` .
* `access_proxy` - IPv4 access proxy. The structure of `access_proxy` block is documented below.

The `access_proxy` block contains:

* `name` - Access Proxy name. This attribute must reference one of the following datasources: `firewall.access-proxy.name` .
* `access_proxy6` - IPv6 access proxy. The structure of `access_proxy6` block is documented below.

The `access_proxy6` block contains:

* `name` - Access proxy name. This attribute must reference one of the following datasources: `firewall.access-proxy6.name` .
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` `firewall.proxy-addrgrp.name` `firewall.vip.name` `firewall.vipgrp.name` `system.external-resource.name` .
* `dstaddr6` - IPv6 destination address objects. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `firewall.vip6.name` `firewall.vipgrp6.name` `system.external-resource.name` .
* `dstintf` - Destination interface names. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `groups` - Names of group objects. The structure of `groups` block is documented below.

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
* `poolname` - Name of IP pool object. The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name. This attribute must reference one of the following datasources: `firewall.ippool.name` .
* `service` - Name of service objects. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.proxy-address.name` `firewall.proxy-addrgrp.name` `system.external-resource.name` .
* `srcaddr6` - IPv6 source address objects. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `srcintf` - Source interface names. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `users` - Names of user objects. The structure of `users` block is documented below.

The `users` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.local.name` `user.certificate.name` .
* `ztna_ems_tag` - ZTNA EMS Tag names. The structure of `ztna_ems_tag` block is documented below.

The `ztna_ems_tag` block contains:

* `name` - EMS Tag name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_proxypolicy can be imported using:
```sh
terraform import fortios_firewall_proxypolicy.labelname {{mkey}}
```
