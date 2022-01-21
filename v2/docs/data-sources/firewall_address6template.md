---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6template"
description: |-
  Get information on a fortios Configure IPv6 address templates.
---

# Data Source: fortios_firewall_address6template
Use this data source to get information on a fortios Configure IPv6 address templates.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) IPv6 address template name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `fabric_object` - Security Fabric global object setting.
* `ip6` - IPv6 address prefix.
* `name` - IPv6 address template name.
* `subnet_segment_count` - Number of IPv6 subnet segments.
* `subnet_segment` - IPv6 subnet segments.The structure of `subnet_segment` block is documented below.

The `subnet_segment` block contains:

* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value.
* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `values` - Subnet segment values.The structure of `values` block is documented below.

The `values` block contains:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.
