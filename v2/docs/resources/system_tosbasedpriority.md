---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_tosbasedpriority"
description: |-
  Configure Type of Service (ToS) based priority table to set network traffic priorities.
---

## fortios_system_tosbasedpriority
Configure Type of Service (ToS) based priority table to set network traffic priorities.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `id` - Item ID.
* `priority` - ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low` `medium` `high` .
* `tos` - Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_tosbasedpriority can be imported using:
```sh
terraform import fortios_system_tosbasedpriority.labelname {{mkey}}
```
