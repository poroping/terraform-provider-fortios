---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp"
description: |-
  Get information on a fortios Configure IPv4 address groups.
---

# Data Source: fortios_firewall_addrgrp
Use this data source to get information on a fortios Configure IPv4 address groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Address group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `allow_routing` - Enable/disable use of this group in the static route configuration.
* `category` - Address group category.
* `color` - Color of icon on the GUI.
* `comment` - Comment.
* `exclude` - Enable/disable address exclusion.
* `fabric_object` - Security Fabric global object setting.
* `name` - Address group name.
* `type` - Address group type.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address visibility in the GUI.
* `exclude_member` - Address exclusion member.The structure of `exclude_member` block is documented below.

The `exclude_member` block contains:

* `name` - Address name.
* `member` - Address objects contained within the group.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Address name.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
