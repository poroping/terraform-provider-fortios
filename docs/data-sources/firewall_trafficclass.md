---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_trafficclass"
description: |-
  Get information on a fortios Configure names for shaping classes.
---

# Data Source: fortios_firewall_trafficclass
Use this data source to get information on a fortios Configure names for shaping classes.


## Example Usage

```hcl

```

## Argument Reference

* `class_id` - (Required) Class ID to be named.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `class_id` - Class ID to be named.
* `class_name` - Define the name for this class-id.
