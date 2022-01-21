---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_address"
description: |-
  Configure the client with its MAC address.
---

## fortios_wirelesscontroller_address
Configure the client with its MAC address.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - ID.
* `mac` - MAC address.
* `policy` - Allow or block the client with this MAC address. Valid values: `allow` `deny` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_address can be imported using:
```sh
terraform import fortios_wirelesscontroller_address.labelname {{mkey}}
```
