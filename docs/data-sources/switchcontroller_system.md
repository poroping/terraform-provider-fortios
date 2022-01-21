---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_system"
description: |-
  Get information on a fortios Configure system-wide switch controller settings.
---

# Data Source: fortios_switchcontroller_system
Use this data source to get information on a fortios Configure system-wide switch controller settings.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `data_sync_interval` - Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
* `dynamic_periodic_interval` - Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
* `iot_holdoff` - MAC entry's creation time. Time must be greater than this value for an entry to be created (0 - 10080 mins, default = 5 mins).
* `iot_mac_idle` - MAC entry's idle time. MAC entry is removed after this value (0 - 10080 mins, default = 1440 mins).
* `iot_scan_interval` - IoT scan interval (2 - 10080 mins, default = 60 mins, 0 = disable).
* `iot_weight_threshold` - MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
* `nac_periodic_interval` - Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
* `parallel_process` - Maximum number of parallel processes.
* `parallel_process_override` - Enable/disable parallel process override.
* `tunnel_mode` - Compatible/strict tunnel mode.
