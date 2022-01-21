---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ippool6"
description: |-
  Configure IPv6 IP pools.
---

## fortios_firewall_ippool6
Configure IPv6 IP pools.

## Example Usage

```hcl
resource "fortios_firewall_ippool6" "example" {
  name    = "foobar"
  startip = "2001::1"
  endip   = "2001::3"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable` `enable` .
* `comments` - Comment.
* `endip` - Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `name` - IPv6 IP pool name.
* `nat46` - Enable/disable NAT46. Valid values: `disable` `enable` .
* `startip` - First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_ippool6 can be imported using:
```sh
terraform import fortios_firewall_ippool6.labelname {{mkey}}
```
