---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_hamonitor"
description: |-
  Configure HA monitor.
---

## fortios_system_hamonitor
Configure HA monitor.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `monitor_vlan` - Enable/disable monitor VLAN interfaces. Valid values: `enable` `disable` .
* `vlan_hb_interval` - Configure heartbeat interval (seconds).
* `vlan_hb_lost_threshold` - VLAN lost heartbeat threshold (1 - 60).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_hamonitor can be imported using:
```sh
terraform import fortios_system_hamonitor.labelname {{mkey}}
```
