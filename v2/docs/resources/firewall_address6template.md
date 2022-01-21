---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6template"
description: |-
  Configure IPv6 address templates.
---

## fortios_firewall_address6template
Configure IPv6 address templates.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `name` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `fabric_object` - Security Fabric global object setting. Valid values: `enable` `disable` .
* `ip6` - IPv6 address prefix.
* `name` - IPv6 address template name.
* `subnet_segment_count` - Number of IPv6 subnet segments.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.

The `subnet_segment` block contains:

* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value. Valid values: `enable` `disable` .
* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `values` - Subnet segment values. The structure of `values` block is documented below.

The `values` block contains:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_address6template can be imported using:
```sh
terraform import fortios_firewall_address6template.labelname {{mkey}}
```
