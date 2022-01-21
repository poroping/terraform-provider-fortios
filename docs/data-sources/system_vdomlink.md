---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomlink"
description: |-
  Get information on a fortios Configure VDOM links.
---

# Data Source: fortios_system_vdomlink
Use this data source to get information on a fortios Configure VDOM links.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) VDOM link name (maximum = 11 characters).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - VDOM link name (maximum = 11 characters).
* `type` - VDOM link type: PPP or Ethernet.
* `vcluster` - Virtual cluster.
