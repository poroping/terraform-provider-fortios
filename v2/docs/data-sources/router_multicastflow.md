---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicastflow"
description: |-
  Get information on a fortios Configure multicast-flow.
---

# Data Source: fortios_router_multicastflow
Use this data source to get information on a fortios Configure multicast-flow.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Comment.
* `name` - Name.
* `flows` - Multicast-flow entries.The structure of `flows` block is documented below.

The `flows` block contains:

* `group_addr` - Multicast group IP address.
* `id` - Flow ID.
* `source_addr` - Multicast source IP address.
