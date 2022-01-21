---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ltemodem"
description: |-
  Configure USB LTE/WIMAX devices.
---

## fortios_system_ltemodem
Configure USB LTE/WIMAX devices.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `apn` - Login APN string for PDP-IP packet data calls.
* `authtype` - Authentication type for PDP-IP packet data calls. Valid values: `none` `pap` `chap` .
* `extra_init` - Extra initialization string for USB LTE/WIMAX devices.
* `holddown_timer` - Hold down timer (10 - 60 sec).
* `interface` - The interface that the modem is acting as a redundant interface for. This attribute must reference one of the following datasources: `system.interface.name` .
* `mode` - Modem operation mode. Valid values: `standalone` `redundant` .
* `modem_port` - Modem port index (0 - 20).
* `passwd` - Authentication password for PDP-IP packet data calls.
* `status` - Enable/disable USB LTE/WIMAX device. Valid values: `enable` `disable` .
* `username` - Authentication username for PDP-IP packet data calls.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ltemodem can be imported using:
```sh
terraform import fortios_system_ltemodem.labelname {{mkey}}
```
