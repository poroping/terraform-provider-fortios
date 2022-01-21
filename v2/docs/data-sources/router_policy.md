---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy"
description: |-
  Get information on a fortios Configure IPv4 routing policies.
---

# Data Source: fortios_router_policy
Use this data source to get information on a fortios Configure IPv4 routing policies.


## Example Usage

```hcl

```

## Argument Reference

* `seq_num` - (Required) Sequence number(1-65535).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Action of the policy route.
* `comments` - Optional comments.
* `dst_negate` - Enable/disable negating destination address match.
* `end_port` - End destination port number (0 - 65535).
* `end_source_port` - End source port number (0 - 65535).
* `gateway` - IP address of the gateway.
* `input_device_negate` - Enable/disable negation of input device match.
* `output_device` - Outgoing interface name.
* `protocol` - Protocol number (0 - 255).
* `seq_num` - Sequence number(1-65535).
* `src_negate` - Enable/disable negating source address match.
* `start_port` - Start destination port number (0 - 65535).
* `start_source_port` - Start source port number (0 - 65535).
* `status` - Enable/disable this policy route.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `dst` - Destination IP and mask (x.x.x.x/x).The structure of `dst` block is documented below.

The `dst` block contains:

* `subnet` - IP and mask.
* `dstaddr` - Destination address name.The structure of `dstaddr` block is documented below.

The `dstaddr` block contains:

* `name` - Address/group name.
* `input_device` - Incoming interface name.The structure of `input_device` block is documented below.

The `input_device` block contains:

* `name` - Interface name.
* `internet_service_custom` - Custom Destination Internet Service name.The structure of `internet_service_custom` block is documented below.

The `internet_service_custom` block contains:

* `name` - Custom Destination Internet Service name.
* `internet_service_id` - Destination Internet Service ID.The structure of `internet_service_id` block is documented below.

The `internet_service_id` block contains:

* `id` - Destination Internet Service ID.
* `src` - Source IP and mask (x.x.x.x/x).The structure of `src` block is documented below.

The `src` block contains:

* `subnet` - IP and mask.
* `srcaddr` - Source address name.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address/group name.
