---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_rip_redistribute"
description: |-
  Get information on a fortios Redistribute configuration.
---

# Data Source: fortios_router_rip_redistribute
Use this data source to get information on a fortios Redistribute configuration.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Redistribute name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `metric` - Redistribute metric setting.
* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status.
