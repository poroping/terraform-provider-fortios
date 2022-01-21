---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vxlan"
description: |-
  Configure VXLAN devices.
---

## fortios_system_vxlan
Configure VXLAN devices.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `interface` - Outgoing interface for VXLAN encapsulated traffic. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_version` - IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast` `ipv6-unicast` `ipv4-multicast` `ipv6-multicast` .
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `name` - VXLAN device or interface name. Must be a unique interface name.
* `vni` - VXLAN network ID.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.

The `remote_ip` block contains:

* `ip` - IPv4 address.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.

The `remote_ip6` block contains:

* `ip6` - IPv6 address.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_vxlan can be imported using:
```sh
terraform import fortios_system_vxlan.labelname {{mkey}}
```
