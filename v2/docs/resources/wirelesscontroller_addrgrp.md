---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_addrgrp"
description: |-
  Configure the MAC address group.
---

## fortios_wirelesscontroller_addrgrp
Configure the MAC address group.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `default_policy` - Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow` `deny` .
* `id` - ID.
* `addresses` - Manually selected group of addresses. The structure of `addresses` block is documented below.

The `addresses` block contains:

* `id` - Address ID. This attribute must reference one of the following datasources: `wireless-controller.address.id` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_addrgrp can be imported using:
```sh
terraform import fortios_wirelesscontroller_addrgrp.labelname {{mkey}}
```
