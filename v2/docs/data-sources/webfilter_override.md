---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_override"
description: |-
  Get information on a fortios Configure FortiGuard Web Filter administrative overrides.
---

# Data Source: fortios_webfilter_override
Use this data source to get information on a fortios Configure FortiGuard Web Filter administrative overrides.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Override rule ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `expires` - Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
* `id` - Override rule ID.
* `initiator` - Initiating user of override (read-only setting).
* `ip` - IPv4 address which the override applies.
* `ip6` - IPv6 address which the override applies.
* `new_profile` - Name of the new web filter profile used by the override.
* `old_profile` - Name of the web filter profile which the override applies.
* `scope` - Override either the specific user, user group, IPv4 address, or IPv6 address.
* `status` - Enable/disable override rule.
* `user` - Name of the user which the override applies.
* `user_group` - Specify the user group for which the override applies.
