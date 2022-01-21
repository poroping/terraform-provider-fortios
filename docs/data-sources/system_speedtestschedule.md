---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestschedule"
description: |-
  Get information on a fortios Speed test schedule for each interface.
---

# Data Source: fortios_system_speedtestschedule
Use this data source to get information on a fortios Speed test schedule for each interface.


## Example Usage

```hcl

```

## Argument Reference

* `interface` - (Required) Interface name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `diffserv` - DSCP used for speed test.
* `dynamic_server` - Enable/disable dynamic server option.
* `interface` - Interface name.
* `server_name` - Speed test server name.
* `status` - Enable/disable scheduled speed test.
* `update_inbandwidth` - Enable/disable bypassing interface's inbound bandwidth setting.
* `update_inbandwidth_maximum` - Maximum downloading bandwidth (kbps) to be used in a speed test.
* `update_inbandwidth_minimum` - Minimum downloading bandwidth (kbps) to be considered effective.
* `update_outbandwidth` - Enable/disable bypassing interface's outbound bandwidth setting.
* `update_outbandwidth_maximum` - Maximum uploading bandwidth (kbps) to be used in a speed test.
* `update_outbandwidth_minimum` - Minimum uploading bandwidth (kbps) to be considered effective.
* `schedules` - Schedules for the interface.The structure of `schedules` block is documented below.

The `schedules` block contains:

* `name` - Name of a firewall recurring schedule.
