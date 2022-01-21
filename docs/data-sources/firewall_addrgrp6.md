---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp6"
description: |-
  Get information on a fortios Configure IPv6 address groups.
---

# Data Source: fortios_firewall_addrgrp6
Use this data source to get information on a fortios Configure IPv6 address groups.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPv6 address group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting.
* `name` - IPv6 address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address group6 visibility in the GUI.
* `member` - Address objects contained within the group.The structure of `member` block is documented below.

The `member` block contains:

* `name` - Address6/addrgrp6 name.
* `tagging` - Config object tagging.The structure of `tagging` block is documented below.

The `tagging` block contains:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
