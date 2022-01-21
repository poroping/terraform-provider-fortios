---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_stp"
description: |-
  Get information on a fortios Configure Spanning Tree Protocol (STP).
---

# Data Source: fortios_system_stp
Use this data source to get information on a fortios Configure Spanning Tree Protocol (STP).


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `forward_delay` - Forward delay (4 - 30 sec, default = 15).
* `hello_time` - Hello time (1 - 10 sec, default = 2).
* `max_age` - Maximum packet age (6 - 40 sec, default = 20).
* `max_hops` - Maximum number of hops (1 - 40, default = 20).
* `switch_priority` - STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344).
