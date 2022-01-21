---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicegroup"
description: |-
  Configure group of Internet Service.
---

## fortios_firewall_internetservicegroup
Configure group of Internet Service.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Comment.
* `direction` - How this service may be used (source, destination or both). Valid values: `source` `destination` `both` .
* `name` - Internet Service group name.
* `member` - Internet Service group member. The structure of `member` block is documented below.

The `member` block contains:

* `id` - Internet Service ID. This attribute must reference one of the following datasources: `firewall.internet-service.id` .
* `name` - Internet Service name. This attribute must reference one of the following datasources: `firewall.internet-service-name.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_internetservicegroup can be imported using:
```sh
terraform import fortios_firewall_internetservicegroup.labelname {{mkey}}
```
