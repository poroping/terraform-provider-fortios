---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp"
description: |-
  Configure IPv4 address groups.
---

## fortios_firewall_addrgrp
Configure IPv4 address groups.

## Example Usage

```hcl
resource "fortios_firewall_address" "example" {
  name    = "foobar"
  type    = "ipmask"
  subnet  = "10.0.1.0/24"
  color   = 6
  comment = "acc testing"
}

resource "fortios_firewall_addrgrp" "example" {
  name = "groupygroup"
  member {
    name = fortios_firewall_address.example.name
  }
}
```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `allow_routing` - Enable/disable use of this group in the static route configuration. Valid values: `enable` `disable` .
* `category` - Address group category. Valid values: `default` `ztna-ems-tag` `ztna-geo-tag` .
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `exclude` - Enable/disable address exclusion. Valid values: `enable` `disable` .
* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `name` - Address group name.
* `type` - Address group type. Valid values: `default` `folder` .
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `enable` `disable` .
* `exclude_member` - Address exclusion member. The structure of `exclude_member` block is documented below.

The `exclude_member` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
* `member` - Address objects contained within the group. The structure of `member` block is documented below.

The `member` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .
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

fortios_firewall_addrgrp can be imported using:
```sh
terraform import fortios_firewall_addrgrp.labelname {{mkey}}
```
