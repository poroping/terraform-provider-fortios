---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ttlpolicy"
description: |-
  Get information on a fortios Configure TTL policies.
---

# Data Source: fortios_firewall_ttlpolicy
Use this data source to get information on a fortios Configure TTL policies.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `action` - Action to be performed on traffic matching this policy (default = deny).
* `id` - ID.
* `schedule` - Schedule object from available options.
* `srcintf` - Source interface name from available interfaces.
* `status` - Enable/disable this TTL policy.
* `ttl` - Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
* `service` - Service object(s) from available options. Separate multiple names with a space.The structure of `service` block is documented below.

The `service` block contains:

* `name` - Service name.
* `srcaddr` - Source address object(s) from available options. Separate multiple names with a space.The structure of `srcaddr` block is documented below.

The `srcaddr` block contains:

* `name` - Address name.
