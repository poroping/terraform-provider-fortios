---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualswitch"
description: |-
  Configure virtual hardware switch interfaces.
---

## fortios_system_virtualswitch
Configure virtual hardware switch interfaces.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Name of the virtual switch.
* `physical_switch` - Physical switch parent. This attribute must reference one of the following datasources: `system.physical-switch.name` .
* `span` - Enable/disable SPAN. Valid values: `disable` `enable` .
* `span_dest_port` - SPAN destination port.
* `span_direction` - SPAN direction. Valid values: `rx` `tx` `both` .
* `span_source_port` - SPAN source port.
* `vlan` - VLAN.
* `port` - Configure member ports. The structure of `port` block is documented below.

The `port` block contains:

* `alias` - Alias.
* `name` - Physical interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `speed` - Interface speed. Valid values: `auto` `10full` `10half` `100full` `100half` `1000full` `1000half` `1000auto` `10000full` `10000auto` `40000full` `25000full` .
* `status` - Interface status. Valid values: `up` `down` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_virtualswitch can be imported using:
```sh
terraform import fortios_system_virtualswitch.labelname {{mkey}}
```
