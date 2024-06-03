---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_shapingpolicy"
description: |-
  Configure shaping policies.
---

## fortios_firewall_shapingpolicy
Configure shaping policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `class_id` - Traffic class ID. This attribute must reference one of the following datasources: `firewall.traffic-class.class-id` .
* `comment` - Comments.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `enable` `disable` .
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `enable` `disable` .
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `id` - Shaping policy ID (0 - 4294967295).
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `enable` `disable` .
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `enable` `disable` .
* `ip_version` - Apply this traffic shaping policy to IPv4 or IPv6 traffic. Valid values: `4` `6` .
* `name` - Shaping policy name.
* `per_ip_shaper` - Per-IP traffic shaper to apply with this policy. This attribute must reference one of the following datasources: `firewall.shaper.per-ip-shaper.name` .
* `schedule` - Schedule name. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `status` - Enable/disable this traffic shaping policy. Valid values: `enable` `disable` .
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `enable` `disable` .
* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the firewall policy. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `traffic_shaper_reverse` - Traffic shaper to apply to response traffic received by the firewall policy. This attribute must reference one of the following datasources: `firewall.shaper.traffic-shaper.name` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `app_category` - IDs of one or more application categories that this shaper applies application control traffic shaping to. The structure of `app_category` block is documented below.

The `app_category` block contains:

* `id` - Category IDs.
* `app_group` - One or more application group names. The structure of `app_group` block is documented below.

The `app_group` block contains:

* `name` - Application group name. This attribute must reference one of the following datasources: `application.group.name` .
* `application` - IDs of one or more applications that this shaper applies application control traffic shaping to. The structure of `application` block is documented below.

The `application` block contains:

* `id` - Application IDs.
* `dstaddr` - IPv4 destination address and address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dstaddr6` - IPv6 destination address and address group names. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `dstintf` - One or more outgoing (egress) interfaces. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `groups` - Apply this traffic shaping policy to user groups that have authenticated with the FortiGate. The structure of `groups` block is documented below.

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
* `internet_service_name` - Internet Service ID. The structure of `internet_service_name` block is documented below.

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
* `service` - Service and service group names. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - IPv4 source address and address group names. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `srcaddr6` - IPv6 source address and address group names. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `srcintf` - One or more incoming (ingress) interfaces. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `url_category` - IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to. The structure of `url_category` block is documented below.

The `url_category` block contains:

* `id` - URL category ID.
* `users` - Apply this traffic shaping policy to individual users that have authenticated with the FortiGate. The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name. This attribute must reference one of the following datasources: `user.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_shapingpolicy can be imported using:
```sh
terraform import fortios_firewall_shapingpolicy.labelname {{mkey}}
```
