---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_intercontroller"
description: |-
  Get information on a fortios Configure inter wireless controller operation.
---

# Data Source: fortios_wirelesscontroller_intercontroller
Use this data source to get information on a fortios Configure inter wireless controller operation.


## Example Usage

```hcl

```

## Argument Reference

* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fast_failover_max` - Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
* `fast_failover_wait` - Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
* `inter_controller_key` - Secret key for inter-controller communications.
* `inter_controller_mode` - Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable).
* `inter_controller_pri` - Configure inter-controller's priority (primary or secondary, default = primary).
* `l3_roaming` - Enable/disable layer 3 roaming (default = disable).
* `inter_controller_peer` - Fast failover peer wireless controller list.The structure of `inter_controller_peer` block is documented below.

The `inter_controller_peer` block contains:

* `id` - ID.
* `peer_ip` - Peer wireless controller's IP address.
* `peer_port` - Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
* `peer_priority` - Peer wireless controller's priority (primary or secondary, default = primary).
