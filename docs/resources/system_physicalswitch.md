---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_physicalswitch"
description: |-
  Configure physical switches.
---

## fortios_system_physicalswitch
Configure physical switches.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `age_enable` - Enable/disable layer 2 age timer. Valid values: `enable` `disable` .
* `age_val` - Layer 2 table age timer value.
* `name` - Name.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_physicalswitch can be imported using:
```sh
terraform import fortios_system_physicalswitch.labelname {{mkey}}
```
