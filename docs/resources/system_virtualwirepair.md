---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualwirepair"
description: |-
  Configure virtual wire pairs.
---

## fortios_system_virtualwirepair
Configure virtual wire pairs.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `name` - Virtual-wire-pair name. Must be a unique interface name.
* `vlan_filter` - VLAN ranges to allow
* `wildcard_vlan` - Enable/disable wildcard VLAN. Valid values: `enable` `disable` .
* `member` - Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `outer_vlan_id` - Outer VLAN ID. The structure of `outer_vlan_id` block is documented below.

The `outer_vlan_id` block contains:

* `vlanid` - VLAN ID (1 - 4094).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_virtualwirepair can be imported using:
```sh
terraform import fortios_system_virtualwirepair.labelname {{mkey}}
```
