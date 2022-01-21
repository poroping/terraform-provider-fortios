---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoscript"
description: |-
  Get information on a fortios Configure auto script.
---

# Data Source: fortios_system_autoscript
Use this data source to get information on a fortios Configure auto script.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Auto script name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `interval` - Repeat interval in seconds.
* `name` - Auto script name.
* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).
* `repeat` - Number of times to repeat this script (0 = infinite).
* `script` - List of FortiOS CLI commands to repeat.
* `start` - Script starting mode.
* `timeout` - Maximum running time for this script in seconds (0 = no timeout).
