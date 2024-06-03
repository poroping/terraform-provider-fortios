---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_securitypolicy"
description: |-
  Configure NGFW IPv4/IPv6 application policies.
---

## fortios_firewall_securitypolicy
Configure NGFW IPv4/IPv6 application policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `policyid` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Policy action (accept/deny). Valid values: `accept` `deny` .
* `application_list` - Name of an existing Application list. This attribute must reference one of the following datasources: `application.list.name` .
* `av_profile` - Name of an existing Antivirus profile. This attribute must reference one of the following datasources: `antivirus.profile.name` .
* `cifs_profile` - Name of an existing CIFS profile. This attribute must reference one of the following datasources: `cifs.profile.name` .
* `comments` - Comment.
* `dlp_profile` - Name of an existing DLP profile. This attribute must reference one of the following datasources: `dlp.profile.name` .
* `dlp_sensor` - Name of an existing DLP sensor. This attribute must reference one of the following datasources: `dlp.sensor.name` .
* `dnsfilter_profile` - Name of an existing DNS filter profile. This attribute must reference one of the following datasources: `dnsfilter.profile.name` .
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable` `disable` .
* `dstaddr6_negate` - When enabled dstaddr6 specifies what the destination address must NOT be. Valid values: `enable` `disable` .
* `emailfilter_profile` - Name of an existing email filter profile. This attribute must reference one of the following datasources: `emailfilter.profile.name` .
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `enable` `disable` .
* `file_filter_profile` - Name of an existing file-filter profile. This attribute must reference one of the following datasources: `file-filter.profile.name` .
* `icap_profile` - Name of an existing ICAP profile. This attribute must reference one of the following datasources: `icap.profile.name` .
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address, service and default application port enforcement are not used. Valid values: `enable` `disable` .
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `enable` `disable` .
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `enable` `disable` .
* `internet_service6` - Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address, service and default application port enforcement are not used. Valid values: `enable` `disable` .
* `internet_service6_negate` - When enabled internet-service6 specifies what the service must NOT be. Valid values: `enable` `disable` .
* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used. Valid values: `enable` `disable` .
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be. Valid values: `enable` `disable` .
* `ips_sensor` - Name of an existing IPS sensor. This attribute must reference one of the following datasources: `ips.sensor.name` .
* `ips_voip_filter` - Name of an existing VoIP (ips) profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated. Valid values: `enable` `disable` .
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `all` `utm` `disable` .
* `logtraffic_start` - Record logs when a session starts. Valid values: `enable` `disable` .
* `name` - Policy name.
* `nat46` - Enable/disable NAT46. Valid values: `enable` `disable` .
* `nat64` - Enable/disable NAT64. Valid values: `enable` `disable` .
* `policyid` - Policy ID.
* `profile_group` - Name of profile group. This attribute must reference one of the following datasources: `firewall.profile-group.name` .
* `profile_protocol_options` - Name of an existing Protocol options profile. This attribute must reference one of the following datasources: `firewall.profile-protocol-options.name` .
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single` `group` .
* `schedule` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `sctp_filter_profile` - Name of an existing SCTP filter profile. This attribute must reference one of the following datasources: `sctp-filter.profile.name` .
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable` `enable` .
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable` `disable` .
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable` `disable` .
* `srcaddr6_negate` - When enabled srcaddr6 specifies what the source address must NOT be. Valid values: `enable` `disable` .
* `ssh_filter_profile` - Name of an existing SSH filter profile. This attribute must reference one of the following datasources: `ssh-filter.profile.name` .
* `ssl_ssh_profile` - Name of an existing SSL SSH profile. This attribute must reference one of the following datasources: `firewall.ssl-ssh-profile.name` .
* `status` - Enable or disable this policy. Valid values: `enable` `disable` .
* `url_category` - URL categories or groups.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile. This attribute must reference one of the following datasources: `videofilter.profile.name` .
* `voip_profile` - Name of an existing VoIP (voipd) profile. This attribute must reference one of the following datasources: `voip.profile.name` .
* `webfilter_profile` - Name of an existing Web filter profile. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - Application group names. The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group names. This attribute must reference one of the following datasources: `application.group.name` .
* `application` - Application ID list. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `dstaddr` - Destination IPv4 address name and address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.vip.name` `firewall.vipgrp.name` `system.external-resource.name` .
* `dstaddr4` - Destination IPv4 address name and address group names. The structure of `dstaddr4` block is documented below.

The `dstaddr4` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `firewall.vip.name` `firewall.vipgrp.name` .
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `firewall.vip6.name` `firewall.vipgrp6.name` `system.external-resource.name` .
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below.

The `fsso_groups` block contains:

* `name` - Names of FSSO groups. This attribute must reference one of the following datasources: `user.adgrp.name` .
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - User group name. This attribute must reference one of the following datasources: `user.group.name` .
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
* `internet_service6_custom` - Custom IPv6 Internet Service name. The structure of `internet_service6_custom` block is documented below.

The `internet_service6_custom` block contains:

* `name` - Custom IPv6 Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-custom.name` .
* `internet_service6_custom_group` - Custom IPv6 Internet Service group name. The structure of `internet_service6_custom_group` block is documented below.

The `internet_service6_custom_group` block contains:

* `name` - Custom IPv6 Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-custom-group.name` .
* `internet_service6_group` - Internet Service group name. The structure of `internet_service6_group` block is documented below.

The `internet_service6_group` block contains:

* `name` - Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-group.name` .
* `internet_service6_name` - IPv6 Internet Service name. The structure of `internet_service6_name` block is documented below.

The `internet_service6_name` block contains:

* `name` - IPv6 Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-name.name` .
* `internet_service6_src_custom` - Custom IPv6 Internet Service source name. The structure of `internet_service6_src_custom` block is documented below.

The `internet_service6_src_custom` block contains:

* `name` - Custom Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-custom.name` .
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name. The structure of `internet_service6_src_custom_group` block is documented below.

The `internet_service6_src_custom_group` block contains:

* `name` - Custom Internet Service6 group name. This attribute must reference one of the following datasources: `firewall.internet-service-custom-group.name` .
* `internet_service6_src_group` - Internet Service6 source group name. The structure of `internet_service6_src_group` block is documented below.

The `internet_service6_src_group` block contains:

* `name` - Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-group.name` .
* `internet_service6_src_name` - IPv6 Internet Service source name. The structure of `internet_service6_src_name` block is documented below.

The `internet_service6_src_name` block contains:

* `name` - Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-name.name` .
* `service` - Service and service group names. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source IPv4 address name and address group names. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` `system.external-resource.name` .
* `srcaddr4` - Source IPv4 address name and address group names. The structure of `srcaddr4` block is documented below.

The `srcaddr4` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` `system.external-resource.name` .
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name. This attribute must reference one of the following datasources: `user.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_securitypolicy can be imported using:
```sh
terraform import fortios_firewall_securitypolicy.labelname {{mkey}}
```
