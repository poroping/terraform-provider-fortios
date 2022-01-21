---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservice"
description: |-
  Show Internet Service application.
---

## fortios_firewall_internetservice
Show Internet Service application.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.

* `database` - Database name this Internet Service belongs to. Valid values: `isdb` `irdb` .
* `direction` - How this service may be used in a firewall policy (source, destination or both). Valid values: `src` `dst` `both` .
* `extra_ip_range_number` - Extra number of IP ranges.
* `icon_id` - Icon ID of Internet Service.
* `id` - Internet Service ID.
* `ip_number` - Total number of IP addresses.
* `ip_range_number` - Number of IP ranges.
* `name` - Internet Service name.
* `obsolete` - Indicates whether the Internet Service can be used.
* `reputation` - Reputation level of the Internet Service.
* `singularity` - Singular level of the Internet Service.
* `sld_id` - Second Level Domain.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_internetservice can be imported using:
```sh
terraform import fortios_firewall_internetservice.labelname {{mkey}}
```
