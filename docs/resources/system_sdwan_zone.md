---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdwan_zone"
description: |-
  Configure SD-WAN zones.
---

## fortios_system_sdwan_zone
Configure SD-WAN zones.

~> This resource is configuring a child table of the parent resource: `fortios_system_sdwan`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `cfg-order` `fib-best-match` `input-device` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_zone can be imported using:
```sh
terraform import fortios_system_zone.labelname {{mkey}}
```
