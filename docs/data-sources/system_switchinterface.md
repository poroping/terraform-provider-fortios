---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_switchinterface"
description: |-
  Get information on a fortios Configure software switch interfaces by grouping physical and WiFi interfaces.
---

# Data Source: fortios_system_switchinterface
Use this data source to get information on a fortios Configure software switch interfaces by grouping physical and WiFi interfaces.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `intra_switch_policy` - Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces.
* `mac_ttl` - Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
* `name` - Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
* `span` - Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port.
* `span_dest_port` - SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
* `span_direction` - The direction in which the SPAN port operates, either: rx, tx, or both.
* `type` - Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members.
* `vdom` - VDOM that the software switch belongs to.
* `member` - Names of the interfaces that belong to the virtual switch.The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Physical interface name.
* `span_source_port` - Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port.The structure of `span_source_port` block is documented below.

The `span_source_port` block contains:

* `interface_name` - Physical interface name.
