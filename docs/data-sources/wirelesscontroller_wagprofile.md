---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wagprofile"
description: |-
  Get information on a fortios Configure wireless access gateway (WAG) profiles used for tunnels on AP.
---

# Data Source: fortios_wirelesscontroller_wagprofile
Use this data source to get information on a fortios Configure wireless access gateway (WAG) profiles used for tunnels on AP.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Tunnel profile name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `dhcp_ip_addr` - IP address of the monitoring DHCP request packet sent through the tunnel.
* `name` - Tunnel profile name.
* `ping_interval` - Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
* `ping_number` - Number of the tunnel monitoring echo packets (1 - 65535, default = 5).
* `return_packet_timeout` - Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
* `tunnel_type` - Tunnel type.
* `wag_ip` - IP Address of the wireless access gateway.
* `wag_port` - UDP port of the wireless access gateway.
