---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipv6neighborcache"
description: |-
  Get information on a fortios Configure IPv6 neighbor cache table.
---

# Data Source: fortios_system_ipv6neighborcache
Use this data source to get information on a fortios Configure IPv6 neighbor cache table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Unique integer ID of the entry.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Unique integer ID of the entry.
* `interface` - Select the associated interface name from available options.
* `ipv6` - IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `mac` - MAC address (format: xx:xx:xx:xx:xx:xx).
