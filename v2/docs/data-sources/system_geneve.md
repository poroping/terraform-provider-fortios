---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geneve"
description: |-
  Get information on a fortios Configure GENEVE devices.
---

# Data Source: fortios_system_geneve
Use this data source to get information on a fortios Configure GENEVE devices.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) GENEVE device or interface name. Must be an unique interface name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `dstport` - GENEVE destination port (1 - 65535, default = 6081).
* `interface` - Outgoing interface for GENEVE encapsulated traffic.
* `ip_version` - IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast.
* `name` - GENEVE device or interface name. Must be an unique interface name.
* `remote_ip` - IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
* `remote_ip6` - IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
* `type` - GENEVE type.
* `vni` - GENEVE network ID.
