---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static6"
description: |-
  Get information on a fortios Configure IPv6 static routing tables.
---

# Data Source: fortios_router_static6
Use this data source to get information on a fortios Configure IPv6 static routing tables.


## Example Usage

```hcl

```

## Argument Reference

* `seq_num` - (Required) Sequence number.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD).
* `blackhole` - Enable/disable black hole.
* `comment` - Optional comments.
* `device` - Gateway out interface or tunnel.
* `devindex` - Device index (0 - 4294967295).
* `distance` - Administrative distance (1 - 255).
* `dst` - Destination IPv6 prefix.
* `dstaddr` - Name of firewall address or address group.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from Router Advertisement (RA).
* `gateway` - IPv6 address of the gateway.
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down.
* `priority` - Administrative priority (1 - 65535).
* `sdwan` - Enable/disable egress through the SD-WAN.
* `seq_num` - Sequence number.
* `status` - Enable/disable this static route.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link.
* `vrf` - Virtual Routing Forwarding ID.
* `weight` - Administrative weight (0 - 255).
* `sdwan_zone` - Choose SD-WAN Zone.The structure of `sdwan_zone` block is documented below.

The `sdwan_zone` block contains:

* `name` - SD-WAN zone name.
