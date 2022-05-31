---
subcategory: "FortiGate Emailfilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_blockallowlist"
description: |-
  Get information on a fortios Configure anti-spam block/allow list.
---

# Data Source: fortios_emailfilter_blockallowlist
Use this data source to get information on a fortios Configure anti-spam block/allow list.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `comment` - Optional comments.
* `id` - ID.
* `name` - Name of table.
* `entries` - Anti-spam block/allow entries.The structure of `entries` block is documented below.

The `entries` block contains:

* `action` - Reject, mark as spam or good email.
* `addr_type` - IP address type.
* `email_pattern` - Email address pattern.
* `id` - Entry ID.
* `ip4_subnet` - IPv4 network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `pattern` - Pattern to match.
* `pattern_type` - Wildcard pattern or regular expression.
* `status` - Enable/disable status.
* `type` - Entry type.
