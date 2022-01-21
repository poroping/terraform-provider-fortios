---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp6"
description: |-
  Configure IPv6 address groups.
---

## fortios_firewall_addrgrp6
Configure IPv6 address groups.

## Example Usage

```hcl
resource "fortios_firewall_address6" "example" {
  name    = "foobar60"
  ip6     = "3004:1234:b00b::/64"
  comment = "acc testing"
}

resource "fortios_firewall_addrgrp6" "example" {
  name = "groupygroup"
  member {
    name = fortios_firewall_address6.example.name
  }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `name` - IPv6 address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address group6 visibility in the GUI. Valid values: `enable` `disable` .
* `member` - Address objects contained within the group. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Address6/addrgrp6 name. This attribute must reference one of the following datasources: `firewall.address6.name` `firewall.addrgrp6.name` .
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

fortios_firewall_addrgrp6 can be imported using:
```sh
terraform import fortios_firewall_addrgrp6.labelname {{mkey}}
```
