---
subcategory: "FortiGate Webfilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_override"
description: |-
  Configure FortiGuard Web Filter administrative overrides.
---

## fortios_webfilter_override
Configure FortiGuard Web Filter administrative overrides.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `expires` - Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
* `id` - Override rule ID.
* `initiator` - Initiating user of override (read-only setting).
* `ip` - IPv4 address which the override applies.
* `ip6` - IPv6 address which the override applies.
* `new_profile` - Name of the new web filter profile used by the override. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `old_profile` - Name of the web filter profile which the override applies. This attribute must reference one of the following datasources: `webfilter.profile.name` .
* `scope` - Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user` `user-group` `ip` `ip6` .
* `status` - Enable/disable override rule. Valid values: `enable` `disable` .
* `user` - Name of the user which the override applies.
* `user_group` - Specify the user group for which the override applies. This attribute must reference one of the following datasources: `user.group.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_webfilter_override can be imported using:
```sh
terraform import fortios_webfilter_override.labelname {{mkey}}
```
