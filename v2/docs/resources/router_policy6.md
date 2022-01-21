---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy6"
description: |-
  Configure IPv6 routing policies.
---

## fortios_router_policy6
Configure IPv6 routing policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `seq_num` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comments` - Optional comments.
* `dst` - Destination IPv6 prefix.
* `end_port` - End destination port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `output_device` - Outgoing interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `protocol` - Protocol number (0 - 255).
* `seq_num` - Sequence number(1-65535).
* `src` - Source IPv6 prefix.
* `start_port` - Start destination port number (1 - 65535).
* `status` - Enable/disable this policy route. Valid values: `enable` `disable` .
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `input_device` - Incoming interface name. The structure of `input_device` block is documented below.

The `input_device` block contains:

* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_policy6 can be imported using:
```sh
terraform import fortios_router_policy6.labelname {{mkey}}
```
