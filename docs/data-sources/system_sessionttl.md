---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionttl"
description: |-
  Get information on a fortios Configure global session TTL timers for this FortiGate.
---

# Data Source: fortios_system_sessionttl
Use this data source to get information on a fortios Configure global session TTL timers for this FortiGate.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `default` - Default timeout.
* `port` - Session TTL port.The structure of `port` block is documented below.

The `port` block contains:

* `end_port` - End port number.
* `id` - Table entry ID.
* `protocol` - Protocol (0 - 255).
* `refresh_direction` - Refresh direction: Both, outgoing, incoming
* `start_port` - Start port number.
* `timeout` - Session timeout (TTL).
