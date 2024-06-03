---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static"
description: |-
  Configure IPv4 static routing tables.
---

## fortios_router_static
Configure IPv4 static routing tables.

## Example Usage

```hcl
resource "fortios_router_static" "example" {
  dst       = "1.1.1.50/32"
  blackhole = "enable"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `seq_num` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `enable` `disable` .
* `blackhole` - Enable/disable black hole. Valid values: `enable` `disable` .
* `comment` - Optional comments.
* `device` - Gateway out interface or tunnel. This attribute must reference one of the following datasources: `system.interface.name` .
* `distance` - Administrative distance (1 - 255).
* `dst` - Destination IP and mask for this route.
* `dstaddr` - Name of firewall address or address group. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from a DHCP or PPP server. Valid values: `enable` `disable` .
* `gateway` - Gateway IP for this route.
* `internet_service` - Application ID in the Internet service database. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `internet_service_custom` - Application name in the Internet service custom database. This attribute must reference one of the following datasources: `firewall.internet-service-custom.name` .
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable` `disable` .
* `priority` - Administrative priority (1 - 65535).
* `sdwan` - Enable/disable egress through SD-WAN. Valid values: `enable` `disable` .
* `seq_num` - Sequence number.
* `src` - Source prefix for this route.
* `status` - Enable/disable this static route. Valid values: `enable` `disable` .
* `tag` - Route tag.
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link. Valid values: `enable` `disable` .
* `vrf` - Virtual Routing Forwarding ID.
* `weight` - Administrative weight (0 - 255).
* `sdwan_zone` - Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.

The `sdwan_zone` block contains:

* `name` - SD-WAN zone name. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_static can be imported using:
```sh
terraform import fortios_router_static.labelname {{mkey}}
```
