---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualswitch"
description: |-
  Get information on a fortios Configure virtual hardware switch interfaces.
---

# Data Source: fortios_system_virtualswitch
Use this data source to get information on a fortios Configure virtual hardware switch interfaces.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name of the virtual switch.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `name` - Name of the virtual switch.
* `physical_switch` - Physical switch parent.
* `span` - Enable/disable SPAN.
* `span_dest_port` - SPAN destination port.
* `span_direction` - SPAN direction.
* `span_source_port` - SPAN source port.
* `vlan` - VLAN.
* `port` - Configure member ports.The structure of `port` block is documented below.

The `port` block contains:

* `alias` - Alias.
* `name` - Physical interface name.
* `speed` - Interface speed.
* `status` - Interface status.
