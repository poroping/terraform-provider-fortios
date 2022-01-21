---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoscript"
description: |-
  Configure auto script.
---

## fortios_system_autoscript
Configure auto script.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `interval` - Repeat interval in seconds.
* `name` - Auto script name.
* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).
* `repeat` - Number of times to repeat this script (0 = infinite).
* `script` - List of FortiOS CLI commands to repeat.
* `start` - Script starting mode. Valid values: `manual` `auto` .
* `timeout` - Maximum running time for this script in seconds (0 = no timeout).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_autoscript can be imported using:
```sh
terraform import fortios_system_autoscript.labelname {{mkey}}
```
