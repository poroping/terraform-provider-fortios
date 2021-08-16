---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_service"
description: |-
  Get information on a fortios Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.
---

# Data Source: fortios_system_sdwan_service
Use this data source to get information on a fortios Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN.

## Example Usage

```hcl

```

## Argument Reference

* `fosid` - (Required) SD-WAN rule ID (1 - 4000).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `addr_mode` - Address mode (IPv4 or IPv6).
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `default` - Enable/disable use of SD-WAN as default service.
* `dscp_forward` - Enable/disable forward traffic DSCP tag.
* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag.
* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `dst_negate` - Enable/disable negation of destination address match.
* `end_port` - End destination port number.
* `gateway` - Enable/disable SD-WAN service gateway.
* `hash_mode` - Hash algorithm for selected priority members for load balance mode.
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `fosid` - SD-WAN rule ID (1 - 4000).
* `input_device_negate` - Enable/disable negation of input device match.
* `internet_service` - Enable/disable use of Internet service for application-based load balancing.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `link_cost_factor` - Link cost factor.
* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `minimum_sla_meet_members` - Minimum number of members which meet SLA.
* `mode` - Control how the SD-WAN rule sets the priority of interfaces in the SD-WAN.
* `name` - SD-WAN rule name.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `protocol` - Protocol number.
* `quality_link` - Quality grade.
* `role` - Service role to work with neighbor.
* `route_tag` - IPv4 route map route-tag.
* `sla_compare_method` - Method to compare SLA value for SLA mode.
* `src_negate` - Enable/disable negation of source address match.
* `standalone_action` - Enable/disable service when selected neighbor role is standalone while service role is not standalone.
* `start_port` - Start destination port number.
* `status` - Enable/disable SD-WAN service.
* `tie_break` - Method of selecting member if more than one meets the SLA.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `use_shortcut_sla` - Enable/disable use of ADVPN shortcut for quality comparison.
* `dst` - Destination address name.The structure of `dst` block is documented below.

The `dst` block contains:

* `name` - Address or address group name.
* `dst6` - Destination address6 name.The structure of `dst6` block is documented below.

The `dst6` block contains:

* `name` - Address6 or address6 group name.
* `groups` - User groups.The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name.
* `health_check` - Health check list.The structure of `health_check` block is documented below.

The `health_check` block contains:

* `name` - Health check name.
* `input_device` - Source interface name.The structure of `input_device` block is documented below.

The `input_device` block contains:

* `name` - Interface name.
* `internet_service_app_ctrl` - Application control based Internet Service ID list.The structure of `internet_service_app_ctrl` block is documented below.

The `internet_service_app_ctrl` block contains:

* `id` - Application control based Internet Service ID.
* `internet_service_app_ctrl_group` - Application control based Internet Service group list.The structure of `internet_service_app_ctrl_group` block is documented below.

The `internet_service_app_ctrl_group` block contains:

* `name` - Application control based Internet Service group name.
* `internet_service_custom` - Custom Internet service name list.The structure of `internet_service_custom` block is documented below.

The `internet_service_custom` block contains:

* `name` - Custom Internet service name.
* `internet_service_custom_group` - Custom Internet Service group list.The structure of `internet_service_custom_group` block is documented below.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name.
* `internet_service_group` - Internet Service group list.The structure of `internet_service_group` block is documented below.

The `internet_service_group` block contains:

* `name` - Internet Service group name.
* `internet_service_name` - Internet service name list.The structure of `internet_service_name` block is documented below.

The `internet_service_name` block contains:

* `name` - Internet service name.
* `priority_members` - Member sequence number list.The structure of `priority_members` block is documented below.

The `priority_members` block contains:

* `seq_num` - Member sequence number.
* `priority_zone` - Priority zone name list.The structure of `priority_zone` block is documented below.

The `priority_zone` block contains:

* `name` - Priority zone name.
* `sla` - Service level agreement (SLA).The structure of `sla` block is documented below.

The `sla` block contains:

* `health_check` - SD-WAN health-check.
* `id` - SLA ID.
* `src` - Source address name.The structure of `src` block is documented below.

The `src` block contains:

* `name` - Address or address group name.
* `src6` - Source address6 name.The structure of `src6` block is documented below.

The `src6` block contains:

* `name` - Address6 or address6 group name.
* `users` - User name.The structure of `users` block is documented below.

The `users` block contains:

* `name` - User name.
