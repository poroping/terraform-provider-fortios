---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionttl"
description: |-
  Configure global session TTL timers for this FortiGate.
---

## fortios_system_sessionttl
Configure global session TTL timers for this FortiGate.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `default` - Default timeout.
* `port` - Session TTL port. The structure of `port` block is documented below.

The `port` block contains:

* `end_port` - End port number.
* `id` - Table entry ID.
* `protocol` - Protocol (0 - 255).
* `refresh_direction` - Refresh direction: Both, outgoing, incoming Valid values: `both` `outgoing` `incoming` .
* `start_port` - Start port number.
* `timeout` - Session timeout (TTL).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_sessionttl can be imported using:
```sh
terraform import fortios_system_sessionttl.labelname {{mkey}}
```
