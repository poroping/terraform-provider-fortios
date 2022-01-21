---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_intercontroller"
description: |-
  Configure inter wireless controller operation.
---

## fortios_wirelesscontroller_intercontroller
Configure inter wireless controller operation.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `fast_failover_max` - Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
* `fast_failover_wait` - Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
* `inter_controller_key` - Secret key for inter-controller communications.
* `inter_controller_mode` - Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable` `l2-roaming` `1+1` .
* `inter_controller_pri` - Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary` `secondary` .
* `inter_controller_peer` - Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.

The `inter_controller_peer` block contains:

* `id` - ID.
* `peer_ip` - Peer wireless controller's IP address.
* `peer_port` - Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
* `peer_priority` - Peer wireless controller's priority (primary or secondary, default = primary). Valid values: `primary` `secondary` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_wirelesscontroller_intercontroller can be imported using:
```sh
terraform import fortios_wirelesscontroller_intercontroller.labelname {{mkey}}
```
