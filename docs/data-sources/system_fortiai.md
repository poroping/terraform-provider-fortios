---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortiai"
description: |-
  Get information on a fortios Configure FortiAI.
---

# Data Source: fortios_system_fortiai
Use this data source to get information on a fortios Configure FortiAI.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `source_ip` - Source IP address for communications to FortiAI.
* `status` - Enable/disable FortiAI.
