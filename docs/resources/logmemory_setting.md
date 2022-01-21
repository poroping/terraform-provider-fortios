---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logmemory_setting"
description: |-
  Settings for memory buffer.
---

## fortios_logmemory_setting
Settings for memory buffer.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `diskfull` - Action to take when memory is full. Valid values: `overwrite` .
* `status` - Enable/disable logging to the FortiGate's memory. Valid values: `enable` `disable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_logmemory_setting can be imported using:
```sh
terraform import fortios_logmemory_setting.labelname {{mkey}}
```
