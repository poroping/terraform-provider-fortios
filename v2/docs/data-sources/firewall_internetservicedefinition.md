---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicedefinition"
description: |-
  Get information on a fortios Configure Internet Service definition.
---

# Data Source: fortios_firewall_internetservicedefinition
Use this data source to get information on a fortios Configure Internet Service definition.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Internet Service application list ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `id` - Internet Service application list ID.
* `entry` - Protocol and port information in an Internet Service entry.The structure of `entry` block is documented below.

The `entry` block contains:

* `category_id` - Internet Service category ID.
* `name` - Internet Service name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `seq_num` - Entry sequence number.
* `port_range` - Port ranges in the definition entry.The structure of `port_range` block is documented below.

The `port_range` block contains:

* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
