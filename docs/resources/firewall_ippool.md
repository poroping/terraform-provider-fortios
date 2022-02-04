---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ippool"
description: |-
  Configure IPv4 IP pools.
---

## fortios_firewall_ippool
Configure IPv4 IP pools.

## Example Usage

```hcl
resource "fortios_firewall_ippool" "example" {
  name    = "foobar"
  startip = "1.1.1.1"
  endip   = "1.1.1.2"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable` `enable` .
* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected). This attribute must reference one of the following datasources: `system.interface.name` .
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable` `enable` .
* `associated_interface` - Associated interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `block_size` - Number of addresses in a block (64 - 4096, default = 128).
* `comments` - Comment.
* `endip` - Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endport` - Final port number (inclusive) in the range for the address pool (Default: 65533).
* `name` - IP pool name.
* `nat64` - Enable/disable NAT64. Valid values: `disable` `enable` .
* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable` `enable` .
* `port_per_user` - Number of port for each user (32 - 60416, default = 0, which is auto).
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_startip` - First IPv4 address (inclusive) in the range of the source addresses to be translated (format = xxx.xxx.xxx.xxx, default = 0.0.0.0).
* `startip` - First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - First port number (inclusive) in the range for the address pool (Default: 5117).
* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload` `one-to-one` `fixed-port-range` `port-block-allocation` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_ippool can be imported using:
```sh
terraform import fortios_firewall_ippool.labelname {{mkey}}
```
