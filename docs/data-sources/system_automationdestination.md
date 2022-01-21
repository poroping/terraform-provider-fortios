---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationdestination"
description: |-
  Get information on a fortios Automation destinations.
---

# Data Source: fortios_system_automationdestination
Use this data source to get information on a fortios Automation destinations.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `ha_group_id` - Cluster group ID set for this destination (default = 0).
* `name` - Name.
* `type` - Destination type.
* `destination` - Destinations.The structure of `destination` block is documented below.

The `destination` block contains:

* `name` - Destination.
