---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomexception"
description: |-
  Get information on a fortios Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.
---

# Data Source: fortios_system_vdomexception
Use this data source to get information on a fortios Global configuration objects that can be configured independently across different ha peers for all VDOMs or for the defined VDOM scope.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Index (1 - 4096).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Index (1 - 4096).
* `object` - Name of the configuration object that can be configured independently for all VDOMs.
* `scope` - Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.
* `vdom` - Names of the VDOMs.The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - VDOM name.
