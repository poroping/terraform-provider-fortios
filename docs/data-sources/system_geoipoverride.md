---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geoipoverride"
description: |-
  Get information on a fortios Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
---

# Data Source: fortios_system_geoipoverride
Use this data source to get information on a fortios Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.


## Example Usage

```hcl

```

## Argument Reference

* `name` - (Required) Location name.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

## Attribute Reference

The following attributes are exported:

* `country_id` - Two character Country ID code.
* `description` - Description.
* `name` - Location name.
* `ip_range` - Table of IP ranges assigned to country.The structure of `ip_range` block is documented below.

The `ip_range` block contains:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `id` - ID of individual entry in the IP range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `ip6_range` - Table of IPv6 ranges assigned to country.The structure of `ip6_range` block is documented below.

The `ip6_range` block contains:

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `id` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
