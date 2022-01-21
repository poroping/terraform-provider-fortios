---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpipaddresstype"
description: |-
  Configure IP address type availability.
---

## fortios_wirelesscontrollerhotspot20_anqpipaddresstype
Configure IP address type availability.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `ipv4_address_type` - IPv4 address type. Valid values: `not-available` `public` `port-restricted` `single-NATed-private` `double-NATed-private` `port-restricted-and-single-NATed` `port-restricted-and-double-NATed` `not-known` .
* `ipv6_address_type` - IPv6 address type. Valid values: `not-available` `available` `not-known` .
* `name` - IP type name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_anqpipaddresstype can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_anqpipaddresstype.labelname {{mkey}}
```
