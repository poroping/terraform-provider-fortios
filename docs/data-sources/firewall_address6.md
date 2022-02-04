---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6"
description: |-
  Get information on a fortios Configure IPv6 firewall addresses.
---

# Data Source: fortios_firewall_address6
Use this data source to get information on a fortios Configure IPv6 firewall addresses.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Address name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `country` - IPv6 addresses associated to a specific country.
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_mac` - Last MAC address in the range.
* `fabric_object` - Security Fabric global object setting.
* `fqdn` - Fully qualified domain name.
* `host` - Host Address.
* `host_type` - Host type.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `name` - Address name.
* `obj_id` - Object ID for NSX.
* `sdn` - SDN.
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `start_mac` - First MAC address in the range.
* `template` - IPv6 address template.
* `type` - Type of IPv6 address object (default = ipprefix).
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable the visibility of the object in the GUI.
* `list` - IP address list.The structure of `list` block is documented below.

The `list` block contains:

* `ip` - IP.
* `macaddr` - Multiple MAC address ranges.The structure of `macaddr` block is documented below.

The `macaddr` block contains:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.
* `subnet_segment` - IPv6 subnet segments.The structure of `subnet_segment` block is documented below.

The `subnet_segment` block contains:

* `name` - Name.
* `type` - Subnet segment type.
* `value` - Subnet segment value.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
