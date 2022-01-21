---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geneve"
description: |-
  Configure GENEVE devices.
---

## fortios_system_geneve
Configure GENEVE devices.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.

* `dstport` - GENEVE destination port (1 - 65535, default = 6081).
* `interface` - Outgoing interface for GENEVE encapsulated traffic. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_version` - IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast` `ipv6-unicast` .
* `name` - GENEVE device or interface name. Must be an unique interface name.
* `remote_ip` - IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
* `remote_ip6` - IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
* `type` - GENEVE type. Valid values: `ethernet` `ppp` .
* `vni` - GENEVE network ID.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_geneve can be imported using:
```sh
terraform import fortios_system_geneve.labelname {{mkey}}
```
