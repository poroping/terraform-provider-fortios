---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicast6"
description: |-
  Configure IPv6 multicast.
---

## fortios_router_multicast6
Configure IPv6 multicast.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `multicast_pmtu` - Enable/disable PMTU for IPv6 multicast. Valid values: `enable` `disable` .
* `multicast_routing` - Enable/disable IPv6 multicast routing. Valid values: `enable` `disable` .
* `interface` - Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.

The `interface` block contains:

* `hello_holdtime` - Time before old neighbor information expires in seconds (1 - 65535, default = 105).
* `hello_interval` - Interval between sending PIM hello messages in seconds (1 - 65535, default = 30).
* `name` - Interface name. This attribute must reference one of the following datasources: `system.interface.name` .
* `pim_sm_global` - PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.

The `pim_sm_global` block contains:

* `register_rate_limit` - Limit of packets/sec per source registered through this RP (0 means unlimited).
* `rp_address` - Statically configured RP addresses. The structure of `rp_address` block is documented below.

The `rp_address` block contains:

* `id` - ID of the entry.
* `ip6_address` - RP router IPv6 address.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_multicast6 can be imported using:
```sh
terraform import fortios_router_multicast6.labelname {{mkey}}
```
