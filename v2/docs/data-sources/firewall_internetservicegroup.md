---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicegroup"
description: |-
  Get information on a fortios Configure group of Internet Service.
---

# Data Source: fortios_firewall_internetservicegroup
Use this data source to get information on a fortios Configure group of Internet Service.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Internet Service group name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both).
* `name` - Internet Service group name.
* `member` - Internet Service group member.The structure of `member` block is documented below.

The `member` block contains:

* `id` - Internet Service ID.
* `name` - Internet Service name.
