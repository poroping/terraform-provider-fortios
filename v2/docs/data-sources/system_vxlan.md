---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vxlan"
description: |-
  Get information on a fortios Configure VXLAN devices.
---

# Data Source: fortios_system_vxlan
Use this data source to get information on a fortios Configure VXLAN devices.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) VXLAN device or interface name. Must be a unique interface name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `interface` - Outgoing interface for VXLAN encapsulated traffic.
* `ip_version` - IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `name` - VXLAN device or interface name. Must be a unique interface name.
* `vni` - VXLAN network ID.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN.The structure of `remote_ip` block is documented below.

The `remote_ip` block contains:

* `ip` - IPv4 address.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN.The structure of `remote_ip6` block is documented below.

The `remote_ip6` block contains:

* `ip6` - IPv6 address.
