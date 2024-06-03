---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipam"
description: |-
  Configure IP address management services.
---

## fortios_system_ipam
Configure IP address management services.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `pool_subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `server_type` - Configure the type of IPAM server to use. Valid values: `fabric-root` .
* `status` - Enable/disable IP address management services. Valid values: `enable` `disable` .
* `pools` - Configure IPAM pools. The structure of `pools` block is documented below.

The `pools` block contains:

* `description` - Description.
* `name` - IPAM pool name.
* `subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `rules` - Configure IPAM allocation rules. The structure of `rules` block is documented below.

The `rules` block contains:

* `description` - Description.
* `dhcp` - Enable/disable DHCP server for matching IPAM interfaces. Valid values: `enable` `disable` .
* `name` - IPAM rule name.
* `role` - Configure role of interface to match. Valid values: `any` `lan` `wan` `dmz` `undefined` .
* `device` - Configure serial number or wildcard of FortiGate to match. The structure of `device` block is documented below.

The `device` block contains:

* `name` - FortiGate serial number or wildcard.
* `interface` - Configure name or wildcard of interface to match. The structure of `interface` block is documented below.

The `interface` block contains:

* `name` - Interface name or wildcard.
* `pool` - Configure name of IPAM pool to use. The structure of `pool` block is documented below.

The `pool` block contains:

* `name` - Ipam pool name. This attribute must reference one of the following datasources: `system.ipam.pools.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ipam can be imported using:
```sh
terraform import fortios_system_ipam.labelname {{mkey}}
```
