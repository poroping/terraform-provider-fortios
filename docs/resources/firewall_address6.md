---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6"
description: |-
  Configure IPv6 firewall addresses.
---

## fortios_firewall_address6
Configure IPv6 firewall addresses.

## Example Usage

```hcl
resource "fortios_firewall_address6" "example" {
  name    = "foobar6"
  ip6     = "3004:1234:b00b::/64"
  comment = "acc testing"
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `country` - IPv6 addresses associated to a specific country.
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_mac` - Last MAC address in the range.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `fqdn` - Fully qualified domain name.
* `host` - Host Address.
* `host_type` - Host type. Valid values: `any` `specific` .
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `name` - Address name.
* `obj_id` - Object ID for NSX.
* `sdn` - SDN. This attribute must reference one of the following datasources: `system.sdn-connector.name` .
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `start_mac` - First MAC address in the range.
* `template` - IPv6 address template. This attribute must reference one of the following datasources: `firewall.address6-template.name` .
* `type` - Type of IPv6 address object (default = ipprefix). Valid values: `ipprefix` `iprange` `fqdn` `geography` `dynamic` `template` `mac` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable the visibility of the object in the GUI. Valid values: `enable` `disable` .
* `list` - IP address list. The structure of `list` block is documented below.

The `list` block contains:

* `ip` - IP.
* `macaddr` - Multiple MAC address ranges. The structure of `macaddr` block is documented below.

The `macaddr` block contains:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.

The `subnet_segment` block contains:

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any` `specific` .
* `value` - Subnet segment value.
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

fortios_firewall_address6 can be imported using:
```sh
terraform import fortios_firewall_address6.labelname {{mkey}}
```
