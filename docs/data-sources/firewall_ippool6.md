---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ippool6"
description: |-
  Get information on a fortios Configure IPv6 IP pools.
---

# Data Source: fortios_firewall_ippool6
Use this data source to get information on a fortios Configure IPv6 IP pools.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPv6 IP pool name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `add_nat46_route` - Enable/disable adding NAT46 route.
* `comments` - Comment.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, default = ::).
* `name` - IPv6 IP pool name.
* `nat46` - Enable/disable NAT46.
* `startip` - First IPv6 address (inclusive) in the range for the address pool (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, default = ::).
