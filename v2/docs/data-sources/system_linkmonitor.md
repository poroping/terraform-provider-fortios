---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_linkmonitor"
description: |-
  Get information on a fortios Configure Link Health Monitor.
---

# Data Source: fortios_system_linkmonitor
Use this data source to get information on a fortios Configure Link Health Monitor.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Link monitor name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `addr_mode` - Address mode (IPv4 or IPv6).
* `class_id` - Traffic class ID.
* `diffservcode` - Differentiated services code point (DSCP) in the IP header of the probe packet.
* `fail_weight` - Threshold weight to trigger link failure alert.
* `failtime` - Number of retry attempts before the server is considered down (1 - 10, default = 5)
* `gateway_ip` - Gateway IP address used to probe the server.
* `gateway_ip6` - Gateway IPv6 address used to probe the server.
* `ha_priority` - HA election priority (1 - 50).
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_get` - If you are monitoring an HTML server you can send an HTTP-GET request with a custom string. Use this option to define the string.
* `http_match` - String that you expect to see in the HTTP-GET requests of the traffic to be monitored.
* `interval` - Detection interval in milliseconds (500 - 3600 * 1000 msec, default = 500).
* `name` - Link monitor name.
* `packet_size` - Packet size of a twamp test session,
* `password` - Twamp controller password in authentication mode
* `port` - Port number of the traffic to be used to monitor the server.
* `probe_count` - Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
* `probe_timeout` - Time to wait before a probe packet is considered lost (500 - 5000 msec, default = 500).
* `protocol` - Protocols used to monitor the server.
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 10, default = 5).
* `security_mode` - Twamp controller security mode.
* `server_config` - Mode of server configuration.
* `service_detection` - Only use monitor to read quality values. If enabled, static routes and cascade interfaces will not be updated.
* `source_ip` - Source IP address used in packet to the server.
* `source_ip6` - Source IPv6 address used in packet to the server.
* `srcintf` - Interface that receives the traffic to be monitored.
* `status` - Enable/disable this link monitor.
* `update_cascade_interface` - Enable/disable update cascade interface.
* `update_policy_route` - Enable/disable updating the policy route.
* `update_static_route` - Enable/disable updating the static route.
* `route` - Subnet to monitor.The structure of `route` block is documented below.

The `route` block contains:

* `subnet` - IP and netmask (x.x.x.x/y).
* `server` - IP address of the server(s) to be monitored.The structure of `server` block is documented below.

The `server` block contains:

* `address` - Server address.
* `server_list` - Servers for link-monitor to monitor.The structure of `server_list` block is documented below.

The `server_list` block contains:

* `dst` - IP address of the server to be monitored.
* `id` - Server ID.
* `port` - Port number of the traffic to be used to monitor the server.
* `protocol` - Protocols used to monitor the server.
* `weight` - Weight of the monitor to this dst (0 - 255).
