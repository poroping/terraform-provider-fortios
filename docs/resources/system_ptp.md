---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ptp"
description: |-
  Configure system PTP information.
---

## fortios_system_ptp
Configure system PTP information.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E` `P2P` .
* `interface` - PTP client will reply through this interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `mode` - Multicast transmission or hybrid transmission. Valid values: `multicast` `hybrid` .
* `request_interval` - The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
* `server_mode` - Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable` `disable` .
* `status` - Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable` `disable` .
* `server_interface` - FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.

The `server_interface` block contains:

* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E` `P2P` .
* `id` - ID.
* `server_interface_name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ptp can be imported using:
```sh
terraform import fortios_system_ptp.labelname {{mkey}}
```
