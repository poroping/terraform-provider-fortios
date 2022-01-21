---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_ospf6_summaryaddress"
description: |-
  IPv6 address summary configuration.
---

## fortios_router_ospf6_summaryaddress
IPv6 address summary configuration.

~> This resource is configuring a child table of the parent resource: `fortios_router_ospf6`. If this resource is used the parent resource must NOT modify this child table or state inconsistencies will occur.


## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `advertise` - Enable/disable advertise status. Valid values: `disable` `enable` .
* `id` - Summary address entry ID.
* `prefix6` - IPv6 prefix.
* `tag` - Tag value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_router_summaryaddress can be imported using:
```sh
terraform import fortios_router_summaryaddress.labelname {{mkey}}
```
