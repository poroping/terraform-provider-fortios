---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipv6tunnel"
description: |-
  Get information on a fortios Configure IPv6/IPv4 in IPv6 tunnel.
---

# Data Source: fortios_system_ipv6tunnel
Use this data source to get information on a fortios Configure IPv6/IPv4 in IPv6 tunnel.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPv6 tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `auto_asic_offload` - Enable/disable tunnel ASIC offloading.
* `destination` - Remote IPv6 address of the tunnel.
* `interface` - Interface name.
* `name` - IPv6 tunnel name.
* `source` - Local IPv6 address of the tunnel.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway.
