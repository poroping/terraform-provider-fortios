---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicelist"
description: |-
  Get information on a fortios Internet Service list.
---

# Data Source: fortios_firewall_internetservicelist
Use this data source to get information on a fortios Internet Service list.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Internet Service category ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Internet Service category ID.
* `name` - Internet Service category name.
