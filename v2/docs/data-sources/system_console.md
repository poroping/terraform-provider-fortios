---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_console"
description: |-
  Get information on a fortios Configure console.
---

# Data Source: fortios_system_console
Use this data source to get information on a fortios Configure console.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `baudrate` - Console baud rate.
* `fortiexplorer` - Enable/disable access for FortiExplorer.
* `login` - Enable/disable serial console and FortiExplorer.
* `mode` - Console mode.
* `output` - Console output mode.
