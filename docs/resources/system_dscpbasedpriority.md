---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dscpbasedpriority"
description: |-
  Configure DSCP based priority table.
---

## fortios_system_dscpbasedpriority
Configure DSCP based priority table.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `ds` - DSCP(DiffServ) DS value (0 - 63).
* `id` - Item ID.
* `priority` - DSCP based priority level. Valid values: `low` `medium` `high` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_dscpbasedpriority can be imported using:
```sh
terraform import fortios_system_dscpbasedpriority.labelname {{mkey}}
```
