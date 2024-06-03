---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxypolicy"
description: |-
  Get information on a fortios Configure proxy policies.
---

# Data Source: fortios_firewall_proxypolicy
Use this data source to get information on a fortios Configure proxy policies.


## Example Usage

```hcl

```

## Argument Reference

* `policyid` - (Required) Policy ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Accept or deny traffic matching the policy parameters.
* `application_list` - Name of an existing Application list.
* `av_profile` - Name of an existing Antivirus profile.
* `block_notification` - Enable/disable block notification.
* `cifs_profile` - Name of an existing CIFS profile.
* `comments` - Optional comments.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `device_ownership` - When enabled, the ownership enforcement will be done at policy level.
* `disclaimer` - Web proxy disclaimer setting: by domain, policy, or user.
* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `dstaddr_negate` - When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
* `emailfilter_profile` - Name of an existing email filter profile.
* `file_filter_profile` - Name of an existing file-filter profile.
* `http_tunnel_auth` - Enable/disable HTTP tunnel authentication.
* `icap_profile` - Name of an existing ICAP profile.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
* `internet_service_negate` - When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
* `internet_service6` - Enable/disable use of Internet Services IPv6 for this policy. If enabled, destination IPv6 address and service are not used.
* `internet_service6_negate` - When enabled, Internet Services match against any internet service IPv6 EXCEPT the selected Internet Service IPv6.
* `ips_sensor` - Name of an existing IPS sensor.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `logtraffic` - Enable/disable logging traffic through the policy.
* `logtraffic_start` - Enable/disable policy log traffic start.
* `name` - Policy name.
* `policyid` - Policy ID.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `proxy` - Type of explicit proxy.
* `redirect_url` - Redirect URL for further explicit web proxy processing.
* `replacemsg_override_group` - Authentication replacement message override group.
* `schedule` - Name of schedule object.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `service_negate` - When enabled, services match against any service EXCEPT the specified destination services.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr_negate` - When enabled, source addresses match against any address EXCEPT the specified source addresses.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `status` - Enable/disable the active status of the policy.
* `transparent` - Enable to use the IP address of the client to connect to the server.
* `utm_status` - Enable the use of UTM profiles/sensors/lists.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `voip_profile` - Name of an existing VoIP profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `webcache` - Enable/disable web caching.
* `webcache_https` - Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
* `webfilter_profile` - Name of an existing Web filter profile.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Name of web proxy profile.
* `ztna_tags_match_logic` - ZTNA tag matching logic.
* `access_proxy` - IPv4 access proxy.The structure of `access_proxy` block is documented below.

The `access_proxy` block contains:

* `name` - Access Proxy name.
* `access_proxy6` - IPv6 access proxy.The structure of `access_proxy6` block is documented below.

The `access_proxy6` block contains:

* `name` - Access proxy name.
* `dstaddr` - Destination address objects.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `dstaddr6` - IPv6 destination address objects.The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name.
* `dstintf` - Destination interface names.The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name.
* `groups` - Names of group objects.The structure of `groups` block is documented below.

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
* `internet_service6_custom` - Custom Internet Service IPv6 name.The structure of `internet_service6_custom` block is documented below.

The `internet_service6_custom` block contains:

* `name` - Custom Internet Service IPv6 name.
* `internet_service6_custom_group` - Custom Internet Service IPv6 group name.The structure of `internet_service6_custom_group` block is documented below.

The `internet_service6_custom_group` block contains:

* `name` - Custom Internet Service IPv6 group name.
* `internet_service6_group` - Internet Service IPv6 group name.The structure of `internet_service6_group` block is documented below.

The `internet_service6_group` block contains:

* `name` - Internet Service IPv6 group name.
* `internet_service6_name` - Internet Service IPv6 name.The structure of `internet_service6_name` block is documented below.

The `internet_service6_name` block contains:

* `name` - Internet Service IPv6 name.
* `poolname` - Name of IP pool object.The structure of `poolname` block is documented below.

The `poolname` block contains:

* `name` - IP pool name.
* `service` - Name of service objects.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source address objects.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
* `srcaddr6` - IPv6 source address objects.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name.
* `srcintf` - Source interface names.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
* `users` - Names of user objects.The structure of `users` block is documented below.

The `users` block contains:

* `name` - Group name.
* `ztna_ems_tag` - ZTNA EMS Tag names.The structure of `ztna_ems_tag` block is documented below.

The `ztna_ems_tag` block contains:

* `name` - EMS Tag name.
