---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vendormac"
description: |-
  Show vendor and the MAC address they have.
---

## fortios_firewall_vendormac
Show vendor and the MAC address they have.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - Vendor ID.
* `mac_number` - Total number of MAC addresses.
* `name` - Vendor name.
* `obsolete` - Indicates whether the Vendor ID can be used.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_vendormac can be imported using:
```sh
terraform import fortios_firewall_vendormac.labelname {{mkey}}
```
