---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address"
description: |-
  Get information on a fortios Configure IPv4 addresses.
---

# Data Source: fortios_firewall_address
Use this data source to get information on a fortios Configure IPv4 addresses.


## Example Usage

```hcl
data "fortios_firewall_address" "example" {
  vdomparam = "root"

  name = "all"
}

output "example" {
  value = data.fortios_firewall_address.example
}
```

## Argument Reference

* `name` - (Required) Address name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allow_routing` - Enable/disable use of this address in the static route configuration.
* `associated_interface` - Network interface associated with address.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value.
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting.
* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `interface` - Name of interface whose IP address is to be used.
* `name` - Address name.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes.
* `obj_id` - Object ID for NSX.
* `obj_tag` - Tag of dynamic address object.
* `obj_type` - Object type.
* `organization` - Organization domain name (Syntax: organization/domain).
* `policy_group` - Policy group name.
* `sdn` - SDN.
* `sdn_addr_type` - Type of addresses to collect.
* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - First MAC address in the range.
* `sub_type` - Sub-type of address.
* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `tenant` - Tenant.
* `type` - Type of address.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI.
* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `fsso_group` - FSSO group(s).The structure of `fsso_group` block is documented below.

The `fsso_group` block contains:

* `name` - FSSO group name.
* `list` - IP address list.The structure of `list` block is documented below.

The `list` block contains:

* `ip` - IP.
* `macaddr` - Multiple MAC address ranges.The structure of `macaddr` block is documented below.

The `macaddr` block contains:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
