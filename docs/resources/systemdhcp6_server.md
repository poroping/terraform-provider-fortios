---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemdhcp6_server"
description: |-
  Configure DHCPv6 servers.
---

## fortios_systemdhcp6_server
Configure DHCPv6 servers.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `dns_search_list` - DNS search list options. Valid values: `delegated` `specify` .
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` - Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated` `default` `specify` .
* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `id` - ID.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_mode` - Method used to assign client IP. Valid values: `range` `delegated` .
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `option1` - Option 1.
* `option2` - Option 2.
* `option3` - Option 3.
* `prefix_mode` - Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6` `ra` .
* `rapid_commit` - Enable/disable allow/disallow rapid commit. Valid values: `disable` `enable` .
* `status` - Enable/disable this DHCPv6 configuration. Valid values: `disable` `enable` .
* `subnet` - Subnet or subnet-id if the IP mode is delegated.
* `upstream_interface` - Interface name from where delegated information is provided. This attribute must reference one of the following datasources: `system.interface.name` .
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.
* `prefix_range` - DHCP prefix configuration. The structure of `prefix_range` block is documented below.

The `prefix_range` block contains:

* `end_prefix` - End of prefix range.
* `id` - ID.
* `prefix_length` - Prefix length.
* `start_prefix` - Start of prefix range.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_systemdhcp6_server can be imported using:
```sh
terraform import fortios_systemdhcp6_server.labelname {{mkey}}
```
