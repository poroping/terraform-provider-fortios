---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis_isisnet"
description: |-
  Get information on a fortios IS-IS net configuration.
---

# Data Source: fortios_router_isis_isisnet
Use this data source to get information on a fortios IS-IS net configuration.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) isis-net ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - isis-net ID.
* `net` - IS-IS net xx.xxxx. ... .xxxx.xx.
