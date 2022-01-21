---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qpwanmetric"
description: |-
  Configure WAN metrics.
---

## fortios_wirelesscontrollerhotspot20_h2qpwanmetric
Configure WAN metrics.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `downlink_load` - Downlink load.
* `downlink_speed` - Downlink speed (in kilobits/s).
* `link_at_capacity` - Link at capacity. Valid values: `enable` `disable` .
* `link_status` - Link status. Valid values: `up` `down` `in-test` .
* `load_measurement_duration` - Load measurement duration (in tenths of a second).
* `name` - WAN metric name.
* `symmetric_wan_link` - WAN link symmetry. Valid values: `symmetric` `asymmetric` .
* `uplink_load` - Uplink load.
* `uplink_speed` - Uplink speed (in kilobits/s).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontrollerhotspot20_h2qpwanmetric can be imported using:
```sh
terraform import fortios_wirelesscontrollerhotspot20_h2qpwanmetric.labelname {{mkey}}
```
