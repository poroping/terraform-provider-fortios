---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservice"
description: |-
  Get information on a fortios Show Internet Service application.
---

# Data Source: fortios_firewall_internetservice
Use this data source to get information on a fortios Show Internet Service application.


## Example Usage

```hcl

```

## Argument Reference

* `id` - (Required) Internet Service ID.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `database` - Database name this Internet Service belongs to.
* `direction` - How this service may be used in a firewall policy (source, destination or both).
* `extra_ip_range_number` - Extra number of IPv4 ranges.
* `extra_ip6_range_number` - Extra number of IPv6 ranges.
* `icon_id` - Icon ID of Internet Service.
* `id` - Internet Service ID.
* `ip_number` - Total number of IPv4 addresses.
* `ip_range_number` - Number of IPv4 ranges.
* `ip6_range_number` - Number of IPv6 ranges.
* `name` - Internet Service name.
* `obsolete` - Indicates whether the Internet Service can be used.
* `reputation` - Reputation level of the Internet Service.
* `singularity` - Singular level of the Internet Service.
* `sld_id` - Second Level Domain.
