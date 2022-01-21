---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_log"
description: |-
  Configure wireless controller event log filters.
---

## fortios_wirelesscontroller_log
Configure wireless controller event log filters.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `addrgrp_log` - Lowest severity level to log address group message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `ble_log` - Lowest severity level to log BLE detection message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `clb_log` - Lowest severity level to log client load balancing message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `dhcp_starv_log` - Lowest severity level to log DHCP starvation event message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `led_sched_log` - Lowest severity level to log LED schedule event message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `radio_event_log` - Lowest severity level to log radio event message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `rogue_event_log` - Lowest severity level to log rogue AP event message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `sta_event_log` - Lowest severity level to log station event message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `sta_locate_log` - Lowest severity level to log station locate message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `status` - Enable/disable wireless event logging. Valid values: `enable` `disable` .
* `wids_log` - Lowest severity level to log WIDS message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .
* `wtp_event_log` - Lowest severity level to log WTP event message. Valid values: `emergency` `alert` `critical` `error` `warning` `notification` `information` `debug` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_log can be imported using:
```sh
terraform import fortios_wirelesscontroller_log.labelname {{mkey}}
```
