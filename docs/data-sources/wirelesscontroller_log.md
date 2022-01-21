---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_log"
description: |-
  Get information on a fortios Configure wireless controller event log filters.
---

# Data Source: fortios_wirelesscontroller_log
Use this data source to get information on a fortios Configure wireless controller event log filters.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `addrgrp_log` - Lowest severity level to log address group message.
* `ble_log` - Lowest severity level to log BLE detection message.
* `clb_log` - Lowest severity level to log client load balancing message.
* `dhcp_starv_log` - Lowest severity level to log DHCP starvation event message.
* `led_sched_log` - Lowest severity level to log LED schedule event message.
* `radio_event_log` - Lowest severity level to log radio event message.
* `rogue_event_log` - Lowest severity level to log rogue AP event message.
* `sta_event_log` - Lowest severity level to log station event message.
* `sta_locate_log` - Lowest severity level to log station locate message.
* `status` - Enable/disable wireless event logging.
* `wids_log` - Lowest severity level to log WIDS message.
* `wtp_event_log` - Lowest severity level to log WTP event message.
