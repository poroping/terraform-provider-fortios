---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy6"
description: |-
  Get information on a fortios Configure IPv6 routing policies.
---

# Data Source: fortios_router_policy6
Use this data source to get information on a fortios Configure IPv6 routing policies.


## Example Usage

```hcl

```

## Argument Reference

* `seq_num` - (Required) Sequence number(1-65535).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Optional comments.
* `dst` - Destination IPv6 prefix.
* `end_port` - End destination port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `output_device` - Outgoing interface name.
* `protocol` - Protocol number (0 - 255).
* `seq_num` - Sequence number(1-65535).
* `src` - Source IPv6 prefix.
* `start_port` - Start destination port number (1 - 65535).
* `status` - Enable/disable this policy route.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `input_device` - Incoming interface name.The structure of `input_device` block is documented below.

The `input_device` block contains:

* `name` - Interface name.
