---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist6"
description: |-
  Get information on a fortios Configure IPv6 prefix lists.
---

# Data Source: fortios_router_prefixlist6
Use this data source to get information on a fortios Configure IPv6 prefix lists.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comments` - Comment.
* `name` - Name.
* `rule` - IPv6 prefix list rule.The structure of `rule` block is documented below.

The `rule` block contains:

* `action` - Permit or deny packets that match this rule.
* `flags` - Flags.
* `ge` - Minimum prefix length to be matched (0 - 128).
* `id` - Rule ID.
* `le` - Maximum prefix length to be matched (0 - 128).
* `prefix6` - IPv6 prefix to define regular filter criteria, such as "any" or subnets.
