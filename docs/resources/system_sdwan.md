---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan"
description: |-
  Configure redundant Internet connections with multiple outbound links and health-check profiles.
---

## fortios_system_sdwan
Configure redundant Internet connections with multiple outbound links and health-check profiles.

## Example Usage

```hcl
resource "fortios_system_sdwan" "example" {
    status = "enable"

    // ignore defaults
    lifecycle {
        ignore_changes = [
          health_check,
          members,
          zone
        ]
    }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `duplication_max_num` - Maximum number of interface members a packet is duplicated in the SD-WAN zone (2 - 4, default = 2; if set to 3, the original packet plus 2 more copies are created).
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable` `disable` .
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based` `weight-based` `usage-based` `source-dest-ip-based` `measured-volume-based` .
* `neighbor_hold_boot_time` - Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
* `neighbor_hold_down` - Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable` `disable` .
* `neighbor_hold_down_time` - Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
* `speedtest_bypass_routing` - Enable/disable bypass routing when speedtest on a SD-WAN member. Valid values: `disable` `enable` .
* `status` - Enable/disable SD-WAN. Valid values: `disable` `enable` .
* `duplication` - Create SD-WAN duplication rule. The structure of `duplication` block is documented below.

The `duplication` block contains:

* `id` - Duplication rule ID (1 - 255).
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated. Valid values: `enable` `disable` .
* `packet_duplication` - Configure packet duplication method. Valid values: `disable` `force` `on-demand` .
* `sla_match_service` - Enable/disable packet duplication matching health-check SLAs in service rule. Valid values: `enable` `disable` .
* `dstaddr` - Destination address or address group names. The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address or address group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dstaddr6` - Destination address6 or address6 group names. The structure of `dstaddr6` block is documented below.

The `dstaddr6` block contains:

* `name` - Address6 or address6 group name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `dstintf` - Outgoing (egress) interfaces or zones. The structure of `dstintf` block is documented below.

The `dstintf` block contains:

* `name` - Interface, zone or SDWAN zone name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `service` - Service and service group name. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service and service group name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `service_id` - SD-WAN service rule ID list. The structure of `service_id` block is documented below.

The `service_id` block contains:

* `id` - SD-WAN service rule ID. This attribute must reference one of the following datasources: `system.sdwan.service.id` .
* `srcaddr` - Source address or address group names. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address or address group name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `srcaddr6` - Source address6 or address6 group names. The structure of `srcaddr6` block is documented below.

The `srcaddr6` block contains:

* `name` - Address6 or address6 group name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
* `srcintf` - Incoming (ingress) interfaces or zones. The structure of `srcintf` block is documented below.

The `srcintf` block contains:

* `name` - Interface, zone or SDWAN zone name. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` `system.sdwan.zone.name` .
* `fail_alert_interfaces` - Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.

The `fail_alert_interfaces` block contains:

* `name` - Physical interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `health_check` - SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `health_check` block is documented below.

The `health_check` block contains:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4` `ipv6` .
* `detect_mode` - The mode determining how to detect the server. Valid values: `active` `passive` `prefer-passive` .
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `dns_match_ip` - Response IP expected from DNS server if the protocol is DNS.
* `dns_request_domain` - Fully qualified domain name to resolve for the DNS probe.
* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `ftp_file` - Full path and file name on the FTP server to download for FTP health-check to probe.
* `ftp_mode` - FTP mode. Valid values: `passive` `port` .
* `ha_priority` - HA election priority (1 - 50).
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `interval` - Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).
* `mos_codec` - Codec to use for MOS calculation (default = g711). Valid values: `g711` `g722` `g729` .
* `name` - Status check or health check name.
* `packet_size` - Packet size of a TWAMP test session.
* `password` - TWAMP controller password in authentication mode.
* `port` - Port number used to communicate with the server over the selected protocol (0 - 65535, default = 0, auto select. http, tcp-connect: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21, twamp: 862).
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `probe_packets` - Enable/disable transmission of probe packets. Valid values: `disable` `enable` .
* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server. Valid values: `ping` `tcp-echo` `udp-echo` `http` `twamp` `dns` `tcp-connect` `ftp` .
* `quality_measured_method` - Method to measure the quality of tcp-connect. Valid values: `half-open` `half-close` .
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `security_mode` - Twamp controller security mode. Valid values: `none` `authentication` .
* `server` - IP address or FQDN name of the server.
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `source` - Source IP address used in the health-check packet to the server.
* `system_dns` - Enable/disable system DNS as the probe server. Valid values: `disable` `enable` .
* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `update_cascade_interface` - Enable/disable update cascade interface. Valid values: `enable` `disable` .
* `update_static_route` - Enable/disable updating the static route. Valid values: `enable` `disable` .
* `user` - The user name to access probe server.
* `vrf` - Virtual Routing Forwarding ID.
* `members` - Member sequence number list. The structure of `members` block is documented below.

The `members` block contains:

* `seq_num` - Member sequence number. This attribute must reference one of the following datasources: `system.sdwan.members.seq-num` .
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `sla` block contains:

* `id` - SLA ID.
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency` `jitter` `packet-loss` `mos` .
* `mos_threshold` - Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
* `members` - FortiGate interfaces added to the SD-WAN. The structure of `members` block is documented below.

The `members` block contains:

* `comment` - Comments.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `gateway6` - IPv6 gateway.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `priority` - Priority of the interface for IPv4 (1 - 65535, default = 1). Used for SD-WAN rules or priority rules.
* `priority6` - Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.
* `seq_num` - Sequence number(1-512).
* `source` - Source IP address used in the health-check packet to the server.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `status` - Enable/disable this interface in the SD-WAN. Valid values: `disable` `enable` .
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).
* `weight` - Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.
* `zone` - Zone name. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .
* `neighbor` - Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `health_check` - SD-WAN health-check name. This attribute must reference one of the following datasources: `system.sdwan.health-check.name` .
* `ip` - IP/IPv6 address of neighbor. This attribute must reference one of the following datasources: `router.bgp.neighbor.ip` .
* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `mode` - What metric to select the neighbor. Valid values: `sla` `speedtest` .
* `role` - Role of neighbor. Valid values: `standalone` `primary` `secondary` .
* `sla_id` - SLA ID.
* `member` - Member sequence number list. The structure of `member` block is documented below.

The `member` block contains:

* `seq_num` - Member sequence number. This attribute must reference one of the following datasources: `system.sdwan.members.seq-num` .
* `service` - Create SD-WAN rules (also called services) to control how sessions are distributed to interfaces in the SD-WAN. The structure of `service` block is documented below.

The `service` block contains:

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
* `zone` - Configure SD-WAN zones. The structure of `zone` block is documented below.

The `zone` block contains:

* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `cfg-order` `fib-best-match` `input-device` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_sdwan can be imported using:
```sh
terraform import fortios_system_sdwan.labelname {{mkey}}
```
