---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_switchinterface"
description: |-
  Configure software switch interfaces by grouping physical and WiFi interfaces.
---

## fortios_system_switchinterface
Configure software switch interfaces by grouping physical and WiFi interfaces.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `intra_switch_policy` - Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit` `explicit` .
* `mac_ttl` - Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
* `name` - Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
* `span` - Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable` `enable` .
* `span_dest_port` - SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port. This attribute must reference one of the following datasources: `system.interface.name` .
* `span_direction` - The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx` `tx` `both` .
* `type` - Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch` `hub` .
* `vdom` - VDOM that the software switch belongs to. This attribute must reference one of the following datasources: `system.vdom.name` .
* `member` - Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Physical interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `span_source_port` - Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.

The `span_source_port` block contains:

* `interface_name` - Physical interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_switchinterface can be imported using:
```sh
terraform import fortios_system_switchinterface.labelname {{mkey}}
```
