---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_securitypolicy"
description: |-
  Get information on a fortios Configure NGFW IPv4/IPv6 application policies.
---

# Data Source: fortios_firewall_securitypolicy
Use this data source to get information on a fortios Configure NGFW IPv4/IPv6 application policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Policy action (accept/deny).
* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Comment.
* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be.
* `dstaddr6_negate` - When enabled dstaddr6 specifies what the destination address must NOT be.
* `emailfilter_profile` - Name of an existing email filter profile.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications.
* `file_filter_profile` - Name of an existing file-filter profile.
* `icap_profile` - Name of an existing ICAP profile.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address, service and default application port enforcement are not used.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be.
* `internet_service6` - Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address, service and default application port enforcement are not used.
* `internet_service6_negate` - When enabled internet-service6 specifies what the service must NOT be.
* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be.
* `ips_sensor` - Name of an existing IPS sensor.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `name` - Policy name.
* `nat46` - Enable/disable NAT46.
* `nat64` - Enable/disable NAT64.
* `policyid` - Policy ID.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `schedule` - Schedule name.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be.
* `srcaddr6_negate` - When enabled srcaddr6 specifies what the source address must NOT be.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable or disable this policy.
* `url_category` - URL categories or groups.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `voip_profile` - Name of an existing VoIP (voipd) profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `app_category` - Application category ID list.The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - Application group names.The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group names.
* `application` - Application ID list.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `dstaddr` - Destination IPv4 address name and address group names.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
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

* `name` - User group name.
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
* `internet_service6_custom` - Custom IPv6 Internet Service name.The structure of `internet_service6_custom` block is documented below.

The `internet_service6_custom` block contains:

* `name` - Custom IPv6 Internet Service name.
* `internet_service6_custom_group` - Custom IPv6 Internet Service group name.The structure of `internet_service6_custom_group` block is documented below.

The `internet_service6_custom_group` block contains:

* `name` - Custom IPv6 Internet Service group name.
* `internet_service6_group` - Internet Service group name.The structure of `internet_service6_group` block is documented below.

The `internet_service6_group` block contains:

* `name` - Internet Service group name.
* `internet_service6_name` - IPv6 Internet Service name.The structure of `internet_service6_name` block is documented below.

The `internet_service6_name` block contains:

* `name` - IPv6 Internet Service name.
* `internet_service6_src_custom` - Custom IPv6 Internet Service source name.The structure of `internet_service6_src_custom` block is documented below.

The `internet_service6_src_custom` block contains:

* `name` - Custom Internet Service name.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name.The structure of `internet_service6_src_custom_group` block is documented below.

The `internet_service6_src_custom_group` block contains:

* `name` - Custom Internet Service6 group name.
* `internet_service6_src_group` - Internet Service6 source group name.The structure of `internet_service6_src_group` block is documented below.

The `internet_service6_src_group` block contains:

* `name` - Internet Service group name.
* `internet_service6_src_name` - IPv6 Internet Service source name.The structure of `internet_service6_src_name` block is documented below.

The `internet_service6_src_name` block contains:

* `name` - Internet Service name.
* `service` - Service and service group names.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source IPv4 address name and address group names.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
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
