---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_schedule"
description: |-
  Configure update schedule.
---

## fortios_systemautoupdate_schedule
Configure update schedule.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `day` - Update day. Valid values: `Sunday` `Monday` `Tuesday` `Wednesday` `Thursday` `Friday` `Saturday` .
* `frequency` - Update frequency. Valid values: `every` `daily` `weekly` `automatic` .
* `status` - Enable/disable scheduled updates. Valid values: `enable` `disable` .
* `time` - Update time.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemautoupdate_schedule can be imported using:
```sh
terraform import fortios_systemautoupdate_schedule.labelname {{mkey}}
```
