---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_stpsettings"
description: |-
  Get information on a fortios Configure FortiSwitch spanning tree protocol (STP).
---

# Data Source: fortios_switchcontroller_stpsettings
Use this data source to get information on a fortios Configure FortiSwitch spanning tree protocol (STP).


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `forward_time` - Period of time a port is in listening and learning state (4 - 30 sec, default = 15).
* `hello_time` - Period of time between successive STP frame Bridge Protocol Data Units (BPDUs) sent on a port (1 - 10 sec, default = 2).
* `max_age` - Maximum time before a bridge port expires its configuration BPDU information (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops between the root bridge and the furthest bridge (1- 40, default = 20).
* `name` - Name of global STP settings configuration.
* `pending_timer` - Pending time (1 - 15 sec, default = 4).
* `revision` - STP revision number (0 - 65535).
