---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_static6"
description: |-
  Configure IPv6 static routing tables.
---

## fortios_router_static6
Configure IPv6 static routing tables.

## Example Usage

```hcl
resource "fortios_router_static6" "example" {
  dst       = "2001:b00b:fa57::/64"
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
* `devindex` - Device index (0 - 4294967295).
* `distance` - Administrative distance (1 - 255).
* `dst` - Destination IPv6 prefix.
* `dynamic_gateway` - Enable use of dynamic gateway retrieved from Router Advertisement (RA). Valid values: `enable` `disable` .
* `gateway` - IPv6 address of the gateway.
* `link_monitor_exempt` - Enable/disable withdrawal of this static route when link monitor or health check is down. Valid values: `enable` `disable` .
* `priority` - Administrative priority (1 - 65535).
* `sdwan` - Enable/disable egress through the SD-WAN. Valid values: `enable` `disable` .
* `seq_num` - Sequence number.
* `status` - Enable/disable this static route. Valid values: `enable` `disable` .
* `virtual_wan_link` - Enable/disable egress through the virtual-wan-link. Valid values: `enable` `disable` .
* `vrf` - Virtual Routing Forwarding ID.
* `sdwan_zone` - Choose SD-WAN Zone. The structure of `sdwan_zone` block is documented below.

The `sdwan_zone` block contains:

* `name` - SD-WAN zone name. This attribute must reference one of the following datasources: `system.sdwan.zone.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_static6 can be imported using:
```sh
terraform import fortios_router_static6.labelname {{mkey}}
```
