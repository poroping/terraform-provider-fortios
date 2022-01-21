---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ptp"
description: |-
  Get information on a fortios Configure system PTP information.
---

# Data Source: fortios_system_ptp
Use this data source to get information on a fortios Configure system PTP information.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `delay_mechanism` - End to end delay detection or peer to peer delay detection.
* `interface` - PTP client will reply through this interface.
* `mode` - Multicast transmission or hybrid transmission.
* `request_interval` - The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
* `server_mode` - Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network.
* `status` - Enable/disable setting the FortiGate system time by synchronizing with an PTP Server.
* `server_interface` - FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services.The structure of `server_interface` block is documented below.

The `server_interface` block contains:

* `delay_mechanism` - End to end delay detection or peer to peer delay detection.
* `id` - ID.
* `server_interface_name` - Interface name.
