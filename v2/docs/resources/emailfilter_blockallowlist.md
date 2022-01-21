---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_blockallowlist"
description: |-
  Configure anti-spam block/allow list.
---

## fortios_emailfilter_blockallowlist
Configure anti-spam block/allow list.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - Anti-spam block/allow entries. The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Reject, mark as spam or good email. Valid values: `reject` `spam` `clear` .
* `addr_type` - IP address type. Valid values: `ipv4` `ipv6` .
* `email_pattern` - Email address pattern.
* `id` - Entry ID.
* `ip4_subnet` - IPv4 network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard` `regexp` .
* `status` - Enable/disable status. Valid values: `enable` `disable` .
* `type` - Entry type. Valid values: `ip` `email` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_emailfilter_blockallowlist can be imported using:
```sh
terraform import fortios_emailfilter_blockallowlist.labelname {{mkey}}
```
