---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_country"
description: |-
  Get information on a fortios Define country table.
---

# Data Source: fortios_firewall_country
Use this data source to get information on a fortios Define country table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Country ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Country ID.
* `name` - Country name.
* `region` - Region ID list.The structure of `region` block is documented below.

The `region` block contains:

* `id` - Region ID.
