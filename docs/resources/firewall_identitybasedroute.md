---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_identitybasedroute"
description: |-
  Configure identity based routing.
---

## fortios_firewall_identitybasedroute
Configure identity based routing.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comments` - Comments.
* `name` - Name.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `device` - Outgoing interface for the rule. This attribute must reference one of the following datasources: `system.interface.name` .
* `gateway` - IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
* `id` - Rule ID.
* `groups` - Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space. The structure of `groups` block is documented below.

The `groups` block contains:

* `name` - Group name. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_identitybasedroute can be imported using:
```sh
terraform import fortios_firewall_identitybasedroute.labelname {{mkey}}
```
