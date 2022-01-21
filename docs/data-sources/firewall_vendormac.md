---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vendormac"
description: |-
  Get information on a fortios Show vendor and the MAC address they have.
---

# Data Source: fortios_firewall_vendormac
Use this data source to get information on a fortios Show vendor and the MAC address they have.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Vendor ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Vendor ID.
* `mac_number` - Total number of MAC addresses.
* `name` - Vendor name.
* `obsolete` - Indicates whether the Vendor ID can be used.
