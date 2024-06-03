---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestschedule"
description: |-
  Speed test schedule for each interface.
---

## fortios_system_speedtestschedule
Speed test schedule for each interface.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `interface` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `diffserv` - DSCP used for speed test.
* `dynamic_server` - Enable/disable dynamic server option. Valid values: `disable` `enable` .
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `mode` - Protocol Auto(default), TCP or UDP used for speed test. Valid values: `UDP` `TCP` `Auto` .
* `server_name` - Speed test server name.
* `status` - Enable/disable scheduled speed test. Valid values: `disable` `enable` .
* `update_inbandwidth` - Enable/disable bypassing interface's inbound bandwidth setting. Valid values: `disable` `enable` .
* `update_inbandwidth_maximum` - Maximum downloading bandwidth (kbps) to be used in a speed test.
* `update_inbandwidth_minimum` - Minimum downloading bandwidth (kbps) to be considered effective.
* `update_outbandwidth` - Enable/disable bypassing interface's outbound bandwidth setting. Valid values: `disable` `enable` .
* `update_outbandwidth_maximum` - Maximum uploading bandwidth (kbps) to be used in a speed test.
* `update_outbandwidth_minimum` - Minimum uploading bandwidth (kbps) to be considered effective.
* `schedules` - Schedules for the interface. The structure of `schedules` block is documented below.

The `schedules` block contains:

* `name` - Name of a firewall recurring schedule. This attribute must reference one of the following datasources: `firewall.schedule.recurring.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_speedtestschedule can be imported using:
```sh
terraform import fortios_system_speedtestschedule.labelname {{mkey}}
```
