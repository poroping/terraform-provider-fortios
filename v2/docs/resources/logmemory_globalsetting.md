---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_globalsetting"
description: |-
  Global settings for memory logging.
---

## fortios_logmemory_globalsetting
Global settings for memory logging.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `max_size` - Maximum amount of memory that can be used for memory logging in bytes.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logmemory_globalsetting can be imported using:
```sh
terraform import fortios_logmemory_globalsetting.labelname {{mkey}}
```
