---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_nat64"
description: |-
  Configure NAT64.
---

## fortios_system_nat64
Configure NAT64.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable). Valid values: `enable` `disable` .
* `generate_ipv6_fragment_header` - Enable/disable IPv6 fragment header generation. Valid values: `enable` `disable` .
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable` `disable` .
* `nat64_prefix` - NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
* `secondary_prefix_status` - Enable/disable secondary NAT64 prefix. Valid values: `enable` `disable` .
* `status` - Enable/disable NAT64 (default = disable). Valid values: `enable` `disable` .
* `secondary_prefix` - Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.

The `secondary_prefix` block contains:

* `name` - NAT64 prefix name.
* `nat64_prefix` - NAT64 prefix.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_nat64 can be imported using:
```sh
terraform import fortios_system_nat64.labelname {{mkey}}
```
