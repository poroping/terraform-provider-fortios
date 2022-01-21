---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemdhcp6_server"
description: |-
  Get information on a fortios Configure DHCPv6 servers.
---

# Data Source: fortios_systemdhcp6_server
Use this data source to get information on a fortios Configure DHCPv6 servers.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `delegated_prefix_iaid` - IAID of obtained delegated-prefix from the upstream interface.
* `dns_search_list` - DNS search list options.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `dns_service` -  Options for assigning DNS servers to DHCPv6 clients.
* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `id` - ID.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_mode` - Method used to assign client IP.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `option1` - Option 1.
* `option2` - Option 2.
* `option3` - Option 3.
* `prefix_mode` - Assigning a prefix from a DHCPv6 client or RA.
* `rapid_commit` - Enable/disable allow/disallow rapid commit.
* `status` - Enable/disable this DHCPv6 configuration.
* `subnet` - Subnet or subnet-id if the IP mode is delegated.
* `upstream_interface` - Interface name from where delegated information is provided.
* `ip_range` - DHCP IP range configuration.The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.
* `prefix_range` - DHCP prefix configuration.The structure of `prefix_range` block is documented below.

The `prefix_range` block contains:

* `end_prefix` - End of prefix range.
* `id` - ID.
* `prefix_length` - Prefix length.
* `start_prefix` - Start of prefix range.
