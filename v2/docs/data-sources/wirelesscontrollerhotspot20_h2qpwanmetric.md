---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpwanmetric"
description: |-
  Get information on a fortios Configure WAN metrics.
---

# Data Source: fortios_wirelesscontrollerhotspot20_h2qpwanmetric
Use this data source to get information on a fortios Configure WAN metrics.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) WAN metric name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `downlink_load` - Downlink load.
* `downlink_speed` - Downlink speed (in kilobits/s).
* `link_at_capacity` - Link at capacity.
* `link_status` - Link status.
* `load_measurement_duration` - Load measurement duration (in tenths of a second).
* `name` - WAN metric name.
* `symmetric_wan_link` - WAN link symmetry.
* `uplink_load` - Uplink load.
* `uplink_speed` - Uplink speed (in kilobits/s).
