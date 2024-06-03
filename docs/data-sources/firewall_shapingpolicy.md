---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_shapingpolicy"
description: |-
  Get information on a fortios Configure shaping policies.
---

# Data Source: fortios_firewall_shapingpolicy
Use this data source to get information on a fortios Configure shaping policies.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Shaping policy ID (0 - 4294967295).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `class_id` - Traffic class ID.
* `comment` - Comments.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `id` - Shaping policy ID (0 - 4294967295).
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.
* `ip_version` - Apply this traffic shaping policy to IPv4 or IPv6 traffic.
* `name` - Shaping policy name.
* `per_ip_shaper` - Per-IP traffic shaper to apply with this policy.
* `schedule` - Schedule name.
* `status` - Enable/disable this traffic shaping policy.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match.
* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the firewall policy.
* `traffic_shaper_reverse` - Traffic shaper to apply to response traffic received by the firewall policy.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `app_category` - IDs of one or more application categories that this shaper applies application control traffic shaping to.The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - One or more application group names.The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group name.
* `application` - IDs of one or more applications that this shaper applies application control traffic shaping to.The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `dstaddr` - IPv4 destination address and address group names.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name.
* `dstaddr6` - IPv6 destination address and address group names.The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name.
* `dstintf` - One or more outgoing (egress) interfaces.The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name.
* `groups` - Apply this traffic shaping policy to user groups that have authenticated with the FortiGate.The structure of `groups` block is documented below.

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
* `internet_service_name` - Internet Service ID.The structure of `internet_service_name` block is documented below.

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
* `service` - Service and service group names.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - IPv4 source address and address group names.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
* `srcaddr6` - IPv6 source address and address group names.The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name.
* `srcintf` - One or more incoming (ingress) interfaces.The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.
* `url_category` - IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to.The structure of `url_category` block is documented below.

The `url_category` block contains:

* `id` - URL category ID.
* `users` - Apply this traffic shaping policy to individual users that have authenticated with the FortiGate.The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name.
