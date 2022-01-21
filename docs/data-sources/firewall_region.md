---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_region"
description: |-
  Get information on a fortios Define region table.
---

# Data Source: fortios_firewall_region
Use this data source to get information on a fortios Define region table.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Region ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Region ID.
* `name` - Region name.
* `city` - City ID list.The structure of `city` block is documented below.

The `city` block contains:

* `id` - City ID.
