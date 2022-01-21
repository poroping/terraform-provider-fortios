---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_globalsetting"
description: |-
  Get information on a fortios Global settings for memory logging.
---

# Data Source: fortios_logmemory_globalsetting
Use this data source to get information on a fortios Global settings for memory logging.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `max_size` - Maximum amount of memory that can be used for memory logging in bytes.
