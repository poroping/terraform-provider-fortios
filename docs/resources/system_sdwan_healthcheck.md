---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_healthcheck"
description: |-
  SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.
---

## fortios_system_sdwan_healthcheck
SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.

~> This resource is configuring a child table of the parent resource: `fortios_system_sdwan`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

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
* `members` - Member sequence number list. The structure of `members` block is documented below.

The `members` block contains:

* `seq_num` - Member sequence number. This attribute must reference one of the following datasources: `system.sdwan.members.seq-num` .
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `sla` block contains:

* `id` - SLA ID.
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency` `jitter` `packet-loss` .
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_healthcheck can be imported using:
```sh
terraform import fortios_system_healthcheck.labelname {{mkey}}
```
