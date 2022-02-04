---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ips"
description: |-
  Configure IPS system settings.
---

## fortios_system_ips
Configure IPS system settings.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

* `override_signature_hold_by_id` - Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold. Valid values: `enable` `disable` .
* `signature_hold_time` - Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_system_ips can be imported using:
```sh
terraform import fortios_system_ips.labelname {{mkey}}
```
