---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipv6neighborcache"
description: |-
  Configure IPv6 neighbor cache table.
---

## fortios_system_ipv6neighborcache
Configure IPv6 neighbor cache table.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - Unique integer ID of the entry.
* `interface` - Select the associated interface name from available options. This attribute must reference one of the following datasources: `system.interface.name` .
* `ipv6` - IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `mac` - MAC address (format: xx:xx:xx:xx:xx:xx).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ipv6neighborcache can be imported using:
```sh
terraform import fortios_system_ipv6neighborcache.labelname {{mkey}}
```
