---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_timers"
description: |-
  Configure CAPWAP timers.
---

## fortios_wirelesscontroller_timers
Configure CAPWAP timers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `ble_scan_report_intv` - Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
* `client_idle_timeout` - Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
* `discovery_interval` - Time between discovery requests (2 - 180 sec, default = 5).
* `drma_interval` - Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
* `echo_interval` - Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
* `fake_ap_log` - Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
* `ipsec_intf_cleanup` - Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
* `radio_stats_interval` - Time between running radio reports (1 - 255 sec, default = 15).
* `rogue_ap_log` - Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
* `sta_capability_interval` - Time between running station capability reports (1 - 255 sec, default = 30).
* `sta_locate_timer` - Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
* `sta_stats_interval` - Time between running client (station) reports (1 - 255 sec, default = 1).
* `vap_stats_interval` - Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_timers can be imported using:
```sh
terraform import fortios_wirelesscontroller_timers.labelname {{mkey}}
```
