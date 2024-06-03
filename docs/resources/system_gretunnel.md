---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_gretunnel"
description: |-
  Configure GRE tunnel.
---

## fortios_system_gretunnel
Configure GRE tunnel.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `auto_asic_offload` - Enable/disable automatic ASIC offloading. Valid values: `enable` `disable` .
* `checksum_reception` - Enable/disable validating checksums in received GRE packets. Valid values: `disable` `enable` .
* `checksum_transmission` - Enable/disable including checksums in transmitted GRE packets. Valid values: `disable` `enable` .
* `diffservcode` - DiffServ setting to be applied to GRE tunnel outer IP header.
* `dscp_copying` - Enable/disable DSCP copying. Valid values: `disable` `enable` .
* `interface` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_version` - IP version to use for VPN interface. Valid values: `4` `6` .
* `keepalive_failtimes` - Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
* `keepalive_interval` - Keepalive message interval (0 - 32767, 0 = disabled).
* `key_inbound` - Require received GRE packets contain this key (0 - 4294967295).
* `key_outbound` - Include this key in transmitted GRE packets (0 - 4294967295).
* `local_gw` - IP address of the local gateway.
* `local_gw6` - IPv6 address of the local gateway.
* `name` - Tunnel name.
* `remote_gw` - IP address of the remote gateway.
* `remote_gw6` - IPv6 address of the remote gateway.
* `sequence_number_reception` - Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable` `enable` .
* `sequence_number_transmission` - Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable` `enable` .
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable` `enable` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_gretunnel can be imported using:
```sh
terraform import fortios_system_gretunnel.labelname {{mkey}}
```
