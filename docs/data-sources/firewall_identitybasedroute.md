---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_identitybasedroute"
description: |-
  Get information on a fortios Configure identity based routing.
---

# Data Source: fortios_firewall_identitybasedroute
Use this data source to get information on a fortios Configure identity based routing.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Comments.
* `name` - Name.
* `rule` - Rule.The structure of `rule` block is documented below.

The `rule` block contains:

* `device` - Outgoing interface for the rule.
* `gateway` - IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
* `id` - Rule ID.
* `groups` - Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space.The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name.
