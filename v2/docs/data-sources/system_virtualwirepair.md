---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualwirepair"
description: |-
  Get information on a fortios Configure virtual wire pairs.
---

# Data Source: fortios_system_virtualwirepair
Use this data source to get information on a fortios Configure virtual wire pairs.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Virtual-wire-pair name. Must be a unique interface name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Virtual-wire-pair name. Must be a unique interface name.
* `vlan_filter` - Set VLAN filters.
* `wildcard_vlan` - Enable/disable wildcard VLAN.
* `member` - Interfaces belong to the virtual-wire-pair.The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Interface name.
