---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_gretunnel"
description: |-
  Get information on a fortios Configure GRE tunnel.
---

# Data Source: fortios_system_gretunnel
Use this data source to get information on a fortios Configure GRE tunnel.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Tunnel name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `checksum_reception` - Enable/disable validating checksums in received GRE packets.
* `checksum_transmission` - Enable/disable including checksums in transmitted GRE packets.
* `diffservcode` - DiffServ setting to be applied to GRE tunnel outer IP header.
* `dscp_copying` - Enable/disable DSCP copying.
* `interface` - Interface name.
* `ip_version` - IP version to use for VPN interface.
* `keepalive_failtimes` - Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
* `keepalive_interval` - Keepalive message interval (0 - 32767, 0 = disabled).
* `key_inbound` - Require received GRE packets contain this key (0 - 4294967295).
* `key_outbound` - Include this key in transmitted GRE packets (0 - 4294967295).
* `local_gw` - IP address of the local gateway.
* `local_gw6` - IPv6 address of the local gateway.
* `name` - Tunnel name.
* `remote_gw` - IP address of the remote gateway.
* `remote_gw6` - IPv6 address of the remote gateway.
* `sequence_number_reception` - Enable/disable validating sequence numbers in received GRE packets.
* `sequence_number_transmission` - Enable/disable including of sequence numbers in transmitted GRE packets.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway.
