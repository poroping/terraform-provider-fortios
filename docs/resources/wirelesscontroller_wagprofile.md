---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wagprofile"
description: |-
  Configure wireless access gateway (WAG) profiles used for tunnels on AP.
---

## fortios_wirelesscontroller_wagprofile
Configure wireless access gateway (WAG) profiles used for tunnels on AP.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `comment` - Comment.
* `dhcp_ip_addr` - IP address of the monitoring DHCP request packet sent through the tunnel
* `name` - Tunnel profile name.
* `ping_interval` - Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).
* `ping_number` - Number of the tunnel mointoring echo packets (1 - 65535, default = 5).
* `return_packet_timeout` - Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).
* `tunnel_type` - Tunnel type. Valid values: `l2tpv3` `gre` .
* `wag_ip` - IP Address of the wireless access gateway.
* `wag_port` - UDP port of the wireless access gateway.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_wagprofile can be imported using:
```sh
terraform import fortios_wirelesscontroller_wagprofile.labelname {{mkey}}
```
