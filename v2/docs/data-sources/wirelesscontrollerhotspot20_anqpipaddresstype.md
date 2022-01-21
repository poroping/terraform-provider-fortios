---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpipaddresstype"
description: |-
  Get information on a fortios Configure IP address type availability.
---

# Data Source: fortios_wirelesscontrollerhotspot20_anqpipaddresstype
Use this data source to get information on a fortios Configure IP address type availability.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IP type name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ipv4_address_type` - IPv4 address type.
* `ipv6_address_type` - IPv6 address type.
* `name` - IP type name.
