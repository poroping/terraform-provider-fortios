---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_service"
description: |-
  Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.
---

## fortios_system_sdwan_service
Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

~> This resource is configuring a child table of the parent resource: `fortios_system_sdwan`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4` `ipv6` .
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `default` - Enable/disable use of SD-WAN as default service. Valid values: `enable` `disable` .
* `dscp_forward` - Enable/disable forward traffic DSCP tag. Valid values: `enable` `disable` .
* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag. Valid values: `enable` `disable` .
* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `dst_negate` - Enable/disable negation of destination address match. Valid values: `enable` `disable` .
* `end_port` - End destination port number.
* `gateway` - Enable/disable SD-WAN service gateway. Valid values: `enable` `disable` .
* `hash_mode` - Hash algorithm for selected priority members for load balance mode. Valid values: `round-robin` `source-ip-based` `source-dest-ip-based` `inbandwidth` `outbandwidth` `bibandwidth` .
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `id` - SD-WAN rule ID (1 - 4000).
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `enable` `disable` .
* `internet_service` - Enable/disable use of Internet service for application-based load balancing. Valid values: `enable` `disable` .
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `link_cost_factor` - Link cost factor. Valid values: `latency` `jitter` `packet-loss` `inbandwidth` `outbandwidth` `bibandwidth` `custom-profile-1` .
* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `minimum_sla_meet_members` - Minimum number of members which meet SLA.
* `mode` - Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN. Valid values: `auto` `manual` `priority` `sla` `load-balance` .
* `name` - SD-WAN rule name.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `passive_measurement` - Enable/disable passive measurement based on the service criteria. Valid values: `enable` `disable` .
* `protocol` - Protocol number.
* `quality_link` - Quality grade.
* `role` - Service role to work with neighbor. Valid values: `standalone` `primary` `secondary` .
* `route_tag` - IPv4 route map route-tag.
* `sla_compare_method` - Method to compare SLA value for SLA mode. Valid values: `order` `number` .
* `src_negate` - Enable/disable negation of source address match. Valid values: `enable` `disable` .
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone. Valid values: `enable` `disable` .
* `start_port` - Start destination port number.
* `status` - Enable/disable SD-WAN service. Valid values: `enable` `disable` .
* `tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `zone` `cfg-order` `fib-best-match` `input-device` .
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `use_shortcut_sla` - Enable/disable use of ADVPN shortcut for quality comparison. Valid values: `enable` `disable` .
* `dst` - Destination address name. The structure of `dst` block is documented below.

The `dst` block contains:

* `name` - Address or address group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dst6` - Destination address6 name. The structure of `dst6` block is documented below.

The `dst6` block contains:

* `name` - Address6 or address6 group name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `groups` - User groups. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .
* `health_check` - Health check list. The structure of `health_check` block is documented below.

The `health_check` block contains:

* `name` - Health check name. This attribute must reference one of the following datasources: `system.sdwan.health-check.name` .
* `input_device` - Source interface name. The structure of `input_device` block is documented below.

The `input_device` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `input_zone` - Source input-zone name. The structure of `input_zone` block is documented below.

The `input_zone` block contains:

* `name` - Zone. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .
* `internet_service_app_ctrl` - Application control based Internet Service ID list. The structure of `internet_service_app_ctrl` block is documented below.

The `internet_service_app_ctrl` block contains:

* `id` - Application control based Internet Service ID.
* `internet_service_app_ctrl_category` - IDs of one or more application control categories. The structure of `internet_service_app_ctrl_category` block is documented below.

The `internet_service_app_ctrl_category` block contains:

* `id` - Application control category ID.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list. The structure of `internet_service_app_ctrl_group` block is documented below.

The `internet_service_app_ctrl_group` block contains:

* `name` - Application control based Internet Service group name. This attribute must reference one of the following datasources: `application.group.name` .
* `internet_service_custom` - Custom Internet service name list. The structure of `internet_service_custom` block is documented below.

The `internet_service_custom` block contains:

* `name` - Custom Internet service name. This attribute must reference one of the following datasources: `firewall.internet-service-custom.name` .
* `internet_service_custom_group` - Custom Internet Service group list. The structure of `internet_service_custom_group` block is documented below.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-custom-group.name` .
* `internet_service_group` - Internet Service group list. The structure of `internet_service_group` block is documented below.

The `internet_service_group` block contains:

* `name` - Internet Service group name. This attribute must reference one of the following datasources: `firewall.internet-service-group.name` .
* `internet_service_name` - Internet service name list. The structure of `internet_service_name` block is documented below.

The `internet_service_name` block contains:

* `name` - Internet service name. This attribute must reference one of the following datasources: `firewall.internet-service-name.name` .
* `priority_members` - Member sequence number list. The structure of `priority_members` block is documented below.

The `priority_members` block contains:

* `seq_num` - Member sequence number. This attribute must reference one of the following datasources: `system.sdwan.members.seq-num` .
* `priority_zone` - Priority zone name list. The structure of `priority_zone` block is documented below.

The `priority_zone` block contains:

* `name` - Priority zone name. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `sla` block contains:

* `health_check` - SD-WAN health-check. This attribute must reference one of the following datasources: `system.sdwan.health-check.name` .
* `id` - SLA ID.
* `src` - Source address name. The structure of `src` block is documented below.

The `src` block contains:

* `name` - Address or address group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `src6` - Source address6 name. The structure of `src6` block is documented below.

The `src6` block contains:

* `name` - Address6 or address6 group name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `users` - User name. The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name. This attribute must reference one of the following datasources: `user.local.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_service can be imported using:
```sh
terraform import fortios_system_service.labelname {{mkey}}
```
