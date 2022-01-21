---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ltemodem"
description: |-
  Get information on a fortios Configure USB LTE/WIMAX devices.
---

# Data Source: fortios_system_ltemodem
Use this data source to get information on a fortios Configure USB LTE/WIMAX devices.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `apn` - Login APN string for PDP-IP packet data calls.
* `authtype` - Authentication type for PDP-IP packet data calls.
* `extra_init` - Extra initialization string for USB LTE/WIMAX devices.
* `holddown_timer` - Hold down timer (10 - 60 sec).
* `interface` - The interface that the modem is acting as a redundant interface for.
* `mode` - Modem operation mode.
* `modem_port` - Modem port index (0 - 20).
* `passwd` - Authentication password for PDP-IP packet data calls.
* `status` - Enable/disable USB LTE/WIMAX device.
* `username` - Authentication username for PDP-IP packet data calls.
