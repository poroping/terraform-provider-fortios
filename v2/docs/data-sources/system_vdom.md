---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdom"
description: |-
  Get information on a fortios Configure virtual domain.
---

# Data Source: fortios_system_vdom
Use this data source to get information on a fortios Configure virtual domain.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) VDOM name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `flag` - Flag.
* `name` - VDOM name.
* `short_name` - VDOM short name.
* `vcluster_id` - Virtual cluster ID (0 - 4294967295).
