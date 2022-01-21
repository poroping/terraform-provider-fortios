---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address"
description: |-
  Configure IPv4 addresses.
---

## fortios_firewall_address
Configure IPv4 addresses.

## Example Usage

```hcl
resource "fortios_firewall_address" "example" {
  name    = "foobar"
  type    = "ipmask"
  subnet  = "10.0.1.0/24"
  color   = 6
  comment = "acc testing"
}

resource "fortios_firewall_address" "example2" {
  name    = "fqdn"
  type    = "fqdn"
  fqdn    = "example.com"
  color   = 3
  comment = "acc testing"
}

resource "fortios_firewall_address" "example3" {
  name     = "iprange"
  type     = "iprange"
  start_ip = "192.168.1.4"
  end_ip   = "192.168.1.76"
  color    = 32
  comment  = "acc testing"
}

resource "fortios_firewall_address" "example4" {
  name    = "country"
  type    = "geography"
  country = "IS"
  color   = 15
  comment = "acc testing"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allow_routing` - Enable/disable use of this address in the static route configuration. Valid values: `enable` `disable` .
* `associated_interface` - Network interface associated with address. This attribute must reference one of the following datasources: `system.interface.name` `system.zone.name` .
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown` `healthy` `quarantine` `checkup` `transient` `infected` .
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `country` - IP addresses associated to a specific country.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `end_mac` - Last MAC address in the range.
* `epg_name` - Endpoint group name.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `filter` - Match criteria filter.
* `fqdn` - Fully Qualified Domain Name address.
* `interface` - Name of interface whose IP address is to be used. This attribute must reference one of the following datasources: `system.interface.name` .
* `name` - Address name.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable` `disable` .
* `obj_id` - Object ID for NSX.
* `obj_tag` - Tag of dynamic address object.
* `obj_type` - Object type. Valid values: `ip` `mac` .
* `organization` - Organization domain name (Syntax: organization/domain).
* `policy_group` - Policy group name.
* `sdn` - SDN. This attribute must reference one of the following datasources: `system.sdn-connector.name` .
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private` `public` `all` .
* `sdn_tag` - SDN Tag.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `start_mac` - First MAC address in the range.
* `sub_type` - Sub-type of address. Valid values: `sdn` `clearpass-spt` `fsso` `ems-tag` `swc-tag` .
* `subnet` - IP address and subnet mask of address.
* `subnet_name` - Subnet name.
* `tenant` - Tenant.
* `type` - Type of address. Valid values: `ipmask` `iprange` `fqdn` `geography` `wildcard` `dynamic` `interface-subnet` `mac` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `enable` `disable` .
* `wildcard` - IP address and wildcard netmask.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `fsso_group` - FSSO group(s). The structure of `fsso_group` block is documented below.

The `fsso_group` block contains:

* `name` - FSSO group name. This attribute must reference one of the following datasources: `user.adgrp.name` .
* `list` - IP address list. The structure of `list` block is documented below.

The `list` block contains:

* `ip` - IP.
* `macaddr` - Multiple MAC address ranges. The structure of `macaddr` block is documented below.

The `macaddr` block contains:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category. This attribute must reference one of the following datasources: `system.object-tagging.category` .
* `name` - Tagging entry name.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name. This attribute must reference one of the following datasources: `system.object-tagging.tags.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_address can be imported using:
```sh
terraform import fortios_firewall_address.labelname {{mkey}}
```
