---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ttlpolicy"
description: |-
  Configure TTL policies.
---

## fortios_firewall_ttlpolicy
Configure TTL policies.

## Example Usage

```hcl

```

## Argument Reference
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `allow_append` - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requires `id` to be defined.
* `dynamic_sort_table` - `true` or `false`, set this parameter to `true` when using dynamic for_each + toset to configure and sort sub-tables, if set to `true` static sub-tables must be ordered.

* `action` - Action to be performed on traffic matching this policy (default = deny). Valid values: `accept` `deny` .
* `id` - ID.
* `schedule` - Schedule object from available options. This attribute must reference one of the following datasources: `firewall.schedule.onetime.name` `firewall.schedule.recurring.name` `firewall.schedule.group.name` .
* `srcintf` - Source interface name from available interfaces. This attribute must reference one of the following datasources: `system.zone.name` `system.interface.name` .
* `status` - Enable/disable this TTL policy. Valid values: `enable` `disable` .
* `ttl` - Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
* `service` - Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name. This attribute must reference one of the following datasources: `firewall.service.custom.name` `firewall.service.group.name` .
* `srcaddr` - Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name. This attribute must reference one of the following datasources: `firewall.address.name` `firewall.addrgrp.name` .

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Check out `allow_append` to auto import upon resource creation.

fortios_firewall_ttlpolicy can be imported using:
```sh
terraform import fortios_firewall_ttlpolicy.labelname {{mkey}}
```
